// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetIssuesRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetIssuesRequest
	GetBizModule() *string
	SetFilter(v *GetIssuesRequestFilter) *GetIssuesRequest
	GetFilter() *GetIssuesRequestFilter
	SetName(v string) *GetIssuesRequest
	GetName() *string
	SetOrderBy(v string) *GetIssuesRequest
	GetOrderBy() *string
	SetOrderType(v string) *GetIssuesRequest
	GetOrderType() *string
	SetOs(v string) *GetIssuesRequest
	GetOs() *string
	SetPageIndex(v int32) *GetIssuesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetIssuesRequest
	GetPageSize() *int32
	SetStatus(v int32) *GetIssuesRequest
	GetStatus() *int32
	SetTimeRange(v *GetIssuesRequestTimeRange) *GetIssuesRequest
	GetTimeRange() *GetIssuesRequestTimeRange
}

type GetIssuesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// h5Resource
	BizModule *string                 `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	Filter    *GetIssuesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// GUI-TEST1711072832000
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// instances
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// h5
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FE_RUNNING
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is required.
	TimeRange *GetIssuesRequestTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
}

func (s GetIssuesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesRequest) GoString() string {
	return s.String()
}

func (s *GetIssuesRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetIssuesRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetIssuesRequest) GetFilter() *GetIssuesRequestFilter {
	return s.Filter
}

func (s *GetIssuesRequest) GetName() *string {
	return s.Name
}

func (s *GetIssuesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetIssuesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetIssuesRequest) GetOs() *string {
	return s.Os
}

func (s *GetIssuesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetIssuesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetIssuesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetIssuesRequest) GetTimeRange() *GetIssuesRequestTimeRange {
	return s.TimeRange
}

func (s *GetIssuesRequest) SetAppKey(v int64) *GetIssuesRequest {
	s.AppKey = &v
	return s
}

func (s *GetIssuesRequest) SetBizModule(v string) *GetIssuesRequest {
	s.BizModule = &v
	return s
}

func (s *GetIssuesRequest) SetFilter(v *GetIssuesRequestFilter) *GetIssuesRequest {
	s.Filter = v
	return s
}

func (s *GetIssuesRequest) SetName(v string) *GetIssuesRequest {
	s.Name = &v
	return s
}

func (s *GetIssuesRequest) SetOrderBy(v string) *GetIssuesRequest {
	s.OrderBy = &v
	return s
}

func (s *GetIssuesRequest) SetOrderType(v string) *GetIssuesRequest {
	s.OrderType = &v
	return s
}

func (s *GetIssuesRequest) SetOs(v string) *GetIssuesRequest {
	s.Os = &v
	return s
}

func (s *GetIssuesRequest) SetPageIndex(v int32) *GetIssuesRequest {
	s.PageIndex = &v
	return s
}

func (s *GetIssuesRequest) SetPageSize(v int32) *GetIssuesRequest {
	s.PageSize = &v
	return s
}

func (s *GetIssuesRequest) SetStatus(v int32) *GetIssuesRequest {
	s.Status = &v
	return s
}

func (s *GetIssuesRequest) SetTimeRange(v *GetIssuesRequestTimeRange) *GetIssuesRequest {
	s.TimeRange = v
	return s
}

func (s *GetIssuesRequest) Validate() error {
	return dara.Validate(s)
}

type GetIssuesRequestFilter struct {
	// example:
	//
	// erConfig
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// BeginWith
	Operator   *string       `json:"Operator,omitempty" xml:"Operator,omitempty"`
	SubFilters []*string     `json:"SubFilters,omitempty" xml:"SubFilters,omitempty" type:"Repeated"`
	Values     []interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetIssuesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesRequestFilter) GoString() string {
	return s.String()
}

func (s *GetIssuesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetIssuesRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *GetIssuesRequestFilter) GetSubFilters() []*string {
	return s.SubFilters
}

func (s *GetIssuesRequestFilter) GetValues() []interface{} {
	return s.Values
}

func (s *GetIssuesRequestFilter) SetKey(v string) *GetIssuesRequestFilter {
	s.Key = &v
	return s
}

func (s *GetIssuesRequestFilter) SetOperator(v string) *GetIssuesRequestFilter {
	s.Operator = &v
	return s
}

func (s *GetIssuesRequestFilter) SetSubFilters(v []*string) *GetIssuesRequestFilter {
	s.SubFilters = v
	return s
}

func (s *GetIssuesRequestFilter) SetValues(v []interface{}) *GetIssuesRequestFilter {
	s.Values = v
	return s
}

func (s *GetIssuesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type GetIssuesRequestTimeRange struct {
	// example:
	//
	// 2024-09-04T02:15:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// Host
	Granularity *int32 `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// example:
	//
	// day
	GranularityUnit *string `json:"GranularityUnit,omitempty" xml:"GranularityUnit,omitempty"`
	// example:
	//
	// 2024-11-05T16:00:00Z
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetIssuesRequestTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesRequestTimeRange) GoString() string {
	return s.String()
}

func (s *GetIssuesRequestTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetIssuesRequestTimeRange) GetGranularity() *int32 {
	return s.Granularity
}

func (s *GetIssuesRequestTimeRange) GetGranularityUnit() *string {
	return s.GranularityUnit
}

func (s *GetIssuesRequestTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetIssuesRequestTimeRange) SetEndTime(v int64) *GetIssuesRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetIssuesRequestTimeRange) SetGranularity(v int32) *GetIssuesRequestTimeRange {
	s.Granularity = &v
	return s
}

func (s *GetIssuesRequestTimeRange) SetGranularityUnit(v string) *GetIssuesRequestTimeRange {
	s.GranularityUnit = &v
	return s
}

func (s *GetIssuesRequestTimeRange) SetStartTime(v int64) *GetIssuesRequestTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetIssuesRequestTimeRange) Validate() error {
	return dara.Validate(s)
}
