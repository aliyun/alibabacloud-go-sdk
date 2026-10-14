// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *RunInstancesRequest
	GetAmount() *int64
	SetAutoInstallCloudAssistant(v bool) *RunInstancesRequest
	GetAutoInstallCloudAssistant() *bool
	SetAutoReleaseTime(v string) *RunInstancesRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *RunInstancesRequest
	GetAutoRenew() *bool
	SetAutoUseCoupon(v string) *RunInstancesRequest
	GetAutoUseCoupon() *string
	SetBillingCycle(v string) *RunInstancesRequest
	GetBillingCycle() *string
	SetCarrier(v string) *RunInstancesRequest
	GetCarrier() *string
	SetCloudAssistantRegion(v string) *RunInstancesRequest
	GetCloudAssistantRegion() *string
	SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest
	GetDataDisk() []*RunInstancesRequestDataDisk
	SetDeletionProtection(v bool) *RunInstancesRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *RunInstancesRequest
	GetDeploymentSetId() *string
	SetEnsRegionId(v string) *RunInstancesRequest
	GetEnsRegionId() *string
	SetHostName(v string) *RunInstancesRequest
	GetHostName() *string
	SetImageId(v string) *RunInstancesRequest
	GetImageId() *string
	SetInstanceChargeStrategy(v string) *RunInstancesRequest
	GetInstanceChargeStrategy() *string
	SetInstanceChargeType(v string) *RunInstancesRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunInstancesRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunInstancesRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunInstancesRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthOut(v int64) *RunInstancesRequest
	GetInternetMaxBandwidthOut() *int64
	SetIpType(v string) *RunInstancesRequest
	GetIpType() *string
	SetIpv6AddressCount(v int64) *RunInstancesRequest
	GetIpv6AddressCount() *int64
	SetKeyPairName(v string) *RunInstancesRequest
	GetKeyPairName() *string
	SetLaunchTemplateId(v string) *RunInstancesRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *RunInstancesRequest
	GetLaunchTemplateName() *string
	SetLaunchTemplateVersion(v int32) *RunInstancesRequest
	GetLaunchTemplateVersion() *int32
	SetNetDistrictCode(v string) *RunInstancesRequest
	GetNetDistrictCode() *string
	SetNetWorkId(v string) *RunInstancesRequest
	GetNetWorkId() *string
	SetPassword(v string) *RunInstancesRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunInstancesRequest
	GetPasswordInherit() *bool
	SetPeriod(v int64) *RunInstancesRequest
	GetPeriod() *int64
	SetPeriodUnit(v string) *RunInstancesRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *RunInstancesRequest
	GetPrivateIpAddress() *string
	SetPublicIpIdentification(v bool) *RunInstancesRequest
	GetPublicIpIdentification() *bool
	SetScheduleAreaLevel(v string) *RunInstancesRequest
	GetScheduleAreaLevel() *string
	SetSchedulingPriceStrategy(v string) *RunInstancesRequest
	GetSchedulingPriceStrategy() *string
	SetSchedulingStrategy(v string) *RunInstancesRequest
	GetSchedulingStrategy() *string
	SetSecurityId(v string) *RunInstancesRequest
	GetSecurityId() *string
	SetSpotDuration(v int32) *RunInstancesRequest
	GetSpotDuration() *int32
	SetSpotStrategy(v string) *RunInstancesRequest
	GetSpotStrategy() *string
	SetSystemDisk(v *RunInstancesRequestSystemDisk) *RunInstancesRequest
	GetSystemDisk() *RunInstancesRequestSystemDisk
	SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest
	GetTag() []*RunInstancesRequestTag
	SetUniqueSuffix(v bool) *RunInstancesRequest
	GetUniqueSuffix() *bool
	SetUserData(v string) *RunInstancesRequest
	GetUserData() *string
	SetVSwitchId(v string) *RunInstancesRequest
	GetVSwitchId() *string
	SetVpdId(v string) *RunInstancesRequest
	GetVpdId() *string
	SetVpdVSwitchIds(v []*string) *RunInstancesRequest
	GetVpdVSwitchIds() []*string
}

type RunInstancesRequest struct {
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
	DataDisk []*RunInstancesRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
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
	SystemDisk *RunInstancesRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The tags that are bound to the instance.
	Tag []*RunInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	VpdVSwitchIds []*string `json:"VpdVSwitchIds,omitempty" xml:"VpdVSwitchIds,omitempty" type:"Repeated"`
}

func (s RunInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequest) GoString() string {
	return s.String()
}

func (s *RunInstancesRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *RunInstancesRequest) GetAutoInstallCloudAssistant() *bool {
	return s.AutoInstallCloudAssistant
}

