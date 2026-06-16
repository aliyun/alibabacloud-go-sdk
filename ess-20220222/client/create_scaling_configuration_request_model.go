// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageOptions(v *CreateScalingConfigurationRequestImageOptions) *CreateScalingConfigurationRequest
	GetImageOptions() *CreateScalingConfigurationRequestImageOptions
	SetPrivatePoolOptions(v *CreateScalingConfigurationRequestPrivatePoolOptions) *CreateScalingConfigurationRequest
	GetPrivatePoolOptions() *CreateScalingConfigurationRequestPrivatePoolOptions
	SetSystemDisk(v *CreateScalingConfigurationRequestSystemDisk) *CreateScalingConfigurationRequest
	GetSystemDisk() *CreateScalingConfigurationRequestSystemDisk
	SetAffinity(v string) *CreateScalingConfigurationRequest
	GetAffinity() *string
	SetClientToken(v string) *CreateScalingConfigurationRequest
	GetClientToken() *string
	SetCpu(v int32) *CreateScalingConfigurationRequest
	GetCpu() *int32
	SetCreditSpecification(v string) *CreateScalingConfigurationRequest
	GetCreditSpecification() *string
	SetCustomPriorities(v []*CreateScalingConfigurationRequestCustomPriorities) *CreateScalingConfigurationRequest
	GetCustomPriorities() []*CreateScalingConfigurationRequestCustomPriorities
	SetDataDisks(v []*CreateScalingConfigurationRequestDataDisks) *CreateScalingConfigurationRequest
	GetDataDisks() []*CreateScalingConfigurationRequestDataDisks
	SetDedicatedHostClusterId(v string) *CreateScalingConfigurationRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *CreateScalingConfigurationRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *CreateScalingConfigurationRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *CreateScalingConfigurationRequest
	GetDeploymentSetId() *string
	SetHostName(v string) *CreateScalingConfigurationRequest
	GetHostName() *string
	SetHpcClusterId(v string) *CreateScalingConfigurationRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *CreateScalingConfigurationRequest
	GetHttpEndpoint() *string
	SetHttpTokens(v string) *CreateScalingConfigurationRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *CreateScalingConfigurationRequest
	GetImageFamily() *string
	SetImageId(v string) *CreateScalingConfigurationRequest
	GetImageId() *string
	SetImageName(v string) *CreateScalingConfigurationRequest
	GetImageName() *string
	SetInstanceDescription(v string) *CreateScalingConfigurationRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *CreateScalingConfigurationRequest
	GetInstanceName() *string
	SetInstancePatternInfos(v []*CreateScalingConfigurationRequestInstancePatternInfos) *CreateScalingConfigurationRequest
	GetInstancePatternInfos() []*CreateScalingConfigurationRequestInstancePatternInfos
	SetInstanceType(v string) *CreateScalingConfigurationRequest
	GetInstanceType() *string
	SetInstanceTypeCandidateOptions(v *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) *CreateScalingConfigurationRequest
	GetInstanceTypeCandidateOptions() *CreateScalingConfigurationRequestInstanceTypeCandidateOptions
	SetInstanceTypeOverrides(v []*CreateScalingConfigurationRequestInstanceTypeOverrides) *CreateScalingConfigurationRequest
	GetInstanceTypeOverrides() []*CreateScalingConfigurationRequestInstanceTypeOverrides
	SetInstanceTypes(v []*string) *CreateScalingConfigurationRequest
	GetInstanceTypes() []*string
	SetInternetChargeType(v string) *CreateScalingConfigurationRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateScalingConfigurationRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *CreateScalingConfigurationRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *CreateScalingConfigurationRequest
	GetKeyPairName() *string
	SetLoadBalancerWeight(v int32) *CreateScalingConfigurationRequest
	GetLoadBalancerWeight() *int32
	SetMemory(v int32) *CreateScalingConfigurationRequest
	GetMemory() *int32
	SetNetworkInterfaces(v []*CreateScalingConfigurationRequestNetworkInterfaces) *CreateScalingConfigurationRequest
	GetNetworkInterfaces() []*CreateScalingConfigurationRequestNetworkInterfaces
	SetOwnerAccount(v string) *CreateScalingConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateScalingConfigurationRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateScalingConfigurationRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *CreateScalingConfigurationRequest
	GetPasswordInherit() *bool
	SetRamRoleName(v string) *CreateScalingConfigurationRequest
	GetRamRoleName() *string
	SetResourceGroupId(v string) *CreateScalingConfigurationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateScalingConfigurationRequest
	GetResourceOwnerAccount() *string
	SetResourcePoolOptions(v *CreateScalingConfigurationRequestResourcePoolOptions) *CreateScalingConfigurationRequest
	GetResourcePoolOptions() *CreateScalingConfigurationRequestResourcePoolOptions
	SetScalingConfigurationName(v string) *CreateScalingConfigurationRequest
	GetScalingConfigurationName() *string
	SetScalingGroupId(v string) *CreateScalingConfigurationRequest
	GetScalingGroupId() *string
	SetSchedulerOptions(v map[string]interface{}) *CreateScalingConfigurationRequest
	GetSchedulerOptions() map[string]interface{}
	SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateScalingConfigurationRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateScalingConfigurationRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *CreateScalingConfigurationRequestSecurityOptions) *CreateScalingConfigurationRequest
	GetSecurityOptions() *CreateScalingConfigurationRequestSecurityOptions
	SetSpotDuration(v int32) *CreateScalingConfigurationRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *CreateScalingConfigurationRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimits(v []*CreateScalingConfigurationRequestSpotPriceLimits) *CreateScalingConfigurationRequest
	GetSpotPriceLimits() []*CreateScalingConfigurationRequestSpotPriceLimits
	SetSpotStrategy(v string) *CreateScalingConfigurationRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *CreateScalingConfigurationRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateScalingConfigurationRequest
	GetStorageSetPartitionNumber() *int32
	SetSystemDiskCategories(v []*string) *CreateScalingConfigurationRequest
	GetSystemDiskCategories() []*string
	SetTags(v string) *CreateScalingConfigurationRequest
	GetTags() *string
	SetTenancy(v string) *CreateScalingConfigurationRequest
	GetTenancy() *string
	SetUserData(v string) *CreateScalingConfigurationRequest
	GetUserData() *string
	SetZoneId(v string) *CreateScalingConfigurationRequest
	GetZoneId() *string
}

