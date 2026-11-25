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
	SetSecondaryEniVSwitchConstraint(v string) *CreateScalingConfigurationRequest
	GetSecondaryEniVSwitchConstraint() *string
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
	// Specifies whether to associate the instance on a dedicated host with the dedicated host. Valid values:
	//
	// - default: The instance is not associated with the dedicated host. If an instance with the economical mode enabled is stopped and then started, and the resources of the original dedicated host are insufficient, the instance is automatically deployed to another dedicated host in the automatic deployment resource pool.
	//
	// - host: The instance is associated with the dedicated host. If an instance with the economical mode enabled is stopped and then started, it remains on the original dedicated host. If the resources of the original dedicated host are insufficient, the instance fails to start.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of vCPUs. Unit: cores.
	//
	// You can specify both Cpu and Memory to define the range of instance types. For example, if you set Cpu to 2 and Memory to 16, all instance types with 2 vCPUs and 16 GiB of memory are matched. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and creates instances with the lowest price.
	//
	// > This configuration takes effect only when the scaling mode is set to cost optimization and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// Specifies the running pattern of the burstable instance. Valid values:
	//
	// - Standard: standard pattern.
	//
	// - Unlimited: unlimited pattern.
	//
	// For more information, see the performance schema section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The custom priority of the ECS instance type and vSwitch combination.
	//
	// 	Notice: This parameter takes effect only when the scaling policy of the scaling group is set to the priority-based policy.
	//
	// If an instance cannot be created based on the instance type and vSwitch combination with a higher priority, Auto Scaling automatically selects the combination with the next highest priority to create the instance.
	//
	// > If you specify custom priorities for only some instance type and vSwitch combinations, the unspecified combinations have lower priorities than the specified ones. The priorities of the unspecified combinations are determined based on the vSwitch sequence in the scaling group and the instance type sequence in the scaling configuration.
	//
	// > - Example: If the vSwitch sequence in the scaling group is vsw1 and vsw2, the instance type sequence in the scaling configuration is type1 and type2, and the custom priority sequence is ["vsw2+type2", "vsw1+type2"], the final priority sequence is: "vsw2+type2" > "vsw1+type2" > "vsw1+type1" > "vsw2+type1".
	CustomPriorities []*CreateScalingConfigurationRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The collection of data cloud disk information.
	DataDisks []*CreateScalingConfigurationRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-2zedxc67zqzt7lb4****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// Specifies whether to create the ECS instance on a dedicated host. Because dedicated hosts do not support spot instances, the SpotStrategy and SpotPriceLimit parameters are automatically ignored if DedicatedHostId is specified.
	//
	// You can call the DescribeDedicatedHosts operation to query the list of dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection property of the instance. This property specifies whether the instance can be directly released by using the ECS console or the DeleteInstance API operation, which prevents accidental deletion of the instance. Valid values:
	//
	// - true: enables release protection for the instance. The instance cannot be directly released by using the ECS console or the DeleteInstance API operation.
	//
	// - false: shutdown release protection for the instance. The instance can be directly released by using the ECS console or the DeleteInstance API operation.
	//
	// Default value: false.
	//
	// > This property applies only to pay-as-you-go instances to prevent instances created by Auto Scaling during scale-out from being accidentally deleted. It does not affect normal scale-in activities. Instances with release protection enabled can still be released during scale-in activities.
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
	// The hostname of the ECS instance. The hostname cannot start or end with a period (.) or a hyphen (-) and cannot contain consecutive periods (.) or hyphens (-). The naming requirements for different instance types are as follows:
	//
	// - Windows instances: The hostname must be 2 to 15 characters in length and can contain letters, digits, and hyphens (-). It cannot contain periods (.) or contain only digits.
	//
	// - Other instance types such as Linux instances: The hostname must be 2 to 64 characters in length and can contain multiple periods (.). The string between two periods is a segment, which can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// host****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the HPC cluster to which the ECS instance belongs.
	//
	// example:
	//
	// hpc-clusterid
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enables the access channel.
	//
	// - disabled: disables the access channel.
	//
	// Default value: enabled.
	//
	// > For more information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// Specifies whether to forcibly use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// - optional: does not forcibly use the security-hardened mode.
	//
	// - required: forcibly uses the security-hardened mode. After you set this value, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// > For more information about the modes for accessing instance metadata, see [Access modes of instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image in the image family to create instances. If `ImageId` is specified, you cannot specify this parameter.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image file. The image resource used during automatic creation of instances.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image file. The image name must be unique within a region. If ImageId is specified, ImageName is ignored.
	//
	//
	// You cannot use ImageName to specify an Alibaba Cloud Marketplace image.
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
	// The name of the ECS instance during automatic creation by using this scaling configuration.
	//
	// example:
	//
	// instance****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The collection of intelligent configuration information used to filter instance types that meet your requirements.
	InstancePatternInfos []*CreateScalingConfigurationRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance type of the ECS instance. For more information, see instance families.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// After you enable the candidate mode, if issues such as insufficient stock occur, the system supplements the currently selected instance types with similar instance types of the same size, or creates vSwitches in candidate zones and adds them to the scaling group.
	InstanceTypeCandidateOptions *CreateScalingConfigurationRequestInstanceTypeCandidateOptions `json:"InstanceTypeCandidateOptions,omitempty" xml:"InstanceTypeCandidateOptions,omitempty" type:"Struct"`
	// The information used to specify instance types.
	InstanceTypeOverrides []*CreateScalingConfigurationRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	// The multiple instance types. If InstanceTypes is specified, InstanceType is ignored.
	//
	// If an instance cannot be created based on the instance type with a higher priority, Auto Scaling automatically selects the instance type with the next highest priority to create the instance.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The network billing method. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth. In this case, InternetMaxBandwidthOut specifies the fixed bandwidth value.
	//
	// - PayByTraffic: pay-by-traffic. In this case, InternetMaxBandwidthOut specifies only the maximum bandwidth, and you are billed based on the actual network traffic.
	//
	// If this parameter is not specified, the default value is PayByBandwidth for the classic network and PayByTraffic for a VPC.
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
	// Specifies whether the instance is an I/O optimization instance. Valid values:
	//
	// - none: The instance is not an I/O optimization instance.
	//
	// - optimized: The instance is an I/O optimization instance.
	//
	// The default value is none for retired instance types and optimized for other instance types.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of IPv6 addresses that are randomly generated for the network interface controller (NIC).
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair used to log on to the ECS instance.
	//
	//
	//
	// - For Windows instances, this parameter is ignored and left empty by default.
	//
	// - For Linux instances, password-based logon is disabled by default.
	//
	// example:
	//
	// KeyPairTest
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The weight of the ECS instance as a load balancing backend server. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify both Cpu and Memory to define the range of instance types. For example, if you set Cpu to 2 and Memory to 16, all instance types with 2 vCPUs and 16 GiB of memory are matched. Auto Scaling determines the available instance types based on factors such as I/O optimization and zone, and creates instances with the lowest price.
	//
	// > This configuration takes effect only when the scaling mode is set to cost optimization and no instance type is specified in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The list of network interface controller (NIC) devices.
	NetworkInterfaces []*CreateScalingConfigurationRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	OwnerAccount      *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the ECS instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	// \\`()` ~!@#$%^&*-_+=\\|{}[]:;\\"<>,.?/`
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify the Password parameter, send requests over HTTPS to prevent password leaks.
	//
	// example:
	//
	// 123abc****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image. When you use this parameter, make sure that a password is preset in the image. Valid values:
	//
	// - true: uses the preset password of the image.
	//
	// - false: does not use the preset password of the image.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
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
	// rg-resource****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource pool policy used when creating the instance. Take note of the following items when you configure this parameter:
	//
	// - This parameter takes effect only when you create pay-as-you-go instances.
	//
	// - You cannot configure this parameter together with PrivatePoolOptions.MatchCriteria or PrivatePoolOptions.Id.
	ResourcePoolOptions *CreateScalingConfigurationRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
	// The name of the scaling configuration. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter, digit, or Chinese character.
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
	// The scheduling options.
	//
	// example:
	//
	// ["testManagedPrivateSpaceId****"]
	SchedulerOptions map[string]interface{} `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// Specifies whether a vSwitch available for secondary elastic network interfaces (ENIs) must be selected from the list when the VSwitchIds parameter is specified for the secondary ENIs. Valid values:
	//
	// - Strict: strict mode. If no vSwitch in the VSwitchIds list meets the requirements, instance creation fails.
	//
	// - Tolerant: fault-tolerant mode. If no vSwitch in the VSwitchIds list meets the requirements, no secondary ENI is configured, and the instance is still created.
	//
	// Default value: Strict.
	//
	// example:
	//
	// Strict
	SecondaryEniVSwitchConstraint *string `json:"SecondaryEniVSwitchConstraint,omitempty" xml:"SecondaryEniVSwitchConstraint,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// - Active: enables security hardening. This value applies only to public images.
	//
	// - Deactive: disables security hardening. This value applies to all image types.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the ECS instance belongs. ECS instances in the same security group can access each other.
	//
	// example:
	//
	// sg-280ih****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Adds the ECS instance to multiple security groups at the same time. For more information, see the Security groups section in [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > You cannot specify both SecurityGroupId and SecurityGroupIds at the same time.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *CreateScalingConfigurationRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Valid values:
	//
	// - 1: Alibaba Cloud guarantees that the instance runs for 1 hour after it is created and is not automatically released. After 1 hour, the system automatically compares your bid with the market price and checks resource inventory to determine whether to retain automatic release the instance.
	//
	// - 0: Alibaba Cloud does not guarantee that the instance runs for 1 hour after it is created. The system automatically compares your bid with the market price and checks resource inventory to determine whether to retain automatic release the instance.
	//
	// > Alibaba Cloud sends you a notification by using ECS system events 5 minutes before the instance is released. Spot instances are billed by the second. Select an appropriate protection period based on the execution duration of your tasks.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The break mode of the spot instance. Currently, only terminate (default) is supported, which releases the instance directly.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The collection of billing information for spot instances.
	SpotPriceLimits []*CreateScalingConfigurationRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy for the pay-as-you-go instance. Valid values:
	//
	//
	//
	// - NoSpot: regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: spot instance with a maximum hourly price.
	//
	// - SpotAsPriceGo: spot instance whose price is automatically set based on the current market price.
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
	// The maximum number of partitions in the storage set. Valid values: integers greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The disk types for the system cloud disk. If a disk type with a higher priority is unavailable, Auto Scaling automatically attempts to use the disk type with the next highest priority to create the system cloud disk. Valid values:
	//
	// - cloud: basic cloud disk.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD cloud disk.
	//
	// - cloud_essd: ESSD cloud disk.
	//
	// > You cannot specify both this parameter and `SystemDisk.Category` at the same time.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The tags of the ECS instance. Tags are passed in as key-value pairs. You can add up to 20 tags. The requirements for keys and values are as follows:
	//
	//
	//
	// - A key can be up to 64 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`. A key cannot be an empty string.
	//
	// - A value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`. A value can be an empty string.
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
	// The user data of the ECS instance. It must be encoded in Base64. The maximum size of the raw data before encoding is 32 KB.
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

