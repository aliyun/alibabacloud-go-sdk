// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *RunInstancesShrinkRequest
	GetAmount() *int64
	SetAutoInstallCloudAssistant(v bool) *RunInstancesShrinkRequest
	GetAutoInstallCloudAssistant() *bool
	SetAutoReleaseTime(v string) *RunInstancesShrinkRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *RunInstancesShrinkRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v string) *RunInstancesShrinkRequest
	GetAutoUseCoupon() *string
	SetBillingCycle(v string) *RunInstancesShrinkRequest
	GetBillingCycle() *string
	SetCarrier(v string) *RunInstancesShrinkRequest
	GetCarrier() *string
	SetCloudAssistantRegion(v string) *RunInstancesShrinkRequest
	GetCloudAssistantRegion() *string
	SetDataDiskShrink(v string) *RunInstancesShrinkRequest
	GetDataDiskShrink() *string
	SetDeletionProtection(v bool) *RunInstancesShrinkRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *RunInstancesShrinkRequest
	GetDeploymentSetId() *string
	SetEnsRegionId(v string) *RunInstancesShrinkRequest
	GetEnsRegionId() *string
	SetHostName(v string) *RunInstancesShrinkRequest
	GetHostName() *string
	SetImageId(v string) *RunInstancesShrinkRequest
	GetImageId() *string
	SetInstanceChargeStrategy(v string) *RunInstancesShrinkRequest
	GetInstanceChargeStrategy() *string
	SetInstanceChargeType(v string) *RunInstancesShrinkRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunInstancesShrinkRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunInstancesShrinkRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunInstancesShrinkRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int64) *RunInstancesShrinkRequest
	GetInternetMaxBandwidthOut() *int64
	SetIpType(v string) *RunInstancesShrinkRequest
	GetIpType() *string
	SetIpv6AddressCount(v int64) *RunInstancesShrinkRequest
	GetIpv6AddressCount() *int64
	SetKeyPairName(v string) *RunInstancesShrinkRequest
	GetKeyPairName() *string
	SetLaunchTemplateId(v string) *RunInstancesShrinkRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *RunInstancesShrinkRequest
	GetLaunchTemplateName() *string
	SetLaunchTemplateVersion(v int32) *RunInstancesShrinkRequest
	GetLaunchTemplateVersion() *int32
	SetNetDistrictCode(v string) *RunInstancesShrinkRequest
	GetNetDistrictCode() *string
	SetNetWorkId(v string) *RunInstancesShrinkRequest
	GetNetWorkId() *string
	SetPassword(v string) *RunInstancesShrinkRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunInstancesShrinkRequest
	GetPasswordInherit() *bool
	SetPeriod(v int64) *RunInstancesShrinkRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *RunInstancesShrinkRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *RunInstancesShrinkRequest
	GetPrivateIpAddress() *string
	SetPublicIpIdentification(v bool) *RunInstancesShrinkRequest
	GetPublicIpIdentification() *bool
	SetScheduleAreaLevel(v string) *RunInstancesShrinkRequest
	GetScheduleAreaLevel() *string
	SetSchedulingPriceStrategy(v string) *RunInstancesShrinkRequest
	GetSchedulingPriceStrategy() *string
	SetSchedulingStrategy(v string) *RunInstancesShrinkRequest
	GetSchedulingStrategy() *string
	SetSecurityId(v string) *RunInstancesShrinkRequest
	GetSecurityId() *string
	SetSpotDuration(v int32) *RunInstancesShrinkRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *RunInstancesShrinkRequest
	GetSpotStrategy() *string
	SetSystemDiskShrink(v string) *RunInstancesShrinkRequest
	GetSystemDiskShrink() *string
	SetTag(v []*RunInstancesShrinkRequestTag) *RunInstancesShrinkRequest
	GetTag() []*RunInstancesShrinkRequestTag
	SetUniqueSuffix(v bool) *RunInstancesShrinkRequest
	GetUniqueSuffix() *bool
	SetUserData(v string) *RunInstancesShrinkRequest
	GetUserData() *string
	SetVSwitchId(v string) *RunInstancesShrinkRequest
	GetVSwitchId() *string
	SetVpdId(v string) *RunInstancesShrinkRequest
	GetVpdId() *string
	SetVpdVSwitchIdsShrink(v string) *RunInstancesShrinkRequest
	GetVpdVSwitchIdsShrink() *string
}

