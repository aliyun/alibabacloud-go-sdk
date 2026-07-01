// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHibernationOptions(v *CreateInstanceRequestHibernationOptions) *CreateInstanceRequest
	GetHibernationOptions() *CreateInstanceRequestHibernationOptions
	SetPrivatePoolOptions(v *CreateInstanceRequestPrivatePoolOptions) *CreateInstanceRequest
	GetPrivatePoolOptions() *CreateInstanceRequestPrivatePoolOptions
	SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest
	GetSystemDisk() *CreateInstanceRequestSystemDisk
	SetAffinity(v string) *CreateInstanceRequest
	GetAffinity() *string
	SetArn(v []*CreateInstanceRequestArn) *CreateInstanceRequest
	GetArn() []*CreateInstanceRequestArn
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstanceRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateInstanceRequest
	GetClusterId() *string
	SetCreditSpecification(v string) *CreateInstanceRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest
	GetDataDisk() []*CreateInstanceRequestDataDisk
	SetDedicatedHostId(v string) *CreateInstanceRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *CreateInstanceRequest
	GetDeletionProtection() *bool
	SetDeploymentSetGroupNo(v int32) *CreateInstanceRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *CreateInstanceRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateInstanceRequest
	GetDryRun() *bool
	SetHostName(v string) *CreateInstanceRequest
	GetHostName() *string
	SetHpcClusterId(v string) *CreateInstanceRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *CreateInstanceRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateInstanceRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateInstanceRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *CreateInstanceRequest
	GetImageFamily() *string
	SetImageId(v string) *CreateInstanceRequest
	GetImageId() *string
	SetInnerIpAddress(v string) *CreateInstanceRequest
	GetInnerIpAddress() *string
	SetInstanceChargeType(v string) *CreateInstanceRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateInstanceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateInstanceRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *CreateInstanceRequest
	GetKeyPairName() *string
	SetNodeControllerId(v string) *CreateInstanceRequest
	GetNodeControllerId() *string
	SetOwnerAccount(v string) *CreateInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateInstanceRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *CreateInstanceRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateInstanceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateInstanceRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateInstanceRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateInstanceRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateInstanceRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateInstanceRequest
	GetSecurityGroupId() *string
	SetSpotDuration(v int32) *CreateInstanceRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *CreateInstanceRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimit(v float32) *CreateInstanceRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateInstanceRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *CreateInstanceRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateInstanceRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetTenancy(v string) *CreateInstanceRequest
	GetTenancy() *string
	SetUseAdditionalService(v bool) *CreateInstanceRequest
	GetUseAdditionalService() *bool
	SetUserData(v string) *CreateInstanceRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateInstanceRequest
	GetVSwitchId() *string
	SetVlanId(v string) *CreateInstanceRequest
	GetVlanId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
}

