// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateBillingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingType(v string) *ModelRouterUpdateBillingRuleRequest
	GetBillingType() *string
	SetEffectiveTime(v string) *ModelRouterUpdateBillingRuleRequest
	GetEffectiveTime() *string
	SetExpireTime(v string) *ModelRouterUpdateBillingRuleRequest
	GetExpireTime() *string
	SetPricingConfig(v interface{}) *ModelRouterUpdateBillingRuleRequest
	GetPricingConfig() interface{}
	SetStatus(v int32) *ModelRouterUpdateBillingRuleRequest
	GetStatus() *int32
	SetVersion(v int32) *ModelRouterUpdateBillingRuleRequest
	GetVersion() *int32
}

type ModelRouterUpdateBillingRuleRequest struct {
	// example:
	//
	// token_tiered
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveTime *string `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// example:
	//
	// 2025-01-01T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// {}
	PricingConfig interface{} `json:"pricingConfig,omitempty" xml:"pricingConfig,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelRouterUpdateBillingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateBillingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateBillingRuleRequest) GetBillingType() *string {
	return s.BillingType
}

func (s *ModelRouterUpdateBillingRuleRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterUpdateBillingRuleRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterUpdateBillingRuleRequest) GetPricingConfig() interface{} {
	return s.PricingConfig
}

func (s *ModelRouterUpdateBillingRuleRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModelRouterUpdateBillingRuleRequest) GetVersion() *int32 {
	return s.Version
}

func (s *ModelRouterUpdateBillingRuleRequest) SetBillingType(v string) *ModelRouterUpdateBillingRuleRequest {
	s.BillingType = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleRequest) SetEffectiveTime(v string) *ModelRouterUpdateBillingRuleRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleRequest) SetExpireTime(v string) *ModelRouterUpdateBillingRuleRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleRequest) SetPricingConfig(v interface{}) *ModelRouterUpdateBillingRuleRequest {
	s.PricingConfig = v
	return s
}

func (s *ModelRouterUpdateBillingRuleRequest) SetStatus(v int32) *ModelRouterUpdateBillingRuleRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleRequest) SetVersion(v int32) *ModelRouterUpdateBillingRuleRequest {
	s.Version = &v
	return s
}

func (s *ModelRouterUpdateBillingRuleRequest) Validate() error {
	return dara.Validate(s)
}
