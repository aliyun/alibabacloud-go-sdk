// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *DescribePriceRequest
	GetAmount() *int32
	SetBandwidth(v int32) *DescribePriceRequest
	GetBandwidth() *int32
	SetDuration(v int32) *DescribePriceRequest
	GetDuration() *int32
	SetGroupDesktopCount(v int32) *DescribePriceRequest
	GetGroupDesktopCount() *int32
	SetInstanceType(v string) *DescribePriceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *DescribePriceRequest
	GetInternetChargeType() *string
	SetOsType(v string) *DescribePriceRequest
	GetOsType() *string
	SetPeriod(v int32) *DescribePriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribePriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribePriceRequest
	GetPromotionId() *string
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DescribePriceRequest
	GetResellerOwnerUid() *int64
	SetResourceType(v string) *DescribePriceRequest
	GetResourceType() *string
	SetRootDiskCategory(v string) *DescribePriceRequest
	GetRootDiskCategory() *string
	SetRootDiskPerformanceLevel(v string) *DescribePriceRequest
	GetRootDiskPerformanceLevel() *string
	SetRootDiskSizeGib(v int32) *DescribePriceRequest
	GetRootDiskSizeGib() *int32
	SetUserDiskCategory(v string) *DescribePriceRequest
	GetUserDiskCategory() *string
	SetUserDiskPerformanceLevel(v string) *DescribePriceRequest
	GetUserDiskPerformanceLevel() *string
	SetUserDiskSizeGib(v int32) *DescribePriceRequest
	GetUserDiskSizeGib() *int32
}