type CreateInstanceRequest struct {
	HibernationOptions *CreateInstanceRequestHibernationOptions `json:"HibernationOptions,omitempty" xml:"HibernationOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *CreateInstanceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *CreateInstanceRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether the instance is associated with a dedicated host. Valid values:
	//
	// - default: The instance is not associated with a dedicated host. When an instance that has economical mode enabled is restarted after it is stopped, the instance is deployed to another dedicated host in the automatic deployment resource pool if the resources of the original dedicated host are insufficient.
	//
	// - host: The instance is associated with a dedicated host. When an instance that has economical mode enabled is restarted after it is stopped, the instance remains on the original dedicated host. If the resources of the original dedicated host are insufficient, the instance fails to restart.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	Arn []*CreateInstanceRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false (default): disables auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. This parameter is required when AutoRenew is set to True.
	//
	// <props="china">If PeriodUnit is set to Week, valid values of AutoRenewPeriod are 1, 2, and 3.
	//
	// If PeriodUnit is set to Month, valid values of AutoRenewPeriod are 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID of the instance.
	//
	// > This parameter will be deprecated. To improve compatibility, use other parameters instead.
	//
	// example:
	//
	// c-bp67acfmxazb4p****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information, see the performance constrained mode section in [What are burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information, see the unlimited performance mode section in [What are burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The list of data disks.
	DataDisk []*CreateInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the dedicated host.
	//
	// <props="china">You can call [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) to query the list of dedicated host IDs.
	//
	// <props="intl">You can call [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) to query the list of dedicated host IDs.
	//
	// 	Notice: Spot instances cannot be created on dedicated hosts. If you specify `DedicatedHostId`, the `SpotStrategy` and `SpotPriceLimit` settings in the request are automatically ignored.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection attribute of the instance. Specifies whether the instance can be released from the ECS console or by calling [DeleteInstance](https://help.aliyun.com/document_detail/25507.html).
	//
	// -   true: enables release protection.
	//
	// -   false (default): disables release protection.
	//
	// > This attribute is applicable only to pay-as-you-go instances. It can only prevent manual release, not system-initiated release.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The group number of the instance in the deployment set. This parameter takes effect only when the deployment set uses the high availability group strategy (AvailabilityGroup). Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1brhwhoqinyjd6****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// example:
	//
	// InstanceTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run. The system checks whether the required parameters are specified, whether the request format is valid, whether service limits are met, and whether the specified ECS resources are available. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - false (default): performs a dry run and sends the request. If the check succeeds, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the server.
	//
	// - The hostname cannot start or end with a period (.) or hyphen (-), and cannot contain consecutive periods or hyphens.
	//
	// - Windows instances: The hostname must be 2 to 15 characters in length and cannot contain periods (.) or consist entirely of digits. It can contain uppercase and lowercase letters, digits, and hyphens (-).
	//
	// - Other instances (such as Linux): The hostname must be 2 to 64 characters in length and can contain multiple periods (.). Each segment separated by a period can contain uppercase and lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// LocalHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the HPC cluster to which the instance belongs.
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enabled.
	//
	// - disabled: disabled.
	//
	// Default value: enabled.
	//
	// > For more information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/49122.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// - optional: does not forcefully use the security-hardened mode.
	//
	// - required: forcefully uses the security-hardened mode. After this value is set, instance metadata cannot be accessed in normal mode.
	//
	// Default value: optional.
	//
	// > For more information about the modes for accessing instance metadata, see [Instance metadata access modes](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. Set this parameter to obtain the latest available image from the specified image family to create the instance.
	//
	// - If `ImageId` is specified, you cannot set this parameter.
	//
	// - If `ImageId` is not specified, you can set this parameter.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image used to start the instance. To use an Alibaba Cloud Marketplace image, you can view the `ImageId` on the image product page. If you do not use `ImageFamily` to select the latest available image from an image family, this parameter is required.
	//
	// example:
	//
	// ubuntu_18_04_64_20G_alibase_20190624.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address of the instance.
	//
	// example:
	//
	// ``192.168.**.**``
	InnerIpAddress *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, make sure that your account supports credit payment or balance payment. Otherwise, an `InvalidPayMethod` error is returned.
	//
	// - PostPaid (default): pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length and can contain letters in the Unicode letter category (including English and Chinese characters) and digits. The name can contain colons (:), underscores (_), periods (.), or hyphens (-). If this parameter is not specified, the default value is the instance ID.
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// - Instance type selection: See [Instance family](https://help.aliyun.com/document_detail/25378.html) or invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the performance data of the target instance type. You can also see [Best practices for instance type selection](https://help.aliyun.com/document_detail/58291.html) to learn how to select an instance type.
	//
	// - Check active resources: Invoke [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) to query active resources in a specific region or zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic (default): pay-by-traffic.
	//
	// > In **pay-by-traffic*	- mode, the peak inbound and outbound bandwidths are both upper limits and are not guaranteed. When resource contention occurs, the peak bandwidth may be throttled. If your workloads require guaranteed bandwidth, use **pay-by-bandwidth*	- mode.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth, in Mbit/s. Valid values:
	//
	// - If the purchased outbound bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - If the purchased outbound bandwidth is greater than 10 Mbit/s: 1 to the value of `InternetMaxBandwidthOut`. Default value: the value of `InternetMaxBandwidthOut`.
	//
	// example:
	//
	// 50
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth, in Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// - none: not I/O optimized.
	//
	// - optimized: I/O optimized.
	//
	// The default value for [retired instance types](https://help.aliyun.com/document_detail/55263.html) is none.
	//
	// The default value for other instance types is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// > For Windows instances, this parameter is ignored. The default value is empty. Even if this parameter is specified, only the `Password` content is used.
	//
	// example:
	//
	// KeyPairTestName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	NodeControllerId *string `json:"NodeControllerId,omitempty" xml:"NodeControllerId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// ```
	//
	// ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// ```
	//
	// Note the following items:
	//
	// - For security purposes, use HTTPS to send requests if the Password parameter is specified.
	//
	// - For Windows instances, the password cannot start with a forward slash (/).
	//
	// - For instances that run certain operating systems, passwords are not supported. Only key pairs are supported. Examples: Others Linux and Fedora CoreOS.
	//
	// example:
	//
	// TestEcs123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. When you use this parameter, the Password parameter must be empty. Make sure that the image has a password configured.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the instance. The unit is specified by `PeriodUnit`. This parameter is required and takes effect only when `InstanceChargeType` is set to `PrePaid`. If `DedicatedHostId` is specified, the value of this parameter cannot exceed the subscription duration of the dedicated host. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week, valid values of Period are 1, 2, 3, and 4.
	//
	// - If PeriodUnit is set to Month, valid values of Period are 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">If PeriodUnit is set to Month, valid values of Period are 1, 2, 3, 6, and 12.
	//
	// <props="partner">If PeriodUnit is set to Month, valid values of Period are 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// <props="china">
	//
	// - Week.
	//
	// - Month.
	//
	//
	//
	// <props="intl">Month.
	//
	// <props="partner">Month.
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address of the instance. The IP address must be an available address within the CIDR block of the specified vSwitch (VSwitchId).
	//
	// example:
	//
	// 172.16.236.*
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance RAM role. You can call the RAM API [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query the instance RAM roles that you have created.
	//
	// example:
	//
	// RAMTestName
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	// The ID of the security group to which the instance belongs.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The protection period of the spot instance, in hours. Default value: 1. Valid values:
	//
	// - 1: After the instance is created, Alibaba Cloud ensures that the instance is not automatically released for 1 hour. After 1 hour, the system automatically compares the bid price with the market price and checks resource availability to determine whether to retain automatic release the instance.
	//
	// - 0: After the instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system automatically compares the bid price with the market price and checks resource availability to determine whether to retain automatic release the instance.
	//
	// >
	//
	// > - This parameter supports only the value 0 or 1.
	//
	// > - Spot instances are billed by second. Set the protection period based on the expected task execution duration.
	//
	// > - Alibaba Cloud sends an ECS system event notification 5 minutes before the instance is released.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The break mode of the spot instance. Valid values:
	//
	// - Terminate: The instance is directly released.
	//
	// - Stop: The instance enters economical mode.
	//
	//   For more information about economical mode, see [Economical mode for pay-as-you-go instances](https://help.aliyun.com/document_detail/63353.html).
	//
	// Default value: Terminate.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the instance. This value can be accurate to three decimal places. This parameter takes effect only when `SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.98
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the instance. This parameter takes effect only when `InstanceChargeType` is set to `PostPaid`. Valid values:
	//
	// - NoSpot (default): The instance is a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: The instance is a spot instance with a user-defined maximum hourly price.
	//
	// - SpotAsPriceGo: The instance is a spot instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp1j4i2jdf3owlhe****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. Valid values: greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The list of tags.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// - default: creates the instance on a non-dedicated host.
	//
	// - host: creates the instance on a dedicated host. If you do not specify `DedicatedHostId`, Alibaba Cloud automatically selects a dedicated host for the instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// Specifies whether to use the virtual machine system configuration provided by Alibaba Cloud (Windows: NTP and KMS. Linux: NTP and YUM).
	//
	// example:
	//
	// true
	UseAdditionalService *bool `json:"UseAdditionalService,omitempty" xml:"UseAdditionalService,omitempty"`
	// The instance user data. The data must be encoded in Base64. The maximum size of the raw data is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID. This parameter is required when you create a VPC-connected instance. You can call [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) to query available vSwitches.
	//
	// > If you specify `VSwitchId`, the specified `ZoneId` must be in the same zone as the vSwitch. You can also leave `ZoneId` empty, and the system automatically selects the zone of the specified vSwitch.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual local area network ID.
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The zone ID of the instance. For more information, call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zone list.
	//
	// > If you specify `VSwitchId`, the specified `ZoneId` must be in the same zone as the vSwitch. You can also leave `ZoneId` empty, and the system automatically selects the zone of the specified vSwitch.
	//
	// Default value: empty. The system automatically selects a zone.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetHibernationOptions() *CreateInstanceRequestHibernationOptions {
	return s.HibernationOptions
}

func (s *CreateInstanceRequest) GetPrivatePoolOptions() *CreateInstanceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateInstanceRequest) GetSystemDisk() *CreateInstanceRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateInstanceRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *CreateInstanceRequest) GetArn() []*CreateInstanceRequestArn {
	return s.Arn
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateInstanceRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateInstanceRequest) GetDataDisk() []*CreateInstanceRequestDataDisk {
	return s.DataDisk
}

