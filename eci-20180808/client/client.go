// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateContainerGroupRequest struct {
	DnsConfig                     *CreateContainerGroupRequestDnsConfig                 `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	HostSecurityContext           *CreateContainerGroupRequestHostSecurityContext       `json:"HostSecurityContext,omitempty" xml:"HostSecurityContext,omitempty" type:"Struct"`
	SecurityContext               *CreateContainerGroupRequestSecurityContext           `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	AcrRegistryInfo               []*CreateContainerGroupRequestAcrRegistryInfo         `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	ActiveDeadlineSeconds         *int64                                                `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	Arn                           []*CreateContainerGroupRequestArn                     `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	AutoCreateEip                 *bool                                                 `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	AutoMatchImageCache           *bool                                                 `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ClientToken                   *string                                               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Container                     []*CreateContainerGroupRequestContainer               `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	ContainerGroupName            *string                                               `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	CorePattern                   *string                                               `json:"CorePattern,omitempty" xml:"CorePattern,omitempty"`
	Cpu                           *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CpuOptionsCore                *int32                                                `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	CpuOptionsNuma                *string                                               `json:"CpuOptionsNuma,omitempty" xml:"CpuOptionsNuma,omitempty"`
	CpuOptionsThreadsPerCore      *int32                                                `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	DnsPolicy                     *string                                               `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	EgressBandwidth               *int64                                                `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	EipBandwidth                  *int32                                                `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	EipCommonBandwidthPackage     *string                                               `json:"EipCommonBandwidthPackage,omitempty" xml:"EipCommonBandwidthPackage,omitempty"`
	EipISP                        *string                                               `json:"EipISP,omitempty" xml:"EipISP,omitempty"`
	EipInstanceId                 *string                                               `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	EphemeralStorage              *int32                                                `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	HostAliase                    []*CreateContainerGroupRequestHostAliase              `json:"HostAliase,omitempty" xml:"HostAliase,omitempty" type:"Repeated"`
	HostName                      *string                                               `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageAccelerateMode           *string                                               `json:"ImageAccelerateMode,omitempty" xml:"ImageAccelerateMode,omitempty"`
	ImageRegistryCredential       []*CreateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	ImageSnapshotId               *string                                               `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	IngressBandwidth              *int64                                                `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	InitContainer                 []*CreateContainerGroupRequestInitContainer           `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	InsecureRegistry              *string                                               `json:"InsecureRegistry,omitempty" xml:"InsecureRegistry,omitempty"`
	InstanceType                  *string                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Ipv6AddressCount              *int32                                                `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	Ipv6GatewayBandwidth          *string                                               `json:"Ipv6GatewayBandwidth,omitempty" xml:"Ipv6GatewayBandwidth,omitempty"`
	Ipv6GatewayBandwidthEnable    *bool                                                 `json:"Ipv6GatewayBandwidthEnable,omitempty" xml:"Ipv6GatewayBandwidthEnable,omitempty"`
	Memory                        *float32                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NtpServer                     []*string                                             `json:"NtpServer,omitempty" xml:"NtpServer,omitempty" type:"Repeated"`
	OwnerAccount                  *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                       *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlainHttpRegistry             *string                                               `json:"PlainHttpRegistry,omitempty" xml:"PlainHttpRegistry,omitempty"`
	ProductOnEciMode              *string                                               `json:"ProductOnEciMode,omitempty" xml:"ProductOnEciMode,omitempty"`
	RamRoleName                   *string                                               `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RegionId                      *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId               *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount          *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId               *int64                                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestartPolicy                 *string                                               `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ScheduleStrategy              *string                                               `json:"ScheduleStrategy,omitempty" xml:"ScheduleStrategy,omitempty"`
	SecondaryENIPolicy            *string                                               `json:"SecondaryENIPolicy,omitempty" xml:"SecondaryENIPolicy,omitempty"`
	SecurityGroupId               *string                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	ShareProcessNamespace         *bool                                                 `json:"ShareProcessNamespace,omitempty" xml:"ShareProcessNamespace,omitempty"`
	SlsEnable                     *bool                                                 `json:"SlsEnable,omitempty" xml:"SlsEnable,omitempty"`
	SpotDuration                  *int64                                                `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotPriceLimit                *float32                                              `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy                  *string                                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	StrictSpot                    *bool                                                 `json:"StrictSpot,omitempty" xml:"StrictSpot,omitempty"`
	Tag                           []*CreateContainerGroupRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TenantSecurityGroupId         *string                                               `json:"TenantSecurityGroupId,omitempty" xml:"TenantSecurityGroupId,omitempty"`
	TenantVSwitchId               *string                                               `json:"TenantVSwitchId,omitempty" xml:"TenantVSwitchId,omitempty"`
	TerminationGracePeriodSeconds *int64                                                `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	VSwitchId                     *string                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Volume                        []*CreateContainerGroupRequestVolume                  `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
	ZoneId                        *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequest) SetDnsConfig(v *CreateContainerGroupRequestDnsConfig) *CreateContainerGroupRequest {
	s.DnsConfig = v
	return s
}

func (s *CreateContainerGroupRequest) SetHostSecurityContext(v *CreateContainerGroupRequestHostSecurityContext) *CreateContainerGroupRequest {
	s.HostSecurityContext = v
	return s
}

func (s *CreateContainerGroupRequest) SetSecurityContext(v *CreateContainerGroupRequestSecurityContext) *CreateContainerGroupRequest {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequest) SetAcrRegistryInfo(v []*CreateContainerGroupRequestAcrRegistryInfo) *CreateContainerGroupRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *CreateContainerGroupRequest) SetActiveDeadlineSeconds(v int64) *CreateContainerGroupRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *CreateContainerGroupRequest) SetArn(v []*CreateContainerGroupRequestArn) *CreateContainerGroupRequest {
	s.Arn = v
	return s
}

func (s *CreateContainerGroupRequest) SetAutoCreateEip(v bool) *CreateContainerGroupRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *CreateContainerGroupRequest) SetAutoMatchImageCache(v bool) *CreateContainerGroupRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateContainerGroupRequest) SetClientToken(v string) *CreateContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateContainerGroupRequest) SetContainer(v []*CreateContainerGroupRequestContainer) *CreateContainerGroupRequest {
	s.Container = v
	return s
}

func (s *CreateContainerGroupRequest) SetContainerGroupName(v string) *CreateContainerGroupRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCorePattern(v string) *CreateContainerGroupRequest {
	s.CorePattern = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpu(v float32) *CreateContainerGroupRequest {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsCore(v int32) *CreateContainerGroupRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsNuma(v string) *CreateContainerGroupRequest {
	s.CpuOptionsNuma = &v
	return s
}

func (s *CreateContainerGroupRequest) SetCpuOptionsThreadsPerCore(v int32) *CreateContainerGroupRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *CreateContainerGroupRequest) SetDnsPolicy(v string) *CreateContainerGroupRequest {
	s.DnsPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEgressBandwidth(v int64) *CreateContainerGroupRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipBandwidth(v int32) *CreateContainerGroupRequest {
	s.EipBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipCommonBandwidthPackage(v string) *CreateContainerGroupRequest {
	s.EipCommonBandwidthPackage = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipISP(v string) *CreateContainerGroupRequest {
	s.EipISP = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEipInstanceId(v string) *CreateContainerGroupRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetEphemeralStorage(v int32) *CreateContainerGroupRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *CreateContainerGroupRequest) SetHostAliase(v []*CreateContainerGroupRequestHostAliase) *CreateContainerGroupRequest {
	s.HostAliase = v
	return s
}

func (s *CreateContainerGroupRequest) SetHostName(v string) *CreateContainerGroupRequest {
	s.HostName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetImageAccelerateMode(v string) *CreateContainerGroupRequest {
	s.ImageAccelerateMode = &v
	return s
}

func (s *CreateContainerGroupRequest) SetImageRegistryCredential(v []*CreateContainerGroupRequestImageRegistryCredential) *CreateContainerGroupRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *CreateContainerGroupRequest) SetImageSnapshotId(v string) *CreateContainerGroupRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIngressBandwidth(v int64) *CreateContainerGroupRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetInitContainer(v []*CreateContainerGroupRequestInitContainer) *CreateContainerGroupRequest {
	s.InitContainer = v
	return s
}

func (s *CreateContainerGroupRequest) SetInsecureRegistry(v string) *CreateContainerGroupRequest {
	s.InsecureRegistry = &v
	return s
}

func (s *CreateContainerGroupRequest) SetInstanceType(v string) *CreateContainerGroupRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6AddressCount(v int32) *CreateContainerGroupRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6GatewayBandwidth(v string) *CreateContainerGroupRequest {
	s.Ipv6GatewayBandwidth = &v
	return s
}

func (s *CreateContainerGroupRequest) SetIpv6GatewayBandwidthEnable(v bool) *CreateContainerGroupRequest {
	s.Ipv6GatewayBandwidthEnable = &v
	return s
}

func (s *CreateContainerGroupRequest) SetMemory(v float32) *CreateContainerGroupRequest {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequest) SetNtpServer(v []*string) *CreateContainerGroupRequest {
	s.NtpServer = v
	return s
}

func (s *CreateContainerGroupRequest) SetOwnerAccount(v string) *CreateContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetOwnerId(v int64) *CreateContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetPlainHttpRegistry(v string) *CreateContainerGroupRequest {
	s.PlainHttpRegistry = &v
	return s
}

func (s *CreateContainerGroupRequest) SetProductOnEciMode(v string) *CreateContainerGroupRequest {
	s.ProductOnEciMode = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRamRoleName(v string) *CreateContainerGroupRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRegionId(v string) *CreateContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceGroupId(v string) *CreateContainerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceOwnerAccount(v string) *CreateContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateContainerGroupRequest) SetResourceOwnerId(v int64) *CreateContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetRestartPolicy(v string) *CreateContainerGroupRequest {
	s.RestartPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetScheduleStrategy(v string) *CreateContainerGroupRequest {
	s.ScheduleStrategy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSecondaryENIPolicy(v string) *CreateContainerGroupRequest {
	s.SecondaryENIPolicy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSecurityGroupId(v string) *CreateContainerGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetShareProcessNamespace(v bool) *CreateContainerGroupRequest {
	s.ShareProcessNamespace = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSlsEnable(v bool) *CreateContainerGroupRequest {
	s.SlsEnable = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotDuration(v int64) *CreateContainerGroupRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotPriceLimit(v float32) *CreateContainerGroupRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateContainerGroupRequest) SetSpotStrategy(v string) *CreateContainerGroupRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateContainerGroupRequest) SetStrictSpot(v bool) *CreateContainerGroupRequest {
	s.StrictSpot = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTag(v []*CreateContainerGroupRequestTag) *CreateContainerGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateContainerGroupRequest) SetTenantSecurityGroupId(v string) *CreateContainerGroupRequest {
	s.TenantSecurityGroupId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTenantVSwitchId(v string) *CreateContainerGroupRequest {
	s.TenantVSwitchId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetTerminationGracePeriodSeconds(v int64) *CreateContainerGroupRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequest) SetVSwitchId(v string) *CreateContainerGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateContainerGroupRequest) SetVolume(v []*CreateContainerGroupRequestVolume) *CreateContainerGroupRequest {
	s.Volume = v
	return s
}

func (s *CreateContainerGroupRequest) SetZoneId(v string) *CreateContainerGroupRequest {
	s.ZoneId = &v
	return s
}

type CreateContainerGroupRequestDnsConfig struct {
	NameServer []*string                                     `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	Option     []*CreateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	Search     []*string                                     `json:"Search,omitempty" xml:"Search,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestDnsConfig) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestDnsConfig) SetNameServer(v []*string) *CreateContainerGroupRequestDnsConfig {
	s.NameServer = v
	return s
}

func (s *CreateContainerGroupRequestDnsConfig) SetOption(v []*CreateContainerGroupRequestDnsConfigOption) *CreateContainerGroupRequestDnsConfig {
	s.Option = v
	return s
}

func (s *CreateContainerGroupRequestDnsConfig) SetSearch(v []*string) *CreateContainerGroupRequestDnsConfig {
	s.Search = v
	return s
}

type CreateContainerGroupRequestDnsConfigOption struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestDnsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestDnsConfigOption) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestDnsConfigOption) SetName(v string) *CreateContainerGroupRequestDnsConfigOption {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestDnsConfigOption) SetValue(v string) *CreateContainerGroupRequestDnsConfigOption {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestHostSecurityContext struct {
	Sysctl []*CreateContainerGroupRequestHostSecurityContextSysctl `json:"Sysctl,omitempty" xml:"Sysctl,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestHostSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostSecurityContext) SetSysctl(v []*CreateContainerGroupRequestHostSecurityContextSysctl) *CreateContainerGroupRequestHostSecurityContext {
	s.Sysctl = v
	return s
}

