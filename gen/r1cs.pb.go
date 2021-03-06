// Code generated by protoc-gen-go. DO NOT EDIT.
// source: r1cs.proto

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	r1cs.proto

It has these top-level messages:
	Constraint
	ConstraintSystem
*/
package gen

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

type Constraint struct {
	A          map[int32]int32 `protobuf:"bytes,1,rep,name=a" json:"a,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	B          map[int32]int32 `protobuf:"bytes,2,rep,name=b" json:"b,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	C          map[int32]int32 `protobuf:"bytes,3,rep,name=c" json:"c,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Annotation string          `protobuf:"bytes,4,opt,name=annotation" json:"annotation,omitempty"`
}

func (m *Constraint) Reset()                    { *m = Constraint{} }
func (m *Constraint) String() string            { return proto.CompactTextString(m) }
func (*Constraint) ProtoMessage()               {}
func (*Constraint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Constraint) GetA() map[int32]int32 {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Constraint) GetB() map[int32]int32 {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *Constraint) GetC() map[int32]int32 {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *Constraint) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

type ConstraintSystem struct {
	Variables   []string      `protobuf:"bytes,1,rep,name=variables" json:"variables,omitempty"`
	Constraints []*Constraint `protobuf:"bytes,2,rep,name=constraints" json:"constraints,omitempty"`
}

func (m *ConstraintSystem) Reset()                    { *m = ConstraintSystem{} }
func (m *ConstraintSystem) String() string            { return proto.CompactTextString(m) }
func (*ConstraintSystem) ProtoMessage()               {}
func (*ConstraintSystem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConstraintSystem) GetVariables() []string {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *ConstraintSystem) GetConstraints() []*Constraint {
	if m != nil {
		return m.Constraints
	}
	return nil
}

func init() {
	proto.RegisterType((*Constraint)(nil), "R1CS.serializer.Constraint")
	proto.RegisterType((*ConstraintSystem)(nil), "R1CS.serializer.ConstraintSystem")
}

func init() { proto.RegisterFile("r1cs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0xd2, 0xba, 0x41, 0xdf, 0x0e, 0x8e, 0xe0, 0x21, 0xa8, 0x48, 0xe9, 0xa9, 0x17,
	0x53, 0xa7, 0x1e, 0x44, 0xf0, 0x60, 0xeb, 0x3e, 0x80, 0xd9, 0xcd, 0x8b, 0xa4, 0x21, 0xcc, 0x60,
	0x97, 0x48, 0x92, 0x0d, 0xea, 0xe7, 0xf5, 0x83, 0x48, 0xba, 0x62, 0x87, 0x07, 0x87, 0xbb, 0xbd,
	0x3e, 0xfe, 0x3f, 0x68, 0xfe, 0x3c, 0x00, 0x3b, 0x13, 0x8e, 0x7e, 0x58, 0xe3, 0x0d, 0x3e, 0x66,
	0xb3, 0x6a, 0x41, 0x9d, 0xb4, 0x8a, 0x37, 0xea, 0x53, 0xda, 0xec, 0x2b, 0x02, 0xa8, 0x8c, 0x76,
	0xde, 0x72, 0xa5, 0x3d, 0xbe, 0x02, 0xc4, 0x09, 0x4a, 0xe3, 0x7c, 0x72, 0x9d, 0xd1, 0x5f, 0x59,
	0x3a, 0xe4, 0xe8, 0xe3, 0x5c, 0x7b, 0xdb, 0x32, 0xc4, 0x83, 0xa8, 0x49, 0xb4, 0x5f, 0x94, 0xbd,
	0xa8, 0x83, 0x10, 0x24, 0xde, 0x2f, 0xaa, 0x5e, 0x08, 0x7c, 0x01, 0xc0, 0xb5, 0x36, 0x9e, 0x7b,
	0x65, 0x34, 0x39, 0x4a, 0x51, 0x9e, 0xb0, 0x9d, 0xcd, 0xe9, 0x2d, 0x8c, 0xb7, 0x3f, 0x84, 0xa7,
	0x10, 0xbf, 0xcb, 0x96, 0xa0, 0x14, 0xe5, 0x23, 0x16, 0x46, 0x7c, 0x02, 0xa3, 0x0d, 0x6f, 0xd6,
	0x92, 0x44, 0xdd, 0x6e, 0xfb, 0x71, 0x1f, 0xdd, 0xa1, 0xa0, 0xca, 0x83, 0x54, 0xf5, 0x6f, 0x95,
	0x19, 0x98, 0x0e, 0x2f, 0x5b, 0xb4, 0xce, 0xcb, 0x15, 0x3e, 0x87, 0x64, 0xc3, 0xad, 0xe2, 0x75,
	0x23, 0x5d, 0xd7, 0x79, 0xc2, 0x86, 0x05, 0x7e, 0x80, 0x89, 0xf8, 0x11, 0xae, 0x6f, 0xf8, 0xec,
	0x8f, 0xbe, 0xd8, 0x6e, 0xbe, 0xcc, 0x5e, 0xd2, 0xa5, 0xf2, 0x6f, 0xeb, 0x9a, 0x0a, 0xb3, 0x2a,
	0x9e, 0xe7, 0x4f, 0x97, 0xca, 0x17, 0xe1, 0x10, 0x5e, 0xbb, 0x43, 0x28, 0x96, 0x52, 0xd7, 0xe3,
	0x6e, 0xbc, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xba, 0x7d, 0x34, 0x01, 0x21, 0x02, 0x00, 0x00,
}
