// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlansInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateSavingsPlansInstanceShrinkRequest
	GetAutoPay() *bool
	SetCommodityCode(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetCommodityCode() *string
	SetDuration(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetDuration() *string
	SetEffectiveDate(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetEffectiveDate() *string
	SetExtendMapShrink(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetExtendMapShrink() *string
	SetPayMode(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetPayMode() *string
	SetPoolValue(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetPoolValue() *string
	SetPricingCycle(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetPricingCycle() *string
	SetRegion(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetRegion() *string
	SetSpecType(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetSpecType() *string
	SetSpecification(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetSpecification() *string
	SetType(v string) *CreateSavingsPlansInstanceShrinkRequest
	GetType() *string
}

type CreateSavingsPlansInstanceShrinkRequest struct {
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// This parameter is required.
	Duration      *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	// if can be null:
	// true
	ExtendMapShrink *string `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	// This parameter is required.
	PayMode *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	// This parameter is required.
	PoolValue *string `json:"PoolValue,omitempty" xml:"PoolValue,omitempty"`
	// This parameter is required.
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SpecType      *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateSavingsPlansInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetEffectiveDate() *string {
	return s.EffectiveDate
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetExtendMapShrink() *string {
	return s.ExtendMapShrink
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetPayMode() *string {
	return s.PayMode
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetPoolValue() *string {
	return s.PoolValue
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetSpecification() *string {
	return s.Specification
}

func (s *CreateSavingsPlansInstanceShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetAutoPay(v bool) *CreateSavingsPlansInstanceShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetCommodityCode(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetDuration(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetEffectiveDate(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.EffectiveDate = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetExtendMapShrink(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.ExtendMapShrink = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetPayMode(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.PayMode = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetPoolValue(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.PoolValue = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetPricingCycle(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetRegion(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetSpecType(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetSpecification(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.Specification = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) SetType(v string) *CreateSavingsPlansInstanceShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateSavingsPlansInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
