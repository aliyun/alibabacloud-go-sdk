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
	// example:
	//
	// True
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
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
	// prePay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 2345
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
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
