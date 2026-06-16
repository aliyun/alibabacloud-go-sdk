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
	// - default: The instance is not associated with the dedicated host. If you restart an instance that was stopped in economical mode, the instance may be placed on a different dedicated host in the automatic deployment resource pool if the resources of the original dedicated host are insufficient.
	//
	// - host: The instance is associated with the dedicated host. If you restart an instance that was stopped in economical mode, the instance is still placed on the original dedicated host. If the resources of the original dedicated host are insufficient, the instance fails to restart.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of vCPUs. Unit: cores.
	//
	// You can specify both \\`Cpu\\` and \\`Memory\\` to define a range of instance types. For example, if you set \\`Cpu\\` to 2 and \\`Memory\\` to 16, all instance types with 2 vCPUs and 16 GiB of memory are matched. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and then creates the instance of the lowest-priced instance type.
	//
	// > This configuration is effective only when the cost optimization mode is enabled and no instance types are specified in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// - Standard: the standard mode. For more information about the instance performance, see the "Performance modes" section in [What is a burstable instance?](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: the unlimited mode. For more information about the instance performance, see the "Performance modes" section in [What is a burstable instance?](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The custom priority of the ECS instance type and vSwitch.
	//
	// 	Notice: This parameter is in effect only when the scaling policy of the scaling group is set to the priority-based policy.
	//
	// If an instance cannot be created using the instance type and vSwitch with a higher priority, Auto Scaling automatically uses the instance type and vSwitch combination with the next priority to create the instance.
	//
	// > If you specify custom priorities for only some instance type and vSwitch combinations, the combinations for which you do not specify custom priorities have a lower priority than the combinations for which you specify custom priorities. The priority of the combinations for which you do not specify custom priorities is determined by the order of vSwitches in the scaling group and the order of instance types in the scaling configuration.
	//
	// >
	//
	// > - For example, if the vSwitches in the scaling group are ordered as vsw1 and vsw2, the instance types in the scaling configuration are ordered as type1 and type2, and the custom priority is set to ["vsw2+type2", "vsw1+type2"], the final priority is: "vsw2+type2" > "vsw1+type2" > "vsw1+type1" > "vsw2+type1".
	CustomPriorities []*ModifyScalingConfigurationShrinkRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The collection of data disk information.
	DataDisks []*ModifyScalingConfigurationShrinkRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-zm04u8r3lohsq****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// Specifies whether to create the ECS instance on a dedicated host. If you specify \\`DedicatedHostId\\`, the \\`SpotStrategy\\` and \\`SpotPriceLimit\\` settings in the request are ignored. This is because dedicated hosts do not support spot instances.
	//
	// You can call the DescribeDedicatedHosts operation to query the list of dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection attribute of the instance. This parameter specifies whether you can release the instance using the ECS console or by calling the DeleteInstance operation. This prevents the instance from being accidentally released. Valid values:
	//
	// - true: enables release protection. You cannot release the instance using the ECS console or by calling the DeleteInstance operation.
	//
	// - false: disables release protection. You can release the instance using the ECS console or by calling the DeleteInstance operation.
	//
	// > This attribute applies only to pay-as-you-go instances. It prevents the instances that are scaled out by Auto Scaling from being accidentally released. This attribute does not affect normal scale-in activities. Instances for which release protection is enabled can be released during scale-in activities.
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
	// The hostname of the ECS instance. A period (.) or a hyphen (-) cannot be used as the first or last character of the hostname. Consecutive periods (.) or hyphens (-) are not allowed. The naming conventions for hostnames vary based on the instance operating system:
	//
	// - For Windows instances, the hostname must be 2 to 15 characters in length and can contain letters, digits, and hyphens (-). It cannot contain periods (.) or consist of only digits.
	//
	// - For other instance types, such as Linux, the hostname must be 2 to 64 characters in length. You can use periods (.) to separate the hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-).
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
	// - enabled: enable.
	//
	// - disabled: disable.
	//
	// Default value: enabled.
	//
	// > For more information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Specifies whether to enforce the security-hardened mode (IMDSv2) when you access instance metadata. Valid values:
	//
	// - optional: does not enforce the use of IMDSv2.
	//
	// - required: enforces the use of IMDSv2. If you set the value to \\`required\\`, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// > For more information about how to access instance metadata, see [Access modes of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image from the specified image family to create instances. If you have set the `ImageId` parameter, you cannot set this parameter.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file that is used to create the instances.
	//
	// > If the image that was previously used in the scaling configuration includes a system disk and data disks, the original data disk information is cleared after you change the image.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis_2014****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image file. The image name must be unique in a region. If you specify \\`ImageId\\`, \\`ImageName\\` is ignored.
	//
	// You cannot use \\`ImageName\\` to specify a Marketplace image.
	//
	// example:
	//
	// suse11sp3_64_20G_aliaegis_2015****.vhd
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The description of the ECS instance. The description must be 2 to 256 English or Chinese characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the ECS instances that are automatically created using this scaling configuration.
	//
	// example:
	//
	// inst****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The collection of intelligent configuration information that is used to filter instance types that meet the specified requirements.
	InstancePatternInfos []*ModifyScalingConfigurationShrinkRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// After you enable the alternative mode, if issues such as insufficient inventory occur, the system supplements the selected instance types with similar instance types of the same size, or creates vSwitches in alternative zones and adds them to the scaling group.
	InstanceTypeCandidateOptions *ModifyScalingConfigurationShrinkRequestInstanceTypeCandidateOptions `json:"InstanceTypeCandidateOptions,omitempty" xml:"InstanceTypeCandidateOptions,omitempty" type:"Struct"`
	// The information about the specified instance types.
	InstanceTypeOverrides []*ModifyScalingConfigurationShrinkRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	// The instance types. If you use \\`InstanceTypes\\`, \\`InstanceType\\` is ignored.
	//
	// If an instance cannot be created using the instance type with a higher priority, Auto Scaling automatically uses the instance type with the next priority to create the instance.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth. If you set the value to PayByBandwidth, the value of \\`InternetMaxBandwidthOut\\` is the selected fixed bandwidth.
	//
	// - PayByTraffic: pay-by-data-transfer. If you set the value to PayByTraffic, the value of \\`InternetMaxBandwidthOut\\` is the maximum bandwidth, and the billing is based on the actual network traffic.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Value range:
	//
	// - If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. The default value is 10.
	//
	// - If the purchased outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `InternetMaxBandwidthOut`. The default value is the value of `InternetMaxBandwidthOut`.
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
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// - none: The instance is not I/O optimized.
	//
	// - optimized: The instance is I/O optimized.
	//
	// example:
	//
	// none
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses to be assigned to the ENI.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair that is used to log on to the ECS instance.
	//
	// - For Windows instances, this parameter is ignored. The default value is empty.
	//
	// - For Linux instances, password-based logon is disabled by default.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The weight of the backend server. Valid values: 1 to 100.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify both \\`Cpu\\` and \\`Memory\\` to define a range of instance types. For example, if you set \\`Cpu\\` to 2 and \\`Memory\\` to 16, all instance types with 2 vCPUs and 16 GiB of memory are matched. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and then creates the instance of the lowest-priced instance type.
	//
	// > This configuration is effective only when the cost optimization mode is enabled and no instance types are specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The list of ENIs.
	NetworkInterfaces []*ModifyScalingConfigurationShrinkRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	// Specifies whether to overwrite the parameter. Valid values:
	//
	// - true: Overwrite the parameter.
	//
	// - false: Do not overwrite the parameter.
	//
	// example:
	//
	// true
	Override     *bool   `json:"Override,omitempty" xml:"Override,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters can be:
	//
	// \\`()\\~!@#$%^&\\*-_+=|{}[]:;\\"<>,.?/
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify the \\`Password\\` parameter, we recommend that you send requests over HTTPS to prevent password leaks.
	//
	// example:
	//
	// 123abc****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password that is preset in the image. If you use this parameter, make sure that a password is set for the image.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the RAM role of the ECS instance. The RAM role is provided and maintained by RAM. You can call the ListRoles operation to query the available RAM roles. For information about how to create a RAM role, see API CreateRole.
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
	// The resource pool policy to use when creating an instance. Note the following when you set this parameter:
	//
	// - This parameter is in effect only when you create a pay-as-you-go instance.
	//
	// - You cannot set this parameter and \\`PrivatePoolOptions.MatchCriteria\\` or \\`PrivatePoolOptions.Id\\` at the same time.
	ResourcePoolOptions *ModifyScalingConfigurationShrinkRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
	// The ID of the scaling configuration that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// asc-bp16har3jpj6fjbx****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The name of the scaling configuration. The name must be 2 to 64 English or Chinese characters in length. It must start with a digit, a letter, or a Chinese character. The name can contain digits, underscores (_), hyphens (-), and periods (.).
	//
	// The name of a scaling configuration must be unique within a scaling group in the same region. If you do not specify this parameter, the ID of the scaling configuration is used by default.
	//
	// example:
	//
	// test-modify
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The scheduling options.
	//
	// example:
	//
	// ["testManagedPrivateSpaceId****"]
	SchedulerOptionsShrink *string `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// The ID of the security group to which the ECS instance belongs. ECS instances in the same security group can access each other.
	//
	// example:
	//
	// sg-F876F****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the security group.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *ModifyScalingConfigurationShrinkRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Valid values:
	//
	// - 1: Alibaba Cloud ensures that the instance runs for 1 hour and is not automatically released. After 1 hour, the system automatically compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// - 0: Alibaba Cloud does not ensure that the instance runs for 1 hour after it is created. The system automatically compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// > Alibaba Cloud sends a notification to you through ECS system events 5 minutes before the instance is released. Spot instances are billed by the second. Select a protection period based on the time required to complete your task.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the spot instance. Currently, only Terminate is supported, which is the default value. This value indicates that the instance is directly released.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The information about the spot instance types.
	SpotPriceLimits []*ModifyScalingConfigurationShrinkRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy for the pay-as-you-go instance. Valid values:
	//
	// - NoSpot: a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: a spot instance for which you specify the maximum hourly price.
	//
	// - SpotAsPriceGo: a spot instance for which the system automatically bids based on the current market price.
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
	// The maximum number of partitions in the storage set. The value must be an integer that is greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The categories of the system disk. If a disk of a category with a higher priority cannot be created, Auto Scaling automatically tries to create a disk of a category with the next priority. Valid values:
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// > You cannot specify this parameter and `SystemDisk.Category` at the same time.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The tags of the ECS instance. You can specify up to 20 tags in key-value pairs. The following limits apply to keys and values:
	//
	// - A key can be up to 64 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. If you specify tags, the key cannot be an empty string.
	//
	// - A value can be up to 128 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`. The value can be an empty string.
	//
	// example:
	//
	// {"key1":"value1","key2":"value2", ... "key5":"value5"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// - default: Create the instance not on a dedicated host.
	//
	// - host: Create the instance on a dedicated host. If you do not specify \\`DedicatedHostId\\`, Alibaba Cloud automatically selects a dedicated host for the instance.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// The custom data of the ECS instance. The data must be Base64-encoded. The raw data can be up to 32 KB in size.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the zone to which the ECS instance belongs.
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
	// Specifies whether to log on to the ECS instance as the ecs-user user. For more information, see [Manage logon usernames of ECS instances](https://help.aliyun.com/document_detail/388447.html). Valid values:
	//
	// true: yes.
	//
	// false: no.
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
	// The ID of the private pool. The private pool can be an Elastic Assurance service or a Capacity Reservation service.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The capacity option of the private pool for starting the instance. The private pool is generated after an Elastic Assurance service or a Capacity Reservation service takes effect. You can select a private pool to start an instance. Valid values:
	//
	// - Open: open mode. The system automatically matches the instance with an open private pool. If no open private pools are available, the instance is started using public pool resources. You do not need to set the \\`PrivatePoolOptions.Id\\` parameter in this mode.
	//
	// - Target: specified mode. The instance is started using the capacity of a specified private pool. If the specified private pool is unavailable, the instance fails to start. You must specify the private pool ID by setting the \\`PrivatePoolOptions.Id\\` parameter in this mode.
	//
	// - None: no mode. The instance is not started using the capacity of a private pool.
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
	// The ID of the automatic snapshot policy used for the system disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// - true: enable.
	//
	// - false: disable.
	//
	// > This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">
	//
	// For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the system disk. Valid values:
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// You cannot specify this parameter and `SystemDiskCategories` at the same time. If neither this parameter nor `SystemDiskCategories` is specified, this parameter has a default value. If the instance type is from instance family I and the instance is not I/O optimized, the default value is \\`cloud\\`. Otherwise, the default value is \\`cloud_efficiency\\`.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 English or Chinese characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test system disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 English or Chinese characters in length. It must start with a letter or a Chinese character and cannot start with http\\:// or https\\://. It can contain digits, colons (:), underscores (_), and hyphens (-). Default value: empty
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The encryption algorithm used for the system disk. Valid values:
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
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// - true: encrypt the system disk.
	//
	// - false: do not encrypt the system disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key used for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// - PL0: A single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// > For more information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The pre-configured IOPS of the system disk.
	//
	// > IOPS, or input/output operations per second, is the number of I/O operations that a block storage device can process per second. It indicates the read and write performance of the block storage device.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the system disk. Unit: GiB. Value range:
	//
	// - Basic disks: 20 to 500.
	//
	// - ESSDs:
	//
	//   - PL0: 1 to 2048.
	//
	//   - PL1: 20 to 2048.
	//
	//   - PL2: 461 to 2048.
	//
	//   - PL3: 1261 to 2048.
	//
	// - ESSD AutoPL cloud disks: 1 to 2048.
	//
	// - Other disk categories: 20 to 2048.
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
	// > The instance type must be included in the list of instance types in the scaling configuration.
	//
	// example:
	//
	// ecs.c6a.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the vSwitch.
	//
	// > The vSwitch must be included in the list of vSwitches in the scaling group.
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
	// The ID of the automatic snapshot policy used for the data disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable performance burst for the system disk. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// > This parameter takes effect only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">
	//
	// For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The categories of the data disk. Valid values:
	//
	// - cloud: basic disk. The \\`DeleteWithInstance\\` attribute of a basic disk that is created with an instance is \\`true\\`.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// > You cannot specify this parameter and `DataDisk.Category` at the same time.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data disk. Valid values:
	//
	// - cloud: basic disk. The \\`DeleteWithInstance\\` attribute of a basic disk that is created with an instance is \\`true\\`.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_essd: ESSD.
	//
	// You cannot specify this parameter and `DataDisk.Categories` at the same time. If neither this parameter nor `DataDisk.Categories` is specified, this parameter has a default value:
	//
	// - For I/O optimized instances, the default value is \\`cloud_efficiency\\`.
	//
	// - For non-I/O optimized instances, the default value is \\`cloud\\`.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// - true: The disk is released with the instance.
	//
	// - false: The disk is not released with the instance.
	//
	// You can set this parameter only for independent cloud disks (\\`DataDisk.Category\\` is \\`cloud\\`, \\`cloud_efficiency\\`, \\`cloud_ssd\\`, \\`cloud_essd\\`, or \\`cloud_auto\\`). Otherwise, an error occurs.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 English or Chinese characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test data disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. If you do not specify this parameter, the system allocates a mount point when the ECS instance is automatically created. The mount point starts from /dev/xvdb and goes to /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdd
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 English or Chinese characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether the system disk is encrypted. Valid values:
	//
	// - true: The system disk is encrypted.
	//
	// - false: The system disk is not encrypted.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key for the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD that is used as the data disk. Valid values:
	//
	// - PL0: A single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// > For more information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The pre-configured IOPS of the data disk.
	//
	// > IOPS, or input/output operations per second, is the number of I/O operations that a block storage device can process per second. It indicates the read and write performance of the block storage device.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the data disk. Unit: GiB. Value range:
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
	// If you specify this parameter, the disk size must be greater than or equal to the size of the snapshot specified by \\`SnapshotId\\`.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot that is used to create the data disk. If you specify this parameter, \\`DataDisk.Size\\` is ignored, and the size of the created disk is the size of the specified snapshot.
	//
	// If the snapshot was created on or before July 15, 2013, the call is rejected, and the \\`InvalidSnapshot.TooOld\\` error is returned.
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
	// The architecture type of the instance type. Valid values:
	//
	// - X86: x86.
	//
	// - Heterogeneous: heterogeneous computing, such as GPU and FPGA.
	//
	// - BareMental: ECS Bare Metal Instance.
	//
	// - Arm: Arm.
	//
	// Default value: all architecture types.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether to include burstable instance types. Valid values:
	//
	// - Exclude: do not include burstable instance types.
	//
	// - Include: include burstable instance types.
	//
	// - Required: include only burstable instance types.
	//
	// Default value: Include.
	//
	// example:
	//
	// Include
	BurstablePerformance *string `json:"BurstablePerformance,omitempty" xml:"BurstablePerformance,omitempty"`
	// The number of vCPU cores of the instance type in intelligent configuration mode. This parameter is used to filter instance types. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// Note the following information:
	//
	// - The \\`InstancePatternInfo\\` parameter is applicable only to scaling groups whose NetworkType is set to VPC.
	//
	// - You must specify \\`InstancePatternInfo.Cores\\` and \\`InstancePatternInfo.Memory\\` at the same time.
	//
	// - If you specify instance types using the \\`InstanceType\\` or \\`InstanceTypes\\` parameter, Auto Scaling preferentially uses the specified instance types for scale-out activities. If the specified instance types are out of stock, Auto Scaling uses the lowest-priced instance type among those that meet the requirements specified by the \\`InstancePatternInfo\\` parameter for scale-out activities.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architecture of the instance. Valid values:
	//
	// > You can specify up to two CPU architectures.
	//
	// - X86.
	//
	// - ARM.
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The instance types to exclude. You can use a wildcard character (\\*) to exclude an instance type or an entire instance family. Examples:
	//
	// - ecs.c6.large: excludes the ecs.c6.large instance type.
	//
	// - ecs.c6.\\*: excludes all instance types of the c6 family.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The GPU type.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The category of the instance type. Valid values:
	//
	// - General-purpose: General-purpose.
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
	// - Enhanced: enhanced instance families.
	//
	// - Shared: shared-resource instances.
	//
	// - Compute-optimized with GPU: GPU.
	//
	// - Visual Compute-optimized: visual compute-optimized.
	//
	// - Heterogeneous Service: heterogeneous computing.
	//
	// - Compute-optimized with FPGA: FPGA.
	//
	// - Compute-optimized with NPU: NPU-accelerated.
	//
	// - ECS Bare Metal: ECS Bare Metal Instance.
	//
	// - High Performance Compute: high-performance computing (HPC).
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family. This parameter is used to filter instance types. This parameter takes effect only when \\`CostOptimization\\` is enabled. Valid values:
	//
	// - EntryLevel: entry-level instances (shared). This instance type is cost-effective but does not provide stable computing performance. It is suitable for business scenarios that have low CPU utilization. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise-level instances. This instance type provides stable performance and dedicated resources, and is suitable for business scenarios that have high stability requirements. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit entry-level instances (burstable). This instance type provides CPU credits to ensure computing performance. It is suitable for business scenarios that have low and sporadic CPU utilization. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families to query. You can specify up to 10 instance families.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum hourly price that you can accept for a pay-as-you-go or spot instance in intelligent configuration mode. This parameter is used to filter instance types.
	//
	// > This parameter is required when \\`SpotStrategy\\` is set to \\`SpotWithPriceLimit\\`. In other cases, this parameter is optional.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPU cores of the instance type.
	//
	// > The value of \\`MaximumCpuCoreCount\\` cannot be more than four times the value of \\`MinimumCpuCoreCount\\`.
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
	// The memory size of the instance type in intelligent configuration mode. Unit: GiB. This parameter is used to filter instance types.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum baseline vCPU computing performance (for all vCPUs) of a t5 or t6 burstable instance.
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
	// The minimum number of IPv6 addresses that can be assigned to a single ENI of the instance.
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of IPv4 addresses that can be assigned to a single ENI of the instance.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of ENIs that can be attached to the instance.
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
	// The minimum initial vCPU credit for a t5 or t6 burstable instance.
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
	// The processor model of the instance. You can specify up to 10 processor models.
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
	// When supplementing with vSwitches from other zones is allowed, you must specify the optional CIDR blocks for the vSwitches.
	AllowCidrBlocks []*string `json:"AllowCidrBlocks,omitempty" xml:"AllowCidrBlocks,omitempty" type:"Repeated"`
	// Specifies whether to allow supplementing with vSwitches from other zones.
	//
	// > The instance type remains unchanged, and only new zones are selected as alternatives. When the scaling group cannot be scaled out in any of the selected zones due to issues such as insufficient inventory, the system automatically adds a new vSwitch in another zone to the scaling group based on this configuration.
	//
	// > For example, if the scaling group is configured with zones cn-hangzhou-h and cn-hangzhou-g, and scale-out fails in both zones, ESS may create a vSwitch in cn-hangzhou-k and add it to the scaling group based on real-time inventory.
	//
	// example:
	//
	// true
	AllowCrossAz *bool `json:"AllowCrossAz,omitempty" xml:"AllowCrossAz,omitempty"`
	// Specifies whether to allow supplementing with instance types from other generations.
	//
	// - For example, if the current instance type is ecs.c7.large, you can enable this feature to use alternative instance types such as ecs.c6.large and ecs.c8.large.
	//
	// example:
	//
	// true
	AllowDifferentGeneration *bool `json:"AllowDifferentGeneration,omitempty" xml:"AllowDifferentGeneration,omitempty"`
	// Specifies whether to enable the alternative mode.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum price for alternative instance types.
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
	// If you want to specify the capacity of an instance type in a scaling configuration, specify this parameter and \\`InstanceTypeOverride.WeightedCapacity\\`.
	//
	// This parameter is used to specify an instance type. You can specify this parameter multiple times. You can use this parameter with the \\`InstanceTypeOverride.WeightedCapacity\\` parameter to enable custom weights for multiple instance types.
	//
	// > If you specify this parameter, you cannot specify \\`instanceTypes\\`.
	//
	// Valid values of InstanceType: available ECS instance types.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// If you want the scaling group to scale based on the capacity of instance types, specify this parameter after you specify \\`LaunchTemplateOverride.InstanceType\\`.
	//
	// This parameter specifies the weight of an instance type, which indicates the capacity of a single instance of the specified instance type in the scaling group. A higher weight indicates that a smaller number of instances of the specified instance type are required to meet the expected capacity.
	//
	// You can set different weights for different instance types as needed because the performance metrics, such as the number of vCPUs and memory size, vary based on the instance type.
	//
	// For example:
	//
	// - Current capacity: 0.
	//
	// - Expected capacity: 6.
	//
	// - Capacity of the ecs.c5.xlarge instance type: 4.
	//
	// To meet the expected capacity, the scaling group scales out two ecs.c5.xlarge instances.
	//
	// > The capacity of the scaling group after a scale-out activity cannot exceed the sum of the maximum capacity (MaxSize) and the maximum weight of an instance type.
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
	// The type of the ENI. When you use this parameter, you must use \\`NetworkInterfaces\\` to configure the primary ENI. You cannot set the \\`SecurityGroupId\\` or \\`SecurityGroupIds\\` parameter. Valid values:
	//
	// - Primary: primary ENI.
	//
	// - Secondary: secondary ENI.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses to be assigned to the primary ENI. Note:
	//
	// This parameter takes effect only when \\`NetworkInterface.InstanceType\\` is set to \\`Primary\\`. You cannot set this parameter if \\`NetworkInterface.InstanceType\\` is set to \\`Secondary\\` or is empty.
	//
	// After you set this parameter, you cannot set \\`Ipv6AddressCount\\`.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication mode of the NIC. Valid values:
	//
	// - Standard: uses the TCP communication mode.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication mode.
	//
	// Default value: Standard.
	//
	// > The number of ENIs in RDMA mode cannot exceed the limit for the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of secondary private IPv4 addresses to assign to the NIC. Valid values: 1 to 49.
	//
	// - The value cannot exceed the limit on the number of IP addresses for the instance type. For more information, see [Instance families](https://help.aliyun.com/zh/ecs/user-guide/overview-of-instance-families).
	//
	// - \\`NetworkInterface.N.SecondaryPrivateIpAddressCount\\` is used to assign a number of secondary private IPv4 addresses to the NIC, excluding the primary private IP address of the NIC. The system randomly assigns the addresses from the available CIDR block of the vSwitch where the NIC is located (\\`NetworkInterface.N.VSwitchId\\`).
	//
	// example:
	//
	// 6
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The IDs of one or more security groups to which the ENI belongs.
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

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
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