func (s *CreateInstanceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateInstanceRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateInstanceRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *CreateInstanceRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateInstanceRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateInstanceRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *CreateInstanceRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateInstanceRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateInstanceRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateInstanceRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateInstanceRequest) GetInnerIpAddress() *string {
	return s.InnerIpAddress
}

func (s *CreateInstanceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateInstanceRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateInstanceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateInstanceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateInstanceRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateInstanceRequest) GetNodeControllerId() *string {
	return s.NodeControllerId
}

func (s *CreateInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateInstanceRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstanceRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateInstanceRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateInstanceRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateInstanceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateInstanceRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateInstanceRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *CreateInstanceRequest) GetUseAdditionalService() *bool {
	return s.UseAdditionalService
}

func (s *CreateInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequest) GetVlanId() *string {
	return s.VlanId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) SetHibernationOptions(v *CreateInstanceRequestHibernationOptions) *CreateInstanceRequest {
	s.HibernationOptions = v
	return s
}

func (s *CreateInstanceRequest) SetPrivatePoolOptions(v *CreateInstanceRequestPrivatePoolOptions) *CreateInstanceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateInstanceRequest) SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateInstanceRequest) SetAffinity(v string) *CreateInstanceRequest {
	s.Affinity = &v
	return s
}

