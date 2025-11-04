// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEciScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrRegistryInfos(v []*CreateEciScalingConfigurationRequestAcrRegistryInfos) *CreateEciScalingConfigurationRequest
	GetAcrRegistryInfos() []*CreateEciScalingConfigurationRequestAcrRegistryInfos
	SetActiveDeadlineSeconds(v int64) *CreateEciScalingConfigurationRequest
	GetActiveDeadlineSeconds() *int64
	SetAutoCreateEip(v bool) *CreateEciScalingConfigurationRequest
	GetAutoCreateEip() *bool
	SetAutoMatchImageCache(v bool) *CreateEciScalingConfigurationRequest
	GetAutoMatchImageCache() *bool
	SetContainerGroupName(v string) *CreateEciScalingConfigurationRequest
	GetContainerGroupName() *string
	SetContainers(v []*CreateEciScalingConfigurationRequestContainers) *CreateEciScalingConfigurationRequest
	GetContainers() []*CreateEciScalingConfigurationRequestContainers
	SetCostOptimization(v bool) *CreateEciScalingConfigurationRequest
	GetCostOptimization() *bool
	SetCpu(v float32) *CreateEciScalingConfigurationRequest
	GetCpu() *float32
	SetCpuOptionsCore(v int32) *CreateEciScalingConfigurationRequest
	GetCpuOptionsCore() *int32
	SetCpuOptionsThreadsPerCore(v int32) *CreateEciScalingConfigurationRequest
	GetCpuOptionsThreadsPerCore() *int32
	SetDataCacheBucket(v string) *CreateEciScalingConfigurationRequest
	GetDataCacheBucket() *string
	SetDataCacheBurstingEnabled(v bool) *CreateEciScalingConfigurationRequest
	GetDataCacheBurstingEnabled() *bool
	SetDataCachePL(v string) *CreateEciScalingConfigurationRequest
	GetDataCachePL() *string
	SetDataCacheProvisionedIops(v int32) *CreateEciScalingConfigurationRequest
	GetDataCacheProvisionedIops() *int32
	SetDescription(v string) *CreateEciScalingConfigurationRequest
	GetDescription() *string
	SetDnsConfigNameServers(v []*string) *CreateEciScalingConfigurationRequest
	GetDnsConfigNameServers() []*string
	SetDnsConfigOptions(v []*CreateEciScalingConfigurationRequestDnsConfigOptions) *CreateEciScalingConfigurationRequest
	GetDnsConfigOptions() []*CreateEciScalingConfigurationRequestDnsConfigOptions
	SetDnsConfigSearchs(v []*string) *CreateEciScalingConfigurationRequest
	GetDnsConfigSearchs() []*string
	SetDnsPolicy(v string) *CreateEciScalingConfigurationRequest
	GetDnsPolicy() *string
	SetEgressBandwidth(v int64) *CreateEciScalingConfigurationRequest
	GetEgressBandwidth() *int64
	SetEipBandwidth(v int32) *CreateEciScalingConfigurationRequest
	GetEipBandwidth() *int32
	SetEnableSls(v bool) *CreateEciScalingConfigurationRequest
	GetEnableSls() *bool
	SetEphemeralStorage(v int32) *CreateEciScalingConfigurationRequest
	GetEphemeralStorage() *int32
	SetGpuDriverVersion(v string) *CreateEciScalingConfigurationRequest
	GetGpuDriverVersion() *string
	SetHostAliases(v []*CreateEciScalingConfigurationRequestHostAliases) *CreateEciScalingConfigurationRequest
	GetHostAliases() []*CreateEciScalingConfigurationRequestHostAliases
	SetHostName(v string) *CreateEciScalingConfigurationRequest
	GetHostName() *string
	SetImageRegistryCredentials(v []*CreateEciScalingConfigurationRequestImageRegistryCredentials) *CreateEciScalingConfigurationRequest
	GetImageRegistryCredentials() []*CreateEciScalingConfigurationRequestImageRegistryCredentials
	SetImageSnapshotId(v string) *CreateEciScalingConfigurationRequest
	GetImageSnapshotId() *string
	SetIngressBandwidth(v int64) *CreateEciScalingConfigurationRequest
	GetIngressBandwidth() *int64
	SetInitContainers(v []*CreateEciScalingConfigurationRequestInitContainers) *CreateEciScalingConfigurationRequest
	GetInitContainers() []*CreateEciScalingConfigurationRequestInitContainers
	SetInstanceFamilyLevel(v string) *CreateEciScalingConfigurationRequest
	GetInstanceFamilyLevel() *string
	SetInstanceTypes(v []*string) *CreateEciScalingConfigurationRequest
	GetInstanceTypes() []*string
	SetIpv6AddressCount(v int32) *CreateEciScalingConfigurationRequest
	GetIpv6AddressCount() *int32
	SetLoadBalancerWeight(v int32) *CreateEciScalingConfigurationRequest
	GetLoadBalancerWeight() *int32
	SetMemory(v float32) *CreateEciScalingConfigurationRequest
	GetMemory() *float32
	SetNtpServers(v []*string) *CreateEciScalingConfigurationRequest
	GetNtpServers() []*string
	SetOwnerId(v int64) *CreateEciScalingConfigurationRequest
	GetOwnerId() *int64
	SetRamRoleName(v string) *CreateEciScalingConfigurationRequest
	GetRamRoleName() *string
	SetResourceGroupId(v string) *CreateEciScalingConfigurationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateEciScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetRestartPolicy(v string) *CreateEciScalingConfigurationRequest
	GetRestartPolicy() *string
	SetScalingConfigurationName(v string) *CreateEciScalingConfigurationRequest
	GetScalingConfigurationName() *string
	SetScalingGroupId(v string) *CreateEciScalingConfigurationRequest
	GetScalingGroupId() *string
	SetSecurityContextSysctls(v []*CreateEciScalingConfigurationRequestSecurityContextSysctls) *CreateEciScalingConfigurationRequest
	GetSecurityContextSysctls() []*CreateEciScalingConfigurationRequestSecurityContextSysctls
	SetSecurityGroupId(v string) *CreateEciScalingConfigurationRequest
	GetSecurityGroupId() *string
	SetSpotPriceLimit(v float32) *CreateEciScalingConfigurationRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateEciScalingConfigurationRequest
	GetSpotStrategy() *string
	SetTags(v []*CreateEciScalingConfigurationRequestTags) *CreateEciScalingConfigurationRequest
	GetTags() []*CreateEciScalingConfigurationRequestTags
	SetTerminationGracePeriodSeconds(v int64) *CreateEciScalingConfigurationRequest
	GetTerminationGracePeriodSeconds() *int64
	SetVolumes(v []*CreateEciScalingConfigurationRequestVolumes) *CreateEciScalingConfigurationRequest
	GetVolumes() []*CreateEciScalingConfigurationRequestVolumes
}

