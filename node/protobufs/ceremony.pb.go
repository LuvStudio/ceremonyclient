// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ceremony.proto

package protobufs

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

// Describes the transcript of KZG ceremony execution
type CeremonyTranscript struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The active collection of powers over G1
	G1Powers []*BLS48581G1PublicKey `protobuf:"bytes,1,rep,name=g1_powers,json=g1Powers,proto3" json:"g1_powers,omitempty"`
	// The active collection of powers over G2
	G2Powers []*BLS48581G2PublicKey `protobuf:"bytes,2,rep,name=g2_powers,json=g2Powers,proto3" json:"g2_powers,omitempty"`
	// The running s^256 G1 witnesses – the choice of the 256th power is to ensure
	// combinatorial birthday paradox-based attacks are not possible. In common
	// KZG ceremonies, the collection of witnesses to PoT pubkeys produce the
	// relationship of e(w*G1, s*G2) == (s'*G1, G2), where w*s == s'. The problem
	// with this is that there are n powers under G2 (excl. the case where PoT
	// ceremonies _only_ have the first G2 power), and so the chance of collision
	// by combination to a target value for s' is feasible such that a sum of a
	// permutation of valid G2 powers could forge witness values to reach a
	// a desired outcome, as there are matching pairs of the G1 and G2 powers to
	// permute. When the number of G2 powers is low, or one, this reduces to the
	// discrete log assumption and so the only viable attack is of
	// O(sqrt(<bit size>)) per Pollard's Rho (barring any advancements), but in
	// many cases the number of G2 powers is high enough such that n! naive
	// combinations of additions are greater (and cheap, since the additions are
	// first tested in G1) than the required time of testing the discrete log,
	// and combined with many generated target values, significantly reduces the
	// amount of time required to complete the attack. This means that in
	// traditional KZG ceremonies, the last contributor to a ceremony can
	// potentially control the secret. Or, we can just track the witnesses to the
	// highest power in the ceremony and avoid the whole problem. :)
	RunningG1_256Witnesses []*BLS48581G1PublicKey `protobuf:"bytes,3,rep,name=running_g1_256_witnesses,json=runningG1256Witnesses,proto3" json:"running_g1_256_witnesses,omitempty"`
	// The running s^256 G2 powers – see notes on running_g1_256_witnesses for why
	// we do this.
	RunningG2_256Powers []*BLS48581G2PublicKey `protobuf:"bytes,4,rep,name=running_g2_256_powers,json=runningG2256Powers,proto3" json:"running_g2_256_powers,omitempty"`
}

func (x *CeremonyTranscript) Reset() {
	*x = CeremonyTranscript{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ceremony_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CeremonyTranscript) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CeremonyTranscript) ProtoMessage() {}

func (x *CeremonyTranscript) ProtoReflect() protoreflect.Message {
	mi := &file_ceremony_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CeremonyTranscript.ProtoReflect.Descriptor instead.
func (*CeremonyTranscript) Descriptor() ([]byte, []int) {
	return file_ceremony_proto_rawDescGZIP(), []int{0}
}

func (x *CeremonyTranscript) GetG1Powers() []*BLS48581G1PublicKey {
	if x != nil {
		return x.G1Powers
	}
	return nil
}

func (x *CeremonyTranscript) GetG2Powers() []*BLS48581G2PublicKey {
	if x != nil {
		return x.G2Powers
	}
	return nil
}

func (x *CeremonyTranscript) GetRunningG1_256Witnesses() []*BLS48581G1PublicKey {
	if x != nil {
		return x.RunningG1_256Witnesses
	}
	return nil
}

func (x *CeremonyTranscript) GetRunningG2_256Powers() []*BLS48581G2PublicKey {
	if x != nil {
		return x.RunningG2_256Powers
	}
	return nil
}

var File_ceremony_proto protoreflect.FileDescriptor

var file_ceremony_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x65, 0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x63, 0x65, 0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x2e, 0x70, 0x62, 0x1a, 0x0a, 0x6b,
	0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x02, 0x0a, 0x12, 0x43, 0x65,
	0x72, 0x65, 0x6d, 0x6f, 0x6e, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x12, 0x49, 0x0a, 0x09, 0x67, 0x31, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x4c,
	0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47, 0x31, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x52, 0x08, 0x67, 0x31, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x49, 0x0a, 0x09, 0x67,
	0x32, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38,
	0x31, 0x47, 0x32, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x67, 0x32,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x65, 0x0a, 0x18, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x67, 0x31, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x77, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69,
	0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e,
	0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47, 0x31, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x15, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x47,
	0x31, 0x32, 0x35, 0x36, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x5f, 0x0a,
	0x15, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x32, 0x5f, 0x32, 0x35, 0x36, 0x5f,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71,
	0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6b,
	0x65, 0x79, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x4c, 0x53, 0x34, 0x38, 0x35, 0x38, 0x31, 0x47,
	0x32, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x12, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x47, 0x32, 0x32, 0x35, 0x36, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x42, 0x3a,
	0x5a, 0x38, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72,
	0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x6c, 0x69, 0x62, 0x72, 0x69,
	0x75, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ceremony_proto_rawDescOnce sync.Once
	file_ceremony_proto_rawDescData = file_ceremony_proto_rawDesc
)

func file_ceremony_proto_rawDescGZIP() []byte {
	file_ceremony_proto_rawDescOnce.Do(func() {
		file_ceremony_proto_rawDescData = protoimpl.X.CompressGZIP(file_ceremony_proto_rawDescData)
	})
	return file_ceremony_proto_rawDescData
}

var file_ceremony_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ceremony_proto_goTypes = []interface{}{
	(*CeremonyTranscript)(nil),  // 0: quilibrium.node.ceremony.pb.CeremonyTranscript
	(*BLS48581G1PublicKey)(nil), // 1: quilibrium.node.keys.pb.BLS48581G1PublicKey
	(*BLS48581G2PublicKey)(nil), // 2: quilibrium.node.keys.pb.BLS48581G2PublicKey
}
var file_ceremony_proto_depIdxs = []int32{
	1, // 0: quilibrium.node.ceremony.pb.CeremonyTranscript.g1_powers:type_name -> quilibrium.node.keys.pb.BLS48581G1PublicKey
	2, // 1: quilibrium.node.ceremony.pb.CeremonyTranscript.g2_powers:type_name -> quilibrium.node.keys.pb.BLS48581G2PublicKey
	1, // 2: quilibrium.node.ceremony.pb.CeremonyTranscript.running_g1_256_witnesses:type_name -> quilibrium.node.keys.pb.BLS48581G1PublicKey
	2, // 3: quilibrium.node.ceremony.pb.CeremonyTranscript.running_g2_256_powers:type_name -> quilibrium.node.keys.pb.BLS48581G2PublicKey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ceremony_proto_init() }
func file_ceremony_proto_init() {
	if File_ceremony_proto != nil {
		return
	}
	file_keys_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ceremony_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CeremonyTranscript); i {
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
			RawDescriptor: file_ceremony_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ceremony_proto_goTypes,
		DependencyIndexes: file_ceremony_proto_depIdxs,
		MessageInfos:      file_ceremony_proto_msgTypes,
	}.Build()
	File_ceremony_proto = out.File
	file_ceremony_proto_rawDesc = nil
	file_ceremony_proto_goTypes = nil
	file_ceremony_proto_depIdxs = nil
}
