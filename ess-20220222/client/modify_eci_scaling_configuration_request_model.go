// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEciScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrRegistryInfos(v []*ModifyEciScalingConfigurationRequestAcrRegistryInfos) *ModifyEciScalingConfigurationRequest
	GetAcrRegistryInfos() []*ModifyEciScalingConfigurationRequestAcrRegistryInfos
	SetActiveDeadlineSeconds(v int64) *ModifyEciScalingConfigurationRequest
	GetActiveDeadlineSeconds() *int64
	SetAutoCreateEip(v bool) *ModifyEciScalingConfigurationRequest
	GetAutoCreateEip() *bool
	SetAutoMatchImageCache(v bool) *ModifyEciScalingConfigurationRequest
	GetAutoMatchImageCache() *bool
	SetContainerGroupName(v string) *ModifyEciScalingConfigurationRequest
	GetContainerGroupName() *string
	SetContainers(v []*ModifyEciScalingConfigurationRequestContainers) *ModifyEciScalingConfigurationRequest
	GetContainers() []*ModifyEciScalingConfigurationRequestContainers
	SetContainersUpdateType(v string) *ModifyEciScalingConfigurationRequest
	GetContainersUpdateType() *string
	SetCostOptimization(v bool) *ModifyEciScalingConfigurationRequest
	GetCostOptimization() *bool
	SetCpu(v float32) *ModifyEciScalingConfigurationRequest
	GetCpu() *float32
	SetCpuOptionsCore(v int32) *ModifyEciScalingConfigurationRequest
	GetCpuOptionsCore() *int32
	SetCpuOptionsThreadsPerCore(v int32) *ModifyEciScalingConfigurationRequest
	GetCpuOptionsThreadsPerCore() *int32
	SetDataCacheBucket(v string) *ModifyEciScalingConfigurationRequest
	GetDataCacheBucket() *string
	SetDataCacheBurstingEnabled(v bool) *ModifyEciScalingConfigurationRequest
	GetDataCacheBurstingEnabled() *bool
	SetDataCachePL(v string) *ModifyEciScalingConfigurationRequest
	GetDataCachePL() *string
	SetDataCacheProvisionedIops(v int32) *ModifyEciScalingConfigurationRequest
	GetDataCacheProvisionedIops() *int32
	SetDescription(v string) *ModifyEciScalingConfigurationRequest
	GetDescription() *string
	SetDnsConfigNameServers(v []*string) *ModifyEciScalingConfigurationRequest
	GetDnsConfigNameServers() []*string
	SetDnsConfigOptions(v []*ModifyEciScalingConfigurationRequestDnsConfigOptions) *ModifyEciScalingConfigurationRequest
	GetDnsConfigOptions() []*ModifyEciScalingConfigurationRequestDnsConfigOptions
	SetDnsConfigSearchs(v []*string) *ModifyEciScalingConfigurationRequest
	GetDnsConfigSearchs() []*string
	SetDnsPolicy(v string) *ModifyEciScalingConfigurationRequest
	GetDnsPolicy() *string
	SetEgressBandwidth(v int64) *ModifyEciScalingConfigurationRequest
	GetEgressBandwidth() *int64
	SetEipBandwidth(v int32) *ModifyEciScalingConfigurationRequest
	GetEipBandwidth() *int32
	SetEnableSls(v bool) *ModifyEciScalingConfigurationRequest
	GetEnableSls() *bool
	SetEphemeralStorage(v int32) *ModifyEciScalingConfigurationRequest
	GetEphemeralStorage() *int32
	SetGpuDriverVersion(v string) *ModifyEciScalingConfigurationRequest
	GetGpuDriverVersion() *string
	SetHostAliases(v []*ModifyEciScalingConfigurationRequestHostAliases) *ModifyEciScalingConfigurationRequest
	GetHostAliases() []*ModifyEciScalingConfigurationRequestHostAliases
	SetHostName(v string) *ModifyEciScalingConfigurationRequest
	GetHostName() *string
	SetImageRegistryCredentials(v []*ModifyEciScalingConfigurationRequestImageRegistryCredentials) *ModifyEciScalingConfigurationRequest
	GetImageRegistryCredentials() []*ModifyEciScalingConfigurationRequestImageRegistryCredentials
	SetImageSnapshotId(v string) *ModifyEciScalingConfigurationRequest
	GetImageSnapshotId() *string
	SetIngressBandwidth(v int64) *ModifyEciScalingConfigurationRequest
	GetIngressBandwidth() *int64
	SetInitContainers(v []*ModifyEciScalingConfigurationRequestInitContainers) *ModifyEciScalingConfigurationRequest
	GetInitContainers() []*ModifyEciScalingConfigurationRequestInitContainers
	SetInstanceFamilyLevel(v string) *ModifyEciScalingConfigurationRequest
	GetInstanceFamilyLevel() *string
	SetInstanceTypes(v []*string) *ModifyEciScalingConfigurationRequest
	GetInstanceTypes() []*string
	SetIpv6AddressCount(v int32) *ModifyEciScalingConfigurationRequest
	GetIpv6AddressCount() *int32
	SetLoadBalancerWeight(v int32) *ModifyEciScalingConfigurationRequest
	GetLoadBalancerWeight() *int32
	SetMemory(v float32) *ModifyEciScalingConfigurationRequest
	GetMemory() *float32
	SetNtpServers(v []*string) *ModifyEciScalingConfigurationRequest
	GetNtpServers() []*string
	SetOwnerId(v int64) *ModifyEciScalingConfigurationRequest
	GetOwnerId() *int64
	SetRamRoleName(v string) *ModifyEciScalingConfigurationRequest
	GetRamRoleName() *string
	SetResourceGroupId(v string) *ModifyEciScalingConfigurationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyEciScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetRestartPolicy(v string) *ModifyEciScalingConfigurationRequest
	GetRestartPolicy() *string
	SetScalingConfigurationId(v string) *ModifyEciScalingConfigurationRequest
	GetScalingConfigurationId() *string
	SetScalingConfigurationName(v string) *ModifyEciScalingConfigurationRequest
	GetScalingConfigurationName() *string
	SetSecurityContextSysCtls(v []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls) *ModifyEciScalingConfigurationRequest
	GetSecurityContextSysCtls() []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls
	SetSecurityGroupId(v string) *ModifyEciScalingConfigurationRequest
	GetSecurityGroupId() *string
	SetSpotPriceLimit(v float32) *ModifyEciScalingConfigurationRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *ModifyEciScalingConfigurationRequest
	GetSpotStrategy() *string
	SetTags(v []*ModifyEciScalingConfigurationRequestTags) *ModifyEciScalingConfigurationRequest
	GetTags() []*ModifyEciScalingConfigurationRequestTags
	SetTerminationGracePeriodSeconds(v int64) *ModifyEciScalingConfigurationRequest
	GetTerminationGracePeriodSeconds() *int64
	SetVolumes(v []*ModifyEciScalingConfigurationRequestVolumes) *ModifyEciScalingConfigurationRequest
	GetVolumes() []*ModifyEciScalingConfigurationRequestVolumes
}

