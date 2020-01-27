// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1beta1/notification_config.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The type of events.
type NotificationConfig_EventType int32

const (
	// Unspecified event type.
	NotificationConfig_EVENT_TYPE_UNSPECIFIED NotificationConfig_EventType = 0
	// Events for findings.
	NotificationConfig_FINDING NotificationConfig_EventType = 1
)

var NotificationConfig_EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNSPECIFIED",
	1: "FINDING",
}

var NotificationConfig_EventType_value = map[string]int32{
	"EVENT_TYPE_UNSPECIFIED": 0,
	"FINDING":                1,
}

func (x NotificationConfig_EventType) String() string {
	return proto.EnumName(NotificationConfig_EventType_name, int32(x))
}

func (NotificationConfig_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c034a8341e15b36, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) notification configs.
//
// A notification config is a Cloud SCC resource that contains the configuration
// to send notifications for create/update events of findings, assets and etc.
type NotificationConfig struct {
	// The relative resource name of this notification config. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The description of the notification config (max of 1024 characters).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The type of events the config is for, e.g. FINDING.
	EventType NotificationConfig_EventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.cloud.securitycenter.v1p1beta1.NotificationConfig_EventType" json:"event_type,omitempty"`
	// The PubSub topic to send notifications to. Its format is
	// "projects/[project_id]/topics/[topic]".
	PubsubTopic string `protobuf:"bytes,4,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Output only. The service account that needs "pubsub.topics.publish"
	// permission to publish to the PubSub topic.
	ServiceAccount string `protobuf:"bytes,5,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// The config for triggering notifications.
	//
	// Types that are valid to be assigned to NotifyConfig:
	//	*NotificationConfig_StreamingConfig_
	NotifyConfig         isNotificationConfig_NotifyConfig `protobuf_oneof:"notify_config"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *NotificationConfig) Reset()         { *m = NotificationConfig{} }
func (m *NotificationConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig) ProtoMessage()    {}
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c034a8341e15b36, []int{0}
}

func (m *NotificationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationConfig.Unmarshal(m, b)
}
func (m *NotificationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationConfig.Marshal(b, m, deterministic)
}
func (m *NotificationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationConfig.Merge(m, src)
}
func (m *NotificationConfig) XXX_Size() int {
	return xxx_messageInfo_NotificationConfig.Size(m)
}
func (m *NotificationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationConfig proto.InternalMessageInfo

func (m *NotificationConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NotificationConfig) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NotificationConfig) GetEventType() NotificationConfig_EventType {
	if m != nil {
		return m.EventType
	}
	return NotificationConfig_EVENT_TYPE_UNSPECIFIED
}

func (m *NotificationConfig) GetPubsubTopic() string {
	if m != nil {
		return m.PubsubTopic
	}
	return ""
}

func (m *NotificationConfig) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

type isNotificationConfig_NotifyConfig interface {
	isNotificationConfig_NotifyConfig()
}

type NotificationConfig_StreamingConfig_ struct {
	StreamingConfig *NotificationConfig_StreamingConfig `protobuf:"bytes,6,opt,name=streaming_config,json=streamingConfig,proto3,oneof"`
}

func (*NotificationConfig_StreamingConfig_) isNotificationConfig_NotifyConfig() {}

func (m *NotificationConfig) GetNotifyConfig() isNotificationConfig_NotifyConfig {
	if m != nil {
		return m.NotifyConfig
	}
	return nil
}

func (m *NotificationConfig) GetStreamingConfig() *NotificationConfig_StreamingConfig {
	if x, ok := m.GetNotifyConfig().(*NotificationConfig_StreamingConfig_); ok {
		return x.StreamingConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NotificationConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NotificationConfig_StreamingConfig_)(nil),
	}
}

// The config for streaming-based notifications, which send each event as soon
// as it is detected.
type NotificationConfig_StreamingConfig struct {
	// Expression that defines the filter to apply across create/update events
	// of assets or findings as specified by the event type. The expression is a
	// list of zero or more restrictions combined via logical operators `AND`
	// and `OR`. Parentheses are supported, and `OR` has higher precedence than
	// `AND`.
	//
	// Restrictions have the form `<field> <operator> <value>` and may have a
	// `-` character in front of them to indicate negation. The fields map to
	// those defined in the corresponding resource.
	//
	// The supported operators are:
	//
	// * `=` for all value types.
	// * `>`, `<`, `>=`, `<=` for integer values.
	// * `:`, meaning substring matching, for strings.
	//
	// The supported value types are:
	//
	// * string literals in quotes.
	// * integer literals without quotes.
	// * boolean literals `true` and `false` without quotes.
	Filter               string   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotificationConfig_StreamingConfig) Reset()         { *m = NotificationConfig_StreamingConfig{} }
func (m *NotificationConfig_StreamingConfig) String() string { return proto.CompactTextString(m) }
func (*NotificationConfig_StreamingConfig) ProtoMessage()    {}
func (*NotificationConfig_StreamingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c034a8341e15b36, []int{0, 0}
}

func (m *NotificationConfig_StreamingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotificationConfig_StreamingConfig.Unmarshal(m, b)
}
func (m *NotificationConfig_StreamingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotificationConfig_StreamingConfig.Marshal(b, m, deterministic)
}
func (m *NotificationConfig_StreamingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotificationConfig_StreamingConfig.Merge(m, src)
}
func (m *NotificationConfig_StreamingConfig) XXX_Size() int {
	return xxx_messageInfo_NotificationConfig_StreamingConfig.Size(m)
}
func (m *NotificationConfig_StreamingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NotificationConfig_StreamingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NotificationConfig_StreamingConfig proto.InternalMessageInfo

func (m *NotificationConfig_StreamingConfig) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1p1beta1.NotificationConfig_EventType", NotificationConfig_EventType_name, NotificationConfig_EventType_value)
	proto.RegisterType((*NotificationConfig)(nil), "google.cloud.securitycenter.v1p1beta1.NotificationConfig")
	proto.RegisterType((*NotificationConfig_StreamingConfig)(nil), "google.cloud.securitycenter.v1p1beta1.NotificationConfig.StreamingConfig")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1beta1/notification_config.proto", fileDescriptor_4c034a8341e15b36)
}

var fileDescriptor_4c034a8341e15b36 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0x8a, 0xda, 0x40,
	0x14, 0x86, 0x37, 0xab, 0xb5, 0x38, 0xb6, 0xab, 0xcc, 0xc5, 0x92, 0xda, 0x96, 0xca, 0x82, 0xe0,
	0x42, 0x49, 0x70, 0xdb, 0xab, 0xf4, 0xa2, 0x68, 0x36, 0x6e, 0xbd, 0x11, 0x51, 0x2b, 0xb4, 0x08,
	0x61, 0x32, 0x8e, 0xe9, 0x80, 0xce, 0x84, 0xc9, 0x44, 0xb0, 0xcb, 0xbe, 0x50, 0x1f, 0xa5, 0x4f,
	0xd1, 0xeb, 0x7d, 0x84, 0x42, 0x69, 0xc9, 0x4c, 0xdc, 0x26, 0xb6, 0xd0, 0xa5, 0x77, 0x39, 0xe7,
	0xff, 0xfc, 0xcf, 0x3f, 0x33, 0x47, 0xf0, 0x36, 0xe4, 0x3c, 0x5c, 0x13, 0x1b, 0xaf, 0x79, 0xb2,
	0xb4, 0x63, 0x82, 0x13, 0x41, 0xe5, 0x0e, 0x13, 0x26, 0x89, 0xb0, 0xb7, 0xdd, 0xa8, 0x1b, 0x10,
	0x89, 0xba, 0x36, 0xe3, 0x92, 0xae, 0x28, 0x46, 0x92, 0x72, 0xe6, 0x63, 0xce, 0x56, 0x34, 0xb4,
	0x22, 0xc1, 0x25, 0x87, 0x6d, 0x6d, 0x60, 0x29, 0x03, 0xab, 0x68, 0x60, 0xdd, 0x19, 0x34, 0x9f,
	0x65, 0x73, 0x50, 0x44, 0x6d, 0xc4, 0x18, 0x97, 0xca, 0x2a, 0xd6, 0x26, 0xcd, 0x17, 0x39, 0x75,
	0x45, 0xc9, 0x7a, 0xe9, 0x07, 0xe4, 0x13, 0xda, 0x52, 0x2e, 0x32, 0xe0, 0x49, 0x0e, 0x10, 0x24,
	0xe6, 0x89, 0xc0, 0x44, 0x4b, 0x67, 0x3f, 0xcb, 0x00, 0x8e, 0x72, 0xf1, 0x5c, 0x95, 0x0e, 0x42,
	0x50, 0x66, 0x68, 0x43, 0x4c, 0xa3, 0x65, 0x74, 0xaa, 0x13, 0xf5, 0x0d, 0x5b, 0xa0, 0xb6, 0x24,
	0x31, 0x16, 0x34, 0x4a, 0x41, 0xf3, 0x58, 0x49, 0xf9, 0x16, 0x0c, 0x00, 0x20, 0x5b, 0xc2, 0xa4,
	0x2f, 0x77, 0x11, 0x31, 0x4b, 0x2d, 0xa3, 0x73, 0x72, 0xe1, 0x5a, 0xf7, 0x3a, 0xa2, 0xf5, 0x67,
	0x08, 0xcb, 0x4b, 0xbd, 0x66, 0xbb, 0x88, 0x4c, 0xaa, 0x64, 0xff, 0x09, 0x5d, 0xf0, 0x28, 0x4a,
	0x82, 0x38, 0x09, 0x7c, 0xc9, 0x23, 0x8a, 0xcd, 0x72, 0x1a, 0xa3, 0xdf, 0xfa, 0xde, 0x7b, 0x0e,
	0x9e, 0x6a, 0x21, 0x1b, 0x88, 0x22, 0x1a, 0x5b, 0x98, 0x6f, 0xec, 0x59, 0xca, 0x4d, 0x6a, 0x5a,
	0x54, 0x05, 0x7c, 0x09, 0xea, 0x31, 0x11, 0x5b, 0x8a, 0x89, 0x8f, 0x30, 0xe6, 0x09, 0x93, 0xe6,
	0x03, 0xe5, 0x53, 0xfa, 0xd6, 0x2b, 0x4d, 0x4e, 0x32, 0xad, 0xa7, 0x25, 0xb8, 0x05, 0x8d, 0x58,
	0x0a, 0x82, 0x36, 0x94, 0x85, 0xd9, 0xf3, 0x99, 0x95, 0x96, 0xd1, 0xa9, 0x5d, 0x0c, 0xff, 0xff,
	0x70, 0xd3, 0xbd, 0xa3, 0xae, 0xdf, 0x1d, 0x4d, 0xea, 0x71, 0xb1, 0xd5, 0x3c, 0x07, 0xf5, 0x03,
	0x0a, 0x9e, 0x82, 0xca, 0x8a, 0xae, 0x25, 0x11, 0xd9, 0xcb, 0x64, 0xd5, 0xd9, 0x6b, 0x50, 0xbd,
	0xbb, 0x2d, 0xd8, 0x04, 0xa7, 0xde, 0xdc, 0x1b, 0xcd, 0xfc, 0xd9, 0x87, 0xb1, 0xe7, 0xbf, 0x1f,
	0x4d, 0xc7, 0x9e, 0x3b, 0x1c, 0x0c, 0xbd, 0xcb, 0xc6, 0x11, 0xac, 0x81, 0x87, 0x83, 0xe1, 0xe8,
	0x72, 0x38, 0xba, 0x6a, 0x18, 0x8e, 0xb8, 0xed, 0x71, 0xd0, 0x3e, 0x48, 0x7d, 0x70, 0x7b, 0x03,
	0xca, 0x96, 0x94, 0x85, 0x70, 0xc0, 0x45, 0x88, 0x18, 0xfd, 0xac, 0x37, 0xcf, 0xbe, 0xce, 0x97,
	0x37, 0x85, 0x0d, 0xd7, 0x51, 0x63, 0xfb, 0xfa, 0x2f, 0x6b, 0x7f, 0xd3, 0xaf, 0x83, 0xc7, 0xaa,
	0xbf, 0xcb, 0x3a, 0xfd, 0x1f, 0x06, 0x38, 0xc7, 0x7c, 0x73, 0xbf, 0x9b, 0x1c, 0x1b, 0x1f, 0xa7,
	0x19, 0x18, 0xf2, 0x35, 0x62, 0xa1, 0xc5, 0x45, 0x68, 0x87, 0x84, 0xa9, 0x6d, 0xb6, 0x7f, 0x67,
	0xff, 0xc7, 0x5f, 0xf2, 0x4d, 0x51, 0xf8, 0x72, 0xdc, 0xbe, 0xd2, 0xae, 0xae, 0x1a, 0x3f, 0xcd,
	0x54, 0x57, 0x8f, 0x9f, 0x77, 0xc7, 0xdd, 0x7e, 0xfa, 0xb3, 0xaf, 0x7b, 0x6e, 0xa1, 0xb8, 0x45,
	0x91, 0x5b, 0xcc, 0xf7, 0xf6, 0xb7, 0xc7, 0x1d, 0xcd, 0x39, 0x8e, 0x02, 0x1d, 0xa7, 0x48, 0x3a,
	0x4e, 0x8a, 0x2a, 0xcb, 0xa0, 0xa2, 0xa2, 0xbf, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x25, 0x7a,
	0x37, 0xba, 0x4c, 0x04, 0x00, 0x00,
}