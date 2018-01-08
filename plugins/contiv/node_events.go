// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contiv

import (
	"context"
	"fmt"
	"strings"

	"github.com/contiv/vpp/plugins/contiv/model/uid"
	"github.com/ligato/cn-infra/datasync"
)

// handleNodeEvents handles changes in nodes within the k8s cluster (node add / delete) and
// adjusts the vswitch config (routes to the other nodes) accordingly.
func (s *remoteCNIserver) handleNodeEvents(ctx context.Context, resyncChan chan datasync.ResyncEvent, changeChan chan datasync.ChangeEvent) {
	for {
		select {

		case resyncEv := <-resyncChan:
			err := s.nodeResync(resyncEv)
			resyncEv.Done(err)

		case changeEv := <-changeChan:
			err := s.nodeChangePropageteEvent(changeEv)
			changeEv.Done(err)

		case <-ctx.Done():
			return
		}
	}
}

// nodeResync processes all nodes data and configures vswitch (routes to the other nodes) accordingly.
func (s *remoteCNIserver) nodeResync(dataResyncEv datasync.ResyncEvent) error {

	// TODO: implement proper resync (handle deleted routes as well)

	var err error
	txn := s.vppTxnFactory().Put()
	data := dataResyncEv.GetValues()

	for prefix, it := range data {
		if prefix == allocatedIDsKeyPrefix {
			for {
				kv, stop := it.GetNext()
				if stop {
					break
				}
				nodeInfo := &uid.Identifier{}
				err = kv.GetValue(nodeInfo)
				if err != nil {
					return err
				}
				nodeID := uint8(nodeInfo.Id)

				if nodeID != s.ipam.NodeID() {
					s.Logger.Info("Other node discovered: ", nodeID)

					// add routes to the node
					err = s.addRoutesToNode(nodeID)
				}
			}
		}
	}

	return txn.Send().ReceiveReply()
}

// nodeChangePropageteEvent handles change in nodes within the k8s cluster (node add / delete)
// and configures vswitch (routes to the other nodes) accordingly.
func (s *remoteCNIserver) nodeChangePropageteEvent(dataChngEv datasync.ChangeEvent) error {
	var err error
	key := dataChngEv.GetKey()

	if strings.HasPrefix(key, allocatedIDsKeyPrefix) {
		nodeInfo := &uid.Identifier{}
		err = dataChngEv.GetValue(nodeInfo)
		if err != nil {
			return err
		}
		nodeID := uint8(nodeInfo.Id)

		// route := s.getRouteToNode(conf, nodeInfo.Id)
		if dataChngEv.GetChangeType() == datasync.Put {
			s.Logger.Info("New node discovered: ", nodeID)

			// add routes to the node
			err = s.addRoutesToNode(nodeID)
		} else {
			s.Logger.Info("Node removed: ", nodeID)

			// delete routes to the node
			err = s.deleteRoutesToNode(nodeID)
		}
	} else {
		return fmt.Errorf("Unknown key %v", key)
	}

	return err
}

// addRoutesToNode add routes to the node specified by nodeID.
func (s *remoteCNIserver) addRoutesToNode(nodeID uint8) error {
	podsRoute, hostRoute, err := s.computeRoutesForHost(nodeID)
	if err != nil {
		return err
	}
	s.Logger.Info("Adding PODs route: ", podsRoute)
	s.Logger.Info("Adding host route: ", hostRoute)

	err = s.vppTxnFactory().Put().
		StaticRoute(podsRoute).
		StaticRoute(hostRoute).
		Send().ReceiveReply()

	if err != nil {
		return fmt.Errorf("Can't configure vpp to add route to host %v (and its pods): %v ", nodeID, err)
	}
	return nil
}

// deleteRoutesToNode delete routes to the node specified by nodeID.
func (s *remoteCNIserver) deleteRoutesToNode(nodeID uint8) error {
	podsRoute, hostRoute, err := s.computeRoutesForHost(nodeID)
	if err != nil {
		return err
	}
	s.Logger.Info("Deleting PODs route: ", podsRoute)
	s.Logger.Info("Deleting host route: ", hostRoute)

	err = s.vppTxnFactory().Delete().
		StaticRoute(podsRoute.VrfId, podsRoute.DstIpAddr, podsRoute.NextHopAddr).
		StaticRoute(hostRoute.VrfId, hostRoute.DstIpAddr, hostRoute.NextHopAddr).
		Send().ReceiveReply()

	if err != nil {
		return fmt.Errorf("Can't configure vpp to remove route to host %v (and its pods): %v ", nodeID, err)
	}
	return nil
}