type ModifyEciScalingConfigurationRequest struct {
	// The Container Registry Enterprise Edition instances.
	AcrRegistryInfos []*ModifyEciScalingConfigurationRequestAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	// The validity period of the scaling configuration. Unit: seconds.
	//
	// example:
	//
	// 1000
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	// Specifies whether to automatically create elastic IP addresses (EIPs) and bind the EIPs to elastic container instances.
	//
	// example:
	//
	// true
	AutoCreateEip *bool `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	// Specifies whether to automatically match image caches.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The name series of elastic container instances. Naming conventions:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name can contain only lowercase letters, digits, and hyphens (-). It cannot start or end with a hyphen (-).
	//
	// example:
	//
	// nginx-test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The containers.
	Containers []*ModifyEciScalingConfigurationRequestContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The update mode of containers. Valid values:
	//
	// 	- RenewUpdate: full update mode. This value takes effect based on the value of Containers in an update request. This value indicates that the previous setting of Containers is overwritten.
	//
	// 	- IncrementalUpdate: incremental update mode. Container matching is performed based on the Container.name value. Only the parameters that are included in the request parameters are updated.
	//
	// Default value: RenewUpdate.
	//
	// example:
	//
	// RenewUpdate
	ContainersUpdateType *string `json:"ContainersUpdateType,omitempty" xml:"ContainersUpdateType,omitempty"`
	// Specifies whether to enable the Cost Optimization feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	CostOptimization *bool `json:"CostOptimization,omitempty" xml:"CostOptimization,omitempty"`
	// The number of vCPUs per elastic container instance.
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of physical CPU cores. You can specify this parameter for only specific ECS instance types. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// The number of threads per core. You can specify this parameter for only specific instance types. A value of 1 specifies that Hyper-Threading is disabled. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The bucket in which data caches are stored.
	//
	// example:
	//
	// default
	DataCacheBucket *string `json:"DataCacheBucket,omitempty" xml:"DataCacheBucket,omitempty"`
	// Specifies whether to enable the Performance Burst feature for the ESSD AutoPL disk in which data caches are stored. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  For more information about ESSD AutoPL disks, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	DataCacheBurstingEnabled *bool `json:"DataCacheBurstingEnabled,omitempty" xml:"DataCacheBurstingEnabled,omitempty"`
	// The performance level (PL) of the cloud disk in which data caches are stored. We recommend that you use Enterprise SSDs (ESSDs). Valid values:
	//
	// 	- PL0: An ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// >  For more information about ESSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	DataCachePL *string `json:"DataCachePL,omitempty" xml:"DataCachePL,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk in which data caches are stored. Valid values: 0 to min{50,000, 1,000 × *Capacity - Baseline IOPS}. Baseline IOPS = min{1,800+50 x *Capacity, 50,000}.
	//
	// >  For more information about ESSD AutoPL disks, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	DataCacheProvisionedIops *int32 `json:"DataCacheProvisionedIops,omitempty" xml:"DataCacheProvisionedIops,omitempty"`
	// >  This parameter is unavailable for use.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP addresses of DNS servers.
	DnsConfigNameServers []*string `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	// The options. Each option is a name-value pair. The value in the name-value pair is optional.
	DnsConfigOptions []*ModifyEciScalingConfigurationRequestDnsConfigOptions `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	// The search domains of DNS servers.
	DnsConfigSearchs []*string `json:"DnsConfigSearchs,omitempty" xml:"DnsConfigSearchs,omitempty" type:"Repeated"`
	// The Domain Name System (DNS) policy. Valid values:
	//
	// 	- None: uses the DNS that is specified by DnsConfig.
	//
	// 	- Default: uses the DNS that is specified for the runtime environment.
	//
	// example:
	//
	// Default
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// The maximum outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 1024000
	EgressBandwidth *int64 `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	// The EIP bandwidth.
	//
	// Default value: 5. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// >  This parameter is not available for use.
	//
	// example:
	//
	// false
	EnableSls *bool `json:"EnableSls,omitempty" xml:"EnableSls,omitempty"`
	// The size of the temporary storage space. By default, an Enterprise SSD (ESSD) of the PL1 type is used. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The version of the GPU driver. Valid values:
	//
	// 	- tesla=470.82.01 (default)
	//
	// 	- tesla=525.85.12
	//
	// >  You can switch the GPU driver version only for a few Elastic Compute Service (ECS) instance types. For more information, see [Specify GPU-accelerated ECS instance types to create an elastic container instance](https://help.aliyun.com/document_detail/2579486.html).
	//
	// example:
	//
	// tesla=525.85.12
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The hosts.
	HostAliases []*ModifyEciScalingConfigurationRequestHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The hostname series of elastic container instances.
	//
	// example:
	//
	// test
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image repositories.
	ImageRegistryCredentials []*ModifyEciScalingConfigurationRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	// The ID of the image cache.
	//
	// example:
	//
	// imc-2zebxkiifuyzzlhl****
	ImageSnapshotId *string `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	// The maximum inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 1024000
	IngressBandwidth *int64 `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	// The init containers.
	InitContainers []*ModifyEciScalingConfigurationRequestInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The level of the instance family, which is used to filter instance types that meet the specified criteria. This parameter takes effect only if you set `CostOptimization` to true. Valid values:
	//
	// 	- EntryLevel: entry level (shared instance type). Instance types of this level are the most cost-effective but may not provide stable computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise level. Instance types of this level provide stable performance and dedicated resources, and are suitable for business scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit-based entry level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview](https://help.aliyun.com/document_detail/59977.html) of burstable instances.
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The ECS instance types. You can specify up to five instance types.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of IPv6 addresses.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The load balancing weight of each backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size per elastic container instance. Unit: GiB.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The endpoints of Network Time Protocol (NTP) servers.
	NtpServers []*string `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	OwnerId    *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the instance Resource Access Management (RAM) role. You can use the same RAM role to access elastic container instances and Elastic Compute Service (ECS) instances. For more information, see [Use an instance RAM role by calling API operations](https://help.aliyun.com/document_detail/61178.html).
	//
	// example:
	//
	// RamTestRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-uf66jeqopgqa9hdn****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The instance restart policy. Valid values:
	//
	// 	- Always: always restarts elastic container instances.
	//
	// 	- Never: never restarts elastic container instances.
	//
	// 	- OnFailure: restarts elastic container instances upon failures.
	//
	// Default value: Always.
	//
	// example:
	//
	// Always
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The ID of the scaling configuration that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// asc-bp16har3jpj6fjbx****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The name of the scaling configuration. The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). The name must start with a letter or a digit.
	//
	// The name of a scaling configuration must be unique in the specified region. If you do not specify this parameter, the value of ScalingConfigurationId is used.
	//
	// example:
	//
	// test-modify
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The security contexts in which the elastic container instance runs.
	SecurityContextSysCtls []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls `json:"SecurityContextSysCtls,omitempty" xml:"SecurityContextSysCtls,omitempty" type:"Repeated"`
	// The ID of the security group to which elastic container instances belong. Elastic container instances that belong to the same security group can communicate with each other.
	//
	// If you do not specify a security group, the system uses the default security group in the region that you selected. Make sure that the inbound rules of the security group contain the protocols and port numbers of the containers that you want to expose. If you do not have a default security group in the region, the system creates a default security group and then adds the container protocols and port numbers that you specified to the inbound rules of the security group.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The maximum hourly price of preemptible elastic container instances. The value can be accurate to three decimal places.
	//
	// If you set SpotStrategy to SpotWithPriceLimit, you must specify SpotPriceLimit.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The instance bidding policy. Valid values:
	//
	// 	- NoSpot: The instances are created as pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are preemptible instances for which you can specify the maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are created as preemptible instances for which the market price at the time of purchase is used as the bid price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// SpotPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags.
	Tags []*ModifyEciScalingConfigurationRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The buffer period during which the program handles operations before the program is stopped. Unit: seconds.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The volumes.
	Volumes []*ModifyEciScalingConfigurationRequestVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequest) GetAcrRegistryInfos() []*ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	return s.AcrRegistryInfos
}

