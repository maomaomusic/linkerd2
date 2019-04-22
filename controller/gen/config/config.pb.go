// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config // import "github.com/linkerd/linkerd2/controller/gen/config"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type All struct {
	Global               *Global  `protobuf:"bytes,1,opt,name=global,proto3" json:"global,omitempty"`
	Proxy                *Proxy   `protobuf:"bytes,2,opt,name=proxy,proto3" json:"proxy,omitempty"`
	Install              *Install `protobuf:"bytes,3,opt,name=install,proto3" json:"install,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{0}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (dst *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(dst, src)
}
func (m *All) XXX_Size() int {
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetGlobal() *Global {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *All) GetProxy() *Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *All) GetInstall() *Install {
	if m != nil {
		return m.Install
	}
	return nil
}

type Global struct {
	LinkerdNamespace string `protobuf:"bytes,1,opt,name=linkerd_namespace,json=linkerdNamespace,proto3" json:"linkerd_namespace,omitempty"`
	CniEnabled       bool   `protobuf:"varint,2,opt,name=cni_enabled,json=cniEnabled,proto3" json:"cni_enabled,omitempty"`
	// Control plane and proxy-init version
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// If present, configures identity.
	IdentityContext      *IdentityContext `protobuf:"bytes,4,opt,name=identity_context,json=identityContext,proto3" json:"identity_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{1}
}
func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (dst *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(dst, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetLinkerdNamespace() string {
	if m != nil {
		return m.LinkerdNamespace
	}
	return ""
}

func (m *Global) GetCniEnabled() bool {
	if m != nil {
		return m.CniEnabled
	}
	return false
}

func (m *Global) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Global) GetIdentityContext() *IdentityContext {
	if m != nil {
		return m.IdentityContext
	}
	return nil
}

type Proxy struct {
	ProxyImage              *Image                `protobuf:"bytes,1,opt,name=proxy_image,json=proxyImage,proto3" json:"proxy_image,omitempty"`
	ProxyInitImage          *Image                `protobuf:"bytes,2,opt,name=proxy_init_image,json=proxyInitImage,proto3" json:"proxy_init_image,omitempty"`
	ControlPort             *Port                 `protobuf:"bytes,3,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	IgnoreInboundPorts      []*Port               `protobuf:"bytes,4,rep,name=ignore_inbound_ports,json=ignoreInboundPorts,proto3" json:"ignore_inbound_ports,omitempty"`
	IgnoreOutboundPorts     []*Port               `protobuf:"bytes,5,rep,name=ignore_outbound_ports,json=ignoreOutboundPorts,proto3" json:"ignore_outbound_ports,omitempty"`
	InboundPort             *Port                 `protobuf:"bytes,6,opt,name=inbound_port,json=inboundPort,proto3" json:"inbound_port,omitempty"`
	AdminPort               *Port                 `protobuf:"bytes,7,opt,name=admin_port,json=adminPort,proto3" json:"admin_port,omitempty"`
	OutboundPort            *Port                 `protobuf:"bytes,8,opt,name=outbound_port,json=outboundPort,proto3" json:"outbound_port,omitempty"`
	Resource                *ResourceRequirements `protobuf:"bytes,9,opt,name=resource,proto3" json:"resource,omitempty"`
	ProxyUid                int64                 `protobuf:"varint,10,opt,name=proxy_uid,json=proxyUid,proto3" json:"proxy_uid,omitempty"`
	LogLevel                *LogLevel             `protobuf:"bytes,11,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	DisableExternalProfiles bool                  `protobuf:"varint,12,opt,name=disable_external_profiles,json=disableExternalProfiles,proto3" json:"disable_external_profiles,omitempty"`
	ProxyVersion            string                `protobuf:"bytes,13,opt,name=proxy_version,json=proxyVersion,proto3" json:"proxy_version,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{2}
}
func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (dst *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(dst, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetProxyImage() *Image {
	if m != nil {
		return m.ProxyImage
	}
	return nil
}

func (m *Proxy) GetProxyInitImage() *Image {
	if m != nil {
		return m.ProxyInitImage
	}
	return nil
}

func (m *Proxy) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *Proxy) GetIgnoreInboundPorts() []*Port {
	if m != nil {
		return m.IgnoreInboundPorts
	}
	return nil
}

func (m *Proxy) GetIgnoreOutboundPorts() []*Port {
	if m != nil {
		return m.IgnoreOutboundPorts
	}
	return nil
}

func (m *Proxy) GetInboundPort() *Port {
	if m != nil {
		return m.InboundPort
	}
	return nil
}

func (m *Proxy) GetAdminPort() *Port {
	if m != nil {
		return m.AdminPort
	}
	return nil
}

func (m *Proxy) GetOutboundPort() *Port {
	if m != nil {
		return m.OutboundPort
	}
	return nil
}

func (m *Proxy) GetResource() *ResourceRequirements {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Proxy) GetProxyUid() int64 {
	if m != nil {
		return m.ProxyUid
	}
	return 0
}

func (m *Proxy) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *Proxy) GetDisableExternalProfiles() bool {
	if m != nil {
		return m.DisableExternalProfiles
	}
	return false
}

func (m *Proxy) GetProxyVersion() string {
	if m != nil {
		return m.ProxyVersion
	}
	return ""
}

type Image struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	PullPolicy           string   `protobuf:"bytes,2,opt,name=pull_policy,json=pullPolicy,proto3" json:"pull_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{3}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

type Port struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{4}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ResourceRequirements struct {
	RequestCpu           string   `protobuf:"bytes,1,opt,name=request_cpu,json=requestCpu,proto3" json:"request_cpu,omitempty"`
	RequestMemory        string   `protobuf:"bytes,2,opt,name=request_memory,json=requestMemory,proto3" json:"request_memory,omitempty"`
	LimitCpu             string   `protobuf:"bytes,3,opt,name=limit_cpu,json=limitCpu,proto3" json:"limit_cpu,omitempty"`
	LimitMemory          string   `protobuf:"bytes,4,opt,name=limit_memory,json=limitMemory,proto3" json:"limit_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{5}
}
func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (dst *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(dst, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequestCpu() string {
	if m != nil {
		return m.RequestCpu
	}
	return ""
}

func (m *ResourceRequirements) GetRequestMemory() string {
	if m != nil {
		return m.RequestMemory
	}
	return ""
}

func (m *ResourceRequirements) GetLimitCpu() string {
	if m != nil {
		return m.LimitCpu
	}
	return ""
}

func (m *ResourceRequirements) GetLimitMemory() string {
	if m != nil {
		return m.LimitMemory
	}
	return ""
}

type IdentityContext struct {
	TrustDomain          string             `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	TrustAnchorsPem      string             `protobuf:"bytes,2,opt,name=trust_anchors_pem,json=trustAnchorsPem,proto3" json:"trust_anchors_pem,omitempty"`
	IssuanceLifetime     *duration.Duration `protobuf:"bytes,3,opt,name=issuance_lifetime,json=issuanceLifetime,proto3" json:"issuance_lifetime,omitempty"`
	ClockSkewAllowance   *duration.Duration `protobuf:"bytes,4,opt,name=clock_skew_allowance,json=clockSkewAllowance,proto3" json:"clock_skew_allowance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IdentityContext) Reset()         { *m = IdentityContext{} }
func (m *IdentityContext) String() string { return proto.CompactTextString(m) }
func (*IdentityContext) ProtoMessage()    {}
func (*IdentityContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{6}
}
func (m *IdentityContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentityContext.Unmarshal(m, b)
}
func (m *IdentityContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentityContext.Marshal(b, m, deterministic)
}
func (dst *IdentityContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityContext.Merge(dst, src)
}
func (m *IdentityContext) XXX_Size() int {
	return xxx_messageInfo_IdentityContext.Size(m)
}
func (m *IdentityContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityContext.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityContext proto.InternalMessageInfo

func (m *IdentityContext) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

func (m *IdentityContext) GetTrustAnchorsPem() string {
	if m != nil {
		return m.TrustAnchorsPem
	}
	return ""
}

func (m *IdentityContext) GetIssuanceLifetime() *duration.Duration {
	if m != nil {
		return m.IssuanceLifetime
	}
	return nil
}

func (m *IdentityContext) GetClockSkewAllowance() *duration.Duration {
	if m != nil {
		return m.ClockSkewAllowance
	}
	return nil
}

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{7}
}
func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (dst *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(dst, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

// Stores information about the last installation/upgrade.
//
// Useful for driving upgrades.
type Install struct {
	// The unique ID for this installation. Does not change on upgrade.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The CLI version that drove the last install or upgrade.
	CliVersion string `protobuf:"bytes,2,opt,name=cli_version,json=cliVersion,proto3" json:"cli_version,omitempty"`
	// The CLI arguments to the install (or upgrade) command, indicating the
	// installer's intent.
	Flags                []*Install_Flag `protobuf:"bytes,3,rep,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Install) Reset()         { *m = Install{} }
func (m *Install) String() string { return proto.CompactTextString(m) }
func (*Install) ProtoMessage()    {}
func (*Install) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{8}
}
func (m *Install) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install.Unmarshal(m, b)
}
func (m *Install) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install.Marshal(b, m, deterministic)
}
func (dst *Install) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install.Merge(dst, src)
}
func (m *Install) XXX_Size() int {
	return xxx_messageInfo_Install.Size(m)
}
func (m *Install) XXX_DiscardUnknown() {
	xxx_messageInfo_Install.DiscardUnknown(m)
}