type RunInstancesShrinkRequest struct {
	// The number of instances to create. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to automatically install the Cloud Assistant client on the instance.
	AutoInstallCloudAssistant *bool `json:"AutoInstallCloudAssistant,omitempty" xml:"AutoInstallCloudAssistant,omitempty"`
	// The automatic release time of the pay-as-you-go instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// - If the value of seconds (`ss`) is not `00`, the time is automatically set to the start of the current minute (`mm`).
	//
	// - The earliest release time is one hour after the current time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-06-28T14:38:52Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// - **true**: enables auto-renewal.
	//
	// - **false**: does not enable auto-renewal. (Default)
	//
	// >This parameter is invalid for pay-as-you-go instances.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons. Default value: true (AutoUseCoupon=true).
	//
	// example:
	//
	// true
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The billing cycle for instance compute resources. Only instance-level pay-as-you-go is supported. Valid values:
	//
	// - **Day**: daily billing.
	//
	// - **Month**: monthly billing.
	//
	// example:
	//
	// Day
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The carrier.
	//
	// >This parameter is required for region-level scheduling.</br>
	//
	// >For node-level scheduling, if a node has multiple carriers and you do not specify a carrier, instances are created for each carrier on the node. For example, if a node has both China Mobile and China Unicom carriers, instances with two ISPs are created.</br>You can call [DescribeRegionIsps](https://help.aliyun.com/document_detail/2637461.html) to query the carrier information of a node.
	//
	// example:
	//
	// telecom
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The region for Cloud Assistant access. Valid values: cn-zhangjiakou (Zhangjiakou) and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-zhangjiakou
	CloudAssistantRegion *string `json:"CloudAssistantRegion,omitempty" xml:"CloudAssistantRegion,omitempty"`
	// The data cloud disk specifications.
	DataDiskShrink *string `json:"DataDisk,omitempty" xml:"DataDisk,omitempty"`
	// The release protection property of the instance. Specifies whether the instance can be released from the console or by calling an API operation. Valid values:
	//
	// -  true: Enables instance release protection.
	//
	// -  false: Disables instance release protection.
	//
	// Default value: false.
	//
	// > This property applies only to pay-as-you-go instances and can only restrict manual release operations. It does not take effect on system-initiated release operations.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The deployment set ID.
	//
	// example:
	//
	// ds-5s1p6x2l3r4b****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The node ID.
	//
	// > This parameter is required for node-level scheduling and is invalid for region-level scheduling.
	//
	// example:
	//
	// cn-foshan-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The hostname.
	//
	// example:
	//
	// test-HostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image ID. This parameter cannot be specified for ARM-based instances. This parameter is required for other instance types.
	//
	// example:
	//
	// m-5si16wo6simkt267p8b7h****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance billing strategy. Valid values:
	//
	// - **instance**: instance-level billing.
	//
	// - If not specified, the default is user-level billing.
	//
	// example:
	//
	// instance
	InstanceChargeStrategy *string `json:"InstanceChargeStrategy,omitempty" xml:"InstanceChargeStrategy,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PrePaid**: subscription.
	//
	// - **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It cannot start with `http://` or `https://`. It can contain Chinese characters, letters, digits, colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Default value: the InstanceId of the instance.
	//
	// example:
	//
	// TestName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ens.sn1.small
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Billable methods for bandwidth. Valid values:
	//
	// - **BandwidthByDay**: daily peak bandwidth
	//
	// - **95BandwidthByMonth**: monthly 95th percentile bandwidth
	//
	// > You can select a billing method only for the first purchase. Subsequent purchases use the billing method selected during the first purchase by default.
	//
	// example:
	//
	// BandwidthByDay
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum public bandwidth. If the value of this parameter is greater than 0, a public IP address is automatically assigned to the instance.
	//
	// example:
	//
	// 1
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The IP type. Valid values:
	//
	// - **ipv4*	- (default): IPv4.
	//
	// - **ipv6**: IPv6.
	//
	// - **ipv4Andipv6**: IPv4 and IPv6 (single stack).
	//
	// - **ipv4Withipv6**: IPv4 and IPv6 (dual-stack).
	//
	// example:
	//
	// ipv4
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The number of randomly generated IPv6 addresses to assign to the primary ENI of the private network.
	//
	// example:
	//
	// 2
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The key pair name.
	//
	// > Specify at least one of **Password**, **KeyPairName**, and **PasswordInherit**.
	//
	// example:
	//
	// wx2-jumpserver
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The ID of the instance launch template.
	//
	// example:
	//
	// lt-test
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the instance launch template.
	//
	// example:
	//
	// lt-test
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The version of the instance launch template.
	//
	// example:
	//
	// 2
	LaunchTemplateVersion *int32 `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The region code.
	//
	// >This parameter is required for region-level scheduling and is invalid for node-level scheduling.
	//
	// example:
	//
	// 350000
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The network ID.
	//
	// >This parameter can only be used for node-level scheduling. An error is returned if this parameter is specified for region-level scheduling. If you do not specify vSwitch information, only the validity of the network ID is verified. The actual resource creation does not use this network ID.
	//
	// example:
	//
	// net-id
	NetWorkId *string `json:"NetWorkId,omitempty" xml:"NetWorkId,omitempty"`
	// The instance password.
	//
	// > Specify at least one of **Password**, **KeyPairName**, and **PasswordInherit**.
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image. Valid values:
	//
	// - **true**: uses the preset password.
	//
	// - **false**: does not use the preset password.
	//
	// > Specify at least one of **Password**, **KeyPairName**, and **PasswordInherit**.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The duration of the purchased resource.
	//
	// - If **PeriodUnit*	- is set to **Day**, **Period*	- can only be set to **3**.
	//
	// - If **PeriodUnit*	- is set to **Month**, **Period*	- can be set to **1-9*	- or **12**.
	//
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the purchase duration. Valid values:
	//
	// - **Month*	- (default): monthly purchase.
	//
	// - **Day**: daily purchase.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address.
	//
	// >This parameter can only be used for node-level scheduling. An error is returned if this parameter is specified for region-level scheduling.
	//
	// If a private IP address is specified, the number of instances can only be 1. The private IP address takes effect only when both the private IP address and vSwitch ID are specified.
	//
	// example:
	//
	// 10.0.XX.XX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Specifies whether to enable public IP identification. Valid values:
	//
	// - true: enables public IP identification.
	//
	// - false: does not enable public IP identification.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	PublicIpIdentification *bool `json:"PublicIpIdentification,omitempty" xml:"PublicIpIdentification,omitempty"`
	// The scheduling level. Use this parameter to specify node-level or region-level scheduling. Valid values:
	//
	// - **Big**: major area
	//
	// - **Middle**: province
	//
	// - **Small**: city
	//
	// - **Region**: node
	//
	// example:
	//
	// Region
	ScheduleAreaLevel *string `json:"ScheduleAreaLevel,omitempty" xml:"ScheduleAreaLevel,omitempty"`
	// The scheduling price strategy. Valid values:
	//
	// - **PriceHighPriority**: high price first
	//
	// - **PriceLowPriority**: low price first
	//
	// example:
	//
	// PriceHighPriority
	SchedulingPriceStrategy *string `json:"SchedulingPriceStrategy,omitempty" xml:"SchedulingPriceStrategy,omitempty"`
	// The scheduling strategy. Valid values:
	//
	// - **Concentrate**: concentrated
	//
	// - **Disperse**: dispersed
	//
	// > For node-level scheduling, the value is **Concentrate**. For region-level scheduling, select a value based on your requirements.
	//
	// example:
	//
	// Concentrate
	SchedulingStrategy *string `json:"SchedulingStrategy,omitempty" xml:"SchedulingStrategy,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-test
	SecurityId *string `json:"SecurityId,omitempty" xml:"SecurityId,omitempty"`
	// The protection period of the spot instance.
	//
	// Unit: hours.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter takes effect when the `InstanceChargeType` parameter is set to `PostPaid`. Valid values:
	//
	// - NoSpot: a regular pay-as-you-go instance.
	//
	// - SpotAsPriceGo: the system automatically bids at the current market price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// SpotAsPriceGo
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The system cloud disk specifications.
	SystemDiskShrink *string `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The tags that are bound to the instance.
	Tag []*RunInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to add sequential suffixes to **HostName*	- and **InstanceName**. The sequential suffixes start from 001 and cannot exceed 999.
	//
	// example:
	//
	// True
	UniqueSuffix *bool `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
	// The custom user data. Maximum size: 16 KB. You can pass in **UserData*	- information. **UserData*	- is Base64-encoded.
	//
	// example:
	//
	// ZWtest
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID.
	//
	// >This parameter can only be used for node-level scheduling. An error is returned if this parameter is specified for region-level scheduling.
	//
	// example:
	//
	// vsw-5sagnw7m613oulalkd10n****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPD network ID.
	//
	// example:
	//
	// vpd-5s7xxxxxxx
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The list of VPD vSwitch IDs.
	VpdVSwitchIdsShrink *string `json:"VpdVSwitchIds,omitempty" xml:"VpdVSwitchIds,omitempty"`
}

