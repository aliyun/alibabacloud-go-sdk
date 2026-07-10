// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *CreateLaunchTemplateRequestSystemDisk) *CreateLaunchTemplateRequest
	GetSystemDisk() *CreateLaunchTemplateRequestSystemDisk
	SetAutoReleaseTime(v string) *CreateLaunchTemplateRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *CreateLaunchTemplateRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateLaunchTemplateRequest
	GetAutoRenewPeriod() *int32
	SetCreditSpecification(v string) *CreateLaunchTemplateRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateLaunchTemplateRequestDataDisk) *CreateLaunchTemplateRequest
	GetDataDisk() []*CreateLaunchTemplateRequestDataDisk
	SetDeletionProtection(v bool) *CreateLaunchTemplateRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *CreateLaunchTemplateRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateLaunchTemplateRequest
	GetDescription() *string
	SetEnableVmOsConfig(v bool) *CreateLaunchTemplateRequest
	GetEnableVmOsConfig() *bool
	SetHostName(v string) *CreateLaunchTemplateRequest
	GetHostName() *string
	SetHttpEndpoint(v string) *CreateLaunchTemplateRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateLaunchTemplateRequest
	GetHttpTokens() *string
	SetImageId(v string) *CreateLaunchTemplateRequest
	GetImageId() *string
	SetImageOptions(v *CreateLaunchTemplateRequestImageOptions) *CreateLaunchTemplateRequest
	GetImageOptions() *CreateLaunchTemplateRequestImageOptions
	SetImageOwnerAlias(v string) *CreateLaunchTemplateRequest
	GetImageOwnerAlias() *string
	SetInstanceChargeType(v string) *CreateLaunchTemplateRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateLaunchTemplateRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateLaunchTemplateRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateLaunchTemplateRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateLaunchTemplateRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *CreateLaunchTemplateRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *CreateLaunchTemplateRequest
	GetKeyPairName() *string
	SetLaunchTemplateName(v string) *CreateLaunchTemplateRequest
	GetLaunchTemplateName() *string
	SetNetworkInterface(v []*CreateLaunchTemplateRequestNetworkInterface) *CreateLaunchTemplateRequest
	GetNetworkInterface() []*CreateLaunchTemplateRequestNetworkInterface
	SetNetworkType(v string) *CreateLaunchTemplateRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateLaunchTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLaunchTemplateRequest
	GetOwnerId() *int64
	SetPasswordInherit(v bool) *CreateLaunchTemplateRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateLaunchTemplateRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateLaunchTemplateRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateLaunchTemplateRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateLaunchTemplateRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateLaunchTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLaunchTemplateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLaunchTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLaunchTemplateRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateLaunchTemplateRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateLaunchTemplateRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *CreateLaunchTemplateRequestSecurityOptions) *CreateLaunchTemplateRequest
	GetSecurityOptions() *CreateLaunchTemplateRequestSecurityOptions
	SetSpotDuration(v int32) *CreateLaunchTemplateRequest
	GetSpotDuration() *int32
	SetSpotPriceLimit(v float32) *CreateLaunchTemplateRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateLaunchTemplateRequest
	GetSpotStrategy() *string
	SetTag(v []*CreateLaunchTemplateRequestTag) *CreateLaunchTemplateRequest
	GetTag() []*CreateLaunchTemplateRequestTag
	SetTemplateResourceGroupId(v string) *CreateLaunchTemplateRequest
	GetTemplateResourceGroupId() *string
	SetTemplateTag(v []*CreateLaunchTemplateRequestTemplateTag) *CreateLaunchTemplateRequest
	GetTemplateTag() []*CreateLaunchTemplateRequestTemplateTag
	SetUserData(v string) *CreateLaunchTemplateRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateLaunchTemplateRequest
	GetVSwitchId() *string
	SetVersionDescription(v string) *CreateLaunchTemplateRequest
	GetVersionDescription() *string
	SetVpcId(v string) *CreateLaunchTemplateRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateLaunchTemplateRequest
	GetZoneId() *string
}

