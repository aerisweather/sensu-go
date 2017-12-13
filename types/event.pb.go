// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: event.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An Event is the encapsulating type sent across the Sensu websocket transport.
type Event struct {
	// Timestamp is the time in seconds since the Epoch.
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Entity describes the entity in which the event occurred.
	Entity *Entity `protobuf:"bytes,2,opt,name=entity" json:"entity,omitempty"`
	// Check describes the result of a check; if event is associated to check execution.
	Check *Check `protobuf:"bytes,3,opt,name=check" json:"check,omitempty"`
	// Metrics ...
	Metrics *Metrics `protobuf:"bytes,4,opt,name=metrics" json:"metrics,omitempty"`
	// Silenced is a list of silenced entry ids (subscription and check name)
	Silenced []string `protobuf:"bytes,5,rep,name=silenced" json:"silenced,omitempty"`
	// Hooks describes the results of multiple hooks; if event is associated to hook execution.
	Hooks []*Hook `protobuf:"bytes,6,rep,name=hooks" json:"hooks,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorEvent, []int{0} }

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Event) GetCheck() *Check {
	if m != nil {
		return m.Check
	}
	return nil
}

func (m *Event) GetMetrics() *Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Event) GetSilenced() []string {
	if m != nil {
		return m.Silenced
	}
	return nil
}

func (m *Event) GetHooks() []*Hook {
	if m != nil {
		return m.Hooks
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "sensu.types.Event")
}
func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if !this.Entity.Equal(that1.Entity) {
		return false
	}
	if !this.Check.Equal(that1.Check) {
		return false
	}
	if !this.Metrics.Equal(that1.Metrics) {
		return false
	}
	if len(this.Silenced) != len(that1.Silenced) {
		return false
	}
	for i := range this.Silenced {
		if this.Silenced[i] != that1.Silenced[i] {
			return false
		}
	}
	if len(this.Hooks) != len(that1.Hooks) {
		return false
	}
	for i := range this.Hooks {
		if !this.Hooks[i].Equal(that1.Hooks[i]) {
			return false
		}
	}
	return true
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Timestamp))
	}
	if m.Entity != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Entity.Size()))
		n1, err := m.Entity.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Check != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Check.Size()))
		n2, err := m.Check.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Metrics != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEvent(dAtA, i, uint64(m.Metrics.Size()))
		n3, err := m.Metrics.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Silenced) > 0 {
		for _, s := range m.Silenced {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Hooks) > 0 {
		for _, msg := range m.Hooks {
			dAtA[i] = 0x32
			i++
			i = encodeVarintEvent(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedEvent(r randyEvent, easy bool) *Event {
	this := &Event{}
	this.Timestamp = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Timestamp *= -1
	}
	if r.Intn(10) != 0 {
		this.Entity = NewPopulatedEntity(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Check = NewPopulatedCheck(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Metrics = NewPopulatedMetrics(r, easy)
	}
	v1 := r.Intn(10)
	this.Silenced = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Silenced[i] = string(randStringEvent(r))
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(5)
		this.Hooks = make([]*Hook, v2)
		for i := 0; i < v2; i++ {
			this.Hooks[i] = NewPopulatedHook(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyEvent interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneEvent(r randyEvent) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringEvent(r randyEvent) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneEvent(r)
	}
	return string(tmps)
}
func randUnrecognizedEvent(r randyEvent, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldEvent(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldEvent(dAtA []byte, r randyEvent, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateEvent(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateEvent(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Event) Size() (n int) {
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovEvent(uint64(m.Timestamp))
	}
	if m.Entity != nil {
		l = m.Entity.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Check != nil {
		l = m.Check.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Metrics != nil {
		l = m.Metrics.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.Silenced) > 0 {
		for _, s := range m.Silenced {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if len(m.Hooks) > 0 {
		for _, e := range m.Hooks {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func sovEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entity == nil {
				m.Entity = &Entity{}
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Check", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Check == nil {
				m.Check = &Check{}
			}
			if err := m.Check.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metrics == nil {
				m.Metrics = &Metrics{}
			}
			if err := m.Metrics.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Silenced", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Silenced = append(m.Silenced, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hooks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hooks = append(m.Hooks, &Hook{})
			if err := m.Hooks[len(m.Hooks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvent
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
				next, err := skipEvent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("event.proto", fileDescriptorEvent) }

var fileDescriptorEvent = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x71, 0xd3, 0x04, 0xea, 0xc0, 0x80, 0x61, 0xb0, 0x2a, 0x64, 0x22, 0x58, 0xb2, 0xd4,
	0x15, 0x85, 0x13, 0x14, 0x55, 0x62, 0x61, 0xe9, 0xc8, 0x46, 0xc2, 0x23, 0x89, 0x20, 0x71, 0x54,
	0x3b, 0x48, 0x3d, 0x05, 0x2b, 0x47, 0xe0, 0x08, 0x1c, 0xa1, 0x23, 0x27, 0x40, 0x10, 0x2e, 0xc1,
	0x88, 0xf2, 0x62, 0x4a, 0xb3, 0xf9, 0x7f, 0xff, 0xff, 0x3d, 0xfd, 0x7e, 0xd4, 0x87, 0x27, 0x28,
	0x8c, 0x2c, 0x17, 0xca, 0x28, 0xe6, 0x6b, 0x28, 0x74, 0x25, 0xcd, 0xb2, 0x04, 0x3d, 0x1c, 0x25,
	0x99, 0x49, 0xab, 0x48, 0xc6, 0x2a, 0x1f, 0x27, 0x2a, 0x51, 0x63, 0xcc, 0x44, 0xd5, 0x3d, 0x2a,
	0x14, 0xf8, 0x6a, 0xd9, 0xe1, 0x2e, 0x14, 0x26, 0x33, 0x4b, 0xab, 0xfc, 0x38, 0x85, 0xf8, 0xc1,
	0x8a, 0xbd, 0x1c, 0xcc, 0x22, 0x8b, 0xb5, 0x95, 0x34, 0x55, 0xca, 0x5a, 0x27, 0xcf, 0x3d, 0xea,
	0xce, 0x9a, 0x06, 0xec, 0x88, 0x0e, 0x4c, 0x96, 0x83, 0x36, 0xb7, 0x79, 0xc9, 0x49, 0x40, 0x42,
	0x67, 0xfe, 0x3f, 0x60, 0x67, 0xd4, 0x6b, 0xf7, 0xf3, 0x5e, 0x40, 0x42, 0x7f, 0x72, 0x20, 0x37,
	0xaa, 0xca, 0x19, 0x5a, 0xd3, 0xfe, 0xea, 0xe3, 0x98, 0xcc, 0x6d, 0x90, 0x49, 0xea, 0x62, 0x09,
	0xee, 0x20, 0xc1, 0x3a, 0xc4, 0x65, 0xe3, 0x58, 0xa0, 0x8d, 0xb1, 0x0b, 0xba, 0x6d, 0x7b, 0xf2,
	0x3e, 0x12, 0x87, 0x1d, 0xe2, 0xba, 0xf5, 0x2c, 0xf3, 0x17, 0x65, 0x01, 0xdd, 0xd1, 0xd9, 0x23,
	0x14, 0x31, 0xdc, 0x71, 0x37, 0x70, 0xc2, 0x81, 0x0d, 0xac, 0xa7, 0x6c, 0x44, 0xdd, 0xe6, 0xc3,
	0x9a, 0x7b, 0x81, 0x13, 0xfa, 0x93, 0xfd, 0xce, 0xd6, 0x2b, 0xa5, 0xd6, 0x35, 0x30, 0x35, 0x3d,
	0xfd, 0xf9, 0x12, 0xe4, 0xb5, 0x16, 0xe4, 0xad, 0x16, 0x64, 0x55, 0x0b, 0xf2, 0x5e, 0x0b, 0xf2,
	0x59, 0x0b, 0xf2, 0xf2, 0x2d, 0xb6, 0x6e, 0x5c, 0xc4, 0x22, 0x0f, 0xaf, 0x77, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xa9, 0xd8, 0xbf, 0x58, 0xbe, 0x01, 0x00, 0x00,
}
