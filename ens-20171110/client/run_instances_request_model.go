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
	SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest
	GetDataDisk() []*RunInstancesRequestDataDisk
	SetDeletionProtection(v bool) *RunInstancesRequest
	GetDeletionProtection() *bool
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
}

type RunInstancesRequest struct {
	// The number of instances that you want to create. Valid values: 1 to 100.
	//
	// This parameter is required.
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
	DataDisk           []*RunInstancesRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	DeletionProtection *bool                          `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
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
	// This parameter is required.
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
	// This parameter is required.
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
	// This parameter is required.
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
	// This parameter is required.
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
	SystemDisk *RunInstancesRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The tags.
	Tag []*RunInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s RunInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequest) GoString() string {
	return s.String()
}

func (s *RunInstancesRequest) GetAmount() *int64 {
	return s.Amount
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

func (s *RunInstancesRequest) GetDataDisk() []*RunInstancesRequestDataDisk {
	return s.DataDisk
}

func (s *RunInstancesRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
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

func (s *RunInstancesRequest) SetAmount(v int64) *RunInstancesRequest {
	s.Amount = &v
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

func (s *RunInstancesRequest) SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest {
	s.DataDisk = v
	return s
}

func (s *RunInstancesRequest) SetDeletionProtection(v bool) *RunInstancesRequest {
	s.DeletionProtection = &v
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

func (s *RunInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestDataDisk struct {
	// The category of the disk. Valid values:
	//
	// 	- **cloud_efficiency**: ultra disk.
	//
	// 	- **cloud_ssd**: all-flash disk.
	//
	// 	- **local_hdd**: local HDD.
	//
	// 	- **local_ssd**: local SSD.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to encrypt the disk. Valid values:
	//
	// 	- true.
	//
	// 	- false (default).
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used for the disk. Valid values:
	//
	// 	- true.
	//
	// 	- false (default).
	//
	// >  If you set the Encrypted parameter to true, the default service key is used when the KMSKeyId parameter is empty.
	//
	// example:
	//
	// false
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The size of a data disk. Unit: GiB.
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
	// The category of the system disk.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The size of the system disk. Unit: GiB.
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