type CreateLaunchTemplateRequest struct {
	SystemDisk *CreateLaunchTemplateRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The automatic release time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// - If the value of seconds (`ss`) is not `00`, the time is automatically rounded down to the start of the current minute (`mm`).
	//
	// - The earliest release time is 30 minutes after the current time.
	//
	// - The latest release time cannot be more than three years from the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false: does not enable auto-renewal.
	//
	// Default value: false.
	//
	// > This parameter takes effect only when `InstanceChargeType` is set to `PrePaid`.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	//
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week: 1, 2, and 3.
	//
	// - If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">If PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The running mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information, see the performance constrained mode section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information, see the unlimited mode section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The list of data disk information.
	DataDisk []*CreateLaunchTemplateRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The release protection attribute of the instance. Specifies whether the instance can be released from the console or by calling [DeleteInstance](https://help.aliyun.com/document_detail/25507.html). Valid values:
	//
	// - true: enables release protection.
	//
	// - false: disables release protection.
	//
	// Default value: false.
	//
	// > This attribute applies only to pay-as-you-go instances. It can only restrict manual release operations and does not affect system-initiated releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1brhwhoqinyjd6****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testECSDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the operating system configuration of the instance.
	//
	// > This parameter will be deprecated. For better compatibility, use other parameters.
	//
	// example:
	//
	// false
	EnableVmOsConfig *bool `json:"EnableVmOsConfig,omitempty" xml:"EnableVmOsConfig,omitempty"`
	// The hostname of the Elastic Compute Service server.
	//
	// -   The hostname cannot start or end with a period (.) or a hyphen (-). It cannot contain consecutive periods (.) or hyphens (-).
	//
	// -   Windows instances: The hostname must be 2 to 15 characters in length and cannot contain periods (.) or consist entirely of digits. It can contain letters, digits, and hyphens (-).
	//
	// -   Other instances (such as Linux): The hostname must be 2 to 64 characters in length. It can contain multiple periods (.), with each segment between periods allowing letters, digits, and hyphens (-).
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enables the access channel.
	//
	// - disabled: disables the access channel.
	//
	// Default value: enabled.
	//
	// > For more information about instance metadata, see [Overview of ECS instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// > This parameter is not available for use.
	//
	// example:
	//
	// 3
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// - optional: does not forcefully use the security-hardened mode.
	//
	// - required: forcefully uses the security-hardened mode. After you set this value, instance metadata cannot be accessed in normal mode.
	//
	// Default value: optional.
	//
	// > For more information about the modes for
	//
	// [Instance metadata overview](https://help.aliyun.com/document_detail/108460.html)
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The image ID. You can call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available image resources.
	//
	// example:
	//
	// win2008r2_64_ent_sp1_en-us_40G_alibase_20170915.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image-related property information.
	ImageOptions *CreateLaunchTemplateRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The source of the image. Valid values:
	//
	// - system: a public image provided by Alibaba Cloud.
	//
	// - self: a custom image that you created.
	//
	// - others: a shared image from another Alibaba Cloud account.
	//
	// - marketplace: an image from <props="china"><ph>[Alibaba Cloud Marketplace](https://market.aliyun.com/)</ph><props="intl"><ph>[Alibaba Cloud Marketplace](https://marketplace.alibabacloud.com/)</ph>. You can use Alibaba Cloud Marketplace images directly without subscribing to them first. Check the billing details of the Alibaba Cloud Marketplace image on your own.
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, <props="china"><ph>confirm that your account supports balance payments and credit payments</ph><props="intl"><ph>confirm that your account supports credit payments</ph>. Otherwise, an `InvalidPayMethod` error is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length and can contain letters, digits, and characters from the Unicode letter category (including Chinese characters). The name can contain colons (:), underscores (_), periods (.), and hyphens (-). The default value is the `InstanceId` of the instance.
	//
	// When you create multiple ECS instances, you can specify sequential instance names. The names can contain brackets ([]) and commas (,). For more information, see [Specify sequential instance names or hostnames](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. For more information, see [Instance family](https://help.aliyun.com/document_detail/25378.html). You can also invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the most recent instance type list.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// > In **pay-by-traffic*	- mode, the peak inbound and outbound bandwidths are used as bandwidth upper limits instead of guaranteed service metrics. When resource contention occurs, the peak bandwidth may be limited. If your workloads require guaranteed bandwidth, use the **pay-by-bandwidth*	- mode.
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
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// -   none: not I/O optimized.
	//
	// -   optimized: I/O optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of randomly generated IPv6 addresses to assign to the primary ENI. Valid values: 1 to 10.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair.
	//
	// -   For Windows instances, this parameter is ignored. Even if you specify this parameter, only the `Password` content is used.
	//
	// -   For Linux instances, password-based logon is disabled during initialization.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The name of the launch template. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testLaunchTemplateName
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The network interface controller (NIC) information.
	NetworkInterface []*CreateLaunchTemplateRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// - classic: classic network. This feature has been retired. For more information, see [Retirement notice](https://help.aliyun.com/document_detail/2833134.html).
	//
	// - vpc: VPC.
	//
	// example:
	//
	// vpc
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to use the password preset in the image.
	//
	// > When you use this parameter, the Password parameter must be empty. Make sure that the image you use has a password configured.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the resource. Unit: months. This parameter takes effect and is required only when `InstanceChargeType` is set to `PrePaid`. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription billable methods duration. Valid values:
	//
	// <props="china">
	//
	// - Week.
	//
	// - Month (default).
	//
	//
	//
	// <props="intl">Month (default).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address of the instance.
	//
	// When you specify a private IP address for a VPC-connected ECS instance, the IP address must be from the idle CIDR block of the vSwitch (`VSwitchId`).
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance RAM role. You can call the RAM API [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query the instance RAM roles that you have created.
	//
	// example:
	//
	// testRamRoleName
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance, block storage, and elastic network interface controller (NIC) belong.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable security hardening for the operating system. Valid values:
	//
	// -   Active: enables security hardening. This value is applicable only to public images.
	//
	// -   Deactive: does not enable security hardening. This value is applicable to all image types.
	//
	// example:
	//
	// Deactive
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the instance belongs. Instances in the same security group can communicate with each other. A security group can contain up to 1,000 instances.
	//
	// > You cannot specify both `SecurityGroupId` and `SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which the instance belongs. The valid value range of N depends on the maximum number of security groups to which an instance can belong. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > You cannot specify both `SecurityGroupId` and `SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *CreateLaunchTemplateRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain automatic release the instance.
	//
	// - 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain automatic release the instance.
	//
	// Alibaba Cloud sends an ECS system event notification 5 minutes before the instance is released. Spot instances are billed by second. Select an appropriate protection period based on the expected task execution duration.
	//
	// > This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the instance. This parameter supports up to three decimal places and takes effect when `SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy for the pay-as-you-go instance. This parameter takes effect when `InstanceChargeType` is set to `PostPaid`. Valid values:
	//
	// -   NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// -   SpotWithPriceLimit: The instance is a spot instance with a user-defined maximum hourly price.
	//
	// -   SpotAsPriceGo: The instance is a spot instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags for instances, disks, and primary ENIs created from this template version.
	//
	// <details>
	//
	// <summary>Scenarios</summary>
	//
	// After you call CreateLaunchTemplate to create a template, the auto-generated default version uses these tags to tag instances, disks, and primary ENIs when you create instances.
	//
	// </details>
	Tag []*CreateLaunchTemplateRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the resource group to which the launch template belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	TemplateResourceGroupId *string `json:"TemplateResourceGroupId,omitempty" xml:"TemplateResourceGroupId,omitempty"`
	// The tag key-value pairs of the launch template itself.
	//
	// > Currently, you can create and query tags for launch templates only by using API operations. The console does not support creating or viewing these tags.
	TemplateTag []*CreateLaunchTemplateRequestTemplateTag `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Repeated"`
	// The instance user data. The data must be Base64-encoded. The maximum size of the raw data is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID. This parameter is required when you create a VPC-connected instance.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The description of the launch template version. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testVersionDescription
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp12433upq1y5scen****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLaunchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequest) GetSystemDisk() *CreateLaunchTemplateRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateLaunchTemplateRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *CreateLaunchTemplateRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateLaunchTemplateRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateLaunchTemplateRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateLaunchTemplateRequest) GetDataDisk() []*CreateLaunchTemplateRequestDataDisk {
	return s.DataDisk
}

func (s *CreateLaunchTemplateRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateLaunchTemplateRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateLaunchTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequest) GetEnableVmOsConfig() *bool {
	return s.EnableVmOsConfig
}

func (s *CreateLaunchTemplateRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateLaunchTemplateRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateLaunchTemplateRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateLaunchTemplateRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateLaunchTemplateRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateLaunchTemplateRequest) GetImageOptions() *CreateLaunchTemplateRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateLaunchTemplateRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *CreateLaunchTemplateRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateLaunchTemplateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateLaunchTemplateRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateLaunchTemplateRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateLaunchTemplateRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateLaunchTemplateRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateLaunchTemplateRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateLaunchTemplateRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateLaunchTemplateRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *CreateLaunchTemplateRequest) GetNetworkInterface() []*CreateLaunchTemplateRequestNetworkInterface {
	return s.NetworkInterface
}

