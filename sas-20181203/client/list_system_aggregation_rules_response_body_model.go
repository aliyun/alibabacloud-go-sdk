// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemAggregationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationList(v []*ListSystemAggregationRulesResponseBodyAggregationList) *ListSystemAggregationRulesResponseBody
	GetAggregationList() []*ListSystemAggregationRulesResponseBodyAggregationList
	SetPageInfo(v *ListSystemAggregationRulesResponseBodyPageInfo) *ListSystemAggregationRulesResponseBody
	GetPageInfo() *ListSystemAggregationRulesResponseBodyPageInfo
	SetRequestId(v string) *ListSystemAggregationRulesResponseBody
	GetRequestId() *string
}

type ListSystemAggregationRulesResponseBody struct {
	// An array that consists of the details about the aggregation types.
	AggregationList []*ListSystemAggregationRulesResponseBodyAggregationList `json:"AggregationList,omitempty" xml:"AggregationList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListSystemAggregationRulesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6800B790-B10A-5C2F-BEB3-F1D5CE61****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSystemAggregationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAggregationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemAggregationRulesResponseBody) GetAggregationList() []*ListSystemAggregationRulesResponseBodyAggregationList {
	return s.AggregationList
}

func (s *ListSystemAggregationRulesResponseBody) GetPageInfo() *ListSystemAggregationRulesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListSystemAggregationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemAggregationRulesResponseBody) SetAggregationList(v []*ListSystemAggregationRulesResponseBodyAggregationList) *ListSystemAggregationRulesResponseBody {
	s.AggregationList = v
	return s
}

func (s *ListSystemAggregationRulesResponseBody) SetPageInfo(v *ListSystemAggregationRulesResponseBodyPageInfo) *ListSystemAggregationRulesResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListSystemAggregationRulesResponseBody) SetRequestId(v string) *ListSystemAggregationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBody) Validate() error {
	if s.AggregationList != nil {
		for _, item := range s.AggregationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSystemAggregationRulesResponseBodyAggregationList struct {
	// The ID of the aggregation type.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the aggregation type.
	//
	// example:
	//
	// Remote control\\*\\*\\*\\*
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of rules that are of the aggregation type.
	//
	// example:
	//
	// 0
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s ListSystemAggregationRulesResponseBodyAggregationList) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAggregationRulesResponseBodyAggregationList) GoString() string {
	return s.String()
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) GetId() *int32 {
	return s.Id
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) GetName() *string {
	return s.Name
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) SetId(v int32) *ListSystemAggregationRulesResponseBodyAggregationList {
	s.Id = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) SetName(v string) *ListSystemAggregationRulesResponseBodyAggregationList {
	s.Name = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) SetRuleCount(v int32) *ListSystemAggregationRulesResponseBodyAggregationList {
	s.RuleCount = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBodyAggregationList) Validate() error {
	return dara.Validate(s)
}

type ListSystemAggregationRulesResponseBodyPageInfo struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSystemAggregationRulesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSystemAggregationRulesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) SetCurrentPage(v int32) *ListSystemAggregationRulesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) SetPageSize(v int32) *ListSystemAggregationRulesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) SetTotalCount(v int32) *ListSystemAggregationRulesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSystemAggregationRulesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
