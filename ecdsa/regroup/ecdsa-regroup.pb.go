// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob/ecdsa-regroup.proto

package regroup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//
// The Round 1 data is broadcast to peers of the New Committee in this message.
type DGRound1Message struct {
	EcdsaPubX            []byte   `protobuf:"bytes,1,opt,name=ecdsa_pub_x,json=ecdsaPubX,proto3" json:"ecdsa_pub_x,omitempty"`
	EcdsaPubY            []byte   `protobuf:"bytes,2,opt,name=ecdsa_pub_y,json=ecdsaPubY,proto3" json:"ecdsa_pub_y,omitempty"`
	VCommitment          []byte   `protobuf:"bytes,3,opt,name=v_commitment,json=vCommitment,proto3" json:"v_commitment,omitempty"`
	XAndKCommitment      []byte   `protobuf:"bytes,4,opt,name=x_and_k_commitment,json=xAndKCommitment,proto3" json:"x_and_k_commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound1Message) Reset()         { *m = DGRound1Message{} }
func (m *DGRound1Message) String() string { return proto.CompactTextString(m) }
func (*DGRound1Message) ProtoMessage()    {}
func (*DGRound1Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb988ad5f06b19d, []int{0}
}

func (m *DGRound1Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound1Message.Unmarshal(m, b)
}
func (m *DGRound1Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound1Message.Marshal(b, m, deterministic)
}
func (m *DGRound1Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound1Message.Merge(m, src)
}
func (m *DGRound1Message) XXX_Size() int {
	return xxx_messageInfo_DGRound1Message.Size(m)
}
func (m *DGRound1Message) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound1Message.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound1Message proto.InternalMessageInfo

func (m *DGRound1Message) GetEcdsaPubX() []byte {
	if m != nil {
		return m.EcdsaPubX
	}
	return nil
}

func (m *DGRound1Message) GetEcdsaPubY() []byte {
	if m != nil {
		return m.EcdsaPubY
	}
	return nil
}

func (m *DGRound1Message) GetVCommitment() []byte {
	if m != nil {
		return m.VCommitment
	}
	return nil
}

func (m *DGRound1Message) GetXAndKCommitment() []byte {
	if m != nil {
		return m.XAndKCommitment
	}
	return nil
}

//
// The Round 2 data is broadcast to other peers of the New Committee in this message.
type DGRound2Message1 struct {
	PaillierN            []byte   `protobuf:"bytes,1,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	PaillierProof        [][]byte `protobuf:"bytes,2,rep,name=paillier_proof,json=paillierProof,proto3" json:"paillier_proof,omitempty"`
	NTilde               []byte   `protobuf:"bytes,3,opt,name=n_tilde,json=nTilde,proto3" json:"n_tilde,omitempty"`
	H1                   []byte   `protobuf:"bytes,4,opt,name=h1,proto3" json:"h1,omitempty"`
	H2                   []byte   `protobuf:"bytes,5,opt,name=h2,proto3" json:"h2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound2Message1) Reset()         { *m = DGRound2Message1{} }
func (m *DGRound2Message1) String() string { return proto.CompactTextString(m) }
func (*DGRound2Message1) ProtoMessage()    {}
func (*DGRound2Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb988ad5f06b19d, []int{1}
}

func (m *DGRound2Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound2Message1.Unmarshal(m, b)
}
func (m *DGRound2Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound2Message1.Marshal(b, m, deterministic)
}
func (m *DGRound2Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound2Message1.Merge(m, src)
}
func (m *DGRound2Message1) XXX_Size() int {
	return xxx_messageInfo_DGRound2Message1.Size(m)
}
func (m *DGRound2Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound2Message1.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound2Message1 proto.InternalMessageInfo

func (m *DGRound2Message1) GetPaillierN() []byte {
	if m != nil {
		return m.PaillierN
	}
	return nil
}

func (m *DGRound2Message1) GetPaillierProof() [][]byte {
	if m != nil {
		return m.PaillierProof
	}
	return nil
}

func (m *DGRound2Message1) GetNTilde() []byte {
	if m != nil {
		return m.NTilde
	}
	return nil
}

func (m *DGRound2Message1) GetH1() []byte {
	if m != nil {
		return m.H1
	}
	return nil
}

func (m *DGRound2Message1) GetH2() []byte {
	if m != nil {
		return m.H2
	}
	return nil
}

//
// The Round 2 "ACK" is broadcast to peers of the Old Committee in this message.
type DGRound2Message2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound2Message2) Reset()         { *m = DGRound2Message2{} }
func (m *DGRound2Message2) String() string { return proto.CompactTextString(m) }
func (*DGRound2Message2) ProtoMessage()    {}
func (*DGRound2Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb988ad5f06b19d, []int{2}
}

func (m *DGRound2Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound2Message2.Unmarshal(m, b)
}
func (m *DGRound2Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound2Message2.Marshal(b, m, deterministic)
}
func (m *DGRound2Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound2Message2.Merge(m, src)
}
func (m *DGRound2Message2) XXX_Size() int {
	return xxx_messageInfo_DGRound2Message2.Size(m)
}
func (m *DGRound2Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound2Message2.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound2Message2 proto.InternalMessageInfo

//
// The Round 3 data is sent to peers of the New Committee in this message.
type DGRound3Message1 struct {
	Share                []byte   `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound3Message1) Reset()         { *m = DGRound3Message1{} }
