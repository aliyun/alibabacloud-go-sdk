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
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
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
	// Specifies whether the instance on a dedicated host is associated with the dedicated host. Valid values:
	//
	// - default: The instance is not associated with the dedicated host. If the economical mode is enabled and the original dedicated host has insufficient resources when the instance is restarted, the instance is automatically deployed to another dedicated host in the automatic deployment resource pool.
	//
	// - host: The instance is associated with the dedicated host. If the economical mode is enabled and the original dedicated host has insufficient resources when the instance is restarted, the instance fails to be restarted.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of vCPUs.
	//
	// You can specify both the Cpu and Memory parameters to define the instance type range. For example, if you set Cpu to 2 and Memory to 16, all instance types that have 2 vCPUs and 16 GiB of memory are matched. Auto Scaling determines the available instance types based on factors such as I/O optimization and zones, and creates instances of the type that has the lowest price.
	//
	// > This configuration takes effect only when the scaling mode is set to cost optimization and no instance types are specified in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The CPU options.
	CpuOptions *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// The time when the scaling configuration was created. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information, see the "Standard mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information, see the "Unlimited mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The custom priority of ECS instance type and vSwitch combination.
	//
	// 	Notice: This parameter takes effect only when the scaling policy of the scaling group is set to priority-based policy.</notice>
	//
	// If an instance cannot be created by using ECS instance type and vSwitch combination that has the highest priority, Auto Scaling automatically uses the combination that has the next highest priority to create ECS instance.
	//
	// > If you specify custom priorities for only some instance type and vSwitch combinations, the unspecified combinations have lower priorities than the specified ones. The priorities of the unspecified combinations are determined based on the vSwitch sequence in the scaling group and ECS instance type sequence in the scaling configuration.
	//
	// > - For example, if the vSwitch sequence in the scaling group is vsw1 and vsw2, ECS instance type sequence in the scaling configuration is type1 and type2, and the custom priority sequence is ["vsw2+type2", "vsw1+type2"], the final priority sequence is: "vsw2+type2" > "vsw1+type2" > "vsw1+type1" > "vsw2+type1".
	CustomPriorities []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The data cloud disks.
	DataDisks []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-zm04u8r3lohsq****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host on which to create the ECS instance. Dedicated hosts do not support spot instances. If you specify the DedicatedHostId parameter, the SpotStrategy and SpotPriceLimit parameters are ignored.
	//
	// You can call the DescribeDedicatedHosts operation to query the list of dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection property of the instance, which specifies whether the instance can be directly released by using the ECS console or the DeleteInstance API operation. Valid values:
	//
	// - true: Release protection is enabled. The instance cannot be directly released by using the ECS console or the DeleteInstance API operation.
	//
	// - false: Release protection is shutdown. The instance can be directly released by using the ECS console or the DeleteInstance API operation.
	//
	// > This property applies only to pay-as-you-go instances. It prevents instances created by Auto Scaling during scale-out from being accidentally deleted, but does not affect normal scale-in activities. Instances with release protection enabled can still be released during scale-in activities.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set to which the ECS instance belongs.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the ECS instance.
	//
	// example:
	//
	// LocalHost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the High Performance Computing (HPC) cluster to which the ECS instance belongs.
	//
	// example:
	//
	// hpc-clus****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enabled.
	//
	// - disabled: disabled.
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Specifies whether to forcibly use the security hardening mode (IMDSv2) when accessing instance metadata. Valid values:
	//
	// - optional: not forcibly used.
	//
	// - required: forcibly used. If this value is set, instance metadata cannot be accessed in normal mode.
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image in the image family to create instances. If the ImageId parameter is specified, you cannot specify this parameter.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file, which is the image resource used for the automatic creation of instances.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Specifies whether to log on to the ECS instance as the ecs-user user. Valid values:
	//
	// - true: yes.
	//
	// - false: no.
	//
	// example:
	//
	// false
	ImageOptionsLoginAsNonRoot *bool `json:"ImageOptionsLoginAsNonRoot,omitempty" xml:"ImageOptionsLoginAsNonRoot,omitempty"`
	// The source of the image. Valid values:
	//
	// - system: public images provided by Alibaba Cloud.
	//
	// - self: custom images created by you.
	//
	// - others: shared images or community images provided by other Alibaba Cloud users.
	//
	// - marketplace: images provided by Alibaba Cloud Marketplace.
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The description of the ECS instance.
	//
	// example:
	//
	// FinanceDept
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The generation of the ECS instance.
	//
	// example:
	//
	// ecs-3
	InstanceGeneration *string `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// instance****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The collection of intelligent configuration information used to filter instance types that meet your requirements.
	InstancePatternInfos []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance type of the ECS instance.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// After the fallback mode is configured, if issues such as insufficient inventory occur, the system supplements the selected instance types with similar instance types of the same size, or creates vSwitches in fallback zones and adds them to the scaling group.
	InstanceTypeCandidateOptions *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions `json:"InstanceTypeCandidateOptions,omitempty" xml:"InstanceTypeCandidateOptions,omitempty" type:"Struct"`
	// The collection of instance types for ECS instances.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The network billing method. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth. In this case, the InternetMaxBandwidthOut parameter specifies the fixed bandwidth value.
	//
	// - PayByTraffic: pay-by-traffic. In this case, the InternetMaxBandwidthOut parameter specifies the bandwidth cap, and you are billed based on the actual network traffic.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is an I/O optimized instance. Valid values:
	//
	// - none: non-I/O optimization.
	//
	// - optimized: I/O optimization.
	//
	// example:
	//
	// none
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of IPv6 addresses that are randomly generated for the elastic network interface controller (NIC).
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair used to log on to the ECS instance.
	//
	// example:
	//
	// keypair****
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The lifecycle status of the scaling configuration in the scaling group. Valid values:
	//
	// - Active: The scaling configuration is active. The scaling group uses active scaling configurations for the automatic creation of ECS instances.
	//
	// - Inacitve: The scaling configuration is inactive. Inactive scaling configurations exist in the scaling group, but the scaling group does not use them for the automatic creation of ECS instances.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The weight of the ECS instance when it is used as an SLB backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify both the Cpu and Memory parameters to define the instance type range. For example, if you set Cpu to 2 and Memory to 16, all instance types that have 2 vCPUs and 16 GiB of memory are matched. Auto Scaling determines the available instance types based on factors such as I/O optimization and zones, and creates instances of the type that has the lowest price.
	//
	// > This configuration takes effect only when the scaling mode is set to cost optimization and no instance types are specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The list of elastic network interfaces (ENIs). Each ENI is a network interface controller (NIC).
	NetworkInterfaces []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	// Specifies whether to use the password preset in the image.
	//
	// example:
	//
	// true
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// Specifies whether an instance password is set. Valid values:
	//
	// - true: An instance password is set.
	//
	// - false: No instance password is set.
	//
	// example:
	//
	// false
	PasswordSetted *bool `json:"PasswordSetted,omitempty" xml:"PasswordSetted,omitempty"`
	// The ID of the private pool. The ID is the same as the ID of the elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	PrivatePoolOptions_id *string `json:"PrivatePoolOptions.Id,omitempty" xml:"PrivatePoolOptions.Id,omitempty"`
	// The private pool type for starting the instance. After an elasticity assurance or capacity reservation takes effect, a private pool is generated. You can select a private pool type when you start an instance. Valid values:
	//
	// - Open: open mode. The system automatically matches a private pool of the Open type. If no matching private pool is available, the system uses resources from the public pool to start the instance.
	//
	// - Target: targeted mode. The system uses the specified private pool to start the instance. If the specified private pool is unavailable, the instance fails to be started.
	//
	// - None: no private pool. The system does not use private pools to start the instance.
	//
	// example:
	//
	// Open
	PrivatePoolOptions_matchCriteria *string `json:"PrivatePoolOptions.MatchCriteria,omitempty" xml:"PrivatePoolOptions.MatchCriteria,omitempty"`
	// The name of the Resource Access Management (RAM) role for the ECS instance. RAM roles are provided and maintained by RAM. You can call the ListRoles operation to query available RAM roles.
	//
	// example:
	//
	// ramrole****
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the ECS instance belongs.
	//
	// example:
	//
	// rg-aekzn2ou7xo****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource pool policy used when creating instances.
	//
	// - This parameter takes effect only when creating pay-as-you-go instances.
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
	// > This parameter is in invitational preview and is not available for public use.
	SchedulerOptions *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	// Specifies whether a vSwitch available for secondary elastic network interfaces (ENIs) must be selected from the list when the VSwitchIds parameter is specified.
	//
	// example:
	//
	// Strict
	SecondaryEniVSwitchConstraint *string `json:"SecondaryEniVSwitchConstraint,omitempty" xml:"SecondaryEniVSwitchConstraint,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// - Active: enables security hardening. This value is applicable only to public images.
	//
	// - Deactive: disables security hardening. This value is applicable to all image types.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the ECS instance belongs. ECS instances in the same security group can access each other.
	//
	// example:
	//
	// sg-bp18kz60mefs****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups to which the ECS instances belong. ECS instances in the same security group can access each other.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *DescribeScalingConfigurationsResponseBodyScalingConfigurationsSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption pattern of the spot instance.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The spot instances.
	SpotPriceLimits []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy for the pay-as-you-go instance. Valid values:
	//
	// - NoSpot: regular pay-as-you-go instances.
	//
	// - SpotWithPriceLimit: spot instances that have a maximum hourly price.
	//
	// - SpotAsPriceGo: spot instances for which the system automatically bids based on the current market price.
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
	// The maximum number of partitions in the storage set. The value must be an integer greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The ID of the automatic snapshot policy applied to the system cloud disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	SystemDiskAutoSnapshotPolicyId *string `json:"SystemDiskAutoSnapshotPolicyId,omitempty" xml:"SystemDiskAutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burst feature for the system cloud disk. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// > This parameter is supported only when SystemDisk.Category is set to cloud_auto.
	//
	// example:
	//
	// false
	SystemDiskBurstingEnabled *bool `json:"SystemDiskBurstingEnabled,omitempty" xml:"SystemDiskBurstingEnabled,omitempty"`
	// The disk types for the system cloud disk. The first disk type has the highest priority, and the priorities of the other disk types decrease in sequence. If the disk type with a higher priority is unavailable, Auto Scaling automatically attempts to create the system cloud disk by using the disk type with the next highest priority. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD cloud disk.
	//
	// - cloud_essd: ESSD cloud disk.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The category of the system cloud disk. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// example:
	//
	// cloud
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The description of the system cloud disk.
	//
	// example:
	//
	// Test system disk.
	SystemDiskDescription *string `json:"SystemDiskDescription,omitempty" xml:"SystemDiskDescription,omitempty"`
	// The encryption algorithm used for the system cloud disk. Valid values:
	//
	// - AES-256.
	//
	// - SM4-128.
	//
	// example:
	//
	// AES-256
	SystemDiskEncryptAlgorithm *string `json:"SystemDiskEncryptAlgorithm,omitempty" xml:"SystemDiskEncryptAlgorithm,omitempty"`
	// Specifies whether the system cloud disk is encrypted. Valid values:
	//
	// - true: The system cloud disk is encrypted.
	//
	// - false: The system cloud disk is not encrypted.
	//
	// example:
	//
	// false
	SystemDiskEncrypted *bool `json:"SystemDiskEncrypted,omitempty" xml:"SystemDiskEncrypted,omitempty"`
	// The ID of the KMS key used for the system cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	SystemDiskKMSKeyId *string `json:"SystemDiskKMSKeyId,omitempty" xml:"SystemDiskKMSKeyId,omitempty"`
	// The name of the system cloud disk.
	//
	// example:
	//
	// cloud_ssd_Test
	SystemDiskName *string `json:"SystemDiskName,omitempty" xml:"SystemDiskName,omitempty"`
	// The performance level of the system cloud disk when it is an ESSD (standard SSD).
	//
	// example:
	//
	// PL1
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The provisioned IOPS (Input/Output Operations Per Second) performance metrics for the system cloud disk.
	//
	// > IOPS refers to the number of I/O operations that can be processed per second, which indicates the read and write processing capability of the block storage. Unit: operations per second.
	//
	// example:
	//
	// 100
	SystemDiskProvisionedIops *int64 `json:"SystemDiskProvisionedIops,omitempty" xml:"SystemDiskProvisionedIops,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 100
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags.
	Tags []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// - default: creates the instance on a non-dedicated host.
	//
	// - host: creates the instance on a dedicated host. If you do not specify the DedicatedHostId parameter, Alibaba Cloud automatically selects a dedicated host for the instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// The user data of the ECS instance.
	//
	// example:
	//
	// echo hello ecs!
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The weight of the specified instance type, which indicates the capacity that a single instance of the specified instance type provides in the scaling group. A greater weight indicates that fewer instances of the instance type are required to meet the expected capacity.
	WeightedCapacities []*int32 `json:"WeightedCapacities,omitempty" xml:"WeightedCapacities,omitempty" type:"Repeated"`
	// The zone ID of the instance. You can call the DescribeZones operation to query the zone list.
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetCpuOptions() *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions {
	return s.CpuOptions
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetInstanceTypeCandidateOptions() *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions {
	return s.InstanceTypeCandidateOptions
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) GetSecondaryEniVSwitchConstraint() *string {
	return s.SecondaryEniVSwitchConstraint
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetCpuOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.CpuOptions = v
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetInstanceTypeCandidateOptions(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.InstanceTypeCandidateOptions = v
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetSecondaryEniVSwitchConstraint(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.SecondaryEniVSwitchConstraint = &v
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
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
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
	if s.InstanceTypeCandidateOptions != nil {
		if err := s.InstanceTypeCandidateOptions.Validate(); err != nil {
			return err
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

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions struct {
	// Specifies whether to enable hardware-based nested virtualization. Valid values:
	//
	// 	- enabled: Enabled.
	//
	// 	- disabled: Disabled.
	//
	// example:
	//
	// enabled
	NestedVirtualization *string `json:"NestedVirtualization,omitempty" xml:"NestedVirtualization,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions) GetNestedVirtualization() *string {
	return s.NestedVirtualization
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions) SetNestedVirtualization(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions {
	s.NestedVirtualization = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsCpuOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsCustomPriorities struct {
	// The instance type of the ECS instance.
	//
	// example:
	//
	// ecs.c6a.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the vSwitch.
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
	// The ID of the automatic snapshot policy applied to the data cloud disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature for the data cloud disk. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// > This parameter is displayed only when the DataDisk.Category parameter is set to cloud_auto.
	//
	// <props="china">For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The disk types of the data cloud disks. The first disk type has the highest priority, and the priorities of other disk types decrease in sequence. If a disk type with a higher priority is unavailable, Auto Scaling automatically attempts to create the data cloud disk by using the disk type with the next highest priority. Valid values:
	//
	// - cloud: basic cloud disk. The DeleteWithInstance property of a basic cloud disk created along with an instance is set to true.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data cloud disk. Valid values:
	//
	// - cloud: basic cloud disk. The DeleteWithInstance attribute of a basic cloud disk created along with an instance is set to true.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// example:
	//
	// cloud
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data cloud disk when the instance is released. Valid values:
	//
	// - true: releases the data cloud disk when the instance is released.
	//
	// - false: retains the data cloud disk when the instance is released.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data cloud disk.
	//
	// example:
	//
	// FinanceDept
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data cloud disk.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data cloud disk.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the data cloud disk. Valid values:
	//
	// - true: encrypts the data cloud disk.
	//
	// - false: does not encrypt the data cloud disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key used to encrypt the data cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the data cloud disk when it is an ESSD (standard SSD).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS (Input/Output Operations Per Second) performance metrics of the data cloud disk.
	//
	// > IOPS measures the number of I/O operations per second and indicates the read and write capabilities of the block storage. Unit: operations per second.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the data cloud disk. Unit: GiB. Valid values:
	//
	// - cloud: 5 to 2000.
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_ssd: 20 to 32768.
	//
	// - cloud_essd: 20 to 32768.
	//
	// - ephemeral_ssd: 5 to 800.
	//
	// example:
	//
	// 200
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create the data cloud disk.
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
	// The architecture type of the instance type. Valid values:
	//
	// - X86: x86 computing.
	//
	// - Heterogeneous: heterogeneous computing, such as GPU or FPGA.
	//
	// - BareMental: ECS Bare Metal server.
	//
	// - Arm: Arm computing.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether the instance type is a burstable instance type. Valid values:
	//
	// - Exclude: excludes burstable instance types.
	//
	// - Include: includes burstable instance types.
	//
	// - Required: includes only burstable instance types.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The number of vCPU cores of the instance type.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architecture of the instance type. Valid values:
	//
	// > N indicates that you can specify multiple CPU architectures. Valid values of N: 1 to 2.
	//
	// - X86.
	//
	// - ARM.
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The instance types to be excluded. You can use a wildcard (*) to exclude a single instance type or an entire instance family. Examples:
	//
	// - ecs.c6.large: excludes the ecs.c6.large instance type.
	//
	// - ecs.c6.*: excludes all instance types in the c6 instance family.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The GPU type.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The instance categorization. Valid values:
	//
	// > N indicates that you can specify multiple instance categories in the Settings. Valid values of N: 1 to 10.
	//
	// - General-purpose: general-purpose.
	//
	// - Compute-optimized: compute-optimized.
	//
	// - Memory-optimized: memory-optimized.
	//
	// - Big data: big data.
	//
	// - Local SSDs: local SSD.
	//
	// - High Clock Speed: high frequency.
	//
	// - Enhanced: enhanced.
	//
	// - Shared: shared.
	//
	// - Compute-optimized with GPU: GPU computing.
	//
	// - Visual Compute-optimized: visual compute-optimized.
	//
	// - Heterogeneous Service: heterogeneous service.
	//
	// - Compute-optimized with FPGA: FPGA-accelerated compute-optimized.
	//
	// - Compute-optimized with NPU: NPU-accelerated compute-optimized.
	//
	// - ECS Bare Metal: ECS Bare Metal server.
	//
	// - High Performance Compute: high-performance computing (HPC).
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family. Valid values:
	//
	// - EntryLevel: entry-level instance families, which are shared instance families. These instance families are more cost-effective but do not guarantee consistent computing performance. They are suitable for workloads with low CPU utilization. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise-level instance families. These instance families provide stable performance and dedicated resources, and are suitable for workloads that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit entry-level instance families, which are burstable instances. These instance families use CPU credits to maintain computing performance and are suitable for workloads with low CPU utilization but occasional CPU bursts. For more information, see [Burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families to be queried. N indicates that you can specify multiple instance families. Valid values of N: 1 to 10.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum hourly price acceptable for pay-as-you-go or spot instances.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPU cores of the instance type.
	//
	// > The value of MaximumCpuCoreCount cannot exceed four times the value of MinimumCpuCoreCount.
	//
	// example:
	//
	// 4
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum number of GPUs for the instance. Valid values: positive integers.
	//
	// example:
	//
	// 2
	MaximumGpuAmount *int32 `json:"MaximumGpuAmount,omitempty" xml:"MaximumGpuAmount,omitempty"`
	// The maximum memory size of the instance. Unit: GiB.
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
	// The minimum baseline vCPU computing performance (the sum of all vCPUs) for t5 and t6 burstable instances.
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPU cores of the instance type.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum value for the maximum number of IPv6 addresses supported by a single elastic network interfaces (ENIs) on the instance. Each ENI is a network interface controller (NIC).
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum value for the maximum number of IPv4 addresses supported by a single elastic network interfaces (ENIs) on the instance. Each ENI is a network interface controller (NIC).
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum value for the maximum number of elastic network interfaces (ENIs) that can be attached to the instance.
	//
	// example:
	//
	// 2
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of GPUs for the instance. Valid values: positive integers.
	//
	// example:
	//
	// 2
	MinimumGpuAmount *int32 `json:"MinimumGpuAmount,omitempty" xml:"MinimumGpuAmount,omitempty"`
	// The minimum initial vCPU credit value for t5 and t6 burstable instances.
	//
	// example:
	//
	// 12
	MinimumInitialCredit *int32 `json:"MinimumInitialCredit,omitempty" xml:"MinimumInitialCredit,omitempty"`
	// The minimum memory size of the instance. Unit: GiB.
	//
	// example:
	//
	// 4
	MinimumMemorySize *float32 `json:"MinimumMemorySize,omitempty" xml:"MinimumMemorySize,omitempty"`
	// The processor model of the instance. N indicates that you can specify multiple processor models. Valid values of N: 1 to 10.
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

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions struct {
	// The CIDR blocks of the vSwitches that can be selected when adding vSwitches from other zones is allowed.
	AllowCidrBlocks []*string `json:"AllowCidrBlocks,omitempty" xml:"AllowCidrBlocks,omitempty" type:"Repeated"`
	// Specifies whether to allow adding vSwitches from other zones.
	//
	// > The instance types remain unchanged, and only new zones are selected as alternatives. When the scaling group fails to scale out in all selected zones due to inventory shortages or other issues, this configuration automatically adds vSwitches from new zones to the scaling group.
	//
	// > For example, if the scaling group is configured with the cn-hangzhou-h and cn-hangzhou-g zones and fails to scale out in both zones, Auto Scaling may create a vSwitch in the cn-hangzhou-k zone and add it to the scaling group based on real-time inventory.
	//
	// example:
	//
	// true
	AllowCrossAz *bool `json:"AllowCrossAz,omitempty" xml:"AllowCrossAz,omitempty"`
	// Specifies whether to allow supplementing instance types from other generations.
	//
	// - For example, if the current instance type is ecs.c7.large, enabling this option allows fallback to instance types such as ecs.c6.large and ecs.c8.large.
	//
	// example:
	//
	// true
	AllowDifferentGeneration *bool `json:"AllowDifferentGeneration,omitempty" xml:"AllowDifferentGeneration,omitempty"`
	// Specifies whether to enable the fallback mode.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum price for fallback instance types.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) GetAllowCidrBlocks() []*string {
	return s.AllowCidrBlocks
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) GetAllowCrossAz() *bool {
	return s.AllowCrossAz
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) GetAllowDifferentGeneration() *bool {
	return s.AllowDifferentGeneration
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) SetAllowCidrBlocks(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions {
	s.AllowCidrBlocks = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) SetAllowCrossAz(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions {
	s.AllowCrossAz = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) SetAllowDifferentGeneration(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions {
	s.AllowDifferentGeneration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) SetEnabled(v bool) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions {
	s.Enabled = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) SetMaxPrice(v float32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions {
	s.MaxPrice = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsInstanceTypeCandidateOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces struct {
	// The type of the elastic network interfaces (ENIs). Valid values:
	//
	// - Primary: primary network interface controller (NIC).
	//
	// - Secondary: secondary elastic network interfaces (ENIs).
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses specified for the primary ENI.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication pattern of the elastic network interfaces (ENIs), which act as a network interface controller (NIC). Valid values:
	//
	// - Standard: uses the TCP communication pattern.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication pattern.
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of secondary private IPv4 addresses specified for the ENI. Valid values: 1 to 49.
	//
	// - The value cannot exceed the limit on the number of IP addresses for the instance type. For more information, see [Instance families](https://www.alibabacloud.com/help/en/ecs/user-guide/overview-of-instance-families).
	//
	// - NetworkInterface.N.SecondaryPrivateIpAddressCount is used to specify the number of secondary private IPv4 addresses (excluding the primary private IP address of the ENI) to be assigned to the ENI. The system randomly assigns the addresses from the available CIDR block of the vSwitch (NetworkInterface.N.VSwitchId) to which the ENI belongs.
	//
	// example:
	//
	// 2
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The IDs of one or more security groups to which the elastic network interfaces (ENIs) belong. Each ENI is a network interface controller (NIC).
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the elastic network interfaces (ENIs) belong. Each ENI is a network interface controller (NIC).
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) GetVSwitchIds() []*string {
	return s.VSwitchIds
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) SetVSwitchIds(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces {
	s.VSwitchIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions struct {
	// The ID of the private pool, which is the ID of the elasticity assurance or capacity reservation.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// Filters available Target private pools by tag.
	PrivatePoolTags []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags `json:"PrivatePoolTags,omitempty" xml:"PrivatePoolTags,omitempty" type:"Repeated"`
	// Resource pools include private pools generated after the elasticity assurance or capacity reservation service takes effect, as well as public pools, which are available for selection when an instance starts. Valid values:
	//
	// - PrivatePoolFirst: private pool first. When this policy is selected, if ResourcePoolOptions.PrivatePoolIds is specified or the PrivatePoolTags condition is met, the corresponding private pool is used first. If no private pool is specified or the specified private pool has insufficient capacity, the system automatically matches an open private pool. If no eligible private pool is available, the public pool is used to create the instance.
	//
	// - PrivatePoolOnly: private pool only. When this policy is selected, ResourcePoolOptions.PrivatePoolIds must be specified. If the specified private pool has insufficient capacity, the instance fails to start.
	//
	// - PublicPoolFirst: public pool first. The public pool is used first to create the instance. When public pool resources are insufficient, private pool resources are used as a supplement. The system preferentially matches an open private pool. If no eligible private pool is available, the system uses the Target private pool specified by ResourcePoolOptions.PrivatePoolIds or the one that meets the PrivatePoolTags condition.
	//
	// - None: no resource pool policy is used.
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

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) GetPrivatePoolTags() []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags {
	return s.PrivatePoolTags
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) SetPrivatePoolIds(v []*string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) SetPrivatePoolTags(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions {
	s.PrivatePoolTags = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) SetStrategy(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptions) Validate() error {
	if s.PrivatePoolTags != nil {
		for _, item := range s.PrivatePoolTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags struct {
	// The tag key of the private pool.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the private pool.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) GetKey() *string {
	return s.Key
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) GetValue() *string {
	return s.Value
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) SetKey(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags {
	s.Key = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) SetValue(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags {
	s.Value = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsResourcePoolOptionsPrivatePoolTags) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsSchedulerOptions struct {
	// > This parameter is in invitational preview and is not available for public use.
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
	// The confidential computing mode. Valid values:
	//
	// - Enclave: The ECS instance uses Enclave to build a confidential computing environment. For more information, see [Build a confidential computing environment by using Enclave](https://help.aliyun.com/document_detail/203433.html).
	//
	// - TDX: builds a TDX confidential computing environment. For more information, see [Build a TDX confidential computing environment](https://help.aliyun.com/document_detail/479090.html).
	//
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
	// The instance type of the spot instance.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The bid price for the spot instance.
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
	// The key of tag N of the instance. Valid values of N: 1 to 20.
	//
	// If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// binary
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the instance. Valid values of N: 1 to 20.
	//
	// If you specify this parameter, it can be an empty string. The tag value can be up to 128 characters in length. It cannot start with acs:, and cannot contain http:// or https://.
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
