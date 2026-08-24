// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityType(v string) *DescribeFrInstancesRequest
	GetCapacityType() *string
	SetCommodityCode(v string) *DescribeFrInstancesRequest
	GetCommodityCode() *string
	SetCycleType(v string) *DescribeFrInstancesRequest
	GetCycleType() *string
	SetEcIdAccountIds(v []*DescribeFrInstancesRequestEcIdAccountIds) *DescribeFrInstancesRequest
	GetEcIdAccountIds() []*DescribeFrInstancesRequestEcIdAccountIds
	SetEndTime(v int64) *DescribeFrInstancesRequest
	GetEndTime() *int64
	SetGroup(v string) *DescribeFrInstancesRequest
	GetGroup() *string
	SetInstanceId(v string) *DescribeFrInstancesRequest
	GetInstanceId() *string
	SetInstanceTag(v string) *DescribeFrInstancesRequest
	GetInstanceTag() *string
	SetNbid(v string) *DescribeFrInstancesRequest
	GetNbid() *string
	SetPageNum(v int32) *DescribeFrInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeFrInstancesRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeFrInstancesRequest
	GetProductCode() *string
	SetSortField(v string) *DescribeFrInstancesRequest
	GetSortField() *string
	SetSortRule(v string) *DescribeFrInstancesRequest
	GetSortRule() *string
	SetSpec(v string) *DescribeFrInstancesRequest
	GetSpec() *string
	SetStartTime(v int64) *DescribeFrInstancesRequest
	GetStartTime() *int64
	SetStatus(v string) *DescribeFrInstancesRequest
	GetStatus() *string
	SetTemplateCode(v string) *DescribeFrInstancesRequest
	GetTemplateCode() *string
}

type DescribeFrInstancesRequest struct {
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
	// The enterprise and account list. If empty, the current account is queried.
	EcIdAccountIds []*DescribeFrInstancesRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// The end time.
	//
	// example:
	//
	// 1710604800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The resource dimension for the query.
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
	// The instance tag label value of the resource plan instance.
	//
	// example:
	//
	// FR-***
	InstanceTag *string `json:"InstanceTag,omitempty" xml:"InstanceTag,omitempty"`
	// The primary marketplace ID. If empty, the marketplace ID of the current user is used by default.
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
	// The collation for sorting.
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

func (s DescribeFrInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesRequest) GetCapacityType() *string {
	return s.CapacityType
}

func (s *DescribeFrInstancesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeFrInstancesRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *DescribeFrInstancesRequest) GetEcIdAccountIds() []*DescribeFrInstancesRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *DescribeFrInstancesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFrInstancesRequest) GetGroup() *string {
	return s.Group
}

func (s *DescribeFrInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFrInstancesRequest) GetInstanceTag() *string {
	return s.InstanceTag
}

func (s *DescribeFrInstancesRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeFrInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeFrInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFrInstancesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeFrInstancesRequest) GetSortField() *string {
	return s.SortField
}

func (s *DescribeFrInstancesRequest) GetSortRule() *string {
	return s.SortRule
}

func (s *DescribeFrInstancesRequest) GetSpec() *string {
	return s.Spec
}

func (s *DescribeFrInstancesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFrInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeFrInstancesRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeFrInstancesRequest) SetCapacityType(v string) *DescribeFrInstancesRequest {
	s.CapacityType = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetCommodityCode(v string) *DescribeFrInstancesRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetCycleType(v string) *DescribeFrInstancesRequest {
	s.CycleType = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetEcIdAccountIds(v []*DescribeFrInstancesRequestEcIdAccountIds) *DescribeFrInstancesRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *DescribeFrInstancesRequest) SetEndTime(v int64) *DescribeFrInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetGroup(v string) *DescribeFrInstancesRequest {
	s.Group = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetInstanceId(v string) *DescribeFrInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetInstanceTag(v string) *DescribeFrInstancesRequest {
	s.InstanceTag = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetNbid(v string) *DescribeFrInstancesRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetPageNum(v int32) *DescribeFrInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetPageSize(v int32) *DescribeFrInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetProductCode(v string) *DescribeFrInstancesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetSortField(v string) *DescribeFrInstancesRequest {
	s.SortField = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetSortRule(v string) *DescribeFrInstancesRequest {
	s.SortRule = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetSpec(v string) *DescribeFrInstancesRequest {
	s.Spec = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetStartTime(v int64) *DescribeFrInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetStatus(v string) *DescribeFrInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeFrInstancesRequest) SetTemplateCode(v string) *DescribeFrInstancesRequest {
	s.TemplateCode = &v
	return s
}

func (s *DescribeFrInstancesRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFrInstancesRequestEcIdAccountIds struct {
	// The account list to access. If empty, all accounts under the current entity ID are selected.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The enterprise entity ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s DescribeFrInstancesRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *DescribeFrInstancesRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *DescribeFrInstancesRequestEcIdAccountIds) SetAccountIds(v []*int64) *DescribeFrInstancesRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *DescribeFrInstancesRequestEcIdAccountIds) SetEcId(v string) *DescribeFrInstancesRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *DescribeFrInstancesRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
