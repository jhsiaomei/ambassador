// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/admin/v2alpha/listeners.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/go/apis/envoy/api/v2/core"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Admin endpoint uses this wrapper for `/listeners` to display listener status information.
// See :ref:`/listeners <operations_admin_interface_listeners>` for more information.
type Listeners struct {
	// List of listener statuses.
	ListenerStatuses     []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Listeners) Reset()         { *m = Listeners{} }
func (m *Listeners) String() string { return proto.CompactTextString(m) }
func (*Listeners) ProtoMessage()    {}
func (*Listeners) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3cc5a2daffd3f5, []int{0}
}
func (m *Listeners) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Listeners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Listeners.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Listeners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listeners.Merge(m, src)
}
func (m *Listeners) XXX_Size() int {
	return m.Size()
}
func (m *Listeners) XXX_DiscardUnknown() {
	xxx_messageInfo_Listeners.DiscardUnknown(m)
}

var xxx_messageInfo_Listeners proto.InternalMessageInfo

func (m *Listeners) GetListenerStatuses() []*ListenerStatus {
	if m != nil {
		return m.ListenerStatuses
	}
	return nil
}

// Details an individual listener's current status.
type ListenerStatus struct {
	// Name of the listener
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The actual local address that the listener is listening on. If a listener was configured
	// to listen on port 0, then this address has the port that was allocated by the OS.
	LocalAddress         *core.Address `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListenerStatus) Reset()         { *m = ListenerStatus{} }
func (m *ListenerStatus) String() string { return proto.CompactTextString(m) }
func (*ListenerStatus) ProtoMessage()    {}
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd3cc5a2daffd3f5, []int{1}
}
func (m *ListenerStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListenerStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListenerStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListenerStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerStatus.Merge(m, src)
}
func (m *ListenerStatus) XXX_Size() int {
	return m.Size()
}
func (m *ListenerStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerStatus proto.InternalMessageInfo

func (m *ListenerStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListenerStatus) GetLocalAddress() *core.Address {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Listeners)(nil), "envoy.admin.v2alpha.Listeners")
	proto.RegisterType((*ListenerStatus)(nil), "envoy.admin.v2alpha.ListenerStatus")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/listeners.proto", fileDescriptor_cd3cc5a2daffd3f5)
}

var fileDescriptor_cd3cc5a2daffd3f5 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0x89, 0x8a, 0xd0, 0x54, 0x45, 0xe3, 0x66, 0xe8, 0x62, 0x1c, 0xdb, 0xcd, 0xac, 0x5e,
	0x20, 0x7e, 0x80, 0xb4, 0x6b, 0x17, 0xc3, 0xb8, 0x96, 0x12, 0x3b, 0x0f, 0x0c, 0xa4, 0x93, 0x90,
	0x37, 0x0e, 0xf6, 0x0f, 0x5d, 0xfa, 0x09, 0x32, 0x5f, 0x22, 0x4d, 0x33, 0x42, 0xa1, 0xbb, 0x3c,
	0x38, 0xe7, 0xe6, 0x72, 0xf9, 0x02, 0xdb, 0xde, 0xed, 0xa4, 0x6e, 0xb6, 0xa6, 0x95, 0xbd, 0xd2,
	0xd6, 0x7f, 0x68, 0x69, 0x0d, 0x75, 0xd8, 0x62, 0x20, 0xf0, 0xc1, 0x75, 0x4e, 0xdc, 0x47, 0x08,
	0x22, 0x04, 0x09, 0x9a, 0x3d, 0x24, 0xd3, 0x1b, 0xd9, 0x2b, 0xb9, 0x71, 0x01, 0xa5, 0x6e, 0x9a,
	0x80, 0x94, 0xac, 0xf9, 0x1b, 0x9f, 0xbc, 0x8c, 0x41, 0xa2, 0xe2, 0x77, 0x63, 0xea, 0x9a, 0x3a,
	0xdd, 0x7d, 0x12, 0x52, 0xc6, 0x8a, 0xf3, 0x72, 0xaa, 0x16, 0x70, 0x22, 0x1e, 0x46, 0xf5, 0x35,
	0xc2, 0xf5, 0xad, 0x3d, 0xba, 0x91, 0xe6, 0xc8, 0x6f, 0x8e, 0x19, 0x21, 0xf8, 0x45, 0xab, 0xb7,
	0x98, 0xb1, 0x82, 0x95, 0x93, 0x3a, 0xbe, 0xc5, 0x33, 0xbf, 0xb6, 0x6e, 0xa3, 0xed, 0x3a, 0x75,
	0xcb, 0xce, 0x0a, 0x56, 0x4e, 0xd5, 0x6c, 0xfc, 0xd3, 0x1b, 0xe8, 0x15, 0xec, 0xdb, 0xc3, 0xf2,
	0x40, 0xd4, 0x57, 0x51, 0x48, 0xd7, 0x6a, 0xf9, 0x3d, 0xe4, 0xec, 0x67, 0xc8, 0xd9, 0xef, 0x90,
	0x33, 0xfe, 0x68, 0xdc, 0xc1, 0xf4, 0xc1, 0x7d, 0xed, 0x4e, 0x15, 0x5f, 0xfd, 0xb7, 0xa2, 0x6a,
	0x3f, 0x43, 0xc5, 0xde, 0x2f, 0xe3, 0x1e, 0x4f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xc0,
	0xab, 0x67, 0x6c, 0x01, 0x00, 0x00,
}

func (m *Listeners) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Listeners) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ListenerStatuses) > 0 {
		for _, msg := range m.ListenerStatuses {
			dAtA[i] = 0xa
			i++
			i = encodeVarintListeners(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ListenerStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListenerStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintListeners(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.LocalAddress != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintListeners(dAtA, i, uint64(m.LocalAddress.Size()))
		n1, err := m.LocalAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintListeners(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Listeners) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ListenerStatuses) > 0 {
		for _, e := range m.ListenerStatuses {
			l = e.Size()
			n += 1 + l + sovListeners(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListenerStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovListeners(uint64(l))
	}
	if m.LocalAddress != nil {
		l = m.LocalAddress.Size()
		n += 1 + l + sovListeners(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovListeners(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozListeners(x uint64) (n int) {
	return sovListeners(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Listeners) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListeners
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
			return fmt.Errorf("proto: Listeners: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Listeners: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenerStatuses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenerStatuses = append(m.ListenerStatuses, &ListenerStatus{})
			if err := m.ListenerStatuses[len(m.ListenerStatuses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListeners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListenerStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowListeners
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
			return fmt.Errorf("proto: ListenerStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListenerStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
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
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowListeners
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthListeners
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthListeners
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LocalAddress == nil {
				m.LocalAddress = &core.Address{}
			}
			if err := m.LocalAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipListeners(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthListeners
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipListeners(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowListeners
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
					return 0, ErrIntOverflowListeners
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowListeners
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
				return 0, ErrInvalidLengthListeners
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthListeners
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowListeners
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipListeners(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthListeners
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthListeners = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowListeners   = fmt.Errorf("proto: integer overflow")
)