func (s *ModifyEciScalingConfigurationRequest) GetActiveDeadlineSeconds() *int64 {
	return s.ActiveDeadlineSeconds
}

func (s *ModifyEciScalingConfigurationRequest) GetAutoCreateEip() *bool {
	return s.AutoCreateEip
}

func (s *ModifyEciScalingConfigurationRequest) GetAutoMatchImageCache() *bool {
	return s.AutoMatchImageCache
}

func (s *ModifyEciScalingConfigurationRequest) GetContainerGroupName() *string {
	return s.ContainerGroupName
}

func (s *ModifyEciScalingConfigurationRequest) GetContainers() []*ModifyEciScalingConfigurationRequestContainers {
	return s.Containers
}

func (s *ModifyEciScalingConfigurationRequest) GetContainersUpdateType() *string {
	return s.ContainersUpdateType
}

func (s *ModifyEciScalingConfigurationRequest) GetCostOptimization() *bool {
	return s.CostOptimization
}

func (s *ModifyEciScalingConfigurationRequest) GetCpu() *float32 {
	return s.Cpu
}

func (s *ModifyEciScalingConfigurationRequest) GetCpuOptionsCore() *int32 {
	return s.CpuOptionsCore
}

func (s *ModifyEciScalingConfigurationRequest) GetCpuOptionsThreadsPerCore() *int32 {
	return s.CpuOptionsThreadsPerCore
}

func (s *ModifyEciScalingConfigurationRequest) GetDataCacheBucket() *string {
	return s.DataCacheBucket
}

func (s *ModifyEciScalingConfigurationRequest) GetDataCacheBurstingEnabled() *bool {
	return s.DataCacheBurstingEnabled
}

func (s *ModifyEciScalingConfigurationRequest) GetDataCachePL() *string {
	return s.DataCachePL
}

func (s *ModifyEciScalingConfigurationRequest) GetDataCacheProvisionedIops() *int32 {
	return s.DataCacheProvisionedIops
}

func (s *ModifyEciScalingConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyEciScalingConfigurationRequest) GetDnsConfigNameServers() []*string {
	return s.DnsConfigNameServers
}

func (s *ModifyEciScalingConfigurationRequest) GetDnsConfigOptions() []*ModifyEciScalingConfigurationRequestDnsConfigOptions {
	return s.DnsConfigOptions
}

func (s *ModifyEciScalingConfigurationRequest) GetDnsConfigSearchs() []*string {
	return s.DnsConfigSearchs
}

func (s *ModifyEciScalingConfigurationRequest) GetDnsPolicy() *string {
	return s.DnsPolicy
}

func (s *ModifyEciScalingConfigurationRequest) GetEgressBandwidth() *int64 {
	return s.EgressBandwidth
}

func (s *ModifyEciScalingConfigurationRequest) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *ModifyEciScalingConfigurationRequest) GetEnableSls() *bool {
	return s.EnableSls
}

func (s *ModifyEciScalingConfigurationRequest) GetEphemeralStorage() *int32 {
	return s.EphemeralStorage
}

func (s *ModifyEciScalingConfigurationRequest) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *ModifyEciScalingConfigurationRequest) GetHostAliases() []*ModifyEciScalingConfigurationRequestHostAliases {
	return s.HostAliases
}

func (s *ModifyEciScalingConfigurationRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyEciScalingConfigurationRequest) GetImageRegistryCredentials() []*ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	return s.ImageRegistryCredentials
}

func (s *ModifyEciScalingConfigurationRequest) GetImageSnapshotId() *string {
	return s.ImageSnapshotId
}

func (s *ModifyEciScalingConfigurationRequest) GetIngressBandwidth() *int64 {
	return s.IngressBandwidth
}

func (s *ModifyEciScalingConfigurationRequest) GetInitContainers() []*ModifyEciScalingConfigurationRequestInitContainers {
	return s.InitContainers
}

func (s *ModifyEciScalingConfigurationRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *ModifyEciScalingConfigurationRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ModifyEciScalingConfigurationRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *ModifyEciScalingConfigurationRequest) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *ModifyEciScalingConfigurationRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *ModifyEciScalingConfigurationRequest) GetNtpServers() []*string {
	return s.NtpServers
}

func (s *ModifyEciScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyEciScalingConfigurationRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *ModifyEciScalingConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyEciScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyEciScalingConfigurationRequest) GetRestartPolicy() *string {
	return s.RestartPolicy
}

func (s *ModifyEciScalingConfigurationRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *ModifyEciScalingConfigurationRequest) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *ModifyEciScalingConfigurationRequest) GetSecurityContextSysCtls() []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls {
	return s.SecurityContextSysCtls
}

func (s *ModifyEciScalingConfigurationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyEciScalingConfigurationRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *ModifyEciScalingConfigurationRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ModifyEciScalingConfigurationRequest) GetTags() []*ModifyEciScalingConfigurationRequestTags {
	return s.Tags
}

func (s *ModifyEciScalingConfigurationRequest) GetTerminationGracePeriodSeconds() *int64 {
	return s.TerminationGracePeriodSeconds
}

func (s *ModifyEciScalingConfigurationRequest) GetVolumes() []*ModifyEciScalingConfigurationRequestVolumes {
	return s.Volumes
}

func (s *ModifyEciScalingConfigurationRequest) SetAcrRegistryInfos(v []*ModifyEciScalingConfigurationRequestAcrRegistryInfos) *ModifyEciScalingConfigurationRequest {
	s.AcrRegistryInfos = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetActiveDeadlineSeconds(v int64) *ModifyEciScalingConfigurationRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetAutoCreateEip(v bool) *ModifyEciScalingConfigurationRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetAutoMatchImageCache(v bool) *ModifyEciScalingConfigurationRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetContainerGroupName(v string) *ModifyEciScalingConfigurationRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetContainers(v []*ModifyEciScalingConfigurationRequestContainers) *ModifyEciScalingConfigurationRequest {
	s.Containers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetContainersUpdateType(v string) *ModifyEciScalingConfigurationRequest {
	s.ContainersUpdateType = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCostOptimization(v bool) *ModifyEciScalingConfigurationRequest {
	s.CostOptimization = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCpu(v float32) *ModifyEciScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCpuOptionsCore(v int32) *ModifyEciScalingConfigurationRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetCpuOptionsThreadsPerCore(v int32) *ModifyEciScalingConfigurationRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDataCacheBucket(v string) *ModifyEciScalingConfigurationRequest {
	s.DataCacheBucket = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDataCacheBurstingEnabled(v bool) *ModifyEciScalingConfigurationRequest {
	s.DataCacheBurstingEnabled = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDataCachePL(v string) *ModifyEciScalingConfigurationRequest {
	s.DataCachePL = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDataCacheProvisionedIops(v int32) *ModifyEciScalingConfigurationRequest {
	s.DataCacheProvisionedIops = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDescription(v string) *ModifyEciScalingConfigurationRequest {
	s.Description = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsConfigNameServers(v []*string) *ModifyEciScalingConfigurationRequest {
	s.DnsConfigNameServers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsConfigOptions(v []*ModifyEciScalingConfigurationRequestDnsConfigOptions) *ModifyEciScalingConfigurationRequest {
	s.DnsConfigOptions = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsConfigSearchs(v []*string) *ModifyEciScalingConfigurationRequest {
	s.DnsConfigSearchs = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetDnsPolicy(v string) *ModifyEciScalingConfigurationRequest {
	s.DnsPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEgressBandwidth(v int64) *ModifyEciScalingConfigurationRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEipBandwidth(v int32) *ModifyEciScalingConfigurationRequest {
	s.EipBandwidth = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEnableSls(v bool) *ModifyEciScalingConfigurationRequest {
	s.EnableSls = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetEphemeralStorage(v int32) *ModifyEciScalingConfigurationRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetGpuDriverVersion(v string) *ModifyEciScalingConfigurationRequest {
	s.GpuDriverVersion = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetHostAliases(v []*ModifyEciScalingConfigurationRequestHostAliases) *ModifyEciScalingConfigurationRequest {
	s.HostAliases = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetHostName(v string) *ModifyEciScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetImageRegistryCredentials(v []*ModifyEciScalingConfigurationRequestImageRegistryCredentials) *ModifyEciScalingConfigurationRequest {
	s.ImageRegistryCredentials = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetImageSnapshotId(v string) *ModifyEciScalingConfigurationRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetIngressBandwidth(v int64) *ModifyEciScalingConfigurationRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetInitContainers(v []*ModifyEciScalingConfigurationRequestInitContainers) *ModifyEciScalingConfigurationRequest {
	s.InitContainers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetInstanceFamilyLevel(v string) *ModifyEciScalingConfigurationRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetInstanceTypes(v []*string) *ModifyEciScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetIpv6AddressCount(v int32) *ModifyEciScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *ModifyEciScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetMemory(v float32) *ModifyEciScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetNtpServers(v []*string) *ModifyEciScalingConfigurationRequest {
	s.NtpServers = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetOwnerId(v int64) *ModifyEciScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetRamRoleName(v string) *ModifyEciScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetResourceGroupId(v string) *ModifyEciScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyEciScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetRestartPolicy(v string) *ModifyEciScalingConfigurationRequest {
	s.RestartPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetScalingConfigurationId(v string) *ModifyEciScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetScalingConfigurationName(v string) *ModifyEciScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSecurityContextSysCtls(v []*ModifyEciScalingConfigurationRequestSecurityContextSysCtls) *ModifyEciScalingConfigurationRequest {
	s.SecurityContextSysCtls = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSecurityGroupId(v string) *ModifyEciScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSpotPriceLimit(v float32) *ModifyEciScalingConfigurationRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetSpotStrategy(v string) *ModifyEciScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetTags(v []*ModifyEciScalingConfigurationRequestTags) *ModifyEciScalingConfigurationRequest {
	s.Tags = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetTerminationGracePeriodSeconds(v int64) *ModifyEciScalingConfigurationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) SetVolumes(v []*ModifyEciScalingConfigurationRequestVolumes) *ModifyEciScalingConfigurationRequest {
	s.Volumes = v
	return s
}

func (s *ModifyEciScalingConfigurationRequest) Validate() error {
	if s.AcrRegistryInfos != nil {
		for _, item := range s.AcrRegistryInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Containers != nil {
		for _, item := range s.Containers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DnsConfigOptions != nil {
		for _, item := range s.DnsConfigOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HostAliases != nil {
		for _, item := range s.HostAliases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageRegistryCredentials != nil {
		for _, item := range s.ImageRegistryCredentials {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InitContainers != nil {
		for _, item := range s.InitContainers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityContextSysCtls != nil {
		for _, item := range s.SecurityContextSysCtls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Volumes != nil {
		for _, item := range s.Volumes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestAcrRegistryInfos struct {
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. Separate multiple domain names with commas (,).
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// acr-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestAcrRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) GetDomains() []*string {
	return s.Domains
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetDomains(v []*string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceId(v string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceName(v string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) SetRegionId(v string) *ModifyEciScalingConfigurationRequestAcrRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestAcrRegistryInfos) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainers struct {
	LivenessProbe   *ModifyEciScalingConfigurationRequestContainersLivenessProbe   `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	ReadinessProbe  *ModifyEciScalingConfigurationRequestContainersReadinessProbe  `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	SecurityContext *ModifyEciScalingConfigurationRequestContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The container startup arguments. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The commands that you can run in the container when you use the CLI to perform a liveness probe.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs per container.
	//
	// example:
	//
	// 0.25
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables.
	EnvironmentVars []*ModifyEciScalingConfigurationRequestContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs per container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The container image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Image pulling is performed each time instances are created.
	//
	// 	- IfNotPresent: Image pulling is performed as needed. On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The commands that you can run within the container to configure the postStart callback function.
	LifecyclePostStartHandlerExecs []*string `json:"LifecyclePostStartHandlerExecs,omitempty" xml:"LifecyclePostStartHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which you want to send the HTTP GET request to configure the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The path to which you want to send the HTTP GET request to configure the postStart callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port over which you want to send the HTTP GET request to configure the postStart callback function.
	//
	// example:
	//
	// 5050
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP GET request that you want to send to configure the postStart callback function. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTPS
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP socket that you want to use to configure the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP socket that you want to use to configure the postStart callback function.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The commands that you can run within the container to configure the preStop callback function.
	LifecyclePreStopHandlerExecs []*string `json:"LifecyclePreStopHandlerExecs,omitempty" xml:"LifecyclePreStopHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which you want to send the HTTP GET request to configure the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The path to which you want to send the HTTP GET request to configure the preStop callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port over which you want to send the HTTP GET request to configure the preStop callback function.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP GET request that you want to send to configure the preStop callback function. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP socket that you want to use to configure the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP socket that you want to use to configure the preStop callback function.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size per container. Unit: GiB.
	//
	// example:
	//
	// 0.5
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container image.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ports.
	Ports []*ModifyEciScalingConfigurationRequestContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// Specifies whether the container allocates buffer resources to standard input streams during its active runtime. If you do not specify this parameter, an end-of-file (EOF) error occurs when standard input streams in the container are read.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether standard input streams remain connected during multiple sessions when StdinOnce is set to true.
	//
	// If you set StdinOnce to true, standard input streams are connected after the container is started, and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected and remain disconnected until the container is restarted.
	//
	// example:
	//
	// false
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable Interaction. Default value: false.
	//
	// If the command is a /bin/bash command, set this parameter to true.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The volume mounts of the container.
	VolumeMounts []*ModifyEciScalingConfigurationRequestContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/local/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainers) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainers) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLivenessProbe() *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	return s.LivenessProbe
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetReadinessProbe() *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	return s.ReadinessProbe
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetSecurityContext() *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	return s.SecurityContext
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetArgs() []*string {
	return s.Args
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetCommands() []*string {
	return s.Commands
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetEnvironmentVars() []*ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	return s.EnvironmentVars
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetImage() *string {
	return s.Image
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerExecs() []*string {
	return s.LifecyclePostStartHandlerExecs
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetHost() *string {
	return s.LifecyclePostStartHandlerHttpGetHost
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetPath() *string {
	return s.LifecyclePostStartHandlerHttpGetPath
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetPort() *int32 {
	return s.LifecyclePostStartHandlerHttpGetPort
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetScheme() *string {
	return s.LifecyclePostStartHandlerHttpGetScheme
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerTcpSocketHost() *string {
	return s.LifecyclePostStartHandlerTcpSocketHost
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerTcpSocketPort() *int32 {
	return s.LifecyclePostStartHandlerTcpSocketPort
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerExecs() []*string {
	return s.LifecyclePreStopHandlerExecs
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetHost() *string {
	return s.LifecyclePreStopHandlerHttpGetHost
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetPath() *string {
	return s.LifecyclePreStopHandlerHttpGetPath
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetPort() *int32 {
	return s.LifecyclePreStopHandlerHttpGetPort
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetScheme() *string {
	return s.LifecyclePreStopHandlerHttpGetScheme
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerTcpSocketHost() *string {
	return s.LifecyclePreStopHandlerTcpSocketHost
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerTcpSocketPort() *int32 {
	return s.LifecyclePreStopHandlerTcpSocketPort
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetPorts() []*ModifyEciScalingConfigurationRequestContainersPorts {
	return s.Ports
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetStdin() *bool {
	return s.Stdin
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetStdinOnce() *bool {
	return s.StdinOnce
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetTty() *bool {
	return s.Tty
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetVolumeMounts() []*ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	return s.VolumeMounts
}

func (s *ModifyEciScalingConfigurationRequestContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLivenessProbe(v *ModifyEciScalingConfigurationRequestContainersLivenessProbe) *ModifyEciScalingConfigurationRequestContainers {
	s.LivenessProbe = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetReadinessProbe(v *ModifyEciScalingConfigurationRequestContainersReadinessProbe) *ModifyEciScalingConfigurationRequestContainers {
	s.ReadinessProbe = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetSecurityContext(v *ModifyEciScalingConfigurationRequestContainersSecurityContext) *ModifyEciScalingConfigurationRequestContainers {
	s.SecurityContext = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetArgs(v []*string) *ModifyEciScalingConfigurationRequestContainers {
	s.Args = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestContainers {
	s.Commands = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetCpu(v float32) *ModifyEciScalingConfigurationRequestContainers {
	s.Cpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetEnvironmentVars(v []*ModifyEciScalingConfigurationRequestContainersEnvironmentVars) *ModifyEciScalingConfigurationRequestContainers {
	s.EnvironmentVars = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetGpu(v int32) *ModifyEciScalingConfigurationRequestContainers {
	s.Gpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetImage(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.Image = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetImagePullPolicy(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerExecs(v []*string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerExecs = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetHost(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetPath(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetPort(v int32) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetScheme(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerTcpSocketHost(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerExecs(v []*string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerExecs = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetHost(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetPath(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetPort(v int32) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetScheme(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerTcpSocketHost(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *ModifyEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetMemory(v float32) *ModifyEciScalingConfigurationRequestContainers {
	s.Memory = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetName(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetPorts(v []*ModifyEciScalingConfigurationRequestContainersPorts) *ModifyEciScalingConfigurationRequestContainers {
	s.Ports = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetStdin(v bool) *ModifyEciScalingConfigurationRequestContainers {
	s.Stdin = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetStdinOnce(v bool) *ModifyEciScalingConfigurationRequestContainers {
	s.StdinOnce = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetTty(v bool) *ModifyEciScalingConfigurationRequestContainers {
	s.Tty = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetVolumeMounts(v []*ModifyEciScalingConfigurationRequestContainersVolumeMounts) *ModifyEciScalingConfigurationRequestContainers {
	s.VolumeMounts = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) SetWorkingDir(v string) *ModifyEciScalingConfigurationRequestContainers {
	s.WorkingDir = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainers) Validate() error {
	if s.LivenessProbe != nil {
		if err := s.LivenessProbe.Validate(); err != nil {
			return err
		}
	}
	if s.ReadinessProbe != nil {
		if err := s.ReadinessProbe.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityContext != nil {
		if err := s.SecurityContext.Validate(); err != nil {
			return err
		}
	}
	if s.EnvironmentVars != nil {
		for _, item := range s.EnvironmentVars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VolumeMounts != nil {
		for _, item := range s.VolumeMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbe struct {
	Exec                *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                                `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                                `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	TimeoutSeconds      *int32                                                                `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbe) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetExec() *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec {
	return s.Exec
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetHttpGet() *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	return s.HttpGet
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetTcpSocket() *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket {
	return s.TcpSocket
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetExec(v *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.Exec = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetFailureThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetHttpGet(v *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetInitialDelaySeconds(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetPeriodSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetSuccessThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetTcpSocket(v *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) SetTimeoutSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbe) Validate() error {
	if s.Exec != nil {
		if err := s.Exec.Validate(); err != nil {
			return err
		}
	}
	if s.HttpGet != nil {
		if err := s.HttpGet.Validate(); err != nil {
			return err
		}
	}
	if s.TcpSocket != nil {
		if err := s.TcpSocket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) GetCommands() []*string {
	return s.Commands
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec {
	s.Commands = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeExec) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GetPath() *string {
	return s.Path
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GetPort() *int32 {
	return s.Port
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GetScheme() *string {
	return s.Scheme
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPath(v string) *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetScheme(v string) *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeHttpGet) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) GetPort() *int32 {
	return s.Port
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbe struct {
	Exec                *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                                 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                                 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	TimeoutSeconds      *int32                                                                 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbe) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetExec() *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec {
	return s.Exec
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetHttpGet() *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	return s.HttpGet
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetTcpSocket() *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket {
	return s.TcpSocket
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetExec(v *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.Exec = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetFailureThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetHttpGet(v *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetInitialDelaySeconds(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetPeriodSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetSuccessThreshold(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetTcpSocket(v *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) SetTimeoutSeconds(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbe) Validate() error {
	if s.Exec != nil {
		if err := s.Exec.Validate(); err != nil {
			return err
		}
	}
	if s.HttpGet != nil {
		if err := s.HttpGet.Validate(); err != nil {
			return err
		}
	}
	if s.TcpSocket != nil {
		if err := s.TcpSocket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) GetCommands() []*string {
	return s.Commands
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec {
	s.Commands = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeExec) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GetPath() *string {
	return s.Path
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GetPort() *int32 {
	return s.Port
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GetScheme() *string {
	return s.Scheme
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPath(v string) *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetScheme(v string) *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeHttpGet) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) GetPort() *int32 {
	return s.Port
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersSecurityContext struct {
	Capability             *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                    `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                   `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContext) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) GetCapability() *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability {
	return s.Capability
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) GetReadOnlyRootFilesystem() *bool {
	return s.ReadOnlyRootFilesystem
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) GetRunAsUser() *int64 {
	return s.RunAsUser
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) SetCapability(v *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) SetRunAsUser(v int64) *ModifyEciScalingConfigurationRequestContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContext) Validate() error {
	if s.Capability != nil {
		if err := s.Capability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) GetAdds() []*string {
	return s.Adds
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) SetAdds(v []*string) *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability {
	s.Adds = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersSecurityContextCapability) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersEnvironmentVars struct {
	FieldRef *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
	// The name of the environment variable. The name can be 1 to 128 characters in length, and can contain letters, underscores (_), and digits. The name cannot start with a digit. Specify the value in the `[0-9a-zA-Z]` format.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value can be up to 256 characters in length.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) GetFieldRef() *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef {
	return s.FieldRef
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) SetFieldRef(v *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) *ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	s.FieldRef = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) SetKey(v string) *ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) SetValue(v string) *ModifyEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVars) Validate() error {
	if s.FieldRef != nil {
		if err := s.FieldRef.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) GetFieldPath() *string {
	return s.FieldPath
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) SetFieldPath(v string) *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef {
	s.FieldPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersEnvironmentVarsFieldRef) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersPorts struct {
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersPorts) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersPorts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) GetPort() *int32 {
	return s.Port
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) SetPort(v int32) *ModifyEciScalingConfigurationRequestContainersPorts {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) SetProtocol(v string) *ModifyEciScalingConfigurationRequestContainersPorts {
	s.Protocol = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersPorts) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestContainersVolumeMounts struct {
	// The directory within the container onto which you want to mount the volume.
	//
	// >  The information stored within this directory is overwritten by the data on the mounted volume. Exercise caution when you specify this parameter.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation enables volumes mounted on one container to be shared among other containers within the same pod or across distinct pods residing on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed on the volume or its subdirectories do not propagate to the volume.
	//
	// 	- HostToCotainer: Subsequent mounts executed on the volume or its subdirectories propagate to the volume.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed on the volume or its subdirectories propagate to the volume. In addition, volume mounts executed on the container propagate back to the underlying instance and to all containers across every pod that uses the same volume.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The volume name. The value of this parameter is the same as the name of the volume that is mounted to containers.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The volume subdirectory.
	//
	// example:
	//
	// data2/
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestContainersVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetMountPath(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetMountPropagation(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetName(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetReadOnly(v bool) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) SetSubPath(v string) *ModifyEciScalingConfigurationRequestContainersVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestContainersVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestDnsConfigOptions struct {
	// The variable name of the option.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The variable value of the option.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestDnsConfigOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) GetValue() *string {
	return s.Value
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) SetName(v string) *ModifyEciScalingConfigurationRequestDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) SetValue(v string) *ModifyEciScalingConfigurationRequestDnsConfigOptions {
	s.Value = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestDnsConfigOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestHostAliases struct {
	// The names of the hosts that you want to add.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The IP address that you want to add.
	//
	// example:
	//
	// 192.0.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestHostAliases) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestHostAliases) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) GetHostnames() []*string {
	return s.Hostnames
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) GetIp() *string {
	return s.Ip
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) SetHostnames(v []*string) *ModifyEciScalingConfigurationRequestHostAliases {
	s.Hostnames = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) SetIp(v string) *ModifyEciScalingConfigurationRequestHostAliases {
	s.Ip = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestHostAliases) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestImageRegistryCredentials struct {
	// The password of the image repository.
	//
	// example:
	//
	// yourpaasword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The address of the image repository.
	//
	// example:
	//
	// registry-vpc.cn-shanghai.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username of the image repository.
	//
	// example:
	//
	// yourusername
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestImageRegistryCredentials) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) GetPassword() *string {
	return s.Password
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) GetServer() *string {
	return s.Server
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) GetUserName() *string {
	return s.UserName
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) SetPassword(v string) *ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) SetServer(v string) *ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) SetUserName(v string) *ModifyEciScalingConfigurationRequestImageRegistryCredentials {
	s.UserName = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestImageRegistryCredentials) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestInitContainers struct {
	SecurityContext *ModifyEciScalingConfigurationRequestInitContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The container startup arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The commands that you can run to start the init container.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs per init container.
	//
	// example:
	//
	// 0.5
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs per init container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the init container.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Image pulling is performed each time instances are created.
	//
	// 	- IfNotPresent: Image pulling is performed as needed. On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The environment variables of the init container.
	InitContainerEnvironmentVars []*ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	// The ports of the init container.
	InitContainerPorts []*ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	// The volume mounts of the init container.
	InitContainerVolumeMounts []*ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	// The memory size per init container. Unit: GiB.
	//
	// example:
	//
	// 1.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the init container.
	//
	// example:
	//
	// test-init
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The working directory.
	//
	// example:
	//
	// /usr/local
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainers) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainers) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetSecurityContext() *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	return s.SecurityContext
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetArgs() []*string {
	return s.Args
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetCommands() []*string {
	return s.Commands
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetImage() *string {
	return s.Image
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetInitContainerEnvironmentVars() []*ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	return s.InitContainerEnvironmentVars
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetInitContainerPorts() []*ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts {
	return s.InitContainerPorts
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetInitContainerVolumeMounts() []*ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	return s.InitContainerVolumeMounts
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetSecurityContext(v *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) *ModifyEciScalingConfigurationRequestInitContainers {
	s.SecurityContext = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetArgs(v []*string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Args = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetCommands(v []*string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Commands = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetCpu(v float32) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Cpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetGpu(v int32) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Gpu = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetImage(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Image = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetImagePullPolicy(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetInitContainerEnvironmentVars(v []*ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) *ModifyEciScalingConfigurationRequestInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetInitContainerPorts(v []*ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) *ModifyEciScalingConfigurationRequestInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetInitContainerVolumeMounts(v []*ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) *ModifyEciScalingConfigurationRequestInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetMemory(v float32) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Memory = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetName(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) SetWorkingDir(v string) *ModifyEciScalingConfigurationRequestInitContainers {
	s.WorkingDir = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainers) Validate() error {
	if s.SecurityContext != nil {
		if err := s.SecurityContext.Validate(); err != nil {
			return err
		}
	}
	if s.InitContainerEnvironmentVars != nil {
		for _, item := range s.InitContainerEnvironmentVars {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InitContainerPorts != nil {
		for _, item := range s.InitContainerPorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InitContainerVolumeMounts != nil {
		for _, item := range s.InitContainerVolumeMounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestInitContainersSecurityContext struct {
	Capability             *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                        `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                       `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContext) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) GetCapability() *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability {
	return s.Capability
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) GetReadOnlyRootFilesystem() *bool {
	return s.ReadOnlyRootFilesystem
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) GetRunAsUser() *int64 {
	return s.RunAsUser
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) SetCapability(v *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) SetRunAsUser(v int64) *ModifyEciScalingConfigurationRequestInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContext) Validate() error {
	if s.Capability != nil {
		if err := s.Capability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) GetAdds() []*string {
	return s.Adds
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) SetAdds(v []*string) *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersSecurityContextCapability) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars struct {
	FieldRef *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef `json:"FieldRef,omitempty" xml:"FieldRef,omitempty" type:"Struct"`
	// The name of the environment variable. The name can be 1 to 128 characters in length, and can contain letters, underscores (_), and digits. The name cannot start with a digit. Specify the value in the `[0-9a-zA-Z]` format.
	//
	// example:
	//
	// Path
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value can be up to 256 characters in length.
	//
	// example:
	//
	// /usr/bin/
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GetFieldRef() *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef {
	return s.FieldRef
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetFieldRef(v *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.FieldRef = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetKey(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetValue(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) Validate() error {
	if s.FieldRef != nil {
		if err := s.FieldRef.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef struct {
	FieldPath *string `json:"FieldPath,omitempty" xml:"FieldPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) GetFieldPath() *string {
	return s.FieldPath
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) SetFieldPath(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef {
	s.FieldPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVarsFieldRef) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts struct {
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) GetPort() *int32 {
	return s.Port
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) SetPort(v int32) *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) SetProtocol(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerPorts) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts struct {
	// The directory within the init container onto which you want to mount the volume.
	//
	// >  The information stored within this directory is overwritten by the data on the mounted volume. Exercise caution when you specify this parameter.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation enables volumes mounted on one container to be shared among other containers within the same pod or across distinct pods residing on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed on the volume or its subdirectories do not propagate to the volume.
	//
	// 	- HostToCotainer: Subsequent mounts executed on the volume or its subdirectories propagate to the volume.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed on the volume or its subdirectories propagate to the volume. In addition, volume mounts executed on the container propagate back to the underlying instance and to all containers across every pod that uses the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the mount path is read-only.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The volume subdirectory. The pod can mount different directories of the same volume to different subdirectories of init containers.
	//
	// example:
	//
	// Always
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPath(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetName(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetSubPath(v string) *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestSecurityContextSysCtls struct {
	// The variable name of the security context in which the elastic container instance runs.
	//
	// example:
	//
	// kernel.msgmax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The variable value of the security context in which the elastic container instance runs.
	//
	// example:
	//
	// 65536
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestSecurityContextSysCtls) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestSecurityContextSysCtls) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) GetValue() *string {
	return s.Value
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) SetName(v string) *ModifyEciScalingConfigurationRequestSecurityContextSysCtls {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) SetValue(v string) *ModifyEciScalingConfigurationRequestSecurityContextSysCtls {
	s.Value = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestSecurityContextSysCtls) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestTags) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestTags) GetKey() *string {
	return s.Key
}

func (s *ModifyEciScalingConfigurationRequestTags) GetValue() *string {
	return s.Value
}

func (s *ModifyEciScalingConfigurationRequestTags) SetKey(v string) *ModifyEciScalingConfigurationRequestTags {
	s.Key = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestTags) SetValue(v string) *ModifyEciScalingConfigurationRequestTags {
	s.Value = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestTags) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestVolumes struct {
	DiskVolume     *ModifyEciScalingConfigurationRequestVolumesDiskVolume     `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" type:"Struct"`
	EmptyDirVolume *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" type:"Struct"`
	FlexVolume     *ModifyEciScalingConfigurationRequestVolumesFlexVolume     `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" type:"Struct"`
	HostPathVolume *ModifyEciScalingConfigurationRequestVolumesHostPathVolume `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" type:"Struct"`
	NFSVolume      *ModifyEciScalingConfigurationRequestVolumesNFSVolume      `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" type:"Struct"`
	// The paths to the configuration files.
	ConfigFileVolumeConfigFileToPath []*ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath `json:"ConfigFileVolumeConfigFileToPath,omitempty" xml:"ConfigFileVolumeConfigFileToPath,omitempty" type:"Repeated"`
	// The default permissions on the ConfigFile volume.
	//
	// example:
	//
	// 0644
	ConfigFileVolumeDefaultMode *int32 `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	// The volume name.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the Host directory. Examples: File, Directory, and Socket.
	//
	// example:
	//
	// EmptyDirVolume
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumes) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumes) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetDiskVolume() *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	return s.DiskVolume
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetEmptyDirVolume() *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume {
	return s.EmptyDirVolume
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetFlexVolume() *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	return s.FlexVolume
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetHostPathVolume() *ModifyEciScalingConfigurationRequestVolumesHostPathVolume {
	return s.HostPathVolume
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetNFSVolume() *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	return s.NFSVolume
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetConfigFileVolumeConfigFileToPath() []*ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	return s.ConfigFileVolumeConfigFileToPath
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetConfigFileVolumeDefaultMode() *int32 {
	return s.ConfigFileVolumeDefaultMode
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetName() *string {
	return s.Name
}

func (s *ModifyEciScalingConfigurationRequestVolumes) GetType() *string {
	return s.Type
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetDiskVolume(v *ModifyEciScalingConfigurationRequestVolumesDiskVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.DiskVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetEmptyDirVolume(v *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.EmptyDirVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetFlexVolume(v *ModifyEciScalingConfigurationRequestVolumesFlexVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.FlexVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetHostPathVolume(v *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.HostPathVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetNFSVolume(v *ModifyEciScalingConfigurationRequestVolumesNFSVolume) *ModifyEciScalingConfigurationRequestVolumes {
	s.NFSVolume = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetConfigFileVolumeConfigFileToPath(v []*ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) *ModifyEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeConfigFileToPath = v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetConfigFileVolumeDefaultMode(v int32) *ModifyEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetName(v string) *ModifyEciScalingConfigurationRequestVolumes {
	s.Name = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) SetType(v string) *ModifyEciScalingConfigurationRequestVolumes {
	s.Type = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumes) Validate() error {
	if s.DiskVolume != nil {
		if err := s.DiskVolume.Validate(); err != nil {
			return err
		}
	}
	if s.EmptyDirVolume != nil {
		if err := s.EmptyDirVolume.Validate(); err != nil {
			return err
		}
	}
	if s.FlexVolume != nil {
		if err := s.FlexVolume.Validate(); err != nil {
			return err
		}
	}
	if s.HostPathVolume != nil {
		if err := s.HostPathVolume.Validate(); err != nil {
			return err
		}
	}
	if s.NFSVolume != nil {
		if err := s.NFSVolume.Validate(); err != nil {
			return err
		}
	}
	if s.ConfigFileVolumeConfigFileToPath != nil {
		for _, item := range s.ConfigFileVolumeConfigFileToPath {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyEciScalingConfigurationRequestVolumesDiskVolume struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesDiskVolume) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesDiskVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) GetFsType() *string {
	return s.FsType
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) SetDiskId(v string) *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskId = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) SetDiskSize(v int32) *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) SetFsType(v string) *ModifyEciScalingConfigurationRequestVolumesDiskVolume {
	s.FsType = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesDiskVolume) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume struct {
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) GetMedium() *string {
	return s.Medium
}

func (s *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) GetSizeLimit() *string {
	return s.SizeLimit
}

func (s *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) SetMedium(v string) *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume {
	s.Medium = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) SetSizeLimit(v string) *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume {
	s.SizeLimit = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesEmptyDirVolume) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestVolumesFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesFlexVolume) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesFlexVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) GetDriver() *string {
	return s.Driver
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) GetFsType() *string {
	return s.FsType
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) GetOptions() *string {
	return s.Options
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) SetDriver(v string) *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	s.Driver = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) SetFsType(v string) *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	s.FsType = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) SetOptions(v string) *ModifyEciScalingConfigurationRequestVolumesFlexVolume {
	s.Options = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesFlexVolume) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestVolumesHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesHostPathVolume) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesHostPathVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) GetPath() *string {
	return s.Path
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) GetType() *string {
	return s.Type
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) SetPath(v string) *ModifyEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) SetType(v string) *ModifyEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Type = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesHostPathVolume) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestVolumesNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesNFSVolume) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesNFSVolume) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) GetPath() *string {
	return s.Path
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) GetServer() *string {
	return s.Server
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) SetPath(v string) *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) SetReadOnly(v bool) *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) SetServer(v string) *ModifyEciScalingConfigurationRequestVolumesNFSVolume {
	s.Server = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesNFSVolume) Validate() error {
	return dara.Validate(s)
}

type ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath struct {
	// The content of the configuration file (32 KB).
	//
	// example:
	//
	// bGl1bWk=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The permissions on the ConfigFile volume.
	//
	// example:
	//
	// 0644
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The relative path to the configuration file.
	//
	// example:
	//
	// /usr/bin/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) GetContent() *string {
	return s.Content
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) GetMode() *int32 {
	return s.Mode
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) GetPath() *string {
	return s.Path
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) SetContent(v string) *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	s.Content = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) SetMode(v int32) *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	s.Mode = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) SetPath(v string) *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath {
	s.Path = &v
	return s
}

func (s *ModifyEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPath) Validate() error {
	return dara.Validate(s)
}
