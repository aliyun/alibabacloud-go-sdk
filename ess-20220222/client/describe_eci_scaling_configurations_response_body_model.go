// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEciScalingConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeEciScalingConfigurationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEciScalingConfigurationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEciScalingConfigurationsResponseBody
	GetRequestId() *string
	SetScalingConfigurations(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) *DescribeEciScalingConfigurationsResponseBody
	GetScalingConfigurations() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations
	SetTotalCount(v int32) *DescribeEciScalingConfigurationsResponseBody
	GetTotalCount() *int32
}

type DescribeEciScalingConfigurationsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling configurations.
	ScalingConfigurations []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations `json:"ScalingConfigurations,omitempty" xml:"ScalingConfigurations,omitempty" type:"Repeated"`
	// The total number of scaling configurations.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEciScalingConfigurationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEciScalingConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEciScalingConfigurationsResponseBody) GetScalingConfigurations() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	return s.ScalingConfigurations
}

func (s *DescribeEciScalingConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetPageNumber(v int32) *DescribeEciScalingConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetPageSize(v int32) *DescribeEciScalingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetRequestId(v string) *DescribeEciScalingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetScalingConfigurations(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) *DescribeEciScalingConfigurationsResponseBody {
	s.ScalingConfigurations = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) SetTotalCount(v int32) *DescribeEciScalingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurations struct {
	// The Container Registry Enterprise Edition instances.
	AcrRegistryInfos []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	// The validity period of the scaling configuration. Unit: seconds.
	//
	// example:
	//
	// 1000
	ActiveDeadlineSeconds *int32 `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	// Indicates whether elastic IP addresses (EIPs) are automatically created and bound to elastic container instances.
	//
	// example:
	//
	// true
	AutoCreateEip *bool `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	// Indicates whether the image cache is automatically matched. Default value: false.
	//
	// example:
	//
	// false
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The name series of elastic container instances.
	//
	// example:
	//
	// test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The containers in an elastic container instance.
	Containers []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// Indicates whether the Cost Optimization feature is enabled. Valid values:
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
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of physical CPU cores. This parameter can be specified for only specific instance types. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// The number of threads per core. This parameter can be specified for only specific instance types. A value of 1 indicates that Hyper-Threading is disabled. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The time at which the scaling configuration was created.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The bucket that stores the data cache.
	//
	// example:
	//
	// default
	DataCacheBucket *string `json:"DataCacheBucket,omitempty" xml:"DataCacheBucket,omitempty"`
	// Indicates whether the Performance Burst feature is enabled for the ESSD AutoPL disk that stores the data cache. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  For more information about ESSD AutoPL disks, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	DataCacheBurstingEnabled *bool `json:"DataCacheBurstingEnabled,omitempty" xml:"DataCacheBurstingEnabled,omitempty"`
	// The performance level (PL) of the cloud disk that stores the data cache. We recommend that you use enterprise SSDs (ESSDs). Valid values:
	//
	// 	- PL0: An ESSD can provide up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can provide up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can provide up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD can provide up to 1,000,000 random read/write IOPS.
	//
	// >  For more information about ESSDs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	DataCachePL *string `json:"DataCachePL,omitempty" xml:"DataCachePL,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk that stores the data cache. Valid values: 0 to min{50,000, 1000 x *Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50	- x Capacity, 50,000}.
	//
	// >  For more information about ESSD AutoPL disks, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	DataCacheProvisionedIops *int32 `json:"DataCacheProvisionedIops,omitempty" xml:"DataCacheProvisionedIops,omitempty"`
	// >  This parameter is not available for use.
	//
	// example:
	//
	// This is an example.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP addresses of the Domain Name Service (DNS) server.
	DnsConfigNameServers []*string `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	// The options. Each option is a name-value pair. The value in the name-value pair is optional.
	DnsConfigOptions []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	// The search domains of the DNS servers.
	DnsConfigSearches []*string `json:"DnsConfigSearches,omitempty" xml:"DnsConfigSearches,omitempty" type:"Repeated"`
	// The Domain Name System (DNS) policy.
	//
	// example:
	//
	// Default
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// The maximum outbound public bandwidth. Unit: bit/s.
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
	// The capacity of the ephemeral storage. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The version of the GPU driver.
	//
	// example:
	//
	// tesla=470.82.01
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The hostnames and IP addresses for a container that are added to the hosts file of the elastic container instance.
	HostAliases []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The hostname series.
	//
	// example:
	//
	// [\\"hehe.com\\", \\"haha.com\\"]
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The information about the image repository.
	ImageRegistryCredentials []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
	// The ID of the image cache.
	//
	// example:
	//
	// imc-2zebxkiifuyzzlhl****
	ImageSnapshotId *string `json:"ImageSnapshotId,omitempty" xml:"ImageSnapshotId,omitempty"`
	// The maximum inbound public bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 1024000
	IngressBandwidth *int64 `json:"IngressBandwidth,omitempty" xml:"IngressBandwidth,omitempty"`
	// The init containers.
	InitContainers []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The level of the instance family, which is used to filter the available instance types that meet the specified requirements. This parameter takes effect only if `CostOptimization` is set to true. Valid values:
	//
	// 	- EntryLevel: entry level (shared instance types). Instance types of this level are the most cost-effective but may not provide stable computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise level. Instance types of this level provide stable performance and dedicated resources, and are suitable for business scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit-based entry level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
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
	// The status of the scaling configuration in the scaling group. Valid values:
	//
	// 	- Active: The scaling configuration is active in the scaling group. Auto Scaling uses the scaling configuration that is in the Active state to create instances during scale-out events.
	//
	// 	- Inactive: The scaling configuration is inactive in the scaling group. Scaling configurations that are in the Inactive state are still contained in the scaling group, but Auto Scaling does not use the inactive scaling configurations to create instances during scale-out events.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The load balancing weight of each elastic container instance as a backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size.
	//
	// You can specify CPU and Memory to define the range of instance types. For example, if you set CPU to 2 and Memory to 16, the instance types that have 2 vCPUs and 16 GiB are returned. After you specify CPU and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones and preferentially creates instances by using the lowest-priced instance type.
	//
	// >  You can specify CPU and Memory to define instance types only if you set Scaling Policy to Cost Optimization and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The endpoints of the Network Time Protocol (NTP) server.
	NtpServers []*string `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	// The Resource Access Management (RAM) role of elastic container instances. Elastic container instances and Elastic Compute Service (ECS) instances can share the same RAM role. For more information, see [Use an instance RAM role by calling API operations](https://help.aliyun.com/document_detail/61178.html).
	//
	// example:
	//
	// ram:PassRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of elastic container instances.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-8db03793gfrz****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The restart policy of elastic container instances. Valid values:
	//
	// 	- Never: Elastic container instances are never restarted.
	//
	// 	- Always: Elastic container instances are always restarted.
	//
	// 	- OnFailure: Elastic container instances are restarted upon failures.
	//
	// example:
	//
	// Never
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-bp1ezrfgoyn5kijl****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The name of the scaling configuration.
	//
	// example:
	//
	// scalingconfi****
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The ID of the scaling group to which the scaling configuration belongs.
	//
	// example:
	//
	// asg-bp17pelvl720x3v7****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The system information of the security context in which the elastic container instance runs.
	SecurityContextSysCtls []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls `json:"SecurityContextSysCtls,omitempty" xml:"SecurityContextSysCtls,omitempty" type:"Repeated"`
	// The ID of the security group with which elastic container instances are associated. Elastic container instances that are associated with the same security group can communicate with each other.
	//
	// example:
	//
	// sg-bp18kz60mefs****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// >  This parameter is not available for use.
	//
	// example:
	//
	// False
	SlsEnable *bool `json:"SlsEnable,omitempty" xml:"SlsEnable,omitempty"`
	// The maximum hourly price for preemptible elastic container instances.
	//
	// This parameter is returned only if you set SpotStrategy to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for elastic container instances. Valid values:
	//
	// 	- NoSpot: The instances are created as regular pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are created as preemptible instances for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags of elastic container instances. Tags are specified in the key-value format.
	Tags []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The buffer time during which a program handles operations before the program stops.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The volumes.
	Volumes []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetAcrRegistryInfos() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	return s.AcrRegistryInfos
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetActiveDeadlineSeconds() *int32 {
	return s.ActiveDeadlineSeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetAutoCreateEip() *bool {
	return s.AutoCreateEip
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetAutoMatchImageCache() *bool {
	return s.AutoMatchImageCache
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetContainerGroupName() *string {
	return s.ContainerGroupName
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetContainers() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	return s.Containers
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetCostOptimization() *bool {
	return s.CostOptimization
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetCpuOptionsCore() *int32 {
	return s.CpuOptionsCore
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetCpuOptionsThreadsPerCore() *int32 {
	return s.CpuOptionsThreadsPerCore
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDataCacheBucket() *string {
	return s.DataCacheBucket
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDataCacheBurstingEnabled() *bool {
	return s.DataCacheBurstingEnabled
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDataCachePL() *string {
	return s.DataCachePL
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDataCacheProvisionedIops() *int32 {
	return s.DataCacheProvisionedIops
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDescription() *string {
	return s.Description
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDnsConfigNameServers() []*string {
	return s.DnsConfigNameServers
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDnsConfigOptions() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions {
	return s.DnsConfigOptions
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDnsConfigSearches() []*string {
	return s.DnsConfigSearches
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetDnsPolicy() *string {
	return s.DnsPolicy
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetEgressBandwidth() *int64 {
	return s.EgressBandwidth
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetEphemeralStorage() *int32 {
	return s.EphemeralStorage
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetHostAliases() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases {
	return s.HostAliases
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetHostName() *string {
	return s.HostName
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetImageRegistryCredentials() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	return s.ImageRegistryCredentials
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetImageSnapshotId() *string {
	return s.ImageSnapshotId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetIngressBandwidth() *int64 {
	return s.IngressBandwidth
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetInitContainers() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	return s.InitContainers
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetNtpServers() []*string {
	return s.NtpServers
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetRestartPolicy() *string {
	return s.RestartPolicy
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetSecurityContextSysCtls() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls {
	return s.SecurityContextSysCtls
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetSlsEnable() *bool {
	return s.SlsEnable
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetTags() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags {
	return s.Tags
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) GetVolumes() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	return s.Volumes
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetAcrRegistryInfos(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.AcrRegistryInfos = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetActiveDeadlineSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetAutoCreateEip(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.AutoCreateEip = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetAutoMatchImageCache(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.AutoMatchImageCache = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetContainerGroupName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetContainers(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Containers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCostOptimization(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CostOptimization = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCpu(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCpuOptionsCore(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CpuOptionsCore = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCpuOptionsThreadsPerCore(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetCreationTime(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.CreationTime = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDataCacheBucket(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DataCacheBucket = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDataCacheBurstingEnabled(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DataCacheBurstingEnabled = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDataCachePL(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DataCachePL = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDataCacheProvisionedIops(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DataCacheProvisionedIops = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDescription(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Description = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsConfigNameServers(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsConfigNameServers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsConfigOptions(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsConfigOptions = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsConfigSearches(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsConfigSearches = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetDnsPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.DnsPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetEgressBandwidth(v int64) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.EgressBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetEipBandwidth(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetEphemeralStorage(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetGpuDriverVersion(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetHostAliases(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.HostAliases = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetHostName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.HostName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetImageRegistryCredentials(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageRegistryCredentials = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetImageSnapshotId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageSnapshotId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetIngressBandwidth(v int64) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.IngressBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetInitContainers(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.InitContainers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceFamilyLevel(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceTypes(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceTypes = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetIpv6AddressCount(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetLifecycleState(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.LifecycleState = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetLoadBalancerWeight(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetMemory(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetNtpServers(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.NtpServers = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetRamRoleName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.RamRoleName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetRegionId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetResourceGroupId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetRestartPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.RestartPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetScalingGroupId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityContextSysCtls(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityContextSysCtls = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityGroupId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSlsEnable(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SlsEnable = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSpotPriceLimit(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetSpotStrategy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetTags(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Tags = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetTerminationGracePeriodSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) SetVolumes(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations {
	s.Volumes = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurations) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos struct {
	// The domain names of the Container Registry Enterprise Edition instance. By default, all domain names of the instance are displayed. Multiple domain names are separated by commas (,).
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetDomains(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetInstanceId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetInstanceName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) SetRegionId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsAcrRegistryInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers struct {
	// The container startup arguments. You can specify up to 10 arguments.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container startup commands. You can specify up to 20 commands. Each command can contain up to 256 characters.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs per container.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables.
	EnvironmentVars []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The container image.
	//
	// example:
	//
	// mysql
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy. Valid values:
	//
	// 	- Always: Each time image pulling is performed.
	//
	// 	- IfNotPresent: Image pulling is performed as needed. On-premises images are preferentially used. If no on-premises images are available, image pulling is performed.
	//
	// 	- Never: On-premises images are always used. Image pulling is not performed.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The commands that are run for configuring the postStart callback function by using the CLI within the container.
	LifecyclePostStartHandlerExecs []*string `json:"LifecyclePostStartHandlerExecs,omitempty" xml:"LifecyclePostStartHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which HTTP GET requests for configuring the postStart callback function are sent.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The path to which HTTP GET requests for configuring the postStart callback function are sent.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port over which HTTP GET requests for configuring the postStart callback function are sent.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP GET requests that are sent for configuring the postStart callback function.
	//
	// example:
	//
	// HTTP
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP sockets used for configuring the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP sockets used for configuring the postStart callback function.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The commands that are run for configuring the preStop callback function by using the CLI within the container.
	LifecyclePreStopHandlerExecs []*string `json:"LifecyclePreStopHandlerExecs,omitempty" xml:"LifecyclePreStopHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which HTTP GET requests for configuring the preStop callback function are sent.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The path to which HTTP GET requests for configuring the preStop callback function are sent.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port over which HTTP GET requests for configuring the preStop callback function are sent.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP Get requests that are sent for configuring the preStop callback function.
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP sockets used for configuring the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP sockets used for configuring the preStop callback function.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The commands that are run in the container when you use the CLI to perform liveness probes.
	LivenessProbeExecCommands []*string `json:"LivenessProbeExecCommands,omitempty" xml:"LivenessProbeExecCommands,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures before a successful liveness probe is considered failed.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	LivenessProbeFailureThreshold *int32 `json:"LivenessProbeFailureThreshold,omitempty" xml:"LivenessProbeFailureThreshold,omitempty"`
	// The path to which HTTP GET requests are sent when you use the HTTP GET requests to perform liveness probes.
	//
	// example:
	//
	// /usr/nginx/
	LivenessProbeHttpGetPath *string `json:"LivenessProbeHttpGetPath,omitempty" xml:"LivenessProbeHttpGetPath,omitempty"`
	// The port to which HTTP GET requests are sent to perform liveness probes.
	//
	// example:
	//
	// 80
	LivenessProbeHttpGetPort *int32 `json:"LivenessProbeHttpGetPort,omitempty" xml:"LivenessProbeHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use the HTTP GET requests to perform liveness probes. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	LivenessProbeHttpGetScheme *string `json:"LivenessProbeHttpGetScheme,omitempty" xml:"LivenessProbeHttpGetScheme,omitempty"`
	// The number of seconds that elapses from the startup of the container to the start time of a liveness probe.
	//
	// example:
	//
	// 10
	LivenessProbeInitialDelaySeconds *int32 `json:"LivenessProbeInitialDelaySeconds,omitempty" xml:"LivenessProbeInitialDelaySeconds,omitempty"`
	// The interval at which liveness probes are performed. Unit: seconds. Default value: 10. Minimum value: 1.
	//
	// example:
	//
	// 5
	LivenessProbePeriodSeconds *int32 `json:"LivenessProbePeriodSeconds,omitempty" xml:"LivenessProbePeriodSeconds,omitempty"`
	// The minimum number of consecutive successes before a failed liveness probe is considered successful. Default value: 1. Valid value: 1.
	//
	// example:
	//
	// 1
	LivenessProbeSuccessThreshold *int32 `json:"LivenessProbeSuccessThreshold,omitempty" xml:"LivenessProbeSuccessThreshold,omitempty"`
	// The port detected by TCP sockets when you use the TCP sockets to perform liveness probes.
	//
	// example:
	//
	// 80
	LivenessProbeTcpSocketPort *int32 `json:"LivenessProbeTcpSocketPort,omitempty" xml:"LivenessProbeTcpSocketPort,omitempty"`
	// The timeout period of a liveness probe. Default value: 1. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 10
	LivenessProbeTimeoutSeconds *int32 `json:"LivenessProbeTimeoutSeconds,omitempty" xml:"LivenessProbeTimeoutSeconds,omitempty"`
	// The memory size per container.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The custom name of the container.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The exposed ports and protocols.
	Ports []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The commands that are run in the container when you use the CLI to perform readiness probes.
	ReadinessProbeExecCommands []*string `json:"ReadinessProbeExecCommands,omitempty" xml:"ReadinessProbeExecCommands,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures before a successful readiness probe is considered failed.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	ReadinessProbeFailureThreshold *int32 `json:"ReadinessProbeFailureThreshold,omitempty" xml:"ReadinessProbeFailureThreshold,omitempty"`
	// The path to which HTTP GET requests are sent when you use the HTTP GET requests to perform readiness probes.
	//
	// example:
	//
	// /usr/local
	ReadinessProbeHttpGetPath *string `json:"ReadinessProbeHttpGetPath,omitempty" xml:"ReadinessProbeHttpGetPath,omitempty"`
	// The path to which HTTP GET requests are sent when you use the HTTP GET requests to perform readiness probes.
	//
	// example:
	//
	// /usr/nginx/
	ReadinessProbeHttpGetPort *int32 `json:"ReadinessProbeHttpGetPort,omitempty" xml:"ReadinessProbeHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use the HTTP requests to perform readiness probes. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	ReadinessProbeHttpGetScheme *string `json:"ReadinessProbeHttpGetScheme,omitempty" xml:"ReadinessProbeHttpGetScheme,omitempty"`
	// The number of seconds that elapses from the startup of the container to the start time of a readiness probe.
	//
	// example:
	//
	// 5
	ReadinessProbeInitialDelaySeconds *int32 `json:"ReadinessProbeInitialDelaySeconds,omitempty" xml:"ReadinessProbeInitialDelaySeconds,omitempty"`
	// The interval at which readiness probes are performed. Unit: seconds. Default value: 10. Minimum value: 1.
	//
	// example:
	//
	// 1
	ReadinessProbePeriodSeconds *int32 `json:"ReadinessProbePeriodSeconds,omitempty" xml:"ReadinessProbePeriodSeconds,omitempty"`
	// The minimum number of consecutive successes before a failed readiness probe is considered successful. Default value: 1. Valid value: 1.
	//
	// example:
	//
	// 1
	ReadinessProbeSuccessThreshold *int32 `json:"ReadinessProbeSuccessThreshold,omitempty" xml:"ReadinessProbeSuccessThreshold,omitempty"`
	// The port detected by TCP sockets when you use the TCP sockets to perform readiness probes.
	//
	// example:
	//
	// 8888
	ReadinessProbeTcpSocketPort *int32 `json:"ReadinessProbeTcpSocketPort,omitempty" xml:"ReadinessProbeTcpSocketPort,omitempty"`
	// The timeout period of a readiness probe. Default value: 1. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 5
	ReadinessProbeTimeoutSeconds *int32 `json:"ReadinessProbeTimeoutSeconds,omitempty" xml:"ReadinessProbeTimeoutSeconds,omitempty"`
	// The permissions that are granted to the processes in the container. Valid values: NET_ADMIN and NET_RAW.
	//
	// >  To use NET_RAW, you must submit a ticket.
	SecurityContextCapabilityAdds []*string `json:"SecurityContextCapabilityAdds,omitempty" xml:"SecurityContextCapabilityAdds,omitempty" type:"Repeated"`
	// Indicates whether the root file system on which the container runs is read-only. Valid value: true.
	//
	// example:
	//
	// true
	SecurityContextReadOnlyRootFilesystem *bool `json:"SecurityContextReadOnlyRootFilesystem,omitempty" xml:"SecurityContextReadOnlyRootFilesystem,omitempty"`
	// The ID of the user that runs the entry point of the container process.
	//
	// example:
	//
	// 1000
	SecurityContextRunAsUser *int64 `json:"SecurityContextRunAsUser,omitempty" xml:"SecurityContextRunAsUser,omitempty"`
	// Indicates whether the container allocates buffer resources to standard input streams when the container is running. If this parameter is not specified, an end-of-file (EOF) error may occur when standard input streams in the container are read. Default value: false.
	//
	// example:
	//
	// true
	Stdin *bool `json:"Stdin,omitempty" xml:"Stdin,omitempty"`
	// Indicates whether standard input streams are disconnected after a client is disconnected. If Stdin is set to true, standard input streams remain connected among multiple sessions.
	//
	// If StdinOnce is set to true, standard input streams are connected after the container is started, and remain idle until a client is connected to receive data. After the client is disconnected, streams are also disconnected, and remain disconnected until the container restarts.
	//
	// example:
	//
	// true
	StdinOnce *bool `json:"StdinOnce,omitempty" xml:"StdinOnce,omitempty"`
	// Indicates whether the Interaction feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// If the command is a /bin/bash command, the value of this parameter is true.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The mounted volumes.
	VolumeMounts []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory.
	//
	// example:
	//
	// /usr/local/nginx
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetArgs() []*string {
	return s.Args
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetCommands() []*string {
	return s.Commands
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetEnvironmentVars() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	return s.EnvironmentVars
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetImage() *string {
	return s.Image
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerExecs() []*string {
	return s.LifecyclePostStartHandlerExecs
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerHttpGetHost() *string {
	return s.LifecyclePostStartHandlerHttpGetHost
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerHttpGetPath() *string {
	return s.LifecyclePostStartHandlerHttpGetPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerHttpGetPort() *int32 {
	return s.LifecyclePostStartHandlerHttpGetPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerHttpGetScheme() *string {
	return s.LifecyclePostStartHandlerHttpGetScheme
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerTcpSocketHost() *string {
	return s.LifecyclePostStartHandlerTcpSocketHost
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePostStartHandlerTcpSocketPort() *int32 {
	return s.LifecyclePostStartHandlerTcpSocketPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerExecs() []*string {
	return s.LifecyclePreStopHandlerExecs
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerHttpGetHost() *string {
	return s.LifecyclePreStopHandlerHttpGetHost
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerHttpGetPath() *string {
	return s.LifecyclePreStopHandlerHttpGetPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerHttpGetPort() *int32 {
	return s.LifecyclePreStopHandlerHttpGetPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerHttpGetScheme() *string {
	return s.LifecyclePreStopHandlerHttpGetScheme
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerTcpSocketHost() *string {
	return s.LifecyclePreStopHandlerTcpSocketHost
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLifecyclePreStopHandlerTcpSocketPort() *int32 {
	return s.LifecyclePreStopHandlerTcpSocketPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeExecCommands() []*string {
	return s.LivenessProbeExecCommands
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeFailureThreshold() *int32 {
	return s.LivenessProbeFailureThreshold
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeHttpGetPath() *string {
	return s.LivenessProbeHttpGetPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeHttpGetPort() *int32 {
	return s.LivenessProbeHttpGetPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeHttpGetScheme() *string {
	return s.LivenessProbeHttpGetScheme
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeInitialDelaySeconds() *int32 {
	return s.LivenessProbeInitialDelaySeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbePeriodSeconds() *int32 {
	return s.LivenessProbePeriodSeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeSuccessThreshold() *int32 {
	return s.LivenessProbeSuccessThreshold
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeTcpSocketPort() *int32 {
	return s.LivenessProbeTcpSocketPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetLivenessProbeTimeoutSeconds() *int32 {
	return s.LivenessProbeTimeoutSeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetPorts() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts {
	return s.Ports
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeExecCommands() []*string {
	return s.ReadinessProbeExecCommands
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeFailureThreshold() *int32 {
	return s.ReadinessProbeFailureThreshold
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeHttpGetPath() *string {
	return s.ReadinessProbeHttpGetPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeHttpGetPort() *int32 {
	return s.ReadinessProbeHttpGetPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeHttpGetScheme() *string {
	return s.ReadinessProbeHttpGetScheme
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeInitialDelaySeconds() *int32 {
	return s.ReadinessProbeInitialDelaySeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbePeriodSeconds() *int32 {
	return s.ReadinessProbePeriodSeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeSuccessThreshold() *int32 {
	return s.ReadinessProbeSuccessThreshold
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeTcpSocketPort() *int32 {
	return s.ReadinessProbeTcpSocketPort
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetReadinessProbeTimeoutSeconds() *int32 {
	return s.ReadinessProbeTimeoutSeconds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetSecurityContextCapabilityAdds() []*string {
	return s.SecurityContextCapabilityAdds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetSecurityContextReadOnlyRootFilesystem() *bool {
	return s.SecurityContextReadOnlyRootFilesystem
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetSecurityContextRunAsUser() *int64 {
	return s.SecurityContextRunAsUser
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetStdin() *bool {
	return s.Stdin
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetStdinOnce() *bool {
	return s.StdinOnce
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetTty() *bool {
	return s.Tty
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetVolumeMounts() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	return s.VolumeMounts
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetArgs(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Args = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Commands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetCpu(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetEnvironmentVars(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetGpu(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetImage(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Image = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetImagePullPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerExecs(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerExecs = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerHttpGetHost(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerHttpGetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerHttpGetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerHttpGetScheme(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerTcpSocketHost(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerExecs(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerExecs = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerHttpGetHost(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerHttpGetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerHttpGetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerHttpGetScheme(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerTcpSocketHost(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeExecCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeExecCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeFailureThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeFailureThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeHttpGetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeHttpGetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeHttpGetScheme(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeInitialDelaySeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeInitialDelaySeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbePeriodSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeSuccessThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeSuccessThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeTcpSocketPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetLivenessProbeTimeoutSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.LivenessProbeTimeoutSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetMemory(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetPorts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Ports = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeExecCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeExecCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeFailureThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeFailureThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeHttpGetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeHttpGetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeHttpGetScheme(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeInitialDelaySeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeInitialDelaySeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbePeriodSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeSuccessThreshold(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeSuccessThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeTcpSocketPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetReadinessProbeTimeoutSeconds(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.ReadinessProbeTimeoutSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetSecurityContextCapabilityAdds(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.SecurityContextCapabilityAdds = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetSecurityContextReadOnlyRootFilesystem(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.SecurityContextReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetSecurityContextRunAsUser(v int64) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.SecurityContextRunAsUser = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetStdin(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Stdin = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetStdinOnce(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.StdinOnce = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetTty(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.Tty = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetVolumeMounts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) SetWorkingDir(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers {
	s.WorkingDir = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainers) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars struct {
	// >  This parameter is not available for use.
	//
	// example:
	//
	// path
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// /usr/bin/
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) SetKey(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts struct {
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8888
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) SetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) SetProtocol(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts struct {
	// The directory to which the volume is mounted.
	//
	// >  Data in this directory is overwritten by the data on the volume. Proceed with caution if you specify this parameter.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation settings. Mount propagation enables volumes mounted on one container to be shared among other containers within the same pod or across distinct pods residing on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed either on the volume itself or its subdirectories do not propagate to the already established volume mount.
	//
	// 	- HostToCotainer: Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount. In addition, any volume mounts executed on a container not only propagate back to the underlying host but also to all containers across every pod that uses the same volume.
	//
	// Default value: None.
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
	// Indicates whether the volume is read-only.
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetMountPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetMountPropagation(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetReadOnly(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) SetSubPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsContainersVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions struct {
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsDnsConfigOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases struct {
	// The added hostnames.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The added IP address.
	//
	// example:
	//
	// 192.0.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) GetHostnames() []*string {
	return s.Hostnames
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) GetIp() *string {
	return s.Ip
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) SetHostnames(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases {
	s.Hostnames = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) SetIp(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases {
	s.Ip = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsHostAliases) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials struct {
	// The password of the image repository.
	//
	// example:
	//
	// yourpaasword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The domain name of the image repository.
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) GetPassword() *string {
	return s.Password
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) GetServer() *string {
	return s.Server
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) GetUserName() *string {
	return s.UserName
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) SetPassword(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) SetServer(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) SetUserName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials {
	s.UserName = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsImageRegistryCredentials) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers struct {
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
	// The container startup arguments.
	InitContainerArgs []*string `json:"InitContainerArgs,omitempty" xml:"InitContainerArgs,omitempty" type:"Repeated"`
	// The container startup commands.
	InitContainerCommands []*string `json:"InitContainerCommands,omitempty" xml:"InitContainerCommands,omitempty" type:"Repeated"`
	// The environment variables.
	InitContainerEnvironmentVars []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	// The ports of the init container.
	InitContainerPorts []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	// The volumes that are mounted to the init container.
	InitContainerVolumeMounts []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	// The memory size per init container.
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
	// The permissions that are granted to the processes in the init container. Valid values: NET_ADMIN and NET_RAW.
	//
	// >  To use NET_RAW, you must submit a ticket.
	SecurityContextCapabilityAdds []*string `json:"SecurityContextCapabilityAdds,omitempty" xml:"SecurityContextCapabilityAdds,omitempty" type:"Repeated"`
	// Indicates whether the root file system is read-only. Valid value: true.
	//
	// example:
	//
	// true
	SecurityContextReadOnlyRootFilesystem *bool `json:"SecurityContextReadOnlyRootFilesystem,omitempty" xml:"SecurityContextReadOnlyRootFilesystem,omitempty"`
	// The ID of the user that runs the init container.
	//
	// example:
	//
	// 587
	SecurityContextRunAsUser *string `json:"SecurityContextRunAsUser,omitempty" xml:"SecurityContextRunAsUser,omitempty"`
	// The working directory.
	//
	// example:
	//
	// /usr/local
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetImage() *string {
	return s.Image
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetInitContainerArgs() []*string {
	return s.InitContainerArgs
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetInitContainerCommands() []*string {
	return s.InitContainerCommands
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetInitContainerEnvironmentVars() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	return s.InitContainerEnvironmentVars
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetInitContainerPorts() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts {
	return s.InitContainerPorts
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetInitContainerVolumeMounts() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	return s.InitContainerVolumeMounts
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetSecurityContextCapabilityAdds() []*string {
	return s.SecurityContextCapabilityAdds
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetSecurityContextReadOnlyRootFilesystem() *bool {
	return s.SecurityContextReadOnlyRootFilesystem
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetSecurityContextRunAsUser() *string {
	return s.SecurityContextRunAsUser
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetCpu(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetGpu(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetImage(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Image = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetImagePullPolicy(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerArgs(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerArgs = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerCommands(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerEnvironmentVars(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerPorts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetInitContainerVolumeMounts(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetMemory(v float32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetSecurityContextCapabilityAdds(v []*string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.SecurityContextCapabilityAdds = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetSecurityContextReadOnlyRootFilesystem(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.SecurityContextReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetSecurityContextRunAsUser(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.SecurityContextRunAsUser = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) SetWorkingDir(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers {
	s.WorkingDir = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainers) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars struct {
	// >  This parameter is not available for use.
	//
	// example:
	//
	// path
	FieldRefFieldPath *string `json:"FieldRefFieldPath,omitempty" xml:"FieldRefFieldPath,omitempty"`
	// The name of the environment variable.
	//
	// example:
	//
	// PATH
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the environment variable.
	//
	// example:
	//
	// /usr/local/bin
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) SetKey(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts struct {
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) SetPort(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) SetProtocol(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts struct {
	// The directory to which the volume is mounted. Data under this directory is overwritten by the data on the volume.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting. Mount propagation enables volumes mounted on one container to be shared among other containers within the same pod or across distinct pods residing on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed either on the volume itself or its subdirectories do not propagate to the already established volume mount.
	//
	// 	- HostToCotainer: Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount. In addition, any volume mounts executed on a container not only propagate back to the underlying host but also to all containers across every pod that uses the same volume.
	//
	// Default value: None.
	//
	// example:
	//
	// None
	MountPropagation *string `json:"MountPropagation,omitempty" xml:"MountPropagation,omitempty"`
	// The volume name.
	//
	// example:
	//
	// test-empty
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the mount directory is read-only.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The volume subdirectory. A pod can mount different directories of the same volume to different directories of the init container.
	//
	// example:
	//
	// /usr/sub/
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetMountPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) SetSubPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsInitContainersInitContainerVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls struct {
	// The system name of the security context in which the elastic container instance runs.
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsSecurityContextSysCtls) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags struct {
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) SetKey(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) SetValue(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsTags) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes struct {
	// The paths to configuration files.
	ConfigFileVolumeConfigFileToPaths []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	// The default permissions on the ConfigFile volume.
	//
	// example:
	//
	// 0644
	ConfigFileVolumeDefaultMode *int32 `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	// The size of the disk volume. Unit: GiB.
	//
	// example:
	//
	// 15
	DiskVolumeDiskId *string `json:"DiskVolumeDiskId,omitempty" xml:"DiskVolumeDiskId,omitempty"`
	// The size of the disk volume. Unit: GiB.
	//
	// example:
	//
	// 15
	DiskVolumeDiskSize *int32 `json:"DiskVolumeDiskSize,omitempty" xml:"DiskVolumeDiskSize,omitempty"`
	// The system type of the disk volume.
	//
	// example:
	//
	// xfs
	DiskVolumeFsType *string `json:"DiskVolumeFsType,omitempty" xml:"DiskVolumeFsType,omitempty"`
	// The storage medium of the emptyDir volume. If you do not specify a storage medium for the emptyDir volume, the volume stores data in the file system of a node by default. We recommend that you set this parameter to memory. In this case, the emptyDir volume stores data in the specified memory.
	//
	// example:
	//
	// memory
	EmptyDirVolumeMedium *string `json:"EmptyDirVolumeMedium,omitempty" xml:"EmptyDirVolumeMedium,omitempty"`
	// The storage size of the emptyDir volume.
	//
	// example:
	//
	// 256Mi
	EmptyDirVolumeSizeLimit *string `json:"EmptyDirVolumeSizeLimit,omitempty" xml:"EmptyDirVolumeSizeLimit,omitempty"`
	// The name of the FlexVolume driver.
	//
	// example:
	//
	// flexvolume
	FlexVolumeDriver *string `json:"FlexVolumeDriver,omitempty" xml:"FlexVolumeDriver,omitempty"`
	// The type of the mounted file system. The default value is determined by the script of FlexVolume.
	//
	// example:
	//
	// ext4
	FlexVolumeFsType *string `json:"FlexVolumeFsType,omitempty" xml:"FlexVolumeFsType,omitempty"`
	// The FlexVolume options. Each option is a key-value pair in a JSON string.
	//
	// For example, if you use FlexVolume to mount a disk, the format of Options is `{"volumeId":"d-2zehdahrwoa7srg****","performanceLevel": "PL2"}`.
	//
	// example:
	//
	// {"volumeId":"d-2zehdahrwoa7srg****","performanceLevel": "PL2"}
	FlexVolumeOptions *string `json:"FlexVolumeOptions,omitempty" xml:"FlexVolumeOptions,omitempty"`
	// The path to the HostPath volume on the host.
	//
	// example:
	//
	// /pod/data
	HostPathVolumePath *string `json:"HostPathVolumePath,omitempty" xml:"HostPathVolumePath,omitempty"`
	// The type of the HostPath volume.
	//
	// example:
	//
	// Directory
	HostPathVolumeType *string `json:"HostPathVolumeType,omitempty" xml:"HostPathVolumeType,omitempty"`
	// The path to the Network File System (NFS) volume.
	//
	// example:
	//
	// /share
	NFSVolumePath *string `json:"NFSVolumePath,omitempty" xml:"NFSVolumePath,omitempty"`
	// Indicates whether the NFS volume is read-only.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	NFSVolumeReadOnly *bool `json:"NFSVolumeReadOnly,omitempty" xml:"NFSVolumeReadOnly,omitempty"`
	// The endpoint of the NFS server.
	//
	// example:
	//
	// 3f9cd4a596-naw76.cn-shanghai.nas.aliyuncs.com
	NFSVolumeServer *string `json:"NFSVolumeServer,omitempty" xml:"NFSVolumeServer,omitempty"`
	// The volume name.
	//
	// example:
	//
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The volume type. Valid values:
	//
	// 	- EmptyDirVolume
	//
	// 	- NFSVolume
	//
	// 	- ConfigFileVolume
	//
	// 	- FlexVolume
	//
	// example:
	//
	// EmptyDirVolume
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetConfigFileVolumeConfigFileToPaths() []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	return s.ConfigFileVolumeConfigFileToPaths
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetConfigFileVolumeDefaultMode() *int32 {
	return s.ConfigFileVolumeDefaultMode
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetDiskVolumeDiskId() *string {
	return s.DiskVolumeDiskId
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetDiskVolumeDiskSize() *int32 {
	return s.DiskVolumeDiskSize
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetDiskVolumeFsType() *string {
	return s.DiskVolumeFsType
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetEmptyDirVolumeMedium() *string {
	return s.EmptyDirVolumeMedium
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetEmptyDirVolumeSizeLimit() *string {
	return s.EmptyDirVolumeSizeLimit
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetFlexVolumeDriver() *string {
	return s.FlexVolumeDriver
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetFlexVolumeFsType() *string {
	return s.FlexVolumeFsType
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetFlexVolumeOptions() *string {
	return s.FlexVolumeOptions
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetHostPathVolumePath() *string {
	return s.HostPathVolumePath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetHostPathVolumeType() *string {
	return s.HostPathVolumeType
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetNFSVolumePath() *string {
	return s.NFSVolumePath
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetNFSVolumeReadOnly() *bool {
	return s.NFSVolumeReadOnly
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetNFSVolumeServer() *string {
	return s.NFSVolumeServer
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) GetType() *string {
	return s.Type
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetConfigFileVolumeConfigFileToPaths(v []*DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetConfigFileVolumeDefaultMode(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetDiskVolumeDiskId(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.DiskVolumeDiskId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetDiskVolumeDiskSize(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.DiskVolumeDiskSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetDiskVolumeFsType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.DiskVolumeFsType = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetEmptyDirVolumeMedium(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.EmptyDirVolumeMedium = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetEmptyDirVolumeSizeLimit(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.EmptyDirVolumeSizeLimit = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetFlexVolumeDriver(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.FlexVolumeDriver = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetFlexVolumeFsType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.FlexVolumeFsType = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetFlexVolumeOptions(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.FlexVolumeOptions = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetHostPathVolumePath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.HostPathVolumePath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetHostPathVolumeType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.HostPathVolumeType = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetNFSVolumePath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.NFSVolumePath = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetNFSVolumeReadOnly(v bool) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.NFSVolumeReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetNFSVolumeServer(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.NFSVolumeServer = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetName(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) SetType(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes {
	s.Type = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumes) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths struct {
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

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) GetContent() *string {
	return s.Content
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) GetMode() *int32 {
	return s.Mode
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) GetPath() *string {
	return s.Path
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) SetMode(v int32) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	s.Mode = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

func (s *DescribeEciScalingConfigurationsResponseBodyScalingConfigurationsVolumesConfigFileVolumeConfigFileToPaths) Validate() error {
	return dara.Validate(s)
}
