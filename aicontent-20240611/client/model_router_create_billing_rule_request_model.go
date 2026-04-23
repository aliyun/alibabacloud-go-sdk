// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBillingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingType(v string) *ModelRouterCreateBillingRuleRequest
	GetBillingType() *string
	SetEffectiveTime(v string) *ModelRouterCreateBillingRuleRequest
	GetEffectiveTime() *string
	SetExpireTime(v string) *ModelRouterCreateBillingRuleRequest
	GetExpireTime() *string
	SetModelId(v int64) *ModelRouterCreateBillingRuleRequest
	GetModelId() *int64
	SetPricingConfig(v interface{}) *ModelRouterCreateBillingRuleRequest
	GetPricingConfig() interface{}
	SetVersion(v int32) *ModelRouterCreateBillingRuleRequest
	GetVersion() *int32
}

type ModelRouterCreateBillingRuleRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// {}
	PricingConfig interface{} `json:"pricingConfig,omitempty" xml:"pricingConfig,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModelRouterCreateBillingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBillingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBillingRuleRequest) GetBillingType() *string {
	return s.BillingType
}

func (s *ModelRouterCreateBillingRuleRequest) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ModelRouterCreateBillingRuleRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ModelRouterCreateBillingRuleRequest) GetModelId() *int64 {
	return s.ModelId
}

func (s *ModelRouterCreateBillingRuleRequest) GetPricingConfig() interface{} {
	return s.PricingConfig
}

func (s *ModelRouterCreateBillingRuleRequest) GetVersion() *int32 {
	return s.Version
}

func (s *ModelRouterCreateBillingRuleRequest) SetBillingType(v string) *ModelRouterCreateBillingRuleRequest {
	s.BillingType = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetEffectiveTime(v string) *ModelRouterCreateBillingRuleRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetExpireTime(v string) *ModelRouterCreateBillingRuleRequest {
	s.ExpireTime = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetModelId(v int64) *ModelRouterCreateBillingRuleRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetPricingConfig(v interface{}) *ModelRouterCreateBillingRuleRequest {
	s.PricingConfig = v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) SetVersion(v int32) *ModelRouterCreateBillingRuleRequest {
	s.Version = &v
	return s
}

func (s *ModelRouterCreateBillingRuleRequest) Validate() error {
	return dara.Validate(s)
}
