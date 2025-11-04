// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageOptions(v *ModifyScalingConfigurationRequestImageOptions) *ModifyScalingConfigurationRequest
	GetImageOptions() *ModifyScalingConfigurationRequestImageOptions
	SetPrivatePoolOptions(v *ModifyScalingConfigurationRequestPrivatePoolOptions) *ModifyScalingConfigurationRequest
	GetPrivatePoolOptions() *ModifyScalingConfigurationRequestPrivatePoolOptions
	SetSystemDisk(v *ModifyScalingConfigurationRequestSystemDisk) *ModifyScalingConfigurationRequest
	GetSystemDisk() *ModifyScalingConfigurationRequestSystemDisk
	SetAffinity(v string) *ModifyScalingConfigurationRequest
	GetAffinity() *string
	SetCpu(v int32) *ModifyScalingConfigurationRequest
	GetCpu() *int32
	SetCreditSpecification(v string) *ModifyScalingConfigurationRequest
	GetCreditSpecification() *string
	SetCustomPriorities(v []*ModifyScalingConfigurationRequestCustomPriorities) *ModifyScalingConfigurationRequest
	GetCustomPriorities() []*ModifyScalingConfigurationRequestCustomPriorities
	SetDataDisks(v []*ModifyScalingConfigurationRequestDataDisks) *ModifyScalingConfigurationRequest
	GetDataDisks() []*ModifyScalingConfigurationRequestDataDisks
	SetDedicatedHostClusterId(v string) *ModifyScalingConfigurationRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *ModifyScalingConfigurationRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *ModifyScalingConfigurationRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *ModifyScalingConfigurationRequest
	GetDeploymentSetId() *string
	SetHostName(v string) *ModifyScalingConfigurationRequest
	GetHostName() *string
	SetHpcClusterId(v string) *ModifyScalingConfigurationRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *ModifyScalingConfigurationRequest
	GetHttpEndpoint() *string
	SetHttpTokens(v string) *ModifyScalingConfigurationRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *ModifyScalingConfigurationRequest
	GetImageFamily() *string
	SetImageId(v string) *ModifyScalingConfigurationRequest
	GetImageId() *string
	SetImageName(v string) *ModifyScalingConfigurationRequest
	GetImageName() *string
	SetInstanceDescription(v string) *ModifyScalingConfigurationRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *ModifyScalingConfigurationRequest
	GetInstanceName() *string
	SetInstancePatternInfos(v []*ModifyScalingConfigurationRequestInstancePatternInfos) *ModifyScalingConfigurationRequest
	GetInstancePatternInfos() []*ModifyScalingConfigurationRequestInstancePatternInfos
	SetInstanceTypeOverrides(v []*ModifyScalingConfigurationRequestInstanceTypeOverrides) *ModifyScalingConfigurationRequest
	GetInstanceTypeOverrides() []*ModifyScalingConfigurationRequestInstanceTypeOverrides
	SetInstanceTypes(v []*string) *ModifyScalingConfigurationRequest
	GetInstanceTypes() []*string
	SetInternetChargeType(v string) *ModifyScalingConfigurationRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *ModifyScalingConfigurationRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *ModifyScalingConfigurationRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *ModifyScalingConfigurationRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *ModifyScalingConfigurationRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *ModifyScalingConfigurationRequest
	GetKeyPairName() *string
	SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationRequest
	GetLoadBalancerWeight() *int32
	SetMemory(v int32) *ModifyScalingConfigurationRequest
	GetMemory() *int32
	SetNetworkInterfaces(v []*ModifyScalingConfigurationRequestNetworkInterfaces) *ModifyScalingConfigurationRequest
	GetNetworkInterfaces() []*ModifyScalingConfigurationRequestNetworkInterfaces
	SetOverride(v bool) *ModifyScalingConfigurationRequest
	GetOverride() *bool
	SetOwnerAccount(v string) *ModifyScalingConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyScalingConfigurationRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifyScalingConfigurationRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *ModifyScalingConfigurationRequest
	GetPasswordInherit() *bool
	SetRamRoleName(v string) *ModifyScalingConfigurationRequest
	GetRamRoleName() *string
	SetResourceGroupId(v string) *ModifyScalingConfigurationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetResourcePoolOptions(v *ModifyScalingConfigurationRequestResourcePoolOptions) *ModifyScalingConfigurationRequest
	GetResourcePoolOptions() *ModifyScalingConfigurationRequestResourcePoolOptions
	SetScalingConfigurationId(v string) *ModifyScalingConfigurationRequest
	GetScalingConfigurationId() *string
	SetScalingConfigurationName(v string) *ModifyScalingConfigurationRequest
	GetScalingConfigurationName() *string
	SetSchedulerOptions(v map[string]interface{}) *ModifyScalingConfigurationRequest
	GetSchedulerOptions() map[string]interface{}
	SetSecurityGroupId(v string) *ModifyScalingConfigurationRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *ModifyScalingConfigurationRequestSecurityOptions) *ModifyScalingConfigurationRequest
	GetSecurityOptions() *ModifyScalingConfigurationRequestSecurityOptions
	SetSpotDuration(v int32) *ModifyScalingConfigurationRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *ModifyScalingConfigurationRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimits(v []*ModifyScalingConfigurationRequestSpotPriceLimits) *ModifyScalingConfigurationRequest
	GetSpotPriceLimits() []*ModifyScalingConfigurationRequestSpotPriceLimits
	SetSpotStrategy(v string) *ModifyScalingConfigurationRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *ModifyScalingConfigurationRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *ModifyScalingConfigurationRequest
	GetStorageSetPartitionNumber() *int32
	SetSystemDiskCategories(v []*string) *ModifyScalingConfigurationRequest
	GetSystemDiskCategories() []*string
	SetTags(v string) *ModifyScalingConfigurationRequest
	GetTags() *string
	SetTenancy(v string) *ModifyScalingConfigurationRequest
	GetTenancy() *string
	SetUserData(v string) *ModifyScalingConfigurationRequest
	GetUserData() *string
	SetZoneId(v string) *ModifyScalingConfigurationRequest
	GetZoneId() *string
}

