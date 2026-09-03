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
	// The resource count. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The peak Internet bandwidth. Unit: Mbit/s.
	//
	// - For pay-by-bandwidth, valid values are 10 to 1000.
	//
	// - For pay-by-traffic, valid values are 10 to 200.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The duration package type for monthly cloud desktop purchases. If ResourceType is set to DesktopMonthPackage, this parameter is required.
	//
	// example:
	//
	// 120
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of shared cloud desktops. Default value: 1.
	//
	// > This parameter takes effect only when ResourceType is set to DesktopGroup.
	//
	// example:
	//
	// 1
	GroupDesktopCount *int32 `json:"GroupDesktopCount,omitempty" xml:"GroupDesktopCount,omitempty"`
	// The resource specification.
	//
	// - If ResourceType is set to Desktop, this parameter is required. You can call [DescribeDesktopTypes](~~DescribeDesktopTypes~~) to query available values (corresponding to the DesktopTypeId value).
	//
	// - If ResourceType is set to DesktopGroup, set this parameter to `large`.
	//
	// - If ResourceType is set to Bandwidth, you do not need to specify this parameter.
	//
	// example:
	//
	// eds.general.2c2g
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method of the Internet access package.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The subscription duration. Valid values are determined by the PeriodUnit parameter.
	//
	// - If PeriodUnit is set to Hour, the valid value is 1.
	//
	// - If PeriodUnit is set to Month, valid values are 1, 2, 3, and 6.
	//
	// - If PeriodUnit is set to Year, valid values are 1, 2, and 3.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle.
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
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID for resource ownership in reseller mode. You do not need to specify this parameter in non-reseller mode.
	//
	// example:
	//
	// 1422724566551XXX
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The system cloud disk type.
	//
	// example:
	//
	// 40
	RootDiskCategory *string `json:"RootDiskCategory,omitempty" xml:"RootDiskCategory,omitempty"`
	// The performance level (PL) of the system cloud disk. You can set the disk performance level when the cloud desktop specification is set to Graphics or High Frequency. For more information about the differences between performance levels, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	RootDiskPerformanceLevel *string `json:"RootDiskPerformanceLevel,omitempty" xml:"RootDiskPerformanceLevel,omitempty"`
	// The system cloud disk size. Unit: GiB. If ResourceType is set to Desktop, this parameter is required.
	//
	// example:
	//
	// 80
	RootDiskSizeGib *int32 `json:"RootDiskSizeGib,omitempty" xml:"RootDiskSizeGib,omitempty"`
	// The data cloud disk type.
	//
	// example:
	//
	// 80
	UserDiskCategory *string `json:"UserDiskCategory,omitempty" xml:"UserDiskCategory,omitempty"`
	// The performance level (PL) of the data cloud disk. You can set the disk performance level when the cloud desktop specification is set to Graphics or High Frequency. For more information about the differences between performance levels, see [ESSD cloud disks](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	UserDiskPerformanceLevel *string `json:"UserDiskPerformanceLevel,omitempty" xml:"UserDiskPerformanceLevel,omitempty"`
	// The data cloud disk size. Unit: GiB.
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
