// Code generated by protoc-gen-go.
// source: response.proto
// DO NOT EDIT!

package mypackage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Response_Type int32

const (
	Response_GET Response_Type = 2
)

var Response_Type_name = map[int32]string{
	2: "GET",
}
var Response_Type_value = map[string]int32{
	"GET": 2,
}

func (x Response_Type) Enum() *Response_Type {
	p := new(Response_Type)
	*p = x
	return p
}
func (x Response_Type) String() string {
	return proto.EnumName(Response_Type_name, int32(x))
}
func (x *Response_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Response_Type_value, data, "Response_Type")
	if err != nil {
		return err
	}
	*x = Response_Type(value)
	return nil
}
func (Response_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type Response struct {
	Type             *Response_Type `protobuf:"varint,1,req,name=type,enum=dht.Response_Type" json:"type,omitempty"`
	Get              *GetResponse   `protobuf:"bytes,3,opt,name=get" json:"get,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Response) GetType() Response_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Response_GET
}

func (m *Response) GetGet() *GetResponse {
	if m != nil {
		return m.Get
	}
	return nil
}

type GetResponse struct {
	Value            *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	Success          *bool   `protobuf:"varint,2,req,name=success" json:"success,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetResponse) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *GetResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Response)(nil), "dht.Response")
	proto.RegisterType((*GetResponse)(nil), "dht.GetResponse")
	proto.RegisterEnum("dht.Response_Type", Response_Type_name, Response_Type_value)
}

func init() { proto.RegisterFile("response.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xc9, 0x28, 0x51,
	0x8a, 0xe1, 0xe2, 0x08, 0x82, 0x0a, 0x0b, 0x29, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30,
	0x2a, 0x30, 0x69, 0xf0, 0x19, 0x09, 0xe9, 0x01, 0xe5, 0xf5, 0x60, 0x92, 0x7a, 0x21, 0x40, 0x19,
	0x21, 0x59, 0x2e, 0xe6, 0xf4, 0xd4, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x01, 0xb0,
	0x02, 0xf7, 0xd4, 0x12, 0x98, 0x1a, 0x25, 0x7e, 0x2e, 0x16, 0xb0, 0x32, 0x76, 0x2e, 0x66, 0x77,
	0xd7, 0x10, 0x01, 0x26, 0x25, 0x5d, 0x2e, 0x6e, 0x24, 0x79, 0x21, 0x5e, 0x2e, 0xd6, 0xb2, 0xc4,
	0x9c, 0x52, 0x88, 0x0d, 0x9c, 0x42, 0xfc, 0x5c, 0xec, 0xc5, 0xa5, 0xc9, 0xc9, 0xa9, 0xc5, 0xc5,
	0x12, 0x4c, 0x40, 0x01, 0x0e, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x73, 0xb2, 0xea, 0xa2,
	0x00, 0x00, 0x00,
}
