package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Language int32

const (
	Language_UNKNOWN Language = 0
	Language_VI      Language = 1
	Language_EN      Language = 2
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "UNKNOWN",
		1: "VI",
		2: "EN",
	}
	Language_value = map[string]int32{
		"UNKNOWN": 0,
		"VI":      1,
		"EN":      2,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_seastone_api_language_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_seastone_api_language_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_seastone_api_language_proto_rawDescGZIP(), []int{0}
}

var File_seastone_api_language_proto protoreflect.FileDescriptor

var file_seastone_api_language_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x61, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6f,
	0x6e, 0x65, 0x70, 0x69, 0x65, 0x63, 0x65, 0x2e, 0x73, 0x65, 0x61, 0x73, 0x74, 0x6f, 0x6e, 0x65,
	0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2a, 0x27, 0x0a, 0x08, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x49, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x4e,
	0x10, 0x02, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x75, 0x67, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x61, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x3b, 0x73, 0x65, 0x61, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_seastone_api_language_proto_rawDescOnce sync.Once
	file_seastone_api_language_proto_rawDescData = file_seastone_api_language_proto_rawDesc
)

func file_seastone_api_language_proto_rawDescGZIP() []byte {
	file_seastone_api_language_proto_rawDescOnce.Do(func() {
		file_seastone_api_language_proto_rawDescData = protoimpl.X.CompressGZIP(file_seastone_api_language_proto_rawDescData)
	})
	return file_seastone_api_language_proto_rawDescData
}

var file_seastone_api_language_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_seastone_api_language_proto_goTypes = []interface{}{
	(Language)(0), // 0: onepiece.seastone.language.Language
}
var file_seastone_api_language_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_seastone_api_language_proto_init() }
func file_seastone_api_language_proto_init() {
	if File_seastone_api_language_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_seastone_api_language_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_seastone_api_language_proto_goTypes,
		DependencyIndexes: file_seastone_api_language_proto_depIdxs,
		EnumInfos:         file_seastone_api_language_proto_enumTypes,
	}.Build()
	File_seastone_api_language_proto = out.File
	file_seastone_api_language_proto_rawDesc = nil
	file_seastone_api_language_proto_goTypes = nil
	file_seastone_api_language_proto_depIdxs = nil
}