func (s *CreateInstanceRequest) SetArn(v []*CreateInstanceRequestArn) *CreateInstanceRequest {
	s.Arn = v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetClusterId(v string) *CreateInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetCreditSpecification(v string) *CreateInstanceRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateInstanceRequest) SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequest) SetDedicatedHostId(v string) *CreateInstanceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateInstanceRequest) SetDeletionProtection(v bool) *CreateInstanceRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateInstanceRequest) SetDeploymentSetGroupNo(v int32) *CreateInstanceRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *CreateInstanceRequest) SetDeploymentSetId(v string) *CreateInstanceRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetDryRun(v bool) *CreateInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateInstanceRequest) SetHostName(v string) *CreateInstanceRequest {
	s.HostName = &v
	return s
}

func (s *CreateInstanceRequest) SetHpcClusterId(v string) *CreateInstanceRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpEndpoint(v string) *CreateInstanceRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpPutResponseHopLimit(v int32) *CreateInstanceRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpTokens(v string) *CreateInstanceRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateInstanceRequest) SetImageFamily(v string) *CreateInstanceRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetInnerIpAddress(v string) *CreateInstanceRequest {
	s.InnerIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceChargeType(v string) *CreateInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateInstanceRequest) SetIoOptimized(v string) *CreateInstanceRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateInstanceRequest) SetKeyPairName(v string) *CreateInstanceRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceRequest) SetNodeControllerId(v string) *CreateInstanceRequest {
	s.NodeControllerId = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerAccount(v string) *CreateInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPasswordInherit(v bool) *CreateInstanceRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetRamRoleName(v string) *CreateInstanceRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerAccount(v string) *CreateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerId(v int64) *CreateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityEnhancementStrategy(v string) *CreateInstanceRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityGroupId(v string) *CreateInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotDuration(v int32) *CreateInstanceRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotInterruptionBehavior(v string) *CreateInstanceRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotPriceLimit(v float32) *CreateInstanceRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotStrategy(v string) *CreateInstanceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSetId(v string) *CreateInstanceRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSetPartitionNumber(v int32) *CreateInstanceRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetTenancy(v string) *CreateInstanceRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateInstanceRequest) SetUseAdditionalService(v bool) *CreateInstanceRequest {
	s.UseAdditionalService = &v
	return s
}

