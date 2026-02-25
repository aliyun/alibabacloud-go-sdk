// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *ModifyChargeTypeRequest
	GetAutoRenew() *bool
	SetBillingInstanceIds(v string) *ModifyChargeTypeRequest
	GetBillingInstanceIds() *string
	SetDuration(v string) *ModifyChargeTypeRequest
	GetDuration() *string
	SetInstanceId(v string) *ModifyChargeTypeRequest
	GetInstanceId() *string
	SetPricingCycle(v string) *ModifyChargeTypeRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *ModifyChargeTypeRequest
	GetPromotionOptionNo() *string
}

type ModifyChargeTypeRequest struct {
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
	// 2345
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s ModifyChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyChargeTypeRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ModifyChargeTypeRequest) GetBillingInstanceIds() *string {
	return s.BillingInstanceIds
}

func (s *ModifyChargeTypeRequest) GetDuration() *string {
	return s.Duration
}

func (s *ModifyChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyChargeTypeRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ModifyChargeTypeRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *ModifyChargeTypeRequest) SetAutoRenew(v bool) *ModifyChargeTypeRequest {
	s.AutoRenew = &v
	return s
}

func (s *ModifyChargeTypeRequest) SetBillingInstanceIds(v string) *ModifyChargeTypeRequest {
	s.BillingInstanceIds = &v
	return s
}

func (s *ModifyChargeTypeRequest) SetDuration(v string) *ModifyChargeTypeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyChargeTypeRequest) SetInstanceId(v string) *ModifyChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyChargeTypeRequest) SetPricingCycle(v string) *ModifyChargeTypeRequest {
	s.PricingCycle = &v
	return s
}

func (s *ModifyChargeTypeRequest) SetPromotionOptionNo(v string) *ModifyChargeTypeRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *ModifyChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
