// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageOptions(v *CreateScalingConfigurationShrinkRequestImageOptions) *CreateScalingConfigurationShrinkRequest
	GetImageOptions() *CreateScalingConfigurationShrinkRequestImageOptions
	SetPrivatePoolOptions(v *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) *CreateScalingConfigurationShrinkRequest
	GetPrivatePoolOptions() *CreateScalingConfigurationShrinkRequestPrivatePoolOptions
	SetSystemDisk(v *CreateScalingConfigurationShrinkRequestSystemDisk) *CreateScalingConfigurationShrinkRequest
	GetSystemDisk() *CreateScalingConfigurationShrinkRequestSystemDisk
	SetAffinity(v string) *CreateScalingConfigurationShrinkRequest
	GetAffinity() *string
	SetClientToken(v string) *CreateScalingConfigurationShrinkRequest
	GetClientToken() *string
	SetCpu(v int32) *CreateScalingConfigurationShrinkRequest
	GetCpu() *int32
	SetCreditSpecification(v string) *CreateScalingConfigurationShrinkRequest
	GetCreditSpecification() *string
	SetCustomPriorities(v []*CreateScalingConfigurationShrinkRequestCustomPriorities) *CreateScalingConfigurationShrinkRequest
	GetCustomPriorities() []*CreateScalingConfigurationShrinkRequestCustomPriorities
	SetDataDisks(v []*CreateScalingConfigurationShrinkRequestDataDisks) *CreateScalingConfigurationShrinkRequest
	GetDataDisks() []*CreateScalingConfigurationShrinkRequestDataDisks
	SetDedicatedHostClusterId(v string) *CreateScalingConfigurationShrinkRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostId(v string) *CreateScalingConfigurationShrinkRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *CreateScalingConfigurationShrinkRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *CreateScalingConfigurationShrinkRequest
	GetDeploymentSetId() *string
	SetHostName(v string) *CreateScalingConfigurationShrinkRequest
	GetHostName() *string
	SetHpcClusterId(v string) *CreateScalingConfigurationShrinkRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *CreateScalingConfigurationShrinkRequest
	GetHttpEndpoint() *string
	SetHttpTokens(v string) *CreateScalingConfigurationShrinkRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *CreateScalingConfigurationShrinkRequest
	GetImageFamily() *string
	SetImageId(v string) *CreateScalingConfigurationShrinkRequest
	GetImageId() *string
	SetImageName(v string) *CreateScalingConfigurationShrinkRequest
	GetImageName() *string
	SetInstanceDescription(v string) *CreateScalingConfigurationShrinkRequest
	GetInstanceDescription() *string
	SetInstanceName(v string) *CreateScalingConfigurationShrinkRequest
	GetInstanceName() *string
	SetInstancePatternInfos(v []*CreateScalingConfigurationShrinkRequestInstancePatternInfos) *CreateScalingConfigurationShrinkRequest
	GetInstancePatternInfos() []*CreateScalingConfigurationShrinkRequestInstancePatternInfos
	SetInstanceType(v string) *CreateScalingConfigurationShrinkRequest
	GetInstanceType() *string
	SetInstanceTypeOverrides(v []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) *CreateScalingConfigurationShrinkRequest
	GetInstanceTypeOverrides() []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides
	SetInstanceTypes(v []*string) *CreateScalingConfigurationShrinkRequest
	GetInstanceTypes() []*string
	SetInternetChargeType(v string) *CreateScalingConfigurationShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationShrinkRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationShrinkRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateScalingConfigurationShrinkRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *CreateScalingConfigurationShrinkRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *CreateScalingConfigurationShrinkRequest
	GetKeyPairName() *string
	SetLoadBalancerWeight(v int32) *CreateScalingConfigurationShrinkRequest
	GetLoadBalancerWeight() *int32
	SetMemory(v int32) *CreateScalingConfigurationShrinkRequest
	GetMemory() *int32
	SetNetworkInterfaces(v []*CreateScalingConfigurationShrinkRequestNetworkInterfaces) *CreateScalingConfigurationShrinkRequest
	GetNetworkInterfaces() []*CreateScalingConfigurationShrinkRequestNetworkInterfaces
	SetOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateScalingConfigurationShrinkRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateScalingConfigurationShrinkRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *CreateScalingConfigurationShrinkRequest
	GetPasswordInherit() *bool
	SetRamRoleName(v string) *CreateScalingConfigurationShrinkRequest
	GetRamRoleName() *string
	SetResourceGroupId(v string) *CreateScalingConfigurationShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourcePoolOptions(v *CreateScalingConfigurationShrinkRequestResourcePoolOptions) *CreateScalingConfigurationShrinkRequest
	GetResourcePoolOptions() *CreateScalingConfigurationShrinkRequestResourcePoolOptions
	SetScalingConfigurationName(v string) *CreateScalingConfigurationShrinkRequest
	GetScalingConfigurationName() *string
	SetScalingGroupId(v string) *CreateScalingConfigurationShrinkRequest
	GetScalingGroupId() *string
	SetSchedulerOptionsShrink(v string) *CreateScalingConfigurationShrinkRequest
	GetSchedulerOptionsShrink() *string
	SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationShrinkRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateScalingConfigurationShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateScalingConfigurationShrinkRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *CreateScalingConfigurationShrinkRequestSecurityOptions) *CreateScalingConfigurationShrinkRequest
	GetSecurityOptions() *CreateScalingConfigurationShrinkRequestSecurityOptions
	SetSpotDuration(v int32) *CreateScalingConfigurationShrinkRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *CreateScalingConfigurationShrinkRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimits(v []*CreateScalingConfigurationShrinkRequestSpotPriceLimits) *CreateScalingConfigurationShrinkRequest
	GetSpotPriceLimits() []*CreateScalingConfigurationShrinkRequestSpotPriceLimits
	SetSpotStrategy(v string) *CreateScalingConfigurationShrinkRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *CreateScalingConfigurationShrinkRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateScalingConfigurationShrinkRequest
	GetStorageSetPartitionNumber() *int32
	SetSystemDiskCategories(v []*string) *CreateScalingConfigurationShrinkRequest
	GetSystemDiskCategories() []*string
	SetTags(v string) *CreateScalingConfigurationShrinkRequest
	GetTags() *string
	SetTenancy(v string) *CreateScalingConfigurationShrinkRequest
	GetTenancy() *string
	SetUserData(v string) *CreateScalingConfigurationShrinkRequest
	GetUserData() *string
	SetZoneId(v string) *CreateScalingConfigurationShrinkRequest
	GetZoneId() *string
}

