// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: common/matcher/geoip/geoip.proto

package geoip

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// IP for routing decision, in CIDR form.
type CIDR struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP address, should be either 4 or 16 bytes.
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// Number of leading ones in the network mask.
	Prefix uint32 `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *CIDR) Reset() {
	*x = CIDR{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_matcher_geoip_geoip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIDR) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIDR) ProtoMessage() {}

func (x *CIDR) ProtoReflect() protoreflect.Message {
	mi := &file_common_matcher_geoip_geoip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIDR.ProtoReflect.Descriptor instead.
func (*CIDR) Descriptor() ([]byte, []int) {
	return file_common_matcher_geoip_geoip_proto_rawDescGZIP(), []int{0}
}

func (x *CIDR) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *CIDR) GetPrefix() uint32 {
	if x != nil {
		return x.Prefix
	}
	return 0
}

type GeoIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode string  `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Cidr        []*CIDR `protobuf:"bytes,2,rep,name=cidr,proto3" json:"cidr,omitempty"`
}

func (x *GeoIP) Reset() {
	*x = GeoIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_matcher_geoip_geoip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIP) ProtoMessage() {}

func (x *GeoIP) ProtoReflect() protoreflect.Message {
	mi := &file_common_matcher_geoip_geoip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIP.ProtoReflect.Descriptor instead.
func (*GeoIP) Descriptor() ([]byte, []int) {
	return file_common_matcher_geoip_geoip_proto_rawDescGZIP(), []int{1}
}

func (x *GeoIP) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GeoIP) GetCidr() []*CIDR {
	if x != nil {
		return x.Cidr
	}
	return nil
}

type GeoIPList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry []*GeoIP `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
}

func (x *GeoIPList) Reset() {
	*x = GeoIPList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_matcher_geoip_geoip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoIPList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoIPList) ProtoMessage() {}

func (x *GeoIPList) ProtoReflect() protoreflect.Message {
	mi := &file_common_matcher_geoip_geoip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoIPList.ProtoReflect.Descriptor instead.
func (*GeoIPList) Descriptor() ([]byte, []int) {
	return file_common_matcher_geoip_geoip_proto_rawDescGZIP(), []int{2}
}

func (x *GeoIPList) GetEntry() []*GeoIP {
	if x != nil {
		return x.Entry
	}
	return nil
}

var File_common_matcher_geoip_geoip_proto protoreflect.FileDescriptor

var file_common_matcher_geoip_geoip_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x22, 0x2e, 0x0a,
	0x04, 0x43, 0x49, 0x44, 0x52, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x5f, 0x0a,
	0x05, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x69, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x65,
	0x6f, 0x69, 0x70, 0x2e, 0x43, 0x49, 0x44, 0x52, 0x52, 0x04, 0x63, 0x69, 0x64, 0x72, 0x22, 0x43,
	0x0a, 0x09, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x67, 0x65, 0x6f, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x6f, 0x49, 0x50, 0x52, 0x05, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x6d, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67,
	0x65, 0x6f, 0x69, 0x70, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x67, 0x65, 0x6f, 0x69, 0x70, 0xaa, 0x02, 0x19, 0x58, 0x72, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6f,
	0x49, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_matcher_geoip_geoip_proto_rawDescOnce sync.Once
	file_common_matcher_geoip_geoip_proto_rawDescData = file_common_matcher_geoip_geoip_proto_rawDesc
)

func file_common_matcher_geoip_geoip_proto_rawDescGZIP() []byte {
	file_common_matcher_geoip_geoip_proto_rawDescOnce.Do(func() {
		file_common_matcher_geoip_geoip_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_matcher_geoip_geoip_proto_rawDescData)
	})
	return file_common_matcher_geoip_geoip_proto_rawDescData
}

var file_common_matcher_geoip_geoip_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_matcher_geoip_geoip_proto_goTypes = []interface{}{
	(*CIDR)(nil),      // 0: xray.common.matcher.geoip.CIDR
	(*GeoIP)(nil),     // 1: xray.common.matcher.geoip.GeoIP
	(*GeoIPList)(nil), // 2: xray.common.matcher.geoip.GeoIPList
}
var file_common_matcher_geoip_geoip_proto_depIdxs = []int32{
	0, // 0: xray.common.matcher.geoip.GeoIP.cidr:type_name -> xray.common.matcher.geoip.CIDR
	1, // 1: xray.common.matcher.geoip.GeoIPList.entry:type_name -> xray.common.matcher.geoip.GeoIP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_matcher_geoip_geoip_proto_init() }
func file_common_matcher_geoip_geoip_proto_init() {
	if File_common_matcher_geoip_geoip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_matcher_geoip_geoip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIDR); i {
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
		file_common_matcher_geoip_geoip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoIP); i {
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
		file_common_matcher_geoip_geoip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoIPList); i {
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
			RawDescriptor: file_common_matcher_geoip_geoip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_matcher_geoip_geoip_proto_goTypes,
		DependencyIndexes: file_common_matcher_geoip_geoip_proto_depIdxs,
		MessageInfos:      file_common_matcher_geoip_geoip_proto_msgTypes,
	}.Build()
	File_common_matcher_geoip_geoip_proto = out.File
	file_common_matcher_geoip_geoip_proto_rawDesc = nil
	file_common_matcher_geoip_geoip_proto_goTypes = nil
	file_common_matcher_geoip_geoip_proto_depIdxs = nil
}