type CreateScalingConfigurationRequest struct {
	ImageOptions       *CreateScalingConfigurationRequestImageOptions       `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *CreateScalingConfigurationRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *CreateScalingConfigurationRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether to associate the instance with the dedicated host. Valid values:
	//
	// - default: The instance is not associated with the dedicated host. When you restart an instance that was stopped in economical mode, the instance may be placed on a different dedicated host in the automatic deployment resource pool if the original dedicated host has insufficient resources.
	//
	// - host: The instance is associated with the dedicated host. When you restart an instance that was stopped in economical mode, the instance is still placed on the original dedicated host. If the original dedicated host has insufficient resources, the instance fails to restart.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// Ensures the idempotence of the request. You can use the client to generate a unique parameter value to make sure that the same request is not repeated. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of vCPUs. Unit: cores.
	//
	// You can specify both Cpu and Memory to define a range of instance types. For example, if you set Cpu to 2 and Memory to 16, all instance types with 2 vCPUs and 16 GiB of memory are selected. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and then creates an instance of the instance type that has the lowest price.
	//
	// > This configuration is effective only when the cost optimization policy is enabled for the scaling group and no instance types are specified in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode.
	//
	// - Unlimited: unlimited mode.
	//
	// For more information, see the Performance modes section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The custom priority for the combination of an ECS instance type and a vSwitch.
	//
	// 	Notice: This is effective only when the scaling policy of the scaling group is set to the priority policy.
	//
	// If an instance cannot be created from a combination of an instance type and a vSwitch with a higher priority, Auto Scaling automatically tries the next combination in the list.
	//
	// > If you specify custom priorities for only some combinations of instance types and vSwitches, the unspecified combinations have a lower priority than the specified ones. The unspecified combinations are still prioritized based on the order of vSwitches in the scaling group and the order of instance types in the scaling configuration.
	//
	// >
	//
	// > - For example, if the vSwitch order in the scaling group is vsw1, vsw2, the instance type order in the scaling configuration is type1, type2, and the custom priority order is ["vsw2+type2", "vsw1+type2"], the final priority will be: "vsw2+type2" > "vsw1+type2" > "vsw1+type1" > "vsw2+type1".
	CustomPriorities []*CreateScalingConfigurationRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// A collection of data disk information.
	DataDisks []*CreateScalingConfigurationRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-2zedxc67zqzt7lb4****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// Specifies whether to create the ECS instance on a dedicated host. If you specify the DedicatedHostId parameter, the SpotStrategy and SpotPriceLimit settings in the request are automatically ignored. This is because dedicated hosts do not support creating preemptible instances.
	//
	// You can call the DescribeDedicatedHosts operation to query the list of dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection attribute of the instance. This specifies whether you can release the instance through the ECS console or by calling the DeleteInstance API. This prevents accidental release of the instance. Valid values:
	//
	// - true: enables release protection. You cannot release the instance through the ECS console or by calling the DeleteInstance API.
	//
	// - false: disables release protection. You can release the instance through the ECS console or by calling the DeleteInstance API.
	//
	// Default value: false.
	//
	// > This attribute applies only to pay-as-you-go instances to prevent accidental release of instances scaled out by Auto Scaling. It does not affect normal scale-in activities. Instances with release protection enabled can be released during scale-in activities.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set to which the ECS instance belongs.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pz****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the ECS instance. A period (.) or a hyphen (-) cannot be used as the first or last character of a hostname. Consecutive periods (.) or hyphens (-) are not allowed. The naming conventions for different instance types are as follows:
	//
	// - For Windows instances, the hostname must be 2 to 15 characters in length and can contain letters, digits, and hyphens (-). It cannot contain periods (.) or be all digits.
	//
	// - For other instance types, such as Linux, the hostname must be 2 to 64 characters in length. You can use periods (.) to separate the hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// host****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the High Performance Computing (HPC) cluster to which the ECS instance belongs.
	//
	// example:
	//
	// hpc-clusterid
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enables the channel.
	//
	// - disabled: disables the channel.
	//
	// Default value: enabled.
	//
	// > For information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Specifies whether to enforce the security-hardened mode (IMDSv2) when accessing instance metadata. Valid values:
	//
	// - optional: does not enforce the mode.
	//
	// - required: enforces the mode. If you set this value, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// > For information about instance metadata access modes, see [Instance metadata access modes](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image from the specified image family to create the instance. If you have specified the `ImageId` parameter, you cannot set this parameter.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file to use for creating instances.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image file. Image names must be unique within a region. If you specify ImageId, ImageName is ignored.
	//
	// You cannot use ImageName to specify a Marketplace image.
	//
	// example:
	//
	// image****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The description of the ECS instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the ECS instances that are created using the scaling configuration.
	//
	// example:
	//
	// instance****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// A collection of intelligent configuration information used to filter instance types that meet the requirements.
	InstancePatternInfos []*CreateScalingConfigurationRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance type of the ECS instance. For more information, see Instance families.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// After you enable the candidate mode, if issues such as insufficient inventory occur, the system supplements the currently selected instance types with similar-sized alternatives or creates vSwitches in candidate zones and adds them to the scaling group.
	InstanceTypeCandidateOptions *CreateScalingConfigurationRequestInstanceTypeCandidateOptions `json:"InstanceTypeCandidateOptions,omitempty" xml:"InstanceTypeCandidateOptions,omitempty" type:"Struct"`
	// Information about the specified instance types.
	InstanceTypeOverrides []*CreateScalingConfigurationRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	// Multiple instance types. If you use InstanceTypes, InstanceType is ignored.
	//
	// If an instance cannot be created from an instance type with a higher priority, Auto Scaling automatically tries the next instance type in the list.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth. InternetMaxBandwidthOut specifies the fixed bandwidth.
	//
	// - PayByTraffic: pay-by-data-transfer. InternetMaxBandwidthOut specifies the maximum bandwidth. You are charged for the actual data transfer.
	//
	// If you do not specify this parameter, the default value is PayByBandwidth for classic network and PayByTraffic for VPC.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// - If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - If the purchased outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `InternetMaxBandwidthOut`. Default value: the value of `InternetMaxBandwidthOut`.
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
	// none: The instance is not I/O optimized.
	//
	// optimized: The instance is I/O optimized.
	//
	// For retired instance types, the default value is none. For other instance types, the default value is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses to assign to the ENI.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair to use to log on to the ECS instance.
	//
	// - For Windows instances, this parameter is ignored. The default value is empty.
	//
	// - For Linux instances, password-based logon is disabled.
	//
	// example:
	//
	// KeyPairTest
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The weight of the ECS instance as a backend server of the load balancer. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify both Cpu and Memory to define a range of instance types. For example, if you set Cpu to 2 and Memory to 16, all instance types with 2 vCPUs and 16 GiB of memory are selected. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and then creates an instance of the instance type that has the lowest price.
	//
	// > This configuration is effective only when the cost optimization policy is enabled for the scaling group and no instance types are specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The list of ENIs.
	NetworkInterfaces []*CreateScalingConfigurationRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	OwnerAccount      *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters can be:
	//
	// \\`()` ~!@#$%^&*-_+=\\|{}[]:;\\"<>,.?/`
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify the Password parameter, we recommend that you send the request over HTTPS to prevent password leaks.
	//
	// example:
	//
	// 123abc****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. If you use this parameter, make sure that a password is preset in the image. Valid values:
	//
	// - true: uses the preset password.
	//
	// - false: does not use the preset password.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the RAM role for the ECS instance. The name is provided and maintained by RAM. You can call the ListRoles operation to query the available RAM roles.
	//
	// example:
	//
	// ramrole****
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the ECS instance belongs.
	//
	// example:
	//
	// rg-resource****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource pool policy to use when creating an instance. Note the following when you set this parameter:
	//
	// - This parameter takes effect only when creating pay-as-you-go instances.
	//
	// - You cannot set this parameter at the same time as PrivatePoolOptions.MatchCriteria and PrivatePoolOptions.Id.
	ResourcePoolOptions *CreateScalingConfigurationRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
	// The name of the scaling configuration. The name must be 2 to 64 characters in length, and can contain digits, underscores (_), hyphens (-), and periods (.). It must start with a digit, a letter, or a Chinese character.
	//
	// The name of a scaling configuration must be unique within a scaling group in a region. If you do not specify this parameter, the ID of the scaling configuration is used.
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
	// The scheduling options.
	//
	// example:
	//
	// ["testManagedPrivateSpaceId****"]
	SchedulerOptions map[string]interface{} `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// - Active: enables security hardening. This setting is valid only for public images.
	//
	// - Deactive: disables security hardening. This setting is valid for all image types.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the ECS instance belongs. ECS instances in the same security group can communicate with each other.
	//
	// example:
	//
	// sg-280ih****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Adds the ECS instance to multiple security groups. For more information, see the Security groups section in [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > You cannot specify both SecurityGroupId and SecurityGroupIds.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *CreateScalingConfigurationRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the preemptible instance. Unit: hours. Valid values:
	//
	// - 1: The instance is retained for 1 hour after it is created. After 1 hour, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// - 0: The instance is not guaranteed to be retained for 1 hour after it is created. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// > Alibaba Cloud sends a notification to you through an ECS system event 5 minutes before the instance is released. Preemptible instances are billed by the second. We recommend that you select a protection period based on the time required for your task to complete.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the preemptible instance. Currently, only terminate is supported, which releases the instance by default.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// A collection of billing information for preemptible instances.
	SpotPriceLimits []*CreateScalingConfigurationRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy for pay-as-you-go instances. Valid values:
	//
	// - NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: The instance is a preemptible instance with a user-defined maximum hourly price.
	//
	// - SpotAsPriceGo: The instance is a preemptible instance for which the price is automatically bid based on the current market price.
	//
	// Default value: NoSpot.
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
	// The maximum number of partitions in the storage set. The value must be greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// Multiple disk categories for the system disk. When a disk of a higher-priority category is unavailable, Auto Scaling automatically tries the next lower-priority category to create the system disk. Valid values:
	//
	// - cloud: basic disk.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// > If you specify this parameter, you cannot specify `SystemDisk.Category`.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The tags of the ECS instance. You can specify up to 20 tags in key-value pairs. The following limits apply to keys and values:
	//
	// - A key can be up to 64 characters in length and cannot start with aliyun or acs:. It cannot contain `http://` or `https://`. If you use tags, the key cannot be an empty string.
	//
	// - A value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain `http://` or `https://`. The value can be an empty string.
	//
	// example:
	//
	// {"key1":"value1","key2":"value2", ... "key5":"value5"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// - default: creates the instance on a non-dedicated host.
	//
	// - host: creates the instance on a dedicated host. If you do not specify DedicatedHostId, Alibaba Cloud automatically selects a dedicated host for the instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// The custom data of the ECS instance. The data must be Base64 encoded. The raw data can be up to 32 KB in size.
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

func (s CreateScalingConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequest) GetImageOptions() *CreateScalingConfigurationRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateScalingConfigurationRequest) GetPrivatePoolOptions() *CreateScalingConfigurationRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateScalingConfigurationRequest) GetSystemDisk() *CreateScalingConfigurationRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateScalingConfigurationRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *CreateScalingConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateScalingConfigurationRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateScalingConfigurationRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateScalingConfigurationRequest) GetCustomPriorities() []*CreateScalingConfigurationRequestCustomPriorities {
	return s.CustomPriorities
}