type CreateEciScalingConfigurationRequest struct {
	// The Container Registry Enterprise Edition instances.
	AcrRegistryInfos []*CreateEciScalingConfigurationRequestAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
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
	// Specifies whether to automatically match the image cache. Valid values:
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
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The name series of elastic container instances.
	//
	// If you want to use an ordered instance name, specify the value for this parameter in the following format: name_prefix[begin_number,bits]name_suffix.
	//
	// example:
	//
	// nginx-test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The containers per elastic container instance.
	Containers []*CreateEciScalingConfigurationRequestContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// Specifies whether to enable the Cost Optimization feature. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The number of physical CPU cores. You can specify this parameter for specific instance types. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// The number of threads per core. You can specify this parameter for specific instance types. A value of 1 specifies that Hyper-Threading is disabled. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The bucket that you want to use to store data caches.
	//
	// example:
	//
	// default
	DataCacheBucket *string `json:"DataCacheBucket,omitempty" xml:"DataCacheBucket,omitempty"`
	// Specifies whether to enable the performance burst feature when ESSD AutoPL disks are used to store data caches. Valid values:
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
	// The PL of the cloud disk that you want to use to store data caches. We recommend that you use ESSDs. Valid values:
	//
	// 	- PL0: An ESSD can provide up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can provide up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can provide up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD can provide up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// >  For more information about ESSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	DataCachePL *string `json:"DataCachePL,omitempty" xml:"DataCachePL,omitempty"`
	// The provisioned IOPS of the ESSD AutoPL disk that you want to use to store data caches. Valid values: 0 to min{50,000, 1,000 × *Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × *Capacity, 50,000}.
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
	// The IP addresses of the DNS servers.
	DnsConfigNameServers []*string `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	// The options. Each option is a name-value pair. The value in the name-value pair is optional.
	DnsConfigOptions []*CreateEciScalingConfigurationRequestDnsConfigOptions `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	// The search domains of the DNS servers.
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
	// The EIP bandwidth. Default value: 5. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// >  This parameter is unavailable for use.
	//
	// example:
	//
	// false
	EnableSls *bool `json:"EnableSls,omitempty" xml:"EnableSls,omitempty"`
	// The size of the temporary storage space. By default, an Enterprise SSD (ESSD) of performance level 1 (PL1) is used. Unit: GiB.
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
	// The custom hostnames of the containers.
	HostAliases []*CreateEciScalingConfigurationRequestHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The hostname series of elastic container instances.
	//
	// example:
	//
	// test
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image repositories.
	ImageRegistryCredentials []*CreateEciScalingConfigurationRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
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
	InitContainers []*CreateEciScalingConfigurationRequestInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The level of the instance family. You can specify this parameter to match the available instance types. This parameter takes effect only if you set `CostOptimization` to true. Valid values:
	//
	// 	- EntryLevel: entry level (shared instance types). Instance types of this level are the most cost-effective, but may not ensure stable computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise level. Instance types of this level provide stable performance and dedicated resources, and are suitable for business scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit-based entry level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview](https://help.aliyun.com/document_detail/59977.html) of burstable instances.
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The ECS instance types that you want to use to create elastic container instances. You can specify up to five ECS instance types.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of IPv6 addresses.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The load balancing weight of each elastic container instance. Valid values: 1 to 100.
	//
	// Default value: 50.
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
	// The endpoints of the Network Time Protocol (NTP) servers.
	NtpServers []*string `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	OwnerId    *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the instance Resource Access Management (RAM) role. Elastic container instances and Elastic Compute Service (ECS) instances can share the same RAM role. For more information, see [RAM roles](https://help.aliyun.com/document_detail/61175.html).
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
	// The restart policy of elastic container instances. Valid values:
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
	// The name of the scaling configuration. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter or a digit.
	//
	// The name of the scaling configuration must be unique in a scaling group within a region. If you do not specify this parameter, the value of ScalingConfigurationId is used.
	//
	// example:
	//
	// scalingconfig****
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The ID of the scaling group to which the scaling configuration belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The system information of the security context in which the elastic container instance runs.
	SecurityContextSysctls []*CreateEciScalingConfigurationRequestSecurityContextSysctls `json:"SecurityContextSysctls,omitempty" xml:"SecurityContextSysctls,omitempty" type:"Repeated"`
	// The ID of the security group to which elastic container instances belong. Elastic container instances that belong to the same security group can communicate with each other.
	//
	// If you do not specify a security group, the system uses the default security group in the region that you selected. Make sure that the inbound rules of the security group contain the protocols and port numbers of the containers that you want to expose. If you do not have a default security group in the region, the system creates a default security group, and then adds the container protocols and port numbers that you specified to the inbound rules of the security group.
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
	// The tags of elastic container instances. Tags must be specified as key-value pairs. You can specify up to 20 tags for each elastic container instance. When you specify Key and Value, take note of the following items:
	//
	// 	- A tag key can be up to 64 characters in length. The key cannot start with acs: or aliyun or contain `http://` or `https://`. You cannot specify an empty string as a tag key.
	//
	// 	- A tag value can be up to 128 characters in length. The value cannot start with acs: or aliyun or contain `http://` or `https://`. You can specify an empty string as a tag value.
	Tags []*CreateEciScalingConfigurationRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The buffer time during which a program handles operations before the program stops. Unit: seconds.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The volumes.
	Volumes []*CreateEciScalingConfigurationRequestVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequest) GetAcrRegistryInfos() []*CreateEciScalingConfigurationRequestAcrRegistryInfos {
	return s.AcrRegistryInfos
}

func (s *CreateEciScalingConfigurationRequest) GetActiveDeadlineSeconds() *int64 {
	return s.ActiveDeadlineSeconds
}

func (s *CreateEciScalingConfigurationRequest) GetAutoCreateEip() *bool {
	return s.AutoCreateEip
}

func (s *CreateEciScalingConfigurationRequest) GetAutoMatchImageCache() *bool {
	return s.AutoMatchImageCache
}

func (s *CreateEciScalingConfigurationRequest) GetContainerGroupName() *string {
	return s.ContainerGroupName
}