type DescribePriceRequest struct {
	// The number of resources. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The maximum public bandwidth. Unit: Mbit/s.
	//
	// 	- Valid values if you set InternetChargeType to PayByBandwidth: 10 to 1000.
	//
	// 	- Valid values if you set InternetChargeType to InternetChargeType: 10 to 200.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of hourly plan if you use the Monthly Subscription billing method. If you set `ResourceType` to `DesktopMonthPackage`, you must specify this parameter.
	//
	// Valid values:
	//
	// 	- 120: the 120-hour computing plan.
	//
	// 	- 250: the 250-hour computing plan.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of cloud computer shares. Default value: 1.
	//
	// >  This parameter takes effect only if you set `ResourceType` to `DesktopGroup`.
	//
	// example:
	//
	// 1
	GroupDesktopCount *int32 `json:"GroupDesktopCount,omitempty" xml:"GroupDesktopCount,omitempty"`
	// The specifications of the resource.
	//
	// 	- This parameter is required if you set `ResourceType` to `Desktop`. You can call the [DescribeDesktopTypes](~~DescribeDesktopTypes~~) to query the available cloud computer types that correspond to the value of `DesktopTypeId`.
	//
	// 	- If you set `ResourceType` to `DesktopGroup`, set the value of this parameter to `large`.
	//
	// 	- If you set `ResourceType` to `Bandwidth`, you can leave this parameter empty.
	//
	// example:
	//
	// eds.general.2c2g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The metering method for network traffic.
	//
	// Valid values:
	//
	// 	- PayByTraffic: You are charged for the actually consumed traffic.
	//
	// 	- PayByBandwidth: You are charged by a fixed bandwidth.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The OS type.
	//
	// Valid values:
	//
	// 	- Linux
	//
	// 	- Windows (default)
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The subscription duration. The valid values of this parameter vary based on the value of `PeriodUnit`.
	//
	// 	- If you set `PeriodUnit` to `Hour`, set the value of this parameter to 1.
	//
	// 	- If you set `PeriodUnit` to `Month`, set the value of this parameter to 1, 2, 3, or 6.
	//
	// 	- If you set `PeriodUnit` to `Year`, set the value of this parameter to 1, 2, or 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Hour (default)
	//
	// example:
	//
	// Hour
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 123456
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by EDS.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The type of the resource.
	//
	// Valid values:
	//
	// 	- DesktopMonthPackage: monthly subscription cloud computers that use hourly limit plans.
	//
	// 	- Desktop (default): pay-as-you-go cloud computers/monthly subscription cloud computers that use unlimited plans.
	//
	// 	- Bandwidth: premium bandwidth plans.
	//
	// 	- DesktopGroup: cloud computer shares.
	//
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The category of the system disk.
	//
	// Valid values:
	//
	// 	- cloud_efficiency: the ultra disk
	//
	// 	- cloud_auto: the standard SSD.
	//
	// 	- cloud_essd: the Enterprise SSD (ESSD). Take note that only specific cloud computer types support ESSDs.
	//
	// example:
	//
	// 40
	RootDiskCategory         *string `json:"RootDiskCategory,omitempty" xml:"RootDiskCategory,omitempty"`
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. If you set `ResourceType` to `Desktop`, you must specify this parameter.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The category of the data disk.
	//
	// Valid values:
	//
	// 	- cloud_efficiency: the ultra disk
	//
	// 	- cloud_auto: the standard SSD.
	//
	// 	- cloud_essd: the ESSD. Take note that only specific cloud computer types support ESSDs.
	//
	// example:
	//
	// 80
	UserDiskCategory         *string `json:"UserDiskCategory,omitempty" xml:"UserDiskCategory,omitempty"`
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The size of the data disk. Unit: GiB.
	//
	// example:
	//
	// 100
	UserDiskSizeGib *int32 `json:"UserDiskSizeGib,omitempty" xml:"UserDiskSizeGib,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribePriceRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribePriceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribePriceRequest) GetGroupDesktopCount() *int32 {
	return s.GroupDesktopCount
}

func (s *DescribePriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePriceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribePriceRequest) GetOsType() *string {
	return s.OsType
}

func (s *DescribePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribePriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DescribePriceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePriceRequest) GetRootDiskCategory() *string {
	return s.RootDiskCategory
}

func (s *DescribePriceRequest) GetRootDiskPerformanceLevel() *string {
	return s.RootDiskPerformanceLevel
}

func (s *DescribePriceRequest) GetRootDiskSizeGib() *int32 {
	return s.RootDiskSizeGib
}

func (s *DescribePriceRequest) GetUserDiskCategory() *string {
	return s.UserDiskCategory
}

func (s *DescribePriceRequest) GetUserDiskPerformanceLevel() *string {
	return s.UserDiskPerformanceLevel
}

func (s *DescribePriceRequest) GetUserDiskSizeGib() *int32 {
	return s.UserDiskSizeGib
}

func (s *DescribePriceRequest) SetAmount(v int32) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetBandwidth(v int32) *DescribePriceRequest {
	s.Bandwidth = &v
	return s
}

func (s *DescribePriceRequest) SetDuration(v int32) *DescribePriceRequest {
	s.Duration = &v
	return s
}

func (s *DescribePriceRequest) SetGroupDesktopCount(v int32) *DescribePriceRequest {
	s.GroupDesktopCount = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetOsType(v string) *DescribePriceRequest {
	s.OsType = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetPeriodUnit(v string) *DescribePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceRequest) SetPromotionId(v string) *DescribePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResellerOwnerUid(v int64) *DescribePriceRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DescribePriceRequest) SetResourceType(v string) *DescribePriceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskCategory(v string) *DescribePriceRequest {
	s.RootDiskCategory = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskPerformanceLevel(v string) *DescribePriceRequest {
	s.RootDiskPerformanceLevel = &v
	return s
}

func (s *DescribePriceRequest) SetRootDiskSizeGib(v int32) *DescribePriceRequest {
	s.RootDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskCategory(v string) *DescribePriceRequest {
	s.UserDiskCategory = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskPerformanceLevel(v string) *DescribePriceRequest {
	s.UserDiskPerformanceLevel = &v
	return s
}

func (s *DescribePriceRequest) SetUserDiskSizeGib(v int32) *DescribePriceRequest {
	s.UserDiskSizeGib = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	return dara.Validate(s)
}
