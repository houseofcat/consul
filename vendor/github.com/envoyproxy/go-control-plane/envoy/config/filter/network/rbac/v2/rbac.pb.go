// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rbac/v2/rbac.proto

package envoy_config_filter_network_rbac_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RBAC_EnforcementType int32

const (
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	RBAC_CONTINUOUS             RBAC_EnforcementType = 1
)

var RBAC_EnforcementType_name = map[int32]string{
	0: "ONE_TIME_ON_FIRST_BYTE",
	1: "CONTINUOUS",
}

var RBAC_EnforcementType_value = map[string]int32{
	"ONE_TIME_ON_FIRST_BYTE": 0,
	"CONTINUOUS":             1,
}

func (x RBAC_EnforcementType) String() string {
	return proto.EnumName(RBAC_EnforcementType_name, int32(x))
}

func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ec60cc393c44598, []int{0, 0}
}

type RBAC struct {
	Rules                *v2.RBAC             `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v2.RBAC             `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	StatPrefix           string               `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	EnforcementType      RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.config.filter.network.rbac.v2.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec60cc393c44598, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v2.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v2.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

func (m *RBAC) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if m != nil {
		return m.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.rbac.v2.RBAC_EnforcementType", RBAC_EnforcementType_name, RBAC_EnforcementType_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.network.rbac.v2.RBAC")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rbac/v2/rbac.proto", fileDescriptor_8ec60cc393c44598)
}

var fileDescriptor_8ec60cc393c44598 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x71, 0xe8, 0xbd, 0xe8, 0xba, 0xe8, 0xde, 0x28, 0x03, 0x54, 0x91, 0x80, 0xa8, 0x2c,
	0x11, 0x83, 0x0d, 0xe9, 0xc4, 0xd0, 0x81, 0x54, 0x41, 0xea, 0x40, 0x52, 0xa5, 0xe9, 0xc0, 0x14,
	0xb9, 0x89, 0x53, 0x22, 0x5a, 0x3b, 0x72, 0xdc, 0x34, 0xd9, 0x78, 0x03, 0x56, 0x9e, 0x85, 0x07,
	0x40, 0xac, 0xbc, 0x0a, 0x23, 0x03, 0x42, 0x89, 0x83, 0x50, 0xdb, 0x81, 0x4e, 0x89, 0xfc, 0x7f,
	0xe7, 0xf8, 0x3b, 0x3e, 0x10, 0x51, 0x56, 0xf1, 0x06, 0x27, 0x9c, 0x65, 0xf9, 0x06, 0x67, 0xf9,
	0x56, 0x52, 0x81, 0x19, 0x95, 0x07, 0x2e, 0x3e, 0x62, 0xb1, 0x26, 0x09, 0xae, 0x9c, 0xee, 0x8b,
	0x0a, 0xc1, 0x25, 0x37, 0x9e, 0x77, 0x3c, 0x52, 0x3c, 0x52, 0x3c, 0xea, 0x79, 0xd4, 0x71, 0x95,
	0x63, 0x3e, 0x3b, 0x6a, 0x7a, 0xde, 0xc5, 0x7c, 0xba, 0x4f, 0x0b, 0x82, 0x09, 0x63, 0x5c, 0x12,
	0x99, 0x73, 0x56, 0xe2, 0x5d, 0xbe, 0x11, 0x44, 0xd2, 0x3e, 0x7f, 0x72, 0x96, 0x97, 0x92, 0xc8,
	0x7d, 0xd9, 0xc7, 0x8f, 0x2b, 0xb2, 0xcd, 0x53, 0x22, 0x29, 0xfe, 0xfb, 0xa3, 0x82, 0xf1, 0x37,
	0x0d, 0x0e, 0x42, 0xf7, 0xcd, 0xcc, 0x78, 0x09, 0xaf, 0xc4, 0x7e, 0x4b, 0xcb, 0x11, 0xb0, 0x80,
	0x3d, 0x74, 0x4c, 0x74, 0xa4, 0xdd, 0x7b, 0xa2, 0x16, 0x0d, 0x15, 0x68, 0x4c, 0xe1, 0xc3, 0xf2,
	0x03, 0x49, 0xf9, 0x21, 0x56, 0x85, 0xda, 0x7f, 0x0b, 0x87, 0x8a, 0x0f, 0xbb, 0x72, 0x1b, 0x0e,
	0x5b, 0xc5, 0xb8, 0x10, 0x34, 0xcb, 0xeb, 0xd1, 0x7d, 0x0b, 0xd8, 0x37, 0xee, 0x83, 0x5f, 0xee,
	0x40, 0x68, 0x16, 0x08, 0x61, 0x9b, 0x2d, 0xba, 0xc8, 0x48, 0xa1, 0x4e, 0x59, 0xc6, 0x45, 0x42,
	0x77, 0x94, 0xc9, 0x58, 0x36, 0x05, 0x1d, 0x0d, 0x2c, 0x60, 0xdf, 0x3a, 0xaf, 0xd1, 0x05, 0x8f,
	0xdb, 0xdd, 0x8d, 0xbc, 0x7f, 0x1d, 0xa2, 0xa6, 0xa0, 0xe1, 0x1d, 0x3d, 0x3e, 0x18, 0x4f, 0xe1,
	0xdd, 0x09, 0x63, 0x98, 0xf0, 0x51, 0xe0, 0x7b, 0x71, 0x34, 0x7f, 0xe7, 0xc5, 0x81, 0x1f, 0xbf,
	0x9d, 0x87, 0xcb, 0x28, 0x76, 0xdf, 0x47, 0x9e, 0x7e, 0xcf, 0xb8, 0x85, 0x70, 0x16, 0xf8, 0xd1,
	0xdc, 0x5f, 0x05, 0xab, 0xa5, 0x0e, 0xdc, 0xfa, 0xe7, 0x97, 0xdf, 0x9f, 0xaf, 0x5e, 0x18, 0xb6,
	0x32, 0xa2, 0xb5, 0xa4, 0xac, 0x6c, 0x17, 0xd1, 0x5b, 0x95, 0x27, 0x5a, 0x93, 0xaf, 0x9f, 0xbe,
	0xff, 0xb8, 0xd6, 0x74, 0x00, 0x5f, 0xe5, 0x5c, 0x8d, 0x51, 0x08, 0x5e, 0x37, 0x97, 0x4c, 0xe4,
	0xde, 0x84, 0x6b, 0x92, 0x2c, 0xda, 0x05, 0x2e, 0xc0, 0xfa, 0xba, 0xdb, 0xe4, 0xe4, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x25, 0xdd, 0x2f, 0xcb, 0x99, 0x02, 0x00, 0x00,
}