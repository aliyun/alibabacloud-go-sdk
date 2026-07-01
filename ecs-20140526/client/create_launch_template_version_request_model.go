// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *CreateLaunchTemplateVersionRequestSystemDisk) *CreateLaunchTemplateVersionRequest
	GetSystemDisk() *CreateLaunchTemplateVersionRequestSystemDisk
	SetAutoReleaseTime(v string) *CreateLaunchTemplateVersionRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *CreateLaunchTemplateVersionRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateLaunchTemplateVersionRequest
	GetAutoRenewPeriod() *int32
	SetCreditSpecification(v string) *CreateLaunchTemplateVersionRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateLaunchTemplateVersionRequestDataDisk) *CreateLaunchTemplateVersionRequest
	GetDataDisk() []*CreateLaunchTemplateVersionRequestDataDisk
	SetDeletionProtection(v bool) *CreateLaunchTemplateVersionRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *CreateLaunchTemplateVersionRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateLaunchTemplateVersionRequest
	GetDescription() *string
	SetEnableVmOsConfig(v bool) *CreateLaunchTemplateVersionRequest
	GetEnableVmOsConfig() *bool
	SetHostName(v string) *CreateLaunchTemplateVersionRequest
	GetHostName() *string
	SetHttpEndpoint(v string) *CreateLaunchTemplateVersionRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateVersionRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateLaunchTemplateVersionRequest
	GetHttpTokens() *string
	SetImageId(v string) *CreateLaunchTemplateVersionRequest
	GetImageId() *string
	SetImageOptions(v *CreateLaunchTemplateVersionRequestImageOptions) *CreateLaunchTemplateVersionRequest
	GetImageOptions() *CreateLaunchTemplateVersionRequestImageOptions
	SetImageOwnerAlias(v string) *CreateLaunchTemplateVersionRequest
	GetImageOwnerAlias() *string
	SetInstanceChargeType(v string) *CreateLaunchTemplateVersionRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateLaunchTemplateVersionRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateLaunchTemplateVersionRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateLaunchTemplateVersionRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateVersionRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateVersionRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateLaunchTemplateVersionRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *CreateLaunchTemplateVersionRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *CreateLaunchTemplateVersionRequest
	GetKeyPairName() *string
	SetLaunchTemplateId(v string) *CreateLaunchTemplateVersionRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *CreateLaunchTemplateVersionRequest
	GetLaunchTemplateName() *string
	SetNetworkInterface(v []*CreateLaunchTemplateVersionRequestNetworkInterface) *CreateLaunchTemplateVersionRequest
	GetNetworkInterface() []*CreateLaunchTemplateVersionRequestNetworkInterface
	SetNetworkType(v string) *CreateLaunchTemplateVersionRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateLaunchTemplateVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLaunchTemplateVersionRequest
	GetOwnerId() *int64
	SetPasswordInherit(v bool) *CreateLaunchTemplateVersionRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateLaunchTemplateVersionRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateLaunchTemplateVersionRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateLaunchTemplateVersionRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateLaunchTemplateVersionRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateLaunchTemplateVersionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLaunchTemplateVersionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLaunchTemplateVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLaunchTemplateVersionRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateVersionRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateLaunchTemplateVersionRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateLaunchTemplateVersionRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *CreateLaunchTemplateVersionRequestSecurityOptions) *CreateLaunchTemplateVersionRequest
	GetSecurityOptions() *CreateLaunchTemplateVersionRequestSecurityOptions
	SetSpotDuration(v int32) *CreateLaunchTemplateVersionRequest
	GetSpotDuration() *int32
	SetSpotPriceLimit(v float32) *CreateLaunchTemplateVersionRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateLaunchTemplateVersionRequest
	GetSpotStrategy() *string
	SetTag(v []*CreateLaunchTemplateVersionRequestTag) *CreateLaunchTemplateVersionRequest
	GetTag() []*CreateLaunchTemplateVersionRequestTag
	SetUserData(v string) *CreateLaunchTemplateVersionRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateLaunchTemplateVersionRequest
	GetVSwitchId() *string
	SetVersionDescription(v string) *CreateLaunchTemplateVersionRequest
	GetVersionDescription() *string
	SetVpcId(v string) *CreateLaunchTemplateVersionRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateLaunchTemplateVersionRequest
	GetZoneId() *string
}

