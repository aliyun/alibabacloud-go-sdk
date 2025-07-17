// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyntheticDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedFilters(v []*ListSyntheticDetailRequestAdvancedFilters) *ListSyntheticDetailRequest
	GetAdvancedFilters() []*ListSyntheticDetailRequestAdvancedFilters
	SetCategory(v string) *ListSyntheticDetailRequest
	GetCategory() *string
	SetDetail(v string) *ListSyntheticDetailRequest
	GetDetail() *string
	SetEndTime(v int64) *ListSyntheticDetailRequest
	GetEndTime() *int64
	SetExactFilters(v []*ListSyntheticDetailRequestExactFilters) *ListSyntheticDetailRequest
	GetExactFilters() []*ListSyntheticDetailRequestExactFilters
	SetFilters(v map[string]*string) *ListSyntheticDetailRequest
	GetFilters() map[string]*string
	SetOrder(v string) *ListSyntheticDetailRequest
	GetOrder() *string
	SetOrderBy(v string) *ListSyntheticDetailRequest
	GetOrderBy() *string
	SetPage(v int32) *ListSyntheticDetailRequest
	GetPage() *int32
	SetPageSize(v int32) *ListSyntheticDetailRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListSyntheticDetailRequest
	GetRegionId() *string
	SetStartTime(v int64) *ListSyntheticDetailRequest
	GetStartTime() *int64
	SetSyntheticType(v int32) *ListSyntheticDetailRequest
	GetSyntheticType() *int32
}

type ListSyntheticDetailRequest struct {
	// An array of filter conditions. This parameter is required.
	//
	// 	- To query the list of synthetic test results, set this parameter in the following format: [{"Key":"taskType","OpType":"in","Value":[Task type]}].
	//
	// 	- To query the result details of a synthetic monitoring task, set this parameter in the following format: [{"Key":"dataId","OpType":"eq","Value":"dataId"}]. dataId is returned when you query the list of synthetic test results.
	AdvancedFilters []*ListSyntheticDetailRequestAdvancedFilters `json:"AdvancedFilters,omitempty" xml:"AdvancedFilters,omitempty" type:"Repeated"`
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
	ExactFilters []*ListSyntheticDetailRequestExactFilters `json:"ExactFilters,omitempty" xml:"ExactFilters,omitempty" type:"Repeated"`
	// The filter condition. This parameter is required.
	//
	// 	- To query the result of a synthetic monitoring task, set this parameter in the following format: {"taskId":"${taskId}"}.
	//
	// 	- To query the result details of a synthetic monitoring task, set this parameter in the following format: {"taskId":"${taskId}","dataId":"${dataId}"}.
	Filters map[string]*string `json:"Filters,omitempty" xml:"Filters,omitempty"`
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

func (s ListSyntheticDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailRequest) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailRequest) GetAdvancedFilters() []*ListSyntheticDetailRequestAdvancedFilters {
	return s.AdvancedFilters
}

func (s *ListSyntheticDetailRequest) GetCategory() *string {
	return s.Category
}

func (s *ListSyntheticDetailRequest) GetDetail() *string {
	return s.Detail
}

func (s *ListSyntheticDetailRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSyntheticDetailRequest) GetExactFilters() []*ListSyntheticDetailRequestExactFilters {
	return s.ExactFilters
}

func (s *ListSyntheticDetailRequest) GetFilters() map[string]*string {
	return s.Filters
}

func (s *ListSyntheticDetailRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSyntheticDetailRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSyntheticDetailRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListSyntheticDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSyntheticDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSyntheticDetailRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSyntheticDetailRequest) GetSyntheticType() *int32 {
	return s.SyntheticType
}

func (s *ListSyntheticDetailRequest) SetAdvancedFilters(v []*ListSyntheticDetailRequestAdvancedFilters) *ListSyntheticDetailRequest {
	s.AdvancedFilters = v
	return s
}