type CreateContainerGroupRequestHostSecurityContextSysctl struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestHostSecurityContextSysctl) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostSecurityContextSysctl) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostSecurityContextSysctl) SetName(v string) *CreateContainerGroupRequestHostSecurityContextSysctl {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestHostSecurityContextSysctl) SetValue(v string) *CreateContainerGroupRequestHostSecurityContextSysctl {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestSecurityContext struct {
	Sysctl []*CreateContainerGroupRequestSecurityContextSysctl `json:"Sysctl,omitempty" xml:"Sysctl,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestSecurityContext) SetSysctl(v []*CreateContainerGroupRequestSecurityContextSysctl) *CreateContainerGroupRequestSecurityContext {
	s.Sysctl = v
	return s
}

type CreateContainerGroupRequestSecurityContextSysctl struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestSecurityContextSysctl) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestSecurityContextSysctl) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestSecurityContextSysctl) SetName(v string) *CreateContainerGroupRequestSecurityContextSysctl {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestSecurityContextSysctl) SetValue(v string) *CreateContainerGroupRequestSecurityContextSysctl {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestAcrRegistryInfo struct {
	Domain       []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateContainerGroupRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetDomain(v []*string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetInstanceId(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetInstanceName(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *CreateContainerGroupRequestAcrRegistryInfo) SetRegionId(v string) *CreateContainerGroupRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CreateContainerGroupRequestArn struct {
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	RoleArn       *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleType      *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CreateContainerGroupRequestArn) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestArn) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestArn) SetAssumeRoleFor(v string) *CreateContainerGroupRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateContainerGroupRequestArn) SetRoleArn(v string) *CreateContainerGroupRequestArn {
	s.RoleArn = &v
	return s
}

func (s *CreateContainerGroupRequestArn) SetRoleType(v string) *CreateContainerGroupRequestArn {
	s.RoleType = &v
	return s
}

type CreateContainerGroupRequestContainer struct {
	LivenessProbe                              *CreateContainerGroupRequestContainerLivenessProbe                                `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	ReadinessProbe                             *CreateContainerGroupRequestContainerReadinessProbe                               `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext                            *CreateContainerGroupRequestContainerSecurityContext                              `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Arg                                        []*string                                                                         `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command                                    []*string                                                                         `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Cpu                                        *float32                                                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVar                             []*CreateContainerGroupRequestContainerEnvironmentVar                             `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	Gpu                                        *int32                                                                            `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                                      *string                                                                           `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy                            *string                                                                           `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	LifecyclePostStartHandlerExec              []*string                                                                         `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerHttpGetHost       *string                                                                           `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	LifecyclePostStartHandlerHttpGetHttpHeader []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader `json:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerHttpGetPath       *string                                                                           `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	LifecyclePostStartHandlerHttpGetPort       *int32                                                                            `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	LifecyclePostStartHandlerHttpGetScheme     *string                                                                           `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	LifecyclePostStartHandlerTcpSocketHost     *string                                                                           `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	LifecyclePostStartHandlerTcpSocketPort     *int32                                                                            `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	LifecyclePreStopHandlerExec                []*string                                                                         `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetHost         *string                                                                           `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	LifecyclePreStopHandlerHttpGetHttpHeader   []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader   `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetPath         *string                                                                           `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	LifecyclePreStopHandlerHttpGetPort         *int32                                                                            `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	LifecyclePreStopHandlerHttpGetScheme       *string                                                                           `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	LifecyclePreStopHandlerTcpSocketHost       *string                                                                           `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	LifecyclePreStopHandlerTcpSocketPort       *int32                                                                            `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	Memory                                     *float32                                                                          `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                                       *string                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Port                                       []*CreateContainerGroupRequestContainerPort                                       `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	Stdin                                      *bool                                                                             `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce                                  *bool                                                                             `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	TerminationMessagePath                     *string                                                                           `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	TerminationMessagePolicy                   *string                                                                           `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	Tty                                        *bool                                                                             `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMount                                []*CreateContainerGroupRequestContainerVolumeMount                                `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	WorkingDir                                 *string                                                                           `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateContainerGroupRequestContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainer) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainer) SetLivenessProbe(v *CreateContainerGroupRequestContainerLivenessProbe) *CreateContainerGroupRequestContainer {
	s.LivenessProbe = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetReadinessProbe(v *CreateContainerGroupRequestContainerReadinessProbe) *CreateContainerGroupRequestContainer {
	s.ReadinessProbe = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetSecurityContext(v *CreateContainerGroupRequestContainerSecurityContext) *CreateContainerGroupRequestContainer {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetArg(v []*string) *CreateContainerGroupRequestContainer {
	s.Arg = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetCommand(v []*string) *CreateContainerGroupRequestContainer {
	s.Command = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetCpu(v float32) *CreateContainerGroupRequestContainer {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetEnvironmentVar(v []*CreateContainerGroupRequestContainerEnvironmentVar) *CreateContainerGroupRequestContainer {
	s.EnvironmentVar = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetGpu(v int32) *CreateContainerGroupRequestContainer {
	s.Gpu = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetImage(v string) *CreateContainerGroupRequestContainer {
	s.Image = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetImagePullPolicy(v string) *CreateContainerGroupRequestContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerExec(v []*string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerExec = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHttpHeader(v []*CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHttpHeader = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPath(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetScheme(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerExec(v []*string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerExec = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHttpHeader(v []*CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHttpHeader = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPath(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetScheme(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketHost(v string) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *CreateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetMemory(v float32) *CreateContainerGroupRequestContainer {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetName(v string) *CreateContainerGroupRequestContainer {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetPort(v []*CreateContainerGroupRequestContainerPort) *CreateContainerGroupRequestContainer {
	s.Port = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetStdin(v bool) *CreateContainerGroupRequestContainer {
	s.Stdin = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetStdinOnce(v bool) *CreateContainerGroupRequestContainer {
	s.StdinOnce = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTerminationMessagePath(v string) *CreateContainerGroupRequestContainer {
	s.TerminationMessagePath = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTerminationMessagePolicy(v string) *CreateContainerGroupRequestContainer {
	s.TerminationMessagePolicy = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetTty(v bool) *CreateContainerGroupRequestContainer {
	s.Tty = &v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetVolumeMount(v []*CreateContainerGroupRequestContainerVolumeMount) *CreateContainerGroupRequestContainer {
	s.VolumeMount = v
	return s
}

func (s *CreateContainerGroupRequestContainer) SetWorkingDir(v string) *CreateContainerGroupRequestContainer {
	s.WorkingDir = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbe struct {
	Exec                *CreateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                      `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbe) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetExec(v *CreateContainerGroupRequestContainerLivenessProbeExec) *CreateContainerGroupRequestContainerLivenessProbe {
	s.Exec = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetFailureThreshold(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetHttpGet(v *CreateContainerGroupRequestContainerLivenessProbeHttpGet) *CreateContainerGroupRequestContainerLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetInitialDelaySeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetPeriodSeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetSuccessThreshold(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetTcpSocket(v *CreateContainerGroupRequestContainerLivenessProbeTcpSocket) *CreateContainerGroupRequestContainerLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbe) SetTimeoutSeconds(v int32) *CreateContainerGroupRequestContainerLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeExec) SetCommand(v []*string) *CreateContainerGroupRequestContainerLivenessProbeExec {
	s.Command = v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetPath(v string) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetPort(v int32) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLivenessProbeHttpGet) SetScheme(v string) *CreateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type CreateContainerGroupRequestContainerLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLivenessProbeTcpSocket) SetPort(v int32) *CreateContainerGroupRequestContainerLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbe struct {
	Exec                *CreateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                       `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbe) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetExec(v *CreateContainerGroupRequestContainerReadinessProbeExec) *CreateContainerGroupRequestContainerReadinessProbe {
	s.Exec = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetFailureThreshold(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetHttpGet(v *CreateContainerGroupRequestContainerReadinessProbeHttpGet) *CreateContainerGroupRequestContainerReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetInitialDelaySeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetPeriodSeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetSuccessThreshold(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetTcpSocket(v *CreateContainerGroupRequestContainerReadinessProbeTcpSocket) *CreateContainerGroupRequestContainerReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbe) SetTimeoutSeconds(v int32) *CreateContainerGroupRequestContainerReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeExec) SetCommand(v []*string) *CreateContainerGroupRequestContainerReadinessProbeExec {
	s.Command = v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetPath(v string) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetPort(v int32) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerReadinessProbeHttpGet) SetScheme(v string) *CreateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type CreateContainerGroupRequestContainerReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateContainerGroupRequestContainerReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerReadinessProbeTcpSocket) SetPort(v int32) *CreateContainerGroupRequestContainerReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type CreateContainerGroupRequestContainerSecurityContext struct {
	Capability             *CreateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                          `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                         `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateContainerGroupRequestContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetCapability(v *CreateContainerGroupRequestContainerSecurityContextCapability) *CreateContainerGroupRequestContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateContainerGroupRequestContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateContainerGroupRequestContainerSecurityContext) SetRunAsUser(v int64) *CreateContainerGroupRequestContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateContainerGroupRequestContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerSecurityContextCapability) SetAdd(v []*string) *CreateContainerGroupRequestContainerSecurityContextCapability {
	s.Add = v
	return s
}

type CreateContainerGroupRequestContainerEnvironmentVar struct {
	FieldRef *CreateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetFieldRef(v *CreateContainerGroupRequestContainerEnvironmentVarFieldRef) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetKey(v string) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestContainerEnvironmentVar) SetValue(v string) *CreateContainerGroupRequestContainerEnvironmentVar {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s CreateContainerGroupRequestContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerEnvironmentVarFieldRef) SetFieldPath(v string) *CreateContainerGroupRequestContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) SetName(v string) *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader) SetValue(v string) *CreateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetName(v string) *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetValue(v string) *CreateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestContainerPort struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateContainerGroupRequestContainerPort) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerPort) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerPort) SetPort(v int32) *CreateContainerGroupRequestContainerPort {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestContainerPort) SetProtocol(v string) *CreateContainerGroupRequestContainerPort {
	s.Protocol = &v
	return s
}

type CreateContainerGroupRequestContainerVolumeMount struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateContainerGroupRequestContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetMountPath(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetMountPropagation(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetName(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetReadOnly(v bool) *CreateContainerGroupRequestContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestContainerVolumeMount) SetSubPath(v string) *CreateContainerGroupRequestContainerVolumeMount {
	s.SubPath = &v
	return s
}

type CreateContainerGroupRequestHostAliase struct {
	Hostname []*string `json:"Hostname,omitempty" xml:"Hostname,omitempty" type:"Repeated"`
	Ip       *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s CreateContainerGroupRequestHostAliase) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestHostAliase) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestHostAliase) SetHostname(v []*string) *CreateContainerGroupRequestHostAliase {
	s.Hostname = v
	return s
}

func (s *CreateContainerGroupRequestHostAliase) SetIp(v string) *CreateContainerGroupRequestHostAliase {
	s.Ip = &v
	return s
}

type CreateContainerGroupRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateContainerGroupRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetPassword(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetServer(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *CreateContainerGroupRequestImageRegistryCredential) SetUserName(v string) *CreateContainerGroupRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type CreateContainerGroupRequestInitContainer struct {
	SecurityContext          *CreateContainerGroupRequestInitContainerSecurityContext  `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Arg                      []*string                                                 `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command                  []*string                                                 `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Cpu                      *float32                                                  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVar           []*CreateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	Gpu                      *int32                                                    `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                    *string                                                   `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy          *string                                                   `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	Memory                   *float32                                                  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                     *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Port                     []*CreateContainerGroupRequestInitContainerPort           `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	TerminationMessagePath   *string                                                   `json:"TerminationMessagePath,omitempty" xml:"TerminationMessagePath,omitempty"`
	TerminationMessagePolicy *string                                                   `json:"TerminationMessagePolicy,omitempty" xml:"TerminationMessagePolicy,omitempty"`
	VolumeMount              []*CreateContainerGroupRequestInitContainerVolumeMount    `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	WorkingDir               *string                                                   `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateContainerGroupRequestInitContainer) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainer) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainer) SetSecurityContext(v *CreateContainerGroupRequestInitContainerSecurityContext) *CreateContainerGroupRequestInitContainer {
	s.SecurityContext = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetArg(v []*string) *CreateContainerGroupRequestInitContainer {
	s.Arg = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetCommand(v []*string) *CreateContainerGroupRequestInitContainer {
	s.Command = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetCpu(v float32) *CreateContainerGroupRequestInitContainer {
	s.Cpu = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetEnvironmentVar(v []*CreateContainerGroupRequestInitContainerEnvironmentVar) *CreateContainerGroupRequestInitContainer {
	s.EnvironmentVar = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetGpu(v int32) *CreateContainerGroupRequestInitContainer {
	s.Gpu = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetImage(v string) *CreateContainerGroupRequestInitContainer {
	s.Image = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetImagePullPolicy(v string) *CreateContainerGroupRequestInitContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetMemory(v float32) *CreateContainerGroupRequestInitContainer {
	s.Memory = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetName(v string) *CreateContainerGroupRequestInitContainer {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetPort(v []*CreateContainerGroupRequestInitContainerPort) *CreateContainerGroupRequestInitContainer {
	s.Port = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetTerminationMessagePath(v string) *CreateContainerGroupRequestInitContainer {
	s.TerminationMessagePath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetTerminationMessagePolicy(v string) *CreateContainerGroupRequestInitContainer {
	s.TerminationMessagePolicy = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetVolumeMount(v []*CreateContainerGroupRequestInitContainerVolumeMount) *CreateContainerGroupRequestInitContainer {
	s.VolumeMount = v
	return s
}

func (s *CreateContainerGroupRequestInitContainer) SetWorkingDir(v string) *CreateContainerGroupRequestInitContainer {
	s.WorkingDir = &v
	return s
}

type CreateContainerGroupRequestInitContainerSecurityContext struct {
	Capability             *CreateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                              `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                             `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetCapability(v *CreateContainerGroupRequestInitContainerSecurityContextCapability) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerSecurityContext) SetRunAsUser(v int64) *CreateContainerGroupRequestInitContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type CreateContainerGroupRequestInitContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateContainerGroupRequestInitContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerSecurityContextCapability) SetAdd(v []*string) *CreateContainerGroupRequestInitContainerSecurityContextCapability {
	s.Add = v
	return s
}

type CreateContainerGroupRequestInitContainerEnvironmentVar struct {
	FieldRef *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                         `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                         `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetFieldRef(v *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetKey(v string) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVar) SetValue(v string) *CreateContainerGroupRequestInitContainerEnvironmentVar {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef) SetFieldPath(v string) *CreateContainerGroupRequestInitContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type CreateContainerGroupRequestInitContainerPort struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerPort) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerPort) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerPort) SetPort(v int32) *CreateContainerGroupRequestInitContainerPort {
	s.Port = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerPort) SetProtocol(v string) *CreateContainerGroupRequestInitContainerPort {
	s.Protocol = &v
	return s
}

type CreateContainerGroupRequestInitContainerVolumeMount struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateContainerGroupRequestInitContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestInitContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetMountPath(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetMountPropagation(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetName(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetReadOnly(v bool) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestInitContainerVolumeMount) SetSubPath(v string) *CreateContainerGroupRequestInitContainerVolumeMount {
	s.SubPath = &v
	return s
}

type CreateContainerGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateContainerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestTag) SetKey(v string) *CreateContainerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateContainerGroupRequestTag) SetValue(v string) *CreateContainerGroupRequestTag {
	s.Value = &v
	return s
}

type CreateContainerGroupRequestVolume struct {
	ConfigFileVolume *CreateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" require:"true" type:"Struct"`
	DiskVolume       *CreateContainerGroupRequestVolumeDiskVolume       `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume   *CreateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume       *CreateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume   *CreateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume        *CreateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	Name             *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContainerGroupRequestVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolume) SetConfigFileVolume(v *CreateContainerGroupRequestVolumeConfigFileVolume) *CreateContainerGroupRequestVolume {
	s.ConfigFileVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetDiskVolume(v *CreateContainerGroupRequestVolumeDiskVolume) *CreateContainerGroupRequestVolume {
	s.DiskVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetEmptyDirVolume(v *CreateContainerGroupRequestVolumeEmptyDirVolume) *CreateContainerGroupRequestVolume {
	s.EmptyDirVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetFlexVolume(v *CreateContainerGroupRequestVolumeFlexVolume) *CreateContainerGroupRequestVolume {
	s.FlexVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetHostPathVolume(v *CreateContainerGroupRequestVolumeHostPathVolume) *CreateContainerGroupRequestVolume {
	s.HostPathVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetNFSVolume(v *CreateContainerGroupRequestVolumeNFSVolume) *CreateContainerGroupRequestVolume {
	s.NFSVolume = v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetName(v string) *CreateContainerGroupRequestVolume {
	s.Name = &v
	return s
}

func (s *CreateContainerGroupRequestVolume) SetType(v string) *CreateContainerGroupRequestVolume {
	s.Type = &v
	return s
}

type CreateContainerGroupRequestVolumeConfigFileVolume struct {
	ConfigFileToPath []*CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" xml:"ConfigFileToPath,omitempty" type:"Repeated"`
	DefaultMode      *int32                                                               `json:"DefaultMode,omitempty" xml:"DefaultMode,omitempty"`
}

func (s CreateContainerGroupRequestVolumeConfigFileVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeConfigFileVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolume) SetConfigFileToPath(v []*CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) *CreateContainerGroupRequestVolumeConfigFileVolume {
	s.ConfigFileToPath = v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolume) SetDefaultMode(v int32) *CreateContainerGroupRequestVolumeConfigFileVolume {
	s.DefaultMode = &v
	return s
}

type CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Mode    *int32  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetContent(v string) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetMode(v int32) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Mode = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetPath(v string) *CreateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

type CreateContainerGroupRequestVolumeDiskVolume struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
}

func (s CreateContainerGroupRequestVolumeDiskVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeDiskVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetDiskId(v string) *CreateContainerGroupRequestVolumeDiskVolume {
	s.DiskId = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetDiskSize(v int32) *CreateContainerGroupRequestVolumeDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeDiskVolume) SetFsType(v string) *CreateContainerGroupRequestVolumeDiskVolume {
	s.FsType = &v
	return s
}

type CreateContainerGroupRequestVolumeEmptyDirVolume struct {
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateContainerGroupRequestVolumeEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeEmptyDirVolume) SetMedium(v string) *CreateContainerGroupRequestVolumeEmptyDirVolume {
	s.Medium = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeEmptyDirVolume) SetSizeLimit(v string) *CreateContainerGroupRequestVolumeEmptyDirVolume {
	s.SizeLimit = &v
	return s
}

type CreateContainerGroupRequestVolumeFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s CreateContainerGroupRequestVolumeFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeFlexVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetDriver(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.Driver = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetFsType(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.FsType = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeFlexVolume) SetOptions(v string) *CreateContainerGroupRequestVolumeFlexVolume {
	s.Options = &v
	return s
}

type CreateContainerGroupRequestVolumeHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateContainerGroupRequestVolumeHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeHostPathVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeHostPathVolume) SetPath(v string) *CreateContainerGroupRequestVolumeHostPathVolume {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeHostPathVolume) SetType(v string) *CreateContainerGroupRequestVolumeHostPathVolume {
	s.Type = &v
	return s
}

type CreateContainerGroupRequestVolumeNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s CreateContainerGroupRequestVolumeNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupRequestVolumeNFSVolume) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetPath(v string) *CreateContainerGroupRequestVolumeNFSVolume {
	s.Path = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetReadOnly(v bool) *CreateContainerGroupRequestVolumeNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *CreateContainerGroupRequestVolumeNFSVolume) SetServer(v string) *CreateContainerGroupRequestVolumeNFSVolume {
	s.Server = &v
	return s
}

type CreateContainerGroupResponseBody struct {
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupResponseBody) SetContainerGroupId(v string) *CreateContainerGroupResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateContainerGroupResponseBody) SetRequestId(v string) *CreateContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateContainerGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateContainerGroupResponse) SetHeaders(v map[string]*string) *CreateContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateContainerGroupResponse) SetStatusCode(v int32) *CreateContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContainerGroupResponse) SetBody(v *CreateContainerGroupResponseBody) *CreateContainerGroupResponse {
	s.Body = v
	return s
}

type CreateImageCacheRequest struct {
	AcrRegistryInfo         []*CreateImageCacheRequestAcrRegistryInfo         `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	Annotations             *string                                           `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	AutoMatchImageCache     *bool                                             `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ClientToken             *string                                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EipInstanceId           *string                                           `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	EliminationStrategy     *string                                           `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	Flash                   *bool                                             `json:"Flash,omitempty" xml:"Flash,omitempty"`
	FlashCopyCount          *int32                                            `json:"FlashCopyCount,omitempty" xml:"FlashCopyCount,omitempty"`
	Image                   []*string                                         `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	ImageCacheName          *string                                           `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	ImageCacheSize          *int32                                            `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	ImageRegistryCredential []*CreateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	InsecureRegistry        *string                                           `json:"InsecureRegistry,omitempty" xml:"InsecureRegistry,omitempty"`
	OwnerAccount            *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlainHttpRegistry       *string                                           `json:"PlainHttpRegistry,omitempty" xml:"PlainHttpRegistry,omitempty"`
	RegionId                *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId         *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount    *string                                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RetentionDays           *int32                                            `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SecurityGroupId         *string                                           `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	StandardCopyCount       *int32                                            `json:"StandardCopyCount,omitempty" xml:"StandardCopyCount,omitempty"`
	Tag                     []*CreateImageCacheRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId               *string                                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                  *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequest) SetAcrRegistryInfo(v []*CreateImageCacheRequestAcrRegistryInfo) *CreateImageCacheRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *CreateImageCacheRequest) SetAnnotations(v string) *CreateImageCacheRequest {
	s.Annotations = &v
	return s
}

func (s *CreateImageCacheRequest) SetAutoMatchImageCache(v bool) *CreateImageCacheRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateImageCacheRequest) SetClientToken(v string) *CreateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheRequest) SetEipInstanceId(v string) *CreateImageCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateImageCacheRequest) SetEliminationStrategy(v string) *CreateImageCacheRequest {
	s.EliminationStrategy = &v
	return s
}

func (s *CreateImageCacheRequest) SetFlash(v bool) *CreateImageCacheRequest {
	s.Flash = &v
	return s
}

func (s *CreateImageCacheRequest) SetFlashCopyCount(v int32) *CreateImageCacheRequest {
	s.FlashCopyCount = &v
	return s
}

func (s *CreateImageCacheRequest) SetImage(v []*string) *CreateImageCacheRequest {
	s.Image = v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheName(v string) *CreateImageCacheRequest {
	s.ImageCacheName = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageCacheSize(v int32) *CreateImageCacheRequest {
	s.ImageCacheSize = &v
	return s
}

func (s *CreateImageCacheRequest) SetImageRegistryCredential(v []*CreateImageCacheRequestImageRegistryCredential) *CreateImageCacheRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *CreateImageCacheRequest) SetInsecureRegistry(v string) *CreateImageCacheRequest {
	s.InsecureRegistry = &v
	return s
}

func (s *CreateImageCacheRequest) SetOwnerAccount(v string) *CreateImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageCacheRequest) SetOwnerId(v int64) *CreateImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageCacheRequest) SetPlainHttpRegistry(v string) *CreateImageCacheRequest {
	s.PlainHttpRegistry = &v
	return s
}

func (s *CreateImageCacheRequest) SetRegionId(v string) *CreateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceGroupId(v string) *CreateImageCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceOwnerAccount(v string) *CreateImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageCacheRequest) SetResourceOwnerId(v int64) *CreateImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImageCacheRequest) SetRetentionDays(v int32) *CreateImageCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateImageCacheRequest) SetSecurityGroupId(v string) *CreateImageCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateImageCacheRequest) SetStandardCopyCount(v int32) *CreateImageCacheRequest {
	s.StandardCopyCount = &v
	return s
}

func (s *CreateImageCacheRequest) SetTag(v []*CreateImageCacheRequestTag) *CreateImageCacheRequest {
	s.Tag = v
	return s
}

func (s *CreateImageCacheRequest) SetVSwitchId(v string) *CreateImageCacheRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateImageCacheRequest) SetZoneId(v string) *CreateImageCacheRequest {
	s.ZoneId = &v
	return s
}

type CreateImageCacheRequestAcrRegistryInfo struct {
	Domain       []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageCacheRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetDomain(v []*string) *CreateImageCacheRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetInstanceId(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetInstanceName(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *CreateImageCacheRequestAcrRegistryInfo) SetRegionId(v string) *CreateImageCacheRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type CreateImageCacheRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateImageCacheRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetPassword(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetServer(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *CreateImageCacheRequestImageRegistryCredential) SetUserName(v string) *CreateImageCacheRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type CreateImageCacheRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestTag) SetKey(v string) *CreateImageCacheRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImageCacheRequestTag) SetValue(v string) *CreateImageCacheRequestTag {
	s.Value = &v
	return s
}

type CreateImageCacheResponseBody struct {
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ImageCacheId     *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponseBody) SetContainerGroupId(v string) *CreateImageCacheResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetImageCacheId(v string) *CreateImageCacheResponseBody {
	s.ImageCacheId = &v
	return s
}

func (s *CreateImageCacheResponseBody) SetRequestId(v string) *CreateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageCacheResponse) GoString() string {
	return s.String()
}

func (s *CreateImageCacheResponse) SetHeaders(v map[string]*string) *CreateImageCacheResponse {
	s.Headers = v
	return s
}

func (s *CreateImageCacheResponse) SetStatusCode(v int32) *CreateImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageCacheResponse) SetBody(v *CreateImageCacheResponseBody) *CreateImageCacheResponse {
	s.Body = v
	return s
}