func (s *CreateLaunchTemplateRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateLaunchTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLaunchTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLaunchTemplateRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateLaunchTemplateRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateLaunchTemplateRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateLaunchTemplateRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateLaunchTemplateRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateLaunchTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLaunchTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLaunchTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLaunchTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLaunchTemplateRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateLaunchTemplateRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateRequest) GetSecurityOptions() *CreateLaunchTemplateRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateLaunchTemplateRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateLaunchTemplateRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateLaunchTemplateRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateLaunchTemplateRequest) GetTag() []*CreateLaunchTemplateRequestTag {
	return s.Tag
}

func (s *CreateLaunchTemplateRequest) GetTemplateResourceGroupId() *string {
	return s.TemplateResourceGroupId
}

func (s *CreateLaunchTemplateRequest) GetTemplateTag() []*CreateLaunchTemplateRequestTemplateTag {
	return s.TemplateTag
}

func (s *CreateLaunchTemplateRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLaunchTemplateRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateLaunchTemplateRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLaunchTemplateRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLaunchTemplateRequest) SetSystemDisk(v *CreateLaunchTemplateRequestSystemDisk) *CreateLaunchTemplateRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetAutoReleaseTime(v string) *CreateLaunchTemplateRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetAutoRenew(v bool) *CreateLaunchTemplateRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetAutoRenewPeriod(v int32) *CreateLaunchTemplateRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetCreditSpecification(v string) *CreateLaunchTemplateRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDataDisk(v []*CreateLaunchTemplateRequestDataDisk) *CreateLaunchTemplateRequest {
	s.DataDisk = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDeletionProtection(v bool) *CreateLaunchTemplateRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDeploymentSetId(v string) *CreateLaunchTemplateRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDescription(v string) *CreateLaunchTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetEnableVmOsConfig(v bool) *CreateLaunchTemplateRequest {
	s.EnableVmOsConfig = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHostName(v string) *CreateLaunchTemplateRequest {
	s.HostName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHttpEndpoint(v string) *CreateLaunchTemplateRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHttpTokens(v string) *CreateLaunchTemplateRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetImageId(v string) *CreateLaunchTemplateRequest {
	s.ImageId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetImageOptions(v *CreateLaunchTemplateRequestImageOptions) *CreateLaunchTemplateRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetImageOwnerAlias(v string) *CreateLaunchTemplateRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInstanceChargeType(v string) *CreateLaunchTemplateRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInstanceName(v string) *CreateLaunchTemplateRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInstanceType(v string) *CreateLaunchTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInternetChargeType(v string) *CreateLaunchTemplateRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetIoOptimized(v string) *CreateLaunchTemplateRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetIpv6AddressCount(v int32) *CreateLaunchTemplateRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetKeyPairName(v string) *CreateLaunchTemplateRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetLaunchTemplateName(v string) *CreateLaunchTemplateRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetNetworkInterface(v []*CreateLaunchTemplateRequestNetworkInterface) *CreateLaunchTemplateRequest {
	s.NetworkInterface = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetNetworkType(v string) *CreateLaunchTemplateRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetOwnerAccount(v string) *CreateLaunchTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetOwnerId(v int64) *CreateLaunchTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPasswordInherit(v bool) *CreateLaunchTemplateRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPeriod(v int32) *CreateLaunchTemplateRequest {
	s.Period = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPeriodUnit(v string) *CreateLaunchTemplateRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPrivateIpAddress(v string) *CreateLaunchTemplateRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetRamRoleName(v string) *CreateLaunchTemplateRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetRegionId(v string) *CreateLaunchTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetResourceGroupId(v string) *CreateLaunchTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetResourceOwnerAccount(v string) *CreateLaunchTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetResourceOwnerId(v int64) *CreateLaunchTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityGroupId(v string) *CreateLaunchTemplateRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityOptions(v *CreateLaunchTemplateRequestSecurityOptions) *CreateLaunchTemplateRequest {
	s.SecurityOptions = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSpotDuration(v int32) *CreateLaunchTemplateRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSpotPriceLimit(v float32) *CreateLaunchTemplateRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSpotStrategy(v string) *CreateLaunchTemplateRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetTag(v []*CreateLaunchTemplateRequestTag) *CreateLaunchTemplateRequest {
	s.Tag = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetTemplateResourceGroupId(v string) *CreateLaunchTemplateRequest {
	s.TemplateResourceGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetTemplateTag(v []*CreateLaunchTemplateRequestTemplateTag) *CreateLaunchTemplateRequest {
	s.TemplateTag = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetUserData(v string) *CreateLaunchTemplateRequest {
	s.UserData = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetVSwitchId(v string) *CreateLaunchTemplateRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetVersionDescription(v string) *CreateLaunchTemplateRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetVpcId(v string) *CreateLaunchTemplateRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetZoneId(v string) *CreateLaunchTemplateRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
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
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterface != nil {
		for _, item := range s.NetworkInterface {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
			return err
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
	if s.TemplateTag != nil {
		for _, item := range s.TemplateTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLaunchTemplateRequestSystemDisk struct {
	// The ID of the automatic snapshot policy applied to the system disk.
	//
	// example:
	//
	// sp-gc7c37d4ylw7mtnk****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: enables the performance burst feature.
	//
	// - false: does not enable the performance burst feature.
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the system disk. Valid values:
	//
	// -   cloud: basic disk.
	//
	// -   cloud_efficiency: ultra disk.
	//
	// -   cloud_ssd: standard SSD.
	//
	// -   cloud_essd: enterprise SSD (ESSD). You can use the `SystemDisk.PerformanceLevel` parameter to set the performance level of the disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// For retired instance types that are not I/O optimized, the default value is cloud. Otherwise, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the system disk when the instance is released. Valid values:
	//
	// - true: releases the system disk when the instance is released.
	//
	// - false: does not release the system disk when the instance is released.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testSystemDiskDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testSystemDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// - true: encrypts the system disk.
	//
	// - false: does not encrypt the system disk.
	//
	// Default value: false.
	//
	// > Zone D in Hong Kong (China) and Zone A in Singapore do not support system disk encryption when you create instances.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	Iops *int32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The KMS key ID of the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD used as the system disk. Valid values:
	//
	// - PL0 (default): a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline Performance}.
	//
	// Baseline Performance = min{1,800 + 50 × Capacity, 50,000}
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the provisioned performance of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
	//
	// example:
	//
	// 50000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// - cloud: 20 to 500.
	//
	// - Other disk categories: 20 to 2048.
	//
	// The value of this parameter must be greater than or equal to max{20, ImageSize}.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateLaunchTemplateRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetIops() *int32 {
	return s.Iops
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetCategory(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateRequestSystemDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetDescription(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetDiskName(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetEncrypted(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetIops(v int32) *CreateLaunchTemplateRequestSystemDisk {
	s.Iops = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetKMSKeyId(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetSize(v int32) *CreateLaunchTemplateRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestDataDisk struct {
	// The ID of the automatic snapshot policy applied to the data disk.
	//
	// example:
	//
	// sp-m5e7fa9ute44ssa****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: enables the performance burst feature.
	//
	// - false: does not enable the performance burst feature.
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of data disk N. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD.
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	//   > The cloud_essd_entry value is supported only when `InstanceType` is set to the `ecs.u1` or `ecs.e` instance family.
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - standard.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - Premium Edition.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// Default value description:
	//
	// - If InstanceType is a retired instance type that is not I/O optimized, the default value of this parameter is `cloud`.
	//
	// - In other cases, the default value is `cloud_efficiency`.<props="china">After January 30, 2026, if the I/O optimized instance type does not support cloud_auto, the default value is cloud_efficiency. Otherwise, the default value is cloud_auto, and the performance burst feature is enabled by default (which incurs additional fees. For details, see [Billing examples](~~368372#p_75k_2hp_7gp~~)). For more information, see [Change notice](https://www.aliyun.com/notice/117844).
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// - true: releases the data disk when the instance is released.
	//
	// - false: does not release the data disk when the instance is released.
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
	// testDataDiskDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. The naming convention varies based on the number of data disks attached:
	//
	// - 1 to 25 data disks: /dev/xvd`[b-z]`
	//
	// - More than 25 data disks: /dev/xvd`[aa-zz]`. For example, the 26th data disk is named /dev/xvdaa, the 27th data disk is named /dev/xvdab, and so on.
	//
	// > This parameter is intended only for full image (whole-machine image) scenarios. You can set this parameter to the mount point of a data disk in the full image and modify the corresponding `DataDisk.N.Size` and `DataDisk.N.Category` parameters to change the disk category and size of the data disk in the full image.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. The name can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testDataDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the data disk.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The KMS key ID for the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD used as a data disk. The value of N must be the same as that in `DataDisk.N.Category=cloud_essd`. Valid values:
	//
	// - PL0: a single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): a single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: a single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: a single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline Performance}.
	//
	// Baseline Performance = min{1,800 + 50 × Capacity, 50,000}
	//
	// > This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the provisioned performance of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
	//
	// example:
	//
	// 50000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// -   cloud: 5 to 2000.
	//
	// -   cloud_efficiency: 20 to 32768.
	//
	// -   cloud_ssd: 20 to 32768.
	//
	// -   cloud_essd: The valid value range depends on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     - PL0: 1 to 32768.
	//
	//     - PL1: 20 to 32768.
	//
	//     - PL2: 461 to 32768.
	//
	//     - PL3: 1261 to 32768.
	//
	// - cloud_auto: 1 to 32,768.
	//
	// - cloud_essd_entry: 10 to 32,768.
	//
	// The value of this parameter must be greater than or equal to the size of the snapshot specified by `SnapshotId`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create data disk N. Valid values of N: 1 to 16. After you specify `DataDisk.N.SnapshotId`, the `DataDisk.N.Size` parameter is ignored. The actual size of the created disk is the size of the specified snapshot.
	//
	// > You cannot use snapshots created on or before July 15, 2013. Such requests are rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateLaunchTemplateRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateRequestDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateLaunchTemplateRequestDataDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateRequestDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetCategory(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDescription(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDevice(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDiskName(v string) *CreateLaunchTemplateRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetEncrypted(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetKMSKeyId(v string) *CreateLaunchTemplateRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetSize(v int32) *CreateLaunchTemplateRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetSnapshotId(v string) *CreateLaunchTemplateRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestImageOptions struct {
	// Indicates whether the instance that uses this image supports logon with the ecs-user account.
	//
	// Valid values:
	//
	// - true: supported.
	//
	// - false: not supported.
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s CreateLaunchTemplateRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateLaunchTemplateRequestImageOptions) SetLoginAsNonRoot(v bool) *CreateLaunchTemplateRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateLaunchTemplateRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestNetworkInterface struct {
	// Specifies whether to retain the ENI when the instance is released. Valid values:
	//
	// - true: does not retain the ENI.
	//
	// - false: retains the ENI.
	//
	// Default value: true.
	//
	// > This parameter takes effect only for secondary ENIs.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the secondary network interface controller (NIC). The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// testEniDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the network interface controller (NIC). Valid values of N: 1 to 2. If you set N to 1, you can configure a primary or secondary NIC. If you set N to 2, you must configure one primary NIC and one secondary NIC.
	//
	// Valid values:
	//
	// - Primary: primary NIC.
	//
	// - Secondary: secondary NIC.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the network interface controller (NIC).
	//
	// Take note of the following items:
	//
	// - Valid values of N: 1 to 2. If you set N to 1, you can configure a primary or secondary NIC. If you set N to 2, you must configure one primary NIC and one secondary NIC.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you do not need to set this parameter.
	//
	// example:
	//
	// testEniName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication mode of the primary ENI. Valid values:
	//
	// - Standard: uses the TCP communication mode.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication mode.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// Adds a network interface controller (NIC) and sets the primary IP address.
	//
	// Take note of the following items:
	//
	// - Valid values of N: 1 to 2.
	//
	//     - If you set N to 1, you can configure a primary or secondary NIC. If the `Amount` parameter is set to a value greater than 1 and you configure a primary NIC with this parameter specified, consecutive primary IP addresses starting from the specified IP address are allocated to multiple ECS instances in batch. In this case, you cannot attach a secondary NIC to the instances.
	//
	//     - If you set N to 2, you must configure one primary NIC and one secondary NIC. If the `Amount` parameter is set to a value greater than 1 and this parameter is specified for the primary NIC, you cannot configure a secondary NIC (that is, you cannot set `NetworkInterface.2.InstanceType=Secondary`).
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, this parameter has the same effect as `PrivateIpAddress`, but you cannot specify both this parameter and `PrivateIpAddress`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter specifies the primary IP address of the secondary NIC. By default, an IP address is randomly selected from the vSwitch CIDR block to which the NIC belongs.
	//
	// > When you create an ECS instance, you can attach up to one secondary NIC. After the instance is created, you can invoke [CreateNetworkInterface](https://help.aliyun.com/document_detail/58504.html) and [AttachNetworkInterface](https://help.aliyun.com/document_detail/58515.html) to attach more secondary NICs.
	//
	// example:
	//
	// ``192.168.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The ID of the security group to which the network interface controller (NIC) belongs.
	//
	// Take note of the following items:
	//
	// - Valid values of N: 1 to 2. If you set N to 1, you can configure a primary or secondary NIC. If you set N to 2, you must configure one primary NIC and one secondary NIC.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you must set this parameter. This parameter has the same effect as `SecurityGroupId`, but you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the security group of the ECS instance.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which the network interface controller (NIC) belongs.
	//
	// - The first N has a valid value range of 1 to 2. If you set N to 1, you can configure a primary or secondary NIC. If you set N to 2, you must configure one primary NIC and one secondary NIC.
	//
	// - The second N specifies one or more security group IDs. The valid value range of N depends on the maximum number of security groups to which an instance can belong. For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// Take note of the following items:
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you must set this parameter or `NetworkInterface.N.SecurityGroupId`. This parameter has the same effect as `SecurityGroupIds.N`, but you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupId`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the security group of the ECS instance.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The vSwitch ID of the network interface controller (NIC).
	//
	// Take note of the following items:
	//
	// - Valid values of N: 1 to 2. If you set N to 1, you can configure a primary or secondary NIC. If you set N to 2, you must configure one primary NIC and one secondary NIC.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, this parameter is required. This parameter has the same effect as `VSwitchId`, but you cannot specify both this parameter and `VSwitchId`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the vSwitch of the ECS instance.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateLaunchTemplateRequestNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestNetworkInterface) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetDeleteOnRelease(v bool) *CreateLaunchTemplateRequestNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetDescription(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetInstanceType(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetNetworkInterfaceName(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetPrimaryIpAddress(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetSecurityGroupId(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateRequestNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetVSwitchId(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestSecurityOptions struct {
	// The trusted system mode. Set the value to vTPM.
	//
	// The following instance families support trusted system mode:
	//
	// - g7, c7, and r7.
	//
	// - Enhanced instance families (g7t, c7t, and r7t).
	//
	// When you create ECS instances of the preceding instance types, you must configure this parameter. Take note of the following items:
	//
	// - To use Alibaba Cloud Trusted System, set this parameter to vTPM. Then, Alibaba Cloud Trusted System performs trusted verification when the instance starts.
	//
	// - If you do not use Alibaba Cloud Trusted System, you do not need to configure this parameter. However, if the ECS instance that you create uses the Enclave-based confidential computing pattern (`SecurityOptions.ConfidentialComputingMode=Enclave`), the trusted system is also enabled for the instance.
	//
	// - When you use OpenAPI to create a trusted ECS instance, you can only invoke `RunInstances`. `CreateInstance` does not support the `SecurityOptions.TrustedSystemMode` parameter.
	//
	// > If you specify the instance as a trusted instance during creation, you can replace the system disk only with an image that supports the trusted system.
	//
	// For more information about the trusted system, see [Overview of trusted features for security-enhanced instances](https://help.aliyun.com/document_detail/201394.html).
	//
	// example:
	//
	// vTPM
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s CreateLaunchTemplateRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *CreateLaunchTemplateRequestSecurityOptions) SetTrustedSystemMode(v string) *CreateLaunchTemplateRequestSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *CreateLaunchTemplateRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestTag struct {
	// The tag key for instances, disks, and primary ENIs created from this template version. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for instances, disks, and primary ENIs created from this template version. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLaunchTemplateRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLaunchTemplateRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLaunchTemplateRequestTag) SetKey(v string) *CreateLaunchTemplateRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLaunchTemplateRequestTag) SetValue(v string) *CreateLaunchTemplateRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLaunchTemplateRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestTemplateTag struct {
	// The tag key of the launch template. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the launch template. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLaunchTemplateRequestTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestTemplateTag) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestTemplateTag) GetKey() *string {
	return s.Key
}

func (s *CreateLaunchTemplateRequestTemplateTag) GetValue() *string {
	return s.Value
}

func (s *CreateLaunchTemplateRequestTemplateTag) SetKey(v string) *CreateLaunchTemplateRequestTemplateTag {
	s.Key = &v
	return s
}

func (s *CreateLaunchTemplateRequestTemplateTag) SetValue(v string) *CreateLaunchTemplateRequestTemplateTag {
	s.Value = &v
	return s
}

func (s *CreateLaunchTemplateRequestTemplateTag) Validate() error {
	return dara.Validate(s)
}