var xxx_messageInfo_Install proto.InternalMessageInfo

func (m *Install) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Install) GetCliVersion() string {
	if m != nil {
		return m.CliVersion
	}
	return ""
}

func (m *Install) GetFlags() []*Install_Flag {
	if m != nil {
		return m.Flags
	}
	return nil
}

type Install_Flag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Install_Flag) Reset()         { *m = Install_Flag{} }
func (m *Install_Flag) String() string { return proto.CompactTextString(m) }
func (*Install_Flag) ProtoMessage()    {}
func (*Install_Flag) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_851cac74f39a8065, []int{8, 0}
}
func (m *Install_Flag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Install_Flag.Unmarshal(m, b)
}
func (m *Install_Flag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Install_Flag.Marshal(b, m, deterministic)
}
func (dst *Install_Flag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Install_Flag.Merge(dst, src)
}
func (m *Install_Flag) XXX_Size() int {
	return xxx_messageInfo_Install_Flag.Size(m)
}
func (m *Install_Flag) XXX_DiscardUnknown() {
	xxx_messageInfo_Install_Flag.DiscardUnknown(m)
}

var xxx_messageInfo_Install_Flag proto.InternalMessageInfo

func (m *Install_Flag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Install_Flag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*All)(nil), "linkerd2.config.All")
	proto.RegisterType((*Global)(nil), "linkerd2.config.Global")
	proto.RegisterType((*Proxy)(nil), "linkerd2.config.Proxy")
	proto.RegisterType((*Image)(nil), "linkerd2.config.Image")
	proto.RegisterType((*Port)(nil), "linkerd2.config.Port")
	proto.RegisterType((*ResourceRequirements)(nil), "linkerd2.config.ResourceRequirements")
	proto.RegisterType((*IdentityContext)(nil), "linkerd2.config.IdentityContext")
	proto.RegisterType((*LogLevel)(nil), "linkerd2.config.LogLevel")
	proto.RegisterType((*Install)(nil), "linkerd2.config.Install")
	proto.RegisterType((*Install_Flag)(nil), "linkerd2.config.Install.Flag")
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_config_851cac74f39a8065) }

