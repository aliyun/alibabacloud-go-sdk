// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribePriceRequest
	GetAcceptLanguage() *string
	SetAmount(v int32) *DescribePriceRequest
	GetAmount() *int32
	SetCommodityType(v string) *DescribePriceRequest
	GetCommodityType() *string
	SetDataDiskSize(v int32) *DescribePriceRequest
	GetDataDiskSize() *int32
	SetInstanceId(v string) *DescribePriceRequest
	GetInstanceId() *string
	SetOrderType(v string) *DescribePriceRequest
	GetOrderType() *string
	SetPayType(v string) *DescribePriceRequest
	GetPayType() *string
	SetPeriod(v int32) *DescribePriceRequest
	GetPeriod() *int32
	SetPlanId(v string) *DescribePriceRequest
	GetPlanId() *string
	SetPriceUnit(v string) *DescribePriceRequest
	GetPriceUnit() *string
	SetPromotionOptions(v *DescribePriceRequestPromotionOptions) *DescribePriceRequest
	GetPromotionOptions() *DescribePriceRequestPromotionOptions
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
}

type DescribePriceRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh-CN*	- (default): Chinese.
	//
	// - **en-US**: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The number of Simple Application Server instances to create. Valid values: 1 to 20.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The resource type. Valid values:
	//
	// - Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// Server
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// The capacity of the mounted data cloud disk. Unit: GB. Valid values: 0 to 16380. The value must be a multiple of 20.
	//
	// example:
	//
	// 20
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The ID of the Simple Application Server instance.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order type. Valid values:
	//
	// - Buy.
	//
	// Default value: Buy.
	//
	// example:
	//
	// UPLOAD
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - Prepaid: subscription.
	//
	// Example: Prepaid.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing duration. Valid values:
	//
	// - If the PriceUnit parameter is set to Month: 1 to 9.
	//
	// - If the PriceUnit parameter is set to Year: 1 to 5.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The plan ID. You can call the [ListPlans](https://help.aliyun.com/document_detail/189314.html) operation to query all plan information in the specified region.
	//
	// example:
	//
	// swas.s1.cxxxxxx3t05
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The pricing unit. Valid values:
	//
	// - Month: the pricing unit for monthly billing.
	//
	// - Year: the pricing unit for yearly billing.
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The promotion option information.
	PromotionOptions *DescribePriceRequestPromotionOptions `json:"PromotionOptions,omitempty" xml:"PromotionOptions,omitempty" type:"Struct"`
	// The region ID of the specified Simple Application Server instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribePriceRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribePriceRequest) GetCommodityType() *string {
	return s.CommodityType
}

func (s *DescribePriceRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePriceRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DescribePriceRequest) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *DescribePriceRequest) GetPromotionOptions() *DescribePriceRequestPromotionOptions {
	return s.PromotionOptions
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) SetAcceptLanguage(v string) *DescribePriceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribePriceRequest) SetAmount(v int32) *DescribePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequest) SetCommodityType(v string) *DescribePriceRequest {
	s.CommodityType = &v
	return s
}

func (s *DescribePriceRequest) SetDataDiskSize(v int32) *DescribePriceRequest {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceId(v string) *DescribePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetPayType(v string) *DescribePriceRequest {
	s.PayType = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetPlanId(v string) *DescribePriceRequest {
	s.PlanId = &v
	return s
}

func (s *DescribePriceRequest) SetPriceUnit(v string) *DescribePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribePriceRequest) SetPromotionOptions(v *DescribePriceRequestPromotionOptions) *DescribePriceRequest {
	s.PromotionOptions = v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	if s.PromotionOptions != nil {
		if err := s.PromotionOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceRequestPromotionOptions struct {
	// The activity ID.
	//
	// example:
	//
	// 12312399123
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
}

func (s DescribePriceRequestPromotionOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestPromotionOptions) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestPromotionOptions) GetActivityId() *string {
	return s.ActivityId
}

func (s *DescribePriceRequestPromotionOptions) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceRequestPromotionOptions) SetActivityId(v string) *DescribePriceRequestPromotionOptions {
	s.ActivityId = &v
	return s
}

func (s *DescribePriceRequestPromotionOptions) SetCouponNo(v string) *DescribePriceRequestPromotionOptions {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceRequestPromotionOptions) Validate() error {
	return dara.Validate(s)
}
