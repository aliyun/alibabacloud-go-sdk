// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeScalingConfigurationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingConfigurationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeScalingConfigurationsResponseBody
	GetRequestId() *string
	SetScalingConfigurations(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurations) *DescribeScalingConfigurationsResponseBody
	GetScalingConfigurations() []*DescribeScalingConfigurationsResponseBodyScalingConfigurations
	SetTotalCount(v int32) *DescribeScalingConfigurationsResponseBody
	GetTotalCount() *int32
}

type DescribeScalingConfigurationsResponseBody struct {
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
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling configurations.
	ScalingConfigurations []*DescribeScalingConfigurationsResponseBodyScalingConfigurations `json:"ScalingConfigurations,omitempty" xml:"ScalingConfigurations,omitempty" type:"Repeated"`
	// The total number of scaling configurations.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingConfigurationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingConfigurationsResponseBody) GetScalingConfigurations() []*DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	return s.ScalingConfigurations
}

func (s *DescribeScalingConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageNumber(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageSize(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetRequestId(v string) *DescribeScalingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetScalingConfigurations(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurations) *DescribeScalingConfigurationsResponseBody {
	s.ScalingConfigurations = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetTotalCount(v int32) *DescribeScalingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) Validate() error {
	if s.ScalingConfigurations != nil {
		for _, item := range s.ScalingConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurations struct {
	// Indicates whether the ECS instance on a dedicated host is associated with the dedicated host. Valid values:
	//
	// 	- default: The instance is not associated with the dedicated host. If you restart an instance that was stopped in Economical Mode and the original dedicated host of the instance has insufficient resources, the instance is automatically deployed to another dedicated host in the automatic deployment resource pool.
	//
	// 	- host: The instance is associated with the dedicated host. If you restart an instance that was stopped in Economical Mode, the instance remains on the original dedicated host. If the available resources of the original dedicated host are insufficient, the instance cannot be restarted.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of vCPUs.
	//
	// You can specify CPU and Memory to define the range of instance types. For example, if you set CPU to 2 and Memory to 16, the instance types that have 2 vCPUs and 16 GiB are returned. If you specify CPU and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones and preferentially creates instances by using the lowest-priced instance type.
	//
	// >  You can specify CPU and Memory to define instance types only when you set Scaling Policy to Cost Optimization and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time at which the scaling configuration was created.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The performance mode of the burstable instances. Valid values:
	//
	// 	- Standard: the standard mode. For more information, see the "Standard mode" section in the [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html) topic.
	//
	// 	- Unlimited: the unlimited mode. For more information, see the "Unlimited mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The priority of the custom ECS instance type + vSwitch combination.
	//
	// >  This parameter takes effect only when Scaling Policy of the scaling group is set to Priority Policy.
	//
	// If Auto Scaling cannot create ECS instances by using the custom ECS instance type + vSwitch combination of the highest priority, Auto Scaling creates ECS instances by using the custom ECS instance type + vSwitch combination of the next highest priority.
	//
	// >  If you specify the priorities of only a portion of custom ECS instance type + vSwitch combinations, Auto Scaling preferentially creates ECS instances by using the custom combinations that have specified priorities. If the custom combinations that have specified priorities do not provide sufficient resources, Auto Scaling creates ECS instances by using the custom combinations that do not have specified priorities based on the specified orders of vSwitches and instance types.
	//
	// 	- Example: the specified order of vSwitches for your scaling group is vsw1 and vsw2 and the specified order of instance types in your scaling configuration is type1 and type 2. In addition, you use CustomPriorities to specify ["vsw2+type2", "vsw1+type2"]. In this example, the vsw2+type2 combination has the highest priority and the vsw2+type1 combination has the lowest priority. The vsw1+type2 combination has a higher priority than the vsw1+type1 combination.
	CustomPriorities []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The data disks.
	DataDisks []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-zm04u8r3lohsq****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host on which the ECS instance is created. Preemptible instances are not supported by dedicated hosts. Therefore, if you specify DedicatedHostId, SpotStrategy and SpotPriceLimit are ignored.
	//
	// You can call the DescribeDedicatedHosts operation to query the IDs of dedicated hosts.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// Indicates whether Release Protection is enabled for the ECS instances. You can specify this parameter to determine whether the ECS instances can be deleted by using the ECS console or calling the DeleteInstance operation. Valid values:
	//
	// 	- true: Release Protection is enabled for the ECS instances. You cannot delete the ECS instances by using the ECS console or calling the DeleteInstance operation.
	//
	// 	- false: Release Protection is disabled for the ECS instances. You can delete the ECS instances by using the ECS console or calling the DeleteInstance operation.
	//
	// >  You can enable Release Protection for only pay-as-you-go instances to prevent unexpected instance deletion during scale-in events. The Release Protection feature does not affect normal scaling activities. In other words, an instance that meets the criteria of scale-in policies may be removed from a scaling group during a scale-in event even if you enabled Release Protection for the instance.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set to which the Elastic Compute Service (ECS) instances belong.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname series of the ECS instances.
	//
	// example:
	//
	// LocalHost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the High Performance Computing (HPC) cluster to which the ECS instances belong.
	//
	// example:
	//
	// hpc-clus****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Indicates whether the access channel is enabled for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Indicates whether the security hardening mode (IMDSv2) is forcefully used to access instance metadata. Valid values:
	//
	// 	- optional: The security hardening mode IMDSv2 is not forcibly used.
	//
	// 	- required: The security hardening mode (IMDSv2) is forcibly used. After you set this parameter to required, you cannot access instance metadata in normal mode.
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can specify this parameter to obtain the latest available images in the current image family for instance creation. If you specify ImageId, you cannot specify `ImageFamily`.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file that provides the image resource for Auto Scaling to create ECS instances.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image file.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Indicates whether the ecs-user username can be used to log on to an ECS instance created from the scaling configuration. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ImageOptionsLoginAsNonRoot *bool `json:"ImageOptionsLoginAsNonRoot,omitempty" xml:"ImageOptionsLoginAsNonRoot,omitempty"`
	// The image source. Valid values:
	//
	// 	- system: a public image provided by Alibaba Cloud
	//
	// 	- self: a custom image that you created
	//
	// 	- others: a shared image from another Alibaba Cloud account or a community image published by another Alibaba Cloud account
	//
	// 	- marketplace: an Alibaba Cloud Marketplace image
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The description of the ECS instances.
	//
	// example:
	//
	// FinanceDept
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The generation of the ECS instances.
	//
	// example:
	//
	// ecs-3
	InstanceGeneration *string `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	// The naming series of the ECS instances.
	//
	// example:
	//
	// instance****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The intelligent configuration settings, which determine the available instance types.
	InstancePatternInfos []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance types of the ECS instances.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ECS instance types.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth. You are charged for the bandwidth that you specified by using InternetMaxBandwidthOut.
	//
	// 	- PayByTraffic: pay-by-traffic. You are charged for the actual traffic that you used. InternetMaxBandwidthOut specifies only the maximum available bandwidth.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Indicates whether the ECS instances are I/O optimized. Valid values:
	//
	// 	- none: The ECS instances are not I/O optimized.
	//
	// 	- optimized: The ECS instances are I/O optimized.
	//
	// example:
	//
	// none
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses that are allocated to the elastic network interface (ENI).
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair that is used to log on to an ECS instance created from the scaling configuration.
	//
	// example:
	//
	// keypair****
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The status of the scaling configuration in the scaling group. Valid values:
	//
	// 	- Active: The scaling configuration is active in the scaling group. Auto Scaling uses the scaling configuration that is in the Active state to create ECS instances during scale-out events.
	//
	// 	- Inactive: The scaling configuration is inactive in the scaling group. Scaling configurations that are in the Inactive state are still contained in the scaling group, but Auto Scaling does not use the inactive scaling configurations to create ECS instances during scale-out events.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The weight of an ECS instance as a backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify CPU and Memory to define the range of instance types. For example, if you set CPU to 2 and Memory to 16, the instance types that have 2 vCPUs and 16 GiB are returned. If you specify CPU and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones and preferentially creates instances by using the lowest-priced instance type.
	//
	// >  You can specify CPU and Memory to define instance types only when you set Scaling Policy to Cost Optimization and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ENIs.
	NetworkInterfaces []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	// Indicates whether the password preconfigured in the image is used.
	//
	// example:
	//
	// true
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// Indicates whether a password is configured for the instance.
	//
	// example:
	//
	// false
	PasswordSetted                   *bool   `json:"PasswordSetted,omitempty" xml:"PasswordSetted,omitempty"`
	PrivatePoolOptions_id            *string `json:"PrivatePoolOptions.Id,omitempty" xml:"PrivatePoolOptions.Id,omitempty"`
	PrivatePoolOptions_matchCriteria *string `json:"PrivatePoolOptions.MatchCriteria,omitempty" xml:"PrivatePoolOptions.MatchCriteria,omitempty"`
	// The name of the Resource Access Management (RAM) role assumed by the ECS instances. This name is provided and maintained by RAM. You can call the ListRoles operation to query the available RAM roles.
	//
	// example:
	//
	// ramrole****
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the ECS instances belong.
	//
	// example:
	//
	// rg-aekzn2ou7xo****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource pools used for instance creation, which can be the public pool or a private pool associated with any active elasticity assurance or capacity reservation.
	//
	// 	- This parameter takes effect only when you create pay-as-you-go instances.
	ResourcePoolOptions *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
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
	// scalingconfigura****
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The ID of the scaling group to which the scaling configuration belongs.
	//
	// example:
	//
	// asg-bp17pelvl720x3v7****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// >  This parameter is in invitational preview and is not available for use.
	SchedulerOptions *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	// Indicates whether Security Hardening is enabled. Valid values:
	//
	// 	- Active: Security Hardening is enabled. This value is applicable to only public images.
	//
	// 	- Deactive: Security Hardening is disabled. This value is applicable to all images.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the ECS instances belong. ECS instances that belong to the same security group can communicate with each other.
	//
	// example:
	//
	// sg-bp18kz60mefs****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups to which the ECS instances belong. ECS instances that belong to the same security group can communicate with each other.
	SecurityGroupIds []*string                                                                      `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the preemptible instances. Unit: hours.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption event of the preemptible instances.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The preemptible instances.
	SpotPriceLimits []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy that is applied to pay-as-you-go instances. Valid values:
	//
	// 	- NoSpot: The instances are created as regular pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are created as preemptible instances that have a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are preemptible instances for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. The value is an integer that is greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The ID of the automatic snapshot policy that is applied to the system disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	SystemDiskAutoSnapshotPolicyId *string `json:"SystemDiskAutoSnapshotPolicyId,omitempty" xml:"SystemDiskAutoSnapshotPolicyId,omitempty"`
	// Indicates whether the Performance Burst feature is enabled for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is available only when you set SystemDisk.Category to cloud_auto.
	//
	// example:
	//
	// false
	SystemDiskBurstingEnabled *bool `json:"SystemDiskBurstingEnabled,omitempty" xml:"SystemDiskBurstingEnabled,omitempty"`
	// The categories of the system disks. The values are sorted based on their priorities. The first value has the highest priority. If Auto Scaling cannot create instances by using the disk category of the highest priority, Auto Scaling creates instances by using the disk category of the next highest priority. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The category of the system disk. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- ephemeral_ssd: local SSD
	//
	// 	- cloud_essd: enterprise SSD (ESSD)
	//
	// 	- cloud_auto: ESSD AutoPL
	//
	// example:
	//
	// cloud
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The description of the system disk.
	//
	// example:
	//
	// Test system disk.
	SystemDiskDescription *string `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	// The encryption algorithm that is applied to the system disk. Valid values:
	//
	// 	- AES-256
	//
	// 	- SM4-128
	//
	// example:
	//
	// AES-256
	SystemDiskEncryptAlgorithm *string `json:"SystemDiskEncryptAlgorithm,omitempty" xml:"SystemDiskEncryptAlgorithm,omitempty"`
	// Indicates whether the system disk is encrypted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"SystemDiskEncrypted,omitempty" xml:"SystemDiskEncrypted,omitempty"`
	// The ID of the KMS key that is applied to the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKMSKeyId *string `json:"SystemDiskKMSKeyId,omitempty" xml:"SystemDiskKMSKeyId,omitempty"`
	// The name of the system disk.
	//
	// example:
	//
	// cloud_ssd_Test
	SystemDiskName *string `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	// The performance level (PL) of the system disk that is an ESSD.
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The provisioned IOPS of the system disk.
	//
	// >  IOPS measures the number of read and write operations that an EBS device can process per second.
	//
	// example:
	//
	// 100
	SystemDiskProvisionedIops *int64 `json:"SystemDiskProvisionedIops,omitempty" xml:"SystemDiskProvisionedIops,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 100
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags.
	Tags []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether the ECS instance is created on a dedicated host. Valid values:
	//
	// 	- default: The ECS instance is created on a non-dedicated host.
	//
	// 	- host: The ECS instance is created on a dedicated host. If you do not specify DedicatedHostId, the system selects a dedicated host for the ECS instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// The user data of the ECS instances.
	//
	// example:
	//
	// echo hello ecs!
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The weights of the instance types. The value of this parameter indicates the capacity of an instance of the specified instance type in the scaling group. A higher weight indicates that a smaller number of instances of the instance type are required to meet the expected capacity requirement.
	WeightedCapacities []*int32 `json:"WeightedCapacities,omitempty" xml:"WeightedCapacities,omitempty" type:"Repeated"`
	// The ID of the zone in which the ECS instances are created. You can call the DescribeZones operation to query the zone IDs.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetAffinity() *string {
	return s.Affinity
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetCustomPriorities() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities {
	return s.CustomPriorities
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetDataDisks() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	return s.DataDisks
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetHostName() *string {
	return s.HostName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetImageOptionsLoginAsNonRoot() *bool {
	return s.ImageOptionsLoginAsNonRoot
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceGeneration() *string {
	return s.InstanceGeneration
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstancePatternInfos() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	return s.InstancePatternInfos
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetNetworkInterfaces() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetPasswordSetted() *bool {
	return s.PasswordSetted
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetPrivatePoolOptions_id() *string {
	return s.PrivatePoolOptions_id
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetPrivatePoolOptions_matchCriteria() *string {
	return s.PrivatePoolOptions_matchCriteria
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetResourcePoolOptions() *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSchedulerOptions() *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions {
	return s.SchedulerOptions
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSecurityOptions() *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions {
	return s.SecurityOptions
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSpotPriceLimits() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits {
	return s.SpotPriceLimits
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskAutoSnapshotPolicyId() *string {
	return s.SystemDiskAutoSnapshotPolicyId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskBurstingEnabled() *bool {
	return s.SystemDiskBurstingEnabled
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskDescription() *string {
	return s.SystemDiskDescription
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskEncryptAlgorithm() *string {
	return s.SystemDiskEncryptAlgorithm
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskEncrypted() *bool {
	return s.SystemDiskEncrypted
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskKMSKeyId() *string {
	return s.SystemDiskKMSKeyId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskName() *string {
	return s.SystemDiskName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskProvisionedIops() *int64 {
	return s.SystemDiskProvisionedIops
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetTags() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags {
	return s.Tags
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetTenancy() *string {
	return s.Tenancy
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetUserData() *string {
	return s.UserData
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetWeightedCapacities() []*int32 {
	return s.WeightedCapacities
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetAffinity(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Affinity = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCpu(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Cpu = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCreationTime(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCreditSpecification(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.CreditSpecification = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCustomPriorities(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.CustomPriorities = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDataDisks(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DataDisks = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDedicatedHostClusterId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDedicatedHostId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDeletionProtection(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetDeploymentSetId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.DeploymentSetId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetHostName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.HostName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetHpcClusterId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.HpcClusterId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetHttpEndpoint(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.HttpEndpoint = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetHttpTokens(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.HttpTokens = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageFamily(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageFamily = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageOptionsLoginAsNonRoot(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageOptionsLoginAsNonRoot = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetImageOwnerAlias(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceGeneration(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstancePatternInfos(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstancePatternInfos = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceTypes(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceTypes = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInternetChargeType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInternetMaxBandwidthIn(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInternetMaxBandwidthOut(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetIoOptimized(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.IoOptimized = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetIpv6AddressCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetKeyPairName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.KeyPairName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetLifecycleState(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetLoadBalancerWeight(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetMemory(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Memory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetNetworkInterfaces(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.NetworkInterfaces = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetPasswordInherit(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.PasswordInherit = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetPasswordSetted(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.PasswordSetted = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetPrivatePoolOptions_id(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.PrivatePoolOptions_id = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetPrivatePoolOptions_matchCriteria(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.PrivatePoolOptions_matchCriteria = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetRamRoleName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.RamRoleName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetResourceGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetResourcePoolOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ResourcePoolOptions = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfigurationName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSchedulerOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SchedulerOptions = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityEnhancementStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityGroupIds(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecurityOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecurityOptions = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotDuration(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotDuration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotInterruptionBehavior(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotPriceLimits(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotPriceLimits = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSpotStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetStorageSetId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.StorageSetId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetStorageSetPartitionNumber(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskAutoSnapshotPolicyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskAutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskBurstingEnabled(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskBurstingEnabled = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskCategories(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskCategories = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskDescription = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskEncryptAlgorithm(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskEncryptAlgorithm = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskEncrypted(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskEncrypted = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskKMSKeyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskKMSKeyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskPerformanceLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskProvisionedIops(v int64) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskProvisionedIops = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSystemDiskSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetTags(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Tags = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetTenancy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.Tenancy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetUserData(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.UserData = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetWeightedCapacities(v []*int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.WeightedCapacities = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetZoneId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ZoneId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) Validate() error {
	if s.CustomPriorities != nil {
		for _, item := range s.CustomPriorities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstancePatternInfos != nil {
		for _, item := range s.InstancePatternInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkInterfaces != nil {
		for _, item := range s.NetworkInterfaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourcePoolOptions != nil {
		if err := s.ResourcePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SpotPriceLimits != nil {
		for _, item := range s.SpotPriceLimits {
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
	return nil
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities struct {
	// The ECS instance type.
	//
	// example:
	//
	// ecs.c6a.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp14zolna43z266bq****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) SetVswitchId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities {
	s.VswitchId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks struct {
	// The ID of the automatic snapshot policy that is applied to the data disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Indicates whether the Performance Burst feature is enabled for the data disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is available only when you set `DataDisk.Category` to `cloud_auto`.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The categories of the data disks. The values are sorted based on their priorities. The first value has the highest priority. If Auto Scaling cannot create instances by using the disk category of the highest priority, Auto Scaling creates instances by using the disk category of the next highest priority. Valid values:
	//
	// 	- cloud: basic disk. DeleteWithInstance of a basic disk created along with the ECS instance is set to true.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: ESSD.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data disk. Valid values:
	//
	// 	- cloud: basic disk. DeleteWithInstance of a basic disk created along with the ECS instance is set to true.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- ephemeral_ssd: local SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// 	- cloud_auto: ESSD AutoPL.
	//
	// example:
	//
	// cloud
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether the data disk is released when the instance to which the data disk is attached is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk.
	//
	// example:
	//
	// FinanceDept
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount target of the data disk.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Indicates whether the data disk is encrypted. Valid values:
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
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key that is applied to the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The PL of the data disk that is an ESSD.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS of the data disk.
	//
	// >  IOPS measures the number of read and write operations that an Elastic Block Storage (EBS) device can process per second.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the data disk. Unit: GB. Valid values:
	//
	// 	- 5 to 2000 if you set Category to cloud.
	//
	// 	- 20 to 32768 if you set Category to cloud_efficiency.
	//
	// 	- 20 to 32768 if you set Category to cloud_ssd.
	//
	// 	- 20 to 32768 if you set Category to cloud_essd.
	//
	// 	- 5 to 800 if you set Category to ephemeral_ssd.
	//
	// example:
	//
	// 200
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot based on which the data disk is created.
	//
	// example:
	//
	// s-23f2i****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetCategories() []*string {
	return s.Categories
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetDescription() *string {
	return s.Description
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetDevice() *string {
	return s.Device
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetAutoSnapshotPolicyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetBurstingEnabled(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetCategories(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Categories = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Category = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDeleteWithInstance(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDescription(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Description = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDevice(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Device = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetDiskName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.DiskName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetEncrypted(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Encrypted = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetKMSKeyId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetPerformanceLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetProvisionedIops(v int64) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.Size = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) SetSnapshotId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos struct {
	// The architectures of instance types. Valid values:
	//
	// 	- X86: x86.
	//
	// 	- Heterogeneous: heterogeneous computing, such as GPU-accelerated or FPGA-accelerated.
	//
	// 	- BareMetal: ECS Bare Metal Instance.
	//
	// 	- Arm: Arm.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Indicates whether burstable instance types are included. Valid values:
	//
	// 	- Exclude: Burstable instance types are not included.
	//
	// 	- Include: Burstable instance types are included.
	//
	// 	- Required: Only burstable instance types are included.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The number of vCPUs of the instance type.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architectures of the instance types. Valid values:
	//
	// >  You can specify 1 to 2 CPU architectures.
	//
	// 	- x86
	//
	// 	- Arm
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The instance types that are excluded. You can use wildcard characters, such as an asterisk (\\*), to exclude an instance type or an instance family. Examples:
	//
	// 	- ecs.c6.large: The ecs.c6.large instance type is excluded.
	//
	// 	- ecs.c6.\\*: The c6 instance family is excluded.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The GPU models.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The categories of ECS instances. Valid values:
	//
	// >  Up to 10 categories of ECS instances are supported.
	//
	// 	- General-purpose: general-purpose instance type.
	//
	// 	- Compute-optimized: compute-optimized instance type.
	//
	// 	- Memory-optimized: memory-optimized instance type.
	//
	// 	- Big data: big data instance type.
	//
	// 	- Local SSDs: instance type with local SSDs.
	//
	// 	- High Clock Speed: instance type with high clock speeds.
	//
	// 	- Enhanced: enhanced instance type.
	//
	// 	- Shared: shared instance type.
	//
	// 	- Compute-optimized with GPU: GPU-accelerated compute-optimized instance type.
	//
	// 	- Visual Compute-optimized: visual compute-optimized instance type.
	//
	// 	- Heterogeneous Service: heterogeneous service instance type.
	//
	// 	- Compute-optimized with FPGA: FPGA-accelerated compute-optimized instance type.
	//
	// 	- Compute-optimized with NPU: NPU-accelerated compute-optimized instance type.
	//
	// 	- ECS Bare Metal: ECS Bare Metal Instance type.
	//
	// 	- High Performance Compute: HPC-optimized instance type.
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family.
	//
	// 	- EntryLevel: entry level (shared instance types). Instance types of this level are the most cost-effective but may not provide stable computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise level. Instance types of this level provide stable performance and dedicated resources, and are suitable for scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit entry level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families that are queried. You can query 1 to 10 instance families in each call.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum hourly price for the pay-as-you-go or preemptible instances.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPUs per instance type.
	//
	// >  The value of MaximumCpuCoreCount cannot exceed four times the value of MinimumCpuCoreCount.
	//
	// example:
	//
	// 4
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum number of GPUs per instance. The value must be a positive integer.
	//
	// example:
	//
	// 2
	MaximumGpuAmount *int32 `json:"MaximumGpuAmount,omitempty" xml:"MaximumGpuAmount,omitempty"`
	// The maximum memory size per instance. Unit: GiB.
	//
	// example:
	//
	// 4
	MaximumMemorySize *float32 `json:"MaximumMemorySize,omitempty" xml:"MaximumMemorySize,omitempty"`
	// The memory size of the instance type. Unit: GiB.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The baseline vCPU computing performance (overall baseline performance of all vCPUs) per t5 or t6 burstable instance.
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPUs per instance type.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum number of IPv6 addresses per ENI.
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of IPv4 addresses per ENI.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of elastic network interfaces (ENIs) per instance.
	//
	// example:
	//
	// 2
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of GPUs per instance. The value must be a positive integer.
	//
	// example:
	//
	// 2
	MinimumGpuAmount *int32 `json:"MinimumGpuAmount,omitempty" xml:"MinimumGpuAmount,omitempty"`
	// The initial vCPU credits per t5 or t6 burstable instance.
	//
	// example:
	//
	// 12
	MinimumInitialCredit *int32 `json:"MinimumInitialCredit,omitempty" xml:"MinimumInitialCredit,omitempty"`
	// The minimum memory size per instance. Unit: GiB.
	//
	// example:
	//
	// 4
	MinimumMemorySize *float32 `json:"MinimumMemorySize,omitempty" xml:"MinimumMemorySize,omitempty"`
	// The processor models of the instance types. You can specify 1 to 10 processor models.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetArchitectures() []*string {
	return s.Architectures
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMaximumGpuAmount() *int32 {
	return s.MaximumGpuAmount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumGpuAmount() *int32 {
	return s.MinimumGpuAmount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetArchitectures(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetBurstablePerformance(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetCores(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetCpuArchitectures(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.CpuArchitectures = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetGpuSpecs(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.GpuSpecs = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetInstanceCategories(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.InstanceCategories = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetInstanceFamilyLevel(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetInstanceTypeFamilies(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMaxPrice(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMaximumCpuCoreCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMaximumGpuAmount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MaximumGpuAmount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMaximumMemorySize(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MaximumMemorySize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMemory(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.Memory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumBaselineCredit(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumCpuCoreCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumEniIpv6AddressQuantity(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumEniPrivateIpAddressQuantity(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumEniQuantity(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumEniQuantity = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumGpuAmount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumGpuAmount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumInitialCredit(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumInitialCredit = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetMinimumMemorySize(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.MinimumMemorySize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) SetPhysicalProcessorModels(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos {
	s.PhysicalProcessorModels = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces struct {
	// The ENI type. Valid values:
	//
	// 	- Primary: the primary ENI
	//
	// 	- Secondary: the secondary ENI
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses that are allocated to the primary ENI.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication mode of the ENI. Valid values:
	//
	// 	- Standard: The TCP communication mode is used.
	//
	// 	- HighPerformance: The Elastic RDMA Interface (ERI) is enabled and the remote direct memory access (RDMA) communication mode is used.
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode    *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	SecondaryPrivateIpAddressCount *int32  `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The IDs of the security groups to which the ENIs belong.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) SetIpv6AddressCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	s.Ipv6AddressCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) SetNetworkInterfaceTrafficMode(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) SetSecondaryPrivateIpAddressCount(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) SetSecurityGroupIds(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions struct {
	// The IDs of private pools. The ID of a private pool is the same as the ID of the elasticity assurance or capacity reservation that is associated with the private pool.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// The resource pool used for instance creation, which can be the public pool or a private pool associated with any active elasticity assurance or capacity reservation. Valid values:
	//
	// 	- PrivatePoolFirst: prioritizes private pools. When this option is set along with ResourcePoolOptions.PrivatePoolIds, the specified private pools are used first. If you leave ResourcePoolOptions.PrivatePoolIds empty or if the specified private pools lack sufficient capacity, the system will automatically use available open private pools instead. If no matching private pools are available, the system defaults to the public pool.
	//
	// 	- PrivatePoolOnly: uses only private pools. If you use this value, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pools lack sufficient capacity, instance creation will fail.
	//
	// 	- None: uses no resource pools.
	//
	// example:
	//
	// PrivatePoolFirst
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) SetPrivatePoolIds(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) SetStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions struct {
	// >  This parameter is in invitational preview and is not available for use.
	//
	// example:
	//
	// testManagedPrivateSpaceId
	ManagedPrivateSpaceId *string `json:"ManagedPrivateSpaceId,omitempty" xml:"ManagedPrivateSpaceId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) GetManagedPrivateSpaceId() *string {
	return s.ManagedPrivateSpaceId
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) SetManagedPrivateSpaceId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions {
	s.ManagedPrivateSpaceId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions struct {
	// example:
	//
	// TDX
	ConfidentialComputingMode *string `json:"ConfidentialComputingMode,omitempty" xml:"ConfidentialComputingMode,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions) SetConfidentialComputingMode(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits struct {
	// The instance type of the preemptible instances.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The price limit of the preemptible instances.
	//
	// example:
	//
	// 0.125
	PriceLimit *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) GetPriceLimit() *float32 {
	return s.PriceLimit
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) SetPriceLimit(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags struct {
	// The tag key of the ECS instance. You can specify up to 20 tags for each ECS instance.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with `acs:` or `aliyun` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// binary
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the ECS instance. You can specify up to 20 tags for each ECS instance.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// alterTable
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) SetKey(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Key = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) SetValue(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags {
	s.Value = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags) Validate() error {
	return dara.Validate(s)
}
