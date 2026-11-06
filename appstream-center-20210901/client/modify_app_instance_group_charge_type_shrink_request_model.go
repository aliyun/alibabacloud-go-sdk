// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupChargeTypeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupIdsShrink(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetAppInstanceGroupIdsShrink() *string
	SetAutoPay(v bool) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetBizRegionId() *string
	SetChargeType(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetChargeType() *string
	SetPeriod(v int32) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetPeriodUnit() *string
	SetProductType(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetProductType() *string
	SetPromotionId(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest
	GetPromotionId() *string
}

type ModifyAppInstanceGroupChargeTypeShrinkRequest struct {
	// The list of delivery group IDs to be modified, in JSON array format. **You must and can only specify one item per request**. The resources must belong to the current account, the specified product, and the business region.
	//
	// This parameter is required.
	AppInstanceGroupIdsShrink *string `json:"AppInstanceGroupIds,omitempty" xml:"AppInstanceGroupIds,omitempty"`
	// Specifies whether to allow automatic payment for the conversion order. Default value: `false`. Valid values:
	//
	// - `true`: Automatic payment is allowed.
	//
	// - `false`: Automatic payment is not allowed. If an unpaid order is generated, you must complete the payment separately.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the order. We recommend that you explicitly specify this parameter for subscription orders to determine whether to enable auto-renewal. Valid values:
	//
	// - `true`: Auto-renewal is enabled.
	//
	// - `false`: Auto-renewal is not enabled.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The business region ID of the target delivery group. It must be the same as the actual region to which the resource belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The target billing method after conversion. It must be different from the current billing method. Valid values:
	//
	// - `PrePaid`: subscription
	//
	// - `PostPaid`: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The numeric part of the purchase duration, used together with `PeriodUnit`. When converting to subscription and a purchase duration is required, specify both parameters. For example, `1` and `Month` indicate a purchase duration of one month. The available purchase durations are determined by the sales rules of the target resource.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the purchase duration, used together with `Period`. This parameter is case-sensitive. Valid values:
	//
	// - `Hour`: hour
	//
	// - `Day`: day
	//
	// - `Week`: week
	//
	// - `Month`: month
	//
	// - `Year`: year
	//
	// These units do not mean that any combination is available for every product. When converting to subscription, select a period supported by the target resource. For example, for a resource that supports monthly purchase, you can set `Period` to `1` and `PeriodUnit` to `Month`.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The product type. Set this parameter to the following fixed value when you call the operation:
	//
	// - `CloudApp`: WUYING Cloud Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the promotion used for the conversion order. It takes effect only when an order is created and the promotion is applicable. If omitted, no specific coupon is specified. The availability of the promotion is subject to the order result.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
}

func (s ModifyAppInstanceGroupChargeTypeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupChargeTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetAppInstanceGroupIdsShrink() *string {
	return s.AppInstanceGroupIdsShrink
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetAppInstanceGroupIdsShrink(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.AppInstanceGroupIdsShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetAutoPay(v bool) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetAutoRenew(v bool) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetBizRegionId(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetChargeType(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetPeriod(v int32) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.Period = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetPeriodUnit(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetProductType(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) SetPromotionId(v string) *ModifyAppInstanceGroupChargeTypeShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyAppInstanceGroupChargeTypeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
