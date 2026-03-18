// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyChargeTypePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *QueryModifyChargeTypePriceRequest
	GetAutoRenew() *bool
	SetBillingInstanceIds(v string) *QueryModifyChargeTypePriceRequest
	GetBillingInstanceIds() *string
	SetDuration(v string) *QueryModifyChargeTypePriceRequest
	GetDuration() *string
	SetInstanceId(v string) *QueryModifyChargeTypePriceRequest
	GetInstanceId() *string
	SetPricingCycle(v string) *QueryModifyChargeTypePriceRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *QueryModifyChargeTypePriceRequest
	GetPromotionOptionNo() *string
}

type QueryModifyChargeTypePriceRequest struct {
	// example:
	//
	// True
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// c-3d6dc31ba67b1839
	BillingInstanceIds *string `json:"BillingInstanceIds,omitempty" xml:"BillingInstanceIds,omitempty"`
	// example:
	//
	// 2
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s QueryModifyChargeTypePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyChargeTypePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyChargeTypePriceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *QueryModifyChargeTypePriceRequest) GetBillingInstanceIds() *string {
	return s.BillingInstanceIds
}

func (s *QueryModifyChargeTypePriceRequest) GetDuration() *string {
	return s.Duration
}

func (s *QueryModifyChargeTypePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyChargeTypePriceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryModifyChargeTypePriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryModifyChargeTypePriceRequest) SetAutoRenew(v bool) *QueryModifyChargeTypePriceRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryModifyChargeTypePriceRequest) SetBillingInstanceIds(v string) *QueryModifyChargeTypePriceRequest {
	s.BillingInstanceIds = &v
	return s
}

func (s *QueryModifyChargeTypePriceRequest) SetDuration(v string) *QueryModifyChargeTypePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryModifyChargeTypePriceRequest) SetInstanceId(v string) *QueryModifyChargeTypePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyChargeTypePriceRequest) SetPricingCycle(v string) *QueryModifyChargeTypePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryModifyChargeTypePriceRequest) SetPromotionOptionNo(v string) *QueryModifyChargeTypePriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryModifyChargeTypePriceRequest) Validate() error {
	return dara.Validate(s)
}
