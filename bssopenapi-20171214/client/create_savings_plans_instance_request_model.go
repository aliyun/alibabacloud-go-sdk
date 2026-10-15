// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavingsPlansInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateSavingsPlansInstanceRequest
	GetAutoPay() *bool
	SetCommodityCode(v string) *CreateSavingsPlansInstanceRequest
	GetCommodityCode() *string
	SetDuration(v string) *CreateSavingsPlansInstanceRequest
	GetDuration() *string
	SetEffectiveDate(v string) *CreateSavingsPlansInstanceRequest
	GetEffectiveDate() *string
	SetExtendMap(v map[string]*string) *CreateSavingsPlansInstanceRequest
	GetExtendMap() map[string]*string
	SetPayMode(v string) *CreateSavingsPlansInstanceRequest
	GetPayMode() *string
	SetPoolValue(v string) *CreateSavingsPlansInstanceRequest
	GetPoolValue() *string
	SetPricingCycle(v string) *CreateSavingsPlansInstanceRequest
	GetPricingCycle() *string
	SetRegion(v string) *CreateSavingsPlansInstanceRequest
	GetRegion() *string
	SetSpecType(v string) *CreateSavingsPlansInstanceRequest
	GetSpecType() *string
	SetSpecification(v string) *CreateSavingsPlansInstanceRequest
	GetSpecification() *string
	SetType(v string) *CreateSavingsPlansInstanceRequest
	GetType() *string
}

type CreateSavingsPlansInstanceRequest struct {
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// This parameter is required.
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// This parameter is required.
	Duration      *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	// if can be null:
	// true
	ExtendMap map[string]*string `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
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

func (s CreateSavingsPlansInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSavingsPlansInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSavingsPlansInstanceRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateSavingsPlansInstanceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateSavingsPlansInstanceRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateSavingsPlansInstanceRequest) GetEffectiveDate() *string {
	return s.EffectiveDate
}

func (s *CreateSavingsPlansInstanceRequest) GetExtendMap() map[string]*string {
	return s.ExtendMap
}

func (s *CreateSavingsPlansInstanceRequest) GetPayMode() *string {
	return s.PayMode
}

func (s *CreateSavingsPlansInstanceRequest) GetPoolValue() *string {
	return s.PoolValue
}

func (s *CreateSavingsPlansInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateSavingsPlansInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateSavingsPlansInstanceRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreateSavingsPlansInstanceRequest) GetSpecification() *string {
	return s.Specification
}

func (s *CreateSavingsPlansInstanceRequest) GetType() *string {
	return s.Type
}

func (s *CreateSavingsPlansInstanceRequest) SetAutoPay(v bool) *CreateSavingsPlansInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetCommodityCode(v string) *CreateSavingsPlansInstanceRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetDuration(v string) *CreateSavingsPlansInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetEffectiveDate(v string) *CreateSavingsPlansInstanceRequest {
	s.EffectiveDate = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetExtendMap(v map[string]*string) *CreateSavingsPlansInstanceRequest {
	s.ExtendMap = v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetPayMode(v string) *CreateSavingsPlansInstanceRequest {
	s.PayMode = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetPoolValue(v string) *CreateSavingsPlansInstanceRequest {
	s.PoolValue = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetPricingCycle(v string) *CreateSavingsPlansInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetRegion(v string) *CreateSavingsPlansInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetSpecType(v string) *CreateSavingsPlansInstanceRequest {
	s.SpecType = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetSpecification(v string) *CreateSavingsPlansInstanceRequest {
	s.Specification = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) SetType(v string) *CreateSavingsPlansInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateSavingsPlansInstanceRequest) Validate() error {
	return dara.Validate(s)
}
