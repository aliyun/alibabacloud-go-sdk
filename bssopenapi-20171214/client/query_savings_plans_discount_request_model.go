// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansDiscountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *QuerySavingsPlansDiscountRequest
	GetCommodityCode() *string
	SetCycle(v string) *QuerySavingsPlansDiscountRequest
	GetCycle() *string
	SetLocale(v string) *QuerySavingsPlansDiscountRequest
	GetLocale() *string
	SetModuleCode(v string) *QuerySavingsPlansDiscountRequest
	GetModuleCode() *string
	SetPageNum(v int32) *QuerySavingsPlansDiscountRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySavingsPlansDiscountRequest
	GetPageSize() *int32
	SetPayMode(v string) *QuerySavingsPlansDiscountRequest
	GetPayMode() *string
	SetRegion(v string) *QuerySavingsPlansDiscountRequest
	GetRegion() *string
	SetSpec(v string) *QuerySavingsPlansDiscountRequest
	GetSpec() *string
	SetSpnCommodityCode(v string) *QuerySavingsPlansDiscountRequest
	GetSpnCommodityCode() *string
	SetSpnType(v string) *QuerySavingsPlansDiscountRequest
	GetSpnType() *string
}

type QuerySavingsPlansDiscountRequest struct {
	// The code of the service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The cycle based on which queries are performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1:Year
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The identifier of the language.
	//
	// Valid values:
	//
	// 	- EN: English.
	//
	// 	- ZH: Chinese.
	//
	// example:
	//
	// ZH
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The code of the pricing module.
	//
	// example:
	//
	// instance_type
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The payment mode. Valid values: total: all upfront. half: half upfront. zero: no upfront.
	//
	// This parameter is required.
	//
	// example:
	//
	// total
	PayMode *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ecs.g6
	Spec             *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	SpnCommodityCode *string `json:"SpnCommodityCode,omitempty" xml:"SpnCommodityCode,omitempty"`
	// The type of the savings plan. Valid values: ecs: Elastic Compute Service (ECS) compute type. universal: general-purpose type.
	//
	// This parameter is required.
	//
	// example:
	//
	// universal
	SpnType *string `json:"SpnType,omitempty" xml:"SpnType,omitempty"`
}

func (s QuerySavingsPlansDiscountRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDiscountRequest) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDiscountRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySavingsPlansDiscountRequest) GetCycle() *string {
	return s.Cycle
}

func (s *QuerySavingsPlansDiscountRequest) GetLocale() *string {
	return s.Locale
}

func (s *QuerySavingsPlansDiscountRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *QuerySavingsPlansDiscountRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySavingsPlansDiscountRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySavingsPlansDiscountRequest) GetPayMode() *string {
	return s.PayMode
}

func (s *QuerySavingsPlansDiscountRequest) GetRegion() *string {
	return s.Region
}

func (s *QuerySavingsPlansDiscountRequest) GetSpec() *string {
	return s.Spec
}

func (s *QuerySavingsPlansDiscountRequest) GetSpnCommodityCode() *string {
	return s.SpnCommodityCode
}

func (s *QuerySavingsPlansDiscountRequest) GetSpnType() *string {
	return s.SpnType
}

func (s *QuerySavingsPlansDiscountRequest) SetCommodityCode(v string) *QuerySavingsPlansDiscountRequest {
	s.CommodityCode = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetCycle(v string) *QuerySavingsPlansDiscountRequest {
	s.Cycle = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetLocale(v string) *QuerySavingsPlansDiscountRequest {
	s.Locale = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetModuleCode(v string) *QuerySavingsPlansDiscountRequest {
	s.ModuleCode = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetPageNum(v int32) *QuerySavingsPlansDiscountRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetPageSize(v int32) *QuerySavingsPlansDiscountRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetPayMode(v string) *QuerySavingsPlansDiscountRequest {
	s.PayMode = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetRegion(v string) *QuerySavingsPlansDiscountRequest {
	s.Region = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetSpec(v string) *QuerySavingsPlansDiscountRequest {
	s.Spec = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetSpnCommodityCode(v string) *QuerySavingsPlansDiscountRequest {
	s.SpnCommodityCode = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) SetSpnType(v string) *QuerySavingsPlansDiscountRequest {
	s.SpnType = &v
	return s
}

func (s *QuerySavingsPlansDiscountRequest) Validate() error {
	return dara.Validate(s)
}
