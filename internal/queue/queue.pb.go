// Code generated by protoc-gen-go.
// source: queue.proto
// DO NOT EDIT!

/*
Package queue is a generated protocol buffer package.

It is generated from these files:
	queue.proto

It has these top-level messages:
	Message
	Recipient
*/
package queue

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Recipient_Type int32

const (
	Recipient_EMAIL Recipient_Type = 0
	Recipient_PIPE  Recipient_Type = 1
)

var Recipient_Type_name = map[int32]string{
	0: "EMAIL",
	1: "PIPE",
}
var Recipient_Type_value = map[string]int32{
	"EMAIL": 0,
	"PIPE":  1,
}

func (x Recipient_Type) String() string {
	return proto.EnumName(Recipient_Type_name, int32(x))
}
func (Recipient_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Recipient_Status int32

const (
	Recipient_PENDING Recipient_Status = 0
	Recipient_SENT    Recipient_Status = 1
	Recipient_FAILED  Recipient_Status = 2
)

var Recipient_Status_name = map[int32]string{
	0: "PENDING",
	1: "SENT",
	2: "FAILED",
}
var Recipient_Status_value = map[string]int32{
	"PENDING": 0,
	"SENT":    1,
	"FAILED":  2,
}

func (x Recipient_Status) String() string {
	return proto.EnumName(Recipient_Status_name, int32(x))
}
func (Recipient_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type Message struct {
	// Message ID. Uniquely identifies this message, it is used for
	// auditing and troubleshooting.
	ID string `protobuf:"bytes,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	// The envelope for this message.
	From string       `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	To   []string     `protobuf:"bytes,3,rep,name=To,json=to" json:"To,omitempty"`
	Rcpt []*Recipient `protobuf:"bytes,4,rep,name=rcpt" json:"rcpt,omitempty"`
	Data []byte       `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// Creation timestamp.
	CreatedAtTs *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created_at_ts,json=createdAtTs" json:"created_at_ts,omitempty"`
	// Hostname of the server receiving this message.
	Hostname string `protobuf:"bytes,7,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetRcpt() []*Recipient {
	if m != nil {
		return m.Rcpt
	}
	return nil
}

func (m *Message) GetCreatedAtTs() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAtTs
	}
	return nil
}

type Recipient struct {
	Address            string           `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Type               Recipient_Type   `protobuf:"varint,2,opt,name=type,enum=queue.Recipient_Type" json:"type,omitempty"`
	Status             Recipient_Status `protobuf:"varint,3,opt,name=status,enum=queue.Recipient_Status" json:"status,omitempty"`
	LastFailureMessage string           `protobuf:"bytes,4,opt,name=last_failure_message,json=lastFailureMessage" json:"last_failure_message,omitempty"`
}

func (m *Recipient) Reset()                    { *m = Recipient{} }
func (m *Recipient) String() string            { return proto.CompactTextString(m) }
func (*Recipient) ProtoMessage()               {}
func (*Recipient) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Message)(nil), "queue.Message")
	proto.RegisterType((*Recipient)(nil), "queue.Recipient")
	proto.RegisterEnum("queue.Recipient_Type", Recipient_Type_name, Recipient_Type_value)
	proto.RegisterEnum("queue.Recipient_Status", Recipient_Status_name, Recipient_Status_value)
}

