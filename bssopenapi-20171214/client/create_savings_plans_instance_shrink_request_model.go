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
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// savingplan_common_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The service duration. This parameter is used together with the PricingCycle parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the savings plan takes effect. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-12-31T00:00:00Z
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	// The extended parameters.
	//
	// if can be null:
	// true
	ExtendMapShrink *string `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	// The payment mode. Valid values:
	//
	// 	- total: all upfront
	//
	// 	- half: partial upfront
	//
	// 	- zero: no upfront
	//
	// This parameter is required.
	//
	// example:
	//
	// total
	PayMode *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	// The contracted amount. unit: CNY
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.1
	PoolValue *string `json:"PoolValue,omitempty" xml:"PoolValue,omitempty"`
	// The unit of the subscription duration. This parameter is used together with Duration. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// This parameter is required.
	//
	// example:
	//
	// Year
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the region in which you create the savings plan. You must specify this parameter if the Type parameter is not set to universal.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The specification type. This parameter is used together with the Specification parameter. You must specify this parameter if the Type parameter is not set to universal. Valid values:
	//
	// 	- group: specification group
	//
	// 	- family: specification family
	//
	// example:
	//
	// family
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The specifications of the savings plan. This parameter is used together with the SpecType parameter.
	//
	// example:
	//
	// ecs.g6
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The type of the savings plan. Valid values:
	//
	// 	- universal: general-purpose type
	//
	// 	- ecs: Elastic Compute Service (ECS) compute type
	//
	// 	- elasticy: elastic type
	//
	// This parameter is required.
	//
	// example:
	//
	// universal
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
