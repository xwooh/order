// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: common.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SortDirection int32

const (
	SortDirection_UNKNOWN    SortDirection = 0
	SortDirection_ASCENDING  SortDirection = 1
	SortDirection_DESCENDING SortDirection = 2
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "UNKNOWN",
		1: "ASCENDING",
		2: "DESCENDING",
	}
	SortDirection_value = map[string]int32{
		"UNKNOWN":    0,
		"ASCENDING":  1,
		"DESCENDING": 2,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type PlatformType int32

const (
	PlatformType_PLATFORM_TYPE_UNKNOWN            PlatformType = 0
	PlatformType_PLATFORM_TYPE_MEITUAN_SHANGOU    PlatformType = 1  // 美团闪购
	PlatformType_PLATFORM_TYPE_ELEME_WAIMAI       PlatformType = 2  // 饿了么外卖
	PlatformType_PLATFORM_TYPE_ELEME_RETAIL       PlatformType = 3  // 饿了么新零售
	PlatformType_PLATFORM_TYPE_JINGDONG_DAOJIA    PlatformType = 4  // 京东到家
	PlatformType_PLATFORM_TYPE_MEITUAN_WAIMAI     PlatformType = 5  // 美团外卖
	PlatformType_PLATFORM_TYPE_TIANMAO_CHAOSHI    PlatformType = 6  // 天猫超市
	PlatformType_PLATFORM_TYPE_YOUZAN             PlatformType = 7  // 有赞微商城
	PlatformType_PLATFORM_TYPE_YOUZAN_RETAIL      PlatformType = 8  // 有赞零售
	PlatformType_PLATFORM_TYPE_JINGDONG_OMNI      PlatformType = 9  // 京东全渠道
	PlatformType_PLATFORM_TYPE_THIRD_PARTY        PlatformType = 10 // 第三方渠道
	PlatformType_PLATFORM_TYPE_DOUDIAN            PlatformType = 11 // 抖店
	PlatformType_PLATFORM_TYPE_TIANMAO_NEW_RETAIL PlatformType = 12 // 天猫新零售
)

// Enum value maps for PlatformType.
var (
	PlatformType_name = map[int32]string{
		0:  "PLATFORM_TYPE_UNKNOWN",
		1:  "PLATFORM_TYPE_MEITUAN_SHANGOU",
		2:  "PLATFORM_TYPE_ELEME_WAIMAI",
		3:  "PLATFORM_TYPE_ELEME_RETAIL",
		4:  "PLATFORM_TYPE_JINGDONG_DAOJIA",
		5:  "PLATFORM_TYPE_MEITUAN_WAIMAI",
		6:  "PLATFORM_TYPE_TIANMAO_CHAOSHI",
		7:  "PLATFORM_TYPE_YOUZAN",
		8:  "PLATFORM_TYPE_YOUZAN_RETAIL",
		9:  "PLATFORM_TYPE_JINGDONG_OMNI",
		10: "PLATFORM_TYPE_THIRD_PARTY",
		11: "PLATFORM_TYPE_DOUDIAN",
		12: "PLATFORM_TYPE_TIANMAO_NEW_RETAIL",
	}
	PlatformType_value = map[string]int32{
		"PLATFORM_TYPE_UNKNOWN":            0,
		"PLATFORM_TYPE_MEITUAN_SHANGOU":    1,
		"PLATFORM_TYPE_ELEME_WAIMAI":       2,
		"PLATFORM_TYPE_ELEME_RETAIL":       3,
		"PLATFORM_TYPE_JINGDONG_DAOJIA":    4,
		"PLATFORM_TYPE_MEITUAN_WAIMAI":     5,
		"PLATFORM_TYPE_TIANMAO_CHAOSHI":    6,
		"PLATFORM_TYPE_YOUZAN":             7,
		"PLATFORM_TYPE_YOUZAN_RETAIL":      8,
		"PLATFORM_TYPE_JINGDONG_OMNI":      9,
		"PLATFORM_TYPE_THIRD_PARTY":        10,
		"PLATFORM_TYPE_DOUDIAN":            11,
		"PLATFORM_TYPE_TIANMAO_NEW_RETAIL": 12,
	}
)

func (x PlatformType) Enum() *PlatformType {
	p := new(PlatformType)
	*p = x
	return p
}

func (x PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (PlatformType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformType.Descriptor instead.
func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

// 配送平台类型
type DeliveryPlatformType int32

const (
	// 其他，未知
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_UNKNOWN DeliveryPlatformType = 0
	// 蜂鸟跑腿
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_FENGNIAO DeliveryPlatformType = 1
	// 美团跑腿
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_MEITUAN DeliveryPlatformType = 2
	// 达达
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_DADA DeliveryPlatformType = 3
	// 闪送
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_SHANSONG DeliveryPlatformType = 4
	// 顺丰同城
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_SHUNFENG_TONGCHENG DeliveryPlatformType = 5
	// 蜂鸟独立结算
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_FENGNIAO_IS DeliveryPlatformType = 6
	// 美团配送
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_MEITUAN_PEISONG DeliveryPlatformType = 7
	// 土星配送
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_SATURN DeliveryPlatformType = 8
	// 蜂鸟月结
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_FENGNIAO_NEW DeliveryPlatformType = 9
	// 顺丰快递
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_SHUNFENG DeliveryPlatformType = 10
	// 快递100
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_KUAIDI100 DeliveryPlatformType = 11
	// 京东快递
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_JINGDONG DeliveryPlatformType = 12
	// 美团专送
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_MEITUAN_MANAGED DeliveryPlatformType = 13
	// 饿百专送
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_ELEME_MANAGED DeliveryPlatformType = 14
	// 京东到家专送
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_JINGDONG_MANAGED DeliveryPlatformType = 15
	// 申通
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_SHENTONG DeliveryPlatformType = 16
	// 圆通
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_YUANTONG DeliveryPlatformType = 17
	// 中通
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_ZHONGTONG DeliveryPlatformType = 18
	// 韵达
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_YUNDA DeliveryPlatformType = 19
	// 德邦
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_DEBANG DeliveryPlatformType = 20
	// 极兔
	DeliveryPlatformType_DELIVERY_PLATFORM_TYPE_JITU DeliveryPlatformType = 21
)

// Enum value maps for DeliveryPlatformType.
var (
	DeliveryPlatformType_name = map[int32]string{
		0:  "DELIVERY_PLATFORM_TYPE_UNKNOWN",
		1:  "DELIVERY_PLATFORM_TYPE_FENGNIAO",
		2:  "DELIVERY_PLATFORM_TYPE_MEITUAN",
		3:  "DELIVERY_PLATFORM_TYPE_DADA",
		4:  "DELIVERY_PLATFORM_TYPE_SHANSONG",
		5:  "DELIVERY_PLATFORM_TYPE_SHUNFENG_TONGCHENG",
		6:  "DELIVERY_PLATFORM_TYPE_FENGNIAO_IS",
		7:  "DELIVERY_PLATFORM_TYPE_MEITUAN_PEISONG",
		8:  "DELIVERY_PLATFORM_TYPE_SATURN",
		9:  "DELIVERY_PLATFORM_TYPE_FENGNIAO_NEW",
		10: "DELIVERY_PLATFORM_TYPE_SHUNFENG",
		11: "DELIVERY_PLATFORM_TYPE_KUAIDI100",
		12: "DELIVERY_PLATFORM_TYPE_JINGDONG",
		13: "DELIVERY_PLATFORM_TYPE_MEITUAN_MANAGED",
		14: "DELIVERY_PLATFORM_TYPE_ELEME_MANAGED",
		15: "DELIVERY_PLATFORM_TYPE_JINGDONG_MANAGED",
		16: "DELIVERY_PLATFORM_TYPE_SHENTONG",
		17: "DELIVERY_PLATFORM_TYPE_YUANTONG",
		18: "DELIVERY_PLATFORM_TYPE_ZHONGTONG",
		19: "DELIVERY_PLATFORM_TYPE_YUNDA",
		20: "DELIVERY_PLATFORM_TYPE_DEBANG",
		21: "DELIVERY_PLATFORM_TYPE_JITU",
	}
	DeliveryPlatformType_value = map[string]int32{
		"DELIVERY_PLATFORM_TYPE_UNKNOWN":            0,
		"DELIVERY_PLATFORM_TYPE_FENGNIAO":           1,
		"DELIVERY_PLATFORM_TYPE_MEITUAN":            2,
		"DELIVERY_PLATFORM_TYPE_DADA":               3,
		"DELIVERY_PLATFORM_TYPE_SHANSONG":           4,
		"DELIVERY_PLATFORM_TYPE_SHUNFENG_TONGCHENG": 5,
		"DELIVERY_PLATFORM_TYPE_FENGNIAO_IS":        6,
		"DELIVERY_PLATFORM_TYPE_MEITUAN_PEISONG":    7,
		"DELIVERY_PLATFORM_TYPE_SATURN":             8,
		"DELIVERY_PLATFORM_TYPE_FENGNIAO_NEW":       9,
		"DELIVERY_PLATFORM_TYPE_SHUNFENG":           10,
		"DELIVERY_PLATFORM_TYPE_KUAIDI100":          11,
		"DELIVERY_PLATFORM_TYPE_JINGDONG":           12,
		"DELIVERY_PLATFORM_TYPE_MEITUAN_MANAGED":    13,
		"DELIVERY_PLATFORM_TYPE_ELEME_MANAGED":      14,
		"DELIVERY_PLATFORM_TYPE_JINGDONG_MANAGED":   15,
		"DELIVERY_PLATFORM_TYPE_SHENTONG":           16,
		"DELIVERY_PLATFORM_TYPE_YUANTONG":           17,
		"DELIVERY_PLATFORM_TYPE_ZHONGTONG":          18,
		"DELIVERY_PLATFORM_TYPE_YUNDA":              19,
		"DELIVERY_PLATFORM_TYPE_DEBANG":             20,
		"DELIVERY_PLATFORM_TYPE_JITU":               21,
	}
)

func (x DeliveryPlatformType) Enum() *DeliveryPlatformType {
	p := new(DeliveryPlatformType)
	*p = x
	return p
}

func (x DeliveryPlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryPlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (DeliveryPlatformType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x DeliveryPlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryPlatformType.Descriptor instead.
func (DeliveryPlatformType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

type FilterKeyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldToFilter string `protobuf:"bytes,1,opt,name=field_to_filter,json=fieldToFilter,proto3" json:"field_to_filter,omitempty"`
	Keyword       string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *FilterKeyword) Reset() {
	*x = FilterKeyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterKeyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterKeyword) ProtoMessage() {}

func (x *FilterKeyword) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterKeyword.ProtoReflect.Descriptor instead.
func (*FilterKeyword) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *FilterKeyword) GetFieldToFilter() string {
	if x != nil {
		return x.FieldToFilter
	}
	return ""
}

func (x *FilterKeyword) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

type SortOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldToSort string        `protobuf:"bytes,1,opt,name=field_to_sort,json=fieldToSort,proto3" json:"field_to_sort,omitempty"`
	Direction   SortDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=order.common.SortDirection" json:"direction,omitempty"`
}

func (x *SortOption) Reset() {
	*x = SortOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortOption) ProtoMessage() {}

func (x *SortOption) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SortOption.ProtoReflect.Descriptor instead.
func (*SortOption) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *SortOption) GetFieldToSort() string {
	if x != nil {
		return x.FieldToSort
	}
	return ""
}

func (x *SortOption) GetDirection() SortDirection {
	if x != nil {
		return x.Direction
	}
	return SortDirection_UNKNOWN
}

type ListRequestWithFilterAndSort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page        int32            `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize    int32            `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Keyword     string           `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	FieldFilter []*FilterKeyword `protobuf:"bytes,4,rep,name=field_filter,json=fieldFilter,proto3" json:"field_filter,omitempty"`
	SortBy      *SortOption      `protobuf:"bytes,5,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
}

func (x *ListRequestWithFilterAndSort) Reset() {
	*x = ListRequestWithFilterAndSort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequestWithFilterAndSort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequestWithFilterAndSort) ProtoMessage() {}

func (x *ListRequestWithFilterAndSort) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequestWithFilterAndSort.ProtoReflect.Descriptor instead.
func (*ListRequestWithFilterAndSort) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequestWithFilterAndSort) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequestWithFilterAndSort) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRequestWithFilterAndSort) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListRequestWithFilterAndSort) GetFieldFilter() []*FilterKeyword {
	if x != nil {
		return x.FieldFilter
	}
	return nil
}