func (s RunInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunInstancesShrinkRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *RunInstancesShrinkRequest) GetAutoInstallCloudAssistant() *bool {
	return s.AutoInstallCloudAssistant
}

func (s *RunInstancesShrinkRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *RunInstancesShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunInstancesShrinkRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *RunInstancesShrinkRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *RunInstancesShrinkRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *RunInstancesShrinkRequest) GetCloudAssistantRegion() *string {
	return s.CloudAssistantRegion
}

func (s *RunInstancesShrinkRequest) GetDataDiskShrink() *string {
	return s.DataDiskShrink
}

func (s *RunInstancesShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunInstancesShrinkRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunInstancesShrinkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *RunInstancesShrinkRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunInstancesShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunInstancesShrinkRequest) GetInstanceChargeStrategy() *string {
	return s.InstanceChargeStrategy
}

func (s *RunInstancesShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunInstancesShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunInstancesShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunInstancesShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunInstancesShrinkRequest) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *RunInstancesShrinkRequest) GetIpType() *string {
	return s.IpType
}

func (s *RunInstancesShrinkRequest) GetIpv6AddressCount() *int64 {
	return s.Ipv6AddressCount
}

func (s *RunInstancesShrinkRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunInstancesShrinkRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *RunInstancesShrinkRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *RunInstancesShrinkRequest) GetLaunchTemplateVersion() *int32 {
	return s.LaunchTemplateVersion
}

