// Code generated by protoc-gen-gogo.
// source: timestamp.proto
// DO NOT EDIT!

/*
	Package timestamp is a generated protocol buffer package.

	It is generated from these files:
		timestamp.proto

	It has these top-level messages:
		Timestamp
*/
package timestamp

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// A Timestamp represents a point in time independent of any time zone
// or calendar, represented as seconds and fractions of seconds at
// nanosecond resolution in UTC Epoch time. It is encoded using the
// Proleptic Gregorian Calendar which extends the Gregorian calendar
// backwards to year one. It is encoded assuming all minutes are 60
// seconds long, i.e. leap seconds are "smeared" so that no leap second
// table is needed for interpretation. Range is from
// 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z.
// By restricting to that range, we ensure that we can convert to
// and from  RFC 3339 date strings.
// See [https://www.ietf.org/rfc/rfc3339.txt](https://www.ietf.org/rfc/rfc3339.txt).
//
// Example 1: Compute Timestamp from POSIX `time()`.
//
//     Timestamp timestamp;
//     timestamp.set_seconds(time(NULL));
//     timestamp.set_nanos(0);
//
// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
//
//     struct timeval tv;
//     gettimeofday(&tv, NULL);
//
//     Timestamp timestamp;
//     timestamp.set_seconds(tv.tv_sec);
//     timestamp.set_nanos(tv.tv_usec * 1000);
//
// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
//
//     FILETIME ft;
//     GetSystemTimeAsFileTime(&ft);
//     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
//
//     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
//     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
//     Timestamp timestamp;
//     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
//     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
//
// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
//
//     long millis = System.currentTimeMillis();
//
//     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
//         .setNanos((int) ((millis % 1000) * 1000000)).build();
//
//
// Example 5: Compute Timestamp from current time in Python.
//
//     now = time.time()
//     seconds = int(now)
//     nanos = int((now - seconds) * 10**9)
//     timestamp = Timestamp(seconds=seconds, nanos=nanos)
//
//
type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptorTimestamp, []int{0} }

func init() {
	proto.RegisterType((*Timestamp)(nil), "docker.cluster.api.Timestamp")
}

func (m *Timestamp) Copy() *Timestamp {
	if m == nil {
		return nil
	}

	o := &Timestamp{
		Seconds: m.Seconds,
		Nanos:   m.Nanos,
	}

	return o
}

func (this *Timestamp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&timestamp.Timestamp{")
	s = append(s, "Seconds: "+fmt.Sprintf("%#v", this.Seconds)+",\n")
	s = append(s, "Nanos: "+fmt.Sprintf("%#v", this.Nanos)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTimestamp(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringTimestamp(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *Timestamp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Timestamp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Seconds != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintTimestamp(data, i, uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintTimestamp(data, i, uint64(m.Nanos))
	}
	return i, nil
}

func encodeFixed64Timestamp(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Timestamp(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTimestamp(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

func (m *Timestamp) Size() (n int) {
	var l int
	_ = l
	if m.Seconds != 0 {
		n += 1 + sovTimestamp(uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		n += 1 + sovTimestamp(uint64(m.Nanos))
	}
	return n
}

func sovTimestamp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimestamp(x uint64) (n int) {
	return sovTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Timestamp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Timestamp{`,
		`Seconds:` + fmt.Sprintf("%v", this.Seconds) + `,`,
		`Nanos:` + fmt.Sprintf("%v", this.Nanos) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTimestamp(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Timestamp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Seconds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Nanos |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimestamp(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimestamp
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
func skipTimestamp(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimestamp
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowTimestamp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimestamp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimestamp
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipTimestamp(data[start:])
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
	ErrInvalidLengthTimestamp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimestamp   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorTimestamp = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcc, 0x4d,
	0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xc9, 0x4f,
	0xce, 0x4e, 0x2d, 0xd2, 0x4b, 0xce, 0x29, 0x2d, 0x2e, 0x01, 0xd2, 0x89, 0x05, 0x99, 0x4a, 0xd6,
	0x5c, 0x9c, 0x21, 0x30, 0x65, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
	0x7e, 0xb1, 0x04, 0x13, 0x50, 0x9c, 0x35, 0x08, 0xc2, 0x71, 0x2a, 0x38, 0xf1, 0x50, 0x8e, 0xe1,
	0x06, 0x10, 0x7f, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0xf1, 0x04, 0x10, 0x5f, 0x00, 0xe2,
	0x07, 0x40, 0xcc, 0x25, 0x9c, 0x9c, 0x9f, 0xab, 0x97, 0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0x0a, 0x71,
	0x40, 0x52, 0x69, 0x9a, 0x13, 0x1f, 0xdc, 0xb6, 0x00, 0x90, 0x50, 0x00, 0xe3, 0x02, 0x46, 0xc6,
	0x1f, 0x8c, 0x8c, 0x8b, 0x98, 0x98, 0xdd, 0x03, 0x9c, 0x56, 0x31, 0xc9, 0xb9, 0x43, 0xd4, 0x07,
	0x40, 0xd5, 0xeb, 0x85, 0xa7, 0xe6, 0xe4, 0x78, 0xe7, 0xe5, 0x97, 0xe7, 0x85, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x0d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf6, 0xa8, 0xfe,
	0xdc, 0x00, 0x00, 0x00,
}