func (x *ListRequestWithFilterAndSort) GetSortBy() *SortOption {
	if x != nil {
		return x.SortBy
	}
	return nil
}

type OperatorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *OperatorInfo) Reset() {
	*x = OperatorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorInfo) ProtoMessage() {}

func (x *OperatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorInfo.ProtoReflect.Descriptor instead.
func (*OperatorInfo) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *OperatorInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0d,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x6f, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x6b, 0x0a, 0x0a, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x6f, 0x53, 0x6f, 0x72,
	0x74, 0x12, 0x39, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xdc, 0x01, 0x0a,
	0x1c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74,
	0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x22, 0x27, 0x0a, 0x0c, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x2a, 0x3b, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x53, 0x43, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x2a, 0xb0, 0x03, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x21, 0x0a,
	0x1d, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x45, 0x49, 0x54, 0x55, 0x41, 0x4e, 0x5f, 0x53, 0x48, 0x41, 0x4e, 0x47, 0x4f, 0x55, 0x10, 0x01,
	0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x5f, 0x57, 0x41, 0x49, 0x4d, 0x41, 0x49, 0x10, 0x02,
	0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x5f, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x03,
	0x12, 0x21, 0x0a, 0x1d, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4a, 0x49, 0x4e, 0x47, 0x44, 0x4f, 0x4e, 0x47, 0x5f, 0x44, 0x41, 0x4f, 0x4a, 0x49,
	0x41, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x49, 0x54, 0x55, 0x41, 0x4e, 0x5f, 0x57, 0x41, 0x49,
	0x4d, 0x41, 0x49, 0x10, 0x05, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x41, 0x4e, 0x4d, 0x41, 0x4f, 0x5f, 0x43,
	0x48, 0x41, 0x4f, 0x53, 0x48, 0x49, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x59, 0x4f, 0x55, 0x5a, 0x41, 0x4e,
	0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x59, 0x4f, 0x55, 0x5a, 0x41, 0x4e, 0x5f, 0x52, 0x45, 0x54, 0x41, 0x49,
	0x4c, 0x10, 0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x49, 0x4e, 0x47, 0x44, 0x4f, 0x4e, 0x47, 0x5f, 0x4f, 0x4d,
	0x4e, 0x49, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54,
	0x59, 0x10, 0x0a, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x44, 0x49, 0x41, 0x4e, 0x10, 0x0b, 0x12, 0x24,
	0x0a, 0x20, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x49, 0x41, 0x4e, 0x4d, 0x41, 0x4f, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x52, 0x45, 0x54, 0x41,
	0x49, 0x4c, 0x10, 0x0c, 0x2a, 0xe1, 0x06, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a,
	0x1e, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x45, 0x4e, 0x47,
	0x4e, 0x49, 0x41, 0x4f, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45,
	0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4d, 0x45, 0x49, 0x54, 0x55, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x44, 0x41, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x41, 0x4e, 0x53, 0x4f, 0x4e, 0x47, 0x10, 0x04,
	0x12, 0x2d, 0x0a, 0x29, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41,
	0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x55, 0x4e, 0x46,
	0x45, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x4e, 0x47, 0x43, 0x48, 0x45, 0x4e, 0x47, 0x10, 0x05, 0x12,
	0x26, 0x0a, 0x22, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x45, 0x4e, 0x47, 0x4e, 0x49,
	0x41, 0x4f, 0x5f, 0x49, 0x53, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x45, 0x49, 0x54, 0x55, 0x41, 0x4e, 0x5f, 0x50, 0x45, 0x49, 0x53, 0x4f, 0x4e,
	0x47, 0x10, 0x07, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x41,
	0x54, 0x55, 0x52, 0x4e, 0x10, 0x08, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45,
	0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x46, 0x45, 0x4e, 0x47, 0x4e, 0x49, 0x41, 0x4f, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x09, 0x12,
	0x23, 0x0a, 0x1f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x55, 0x4e, 0x46, 0x45,
	0x4e, 0x47, 0x10, 0x0a, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59,
	0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b,
	0x55, 0x41, 0x49, 0x44, 0x49, 0x31, 0x30, 0x30, 0x10, 0x0b, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x45,
	0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x49, 0x4e, 0x47, 0x44, 0x4f, 0x4e, 0x47, 0x10, 0x0c, 0x12,
	0x2a, 0x0a, 0x26, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x49, 0x54, 0x55, 0x41,
	0x4e, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x28, 0x0a, 0x24, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x5f, 0x4d, 0x41, 0x4e, 0x41,
	0x47, 0x45, 0x44, 0x10, 0x0e, 0x12, 0x2b, 0x0a, 0x27, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52,
	0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4a, 0x49, 0x4e, 0x47, 0x44, 0x4f, 0x4e, 0x47, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x44,
	0x10, 0x0f, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x48, 0x45,
	0x4e, 0x54, 0x4f, 0x4e, 0x47, 0x10, 0x10, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x59, 0x55, 0x41, 0x4e, 0x54, 0x4f, 0x4e, 0x47, 0x10, 0x11, 0x12, 0x24, 0x0a, 0x20,
	0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x5a, 0x48, 0x4f, 0x4e, 0x47, 0x54, 0x4f, 0x4e, 0x47,
	0x10, 0x12, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x5f, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x59, 0x55, 0x4e,
	0x44, 0x41, 0x10, 0x13, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59,
	0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44,
	0x45, 0x42, 0x41, 0x4e, 0x47, 0x10, 0x14, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x4c, 0x49, 0x56,
	0x45, 0x52, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4a, 0x49, 0x54, 0x55, 0x10, 0x15, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(SortDirection)(0),                   // 0: order.common.SortDirection
	(PlatformType)(0),                    // 1: order.common.PlatformType
	(DeliveryPlatformType)(0),            // 2: order.common.DeliveryPlatformType
	(*FilterKeyword)(nil),                // 3: order.common.FilterKeyword
	(*SortOption)(nil),                   // 4: order.common.SortOption
	(*ListRequestWithFilterAndSort)(nil), // 5: order.common.ListRequestWithFilterAndSort
	(*OperatorInfo)(nil),                 // 6: order.common.OperatorInfo
}
var file_common_proto_depIdxs = []int32{
	0, // 0: order.common.SortOption.direction:type_name -> order.common.SortDirection
	3, // 1: order.common.ListRequestWithFilterAndSort.field_filter:type_name -> order.common.FilterKeyword
	4, // 2: order.common.ListRequestWithFilterAndSort.sort_by:type_name -> order.common.SortOption
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterKeyword); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortOption); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequestWithFilterAndSort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}