func (m *DGRound3Message1) String() string { return proto.CompactTextString(m) }
func (*DGRound3Message1) ProtoMessage()    {}
func (*DGRound3Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb988ad5f06b19d, []int{3}
}

func (m *DGRound3Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound3Message1.Unmarshal(m, b)
}
func (m *DGRound3Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound3Message1.Marshal(b, m, deterministic)
}
func (m *DGRound3Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound3Message1.Merge(m, src)
}
func (m *DGRound3Message1) XXX_Size() int {
	return xxx_messageInfo_DGRound3Message1.Size(m)
}
func (m *DGRound3Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound3Message1.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound3Message1 proto.InternalMessageInfo

func (m *DGRound3Message1) GetShare() []byte {
	if m != nil {
		return m.Share
	}
	return nil
}

//
// The Round 3 data is broadcast to peers of the New Committee in this message.
type DGRound3Message2 struct {
	VDecommitment        [][]byte `protobuf:"bytes,1,rep,name=v_decommitment,json=vDecommitment,proto3" json:"v_decommitment,omitempty"`
	XAndKDecommitment    [][]byte `protobuf:"bytes,2,rep,name=x_and_k_decommitment,json=xAndKDecommitment,proto3" json:"x_and_k_decommitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DGRound3Message2) Reset()         { *m = DGRound3Message2{} }
func (m *DGRound3Message2) String() string { return proto.CompactTextString(m) }
func (*DGRound3Message2) ProtoMessage()    {}
func (*DGRound3Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb988ad5f06b19d, []int{4}
}

func (m *DGRound3Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DGRound3Message2.Unmarshal(m, b)
}
func (m *DGRound3Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DGRound3Message2.Marshal(b, m, deterministic)
}
func (m *DGRound3Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DGRound3Message2.Merge(m, src)
}
func (m *DGRound3Message2) XXX_Size() int {
	return xxx_messageInfo_DGRound3Message2.Size(m)
}
func (m *DGRound3Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_DGRound3Message2.DiscardUnknown(m)
}

var xxx_messageInfo_DGRound3Message2 proto.InternalMessageInfo

func (m *DGRound3Message2) GetVDecommitment() [][]byte {
	if m != nil {
		return m.VDecommitment
	}
	return nil
}

func (m *DGRound3Message2) GetXAndKDecommitment() [][]byte {
	if m != nil {
		return m.XAndKDecommitment
	}
	return nil
}

func init() {
	proto.RegisterType((*DGRound1Message)(nil), "DGRound1Message")
	proto.RegisterType((*DGRound2Message1)(nil), "DGRound2Message1")
	proto.RegisterType((*DGRound2Message2)(nil), "DGRound2Message2")
	proto.RegisterType((*DGRound3Message1)(nil), "DGRound3Message1")
	proto.RegisterType((*DGRound3Message2)(nil), "DGRound3Message2")
}

func init() { proto.RegisterFile("protob/ecdsa-regroup.proto", fileDescriptor_2fb988ad5f06b19d) }

var fileDescriptor_2fb988ad5f06b19d = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x5d, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x49, 0xfa, 0xb5, 0x1f, 0x9e, 0xfe, 0xe9, 0x50, 0x30, 0x08, 0x4a, 0x0d, 0x08, 0x05,
	0xd1, 0x92, 0x74, 0x05, 0x6a, 0xc1, 0x0b, 0x51, 0x4a, 0xf0, 0x42, 0xbd, 0x19, 0x92, 0xce, 0xd8,
	0x44, 0x93, 0x99, 0x30, 0xf9, 0x21, 0x2e, 0xc3, 0x25, 0xb8, 0x53, 0xc9, 0x64, 0x12, 0x13, 0x7a,
	0xf9, 0x3e, 0xe7, 0x1d, 0x78, 0xe6, 0x1c, 0x38, 0x89, 0x05, 0x4f, 0xb9, 0xb7, 0xa4, 0x5b, 0x92,
	0xb8, 0x57, 0x82, 0xee, 0x04, 0xcf, 0xe2, 0x6b, 0x09, 0xcd, 0x1f, 0x0d, 0xa6, 0xeb, 0x7b, 0x87,
	0x67, 0x8c, 0x58, 0x8f, 0x34, 0x49, 0xdc, 0x1d, 0x45, 0x67, 0x30, 0x94, 0x55, 0x1c, 0x67, 0x1e,
	0x2e, 0x0c, 0x6d, 0xae, 0x2d, 0x46, 0xce, 0x81, 0x44, 0x9b, 0xcc, 0x7b, 0xe9, 0xce, 0xbf, 0x0c,
	0xbd, 0x3b, 0x7f, 0x45, 0xe7, 0x30, 0xca, 0xf1, 0x96, 0x47, 0x51, 0x90, 0x46, 0x94, 0xa5, 0x46,
	0x4f, 0x16, 0x86, 0xf9, 0x5d, 0x83, 0xd0, 0x25, 0xa0, 0x02, 0xbb, 0x8c, 0xe0, 0xcf, 0x76, 0xf1,
	0x9f, 0x2c, 0x4e, 0x8b, 0x1b, 0x46, 0x1e, 0xfe, 0xca, 0xe6, 0xb7, 0x06, 0x87, 0xca, 0xd1, 0x56,
	0x8e, 0x16, 0x3a, 0x05, 0x88, 0xdd, 0x20, 0x0c, 0x03, 0x2a, 0x30, 0xab, 0x1d, 0x6b, 0xf2, 0x84,
	0x2e, 0x60, 0xd2, 0x8c, 0x63, 0xc1, 0xf9, 0xbb, 0xa1, 0xcf, 0x7b, 0x8b, 0x91, 0x33, 0xae, 0xe9,
	0xa6, 0x84, 0xe8, 0x18, 0xfe, 0x33, 0x9c, 0x06, 0x21, 0xa1, 0xca, 0x72, 0xc0, 0x9e, 0xcb, 0x84,
	0x26, 0xa0, 0xfb, 0x96, 0x12, 0xd2, 0x7d, 0x4b, 0x66, 0xdb, 0xe8, 0xab, 0x6c, 0x9b, 0x68, 0x4f,
	0xc9, 0x36, 0x17, 0x0d, 0x5b, 0x35, 0x9a, 0x33, 0xe8, 0x27, 0xbe, 0x2b, 0xa8, 0x32, 0xac, 0x82,
	0xf9, 0xb1, 0xd7, 0xb4, 0x4b, 0xe3, 0x1c, 0x13, 0xda, 0x5a, 0x87, 0x56, 0x19, 0xe7, 0xeb, 0x16,
	0x44, 0x4b, 0x98, 0xd5, 0x9b, 0xeb, 0x94, 0xab, 0xef, 0x1d, 0xc9, 0xdd, 0xb5, 0x1f, 0xdc, 0x4e,
	0xdf, 0xc6, 0xf2, 0x34, 0x4b, 0x75, 0x78, 0x6f, 0x20, 0x2f, 0xbf, 0xfa, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x92, 0x63, 0x65, 0x4f, 0x17, 0x02, 0x00, 0x00,
}