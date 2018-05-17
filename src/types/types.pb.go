// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/types.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types/types.proto

It has these top-level messages:
	BluePrintCreateAccountTx
	BluePrintStateTx
	BluePrintAppState
	StateQueryParams
	StateQueryResult
	MapEntry
*/
package types

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BluePrintCreateAccountTx struct {
	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BluePrintCreateAccountTx) Reset()                    { *m = BluePrintCreateAccountTx{} }
func (m *BluePrintCreateAccountTx) String() string            { return proto.CompactTextString(m) }
func (*BluePrintCreateAccountTx) ProtoMessage()               {}
func (*BluePrintCreateAccountTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *BluePrintCreateAccountTx) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BluePrintCreateAccountTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BluePrintCreateAccountTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BluePrintStateTx struct {
	Version int32  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Owner   string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BluePrintStateTx) Reset()                    { *m = BluePrintStateTx{} }
func (m *BluePrintStateTx) String() string            { return proto.CompactTextString(m) }
func (*BluePrintStateTx) ProtoMessage()               {}
func (*BluePrintStateTx) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *BluePrintStateTx) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BluePrintStateTx) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *BluePrintStateTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BluePrintAppState struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Blob    []byte `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *BluePrintAppState) Reset()                    { *m = BluePrintAppState{} }
func (m *BluePrintAppState) String() string            { return proto.CompactTextString(m) }
func (*BluePrintAppState) ProtoMessage()               {}
func (*BluePrintAppState) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *BluePrintAppState) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *BluePrintAppState) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

type StateQueryParams struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *StateQueryParams) Reset()                    { *m = StateQueryParams{} }
func (m *StateQueryParams) String() string            { return proto.CompactTextString(m) }
func (*StateQueryParams) ProtoMessage()               {}
func (*StateQueryParams) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *StateQueryParams) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type StateQueryResult struct {
	State []byte `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *StateQueryResult) Reset()                    { *m = StateQueryResult{} }
func (m *StateQueryResult) String() string            { return proto.CompactTextString(m) }
func (*StateQueryResult) ProtoMessage()               {}
func (*StateQueryResult) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

func (m *StateQueryResult) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

type MapEntry struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *MapEntry) Reset()                    { *m = MapEntry{} }
func (m *MapEntry) String() string            { return proto.CompactTextString(m) }
func (*MapEntry) ProtoMessage()               {}
func (*MapEntry) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{5} }

func (m *MapEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MapEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*BluePrintCreateAccountTx)(nil), "BluePrintCreateAccountTx")
	proto.RegisterType((*BluePrintStateTx)(nil), "BluePrintStateTx")
	proto.RegisterType((*BluePrintAppState)(nil), "BluePrintAppState")
	proto.RegisterType((*StateQueryParams)(nil), "StateQueryParams")
	proto.RegisterType((*StateQueryResult)(nil), "StateQueryResult")
	proto.RegisterType((*MapEntry)(nil), "MapEntry")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x65, 0x4a, 0xf9, 0x73, 0xca, 0x90, 0x5a, 0x0c, 0x1e, 0xa3, 0x4c, 0x99, 0x40, 0x82,
	0x27, 0x08, 0x88, 0x11, 0xa9, 0x18, 0xc4, 0x88, 0xe4, 0x34, 0x37, 0x54, 0x04, 0xdb, 0xb2, 0xcf,
	0x85, 0xbc, 0x3d, 0xf2, 0xb5, 0x09, 0xb0, 0x77, 0xb1, 0x7e, 0x9f, 0x75, 0xf7, 0xf9, 0xac, 0x83,
	0x15, 0x8d, 0x1e, 0xe3, 0x0d, 0x9f, 0xd7, 0x3e, 0x38, 0x72, 0xf5, 0x3b, 0xa8, 0xfb, 0x21, 0xe1,
	0x3a, 0x6c, 0x2d, 0x3d, 0x04, 0x34, 0x84, 0xed, 0x66, 0xe3, 0x92, 0xa5, 0xd7, 0x6f, 0xa9, 0xe0,
	0x7c, 0x87, 0x21, 0x6e, 0x9d, 0x55, 0xa2, 0x12, 0xcd, 0x52, 0x4f, 0x28, 0xaf, 0x60, 0xe9, 0xbe,
	0x2c, 0x06, 0x75, 0x52, 0x89, 0xe6, 0x52, 0xef, 0x41, 0x4a, 0x38, 0xed, 0x0d, 0x19, 0xb5, 0xa8,
	0x44, 0x53, 0x68, 0xce, 0xf5, 0x1b, 0x94, 0xb3, 0xff, 0x85, 0x0c, 0xe1, 0x91, 0xbc, 0x2d, 0xac,
	0x66, 0x6f, 0xeb, 0x3d, 0xab, 0xb3, 0xd8, 0xf4, 0x7d, 0xc0, 0x18, 0x59, 0x5c, 0xe8, 0x09, 0xb3,
	0xa2, 0x1b, 0x5c, 0xc7, 0xde, 0x42, 0x73, 0xae, 0x1b, 0x28, 0xb9, 0xed, 0x39, 0x61, 0x18, 0xd7,
	0x26, 0x98, 0xcf, 0xf8, 0x3b, 0x80, 0xf8, 0x33, 0xc0, 0xff, 0x4a, 0x8d, 0x31, 0x0d, 0x94, 0x2b,
	0x63, 0xbe, 0x3b, 0xbc, 0xb4, 0x87, 0xfa, 0x16, 0x2e, 0x9e, 0x8c, 0x7f, 0xb4, 0x14, 0x46, 0x59,
	0xc2, 0xe2, 0x03, 0xc7, 0x83, 0x29, 0xc7, 0xdc, 0xb3, 0x33, 0x43, 0xc2, 0xe9, 0x7b, 0x0c, 0xdd,
	0x19, 0x6f, 0xe2, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x90, 0xa8, 0x8b, 0xa1, 0x9e, 0x01, 0x00,
	0x00,
}