type ModifyScalingConfigurationRequest struct {
	ImageOptions       *ModifyScalingConfigurationRequestImageOptions       `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *ModifyScalingConfigurationRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *ModifyScalingConfigurationRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
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
	CustomPriorities []*ModifyScalingConfigurationRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The data disks.
	DataDisks []*ModifyScalingConfigurationRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
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
	InstancePatternInfos []*ModifyScalingConfigurationRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance types.
	InstanceTypeOverrides []*ModifyScalingConfigurationRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
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
	NetworkInterfaces []*ModifyScalingConfigurationRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
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
	ResourcePoolOptions *ModifyScalingConfigurationRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
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
	SchedulerOptions map[string]interface{} `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// The ID of the security group with which ECS instances are associated. The ECS instances that are associated with the same security group can access each other.
	//
	// example:
	//
	// sg-F876F****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups.
	SecurityGroupIds []*string                                         `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *ModifyScalingConfigurationRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
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
	SpotPriceLimits []*ModifyScalingConfigurationRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
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

func (s ModifyScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequest) GetImageOptions() *ModifyScalingConfigurationRequestImageOptions {
	return s.ImageOptions
}

func (s *ModifyScalingConfigurationRequest) GetPrivatePoolOptions() *ModifyScalingConfigurationRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyScalingConfigurationRequest) GetSystemDisk() *ModifyScalingConfigurationRequestSystemDisk {
	return s.SystemDisk
}

func (s *ModifyScalingConfigurationRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *ModifyScalingConfigurationRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyScalingConfigurationRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *ModifyScalingConfigurationRequest) GetCustomPriorities() []*ModifyScalingConfigurationRequestCustomPriorities {
	return s.CustomPriorities
}

func (s *ModifyScalingConfigurationRequest) GetDataDisks() []*ModifyScalingConfigurationRequestDataDisks {
	return s.DataDisks
}

func (s *ModifyScalingConfigurationRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *ModifyScalingConfigurationRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ModifyScalingConfigurationRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyScalingConfigurationRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ModifyScalingConfigurationRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyScalingConfigurationRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *ModifyScalingConfigurationRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *ModifyScalingConfigurationRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *ModifyScalingConfigurationRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *ModifyScalingConfigurationRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyScalingConfigurationRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ModifyScalingConfigurationRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *ModifyScalingConfigurationRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyScalingConfigurationRequest) GetInstancePatternInfos() []*ModifyScalingConfigurationRequestInstancePatternInfos {
	return s.InstancePatternInfos
}

func (s *ModifyScalingConfigurationRequest) GetInstanceTypeOverrides() []*ModifyScalingConfigurationRequestInstanceTypeOverrides {
	return s.InstanceTypeOverrides
}

func (s *ModifyScalingConfigurationRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *ModifyScalingConfigurationRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyScalingConfigurationRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *ModifyScalingConfigurationRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyScalingConfigurationRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *ModifyScalingConfigurationRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *ModifyScalingConfigurationRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ModifyScalingConfigurationRequest) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *ModifyScalingConfigurationRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *ModifyScalingConfigurationRequest) GetNetworkInterfaces() []*ModifyScalingConfigurationRequestNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *ModifyScalingConfigurationRequest) GetOverride() *bool {
	return s.Override
}

func (s *ModifyScalingConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyScalingConfigurationRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyScalingConfigurationRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *ModifyScalingConfigurationRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *ModifyScalingConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyScalingConfigurationRequest) GetResourcePoolOptions() *ModifyScalingConfigurationRequestResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *ModifyScalingConfigurationRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *ModifyScalingConfigurationRequest) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *ModifyScalingConfigurationRequest) GetSchedulerOptions() map[string]interface{} {
	return s.SchedulerOptions
}

func (s *ModifyScalingConfigurationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifyScalingConfigurationRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyScalingConfigurationRequest) GetSecurityOptions() *ModifyScalingConfigurationRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *ModifyScalingConfigurationRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *ModifyScalingConfigurationRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *ModifyScalingConfigurationRequest) GetSpotPriceLimits() []*ModifyScalingConfigurationRequestSpotPriceLimits {
	return s.SpotPriceLimits
}

func (s *ModifyScalingConfigurationRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *ModifyScalingConfigurationRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *ModifyScalingConfigurationRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *ModifyScalingConfigurationRequest) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *ModifyScalingConfigurationRequest) GetTags() *string {
	return s.Tags
}

func (s *ModifyScalingConfigurationRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *ModifyScalingConfigurationRequest) GetUserData() *string {
	return s.UserData
}

func (s *ModifyScalingConfigurationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyScalingConfigurationRequest) SetImageOptions(v *ModifyScalingConfigurationRequestImageOptions) *ModifyScalingConfigurationRequest {
	s.ImageOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetPrivatePoolOptions(v *ModifyScalingConfigurationRequestPrivatePoolOptions) *ModifyScalingConfigurationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSystemDisk(v *ModifyScalingConfigurationRequestSystemDisk) *ModifyScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetAffinity(v string) *ModifyScalingConfigurationRequest {
	s.Affinity = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetCpu(v int32) *ModifyScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetCreditSpecification(v string) *ModifyScalingConfigurationRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetCustomPriorities(v []*ModifyScalingConfigurationRequestCustomPriorities) *ModifyScalingConfigurationRequest {
	s.CustomPriorities = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDataDisks(v []*ModifyScalingConfigurationRequestDataDisks) *ModifyScalingConfigurationRequest {
	s.DataDisks = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDedicatedHostClusterId(v string) *ModifyScalingConfigurationRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDedicatedHostId(v string) *ModifyScalingConfigurationRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDeletionProtection(v bool) *ModifyScalingConfigurationRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetDeploymentSetId(v string) *ModifyScalingConfigurationRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHostName(v string) *ModifyScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHpcClusterId(v string) *ModifyScalingConfigurationRequest {
	s.HpcClusterId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHttpEndpoint(v string) *ModifyScalingConfigurationRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetHttpTokens(v string) *ModifyScalingConfigurationRequest {
	s.HttpTokens = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageFamily(v string) *ModifyScalingConfigurationRequest {
	s.ImageFamily = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageId(v string) *ModifyScalingConfigurationRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetImageName(v string) *ModifyScalingConfigurationRequest {
	s.ImageName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceDescription(v string) *ModifyScalingConfigurationRequest {
	s.InstanceDescription = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceName(v string) *ModifyScalingConfigurationRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstancePatternInfos(v []*ModifyScalingConfigurationRequestInstancePatternInfos) *ModifyScalingConfigurationRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceTypeOverrides(v []*ModifyScalingConfigurationRequestInstanceTypeOverrides) *ModifyScalingConfigurationRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInstanceTypes(v []*string) *ModifyScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInternetChargeType(v string) *ModifyScalingConfigurationRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInternetMaxBandwidthIn(v int32) *ModifyScalingConfigurationRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetInternetMaxBandwidthOut(v int32) *ModifyScalingConfigurationRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetIoOptimized(v string) *ModifyScalingConfigurationRequest {
	s.IoOptimized = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetKeyPairName(v string) *ModifyScalingConfigurationRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *ModifyScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetMemory(v int32) *ModifyScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetNetworkInterfaces(v []*ModifyScalingConfigurationRequestNetworkInterfaces) *ModifyScalingConfigurationRequest {
	s.NetworkInterfaces = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOverride(v bool) *ModifyScalingConfigurationRequest {
	s.Override = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOwnerAccount(v string) *ModifyScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetOwnerId(v int64) *ModifyScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetPassword(v string) *ModifyScalingConfigurationRequest {
	s.Password = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetPasswordInherit(v bool) *ModifyScalingConfigurationRequest {
	s.PasswordInherit = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetRamRoleName(v string) *ModifyScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourceGroupId(v string) *ModifyScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetResourcePoolOptions(v *ModifyScalingConfigurationRequestResourcePoolOptions) *ModifyScalingConfigurationRequest {
	s.ResourcePoolOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetScalingConfigurationId(v string) *ModifyScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetScalingConfigurationName(v string) *ModifyScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSchedulerOptions(v map[string]interface{}) *ModifyScalingConfigurationRequest {
	s.SchedulerOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityGroupId(v string) *ModifyScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSecurityOptions(v *ModifyScalingConfigurationRequestSecurityOptions) *ModifyScalingConfigurationRequest {
	s.SecurityOptions = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotDuration(v int32) *ModifyScalingConfigurationRequest {
	s.SpotDuration = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotInterruptionBehavior(v string) *ModifyScalingConfigurationRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotPriceLimits(v []*ModifyScalingConfigurationRequestSpotPriceLimits) *ModifyScalingConfigurationRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSpotStrategy(v string) *ModifyScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetStorageSetId(v string) *ModifyScalingConfigurationRequest {
	s.StorageSetId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetStorageSetPartitionNumber(v int32) *ModifyScalingConfigurationRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetSystemDiskCategories(v []*string) *ModifyScalingConfigurationRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetTags(v string) *ModifyScalingConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetTenancy(v string) *ModifyScalingConfigurationRequest {
	s.Tenancy = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetUserData(v string) *ModifyScalingConfigurationRequest {
	s.UserData = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) SetZoneId(v string) *ModifyScalingConfigurationRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyScalingConfigurationRequest) Validate() error {
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

type ModifyScalingConfigurationRequestImageOptions struct {
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

func (s ModifyScalingConfigurationRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestImageOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *ModifyScalingConfigurationRequestImageOptions) SetLoginAsNonRoot(v bool) *ModifyScalingConfigurationRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *ModifyScalingConfigurationRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestPrivatePoolOptions struct {
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

func (s ModifyScalingConfigurationRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) SetId(v string) *ModifyScalingConfigurationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyScalingConfigurationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyScalingConfigurationRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestSystemDisk struct {
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

func (s ModifyScalingConfigurationRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *ModifyScalingConfigurationRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetBurstingEnabled(v bool) *ModifyScalingConfigurationRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetCategory(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetDescription(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetDiskName(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetEncryptAlgorithm(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetEncrypted(v bool) *ModifyScalingConfigurationRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetKMSKeyId(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetPerformanceLevel(v string) *ModifyScalingConfigurationRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetProvisionedIops(v int64) *ModifyScalingConfigurationRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) SetSize(v int32) *ModifyScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestCustomPriorities struct {
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

func (s ModifyScalingConfigurationRequestCustomPriorities) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestCustomPriorities) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestCustomPriorities) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationRequestCustomPriorities) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ModifyScalingConfigurationRequestCustomPriorities) SetInstanceType(v string) *ModifyScalingConfigurationRequestCustomPriorities {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestCustomPriorities) SetVswitchId(v string) *ModifyScalingConfigurationRequestCustomPriorities {
	s.VswitchId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestCustomPriorities) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestDataDisks struct {
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

func (s ModifyScalingConfigurationRequestDataDisks) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestDataDisks) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetCategories() []*string {
	return s.Categories
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetCategory() *string {
	return s.Category
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetDescription() *string {
	return s.Description
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetDevice() *string {
	return s.Device
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *ModifyScalingConfigurationRequestDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetAutoSnapshotPolicyId(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetBurstingEnabled(v bool) *ModifyScalingConfigurationRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetCategories(v []*string) *ModifyScalingConfigurationRequestDataDisks {
	s.Categories = v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetCategory(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Category = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDeleteWithInstance(v bool) *ModifyScalingConfigurationRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDescription(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Description = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDevice(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Device = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetDiskName(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetEncrypted(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetKMSKeyId(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetPerformanceLevel(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetProvisionedIops(v int64) *ModifyScalingConfigurationRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetSize(v int32) *ModifyScalingConfigurationRequestDataDisks {
	s.Size = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) SetSnapshotId(v string) *ModifyScalingConfigurationRequestDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *ModifyScalingConfigurationRequestDataDisks) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestInstancePatternInfos struct {
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

func (s ModifyScalingConfigurationRequestInstancePatternInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetArchitectures() []*string {
	return s.Architectures
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetCores() *int32 {
	return s.Cores
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMaximumGpuAmount() *int32 {
	return s.MaximumGpuAmount
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMemory() *float32 {
	return s.Memory
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumGpuAmount() *int32 {
	return s.MinimumGpuAmount
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetArchitectures(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetBurstablePerformance(v string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetCores(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetCpuArchitectures(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.CpuArchitectures = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetGpuSpecs(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.GpuSpecs = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetInstanceCategories(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.InstanceCategories = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetInstanceTypeFamilies(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.InstanceTypeFamilies = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMaxPrice(v float32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMaximumCpuCoreCount(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMaximumGpuAmount(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MaximumGpuAmount = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMaximumMemorySize(v float32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MaximumMemorySize = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMemory(v float32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumBaselineCredit(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumCpuCoreCount(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumEniIpv6AddressQuantity(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumEniPrivateIpAddressQuantity(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumEniQuantity(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumEniQuantity = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumGpuAmount(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumGpuAmount = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumInitialCredit(v int32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumInitialCredit = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetMinimumMemorySize(v float32) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.MinimumMemorySize = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) SetPhysicalProcessorModels(v []*string) *ModifyScalingConfigurationRequestInstancePatternInfos {
	s.PhysicalProcessorModels = v
	return s
}

func (s *ModifyScalingConfigurationRequestInstancePatternInfos) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestInstanceTypeOverrides struct {
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

func (s ModifyScalingConfigurationRequestInstanceTypeOverrides) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) SetInstanceType(v string) *ModifyScalingConfigurationRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *ModifyScalingConfigurationRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *ModifyScalingConfigurationRequestInstanceTypeOverrides) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestNetworkInterfaces struct {
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

func (s ModifyScalingConfigurationRequestNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) SetInstanceType(v string) *ModifyScalingConfigurationRequestNetworkInterfaces {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationRequestNetworkInterfaces {
	s.Ipv6AddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) SetNetworkInterfaceTrafficMode(v string) *ModifyScalingConfigurationRequestNetworkInterfaces {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationRequestNetworkInterfaces {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationRequestNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestResourcePoolOptions struct {
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

func (s ModifyScalingConfigurationRequestResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *ModifyScalingConfigurationRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *ModifyScalingConfigurationRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *ModifyScalingConfigurationRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *ModifyScalingConfigurationRequestResourcePoolOptions) SetStrategy(v string) *ModifyScalingConfigurationRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *ModifyScalingConfigurationRequestResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestSecurityOptions struct {
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

func (s ModifyScalingConfigurationRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *ModifyScalingConfigurationRequestSecurityOptions) SetConfidentialComputingMode(v string) *ModifyScalingConfigurationRequestSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationRequestSpotPriceLimits struct {
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

func (s ModifyScalingConfigurationRequestSpotPriceLimits) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) GetPriceLimit() *float32 {
	return s.PriceLimit
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) SetInstanceType(v string) *ModifyScalingConfigurationRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) SetPriceLimit(v float32) *ModifyScalingConfigurationRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

func (s *ModifyScalingConfigurationRequestSpotPriceLimits) Validate() error {
	return dara.Validate(s)
}
