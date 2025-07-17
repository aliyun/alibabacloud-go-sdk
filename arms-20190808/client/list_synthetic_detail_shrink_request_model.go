// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyntheticDetailShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedFiltersShrink(v string) *ListSyntheticDetailShrinkRequest
	GetAdvancedFiltersShrink() *string
	SetCategory(v string) *ListSyntheticDetailShrinkRequest
	GetCategory() *string
	SetDetail(v string) *ListSyntheticDetailShrinkRequest
	GetDetail() *string
	SetEndTime(v int64) *ListSyntheticDetailShrinkRequest
	GetEndTime() *int64
	SetExactFiltersShrink(v string) *ListSyntheticDetailShrinkRequest
	GetExactFiltersShrink() *string
	SetFiltersShrink(v string) *ListSyntheticDetailShrinkRequest
	GetFiltersShrink() *string
	SetOrder(v string) *ListSyntheticDetailShrinkRequest
	GetOrder() *string
	SetOrderBy(v string) *ListSyntheticDetailShrinkRequest
	GetOrderBy() *string
	SetPage(v int32) *ListSyntheticDetailShrinkRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSyntheticDetailShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSyntheticDetailShrinkRequest
	GetRegionId() *string
	SetStartTime(v int64) *ListSyntheticDetailShrinkRequest
	GetStartTime() *int64
	SetSyntheticType(v int32) *ListSyntheticDetailShrinkRequest
	GetSyntheticType() *int32
}

type ListSyntheticDetailShrinkRequest struct {
	// An array of filter conditions. This parameter is required.
	//
	// 	- To query the list of synthetic test results, set this parameter in the following format: [{"Key":"taskType","OpType":"in","Value":[Task type]}].
	//
	// 	- To query the result details of a synthetic monitoring task, set this parameter in the following format: [{"Key":"dataId","OpType":"eq","Value":"dataId"}]. dataId is returned when you query the list of synthetic test results.
	AdvancedFiltersShrink *string `json:"AdvancedFilters,omitempty" xml:"AdvancedFilters,omitempty"`
	// The type of the results. Set the value to SYNTHETIC.
	//
	// example:
	//
	// SYNTHETIC
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The type of the list that contains the results. This parameter is required. Valid values:
	//
	// 	- ICMP_LIST
	//
	// 	- TCP_LIST
	//
	// 	- DNS_LIST
	//
	// 	- HTTP_LIST
	//
	// 	- WEBSITE_LIST
	//
	// 	- DOWNLOAD_LIST
	//
	// 	- ALL
	//
	// example:
	//
	// ICMP_LIST
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The timestamp of the end time of the query. Unit: milliseconds.
	//
	// example:
	//
	// 1684480557772
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A reserved field.
	ExactFiltersShrink *string `json:"ExactFilters,omitempty" xml:"ExactFilters,omitempty"`
	// The filter condition. This parameter is required.
	//
	// 	- To query the result of a synthetic monitoring task, set this parameter in the following format: {"taskId":"${taskId}"}.
	//
	// 	- To query the result details of a synthetic monitoring task, set this parameter in the following format: {"taskId":"${taskId}","dataId":"${dataId}"}.
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The order in which results are sorted. Valid values:
	//
	// - `ASC`: ascending order
	//
	// - `DESC`: descending order
	//
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The field based on which results are sorted. Set the value to timestamp.
	//
	// example:
	//
	// timestamp
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timestamp of the start time of the query. Unit: milliseconds.
	//
	// example:
	//
	// 1684110343127
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the synthetic test. Valid values: 1 and 2. 1 represents an immediate test, and 2 represents a scheduled test.
	//
	// example:
	//
	// 1
	SyntheticType *int32 `json:"SyntheticType,omitempty" xml:"SyntheticType,omitempty"`
}

func (s ListSyntheticDetailShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailShrinkRequest) GetAdvancedFiltersShrink() *string {
	return s.AdvancedFiltersShrink
}

func (s *ListSyntheticDetailShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ListSyntheticDetailShrinkRequest) GetDetail() *string {
	return s.Detail
}

func (s *ListSyntheticDetailShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSyntheticDetailShrinkRequest) GetExactFiltersShrink() *string {
	return s.ExactFiltersShrink
}

func (s *ListSyntheticDetailShrinkRequest) GetFiltersShrink() *string {
	return s.FiltersShrink
}

func (s *ListSyntheticDetailShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSyntheticDetailShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSyntheticDetailShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSyntheticDetailShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSyntheticDetailShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSyntheticDetailShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSyntheticDetailShrinkRequest) GetSyntheticType() *int32 {
	return s.SyntheticType
}

func (s *ListSyntheticDetailShrinkRequest) SetAdvancedFiltersShrink(v string) *ListSyntheticDetailShrinkRequest {
	s.AdvancedFiltersShrink = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetCategory(v string) *ListSyntheticDetailShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetDetail(v string) *ListSyntheticDetailShrinkRequest {
	s.Detail = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetEndTime(v int64) *ListSyntheticDetailShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetExactFiltersShrink(v string) *ListSyntheticDetailShrinkRequest {
	s.ExactFiltersShrink = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetFiltersShrink(v string) *ListSyntheticDetailShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetOrder(v string) *ListSyntheticDetailShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetOrderBy(v string) *ListSyntheticDetailShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetPage(v int32) *ListSyntheticDetailShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetPageSize(v int32) *ListSyntheticDetailShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetRegionId(v string) *ListSyntheticDetailShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetStartTime(v int64) *ListSyntheticDetailShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) SetSyntheticType(v int32) *ListSyntheticDetailShrinkRequest {
	s.SyntheticType = &v
	return s
}

func (s *ListSyntheticDetailShrinkRequest) Validate() error {
	return dara.Validate(s)
}
