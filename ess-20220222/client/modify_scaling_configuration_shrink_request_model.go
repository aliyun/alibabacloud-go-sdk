// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageOptions(v *ModifyScalingConfigurationShrinkRequestImageOptions) *ModifyScalingConfigurationShrinkRequest
	GetImageOptions() *ModifyScalingConfigurationShrinkRequestImageOptions
	SetPrivatePoolOptions(v *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) *ModifyScalingConfigurationShrinkRequest
	GetPrivatePoolOptions() *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions
	SetSystemDisk(v *ModifyScalingConfigurationShrinkRequestSystemDisk) *ModifyScalingConfigurationShrinkRequest
	GetSystemDisk() *ModifyScalingConfigurationShrinkRequestSystemDisk
	SetAffinity(v string) *ModifyScalingConfigurationShrinkRequest
	GetAffinity() *string
	SetCpu(v int32) *ModifyScalingConfigurationShrinkRequest
	GetCpu() *int32
	SetCreditSpecification(v string) *ModifyScalingConfigurationShrinkRequest
	GetCreditSpecification() *string
	SetCustomPriorities(v []*ModifyScalingConfigurationShrinkRequestCustomPriorities) *ModifyScalingConfigurationShrinkRequest
	GetCustomPriorities() []*ModifyScalingConfigurationShrinkRequestCustomPriorities
	SetDataDisks(v []*ModifyScalingConfigurationShrinkRequestDataDisks) *ModifyScalingConfigurationShrinkRequest
	GetDataDisks() []*ModifyScalingConfigurationShrinkRequestDataDisks
	SetDedicatedHostClusterId(v string) *ModifyScalingConfigurationShrinkRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *ModifyScalingConfigurationShrinkRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *ModifyScalingConfigurationShrinkRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *ModifyScalingConfigurationShrinkRequest
	GetDeploymentSetId() *string
	SetHostName(v string) *ModifyScalingConfigurationShrinkRequest
	GetHostName() *string
	SetHpcClusterId(v string) *ModifyScalingConfigurationShrinkRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *ModifyScalingConfigurationShrinkRequest
	GetHttpEndpoint() *string
	SetHttpTokens(v string) *ModifyScalingConfigurationShrinkRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *ModifyScalingConfigurationShrinkRequest
	GetImageFamily() *string
	SetImageId(v string) *ModifyScalingConfigurationShrinkRequest
	GetImageId() *string
	SetImageName(v string) *ModifyScalingConfigurationShrinkRequest
	GetImageName() *string
	SetInstanceDescription(v string) *ModifyScalingConfigurationShrinkRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *ModifyScalingConfigurationShrinkRequest
	GetInstanceName() *string
	SetInstancePatternInfos(v []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos) *ModifyScalingConfigurationShrinkRequest
	GetInstancePatternInfos() []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos
	SetInstanceTypeOverrides(v []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) *ModifyScalingConfigurationShrinkRequest
	GetInstanceTypeOverrides() []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides
	SetInstanceTypes(v []*string) *ModifyScalingConfigurationShrinkRequest
	GetInstanceTypes() []*string
	SetInternetChargeType(v string) *ModifyScalingConfigurationShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *ModifyScalingConfigurationShrinkRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *ModifyScalingConfigurationShrinkRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *ModifyScalingConfigurationShrinkRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *ModifyScalingConfigurationShrinkRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *ModifyScalingConfigurationShrinkRequest
	GetKeyPairName() *string
	SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationShrinkRequest
	GetLoadBalancerWeight() *int32
	SetMemory(v int32) *ModifyScalingConfigurationShrinkRequest
	GetMemory() *int32
	SetNetworkInterfaces(v []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces) *ModifyScalingConfigurationShrinkRequest
	GetNetworkInterfaces() []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces
	SetOverride(v bool) *ModifyScalingConfigurationShrinkRequest
	GetOverride() *bool
	SetOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyScalingConfigurationShrinkRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifyScalingConfigurationShrinkRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *ModifyScalingConfigurationShrinkRequest
	GetPasswordInherit() *bool
	SetRamRoleName(v string) *ModifyScalingConfigurationShrinkRequest
	GetRamRoleName() *string
	SetResourceGroupId(v string) *ModifyScalingConfigurationShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourcePoolOptions(v *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) *ModifyScalingConfigurationShrinkRequest
	GetResourcePoolOptions() *ModifyScalingConfigurationShrinkRequestResourcePoolOptions
	SetScalingConfigurationId(v string) *ModifyScalingConfigurationShrinkRequest
	GetScalingConfigurationId() *string
	SetScalingConfigurationName(v string) *ModifyScalingConfigurationShrinkRequest
	GetScalingConfigurationName() *string
	SetSchedulerOptionsShrink(v string) *ModifyScalingConfigurationShrinkRequest
	GetSchedulerOptionsShrink() *string
	SetSecurityGroupId(v string) *ModifyScalingConfigurationShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationShrinkRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *ModifyScalingConfigurationShrinkRequestSecurityOptions) *ModifyScalingConfigurationShrinkRequest
	GetSecurityOptions() *ModifyScalingConfigurationShrinkRequestSecurityOptions
	SetSpotDuration(v int32) *ModifyScalingConfigurationShrinkRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *ModifyScalingConfigurationShrinkRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimits(v []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits) *ModifyScalingConfigurationShrinkRequest
	GetSpotPriceLimits() []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits
	SetSpotStrategy(v string) *ModifyScalingConfigurationShrinkRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *ModifyScalingConfigurationShrinkRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *ModifyScalingConfigurationShrinkRequest
	GetStorageSetPartitionNumber() *int32
	SetSystemDiskCategories(v []*string) *ModifyScalingConfigurationShrinkRequest
	GetSystemDiskCategories() []*string
	SetTags(v string) *ModifyScalingConfigurationShrinkRequest
	GetTags() *string
	SetTenancy(v string) *ModifyScalingConfigurationShrinkRequest
	GetTenancy() *string
	SetUserData(v string) *ModifyScalingConfigurationShrinkRequest
	GetUserData() *string
	SetZoneId(v string) *ModifyScalingConfigurationShrinkRequest
	GetZoneId() *string
}