func (s *CreateEciScalingConfigurationRequest) GetContainers() []*CreateEciScalingConfigurationRequestContainers {
	return s.Containers
}

func (s *CreateEciScalingConfigurationRequest) GetCostOptimization() *bool {
	return s.CostOptimization
}

func (s *CreateEciScalingConfigurationRequest) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateEciScalingConfigurationRequest) GetCpuOptionsCore() *int32 {
	return s.CpuOptionsCore
}

func (s *CreateEciScalingConfigurationRequest) GetCpuOptionsThreadsPerCore() *int32 {
	return s.CpuOptionsThreadsPerCore
}

func (s *CreateEciScalingConfigurationRequest) GetDataCacheBucket() *string {
	return s.DataCacheBucket
}

func (s *CreateEciScalingConfigurationRequest) GetDataCacheBurstingEnabled() *bool {
	return s.DataCacheBurstingEnabled
}

func (s *CreateEciScalingConfigurationRequest) GetDataCachePL() *string {
	return s.DataCachePL
}

func (s *CreateEciScalingConfigurationRequest) GetDataCacheProvisionedIops() *int32 {
	return s.DataCacheProvisionedIops
}

func (s *CreateEciScalingConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEciScalingConfigurationRequest) GetDnsConfigNameServers() []*string {
	return s.DnsConfigNameServers
}

func (s *CreateEciScalingConfigurationRequest) GetDnsConfigOptions() []*CreateEciScalingConfigurationRequestDnsConfigOptions {
	return s.DnsConfigOptions
}

func (s *CreateEciScalingConfigurationRequest) GetDnsConfigSearchs() []*string {
	return s.DnsConfigSearchs
}

func (s *CreateEciScalingConfigurationRequest) GetDnsPolicy() *string {
	return s.DnsPolicy
}

func (s *CreateEciScalingConfigurationRequest) GetEgressBandwidth() *int64 {
	return s.EgressBandwidth
}

func (s *CreateEciScalingConfigurationRequest) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *CreateEciScalingConfigurationRequest) GetEnableSls() *bool {
	return s.EnableSls
}

func (s *CreateEciScalingConfigurationRequest) GetEphemeralStorage() *int32 {
	return s.EphemeralStorage
}

func (s *CreateEciScalingConfigurationRequest) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *CreateEciScalingConfigurationRequest) GetHostAliases() []*CreateEciScalingConfigurationRequestHostAliases {
	return s.HostAliases
}

func (s *CreateEciScalingConfigurationRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateEciScalingConfigurationRequest) GetImageRegistryCredentials() []*CreateEciScalingConfigurationRequestImageRegistryCredentials {
	return s.ImageRegistryCredentials
}

func (s *CreateEciScalingConfigurationRequest) GetImageSnapshotId() *string {
	return s.ImageSnapshotId
}

func (s *CreateEciScalingConfigurationRequest) GetIngressBandwidth() *int64 {
	return s.IngressBandwidth
}

func (s *CreateEciScalingConfigurationRequest) GetInitContainers() []*CreateEciScalingConfigurationRequestInitContainers {
	return s.InitContainers
}

func (s *CreateEciScalingConfigurationRequest) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *CreateEciScalingConfigurationRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateEciScalingConfigurationRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateEciScalingConfigurationRequest) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *CreateEciScalingConfigurationRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateEciScalingConfigurationRequest) GetNtpServers() []*string {
	return s.NtpServers
}

func (s *CreateEciScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateEciScalingConfigurationRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateEciScalingConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEciScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateEciScalingConfigurationRequest) GetRestartPolicy() *string {
	return s.RestartPolicy
}

func (s *CreateEciScalingConfigurationRequest) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *CreateEciScalingConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateEciScalingConfigurationRequest) GetSecurityContextSysctls() []*CreateEciScalingConfigurationRequestSecurityContextSysctls {
	return s.SecurityContextSysctls
}

func (s *CreateEciScalingConfigurationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateEciScalingConfigurationRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateEciScalingConfigurationRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateEciScalingConfigurationRequest) GetTags() []*CreateEciScalingConfigurationRequestTags {
	return s.Tags
}

func (s *CreateEciScalingConfigurationRequest) GetTerminationGracePeriodSeconds() *int64 {
	return s.TerminationGracePeriodSeconds
}

func (s *CreateEciScalingConfigurationRequest) GetVolumes() []*CreateEciScalingConfigurationRequestVolumes {
	return s.Volumes
}