func (s *RunInstancesRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *RunInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunInstancesRequest) GetAutoUseCoupon() *string {
	return s.AutoUseCoupon
}

func (s *RunInstancesRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *RunInstancesRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *RunInstancesRequest) GetCloudAssistantRegion() *string {
	return s.CloudAssistantRegion
}

func (s *RunInstancesRequest) GetDataDisk() []*RunInstancesRequestDataDisk {
	return s.DataDisk
}

func (s *RunInstancesRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunInstancesRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunInstancesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *RunInstancesRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunInstancesRequest) GetInstanceChargeStrategy() *string {
	return s.InstanceChargeStrategy
}

func (s *RunInstancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunInstancesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunInstancesRequest) GetInternetMaxBandwidthOut() *int64 {
	return s.InternetMaxBandwidthOut
}

func (s *RunInstancesRequest) GetIpType() *string {
	return s.IpType
}

func (s *RunInstancesRequest) GetIpv6AddressCount() *int64 {
	return s.Ipv6AddressCount
}

func (s *RunInstancesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunInstancesRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *RunInstancesRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *RunInstancesRequest) GetLaunchTemplateVersion() *int32 {
	return s.LaunchTemplateVersion
}

func (s *RunInstancesRequest) GetNetDistrictCode() *string {
	return s.NetDistrictCode
}

func (s *RunInstancesRequest) GetNetWorkId() *string {
	return s.NetWorkId
}

func (s *RunInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *RunInstancesRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunInstancesRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *RunInstancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunInstancesRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RunInstancesRequest) GetPublicIpIdentification() *bool {
	return s.PublicIpIdentification
}

func (s *RunInstancesRequest) GetScheduleAreaLevel() *string {
	return s.ScheduleAreaLevel
}

func (s *RunInstancesRequest) GetSchedulingPriceStrategy() *string {
	return s.SchedulingPriceStrategy
}

func (s *RunInstancesRequest) GetSchedulingStrategy() *string {
	return s.SchedulingStrategy
}

func (s *RunInstancesRequest) GetSecurityId() *string {
	return s.SecurityId
}

func (s *RunInstancesRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *RunInstancesRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunInstancesRequest) GetSystemDisk() *RunInstancesRequestSystemDisk {
	return s.SystemDisk
}

func (s *RunInstancesRequest) GetTag() []*RunInstancesRequestTag {
	return s.Tag
}

func (s *RunInstancesRequest) GetUniqueSuffix() *bool {
	return s.UniqueSuffix
}

func (s *RunInstancesRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunInstancesRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *RunInstancesRequest) GetVpdVSwitchIds() []*string {
	return s.VpdVSwitchIds
}

func (s *RunInstancesRequest) SetAmount(v int64) *RunInstancesRequest {
	s.Amount = &v
	return s
}

func (s *RunInstancesRequest) SetAutoInstallCloudAssistant(v bool) *RunInstancesRequest {
	s.AutoInstallCloudAssistant = &v
	return s
}

func (s *RunInstancesRequest) SetAutoReleaseTime(v string) *RunInstancesRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *RunInstancesRequest) SetAutoRenew(v bool) *RunInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunInstancesRequest) SetAutoUseCoupon(v string) *RunInstancesRequest {
	s.AutoUseCoupon = &v
	return s
}

func (s *RunInstancesRequest) SetBillingCycle(v string) *RunInstancesRequest {
	s.BillingCycle = &v
	return s
}

func (s *RunInstancesRequest) SetCarrier(v string) *RunInstancesRequest {
	s.Carrier = &v
	return s
}

func (s *RunInstancesRequest) SetCloudAssistantRegion(v string) *RunInstancesRequest {
	s.CloudAssistantRegion = &v
	return s
}

func (s *RunInstancesRequest) SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest {
	s.DataDisk = v
	return s
}

func (s *RunInstancesRequest) SetDeletionProtection(v bool) *RunInstancesRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunInstancesRequest) SetDeploymentSetId(v string) *RunInstancesRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunInstancesRequest) SetEnsRegionId(v string) *RunInstancesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *RunInstancesRequest) SetHostName(v string) *RunInstancesRequest {
	s.HostName = &v
	return s
}

func (s *RunInstancesRequest) SetImageId(v string) *RunInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceChargeStrategy(v string) *RunInstancesRequest {
	s.InstanceChargeStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceChargeType(v string) *RunInstancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceName(v string) *RunInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceType(v string) *RunInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesRequest) SetInternetChargeType(v string) *RunInstancesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetInternetMaxBandwidthOut(v int64) *RunInstancesRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunInstancesRequest) SetIpType(v string) *RunInstancesRequest {
	s.IpType = &v
	return s
}