func (s *CreateScalingConfigurationRequest) GetSecondaryEniVSwitchConstraint() *string {
	return s.SecondaryEniVSwitchConstraint
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

func (s *CreateScalingConfigurationRequest) SetSecondaryEniVSwitchConstraint(v string) *CreateScalingConfigurationRequest {
	s.SecondaryEniVSwitchConstraint = &v
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
	// Specifies whether to log on to the ECS instance as the ecs-user user. For more information, see [Manage the logon username of an ECS instance](https://help.aliyun.com/document_detail/388447.html). Valid values:
	//
	// - true: yes.
	//
	// - false: no.
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
	// The ID of the private pool. It is the ID of the elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool options for instance startup. After an elasticity assurance or a capacity reservation takes effect, a private pool is generated for you to select when you start instances. Valid values:
	//
	// - Open: open mode. The system automatically matches an open private pool. If no matching private pool is available, the system uses resources from the public pool. In this mode, you do not need to specify the PrivatePoolOptions.Id parameter.
	//
	// - Target: specified mode. The system uses the specified private pool to start the instance. If the specified private pool is unavailable, the instance fails to start. In this mode, you must specify the private pool ID by using the PrivatePoolOptions.Id parameter.
	//
	// - None: no private pool. The system does not use private pools to start the instance.
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
	// The ID of the automatic snapshot policy applied to the system cloud disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burst (performance burst) feature for the system cloud disk. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// > This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
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
	// - cloud_ssd: standard SSD cloud disk.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_essd: ESSD.
	//
	// - cloud_auto: ESSD AutoPL.
	//
	// You cannot specify this parameter together with `SystemDiskCategories`. If neither this parameter nor `SystemDiskCategories` is specified, the default value is used:
	//
	// - For I/O optimized instances, the default value is cloud_efficiency.
	//
	// - For non-I/O optimized instances, the default value is cloud.
	//
	// This parameter specifies only one system cloud disk type. If you want the system to automatically try other types when the preferred disk type is unavailable, use `SystemDiskCategories` and configure multiple disk types in order of priority.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system cloud disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test system disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system cloud disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The encryption algorithm used for the system cloud disk. Valid values:
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
	// The settings of the performance level for the system cloud disk when it is an ESSD. Valid values:
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
	// The provisioned IOPS (Input/Output Operations Per Second) performance metrics for the system cloud disk.
	//
	// > IOPS refers to the number of I/O operations that can be processed per second. It indicates the read and write processing capability of block storage. Unit: operations per second.
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
	// 	Notice: The instance type must be included in the instance type list of the scaling configuration.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the vSwitch.
	//
	// 	Notice: The vSwitch must be included in the vSwitch list of the scaling group.
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
	// The ID of the automatic snapshot policy applied to the data cloud disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burstable performance feature for the data cloud disk. Valid values:
	//
	// - true: enables the feature.
	//
	// - false: disables the feature.
	//
	// > You can set this parameter only when `SystemDisk.Category` is set to `cloud_auto`.
	//
	// <props="china">For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The multiple disk types of the data cloud disk. If the disk type with a higher priority is unavailable, Auto Scaling automatically tries the disk type with the next highest priority to create the data cloud disk. Valid values:
	//
	// - cloud: basic cloud disk. The DeleteWithInstance property of a basic cloud disk created together with the instance is set to true.
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
	// > If you specify this parameter, you cannot specify `DataDisks.Category` at the same time.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data cloud disk. Valid values:
	//
	// - cloud: basic cloud disk. The DeleteWithInstance attribute of a basic cloud disk created together with the instance is set to true.
	//
	// - cloud_efficiency: ultra cloud disk.
	//
	// - cloud_ssd: standard SSD cloud disk.
	//
	// - cloud_essd: ESSD.
	//
	// - ephemeral_ssd: local SSD.
	//
	// - cloud_auto: ESSD AutoPL.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// You cannot specify this parameter together with DataDisk.Categories. If neither this parameter nor DataDisk.Categories is specified, the default value is used:
	//
	// - For I/O optimized instances, the default value is cloud_efficiency.
	//
	// - For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data cloud disk when the instance is released. Valid values:
	//
	// - true: releases the data cloud disk when the instance is released.
	//
	// - false: retains the data cloud disk when the instance is released.
	//
	// You can set this parameter only for independent cloud disks (when DataDisks.Category is set to cloud, cloud_efficiency, cloud_ssd, or cloud_essd). Otherwise, an error is returned.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data cloud disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test data disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data cloud disk. If you do not specify this parameter, the system automatically allocates a mount point by default during automatic creation of the ECS instance. The mount point starts from /dev/xvdb and ends at /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data cloud disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`.
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
	// The ID of the Key Management Service (KMS) key used by the data cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The settings of the performance level for the data cloud disk when it is an ESSD. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// > For more information about how to select the performance level of an ESSD, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned IOPS (Input/Output Operations Per Second) performance metrics of the data cloud disk.
	//
	// > IOPS indicates the number of I/O operations that can be processed per second. It measures the read and write capabilities of block storage. Unit: operations per second.
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
	// The snapshot used to create the data cloud disk. If this parameter is specified, DataDisks.Size is ignored, and the actual size of the created disk is the size of the specified snapshot.
	//
	// If the snapshot was created on or before July 15, 2013, the call is rejected and the InvalidSnapshot.TooOld error code is returned.
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
	// - X86: x86 computing.
	//
	// - Heterogeneous: heterogeneous computing, such as GPU or FPGA.
	//
	// - BareMental: ECS bare metal server.
	//
	// - Arm: Arm computing.
	//
	// Default value: all architecture types.
	Architectures []*string `json:"Architectures,omitempty" xml:"Architectures,omitempty" type:"Repeated"`
	// Specifies whether the instance type is a burstable instance type. Valid values:
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
	// The number of vCPUs of the instance type in intelligent configuration mode. This parameter is used to filter instance types that meet your requirements. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// Take note of the following items:
	//
	// - The InstancePatternInfos parameter is applicable only when the network type of the scaling group is VPC.
	//
	// - You must configure both InstancePatternInfos.Cores and InstancePatternInfos.Memory when you use the InstancePatternInfos parameter.
	//
	// - If you have already specified instance types by using the InstanceType or InstanceTypes parameter, Auto Scaling preferentially uses the specified instance types for scale-out. If the specified instance types are out of stock, Auto Scaling preferentially uses the instance type that has the lowest price among the instance types that meet the requirements specified by the InstancePatternInfos parameter.
	//
	// example:
	//
	// 2
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// The CPU architecture of the instance. Valid values:
	//
	// >N indicates that you can specify multiple CPU architectures. Valid values of N: 1 to 2.
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
	// >N indicates the Settings for multiple instance categories. Valid values of N: 1 to 10.
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
	// - ECS Bare Metal: ECS Bare Metal Instance.
	//
	// - High Performance Compute: high-performance computing (HPC).
	InstanceCategories []*string `json:"InstanceCategories,omitempty" xml:"InstanceCategories,omitempty" type:"Repeated"`
	// The level of the instance family. This parameter is used to filter instance types that meet your requirements and takes effect when `CostOptimization` is enabled. Valid values:
	//
	// - EntryLevel: shared instance types. These instance types are more cost-effective but cannot ensure consistent computing performance. They are suitable for scenarios where the CPU utilization is typically low. For more information, see [Shared instance families](https://help.aliyun.com/document_detail/108489.html).
	//
	// - EnterpriseLevel: enterprise-level instance types. These instance types provide consistent performance and dedicated resources. They are suitable for scenarios that require high stability. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - CreditEntryLevel: credit-based entry-level instance types, which are burstable instances. These instance types use CPU credits to ensure computing performance and are suitable for scenarios where the CPU utilization is typically low but occasionally bursts. For more information, see [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// EnterpriseLevel
	InstanceFamilyLevel *string `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	// The instance families to be queried. N indicates that you can specify multiple instance families. Valid values of N: 1 to 10.
	InstanceTypeFamilies []*string `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
	// The maximum acceptable hourly price for pay-as-you-go or spot instances in intelligent configuration mode. This parameter is used to filter instance types that meet your requirements.
	//
	// > This parameter is required when SpotStrategy is set to SpotWithPriceLimit. In other cases, this parameter is optional.
	//
	// example:
	//
	// 2
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The maximum number of vCPUs of the instance type.
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
	// The memory size of the instance type in intelligent configuration mode. Unit: GiB. This parameter is used to filter instance types that meet your requirements.
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
	// The minimum number of vCPUs of the instance type.
	//
	// example:
	//
	// 2
	MinimumCpuCoreCount *int32 `json:"MinimumCpuCoreCount,omitempty" xml:"MinimumCpuCoreCount,omitempty"`
	// The minimum upper limit on the number of IPv6 addresses per elastic network interfaces (ENIs) or network interface controller (NIC) for the instance.
	//
	// example:
	//
	// 1
	MinimumEniIpv6AddressQuantity *int32 `json:"MinimumEniIpv6AddressQuantity,omitempty" xml:"MinimumEniIpv6AddressQuantity,omitempty"`
	// The minimum upper limit on the number of IPv4 addresses per elastic network interfaces (ENIs) or network interface controller (NIC) for the instance.
	//
	// example:
	//
	// 2
	MinimumEniPrivateIpAddressQuantity *int32 `json:"MinimumEniPrivateIpAddressQuantity,omitempty" xml:"MinimumEniPrivateIpAddressQuantity,omitempty"`
	// The minimum upper limit on the number of network interface controller (NIC) devices that can be mounted to the instance.
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
	// The physical processor model of the instance. N indicates that you can specify multiple processor models. Valid values of N: 1 to 10.
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
	// Specifies the CIDR blocks of the optional vSwitches when you add vSwitches in other zones.
	AllowCidrBlocks []*string `json:"AllowCidrBlocks,omitempty" xml:"AllowCidrBlocks,omitempty" type:"Repeated"`
	// Specifies whether to allow supplementing vSwitches in other zones.
	//
	// > The instance types remain unchanged, and only new zones are used as candidates. If the scaling group cannot scale out in all selected zones due to issues such as insufficient stock, the system automatically adds vSwitches in new zones to the scaling group based on this configuration.
	//
	// > For example, if the scaling group is configured with the cn-hangzhou-h and cn-hangzhou-g zones, and scale-out fails in both zones, Auto Scaling may create a vSwitch in the cn-hangzhou-k zone and add it to the scaling group based on real-time stock availability.
	//
	// example:
	//
	// true
	AllowCrossAz *bool `json:"AllowCrossAz,omitempty" xml:"AllowCrossAz,omitempty"`
	// Specifies whether to allow supplementing instance types of other generations.
	//
	// - For example, if the current instance type is ecs.c7.large, instance types such as ecs.c6.large and ecs.c8.large can be used as candidates after this feature is enabled.
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
	// When you need the scaling group to scale based on the capacity of the instance type, specify both this parameter and WeightedCapacity.
	//
	// This parameter specifies the instance type and overrides the instance type in the launch template. You can specify N instance types to extend the launch template to support N instance types. Valid values of N: 1 to 20.
	//
	// > This parameter takes effect only when LaunchTemplateId is specified.
	//
	// Valid values of InstanceType: available ECS instance types.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// When you need to specify the capacity of the instance type in the scaling configuration, specify this parameter after you specify InstanceTypeOverrides.InstanceType.
	//
	// This parameter specifies the weight of the instance type, which indicates the capacity that a single instance of the instance type provides in the scaling group. A greater weight indicates that fewer instances of this instance type are required to meet the expected capacity.
	//
	// Because instance types have different performance metrics such as the number of vCPUs and memory size, you can configure different weights for different instance types based on your business requirements.
	//
	// Example:
	//
	// - Current capacity: 0.
	//
	// - Expected capacity: 6.
	//
	// - Capacity of the ecs.c5.xlarge instance type: 4.
	//
	// To meet the expected capacity, the scaling group scales out by adding two ecs.c5.xlarge instances.
	//
	// > During scale-out, the capacity of the scaling group cannot exceed the sum of the maximum capacity (MaxSize) and the maximum weight of the instance type.
	//
	// Valid values of WeightedCapacity: 1 to 500.
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
	// The elastic network interfaces (ENIs) type. When you use this parameter, you must use the NetworkInterfaces Settings for the primary network interface controller (NIC), and you cannot set the SecurityGroupId or SecurityGroupIds parameter. Valid values:
	//
	// - Primary: primary elastic network interface (ENI).
	//
	// - Secondary: secondary elastic network interface (ENI).
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The number of randomly generated IPv6 addresses for the primary ENI.
	//
	// Take note of the following items:
	//
	// - This parameter takes effect only when NetworkInterface.InstanceType is set to Primary. If NetworkInterface.InstanceType is set to Secondary or left empty, you cannot configure this parameter.
	//
	// - After you configure this parameter, you cannot configure Ipv6AddressCount.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The communication pattern of the network interface controller (NIC). Valid values of the parameter:
	//
	// - Standard: uses the TCP communication pattern.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication pattern.
	//
	// Default value: Standard.
	//
	// > The number of elastic network interfaces (ENIs) in RDMA pattern cannot exceed the limit of the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The number of secondary private IPv4 addresses to be assigned to the ENI. Valid values: 1 to 49.
	//
	// - The value cannot exceed the limit on the number of IP addresses for the instance type. For more information, see [Instance families](https://www.alibabacloud.com/help/en/ecs/user-guide/overview-of-instance-families).
	//
	// - NetworkInterface.N.SecondaryPrivateIpAddressCount specifies the number of secondary private IPv4 addresses to be assigned to the ENI, excluding the primary private IP address of the ENI. The system randomly assigns these addresses from the available CIDR block of the vSwitch (NetworkInterface.N.VSwitchId) to which the ENI belongs.
	//
	// example:
	//
	// 6
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The IDs of one or more security groups to which the elastic network interfaces (ENIs) belong.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the elastic network interfaces (ENIs) belong. N indicates the Settings for multiple vSwitches. Valid values of N: 1 to 10.
	//
	// Take note of the following items:
	//
	// - If InstanceType is set to Primary, you cannot set this parameter. The default value is the vSwitch to which the instance belongs.
	//
	// - If InstanceType is set to Secondary or left empty, this parameter is optional.
	//
	// - The system selects a vSwitch that is in the same zone as the primary network interface controller (NIC) and has the most active IP addresses.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
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

func (s *CreateScalingConfigurationRequestNetworkInterfaces) GetVSwitchIds() []*string {
	return s.VSwitchIds
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

func (s *CreateScalingConfigurationRequestNetworkInterfaces) SetVSwitchIds(v []*string) *CreateScalingConfigurationRequestNetworkInterfaces {
	s.VSwitchIds = v
	return s
}

func (s *CreateScalingConfigurationRequestNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestResourcePoolOptions struct {
	// The ID of the private pool, which is the ID of the elasticity assurance or capacity reservation service. You can only specify the ID of a private pool in Target mode, and you cannot specify this parameter together with PrivatePoolTags.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// Filters available Target private pools by tag. Valid values of N: 1 to 20.
	//
	// Note:
	//
	// - When you configure this parameter, the system filters private pools from the Target private pools associated with your account. The filtered private pools must match the tags and meet the scaling group constraints, such as zone and instance type.
	//
	// - Tag matching rule: The private pool must match all specified tags.
	//
	// - You cannot specify this parameter together with PrivatePoolIds.
	PrivatePoolTags []*CreateScalingConfigurationRequestResourcePoolOptionsPrivatePoolTags `json:"PrivatePoolTags,omitempty" xml:"PrivatePoolTags,omitempty" type:"Repeated"`
	// Resource pools include private pools generated after the elasticity assurance or capacity reservation service takes effect, as well as public pools. They are available for selection when the instance starts. Valid values:
	//
	// - PrivatePoolFirst: private pools take precedence. When you select this policy, if ResourcePoolOptions.PrivatePoolIds is specified or the PrivatePoolTags condition is met, the corresponding private pool is preferentially used. If no private pool is specified or the specified private pool has insufficient capacity, the system automatically matches an open private pool. If no eligible private pool is available, the public pool is used to create the instance.
	//
	// - PrivatePoolOnly: only private pools are used. When you select this policy, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pool has insufficient capacity, the instance fails to start.
	//
	// - PublicPoolFirst: public pool resources take precedence. The public pool is preferentially used to create the instance. When public pool resources are insufficient, private pool resources are used as a supplement. The system preferentially and automatically matches an open private pool. If no eligible private pool is available, the specified ResourcePoolOptions.PrivatePoolIds or a Target private pool that meets the PrivatePoolTags condition is used. (This policy is in invitational preview and is not supported yet.)
	//
	// - None: no resource pool policy is used.
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
	// The confidential computing mode. Valid values:
	//
	// -  Enclave: The ECS instance uses Enclave to build a confidential computing environment. For more information, see [Build a confidential computing environment by using Enclave](https://help.aliyun.com/document_detail/203433.html).
	//
	// - TDX: builds a TDX confidential computing environment. For more information, see [Build a TDX confidential computing environment](https://help.aliyun.com/document_detail/479090.html).
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
	// The instance type of the spot instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum hourly price for the spot instance. This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit.
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
