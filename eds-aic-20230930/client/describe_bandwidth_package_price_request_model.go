// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackagePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *DescribeBandwidthPackagePriceRequest
	GetBandwidth() *int32
	SetBizRegionId(v string) *DescribeBandwidthPackagePriceRequest
	GetBizRegionId() *string
	SetInternetChargeType(v string) *DescribeBandwidthPackagePriceRequest
	GetInternetChargeType() *string
	SetPeriod(v int32) *DescribeBandwidthPackagePriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribeBandwidthPackagePriceRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *DescribeBandwidthPackagePriceRequest
	GetPromotionId() *string
	SetPurchaseChannelType(v string) *DescribeBandwidthPackagePriceRequest
	GetPurchaseChannelType() *string
}

type DescribeBandwidthPackagePriceRequest struct {
	// The peak bandwidth of the network. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method of the bandwidth plan. Valid values:
	//
	// <props="china">
	//
	// - **PayByBandwidth*	- (default): billing by bandwidth.
	//
	// - **PayBy95**: billing by traditional 95th percentile. IPv6 Internet bandwidth does not support billing by traditional 95th percentile by default. To use this billing method, contact your account manager.
	//
	//
	// <props="intl">
	//
	// - **PayByBandwidth*	- (default): billing by bandwidth.
	//
	// example:
	//
	// PayByBandwidth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The subscription duration of the resource. The unit is specified by `PeriodUnit`. Valid values: 1 month, 2 months, 3 months, 6 months, 1 year, 2 years, and 3 years.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration of the resource. Valid values:
	//
	// - **Month**: Month.
	//
	// - **Year**: Year.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the promotion.
	//
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The payment channel.
	//
	// example:
	//
	// JvsMobile
	PurchaseChannelType *string `json:"PurchaseChannelType,omitempty" xml:"PurchaseChannelType,omitempty"`
}

func (s DescribeBandwidthPackagePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagePriceRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeBandwidthPackagePriceRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeBandwidthPackagePriceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeBandwidthPackagePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeBandwidthPackagePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeBandwidthPackagePriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeBandwidthPackagePriceRequest) GetPurchaseChannelType() *string {
	return s.PurchaseChannelType
}

func (s *DescribeBandwidthPackagePriceRequest) SetBandwidth(v int32) *DescribeBandwidthPackagePriceRequest {
	s.Bandwidth = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) SetBizRegionId(v string) *DescribeBandwidthPackagePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) SetInternetChargeType(v string) *DescribeBandwidthPackagePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) SetPeriod(v int32) *DescribeBandwidthPackagePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) SetPeriodUnit(v string) *DescribeBandwidthPackagePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) SetPromotionId(v string) *DescribeBandwidthPackagePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) SetPurchaseChannelType(v string) *DescribeBandwidthPackagePriceRequest {
	s.PurchaseChannelType = &v
	return s
}

func (s *DescribeBandwidthPackagePriceRequest) Validate() error {
	return dara.Validate(s)
}