type CreateLaunchTemplateVersionRequest struct {
	SystemDisk *CreateLaunchTemplateVersionRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The automatic release time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// - If the value of seconds (`ss`) is not `00`, the time is automatically rounded to the start of the current minute (`mm`).
	//
	// - The earliest release time is 30 minutes after the current time.
	//
	// - The latest release time cannot be more than three years from the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// - true: enables auto-renewal.
	//
	// - false: does not enable auto-renewal.
	//
	// Default value: false.
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
	// The list of data disks.
	DataDisk []*CreateLaunchTemplateVersionRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The release protection attribute of the instance. Specifies whether the instance can be released from the console or by calling [DeleteInstance](https://help.aliyun.com/document_detail/25507.html). Valid values:
	//
	// - true: enables release protection.
	//
	// - false: disables release protection.
	//
	// Default value: false.
	//
	// > This attribute is applicable only to pay-as-you-go instances. It can only restrict manual release operations and does not take effect on system-initiated release operations.
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
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the operating system configuration of the instance.
	//
	// example:
	//
	// false
	EnableVmOsConfig *bool `json:"EnableVmOsConfig,omitempty" xml:"EnableVmOsConfig,omitempty"`
	// The hostname of the Elastic Compute Service (ECS) server.
	//
	// -   The hostname cannot start or end with a period (.) or hyphen (-), and cannot contain consecutive periods or hyphens.
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
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 3
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// - optional: does not forcefully use the security-hardened mode.
	//
	// - required: forcefully uses the security-hardened mode. After you set this value, the normal mode cannot be used to access instance metadata.
	//
	// Default value: optional.
	//
	// > For more information about the modes for accessing instance metadata, see [Overview of ECS instance metadata](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The ID of the image used to create the instance. You can call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available image resources.
	//
	// example:
	//
	// win2008r2_64_ent_sp1_en-us_40G_alibase_20170915.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image-related property information.
	ImageOptions *CreateLaunchTemplateVersionRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The source of the image.
	//
	// > This parameter will be deprecated. To improve compatibility, use other parameters.
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// <props="china">
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, confirm that your account supports balance payment or credit payment. Otherwise, an `InvalidPayMethod` fault is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	//
	// <props="intl">
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, confirm that your account supports credit payment. Otherwise, an `InvalidPayMethod` fault is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	//
	// <props="partner">
	//
	// - PrePaid: subscription. If you set this parameter to PrePaid, confirm that your account supports credit payment. Otherwise, an `InvalidPayMethod` fault is returned.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length and can contain letters, digits, and characters from the Unicode letter category (which includes characters from various languages). The name can contain colons (:), underscores (_), periods (.), and hyphens (-). The default value is the `InstanceId` of the instance.
	//
	// When you create multiple ECS instances at a time, you can batch configure sequential instance names that contain brackets ([]) and commas (,). For more information, see [Batch configure sequential names or hostnames for multiple instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. For more information, see [Instance family](https://help.aliyun.com/document_detail/25378.html). You can also call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the most recent instance type list.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for outbound Internet bandwidth. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// > In **pay-by-traffic*	- mode, the peak inbound and outbound bandwidths are used as upper limits of bandwidths instead of guaranteed performance specifications. When resource contention occurs, the peak bandwidths may be limited. If you want guaranteed bandwidth, use the **pay-by-bandwidth*	- mode.
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
	// 50
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is an I/O optimized instance. Valid values:
	//
	// -   none: The instance is not I/O optimized.
	//
	// -   optimized: The instance is I/O optimization enabled.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10.
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
	// The ID of the launch template. For more information, call [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html). You must specify `LaunchTemplateId` or `LaunchTemplateName` to determine the launch template.
	//
	// example:
	//
	// lt-m5eiaupmvm2op9d****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testLaunchTemplateName
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The network interface controller (NIC) information.
	NetworkInterface []*CreateLaunchTemplateVersionRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// - vpc: VPC.
	//
	// - classic: classic network. The classic network has been retired. For more information, see [Retirement notice](https://help.aliyun.com/document_detail/2833134.html).
	//
	// example:
	//
	// vpc
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to use the preset password of the image. Valid values:
	//
	// - true
	//
	// - false
	//
	// Default value: false.
	//
	// > When you use this parameter, the Password parameter must be empty. You must also make sure that the image has a preset password.
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
	// The unit of the subscription duration. Valid values:
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
	// The ID of the resource group.
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
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the instance created by using this version belongs. Instances in the same security group can communicate with each other.
	//
	// > You cannot specify both `SecurityGroupId` and `SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which the instance belongs. The valid values of N depend on the maximum number of security groups to which an instance can belong. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > You cannot specify both `SecurityGroupId` and `SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The security options.
	SecurityOptions *CreateLaunchTemplateVersionRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// - 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain automatic release the instance.
	//
	// - 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain automatic release the instance.
	//
	// Alibaba Cloud sends an ECS system event notification 5 minutes before the instance is released. Spot instances are billed by second. We recommend that you select an appropriate protection period based on the expected task execution duration.
	//
	// > This parameter takes effect only when SpotStrategy is set to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the instance. A maximum of three decimal places are supported.
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
	// The tags of the instances, disks, and primary ENIs created by using this version.
	Tag []*CreateLaunchTemplateVersionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Instance user data of the instance. Instance user data must be encoded in Base64. The raw data can be up to 32 KB in size.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBl****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the vSwitch. You must specify this parameter when you create a VPC-connected instance.
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
	// The ID of the virtual private cloud (VPC) to which the instance belongs.
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

func (s CreateLaunchTemplateVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequest) GetSystemDisk() *CreateLaunchTemplateVersionRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateLaunchTemplateVersionRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *CreateLaunchTemplateVersionRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateLaunchTemplateVersionRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateLaunchTemplateVersionRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateLaunchTemplateVersionRequest) GetDataDisk() []*CreateLaunchTemplateVersionRequestDataDisk {
	return s.DataDisk
}

func (s *CreateLaunchTemplateVersionRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateLaunchTemplateVersionRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateLaunchTemplateVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequest) GetEnableVmOsConfig() *bool {
	return s.EnableVmOsConfig
}

func (s *CreateLaunchTemplateVersionRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateLaunchTemplateVersionRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateLaunchTemplateVersionRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateLaunchTemplateVersionRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateLaunchTemplateVersionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateLaunchTemplateVersionRequest) GetImageOptions() *CreateLaunchTemplateVersionRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateLaunchTemplateVersionRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *CreateLaunchTemplateVersionRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateLaunchTemplateVersionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateLaunchTemplateVersionRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateVersionRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateLaunchTemplateVersionRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateLaunchTemplateVersionRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateLaunchTemplateVersionRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateLaunchTemplateVersionRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateLaunchTemplateVersionRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateLaunchTemplateVersionRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateLaunchTemplateVersionRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *CreateLaunchTemplateVersionRequest) GetNetworkInterface() []*CreateLaunchTemplateVersionRequestNetworkInterface {
	return s.NetworkInterface
}

func (s *CreateLaunchTemplateVersionRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateLaunchTemplateVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLaunchTemplateVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLaunchTemplateVersionRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateLaunchTemplateVersionRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateLaunchTemplateVersionRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateLaunchTemplateVersionRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateLaunchTemplateVersionRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateLaunchTemplateVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLaunchTemplateVersionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLaunchTemplateVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLaunchTemplateVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityOptions() *CreateLaunchTemplateVersionRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateLaunchTemplateVersionRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateLaunchTemplateVersionRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateLaunchTemplateVersionRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateLaunchTemplateVersionRequest) GetTag() []*CreateLaunchTemplateVersionRequestTag {
	return s.Tag
}

func (s *CreateLaunchTemplateVersionRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLaunchTemplateVersionRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateVersionRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateLaunchTemplateVersionRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLaunchTemplateVersionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLaunchTemplateVersionRequest) SetSystemDisk(v *CreateLaunchTemplateVersionRequestSystemDisk) *CreateLaunchTemplateVersionRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetAutoReleaseTime(v string) *CreateLaunchTemplateVersionRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetAutoRenew(v bool) *CreateLaunchTemplateVersionRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetAutoRenewPeriod(v int32) *CreateLaunchTemplateVersionRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetCreditSpecification(v string) *CreateLaunchTemplateVersionRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDataDisk(v []*CreateLaunchTemplateVersionRequestDataDisk) *CreateLaunchTemplateVersionRequest {
	s.DataDisk = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDeletionProtection(v bool) *CreateLaunchTemplateVersionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDeploymentSetId(v string) *CreateLaunchTemplateVersionRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDescription(v string) *CreateLaunchTemplateVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetEnableVmOsConfig(v bool) *CreateLaunchTemplateVersionRequest {
	s.EnableVmOsConfig = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHostName(v string) *CreateLaunchTemplateVersionRequest {
	s.HostName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHttpEndpoint(v string) *CreateLaunchTemplateVersionRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateVersionRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHttpTokens(v string) *CreateLaunchTemplateVersionRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetImageId(v string) *CreateLaunchTemplateVersionRequest {
	s.ImageId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetImageOptions(v *CreateLaunchTemplateVersionRequestImageOptions) *CreateLaunchTemplateVersionRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetImageOwnerAlias(v string) *CreateLaunchTemplateVersionRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInstanceChargeType(v string) *CreateLaunchTemplateVersionRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInstanceName(v string) *CreateLaunchTemplateVersionRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInstanceType(v string) *CreateLaunchTemplateVersionRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInternetChargeType(v string) *CreateLaunchTemplateVersionRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateVersionRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateVersionRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetIoOptimized(v string) *CreateLaunchTemplateVersionRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetIpv6AddressCount(v int32) *CreateLaunchTemplateVersionRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetKeyPairName(v string) *CreateLaunchTemplateVersionRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetLaunchTemplateId(v string) *CreateLaunchTemplateVersionRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetLaunchTemplateName(v string) *CreateLaunchTemplateVersionRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetNetworkInterface(v []*CreateLaunchTemplateVersionRequestNetworkInterface) *CreateLaunchTemplateVersionRequest {
	s.NetworkInterface = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetNetworkType(v string) *CreateLaunchTemplateVersionRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetOwnerAccount(v string) *CreateLaunchTemplateVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetOwnerId(v int64) *CreateLaunchTemplateVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPasswordInherit(v bool) *CreateLaunchTemplateVersionRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPeriod(v int32) *CreateLaunchTemplateVersionRequest {
	s.Period = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPeriodUnit(v string) *CreateLaunchTemplateVersionRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPrivateIpAddress(v string) *CreateLaunchTemplateVersionRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetRamRoleName(v string) *CreateLaunchTemplateVersionRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetRegionId(v string) *CreateLaunchTemplateVersionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetResourceGroupId(v string) *CreateLaunchTemplateVersionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetResourceOwnerAccount(v string) *CreateLaunchTemplateVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetResourceOwnerId(v int64) *CreateLaunchTemplateVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateVersionRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityGroupId(v string) *CreateLaunchTemplateVersionRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateVersionRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityOptions(v *CreateLaunchTemplateVersionRequestSecurityOptions) *CreateLaunchTemplateVersionRequest {
	s.SecurityOptions = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSpotDuration(v int32) *CreateLaunchTemplateVersionRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSpotPriceLimit(v float32) *CreateLaunchTemplateVersionRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSpotStrategy(v string) *CreateLaunchTemplateVersionRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetTag(v []*CreateLaunchTemplateVersionRequestTag) *CreateLaunchTemplateVersionRequest {
	s.Tag = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetUserData(v string) *CreateLaunchTemplateVersionRequest {
	s.UserData = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetVSwitchId(v string) *CreateLaunchTemplateVersionRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetVersionDescription(v string) *CreateLaunchTemplateVersionRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetVpcId(v string) *CreateLaunchTemplateVersionRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetZoneId(v string) *CreateLaunchTemplateVersionRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) Validate() error {
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
	return nil
}

type CreateLaunchTemplateVersionRequestSystemDisk struct {
	// The ID of the automatic snapshot policy applied to the system disk.
	//
	// example:
	//
	// sp-bp1dgzpaxwc4load****
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
	// -   cloud_auto: ESSD AutoPL disk.
	//
	// -   cloud_essd: enterprise SSD (ESSD). You can use the `SystemDisk.PerformanceLevel` parameter to configure the performance level of the disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// For retired instance types that are not I/O optimization instances, the default value is cloud. For other instance types, the default value is cloud_efficiency.
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
	// The name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// - true: encrypts the system disk.
	//
	// - false: does not encrypt the system disk.
	//
	// Default value: false.
	//
	// > System disk encryption is not supported in Zone D of the Hong Kong (China) region or Zone A of the Singapore region when you create an instance.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 30000
	Iops *int32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The KMS key ID of the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD used as the system disk. Configure the performance level based on the following valid values:
	//
	// - PL0 (default): A single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk used as the system disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline Performance}.
	//
	// Baseline Performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// > This parameter is available only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the provisioned performance of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
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

func (s CreateLaunchTemplateVersionRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetIops() *int32 {
	return s.Iops
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetCategory(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetDescription(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetDiskName(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetEncrypted(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetIops(v int32) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Iops = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetKMSKeyId(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetSize(v int32) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestDataDisk struct {
	// The ID of the automatic snapshot policy applied to the data disk.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
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
	//   >The cloud_essd_entry value is supported only when `InstanceType` is configured as an instance type in the `ecs.u1` or `ecs.e` family.
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - Standard.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - Premium Edition.
	//
	// For I/O optimization instances, the default value is cloud_efficiency. For non-I/O optimization instances, the default value is cloud.
	//
	// Default value details:
	//
	// - When InstanceType is set to a retired instance type that is not I/O optimized, the default parameter value is `cloud`.
	//
	// - In other cases, the default value is `cloud_efficiency`.<props="china">After January 30, 2026, when the I/O optimized instance type does not support cloud_auto, the default value is cloud_efficiency. Otherwise, the default value is cloud_auto, and the performance burst feature is enabled by default (which incurs additional fees. For details, see [Billing examples](~~368372#p_75k_2hp_7gp~~)). For more information, see [Change notice](https://www.aliyun.com/notice/117844).
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
	// The mount point of the data disk. The naming conventions for mount points vary based on the number of data disks attached:
	//
	// - 1 to 25 data disks: /dev/xvd`[b-z]`
	//
	// - More than 25 data disks: /dev/xvd`[aa-zz]`. For example, the 26th data disk is named /dev/xvdaa, the 27th data disk is named /dev/xvdab, and so on.
	//
	// > This parameter is applicable only to full image (system image) scenarios. You can set this parameter to the mount point of the data disk in the full image and modify the corresponding `DataDisk.N.Size` and `DataDisk.N.Category` parameters to change the disk category and size of the data disk in the full image.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain digits, colons (:), underscores (_), and hyphens (-).
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
	// The performance level of the ESSD used as a data disk. The value of N must be the same as that in `DataDisk.N.Category=cloud_essd`. Configure the performance level based on the following valid values:
	//
	// - PL0: A single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): A single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk used as the system disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline Performance}.
	//
	// Baseline Performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// > This parameter is available only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the provisioned performance of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
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
	// The ID of the snapshot used to create data disk N. Valid values of N: 1 to 16. When `DataDisk.N.SnapshotId` is specified, `DataDisk.N.Size` is ignored. The actual size of the created disk is the size of the specified snapshot.
	//
	// Snapshots created on or before July 15, 2013 cannot be used. Requests that use such snapshots are rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateVersionRequestDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetCategory(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateVersionRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDescription(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDevice(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDiskName(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetEncrypted(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetKMSKeyId(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateVersionRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetSize(v int32) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetSnapshotId(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestImageOptions struct {
	// Specifies whether instances that use this image support logon with the ecs-user user.
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

func (s CreateLaunchTemplateVersionRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateLaunchTemplateVersionRequestImageOptions) SetLoginAsNonRoot(v bool) *CreateLaunchTemplateVersionRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestNetworkInterface struct {
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
	// testNetworkInterfaceDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the ENI. Valid values of N: 1 to 2. If you configure one ENI, you can configure either a primary network interface controller (NIC) or a secondary ENI. If you configure two ENIs, you must configure one primary NIC and one secondary ENI.
	//
	// Valid values:
	//
	// - Primary: primary NIC.
	//
	// - Secondary: secondary ENI.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the secondary network interface controller (NIC). The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// testNetworkInterfaceName
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
	// The primary private IP address of the secondary network interface controller (NIC). The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// ``192.168.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The ID of the security group to which the secondary network interface controller (NIC) belongs. The security group of the secondary NIC must belong to the same VPC as the instance. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// > You cannot specify both `NetworkInterface.N.SecurityGroupId` and `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which the secondary network interface controller (NIC) belongs. The security groups and the secondary NIC must belong to the same VPC. The valid values of N in `SecurityGroupIds.N` depend on the quota for the maximum number of security groups to which a secondary NIC can belong. For more information, see [Limits](https://help.aliyun.com/document_detail/25412.html). The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// > You cannot specify both `NetworkInterface.N.SecurityGroupId` and `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the secondary network interface controller (NIC) belongs. The instance and the secondary NIC must be in the same VPC and the same active zone but can belong to different vSwitches. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestNetworkInterface) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetDeleteOnRelease(v bool) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetDescription(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetInstanceType(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetNetworkInterfaceName(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetPrimaryIpAddress(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetSecurityGroupId(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetVSwitchId(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestSecurityOptions struct {
	// The trusted system mode. Set the value to vTPM.
	//
	// The following instance families support trusted system mode:
	//
	// - g7, c7, and r7.
	//
	// - Security-enhanced instance family (g7t, c7t, and r7t).
	//
	// When you create ECS instances of the preceding instance families, you must configure this parameter. Details:
	//
	// - To use the Alibaba Cloud Trusted System, set this parameter to vTPM. The trusted verification is completed by the Alibaba Cloud Trusted System when the instance starts.
	//
	// - If you do not use the Alibaba Cloud Trusted System, you can leave this parameter empty. However, if the ECS instance that you create uses the Enclave-based confidential computing mode (`SecurityOptions.ConfidentialComputingMode=Enclave`), the trusted system is also enabled for the instance.
	//
	// - When you create a trusted ECS instance by invoking an API operation, you can only invoke `RunInstances`. `CreateInstance` does not support configuring the `SecurityOptions.TrustedSystemMode` parameter.
	//
	// > If you specify an instance as a trusted instance during creation, you can only use images that support the trusted system when you replace the system disk.
	//
	// For more information about the trusted system, see [Overview of trusted features for security-enhanced instances](https://help.aliyun.com/document_detail/201394.html).
	//
	// example:
	//
	// vTPM
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *CreateLaunchTemplateVersionRequestSecurityOptions) SetTrustedSystemMode(v string) *CreateLaunchTemplateVersionRequestSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestTag struct {
	// The tag key of the instances, disks, and primary ENIs created by using this version. Valid values of N: 1 to 20. The tag key cannot be an empty string. It can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instances, disks, and primary ENIs created by using this version. Valid values of N: 1 to 20. The tag value can be an empty string. It can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLaunchTemplateVersionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLaunchTemplateVersionRequestTag) SetKey(v string) *CreateLaunchTemplateVersionRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestTag) SetValue(v string) *CreateLaunchTemplateVersionRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestTag) Validate() error {
	return dara.Validate(s)
}
