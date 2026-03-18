// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRenewPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingInstanceIds(v string) *QueryRenewPriceRequest
	GetBillingInstanceIds() *string
	SetDuration(v int32) *QueryRenewPriceRequest
	GetDuration() *int32
	SetInstanceId(v string) *QueryRenewPriceRequest
	GetInstanceId() *string
	SetPricingCycle(v string) *QueryRenewPriceRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *QueryRenewPriceRequest
	GetPromotionOptionNo() *string
}

type QueryRenewPriceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-3d6dc31ba67b1839
	BillingInstanceIds *string `json:"BillingInstanceIds,omitempty" xml:"BillingInstanceIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// youhuiquan_12378dfj6
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s QueryRenewPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRenewPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryRenewPriceRequest) GetBillingInstanceIds() *string {
	return s.BillingInstanceIds
}

func (s *QueryRenewPriceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryRenewPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryRenewPriceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryRenewPriceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *QueryRenewPriceRequest) SetBillingInstanceIds(v string) *QueryRenewPriceRequest {
	s.BillingInstanceIds = &v
	return s
}

func (s *QueryRenewPriceRequest) SetDuration(v int32) *QueryRenewPriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryRenewPriceRequest) SetInstanceId(v string) *QueryRenewPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRenewPriceRequest) SetPricingCycle(v string) *QueryRenewPriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryRenewPriceRequest) SetPromotionOptionNo(v string) *QueryRenewPriceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *QueryRenewPriceRequest) Validate() error {
	return dara.Validate(s)
}
