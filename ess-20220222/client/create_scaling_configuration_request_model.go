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
	CustomPriorities []*CreateScalingConfigurationRequestCustomPriorities `json:"CustomPriorities,omitempty" xml:"CustomPriorities,omitempty" type:"Repeated"`
	// The data disks.
	DataDisks []*CreateScalingConfigurationRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
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
	InstancePatternInfos []*CreateScalingConfigurationRequestInstancePatternInfos `json:"InstancePatternInfos,omitempty" xml:"InstancePatternInfos,omitempty" type:"Repeated"`
	// The instance type of the ECS instance. For more information, see the [Instance families](https://help.aliyun.com/document_detail/25378.html) topic.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The information about instance types.
	InstanceTypeOverrides []*CreateScalingConfigurationRequestInstanceTypeOverrides `json:"InstanceTypeOverrides,omitempty" xml:"InstanceTypeOverrides,omitempty" type:"Repeated"`
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
	Memory            *int32                                                `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkInterfaces []*CreateScalingConfigurationRequestNetworkInterfaces `json:"NetworkInterfaces,omitempty" xml:"NetworkInterfaces,omitempty" type:"Repeated"`
	OwnerAccount      *string                                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	ResourcePoolOptions *CreateScalingConfigurationRequestResourcePoolOptions `json:"ResourcePoolOptions,omitempty" xml:"ResourcePoolOptions,omitempty" type:"Struct"`
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
	SchedulerOptions map[string]interface{} `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty"`
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
	SecurityGroupIds []*string                                         `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *CreateScalingConfigurationRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
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
	SpotPriceLimits []*CreateScalingConfigurationRequestSpotPriceLimits `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Repeated"`
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

type CreateScalingConfigurationRequestInstanceTypeOverrides struct {
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

func (s CreateScalingConfigurationRequestResourcePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationRequestResourcePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) GetPrivatePoolIds() []*string {
	return s.PrivatePoolIds
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) SetPrivatePoolIds(v []*string) *CreateScalingConfigurationRequestResourcePoolOptions {
	s.PrivatePoolIds = v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) SetStrategy(v string) *CreateScalingConfigurationRequestResourcePoolOptions {
	s.Strategy = &v
	return s
}

func (s *CreateScalingConfigurationRequestResourcePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingConfigurationRequestSecurityOptions struct {
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