type ModifyScalingConfigurationShrinkRequest struct {
	ImageOptions       *ModifyScalingConfigurationShrinkRequestImageOptions       `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *ModifyScalingConfigurationShrinkRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether to associate the instance on a dedicated host with the dedicated host. Valid values:
	//
	// 	- default: does not associate the instance on the dedicated host with the dedicated host. If you restart an instance that was stopped in Economical Mode and the original dedicated host of the instance has insufficient resources, the instance is automatically deployed to another dedicated host in the automatic deployment resource pool.
	//
	// 	- host: associates the instance on a dedicated host with the dedicated host. If you restart an instance that was stopped in Economical Mode, the instance remains on the original dedicated host. If the original dedicated host has insufficient resources, the instance cannot be started.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of vCPUs.
	//
	// You can specify the number of vCPUs and the memory size to determine the range of instance types. For example, you can set Cpu to 2 and Memory to 16 to specify instance types that have 2 vCPUs and 16 GiB of memory. If you specify Cpu and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones. Then, Auto Scaling preferentially creates instances by using the lowest-priced instance type.
	//
	// > You can specify CPU and Memory to determine the range of instance types only if you set Scaling Policy to Cost Optimization Policy and you do not specify an instance type in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The performance mode of burstable instances. Valid values:
	//
	// 	- Standard: the standard mode. For more information, see the "Standard mode" section in the [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html) topic.
	//
	// 	- Unlimited: the unlimited mode. For more information, see the "Unlimited mode" section in the [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html) topic.
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The priority of the custom "ECS instance type + vSwitch" combination.
	//
	// >  This setting is valid only if the scaling policy of the scaling group is a priority policy.
	//
	// If Auto Scaling cannot create ECS instances by using the custom "ECS instance type + vSwitch" combination of the highest priority, Auto Scaling creates ECS instances by using the custom "ECS instance type + vSwitch" combination of the next highest priority.
	//
	// >  If you specify the priorities of only a part of custom "ECS instance type + vSwitch" combinations, Auto Scaling preferentially creates ECS instances by using the custom combinations that have the specified priorities. If the custom combinations that have the specified priorities do not provide sufficient resources, Auto Scaling creates ECS instances by using the custom combinations that do not have the specified priorities based on the specified orders of vSwitches and instance types.
	//
	// 	- Example: The specified order of vSwitches for your scaling group is vsw1 and vsw2, and the specified order of instance types in your scaling configuration is type1 and type 2. In addition, you use CustomPriorities to specify ["vsw2+type2", "vsw1+type2"]. In this example, the vsw2+type2 combination has the highest priority and the vsw2+type1 combination has the lowest priority. The vsw1+type2 combination has a higher priority than the vsw1+type1 combination.
	CustomPriorities []*ModifyScalingConfigurationShrinkRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The data disks.
	DataDisks []*ModifyScalingConfigurationShrinkRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-zm04u8r3lohsq****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host on which you want to create ECS instances. You cannot create preemptible instances on dedicated hosts. If you specify DedicatedHostId, SpotStrategy and SpotPriceLimit are ignored.
	//
	// You can call the DescribeDedicatedHosts operation to query the most recent list of dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// Specifies whether to enable the Release Protection feature for ECS instances. If you enable this feature, you cannot directly release the ECS instances in the ECS console or by calling the DeleteInstance operation. Valid values:
	//
	// 	- true: enables the Release Protection feature. In this case, you cannot directly release the ECS instances in the ECS console or by calling the DeleteInstance operation.
	//
	// 	- false: disables the Release Protection feature. In this case, you can directly release the ECS instances in the ECS console or by calling the DeleteInstance operation.
	//
	// >  You can enable the Release Protection feature only for pay-as-you-go instances to prevent accidental instance deletion. The Release Protection feature does not affect normal scaling activities. An instance that meets the criteria of scale-in policies can be removed from a scaling group during a scale-in event, regardless of whether you enabled the Release Protection feature for the instance.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set of the ECS instances that are created by using the scaling configuration.
	//
	// example:
	//
	// ds-bp13v7bjnj9gis****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the ECS instance. The hostname cannot start or end with a period (.) or a hyphen (-). The hostname cannot contain consecutive periods (.) or hyphens (-). Naming conventions for different types of instances:
	//
	// 	- Windows instances: The hostname must be 2 to 15 characters in length, and can contain letters, digits, and hyphens (-). The hostname cannot contain periods (.) or contain only digits.
	//
	// 	- Other instances, such as Linux instances: The hostname must be 2 to 64 characters in length. Separate a hostname into multiple segments with periods (.). Each segment can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// hos****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the Elastic High Performance Computing (E-HPC) cluster to which the ECS instances belong.
	//
	// example:
	//
	// hpc-clusterid
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: enabled.
	//
	// >  For information about instance metadata, see [Obtain instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Specifies whether to forcibly use the security hardening mode (IMDSv2) to access instance metadata. Valid values:
	//
	// 	- optional: does not forcibly use the security hardening mode (IMDSv2).
	//
	// 	- required: forcibly uses the security hardening mode (IMDSv2). If you set this parameter to required, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// >  For more information about instance metadata access modes, see [Access modes of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. If you specify this parameter, the latest custom images that are available in the specified image family are returned. Then, you can use the images to create instances. If you specify ImageId, you cannot specify ImageFamily.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image that is used by Auto Scaling to automatically create ECS instances.
	//
	// > If the image that is specified in the scaling configuration contains system disks and data disks, the data that is stored in the data disks is cleared after you modify the image.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. Each image name must be unique in a region. If you specify ImageId, ImageName is ignored.
	//
	// You cannot use ImageName to specify images from Alibaba Cloud Marketplace.
	//
	// example:
	//
	// suse11sp3_64_20G_aliaegis_2015****.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The description of the ECS instance. The description must be 2 to 256 characters in length. The description can contain letters but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the Elastic Compute Service (ECS) instance that is automatically created by using the scaling configuration.
	//
	// example:
	//
	// inst****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The intelligent configuration settings, which determine the available instance types.
	InstancePatternInfos []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance types.
	InstanceTypeOverrides []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	// The instance types. If you specify InstanceTypes, InstanceType is ignored.
	//
	// Auto Scaling creates instances based on a priority list of instance types. If it fails to create instances using the highest-priority type, it automatically moves to the next type in the priority order.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth. You are charged for the bandwidth specified by InternetMaxBandwidthOut.
	//
	// 	- PayByTraffic: pay-by-traffic. You are charged for the actual traffic generated. InternetMaxBandwidthOut specifies only the maximum available bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// 	- If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of this parameter are 1 to 10. The default value is 10.
	//
	// 	- If the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the value of `InternetMaxBandwidthOut`. The default value is the value of `InternetMaxBandwidthOut`.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 50
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether to create I/O optimized instances from the scaling configuration. Valid values:
	//
	// 	- none: creates non-I/O optimized instances from the scaling configuration.
	//
	// 	- optimized: creates I/O optimized instances from the scaling configuration.
	//
	// example:
	//
	// none
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses that you want to allocate to the elastic network interface (ENI).
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair that you can use to log on to an ECS instance.
	//
	// 	- Windows instances do not support this parameter.
	//
	// 	- By default, the username and password authentication method is disabled for Linux instances.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The weight of an ECS instance as a backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify the number of vCPUs and the memory size to determine the range of instance types. For example, you can set Cpu to 2 and Memory to 16 to specify instance types that have 2 vCPUs and 16 GiB of memory. If you specify Cpu and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones. Then, Auto Scaling preferentially creates instances by using the lowest-priced instance type.
	//
	// > You can specify CPU and Memory to determine the range of instance types only if you set Scaling Policy to Cost Optimization Policy and you do not specify an instance type in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ENIs.
	NetworkInterfaces []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	// Specifies whether to overwrite existing data. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Override     *bool   `json:"Override,omitempty" xml:"Override,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// \\`()~!@#$%^&\\*-_+=|{}[]:;\\"<>,.?/
	//
	// The password of a Windows instance cannot start with a forward slash (/).
	//
	// >  We recommend that you use HTTPS to send requests if you specify Password to avoid password leakage.
	//
	// example:
	//
	// 123abc****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password that is preconfigured in the image. Before you use this parameter, make sure that a password is configured in the image.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the RAM role that you want to attach to the ECS instance. The name is provided and maintained by Resource Access Management (RAM). You can call the ListRoles operation to query the available RAM roles. You can call the CreateRole operation to create RAM roles.
	//
	// example:
	//
	// RamRoleTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the ECS instances belong.
	//
	// example:
	//
	// abcd1234abcd****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource pools used for instance creation, which can be the public pool or a private pool associated with any active elasticity assurance or capacity reservation. When you specify this parameter, take note of the following items:
	//
	// 	- This parameter takes effect only when you create pay-as-you-go instances.
	//
	// 	- If you specify this parameter, you cannot specify PrivatePoolOptions.MatchCriteria or PrivatePoolOptions.Id.
	ResourcePoolOptions *ModifyScalingConfigurationShrinkRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
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
	// The name of the scaling configuration must be unique in a region. If you do not specify this parameter, the scaling configuration ID is used.
	//
	// example:
	//
	// test-modify
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The scheduler options.
	//
	// example:
	//
	// ["testManagedPrivateSpaceId****"]
	SchedulerOptionsShrink *string `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// The ID of the security group with which ECS instances are associated. The ECS instances that are associated with the same security group can access each other.
	//
	// example:
	//
	// sg-F876F****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups.
	SecurityGroupIds []*string                                               `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *ModifyScalingConfigurationShrinkRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of preemptible instances. Unit: hours. Valid values:
	//
	// 	- 1: After a preemptible instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, Alibaba Cloud compares the bidding price with the market price and checks the resource inventory to determine whether to release the instance.
	//
	// 	- 0: After a preemptible instance is created, Alibaba Cloud does not ensure that the instance is not automatically released within 1 hour. Alibaba Cloud compares the biding price with the market price and checks the resource inventory to determine whether to release the instance.
	//
	// >  Alibaba Cloud notifies you of ECS system events 5 minutes before an instance is released. Preemptible instances are billed by second. We recommend that you specify a protection period based on your business requirements.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the preemptible instance. Default value: Terminate. Set the value to Terminate. This value specifies that the preemptible instance is to be released.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The preemptible instance types.
	SpotPriceLimits []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy of pay-as-you-go instances. Valid values:
	//
	// 	- NoSpot: The instances are created as regular pay-as-you-go instances.
	//
	// 	- SpotWithPriceLimit: The instances are preemptible instances that have a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are preemptible instances for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The categories of the system disks. If Auto Scaling cannot create disks by using the disk category of the highest priority, Auto Scaling creates disks by using the disk category of the next highest priority. Valid values:
	//
	// 	- cloud: basic disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// >  If you specify this parameter, you cannot specify `SystemDisk.Category`.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The tags of the ECS instance. Specify the tags as key-value pairs. You can specify up to 20 tags. When you specify tag keys and tag values, take note of the following items:
	//
	// 	- A tag key can be up to 64 characters in length. The key cannot start with `acs:` or `aliyun`, and cannot contain `http://` or `https://`. The tag key cannot be an empty string.
	//
	// 	- A tag value can be up to 128 characters in length. The value cannot start with `acs:` or `aliyun`, and cannot contain `http://` or `https://`. The tag value can be an empty string.
	//
	// example:
	//
	// {"key1":"value1","key2":"value2", ... "key5":"value5"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Specifies whether to create ECS instances on dedicated hosts. Valid values:
	//
	// 	- default: creates ECS instances on non-dedicated hosts.
	//
	// 	- host: creates ECS instances on dedicated hosts. If you do not specify DedicatedHostId, the system randomly selects a dedicated host for an ECS instance.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// The user data of the Elastic Compute Service (ECS) instance. The user data must be encoded in Base64 format. The size of raw data before Base64 encoding cannot exceed 32 KB.
	//
	// example:
	//
	// echo hello ecs!
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone ID of the ECS instances that are created by using the scaling configuration.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequest) GetImageOptions() *ModifyScalingConfigurationShrinkRequestImageOptions {
	return s.ImageOptions
}

