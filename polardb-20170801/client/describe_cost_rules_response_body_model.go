// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeCostRulesResponseBodyItems) *DescribeCostRulesResponseBody
	GetItems() []*DescribeCostRulesResponseBodyItems
	SetPageNumber(v int32) *DescribeCostRulesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeCostRulesResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeCostRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCostRulesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeCostRulesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeCostRulesResponseBody struct {
	Items []*DescribeCostRulesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeCostRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesResponseBody) GetItems() []*DescribeCostRulesResponseBodyItems {
	return s.Items
}

func (s *DescribeCostRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCostRulesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeCostRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCostRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostRulesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeCostRulesResponseBody) SetItems(v []*DescribeCostRulesResponseBodyItems) *DescribeCostRulesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeCostRulesResponseBody) SetPageNumber(v int32) *DescribeCostRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetPageRecordCount(v int32) *DescribeCostRulesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetPageSize(v int32) *DescribeCostRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetRequestId(v string) *DescribeCostRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostRulesResponseBody) SetTotalRecordCount(v int32) *DescribeCostRulesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeCostRulesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostRulesResponseBodyItems struct {
	// example:
	//
	// 0
	CacheCostPointsPerMillion *string `json:"CacheCostPointsPerMillion,omitempty" xml:"CacheCostPointsPerMillion,omitempty"`
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// example:
	//
	// 2026-01-04T16:09:29+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2024-10-16 16:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 0
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// example:
	//
	// gpt-4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// ms-xxxxxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// example:
	//
	// 0
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
}

func (s DescribeCostRulesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesResponseBodyItems) GetCacheCostPointsPerMillion() *string {
	return s.CacheCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *DescribeCostRulesResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCostRulesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCostRulesResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeCostRulesResponseBodyItems) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) GetModel() *string {
	return s.Model
}

func (s *DescribeCostRulesResponseBodyItems) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *DescribeCostRulesResponseBodyItems) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *DescribeCostRulesResponseBodyItems) SetCacheCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.CacheCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetCostRuleId(v string) *DescribeCostRulesResponseBodyItems {
	s.CostRuleId = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetGmtCreated(v string) *DescribeCostRulesResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetGmtModified(v string) *DescribeCostRulesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetGwClusterId(v string) *DescribeCostRulesResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetInputCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetModel(v string) *DescribeCostRulesResponseBodyItems {
	s.Model = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetModelServiceId(v string) *DescribeCostRulesResponseBodyItems {
	s.ModelServiceId = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) SetOutputCostPointsPerMillion(v string) *DescribeCostRulesResponseBodyItems {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *DescribeCostRulesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
