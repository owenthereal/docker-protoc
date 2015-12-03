// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestMessage
*/
package test

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TestMessage struct {
	Hello string `protobuf:"bytes,1,opt,name=hello" json:"hello,omitempty"`
	World string `protobuf:"bytes,2,opt,name=world" json:"world,omitempty"`
}

func (m *TestMessage) Reset()                    { *m = TestMessage{} }
func (m *TestMessage) String() string            { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()               {}
func (*TestMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*TestMessage)(nil), "TestMessage")
}

var fileDescriptor0 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe6, 0xe2, 0x0e, 0x01, 0xf2, 0x7c, 0x53, 0x8b,
	0x8b, 0x13, 0xd3, 0x53, 0x85, 0x78, 0xb9, 0x58, 0x33, 0x52, 0x73, 0x72, 0xf2, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x41, 0xdc, 0xf2, 0xfc, 0xa2, 0x9c, 0x14, 0x09, 0x26, 0x10, 0x37, 0x89, 0x0d,
	0xac, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xbb, 0x5c, 0x75, 0x41, 0x00, 0x00, 0x00,
}