func (s *CreateInstanceRequest) SetUserData(v string) *CreateInstanceRequest {
	s.UserData = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetVlanId(v string) *CreateInstanceRequest {
	s.VlanId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.HibernationOptions != nil {
		if err := s.HibernationOptions.Validate(); err != nil {
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
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestHibernationOptions struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
}

func (s CreateInstanceRequestHibernationOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestHibernationOptions) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestHibernationOptions) GetConfigured() *bool {
	return s.Configured
}

func (s *CreateInstanceRequestHibernationOptions) SetConfigured(v bool) *CreateInstanceRequestHibernationOptions {
	s.Configured = &v
	return s
}

func (s *CreateInstanceRequestHibernationOptions) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestPrivatePoolOptions struct {
	// The private pool ID, which is the ID of the elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool option for launching the instance. A private pool is generated when an elasticity assurance or capacity reservation takes effect. You can select a private pool when you start an instance. Valid values:
	//
	// - Open: open mode. The system automatically matches an open private pool. If no matching private pool is available, the instance is launched using public pool resources. You do not need to set `PrivatePoolOptions.Id`.
	//
	// - Target: specified mode. The instance is launched using the specified private pool. If the specified private pool is unavailable, the instance fails to launch. In this mode, you must specify the private pool ID. Set `PrivatePoolOptions.Id` to the private pool ID.
	//
	// - None: does not use a private pool. The instance is launched without using any private pool.
	//
	// Default value: None.
	//
	// In any of the following scenarios, the private pool option can only be set to `None` or left empty:
	//
	// - Creating a spot instance.
	//
	// - Creating an ECS instance on a dedicated host.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateInstanceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *CreateInstanceRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateInstanceRequestPrivatePoolOptions) SetId(v string) *CreateInstanceRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateInstanceRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateInstanceRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateInstanceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestSystemDisk struct {
	// The category of the system disk. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD (ESSD).
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// > The cloud_essd_entry value is supported only when `InstanceType` is set to the [u1, universal instance family](https://help.aliyun.com/document_detail/457079.html) (`ecs.u1`) or the [e, economy instance family](https://help.aliyun.com/document_detail/108489.html) (`ecs.e`). Settings of this parameter determine the computing power and optimization level of the system disk.
	//
	// For retired instance types that are not I/O optimized instances, the default value is cloud. For other instance types, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// Default value: empty.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length and can contain letters in the Unicode letter category (including English and Chinese characters and digits). The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Default value: empty.
	//
	// example:
	//
	// SystemDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The performance level (PL) of the enterprise SSD used as the system disk. Settings of this parameter apply only when the disk category is not standard SSD. Valid values:
	//
	// - PL0: a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk, in GiB. Valid values:
	//
	// - Basic disk: 20 to 500
	//
	// - Other disk types: 20 to 2048
	//
	// The value of this parameter must be greater than or equal to max{20, ImageSize}.
	//
	// Default value: max{40, ImageSize}.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the dedicated block storage cluster. To use a disk in a dedicated block storage cluster as the system disk when you create an ECS instance, set this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s CreateInstanceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateInstanceRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestSystemDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateInstanceRequestSystemDisk) SetCategory(v string) *CreateInstanceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetDescription(v string) *CreateInstanceRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetDiskName(v string) *CreateInstanceRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetPerformanceLevel(v string) *CreateInstanceRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetSize(v int32) *CreateInstanceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetStorageClusterId(v string) *CreateInstanceRequestSystemDisk {
	s.StorageClusterId = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestArn struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 1234567890
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// Primary
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateInstanceRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestArn) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateInstanceRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateInstanceRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateInstanceRequestArn) SetAssumeRoleFor(v int64) *CreateInstanceRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateInstanceRequestArn) SetRoleType(v string) *CreateInstanceRequestArn {
	s.RoleType = &v
	return s
}

