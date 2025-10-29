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
	SetDataDiskShrink(v string) *RunInstancesShrinkRequest
	GetDataDiskShrink() *string
	SetDeletionProtection(v bool) *RunInstancesShrinkRequest
	GetDeletionProtection() *bool
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
}

type RunInstancesShrinkRequest struct {
	// The number of instances that you want to create. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The time when to automatically release the pay-as-you-go instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in Coordinated Universal Time (UTC).
	//
	// 	- If the value of `ss` is not `00`, the start time is automatically rounded down to the nearest minute based on the value of `mm`.
	//
	// 	- The specified time must be at least one hour later than the current time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-06-28T14:38:52Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal for the premium bandwidth plan. Valid values:
	//
	// 	- **true**.
	//
	// 	- **false*	- (default).
	//
	// >  This parameter is not available when InstanceChargeType is set to PostPaid.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Specifies whether to use coupons. Default value: true.
	//
	// example:
	//
	// true
	AutoUseCoupon *string `json:"AutoUseCoupon,omitempty" xml:"AutoUseCoupon,omitempty"`
	// The billing cycle of computing resources of the instance. Only pay-as-you-go instances are supported. Valid values:
	//
	// 	- **Day**.
	//
	// 	- **Month**.
	//
	// example:
	//
	// Day
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The Internet service provider (ISP).
	//
	// >  This parameter required if ScheduleAreaLevel is set to Region.\\
	//
	// If you set ScheduleAreaLevel to Region, a node has multiple ISPs, and you do not specify an ISP, then the create instance uses the ISP of the node. If the node has two ISPs, such as China Mobile and China Unicom, the created instance has two ISPs.\\
	//
	// You can call the DescribeRegionIsps operation to query ISPs of the edge node.[](~~2637461~~)
	//
	// example:
	//
	// telecom
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// The specifications of data disks.
	DataDiskShrink     *string `json:"DataDisk,omitempty" xml:"DataDisk,omitempty"`
	DeletionProtection *bool   `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the node.
	//
	// >  This parameter is required if ScheduleAreaLevel is set to Region and is not available if ScheduleAreaLevel is set to other values.
	//
	// example:
	//
	// cn-foshan-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the host.
	//
	// example:
	//
	// test-HostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the image. For ARM PCB-based server instances, leave this parameter empty. For other instances, this parameter is required.
	//
	// example:
	//
	// m-5si16wo6simkt267p8b7hcmy3
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The billing policy of the instance. Valid values:
	//
	// 	- **instance**: Bills are generated based on instances.
	//
	// 	- If you do not specify this parameter, bills are generated based on users.
	//
	// example:
	//
	// instance
	InstanceChargeStrategy *string `json:"InstanceChargeStrategy,omitempty" xml:"InstanceChargeStrategy,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid:*	- pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// The default value of this parameter is the value of the InstanceId parameter.
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
	// The bandwidth billing method. Valid values:
	//
	// 	- **BandwidthByDay**: pay by daily peak bandwidth
	//
	// 	- **95BandwidthByMonth**: pay by monthly 95th percentile bandwidth
	//
	// >  This parameter is required if you purchase an ENS instance for the first time. The value that you specified is used as the default value for subsequent purchases.
	//
	// example:
	//
	// BandwidthByDay
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum public bandwidth. If the value of this parameter is greater than 0, a public IP address is assigned to the instance.
	//
	// example:
	//
	// 1
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The type of the IP address. Valid values:
	//
	// 	- **ipv4*	- (default).
	//
	// 	- **ipv6**.
	//
	// 	- **ipv4Andipv6*	- (single stack).
	//
	// 	- **ipv4Withipv6*	- (dual stack).
	//
	// example:
	//
	// ipv4
	IpType           *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	Ipv6AddressCount *int64  `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair.
	//
	// >  You need to specify at least one of **Password**, **KeyPairName**, and **PasswordInherit**.
	//
	// example:
	//
	// wx2-jumpserver
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// example:
	//
	// lt-test
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// example:
	//
	// lt-test
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// example:
	//
	// 2
	LaunchTemplateVersion *int32 `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The code of the region.
	//
	// >  This parameter is not available if ScheduleAreaLevel is set to Region and is required if ScheduleAreaLevel is set to other values.
	//
	// example:
	//
	// 350000
	NetDistrictCode *string `json:"NetDistrictCode,omitempty" xml:"NetDistrictCode,omitempty"`
	// The ID of the network.
	//
	// >  This parameter is available only if ScheduleAreaLevel is set to Region and cannot be configured if ScheduleAreaLevel is set to other values. Otherwise, an error occurs.
	//
	// example:
	//
	// net-id
	NetWorkId *string `json:"NetWorkId,omitempty" xml:"NetWorkId,omitempty"`
	// The password that is used to connect to the instance.
	//
	// >  You need to specify at least one of **Password**, **KeyPairName**, and **PasswordInherit**.
	//
	// example:
	//
	// testPassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  You need to specify at least one of **Password**, **KeyPairName**, and **PasswordInherit**.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The unit of the subscription period.
	//
	// 	- If **PeriodUnit*	- is set to **Day**, **Period*	- can only be set to **3**.
	//
	// 	- If **PeriodUnit*	- is **Month**, **Period*	- can be set to **1 to 9*	- or **12**.
	//
	// example:
	//
	// 1-9,12
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// 	- **Month*	- (default).
	//
	// 	- **Day**.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address.
	//
	// >  This parameter is available only if ScheduleAreaLevel is set to Region and cannot be configured if ScheduleAreaLevel is set to other values. Otherwise, an error occurs. If you specify a private IP address, the number of instances must be 1. The private IP address takes effect only when the private IP address and the vSwitch ID are not empty.
	//
	// example:
	//
	// 10.0.0.120
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// Specifies whether to enable public IP address identification. Valid values: true and false. Default value: false.
	//
	// example:
	//
	// true
	PublicIpIdentification *bool `json:"PublicIpIdentification,omitempty" xml:"PublicIpIdentification,omitempty"`
	// The scheduling level. This parameter specifies area-level scheduling or node-level scheduling. Valid values:
	//
	// 	- **Big**: greater area
	//
	// 	- **Middle**: province
	//
	// 	- **Small**: city
	//
	// 	- **Region**: node
	//
	// example:
	//
	// Region
	ScheduleAreaLevel *string `json:"ScheduleAreaLevel,omitempty" xml:"ScheduleAreaLevel,omitempty"`
	// The scheduling price policy. Valid values:
	//
	// 	- **PriceHighPriority**: The high price prevails.
	//
	// 	- **PriceLowPriority**: The low price prevails.
	//
	// example:
	//
	// PriceHighPriority
	SchedulingPriceStrategy *string `json:"SchedulingPriceStrategy,omitempty" xml:"SchedulingPriceStrategy,omitempty"`
	// The scheduling policy of the taint. Valid values:
	//
	// 	- **Concentrate**
	//
	// 	- **Disperse**
	//
	// >  If ScheduleAreaLevel is set to Region, set this parameter to **Concentrate**. If ScheduleAreaLevel is set to other values, set this parameter to Concentrate or Disperse based on your business requirements.
	//
	// example:
	//
	// concentrate
	SchedulingStrategy *string `json:"SchedulingStrategy,omitempty" xml:"SchedulingStrategy,omitempty"`
	// The ID of security group.
	//
	// example:
	//
	// sg-test
	SecurityId *string `json:"SecurityId,omitempty" xml:"SecurityId,omitempty"`
	// The protection period of the preemptible instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a preemptible instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a preemptible instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you 5 minutes before the instance is released. Preemptible instances are billed by second. We recommend that you specify an appropriate protection period based on your business requirements.
	//
	// example:
	//
	// 2
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter is valid only when the `InstanceChargeType` parameter is set to `PostPaid`. Valid values:
	//
	// 	- NoSpot: The elastic container instances are pay-as-you-go instances.
	//
	// 	- SpotAsPriceGo: The instance is a preemptible instance for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// SpotAsPriceGo
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The specification of the system disk.
	SystemDiskShrink *string `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty"`
	// The tags.
	Tag []*RunInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to append sequential suffixes to the hostname specified by the **HostName*	- parameter and to the instance name specified by the **InstanceName*	- parameter. The sequential suffixes range from 001 to 999.
	//
	// example:
	//
	// True
	UniqueSuffix *bool `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
	// The custom data. The maximum data size is 16 KB. You can specify **UserData**. **UserData*	- must be Base64-encoded.
	//
	// example:
	//
	// ZWtest
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the vSwitch.
	//
	// >  This parameter is available only if ScheduleAreaLevel is set to Region and cannot be configured if ScheduleAreaLevel is set to other values. Otherwise, an error occurs.
	//
	// example:
	//
	// vsw-5sagnw7m613oulalkd10nv0ob
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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

func (s *RunInstancesShrinkRequest) GetDataDiskShrink() *string {
	return s.DataDiskShrink
}

func (s *RunInstancesShrinkRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
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

func (s *RunInstancesShrinkRequest) SetAmount(v int64) *RunInstancesShrinkRequest {
	s.Amount = &v
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

func (s *RunInstancesShrinkRequest) SetDataDiskShrink(v string) *RunInstancesShrinkRequest {
	s.DataDiskShrink = &v
	return s
}

func (s *RunInstancesShrinkRequest) SetDeletionProtection(v bool) *RunInstancesShrinkRequest {
	s.DeletionProtection = &v
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
	// The key of the tag.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