func (s *CreateEciScalingConfigurationRequest) SetAcrRegistryInfos(v []*CreateEciScalingConfigurationRequestAcrRegistryInfos) *CreateEciScalingConfigurationRequest {
	s.AcrRegistryInfos = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetActiveDeadlineSeconds(v int64) *CreateEciScalingConfigurationRequest {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetAutoCreateEip(v bool) *CreateEciScalingConfigurationRequest {
	s.AutoCreateEip = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetAutoMatchImageCache(v bool) *CreateEciScalingConfigurationRequest {
	s.AutoMatchImageCache = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetContainerGroupName(v string) *CreateEciScalingConfigurationRequest {
	s.ContainerGroupName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetContainers(v []*CreateEciScalingConfigurationRequestContainers) *CreateEciScalingConfigurationRequest {
	s.Containers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCostOptimization(v bool) *CreateEciScalingConfigurationRequest {
	s.CostOptimization = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCpu(v float32) *CreateEciScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCpuOptionsCore(v int32) *CreateEciScalingConfigurationRequest {
	s.CpuOptionsCore = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetCpuOptionsThreadsPerCore(v int32) *CreateEciScalingConfigurationRequest {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDataCacheBucket(v string) *CreateEciScalingConfigurationRequest {
	s.DataCacheBucket = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDataCacheBurstingEnabled(v bool) *CreateEciScalingConfigurationRequest {
	s.DataCacheBurstingEnabled = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDataCachePL(v string) *CreateEciScalingConfigurationRequest {
	s.DataCachePL = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDataCacheProvisionedIops(v int32) *CreateEciScalingConfigurationRequest {
	s.DataCacheProvisionedIops = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDescription(v string) *CreateEciScalingConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsConfigNameServers(v []*string) *CreateEciScalingConfigurationRequest {
	s.DnsConfigNameServers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsConfigOptions(v []*CreateEciScalingConfigurationRequestDnsConfigOptions) *CreateEciScalingConfigurationRequest {
	s.DnsConfigOptions = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsConfigSearchs(v []*string) *CreateEciScalingConfigurationRequest {
	s.DnsConfigSearchs = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetDnsPolicy(v string) *CreateEciScalingConfigurationRequest {
	s.DnsPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEgressBandwidth(v int64) *CreateEciScalingConfigurationRequest {
	s.EgressBandwidth = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEipBandwidth(v int32) *CreateEciScalingConfigurationRequest {
	s.EipBandwidth = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEnableSls(v bool) *CreateEciScalingConfigurationRequest {
	s.EnableSls = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetEphemeralStorage(v int32) *CreateEciScalingConfigurationRequest {
	s.EphemeralStorage = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetGpuDriverVersion(v string) *CreateEciScalingConfigurationRequest {
	s.GpuDriverVersion = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetHostAliases(v []*CreateEciScalingConfigurationRequestHostAliases) *CreateEciScalingConfigurationRequest {
	s.HostAliases = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetHostName(v string) *CreateEciScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetImageRegistryCredentials(v []*CreateEciScalingConfigurationRequestImageRegistryCredentials) *CreateEciScalingConfigurationRequest {
	s.ImageRegistryCredentials = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetImageSnapshotId(v string) *CreateEciScalingConfigurationRequest {
	s.ImageSnapshotId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetIngressBandwidth(v int64) *CreateEciScalingConfigurationRequest {
	s.IngressBandwidth = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetInitContainers(v []*CreateEciScalingConfigurationRequestInitContainers) *CreateEciScalingConfigurationRequest {
	s.InitContainers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetInstanceFamilyLevel(v string) *CreateEciScalingConfigurationRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetInstanceTypes(v []*string) *CreateEciScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetIpv6AddressCount(v int32) *CreateEciScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *CreateEciScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetMemory(v float32) *CreateEciScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetNtpServers(v []*string) *CreateEciScalingConfigurationRequest {
	s.NtpServers = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetOwnerId(v int64) *CreateEciScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetRamRoleName(v string) *CreateEciScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetResourceGroupId(v string) *CreateEciScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetResourceOwnerAccount(v string) *CreateEciScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetRestartPolicy(v string) *CreateEciScalingConfigurationRequest {
	s.RestartPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetScalingConfigurationName(v string) *CreateEciScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetScalingGroupId(v string) *CreateEciScalingConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSecurityContextSysctls(v []*CreateEciScalingConfigurationRequestSecurityContextSysctls) *CreateEciScalingConfigurationRequest {
	s.SecurityContextSysctls = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSecurityGroupId(v string) *CreateEciScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSpotPriceLimit(v float32) *CreateEciScalingConfigurationRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetSpotStrategy(v string) *CreateEciScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetTags(v []*CreateEciScalingConfigurationRequestTags) *CreateEciScalingConfigurationRequest {
	s.Tags = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetTerminationGracePeriodSeconds(v int64) *CreateEciScalingConfigurationRequest {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequest) SetVolumes(v []*CreateEciScalingConfigurationRequestVolumes) *CreateEciScalingConfigurationRequest {
	s.Volumes = v
	return s
}

func (s *CreateEciScalingConfigurationRequest) Validate() error {
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
	if s.SecurityContextSysctls != nil {
		for _, item := range s.SecurityContextSysctls {
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

type CreateEciScalingConfigurationRequestAcrRegistryInfos struct {
	// The domain names of the Container Registry Enterprise Edition instances. By default, all domain names are displayed. Separate multiple domain names with commas (,).
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
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEciScalingConfigurationRequestAcrRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) GetDomains() []*string {
	return s.Domains
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetDomains(v []*string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceId(v string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetInstanceName(v string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) SetRegionId(v string) *CreateEciScalingConfigurationRequestAcrRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestAcrRegistryInfos) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainers struct {
	LivenessProbe   *CreateEciScalingConfigurationRequestContainersLivenessProbe   `json:"LivenessProbe,omitempty" xml:"LivenessProbe,omitempty" type:"Struct"`
	ReadinessProbe  *CreateEciScalingConfigurationRequestContainersReadinessProbe  `json:"ReadinessProbe,omitempty" xml:"ReadinessProbe,omitempty" type:"Struct"`
	SecurityContext *CreateEciScalingConfigurationRequestContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The startup arguments of the container. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The commands that you want to run by using the CLI for liveness probing within the container.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs per container.
	//
	// example:
	//
	// 0.25
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables.
	EnvironmentVars []*CreateEciScalingConfigurationRequestContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs per container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image in the container.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Each time instances are created, image pulling is performed.
	//
	// 	- IfNotPresent: Image pulling is performed based on your business requirements. On-premises images are used by default. If no on-premises images are available, images are pulled from remote sources.
	//
	// 	- Never: Image pulling is not performed. On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The commands that you want to run by using the CLI to configure the postStart callback function within the container.
	LifecyclePostStartHandlerExecs []*string `json:"LifecyclePostStartHandlerExecs,omitempty" xml:"LifecyclePostStartHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which you want to send HTTP GET requests to configure the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The path to which you want to send HTTP GET requests to configure the postStart callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port over which you want to send HTTP GET requests to configure the postStart callback function.
	//
	// example:
	//
	// 5050
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests that you want to send to configure the postStart callback function. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTPS
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP sockets that you want to use to configure the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP sockets that you want to use to configure the postStart callback function.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The commands that you want to run by using the CLI to configure the preStop callback function within the container.
	LifecyclePreStopHandlerExecs []*string `json:"LifecyclePreStopHandlerExecs,omitempty" xml:"LifecyclePreStopHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which you want to send HTTP GET requests to configure the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The path to which you want to send HTTP GET requests to configure the preStop callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port over which you want to send HTTP GET requests to configure the preStop callback function.
	//
	// example:
	//
	// 88
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP GET requests that you want to send to configure the preStop callback function. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP sockets that you want to use to configure the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP sockets that you want to use to configure the preStop callback function.
	//
	// example:
	//
	// 90
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size that you want to allocate to the container. Unit: GiB.
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
	Ports []*CreateEciScalingConfigurationRequestContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// Specifies whether the container allocates buffer resources to standard input streams during its active runtime. If you do not specify this parameter, an end-of-file (EOF) error occurs when standard input streams in the container are read.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether standard input streams remain connected during multiple sessions if Stdin is set to true.
	//
	// If you set StdinOnce to true, standard input streams are connected after the container is started, and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected and remain disconnected until the container is restarted.
	//
	// example:
	//
	// false
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable interaction. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// If the command is a /bin/bash command, set the value to true.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The volume mounts of the container.
	VolumeMounts []*CreateEciScalingConfigurationRequestContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory of the container.
	//
	// example:
	//
	// /usr/local/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainers) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainers) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLivenessProbe() *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	return s.LivenessProbe
}

func (s *CreateEciScalingConfigurationRequestContainers) GetReadinessProbe() *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	return s.ReadinessProbe
}

func (s *CreateEciScalingConfigurationRequestContainers) GetSecurityContext() *CreateEciScalingConfigurationRequestContainersSecurityContext {
	return s.SecurityContext
}

func (s *CreateEciScalingConfigurationRequestContainers) GetArgs() []*string {
	return s.Args
}

func (s *CreateEciScalingConfigurationRequestContainers) GetCommands() []*string {
	return s.Commands
}

func (s *CreateEciScalingConfigurationRequestContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateEciScalingConfigurationRequestContainers) GetEnvironmentVars() []*CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	return s.EnvironmentVars
}

func (s *CreateEciScalingConfigurationRequestContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *CreateEciScalingConfigurationRequestContainers) GetImage() *string {
	return s.Image
}

func (s *CreateEciScalingConfigurationRequestContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerExecs() []*string {
	return s.LifecyclePostStartHandlerExecs
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetHost() *string {
	return s.LifecyclePostStartHandlerHttpGetHost
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetPath() *string {
	return s.LifecyclePostStartHandlerHttpGetPath
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetPort() *int32 {
	return s.LifecyclePostStartHandlerHttpGetPort
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerHttpGetScheme() *string {
	return s.LifecyclePostStartHandlerHttpGetScheme
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerTcpSocketHost() *string {
	return s.LifecyclePostStartHandlerTcpSocketHost
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePostStartHandlerTcpSocketPort() *int32 {
	return s.LifecyclePostStartHandlerTcpSocketPort
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerExecs() []*string {
	return s.LifecyclePreStopHandlerExecs
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetHost() *string {
	return s.LifecyclePreStopHandlerHttpGetHost
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetPath() *string {
	return s.LifecyclePreStopHandlerHttpGetPath
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetPort() *int32 {
	return s.LifecyclePreStopHandlerHttpGetPort
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerHttpGetScheme() *string {
	return s.LifecyclePreStopHandlerHttpGetScheme
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerTcpSocketHost() *string {
	return s.LifecyclePreStopHandlerTcpSocketHost
}

func (s *CreateEciScalingConfigurationRequestContainers) GetLifecyclePreStopHandlerTcpSocketPort() *int32 {
	return s.LifecyclePreStopHandlerTcpSocketPort
}

func (s *CreateEciScalingConfigurationRequestContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateEciScalingConfigurationRequestContainers) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestContainers) GetPorts() []*CreateEciScalingConfigurationRequestContainersPorts {
	return s.Ports
}

func (s *CreateEciScalingConfigurationRequestContainers) GetStdin() *bool {
	return s.Stdin
}

func (s *CreateEciScalingConfigurationRequestContainers) GetStdinOnce() *bool {
	return s.StdinOnce
}

func (s *CreateEciScalingConfigurationRequestContainers) GetTty() *bool {
	return s.Tty
}

func (s *CreateEciScalingConfigurationRequestContainers) GetVolumeMounts() []*CreateEciScalingConfigurationRequestContainersVolumeMounts {
	return s.VolumeMounts
}

func (s *CreateEciScalingConfigurationRequestContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLivenessProbe(v *CreateEciScalingConfigurationRequestContainersLivenessProbe) *CreateEciScalingConfigurationRequestContainers {
	s.LivenessProbe = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetReadinessProbe(v *CreateEciScalingConfigurationRequestContainersReadinessProbe) *CreateEciScalingConfigurationRequestContainers {
	s.ReadinessProbe = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetSecurityContext(v *CreateEciScalingConfigurationRequestContainersSecurityContext) *CreateEciScalingConfigurationRequestContainers {
	s.SecurityContext = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetArgs(v []*string) *CreateEciScalingConfigurationRequestContainers {
	s.Args = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetCommands(v []*string) *CreateEciScalingConfigurationRequestContainers {
	s.Commands = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetCpu(v float32) *CreateEciScalingConfigurationRequestContainers {
	s.Cpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetEnvironmentVars(v []*CreateEciScalingConfigurationRequestContainersEnvironmentVars) *CreateEciScalingConfigurationRequestContainers {
	s.EnvironmentVars = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetGpu(v int32) *CreateEciScalingConfigurationRequestContainers {
	s.Gpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetImage(v string) *CreateEciScalingConfigurationRequestContainers {
	s.Image = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetImagePullPolicy(v string) *CreateEciScalingConfigurationRequestContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerExecs(v []*string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerExecs = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetHost(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetPath(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetPort(v int32) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerHttpGetScheme(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerTcpSocketHost(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerExecs(v []*string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerExecs = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetHost(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetPath(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetPort(v int32) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerHttpGetScheme(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerTcpSocketHost(v string) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *CreateEciScalingConfigurationRequestContainers {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetMemory(v float32) *CreateEciScalingConfigurationRequestContainers {
	s.Memory = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetName(v string) *CreateEciScalingConfigurationRequestContainers {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetPorts(v []*CreateEciScalingConfigurationRequestContainersPorts) *CreateEciScalingConfigurationRequestContainers {
	s.Ports = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetStdin(v bool) *CreateEciScalingConfigurationRequestContainers {
	s.Stdin = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetStdinOnce(v bool) *CreateEciScalingConfigurationRequestContainers {
	s.StdinOnce = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetTty(v bool) *CreateEciScalingConfigurationRequestContainers {
	s.Tty = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetVolumeMounts(v []*CreateEciScalingConfigurationRequestContainersVolumeMounts) *CreateEciScalingConfigurationRequestContainers {
	s.VolumeMounts = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) SetWorkingDir(v string) *CreateEciScalingConfigurationRequestContainers {
	s.WorkingDir = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainers) Validate() error {
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

type CreateEciScalingConfigurationRequestContainersLivenessProbe struct {
	Exec                *CreateEciScalingConfigurationRequestContainersLivenessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                                `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                                `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	TimeoutSeconds      *int32                                                                `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbe) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbe) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetExec() *CreateEciScalingConfigurationRequestContainersLivenessProbeExec {
	return s.Exec
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetHttpGet() *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	return s.HttpGet
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetTcpSocket() *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket {
	return s.TcpSocket
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetExec(v *CreateEciScalingConfigurationRequestContainersLivenessProbeExec) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.Exec = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetFailureThreshold(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetHttpGet(v *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetInitialDelaySeconds(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetPeriodSeconds(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetSuccessThreshold(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetTcpSocket(v *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) SetTimeoutSeconds(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbe) Validate() error {
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

type CreateEciScalingConfigurationRequestContainersLivenessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeExec) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeExec) GetCommands() []*string {
	return s.Commands
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeExec) SetCommands(v []*string) *CreateEciScalingConfigurationRequestContainersLivenessProbeExec {
	s.Commands = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeExec) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GetPath() *string {
	return s.Path
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GetPort() *int32 {
	return s.Port
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) GetScheme() *string {
	return s.Scheme
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPath(v string) *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) SetScheme(v string) *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) GetPort() *int32 {
	return s.Port
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersReadinessProbe struct {
	Exec                *CreateEciScalingConfigurationRequestContainersReadinessProbeExec      `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	FailureThreshold    *int32                                                                 `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet             *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet   `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	InitialDelaySeconds *int32                                                                 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	PeriodSeconds       *int32                                                                 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	SuccessThreshold    *int32                                                                 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket           *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	TimeoutSeconds      *int32                                                                 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbe) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbe) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetExec() *CreateEciScalingConfigurationRequestContainersReadinessProbeExec {
	return s.Exec
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetFailureThreshold() *int32 {
	return s.FailureThreshold
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetHttpGet() *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	return s.HttpGet
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetInitialDelaySeconds() *int32 {
	return s.InitialDelaySeconds
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetSuccessThreshold() *int32 {
	return s.SuccessThreshold
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetTcpSocket() *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket {
	return s.TcpSocket
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetExec(v *CreateEciScalingConfigurationRequestContainersReadinessProbeExec) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.Exec = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetFailureThreshold(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.FailureThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetHttpGet(v *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.HttpGet = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetInitialDelaySeconds(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.InitialDelaySeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetPeriodSeconds(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetSuccessThreshold(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.SuccessThreshold = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetTcpSocket(v *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.TcpSocket = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) SetTimeoutSeconds(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbe {
	s.TimeoutSeconds = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbe) Validate() error {
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

type CreateEciScalingConfigurationRequestContainersReadinessProbeExec struct {
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeExec) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeExec) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeExec) GetCommands() []*string {
	return s.Commands
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeExec) SetCommands(v []*string) *CreateEciScalingConfigurationRequestContainersReadinessProbeExec {
	s.Commands = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeExec) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet struct {
	Path   *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Port   *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Scheme *string `json:"Scheme,omitempty" xml:"Scheme,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GetPath() *string {
	return s.Path
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GetPort() *int32 {
	return s.Port
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) GetScheme() *string {
	return s.Scheme
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPath(v string) *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) SetScheme(v string) *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet {
	s.Scheme = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket struct {
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) GetPort() *int32 {
	return s.Port
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersSecurityContext struct {
	Capability             *CreateEciScalingConfigurationRequestContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                    `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                   `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContext) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) GetCapability() *CreateEciScalingConfigurationRequestContainersSecurityContextCapability {
	return s.Capability
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) GetReadOnlyRootFilesystem() *bool {
	return s.ReadOnlyRootFilesystem
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) GetRunAsUser() *int64 {
	return s.RunAsUser
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) SetCapability(v *CreateEciScalingConfigurationRequestContainersSecurityContextCapability) *CreateEciScalingConfigurationRequestContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateEciScalingConfigurationRequestContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) SetRunAsUser(v int64) *CreateEciScalingConfigurationRequestContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContext) Validate() error {
	if s.Capability != nil {
		if err := s.Capability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEciScalingConfigurationRequestContainersSecurityContextCapability struct {
	Add []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContextCapability) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContextCapability) GetAdd() []*string {
	return s.Add
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContextCapability) SetAdd(v []*string) *CreateEciScalingConfigurationRequestContainersSecurityContextCapability {
	s.Add = v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersSecurityContextCapability) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersEnvironmentVars struct {
	// >  This parameter is unavailable for use.
	//
	// example:
	//
	// fieldPath
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable. The name can be 1 to 128 characters in length and can contain letters, underscores (_), and digits. It cannot start with a digit. Specify the value in the [0-9a-zA-Z] format.
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

func (s CreateEciScalingConfigurationRequestContainersEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) SetFieldRefFieldPath(v string) *CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) SetKey(v string) *CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) SetValue(v string) *CreateEciScalingConfigurationRequestContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersPorts struct {
	// The port. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- TCP.
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersPorts) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersPorts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) GetPort() *int32 {
	return s.Port
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) SetPort(v int32) *CreateEciScalingConfigurationRequestContainersPorts {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) SetProtocol(v string) *CreateEciScalingConfigurationRequestContainersPorts {
	s.Protocol = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersPorts) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestContainersVolumeMounts struct {
	// The directory in which the container mounts the volume.
	//
	// >  Data in this directory is overwritten by data on the volume. Specify this parameter with caution.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting. Mount propagation enables volume sharing from one container to other containers within the same pod or to containers across separate pods on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed on the volume or its subdirectories are not propagated to the volume.
	//
	// 	- HostToCotainer: Subsequent mounts executed on the volume or its subdirectories are propagated to the volume.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed on the volume or its subdirectories are propagated to the volume. All volume mounts executed on the container are not only propagate back to the underlying host but also to all containers across every pod that uses the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The volume name. The value of this parameter is the same as the value of Volumes.Name.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Valid values:
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
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The volume subdirectory.
	//
	// example:
	//
	// data2/
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateEciScalingConfigurationRequestContainersVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetMountPath(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetMountPropagation(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetName(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetReadOnly(v bool) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) SetSubPath(v string) *CreateEciScalingConfigurationRequestContainersVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestContainersVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestDnsConfigOptions struct {
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

func (s CreateEciScalingConfigurationRequestDnsConfigOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) GetValue() *string {
	return s.Value
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) SetName(v string) *CreateEciScalingConfigurationRequestDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) SetValue(v string) *CreateEciScalingConfigurationRequestDnsConfigOptions {
	s.Value = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestDnsConfigOptions) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestHostAliases struct {
	// The hostnames of the containers that you want to add.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The IP address of the container that you want to add.
	//
	// example:
	//
	// 1.1.1.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s CreateEciScalingConfigurationRequestHostAliases) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestHostAliases) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestHostAliases) GetHostnames() []*string {
	return s.Hostnames
}

func (s *CreateEciScalingConfigurationRequestHostAliases) GetIp() *string {
	return s.Ip
}

func (s *CreateEciScalingConfigurationRequestHostAliases) SetHostnames(v []*string) *CreateEciScalingConfigurationRequestHostAliases {
	s.Hostnames = v
	return s
}

func (s *CreateEciScalingConfigurationRequestHostAliases) SetIp(v string) *CreateEciScalingConfigurationRequestHostAliases {
	s.Ip = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestHostAliases) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestImageRegistryCredentials struct {
	// The password of the image repository.
	//
	// example:
	//
	// yourpaasword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The endpoint of the image repository.
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

func (s CreateEciScalingConfigurationRequestImageRegistryCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) GetPassword() *string {
	return s.Password
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) GetServer() *string {
	return s.Server
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) GetUserName() *string {
	return s.UserName
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) SetPassword(v string) *CreateEciScalingConfigurationRequestImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) SetServer(v string) *CreateEciScalingConfigurationRequestImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) SetUserName(v string) *CreateEciScalingConfigurationRequestImageRegistryCredentials {
	s.UserName = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestImageRegistryCredentials) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestInitContainers struct {
	SecurityContext *CreateEciScalingConfigurationRequestInitContainersSecurityContext `json:"SecurityContext,omitempty" xml:"SecurityContext,omitempty" type:"Struct"`
	// The startup arguments of the init container. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands of the init container.
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
	// The image pulling policy.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The environment variables of the init container.
	InitContainerEnvironmentVars []*CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	// The ports of init containers.
	InitContainerPorts []*CreateEciScalingConfigurationRequestInitContainersInitContainerPorts `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	// The volume mounts of the init container.
	InitContainerVolumeMounts []*CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
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
	// The working directory of the init container.
	//
	// example:
	//
	// /usr/local
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainers) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainers) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetSecurityContext() *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	return s.SecurityContext
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetArgs() []*string {
	return s.Args
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetCommands() []*string {
	return s.Commands
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetImage() *string {
	return s.Image
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetInitContainerEnvironmentVars() []*CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	return s.InitContainerEnvironmentVars
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetInitContainerPorts() []*CreateEciScalingConfigurationRequestInitContainersInitContainerPorts {
	return s.InitContainerPorts
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetInitContainerVolumeMounts() []*CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	return s.InitContainerVolumeMounts
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestInitContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetSecurityContext(v *CreateEciScalingConfigurationRequestInitContainersSecurityContext) *CreateEciScalingConfigurationRequestInitContainers {
	s.SecurityContext = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetArgs(v []*string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Args = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetCommands(v []*string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Commands = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetCpu(v float32) *CreateEciScalingConfigurationRequestInitContainers {
	s.Cpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetGpu(v int32) *CreateEciScalingConfigurationRequestInitContainers {
	s.Gpu = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetImage(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Image = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetImagePullPolicy(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetInitContainerEnvironmentVars(v []*CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) *CreateEciScalingConfigurationRequestInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetInitContainerPorts(v []*CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) *CreateEciScalingConfigurationRequestInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetInitContainerVolumeMounts(v []*CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) *CreateEciScalingConfigurationRequestInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetMemory(v float32) *CreateEciScalingConfigurationRequestInitContainers {
	s.Memory = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetName(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) SetWorkingDir(v string) *CreateEciScalingConfigurationRequestInitContainers {
	s.WorkingDir = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainers) Validate() error {
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

type CreateEciScalingConfigurationRequestInitContainersSecurityContext struct {
	Capability             *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	ReadOnlyRootFilesystem *bool                                                                        `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	RunAsUser              *int64                                                                       `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContext) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContext) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) GetCapability() *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability {
	return s.Capability
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) GetReadOnlyRootFilesystem() *bool {
	return s.ReadOnlyRootFilesystem
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) GetRunAsUser() *int64 {
	return s.RunAsUser
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) SetCapability(v *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	s.Capability = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) SetReadOnlyRootFilesystem(v bool) *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	s.ReadOnlyRootFilesystem = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) SetRunAsUser(v int64) *CreateEciScalingConfigurationRequestInitContainersSecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContext) Validate() error {
	if s.Capability != nil {
		if err := s.Capability.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability struct {
	Adds []*string `json:"Adds,omitempty" xml:"Adds,omitempty" type:"Repeated"`
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) GetAdds() []*string {
	return s.Adds
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) SetAdds(v []*string) *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability {
	s.Adds = v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars struct {
	// >  This parameter is unavailable for use.
	//
	// example:
	//
	// path
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable. The name can be 1 to 128 characters in length and can contain letters, underscores (_), and digits. It cannot start with a digit. Specify the value in the `[0-9a-zA-Z]` format.
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

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetFieldRefFieldPath(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetKey(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) SetValue(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestInitContainersInitContainerPorts struct {
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

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) GetPort() *int32 {
	return s.Port
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) SetPort(v int32) *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) SetProtocol(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerPorts) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts struct {
	// The directory to which the init container mounts the volume. The data stored in this directory is overwritten by the data on the volume. Specify this parameter with caution.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation enables volume sharing from one container to other containers within the same pod or to containers across separate pods on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed on the volume or its subdirectories do not propagate to the volume.
	//
	// 	- HostToCotainer: Subsequent mounts executed on the volume or its subdirectories propagate to the volume.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed on the volume or its subdirectories propagate to the volume. All volume mounts executed on the container not only propagate back to the underlying host but also to all containers across every pod that uses the same volume.
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
	// The subdirectory of the volume. The pod can mount different directories of the same volume to different subdirectories of init containers.
	//
	// example:
	//
	// /usr/sub/
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPath(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetName(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) SetSubPath(v string) *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestSecurityContextSysctls struct {
	// The variable name of the security context in which the container group runs.
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

func (s CreateEciScalingConfigurationRequestSecurityContextSysctls) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestSecurityContextSysctls) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) GetValue() *string {
	return s.Value
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) SetName(v string) *CreateEciScalingConfigurationRequestSecurityContextSysctls {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) SetValue(v string) *CreateEciScalingConfigurationRequestSecurityContextSysctls {
	s.Value = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestSecurityContextSysctls) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestTags struct {
	// The tag key of the elastic container instance.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with `acs:` or `aliyun` or contain `http://` or `https://`.
	//
	// example:
	//
	// version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the elastic container instance.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// 3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEciScalingConfigurationRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestTags) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateEciScalingConfigurationRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateEciScalingConfigurationRequestTags) SetKey(v string) *CreateEciScalingConfigurationRequestTags {
	s.Key = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestTags) SetValue(v string) *CreateEciScalingConfigurationRequestTags {
	s.Value = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestVolumes struct {
	DiskVolume     *CreateEciScalingConfigurationRequestVolumesDiskVolume     `json:"DiskVolume,omitempty" xml:"DiskVolume,omitempty" type:"Struct"`
	EmptyDirVolume *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume `json:"EmptyDirVolume,omitempty" xml:"EmptyDirVolume,omitempty" type:"Struct"`
	FlexVolume     *CreateEciScalingConfigurationRequestVolumesFlexVolume     `json:"FlexVolume,omitempty" xml:"FlexVolume,omitempty" type:"Struct"`
	HostPathVolume *CreateEciScalingConfigurationRequestVolumesHostPathVolume `json:"HostPathVolume,omitempty" xml:"HostPathVolume,omitempty" type:"Struct"`
	NFSVolume      *CreateEciScalingConfigurationRequestVolumesNFSVolume      `json:"NFSVolume,omitempty" xml:"NFSVolume,omitempty" type:"Struct"`
	// The paths to the configuration files.
	ConfigFileVolumeConfigFileToPaths []*CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	// The default permissions on the ConfigFile volume.
	//
	// example:
	//
	// 0644
	ConfigFileVolumeDefaultMode *int32 `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	// The name of the volume.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the Host directory. Examples: File, Directory, and Socket.
	//
	// example:
	//
	// ConfigFileVolume
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumes) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumes) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetDiskVolume() *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	return s.DiskVolume
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetEmptyDirVolume() *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume {
	return s.EmptyDirVolume
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetFlexVolume() *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	return s.FlexVolume
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetHostPathVolume() *CreateEciScalingConfigurationRequestVolumesHostPathVolume {
	return s.HostPathVolume
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetNFSVolume() *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	return s.NFSVolume
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetConfigFileVolumeConfigFileToPaths() []*CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	return s.ConfigFileVolumeConfigFileToPaths
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetConfigFileVolumeDefaultMode() *int32 {
	return s.ConfigFileVolumeDefaultMode
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetName() *string {
	return s.Name
}

func (s *CreateEciScalingConfigurationRequestVolumes) GetType() *string {
	return s.Type
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetDiskVolume(v *CreateEciScalingConfigurationRequestVolumesDiskVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.DiskVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetEmptyDirVolume(v *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.EmptyDirVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetFlexVolume(v *CreateEciScalingConfigurationRequestVolumesFlexVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.FlexVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetHostPathVolume(v *CreateEciScalingConfigurationRequestVolumesHostPathVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.HostPathVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetNFSVolume(v *CreateEciScalingConfigurationRequestVolumesNFSVolume) *CreateEciScalingConfigurationRequestVolumes {
	s.NFSVolume = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetConfigFileVolumeConfigFileToPaths(v []*CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) *CreateEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetConfigFileVolumeDefaultMode(v int32) *CreateEciScalingConfigurationRequestVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetName(v string) *CreateEciScalingConfigurationRequestVolumes {
	s.Name = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) SetType(v string) *CreateEciScalingConfigurationRequestVolumes {
	s.Type = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumes) Validate() error {
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
	if s.ConfigFileVolumeConfigFileToPaths != nil {
		for _, item := range s.ConfigFileVolumeConfigFileToPaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEciScalingConfigurationRequestVolumesDiskVolume struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	FsType   *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesDiskVolume) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesDiskVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) GetFsType() *string {
	return s.FsType
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) SetDiskId(v string) *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskId = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) SetDiskSize(v int32) *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	s.DiskSize = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) SetFsType(v string) *CreateEciScalingConfigurationRequestVolumesDiskVolume {
	s.FsType = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesDiskVolume) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestVolumesEmptyDirVolume struct {
	Medium    *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	SizeLimit *string `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) GetMedium() *string {
	return s.Medium
}

func (s *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) GetSizeLimit() *string {
	return s.SizeLimit
}

func (s *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) SetMedium(v string) *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume {
	s.Medium = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) SetSizeLimit(v string) *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume {
	s.SizeLimit = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesEmptyDirVolume) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestVolumesFlexVolume struct {
	Driver  *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	FsType  *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesFlexVolume) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesFlexVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) GetDriver() *string {
	return s.Driver
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) GetFsType() *string {
	return s.FsType
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) GetOptions() *string {
	return s.Options
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) SetDriver(v string) *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	s.Driver = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) SetFsType(v string) *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	s.FsType = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) SetOptions(v string) *CreateEciScalingConfigurationRequestVolumesFlexVolume {
	s.Options = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesFlexVolume) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestVolumesHostPathVolume struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesHostPathVolume) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesHostPathVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) GetPath() *string {
	return s.Path
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) GetType() *string {
	return s.Type
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) SetPath(v string) *CreateEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) SetType(v string) *CreateEciScalingConfigurationRequestVolumesHostPathVolume {
	s.Type = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesHostPathVolume) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestVolumesNFSVolume struct {
	Path     *string `json:"Path,omitempty" xml:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Server   *string `json:"Server,omitempty" xml:"Server,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesNFSVolume) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesNFSVolume) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) GetPath() *string {
	return s.Path
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) GetServer() *string {
	return s.Server
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) SetPath(v string) *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) SetReadOnly(v bool) *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	s.ReadOnly = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) SetServer(v string) *CreateEciScalingConfigurationRequestVolumesNFSVolume {
	s.Server = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesNFSVolume) Validate() error {
	return dara.Validate(s)
}

type CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths struct {
	// The content of the configuration file (32 KB).
	//
	// example:
	//
	// bGl1bWk=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The permissions on the configuration file.
	//
	// example:
	//
	// 0644
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the environment variable.
	//
	// example:
	//
	// PATH
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) GetContent() *string {
	return s.Content
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) GetMode() *int32 {
	return s.Mode
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) GetPath() *string {
	return s.Path
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) SetMode(v int32) *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	s.Mode = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

func (s *CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths) Validate() error {
	return dara.Validate(s)
}