func (s *ModifyScalingConfigurationShrinkRequestNetworkInterfaces) SetSecondaryPrivateIpAddressCount(v int32) *ModifyScalingConfigurationShrinkRequestNetworkInterfaces {
	s.SecondaryPrivateIpAddressCount = &v
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
	// The ID of the private pool. The private pool can be an Elastic Assurance service or a Capacity Reservation service. You can only specify the ID of a Target private pool. You cannot specify this parameter and the \\`PrivatePoolTags\\` parameter at the same time.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// Filter available Target private pools by tag. You can specify up to 20 tags.
	//
	// Description:
	//
	// - When you configure this parameter, the system filters only the associated Target private pools under your account to find private pools that match the tags and meet the constraints of the scaling group, such as zone and instance type.
	//
	// - Tag matching rule: The private pool must match all specified tags.
	//
	// - You cannot specify this parameter and the \\`PrivatePoolIds\\` parameter at the same time.
	PrivatePoolTags []*ModifyScalingConfigurationShrinkRequestResourcePoolOptionsPrivatePoolTags `json:"PrivatePoolTags,omitempty" xml:"PrivatePoolTags,omitempty" type:"Repeated"`
	// The resource pool includes the private pool generated after an Elastic Assurance service or a Capacity Reservation service takes effect, and the public pool. You can select a resource pool to start an instance. Valid values:
	//
	// - PrivatePoolFirst: The private pool is used first. If you select this policy and specify \\`ResourcePoolOptions.PrivatePoolIds\\` or meet the \\`PrivatePoolTags\\` conditions, the corresponding private pool is used first. If you do not specify a private pool or the specified private pool has insufficient capacity, the system automatically matches an open private pool. If no eligible private pools are available, a public pool is used to create the instance.
	//
	// - PrivatePoolOnly: Only the private pool is used. If you select this policy, you must specify \\`ResourcePoolOptions.PrivatePoolIds\\`. If the specified private pool has insufficient capacity, the instance fails to start.
	//
	// - PublicPoolFirst: The public pool is used first. A public pool is used first to create the instance. If the public pool has insufficient resources, a private pool is used. The system preferentially matches an open private pool. If no eligible private pools are available, the system uses the Target private pool that is specified by \\`ResourcePoolOptions.PrivatePoolIds\\` or meets the \\`PrivatePoolTags\\` conditions. (This policy is in invitational preview and is not yet available for use.)
	//
	// - None: No resource pool policy is used.
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
	// - Enclave: The ECS instance uses an enclave to build a confidential computing environment. For more information, see [Build a confidential computing environment using an enclave](https://help.aliyun.com/document_detail/203433.html).
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
	// The instance type of the spot instance. This parameter is in effect when \\`SpotStrategy\\` is set to \\`SpotWithPriceLimit\\`.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The bid for the spot instance. This parameter is in effect when \\`SpotStrategy\\` is set to \\`SpotWithPriceLimit\\`.
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
