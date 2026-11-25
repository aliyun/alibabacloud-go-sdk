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
	SetInstanceTypeCandidateOptions(v *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) *ModifyScalingConfigurationShrinkRequest
	GetInstanceTypeCandidateOptions() *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions
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
	SetSecondaryEniVSwitchConstraint(v string) *ModifyScalingConfigurationShrinkRequest
	GetSecondaryEniVSwitchConstraint() *string
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
	// Specifies whether the instance on a dedicated host is associated with the dedicated host. Valid values:
	//
	// - default: The instance is not associated with the dedicated host. When an instance that has the economical mode enabled is restarted, the instance may be placed on another dedicated host in the automatic deployment resource pool if the original dedicated host has insufficient resources.
	//
	// - host: The instance is associated with the dedicated host. When an instance that has the economical mode enabled is restarted, the instance remains on the original dedicated host. If the original dedicated host has insufficient resources, the instance fails to restart.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of vCPUs. Unit: cores.
	//
	// You can specify both CPU and Memory to define the range of instance types. For example, CPU=2 and Memory=16 defines all instance types with 2 vCPUs and 16 GiB of memory. Auto Scaling determines the set of active instance types based on factors such as I/O optimization and zones, and then creates instances of the lowest-priced instance type based on sorting by price.
	//
	// > This configuration takes effect only when the cost optimization mode is enabled and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Modifies the running mode of a burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information, see the performance constrained mode section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information, see the unlimited mode section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The custom priorities for ECS instance type + vSwitch combinations.
	//
	// 	Notice: This parameter takes effect only when the scaling policy of the scaling group is set to priority-based policy.
	//
	// When an instance cannot be created based on the instance type and vSwitch combination with a higher priority, Auto Scaling automatically selects the next-priority instance type and vSwitch combination to create the instance.
	//
	// > If you specify custom priorities for only some instance type + vSwitch combinations, the unspecified combinations have lower priorities than the specified ones. The unspecified combinations are still prioritized based on the vSwitch order of the scaling group and the instance type order of the scaling configuration.
	//
	// > - For example, if the vSwitch order of the scaling group is vsw1 and vsw2, the instance type order of the scaling configuration is type1 and type2, and the custom priority order is ["vsw2+type2", "vsw1+type2"], the final priority is: "vsw2+type2" > "vsw1+type2" > "vsw1+type1" > "vsw2+type1".
	CustomPriorities []*ModifyScalingConfigurationShrinkRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The collection of data cloud disk information.
	DataDisks []*ModifyScalingConfigurationShrinkRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-zm04u8r3lohsq****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// Specifies whether to create the ECS instance on a dedicated host. Because dedicated hosts do not support the creation of spot instances, if you specify the DedicatedHostId parameter, the SpotStrategy and SpotPriceLimit settings in the request are automatically ignored.
	//
	// You can call the DescribeDedicatedHosts operation to query the list of dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection property of the instance. Specifies whether the instance can be directly released from the ECS console or by calling the DeleteInstance operation, to prevent accidental deletion. Valid values:
	//
	// - true: Enables release protection. The instance cannot be directly released from the ECS console or by calling the DeleteInstance operation.
	//
	// - false: Disables release protection. The instance can be directly released from the ECS console or by calling the DeleteInstance operation.
	//
	// > This property applies only to pay-as-you-go instances. It prevents instances created by Auto Scaling scale-out activities from being accidentally deleted, but does not affect normal scale-in activities. Instances with release protection enabled can still be released by scale-in activities.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set to which the ECS instance belongs.
	//
	// example:
	//
	// ds-bp13v7bjnj9gis****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the Elastic Compute Service (ECS) server. A period (.) or a hyphen (-) cannot be used as the first or last character, or used consecutively. In addition, the naming conventions for different types of instances are as follows:
	//
	// - Windows instances: The hostname must be 2 to 15 characters in length and can contain letters, digits, and hyphens (-). It cannot contain periods (.) or consist entirely of digits.
	//
	// - Other instances (such as Linux): The hostname must be 2 to 64 characters in length and can contain multiple periods (.). Each segment between two periods (.) can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// hos****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the HPC cluster to which the ECS instance belongs.
	//
	// example:
	//
	// hpc-clusterid
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enabled.
	//
	// - disabled: disabled.
	//
	// Default value: enabled.
	//
	// > For information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Specifies whether to forcefully use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// - optional: does not forcefully use the security-hardened mode.
	//
	// - required: forcefully uses the security-hardened mode. After you set this value, instance metadata cannot be accessed in normal mode.
	//
	// Default value: optional.
	//
	// > For more information about the modes for accessing instance metadata, see [Instance metadata access modes](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image from the specified image family for instance creation. If you have specified the `ImageId` parameter, you cannot specify this parameter.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file used for automatic creation of instances.
	//
	// > If the image previously used in the scaling configuration includes a system cloud disk and data cloud disks, the data cloud disk information is purged after you change the image.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. Image names must be unique within a region. If you specify ImageId, ImageName is ignored.
	//
	// You cannot use ImageName to specify images from Alibaba Cloud Marketplace.
	//
	// example:
	//
	// suse11sp3_64_20G_aliaegis_2015****.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The description of the ECS instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the ECS instance that is automatically created by using the scaling configuration.
	//
	// example:
	//
	// inst****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The collection of intelligent configuration information used to filter instance types that meet the requirements.
	InstancePatternInfos []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// After the fallback mode is enabled, when issues such as insufficient inventory are encountered, similar instance types of the same size are supplemented based on the currently selected instance types, or vSwitches in alternative zones are created and added to the scaling group.
	InstanceTypeCandidateOptions *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions `json:"InstanceTypeCandidateOptions,omitempty" xml:"InstanceTypeCandidateOptions,omitempty" type:"Struct"`
	// The information about specified instance types.
	InstanceTypeOverrides []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	// The instance types for multiple instance type support. If InstanceTypes is specified, InstanceType is ignored.
	//
	// When instances cannot be created by using the instance type with a higher priority, Auto Scaling automatically selects the instance type with the next priority to create instances.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth. In this case, the value of InternetMaxBandwidthOut specifies the fixed bandwidth value.
	//
	// - PayByTraffic: pay-by-traffic. In this case, the value of InternetMaxBandwidthOut specifies only the maximum bandwidth, and billing is based on the actual network traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// - If the purchased maximum outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - If the purchased maximum outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `InternetMaxBandwidthOut`. Default value: the value of `InternetMaxBandwidthOut`.
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
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is an I/O optimized instance. Valid values:
	//
	// - none: The instance is not an I/O optimized instance.
	//
	// - optimized: The instance is an I/O optimization instance.
	//
	// example:
	//
	// none
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses allocated to the elastic network interfaces (ENIs).
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair used to log on to the ECS instance.
	//
	//
	//
	// - For Windows instances, this parameter is ignored and is empty by default.
	//
	// - For Linux instances, password-based logon is disabled during initialization.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The weight of SLB backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify both CPU and Memory to define the range of instance types. For example, CPU=2 and Memory=16 defines all instance types with 2 vCPUs and 16 GiB of memory. Auto Scaling determines the set of active instance types based on factors such as I/O optimization and zones, and then creates instances of the lowest-priced instance type based on sorting by price.
	//
	// > This configuration takes effect only when the cost optimization mode is enabled and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The list of elastic network interfaces (ENIs).
	NetworkInterfaces []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	// Specifies whether to override existing data. Valid values:
	//
	// - true: overrides existing data.
	//
	// - false: does not override existing data.
	//
	// example:
	//
	// true
	Override     *bool   `json:"Override,omitempty" xml:"Override,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// `()~!@#$%^&*-_+=\\|{}[]:;\\"<>,.?/
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify the Password parameter, we recommend that you send the request over HTTPS to prevent password leaks.
	//
	// example:
	//
	// 123abc****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. To use this parameter, ensure that a password is configured for the image.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the RAM role associated with the ECS instance. This name is provided and maintained by RAM. You can call the ListRoles operation to query available RAM roles. For more information about how to create a RAM role, see the CreateRole operation.
	//
	// example:
	//
	// RamRoleTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the ECS instance belongs.
	//
	// example:
	//
	// abcd1234abcd****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource pool policy used when creating instances. After you specify this parameter, note the following items:
	//
	// - This parameter takes effect only when pay-as-you-go instances are created.
	//
	// - This parameter cannot be specified together with PrivatePoolOptions.MatchCriteria or PrivatePoolOptions.Id.
	ResourcePoolOptions *ModifyScalingConfigurationShrinkRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
	// The ID of the scaling configuration that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// asc-bp16har3jpj6fjbx****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The name of the scaling configuration. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter, digit, or Chinese character.
	//
	// The name of the scaling configuration must be unique within a scaling group in the same region. If you do not specify this parameter, the ID of the scaling configuration is used by default.
	//
	// example:
	//
	// test-modify
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The scheduler options.
	//
	// example:
	//
	// {"managedPrivateSpaceId":"testManagedPrivateSpaceId****"}
	SchedulerOptionsShrink *string `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// Specifies whether a vSwitch must be selected from the VSwitchIds list for the secondary elastic network interfaces (ENIs) when the VSwitchIds parameter is set for the secondary network interface controller (NIC).
	//
	// - Strict: strict mode. If no vSwitch in VSwitchIds meets the conditional requirements, instance creation fails.
	//
	// - Tolerant: tolerant mode. If no vSwitch in VSwitchIds meets the conditional requirements, no secondary network interface controller (NIC) is configured, but the instance is still created as active.
	//
	// Default value: Strict.
	//
	// example:
	//
	// Strict
	SecondaryEniVSwitchConstraint *string `json:"SecondaryEniVSwitchConstraint,omitempty" xml:"SecondaryEniVSwitchConstraint,omitempty"`
	// The ID of the security group to which the ECS instance belongs. ECS instances in the same security group can access each other.
	//
	// example:
	//
	// sg-F876F****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups to which the instance belongs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *ModifyScalingConfigurationShrinkRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Valid values:
	//
	// - 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After 1 hour, the system compares the bid price with the market price and checks resource availability to determine whether to retain automatic release the instance.
	//
	// - 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks resource availability to determine whether to retain automatic release the instance.
	//
	// > Alibaba Cloud sends an ECS system event notification 5 minutes before the instance is released. Spot instances are billed by second. We recommend that you select an appropriate protection period based on the expected task execution duration.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the spot instance. Currently, only Terminate (default) is supported, which directly releases the instance.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The information about spot instance types.
	SpotPriceLimits []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy for pay-as-you-go instances. Valid values:
	//
	//
	//
	// - NoSpot: a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: a spot instance with a user-defined maximum hourly price.
	//
	// - SpotAsPriceGo: a spot instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The storage set ID.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. The value must be an integer that is greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The multi-disk types for the system cloud disk. When a disk type with a higher priority is unavailable, Auto Scaling automatically attempts the next-priority disk type to create the system cloud disk. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: Enterprise SSD (ESSD).
	//
	// > If you specify this parameter, you cannot specify `SystemDisk.Category` at the same time.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The tags of the ECS instance. Tags are specified as key-value pairs. You can specify up to 20 tags. The following limits apply to Key and Value:
	//
	//
	//
	// - Key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. Key cannot be an empty string.
	//
	// - Value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. Value can be an empty string.
	//
	// example:
	//
	// {"key1":"value1","key2":"value2", ... "key5":"value5"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// - default: creates an instance that is not on a dedicated host.
	//
	// - host: creates an instance on a dedicated host. If you do not specify DedicatedHostId, Alibaba Cloud automatically selects a dedicated host for the instance.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// The user data of the ECS instance. The data must be encoded in Base64. The maximum size of the raw data before encoding is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The zone ID of the ECS instance.
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

func (s *ModifyScalingConfigurationShrinkRequest) GetInstanceTypeCandidateOptions() *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions {
	return s.InstanceTypeCandidateOptions
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

func (s *ModifyScalingConfigurationShrinkRequest) GetSecondaryEniVSwitchConstraint() *string {
	return s.SecondaryEniVSwitchConstraint
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

func (s *ModifyScalingConfigurationShrinkRequest) SetInstanceTypeCandidateOptions(v *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) *ModifyScalingConfigurationShrinkRequest {
	s.InstanceTypeCandidateOptions = v
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

func (s *ModifyScalingConfigurationShrinkRequest) SetSecondaryEniVSwitchConstraint(v string) *ModifyScalingConfigurationShrinkRequest {
	s.SecondaryEniVSwitchConstraint = &v
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
	if s.InstanceTypeCandidateOptions != nil {
		if err := s.InstanceTypeCandidateOptions.Validate(); err != nil {
			return err
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
	// Specifies whether the ECS instance uses the ecs-user account to log on. For more information, see [Manage the logon name of an ECS instance](https://help.aliyun.com/document_detail/388447.html). Valid values:
	//
	// - true: The ecs-user account is used.
	//
	// - false: The ecs-user account is not used.
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
	// The private pool ID, which is the ID of the elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool option for instance startup. A private pool is generated after an elasticity assurance or capacity reservation takes effect. You can select a private pool when you start an instance. Valid values:
	//
	// - Open: open mode. The system automatically matches an open private pool. If no matching private pool is available, the public pool is used to start the instance. You do not need to specify PrivatePoolOptions.Id in this mode.
	//
	// - Target: specified mode. A specified private pool is used to start the instance. If the specified private pool is unavailable, the instance fails to start. In this mode, you must specify the private pool ID, which means PrivatePoolOptions.Id is required.
	//
	// - None: no private pool is used. The instance does not use any private pool capacity for startup.
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
	// The ID of the automatic snapshot policy applied to the system cloud disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burst feature for the system cloud disk. Valid values:
	//
	// - true: enables the burst feature.
	//
	// - false: does not enable the burst feature.
	//
	// > This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">For more information, see [ESSD AutoPL disks](
	//
	// [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html)
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the system cloud disk. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: Enterprise SSD (ESSD).
	//
	// - ephemeral_ssd: local SSD.
	//
	// This parameter cannot be specified together with `SystemDiskCategories`. If neither this parameter nor `SystemDiskCategories` is specified, a default value is used. If the instance type belongs to Generation I instance families and the instance is not an I/O optimization instance, the default value is cloud. Otherwise, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system cloud disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test system disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system cloud disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter or Chinese character and cannot start with http:// or https://. Default value: empty.
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The encryption algorithm for the system cloud disk. Valid values:
	//
	// - AES-256.
	//
	// - SM4-128.
	//
	// Default value: AES-256.
	//
	// example:
	//
	// AES-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the system cloud disk. Valid values:
	//
	// - true: encrypts the system cloud disk.
	//
	// - false: does not encrypt the system cloud disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key used for the system cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level (PL) of the system cloud disk when the disk is an ESSD. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// > For more information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS (Input/Output Operations Per Second) performance metrics for the system cloud disk.
	//
	// > IOPS (Input/Output Operations Per Second) refers to the number of I/O operations that can be processed per second, representing the read/write (output/input) capability of block storage. Unit: operations per second.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the system cloud disk. Unit: GiB. Valid values:
	//
	// - Basic cloud disk: 20 to 500.
	//
	// - ESSD:
	//
	//   - PL0: 1 to 2048.
	//
	//   - PL1: 20 to 2048.
	//
	//   - PL2: 461 to 2048.
	//
	//   - PL3: 1261 to 2048.
	//
	// - ESSD AutoPL: 1 to 2048.
	//
	// - Other cloud disk types: 20 to 2048.
	//
	// The value of this parameter must be greater than or equal to max{1, ImageSize}.
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
	// The instance type of the ECS instance.
	//
	// > The instance type must be included in the instance type list of the scaling configuration.
	//
	// example:
	//
	// ecs.c6a.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the vSwitch.
	//
	// > The vSwitch must be included in the vSwitch list of the scaling group.
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
	// The ID of the automatic snapshot policy applied to the data cloud disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burst feature for the data cloud disk. Valid values:
	//
	// - true: enables the burst feature.
	//
	// - false: does not enable the burst feature.
	//
	// > This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The multiple disk types of the data cloud disk. Valid values:
	//
	// - cloud: basic cloud disk. The DeleteWithInstance property of a basic cloud disk created along with the instance is set to true.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_auto: ESSD AutoPL.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// > You cannot specify this parameter together with `DataDisk.Category`.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data cloud disk. Valid values:
	//
	// - cloud: basic cloud disk. The DeleteWithInstance attribute of a basic cloud disk created along with the instance is set to true.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_auto: ESSD AutoPL.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// This parameter cannot be specified together with `DataDisk.Categories`. If neither this parameter nor `DataDisk.Categories` is specified, the default value is used:
	//
	// - For I/O optimized instances, the default value is cloud_efficiency.
	//
	// - For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether the data cloud disk is released when the instance is released. Valid values:
	//
	// - true: The data cloud disk is released when the instance is released.
	//
	// - false: The data cloud disk is retained when the instance is released.
	//
	// This parameter can be set only for independent cloud disks (DataDisk.Category is set to cloud, cloud_efficiency, cloud_ssd, cloud_essd, or cloud_auto). Otherwise, an error is returned.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the system cloud disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test data disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data cloud disk. If you do not specify this parameter, the system uses the default behavior to allocate a mount point during automatic creation of the ECS instance. The mount points range from /dev/xvdb to /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdd
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data cloud disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter or Chinese character and cannot start with `http://` or `https://`.
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
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key for the data cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the data cloud disk when the disk is an ESSD. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// > For more information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS (Input/Output Operations Per Second) performance metrics of the data cloud disk.
	//
	// > IOPS (Input/Output Operations Per Second) measures the number of read and write operations that a block storage device can process per second.
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
	// After you specify this parameter, the disk size must be greater than or equal to the size of the snapshot specified by SnapshotId.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create the data cloud disk. After this parameter is specified, DataDisk.Size is ignored, and the size of the created disk is the size of the specified snapshot.
	//
	// If the snapshot was created on or before July 15, 2013, the request is rejected and the error message InvalidSnapshot.TooOld is returned.
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
	// The architecture types of the instance type. Valid values:
	//
	// - X86: x86 computing.
	//
	// - Heterogeneous: heterogeneous computing, such as GPU or FPGA.
	//
	// - BareMental: ECS Bare Metal Instance.
	//
	// - Arm: Arm computing.
	//
	// Default value: all architecture types are included.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether to include burstable instance types. Valid values:
	//
	// - Exclude: excludes burstable instance types.
	//
	// - Include: includes burstable instance types.
	//
	// - Required: includes only burstable instance types.
	//
	// Default value: Include.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The number of vCPU cores of the instance type in intelligent configuration mode, used to filter instance types that meet the requirements. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// Note the following information:
	//
	// - The InstancePatternInfo parameter applies only to scaling groups whose network type is VPC.
	//
	// - You must specify both InstancePatternInfo.Cores and InstancePatternInfo.Memory.
	//
	// - If you have specified instance types by using the InstanceType or InstanceTypes parameter, Auto Scaling preferentially uses the specified instance types for scale-out. When the specified instance types are out of stock, Auto Scaling selects the lowest-priced instance type from those that match the InstancePatternInfo parameter values.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architectures of the instance. Valid values:
	//
	// > N indicates that you can specify multiple CPU architectures. Valid values of N: 1 to 2.
	//
	// - X86.
	//
	// - ARM.
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The instance types to exclude. You can use wildcard characters (*) to exclude a single instance type or an entire instance family. Examples:
	//
	// - ecs.c6.large: excludes the ecs.c6.large instance type.
	//
	// - ecs.c6.*: excludes the entire c6 instance family.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The GPU types.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The categorization of the instance type. Valid values:
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
	// - Compute-optimized with GPU: GPU-accelerated compute-optimized.
	//
	// - Visual Compute-optimized: visual compute-optimized.
	//
	// - Heterogeneous Service: heterogeneous service.
	//
	// - Compute-optimized with FPGA: FPGA compute-optimized.
	//
	// - Compute-optimized with NPU: NPU compute-optimized.
	//
	// - ECS Bare Metal: ECS Bare Metal Instance.
	//
	// - High Performance Compute: high-performance computing (HPC).
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family, used to filter instance types that meet the requirements. This parameter takes effect when `CostOptimization` is enabled. Valid values:
	//
	// - EntryLevel: entry level, which refers to shared instance types. These instance types are more cost-effective but cannot guarantee stable computing performance. They are suitable for scenarios with low CPU utilization. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise level. These instance types provide stable performance and dedicated resources. They are suitable for scenarios that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit entry level, which refers to burstable instances. These instance types use CPU credits to ensure computing performance. They are suitable for scenarios with low average CPU utilization and occasional bursts. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families to query. N indicates that you can specify multiple instance families. Valid values of N: 1 to 10.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum hourly price for pay-as-you-go or spot instances in intelligent configuration mode, used to filter instance types that meet the requirements.
	//
	// > This parameter is required when SpotStrategy is set to SpotWithPriceLimit. In other cases, this parameter is optional.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPU cores of the instance type.
	//
	// > MaximumCpuCoreCount cannot exceed 4 times the value of MinimumCpuCoreCount.
	//
	// example:
	//
	// 4
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum number of GPUs of the instance. Valid values: positive integers.
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
	// The memory size of the instance type in intelligent configuration mode. Unit: GiB. This parameter is used to filter instance types that meet the requirements.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum baseline vCPU computing performance (sum of all vCPUs) for burstable instances (t5 and t6).
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPU cores of the instance.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum upper limit of IPv6 addresses on a single elastic network interfaces (ENIs) of the instance.
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of IPv4 addresses per ENI supported by the instance.
	//
	// [_single.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of elastic network interfaces (ENIs) supported by the instance.
	//
	// example:
	//
	// 2
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of GPUs of the instance. Valid values: positive integers.
	//
	// example:
	//
	// 2
	MinimumGpuAmount *int32 `json:"MinimumGpuAmount,omitempty" xml:"MinimumGpuAmount,omitempty"`
	// The minimum initial vCPU credits for burstable instances (t5 and t6).
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
	// The processor models of the instance. N indicates that you can specify multiple processor models. Valid values of N: 1 to 10.
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

type ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions struct {
	// The CIDR blocks for optional vSwitches when supplementing vSwitches in other zones.
	AllowCidrBlocks []*string `json:"AllowCidrBlocks,omitempty" xml:"AllowCidrBlocks,omitempty" type:"Repeated"`
	// Specifies whether to allow supplementing vSwitches in other zones.
	//
	// > The instance type remains unchanged. Only alternative zones are selected. When all selected zones of the scaling group cannot be used for scale-out due to issues such as insufficient inventory, new zone vSwitches are automatically added to the scaling group based on this configuration.
	//
	// > For example, if the scaling group is configured with the cn-hangzhou-h and cn-hangzhou-g zones, and neither zone can be used for scale-out, Auto Scaling may create a vSwitch in cn-hangzhou-k based on real-time inventory and add it to the scaling group.
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
	// 1.5
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) GetAllowCidrBlocks() []*string {
	return s.AllowCidrBlocks
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) GetAllowCrossAz() *bool {
	return s.AllowCrossAz
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) GetAllowDifferentGeneration() *bool {
	return s.AllowDifferentGeneration
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) SetAllowCidrBlocks(v []*string) *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions {
	s.AllowCidrBlocks = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) SetAllowCrossAz(v bool) *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions {
	s.AllowCrossAz = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) SetAllowDifferentGeneration(v bool) *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions {
	s.AllowDifferentGeneration = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) SetEnabled(v bool) *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions {
	s.Enabled = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) SetMaxPrice(v float32) *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions {
	s.MaxPrice = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides struct {
	// To specify the capacity of an instance type in the scaling configuration, specify both this parameter and InstanceTypeOverride.WeightedCapacity.
	//
	// This parameter specifies the instance type. You can specify multiple instance types and use the InstanceTypeOverride.WeightedCapacity parameter to extend multi-instance-type support with custom weights.
	//
	// > You cannot specify this parameter together with instanceTypes.
	//
	// Valid values of InstanceType: ECS instance types that are available for purchase.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// To scale the scaling group based on instance type capacity, specify this parameter after specifying LaunchTemplateOverride.InstanceType.
	//
	// This parameter specifies the weight of the instance type, which represents the capacity of a single instance of this type in the scaling group. A larger weight means fewer instances of this type are required to meet the expected capacity.
	//
	// Because instance types differ in vCPUs, memory, and other performance metrics, you can assign different weights to different instance types based on your requirements.
	//
	// Example:
	//
	// - Current capacity: 0.
	//
	// - Expected number of instances: 6.
	//
	// - Capacity of ecs.c5.xlarge: 4.
	//
	// To meet the expected capacity, the scaling group scales out 2 ecs.c5.xlarge instances.
	//
	// > During scale-out, the capacity of the scaling group cannot exceed the sum of the maximum capacity (MaxSize) and the maximum weight of the instance types.
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
	// The type of the elastic network interfaces (ENIs). When you use this parameter, you must use NetworkInterfaces to set the primary network interface controller (NIC). You cannot set SecurityGroupId or SecurityGroupIds at the same time. Valid values:
	//
	// - Primary: primary network interface controller (NIC).
	//
	// - Secondary: secondary network interface controller (NIC).
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses for the primary ENI. Note the following items:
	//
	// This parameter takes effect only when NetworkInterface.InstanceType is set to Primary. If NetworkInterface.InstanceType is set to Secondary or left empty, you cannot specify this parameter.
	//
	// After you specify this parameter, you cannot specify Ipv6AddressCount.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication mode of the network interface controller (NIC). Valid values:
	//
	// - Standard: uses the TCP communication mode.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication mode.
	//
	// Default value: Standard.
	//
	// > The number of elastic network interfaces (ENIs) in RDMA mode cannot exceed the limit of the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of secondary private IPv4 addresses to assign to the ENI. Valid values: 1 to 49.
	//
	// - The value cannot exceed the limit on the number of IP addresses for the instance type. For more information, see [Instance families](https://www.alibabacloud.com/help/en/ecs/user-guide/overview-of-instance-families).
	//
	// - NetworkInterface.N.SecondaryPrivateIpAddressCount specifies the number of secondary private IPv4 addresses to assign to the ENI (excluding the primary private IP address of the ENI). The system randomly assigns IP addresses from the available CIDR block of the vSwitch (NetworkInterface.N.VSwitchId) to which the ENI belongs.
	//
	// example:
	//
	// 6
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The IDs of one or more security groups to which the elastic network interfaces (ENIs) belongs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The IDs of the vSwitches to which the elastic network interfaces (ENIs) belongs. N specifies that you can set multiple vSwitches. Valid values of N: 1 to 10.
	//
	// Note the following items:
	//
	// - If InstanceType is set to Primary, you cannot set this parameter. The default value is the vSwitch of the instance.
	//
	// - If InstanceType is set to Secondary or left empty, this parameter is optional.
	//
	// - The system selects a vSwitch that is in the same zone as the primary network interface controller (NIC) and has the most available IP addresses.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
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

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetVSwitchIds() []*string {
	return s.VSwitchIds
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

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetSecondaryPrivateIpAddressCount(v int32) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetVSwitchIds(v []*string) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.VSwitchIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestResourcePoolOptions struct {
	// The private pool IDs, which are the elasticity assurance IDs or capacity reservation IDs. Only Target-type private pool IDs can be specified for this parameter. This parameter cannot be specified together with the PrivatePoolTags parameter.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// Filters available Target private pools by tags. Valid values of N: 1 to 20.
	//
	// Note:
	//
	// - When this parameter is specified, the system filters only the Target private pools associated with the account that match the tags and satisfy the scaling group constraints (such as zone and instance type).
	//
	// - Tag matching rule: A private pool must match all specified tags.
	//
	// - This parameter cannot be specified together with the PrivatePoolIds parameter.
	PrivatePoolTags []*ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags `json:"PrivatePoolTags,omitempty" xml:"PrivatePoolTags,omitempty" type:"Repeated"`
	// Resource pools include private pools generated after elasticity assurance or capacity reservation takes effect, and public pools, which are available for instance startup. Valid values:
	//
	// - PrivatePoolFirst: private pool first. When this strategy is selected, if ResourcePoolOptions.PrivatePoolIds is specified or PrivatePoolTags conditions are met, the corresponding private pools are used first. If no private pool is specified or the specified private pool has insufficient capacity, open-type private pools are automatically matched. If no eligible private pool is available, the public pool is used to create instances.
	//
	// - PrivatePoolOnly: private pool only. When this strategy is selected, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pool has insufficient capacity, the instance fails to start.
	//
	// - PublicPoolFirst: public pool first. The public pool is used first to create instances. When public pool resources are insufficient, private pool resources are used as a supplement. Open-type private pools are automatically matched first. If no eligible private pool is available, the Target-type private pools specified by ResourcePoolOptions.PrivatePoolIds or matching PrivatePoolTags conditions are used. (This strategy is in invitational preview and is not available for use.)
	//
	// - None: does not use a resource pool strategy.
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

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) GetPrivatePoolTags() []*ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags {
	return s.PrivatePoolTags
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *ModifyScalingConfigurationShrinkRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) SetPrivatePoolTags(v []*ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) *ModifyScalingConfigurationShrinkRequestResourcePoolOptions {
	s.PrivatePoolTags = v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) SetStrategy(v string) *ModifyScalingConfigurationShrinkRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptions) Validate() error {
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

type ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags struct {
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

func (s ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) GetKey() *string {
	return s.Key
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) GetValue() *string {
	return s.Value
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) SetKey(v string) *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags {
	s.Key = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) SetValue(v string) *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags {
	s.Value = &v
	return s
}

func (s *ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingConfigurationShrinkRequestSecurityOptions struct {
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
	// The instance type of the spot instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The bid price for the spot instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
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
