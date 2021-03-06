// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stn.proto

/*
Package stn is a generated protocol buffer package.

It is generated from these files:
	stn.proto

It has these top-level messages:
	StnRule
*/
package stn

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StnRule struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
	Interface string `protobuf:"bytes,2,opt,name=interface" json:"interface,omitempty"`
	RuleName  string `protobuf:"bytes,3,opt,name=rule_name,json=ruleName" json:"rule_name,omitempty"`
}

func (m *StnRule) Reset()                    { *m = StnRule{} }
func (m *StnRule) String() string            { return proto.CompactTextString(m) }
func (*StnRule) ProtoMessage()               {}
func (*StnRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StnRule) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *StnRule) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *StnRule) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

func init() {
	proto.RegisterType((*StnRule)(nil), "stn.StnRule")
}

func init() { proto.RegisterFile("stn.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0xc9, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2e, 0xc9, 0x53, 0x4a, 0xe6, 0x62, 0x0f, 0x2e,
	0xc9, 0x0b, 0x2a, 0xcd, 0x49, 0x15, 0x92, 0xe5, 0xe2, 0xca, 0x2c, 0x88, 0x4f, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0xcc, 0x2c, 0x70, 0x84, 0x08,
	0x08, 0xc9, 0x70, 0x71, 0x66, 0xe6, 0x95, 0xa4, 0x16, 0xa5, 0x25, 0x26, 0xa7, 0x4a, 0x30, 0x41,
	0x65, 0x61, 0x02, 0x42, 0xd2, 0x5c, 0x9c, 0x45, 0xa5, 0x39, 0xa9, 0xf1, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0xcc, 0x60, 0x59, 0x0e, 0x90, 0x80, 0x5f, 0x62, 0x6e, 0x6a, 0x12, 0x1b, 0xd8, 0x42, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x9b, 0xa6, 0xdc, 0x7d, 0x00, 0x00, 0x00,
}
