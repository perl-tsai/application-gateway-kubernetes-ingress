// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_group_criterion_status.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// The possible statuses of an AdGroupCriterion.
type AdGroupCriterionStatusEnum_AdGroupCriterionStatus int32

const (
	// No value has been specified.
	AdGroupCriterionStatusEnum_UNSPECIFIED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupCriterionStatusEnum_UNKNOWN AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 1
	// The ad group criterion is enabled.
	AdGroupCriterionStatusEnum_ENABLED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 2
	// The ad group criterion is paused.
	AdGroupCriterionStatusEnum_PAUSED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 3
	// The ad group criterion is removed.
	AdGroupCriterionStatusEnum_REMOVED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 4
)

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}
var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupCriterionStatusEnum_AdGroupCriterionStatus) String() string {
	return proto.EnumName(AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, int32(x))
}
func (AdGroupCriterionStatusEnum_AdGroupCriterionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_criterion_status_22bc270ca4c6126c, []int{0, 0}
}

// Message describing AdGroupCriterion statuses.
type AdGroupCriterionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionStatusEnum) Reset()         { *m = AdGroupCriterionStatusEnum{} }
func (m *AdGroupCriterionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionStatusEnum) ProtoMessage()    {}
func (*AdGroupCriterionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_group_criterion_status_22bc270ca4c6126c, []int{0}
}
func (m *AdGroupCriterionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Marshal(b, m, deterministic)
}
func (dst *AdGroupCriterionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionStatusEnum.Merge(dst, src)
}
func (m *AdGroupCriterionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Size(m)
}
func (m *AdGroupCriterionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdGroupCriterionStatusEnum)(nil), "google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus", AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_group_criterion_status.proto", fileDescriptor_ad_group_criterion_status_22bc270ca4c6126c)
}

var fileDescriptor_ad_group_criterion_status_22bc270ca4c6126c = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0x83, 0x30,
	0x1c, 0xc6, 0x1d, 0x33, 0x33, 0xe9, 0x0e, 0x12, 0x0e, 0x1e, 0x34, 0x3b, 0x6c, 0x0f, 0x50, 0x48,
	0xbc, 0xd5, 0x78, 0x28, 0xa3, 0x92, 0x45, 0x65, 0x44, 0x02, 0x26, 0x86, 0x84, 0xe0, 0x4a, 0x1a,
	0x92, 0x41, 0x09, 0x85, 0x1d, 0x7c, 0x1c, 0x8f, 0x3e, 0x8a, 0x8f, 0xe2, 0xc9, 0x47, 0x30, 0x6d,
	0x85, 0xd3, 0xf4, 0x42, 0x3e, 0xfa, 0xfd, 0x7f, 0xff, 0x7c, 0xff, 0x0f, 0xdc, 0x32, 0xce, 0xd9,
	0xbe, 0xb0, 0x73, 0x2a, 0x6c, 0x2d, 0xa5, 0x3a, 0x38, 0x76, 0x51, 0xf7, 0x95, 0xb0, 0x73, 0x9a,
	0xb1, 0x96, 0xf7, 0x4d, 0xb6, 0x6b, 0xcb, 0xae, 0x68, 0x4b, 0x5e, 0x67, 0xa2, 0xcb, 0xbb, 0x5e,
	0xc0, 0xa6, 0xe5, 0x1d, 0xb7, 0x16, 0x9a, 0x81, 0x39, 0x15, 0x70, 0xc4, 0xe1, 0xc1, 0x81, 0x0a,
	0x5f, 0xbd, 0x81, 0x4b, 0x4c, 0x7d, 0xb9, 0x60, 0x3d, 0xf0, 0x91, 0xc2, 0x49, 0xdd, 0x57, 0xab,
	0x14, 0x5c, 0x1c, 0x77, 0xad, 0x73, 0x30, 0x8f, 0x83, 0x28, 0x24, 0xeb, 0xcd, 0xdd, 0x86, 0x78,
	0xe6, 0x89, 0x35, 0x07, 0x67, 0x71, 0x70, 0x1f, 0x6c, 0x9f, 0x03, 0x73, 0x22, 0x7f, 0x48, 0x80,
	0xdd, 0x07, 0xe2, 0x99, 0x86, 0x05, 0xc0, 0x2c, 0xc4, 0x71, 0x44, 0x3c, 0x73, 0x2a, 0x8d, 0x27,
	0xf2, 0xb8, 0x4d, 0x88, 0x67, 0x9e, 0xba, 0xdf, 0x13, 0xb0, 0xdc, 0xf1, 0x0a, 0xfe, 0x9b, 0xd0,
	0xbd, 0x3a, 0x9e, 0x20, 0x94, 0xd7, 0x85, 0x93, 0x17, 0xf7, 0x97, 0x66, 0x7c, 0x9f, 0xd7, 0x0c,
	0xf2, 0x96, 0xd9, 0xac, 0xa8, 0xd5, 0xed, 0x43, 0x5d, 0x4d, 0x29, 0xfe, 0x68, 0xef, 0x46, 0x7d,
	0xdf, 0x8d, 0xa9, 0x8f, 0xf1, 0x87, 0xb1, 0xf0, 0xf5, 0x2a, 0x4c, 0x05, 0xd4, 0x52, 0xaa, 0xc4,
	0x81, 0xb2, 0x0b, 0xf1, 0x39, 0xf8, 0x29, 0xa6, 0x22, 0x1d, 0xfd, 0x34, 0x71, 0x52, 0xe5, 0x7f,
	0x19, 0x4b, 0xfd, 0x88, 0x10, 0xa6, 0x02, 0xa1, 0x71, 0x02, 0xa1, 0xc4, 0x41, 0x48, 0xcd, 0xbc,
	0xce, 0x54, 0xb0, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x85, 0xd8, 0x8b, 0xd5, 0x01,
	0x00, 0x00,
}