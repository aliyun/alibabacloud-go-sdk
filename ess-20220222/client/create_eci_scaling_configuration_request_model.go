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
	// The information about the Container Registry Enterprise instances.
	AcrRegistryInfos []*CreateEciScalingConfigurationRequestAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	// The validity period. Unit: seconds.
	//
	// example:
	//
	// 1000
	ActiveDeadlineSeconds *int64 `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	// Specifies whether automatic creation of an elastic IP address (EIP) is performed and the EIP is attached to the ECI instance.
	//
	// example:
	//
	// true
	AutoCreateEip *bool `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	// Specifies whether to automatically match an image cache. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The name of the ECI instance, which is the name of the container group.
	//
	// To set a sequential instance name, use the following format: name_prefix(AUTO_INCREMENT)[begin_number,bits]name_suffix.
	//
	// - **name_prefix**: the prefix of the hostname.
	//
	// - **(AUTO_INCREMENT)**: a fixed value that identifies this sorting method.
	//
	// - **begin_number**: the start value of the sequential number. Valid values: 0 to 999999.
	//
	// - **bits**: the number of digits of the sequential number. Valid values: 1 to 6. If the number of digits of the specified begin_number is greater than the value of bits, bits defaults to 6.
	//
	// <notice>We recommend that you set bits to at least 3. Otherwise, the upper limit of the sequential number is easily reached. After the upper limit is reached, if scale-out is still required, the scaling activity reports an error and stops. In this case, you must reset the rule for generating hostnames.</notice>
	//
	// - **name_suffix*	- (optional): the suffix of the hostname.
	//
	// <notice>No spaces are allowed in [begin_number,bits]. The system increments the number sequentially by default. However, if an ECI instance created during scale-out fails to start, the instance is removed from the scaling group and a new ECI instance is created. Therefore, the hostnames of ECI instances in the scaling group may be incremented discontinuously.</notice>
	//
	// **Example:*	- If you set the value to nginx-test-(AUTO_INCREMENT)[0,3], the container group names are nginx-test-000, nginx-test-001, and nginx-test-002.
	//
	// The rules for setting sequential ECI instance names are the same as those for setting sequential ECS instance names in a scaling group. For more information, see [Set a rule for sequential and unique instance names or hostnames](https://help.aliyun.com/document_detail/203831.html).
	//
	// example:
	//
	// nginx-test-(AUTO_INCREMENT)[0,3]
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The list of containers included in the instance.
	Containers []*CreateEciScalingConfigurationRequestContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// Specifies whether to enable cost optimization. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	CostOptimization *bool `json:"CostOptimization,omitempty" xml:"CostOptimization,omitempty"`
	// The number of vCPUs for the instance. Unit: cores.
	//
	// example:
	//
	// 1.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of physical CPU cores. Only specific instance types support custom settings. For more information, see [Customize CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// The number of threads per core. Only specific instance types support custom settings. A value of 1 indicates that hyper-threading is disabled. For more information, see [Customize CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The data cache bucket.
	//
	// example:
	//
	// default
	DataCacheBucket *string `json:"DataCacheBucket,omitempty" xml:"DataCacheBucket,omitempty"`
	// Specifies whether to enable the burst feature for the ESSD AutoPL cloud disk used for data caching. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// > For more information about ESSD AutoPL cloud disks, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	DataCacheBurstingEnabled *bool `json:"DataCacheBurstingEnabled,omitempty" xml:"DataCacheBurstingEnabled,omitempty"`
	// The performance level of the cloud disk used for data caching. ESSDs are recommended. Valid values for the performance level of an ESSD:
	//
	// - PL0: An ESSD at PL0 delivers up to 10,000 random read/write IOPS.
	//
	// - PL1: An ESSD at PL1 delivers up to 50,000 random read/write IOPS.
	//
	// - PL2: An ESSD at PL2 delivers up to 100,000 random read/write IOPS.
	//
	// - PL3: An ESSD at PL3 delivers up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// > For more information about ESSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	DataCachePL *string `json:"DataCachePL,omitempty" xml:"DataCachePL,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL cloud disk used for data caching. Valid values: 0 to min{50000, 1000 × capacity - baseline performance}, where baseline performance = min{1800 + 50 × capacity, 50000}.
	//
	// > For more information about ESSD AutoPL cloud disks, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	DataCacheProvisionedIops *int32 `json:"DataCacheProvisionedIops,omitempty" xml:"DataCacheProvisionedIops,omitempty"`
	// >This parameter is unavailable.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of IP addresses of DNS servers.
	DnsConfigNameServers []*string `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	// The list of object options. Each object consists of a Name and an optional Value.
	DnsConfigOptions []*CreateEciScalingConfigurationRequestDnsConfigOptions `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	// The list of DNS search domains.
	DnsConfigSearchs []*string `json:"DnsConfigSearchs,omitempty" xml:"DnsConfigSearchs,omitempty" type:"Repeated"`
	// The DNS policy. Valid values:
	//
	// - None: Uses the DNS settings specified in the DnsConfig field.
	//
	// - Default: Inherits the DNS settings of the runtime environment.
	//
	// example:
	//
	// Default
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// The outbound bandwidth limit. Unit: bytes.
	//
	// example:
	//
	// 1024000
	EgressBandwidth *int64 `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	// The bandwidth of the EIP. Default value: 5 Mbit/s.
	//
	// example:
	//
	// 5
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// >This parameter is unavailable.
	//
	// example:
	//
	// false
	EnableSls *bool `json:"EnableSls,omitempty" xml:"EnableSls,omitempty"`
	// The size of the ephemeral storage space. By default, a PL1 ESSD cloud disk is used. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The GPU driver version. Valid values:
	//
	// - tesla=470.82.01 (default)
	//
	// - tesla=525.85.12
	//
	// > Only specific instance types support switching the GPU driver version. For more information, see [Create a GPU-accelerated instance](https://help.aliyun.com/document_detail/2579486.html).
	//
	// example:
	//
	// tesla=525.85.12
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The custom hostname mappings for a container in the instance.
	HostAliases []*CreateEciScalingConfigurationRequestHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The hostname, which is the name of the container group.
	//
	// example:
	//
	// test
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The details of the image registry.
	ImageRegistryCredentials []*CreateEciScalingConfigurationRequestImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	// The ID of the image cache.
	//
	// example:
	//
	// imc-2zebxkiifuyzzlhl****
	ImageSnapshotId *string `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	// The inbound bandwidth limit. Unit: bytes.
	//
	// example:
	//
	// 1024000
	IngressBandwidth *int64 `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	// The list of init containers.
	InitContainers []*CreateEciScalingConfigurationRequestInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The instance family level, which is used to filter instance types that meet your requirements. This parameter takes effect only when CostOptimization is enabled. Valid values:
	//
	// - EntryLevel: entry-level, which indicates shared instance types. Shared instance types are more cost-effective but cannot guarantee stable computing performance. They are suitable for business scenarios with low average CPU utilization. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise-level. Enterprise-level instance types provide stable performance and dedicated resources, and are suitable for business scenarios that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit entry-level, which indicates burstable instances. Burstable instances use CPU credits to ensure computing performance and are suitable for business scenarios with low average CPU utilization but occasional CPU bursts. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The specified ECS instance types. You can specify up to five instance types.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The number of IPv6 addresses.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The weight of the ECI instance when it is used as a backend server for load balancing. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size for the instance. Unit: GiB.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The NTP server.
	NtpServers []*string `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	OwnerId    *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the instance RAM role. ECI and ECS share instance RAM roles. For more information, see [Instance RAM roles](https://help.aliyun.com/document_detail/61175.html).
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
	// The restart policy of the ECI instance. Valid values:
	//
	// - Always: Always restarts the instance.
	//
	// - Never: Never restarts the instance.
	//
	// - OnFailure: Restarts the instance upon failure.
	//
	// Default value: Always.
	//
	// example:
	//
	// Always
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The name of the scaling configuration. The name must be 2 to 64 characters in length and start with a digit, letter, or Chinese character. It can contain digits, underscores (_), hyphens (-), and periods (.).
	//
	// The name must be unique within a scaling group in a region. If you do not specify this parameter, the ID of the scaling configuration is used by default.
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
	// The security context system information for the instance.
	SecurityContextSysctls []*CreateEciScalingConfigurationRequestSecurityContextSysctls `json:"SecurityContextSysctls,omitempty" xml:"SecurityContextSysctls,omitempty" type:"Repeated"`
	// The ID of the security group to which the ECI instance belongs. ECI instances in the same security group can access each other.
	//
	// If you do not specify a security group, the system automatically uses the default security group in the selected region. You must ensure that the inbound rules of this security group include the protocols and ports that the containers need to expose. If no default security group exists in the region, the system performs automatic creation of a default security group and adds the declared container protocols and ports to the inbound rules of the security group.
	//
	// example:
	//
	// sg-uf66jeqopgqa9hdn****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The maximum hourly price for the spot instance. The value is accurate to three decimal places.
	//
	// This parameter is required when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy for the instance. Valid values:
	//
	// - NoSpot: Regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: Spot instance with a maximum price limit.
	//
	// - SpotAsPriceGo: Spot instance whose price automatically follows the current market price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags of the ECI instance. Tags are passed as key-value pairs. You can use up to 20 tags. The requirements for keys and values are as follows:
	//
	//
	//
	// - A key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. A key cannot be an empty string.
	//
	// - A value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. A value can be an empty string.
	Tags []*CreateEciScalingConfigurationRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The buffer time for the program to handle operations before shutdown. Unit: seconds.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int64 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The volume information.
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
	// The domain names of the Container Registry Enterprise instance. By default, all domain names of the instance are used. You can specify one or more domain names. Separate multiple domain names with commas (,).
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The instance ID of the Container Registry Enterprise instance.
	//
	// example:
	//
	// cri-nwj395hgf6f3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry Enterprise instance.
	//
	// example:
	//
	// acr-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region where the Container Registry Enterprise instance resides.
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
	// The arguments for the container startup command. A maximum of 10 arguments are supported.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands for the container. A maximum of 20 commands are supported. Each command can contain up to 256 characters.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of CPU cores allocated to the container.
	//
	// example:
	//
	// 0.25
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variable information.
	EnvironmentVars []*CreateEciScalingConfigurationRequestContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs allocated to the container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The container image. You can pull images from an image repository such as Docker Hub and Alibaba Cloud Container Registry (ACR). When you pull an image over the Internet, if no NAT gateway is created in the VPC, you must associate the ECI instance with an EIP.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/eci_open/nginx:latest
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pull policy for the container. Valid values:
	//
	// - Always: Always pulls the image.
	//
	// - IfNotPresent: Pulls the image only if it is not present locally.
	//
	// - Never: Never pulls the image. Only local images are used.
	//
	// > If image caching is enabled, ImagePullPolicy is overridden to Always when the container is created.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The command executed in the container when the command line method is used to configure the postStart callback function.
	LifecyclePostStartHandlerExecs []*string `json:"LifecyclePostStartHandlerExecs,omitempty" xml:"LifecyclePostStartHandlerExecs,omitempty" type:"Repeated"`
	// The host address that receives the HTTP GET request when the HTTP request method is used to configure the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The path accessed by the HTTP GET request when the HTTP request method is used to configure the postStart callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port number checked by the HTTP GET request when the HTTP request method is used to configure the postStart callback function.
	//
	// example:
	//
	// 5050
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type used for the HTTP GET request when the HTTP request method is used to configure the postStart callback function. Valid values:
	//
	// - HTTP
	//
	// - HTTPS
	//
	// example:
	//
	// HTTPS
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The host address checked by the TCP socket when the TCP socket method is used to configure the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port checked by the TCP socket when the TCP socket method is used to configure the postStart callback function.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The command executed in the container when the command line method is used to configure the preStop callback function.
	LifecyclePreStopHandlerExecs []*string `json:"LifecyclePreStopHandlerExecs,omitempty" xml:"LifecyclePreStopHandlerExecs,omitempty" type:"Repeated"`
	// The host address that receives the HTTP GET request when the HTTP request method is used to configure the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The path accessed by the HTTP GET request when the HTTP request method is used to configure the preStop callback function.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port number checked by the HTTP GET request when the HTTP request method is used to configure the preStop callback function.
	//
	// example:
	//
	// 88
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type used for the HTTP GET request when the HTTP request method is used to configure the preStop callback function. Valid values:
	//
	// - HTTP
	//
	// - HTTPS
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The host address checked by the TCP socket when the TCP socket method is used to configure the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port checked by the TCP socket when the TCP socket method is used to configure the preStop callback function.
	//
	// example:
	//
	// 90
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The memory size of the container. Unit: GiB.
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
	// The list of ports.
	Ports []*CreateEciScalingConfigurationRequestContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// Specifies whether to allocate a buffer for standard input (stdin) during container runtime. If this parameter is not set, reads from stdin in the container result in an EOF.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Specifies whether the standard input stream maintains the enabling status across multiple attached sessions when stdin is set to true.
	//
	// If StdinOnce is set to true, standard input is opened when the container starts and remains empty until the first client attaches to it. It then stays open to accept data until the client disconnects, at which point standard input is shutdown and remains shutdown until the container is restarted.
	//
	// example:
	//
	// false
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Specifies whether to enable interactive mode. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// For example, if Command is set to /bin/bash, set this parameter to true.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The volume mount information of the container.
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
	Exec *CreateEciScalingConfigurationRequestContainersLivenessProbeExec `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	// The minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	FailureThreshold *int32                                                              `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet          *CreateEciScalingConfigurationRequestContainersLivenessProbeHttpGet `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	// The number of seconds after the container has started before the probe is initiated.
	//
	// example:
	//
	// 5
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The interval at which the probe is performed in seconds. Default value: 10. Minimum value: 1.
	//
	// example:
	//
	// 5
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The minimum consecutive successes for the probe to be considered successful after having failed. Default value: 1. The current value must be 1.
	//
	// example:
	//
	// 1
	SuccessThreshold *int32                                                                `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket        *CreateEciScalingConfigurationRequestContainersLivenessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	// The timeout period for the probe in seconds. Default value: 1. Minimum value: 1.
	//
	// example:
	//
	// 1
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
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
	// The command executed in the container when the exec probe is used for health checks.
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
	// The path accessed by the HTTP GET request when the HTTP probe is used for health checks.
	//
	// example:
	//
	// /healthyz
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port number on which the HTTP GET probe performs health checks.
	//
	// example:
	//
	// 8888
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type used for the HTTP GET request when the HTTP probe is used for health checks. Valid values:
	//
	// - HTTP
	//
	// - HTTPS
	//
	// example:
	//
	// HTTP
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
	// The port number on which the TCP socket probe performs health checks.
	//
	// example:
	//
	// 8000
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
	Exec *CreateEciScalingConfigurationRequestContainersReadinessProbeExec `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
	// The minimum consecutive failures for the probe to be considered failed after having succeeded. Default value: 3.
	//
	// example:
	//
	// 3
	FailureThreshold *int32                                                               `json:"FailureThreshold,omitempty" xml:"FailureThreshold,omitempty"`
	HttpGet          *CreateEciScalingConfigurationRequestContainersReadinessProbeHttpGet `json:"HttpGet,omitempty" xml:"HttpGet,omitempty" type:"Struct"`
	// The number of seconds after the container has started before the probe is initiated.
	//
	// example:
	//
	// 3
	InitialDelaySeconds *int32 `json:"InitialDelaySeconds,omitempty" xml:"InitialDelaySeconds,omitempty"`
	// The interval at which the probe is performed in seconds. Default value: 10. Minimum value: 1.
	//
	// example:
	//
	// 3
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The minimum consecutive successes for the probe to be considered successful after having failed. Default value: 1. The current value must be 1.
	//
	// example:
	//
	// 1
	SuccessThreshold *int32                                                                 `json:"SuccessThreshold,omitempty" xml:"SuccessThreshold,omitempty"`
	TcpSocket        *CreateEciScalingConfigurationRequestContainersReadinessProbeTcpSocket `json:"TcpSocket,omitempty" xml:"TcpSocket,omitempty" type:"Struct"`
	// The timeout period for the probe in seconds. Default value: 1. Minimum value: 1.
	//
	// example:
	//
	// 5
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
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
	// The command executed in the container when the exec probe is used for health checks.
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
	// The path accessed by the HTTP GET request when the HTTP probe is used for health checks.
	//
	// example:
	//
	// /healthz
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The port number on which the HTTP GET probe performs health checks.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type used for the HTTP GET request when the HTTP probe is used for health checks. Valid values:
	//
	// - HTTP
	//
	// - HTTPS
	//
	// example:
	//
	// HTTP
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
	// The port number on which the TCP socket probe performs health checks.
	//
	// example:
	//
	// 8000
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
	Capability *CreateEciScalingConfigurationRequestContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	// Specifies whether the root file system of the container is read-only. Currently, only true is supported.
	//
	// example:
	//
	// true
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	// The user ID used to run the container.
	//
	// example:
	//
	// 1000
	RunAsUser *int64 `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
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
	// Grants specific privileges to processes in the container. Currently, only NET_ADMIN and NET_RAW are supported.
	//
	// > NET_RAW is not supported by default. To use NET_RAW, submit a ticket.
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
	// >This parameter is unavailable.
	//
	// example:
	//
	// fieldPath
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable. The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (_). It cannot start with a digit.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value must be 0 to 256 characters in length.
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
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// - TCP
	//
	// - UDP
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
	// The directory where the volume is mounted in the container.
	//
	// 	Notice: The contents of the container mount directory are directly overwritten by the contents of the volume. Proceed with caution.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings for the volume. Mount propagation allows volumes mounted by a container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// - None: The volume is not aware of any subsequent mount operations performed on the volume or its subdirectories.
	//
	// - HostToCotainer: The volume is aware of subsequent mount operations performed on the volume or its subdirectories.
	//
	// - Bidirectional: Similar to HostToCotainer, the volume is aware of mount operations. In addition, the volume is propagated back to the host and to all containers of all pods that use the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the volume. The value must be the same as the Name value in Volumes.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the volume is read-only. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subpath of the volume.
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
	// The name of the object variable.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the object variable.
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
	// The hostname to be added.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The IP address to be added.
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
	// The password of the image registry.
	//
	// example:
	//
	// yourpassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The address of the image registry.
	//
	// example:
	//
	// registry-vpc.cn-shanghai.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The username of the image registry.
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
	// The arguments for the container startup command. A maximum of 10 arguments are supported.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The startup commands of the container.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs for the container. Unit: cores.
	//
	// example:
	//
	// 0.5
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs allocated to the container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image used by the container.
	//
	// example:
	//
	// nginx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pull policy.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The environment variable information of the init container.
	InitContainerEnvironmentVars []*CreateEciScalingConfigurationRequestInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	// The list of port mappings for the init container.
	InitContainerPorts []*CreateEciScalingConfigurationRequestInitContainersInitContainerPorts `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	// The array of mount point information for the init container.
	InitContainerVolumeMounts []*CreateEciScalingConfigurationRequestInitContainersInitContainerVolumeMounts `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	// The memory size. Unit: GiB.
	//
	// example:
	//
	// 1.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The name of the container.
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
	Capability *CreateEciScalingConfigurationRequestInitContainersSecurityContextCapability `json:"Capability,omitempty" xml:"Capability,omitempty" type:"Struct"`
	// Specifies whether the root file system where the container runs is read-only. Only true is supported.
	//
	// example:
	//
	// true
	ReadOnlyRootFilesystem *bool `json:"ReadOnlyRootFilesystem,omitempty" xml:"ReadOnlyRootFilesystem,omitempty"`
	// The user ID used to run the container.
	//
	// example:
	//
	// 587
	RunAsUser *int64 `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
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
	// Grants specific permissions to processes in the container. Only NET_ADMIN and NET_RAW are supported.
	//
	// > NET_RAW is not supported by default. Submit a ticket to apply for this permission.
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
	// >This parameter is unavailable.
	//
	// example:
	//
	// path
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable. The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (_). It cannot start with a digit.
	//
	// example:
	//
	// Path
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable. The value must be 0 to 256 characters in length.
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
	// - TCP
	//
	// - UDP
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
	// The mount directory. The content in the mount directory of the container is directly overwritten by the content of the volume. Use this parameter with caution.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings of the volume. Mount propagation allows a volume mounted by a container to be shared with other containers in the same pod, or even with other pods on the same node. Valid values:
	//
	// - None: The volume is not aware of any subsequent mount operations performed on this volume or its subdirectories.
	//
	// - HostToCotainer: The volume is aware of subsequent mount operations performed on this volume or its subdirectories.
	//
	// - Bidirectional: Similar to HostToCotainer, the volume is aware of mount operations. In addition, the volume is propagated back to the host and to all containers of all pods that use the same volume.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The name of the mounted volume.
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
	// The subdirectory under the volume. This allows a pod to mount different directories under the same volume to different directories in the container.
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
	// The system name of the security context for the instance.
	//
	// example:
	//
	// kernel.msgmax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The contextual variable value of the security context for the instance.
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
	// The tag key of the instance.
	//
	// If you specify this parameter, it cannot be an empty string. It can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// version
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance.
	//
	// If you specify this parameter, it can be an empty string. It can be up to 128 characters in length, and cannot start with `acs:`, or contain `http://` or `https://`.
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
	// The list of configuration file paths.
	ConfigFileVolumeConfigFileToPaths []*CreateEciScalingConfigurationRequestVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	// The default permissions of the ConfigFileVolume.
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
	// The volume type. Valid values:
	//
	// - EmptyDirVolume: the EmptyDir volume, which represents an empty directory.
	//
	// - NFSVolume: the NFS volume, which represents a network file system.
	//
	// - ConfigFileVolume: the ConfigFile volume, which represents a configuration file.
	//
	// - FlexVolume: uses the FlexVolume plug-in to extend storage types and supports mounting cloud disks.
	//
	// - HostPathVolume: the HostPath volume, which represents a file or directory on the host node.
	//
	// - DiskVolume (not recommended): the cloud disk volume. Use FlexVolume to mount cloud disks instead.
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
	// The ID of the DiskVolume.
	//
	// example:
	//
	// d-xx
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The size of the DiskVolume. Unit: GiB.
	//
	// example:
	//
	// 15
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// This file system type is not recommended. Use the FlexVolume.FsType parameter instead.
	//
	// example:
	//
	// xfs
	FsType *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
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
	// The storage medium of the EmptyDirVolume. The default value is empty, which indicates that the node file system is used. The value memory is supported, which indicates that memory is used.
	//
	// example:
	//
	// memory
	Medium *string `json:"Medium,omitempty" xml:"Medium,omitempty"`
	// The size of the EmptyDir volume. Specify a unit for the value. Use Gi or Mi.
	//
	// example:
	//
	// 256Mi
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
	// The name of the driver used for the FlexVolume.
	//
	// example:
	//
	// flexvolume
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// The type of the file system to be mounted. The default value depends on the script of the FlexVolume.
	//
	// example:
	//
	// ext4
	FsType *string `json:"FsType,omitempty" xml:"FsType,omitempty"`
	// The list of FlexVolume object options. The options are in key-value pairs and passed in JSON format.
	//
	// For example, when you mount a cloud disk by using FlexVolume, the format of Options is `{"volumeId":"d-2zehdahrwoa7srg****","performanceLevel": "PL2"}`.
	//
	// example:
	//
	// {"volumeId":"d-2zehdahrwoa7srg****","performanceLevel": "PL2"}
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
	// The absolute path on the host.
	//
	// example:
	//
	// /xx/xx/path
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The type of the host directory. Examples: File, Directory, and Socket.
	//
	// example:
	//
	// Directory
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
	// The path of the NFS volume.
	//
	// example:
	//
	// /share
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether the NFS volume is read-only. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The address of the NFS server.
	//
	// example:
	//
	// 3f9cd4a596-naw76.cn-shanghai.nas.aliyuncs.com
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
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
	// The directory permissions of the ConfigFileVolume.
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
