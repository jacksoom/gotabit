// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gotabit/inbox/v1beta1/inbox.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Msg defines the inbox item - msg
type Msg struct {
	// msg id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// msg sender address
	Sender string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	// msg recipient address
	To string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty" yaml:"to_address"`
	// msg topics
	Topics string `protobuf:"bytes,4,opt,name=topics,proto3" json:"topics,omitempty" yaml:"the_message_topics"`
	// msg message
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty" yaml:"message"`
}

func (m *Msg) Reset()      { *m = Msg{} }
func (*Msg) ProtoMessage() {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b4358adc09e2fc, []int{0}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Msg)(nil), "gotabit.inbox.v1beta1.Msg")
}

func init() { proto.RegisterFile("gotabit/inbox/v1beta1/inbox.proto", fileDescriptor_08b4358adc09e2fc) }

var fileDescriptor_08b4358adc09e2fc = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xd0, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0x07, 0xf0, 0xa4, 0x9b, 0x53, 0x03, 0x0e, 0x16, 0x1c, 0x54, 0xc1, 0x64, 0x06, 0x84, 0x09,
	0xb2, 0x32, 0xc4, 0xcb, 0x8e, 0x3b, 0x78, 0xf3, 0xd2, 0xa3, 0x97, 0xd1, 0xae, 0x21, 0x0b, 0xd8,
	0xa5, 0x34, 0x51, 0xb6, 0x37, 0xf0, 0xe8, 0xd1, 0xe3, 0x1e, 0x67, 0xc7, 0x9d, 0xc4, 0x53, 0xd1,
	0xf6, 0x0d, 0xfa, 0x04, 0x62, 0x93, 0x9e, 0x92, 0xef, 0xfb, 0x7e, 0xff, 0xcb, 0x1f, 0x5d, 0x0b,
	0x65, 0xa2, 0x58, 0x9a, 0x40, 0xae, 0x63, 0xb5, 0x09, 0xde, 0xa6, 0x31, 0x37, 0xd1, 0xd4, 0x4e,
	0x93, 0x2c, 0x57, 0x46, 0xe1, 0xa1, 0x23, 0x13, 0xbb, 0x74, 0xe4, 0xf2, 0x5c, 0x28, 0xa1, 0x1a,
	0x11, 0xfc, 0xff, 0x2c, 0x66, 0x5f, 0x10, 0x75, 0x9e, 0xb4, 0xc0, 0x7d, 0xe4, 0xc9, 0xc4, 0x87,
	0x23, 0x38, 0xee, 0x86, 0x9e, 0x4c, 0xf0, 0x2d, 0xea, 0x69, 0xbe, 0x4e, 0x78, 0xee, 0x7b, 0x23,
	0x38, 0x3e, 0x9d, 0x0f, 0xea, 0x82, 0x9e, 0x6d, 0xa3, 0xf4, 0x65, 0xc6, 0xec, 0x9e, 0x85, 0x0e,
	0xe0, 0x1b, 0xe4, 0x19, 0xe5, 0x77, 0x1a, 0x36, 0xac, 0x0b, 0x3a, 0xb0, 0xcc, 0xa8, 0x45, 0x94,
	0x24, 0x39, 0xd7, 0x9a, 0x85, 0x9e, 0x51, 0xf8, 0x01, 0xf5, 0x8c, 0xca, 0xe4, 0x52, 0xfb, 0xdd,
	0x86, 0x5e, 0xd5, 0x05, 0xbd, 0x70, 0x74, 0xc5, 0x17, 0x29, 0xd7, 0x3a, 0x12, 0x7c, 0x61, 0x0d,
	0x0b, 0x1d, 0xc6, 0x77, 0xe8, 0xd8, 0x9d, 0xfc, 0xa3, 0x26, 0x87, 0xeb, 0x82, 0xf6, 0x6d, 0xce,
	0x1d, 0x58, 0xd8, 0x92, 0xd9, 0xc9, 0xfb, 0x8e, 0x82, 0xcf, 0x1d, 0x05, 0xf3, 0xc7, 0xfd, 0x2f,
	0x01, 0xfb, 0x92, 0xc0, 0x43, 0x49, 0xe0, 0x4f, 0x49, 0xe0, 0x47, 0x45, 0xc0, 0xa1, 0x22, 0xe0,
	0xbb, 0x22, 0xe0, 0x79, 0x2c, 0xa4, 0x59, 0xbd, 0xc6, 0x93, 0xa5, 0x4a, 0x83, 0xb6, 0xd1, 0xf6,
	0xdd, 0xb8, 0x6e, 0xcd, 0x36, 0xe3, 0x3a, 0xee, 0x35, 0x3d, 0xdd, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xb5, 0xe4, 0x82, 0x30, 0x79, 0x01, 0x00, 0x00,
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Msg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintInbox(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Topics) > 0 {
		i -= len(m.Topics)
		copy(dAtA[i:], m.Topics)
		i = encodeVarintInbox(dAtA, i, uint64(len(m.Topics)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintInbox(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintInbox(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintInbox(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintInbox(dAtA []byte, offset int, v uint64) int {
	offset -= sovInbox(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovInbox(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovInbox(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovInbox(uint64(l))
	}
	l = len(m.Topics)
	if l > 0 {
		n += 1 + l + sovInbox(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovInbox(uint64(l))
	}
	return n
}

func sovInbox(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInbox(x uint64) (n int) {
	return sovInbox(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInbox
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInbox
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInbox
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInbox
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInbox
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInbox
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInbox
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInbox
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInbox
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInbox(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInbox
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInbox(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInbox
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInbox
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthInbox
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInbox
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInbox
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInbox        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInbox          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInbox = fmt.Errorf("proto: unexpected end of group")
)
