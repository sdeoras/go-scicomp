// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comp.proto

package proto

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

// Checkpoint is a struct storing weights and biases
type Checkpoint struct {
	// weights stores different weights against a string label
	Weights map[string]*Variable `protobuf:"bytes,1,rep,name=weights,proto3" json:"weights,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// biases stores different biases against a string label
	Biases               map[string]*Variable `protobuf:"bytes,2,rep,name=biases,proto3" json:"biases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_comp_bbb1667608aa7b47, []int{0}
}
func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (dst *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(dst, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetWeights() map[string]*Variable {
	if m != nil {
		return m.Weights
	}
	return nil
}

func (m *Checkpoint) GetBiases() map[string]*Variable {
	if m != nil {
		return m.Biases
	}
	return nil
}

// Variable is basically a buffer and its size
type Variable struct {
	Data                 []float32 `protobuf:"fixed32,1,rep,packed,name=data,proto3" json:"data,omitempty"`
	Size                 []int64   `protobuf:"varint,2,rep,packed,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Variable) Reset()         { *m = Variable{} }
func (m *Variable) String() string { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()    {}
func (*Variable) Descriptor() ([]byte, []int) {
	return fileDescriptor_comp_bbb1667608aa7b47, []int{1}
}
func (m *Variable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variable.Unmarshal(m, b)
}
func (m *Variable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variable.Marshal(b, m, deterministic)
}
func (dst *Variable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variable.Merge(dst, src)
}
func (m *Variable) XXX_Size() int {
	return xxx_messageInfo_Variable.Size(m)
}
func (m *Variable) XXX_DiscardUnknown() {
	xxx_messageInfo_Variable.DiscardUnknown(m)
}

var xxx_messageInfo_Variable proto.InternalMessageInfo

func (m *Variable) GetData() []float32 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Variable) GetSize() []int64 {
	if m != nil {
		return m.Size
	}
	return nil
}

func init() {
	proto.RegisterType((*Checkpoint)(nil), "proto.Checkpoint")
	proto.RegisterMapType((map[string]*Variable)(nil), "proto.Checkpoint.BiasesEntry")
	proto.RegisterMapType((map[string]*Variable)(nil), "proto.Checkpoint.WeightsEntry")
	proto.RegisterType((*Variable)(nil), "proto.Variable")
}

func init() { proto.RegisterFile("comp.proto", fileDescriptor_comp_bbb1667608aa7b47) }

var fileDescriptor_comp_bbb1667608aa7b47 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0xcf, 0x2d,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xd3, 0x98, 0xb8, 0xb8, 0x9c,
	0x33, 0x52, 0x93, 0xb3, 0x0b, 0xf2, 0x33, 0xf3, 0x4a, 0x84, 0x2c, 0xb8, 0xd8, 0xcb, 0x53, 0x33,
	0xd3, 0x33, 0x4a, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xe4, 0x20, 0xca, 0xf5, 0x10,
	0x6a, 0xf4, 0xc2, 0x21, 0x0a, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0x60, 0xca, 0x85, 0x4c, 0xb9,
	0xd8, 0x92, 0x32, 0x13, 0x8b, 0x53, 0x8b, 0x25, 0x98, 0xc0, 0x1a, 0x65, 0x31, 0x35, 0x3a, 0x81,
	0xe5, 0x21, 0xfa, 0xa0, 0x8a, 0xa5, 0xbc, 0xb9, 0x78, 0x90, 0xcd, 0x13, 0x12, 0xe0, 0x62, 0xce,
	0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x54, 0xb9, 0x58, 0xcb,
	0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xa1, 0xe6, 0x86, 0x25,
	0x16, 0x65, 0x26, 0x26, 0xe5, 0xa4, 0x06, 0x41, 0x64, 0xad, 0x98, 0x2c, 0x18, 0xa5, 0xbc, 0xb8,
	0xb8, 0x91, 0xec, 0xa0, 0xc8, 0x2c, 0x25, 0x23, 0x2e, 0x0e, 0x98, 0xb0, 0x90, 0x10, 0x17, 0x4b,
	0x4a, 0x62, 0x49, 0x22, 0x38, 0x48, 0x98, 0x82, 0xc0, 0x6c, 0x90, 0x58, 0x71, 0x66, 0x55, 0x2a,
	0xd8, 0xb7, 0xcc, 0x41, 0x60, 0x76, 0x12, 0x1b, 0xd8, 0x38, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x20, 0xc3, 0xaa, 0x68, 0x01, 0x00, 0x00,
}