func (s *RunInstancesShrinkRequest) GetNetDistrictCode() *string {
	return s.NetDistrictCode
}

func (s *RunInstancesShrinkRequest) GetNetWorkId() *string {
	return s.NetWorkId
}

func (s *RunInstancesShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *RunInstancesShrinkRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunInstancesShrinkRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *RunInstancesShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunInstancesShrinkRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RunInstancesShrinkRequest) GetPublicIpIdentification() *bool {
	return s.PublicIpIdentification
}

func (s *RunInstancesShrinkRequest) GetScheduleAreaLevel() *string {
	return s.ScheduleAreaLevel
}

func (s *RunInstancesShrinkRequest) GetSchedulingPriceStrategy() *string {
	return s.SchedulingPriceStrategy
}

func (s *RunInstancesShrinkRequest) GetSchedulingStrategy() *string {
	return s.SchedulingStrategy
}

func (s *RunInstancesShrinkRequest) GetSecurityId() *string {
	return s.SecurityId
}

func (s *RunInstancesShrinkRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *RunInstancesShrinkRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunInstancesShrinkRequest) GetSystemDiskShrink() *string {
	return s.SystemDiskShrink
}

func (s *RunInstancesShrinkRequest) GetTag() []*RunInstancesShrinkRequestTag {
	return s.Tag
}

func (s *RunInstancesShrinkRequest) GetUniqueSuffix() *bool {
	return s.UniqueSuffix
}

func (s *RunInstancesShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunInstancesShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunInstancesShrinkRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *RunInstancesShrinkRequest) GetVpdVSwitchIdsShrink() *string {
	return s.VpdVSwitchIdsShrink
}

