// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateAgentResourceRequest
	GetAutoRenew() *bool
	SetCu(v int32) *CreateAgentResourceRequest
	GetCu() *int32
	SetDuration(v int32) *CreateAgentResourceRequest
	GetDuration() *int32
	SetInstanceId(v string) *CreateAgentResourceRequest
	GetInstanceId() *string
	SetPayType(v string) *CreateAgentResourceRequest
	GetPayType() *string
	SetPricingCycle(v string) *CreateAgentResourceRequest
	GetPricingCycle() *string
	SetPromotionOptionNo(v string) *CreateAgentResourceRequest
	GetPromotionOptionNo() *string
	SetSpecType(v string) *CreateAgentResourceRequest
	GetSpecType() *string
}

type CreateAgentResourceRequest struct {
	// Enable auto-renewal. This parameter is valid only when payType is set to PrePaid. Auto-renewal is disabled by default.
	//
	// example:
	//
	// True
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Number of CUs. A CU (Compute Unit) is the basic unit of service measurement. 1 CU = 1 CPU core + 4 GiB memory. For memory-enhanced instance family, 1 CU = 1 CPU core + 8 GiB memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// Duration. This parameter is valid only when payType is set to PrePaid.
	//
	// example:
	//
	// 2
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Payment type:
	//
	// 1. Subscription (prePaid).
	//
	// 2. Pay-as-you-go (postPaid).
	//
	// This parameter is required.
	//
	// example:
	//
	// prePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Unit of subscription duration:
	//
	// - Month
	//
	// - Year
	//
	// This parameter is valid only when payType is set to PrePaid.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Coupon ID.
	//
	// example:
	//
	// 2345
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// Compute group specification type.
	//
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
}

func (s CreateAgentResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentResourceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateAgentResourceRequest) GetCu() *int32 {
	return s.Cu
}

func (s *CreateAgentResourceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateAgentResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateAgentResourceRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateAgentResourceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateAgentResourceRequest) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *CreateAgentResourceRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreateAgentResourceRequest) SetAutoRenew(v bool) *CreateAgentResourceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAgentResourceRequest) SetCu(v int32) *CreateAgentResourceRequest {
	s.Cu = &v
	return s
}

func (s *CreateAgentResourceRequest) SetDuration(v int32) *CreateAgentResourceRequest {
	s.Duration = &v
	return s
}

func (s *CreateAgentResourceRequest) SetInstanceId(v string) *CreateAgentResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentResourceRequest) SetPayType(v string) *CreateAgentResourceRequest {
	s.PayType = &v
	return s
}

func (s *CreateAgentResourceRequest) SetPricingCycle(v string) *CreateAgentResourceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateAgentResourceRequest) SetPromotionOptionNo(v string) *CreateAgentResourceRequest {
	s.PromotionOptionNo = &v
	return s
}

func (s *CreateAgentResourceRequest) SetSpecType(v string) *CreateAgentResourceRequest {
	s.SpecType = &v
	return s
}

func (s *CreateAgentResourceRequest) Validate() error {
	return dara.Validate(s)
}