func (s *RunInstancesRequest) SetIpv6AddressCount(v int64) *RunInstancesRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *RunInstancesRequest) SetKeyPairName(v string) *RunInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateId(v string) *RunInstancesRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateName(v string) *RunInstancesRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateVersion(v int32) *RunInstancesRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *RunInstancesRequest) SetNetDistrictCode(v string) *RunInstancesRequest {
	s.NetDistrictCode = &v
	return s
}

func (s *RunInstancesRequest) SetNetWorkId(v string) *RunInstancesRequest {
	s.NetWorkId = &v
	return s
}

func (s *RunInstancesRequest) SetPassword(v string) *RunInstancesRequest {
	s.Password = &v
	return s
}

func (s *RunInstancesRequest) SetPasswordInherit(v bool) *RunInstancesRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunInstancesRequest) SetPeriod(v int64) *RunInstancesRequest {
	s.Period = &v
	return s
}

func (s *RunInstancesRequest) SetPeriodUnit(v string) *RunInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunInstancesRequest) SetPrivateIpAddress(v string) *RunInstancesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunInstancesRequest) SetPublicIpIdentification(v bool) *RunInstancesRequest {
	s.PublicIpIdentification = &v
	return s
}

func (s *RunInstancesRequest) SetScheduleAreaLevel(v string) *RunInstancesRequest {
	s.ScheduleAreaLevel = &v
	return s
}

func (s *RunInstancesRequest) SetSchedulingPriceStrategy(v string) *RunInstancesRequest {
	s.SchedulingPriceStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetSchedulingStrategy(v string) *RunInstancesRequest {
	s.SchedulingStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityId(v string) *RunInstancesRequest {
	s.SecurityId = &v
	return s
}

func (s *RunInstancesRequest) SetSpotDuration(v int32) *RunInstancesRequest {
	s.SpotDuration = &v
	return s
}

func (s *RunInstancesRequest) SetSpotStrategy(v string) *RunInstancesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetSystemDisk(v *RunInstancesRequestSystemDisk) *RunInstancesRequest {
	s.SystemDisk = v
	return s
}

func (s *RunInstancesRequest) SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest {
	s.Tag = v
	return s
}

func (s *RunInstancesRequest) SetUniqueSuffix(v bool) *RunInstancesRequest {
	s.UniqueSuffix = &v
	return s
}

func (s *RunInstancesRequest) SetUserData(v string) *RunInstancesRequest {
	s.UserData = &v
	return s
}

func (s *RunInstancesRequest) SetVSwitchId(v string) *RunInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunInstancesRequest) SetVpdId(v string) *RunInstancesRequest {
	s.VpdId = &v
	return s
}

func (s *RunInstancesRequest) SetVpdVSwitchIds(v []*string) *RunInstancesRequest {
	s.VpdVSwitchIds = v
	return s
}

func (s *RunInstancesRequest) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
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

type RunInstancesRequestDataDisk struct {
	// The disk category. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk
	//
	// - **cloud_ssd**: all-flash cloud disk
	//
	// - **local_hdd**: local HDD
	//
	// - **local_ssd**: local SSD
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to encrypt the cloud disk. Valid values:
	//
	// - true: encrypts the cloud disk.
	//
	// - false (default): does not encrypt the cloud disk.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key used by the cloud disk. Valid values:
	//
	// - true: encrypts the cloud disk.
	//
	// - false (default): does not encrypt the cloud disk.
	//
	// > If Encrypted is set to true and KMSKeyId is empty, the default service key is used.
	//
	// example:
	//
	// false
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The size of the data cloud disk. Unit: GiB.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s RunInstancesRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestDataDisk) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *RunInstancesRequestDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *RunInstancesRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *RunInstancesRequestDataDisk) GetSize() *int64 {
	return s.Size
}

func (s *RunInstancesRequestDataDisk) SetCategory(v string) *RunInstancesRequestDataDisk {
	s.Category = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetEncrypted(v bool) *RunInstancesRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetKMSKeyId(v string) *RunInstancesRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetSize(v int64) *RunInstancesRequestDataDisk {
	s.Size = &v
	return s
}

func (s *RunInstancesRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSystemDisk struct {
	// The type of the system cloud disk.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The size of the system cloud disk. Unit: GiB.
	//
	// example:
	//
	// 50
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s RunInstancesRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *RunInstancesRequestSystemDisk) GetSize() *int64 {
	return s.Size
}

func (s *RunInstancesRequestSystemDisk) SetCategory(v string) *RunInstancesRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetSize(v int64) *RunInstancesRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestTag struct {
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

func (s RunInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunInstancesRequestTag) SetKey(v string) *RunInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *RunInstancesRequestTag) SetValue(v string) *RunInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *RunInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