func (s *ListSyntheticDetailRequest) SetCategory(v string) *ListSyntheticDetailRequest {
	s.Category = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetDetail(v string) *ListSyntheticDetailRequest {
	s.Detail = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetEndTime(v int64) *ListSyntheticDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetExactFilters(v []*ListSyntheticDetailRequestExactFilters) *ListSyntheticDetailRequest {
	s.ExactFilters = v
	return s
}

func (s *ListSyntheticDetailRequest) SetFilters(v map[string]*string) *ListSyntheticDetailRequest {
	s.Filters = v
	return s
}

func (s *ListSyntheticDetailRequest) SetOrder(v string) *ListSyntheticDetailRequest {
	s.Order = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetOrderBy(v string) *ListSyntheticDetailRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetPage(v int32) *ListSyntheticDetailRequest {
	s.Page = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetPageSize(v int32) *ListSyntheticDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetRegionId(v string) *ListSyntheticDetailRequest {
	s.RegionId = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetStartTime(v int64) *ListSyntheticDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ListSyntheticDetailRequest) SetSyntheticType(v int32) *ListSyntheticDetailRequest {
	s.SyntheticType = &v
	return s
}

func (s *ListSyntheticDetailRequest) Validate() error {
	return dara.Validate(s)
}

type ListSyntheticDetailRequestAdvancedFilters struct {
	// The filter condition. The taskType and dataId fields are supported.
	//
	// 	- To query the list of synthetic test results, set the key to taskType.
	//
	// 	- To query the result details of a synthetic monitoring task, set the key to dataId.
	//
	// example:
	//
	// taskType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the filter condition. Valid values: eq and in. eq: equal to. in: include.
	//
	// example:
	//
	// eq
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The value of the filter condition. The type of the task. Valid values: 1: ICMP 2: TCP 3: DNS 4: HTTP 5: website speed measurement 6: file download
	//
	// example:
	//
	// 1
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSyntheticDetailRequestAdvancedFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailRequestAdvancedFilters) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailRequestAdvancedFilters) GetKey() *string {
	return s.Key
}

func (s *ListSyntheticDetailRequestAdvancedFilters) GetOpType() *string {
	return s.OpType
}

func (s *ListSyntheticDetailRequestAdvancedFilters) GetValue() interface{} {
	return s.Value
}

func (s *ListSyntheticDetailRequestAdvancedFilters) SetKey(v string) *ListSyntheticDetailRequestAdvancedFilters {
	s.Key = &v
	return s
}

func (s *ListSyntheticDetailRequestAdvancedFilters) SetOpType(v string) *ListSyntheticDetailRequestAdvancedFilters {
	s.OpType = &v
	return s
}

func (s *ListSyntheticDetailRequestAdvancedFilters) SetValue(v interface{}) *ListSyntheticDetailRequestAdvancedFilters {
	s.Value = v
	return s
}

func (s *ListSyntheticDetailRequestAdvancedFilters) Validate() error {
	return dara.Validate(s)
}

type ListSyntheticDetailRequestExactFilters struct {
	// A reserved field.
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// null
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// null
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSyntheticDetailRequestExactFilters) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailRequestExactFilters) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailRequestExactFilters) GetKey() *string {
	return s.Key
}

func (s *ListSyntheticDetailRequestExactFilters) GetOpType() *string {
	return s.OpType
}

func (s *ListSyntheticDetailRequestExactFilters) GetValue() interface{} {
	return s.Value
}

func (s *ListSyntheticDetailRequestExactFilters) SetKey(v string) *ListSyntheticDetailRequestExactFilters {
	s.Key = &v
	return s
}

func (s *ListSyntheticDetailRequestExactFilters) SetOpType(v string) *ListSyntheticDetailRequestExactFilters {
	s.OpType = &v
	return s
}

func (s *ListSyntheticDetailRequestExactFilters) SetValue(v interface{}) *ListSyntheticDetailRequestExactFilters {
	s.Value = v
	return s
}

func (s *ListSyntheticDetailRequestExactFilters) Validate() error {
	return dara.Validate(s)
}