func (s *CreateScalingConfigurationRequest) GetDataDisks() []*CreateScalingConfigurationRequestDataDisks {
	return s.DataDisks
}

func (s *CreateScalingConfigurationRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *CreateScalingConfigurationRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateScalingConfigurationRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateScalingConfigurationRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateScalingConfigurationRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateScalingConfigurationRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *CreateScalingConfigurationRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateScalingConfigurationRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateScalingConfigurationRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateScalingConfigurationRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateScalingConfigurationRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateScalingConfigurationRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateScalingConfigurationRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateScalingConfigurationRequest) GetInstancePatternInfos() []*CreateScalingConfigurationRequestInstancePatternInfos {
	return s.InstancePatternInfos
}

func (s *CreateScalingConfigurationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationRequest) GetInstanceTypeCandidateOptions() *CreateScalingConfigurationRequestInstanceTypeCandidateOptions {
	return s.InstanceTypeCandidateOptions
}

func (s *CreateScalingConfigurationRequest) GetInstanceTypeOverrides() []*CreateScalingConfigurationRequestInstanceTypeOverrides {
	return s.InstanceTypeOverrides
}

func (s *CreateScalingConfigurationRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateScalingConfigurationRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateScalingConfigurationRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateScalingConfigurationRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateScalingConfigurationRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateScalingConfigurationRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateScalingConfigurationRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateScalingConfigurationRequest) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *CreateScalingConfigurationRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateScalingConfigurationRequest) GetNetworkInterfaces() []*CreateScalingConfigurationRequestNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *CreateScalingConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateScalingConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateScalingConfigurationRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateScalingConfigurationRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateScalingConfigurationRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateScalingConfigurationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateScalingConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateScalingConfigurationRequest) GetResourcePoolOptions() *CreateScalingConfigurationRequestResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *CreateScalingConfigurationRequest) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *CreateScalingConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateScalingConfigurationRequest) GetSchedulerOptions() map[string]interface{} {
	return s.SchedulerOptions
}