type CreateScalingConfigurationShrinkRequest struct {
	ImageOptions       *CreateScalingConfigurationShrinkRequestImageOptions       `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *CreateScalingConfigurationShrinkRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *CreateScalingConfigurationShrinkRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether to associate an ECS instance on a dedicated host with the dedicated host. Valid values:
	//
	// 	- default: does not associate the ECS instance with the dedicated host. If you start an ECS instance that was stopped in economical mode and the original dedicated host has insufficient resources, the ECS instance is automatically deployed to another dedicated host in the automatic deployment resource pool.
	//
	// 	- host: associates the ECS instance with the dedicated host. If you start an ECS instance that was stopped in economical mode, the instance remains on the original dedicated host. If the original dedicated host has insufficient resources, the ECS instance fails to start.
	//
	// Default value: default
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see the "[How to ensure the idempotence of a request](https://help.aliyun.com/document_detail/25693.html)" topic.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of vCPUs.
	//
	// You can specify the number of vCPUs and the memory size to determine the range of instance types. For example, you can set CPU to 2 and Memory to 16 to specify instance types that have 2 vCPUs and 16 GiB of memory. If you specify Cpu and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones. Then, Auto Scaling preferentially creates instances by using the lowest-priced instance type.
	//
	// > You can specify Cpu and Memory to determine the range of instance types only if you set Scaling Policy to Cost Optimization Policy and you do not specify instance types in the scaling configuration.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard: standard mode
	//
	// 	- Unlimited: unlimited mode
	//
	// For more information, see the "Performance modes" section in the "[Overview](https://help.aliyun.com/document_detail/59977.html)" topic.
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
	// >  If you specify the priorities of only partial custom ECS instance type + vSwitch combinations, Auto Scaling preferentially creates ECS instances by using the custom combinations that have specified priorities. If the custom combinations that have specified priorities do not provide sufficient resources, Auto Scaling creates ECS instances by using the custom combinations that do not have specified priorities based on the specified orders of vSwitches and instance types.
	//
	// 	- Example: the specified order of vSwitches for your scaling group is vsw1 and vsw2 and the specified order of instance types in your scaling configuration is type1 and type 2. In addition, you use CustomPriorities to specify ["vsw2+type2", "vsw1+type2"]. In this example, the vsw2+type2 combination has the highest priority and the vsw2+type1 combination has the lowest priority. The vsw1+type2 combination has a higher priority than the vsw1+type1 combination.
	CustomPriorities []*CreateScalingConfigurationShrinkRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The data disks.
	DataDisks []*CreateScalingConfigurationShrinkRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-2zedxc67zqzt7lb4****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The ID of the dedicated host on which you want to create an ECS instance. You cannot create preemptible instances on dedicated hosts. If you specify DedicatedHostId, SpotStrategy and SpotPriceLimit are ignored.
	//
	// You can call the DescribeDedicatedHosts operation to query dedicated host IDs.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set of the ECS instances that are created by using the scaling configuration.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4pz****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The hostname of the ECS instance. The hostname cannot start or end with a period (.) or a hyphen (-). The hostname cannot contain consecutive periods (.) or hyphens (-). Naming conventions for different types of instances:
	//
	// 	- Windows instances: The hostname must be 2 to 15 characters in length and can contain letters, digits, and hyphens (-). The hostname cannot contain periods (.) or contain only digits.
	//
	// 	- Other instances, such as Linux instances: The hostname must be 2 to 64 characters in length. You can use periods (.) to separate a hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// host****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the Elastic High Performance Computing (E-HPC) cluster to which the ECS instances that are created by using the scaling configuration belong.
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
	// The name of the image family. If you specify this parameter, the most recent custom images that are available in the specified image family are returned. You can use the images to create instances. If you specify ImageId, you cannot specify ImageFamily.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image that Auto Scaling uses to automatically create ECS instances.
	//
	// example:
	//
	// centos6u5_64_20G_aliaegis****.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image. Each image name must be unique in a region. If you specify ImageId, ImageName is ignored.
	//
	// You cannot use ImageName to specify images that are purchased from Alibaba Cloud Marketplace.
	//
	// example:
	//
	// image****
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The description of the ECS instance. The description must be 2 to 256 characters in length. The description can contain letters and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test instance.
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The name of the ECS instance that Auto Scaling creates based on the scaling configuration.
	//
	// example:
	//
	// instance****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The intelligent configuration settings, which determine the available instance types.
	InstancePatternInfos []*CreateScalingConfigurationShrinkRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance type of the ECS instance. For more information, see the [Instance families](https://help.aliyun.com/document_detail/25378.html) topic.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The information about instance types.
	InstanceTypeOverrides []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
	// The instance types. If you specify InstanceTypes, InstanceType is ignored.
	//
	// Auto Scaling creates instances based on a priority list of instance types. If it fails to create instances using the highest-priority type, it automatically moves to the next type in the priority order.
	InstanceTypes []*string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	// The metering method for network usage. Valid values:
	//
	// 	- PayByBandwidth: You are charged for the maximum available bandwidth that is specified by InternetMaxBandwidthOut.
	//
	// 	- PayByTraffic: You are charged based on the amount of transferred data. InternetMaxBandwidthOut specifies only the maximum available bandwidth.
	//
	// For the classic network, the default value is PayByBandwidth. For VPCs, the default value is PayByTraffic.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// 	- If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of this parameter are 1 to 10, and the default value is 10.
	//
	// 	- If the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the value of `InternetMaxBandwidthOut`, and the default value is the value of `InternetMaxBandwidthOut`.
	//
	// example:
	//
	// 100
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 50
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether to create an I/O optimized instance. Valid values:
	//
	// none: does not create an I/O optimized instance. optimized: creates an I/O optimized instance.
	//
	// For instances of retired instance types, the default value is none. For instances of other instance types, the default value is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses that you want to allocate to the elastic network interface (ENI).
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair that you want to use to log on to an ECS instance.
	//
	// 	- Windows instances do not support this parameter.
	//
	// 	- By default, the username and password authentication method is disabled for Linux instances.
	//
	// example:
	//
	// KeyPairTest
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The weight of an ECS instance as a backend server. Valid values: 1 to 100.
	//
	// Default value: 50
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The memory size. Unit: GiB.
	//
	// You can specify the number of vCPUs and the memory size to determine the range of instance types. For example, you can set Cpu to 2 and Memory to 16 to specify instance types that have 2 vCPUs and 16 GiB of memory. If you specify Cpu and Memory, Auto Scaling determines the available instance types based on factors such as I/O optimization requirements and zones. Then, Auto Scaling preferentially creates instances by using the lowest-priced instance type.
	//
	// > You can specify Cpu and Memory to determine the range of instance types only if you set Scaling Policy to Cost Optimization Policy and you do not specify instance types in the scaling configuration.
	//
	// example:
	//
	// 16
	Memory            *int32                                                      `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkInterfaces []*CreateScalingConfigurationShrinkRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	OwnerAccount      *string                                                     `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password that you want to use to log on to an ECS instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// `` `() ~!@#$%^&*-_+=\\|{}[]:;\\"<>,.?/ ``
	//
	// The password of a Windows instance cannot start with a forward slash (/).
	//
	// > For security reasons, we recommend that you use HTTPS to send requests if you specify Password.
	//
	// example:
	//
	// 123abc****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password that is preconfigured in the image. Before you use this parameter, make sure that a password is configured in the image. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The name of the RAM role that you attach to the ECS instance. The name is provided and maintained by Resource Access Management (RAM). You can call the ListRoles operation to query the available RAM roles.
	//
	// example:
	//
	// ramrole****
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the resource group to which the ECS instances that are created by using the scaling configuration belong.
	//
	// example:
	//
	// rg-resource****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource pools used for instance creation, which can be the public pool or a private pool associated with any active elasticity assurance or capacity reservation. When you specify this parameter, take note of the following items:
	//
	// 	- This parameter takes effect only when you create pay-as-you-go instances.
	//
	// 	- If you specify this parameter, you cannot specify PrivatePoolOptions.MatchCriteria or PrivatePoolOptions.Id.
	ResourcePoolOptions *CreateScalingConfigurationShrinkRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
	// The name of the scaling configuration. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). The name must start with a letter or a digit.
	//
	// The name of the scaling configuration must be unique in a region. If you do not specify this parameter, the scaling configuration ID is used.
	//
	// example:
	//
	// scalingconfig****
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// The ID of the scaling group in which you want to create a scaling configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The scheduler options.
	//
	// example:
	//
	// ["testManagedPrivateSpaceId****"]
	SchedulerOptionsShrink *string `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// 	- Active: enables security hardening. This value is applicable only to public images.
	//
	// 	- Deactive: disables security hardening. This value is applicable to all image types.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group with which ECS instances are associated. ECS instances that are associated with the same security group can access each other.
	//
	// example:
	//
	// sg-280ih****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups with which you want to associate the ECS instances that are created by using the scaling configuration. For more information, see the "Security group limits" section of the "[Limits](https://help.aliyun.com/document_detail/25412.html)" topic.
	//
	// > If you specify SecurityGroupId, you cannot specify SecurityGroupIds.
	SecurityGroupIds []*string                                               `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *CreateScalingConfigurationShrinkRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The retention period of the preemptible instance. Unit: hours. Valid values: 0, 1, 2, 3, 4, 5, and 6.
	//
	// 	- The following retention periods are available in invitational preview: 2, 3, 4, 5, and 6 hours. If you want to set this parameter to one of these values, submit a ticket.
	//
	// 	- If you set this parameter to 0, no protection period is specified for the preemptible instance.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the preemptible instance. Set the value to Terminate. The value specifies that the preemptible instance is to be released.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The billing information of the spot instances.
	SpotPriceLimits []*CreateScalingConfigurationShrinkRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
	// The preemption policy that you want to apply to pay-as-you-go and preemptible instances. Valid values:
	//
	// 	- NoSpot: The instance is created as a pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a preemptible instance that has a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a preemptible instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot
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
	// The categories of the system disks. If Auto Scaling cannot create instances by using the disk category that has the highest priority, Auto Scaling creates instances by using the disk category that has the next highest priority. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	//
	// > If you specify SystemDiskCategories, you cannot specify `SystemDisk.Category`.
	SystemDiskCategories []*string `json:"SystemDiskCategories,omitempty" xml:"SystemDiskCategories,omitempty" type:"Repeated"`
	// The tags of the ECS instance. Tags must be specified as key-value pairs. You can specify up to 20 tags. When you specify tag keys and tag values, take note of the following items:
	//
	// 	- A tag key can be up to 64 characters in length. The key cannot start with acs: or aliyun and cannot contain `http://` or `https://`. You cannot specify an empty string as a tag key.
	//
	// 	- A tag value can be up to 128 characters in length. The value cannot start with acs: or aliyun and cannot contain `http://` or `https://`. You can specify an empty string as a tag value.
	//
	// example:
	//
	// {"key1":"value1","key2":"value2", ... "key5":"value5"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Specifies whether to create an ECS instance on a dedicated host. Valid values:
	//
	// 	- default: does not create an ECS instance on a dedicated host.
	//
	// 	- host: creates an ECS instance on a dedicated host. If you do not specify DedicatedHostId, Alibaba Cloud selects a dedicated host for the ECS instance.
	//
	// Default value: default
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
	// The zone ID of the ECS instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequest) GetImageOptions() *CreateScalingConfigurationShrinkRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateScalingConfigurationShrinkRequest) GetPrivatePoolOptions() *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateScalingConfigurationShrinkRequest) GetSystemDisk() *CreateScalingConfigurationShrinkRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateScalingConfigurationShrinkRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *CreateScalingConfigurationShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateScalingConfigurationShrinkRequest) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateScalingConfigurationShrinkRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateScalingConfigurationShrinkRequest) GetCustomPriorities() []*CreateScalingConfigurationShrinkRequestCustomPriorities {
	return s.CustomPriorities
}

