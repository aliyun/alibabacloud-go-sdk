// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEciScalingConfigurationDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutput(v string) *DescribeEciScalingConfigurationDetailResponseBody
	GetOutput() *string
	SetRequestId(v string) *DescribeEciScalingConfigurationDetailResponseBody
	GetRequestId() *string
	SetScalingConfiguration(v *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) *DescribeEciScalingConfigurationDetailResponseBody
	GetScalingConfiguration() *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration
}

type DescribeEciScalingConfigurationDetailResponseBody struct {
	// The YAML output of the scaling configuration.
	//
	// example:
	//
	// apiVersion: apps/v1
	//
	// kind: Deployment
	//
	// metadata:
	//
	//   name: nginx-deployment
	//
	//   labels:
	//
	//     app: nginx
	//
	//   spec:
	//
	//     replicas: 3
	//
	//     selector:
	//
	//        matchLabels:
	//
	//         app: nginx
	//
	//     template:
	//
	//       metadata:
	//
	//         labels:
	//
	//           app: nginx
	//
	//         annotations:
	//
	//           k8s.aliyun.com/eip-bandwidth: 10
	//
	//           k8s.aliyun.com/eci-with-eip: true
	//
	//         spec:
	//
	//           containers:
	//
	//           - name: nginx
	//
	//             image: nginx:1.14.2
	//
	//             ports:
	//
	//             - containerPort: 80
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EF9BFEE-FE07-4627-B8FB-14326FB9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the scaling configuration.
	ScalingConfiguration *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration `json:"ScalingConfiguration,omitempty" xml:"ScalingConfiguration,omitempty" type:"Struct"`
}

func (s DescribeEciScalingConfigurationDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) GetOutput() *string {
	return s.Output
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) GetScalingConfiguration() *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	return s.ScalingConfiguration
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) SetOutput(v string) *DescribeEciScalingConfigurationDetailResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) SetRequestId(v string) *DescribeEciScalingConfigurationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) SetScalingConfiguration(v *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) *DescribeEciScalingConfigurationDetailResponseBody {
	s.ScalingConfiguration = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBody) Validate() error {
	if s.ScalingConfiguration != nil {
		if err := s.ScalingConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration struct {
	// The information about the Container Registry Enterprise Edition instance.
	AcrRegistryInfos []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos `json:"AcrRegistryInfos,omitempty" xml:"AcrRegistryInfos,omitempty" type:"Repeated"`
	// The validity period of the scaling configuration. Unit: seconds.
	//
	// example:
	//
	// 60
	ActiveDeadlineSeconds *int32 `json:"ActiveDeadlineSeconds,omitempty" xml:"ActiveDeadlineSeconds,omitempty"`
	// Indicates whether an elastic IP address (EIP) is automatically created and bound to the elastic container instance.
	//
	// example:
	//
	// true
	AutoCreateEip *bool `json:"AutoCreateEip,omitempty" xml:"AutoCreateEip,omitempty"`
	// Indicates whether the image cache is automatically matched. Default value: false.
	//
	// example:
	//
	// true
	AutoMatchImageCache *bool `json:"AutoMatchImageCache,omitempty" xml:"AutoMatchImageCache,omitempty"`
	// The computing power types. A value of economy indicates that economic instance types are returned.
	ComputeCategory []*string `json:"ComputeCategory,omitempty" xml:"ComputeCategory,omitempty" type:"Repeated"`
	// The name of the container group.
	//
	// example:
	//
	// test
	ContainerGroupName *string `json:"ContainerGroupName,omitempty" xml:"ContainerGroupName,omitempty"`
	// The containers in the elastic container instance.
	Containers []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
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
	// The number of vCPUs that are allocated to the elastic container instance.
	//
	// example:
	//
	// 2.0
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of physical CPU cores. You can specify this parameter for only specific instance types.
	//
	// example:
	//
	// 2
	CpuOptionsCore *int32 `json:"CpuOptionsCore,omitempty" xml:"CpuOptionsCore,omitempty"`
	// The number of threads per core. You can specify this parameter for only specific instance types. A value of 1 indicates that Hyper-Threading is disabled. For more information, see [Specify CPU options](https://help.aliyun.com/document_detail/197781.html).
	//
	// example:
	//
	// 2
	CpuOptionsThreadsPerCore *int32 `json:"CpuOptionsThreadsPerCore,omitempty" xml:"CpuOptionsThreadsPerCore,omitempty"`
	// The time when the scaling configuration was created.
	//
	// example:
	//
	// 2023-05-10T02:39:15Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The bucket that caches data.
	//
	// example:
	//
	// default
	DataCacheBucket *string `json:"DataCacheBucket,omitempty" xml:"DataCacheBucket,omitempty"`
	// Indicates whether the Performance Burst feature is enabled for the ESSD AutoPL disk that caches data. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  For more information about ESSD AutoPL disks, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// false
	DataCacheBurstingEnabled *bool `json:"DataCacheBurstingEnabled,omitempty" xml:"DataCacheBurstingEnabled,omitempty"`
	// The performance level (PL) of the cloud disk that caches data. We recommend that you use enhanced SSDs (ESSDs). Valid values:
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
	// The provisioned read/write IOPS of the ESSD AutoPL disk that caches data. Valid values: 0 to min{50,000, 1,000 x *Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50	- x Capacity, 50,000}.
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
	// The IP addresses of DNS servers.
	DnsConfigNameServers []*string `json:"DnsConfigNameServers,omitempty" xml:"DnsConfigNameServers,omitempty" type:"Repeated"`
	// The DNS options.
	DnsConfigOptions []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions `json:"DnsConfigOptions,omitempty" xml:"DnsConfigOptions,omitempty" type:"Repeated"`
	// The search domains of the DNS servers.
	DnsConfigSearches []*string `json:"DnsConfigSearches,omitempty" xml:"DnsConfigSearches,omitempty" type:"Repeated"`
	// The Domain Name System (DNS) policy.
	//
	// example:
	//
	// Default
	DnsPolicy *string `json:"DnsPolicy,omitempty" xml:"DnsPolicy,omitempty"`
	// The maximum outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 10485760
	EgressBandwidth *int64 `json:"EgressBandwidth,omitempty" xml:"EgressBandwidth,omitempty"`
	// The bandwidth of the EIP. Default value: 5. Unit: Mbit/s.
	//
	// example:
	//
	// 5
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// The bound EIP bandwidth plan.
	//
	// example:
	//
	// cbwp-bp1rxai1z4b1an454xl8m
	EipCommonBandwidthPackage *string `json:"EipCommonBandwidthPackage,omitempty" xml:"EipCommonBandwidthPackage,omitempty"`
	// The line type of the EIP. Valid values:
	//
	// 	- BGP: BGP (Multi-ISP) lines
	//
	// 	- BGP_PRO: BGP (Multi-ISP) Pro
	//
	// example:
	//
	// BGP
	EipISP *string `json:"EipISP,omitempty" xml:"EipISP,omitempty"`
	// The ID of the IP address pool.
	//
	// example:
	//
	// pippool-bp187arfugi543y1s****
	EipPublicIpAddressPoolId *string `json:"EipPublicIpAddressPoolId,omitempty" xml:"EipPublicIpAddressPoolId,omitempty"`
	// The size of the temporary storage space. Unit: GiB.
	//
	// example:
	//
	// 20
	EphemeralStorage *int32 `json:"EphemeralStorage,omitempty" xml:"EphemeralStorage,omitempty"`
	// The custom hostname mappings of a container in the elastic container instance.
	HostAliases []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// hostname
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image repositories.
	ImageRegistryCredentials []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials `json:"ImageRegistryCredentials,omitempty" xml:"ImageRegistryCredentials,omitempty" type:"Repeated"`
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
	InitContainers []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// The level of the instance family, which is used to filter instance types that meet the specified criteria. This parameter takes effect only if `CostOptimization` is set to true. Valid values:
	//
	// 	- EntryLevel: entry level (shared instance types). Instance types of this level are the most cost-effective but may not provide stable computing performance in a consistent manner. Instance types of this level are suitable for business scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise level. Instance types of this level provide stable performance and dedicated resources and are suitable for business scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit entry level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview](https://help.aliyun.com/document_detail/59977.html) of burstable instances.
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
	// The state of the scaling configuration in the scaling group. Valid values:
	//
	// 	- Active: The scaling configuration is active in the scaling group. Auto Scaling uses the active scaling configuration to automatically create elastic container instances.
	//
	// 	- Inactive: The scaling configuration is inactive in the scaling group. Inactive scaling configurations are retained in scaling groups. However, Auto Scaling does not use inactive scaling groups to create elastic container instances.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The weight of an elastic container instance as a Server Load Balancer (SLB) backend server. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify CPU and Memory to define the range of instance types. For example, if you set CPU to 2 and Memory to 16, the instance types that have 2 vCPUs and 16 GiB are returned. If you specify CPU and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones and preferentially creates instances by using the lowest-priced instance type.
	//
	// >  You can specify CPU and Memory to define instance types only when you set Scaling Policy to Cost Optimization and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The endpoints of the Network Time Protocol (NTP) servers.
	NtpServers []*string `json:"NtpServers,omitempty" xml:"NtpServers,omitempty" type:"Repeated"`
	// The Resource Access Management (RAM) role of the elastic container instance. Elastic container instances and Elastic Compute Service (ECS) instances can share the same RAM role. For more information, see [Use the instance RAM role by calling APIs](https://help.aliyun.com/document_detail/61178.html).
	//
	// example:
	//
	// ram:PassRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the scaling group to which the scaling configuration belongs.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmwozpmmksakq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The restart policy of the container group. Valid values:
	//
	// 	- Never: The container group is never restarted.
	//
	// 	- Always: The container group is always restarted.
	//
	// 	- OnFailure: The container group is restarted upon failures.
	//
	// example:
	//
	// Always
	RestartPolicy *string `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-2zec39vg84usxdocme6a
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The name of the scaling configuration.
	//
	// example:
	//
	// scalingconfi****
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1frlu04fq4zv65b1bh
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The system information of the security context in which the elastic container instance is run.
	SecurityContextSysCtls []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls `json:"SecurityContextSysCtls,omitempty" xml:"SecurityContextSysCtls,omitempty" type:"Repeated"`
	// The ID of the security group with which the elastic container instance is associated. Elastic container instances that are associated with the same security group can access each other.
	//
	// example:
	//
	// sg-bp18kz60mefs****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Indicates whether user logs are collected. Default value: **false**.
	//
	// example:
	//
	// false
	SlsEnable *bool `json:"SlsEnable,omitempty" xml:"SlsEnable,omitempty"`
	// The maximum hourly price for the preemptible instance.
	//
	// This parameter is returned only when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy of the instance. Valid values:
	//
	// 	- NoSpot: The instance is created as a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a preemptible instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags of the elastic container instance. Tags are specified in the key-value format.
	Tags []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The buffer time during which a program handles operations before the program stops.
	//
	// example:
	//
	// 60
	TerminationGracePeriodSeconds *int32 `json:"TerminationGracePeriodSeconds,omitempty" xml:"TerminationGracePeriodSeconds,omitempty"`
	// The volumes.
	Volumes []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetAcrRegistryInfos() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos {
	return s.AcrRegistryInfos
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetActiveDeadlineSeconds() *int32 {
	return s.ActiveDeadlineSeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetAutoCreateEip() *bool {
	return s.AutoCreateEip
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetAutoMatchImageCache() *bool {
	return s.AutoMatchImageCache
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetComputeCategory() []*string {
	return s.ComputeCategory
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetContainerGroupName() *string {
	return s.ContainerGroupName
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetContainers() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	return s.Containers
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetCostOptimization() *bool {
	return s.CostOptimization
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetCpuOptionsCore() *int32 {
	return s.CpuOptionsCore
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetCpuOptionsThreadsPerCore() *int32 {
	return s.CpuOptionsThreadsPerCore
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDataCacheBucket() *string {
	return s.DataCacheBucket
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDataCacheBurstingEnabled() *bool {
	return s.DataCacheBurstingEnabled
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDataCachePL() *string {
	return s.DataCachePL
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDataCacheProvisionedIops() *int32 {
	return s.DataCacheProvisionedIops
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDescription() *string {
	return s.Description
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDnsConfigNameServers() []*string {
	return s.DnsConfigNameServers
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDnsConfigOptions() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions {
	return s.DnsConfigOptions
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDnsConfigSearches() []*string {
	return s.DnsConfigSearches
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetDnsPolicy() *string {
	return s.DnsPolicy
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetEgressBandwidth() *int64 {
	return s.EgressBandwidth
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetEipCommonBandwidthPackage() *string {
	return s.EipCommonBandwidthPackage
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetEipISP() *string {
	return s.EipISP
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetEipPublicIpAddressPoolId() *string {
	return s.EipPublicIpAddressPoolId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetEphemeralStorage() *int32 {
	return s.EphemeralStorage
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetHostAliases() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases {
	return s.HostAliases
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetHostName() *string {
	return s.HostName
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetImageRegistryCredentials() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials {
	return s.ImageRegistryCredentials
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetImageSnapshotId() *string {
	return s.ImageSnapshotId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetIngressBandwidth() *int64 {
	return s.IngressBandwidth
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetInitContainers() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	return s.InitContainers
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetNtpServers() []*string {
	return s.NtpServers
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetRestartPolicy() *string {
	return s.RestartPolicy
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetSecurityContextSysCtls() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls {
	return s.SecurityContextSysCtls
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetSlsEnable() *bool {
	return s.SlsEnable
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetTags() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags {
	return s.Tags
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetTerminationGracePeriodSeconds() *int32 {
	return s.TerminationGracePeriodSeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) GetVolumes() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	return s.Volumes
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetAcrRegistryInfos(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.AcrRegistryInfos = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetActiveDeadlineSeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ActiveDeadlineSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetAutoCreateEip(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.AutoCreateEip = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetAutoMatchImageCache(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.AutoMatchImageCache = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetComputeCategory(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ComputeCategory = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetContainerGroupName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ContainerGroupName = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetContainers(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Containers = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetCostOptimization(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.CostOptimization = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetCpu(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetCpuOptionsCore(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.CpuOptionsCore = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetCpuOptionsThreadsPerCore(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.CpuOptionsThreadsPerCore = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetCreationTime(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.CreationTime = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDataCacheBucket(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DataCacheBucket = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDataCacheBurstingEnabled(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DataCacheBurstingEnabled = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDataCachePL(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DataCachePL = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDataCacheProvisionedIops(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DataCacheProvisionedIops = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDescription(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Description = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDnsConfigNameServers(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DnsConfigNameServers = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDnsConfigOptions(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DnsConfigOptions = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDnsConfigSearches(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DnsConfigSearches = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetDnsPolicy(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.DnsPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetEgressBandwidth(v int64) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.EgressBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetEipBandwidth(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetEipCommonBandwidthPackage(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.EipCommonBandwidthPackage = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetEipISP(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.EipISP = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetEipPublicIpAddressPoolId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.EipPublicIpAddressPoolId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetEphemeralStorage(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.EphemeralStorage = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetHostAliases(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.HostAliases = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetHostName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.HostName = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetImageRegistryCredentials(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ImageRegistryCredentials = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetImageSnapshotId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ImageSnapshotId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetIngressBandwidth(v int64) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.IngressBandwidth = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetInitContainers(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.InitContainers = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetInstanceFamilyLevel(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetInstanceTypes(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.InstanceTypes = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetIpv6AddressCount(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetLifecycleState(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.LifecycleState = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetLoadBalancerWeight(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetMemory(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetNtpServers(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.NtpServers = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetRamRoleName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.RamRoleName = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetRegionId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetResourceGroupId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetRestartPolicy(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.RestartPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetScalingConfigurationId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetScalingConfigurationName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetScalingGroupId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetSecurityContextSysCtls(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.SecurityContextSysCtls = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetSecurityGroupId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetSlsEnable(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.SlsEnable = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetSpotPriceLimit(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetSpotStrategy(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetTags(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Tags = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetTerminationGracePeriodSeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.TerminationGracePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) SetVolumes(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration {
	s.Volumes = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfiguration) Validate() error {
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

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos struct {
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) SetDomains(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos {
	s.Domains = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) SetInstanceId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) SetInstanceName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos {
	s.InstanceName = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) SetRegionId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationAcrRegistryInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers struct {
	// The arguments that are passed to the container startup commands.
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// The container startup commands.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The number of vCPUs that are allocated to the elastic container instance.
	//
	// example:
	//
	// 2
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The environment variables.
	EnvironmentVars []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars `json:"EnvironmentVars,omitempty" xml:"EnvironmentVars,omitempty" type:"Repeated"`
	// The number of GPUs.
	//
	// example:
	//
	// 2
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The container image.
	//
	// example:
	//
	// registry-vpc.aliyuncs.com/eci_open/alpine:3.5
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The commands that are run by using a CLI for configuring the postStart callback function within the container.
	LifecyclePostStartHandlerExecs []*string `json:"LifecyclePostStartHandlerExecs,omitempty" xml:"LifecyclePostStartHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to the HTTP GET requests for configuring the postStart callback function are sent.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerHttpGetHost *string `json:"LifecyclePostStartHandlerHttpGetHost,omitempty" xml:"LifecyclePostStartHandlerHttpGetHost,omitempty"`
	// The path to the HTTP GET requests for configuring the postStart callback function are sent.
	//
	// example:
	//
	// /healthyz
	LifecyclePostStartHandlerHttpGetPath *string `json:"LifecyclePostStartHandlerHttpGetPath,omitempty" xml:"LifecyclePostStartHandlerHttpGetPath,omitempty"`
	// The port over which the HTTP GET requests for configuring the postStart callback function are sent.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerHttpGetPort *int32 `json:"LifecyclePostStartHandlerHttpGetPort,omitempty" xml:"LifecyclePostStartHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP Get requests that are used for configuring the postStart callback function.
	//
	// example:
	//
	// HTTP
	LifecyclePostStartHandlerHttpGetScheme *string `json:"LifecyclePostStartHandlerHttpGetScheme,omitempty" xml:"LifecyclePostStartHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP sockets that are used for configuring the postStart callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePostStartHandlerTcpSocketHost *string `json:"LifecyclePostStartHandlerTcpSocketHost,omitempty" xml:"LifecyclePostStartHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP sockets that are used for configuring the postStart callback function.
	//
	// example:
	//
	// 80
	LifecyclePostStartHandlerTcpSocketPort *int32 `json:"LifecyclePostStartHandlerTcpSocketPort,omitempty" xml:"LifecyclePostStartHandlerTcpSocketPort,omitempty"`
	// The commands that are run by using a CLI for configuring the preStop callback function within the container.
	LifecyclePreStopHandlerExecs []*string `json:"LifecyclePreStopHandlerExecs,omitempty" xml:"LifecyclePreStopHandlerExecs,omitempty" type:"Repeated"`
	// The IP address of the host to which the HTTP GET requests for configuring the preStop callback function are sent.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerHttpGetHost *string `json:"LifecyclePreStopHandlerHttpGetHost,omitempty" xml:"LifecyclePreStopHandlerHttpGetHost,omitempty"`
	// The path to which the HTTP GET requests for configuring the preStop callback function are sent.
	//
	// example:
	//
	// /healthyz
	LifecyclePreStopHandlerHttpGetPath *string `json:"LifecyclePreStopHandlerHttpGetPath,omitempty" xml:"LifecyclePreStopHandlerHttpGetPath,omitempty"`
	// The port over which the HTTP GET requests for configuring the preStop callback function are sent.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerHttpGetPort *int32 `json:"LifecyclePreStopHandlerHttpGetPort,omitempty" xml:"LifecyclePreStopHandlerHttpGetPort,omitempty"`
	// The protocol type of the HTTP Get requests that are used for configuring the preStop callback function.
	//
	// example:
	//
	// HTTP
	LifecyclePreStopHandlerHttpGetScheme *string `json:"LifecyclePreStopHandlerHttpGetScheme,omitempty" xml:"LifecyclePreStopHandlerHttpGetScheme,omitempty"`
	// The IP address of the host detected by the TCP sockets that are used for configuring the preStop callback function.
	//
	// example:
	//
	// 10.0.XX.XX
	LifecyclePreStopHandlerTcpSocketHost *string `json:"LifecyclePreStopHandlerTcpSocketHost,omitempty" xml:"LifecyclePreStopHandlerTcpSocketHost,omitempty"`
	// The port detected by the TCP sockets that are used for configuring the preStop callback function.
	//
	// example:
	//
	// 80
	LifecyclePreStopHandlerTcpSocketPort *int32 `json:"LifecyclePreStopHandlerTcpSocketPort,omitempty" xml:"LifecyclePreStopHandlerTcpSocketPort,omitempty"`
	// The commands that are run in the container when you use a CLI to perform liveness probes.
	LivenessProbeExecCommands []*string `json:"LivenessProbeExecCommands,omitempty" xml:"LivenessProbeExecCommands,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures before a successful liveness probe is considered failed.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	LivenessProbeFailureThreshold *int32 `json:"LivenessProbeFailureThreshold,omitempty" xml:"LivenessProbeFailureThreshold,omitempty"`
	// The path to which HTTP Get requests are sent when you use the HTTP requests to perform liveness probes.
	//
	// example:
	//
	// /usr/nginx/
	LivenessProbeHttpGetPath *string `json:"LivenessProbeHttpGetPath,omitempty" xml:"LivenessProbeHttpGetPath,omitempty"`
	// The port detected by HTTP Get requests when you use the HTTP requests to perform liveness probes.
	//
	// example:
	//
	// 8080
	LivenessProbeHttpGetPort *int32 `json:"LivenessProbeHttpGetPort,omitempty" xml:"LivenessProbeHttpGetPort,omitempty"`
	// The protocol type of HTTP GET requests when you use the HTTP requests to perform liveness probes. Valid values:
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
	// The memory size.
	//
	// example:
	//
	// 2.0
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The container name.
	//
	// example:
	//
	// nginx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The exposed ports and protocols.
	Ports []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The commands that are run in the container when you use a CLI to perform readiness probes.
	ReadinessProbeExecCommands []*string `json:"ReadinessProbeExecCommands,omitempty" xml:"ReadinessProbeExecCommands,omitempty" type:"Repeated"`
	// The minimum number of consecutive failures before a successful readiness probe is considered failed.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	ReadinessProbeFailureThreshold *int32 `json:"ReadinessProbeFailureThreshold,omitempty" xml:"ReadinessProbeFailureThreshold,omitempty"`
	// The path to which HTTP Get requests are sent when you use the HTTP requests to perform readiness probes.
	//
	// example:
	//
	// /usr/local
	ReadinessProbeHttpGetPath *string `json:"ReadinessProbeHttpGetPath,omitempty" xml:"ReadinessProbeHttpGetPath,omitempty"`
	// The path to which HTTP Get requests are sent when you use the HTTP Get requests to perform readiness probes.
	//
	// example:
	//
	// 80
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
	// 5
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
	// 80
	ReadinessProbeTcpSocketPort *int32 `json:"ReadinessProbeTcpSocketPort,omitempty" xml:"ReadinessProbeTcpSocketPort,omitempty"`
	// The timeout period of a readiness probe. Default value: 1. Minimum value: 1. Unit: seconds.
	//
	// example:
	//
	// 5
	ReadinessProbeTimeoutSeconds *int32 `json:"ReadinessProbeTimeoutSeconds,omitempty" xml:"ReadinessProbeTimeoutSeconds,omitempty"`
	// The permissions that are granted to the processes in the container. Valid values: NET_ADMIN and NET_RAW.
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
	// Specifies whether to enable the Interaction feature. Valid values:
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
	// true
	Tty *bool `json:"Tty,omitempty" xml:"Tty,omitempty"`
	// The volumes that are mounted to the container.
	VolumeMounts []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts `json:"VolumeMounts,omitempty" xml:"VolumeMounts,omitempty" type:"Repeated"`
	// The working directory in the container.
	//
	// example:
	//
	// /usr/local/nginx
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetArgs() []*string {
	return s.Args
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetCommands() []*string {
	return s.Commands
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetEnvironmentVars() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars {
	return s.EnvironmentVars
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetImage() *string {
	return s.Image
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerExecs() []*string {
	return s.LifecyclePostStartHandlerExecs
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerHttpGetHost() *string {
	return s.LifecyclePostStartHandlerHttpGetHost
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerHttpGetPath() *string {
	return s.LifecyclePostStartHandlerHttpGetPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerHttpGetPort() *int32 {
	return s.LifecyclePostStartHandlerHttpGetPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerHttpGetScheme() *string {
	return s.LifecyclePostStartHandlerHttpGetScheme
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerTcpSocketHost() *string {
	return s.LifecyclePostStartHandlerTcpSocketHost
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePostStartHandlerTcpSocketPort() *int32 {
	return s.LifecyclePostStartHandlerTcpSocketPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerExecs() []*string {
	return s.LifecyclePreStopHandlerExecs
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerHttpGetHost() *string {
	return s.LifecyclePreStopHandlerHttpGetHost
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerHttpGetPath() *string {
	return s.LifecyclePreStopHandlerHttpGetPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerHttpGetPort() *int32 {
	return s.LifecyclePreStopHandlerHttpGetPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerHttpGetScheme() *string {
	return s.LifecyclePreStopHandlerHttpGetScheme
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerTcpSocketHost() *string {
	return s.LifecyclePreStopHandlerTcpSocketHost
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLifecyclePreStopHandlerTcpSocketPort() *int32 {
	return s.LifecyclePreStopHandlerTcpSocketPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeExecCommands() []*string {
	return s.LivenessProbeExecCommands
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeFailureThreshold() *int32 {
	return s.LivenessProbeFailureThreshold
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeHttpGetPath() *string {
	return s.LivenessProbeHttpGetPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeHttpGetPort() *int32 {
	return s.LivenessProbeHttpGetPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeHttpGetScheme() *string {
	return s.LivenessProbeHttpGetScheme
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeInitialDelaySeconds() *int32 {
	return s.LivenessProbeInitialDelaySeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbePeriodSeconds() *int32 {
	return s.LivenessProbePeriodSeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeSuccessThreshold() *int32 {
	return s.LivenessProbeSuccessThreshold
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeTcpSocketPort() *int32 {
	return s.LivenessProbeTcpSocketPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetLivenessProbeTimeoutSeconds() *int32 {
	return s.LivenessProbeTimeoutSeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetPorts() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts {
	return s.Ports
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeExecCommands() []*string {
	return s.ReadinessProbeExecCommands
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeFailureThreshold() *int32 {
	return s.ReadinessProbeFailureThreshold
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeHttpGetPath() *string {
	return s.ReadinessProbeHttpGetPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeHttpGetPort() *int32 {
	return s.ReadinessProbeHttpGetPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeHttpGetScheme() *string {
	return s.ReadinessProbeHttpGetScheme
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeInitialDelaySeconds() *int32 {
	return s.ReadinessProbeInitialDelaySeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbePeriodSeconds() *int32 {
	return s.ReadinessProbePeriodSeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeSuccessThreshold() *int32 {
	return s.ReadinessProbeSuccessThreshold
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeTcpSocketPort() *int32 {
	return s.ReadinessProbeTcpSocketPort
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetReadinessProbeTimeoutSeconds() *int32 {
	return s.ReadinessProbeTimeoutSeconds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetSecurityContextCapabilityAdds() []*string {
	return s.SecurityContextCapabilityAdds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetSecurityContextReadOnlyRootFilesystem() *bool {
	return s.SecurityContextReadOnlyRootFilesystem
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetSecurityContextRunAsUser() *int64 {
	return s.SecurityContextRunAsUser
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetStdin() *bool {
	return s.Stdin
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetStdinOnce() *bool {
	return s.StdinOnce
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetTty() *bool {
	return s.Tty
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetVolumeMounts() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts {
	return s.VolumeMounts
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetArgs(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Args = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetCommands(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Commands = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetCpu(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetEnvironmentVars(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.EnvironmentVars = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetGpu(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetImage(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Image = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetImagePullPolicy(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerExecs(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerExecs = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerHttpGetHost(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerHttpGetHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerHttpGetPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerHttpGetPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerHttpGetScheme(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerTcpSocketHost(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerTcpSocketHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePostStartHandlerTcpSocketPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePostStartHandlerTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerExecs(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerExecs = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerHttpGetHost(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerHttpGetHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerHttpGetPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerHttpGetPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerHttpGetScheme(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerTcpSocketHost(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerTcpSocketHost = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLifecyclePreStopHandlerTcpSocketPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LifecyclePreStopHandlerTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeExecCommands(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeExecCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeFailureThreshold(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeFailureThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeHttpGetPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeHttpGetPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeHttpGetScheme(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeInitialDelaySeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeInitialDelaySeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbePeriodSeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeSuccessThreshold(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeSuccessThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeTcpSocketPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetLivenessProbeTimeoutSeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.LivenessProbeTimeoutSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetMemory(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetPorts(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Ports = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeExecCommands(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeExecCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeFailureThreshold(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeFailureThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeHttpGetPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeHttpGetPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeHttpGetPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeHttpGetPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeHttpGetScheme(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeHttpGetScheme = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeInitialDelaySeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeInitialDelaySeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbePeriodSeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbePeriodSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeSuccessThreshold(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeSuccessThreshold = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeTcpSocketPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeTcpSocketPort = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetReadinessProbeTimeoutSeconds(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.ReadinessProbeTimeoutSeconds = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetSecurityContextCapabilityAdds(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.SecurityContextCapabilityAdds = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetSecurityContextReadOnlyRootFilesystem(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.SecurityContextReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetSecurityContextRunAsUser(v int64) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.SecurityContextRunAsUser = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetStdin(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Stdin = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetStdinOnce(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.StdinOnce = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetTty(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.Tty = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetVolumeMounts(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.VolumeMounts = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) SetWorkingDir(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers {
	s.WorkingDir = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainers) Validate() error {
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

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars struct {
	// >  This parameter is not available for use.
	//
	// example:
	//
	// fieldPath
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) SetKey(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) SetValue(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts struct {
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8083
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) SetPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts {
	s.Port = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) SetProtocol(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts struct {
	// The directory in which the container mounts the volume.
	//
	// >  Data in this directory is overwritten by the data on the volume. Proceed with caution if you specify this parameter.
	//
	// example:
	//
	// /pod/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation enables volumes mounted on one container to be shared among other containers within the same pod or across distinct pods residing on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed either on the volume itself or its subdirectories do not propagate to the already established volume mount.
	//
	// 	- HostToCotainer: Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount. In addition, any volume mounts executed on the container not only propagate back to the underlying host but also to all containers across every pod that uses the same volume.
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
	// default-volume1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the volume is read-only.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The subdirectory of the volume.
	//
	// example:
	//
	// data2/
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) SetMountPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) SetMountPropagation(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) SetReadOnly(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) SetSubPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationContainersVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions struct {
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) SetValue(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationDnsConfigOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases struct {
	// The added hostnames.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The added IP address.
	//
	// example:
	//
	// 192.0.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) GetHostnames() []*string {
	return s.Hostnames
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) GetIp() *string {
	return s.Ip
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) SetHostnames(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases {
	s.Hostnames = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) SetIp(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases {
	s.Ip = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationHostAliases) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials struct {
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) GetPassword() *string {
	return s.Password
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) GetServer() *string {
	return s.Server
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) GetUserName() *string {
	return s.UserName
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) SetPassword(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials {
	s.Password = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) SetServer(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials {
	s.Server = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) SetUserName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials {
	s.UserName = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationImageRegistryCredentials) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers struct {
	// The number of vCPUs that are allocated to the init container.
	//
	// example:
	//
	// 0.5
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs that are allocated to the init container.
	//
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The image of the init container.
	//
	// example:
	//
	// registry-vpc.cn-hongkong.aliyuncs.com/eci_open/nginx:alpine
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The image pulling policy.
	//
	// example:
	//
	// Always
	ImagePullPolicy *string `json:"ImagePullPolicy,omitempty" xml:"ImagePullPolicy,omitempty"`
	// The arguments that are passed to the startup commands of the init container.
	InitContainerArgs []*string `json:"InitContainerArgs,omitempty" xml:"InitContainerArgs,omitempty" type:"Repeated"`
	// The commands that are used to start the init container.
	InitContainerCommands []*string `json:"InitContainerCommands,omitempty" xml:"InitContainerCommands,omitempty" type:"Repeated"`
	// The environment variables of the init container.
	InitContainerEnvironmentVars []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars `json:"InitContainerEnvironmentVars,omitempty" xml:"InitContainerEnvironmentVars,omitempty" type:"Repeated"`
	// The ports of the init container.
	InitContainerPorts []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts `json:"InitContainerPorts,omitempty" xml:"InitContainerPorts,omitempty" type:"Repeated"`
	// The volume mounts of the init container.
	InitContainerVolumeMounts []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts `json:"InitContainerVolumeMounts,omitempty" xml:"InitContainerVolumeMounts,omitempty" type:"Repeated"`
	// The memory size of the init container.
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
	SecurityContextCapabilityAdds []*string `json:"SecurityContextCapabilityAdds,omitempty" xml:"SecurityContextCapabilityAdds,omitempty" type:"Repeated"`
	// Indicates whether the root file system on which the init container runs is read-only. Valid value: true.
	SecurityContextReadOnlyRootFilesystem *bool `json:"SecurityContextReadOnlyRootFilesystem,omitempty" xml:"SecurityContextReadOnlyRootFilesystem,omitempty"`
	// The ID of the user that runs the init container.
	//
	// example:
	//
	// 1000
	SecurityContextRunAsUser *string `json:"SecurityContextRunAsUser,omitempty" xml:"SecurityContextRunAsUser,omitempty"`
	// The working directory of the init container.
	//
	// example:
	//
	// /www
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetCpu() *float32 {
	return s.Cpu
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetGpu() *int32 {
	return s.Gpu
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetImage() *string {
	return s.Image
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetImagePullPolicy() *string {
	return s.ImagePullPolicy
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetInitContainerArgs() []*string {
	return s.InitContainerArgs
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetInitContainerCommands() []*string {
	return s.InitContainerCommands
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetInitContainerEnvironmentVars() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars {
	return s.InitContainerEnvironmentVars
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetInitContainerPorts() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts {
	return s.InitContainerPorts
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetInitContainerVolumeMounts() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts {
	return s.InitContainerVolumeMounts
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetSecurityContextCapabilityAdds() []*string {
	return s.SecurityContextCapabilityAdds
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetSecurityContextReadOnlyRootFilesystem() *bool {
	return s.SecurityContextReadOnlyRootFilesystem
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetSecurityContextRunAsUser() *string {
	return s.SecurityContextRunAsUser
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetCpu(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.Cpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetGpu(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.Gpu = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetImage(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.Image = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetImagePullPolicy(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.ImagePullPolicy = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetInitContainerArgs(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.InitContainerArgs = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetInitContainerCommands(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.InitContainerCommands = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetInitContainerEnvironmentVars(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.InitContainerEnvironmentVars = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetInitContainerPorts(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.InitContainerPorts = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetInitContainerVolumeMounts(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.InitContainerVolumeMounts = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetMemory(v float32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.Memory = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetSecurityContextCapabilityAdds(v []*string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.SecurityContextCapabilityAdds = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetSecurityContextReadOnlyRootFilesystem(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.SecurityContextReadOnlyRootFilesystem = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetSecurityContextRunAsUser(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.SecurityContextRunAsUser = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) SetWorkingDir(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers {
	s.WorkingDir = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainers) Validate() error {
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

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars struct {
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) GetFieldRefFieldPath() *string {
	return s.FieldRefFieldPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) GetKey() *string {
	return s.Key
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) SetFieldRefFieldPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars {
	s.FieldRefFieldPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) SetKey(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) SetValue(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerEnvironmentVars) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts struct {
	// The port number. Valid values: 1 to 65535.
	//
	// example:
	//
	// 1024
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// UDP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) SetPort(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts {
	s.Port = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) SetProtocol(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts struct {
	// The directory to which the init container mounts the volume.
	//
	// >  Data in this directory is overwritten by the data on the volume. Proceed with caution if you specify this parameter.
	//
	// example:
	//
	// /usr/share/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount propagation setting of the volume. Mount propagation enables volumes mounted on one container to be shared among other containers within the same pod or across distinct pods residing on the same node. Valid values:
	//
	// 	- None: Subsequent mounts executed either on the volume itself or its subdirectories do not propagate to the already established volume mount.
	//
	// 	- HostToCotainer: Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount.
	//
	// 	- Bidirectional: This value is similar to HostToCotainer. Subsequent mounts executed either on the volume itself or its subdirectories propagate to the already established volume mount. In addition, any volume mounts executed on the container not only propagate back to the underlying host but also to all containers across every pod that uses the same volume.
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
	// The subdirectory of the volume.
	//
	// example:
	//
	// /usr/sub/
	SubPath *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) GetMountPath() *string {
	return s.MountPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) GetMountPropagation() *string {
	return s.MountPropagation
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) GetSubPath() *string {
	return s.SubPath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) SetMountPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts {
	s.MountPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) SetMountPropagation(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts {
	s.MountPropagation = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) SetReadOnly(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts {
	s.ReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) SetSubPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts {
	s.SubPath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationInitContainersInitContainerVolumeMounts) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls struct {
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) SetValue(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationSecurityContextSysCtls) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags struct {
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) GetKey() *string {
	return s.Key
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) GetValue() *string {
	return s.Value
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) SetKey(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags {
	s.Key = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) SetValue(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags {
	s.Value = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationTags) Validate() error {
	return dara.Validate(s)
}

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes struct {
	// The paths to the configuration files.
	ConfigFileVolumeConfigFileToPaths []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths `json:"ConfigFileVolumeConfigFileToPaths,omitempty" xml:"ConfigFileVolumeConfigFileToPaths,omitempty" type:"Repeated"`
	// The default permissions on the ConfigFile volume.
	//
	// example:
	//
	// 0644
	ConfigFileVolumeDefaultMode *int32 `json:"ConfigFileVolumeDefaultMode,omitempty" xml:"ConfigFileVolumeDefaultMode,omitempty"`
	// The ID of the disk volume.
	//
	// example:
	//
	// d-xx
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
	// The storage medium of the emptyDir volume. If you do not specify a storage medium for the emptyDir volume, the volume stores data in the file system of the node by default. We recommend that you set this parameter to memory. In this case, the emptyDir volume stores data in the specified memory.
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
	// The FlexVolume options.
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

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetConfigFileVolumeConfigFileToPaths() []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths {
	return s.ConfigFileVolumeConfigFileToPaths
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetConfigFileVolumeDefaultMode() *int32 {
	return s.ConfigFileVolumeDefaultMode
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetDiskVolumeDiskId() *string {
	return s.DiskVolumeDiskId
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetDiskVolumeDiskSize() *int32 {
	return s.DiskVolumeDiskSize
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetDiskVolumeFsType() *string {
	return s.DiskVolumeFsType
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetEmptyDirVolumeMedium() *string {
	return s.EmptyDirVolumeMedium
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetEmptyDirVolumeSizeLimit() *string {
	return s.EmptyDirVolumeSizeLimit
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetFlexVolumeDriver() *string {
	return s.FlexVolumeDriver
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetFlexVolumeFsType() *string {
	return s.FlexVolumeFsType
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetFlexVolumeOptions() *string {
	return s.FlexVolumeOptions
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetHostPathVolumePath() *string {
	return s.HostPathVolumePath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetHostPathVolumeType() *string {
	return s.HostPathVolumeType
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetNFSVolumePath() *string {
	return s.NFSVolumePath
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetNFSVolumeReadOnly() *bool {
	return s.NFSVolumeReadOnly
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetNFSVolumeServer() *string {
	return s.NFSVolumeServer
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetName() *string {
	return s.Name
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) GetType() *string {
	return s.Type
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetConfigFileVolumeConfigFileToPaths(v []*DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.ConfigFileVolumeConfigFileToPaths = v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetConfigFileVolumeDefaultMode(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.ConfigFileVolumeDefaultMode = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetDiskVolumeDiskId(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.DiskVolumeDiskId = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetDiskVolumeDiskSize(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.DiskVolumeDiskSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetDiskVolumeFsType(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.DiskVolumeFsType = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetEmptyDirVolumeMedium(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.EmptyDirVolumeMedium = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetEmptyDirVolumeSizeLimit(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.EmptyDirVolumeSizeLimit = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetFlexVolumeDriver(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.FlexVolumeDriver = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetFlexVolumeFsType(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.FlexVolumeFsType = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetFlexVolumeOptions(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.FlexVolumeOptions = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetHostPathVolumePath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.HostPathVolumePath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetHostPathVolumeType(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.HostPathVolumeType = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetNFSVolumePath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.NFSVolumePath = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetNFSVolumeReadOnly(v bool) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.NFSVolumeReadOnly = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetNFSVolumeServer(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.NFSVolumeServer = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetName(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.Name = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) SetType(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes {
	s.Type = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumes) Validate() error {
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

type DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths struct {
	// The content of the configuration file.
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
	// The path to the configuration file.
	//
	// example:
	//
	// /usr/bin/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) GetContent() *string {
	return s.Content
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) GetMode() *int32 {
	return s.Mode
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) GetPath() *string {
	return s.Path
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) SetContent(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths {
	s.Content = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) SetMode(v int32) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths {
	s.Mode = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) SetPath(v string) *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths {
	s.Path = &v
	return s
}

func (s *DescribeEciScalingConfigurationDetailResponseBodyScalingConfigurationVolumesConfigFileVolumeConfigFileToPaths) Validate() error {
	return dara.Validate(s)
}
