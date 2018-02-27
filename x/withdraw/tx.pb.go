// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tx.proto

/*
Package tx is a generated protocol buffer package.

It is generated from these files:
	tx.proto

It has these top-level messages:
	WithdrawTx
	LockMsg
*/
package withdraw

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

type WithdrawTx struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Sequence  uint64 `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
	// Types that are valid to be assigned to Tx:
	//	*WithdrawTx_Lock
	Tx isWithdrawTx_Tx `protobuf_oneof:"tx"`
}

func (m *WithdrawTx) Reset()                    { *m = WithdrawTx{} }
func (m *WithdrawTx) String() string            { return proto.CompactTextString(m) }
func (*WithdrawTx) ProtoMessage()               {}
func (*WithdrawTx) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isWithdrawTx_Tx interface {
	isWithdrawTx_Tx()
}

type WithdrawTx_Lock struct {
	Lock *LockMsg `protobuf:"bytes,3,opt,name=lock,oneof"`
}

func (*WithdrawTx_Lock) isWithdrawTx_Tx() {}

func (m *WithdrawTx) GetTx() isWithdrawTx_Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *WithdrawTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *WithdrawTx) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *WithdrawTx) GetLock() *LockMsg {
	if x, ok := m.GetTx().(*WithdrawTx_Lock); ok {
		return x.Lock
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WithdrawTx) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WithdrawTx_OneofMarshaler, _WithdrawTx_OneofUnmarshaler, _WithdrawTx_OneofSizer, []interface{}{
		(*WithdrawTx_Lock)(nil),
	}
}

func _WithdrawTx_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WithdrawTx)
	// tx
	switch x := m.Tx.(type) {
	case *WithdrawTx_Lock:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lock); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WithdrawTx.Tx has unexpected type %T", x)
	}
	return nil
}

func _WithdrawTx_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WithdrawTx)
	switch tag {
	case 3: // tx.lock
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LockMsg)
		err := b.DecodeMessage(msg)
		m.Tx = &WithdrawTx_Lock{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WithdrawTx_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WithdrawTx)
	// tx
	switch x := m.Tx.(type) {
	case *WithdrawTx_Lock:
		s := proto.Size(x.Lock)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LockMsg struct {
	Dest  []byte `protobuf:"bytes,1,opt,name=dest,proto3" json:"dest,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	Token []byte `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *LockMsg) Reset()                    { *m = LockMsg{} }
func (m *LockMsg) String() string            { return proto.CompactTextString(m) }
func (*LockMsg) ProtoMessage()               {}
func (*LockMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LockMsg) GetDest() []byte {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *LockMsg) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *LockMsg) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *LockMsg) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*WithdrawTx)(nil), "WithdrawTx")
	proto.RegisterType((*LockMsg)(nil), "LockMsg")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0xa9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0xca, 0xe6, 0xe2, 0x0c, 0xcf, 0x2c, 0xc9, 0x4b, 0x2d, 0x2e, 0x0e,
	0xa9, 0x10, 0x92, 0xe1, 0xe2, 0x2c, 0xce, 0x4c, 0xcf, 0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x09, 0x42, 0x08, 0x08, 0x49, 0x71, 0x71, 0x14, 0xa7, 0x16, 0x96, 0xa6,
	0xe6, 0x25, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0xc1, 0xf9, 0x42, 0x72, 0x5c, 0x2c,
	0x39, 0xf9, 0xc9, 0xd9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x1c, 0x7a, 0x3e, 0xf9, 0xc9,
	0xd9, 0xbe, 0xc5, 0xe9, 0x1e, 0x0c, 0x41, 0x60, 0x71, 0x27, 0x16, 0x2e, 0xa6, 0x92, 0x0a, 0xa5,
	0x78, 0x2e, 0x76, 0xa8, 0x84, 0x90, 0x10, 0x17, 0x4b, 0x4a, 0x6a, 0x71, 0x09, 0xd4, 0x16, 0x30,
	0x5b, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x14, 0x66, 0x3a, 0x84, 0x03, 0x12, 0x2d, 0xc9,
	0xcf, 0x4e, 0xcd, 0x03, 0x9b, 0xcd, 0x13, 0x04, 0xe1, 0x80, 0x44, 0xf3, 0xf2, 0x41, 0x2e, 0x61,
	0x81, 0xa8, 0x05, 0x73, 0x92, 0xd8, 0xc0, 0x9e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x00,
	0x68, 0x3e, 0x5f, 0xe0, 0x00, 0x00, 0x00,
}