func (s *RunInstancesShrinkRequest) SetAmount(v int64) *RunInstancesShrinkRequest {
	s.Amount = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetAutoInstallCloudAssistant(v bool) *RunInstancesShrinkRequest {
	s.AutoInstallCloudAssistant = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetAutoReleaseTime(v string) *RunInstancesShrinkRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetAutoRenew(v bool) *RunInstancesShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetAutoUseCoupon(v string) *RunInstancesShrinkRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetBillingCycle(v string) *RunInstancesShrinkRequest {
	s.BillingCycle = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetCarrier(v string) *RunInstancesShrinkRequest {
	s.Carrier = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetCloudAssistantRegion(v string) *RunInstancesShrinkRequest {
	s.CloudAssistantRegion = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetDataDiskShrink(v string) *RunInstancesShrinkRequest {
	s.DataDiskShrink = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetDeletionProtection(v bool) *RunInstancesShrinkRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetDeploymentSetId(v string) *RunInstancesShrinkRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetEnsRegionId(v string) *RunInstancesShrinkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetHostName(v string) *RunInstancesShrinkRequest {
	s.HostName = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetImageId(v string) *RunInstancesShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetInstanceChargeStrategy(v string) *RunInstancesShrinkRequest {
	s.InstanceChargeStrategy = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetInstanceChargeType(v string) *RunInstancesShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetInstanceName(v string) *RunInstancesShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetInstanceType(v string) *RunInstancesShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetInternetChargeType(v string) *RunInstancesShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetInternetMaxBandwidthOut(v int64) *RunInstancesShrinkRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetIpType(v string) *RunInstancesShrinkRequest {
	s.IpType = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetIpv6AddressCount(v int64) *RunInstancesShrinkRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetKeyPairName(v string) *RunInstancesShrinkRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetLaunchTemplateId(v string) *RunInstancesShrinkRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetLaunchTemplateName(v string) *RunInstancesShrinkRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetLaunchTemplateVersion(v int32) *RunInstancesShrinkRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetNetDistrictCode(v string) *RunInstancesShrinkRequest {
	s.NetDistrictCode = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetNetWorkId(v string) *RunInstancesShrinkRequest {
	s.NetWorkId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetPassword(v string) *RunInstancesShrinkRequest {
	s.Password = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetPasswordInherit(v bool) *RunInstancesShrinkRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetPeriod(v int64) *RunInstancesShrinkRequest {
	s.Period = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetPeriodUnit(v string) *RunInstancesShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetPrivateIpAddress(v string) *RunInstancesShrinkRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetPublicIpIdentification(v bool) *RunInstancesShrinkRequest {
	s.PublicIpIdentification = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetScheduleAreaLevel(v string) *RunInstancesShrinkRequest {
	s.ScheduleAreaLevel = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetSchedulingPriceStrategy(v string) *RunInstancesShrinkRequest {
	s.SchedulingPriceStrategy = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetSchedulingStrategy(v string) *RunInstancesShrinkRequest {
	s.SchedulingStrategy = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetSecurityId(v string) *RunInstancesShrinkRequest {
	s.SecurityId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetSpotDuration(v int32) *RunInstancesShrinkRequest {
	s.SpotDuration = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetSpotStrategy(v string) *RunInstancesShrinkRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetSystemDiskShrink(v string) *RunInstancesShrinkRequest {
	s.SystemDiskShrink = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetTag(v []*RunInstancesShrinkRequestTag) *RunInstancesShrinkRequest {
	s.Tag = v
	return s
}

func (s *RunInstancesShrinkRequest) SetUniqueSuffix(v bool) *RunInstancesShrinkRequest {
	s.UniqueSuffix = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetUserData(v string) *RunInstancesShrinkRequest {
	s.UserData = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetVSwitchId(v string) *RunInstancesShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetVpdId(v string) *RunInstancesShrinkRequest {
	s.VpdId = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetVpdVSwitchIdsShrink(v string) *RunInstancesShrinkRequest {
	s.VpdVSwitchIdsShrink = &v
	return s
}

func (s *RunInstancesShrinkRequest) Validate() error {
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

type RunInstancesShrinkRequestTag struct {
	// The custom tag key.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value cannot exceed 128 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunInstancesShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *RunInstancesShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunInstancesShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunInstancesShrinkRequestTag) SetKey(v string) *RunInstancesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *RunInstancesShrinkRequestTag) SetValue(v string) *RunInstancesShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *RunInstancesShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