type CreateInstanceOpsTaskRequest struct {
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OpsType              *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	OpsValue             *string `json:"OpsValue,omitempty" xml:"OpsValue,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateInstanceOpsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceOpsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceOpsTaskRequest) SetContainerGroupId(v string) *CreateInstanceOpsTaskRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOpsType(v string) *CreateInstanceOpsTaskRequest {
	s.OpsType = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOpsValue(v string) *CreateInstanceOpsTaskRequest {
	s.OpsValue = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOwnerAccount(v string) *CreateInstanceOpsTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetOwnerId(v int64) *CreateInstanceOpsTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetRegionId(v string) *CreateInstanceOpsTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetResourceOwnerAccount(v string) *CreateInstanceOpsTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceOpsTaskRequest) SetResourceOwnerId(v int64) *CreateInstanceOpsTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateInstanceOpsTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateInstanceOpsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceOpsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceOpsTaskResponseBody) SetRequestId(v string) *CreateInstanceOpsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceOpsTaskResponseBody) SetResult(v string) *CreateInstanceOpsTaskResponseBody {
	s.Result = &v
	return s
}

type CreateInstanceOpsTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceOpsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceOpsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceOpsTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceOpsTaskResponse) SetHeaders(v map[string]*string) *CreateInstanceOpsTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceOpsTaskResponse) SetStatusCode(v int32) *CreateInstanceOpsTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceOpsTaskResponse) SetBody(v *CreateInstanceOpsTaskResponseBody) *CreateInstanceOpsTaskResponse {
	s.Body = v
	return s
}

type CreateVirtualNodeRequest struct {
	ClientToken              *string                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EipInstanceId            *string                          `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	EnablePublicNetwork      *bool                            `json:"EnablePublicNetwork,omitempty" xml:"EnablePublicNetwork,omitempty"`
	KubeConfig               *string                          `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	OwnerAccount             *string                          `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                 *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId          *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount     *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RotateCertificateEnabled *bool                            `json:"RotateCertificateEnabled,omitempty" xml:"RotateCertificateEnabled,omitempty"`
	SecurityGroupId          *string                          `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag                      []*CreateVirtualNodeRequestTag   `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Taint                    []*CreateVirtualNodeRequestTaint `json:"Taint,omitempty" xml:"Taint,omitempty" type:"Repeated"`
	TlsBootstrapEnabled      *bool                            `json:"TlsBootstrapEnabled,omitempty" xml:"TlsBootstrapEnabled,omitempty"`
	VSwitchId                *string                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VirtualNodeName          *string                          `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	ZoneId                   *string                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVirtualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeRequest) SetClientToken(v string) *CreateVirtualNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetEipInstanceId(v string) *CreateVirtualNodeRequest {
	s.EipInstanceId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetEnablePublicNetwork(v bool) *CreateVirtualNodeRequest {
	s.EnablePublicNetwork = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetKubeConfig(v string) *CreateVirtualNodeRequest {
	s.KubeConfig = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetOwnerAccount(v string) *CreateVirtualNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetOwnerId(v int64) *CreateVirtualNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetRegionId(v string) *CreateVirtualNodeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetResourceGroupId(v string) *CreateVirtualNodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetResourceOwnerAccount(v string) *CreateVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetResourceOwnerId(v int64) *CreateVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetRotateCertificateEnabled(v bool) *CreateVirtualNodeRequest {
	s.RotateCertificateEnabled = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetSecurityGroupId(v string) *CreateVirtualNodeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetTag(v []*CreateVirtualNodeRequestTag) *CreateVirtualNodeRequest {
	s.Tag = v
	return s
}

func (s *CreateVirtualNodeRequest) SetTaint(v []*CreateVirtualNodeRequestTaint) *CreateVirtualNodeRequest {
	s.Taint = v
	return s
}

func (s *CreateVirtualNodeRequest) SetTlsBootstrapEnabled(v bool) *CreateVirtualNodeRequest {
	s.TlsBootstrapEnabled = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetVSwitchId(v string) *CreateVirtualNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetVirtualNodeName(v string) *CreateVirtualNodeRequest {
	s.VirtualNodeName = &v
	return s
}

func (s *CreateVirtualNodeRequest) SetZoneId(v string) *CreateVirtualNodeRequest {
	s.ZoneId = &v
	return s
}

type CreateVirtualNodeRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVirtualNodeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeRequestTag) SetKey(v string) *CreateVirtualNodeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVirtualNodeRequestTag) SetValue(v string) *CreateVirtualNodeRequestTag {
	s.Value = &v
	return s
}

type CreateVirtualNodeRequestTaint struct {
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVirtualNodeRequestTaint) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeRequestTaint) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeRequestTaint) SetEffect(v string) *CreateVirtualNodeRequestTaint {
	s.Effect = &v
	return s
}

func (s *CreateVirtualNodeRequestTaint) SetKey(v string) *CreateVirtualNodeRequestTaint {
	s.Key = &v
	return s
}

func (s *CreateVirtualNodeRequestTaint) SetValue(v string) *CreateVirtualNodeRequestTaint {
	s.Value = &v
	return s
}

type CreateVirtualNodeResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VirtualNodeId *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
}

func (s CreateVirtualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeResponseBody) SetRequestId(v string) *CreateVirtualNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualNodeResponseBody) SetVirtualNodeId(v string) *CreateVirtualNodeResponseBody {
	s.VirtualNodeId = &v
	return s
}

type CreateVirtualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVirtualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualNodeResponse) SetHeaders(v map[string]*string) *CreateVirtualNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualNodeResponse) SetStatusCode(v int32) *CreateVirtualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualNodeResponse) SetBody(v *CreateVirtualNodeResponseBody) *CreateVirtualNodeResponse {
	s.Body = v
	return s
}

type DeleteContainerGroupRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupRequest) SetClientToken(v string) *DeleteContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetContainerGroupId(v string) *DeleteContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetOwnerAccount(v string) *DeleteContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetOwnerId(v int64) *DeleteContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetRegionId(v string) *DeleteContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetResourceOwnerAccount(v string) *DeleteContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteContainerGroupRequest) SetResourceOwnerId(v int64) *DeleteContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteContainerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupResponseBody) SetRequestId(v string) *DeleteContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContainerGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerGroupResponse) SetHeaders(v map[string]*string) *DeleteContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerGroupResponse) SetStatusCode(v int32) *DeleteContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContainerGroupResponse) SetBody(v *DeleteContainerGroupResponseBody) *DeleteContainerGroupResponse {
	s.Body = v
	return s
}

type DeleteImageCacheRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ImageCacheId         *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheRequest) SetClientToken(v string) *DeleteImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteImageCacheRequest) SetImageCacheId(v string) *DeleteImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetOwnerAccount(v string) *DeleteImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImageCacheRequest) SetOwnerId(v int64) *DeleteImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetRegionId(v string) *DeleteImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetResourceOwnerAccount(v string) *DeleteImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImageCacheRequest) SetResourceOwnerId(v int64) *DeleteImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteImageCacheResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponseBody) SetRequestId(v string) *DeleteImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponse) SetHeaders(v map[string]*string) *DeleteImageCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageCacheResponse) SetStatusCode(v int32) *DeleteImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageCacheResponse) SetBody(v *DeleteImageCacheResponseBody) *DeleteImageCacheResponse {
	s.Body = v
	return s
}

type DeleteVirtualNodeRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VirtualNodeId        *string `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
}

func (s DeleteVirtualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualNodeRequest) SetClientToken(v string) *DeleteVirtualNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetOwnerAccount(v string) *DeleteVirtualNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetOwnerId(v int64) *DeleteVirtualNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetRegionId(v string) *DeleteVirtualNodeRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetResourceOwnerAccount(v string) *DeleteVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetResourceOwnerId(v int64) *DeleteVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVirtualNodeRequest) SetVirtualNodeId(v string) *DeleteVirtualNodeRequest {
	s.VirtualNodeId = &v
	return s
}

type DeleteVirtualNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualNodeResponseBody) SetRequestId(v string) *DeleteVirtualNodeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVirtualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualNodeResponse) SetHeaders(v map[string]*string) *DeleteVirtualNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualNodeResponse) SetStatusCode(v int32) *DeleteVirtualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualNodeResponse) SetBody(v *DeleteVirtualNodeResponseBody) *DeleteVirtualNodeResponse {
	s.Body = v
	return s
}

type DescribeAvailableResourceRequest struct {
	DestinationResource  *DescribeAvailableResourceRequestDestinationResource `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty" type:"Struct"`
	OwnerAccount         *string                                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpotResource         *DescribeAvailableResourceRequestSpotResource        `json:"SpotResource,omitempty" xml:"SpotResource,omitempty" type:"Struct"`
	ZoneId               *string                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequest) SetDestinationResource(v *DescribeAvailableResourceRequestDestinationResource) *DescribeAvailableResourceRequest {
	s.DestinationResource = v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetRegionId(v string) *DescribeAvailableResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerAccount(v string) *DescribeAvailableResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetResourceOwnerId(v int64) *DescribeAvailableResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAvailableResourceRequest) SetSpotResource(v *DescribeAvailableResourceRequestSpotResource) *DescribeAvailableResourceRequest {
	s.SpotResource = v
	return s
}