func (s *CreateInstanceRequestArn) SetRolearn(v string) *CreateInstanceRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CreateInstanceRequestArn) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestDataDisk struct {
	// The category of data disk N. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD (ESSD).
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	//   > The cloud_essd_entry value is supported only when `InstanceType` is set to the `ecs.u1` or `ecs.e` instance family. Settings of this parameter determine the optimization level of the data disk.
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - standard.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - Premium Edition.
	//
	// The default value is cloud_efficiency for I/O optimized instances and cloud for non-I/O optimized instances.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether the data disk is released when the instance is released.
	//
	// - true: The data disk is released when the instance is released.
	//
	// - false: The data disk is not released when the instance is released.
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
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk.
	//
	// > This parameter is applicable only to full image (system image) scenarios. You can set this parameter to the mount point of the data disk in the full image and modify the corresponding `DataDisk.N.Size` and `DataDisk.N.Category` parameters to change the disk category and size of the data disk in the full image.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters in the Unicode letter category (including English and Chinese characters and digits). The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// DataDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether data disk N is encrypted.
	//
	// - true: encrypted.
	//
	// - false: not encrypted.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The KMS key ID used by the disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level (PL) of the enterprise SSD used as a data disk. Settings of this parameter apply only when the disk category is not standard SSD. The value of N must be the same as that in `DataDisk.N.Category=cloud_essd`. Valid values:
	//
	// - PL0: a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL2
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of data disk N, in GiB. Valid values of N: 1 to 16. Valid values:
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_ssd: 20 to 32768.
	//
	// - cloud_essd: The valid values depend on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     - PL0: 1 to 65,536.
	//
	//     - PL1: 20 to 65,536.
	//
	//     - PL2: 461 to 65,536.
	//
	//     - PL3: 1261 to 65,536.
	//
	// - cloud: 5 to 2000.
	//
	// > The value of this parameter must be greater than or equal to the size of the snapshot specified by `SnapshotId`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID used to create data disk N. Valid values of N: 1 to 16.
	//
	// - If `DataDisk.N.SnapshotId` is specified, `DataDisk.N.Size` is ignored. The actual size of the created disk is the size of the specified snapshot.
	//
	// - Snapshots created on or before July 15, 2013 cannot be used. Requests with such snapshots are rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster. To use a disk in a dedicated block storage cluster as a data disk when you create an ECS instance, set this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s CreateInstanceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateInstanceRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateInstanceRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateInstanceRequestDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateInstanceRequestDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateInstanceRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateInstanceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateInstanceRequestDataDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateInstanceRequestDataDisk) SetCategory(v string) *CreateInstanceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDeleteWithInstance(v bool) *CreateInstanceRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDescription(v string) *CreateInstanceRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDevice(v string) *CreateInstanceRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDiskName(v string) *CreateInstanceRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetEncryptAlgorithm(v string) *CreateInstanceRequestDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetEncrypted(v bool) *CreateInstanceRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetKMSKeyId(v string) *CreateInstanceRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetPerformanceLevel(v string) *CreateInstanceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetSize(v int32) *CreateInstanceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetSnapshotId(v string) *CreateInstanceRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetStorageClusterId(v string) *CreateInstanceRequestDataDisk {
	s.StorageClusterId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTag struct {
	// The tag key for the instance, disk, and primary ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for the instance, disk, and primary ENI. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
