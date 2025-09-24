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
	ExtendMap map[string]*string `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
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