var fileDescriptor_config_851cac74f39a8065 = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5f, 0x6f, 0xdc, 0x44,
	0x10, 0xd7, 0x35, 0x77, 0xc9, 0xdd, 0xdc, 0x5d, 0x93, 0x6c, 0x53, 0xea, 0x04, 0x15, 0x0e, 0xa3,
	0x4a, 0x15, 0x20, 0x1f, 0x24, 0x08, 0xaa, 0x3e, 0x11, 0xfa, 0x27, 0x8a, 0x1a, 0x20, 0x5a, 0x04,
	0x0f, 0xbc, 0x58, 0x3e, 0x7b, 0xcf, 0x5d, 0x65, 0xbd, 0xeb, 0xae, 0xd7, 0x49, 0xfa, 0x4d, 0x78,
	0xe2, 0x8d, 0x8f, 0xc1, 0xc7, 0xea, 0x3b, 0xda, 0xd9, 0x71, 0x49, 0x7b, 0xf4, 0x9e, 0xbc, 0xfb,
	0x9b, 0xdf, 0x6f, 0x66, 0x3c, 0x9e, 0x19, 0xc3, 0x9d, 0xdc, 0xe8, 0xa5, 0x2c, 0xe7, 0xe1, 0x91,
	0xd4, 0xd6, 0x38, 0xc3, 0xb6, 0x95, 0xd4, 0x17, 0xc2, 0x16, 0x87, 0x49, 0x80, 0x0f, 0x3e, 0x29,
	0x8d, 0x29, 0x95, 0x98, 0xa3, 0x79, 0xd1, 0x2e, 0xe7, 0x45, 0x6b, 0x33, 0x27, 0x8d, 0x0e, 0x82,
	0xf8, 0xcf, 0x1e, 0x6c, 0x1c, 0x2b, 0xc5, 0xe6, 0xb0, 0x59, 0x2a, 0xb3, 0xc8, 0x54, 0xd4, 0x9b,
	0xf5, 0x1e, 0x8e, 0x0f, 0xef, 0x25, 0xef, 0x79, 0x4a, 0x4e, 0xd0, 0xcc, 0x89, 0xc6, 0xbe, 0x82,
	0x41, 0x6d, 0xcd, 0xf5, 0xeb, 0xe8, 0x16, 0xf2, 0x3f, 0x5a, 0xe1, 0x9f, 0x7b, 0x2b, 0x0f, 0x24,
	0x76, 0x08, 0x5b, 0x52, 0x37, 0x2e, 0x53, 0x2a, 0xda, 0x40, 0x7e, 0xb4, 0xc2, 0x3f, 0x0d, 0x76,
	0xde, 0x11, 0xe3, 0x7f, 0x7a, 0xb0, 0x19, 0x82, 0xb2, 0x2f, 0x61, 0x97, 0xe8, 0xa9, 0xce, 0x2a,
	0xd1, 0xd4, 0x59, 0x2e, 0x30, 0xd1, 0x11, 0xdf, 0x21, 0xc3, 0xcf, 0x1d, 0xce, 0x3e, 0x85, 0x71,
	0xae, 0x65, 0x2a, 0x74, 0xb6, 0x50, 0xa2, 0xc0, 0xfc, 0x86, 0x1c, 0x72, 0x2d, 0x9f, 0x05, 0x84,
	0x45, 0xb0, 0x75, 0x29, 0x6c, 0x23, 0x8d, 0xc6, 0x64, 0x46, 0xbc, 0xbb, 0xb2, 0x17, 0xb0, 0x23,
	0x0b, 0xa1, 0x9d, 0x74, 0xaf, 0xd3, 0xdc, 0x68, 0x27, 0xae, 0x5d, 0xd4, 0xc7, 0x7c, 0x67, 0xab,
	0xf9, 0x12, 0xf1, 0x49, 0xe0, 0xf1, 0x6d, 0xf9, 0x2e, 0x10, 0xbf, 0x19, 0xc0, 0x00, 0x8b, 0xc0,
	0xbe, 0x87, 0x31, 0x96, 0x21, 0x95, 0x55, 0x56, 0x0a, 0xaa, 0xf0, 0x6a, 0xc5, 0x4e, 0xbd, 0x95,
	0x03, 0x52, 0xf1, 0xcc, 0x7e, 0x80, 0x1d, 0x12, 0x6a, 0xe9, 0x48, 0x7d, 0x6b, 0xad, 0xfa, 0x76,
	0x50, 0x6b, 0xe9, 0x82, 0x87, 0x47, 0x30, 0xf1, 0x2f, 0x62, 0x8d, 0x4a, 0x6b, 0x63, 0x1d, 0x55,
	0xff, 0xee, 0xea, 0xd7, 0x32, 0xd6, 0xf1, 0x31, 0x51, 0xfd, 0x85, 0x9d, 0xc0, 0x9e, 0x2c, 0xb5,
	0xb1, 0x22, 0x95, 0x7a, 0x61, 0x5a, 0x5d, 0xa0, 0x83, 0x26, 0xea, 0xcf, 0x36, 0x3e, 0xec, 0x81,
	0x05, 0xc9, 0x69, 0x50, 0x78, 0xa8, 0x61, 0xa7, 0x70, 0x97, 0x1c, 0x99, 0xd6, 0xdd, 0xf4, 0x34,
	0x58, 0xe7, 0xe9, 0x4e, 0xd0, 0xfc, 0x42, 0x92, 0xe0, 0xea, 0x11, 0x4c, 0x6e, 0x26, 0x13, 0x6d,
	0xae, 0x7d, 0x1b, 0xf9, 0x5f, 0x16, 0xec, 0x5b, 0x80, 0xac, 0xa8, 0xa4, 0x0e, 0xba, 0xad, 0x75,
	0xba, 0x11, 0x12, 0x51, 0xf5, 0x18, 0xa6, 0xef, 0xe4, 0x1c, 0x0d, 0xd7, 0x09, 0x27, 0xe6, 0x46,
	0xb2, 0xec, 0x18, 0x86, 0x56, 0x34, 0xa6, 0xb5, 0xb9, 0x88, 0x46, 0x28, 0x7b, 0xb0, 0x22, 0xe3,
	0x44, 0xe0, 0xe2, 0x55, 0x2b, 0xad, 0xa8, 0x84, 0x76, 0x0d, 0x7f, 0x2b, 0x63, 0x1f, 0xc3, 0x28,
	0x7c, 0xfe, 0x56, 0x16, 0x11, 0xcc, 0x7a, 0x0f, 0x37, 0xf8, 0x10, 0x81, 0xdf, 0x64, 0xc1, 0xbe,
	0x83, 0x91, 0x32, 0x65, 0xaa, 0xc4, 0xa5, 0x50, 0xd1, 0x18, 0x03, 0xec, 0xaf, 0x04, 0x38, 0x33,
	0xe5, 0x99, 0x27, 0xf0, 0xa1, 0xa2, 0x13, 0x7b, 0x0c, 0xfb, 0x85, 0x6c, 0xfc, 0x24, 0xa4, 0xe2,
	0xda, 0x09, 0xab, 0x33, 0x95, 0xd6, 0xd6, 0x2c, 0xa5, 0x12, 0x4d, 0x34, 0xc1, 0x61, 0xb9, 0x47,
	0x84, 0x67, 0x64, 0x3f, 0x27, 0x33, 0xfb, 0x1c, 0xa6, 0x21, 0xa1, 0x6e, 0x7e, 0xa6, 0x38, 0x3f,
	0x13, 0x04, 0x7f, 0x0f, 0x58, 0x7c, 0x02, 0x83, 0xd0, 0x7b, 0xf7, 0x01, 0xb0, 0x65, 0x71, 0x66,
	0x69, 0x5c, 0x47, 0x88, 0xf8, 0x61, 0xf5, 0x73, 0x5a, 0xb7, 0xca, 0xf7, 0xa5, 0x92, 0x79, 0xd8,
	0x23, 0x23, 0x0e, 0x1e, 0x3a, 0x47, 0x24, 0x3e, 0x80, 0x3e, 0x56, 0x92, 0x41, 0x1f, 0x8b, 0xef,
	0x3d, 0x4c, 0x39, 0x9e, 0xe3, 0xbf, 0x7a, 0xb0, 0xf7, 0x7f, 0xd5, 0xf3, 0x5e, 0xad, 0x78, 0xd5,
	0x8a, 0xc6, 0xa5, 0x79, 0xdd, 0x52, 0x54, 0x20, 0xe8, 0x49, 0xdd, 0xb2, 0x07, 0x70, 0xbb, 0x23,
	0x54, 0xa2, 0x32, 0xb6, 0x8b, 0x3c, 0x25, 0xf4, 0x27, 0x04, 0x7d, 0xed, 0x95, 0xac, 0x64, 0xf0,
	0x12, 0xd6, 0xc4, 0x10, 0x01, 0xef, 0xe3, 0x33, 0x98, 0x04, 0x23, 0x79, 0xe8, 0xa3, 0x7d, 0x8c,
	0x58, 0xd0, 0xc7, 0x6f, 0x7a, 0xb0, 0xfd, 0xde, 0x8a, 0xf0, 0x32, 0x67, 0xdb, 0xc6, 0xa5, 0x85,
	0xa9, 0x32, 0xa9, 0x29, 0xb9, 0x31, 0x62, 0x4f, 0x11, 0x62, 0x5f, 0xc0, 0x6e, 0xa0, 0x64, 0x3a,
	0x7f, 0x69, 0x6c, 0x93, 0xd6, 0xa2, 0xa2, 0x04, 0xb7, 0xd1, 0x70, 0x1c, 0xf0, 0x73, 0x51, 0xb1,
	0xe7, 0xb0, 0x2b, 0x9b, 0xa6, 0xcd, 0x74, 0x2e, 0x52, 0x25, 0x97, 0xc2, 0xc9, 0x4a, 0xd0, 0x80,
	0xef, 0x27, 0x61, 0xef, 0x27, 0xdd, 0xde, 0x4f, 0x9e, 0xd2, 0xde, 0xe7, 0x3b, 0x9d, 0xe6, 0x8c,
	0x24, 0xec, 0x05, 0xec, 0xe5, 0xca, 0xe4, 0x17, 0x69, 0x73, 0x21, 0xae, 0xd2, 0x4c, 0x29, 0x73,
	0xe5, 0xed, 0xb4, 0xf9, 0xd6, 0xb8, 0x62, 0x28, 0xfb, 0xf5, 0x42, 0x5c, 0x1d, 0x77, 0xa2, 0x78,
	0x06, 0xc3, 0xae, 0xe9, 0xd8, 0x1e, 0x0c, 0x42, 0x7b, 0x86, 0x17, 0x0d, 0x97, 0xf8, 0xef, 0x1e,
	0x6c, 0xd1, 0xb2, 0xf7, 0x9f, 0xb6, 0xf5, 0xcd, 0x1d, 0x08, 0x78, 0xc6, 0xfd, 0xad, 0xe4, 0xdb,
	0x16, 0xa3, 0xbe, 0xc8, 0x95, 0xa4, 0x06, 0x63, 0x47, 0x30, 0x58, 0xaa, 0xac, 0x6c, 0xa2, 0x0d,
	0x5c, 0x20, 0xf7, 0x3f, 0xf4, 0x2b, 0x49, 0x9e, 0xab, 0xac, 0xe4, 0x81, 0x7b, 0xf0, 0x35, 0xf4,
	0xfd, 0xd5, 0x47, 0xbc, 0xd1, 0x8e, 0x78, 0xf6, 0x79, 0x5e, 0x66, 0xaa, 0x15, 0x14, 0x2b, 0x5c,
	0x7e, 0x3c, 0xfa, 0xe3, 0x9b, 0x52, 0xba, 0x97, 0xed, 0x22, 0xc9, 0x4d, 0x35, 0xa7, 0x18, 0xdd,
	0xf3, 0x70, 0x4e, 0xbb, 0x52, 0x09, 0x3b, 0x2f, 0x85, 0xa6, 0xdf, 0xf0, 0x62, 0x13, 0xab, 0x74,
	0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x5f, 0xe9, 0xac, 0x9e, 0x07, 0x00, 0x00,
}