func (s *CreateScalingConfigurationShrinkRequest) GetDataDisks() []*CreateScalingConfigurationShrinkRequestDataDisks {
	return s.DataDisks
}

func (s *CreateScalingConfigurationShrinkRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *CreateScalingConfigurationShrinkRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateScalingConfigurationShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateScalingConfigurationShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateScalingConfigurationShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateScalingConfigurationShrinkRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *CreateScalingConfigurationShrinkRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateScalingConfigurationShrinkRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateScalingConfigurationShrinkRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateScalingConfigurationShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateScalingConfigurationShrinkRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateScalingConfigurationShrinkRequest) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *CreateScalingConfigurationShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateScalingConfigurationShrinkRequest) GetInstancePatternInfos() []*CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	return s.InstancePatternInfos
}

func (s *CreateScalingConfigurationShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationShrinkRequest) GetInstanceTypeOverrides() []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides {
	return s.InstanceTypeOverrides
}

func (s *CreateScalingConfigurationShrinkRequest) GetInstanceTypes() []*string {
	return s.InstanceTypes
}

func (s *CreateScalingConfigurationShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateScalingConfigurationShrinkRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateScalingConfigurationShrinkRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateScalingConfigurationShrinkRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateScalingConfigurationShrinkRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateScalingConfigurationShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateScalingConfigurationShrinkRequest) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *CreateScalingConfigurationShrinkRequest) GetMemory() *int32 {
	return s.Memory
}

