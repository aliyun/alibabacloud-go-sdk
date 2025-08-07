// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIssuesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetIssuesShrinkRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetIssuesShrinkRequest
	GetBizModule() *string
	SetFilterShrink(v string) *GetIssuesShrinkRequest
	GetFilterShrink() *string
	SetName(v string) *GetIssuesShrinkRequest
	GetName() *string
	SetOrderBy(v string) *GetIssuesShrinkRequest
	GetOrderBy() *string
	SetOrderType(v string) *GetIssuesShrinkRequest
	GetOrderType() *string
	SetOs(v string) *GetIssuesShrinkRequest
	GetOs() *string
	SetPageIndex(v int32) *GetIssuesShrinkRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetIssuesShrinkRequest
	GetPageSize() *int32
	SetStatus(v int32) *GetIssuesShrinkRequest
	GetStatus() *int32
	SetTimeRange(v *GetIssuesShrinkRequestTimeRange) *GetIssuesShrinkRequest
	GetTimeRange() *GetIssuesShrinkRequestTimeRange
}

type GetIssuesShrinkRequest struct {
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
	BizModule    *string `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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
	TimeRange *GetIssuesShrinkRequestTimeRange `json:"TimeRange,omitempty" xml:"TimeRange,omitempty" type:"Struct"`
}

func (s GetIssuesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetIssuesShrinkRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetIssuesShrinkRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetIssuesShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetIssuesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *GetIssuesShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetIssuesShrinkRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetIssuesShrinkRequest) GetOs() *string {
	return s.Os
}

func (s *GetIssuesShrinkRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetIssuesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetIssuesShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetIssuesShrinkRequest) GetTimeRange() *GetIssuesShrinkRequestTimeRange {
	return s.TimeRange
}

func (s *GetIssuesShrinkRequest) SetAppKey(v int64) *GetIssuesShrinkRequest {
	s.AppKey = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetBizModule(v string) *GetIssuesShrinkRequest {
	s.BizModule = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetFilterShrink(v string) *GetIssuesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetName(v string) *GetIssuesShrinkRequest {
	s.Name = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetOrderBy(v string) *GetIssuesShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetOrderType(v string) *GetIssuesShrinkRequest {
	s.OrderType = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetOs(v string) *GetIssuesShrinkRequest {
	s.Os = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetPageIndex(v int32) *GetIssuesShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetPageSize(v int32) *GetIssuesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetStatus(v int32) *GetIssuesShrinkRequest {
	s.Status = &v
	return s
}

func (s *GetIssuesShrinkRequest) SetTimeRange(v *GetIssuesShrinkRequestTimeRange) *GetIssuesShrinkRequest {
	s.TimeRange = v
	return s
}

func (s *GetIssuesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type GetIssuesShrinkRequestTimeRange struct {
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

func (s GetIssuesShrinkRequestTimeRange) String() string {
	return dara.Prettify(s)
}

func (s GetIssuesShrinkRequestTimeRange) GoString() string {
	return s.String()
}

func (s *GetIssuesShrinkRequestTimeRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetIssuesShrinkRequestTimeRange) GetGranularity() *int32 {
	return s.Granularity
}

func (s *GetIssuesShrinkRequestTimeRange) GetGranularityUnit() *string {
	return s.GranularityUnit
}

func (s *GetIssuesShrinkRequestTimeRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetIssuesShrinkRequestTimeRange) SetEndTime(v int64) *GetIssuesShrinkRequestTimeRange {
	s.EndTime = &v
	return s
}

func (s *GetIssuesShrinkRequestTimeRange) SetGranularity(v int32) *GetIssuesShrinkRequestTimeRange {
	s.Granularity = &v
	return s
}

func (s *GetIssuesShrinkRequestTimeRange) SetGranularityUnit(v string) *GetIssuesShrinkRequestTimeRange {
	s.GranularityUnit = &v
	return s
}

func (s *GetIssuesShrinkRequestTimeRange) SetStartTime(v int64) *GetIssuesShrinkRequestTimeRange {
	s.StartTime = &v
	return s
}

func (s *GetIssuesShrinkRequestTimeRange) Validate() error {
	return dara.Validate(s)
}