func (s *CreateScalingConfigurationRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateScalingConfigurationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateScalingConfigurationRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateScalingConfigurationRequest) GetSecurityOptions() *CreateScalingConfigurationRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateScalingConfigurationRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateScalingConfigurationRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateScalingConfigurationRequest) GetSpotPriceLimits() []*CreateScalingConfigurationRequestSpotPriceLimits {
	return s.SpotPriceLimits
}

func (s *CreateScalingConfigurationRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateScalingConfigurationRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateScalingConfigurationRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateScalingConfigurationRequest) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *CreateScalingConfigurationRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateScalingConfigurationRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *CreateScalingConfigurationRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateScalingConfigurationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateScalingConfigurationRequest) SetImageOptions(v *CreateScalingConfigurationRequestImageOptions) *CreateScalingConfigurationRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPrivatePoolOptions(v *CreateScalingConfigurationRequestPrivatePoolOptions) *CreateScalingConfigurationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSystemDisk(v *CreateScalingConfigurationRequestSystemDisk) *CreateScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetAffinity(v string) *CreateScalingConfigurationRequest {
	s.Affinity = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetClientToken(v string) *CreateScalingConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCpu(v int32) *CreateScalingConfigurationRequest {
	s.Cpu = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCreditSpecification(v string) *CreateScalingConfigurationRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetCustomPriorities(v []*CreateScalingConfigurationRequestCustomPriorities) *CreateScalingConfigurationRequest {
	s.CustomPriorities = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDataDisks(v []*CreateScalingConfigurationRequestDataDisks) *CreateScalingConfigurationRequest {
	s.DataDisks = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDedicatedHostClusterId(v string) *CreateScalingConfigurationRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDedicatedHostId(v string) *CreateScalingConfigurationRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDeletionProtection(v bool) *CreateScalingConfigurationRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetDeploymentSetId(v string) *CreateScalingConfigurationRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHostName(v string) *CreateScalingConfigurationRequest {
	s.HostName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHpcClusterId(v string) *CreateScalingConfigurationRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHttpEndpoint(v string) *CreateScalingConfigurationRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetHttpTokens(v string) *CreateScalingConfigurationRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageFamily(v string) *CreateScalingConfigurationRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageId(v string) *CreateScalingConfigurationRequest {
	s.ImageId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageName(v string) *CreateScalingConfigurationRequest {
	s.ImageName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceDescription(v string) *CreateScalingConfigurationRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceName(v string) *CreateScalingConfigurationRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstancePatternInfos(v []*CreateScalingConfigurationRequestInstancePatternInfos) *CreateScalingConfigurationRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceType(v string) *CreateScalingConfigurationRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceTypeCandidateOptions(v *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) *CreateScalingConfigurationRequest {
	s.InstanceTypeCandidateOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceTypeOverrides(v []*CreateScalingConfigurationRequestInstanceTypeOverrides) *CreateScalingConfigurationRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceTypes(v []*string) *CreateScalingConfigurationRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetChargeType(v string) *CreateScalingConfigurationRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIoOptimized(v string) *CreateScalingConfigurationRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIpv6AddressCount(v int32) *CreateScalingConfigurationRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetKeyPairName(v string) *CreateScalingConfigurationRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetLoadBalancerWeight(v int32) *CreateScalingConfigurationRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetMemory(v int32) *CreateScalingConfigurationRequest {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetNetworkInterfaces(v []*CreateScalingConfigurationRequestNetworkInterfaces) *CreateScalingConfigurationRequest {
	s.NetworkInterfaces = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerId(v int64) *CreateScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPassword(v string) *CreateScalingConfigurationRequest {
	s.Password = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetPasswordInherit(v bool) *CreateScalingConfigurationRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetRamRoleName(v string) *CreateScalingConfigurationRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceGroupId(v string) *CreateScalingConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourcePoolOptions(v *CreateScalingConfigurationRequestResourcePoolOptions) *CreateScalingConfigurationRequest {
	s.ResourcePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingGroupId(v string) *CreateScalingConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSchedulerOptions(v map[string]interface{}) *CreateScalingConfigurationRequest {
	s.SchedulerOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityOptions(v *CreateScalingConfigurationRequestSecurityOptions) *CreateScalingConfigurationRequest {
	s.SecurityOptions = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotDuration(v int32) *CreateScalingConfigurationRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotInterruptionBehavior(v string) *CreateScalingConfigurationRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotPriceLimits(v []*CreateScalingConfigurationRequestSpotPriceLimits) *CreateScalingConfigurationRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSpotStrategy(v string) *CreateScalingConfigurationRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetStorageSetId(v string) *CreateScalingConfigurationRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetStorageSetPartitionNumber(v int32) *CreateScalingConfigurationRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSystemDiskCategories(v []*string) *CreateScalingConfigurationRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetTags(v string) *CreateScalingConfigurationRequest {
	s.Tags = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetTenancy(v string) *CreateScalingConfigurationRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetUserData(v string) *CreateScalingConfigurationRequest {
	s.UserData = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetZoneId(v string) *CreateScalingConfigurationRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) Validate() error {
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

type CreateScalingConfigurationRequestImageOptions struct {
	// Specifies whether to use the ecs-user user to log on to the ECS instance. For more information, see [Manage logon usernames of ECS instances](https://help.aliyun.com/document_detail/388447.html). Valid values:
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
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s CreateScalingConfigurationRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateScalingConfigurationRequestImageOptions) SetLoginAsNonRoot(v bool) *CreateScalingConfigurationRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateScalingConfigurationRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestPrivatePoolOptions struct {
	// The ID of the private pool. This is the ID of the Elastic Assurance or Capacity Reservation service.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool capacity option for instance startup. After an Elastic Assurance or a Capacity Reservation service takes effect, a private pool is generated. You can select a private pool to start an instance. Valid values:
	//
	// - Open: open mode. The system automatically matches the instance with an open private pool. If no eligible private pool is found, the instance is started using public pool resources. You do not need to set the PrivatePoolOptions.Id parameter in this mode.
	//
	// - Target: specified mode. The instance is started using the capacity of a specified private pool. If the specified private pool is unavailable, the instance fails to start. You must set the PrivatePoolOptions.Id parameter in this mode.
	//
	// - None: no private pool is used. The instance is started without using a private pool.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateScalingConfigurationRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) SetId(v string) *CreateScalingConfigurationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateScalingConfigurationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateScalingConfigurationRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestSystemDisk struct {
	// The ID of the automatic snapshot policy to apply to the system disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// > You can set this parameter only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">
	//
	// For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
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
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// You cannot specify this parameter and `SystemDiskCategories` at the same time. If you do not specify this parameter or `SystemDiskCategories`, the default value of this parameter is used:
	//
	// - For I/O optimized instances, the default value is cloud_efficiency.
	//
	// - For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test system disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The encryption algorithm for the system disk. Valid values:
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
	// - true: encrypts the disk.
	//
	// - false: does not encrypt the disk.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key to use for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS of the system disk.
	//
	// > IOPS (input/output operations per second) indicates the number of I/O operations that a block storage device can process per second. It represents the read and write capabilities of the device.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// - Basic disk: 20 to 500.
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
	// - ESSD AutoPL disk: 1 to 2048.
	//
	// - Other disk categories: 20 to 2048.
	//
	// The value of this parameter must be greater than or equal to max{1, ImageSize}.
	//
	// Default value: max{40, ImageSize}.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateScalingConfigurationRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateScalingConfigurationRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetBurstingEnabled(v bool) *CreateScalingConfigurationRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetDescription(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetDiskName(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetEncryptAlgorithm(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetEncrypted(v bool) *CreateScalingConfigurationRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetKMSKeyId(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetProvisionedIops(v int64) *CreateScalingConfigurationRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestCustomPriorities struct {
	// The instance type of the ECS instance.
	//
	// 	Notice:
	//
	// Must be included in the list of instance types in the scaling configuration.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the vSwitch.
	//
	// 	Notice:
	//
	// Must be included in the list of vSwitches in the scaling group.
	//
	// example:
	//
	// vsw-bp14zolna43z266bq****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateScalingConfigurationRequestCustomPriorities) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestCustomPriorities) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestCustomPriorities) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationRequestCustomPriorities) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateScalingConfigurationRequestCustomPriorities) SetInstanceType(v string) *CreateScalingConfigurationRequestCustomPriorities {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestCustomPriorities) SetVswitchId(v string) *CreateScalingConfigurationRequestCustomPriorities {
	s.VswitchId = &v
	return s
}

func (s *CreateScalingConfigurationRequestCustomPriorities) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestDataDisks struct {
	// The ID of the automatic snapshot policy to apply to the data disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable performance burst for the system disk. Valid values:
	//
	// - true: Enables performance burst.
	//
	// - false: Disables performance burst.
	//
	// > This parameter is available only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">
	//
	// For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Multiple disk categories for the data disk. When a disk of a higher-priority category is unavailable, Auto Scaling automatically tries the next lower-priority category. Valid values:
	//
	// - cloud: basic disk. The DeleteWithInstance attribute of a basic disk created with an instance is true.
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: ESSD.
	//
	// > If you specify this parameter, you cannot specify `DataDisks.Category`.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data disk. Valid values:
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
	// - cloud_auto: ESSD AutoPL disk.
	//
	// You cannot specify this parameter and DataDisk.Categories at the same time. If you do not specify this parameter or DataDisk.Categories, the default value of this parameter is used:
	//
	// - For I/O optimized instances, the default value is cloud_efficiency.
	//
	// - For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// - true: releases the disk when the instance is released.
	//
	// - false: retains the disk when the instance is released.
	//
	// This parameter can be set only for separately purchased cloud disks (DataDisks.Category is cloud, cloud_efficiency, cloud_ssd, or cloud_essd). Otherwise, an error occurs.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test data disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. If you do not specify this parameter, the system automatically assigns a mount point when creating the ECS instance, starting from /dev/xvdb and ending at /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// - true: encrypts the disk.
	//
	// - false: does not encrypt the disk.
	//
	// Default value: false.
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
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// > For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS of the data disk.
	//
	// > IOPS (input/output operations per second) indicates the number of I/O operations that a block storage device can process per second. It represents the read and write capabilities of the device.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the data disk. Unit: GiB. Valid values:
	//
	// - cloud: 5 to 2000.
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_essd: 20 to 32768.
	//
	// - ephemeral_ssd: 5 to 800.
	//
	// If you specify this parameter, the disk size must be greater than or equal to the snapshot size specified by SnapshotId.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot to use to create the data disk. If you specify this parameter, DataDisks.Size is ignored. The size of the created disk is the size of the specified snapshot.
	//
	// If the snapshot was created on or before July 15, 2013, the call is rejected and an InvalidSnapshot.TooOld error is returned.
	//
	// example:
	//
	// s-280s7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateScalingConfigurationRequestDataDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestDataDisks) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestDataDisks) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateScalingConfigurationRequestDataDisks) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateScalingConfigurationRequestDataDisks) GetCategories() []*string {
	return s.Categories
}

func (s *CreateScalingConfigurationRequestDataDisks) GetCategory() *string {
	return s.Category
}

func (s *CreateScalingConfigurationRequestDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateScalingConfigurationRequestDataDisks) GetDescription() *string {
	return s.Description
}

func (s *CreateScalingConfigurationRequestDataDisks) GetDevice() *string {
	return s.Device
}

func (s *CreateScalingConfigurationRequestDataDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateScalingConfigurationRequestDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateScalingConfigurationRequestDataDisks) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateScalingConfigurationRequestDataDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateScalingConfigurationRequestDataDisks) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateScalingConfigurationRequestDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *CreateScalingConfigurationRequestDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateScalingConfigurationRequestDataDisks) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetBurstingEnabled(v bool) *CreateScalingConfigurationRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetCategories(v []*string) *CreateScalingConfigurationRequestDataDisks {
	s.Categories = v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetCategory(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDeleteWithInstance(v bool) *CreateScalingConfigurationRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDescription(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDevice(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetDiskName(v string) *CreateScalingConfigurationRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetEncrypted(v string) *CreateScalingConfigurationRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetKMSKeyId(v string) *CreateScalingConfigurationRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetPerformanceLevel(v string) *CreateScalingConfigurationRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetProvisionedIops(v int64) *CreateScalingConfigurationRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetSize(v int32) *CreateScalingConfigurationRequestDataDisks {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) SetSnapshotId(v string) *CreateScalingConfigurationRequestDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisks) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestInstancePatternInfos struct {
	// The architecture type of the instance type. Valid values:
	//
	// - X86: X86 compute.
	//
	// - Heterogeneous: heterogeneous computing, such as GPU or FPGA.
	//
	// - BareMental: ECS Bare Metal Instance.
	//
	// - Arm: Arm compute.
	//
	// Default value: all architecture types are included.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether to include burstable instance types. Valid values:
	//
	// - Exclude: does not include burstable instance types.
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
	// In intelligent configuration mode, the number of vCPU cores of the instance type. This is used to filter instance types that meet the requirements. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// Note the following:
	//
	// - The InstancePatternInfos parameter applies only when the network type of the scaling group is VPC.
	//
	// - When using InstancePatternInfos, you must specify both InstancePatternInfos.Cores and InstancePatternInfos.Memory.
	//
	// - If you have already specified instance types using the InstanceType or InstanceTypes parameter, Auto Scaling prioritizes the specified instance types for scale-outs. If the specified instance types are out of stock, Auto Scaling uses the instance type with the lowest price from the instance types that match the InstancePatternInfos parameter.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architecture of the instance. Valid values:
	//
	// > You can specify multiple CPU architectures. N is an integer from 1 to 2.
	//
	// - X86.
	//
	// - ARM.
	CpuArchitectures []*string `json:"CpuArchitectures,omitempty" xml:"CpuArchitectures,omitempty" type:"Repeated"`
	// The instance types to exclude. You can use a wildcard character (\\*) to exclude a single instance type or an entire instance family. For example:
	//
	// - ecs.c6.large: excludes the ecs.c6.large instance type.
	//
	// - ecs.c6.\\*: excludes all instance types in the c6 family.
	ExcludedInstanceTypes []*string `json:"ExcludedInstanceTypes,omitempty" xml:"ExcludedInstanceTypes,omitempty" type:"Repeated"`
	// The GPU type.
	GpuSpecs []*string `json:"GpuSpecs,omitempty" xml:"GpuSpecs,omitempty" type:"Repeated"`
	// The category of the instance. Valid values:
	//
	// > You can specify multiple instance categories. N is an integer from 1 to 10.
	//
	// - General-purpose: General-purpose.
	//
	// - Compute-optimized: compute-optimized.
	//
	// - Memory-optimized: memory-optimized.
	//
	// - Big data: big data.
	//
	// - Local SSDs : local SSDs.
	//
	// - High Clock Speed : high frequency.
	//
	// - Enhanced : enhanced instance families.
	//
	// - Shared: shared-resource instances.
	//
	// - Compute-optimized with GPU : GPU compute-optimized.
	//
	// - Visual Compute-optimized : visual compute-optimized.
	//
	// - Heterogeneous Service : heterogeneous service.
	//
	// - Compute-optimized with FPGA : FPGA compute-optimized.
	//
	// - Compute-optimized with NPU : NPU compute-optimized.
	//
	// - ECS Bare Metal : ECS Bare Metal Instance.
	//
	// - High Performance Compute: high-performance computing.
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family, used to filter instance types that meet the requirements. This parameter takes effect when `CostOptimization` is enabled. Valid values:
	//
	// - EntryLevel: entry-level, which refers to shared-resource instances. This level offers lower costs but cannot guarantee stable computing performance. It is suitable for business scenarios with low CPU utilization. For more information, see [Shared-resource instances](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise-level. This level provides stable performance and dedicated resources, suitable for business scenarios that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit entry-level, which refers to burstable instances. This level uses CPU credits to ensure computing performance and is suitable for business scenarios with low CPU utilization and occasional CPU bursts. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families to query. You can specify multiple instance families. N is an integer from 1 to 10.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// In intelligent configuration mode, the maximum hourly price you are willing to pay for a pay-as-you-go or preemptible instance. This is used to filter instance types that meet the requirements.
	//
	// > This parameter is required when SpotStrategy is set to SpotWithPriceLimit. In other cases, this parameter is optional.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPU cores for the instance type.
	//
	// > MaximumCpuCoreCount cannot be more than four times the value of MinimumCpuCoreCount.
	//
	// example:
	//
	// 4
	MaximumCpuCoreCount *int32 `json:"MaximumCpuCoreCount,omitempty" xml:"MaximumCpuCoreCount,omitempty"`
	// The maximum number of GPUs for the instance. The value must be a positive integer.
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
	// In intelligent configuration mode, the memory size of the instance type. Unit: GiB. This is used to filter instance types that meet the requirements.
	//
	// example:
	//
	// 4
	Memory *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The minimum baseline vCPU performance (the sum of all vCPUs) for burstable instances of the t5 or t6 family.
	//
	// example:
	//
	// 12
	MinimumBaselineCredit *int32 `json:"MinimumBaselineCredit,omitempty" xml:"MinimumBaselineCredit,omitempty"`
	// The minimum number of vCPU cores for the instance type.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum number of IPv6 addresses that can be assigned to a single ENI.
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum number of private IPv4 addresses to assign to a single elastic network interface (ENI) of an instance.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum number of elastic network interfaces (ENIs) that can be attached to an instance.
	//
	// example:
	//
	// 2
	MinimumEniQuantity *int32 `json:"MinimumEniQuantity,omitempty" xml:"MinimumEniQuantity,omitempty"`
	// The minimum number of GPUs for the instance. The value must be a positive integer.
	//
	// example:
	//
	// 2
	MinimumGpuAmount *int32 `json:"MinimumGpuAmount,omitempty" xml:"MinimumGpuAmount,omitempty"`
	// The minimum initial vCPU credit for burstable instances of the t5 or t6 family.
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
	// The processor model of the instance. You can specify multiple processor models. N is an integer from 1 to 10.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
}

func (s CreateScalingConfigurationRequestInstancePatternInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetArchitectures() []*string {
	return s.Architectures
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetCores() *int32 {
	return s.Cores
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMaximumGpuAmount() *int32 {
	return s.MaximumGpuAmount
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumGpuAmount() *int32 {
	return s.MinimumGpuAmount
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetArchitectures(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetBurstablePerformance(v string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetCores(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetCpuArchitectures(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.CpuArchitectures = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetGpuSpecs(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.GpuSpecs = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetInstanceCategories(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.InstanceCategories = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetInstanceTypeFamilies(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.InstanceTypeFamilies = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMaxPrice(v float32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMaximumCpuCoreCount(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMaximumGpuAmount(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MaximumGpuAmount = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMaximumMemorySize(v float32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MaximumMemorySize = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMemory(v float32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumBaselineCredit(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumCpuCoreCount(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumEniIpv6AddressQuantity(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumEniPrivateIpAddressQuantity(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumEniQuantity(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumEniQuantity = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumGpuAmount(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumGpuAmount = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumInitialCredit(v int32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumInitialCredit = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetMinimumMemorySize(v float32) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.MinimumMemorySize = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) SetPhysicalProcessorModels(v []*string) *CreateScalingConfigurationRequestInstancePatternInfos {
	s.PhysicalProcessorModels = v
	return s
}

func (s *CreateScalingConfigurationRequestInstancePatternInfos) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestInstanceTypeCandidateOptions struct {
	// When supplementing with vSwitches from other zones is allowed, you must specify the CIDR blocks for the candidate vSwitches.
	AllowCidrBlocks []*string `json:"AllowCidrBlocks,omitempty" xml:"AllowCidrBlocks,omitempty" type:"Repeated"`
	// Specifies whether to allow supplementing with vSwitches from other zones.
	//
	// > The instance type remains unchanged; only new zones are considered as candidates. When a scaling group cannot scale out in any of the selected zones due to issues like insufficient inventory, this configuration allows ESS to automatically add a vSwitch from a new zone to the scaling group.
	//
	// > For example, if a scaling group is configured for zones cn-hangzhou-h and cn-hangzhou-g, and neither can be scaled out, ESS might create and add a vSwitch in cn-hangzhou-k based on real-time inventory.
	//
	// example:
	//
	// true
	AllowCrossAz *bool `json:"AllowCrossAz,omitempty" xml:"AllowCrossAz,omitempty"`
	// Specifies whether to allow supplementing with instance types from other generations.
	//
	// - For example, if the current instance type is ecs.c7.large, enabling this allows ecs.c6.large and ecs.c8.large as candidate types.
	//
	// example:
	//
	// true
	AllowDifferentGeneration *bool `json:"AllowDifferentGeneration,omitempty" xml:"AllowDifferentGeneration,omitempty"`
	// Specifies whether to enable the candidate mode.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The maximum price for candidate instance types.
	//
	// example:
	//
	// 2.10
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
}

func (s CreateScalingConfigurationRequestInstanceTypeCandidateOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstanceTypeCandidateOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) GetAllowCidrBlocks() []*string {
	return s.AllowCidrBlocks
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) GetAllowCrossAz() *bool {
	return s.AllowCrossAz
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) GetAllowDifferentGeneration() *bool {
	return s.AllowDifferentGeneration
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) SetAllowCidrBlocks(v []*string) *CreateScalingConfigurationRequestInstanceTypeCandidateOptions {
	s.AllowCidrBlocks = v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) SetAllowCrossAz(v bool) *CreateScalingConfigurationRequestInstanceTypeCandidateOptions {
	s.AllowCrossAz = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) SetAllowDifferentGeneration(v bool) *CreateScalingConfigurationRequestInstanceTypeCandidateOptions {
	s.AllowDifferentGeneration = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) SetEnabled(v bool) *CreateScalingConfigurationRequestInstanceTypeCandidateOptions {
	s.Enabled = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) SetMaxPrice(v float32) *CreateScalingConfigurationRequestInstanceTypeCandidateOptions {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeCandidateOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestInstanceTypeOverrides struct {
	// If you want the scaling group to scale based on the capacity of instance types, specify both this parameter and WeightedCapacity.
	//
	// This parameter specifies the instance type and overwrites the instance type in the launch template. You can specify this parameter up to 20 times. N is an integer from 1 to 20.
	//
	// > This parameter takes effect only when a launch template is specified by the LaunchTemplateId parameter.
	//
	// Valid values for InstanceType: available ECS instance types.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// To specify the capacity of an instance type in the scaling configuration, specify this parameter after you specify InstanceTypeOverrides.InstanceType.
	//
	// This parameter specifies the weight of the instance type. The weight indicates the capacity of a single instance of the specified instance type in the scaling group. A higher weight means that fewer instances of that type are needed to meet the desired capacity.
	//
	// You can configure different weights for different instance types based on their performance metrics, such as the number of vCPUs and memory size.
	//
	// For example:
	//
	// - Current capacity: 0.
	//
	// - Expected capacity: 6.
	//
	// - Capacity of ecs.c5.xlarge: 4.
	//
	// To meet the expected capacity, the scaling group will scale out two ecs.c5.xlarge instances.
	//
	// > During a scale-out, the capacity of the scaling group cannot exceed the sum of MaxSize and the maximum weight of the instance type.
	//
	// Valid values for WeightedCapacity: 1 to 500.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingConfigurationRequestInstanceTypeOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) SetInstanceType(v string) *CreateScalingConfigurationRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *CreateScalingConfigurationRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *CreateScalingConfigurationRequestInstanceTypeOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestNetworkInterfaces struct {
	// The type of ENI. When using this parameter, you must use NetworkInterfaces to set the primary NIC and cannot set the SecurityGroupId or SecurityGroupIds parameters. Valid values:
	//
	// - Primary: primary NIC.
	//
	// - Secondary: secondary NIC.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses to assign to the primary NIC.
	//
	// Note:
	//
	// - This parameter takes effect only when NetworkInterface.InstanceType is set to Primary. If NetworkInterface.InstanceType is set to Secondary or is empty, you cannot set this parameter.
	//
	// - If you set this parameter, you cannot set Ipv6AddressCount.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication mode of the NIC. Valid values:
	//
	// - Standard: uses TCP communication mode.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses RDMA communication mode.
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
	// - The value cannot exceed the IP address limit for the instance type. For more information, see [Instance families](https://help.aliyun.com/en/ecs/user-guide/overview-of-instance-families).
	//
	// - NetworkInterface.N.SecondaryPrivateIpAddressCount is used to assign a number of secondary private IPv4 addresses to the NIC (excluding the primary private IP address). The system randomly assigns these addresses from the available CIDR block of the vSwitch (NetworkInterface.N.VSwitchId) where the NIC is located.
	//
	// example:
	//
	// 6
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// One or more security group IDs to which the ENI belongs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s CreateScalingConfigurationRequestNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) SetInstanceType(v string) *CreateScalingConfigurationRequestNetworkInterfaces {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) SetIpv6AddressCount(v int32) *CreateScalingConfigurationRequestNetworkInterfaces {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) SetNetworkInterfaceTrafficMode(v string) *CreateScalingConfigurationRequestNetworkInterfaces {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) SetSecondaryPrivateIpAddressCount(v int32) *CreateScalingConfigurationRequestNetworkInterfaces {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationRequestNetworkInterfaces {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestResourcePoolOptions struct {
	// The ID of the private pool. This is the ID of the Elastic Assurance or Capacity Reservation service. This parameter can only accept Target mode private pool IDs and cannot be specified at the same time as the PrivatePoolTags parameter.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// Filter available Target private pools by tags. N is an integer from 1 to 20.
	//
	// Note:
	//
	// - When this parameter is configured, the system only selects from the associated Target private pools under the account that have matching tags and also meet the scaling group constraints (such as zone, instance type, etc.).
	//
	// - Tag matching rule: The private pool must match all specified tags.
	//
	// - This parameter cannot be specified at the same time as the PrivatePoolIds parameter.
	PrivatePoolTags []*CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags `json:"PrivatePoolTags,omitempty" xml:"PrivatePoolTags,omitempty" type:"Repeated"`
	// The resource pool, which includes private pools generated after an Elastic Assurance or Capacity Reservation service takes effect, and the public pool, can be selected for instance startup. Valid values:
	//
	// - PrivatePoolFirst: Private pool first. When this policy is selected, if you specify ResourcePoolOptions.PrivatePoolIds or meet the PrivatePoolTags conditions, the corresponding private pool is used first. If no private pool is specified or the specified private pool has insufficient capacity, the system automatically matches an open-type private pool. If no eligible private pool is found, the instance is created using the public pool.
	//
	// - PrivatePoolOnly: Private pool only. When this policy is selected, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pool has insufficient capacity, the instance fails to start.
	//
	// - PublicPoolFirst: Public pool resources first. The system prioritizes creating the instance using the public pool. When public pool resources are insufficient, private pool resources are used as a supplement. The system first automatically matches an open-type private pool. If no eligible private pool is found, it uses a Target-type private pool that is specified by ResourcePoolOptions.PrivatePoolIds or meets the PrivatePoolTags conditions. (This policy is in invitational preview and is not yet available for use.)
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

func (s CreateScalingConfigurationRequestResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) GetPrivatePoolTags() []*CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags {
	return s.PrivatePoolTags
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *CreateScalingConfigurationRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) SetPrivatePoolTags(v []*CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) *CreateScalingConfigurationRequestResourcePoolOptions {
	s.PrivatePoolTags = v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) SetStrategy(v string) *CreateScalingConfigurationRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) Validate() error {
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

type CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags struct {
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

func (s CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) GetKey() *string {
	return s.Key
}

func (s *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) GetValue() *string {
	return s.Value
}

func (s *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) SetKey(v string) *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags {
	s.Key = &v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) SetValue(v string) *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags {
	s.Value = &v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestSecurityOptions struct {
	// The confidential computing mode. Possible values:
	//
	// - Enclave: The ECS instance uses an Enclave to build a confidential computing environment. For more information, see [Build a confidential computing environment using an Enclave](https://help.aliyun.com/document_detail/203433.html).
	//
	// - TDX: Builds a TDX confidential computing environment. For more information, see [Build a TDX confidential computing environment](https://help.aliyun.com/document_detail/479090.html).
	//
	// example:
	//
	// TDX
	ConfidentialComputingMode *string `json:"ConfidentialComputingMode,omitempty" xml:"ConfidentialComputingMode,omitempty"`
}

func (s CreateScalingConfigurationRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *CreateScalingConfigurationRequestSecurityOptions) SetConfidentialComputingMode(v string) *CreateScalingConfigurationRequestSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *CreateScalingConfigurationRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestSpotPriceLimits struct {
	// The instance type of the preemptible instance. This parameter takes effect when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The bid price for the preemptible instance. This parameter takes effect when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.5
	PriceLimit *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s CreateScalingConfigurationRequestSpotPriceLimits) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) GetPriceLimit() *float32 {
	return s.PriceLimit
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) SetInstanceType(v string) *CreateScalingConfigurationRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) SetPriceLimit(v float32) *CreateScalingConfigurationRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

func (s *CreateScalingConfigurationRequestSpotPriceLimits) Validate() error {
	return dara.Validate(s)
}