func init() { proto.RegisterFile("queue.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x50, 0x41, 0xcb, 0x9b, 0x40,
	0x10, 0xad, 0x7e, 0x46, 0x93, 0xb1, 0x0d, 0xb2, 0xb4, 0x74, 0x49, 0x2f, 0x41, 0x7a, 0x68, 0x29,
	0x68, 0x49, 0x8f, 0x85, 0x42, 0x40, 0x53, 0x84, 0x26, 0x04, 0xe3, 0x5d, 0x36, 0xba, 0x31, 0x82,
	0x66, 0xad, 0xbb, 0x1e, 0xfa, 0x3b, 0xfb, 0x7b, 0x0a, 0xdd, 0x5d, 0x4d, 0x0a, 0xfd, 0x6e, 0x33,
	0xf3, 0xde, 0xcc, 0xbc, 0xf7, 0xc0, 0xfd, 0x39, 0xd0, 0x81, 0x06, 0x5d, 0xcf, 0x04, 0x43, 0x33,
	0xdd, 0xac, 0xbe, 0x56, 0xb5, 0xb8, 0x0e, 0xe7, 0xa0, 0x60, 0x6d, 0x58, 0xb1, 0x86, 0xdc, 0xaa,
	0x50, 0xe3, 0xe7, 0xe1, 0x12, 0x76, 0xe2, 0x57, 0x47, 0x79, 0x28, 0xea, 0x96, 0x72, 0x41, 0xda,
	0xee, 0x5f, 0x35, 0xde, 0xf0, 0x7f, 0x1b, 0xe0, 0xec, 0x29, 0xe7, 0xa4, 0xa2, 0x68, 0x09, 0x66,
	0x12, 0x61, 0x63, 0x6d, 0x7c, 0x58, 0xa4, 0x66, 0x1d, 0x21, 0x04, 0xd6, 0xa5, 0x67, 0x2d, 0x36,
	0xf5, 0x44, 0xd7, 0x8a, 0x93, 0x31, 0xfc, 0xb4, 0x7e, 0x52, 0x1c, 0xa9, 0xe1, 0x3d, 0x58, 0x7d,
	0xd1, 0x09, 0x6c, 0xc9, 0x89, 0xbb, 0xf1, 0x82, 0x51, 0x5f, 0x4a, 0x8b, 0xba, 0xab, 0xe9, 0x4d,
	0xa4, 0x1a, 0x55, 0x97, 0x4a, 0x22, 0x08, 0x9e, 0xc9, 0x4b, 0x2f, 0x53, 0x5d, 0xa3, 0x6f, 0xf0,
	0xaa, 0xe8, 0x29, 0x11, 0xb4, 0xcc, 0x89, 0xc8, 0x05, 0xc7, 0xb6, 0x04, 0xdd, 0xcd, 0x2a, 0xa8,
	0x18, 0xab, 0x9a, 0xc9, 0xa3, 0xf4, 0x10, 0x64, 0x77, 0xc9, 0xa9, 0x3b, 0x2d, 0x6c, 0x45, 0xc6,
	0xd1, 0x0a, 0xe6, 0x57, 0xc6, 0xc5, 0x8d, 0xb4, 0x14, 0x3b, 0x5a, 0xe1, 0xa3, 0xf7, 0xff, 0x18,
	0xb0, 0x78, 0x68, 0x40, 0x18, 0x1c, 0x52, 0x96, 0xbd, 0x74, 0x39, 0x99, 0xbb, 0xb7, 0xe8, 0x23,
	0x58, 0x2a, 0x20, 0xed, 0x70, 0xb9, 0x79, 0xf3, 0xbf, 0xfa, 0x20, 0x93, 0x60, 0xaa, 0x29, 0x28,
	0x04, 0x5b, 0x8a, 0x10, 0x03, 0x97, 0xe6, 0x15, 0xf9, 0xed, 0x33, 0xf2, 0x49, 0xc3, 0xe9, 0x44,
	0x43, 0x9f, 0xe1, 0x75, 0x43, 0xb8, 0xc8, 0x2f, 0xa4, 0x6e, 0x86, 0x9e, 0xe6, 0xed, 0x98, 0xb2,
	0x4c, 0x4a, 0x49, 0x40, 0x0a, 0xdb, 0x8d, 0xd0, 0x94, 0xbf, 0xff, 0x0e, 0x2c, 0xf5, 0x10, 0x2d,
	0x60, 0x16, 0xef, 0xb7, 0xc9, 0x0f, 0xef, 0x05, 0x9a, 0x83, 0x75, 0x4c, 0x8e, 0xb1, 0x67, 0xf8,
	0x9f, 0xc0, 0x1e, 0x1f, 0x20, 0x17, 0x9c, 0x63, 0x7c, 0x88, 0x92, 0xc3, 0xf7, 0x91, 0x70, 0x8a,
	0x0f, 0x99, 0x67, 0x20, 0x00, 0x7b, 0x27, 0x97, 0xe2, 0xc8, 0x33, 0xcf, 0xb6, 0x0e, 0xef, 0xcb,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x32, 0x94, 0x20, 0x2f, 0x02, 0x00, 0x00,
}