func (s *DescribeAvailableResourceRequest) SetZoneId(v string) *DescribeAvailableResourceRequest {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceRequestDestinationResource struct {
	Category *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	Cores    *float32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Memory   *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Value    *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAvailableResourceRequestDestinationResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequestDestinationResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetCategory(v string) *DescribeAvailableResourceRequestDestinationResource {
	s.Category = &v
	return s
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetCores(v float32) *DescribeAvailableResourceRequestDestinationResource {
	s.Cores = &v
	return s
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetMemory(v float32) *DescribeAvailableResourceRequestDestinationResource {
	s.Memory = &v
	return s
}

func (s *DescribeAvailableResourceRequestDestinationResource) SetValue(v string) *DescribeAvailableResourceRequestDestinationResource {
	s.Value = &v
	return s
}

type DescribeAvailableResourceRequestSpotResource struct {
	SpotDuration   *int32   `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy   *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s DescribeAvailableResourceRequestSpotResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceRequestSpotResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceRequestSpotResource) SetSpotDuration(v int32) *DescribeAvailableResourceRequestSpotResource {
	s.SpotDuration = &v
	return s
}

func (s *DescribeAvailableResourceRequestSpotResource) SetSpotPriceLimit(v float64) *DescribeAvailableResourceRequestSpotResource {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeAvailableResourceRequestSpotResource) SetSpotStrategy(v string) *DescribeAvailableResourceRequestSpotResource {
	s.SpotStrategy = &v
	return s
}

type DescribeAvailableResourceResponseBody struct {
	AvailableZones *DescribeAvailableResourceResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Struct"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBody) SetAvailableZones(v *DescribeAvailableResourceResponseBodyAvailableZones) *DescribeAvailableResourceResponseBody {
	s.AvailableZones = v
	return s
}

func (s *DescribeAvailableResourceResponseBody) SetRequestId(v string) *DescribeAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZones struct {
	AvailableZone []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone `json:"AvailableZone,omitempty" xml:"AvailableZone,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZones) SetAvailableZone(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) *DescribeAvailableResourceResponseBodyAvailableZones {
	s.AvailableZone = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone struct {
	AvailableResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Struct"`
	RegionId           *string                                                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId             *string                                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetAvailableResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.AvailableResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetRegionId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone) SetZoneId(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZone {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources struct {
	AvailableResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource `json:"AvailableResource,omitempty" xml:"AvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources) SetAvailableResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResources {
	s.AvailableResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource struct {
	SupportedResources *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources `json:"SupportedResources,omitempty" xml:"SupportedResources,omitempty" type:"Struct"`
	Type               *string                                                                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetSupportedResources(v *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.SupportedResources = v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource) SetType(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResource {
	s.Type = &v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources struct {
	SupportedResource []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource `json:"SupportedResource,omitempty" xml:"SupportedResource,omitempty" type:"Repeated"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources) SetSupportedResource(v []*DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResources {
	s.SupportedResource = v
	return s
}

type DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource struct {
	StatusCategory *string `json:"StatusCategory,omitempty" xml:"StatusCategory,omitempty"`
	Value          *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetStatusCategory(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.StatusCategory = &v
	return s
}

func (s *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource) SetValue(v string) *DescribeAvailableResourceResponseBodyAvailableZonesAvailableZoneAvailableResourcesAvailableResourceSupportedResourcesSupportedResource {
	s.Value = &v
	return s
}

type DescribeAvailableResourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceResponse) SetStatusCode(v int32) *DescribeAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourceResponse) SetBody(v *DescribeAvailableResourceResponseBody) *DescribeAvailableResourceResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupEventsRequest struct {
	ContainerGroupIds *string                                   `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	EventSource       *string                                   `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	Limit             *int32                                    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken         *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId          *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SinceSecond       *int32                                    `json:"SinceSecond,omitempty" xml:"SinceSecond,omitempty"`
	Tag               []*DescribeContainerGroupEventsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId         *string                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId            *string                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsRequest) SetContainerGroupIds(v string) *DescribeContainerGroupEventsRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetEventSource(v string) *DescribeContainerGroupEventsRequest {
	s.EventSource = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetLimit(v int32) *DescribeContainerGroupEventsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetNextToken(v string) *DescribeContainerGroupEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetRegionId(v string) *DescribeContainerGroupEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetResourceGroupId(v string) *DescribeContainerGroupEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetSinceSecond(v int32) *DescribeContainerGroupEventsRequest {
	s.SinceSecond = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetTag(v []*DescribeContainerGroupEventsRequestTag) *DescribeContainerGroupEventsRequest {
	s.Tag = v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetVSwitchId(v string) *DescribeContainerGroupEventsRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupEventsRequest) SetZoneId(v string) *DescribeContainerGroupEventsRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupEventsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupEventsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsRequestTag) SetKey(v string) *DescribeContainerGroupEventsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupEventsRequestTag) SetValue(v string) *DescribeContainerGroupEventsRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupEventsResponseBody struct {
	Data       []*DescribeContainerGroupEventsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBody) SetData(v []*DescribeContainerGroupEventsResponseBodyData) *DescribeContainerGroupEventsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBody) SetRequestId(v string) *DescribeContainerGroupEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBody) SetTotalCount(v int32) *DescribeContainerGroupEventsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyData struct {
	ContainerGroupId *string                                               `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Events           []*DescribeContainerGroupEventsResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetContainerGroupId(v string) *DescribeContainerGroupEventsResponseBodyData {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyData) SetEvents(v []*DescribeContainerGroupEventsResponseBodyDataEvents) *DescribeContainerGroupEventsResponseBodyData {
	s.Events = v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEvents struct {
	Count              *int32                                                            `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstTimestamp     *string                                                           `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp      *string                                                           `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message            *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Metadata           *DescribeContainerGroupEventsResponseBodyDataEventsMetadata       `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	Reason             *string                                                           `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ReportingComponent *string                                                           `json:"ReportingComponent,omitempty" xml:"ReportingComponent,omitempty"`
	ReportingInstance  *string                                                           `json:"ReportingInstance,omitempty" xml:"ReportingInstance,omitempty"`
	Source             *DescribeContainerGroupEventsResponseBodyDataEventsSource         `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	Type               *string                                                           `json:"Type,omitempty" xml:"Type,omitempty"`
	InvolvedObject     *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject `json:"involvedObject,omitempty" xml:"involvedObject,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetCount(v int32) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Count = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetFirstTimestamp(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetLastTimestamp(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetMessage(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetMetadata(v *DescribeContainerGroupEventsResponseBodyDataEventsMetadata) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Metadata = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetReason(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetReportingComponent(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.ReportingComponent = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetReportingInstance(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.ReportingInstance = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetSource(v *DescribeContainerGroupEventsResponseBodyDataEventsSource) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Source = v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetType(v string) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.Type = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEvents) SetInvolvedObject(v *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) *DescribeContainerGroupEventsResponseBodyDataEvents {
	s.InvolvedObject = v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEventsMetadata struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsMetadata) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsMetadata) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsMetadata) SetName(v string) *DescribeContainerGroupEventsResponseBodyDataEventsMetadata {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsMetadata) SetNamespace(v string) *DescribeContainerGroupEventsResponseBodyDataEventsMetadata {
	s.Namespace = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEventsSource struct {
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsSource) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsSource) SetComponent(v string) *DescribeContainerGroupEventsResponseBodyDataEventsSource {
	s.Component = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsSource) SetHost(v string) *DescribeContainerGroupEventsResponseBodyDataEventsSource {
	s.Host = &v
	return s
}

type DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject struct {
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	Kind       *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Uid        *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetApiVersion(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.ApiVersion = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetKind(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Kind = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetName(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetNamespace(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject) SetUid(v string) *DescribeContainerGroupEventsResponseBodyDataEventsInvolvedObject {
	s.Uid = &v
	return s
}

type DescribeContainerGroupEventsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupEventsResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupEventsResponse) SetStatusCode(v int32) *DescribeContainerGroupEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupEventsResponse) SetBody(v *DescribeContainerGroupEventsResponseBody) *DescribeContainerGroupEventsResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupMetricRequest struct {
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeContainerGroupMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricRequest) SetContainerGroupId(v string) *DescribeContainerGroupMetricRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetEndTime(v string) *DescribeContainerGroupMetricRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetOwnerAccount(v string) *DescribeContainerGroupMetricRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetOwnerId(v int64) *DescribeContainerGroupMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetPeriod(v string) *DescribeContainerGroupMetricRequest {
	s.Period = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetRegionId(v string) *DescribeContainerGroupMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupMetricRequest) SetStartTime(v string) *DescribeContainerGroupMetricRequest {
	s.StartTime = &v
	return s
}

type DescribeContainerGroupMetricResponseBody struct {
	ContainerGroupId *string                                            `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Records          []*DescribeContainerGroupMetricResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBody) SetContainerGroupId(v string) *DescribeContainerGroupMetricResponseBody {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBody) SetRecords(v []*DescribeContainerGroupMetricResponseBodyRecords) *DescribeContainerGroupMetricResponseBody {
	s.Records = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBody) SetRequestId(v string) *DescribeContainerGroupMetricResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecords struct {
	CPU        *DescribeContainerGroupMetricResponseBodyRecordsCPU          `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Containers []*DescribeContainerGroupMetricResponseBodyRecordsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	Disk       []*DescribeContainerGroupMetricResponseBodyRecordsDisk       `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	Filesystem []*DescribeContainerGroupMetricResponseBodyRecordsFilesystem `json:"Filesystem,omitempty" xml:"Filesystem,omitempty" type:"Repeated"`
	Memory     *DescribeContainerGroupMetricResponseBodyRecordsMemory       `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Network    *DescribeContainerGroupMetricResponseBodyRecordsNetwork      `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Timestamp  *string                                                      `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetCPU(v *DescribeContainerGroupMetricResponseBodyRecordsCPU) *DescribeContainerGroupMetricResponseBodyRecords {
	s.CPU = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetContainers(v []*DescribeContainerGroupMetricResponseBodyRecordsContainers) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Containers = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetDisk(v []*DescribeContainerGroupMetricResponseBodyRecordsDisk) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Disk = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetFilesystem(v []*DescribeContainerGroupMetricResponseBodyRecordsFilesystem) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Filesystem = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetMemory(v *DescribeContainerGroupMetricResponseBodyRecordsMemory) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Memory = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetNetwork(v *DescribeContainerGroupMetricResponseBodyRecordsNetwork) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Network = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecords) SetTimestamp(v string) *DescribeContainerGroupMetricResponseBodyRecords {
	s.Timestamp = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsCPU struct {
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsCPU) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetLimit(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetLoad(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.Load = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetUsageCoreNanoSeconds(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsCPU) SetUsageNanoCores(v int64) *DescribeContainerGroupMetricResponseBodyRecordsCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainers struct {
	CPU    *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU    `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Memory *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetCPU(v *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.CPU = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetMemory(v *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.Memory = v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainers) SetName(v string) *DescribeContainerGroupMetricResponseBodyRecordsContainers {
	s.Name = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainersCPU struct {
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetLimit(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetLoad(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.Load = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetUsageCoreNanoSeconds(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU) SetUsageNanoCores(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsContainersMemory struct {
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetAvailableBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetCache(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.Cache = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetRss(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.Rss = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetUsageBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory) SetWorkingSet(v int64) *DescribeContainerGroupMetricResponseBodyRecordsContainersMemory {
	s.WorkingSet = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsDisk struct {
	Device     *string `json:"Device,omitempty" xml:"Device,omitempty"`
	ReadBytes  *int64  `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	ReadIO     *int64  `json:"ReadIO,omitempty" xml:"ReadIO,omitempty"`
	WriteBytes *int64  `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	WriteIO    *int64  `json:"WriteIO,omitempty" xml:"WriteIO,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsDisk) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetDevice(v string) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.Device = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetReadBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.ReadBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetReadIO(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.ReadIO = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetWriteBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.WriteBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsDisk) SetWriteIO(v int64) *DescribeContainerGroupMetricResponseBodyRecordsDisk {
	s.WriteIO = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsFilesystem struct {
	Available *int64  `json:"Available,omitempty" xml:"Available,omitempty"`
	Capacity  *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Category  *string `json:"Category,omitempty" xml:"Category,omitempty"`
	FsName    *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	Usage     *int64  `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsFilesystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsFilesystem) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetAvailable(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Available = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetCapacity(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Capacity = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetCategory(v string) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Category = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetFsName(v string) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.FsName = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsFilesystem) SetUsage(v int64) *DescribeContainerGroupMetricResponseBodyRecordsFilesystem {
	s.Usage = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsMemory struct {
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsMemory) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetAvailableBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetCache(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.Cache = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetRss(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.Rss = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetUsageBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsMemory) SetWorkingSet(v int64) *DescribeContainerGroupMetricResponseBodyRecordsMemory {
	s.WorkingSet = &v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsNetwork struct {
	Interfaces []*DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetwork) SetInterfaces(v []*DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) *DescribeContainerGroupMetricResponseBodyRecordsNetwork {
	s.Interfaces = v
	return s
}

type DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RxBytes   *int64  `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	RxDrops   *int64  `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
	RxErrors  *int64  `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	RxPackets *int64  `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	TxBytes   *int64  `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	TxDrops   *int64  `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	TxErrors  *int64  `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	TxPackets *int64  `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetName(v string) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxDrops(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxDrops = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxErrors(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxErrors = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetRxPackets(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.RxPackets = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxBytes(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxBytes = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxDrops(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxDrops = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxErrors(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxErrors = &v
	return s
}

func (s *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces) SetTxPackets(v int64) *DescribeContainerGroupMetricResponseBodyRecordsNetworkInterfaces {
	s.TxPackets = &v
	return s
}

type DescribeContainerGroupMetricResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupMetricResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupMetricResponse) SetStatusCode(v int32) *DescribeContainerGroupMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupMetricResponse) SetBody(v *DescribeContainerGroupMetricResponseBody) *DescribeContainerGroupMetricResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupPriceRequest struct {
	Cpu                  *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EphemeralStorage     *int32   `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	InstanceType         *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory               *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount         *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SpotPriceLimit       *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy         *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	ZoneId               *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceRequest) SetCpu(v float32) *DescribeContainerGroupPriceRequest {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetEphemeralStorage(v int32) *DescribeContainerGroupPriceRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetInstanceType(v string) *DescribeContainerGroupPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetMemory(v float32) *DescribeContainerGroupPriceRequest {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetOwnerAccount(v string) *DescribeContainerGroupPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetOwnerId(v int64) *DescribeContainerGroupPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetRegionId(v string) *DescribeContainerGroupPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotPriceLimit(v float32) *DescribeContainerGroupPriceRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetSpotStrategy(v string) *DescribeContainerGroupPriceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeContainerGroupPriceRequest) SetZoneId(v string) *DescribeContainerGroupPriceRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupPriceResponseBody struct {
	PriceInfo *DescribeContainerGroupPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBody) SetPriceInfo(v *DescribeContainerGroupPriceResponseBodyPriceInfo) *DescribeContainerGroupPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBody) SetRequestId(v string) *DescribeContainerGroupPriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfo struct {
	Price      *DescribeContainerGroupPriceResponseBodyPriceInfoPrice      `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules      *DescribeContainerGroupPriceResponseBodyPriceInfoRules      `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	SpotPrices *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices `json:"SpotPrices,omitempty" xml:"SpotPrices,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetPrice(v *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetRules(v *DescribeContainerGroupPriceResponseBodyPriceInfoRules) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfo) SetSpotPrices(v *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) *DescribeContainerGroupPriceResponseBodyPriceInfo {
	s.SpotPrices = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPrice struct {
	Currency      *string                                                           `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DetailInfos   *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Struct"`
	DiscountPrice *float32                                                          `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32                                                          `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	TradePrice    *float32                                                          `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetDetailInfos(v *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.DetailInfos = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos struct {
	DetailInfo []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos) SetDetailInfo(v []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfos {
	s.DetailInfo = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo struct {
	DiscountPrice *float32                                                                         `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *float32                                                                         `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Resource      *string                                                                          `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Rules         *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	TradePrice    *float32                                                                         `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetDiscountPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetOriginalPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetResource(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Resource = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetRules(v *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.Rules = v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo) SetTradePrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfo {
	s.TradePrice = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules struct {
	Rule []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules) SetRule(v []*DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRules {
	s.Rule = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) SetDescription(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule) SetRuleId(v int64) *DescribeContainerGroupPriceResponseBodyPriceInfoPriceDetailInfosDetailInfoRulesRule {
	s.RuleId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoRules struct {
	Rule []*DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRules) SetRule(v []*DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) *DescribeContainerGroupPriceResponseBodyPriceInfoRules {
	s.Rule = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) SetDescription(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule) SetRuleId(v int64) *DescribeContainerGroupPriceResponseBodyPriceInfoRulesRule {
	s.RuleId = &v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices struct {
	SpotPrice []*DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices) SetSpotPrice(v []*DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPrices {
	s.SpotPrice = v
	return s
}

type DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OriginPrice  *float32 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	SpotPrice    *float32 `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	ZoneId       *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetInstanceType(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetOriginPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.OriginPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetSpotPrice(v float32) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.SpotPrice = &v
	return s
}

func (s *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice) SetZoneId(v string) *DescribeContainerGroupPriceResponseBodyPriceInfoSpotPricesSpotPrice {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupPriceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupPriceResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupPriceResponse) SetStatusCode(v int32) *DescribeContainerGroupPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupPriceResponse) SetBody(v *DescribeContainerGroupPriceResponseBody) *DescribeContainerGroupPriceResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupStatusRequest struct {
	ContainerGroupIds *string                                   `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	Limit             *int32                                    `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken         *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId          *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SinceSecond       *int32                                    `json:"SinceSecond,omitempty" xml:"SinceSecond,omitempty"`
	Tag               []*DescribeContainerGroupStatusRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId         *string                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId            *string                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusRequest) SetContainerGroupIds(v string) *DescribeContainerGroupStatusRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetLimit(v int32) *DescribeContainerGroupStatusRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetNextToken(v string) *DescribeContainerGroupStatusRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetRegionId(v string) *DescribeContainerGroupStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetResourceGroupId(v string) *DescribeContainerGroupStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetSinceSecond(v int32) *DescribeContainerGroupStatusRequest {
	s.SinceSecond = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetTag(v []*DescribeContainerGroupStatusRequestTag) *DescribeContainerGroupStatusRequest {
	s.Tag = v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetVSwitchId(v string) *DescribeContainerGroupStatusRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupStatusRequest) SetZoneId(v string) *DescribeContainerGroupStatusRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupStatusRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupStatusRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusRequestTag) SetKey(v string) *DescribeContainerGroupStatusRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupStatusRequestTag) SetValue(v string) *DescribeContainerGroupStatusRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupStatusResponseBody struct {
	Data       []*DescribeContainerGroupStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken  *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBody) SetData(v []*DescribeContainerGroupStatusResponseBodyData) *DescribeContainerGroupStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBody) SetNextToken(v string) *DescribeContainerGroupStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBody) SetRequestId(v string) *DescribeContainerGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBody) SetTotalCount(v int32) *DescribeContainerGroupStatusResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyData struct {
	ContainerGroupId *string                                                `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Name             *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace        *string                                                `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	PodStatus        *DescribeContainerGroupStatusResponseBodyDataPodStatus `json:"PodStatus,omitempty" xml:"PodStatus,omitempty" type:"Struct"`
	Status           *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Uuid             *string                                                `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetContainerGroupId(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetName(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetNamespace(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetPodStatus(v *DescribeContainerGroupStatusResponseBodyDataPodStatus) *DescribeContainerGroupStatusResponseBodyData {
	s.PodStatus = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetStatus(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyData) SetUuid(v string) *DescribeContainerGroupStatusResponseBodyData {
	s.Uuid = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatus struct {
	Conditions        []*DescribeContainerGroupStatusResponseBodyDataPodStatusConditions        `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	ContainerStatuses []*DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses `json:"ContainerStatuses,omitempty" xml:"ContainerStatuses,omitempty" type:"Repeated"`
	HostIp            *string                                                                   `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	Phase             *string                                                                   `json:"Phase,omitempty" xml:"Phase,omitempty"`
	PodIp             *string                                                                   `json:"PodIp,omitempty" xml:"PodIp,omitempty"`
	PodIps            []*DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps            `json:"PodIps,omitempty" xml:"PodIps,omitempty" type:"Repeated"`
	QosClass          *string                                                                   `json:"QosClass,omitempty" xml:"QosClass,omitempty"`
	StartTime         *string                                                                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatus) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetConditions(v []*DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.Conditions = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetContainerStatuses(v []*DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.ContainerStatuses = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetHostIp(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.HostIp = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetPhase(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.Phase = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetPodIp(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.PodIp = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetPodIps(v []*DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.PodIps = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetQosClass(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.QosClass = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatus) SetStartTime(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatus {
	s.StartTime = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusConditions struct {
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason             *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetLastTransitionTime(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetStatus(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions) SetType(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusConditions {
	s.Type = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses struct {
	Image        *string                                                                          `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageID      *string                                                                          `json:"ImageID,omitempty" xml:"ImageID,omitempty"`
	LastState    *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState `json:"LastState,omitempty" xml:"LastState,omitempty" type:"Struct"`
	Name         *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Ready        *bool                                                                            `json:"Ready,omitempty" xml:"Ready,omitempty"`
	RestartCount *int32                                                                           `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	Started      *bool                                                                            `json:"Started,omitempty" xml:"Started,omitempty"`
	State        *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState     `json:"State,omitempty" xml:"State,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetImage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetImageID(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.ImageID = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetLastState(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.LastState = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetName(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetReady(v bool) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetRestartCount(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetStarted(v bool) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.Started = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses) SetState(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatuses {
	s.State = v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState struct {
	Running    *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning    `json:"Running,omitempty" xml:"Running,omitempty" type:"Struct"`
	Terminated *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated `json:"Terminated,omitempty" xml:"Terminated,omitempty" type:"Struct"`
	Waiting    *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting    `json:"Waiting,omitempty" xml:"Waiting,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) SetRunning(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState {
	s.Running = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) SetTerminated(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState {
	s.Terminated = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState) SetWaiting(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastState {
	s.Waiting = v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning struct {
	StartedAtstartedAt *string `json:"StartedAtstartedAt,omitempty" xml:"StartedAtstartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning) SetStartedAtstartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateRunning {
	s.StartedAtstartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated struct {
	ContainerID *string `json:"ContainerID,omitempty" xml:"ContainerID,omitempty"`
	ExitCode    *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishedAt  *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Signal      *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	StartedAt   *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetContainerID(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.ContainerID = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetExitCode(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetFinishedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.FinishedAt = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetSignal(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated) SetStartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateTerminated {
	s.StartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason  *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesLastStateWaiting {
	s.Reason = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState struct {
	Running    *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning    `json:"Running,omitempty" xml:"Running,omitempty" type:"Struct"`
	Terminated *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated `json:"Terminated,omitempty" xml:"Terminated,omitempty" type:"Struct"`
	Waiting    *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting    `json:"Waiting,omitempty" xml:"Waiting,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) SetRunning(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState {
	s.Running = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) SetTerminated(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState {
	s.Terminated = v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState) SetWaiting(v *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesState {
	s.Waiting = v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning struct {
	StartedAtstartedAt *string `json:"StartedAtstartedAt,omitempty" xml:"StartedAtstartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning) SetStartedAtstartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateRunning {
	s.StartedAtstartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated struct {
	ContainerID *string `json:"ContainerID,omitempty" xml:"ContainerID,omitempty"`
	ExitCode    *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishedAt  *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Signal      *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	StartedAt   *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetContainerID(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.ContainerID = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetExitCode(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetFinishedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.FinishedAt = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetSignal(v int32) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated) SetStartedAt(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateTerminated {
	s.StartedAt = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason  *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) SetMessage(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting) SetReason(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusContainerStatusesStateWaiting {
	s.Reason = &v
	return s
}

type DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps struct {
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps) SetIp(v string) *DescribeContainerGroupStatusResponseBodyDataPodStatusPodIps {
	s.Ip = &v
	return s
}

type DescribeContainerGroupStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupStatusResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupStatusResponse) SetStatusCode(v int32) *DescribeContainerGroupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupStatusResponse) SetBody(v *DescribeContainerGroupStatusResponseBody) *DescribeContainerGroupStatusResponse {
	s.Body = v
	return s
}

type DescribeContainerGroupsRequest struct {
	ContainerGroupIds    *string                              `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	ContainerGroupName   *string                              `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	Limit                *int32                               `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken            *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag                  []*DescribeContainerGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId            *string                              `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	WithEvent            *bool                                `json:"WithEvent,omitempty" xml:"WithEvent,omitempty"`
	ZoneId               *string                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequest) SetContainerGroupIds(v string) *DescribeContainerGroupsRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetContainerGroupName(v string) *DescribeContainerGroupsRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetLimit(v int32) *DescribeContainerGroupsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetNextToken(v string) *DescribeContainerGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetOwnerAccount(v string) *DescribeContainerGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetOwnerId(v int64) *DescribeContainerGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetRegionId(v string) *DescribeContainerGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceGroupId(v string) *DescribeContainerGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceOwnerAccount(v string) *DescribeContainerGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetResourceOwnerId(v int64) *DescribeContainerGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetStatus(v string) *DescribeContainerGroupsRequest {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetTag(v []*DescribeContainerGroupsRequestTag) *DescribeContainerGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeContainerGroupsRequest) SetVSwitchId(v string) *DescribeContainerGroupsRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetWithEvent(v bool) *DescribeContainerGroupsRequest {
	s.WithEvent = &v
	return s
}

func (s *DescribeContainerGroupsRequest) SetZoneId(v string) *DescribeContainerGroupsRequest {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsRequestTag) SetKey(v string) *DescribeContainerGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsRequestTag) SetValue(v string) *DescribeContainerGroupsRequestTag {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBody struct {
	ContainerGroups []*DescribeContainerGroupsResponseBodyContainerGroups `json:"ContainerGroups,omitempty" xml:"ContainerGroups,omitempty" type:"Repeated"`
	NextToken       *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeContainerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBody) SetContainerGroups(v []*DescribeContainerGroupsResponseBodyContainerGroups) *DescribeContainerGroupsResponseBody {
	s.ContainerGroups = v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetNextToken(v string) *DescribeContainerGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetRequestId(v string) *DescribeContainerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBody) SetTotalCount(v int32) *DescribeContainerGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroups struct {
	ContainerGroupId      *string                                                               `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ContainerGroupName    *string                                                               `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	Containers            []*DescribeContainerGroupsResponseBodyContainerGroupsContainers       `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	Cpu                   *float32                                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime          *string                                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Discount              *int32                                                                `json:"Discount,omitempty" xml:"Discount,omitempty"`
	DnsConfig             *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig          `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	EciSecurityContext    *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext `json:"EciSecurityContext,omitempty" xml:"EciSecurityContext,omitempty" type:"Struct"`
	EniInstanceId         *string                                                               `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	EphemeralStorage      *int32                                                                `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	Events                []*DescribeContainerGroupsResponseBodyContainerGroupsEvents           `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	ExpiredTime           *string                                                               `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FailedTime            *string                                                               `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
	HostAliases           []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases      `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	InitContainers        []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers   `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	InstanceType          *string                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetIp            *string                                                               `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IntranetIp            *string                                                               `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Ipv6Address           *string                                                               `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	Memory                *float32                                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	RamRoleName           *string                                                               `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RegionId              *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string                                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RestartPolicy         *string                                                               `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	SecurityGroupId       *string                                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SpotPriceLimit        *float64                                                              `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	SpotStrategy          *string                                                               `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	Status                *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	SucceededTime         *string                                                               `json:"SucceededTime,omitempty" xml:"SucceededTime,omitempty"`
	Tags                  []*DescribeContainerGroupsResponseBodyContainerGroupsTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TenantEniInstanceId   *string                                                               `json:"TenantEniInstanceId,omitempty" xml:"TenantEniInstanceId,omitempty"`
	TenantEniIp           *string                                                               `json:"TenantEniIp,omitempty" xml:"TenantEniIp,omitempty"`
	TenantSecurityGroupId *string                                                               `json:"TenantSecurityGroupId,omitempty" xml:"TenantSecurityGroupId,omitempty"`
	TenantVSwitchId       *string                                                               `json:"TenantVSwitchId,omitempty" xml:"TenantVSwitchId,omitempty"`
	VSwitchId             *string                                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Volumes               []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes          `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	VpcId                 *string                                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                *string                                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroups) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainerGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainerGroupName(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetContainers(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainers) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Containers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetCreationTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.CreationTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDiscount(v int32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Discount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetDnsConfig(v *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.DnsConfig = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEciSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EciSecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEniInstanceId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEphemeralStorage(v int32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetEvents(v []*DescribeContainerGroupsResponseBodyContainerGroupsEvents) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Events = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetExpiredTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetFailedTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.FailedTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetHostAliases(v []*DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.HostAliases = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInitContainers(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InitContainers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInstanceType(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InstanceType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetInternetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.InternetIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetIntranetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.IntranetIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetIpv6Address(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRamRoleName(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RamRoleName = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRegionId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetResourceGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetRestartPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.RestartPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSecurityGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSpotPriceLimit(v float64) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSpotStrategy(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Status = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetSucceededTime(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.SucceededTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTags(v []*DescribeContainerGroupsResponseBodyContainerGroupsTags) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Tags = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantEniInstanceId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantEniInstanceId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantEniIp(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantEniIp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantSecurityGroupId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantSecurityGroupId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetTenantVSwitchId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.TenantVSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVSwitchId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.VSwitchId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVolumes(v []*DescribeContainerGroupsResponseBodyContainerGroupsVolumes) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.Volumes = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetVpcId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.VpcId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroups) SetZoneId(v string) *DescribeContainerGroupsResponseBodyContainerGroups {
	s.ZoneId = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainers struct {
	Args            []*string                                                                      `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Commands        []*string                                                                      `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Cpu             *float32                                                                       `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CurrentState    *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState      `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	Gpu             *int32                                                                         `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image           *string                                                                        `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy *string                                                                        `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	LivenessProbe   *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe     `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	Memory          *float32                                                                       `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name            *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Ports           []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	PreviousState   *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState     `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	ReadinessProbe  *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe    `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	Ready           *bool                                                                          `json:"Ready,omitempty" xml:"Ready,omitempty"`
	RestartCount    *int32                                                                         `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext   `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	Stdin           *bool                                                                          `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce       *bool                                                                          `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Tty             *bool                                                                          `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMounts    []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	WorkingDir      *string                                                                        `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetArgs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Args = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCommands(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Commands = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetCurrentState(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.CurrentState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetEnvironmentVars(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetGpu(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetImage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetImagePullPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetLivenessProbe(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.LivenessProbe = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetPorts(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Ports = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetPreviousState(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.PreviousState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetReadinessProbe(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.ReadinessProbe = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetReady(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetRestartCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.SecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetStdin(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Stdin = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetStdinOnce(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.StdinOnce = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetTty(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.Tty = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetVolumeMounts(v []*DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainers) SetWorkingDir(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainers {
	s.WorkingDir = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState struct {
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersCurrentState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars struct {
	Key       *string                                                                               `json:"Key,omitempty" xml:"Key,omitempty"`
	Value     *string                                                                               `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueFrom *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars) SetValueFrom(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVars {
	s.ValueFrom = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom struct {
	FieldRef *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom) SetFieldRef(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFrom {
	s.FieldRef = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef) SetFieldPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersEnvironmentVarsValueFromFieldRef {
	s.FieldPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe struct {
	Execs               []*string                                                                           `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	FailureThreshold    *int32                                                                              `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                                              `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                              `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                              `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	TimeoutSeconds      *int32                                                                              `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetExecs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.Execs = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetFailureThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetHttpGet(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetInitialDelaySeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetPeriodSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetSuccessThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetTcpSocket(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe) SetTimeoutSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet) SetScheme(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket struct {
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) SetHost(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts) SetProtocol(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPorts {
	s.Protocol = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState struct {
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersPreviousState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe struct {
	Execs               []*string                                                                            `json:"Execs,omitempty" xml:"Execs,omitempty" type:"Repeated"`
	FailureThreshold    *int32                                                                               `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                                               `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                               `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                               `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	TimeoutSeconds      *int32                                                                               `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetExecs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.Execs = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetFailureThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetHttpGet(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetInitialDelaySeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetPeriodSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetSuccessThreshold(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetTcpSocket(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe) SetTimeoutSeconds(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet) SetScheme(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket struct {
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Port *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) SetHost(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket {
	s.Host = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext struct {
	Capability             *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                                  `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                                 `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetCapability(v *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext) SetRunAsUser(v int64) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability) SetAdds(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetMountPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetMountPropagation(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts) SetReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig struct {
	NameServers []*string                                                             `json:"NameServers,omitempty" xml:"NameServers,omitempty" type:"Repeated"`
	Options     []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	Searches    []*string                                                             `json:"Searches,omitempty" xml:"Searches,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetNameServers(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.NameServers = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetOptions(v []*DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.Options = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig) SetSearches(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfig {
	s.Searches = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsDnsConfigOptions {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext struct {
	Sysctls []*DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls `json:"Sysctls,omitempty" xml:"Sysctls,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext) SetSysctls(v []*DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContext {
	s.Sysctls = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEciSecurityContextSysctls {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsEvents struct {
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsEvents) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Count = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetFirstTimestamp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetLastTimestamp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsEvents) SetType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsEvents {
	s.Type = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsHostAliases struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Ip        *string   `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) SetHostnames(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases {
	s.Hostnames = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases) SetIp(v string) *DescribeContainerGroupsResponseBodyContainerGroupsHostAliases {
	s.Ip = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainers struct {
	Args            []*string                                                                          `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Command         []*string                                                                          `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Cpu             *float32                                                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CurrentState    *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState      `json:"CurrentState,omitempty" xml:"CurrentState,omitempty" type:"Struct"`
	EnvironmentVars []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	Gpu             *int32                                                                             `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image           *string                                                                            `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy *string                                                                            `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	Memory          *float32                                                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name            *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Ports           []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts           `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	PreviousState   *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState     `json:"PreviousState,omitempty" xml:"PreviousState,omitempty" type:"Struct"`
	Ready           *bool                                                                              `json:"Ready,omitempty" xml:"Ready,omitempty"`
	RestartCount    *int32                                                                             `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	SecurityContext *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext   `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	VolumeMounts    []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts    `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	WorkingDir      *string                                                                            `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetArgs(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Args = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCommand(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Command = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCpu(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetCurrentState(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.CurrentState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetEnvironmentVars(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetGpu(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetImage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Image = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetImagePullPolicy(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetMemory(v float32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Memory = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetPorts(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Ports = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetPreviousState(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.PreviousState = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetReady(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.Ready = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetRestartCount(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.RestartCount = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetSecurityContext(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.SecurityContext = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetVolumeMounts(v []*DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers) SetWorkingDir(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainers {
	s.WorkingDir = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState struct {
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersCurrentState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars struct {
	Key       *string                                                                                   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value     *string                                                                                   `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueFrom *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom `json:"ValueFrom,omitempty" xml:"ValueFrom,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars) SetValueFrom(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVars {
	s.ValueFrom = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom struct {
	FieldRef *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom) SetFieldRef(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFrom {
	s.FieldRef = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef) SetFieldPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersEnvironmentVarsValueFromFieldRef {
	s.FieldPath = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) SetPort(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts) SetProtocol(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPorts {
	s.Protocol = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState struct {
	DetailStatus *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	ExitCode     *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime   *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Signal       *int32  `json:"Signal,omitempty" xml:"Signal,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetDetailStatus(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.DetailStatus = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetExitCode(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.ExitCode = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetFinishTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.FinishTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetMessage(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Message = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetReason(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Reason = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetSignal(v int32) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.Signal = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetStartTime(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState) SetState(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersPreviousState {
	s.State = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext struct {
	Capability             *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                                      `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                                     `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetCapability(v *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext) SetRunAsUser(v int64) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability) SetAdds(v []*string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetMountPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetMountPropagation(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts) SetReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsInitContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsTags) SetKey(v string) *DescribeContainerGroupsResponseBodyContainerGroupsTags {
	s.Key = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsTags) SetValue(v string) *DescribeContainerGroupsResponseBodyContainerGroupsTags {
	s.Value = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsVolumes struct {
	ConfigFileVolumeConfigFileToPaths []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	DiskVolumeDiskId                  *string                                                                                       `json:"DiskVolumeDiskId,omitempty" xml:"DiskVolumeDiskId,omitempty"`
	DiskVolumeFsType                  *string                                                                                       `json:"DiskVolumeFsType,omitempty" xml:"DiskVolumeFsType,omitempty"`
	FlexVolumeDriver                  *string                                                                                       `json:"FlexVolumeDriver,omitempty" xml:"FlexVolumeDriver,omitempty"`
	FlexVolumeFsType                  *string                                                                                       `json:"FlexVolumeFsType,omitempty" xml:"FlexVolumeFsType,omitempty"`
	FlexVolumeOptions                 *string                                                                                       `json:"FlexVolumeOptions,omitempty" xml:"FlexVolumeOptions,omitempty"`
	NFSVolumePath                     *string                                                                                       `json:"NFSVolumePath,omitempty" xml:"NFSVolumePath,omitempty"`
	NFSVolumeReadOnly                 *bool                                                                                         `json:"NFSVolumeReadOnly,omitempty" xml:"NFSVolumeReadOnly,omitempty"`
	NFSVolumeServer                   *string                                                                                       `json:"NFSVolumeServer,omitempty" xml:"NFSVolumeServer,omitempty"`
	Name                              *string                                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Type                              *string                                                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumes) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumes) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetConfigFileVolumeConfigFileToPaths(v []*DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetDiskVolumeDiskId(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.DiskVolumeDiskId = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetDiskVolumeFsType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.DiskVolumeFsType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeDriver(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeDriver = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeFsType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeFsType = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetFlexVolumeOptions(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.FlexVolumeOptions = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumePath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumePath = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumeReadOnly(v bool) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumeReadOnly = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetNFSVolumeServer(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.NFSVolumeServer = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetName(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.Name = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumes) SetType(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumes {
	s.Type = &v
	return s
}

type DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *DescribeContainerGroupsResponseBodyContainerGroupsVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

type DescribeContainerGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerGroupsResponse) SetHeaders(v map[string]*string) *DescribeContainerGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerGroupsResponse) SetStatusCode(v int32) *DescribeContainerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerGroupsResponse) SetBody(v *DescribeContainerGroupsResponseBody) *DescribeContainerGroupsResponse {
	s.Body = v
	return s
}

type DescribeContainerLogRequest struct {
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ContainerName        *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	LastTime             *bool   `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	LimitBytes           *int64  `json:"LimitBytes,omitempty" xml:"LimitBytes,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SinceSeconds         *int32  `json:"SinceSeconds,omitempty" xml:"SinceSeconds,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tail                 *int32  `json:"Tail,omitempty" xml:"Tail,omitempty"`
	Timestamps           *bool   `json:"Timestamps,omitempty" xml:"Timestamps,omitempty"`
}

func (s DescribeContainerLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogRequest) SetContainerGroupId(v string) *DescribeContainerLogRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetContainerName(v string) *DescribeContainerLogRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeContainerLogRequest) SetLastTime(v bool) *DescribeContainerLogRequest {
	s.LastTime = &v
	return s
}

func (s *DescribeContainerLogRequest) SetLimitBytes(v int64) *DescribeContainerLogRequest {
	s.LimitBytes = &v
	return s
}

func (s *DescribeContainerLogRequest) SetOwnerAccount(v string) *DescribeContainerLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeContainerLogRequest) SetOwnerId(v int64) *DescribeContainerLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetRegionId(v string) *DescribeContainerLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetResourceOwnerAccount(v string) *DescribeContainerLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeContainerLogRequest) SetResourceOwnerId(v int64) *DescribeContainerLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeContainerLogRequest) SetSinceSeconds(v int32) *DescribeContainerLogRequest {
	s.SinceSeconds = &v
	return s
}

func (s *DescribeContainerLogRequest) SetStartTime(v string) *DescribeContainerLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeContainerLogRequest) SetTail(v int32) *DescribeContainerLogRequest {
	s.Tail = &v
	return s
}

func (s *DescribeContainerLogRequest) SetTimestamps(v bool) *DescribeContainerLogRequest {
	s.Timestamps = &v
	return s
}

type DescribeContainerLogResponseBody struct {
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogResponseBody) SetContainerName(v string) *DescribeContainerLogResponseBody {
	s.ContainerName = &v
	return s
}

func (s *DescribeContainerLogResponseBody) SetContent(v string) *DescribeContainerLogResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeContainerLogResponseBody) SetRequestId(v string) *DescribeContainerLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContainerLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerLogResponse) SetHeaders(v map[string]*string) *DescribeContainerLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerLogResponse) SetStatusCode(v int32) *DescribeContainerLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerLogResponse) SetBody(v *DescribeContainerLogResponseBody) *DescribeContainerLogResponse {
	s.Body = v
	return s
}

type DescribeImageCachesRequest struct {
	Image                *string                          `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageCacheId         *string                          `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	ImageCacheName       *string                          `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	MatchImage           []*string                        `json:"MatchImage,omitempty" xml:"MatchImage,omitempty" type:"Repeated"`
	OwnerAccount         *string                          `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SnapshotId           *string                          `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Tag                  []*DescribeImageCachesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImageCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesRequest) SetImage(v string) *DescribeImageCachesRequest {
	s.Image = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageCacheId(v string) *DescribeImageCachesRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetImageCacheName(v string) *DescribeImageCachesRequest {
	s.ImageCacheName = &v
	return s
}

func (s *DescribeImageCachesRequest) SetMatchImage(v []*string) *DescribeImageCachesRequest {
	s.MatchImage = v
	return s
}

func (s *DescribeImageCachesRequest) SetOwnerAccount(v string) *DescribeImageCachesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImageCachesRequest) SetOwnerId(v int64) *DescribeImageCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetRegionId(v string) *DescribeImageCachesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceGroupId(v string) *DescribeImageCachesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceOwnerAccount(v string) *DescribeImageCachesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageCachesRequest) SetResourceOwnerId(v int64) *DescribeImageCachesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetSnapshotId(v string) *DescribeImageCachesRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageCachesRequest) SetTag(v []*DescribeImageCachesRequestTag) *DescribeImageCachesRequest {
	s.Tag = v
	return s
}

type DescribeImageCachesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCachesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesRequestTag) SetKey(v string) *DescribeImageCachesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImageCachesRequestTag) SetValue(v string) *DescribeImageCachesRequestTag {
	s.Value = &v
	return s
}

type DescribeImageCachesResponseBody struct {
	ImageCaches []*DescribeImageCachesResponseBodyImageCaches `json:"ImageCaches,omitempty" xml:"ImageCaches,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBody) SetImageCaches(v []*DescribeImageCachesResponseBodyImageCaches) *DescribeImageCachesResponseBody {
	s.ImageCaches = v
	return s
}

func (s *DescribeImageCachesResponseBody) SetRequestId(v string) *DescribeImageCachesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeImageCachesResponseBodyImageCaches struct {
	ContainerGroupId    *string                                             `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	CreationTime        *string                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EliminationStrategy *string                                             `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	Events              []*DescribeImageCachesResponseBodyImageCachesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	ExpireDateTime      *string                                             `json:"ExpireDateTime,omitempty" xml:"ExpireDateTime,omitempty"`
	FlashSnapshotId     *string                                             `json:"FlashSnapshotId,omitempty" xml:"FlashSnapshotId,omitempty"`
	ImageCacheId        *string                                             `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	ImageCacheName      *string                                             `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	ImageCacheSize      *int32                                              `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	Images              []*string                                           `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LastMatchedTime     *string                                             `json:"LastMatchedTime,omitempty" xml:"LastMatchedTime,omitempty"`
	Progress            *string                                             `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionId            *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId     *string                                             `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SnapshotId          *string                                             `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	Status              *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                []*DescribeImageCachesResponseBodyImageCachesTags   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeImageCachesResponseBodyImageCaches) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCaches) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetContainerGroupId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetCreationTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.CreationTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetEliminationStrategy(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.EliminationStrategy = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetEvents(v []*DescribeImageCachesResponseBodyImageCachesEvents) *DescribeImageCachesResponseBodyImageCaches {
	s.Events = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetExpireDateTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ExpireDateTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetFlashSnapshotId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.FlashSnapshotId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheName(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheName = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImageCacheSize(v int32) *DescribeImageCachesResponseBodyImageCaches {
	s.ImageCacheSize = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetImages(v []*string) *DescribeImageCachesResponseBodyImageCaches {
	s.Images = v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetLastMatchedTime(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.LastMatchedTime = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetProgress(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.Progress = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetRegionId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.RegionId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetResourceGroupId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetSnapshotId(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.SnapshotId = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetStatus(v string) *DescribeImageCachesResponseBodyImageCaches {
	s.Status = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCaches) SetTags(v []*DescribeImageCachesResponseBodyImageCachesTags) *DescribeImageCachesResponseBodyImageCaches {
	s.Tags = v
	return s
}

type DescribeImageCachesResponseBodyImageCachesEvents struct {
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCachesEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCachesEvents) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetCount(v int32) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Count = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetFirstTimestamp(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetLastTimestamp(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetMessage(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Message = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetName(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Name = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesEvents) SetType(v string) *DescribeImageCachesResponseBodyImageCachesEvents {
	s.Type = &v
	return s
}

type DescribeImageCachesResponseBodyImageCachesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageCachesResponseBodyImageCachesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponseBodyImageCachesTags) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponseBodyImageCachesTags) SetKey(v string) *DescribeImageCachesResponseBodyImageCachesTags {
	s.Key = &v
	return s
}

func (s *DescribeImageCachesResponseBodyImageCachesTags) SetValue(v string) *DescribeImageCachesResponseBodyImageCachesTags {
	s.Value = &v
	return s
}

type DescribeImageCachesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeImageCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageCachesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageCachesResponse) SetHeaders(v map[string]*string) *DescribeImageCachesResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageCachesResponse) SetStatusCode(v int32) *DescribeImageCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageCachesResponse) SetBody(v *DescribeImageCachesResponseBody) *DescribeImageCachesResponse {
	s.Body = v
	return s
}

type DescribeInstanceOpsRecordsRequest struct {
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OpsType              *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceOpsRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsRequest) SetContainerGroupId(v string) *DescribeInstanceOpsRecordsRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetOpsType(v string) *DescribeInstanceOpsRecordsRequest {
	s.OpsType = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetOwnerAccount(v string) *DescribeInstanceOpsRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetOwnerId(v int64) *DescribeInstanceOpsRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetRegionId(v string) *DescribeInstanceOpsRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetResourceOwnerAccount(v string) *DescribeInstanceOpsRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceOpsRecordsRequest) SetResourceOwnerId(v int64) *DescribeInstanceOpsRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeInstanceOpsRecordsResponseBody struct {
	Records   []*DescribeInstanceOpsRecordsResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceOpsRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsResponseBody) SetRecords(v []*DescribeInstanceOpsRecordsResponseBodyRecords) *DescribeInstanceOpsRecordsResponseBody {
	s.Records = v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBody) SetRequestId(v string) *DescribeInstanceOpsRecordsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceOpsRecordsResponseBodyRecords struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime    *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	OpsStatus     *string `json:"OpsStatus,omitempty" xml:"OpsStatus,omitempty"`
	OpsType       *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	ResultType    *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
}

func (s DescribeInstanceOpsRecordsResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetCreateTime(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetExpireTime(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetOpsStatus(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.OpsStatus = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetOpsType(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.OpsType = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetResultContent(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.ResultContent = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponseBodyRecords) SetResultType(v string) *DescribeInstanceOpsRecordsResponseBodyRecords {
	s.ResultType = &v
	return s
}

type DescribeInstanceOpsRecordsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceOpsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceOpsRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceOpsRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceOpsRecordsResponse) SetHeaders(v map[string]*string) *DescribeInstanceOpsRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceOpsRecordsResponse) SetStatusCode(v int32) *DescribeInstanceOpsRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceOpsRecordsResponse) SetBody(v *DescribeInstanceOpsRecordsResponseBody) *DescribeInstanceOpsRecordsResponse {
	s.Body = v
	return s
}

type DescribeMultiContainerGroupMetricRequest struct {
	ContainerGroupIds    *string `json:"ContainerGroupIds,omitempty" xml:"ContainerGroupIds,omitempty"`
	MetricType           *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeMultiContainerGroupMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricRequest) SetContainerGroupIds(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ContainerGroupIds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetMetricType(v string) *DescribeMultiContainerGroupMetricRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetOwnerAccount(v string) *DescribeMultiContainerGroupMetricRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetOwnerId(v int64) *DescribeMultiContainerGroupMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetRegionId(v string) *DescribeMultiContainerGroupMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceGroupId(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceOwnerAccount(v string) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricRequest) SetResourceOwnerId(v int64) *DescribeMultiContainerGroupMetricRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBody struct {
	MonitorDatas []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas `json:"MonitorDatas,omitempty" xml:"MonitorDatas,omitempty" type:"Repeated"`
	RequestId    *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBody) SetMonitorDatas(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) *DescribeMultiContainerGroupMetricResponseBody {
	s.MonitorDatas = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBody) SetRequestId(v string) *DescribeMultiContainerGroupMetricResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatas struct {
	ContainerGroupId *string                                                             `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Records          []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetContainerGroupId(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.ContainerGroupId = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas) SetRecords(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatas {
	s.Records = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords struct {
	CPU        *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU          `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Containers []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	Disk       []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk       `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	Filesystem []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem `json:"Filesystem,omitempty" xml:"Filesystem,omitempty" type:"Repeated"`
	Memory     *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory       `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Network    *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork      `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	Timestamp  *string                                                                       `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetCPU(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.CPU = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetContainers(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Containers = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetDisk(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Disk = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetFilesystem(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Filesystem = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Memory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetNetwork(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Network = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords) SetTimestamp(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecords {
	s.Timestamp = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU struct {
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetLoad(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.Load = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetUsageCoreNanoSeconds(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU) SetUsageNanoCores(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers struct {
	CPU    *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU    `json:"CPU,omitempty" xml:"CPU,omitempty" type:"Struct"`
	Memory *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Name   *string                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetCPU(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.CPU = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetMemory(v *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.Memory = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainers {
	s.Name = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU struct {
	Limit                *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Load                 *int64 `json:"Load,omitempty" xml:"Load,omitempty"`
	UsageCoreNanoSeconds *int64 `json:"UsageCoreNanoSeconds,omitempty" xml:"UsageCoreNanoSeconds,omitempty"`
	UsageNanoCores       *int64 `json:"UsageNanoCores,omitempty" xml:"UsageNanoCores,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetLimit(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.Limit = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetLoad(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.Load = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetUsageCoreNanoSeconds(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.UsageCoreNanoSeconds = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU) SetUsageNanoCores(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersCPU {
	s.UsageNanoCores = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory struct {
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetAvailableBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.Cache = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetUsageBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsContainersMemory {
	s.WorkingSet = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk struct {
	Device     *string `json:"Device,omitempty" xml:"Device,omitempty"`
	ReadBytes  *int64  `json:"ReadBytes,omitempty" xml:"ReadBytes,omitempty"`
	ReadIo     *int64  `json:"ReadIo,omitempty" xml:"ReadIo,omitempty"`
	WriteBytes *int64  `json:"WriteBytes,omitempty" xml:"WriteBytes,omitempty"`
	WriteIo    *int64  `json:"WriteIo,omitempty" xml:"WriteIo,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetDevice(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.Device = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetReadBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.ReadBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetReadIo(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.ReadIo = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetWriteBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.WriteBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk) SetWriteIo(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsDisk {
	s.WriteIo = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem struct {
	Available *int64  `json:"Available,omitempty" xml:"Available,omitempty"`
	Capacity  *int64  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	FsName    *string `json:"FsName,omitempty" xml:"FsName,omitempty"`
	Usage     *int64  `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetAvailable(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Available = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetCapacity(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Capacity = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetFsName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.FsName = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem) SetUsage(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsFilesystem {
	s.Usage = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory struct {
	AvailableBytes *int64 `json:"AvailableBytes,omitempty" xml:"AvailableBytes,omitempty"`
	Cache          *int64 `json:"Cache,omitempty" xml:"Cache,omitempty"`
	Rss            *int64 `json:"Rss,omitempty" xml:"Rss,omitempty"`
	UsageBytes     *int64 `json:"UsageBytes,omitempty" xml:"UsageBytes,omitempty"`
	WorkingSet     *int64 `json:"WorkingSet,omitempty" xml:"WorkingSet,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetAvailableBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.AvailableBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetCache(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.Cache = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetRss(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.Rss = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetUsageBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.UsageBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory) SetWorkingSet(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsMemory {
	s.WorkingSet = &v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork struct {
	Interfaces []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork) SetInterfaces(v []*DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetwork {
	s.Interfaces = v
	return s
}

type DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RxBytes   *int64  `json:"RxBytes,omitempty" xml:"RxBytes,omitempty"`
	RxDrops   *int64  `json:"RxDrops,omitempty" xml:"RxDrops,omitempty"`
	RxErrors  *int64  `json:"RxErrors,omitempty" xml:"RxErrors,omitempty"`
	RxPackets *int64  `json:"RxPackets,omitempty" xml:"RxPackets,omitempty"`
	TxBytes   *int64  `json:"TxBytes,omitempty" xml:"TxBytes,omitempty"`
	TxDrops   *int64  `json:"TxDrops,omitempty" xml:"TxDrops,omitempty"`
	TxErrors  *int64  `json:"TxErrors,omitempty" xml:"TxErrors,omitempty"`
	TxPackets *int64  `json:"TxPackets,omitempty" xml:"TxPackets,omitempty"`
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetName(v string) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.Name = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxDrops(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxDrops = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetRxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.RxPackets = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxBytes(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxBytes = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxDrops(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxDrops = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxErrors(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxErrors = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces) SetTxPackets(v int64) *DescribeMultiContainerGroupMetricResponseBodyMonitorDatasRecordsNetworkInterfaces {
	s.TxPackets = &v
	return s
}

type DescribeMultiContainerGroupMetricResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMultiContainerGroupMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMultiContainerGroupMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMultiContainerGroupMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiContainerGroupMetricResponse) SetHeaders(v map[string]*string) *DescribeMultiContainerGroupMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponse) SetStatusCode(v int32) *DescribeMultiContainerGroupMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiContainerGroupMetricResponse) SetBody(v *DescribeMultiContainerGroupMetricResponseBody) *DescribeMultiContainerGroupMetricResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RecommendZones []*string `json:"RecommendZones,omitempty" xml:"RecommendZones,omitempty" type:"Repeated"`
	RegionEndpoint *string   `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones          []*string `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRecommendZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.RecommendZones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*string) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeVirtualNodesRequest struct {
	ClientToken          *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Limit                *int64                            `json:"Limit,omitempty" xml:"Limit,omitempty"`
	NextToken            *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityGroupId      *string                           `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status               *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag                  []*DescribeVirtualNodesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId            *string                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VirtualNodeIds       *string                           `json:"VirtualNodeIds,omitempty" xml:"VirtualNodeIds,omitempty"`
	VirtualNodeName      *string                           `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	ZoneId               *string                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVirtualNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesRequest) SetClientToken(v string) *DescribeVirtualNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetLimit(v int64) *DescribeVirtualNodesRequest {
	s.Limit = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetNextToken(v string) *DescribeVirtualNodesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetOwnerAccount(v string) *DescribeVirtualNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetOwnerId(v int64) *DescribeVirtualNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetRegionId(v string) *DescribeVirtualNodesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetResourceGroupId(v string) *DescribeVirtualNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetResourceOwnerAccount(v string) *DescribeVirtualNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetResourceOwnerId(v int64) *DescribeVirtualNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetSecurityGroupId(v string) *DescribeVirtualNodesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetStatus(v string) *DescribeVirtualNodesRequest {
	s.Status = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetTag(v []*DescribeVirtualNodesRequestTag) *DescribeVirtualNodesRequest {
	s.Tag = v
	return s
}

func (s *DescribeVirtualNodesRequest) SetVSwitchId(v string) *DescribeVirtualNodesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetVirtualNodeIds(v string) *DescribeVirtualNodesRequest {
	s.VirtualNodeIds = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetVirtualNodeName(v string) *DescribeVirtualNodesRequest {
	s.VirtualNodeName = &v
	return s
}

func (s *DescribeVirtualNodesRequest) SetZoneId(v string) *DescribeVirtualNodesRequest {
	s.ZoneId = &v
	return s
}

type DescribeVirtualNodesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVirtualNodesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesRequestTag) SetKey(v string) *DescribeVirtualNodesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVirtualNodesRequestTag) SetValue(v string) *DescribeVirtualNodesRequestTag {
	s.Value = &v
	return s
}

type DescribeVirtualNodesResponseBody struct {
	NextToken    *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VirtualNodes []*DescribeVirtualNodesResponseBodyVirtualNodes `json:"VirtualNodes,omitempty" xml:"VirtualNodes,omitempty" type:"Repeated"`
}

func (s DescribeVirtualNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBody) SetNextToken(v string) *DescribeVirtualNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVirtualNodesResponseBody) SetRequestId(v string) *DescribeVirtualNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBody) SetTotalCount(v int32) *DescribeVirtualNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVirtualNodesResponseBody) SetVirtualNodes(v []*DescribeVirtualNodesResponseBodyVirtualNodes) *DescribeVirtualNodesResponseBody {
	s.VirtualNodes = v
	return s
}

type DescribeVirtualNodesResponseBodyVirtualNodes struct {
	ClusterId       *string                                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Cpu             *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime    *string                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	EniInstanceId   *string                                               `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	Events          []*DescribeVirtualNodesResponseBodyVirtualNodesEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	InternetIp      *string                                               `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	IntranetIp      *string                                               `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	Memory          *float32                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	RamRoleName     *string                                               `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	RegionId        *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupId *string                                               `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status          *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*DescribeVirtualNodesResponseBodyVirtualNodesTags   `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VSwitchId       *string                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VirtualNodeId   *string                                               `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
	VirtualNodeName *string                                               `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
	VpcId           *string                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId          *string                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodes) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetClusterId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ClusterId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetCpu(v float32) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Cpu = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetCreationTime(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetEniInstanceId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetEvents(v []*DescribeVirtualNodesResponseBodyVirtualNodesEvents) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Events = v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetInternetIp(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.InternetIp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetIntranetIp(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetMemory(v float32) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Memory = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetRamRoleName(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.RamRoleName = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetRegionId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetResourceGroupId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetSecurityGroupId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetStatus(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Status = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetTags(v []*DescribeVirtualNodesResponseBodyVirtualNodesTags) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.Tags = v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVSwitchId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVirtualNodeId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VirtualNodeId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVirtualNodeName(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VirtualNodeName = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetVpcId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.VpcId = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodes) SetZoneId(v string) *DescribeVirtualNodesResponseBodyVirtualNodes {
	s.ZoneId = &v
	return s
}

type DescribeVirtualNodesResponseBodyVirtualNodesEvents struct {
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	FirstTimestamp *string `json:"FirstTimestamp,omitempty" xml:"FirstTimestamp,omitempty"`
	LastTimestamp  *string `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesEvents) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetCount(v int32) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Count = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetFirstTimestamp(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.FirstTimestamp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetLastTimestamp(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.LastTimestamp = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetMessage(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Message = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetName(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Name = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetReason(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Reason = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesEvents) SetType(v string) *DescribeVirtualNodesResponseBodyVirtualNodesEvents {
	s.Type = &v
	return s
}

type DescribeVirtualNodesResponseBodyVirtualNodesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponseBodyVirtualNodesTags) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesTags) SetKey(v string) *DescribeVirtualNodesResponseBodyVirtualNodesTags {
	s.Key = &v
	return s
}

func (s *DescribeVirtualNodesResponseBodyVirtualNodesTags) SetValue(v string) *DescribeVirtualNodesResponseBodyVirtualNodesTags {
	s.Value = &v
	return s
}

type DescribeVirtualNodesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVirtualNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVirtualNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVirtualNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVirtualNodesResponse) SetHeaders(v map[string]*string) *DescribeVirtualNodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVirtualNodesResponse) SetStatusCode(v int32) *DescribeVirtualNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVirtualNodesResponse) SetBody(v *DescribeVirtualNodesResponseBody) *DescribeVirtualNodesResponse {
	s.Body = v
	return s
}

type ExecContainerCommandRequest struct {
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	ContainerName        *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Stdin                *bool   `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	Sync                 *bool   `json:"Sync,omitempty" xml:"Sync,omitempty"`
	TTY                  *bool   `json:"TTY,omitempty" xml:"TTY,omitempty"`
}

func (s ExecContainerCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandRequest) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandRequest) SetCommand(v string) *ExecContainerCommandRequest {
	s.Command = &v
	return s
}

func (s *ExecContainerCommandRequest) SetContainerGroupId(v string) *ExecContainerCommandRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetContainerName(v string) *ExecContainerCommandRequest {
	s.ContainerName = &v
	return s
}

func (s *ExecContainerCommandRequest) SetOwnerAccount(v string) *ExecContainerCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ExecContainerCommandRequest) SetOwnerId(v int64) *ExecContainerCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetRegionId(v string) *ExecContainerCommandRequest {
	s.RegionId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetResourceOwnerAccount(v string) *ExecContainerCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExecContainerCommandRequest) SetResourceOwnerId(v int64) *ExecContainerCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExecContainerCommandRequest) SetStdin(v bool) *ExecContainerCommandRequest {
	s.Stdin = &v
	return s
}

func (s *ExecContainerCommandRequest) SetSync(v bool) *ExecContainerCommandRequest {
	s.Sync = &v
	return s
}

func (s *ExecContainerCommandRequest) SetTTY(v bool) *ExecContainerCommandRequest {
	s.TTY = &v
	return s
}

type ExecContainerCommandResponseBody struct {
	HttpUrl      *string `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SyncResponse *string `json:"SyncResponse,omitempty" xml:"SyncResponse,omitempty"`
	WebSocketUri *string `json:"WebSocketUri,omitempty" xml:"WebSocketUri,omitempty"`
}

func (s ExecContainerCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandResponseBody) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandResponseBody) SetHttpUrl(v string) *ExecContainerCommandResponseBody {
	s.HttpUrl = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetRequestId(v string) *ExecContainerCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetSyncResponse(v string) *ExecContainerCommandResponseBody {
	s.SyncResponse = &v
	return s
}

func (s *ExecContainerCommandResponseBody) SetWebSocketUri(v string) *ExecContainerCommandResponseBody {
	s.WebSocketUri = &v
	return s
}

type ExecContainerCommandResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecContainerCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecContainerCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecContainerCommandResponse) GoString() string {
	return s.String()
}

func (s *ExecContainerCommandResponse) SetHeaders(v map[string]*string) *ExecContainerCommandResponse {
	s.Headers = v
	return s
}

func (s *ExecContainerCommandResponse) SetStatusCode(v int32) *ExecContainerCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecContainerCommandResponse) SetBody(v *ExecContainerCommandResponseBody) *ExecContainerCommandResponse {
	s.Body = v
	return s
}

type ListUsageRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsageRequest) GoString() string {
	return s.String()
}

func (s *ListUsageRequest) SetOwnerAccount(v string) *ListUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListUsageRequest) SetOwnerId(v int64) *ListUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *ListUsageRequest) SetRegionId(v string) *ListUsageRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsageRequest) SetResourceOwnerAccount(v string) *ListUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListUsageRequest) SetResourceOwnerId(v int64) *ListUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListUsageResponseBody struct {
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsageResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsageResponseBody) SetAttributes(v map[string]interface{}) *ListUsageResponseBody {
	s.Attributes = v
	return s
}

func (s *ListUsageResponseBody) SetRequestId(v string) *ListUsageResponseBody {
	s.RequestId = &v
	return s
}

type ListUsageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsageResponse) GoString() string {
	return s.String()
}

func (s *ListUsageResponse) SetHeaders(v map[string]*string) *ListUsageResponse {
	s.Headers = v
	return s
}

func (s *ListUsageResponse) SetStatusCode(v int32) *ListUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsageResponse) SetBody(v *ListUsageResponseBody) *ListUsageResponse {
	s.Body = v
	return s
}

type ResizeContainerGroupVolumeRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	NewSize              *int64  `json:"NewSize,omitempty" xml:"NewSize,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VolumeName           *string `json:"VolumeName,omitempty" xml:"VolumeName,omitempty"`
}

func (s ResizeContainerGroupVolumeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeContainerGroupVolumeRequest) GoString() string {
	return s.String()
}

func (s *ResizeContainerGroupVolumeRequest) SetClientToken(v string) *ResizeContainerGroupVolumeRequest {
	s.ClientToken = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetContainerGroupId(v string) *ResizeContainerGroupVolumeRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetNewSize(v int64) *ResizeContainerGroupVolumeRequest {
	s.NewSize = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetOwnerAccount(v string) *ResizeContainerGroupVolumeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetOwnerId(v int64) *ResizeContainerGroupVolumeRequest {
	s.OwnerId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetRegionId(v string) *ResizeContainerGroupVolumeRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetResourceOwnerAccount(v string) *ResizeContainerGroupVolumeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetResourceOwnerId(v int64) *ResizeContainerGroupVolumeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResizeContainerGroupVolumeRequest) SetVolumeName(v string) *ResizeContainerGroupVolumeRequest {
	s.VolumeName = &v
	return s
}

type ResizeContainerGroupVolumeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeContainerGroupVolumeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeContainerGroupVolumeResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeContainerGroupVolumeResponseBody) SetRequestId(v string) *ResizeContainerGroupVolumeResponseBody {
	s.RequestId = &v
	return s
}

type ResizeContainerGroupVolumeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResizeContainerGroupVolumeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeContainerGroupVolumeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeContainerGroupVolumeResponse) GoString() string {
	return s.String()
}

func (s *ResizeContainerGroupVolumeResponse) SetHeaders(v map[string]*string) *ResizeContainerGroupVolumeResponse {
	s.Headers = v
	return s
}

func (s *ResizeContainerGroupVolumeResponse) SetStatusCode(v int32) *ResizeContainerGroupVolumeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResizeContainerGroupVolumeResponse) SetBody(v *ResizeContainerGroupVolumeResponseBody) *ResizeContainerGroupVolumeResponse {
	s.Body = v
	return s
}

type RestartContainerGroupRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ContainerGroupId     *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RestartContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupRequest) SetClientToken(v string) *RestartContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartContainerGroupRequest) SetContainerGroupId(v string) *RestartContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetOwnerAccount(v string) *RestartContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartContainerGroupRequest) SetOwnerId(v int64) *RestartContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetRegionId(v string) *RestartContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RestartContainerGroupRequest) SetResourceOwnerAccount(v string) *RestartContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartContainerGroupRequest) SetResourceOwnerId(v int64) *RestartContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

type RestartContainerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupResponseBody) SetRequestId(v string) *RestartContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type RestartContainerGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *RestartContainerGroupResponse) SetHeaders(v map[string]*string) *RestartContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *RestartContainerGroupResponse) SetStatusCode(v int32) *RestartContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartContainerGroupResponse) SetBody(v *RestartContainerGroupResponseBody) *RestartContainerGroupResponse {
	s.Body = v
	return s
}

type UpdateContainerGroupRequest struct {
	DnsConfig               *UpdateContainerGroupRequestDnsConfig                 `json:"DnsConfig,omitempty" xml:"DnsConfig,omitempty" type:"Struct"`
	AcrRegistryInfo         []*UpdateContainerGroupRequestAcrRegistryInfo         `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	ClientToken             *string                                               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Container               []*UpdateContainerGroupRequestContainer               `json:"Container,omitempty" xml:"Container,omitempty" type:"Repeated"`
	ContainerGroupId        *string                                               `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	Cpu                     *float32                                              `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ImageRegistryCredential []*UpdateContainerGroupRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	InitContainer           []*UpdateContainerGroupRequestInitContainer           `json:"InitContainer,omitempty" xml:"InitContainer,omitempty" type:"Repeated"`
	Memory                  *float32                                              `json:"Memory,omitempty" xml:"Memory,omitempty"`
	OwnerAccount            *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId         *string                                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount    *string                                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RestartPolicy           *string                                               `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	Tag                     []*UpdateContainerGroupRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UpdateType              *string                                               `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
	Volume                  []*UpdateContainerGroupRequestVolume                  `json:"Volume,omitempty" xml:"Volume,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequest) SetDnsConfig(v *UpdateContainerGroupRequestDnsConfig) *UpdateContainerGroupRequest {
	s.DnsConfig = v
	return s
}

func (s *UpdateContainerGroupRequest) SetAcrRegistryInfo(v []*UpdateContainerGroupRequestAcrRegistryInfo) *UpdateContainerGroupRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *UpdateContainerGroupRequest) SetClientToken(v string) *UpdateContainerGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetContainer(v []*UpdateContainerGroupRequestContainer) *UpdateContainerGroupRequest {
	s.Container = v
	return s
}

func (s *UpdateContainerGroupRequest) SetContainerGroupId(v string) *UpdateContainerGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetCpu(v float32) *UpdateContainerGroupRequest {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetImageRegistryCredential(v []*UpdateContainerGroupRequestImageRegistryCredential) *UpdateContainerGroupRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *UpdateContainerGroupRequest) SetInitContainer(v []*UpdateContainerGroupRequestInitContainer) *UpdateContainerGroupRequest {
	s.InitContainer = v
	return s
}

func (s *UpdateContainerGroupRequest) SetMemory(v float32) *UpdateContainerGroupRequest {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetOwnerAccount(v string) *UpdateContainerGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetOwnerId(v int64) *UpdateContainerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetRegionId(v string) *UpdateContainerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceGroupId(v string) *UpdateContainerGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceOwnerAccount(v string) *UpdateContainerGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetResourceOwnerId(v int64) *UpdateContainerGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetRestartPolicy(v string) *UpdateContainerGroupRequest {
	s.RestartPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetTag(v []*UpdateContainerGroupRequestTag) *UpdateContainerGroupRequest {
	s.Tag = v
	return s
}

func (s *UpdateContainerGroupRequest) SetUpdateType(v string) *UpdateContainerGroupRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateContainerGroupRequest) SetVolume(v []*UpdateContainerGroupRequestVolume) *UpdateContainerGroupRequest {
	s.Volume = v
	return s
}

type UpdateContainerGroupRequestDnsConfig struct {
	NameServer []*string                                     `json:"NameServer,omitempty" xml:"NameServer,omitempty" type:"Repeated"`
	Option     []*UpdateContainerGroupRequestDnsConfigOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	Search     []*string                                     `json:"Search,omitempty" xml:"Search,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestDnsConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestDnsConfig) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestDnsConfig) SetNameServer(v []*string) *UpdateContainerGroupRequestDnsConfig {
	s.NameServer = v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfig) SetOption(v []*UpdateContainerGroupRequestDnsConfigOption) *UpdateContainerGroupRequestDnsConfig {
	s.Option = v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfig) SetSearch(v []*string) *UpdateContainerGroupRequestDnsConfig {
	s.Search = v
	return s
}

type UpdateContainerGroupRequestDnsConfigOption struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestDnsConfigOption) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestDnsConfigOption) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestDnsConfigOption) SetName(v string) *UpdateContainerGroupRequestDnsConfigOption {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestDnsConfigOption) SetValue(v string) *UpdateContainerGroupRequestDnsConfigOption {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestAcrRegistryInfo struct {
	Domain       []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateContainerGroupRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetDomain(v []*string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetInstanceId(v string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetInstanceName(v string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *UpdateContainerGroupRequestAcrRegistryInfo) SetRegionId(v string) *UpdateContainerGroupRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type UpdateContainerGroupRequestContainer struct {
	LivenessProbe                               *UpdateContainerGroupRequestContainerLivenessProbe                                 `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" require:"true" type:"Struct"`
	ReadinessProbe                              *UpdateContainerGroupRequestContainerReadinessProbe                                `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" require:"true" type:"Struct"`
	SecurityContext                             *UpdateContainerGroupRequestContainerSecurityContext                               `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Arg                                         []*string                                                                          `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command                                     []*string                                                                          `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Cpu                                         *float32                                                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVar                              []*UpdateContainerGroupRequestContainerEnvironmentVar                              `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	Gpu                                         *int32                                                                             `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image                                       *string                                                                            `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy                             *string                                                                            `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	LifecyclePostStartHandlerExec               []*string                                                                          `json:"LifecyclePostStartHandlerExec,omitempty" xml:"LifecyclePostStartHandlerExec,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerHttpGetHost        *string                                                                            `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	LifecyclePostStartHandlerHttpGetHttpHeaders []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders `json:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" xml:"LifecyclePostStartHandlerHttpGetHttpHeaders,omitempty" type:"Repeated"`
	LifecyclePostStartHandlerHttpGetPath        *string                                                                            `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	LifecyclePostStartHandlerHttpGetPort        *int32                                                                             `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	LifecyclePostStartHandlerHttpGetScheme      *string                                                                            `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	LifecyclePostStartHandlerTcpSocketHost      *string                                                                            `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	LifecyclePostStartHandlerTcpSocketPort      *int32                                                                             `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	LifecyclePreStopHandlerExec                 []*string                                                                          `json:"LifecyclePreStopHandlerExec,omitempty" xml:"LifecyclePreStopHandlerExec,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetHost          *string                                                                            `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	LifecyclePreStopHandlerHttpGetHttpHeader    []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader    `json:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" xml:"LifecyclePreStopHandlerHttpGetHttpHeader,omitempty" type:"Repeated"`
	LifecyclePreStopHandlerHttpGetPath          *string                                                                            `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	LifecyclePreStopHandlerHttpGetPort          *int32                                                                             `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	LifecyclePreStopHandlerHttpGetScheme        *string                                                                            `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	LifecyclePreStopHandlerTcpSocketHost        *string                                                                            `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	LifecyclePreStopHandlerTcpSocketPort        *int32                                                                             `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	Memory                                      *float32                                                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name                                        *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Port                                        []*UpdateContainerGroupRequestContainerPort                                        `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	Stdin                                       *bool                                                                              `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce                                   *bool                                                                              `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Tty                                         *bool                                                                              `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMount                                 []*UpdateContainerGroupRequestContainerVolumeMount                                 `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	WorkingDir                                  *string                                                                            `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s UpdateContainerGroupRequestContainer) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainer) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainer) SetLivenessProbe(v *UpdateContainerGroupRequestContainerLivenessProbe) *UpdateContainerGroupRequestContainer {
	s.LivenessProbe = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetReadinessProbe(v *UpdateContainerGroupRequestContainerReadinessProbe) *UpdateContainerGroupRequestContainer {
	s.ReadinessProbe = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetSecurityContext(v *UpdateContainerGroupRequestContainerSecurityContext) *UpdateContainerGroupRequestContainer {
	s.SecurityContext = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetArg(v []*string) *UpdateContainerGroupRequestContainer {
	s.Arg = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetCommand(v []*string) *UpdateContainerGroupRequestContainer {
	s.Command = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetCpu(v float32) *UpdateContainerGroupRequestContainer {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetEnvironmentVar(v []*UpdateContainerGroupRequestContainerEnvironmentVar) *UpdateContainerGroupRequestContainer {
	s.EnvironmentVar = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetGpu(v int32) *UpdateContainerGroupRequestContainer {
	s.Gpu = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetImage(v string) *UpdateContainerGroupRequestContainer {
	s.Image = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetImagePullPolicy(v string) *UpdateContainerGroupRequestContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerExec(v []*string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerExec = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetHttpHeaders(v []*UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetHttpHeaders = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPath(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerHttpGetScheme(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerExec(v []*string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerExec = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetHttpHeader(v []*UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetHttpHeader = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPath(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerHttpGetScheme(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketHost(v string) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *UpdateContainerGroupRequestContainer {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetMemory(v float32) *UpdateContainerGroupRequestContainer {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetName(v string) *UpdateContainerGroupRequestContainer {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetPort(v []*UpdateContainerGroupRequestContainerPort) *UpdateContainerGroupRequestContainer {
	s.Port = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetStdin(v bool) *UpdateContainerGroupRequestContainer {
	s.Stdin = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetStdinOnce(v bool) *UpdateContainerGroupRequestContainer {
	s.StdinOnce = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetTty(v bool) *UpdateContainerGroupRequestContainer {
	s.Tty = &v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetVolumeMount(v []*UpdateContainerGroupRequestContainerVolumeMount) *UpdateContainerGroupRequestContainer {
	s.VolumeMount = v
	return s
}

func (s *UpdateContainerGroupRequestContainer) SetWorkingDir(v string) *UpdateContainerGroupRequestContainer {
	s.WorkingDir = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbe struct {
	Exec                *UpdateContainerGroupRequestContainerLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                      `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *UpdateContainerGroupRequestContainerLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                      `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                      `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                      `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                      `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbe) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbe) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetExec(v *UpdateContainerGroupRequestContainerLivenessProbeExec) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.Exec = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetFailureThreshold(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetHttpGet(v *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetInitialDelaySeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetPeriodSeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetSuccessThreshold(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetTcpSocket(v *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbe) SetTimeoutSeconds(v int32) *UpdateContainerGroupRequestContainerLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeExec) SetCommand(v []*string) *UpdateContainerGroupRequestContainerLivenessProbeExec {
	s.Command = v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetPath(v string) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetPort(v int32) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeHttpGet) SetScheme(v string) *UpdateContainerGroupRequestContainerLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

type UpdateContainerGroupRequestContainerLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket) SetPort(v int32) *UpdateContainerGroupRequestContainerLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbe struct {
	Exec                *UpdateContainerGroupRequestContainerReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" require:"true" type:"Struct"`
	FailureThreshold    *int32                                                       `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *UpdateContainerGroupRequestContainerReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" require:"true" type:"Struct"`
	InitialDelaySeconds *int32                                                       `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                       `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                       `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" require:"true" type:"Struct"`
	TimeoutSeconds      *int32                                                       `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbe) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbe) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetExec(v *UpdateContainerGroupRequestContainerReadinessProbeExec) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.Exec = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetFailureThreshold(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetHttpGet(v *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetInitialDelaySeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetPeriodSeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetSuccessThreshold(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetTcpSocket(v *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbe) SetTimeoutSeconds(v int32) *UpdateContainerGroupRequestContainerReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeExec) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeExec) SetCommand(v []*string) *UpdateContainerGroupRequestContainerReadinessProbeExec {
	s.Command = v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeHttpGet) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetPath(v string) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetPort(v int32) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeHttpGet) SetScheme(v string) *UpdateContainerGroupRequestContainerReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

type UpdateContainerGroupRequestContainerReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket) SetPort(v int32) *UpdateContainerGroupRequestContainerReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

type UpdateContainerGroupRequestContainerSecurityContext struct {
	Capability             *UpdateContainerGroupRequestContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                          `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                         `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s UpdateContainerGroupRequestContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetCapability(v *UpdateContainerGroupRequestContainerSecurityContextCapability) *UpdateContainerGroupRequestContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *UpdateContainerGroupRequestContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerSecurityContext) SetRunAsUser(v int64) *UpdateContainerGroupRequestContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type UpdateContainerGroupRequestContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerSecurityContextCapability) SetAdd(v []*string) *UpdateContainerGroupRequestContainerSecurityContextCapability {
	s.Add = v
	return s
}

type UpdateContainerGroupRequestContainerEnvironmentVar struct {
	FieldRef *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                     `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetFieldRef(v *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetKey(v string) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVar) SetValue(v string) *UpdateContainerGroupRequestContainerEnvironmentVar {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef) SetFieldPath(v string) *UpdateContainerGroupRequestContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) SetName(v string) *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders) SetValue(v string) *UpdateContainerGroupRequestContainerLifecyclePostStartHandlerHttpGetHttpHeaders {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetName(v string) *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader) SetValue(v string) *UpdateContainerGroupRequestContainerLifecyclePreStopHandlerHttpGetHttpHeader {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestContainerPort struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UpdateContainerGroupRequestContainerPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerPort) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerPort) SetPort(v int32) *UpdateContainerGroupRequestContainerPort {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerPort) SetProtocol(v string) *UpdateContainerGroupRequestContainerPort {
	s.Protocol = &v
	return s
}

type UpdateContainerGroupRequestContainerVolumeMount struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s UpdateContainerGroupRequestContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetMountPath(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetMountPropagation(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetName(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetReadOnly(v bool) *UpdateContainerGroupRequestContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestContainerVolumeMount) SetSubPath(v string) *UpdateContainerGroupRequestContainerVolumeMount {
	s.SubPath = &v
	return s
}

type UpdateContainerGroupRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateContainerGroupRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetPassword(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetServer(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *UpdateContainerGroupRequestImageRegistryCredential) SetUserName(v string) *UpdateContainerGroupRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type UpdateContainerGroupRequestInitContainer struct {
	SecurityContext *UpdateContainerGroupRequestInitContainerSecurityContext  `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" require:"true" type:"Struct"`
	Arg             []*string                                                 `json:"Arg,omitempty" xml:"Arg,omitempty" type:"Repeated"`
	Command         []*string                                                 `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Cpu             *float32                                                  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	EnvironmentVar  []*UpdateContainerGroupRequestInitContainerEnvironmentVar `json:"EnvironmentVar,omitempty" xml:"EnvironmentVar,omitempty" type:"Repeated"`
	Gpu             *int32                                                    `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Image           *string                                                   `json:"Image,omitempty" xml:"Image,omitempty"`
	ImagePullPolicy *string                                                   `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	Memory          *float32                                                  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Name            *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Port            []*UpdateContainerGroupRequestInitContainerPort           `json:"Port,omitempty" xml:"Port,omitempty" type:"Repeated"`
	Stdin           *bool                                                     `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	StdinOnce       *bool                                                     `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	Tty             *bool                                                     `json:"Tty,omitempty" xml:"Tty,omitempty"`
	VolumeMount     []*UpdateContainerGroupRequestInitContainerVolumeMount    `json:"VolumeMount,omitempty" xml:"VolumeMount,omitempty" type:"Repeated"`
	WorkingDir      *string                                                   `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainer) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainer) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainer) SetSecurityContext(v *UpdateContainerGroupRequestInitContainerSecurityContext) *UpdateContainerGroupRequestInitContainer {
	s.SecurityContext = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetArg(v []*string) *UpdateContainerGroupRequestInitContainer {
	s.Arg = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetCommand(v []*string) *UpdateContainerGroupRequestInitContainer {
	s.Command = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetCpu(v float32) *UpdateContainerGroupRequestInitContainer {
	s.Cpu = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetEnvironmentVar(v []*UpdateContainerGroupRequestInitContainerEnvironmentVar) *UpdateContainerGroupRequestInitContainer {
	s.EnvironmentVar = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetGpu(v int32) *UpdateContainerGroupRequestInitContainer {
	s.Gpu = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetImage(v string) *UpdateContainerGroupRequestInitContainer {
	s.Image = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetImagePullPolicy(v string) *UpdateContainerGroupRequestInitContainer {
	s.ImagePullPolicy = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetMemory(v float32) *UpdateContainerGroupRequestInitContainer {
	s.Memory = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetName(v string) *UpdateContainerGroupRequestInitContainer {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetPort(v []*UpdateContainerGroupRequestInitContainerPort) *UpdateContainerGroupRequestInitContainer {
	s.Port = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetStdin(v bool) *UpdateContainerGroupRequestInitContainer {
	s.Stdin = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetStdinOnce(v bool) *UpdateContainerGroupRequestInitContainer {
	s.StdinOnce = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetTty(v bool) *UpdateContainerGroupRequestInitContainer {
	s.Tty = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetVolumeMount(v []*UpdateContainerGroupRequestInitContainerVolumeMount) *UpdateContainerGroupRequestInitContainer {
	s.VolumeMount = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainer) SetWorkingDir(v string) *UpdateContainerGroupRequestInitContainer {
	s.WorkingDir = &v
	return s
}

type UpdateContainerGroupRequestInitContainerSecurityContext struct {
	Capability             *UpdateContainerGroupRequestInitContainerSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" require:"true" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                              `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                             `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerSecurityContext) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerSecurityContext) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetCapability(v *UpdateContainerGroupRequestInitContainerSecurityContextCapability) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.Capability = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetReadOnlyRootFilesystem(v bool) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContext) SetRunAsUser(v int64) *UpdateContainerGroupRequestInitContainerSecurityContext {
	s.RunAsUser = &v
	return s
}

type UpdateContainerGroupRequestInitContainerSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestInitContainerSecurityContextCapability) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerSecurityContextCapability) SetAdd(v []*string) *UpdateContainerGroupRequestInitContainerSecurityContextCapability {
	s.Add = v
	return s
}

type UpdateContainerGroupRequestInitContainerEnvironmentVar struct {
	FieldRef *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" require:"true" type:"Struct"`
	Key      *string                                                         `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string                                                         `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVar) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVar) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetFieldRef(v *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.FieldRef = v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetKey(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVar) SetValue(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVar {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef) SetFieldPath(v string) *UpdateContainerGroupRequestInitContainerEnvironmentVarFieldRef {
	s.FieldPath = &v
	return s
}

type UpdateContainerGroupRequestInitContainerPort struct {
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerPort) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerPort) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerPort) SetPort(v int32) *UpdateContainerGroupRequestInitContainerPort {
	s.Port = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerPort) SetProtocol(v string) *UpdateContainerGroupRequestInitContainerPort {
	s.Protocol = &v
	return s
}

type UpdateContainerGroupRequestInitContainerVolumeMount struct {
	MountPath        *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ReadOnly         *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	SubPath          *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s UpdateContainerGroupRequestInitContainerVolumeMount) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestInitContainerVolumeMount) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetMountPath(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.MountPath = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetMountPropagation(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.MountPropagation = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetName(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetReadOnly(v bool) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestInitContainerVolumeMount) SetSubPath(v string) *UpdateContainerGroupRequestInitContainerVolumeMount {
	s.SubPath = &v
	return s
}

type UpdateContainerGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateContainerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestTag) SetKey(v string) *UpdateContainerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateContainerGroupRequestTag) SetValue(v string) *UpdateContainerGroupRequestTag {
	s.Value = &v
	return s
}

type UpdateContainerGroupRequestVolume struct {
	ConfigFileVolume *UpdateContainerGroupRequestVolumeConfigFileVolume `json:"ConfigFileVolume,omitempty" xml:"ConfigFileVolume,omitempty" require:"true" type:"Struct"`
	EmptyDirVolume   *UpdateContainerGroupRequestVolumeEmptyDirVolume   `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" require:"true" type:"Struct"`
	FlexVolume       *UpdateContainerGroupRequestVolumeFlexVolume       `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" require:"true" type:"Struct"`
	HostPathVolume   *UpdateContainerGroupRequestVolumeHostPathVolume   `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" require:"true" type:"Struct"`
	NFSVolume        *UpdateContainerGroupRequestVolumeNFSVolume        `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" require:"true" type:"Struct"`
	Name             *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Type             *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateContainerGroupRequestVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolume) SetConfigFileVolume(v *UpdateContainerGroupRequestVolumeConfigFileVolume) *UpdateContainerGroupRequestVolume {
	s.ConfigFileVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetEmptyDirVolume(v *UpdateContainerGroupRequestVolumeEmptyDirVolume) *UpdateContainerGroupRequestVolume {
	s.EmptyDirVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetFlexVolume(v *UpdateContainerGroupRequestVolumeFlexVolume) *UpdateContainerGroupRequestVolume {
	s.FlexVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetHostPathVolume(v *UpdateContainerGroupRequestVolumeHostPathVolume) *UpdateContainerGroupRequestVolume {
	s.HostPathVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetNFSVolume(v *UpdateContainerGroupRequestVolumeNFSVolume) *UpdateContainerGroupRequestVolume {
	s.NFSVolume = v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetName(v string) *UpdateContainerGroupRequestVolume {
	s.Name = &v
	return s
}

func (s *UpdateContainerGroupRequestVolume) SetType(v string) *UpdateContainerGroupRequestVolume {
	s.Type = &v
	return s
}

type UpdateContainerGroupRequestVolumeConfigFileVolume struct {
	ConfigFileToPath []*UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath `json:"ConfigFileToPath,omitempty" xml:"ConfigFileToPath,omitempty" type:"Repeated"`
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolume) SetConfigFileToPath(v []*UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) *UpdateContainerGroupRequestVolumeConfigFileVolume {
	s.ConfigFileToPath = v
	return s
}

type UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetContent(v string) *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath) SetPath(v string) *UpdateContainerGroupRequestVolumeConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

type UpdateContainerGroupRequestVolumeEmptyDirVolume struct {
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeEmptyDirVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeEmptyDirVolume) SetMedium(v string) *UpdateContainerGroupRequestVolumeEmptyDirVolume {
	s.Medium = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeEmptyDirVolume) SetSizeLimit(v string) *UpdateContainerGroupRequestVolumeEmptyDirVolume {
	s.SizeLimit = &v
	return s
}

type UpdateContainerGroupRequestVolumeFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeFlexVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeFlexVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeFlexVolume) SetDriver(v string) *UpdateContainerGroupRequestVolumeFlexVolume {
	s.Driver = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeFlexVolume) SetFsType(v string) *UpdateContainerGroupRequestVolumeFlexVolume {
	s.FsType = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeFlexVolume) SetOptions(v string) *UpdateContainerGroupRequestVolumeFlexVolume {
	s.Options = &v
	return s
}

type UpdateContainerGroupRequestVolumeHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeHostPathVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeHostPathVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeHostPathVolume) SetPath(v string) *UpdateContainerGroupRequestVolumeHostPathVolume {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeHostPathVolume) SetType(v string) *UpdateContainerGroupRequestVolumeHostPathVolume {
	s.Type = &v
	return s
}

type UpdateContainerGroupRequestVolumeNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s UpdateContainerGroupRequestVolumeNFSVolume) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupRequestVolumeNFSVolume) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetPath(v string) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.Path = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetReadOnly(v bool) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *UpdateContainerGroupRequestVolumeNFSVolume) SetServer(v string) *UpdateContainerGroupRequestVolumeNFSVolume {
	s.Server = &v
	return s
}

type UpdateContainerGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateContainerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupResponseBody) SetRequestId(v string) *UpdateContainerGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateContainerGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateContainerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateContainerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateContainerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerGroupResponse) SetHeaders(v map[string]*string) *UpdateContainerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerGroupResponse) SetStatusCode(v int32) *UpdateContainerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContainerGroupResponse) SetBody(v *UpdateContainerGroupResponseBody) *UpdateContainerGroupResponse {
	s.Body = v
	return s
}

type UpdateImageCacheRequest struct {
	AcrRegistryInfo         []*UpdateImageCacheRequestAcrRegistryInfo         `json:"AcrRegistryInfo,omitempty" xml:"AcrRegistryInfo,omitempty" type:"Repeated"`
	AutoMatchImageCache     *bool                                             `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	ClientToken             *string                                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EipInstanceId           *string                                           `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
	EliminationStrategy     *string                                           `json:"EliminationStrategy,omitempty" xml:"EliminationStrategy,omitempty"`
	Flash                   *bool                                             `json:"Flash,omitempty" xml:"Flash,omitempty"`
	FlashCopyCount          *int32                                            `json:"FlashCopyCount,omitempty" xml:"FlashCopyCount,omitempty"`
	Image                   []*string                                         `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	ImageCacheId            *string                                           `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	ImageCacheName          *string                                           `json:"ImageCacheName,omitempty" xml:"ImageCacheName,omitempty"`
	ImageCacheSize          *int32                                            `json:"ImageCacheSize,omitempty" xml:"ImageCacheSize,omitempty"`
	ImageRegistryCredential []*UpdateImageCacheRequestImageRegistryCredential `json:"ImageRegistryCredential,omitempty" xml:"ImageRegistryCredential,omitempty" type:"Repeated"`
	OwnerAccount            *string                                           `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId                *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId         *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount    *string                                           `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                            `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RetentionDays           *int32                                            `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	SecurityGroupId         *string                                           `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	StandardCopyCount       *int32                                            `json:"StandardCopyCount,omitempty" xml:"StandardCopyCount,omitempty"`
	Tag                     []*UpdateImageCacheRequestTag                     `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId               *string                                           `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateImageCacheRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequest) SetAcrRegistryInfo(v []*UpdateImageCacheRequestAcrRegistryInfo) *UpdateImageCacheRequest {
	s.AcrRegistryInfo = v
	return s
}

func (s *UpdateImageCacheRequest) SetAutoMatchImageCache(v bool) *UpdateImageCacheRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *UpdateImageCacheRequest) SetClientToken(v string) *UpdateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateImageCacheRequest) SetEipInstanceId(v string) *UpdateImageCacheRequest {
	s.EipInstanceId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetEliminationStrategy(v string) *UpdateImageCacheRequest {
	s.EliminationStrategy = &v
	return s
}

func (s *UpdateImageCacheRequest) SetFlash(v bool) *UpdateImageCacheRequest {
	s.Flash = &v
	return s
}

func (s *UpdateImageCacheRequest) SetFlashCopyCount(v int32) *UpdateImageCacheRequest {
	s.FlashCopyCount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImage(v []*string) *UpdateImageCacheRequest {
	s.Image = v
	return s
}

func (s *UpdateImageCacheRequest) SetImageCacheId(v string) *UpdateImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImageCacheName(v string) *UpdateImageCacheRequest {
	s.ImageCacheName = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImageCacheSize(v int32) *UpdateImageCacheRequest {
	s.ImageCacheSize = &v
	return s
}

func (s *UpdateImageCacheRequest) SetImageRegistryCredential(v []*UpdateImageCacheRequestImageRegistryCredential) *UpdateImageCacheRequest {
	s.ImageRegistryCredential = v
	return s
}

func (s *UpdateImageCacheRequest) SetOwnerAccount(v string) *UpdateImageCacheRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetOwnerId(v int64) *UpdateImageCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetRegionId(v string) *UpdateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetResourceGroupId(v string) *UpdateImageCacheRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetResourceOwnerAccount(v string) *UpdateImageCacheRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetResourceOwnerId(v int64) *UpdateImageCacheRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetRetentionDays(v int32) *UpdateImageCacheRequest {
	s.RetentionDays = &v
	return s
}

func (s *UpdateImageCacheRequest) SetSecurityGroupId(v string) *UpdateImageCacheRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateImageCacheRequest) SetStandardCopyCount(v int32) *UpdateImageCacheRequest {
	s.StandardCopyCount = &v
	return s
}

func (s *UpdateImageCacheRequest) SetTag(v []*UpdateImageCacheRequestTag) *UpdateImageCacheRequest {
	s.Tag = v
	return s
}

func (s *UpdateImageCacheRequest) SetVSwitchId(v string) *UpdateImageCacheRequest {
	s.VSwitchId = &v
	return s
}

type UpdateImageCacheRequestAcrRegistryInfo struct {
	Domain       []*string `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
	InstanceId   *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateImageCacheRequestAcrRegistryInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequestAcrRegistryInfo) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetDomain(v []*string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.Domain = v
	return s
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetInstanceId(v string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.InstanceId = &v
	return s
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetInstanceName(v string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.InstanceName = &v
	return s
}

func (s *UpdateImageCacheRequestAcrRegistryInfo) SetRegionId(v string) *UpdateImageCacheRequestAcrRegistryInfo {
	s.RegionId = &v
	return s
}

type UpdateImageCacheRequestImageRegistryCredential struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateImageCacheRequestImageRegistryCredential) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequestImageRegistryCredential) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequestImageRegistryCredential) SetPassword(v string) *UpdateImageCacheRequestImageRegistryCredential {
	s.Password = &v
	return s
}

func (s *UpdateImageCacheRequestImageRegistryCredential) SetServer(v string) *UpdateImageCacheRequestImageRegistryCredential {
	s.Server = &v
	return s
}

func (s *UpdateImageCacheRequestImageRegistryCredential) SetUserName(v string) *UpdateImageCacheRequestImageRegistryCredential {
	s.UserName = &v
	return s
}

type UpdateImageCacheRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateImageCacheRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheRequestTag) SetKey(v string) *UpdateImageCacheRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateImageCacheRequestTag) SetValue(v string) *UpdateImageCacheRequestTag {
	s.Value = &v
	return s
}

type UpdateImageCacheResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateImageCacheResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheResponseBody) SetRequestId(v string) *UpdateImageCacheResponseBody {
	s.RequestId = &v
	return s
}

type UpdateImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageCacheResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageCacheResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageCacheResponse) SetHeaders(v map[string]*string) *UpdateImageCacheResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageCacheResponse) SetStatusCode(v int32) *UpdateImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageCacheResponse) SetBody(v *UpdateImageCacheResponseBody) *UpdateImageCacheResponse {
	s.Body = v
	return s
}

type UpdateVirtualNodeRequest struct {
	ClientToken          *string                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityGroupId      *string                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag                  []*UpdateVirtualNodeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId            *string                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VirtualNodeId        *string                        `json:"VirtualNodeId,omitempty" xml:"VirtualNodeId,omitempty"`
	VirtualNodeName      *string                        `json:"VirtualNodeName,omitempty" xml:"VirtualNodeName,omitempty"`
}

func (s UpdateVirtualNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeRequest) SetClientToken(v string) *UpdateVirtualNodeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetOwnerAccount(v string) *UpdateVirtualNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetOwnerId(v int64) *UpdateVirtualNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetRegionId(v string) *UpdateVirtualNodeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceGroupId(v string) *UpdateVirtualNodeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceOwnerAccount(v string) *UpdateVirtualNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetResourceOwnerId(v int64) *UpdateVirtualNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetSecurityGroupId(v string) *UpdateVirtualNodeRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetTag(v []*UpdateVirtualNodeRequestTag) *UpdateVirtualNodeRequest {
	s.Tag = v
	return s
}

func (s *UpdateVirtualNodeRequest) SetVSwitchId(v string) *UpdateVirtualNodeRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetVirtualNodeId(v string) *UpdateVirtualNodeRequest {
	s.VirtualNodeId = &v
	return s
}

func (s *UpdateVirtualNodeRequest) SetVirtualNodeName(v string) *UpdateVirtualNodeRequest {
	s.VirtualNodeName = &v
	return s
}

type UpdateVirtualNodeRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateVirtualNodeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeRequestTag) SetKey(v string) *UpdateVirtualNodeRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateVirtualNodeRequestTag) SetValue(v string) *UpdateVirtualNodeRequestTag {
	s.Value = &v
	return s
}

type UpdateVirtualNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVirtualNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeResponseBody) SetRequestId(v string) *UpdateVirtualNodeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVirtualNodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateVirtualNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVirtualNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirtualNodeResponse) SetHeaders(v map[string]*string) *UpdateVirtualNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirtualNodeResponse) SetStatusCode(v int32) *UpdateVirtualNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirtualNodeResponse) SetBody(v *UpdateVirtualNodeResponseBody) *UpdateVirtualNodeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eci"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateContainerGroupWithOptions(request *CreateContainerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ActiveDeadlineSeconds)) {
		query["ActiveDeadlineSeconds"] = request.ActiveDeadlineSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.Arn)) {
		query["Arn"] = request.Arn
	}

	if !tea.BoolValue(util.IsUnset(request.AutoCreateEip)) {
		query["AutoCreateEip"] = request.AutoCreateEip
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Container)) {
		query["Container"] = request.Container
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupName)) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.CorePattern)) {
		query["CorePattern"] = request.CorePattern
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsCore)) {
		query["CpuOptionsCore"] = request.CpuOptionsCore
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsNuma)) {
		query["CpuOptionsNuma"] = request.CpuOptionsNuma
	}

	if !tea.BoolValue(util.IsUnset(request.CpuOptionsThreadsPerCore)) {
		query["CpuOptionsThreadsPerCore"] = request.CpuOptionsThreadsPerCore
	}

	if !tea.BoolValue(util.IsUnset(request.DnsPolicy)) {
		query["DnsPolicy"] = request.DnsPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.EgressBandwidth)) {
		query["EgressBandwidth"] = request.EgressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.EipCommonBandwidthPackage)) {
		query["EipCommonBandwidthPackage"] = request.EipCommonBandwidthPackage
	}

	if !tea.BoolValue(util.IsUnset(request.EipISP)) {
		query["EipISP"] = request.EipISP
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EphemeralStorage)) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !tea.BoolValue(util.IsUnset(request.HostAliase)) {
		query["HostAliase"] = request.HostAliase
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageAccelerateMode)) {
		query["ImageAccelerateMode"] = request.ImageAccelerateMode
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.ImageSnapshotId)) {
		query["ImageSnapshotId"] = request.ImageSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.IngressBandwidth)) {
		query["IngressBandwidth"] = request.IngressBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainer)) {
		query["InitContainer"] = request.InitContainer
	}

	if !tea.BoolValue(util.IsUnset(request.InsecureRegistry)) {
		query["InsecureRegistry"] = request.InsecureRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6AddressCount)) {
		query["Ipv6AddressCount"] = request.Ipv6AddressCount
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6GatewayBandwidth)) {
		query["Ipv6GatewayBandwidth"] = request.Ipv6GatewayBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.Ipv6GatewayBandwidthEnable)) {
		query["Ipv6GatewayBandwidthEnable"] = request.Ipv6GatewayBandwidthEnable
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.NtpServer)) {
		query["NtpServer"] = request.NtpServer
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlainHttpRegistry)) {
		query["PlainHttpRegistry"] = request.PlainHttpRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.ProductOnEciMode)) {
		query["ProductOnEciMode"] = request.ProductOnEciMode
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RestartPolicy)) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleStrategy)) {
		query["ScheduleStrategy"] = request.ScheduleStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SecondaryENIPolicy)) {
		query["SecondaryENIPolicy"] = request.SecondaryENIPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareProcessNamespace)) {
		query["ShareProcessNamespace"] = request.ShareProcessNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.SlsEnable)) {
		query["SlsEnable"] = request.SlsEnable
	}

	if !tea.BoolValue(util.IsUnset(request.SpotDuration)) {
		query["SpotDuration"] = request.SpotDuration
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimit)) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.StrictSpot)) {
		query["StrictSpot"] = request.StrictSpot
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TenantSecurityGroupId)) {
		query["TenantSecurityGroupId"] = request.TenantSecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantVSwitchId)) {
		query["TenantVSwitchId"] = request.TenantVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationGracePeriodSeconds)) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DnsConfig))) {
		query["DnsConfig"] = request.DnsConfig
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.HostSecurityContext))) {
		query["HostSecurityContext"] = request.HostSecurityContext
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SecurityContext))) {
		query["SecurityContext"] = request.SecurityContext
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateContainerGroup(request *CreateContainerGroupRequest) (_result *CreateContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateContainerGroupResponse{}
	_body, _err := client.CreateContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageCacheWithOptions(request *CreateImageCacheRequest, runtime *util.RuntimeOptions) (_result *CreateImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EliminationStrategy)) {
		query["EliminationStrategy"] = request.EliminationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Flash)) {
		query["Flash"] = request.Flash
	}

	if !tea.BoolValue(util.IsUnset(request.FlashCopyCount)) {
		query["FlashCopyCount"] = request.FlashCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheName)) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheSize)) {
		query["ImageCacheSize"] = request.ImageCacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.InsecureRegistry)) {
		query["InsecureRegistry"] = request.InsecureRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlainHttpRegistry)) {
		query["PlainHttpRegistry"] = request.PlainHttpRegistry
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StandardCopyCount)) {
		query["StandardCopyCount"] = request.StandardCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImageCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageCache(request *CreateImageCacheRequest) (_result *CreateImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CreateImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceOpsTaskWithOptions(request *CreateInstanceOpsTaskRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceOpsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OpsType)) {
		query["OpsType"] = request.OpsType
	}

	if !tea.BoolValue(util.IsUnset(request.OpsValue)) {
		query["OpsValue"] = request.OpsValue
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceOpsTask"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceOpsTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstanceOpsTask(request *CreateInstanceOpsTaskRequest) (_result *CreateInstanceOpsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceOpsTaskResponse{}
	_body, _err := client.CreateInstanceOpsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVirtualNodeWithOptions(request *CreateVirtualNodeRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePublicNetwork)) {
		query["EnablePublicNetwork"] = request.EnablePublicNetwork
	}

	if !tea.BoolValue(util.IsUnset(request.KubeConfig)) {
		query["KubeConfig"] = request.KubeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RotateCertificateEnabled)) {
		query["RotateCertificateEnabled"] = request.RotateCertificateEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Taint)) {
		query["Taint"] = request.Taint
	}

	if !tea.BoolValue(util.IsUnset(request.TlsBootstrapEnabled)) {
		query["TlsBootstrapEnabled"] = request.TlsBootstrapEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeName)) {
		query["VirtualNodeName"] = request.VirtualNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVirtualNode"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVirtualNode(request *CreateVirtualNodeRequest) (_result *CreateVirtualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualNodeResponse{}
	_body, _err := client.CreateVirtualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContainerGroupWithOptions(request *DeleteContainerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (_result *DeleteContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContainerGroupResponse{}
	_body, _err := client.DeleteContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageCacheWithOptions(request *DeleteImageCacheRequest, runtime *util.RuntimeOptions) (_result *DeleteImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheId)) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImageCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageCache(request *DeleteImageCacheRequest) (_result *DeleteImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.DeleteImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualNodeWithOptions(request *DeleteVirtualNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeId)) {
		query["VirtualNodeId"] = request.VirtualNodeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualNode"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualNode(request *DeleteVirtualNodeRequest) (_result *DeleteVirtualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualNodeResponse{}
	_body, _err := client.DeleteVirtualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DestinationResource))) {
		query["DestinationResource"] = request.DestinationResource
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SpotResource))) {
		query["SpotResource"] = request.SpotResource
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableResource"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupEventsWithOptions(request *DescribeContainerGroupEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.EventSource)) {
		query["EventSource"] = request.EventSource
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SinceSecond)) {
		query["SinceSecond"] = request.SinceSecond
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupEvents"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupEvents(request *DescribeContainerGroupEventsRequest) (_result *DescribeContainerGroupEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupEventsResponse{}
	_body, _err := client.DescribeContainerGroupEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupMetricWithOptions(request *DescribeContainerGroupMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupMetric"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupMetric(request *DescribeContainerGroupMetricRequest) (_result *DescribeContainerGroupMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupMetricResponse{}
	_body, _err := client.DescribeContainerGroupMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupPriceWithOptions(request *DescribeContainerGroupPriceRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.EphemeralStorage)) {
		query["EphemeralStorage"] = request.EphemeralStorage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpotPriceLimit)) {
		query["SpotPriceLimit"] = request.SpotPriceLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SpotStrategy)) {
		query["SpotStrategy"] = request.SpotStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupPrice"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupPrice(request *DescribeContainerGroupPriceRequest) (_result *DescribeContainerGroupPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupPriceResponse{}
	_body, _err := client.DescribeContainerGroupPriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupStatusWithOptions(request *DescribeContainerGroupStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SinceSecond)) {
		query["SinceSecond"] = request.SinceSecond
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroupStatus"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroupStatus(request *DescribeContainerGroupStatusRequest) (_result *DescribeContainerGroupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupStatusResponse{}
	_body, _err := client.DescribeContainerGroupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerGroupsWithOptions(request *DescribeContainerGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupName)) {
		query["ContainerGroupName"] = request.ContainerGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.WithEvent)) {
		query["WithEvent"] = request.WithEvent
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerGroups"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (_result *DescribeContainerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerGroupsResponse{}
	_body, _err := client.DescribeContainerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerLogWithOptions(request *DescribeContainerLogRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		query["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.LastTime)) {
		query["LastTime"] = request.LastTime
	}

	if !tea.BoolValue(util.IsUnset(request.LimitBytes)) {
		query["LimitBytes"] = request.LimitBytes
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SinceSeconds)) {
		query["SinceSeconds"] = request.SinceSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tail)) {
		query["Tail"] = request.Tail
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamps)) {
		query["Timestamps"] = request.Timestamps
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContainerLog"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContainerLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerLog(request *DescribeContainerLogRequest) (_result *DescribeContainerLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerLogResponse{}
	_body, _err := client.DescribeContainerLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageCachesWithOptions(request *DescribeImageCachesRequest, runtime *util.RuntimeOptions) (_result *DescribeImageCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheId)) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheName)) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !tea.BoolValue(util.IsUnset(request.MatchImage)) {
		query["MatchImage"] = request.MatchImage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeImageCaches"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeImageCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageCaches(request *DescribeImageCachesRequest) (_result *DescribeImageCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageCachesResponse{}
	_body, _err := client.DescribeImageCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceOpsRecordsWithOptions(request *DescribeInstanceOpsRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceOpsRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OpsType)) {
		query["OpsType"] = request.OpsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceOpsRecords"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceOpsRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceOpsRecords(request *DescribeInstanceOpsRecordsRequest) (_result *DescribeInstanceOpsRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceOpsRecordsResponse{}
	_body, _err := client.DescribeInstanceOpsRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMultiContainerGroupMetricWithOptions(request *DescribeMultiContainerGroupMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeMultiContainerGroupMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerGroupIds)) {
		query["ContainerGroupIds"] = request.ContainerGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMultiContainerGroupMetric"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMultiContainerGroupMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMultiContainerGroupMetric(request *DescribeMultiContainerGroupMetricRequest) (_result *DescribeMultiContainerGroupMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMultiContainerGroupMetricResponse{}
	_body, _err := client.DescribeMultiContainerGroupMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVirtualNodesWithOptions(request *DescribeVirtualNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeVirtualNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeIds)) {
		query["VirtualNodeIds"] = request.VirtualNodeIds
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeName)) {
		query["VirtualNodeName"] = request.VirtualNodeName
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVirtualNodes"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVirtualNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVirtualNodes(request *DescribeVirtualNodesRequest) (_result *DescribeVirtualNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVirtualNodesResponse{}
	_body, _err := client.DescribeVirtualNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecContainerCommandWithOptions(request *ExecContainerCommandRequest, runtime *util.RuntimeOptions) (_result *ExecContainerCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerName)) {
		query["ContainerName"] = request.ContainerName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Stdin)) {
		query["Stdin"] = request.Stdin
	}

	if !tea.BoolValue(util.IsUnset(request.Sync)) {
		query["Sync"] = request.Sync
	}

	if !tea.BoolValue(util.IsUnset(request.TTY)) {
		query["TTY"] = request.TTY
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecContainerCommand"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecContainerCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecContainerCommand(request *ExecContainerCommandRequest) (_result *ExecContainerCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecContainerCommandResponse{}
	_body, _err := client.ExecContainerCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsageWithOptions(request *ListUsageRequest, runtime *util.RuntimeOptions) (_result *ListUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsage"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsage(request *ListUsageRequest) (_result *ListUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsageResponse{}
	_body, _err := client.ListUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeContainerGroupVolumeWithOptions(request *ResizeContainerGroupVolumeRequest, runtime *util.RuntimeOptions) (_result *ResizeContainerGroupVolumeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewSize)) {
		query["NewSize"] = request.NewSize
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VolumeName)) {
		query["VolumeName"] = request.VolumeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResizeContainerGroupVolume"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResizeContainerGroupVolumeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeContainerGroupVolume(request *ResizeContainerGroupVolumeRequest) (_result *ResizeContainerGroupVolumeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeContainerGroupVolumeResponse{}
	_body, _err := client.ResizeContainerGroupVolumeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartContainerGroupWithOptions(request *RestartContainerGroupRequest, runtime *util.RuntimeOptions) (_result *RestartContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartContainerGroup(request *RestartContainerGroupRequest) (_result *RestartContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartContainerGroupResponse{}
	_body, _err := client.RestartContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateContainerGroupWithOptions(request *UpdateContainerGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateContainerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Container)) {
		query["Container"] = request.Container
	}

	if !tea.BoolValue(util.IsUnset(request.ContainerGroupId)) {
		query["ContainerGroupId"] = request.ContainerGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		query["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.InitContainer)) {
		query["InitContainer"] = request.InitContainer
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		query["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RestartPolicy)) {
		query["RestartPolicy"] = request.RestartPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		query["UpdateType"] = request.UpdateType
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DnsConfig))) {
		query["DnsConfig"] = request.DnsConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateContainerGroup"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateContainerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateContainerGroup(request *UpdateContainerGroupRequest) (_result *UpdateContainerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateContainerGroupResponse{}
	_body, _err := client.UpdateContainerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateImageCacheWithOptions(request *UpdateImageCacheRequest, runtime *util.RuntimeOptions) (_result *UpdateImageCacheResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcrRegistryInfo)) {
		query["AcrRegistryInfo"] = request.AcrRegistryInfo
	}

	if !tea.BoolValue(util.IsUnset(request.AutoMatchImageCache)) {
		query["AutoMatchImageCache"] = request.AutoMatchImageCache
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EipInstanceId)) {
		query["EipInstanceId"] = request.EipInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EliminationStrategy)) {
		query["EliminationStrategy"] = request.EliminationStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Flash)) {
		query["Flash"] = request.Flash
	}

	if !tea.BoolValue(util.IsUnset(request.FlashCopyCount)) {
		query["FlashCopyCount"] = request.FlashCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheId)) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheName)) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageCacheSize)) {
		query["ImageCacheSize"] = request.ImageCacheSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageRegistryCredential)) {
		query["ImageRegistryCredential"] = request.ImageRegistryCredential
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionDays)) {
		query["RetentionDays"] = request.RetentionDays
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StandardCopyCount)) {
		query["StandardCopyCount"] = request.StandardCopyCount
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImageCache"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImageCache(request *UpdateImageCacheRequest) (_result *UpdateImageCacheResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageCacheResponse{}
	_body, _err := client.UpdateImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVirtualNodeWithOptions(request *UpdateVirtualNodeRequest, runtime *util.RuntimeOptions) (_result *UpdateVirtualNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeId)) {
		query["VirtualNodeId"] = request.VirtualNodeId
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualNodeName)) {
		query["VirtualNodeName"] = request.VirtualNodeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVirtualNode"),
		Version:     tea.String("2018-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVirtualNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVirtualNode(request *UpdateVirtualNodeRequest) (_result *UpdateVirtualNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVirtualNodeResponse{}
	_body, _err := client.UpdateVirtualNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
