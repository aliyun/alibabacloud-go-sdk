// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityType(v string) *DescribeFrInstancesShrinkRequest
	GetCapacityType() *string
	SetCommodityCode(v string) *DescribeFrInstancesShrinkRequest
	GetCommodityCode() *string
	SetCycleType(v string) *DescribeFrInstancesShrinkRequest
	GetCycleType() *string
	SetEcIdAccountIdsShrink(v string) *DescribeFrInstancesShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetEndTime(v int64) *DescribeFrInstancesShrinkRequest
	GetEndTime() *int64
	SetGroup(v string) *DescribeFrInstancesShrinkRequest
	GetGroup() *string
	SetInstanceId(v string) *DescribeFrInstancesShrinkRequest
	GetInstanceId() *string
	SetInstanceTag(v string) *DescribeFrInstancesShrinkRequest
	GetInstanceTag() *string
	SetNbid(v string) *DescribeFrInstancesShrinkRequest
	GetNbid() *string
	SetPageNum(v int32) *DescribeFrInstancesShrinkRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeFrInstancesShrinkRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeFrInstancesShrinkRequest
	GetProductCode() *string
	SetSortField(v string) *DescribeFrInstancesShrinkRequest
	GetSortField() *string
	SetSortRule(v string) *DescribeFrInstancesShrinkRequest
	GetSortRule() *string
	SetSpec(v string) *DescribeFrInstancesShrinkRequest
	GetSpec() *string
	SetStartTime(v int64) *DescribeFrInstancesShrinkRequest
	GetStartTime() *int64
	SetStatus(v string) *DescribeFrInstancesShrinkRequest
	GetStatus() *string
	SetTemplateCode(v string) *DescribeFrInstancesShrinkRequest
	GetTemplateCode() *string
}

type DescribeFrInstancesShrinkRequest struct {
	// The capacity type.
	//
	// example:
	//
	// deadlineAcc
	CapacityType *string `json:"CapacityType,omitempty" xml:"CapacityType,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// slb_albcubag_dp_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The cycle type.
	//
	// example:
	//
	// dynamicMonth
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The enterprise and account list. If this parameter is empty, the current account is queried.
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1710604800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The resource dimension to query.
	//
	// example:
	//
	// cu
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The instance name.
	//
	// example:
	//
	// alb_cubag*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance label value of the resource plan.
	//
	// example:
	//
	// FR-***
	InstanceTag *string `json:"InstanceTag,omitempty" xml:"InstanceTag,omitempty"`
	// The primary marketplace ID. If this parameter is empty, the marketplace ID of the current user is used by default.
	//
	// example:
	//
	// 2684202000018
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product code.
	//
	// example:
	//
	// slb
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The sort field.
	//
	// example:
	//
	// startTime
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The sorting rule.
	//
	// example:
	//
	// asc
	SortRule *string `json:"SortRule,omitempty" xml:"SortRule,omitempty"`
	// The specification.
	//
	// example:
	//
	// *
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1678939035000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The resource status.
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template code.
	//
	// example:
	//
	// slb_albcubag*******
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s DescribeFrInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesShrinkRequest) GetCapacityType() *string {
	return s.CapacityType
}

func (s *DescribeFrInstancesShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeFrInstancesShrinkRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *DescribeFrInstancesShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *DescribeFrInstancesShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFrInstancesShrinkRequest) GetGroup() *string {
	return s.Group
}

func (s *DescribeFrInstancesShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFrInstancesShrinkRequest) GetInstanceTag() *string {
	return s.InstanceTag
}

func (s *DescribeFrInstancesShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeFrInstancesShrinkRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeFrInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFrInstancesShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeFrInstancesShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *DescribeFrInstancesShrinkRequest) GetSortRule() *string {
	return s.SortRule
}

func (s *DescribeFrInstancesShrinkRequest) GetSpec() *string {
	return s.Spec
}

func (s *DescribeFrInstancesShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFrInstancesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeFrInstancesShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeFrInstancesShrinkRequest) SetCapacityType(v string) *DescribeFrInstancesShrinkRequest {
	s.CapacityType = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetCommodityCode(v string) *DescribeFrInstancesShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetCycleType(v string) *DescribeFrInstancesShrinkRequest {
	s.CycleType = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeFrInstancesShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetEndTime(v int64) *DescribeFrInstancesShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetGroup(v string) *DescribeFrInstancesShrinkRequest {
	s.Group = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetInstanceId(v string) *DescribeFrInstancesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetInstanceTag(v string) *DescribeFrInstancesShrinkRequest {
	s.InstanceTag = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetNbid(v string) *DescribeFrInstancesShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetPageNum(v int32) *DescribeFrInstancesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetPageSize(v int32) *DescribeFrInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetProductCode(v string) *DescribeFrInstancesShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetSortField(v string) *DescribeFrInstancesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetSortRule(v string) *DescribeFrInstancesShrinkRequest {
	s.SortRule = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetSpec(v string) *DescribeFrInstancesShrinkRequest {
	s.Spec = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetStartTime(v int64) *DescribeFrInstancesShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetStatus(v string) *DescribeFrInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) SetTemplateCode(v string) *DescribeFrInstancesShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *DescribeFrInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