func (s *ModifyScalingConfigurationShrinkRequest) GetPrivatePoolOptions() *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSystemDisk() *ModifyScalingConfigurationShrinkRequestSystemDisk {
	return s.SystemDisk
}

func (s *ModifyScalingConfigurationShrinkRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *ModifyScalingConfigurationShrinkRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyScalingConfigurationShrinkRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *ModifyScalingConfigurationShrinkRequest) GetCustomPriorities() []*ModifyScalingConfigurationShrinkRequestCustomPriorities {
	return s.CustomPriorities
}

func (s *ModifyScalingConfigurationShrinkRequest) GetDataDisks() []*ModifyScalingConfigurationShrinkRequestDataDisks {
	return s.DataDisks
}

func (s *ModifyScalingConfigurationShrinkRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyScalingConfigurationShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyScalingConfigurationShrinkRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *ModifyScalingConfigurationShrinkRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *ModifyScalingConfigurationShrinkRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *ModifyScalingConfigurationShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInstancePatternInfos() []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	return s.InstancePatternInfos
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInstanceTypeOverrides() []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides {
	return s.InstanceTypeOverrides
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *ModifyScalingConfigurationShrinkRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyScalingConfigurationShrinkRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *ModifyScalingConfigurationShrinkRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *ModifyScalingConfigurationShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ModifyScalingConfigurationShrinkRequest) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *ModifyScalingConfigurationShrinkRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *ModifyScalingConfigurationShrinkRequest) GetNetworkInterfaces() []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *ModifyScalingConfigurationShrinkRequest) GetOverride() *bool {
	return s.Override
}

func (s *ModifyScalingConfigurationShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyScalingConfigurationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyScalingConfigurationShrinkRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *ModifyScalingConfigurationShrinkRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *ModifyScalingConfigurationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyScalingConfigurationShrinkRequest) GetResourcePoolOptions() *ModifyScalingConfigurationShrinkRequestResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *ModifyScalingConfigurationShrinkRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSchedulerOptionsShrink() *string {
	return s.SchedulerOptionsShrink
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSecurityOptions() *ModifyScalingConfigurationShrinkRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSpotPriceLimits() []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits {
	return s.SpotPriceLimits
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ModifyScalingConfigurationShrinkRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *ModifyScalingConfigurationShrinkRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *ModifyScalingConfigurationShrinkRequest) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *ModifyScalingConfigurationShrinkRequest) GetTags() *string {
	return s.Tags
}

func (s *ModifyScalingConfigurationShrinkRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *ModifyScalingConfigurationShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *ModifyScalingConfigurationShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageOptions(v *ModifyScalingConfigurationShrinkRequestImageOptions) *ModifyScalingConfigurationShrinkRequest {
	s.ImageOptions = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPrivatePoolOptions(v *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) *ModifyScalingConfigurationShrinkRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSystemDisk(v *ModifyScalingConfigurationShrinkRequestSystemDisk) *ModifyScalingConfigurationShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetAffinity(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetCpu(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetCreditSpecification(v string) *ModifyScalingConfigurationShrinkRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetCustomPriorities(v []*ModifyScalingConfigurationShrinkRequestCustomPriorities) *ModifyScalingConfigurationShrinkRequest {
	s.CustomPriorities = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDataDisks(v []*ModifyScalingConfigurationShrinkRequestDataDisks) *ModifyScalingConfigurationShrinkRequest {
	s.DataDisks = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDedicatedHostClusterId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDedicatedHostId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDeletionProtection(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetDeploymentSetId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHostName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HostName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHpcClusterId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHttpEndpoint(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetHttpTokens(v string) *ModifyScalingConfigurationShrinkRequest {
	s.HttpTokens = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageFamily(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetImageName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceDescription(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstancePatternInfos(v []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos) *ModifyScalingConfigurationShrinkRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypeOverrides(v []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypes(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInternetChargeType(v string) *ModifyScalingConfigurationShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInternetMaxBandwidthIn(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetInternetMaxBandwidthOut(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetIoOptimized(v string) *ModifyScalingConfigurationShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetKeyPairName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetMemory(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetNetworkInterfaces(v []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces) *ModifyScalingConfigurationShrinkRequest {
	s.NetworkInterfaces = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOverride(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.Override = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetOwnerId(v int64) *ModifyScalingConfigurationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPassword(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Password = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetPasswordInherit(v bool) *ModifyScalingConfigurationShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetRamRoleName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourceGroupId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourceOwnerAccount(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetResourcePoolOptions(v *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) *ModifyScalingConfigurationShrinkRequest {
	s.ResourcePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetScalingConfigurationId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetScalingConfigurationName(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSchedulerOptionsShrink(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SchedulerOptionsShrink = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityGroupId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSecurityOptions(v *ModifyScalingConfigurationShrinkRequestSecurityOptions) *ModifyScalingConfigurationShrinkRequest {
	s.SecurityOptions = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotDuration(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.SpotDuration = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotInterruptionBehavior(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotPriceLimits(v []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits) *ModifyScalingConfigurationShrinkRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSpotStrategy(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetStorageSetId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.StorageSetId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetStorageSetPartitionNumber(v int32) *ModifyScalingConfigurationShrinkRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetSystemDiskCategories(v []*string) *ModifyScalingConfigurationShrinkRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetTags(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetTenancy(v string) *ModifyScalingConfigurationShrinkRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetUserData(v string) *ModifyScalingConfigurationShrinkRequest {
	s.UserData = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) SetZoneId(v string) *ModifyScalingConfigurationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequest) Validate() error {
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
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
	if s.InstanceTypeOverrides != nil {
		for _, item := range s.InstanceTypeOverrides {
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
	return nil
}

type ModifyScalingConfigurationShrinkRequestImageOptions struct {
	// Specifies whether to use ecs-user to log on to an ECS instance created from the scaling configuration. For information about logon usernames, see [Manage the logon username of an instance](https://help.aliyun.com/document_detail/388447.html). Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestImageOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *ModifyScalingConfigurationShrinkRequestImageOptions) SetLoginAsNonRoot(v bool) *ModifyScalingConfigurationShrinkRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestPrivatePoolOptions struct {
	// The ID of the private pool. The ID of a private pool is the same as the ID of the elasticity assurance or capacity reservation for which the private pool is generated.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the private pool that you want to use to start ECS instances. A private pool is generated when an elasticity assurance or a capacity reservation takes effect. You can specify a private pool for Auto Scaling to start ECS instances. Valid values:
	//
	// 	- Open: open private pool. Auto Scaling selects a matching open private pool to start ECS instances. If no matching open private pools exist, the resources in the public pool are used. In this case, you do not need to specify PrivatePoolOptions.Id.
	//
	// 	- Target: specified private pool. Auto Scaling uses the resources in the specified private pool to start ECS instances. If the specified private pool does not exist, Auto Scaling cannot start ECS instances. If you set this parameter to Target, you must specify PrivatePoolOptions.Id.
	//
	// 	- None: no private pool. Auto Scaling does not use the resources of private pools to start ECS instances.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) SetId(v string) *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestSystemDisk struct {
	// The ID of the automatic snapshot policy that you want to apply to the system disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the Burst feature for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  If you set `SystemDisk.Category` to `cloud_auto`, you can specify this parameter.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the system disk. Valid values:
	//
	// 	- cloud: basic disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: Enterprise SSD (ESSD).
	//
	// 	- ephemeral_ssd: local SSD.
	//
	// If you specify SystemDisk.Category, you cannot specify `SystemDiskCategories`. If you do not specify SystemDisk.Category or `SystemDiskCategories`, the default value of SystemDisk.Category is used. The default value for non-I/O optimized instances of Generation I instance families is cloud. The default value for other instances is cloud_efficiency.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length. The description can contain letters but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test system disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http:// `or `https://`.
	//
	// Default value: null.
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The encryption algorithm of the system disk. Valid values:
	//
	// 	- AES-256
	//
	// 	- SM4-128
	//
	// Default value: AES-256.
	//
	// example:
	//
	// AES-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
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
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key that you want to use to encrypt the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level (PL) of the system disk that is an ESSD. Valid values:
	//
	// 	- PL0: An ESSD can provide up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can provide up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can provide up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD can provide up to 1,000,000 random read/write IOPS.
	//
	// >  For more information about how to select ESSD PLs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The IOPS metric that is preconfigured for the system disk.
	//
	// > IOPS measures the number of read and write operations that an EBS device can process per second.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// 	- Basic disk: 20 to 500.
	//
	// 	- ESSD: Valid values vary based on the performance level of the ESSD.
	//
	//     	- PL0 ESSD: 1 to 2048.
	//
	//     	- PL1 ESSD: 20 to 2048.
	//
	//     	- PL2 ESSD: 461 to 2048.
	//
	//     	- PL3 ESSD: 1261 to 2048.
	//
	// 	- ESSD AutoPL disk: 1 to 2048.
	//
	// 	- Other disk categories: 20 to 2048.
	//
	// The value of this parameter must be at least 1 and greater than or equal to the image size.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetBurstingEnabled(v bool) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetCategory(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetDescription(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetDiskName(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetEncryptAlgorithm(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetEncrypted(v bool) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetKMSKeyId(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetProvisionedIops(v int64) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) SetSize(v int32) *ModifyScalingConfigurationShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestCustomPriorities struct {
	// The ECS instance type.
	//
	// >  The ECS instance type must be included in the instance types specified in the scaling configuration.
	//
	// example:
	//
	// ecs.c6a.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The vSwitch ID.
	//
	// >  The vSwitch must be included in the vSwitch list of the scaling group.
	//
	// example:
	//
	// vsw-bp14zolna43z266bq****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestCustomPriorities) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestCustomPriorities) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestCustomPriorities) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationShrinkRequestCustomPriorities) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ModifyScalingConfigurationShrinkRequestCustomPriorities) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestCustomPriorities {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestCustomPriorities) SetVswitchId(v string) *ModifyScalingConfigurationShrinkRequestCustomPriorities {
	s.VswitchId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestCustomPriorities) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestDataDisks struct {
	// The ID of the automatic snapshot policy that you want to apply to the data disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the Burst feature for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  If you set `SystemDisk.Category` to `cloud_auto`, you can specify this parameter.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The categories of data disks. Valid values:
	//
	// 	- cloud: basic disk. The DeleteWithInstance attribute of a basic disk created along with each ECS instance is set to true.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// >  If you specify this parameter, you cannot specify `DataDisk.Category`.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data disk. Valid values:
	//
	// 	- cloud: basic disk. The DeleteWithInstance attribute of a basic disk created along with each ECS instance is set to true.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- ephemeral_ssd: local SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// If you specify this parameter, you cannot specify `DataDisk.Categories`. If you leave this parameter and `DataDisk.Categories` empty at the same time, the default value of this parameter is used.
	//
	// 	- For I/O optimized instances, the default value is cloud_efficiency.
	//
	// 	- For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk if the instance to which the data disk is attached is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// If you set DataDisk.Category to cloud, cloud_efficiency, cloud_ssd, cloud_essd, or cloud_auto, you can specify this parameter. If you specify this parameter for data disks of other categories, an error is returned.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test data disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount target of the data disk. If you do not specify this parameter, the system automatically assigns a mount target when Auto Scaling creates an ECS instance. Valid values: /dev/xvdb to /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdd
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key that you want to apply to the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The PL of the data disk that is an ESSD. Valid values:
	//
	// 	- PL0: An ESSD can provide up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can provide up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can provide up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD can provide up to 1,000,000 random read/write IOPS.
	//
	// >  For more information about how to select ESSD PLs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
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
	// 	- 5 to 2000 if you set DataDisk.Category to cloud.
	//
	// 	- 20 to 32768 if you set DataDisk.Category to cloud_efficiency.
	//
	// 	- 20 to 32768 if you set DataDisk.Category to cloud_ssd.
	//
	// 	- 20 to 32768 if you set DataDisk.Category to cloud_essd.
	//
	// 	- 5 to 800 if you set DataDisk.Category to ephemeral_ssd.
	//
	// Set Size to a value that is greater than or equal to the size of the snapshot specified by SnapshotId.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot that you want to use to create data disks. If you specify this parameter, DataDisk.Size is ignored. The size of the data disk created by using the snapshot is the same as the size of the snapshot.
	//
	// If the snapshot was created on or before July 15, 2013, the API request is rejected and the InvalidSnapshot.TooOld message is returned.
	//
	// example:
	//
	// s-snapshot****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestDataDisks) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestDataDisks) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetCategories() []*string {
	return s.Categories
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetCategory() *string {
	return s.Category
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetDescription() *string {
	return s.Description
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetDevice() *string {
	return s.Device
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetBurstingEnabled(v bool) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetCategories(v []*string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Categories = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetCategory(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDeleteWithInstance(v bool) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDescription(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDevice(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Device = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetDiskName(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetEncrypted(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetKMSKeyId(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetPerformanceLevel(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetProvisionedIops(v int64) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetSize(v int32) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) SetSnapshotId(v string) *ModifyScalingConfigurationShrinkRequestDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestDataDisks) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestInstancePatternInfos struct {
	// The architecture types of the instance types. Valid values:
	//
	// 	- X86: x86.
	//
	// 	- Heterogeneous: heterogeneous computing, such as GPU-accelerated or FPGA-accelerated.
	//
	// 	- BareMetal: ECS Bare Metal Instance.
	//
	// 	- Arm: Arm.
	//
	// By default, all values are selected.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether to include burstable instance types. Valid values:
	//
	// 	- Exclude: excludes burstable instance types.
	//
	// 	- Include: includes burstable instance types.
	//
	// 	- Required: includes only burstable instance types.
	//
	// Default value: Include.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The number of vCPUs per instance type in intelligent configuration mode. You can specify this parameter to filter the available instance types. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// Before you specify this parameter, take note of the following items:
	//
	// 	- You can specify InstancePatternInfo only if your scaling group resides in a virtual private cloud (VPC).
	//
	// 	- If you specify InstancePatternInfo, you must also specify InstancePatternInfo.Cores and InstancePatternInfo.Memory.
	//
	// 	- Auto Scaling preferentially uses the instance type specified by InstanceType or InstanceTypes to create instances. If the specified instance type does not have sufficient inventory, Auto Scaling selects the lowest-priced instance type specified by InstancePatternInfo to create instances.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architectures of the instance types. Valid values:
	//
	// >  You can specify up to two CPU architectures.
	//
	// 	- x86
	//
	// 	- Arm
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The instance types that you want to exclude. You can use an asterisk (\\*) as a wildcard character to exclude an instance type or an instance family. Examples:
	//
	// 	- ecs.c6.large: excludes the ecs.c6.large instance type.
	//
	// 	- ecs.c6.\\*: excludes the c6 instance family.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The GPU models.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The categories of the instance types. Valid values:
	//
	// 	- General-purpose: general-purpose instance type.
	//
	// 	- Compute-optimized: compute-optimized instance type.
	//
	// 	- Memory-optimized: memory-optimized instance type.
	//
	// 	- Big data: big data instance type.
	//
	// 	- Local SSDs: instance type that uses local SSDs.
	//
	// 	- High Clock Speed: instance type that has a high clock speed.
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
	// The level of the instance family. You can specify this parameter to obtain the available instance types. This parameter takes effect only if you set `CostOptimization` to true. Valid values:
	//
	// 	- EntryLevel: entry-level (shared instance types). Instance types of this level are the most cost-effective but may not ensure stable computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// 	- EnterpriseLevel: enterprise-level. Instance types of this level provide stable performance and dedicated resources and are suitable for business scenarios that require high stability. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// 	- CreditEntryLevel: credit entry-level (burstable instance types). CPU credits are used to ensure computing performance. Instance types of this level are suitable for scenarios in which the CPU utilization is low but may fluctuate in specific cases. For more information, see [Overview](https://help.aliyun.com/document_detail/59977.html) of burstable instances.
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families that you want to specify. You can specify up to 10 instance families in each call.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum hourly price of pay-as-you-go or preemptible instances in intelligent configuration mode. You can specify this parameter to obtain the available instance types.
	//
	// >  If you set SpotStrategy to SpotWithPriceLimit, you must specify this parameter. In other cases, this parameter is optional.
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
	// The memory size per instance type in intelligent configuration mode. Unit: GiB. You can specify this parameter to filter the available instance types.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The baseline vCPU computing performance (overall baseline performance of all vCPUs) of each t5 or t6 burstable instance.
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
	// The initial vCPU credits of each t5 or t6 burstable instance.
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
	// The processor models of the instance types. You can specify up to 10 processor models.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
}

func (s ModifyScalingConfigurationShrinkRequestInstancePatternInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetArchitectures() []*string {
	return s.Architectures
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetCores() *int32 {
	return s.Cores
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMaximumGpuAmount() *int32 {
	return s.MaximumGpuAmount
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMemory() *float32 {
	return s.Memory
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumGpuAmount() *int32 {
	return s.MinimumGpuAmount
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetArchitectures(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetBurstablePerformance(v string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetCores(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetCpuArchitectures(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.CpuArchitectures = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetGpuSpecs(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.GpuSpecs = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceCategories(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceCategories = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceTypeFamilies(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceTypeFamilies = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMaxPrice(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMaximumCpuCoreCount(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMaximumGpuAmount(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaximumGpuAmount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMaximumMemorySize(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaximumMemorySize = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMemory(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumBaselineCredit(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumCpuCoreCount(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumEniIpv6AddressQuantity(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumEniPrivateIpAddressQuantity(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumEniQuantity(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumEniQuantity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumGpuAmount(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumGpuAmount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumInitialCredit(v int32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumInitialCredit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumMemorySize(v float32) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumMemorySize = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) SetPhysicalProcessorModels(v []*string) *ModifyScalingConfigurationShrinkRequestInstancePatternInfos {
	s.PhysicalProcessorModels = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstancePatternInfos) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides struct {
	// The instance type. If you want to specify the capacity of instance types in the scaling configuration, specify InstanceType and WeightedCapacity at the same time.
	//
	// You can use InstanceType to specify multiple instance types and WeightedCapacity to specify the weights of the instance types.
	//
	// > If you specify InstanceType, you cannot specify InstanceTypes.
	//
	// You can use InstanceType to specify only instance types that are available for purchase.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The weight of the instance type. The weight specifies the capacity of an instance of the specified instance type in the scaling group. If you want Auto Scaling to scale instances in the scaling group based on the weighted capacity of the instances, specify WeightedCapacity after you specify InstanceType.
	//
	// A higher weight specifies that a smaller number of instances of the specified instance type are required to meet the expected capacity requirement.
	//
	// Performance metrics, such as the number of vCPUs and the memory size of each instance type, may vary. You can specify different weights for different instance types based on your business requirements.
	//
	// Example:
	//
	// 	- Current capacity: 0
	//
	// 	- Expected capacity: 6
	//
	// 	- Capacity of ecs.c5.xlarge: 4
	//
	// To meet the expected capacity requirement, Auto Scaling must create and add two ecs.c5.xlarge instances.
	//
	// > The capacity of the scaling group cannot exceed the sum of the maximum number of instances that is specified by MaxSize and the maximum weight of the instance types.
	//
	// Valid values of WeightedCapacity: 1 to 500.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestNetworkInterfaces struct {
	// The ENI type. If you specify this parameter, you must use NetworkInterfaces to specify a primary ENI. In addition, you cannot specify SecurityGroupId or SecurityGroupIds. Valid values:
	//
	// 	- Primary: the primary ENI.
	//
	// 	- Secondary: the secondary ENI.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses that you want to allocate to the primary ENI. Before you specify this parameter, take note of the following items:
	//
	// This parameter takes effect only if you set NetworkInterface.InstanceType to Primary. If you set NetworkInterface.InstanceType to Secondary or leave it empty, you cannot specify this parameter.
	//
	// After you specify this parameter, you can no longer specify Ipv6AddressCount.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication mode of the ENI. Valid values:
	//
	// 	- Standard: uses the TCP communication mode.
	//
	// 	- HighPerformance: uses the remote direct memory access (RDMA) communication mode with Elastic RDMA Interface (ERI) enabled.
	//
	// Default value: Standard.
	//
	// >  The number of ERIs on an instance cannot exceed the maximum number of ERIs supported by the instance type. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The IDs of the security groups to which you want to assign the ENI.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s ModifyScalingConfigurationShrinkRequestNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetNetworkInterfaceTrafficMode(v string) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestResourcePoolOptions struct {
	// The IDs of private pools. The ID of a private pool is the same as that of the elasticity assurance or capacity reservation for which the private pool is generated. You can specify the IDs of only targeted private pools for this parameter.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// The resource pool used for instance creation, which can be the public pool or a private pool associated with any active elasticity assurance or capacity reservation. Valid values:
	//
	// 	- PrivatePoolFirst: prioritizes private pools. When this option is set along with ResourcePoolOptions.PrivatePoolIds, the specified private pools are used first. If you leave ResourcePoolOptions.PrivatePoolIds empty or if the specified private pools lack sufficient capacity, the system will automatically use available open private pools instead. If no matching private pools are available, the system defaults to the public pool.
	//
	// 	- PrivatePoolOnly: uses only private pools. If you use this value, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pools lack sufficient capacity, instance creation will fail.
	//
	// 	- None: uses no resource pools.
	//
	// Default value: None.
	//
	// example:
	//
	// PrivatePoolFirst
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *ModifyScalingConfigurationShrinkRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) SetStrategy(v string) *ModifyScalingConfigurationShrinkRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestSecurityOptions struct {
	// The confidential computing mode. Valid values:
	//
	// 	- Enclave: An enclave-based confidential computing environment is built on the instance. For more information, see [Build a confidential computing environment by using Enclave](https://help.aliyun.com/document_detail/203433.html).
	//
	// 	- TDX: A Trust Domain Extensions (TDX) confidential computing environment is built on the instance. For more information, see [Build a TDX confidential computing environment](https://help.aliyun.com/document_detail/479090.html).
	//
	// example:
	//
	// TDX
	ConfidentialComputingMode *string `json:"ConfidentialComputingMode,omitempty" xml:"ConfidentialComputingMode,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *ModifyScalingConfigurationShrinkRequestSecurityOptions) SetConfidentialComputingMode(v string) *ModifyScalingConfigurationShrinkRequestSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestSpotPriceLimits struct {
	// The instance type of the preemptible instance. This parameter takes effect only if you set SpotStrategy to SpotWithPriceLimit.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The price limit of the preemptible instance. This parameter takes effect only if you set SpotStrategy to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.125
	PriceLimit *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestSpotPriceLimits) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) GetPriceLimit() *float32 {
	return s.PriceLimit
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) SetInstanceType(v string) *ModifyScalingConfigurationShrinkRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) SetPriceLimit(v float32) *ModifyScalingConfigurationShrinkRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestSpotPriceLimits) Validate() error {
	return dara.Validate(s)
}