func (s *CreateScalingConfigurationShrinkRequest) GetNetworkInterfaces() []*CreateScalingConfigurationShrinkRequestNetworkInterfaces {
	return s.NetworkInterfaces
}

func (s *CreateScalingConfigurationShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateScalingConfigurationShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateScalingConfigurationShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateScalingConfigurationShrinkRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateScalingConfigurationShrinkRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateScalingConfigurationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateScalingConfigurationShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateScalingConfigurationShrinkRequest) GetResourcePoolOptions() *CreateScalingConfigurationShrinkRequestResourcePoolOptions {
	return s.ResourcePoolOptions
}

func (s *CreateScalingConfigurationShrinkRequest) GetScalingConfigurationName() *string {
	return s.ScalingConfigurationName
}

func (s *CreateScalingConfigurationShrinkRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateScalingConfigurationShrinkRequest) GetSchedulerOptionsShrink() *string {
	return s.SchedulerOptionsShrink
}

func (s *CreateScalingConfigurationShrinkRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateScalingConfigurationShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateScalingConfigurationShrinkRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateScalingConfigurationShrinkRequest) GetSecurityOptions() *CreateScalingConfigurationShrinkRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateScalingConfigurationShrinkRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateScalingConfigurationShrinkRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateScalingConfigurationShrinkRequest) GetSpotPriceLimits() []*CreateScalingConfigurationShrinkRequestSpotPriceLimits {
	return s.SpotPriceLimits
}

func (s *CreateScalingConfigurationShrinkRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateScalingConfigurationShrinkRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateScalingConfigurationShrinkRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateScalingConfigurationShrinkRequest) GetSystemDiskCategories() []*string {
	return s.SystemDiskCategories
}

func (s *CreateScalingConfigurationShrinkRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateScalingConfigurationShrinkRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *CreateScalingConfigurationShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateScalingConfigurationShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageOptions(v *CreateScalingConfigurationShrinkRequestImageOptions) *CreateScalingConfigurationShrinkRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPrivatePoolOptions(v *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) *CreateScalingConfigurationShrinkRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSystemDisk(v *CreateScalingConfigurationShrinkRequestSystemDisk) *CreateScalingConfigurationShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetAffinity(v string) *CreateScalingConfigurationShrinkRequest {
	s.Affinity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetClientToken(v string) *CreateScalingConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCpu(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Cpu = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCreditSpecification(v string) *CreateScalingConfigurationShrinkRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetCustomPriorities(v []*CreateScalingConfigurationShrinkRequestCustomPriorities) *CreateScalingConfigurationShrinkRequest {
	s.CustomPriorities = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDataDisks(v []*CreateScalingConfigurationShrinkRequestDataDisks) *CreateScalingConfigurationShrinkRequest {
	s.DataDisks = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDedicatedHostClusterId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDedicatedHostId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDeletionProtection(v bool) *CreateScalingConfigurationShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetDeploymentSetId(v string) *CreateScalingConfigurationShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHostName(v string) *CreateScalingConfigurationShrinkRequest {
	s.HostName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHpcClusterId(v string) *CreateScalingConfigurationShrinkRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHttpEndpoint(v string) *CreateScalingConfigurationShrinkRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetHttpTokens(v string) *CreateScalingConfigurationShrinkRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageFamily(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetImageName(v string) *CreateScalingConfigurationShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceDescription(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceDescription = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceName(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstancePatternInfos(v []*CreateScalingConfigurationShrinkRequestInstancePatternInfos) *CreateScalingConfigurationShrinkRequest {
	s.InstancePatternInfos = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceTypeOverrides(v []*CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) *CreateScalingConfigurationShrinkRequest {
	s.InstanceTypeOverrides = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInstanceTypes(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.InstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInternetChargeType(v string) *CreateScalingConfigurationShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetIoOptimized(v string) *CreateScalingConfigurationShrinkRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetIpv6AddressCount(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetKeyPairName(v string) *CreateScalingConfigurationShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetLoadBalancerWeight(v int32) *CreateScalingConfigurationShrinkRequest {
	s.LoadBalancerWeight = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetMemory(v int32) *CreateScalingConfigurationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetNetworkInterfaces(v []*CreateScalingConfigurationShrinkRequestNetworkInterfaces) *CreateScalingConfigurationShrinkRequest {
	s.NetworkInterfaces = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetOwnerId(v int64) *CreateScalingConfigurationShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPassword(v string) *CreateScalingConfigurationShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetPasswordInherit(v bool) *CreateScalingConfigurationShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetRamRoleName(v string) *CreateScalingConfigurationShrinkRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourceGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetResourcePoolOptions(v *CreateScalingConfigurationShrinkRequestResourcePoolOptions) *CreateScalingConfigurationShrinkRequest {
	s.ResourcePoolOptions = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationShrinkRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetScalingGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSchedulerOptionsShrink(v string) *CreateScalingConfigurationShrinkRequest {
	s.SchedulerOptionsShrink = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityEnhancementStrategy(v string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSecurityOptions(v *CreateScalingConfigurationShrinkRequestSecurityOptions) *CreateScalingConfigurationShrinkRequest {
	s.SecurityOptions = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotDuration(v int32) *CreateScalingConfigurationShrinkRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotInterruptionBehavior(v string) *CreateScalingConfigurationShrinkRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotPriceLimits(v []*CreateScalingConfigurationShrinkRequestSpotPriceLimits) *CreateScalingConfigurationShrinkRequest {
	s.SpotPriceLimits = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSpotStrategy(v string) *CreateScalingConfigurationShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetStorageSetId(v string) *CreateScalingConfigurationShrinkRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetStorageSetPartitionNumber(v int32) *CreateScalingConfigurationShrinkRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetSystemDiskCategories(v []*string) *CreateScalingConfigurationShrinkRequest {
	s.SystemDiskCategories = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetTags(v string) *CreateScalingConfigurationShrinkRequest {
	s.Tags = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetTenancy(v string) *CreateScalingConfigurationShrinkRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetUserData(v string) *CreateScalingConfigurationShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) SetZoneId(v string) *CreateScalingConfigurationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequest) Validate() error {
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

type CreateScalingConfigurationShrinkRequestImageOptions struct {
	// For more information about whether an ECS instance uses the ecs-user user user to log on to an ECS instance, see [Manage the login name of an ECS instance](https://help.aliyun.com/document_detail/388447.html). Value range:
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

func (s CreateScalingConfigurationShrinkRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateScalingConfigurationShrinkRequestImageOptions) SetLoginAsNonRoot(v bool) *CreateScalingConfigurationShrinkRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestPrivatePoolOptions struct {
	// The ID of the private pool. The ID of a private pool is the same as the ID of the elasticity assurance or capacity reservation for which the private pool is generated.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the private pool that you want to use to start ECS instances. A private pool is generated when an elasticity assurance or a capacity reservation takes effect. You can select a private pool to create ECS instances. Valid values:
	//
	// 	- Open: open private pool. Auto Scaling selects a matching open private pool to start instances. If no matching open private pools are found, Auto Scaling uses the resources in the public pool to start instances. In this case, you do not need to specify PrivatePoolOptions.Id.
	//
	// 	- Target: specified private pool. Auto Scaling uses the resources in the specified private pool to start ECS instances. If the specified private pool is unavailable, Auto Scaling cannot start ECS instances. If you set this parameter to Target, you must specify PrivatePoolOptions.Id.
	//
	// 	- None: no private pool. Auto Scaling does not use the resources in private pools to start ECS instances.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) SetId(v string) *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateScalingConfigurationShrinkRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestSystemDisk struct {
	// The ID of the automatic snapshot policy that you want to apply to the system disk.
	//
	// example:
	//
	// sp-bp12m37ccmxvbmi5****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burst feature for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// > This parameter is available only if you set `SystemDisk.Category` to `cloud_auto`.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
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
	// 	- cloud_essd: enhanced SSD (ESSD)
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// If you specify SystemDisk.Category, you cannot specify `SystemDiskCategories`. If you do not specify SystemDisk.Category or `SystemDiskCategories`, the default value of SystemDisk.Category is used.
	//
	// 	- For I/O optimized instances, the default value is cloud_efficiency.
	//
	// 	- For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length. The description can contain letters and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test system disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The encryption algorithm that you want to use to encrypt the system disk. Valid values:
	//
	// 	- AES-256
	//
	// 	- SM4-128
	//
	// Default value: AES-256
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
	// Default value: false
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
	// The performance level (PL) of the system disk that is an enhanced SSD (ESSD). Valid values:
	//
	// 	- PL0: An ESSD can provide up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can provide up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can provide up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD can provide up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
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
	// The size of the system disk. Unit: GiB.
	//
	// 	- Basic disk: 20 to 500.
	//
	// 	- ESSD (cloud_essd): The valid values vary based on the performance level of the ESSD.
	//
	//     	- PL0 ESSD: 1 to 2048.
	//
	//     	- PL1 ESSD: 20 to 2048.
	//
	//     	- PL2 ESSD: 461 to 2048.
	//
	//     	- PL3 ESSD: 1261 to 2048.
	//
	// 	- ESSD AutoPL disk (cloud_auto): 1 to 2048.
	//
	// 	- Other disk categories: 20 to 2048.
	//
	// The value of this parameter must be at least 1 and greater than or equal to the image size.
	//
	// Default value: 40 or the size of the image, whichever is larger.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetBurstingEnabled(v bool) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetDescription(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetDiskName(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetEncryptAlgorithm(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetEncrypted(v bool) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetKMSKeyId(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetPerformanceLevel(v string) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetProvisionedIops(v int64) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestCustomPriorities struct {
	// The ECS instance type.
	//
	// >  The ECS instance type must be included in the instance types specified in the scaling configuration.
	//
	// example:
	//
	// ecs.g6.large
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

func (s CreateScalingConfigurationShrinkRequestCustomPriorities) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestCustomPriorities) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestCustomPriorities) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationShrinkRequestCustomPriorities) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateScalingConfigurationShrinkRequestCustomPriorities) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestCustomPriorities {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestCustomPriorities) SetVswitchId(v string) *CreateScalingConfigurationShrinkRequestCustomPriorities {
	s.VswitchId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestCustomPriorities) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestDataDisks struct {
	// The ID of the automatic snapshot policy that you want to apply to the data disk.
	//
	// example:
	//
	// sp-bp19nq9enxqkomib****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the burst feature for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// > This parameter is available only if you set `SystemDisk.Category` to `cloud_auto`.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The categories of the data disks. If Auto Scaling cannot create instances by using the disk category that has the highest priority, Auto Scaling creates instances by using the disk category that has the next highest priority. Valid values:
	//
	// 	- cloud: basic disk. For a basic disk that is created together with the instance, DeleteWithInstance is set to true.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// > If you specify Categories, you cannot specify `DataDisks.Category`.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The category of the data disk. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	//
	// 	- ephemeral_ssd: local SSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// If you specify this parameter, you cannot specify Categories. If you do not specify Category or Categories, the default value of Category is used.
	//
	// 	- For I/O optimized instances, the default value is cloud_efficiency.
	//
	// 	- For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance to which the data disk is attached is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is available only for independent disks whose value of Category is set to cloud, cloud_efficiency, cloud_ssd, or cloud_essd. If you specify this parameter for other disks, an error is reported.
	//
	// Default value: true
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length. The description can contain letters and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Test data disk.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount target of the data disk. If you do not specify Device, a mount target is automatically assigned when Auto Scaling creates ECS instances. The names of mount targets range from /dev/xvdb to /dev/xvdz.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the data disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key that you want to use to encrypt the data disk.
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
	// > For more information about how to select ESSD PLs, see [ESSD](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The IOPS metric that is preconfigured for the data disk.
	//
	// > IOPS measures the number of read and write operations that an EBS device can process per second.
	//
	// example:
	//
	// 100
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the data disk. Unit: GiB. Valid values:
	//
	// 	- If you set Categories to cloud: 5 to 2000.
	//
	// 	- If you set Categories to cloud_efficiency: 20 to 32768.
	//
	// 	- If you set Categories to cloud_essd: 20 to 32768.
	//
	// 	- If you set Categories to ephemeral_ssd: 5 to 800.
	//
	// The size of the data disk must be greater than or equal to the size of the snapshot that is specified by SnapshotId.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot that you want to use to create data disks. If you specify this parameter, DataDisks.Size is ignored. The size of the data disk is the same as the size of the specified snapshot.
	//
	// If you specify a snapshot that is created on or before July 15, 2013, the operation fails and the system returns InvalidSnapshot.TooOld.
	//
	// example:
	//
	// s-280s7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestDataDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestDataDisks) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetCategories() []*string {
	return s.Categories
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetCategory() *string {
	return s.Category
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetDescription() *string {
	return s.Description
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetDevice() *string {
	return s.Device
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetAutoSnapshotPolicyId(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetBurstingEnabled(v bool) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetCategories(v []*string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Categories = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetCategory(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDeleteWithInstance(v bool) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDescription(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Description = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDevice(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetDiskName(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.DiskName = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetEncrypted(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Encrypted = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetKMSKeyId(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.KMSKeyId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetPerformanceLevel(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetProvisionedIops(v int64) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetSize(v int32) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) SetSnapshotId(v string) *CreateScalingConfigurationShrinkRequestDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestDataDisks) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestInstancePatternInfos struct {
	// The architecture types of the instance types. Valid values:
	//
	// 	- X86: x86 architecture.
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
	// Take note of the following items:
	//
	// 	- InstancePatternInfos applies only to scaling groups that reside in virtual private clouds (VPCs).
	//
	// 	- If you specify InstancePatternInfos, you must also specify InstancePatternInfos.Cores and InstancePatternInfos.Memory.
	//
	// 	- If you specify InstanceType or InstanceTypes, Auto Scaling preferentially uses the instance type specified by InstanceType or InstanceTypes to create instances during scale-out events. If the specified instance type has insufficient inventory, Auto Scaling uses the lowest-priced instance type specified by InstancePatternInfos to create instances during scale-out events.
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
	// >  You can specify up to 10 categories.
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
	// 	- High Clock Speed: instance type that has high clock speeds.
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
	// The level of the instance family. You can specify this parameter to match the available instance types. This parameter takes effect only if you set `CostOptimization` to true. Valid values:
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
	// The maximum hourly price of pay-as-you-go or preemptible instances in intelligent configuration mode. You can specify this parameter to filter the available instance types.
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
	// The processor models of the instance types. You can specify up to 10 processor models.
	PhysicalProcessorModels []*string `json:"PhysicalProcessorModels,omitempty" xml:"PhysicalProcessorModels,omitempty" type:"Repeated"`
}

func (s CreateScalingConfigurationShrinkRequestInstancePatternInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestInstancePatternInfos) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetArchitectures() []*string {
	return s.Architectures
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetBurstablePerformance() *string {
	return s.BurstablePerformance
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetCores() *int32 {
	return s.Cores
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetExcludedInstanceTypes() []*string {
	return s.ExcludedInstanceTypes
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetGpuSpecs() []*string {
	return s.GpuSpecs
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetInstanceCategories() []*string {
	return s.InstanceCategories
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetInstanceFamilyLevel() *string {
	return s.InstanceFamilyLevel
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetInstanceTypeFamilies() []*string {
	return s.InstanceTypeFamilies
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMaxPrice() *float32 {
	return s.MaxPrice
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMaximumCpuCoreCount() *int32 {
	return s.MaximumCpuCoreCount
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMaximumGpuAmount() *int32 {
	return s.MaximumGpuAmount
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMaximumMemorySize() *float32 {
	return s.MaximumMemorySize
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMemory() *float32 {
	return s.Memory
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumBaselineCredit() *int32 {
	return s.MinimumBaselineCredit
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumCpuCoreCount() *int32 {
	return s.MinimumCpuCoreCount
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumEniIpv6AddressQuantity() *int32 {
	return s.MinimumEniIpv6AddressQuantity
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumEniPrivateIpAddressQuantity() *int32 {
	return s.MinimumEniPrivateIpAddressQuantity
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumEniQuantity() *int32 {
	return s.MinimumEniQuantity
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumGpuAmount() *int32 {
	return s.MinimumGpuAmount
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumInitialCredit() *int32 {
	return s.MinimumInitialCredit
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetMinimumMemorySize() *float32 {
	return s.MinimumMemorySize
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) GetPhysicalProcessorModels() []*string {
	return s.PhysicalProcessorModels
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetArchitectures(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Architectures = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetBurstablePerformance(v string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.BurstablePerformance = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetCores(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Cores = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetCpuArchitectures(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.CpuArchitectures = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetExcludedInstanceTypes(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.ExcludedInstanceTypes = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetGpuSpecs(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.GpuSpecs = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceCategories(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceCategories = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceFamilyLevel(v string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetInstanceTypeFamilies(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.InstanceTypeFamilies = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMaxPrice(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaxPrice = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMaximumCpuCoreCount(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaximumCpuCoreCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMaximumGpuAmount(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaximumGpuAmount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMaximumMemorySize(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MaximumMemorySize = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMemory(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.Memory = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumBaselineCredit(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumBaselineCredit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumCpuCoreCount(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumCpuCoreCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumEniIpv6AddressQuantity(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumEniIpv6AddressQuantity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumEniPrivateIpAddressQuantity(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumEniPrivateIpAddressQuantity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumEniQuantity(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumEniQuantity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumGpuAmount(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumGpuAmount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumInitialCredit(v int32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumInitialCredit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetMinimumMemorySize(v float32) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.MinimumMemorySize = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) SetPhysicalProcessorModels(v []*string) *CreateScalingConfigurationShrinkRequestInstancePatternInfos {
	s.PhysicalProcessorModels = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstancePatternInfos) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestInstanceTypeOverrides struct {
	// If you want to scale instances in the scaling group based on the weight of an instance type, you must specify this property and WeightedCapacity.
	//
	// The instance type specified by using this parameter overwrites the instance type of the launch template. You can specify N instance types by using the Extend Launch Template feature. You can specify 1 to 10 memory sizes, indicated by N.
	//
	// >  This parameter takes effect only if you specify LaunchTemplateId.
	//
	// You can use this parameter to specify any instance types that are available for purchase.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// If you need to specify the capacity of the instance type in the scaling configuration, you must specify this parameter after you specify InstanceTypeOverrides.InstanceType.
	//
	// The weight specifies the capacity of an instance of the specified instance type in the scaling group. A higher weight specifies that a smaller number of instances of the specified instance type are required to meet the expected capacity requirement.
	//
	// Performance metrics such as the number of vCPUs and memory size vary with each instance type. You can specify different weights for different instance types based on your business requirements.
	//
	// For example:
	//
	// 	- Current capacity: 0.
	//
	// 	- Expected capacity: 6
	//
	// 	- Capacity of ecs.c5.xlarge: 4.
	//
	// To reach the expected capacity, Auto Scaling must scale out two instances of ecs.c5.xlarge.
	//
	// >  The total capacity of the scaling group is constrained and cannot surpass the combined total of the maximum group size defined by MaxSize and the highest weight assigned to any instance type.
	//
	// Valid values of WeightedCapacity: 1 to 500.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) SetWeightedCapacity(v int32) *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestInstanceTypeOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestNetworkInterfaces struct {
	// example:
	//
	// Primary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// example:
	//
	// HighPerformance
	NetworkInterfaceTrafficMode    *string   `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	SecondaryPrivateIpAddressCount *int32    `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	SecurityGroupIds               []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s CreateScalingConfigurationShrinkRequestNetworkInterfaces) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestNetworkInterfaces) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestNetworkInterfaces {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) SetIpv6AddressCount(v int32) *CreateScalingConfigurationShrinkRequestNetworkInterfaces {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) SetNetworkInterfaceTrafficMode(v string) *CreateScalingConfigurationShrinkRequestNetworkInterfaces {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) SetSecondaryPrivateIpAddressCount(v int32) *CreateScalingConfigurationShrinkRequestNetworkInterfaces {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) SetSecurityGroupIds(v []*string) *CreateScalingConfigurationShrinkRequestNetworkInterfaces {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestNetworkInterfaces) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestResourcePoolOptions struct {
	// The IDs of private pools. The ID of a private pool is the same as the ID of the elasticity assurance or capacity reservation that is associated with the private pool. You can specify the IDs of only targeted private pools for this parameter.
	PrivatePoolIds []*string `json:"PrivatePoolIds,omitempty" xml:"PrivatePoolIds,omitempty" type:"Repeated"`
	// The resource pool used for instance creation, which can be the public pool or a private pool associated with any active elasticity assurance or capacity reservation. Valid values:
	//
	// 	- PrivatePoolFirst: prioritizes private pools. When this option is set along with ResourcePoolOptions.PrivatePoolIds, the specified private pools are used first. If you leave ResourcePoolOptions.PrivatePoolIds empty or if the specified private pools lack sufficient capacity, the system will automatically use available open private pools instead. If no matching private pools are available, the system defaults to the public pool.
	//
	// 	- PrivatePoolOnly: uses only private pools. If you set this value, you must specify ResourcePoolOptions.PrivatePoolIds. If the specified private pools lack sufficient capacity, instance creation will fail.
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

func (s CreateScalingConfigurationShrinkRequestResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *CreateScalingConfigurationShrinkRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateScalingConfigurationShrinkRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *CreateScalingConfigurationShrinkRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestResourcePoolOptions) SetStrategy(v string) *CreateScalingConfigurationShrinkRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestSecurityOptions struct {
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

func (s CreateScalingConfigurationShrinkRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *CreateScalingConfigurationShrinkRequestSecurityOptions) SetConfidentialComputingMode(v string) *CreateScalingConfigurationShrinkRequestSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationShrinkRequestSpotPriceLimits struct {
	// The instance type of the spot instances. This parameter takes effect only if you set SpotStrategy to SpotWithPriceLimit.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The price limit of the spot instances. This parameter takes effect only if you set SpotStrategy to SpotWithPriceLimit.
	//
	// example:
	//
	// 0.5
	PriceLimit *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s CreateScalingConfigurationShrinkRequestSpotPriceLimits) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationShrinkRequestSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) GetPriceLimit() *float32 {
	return s.PriceLimit
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) SetInstanceType(v string) *CreateScalingConfigurationShrinkRequestSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) SetPriceLimit(v float32) *CreateScalingConfigurationShrinkRequestSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

func (s *CreateScalingConfigurationShrinkRequestSpotPriceLimits) Validate() error {
	return dara.Validate(s)
}
