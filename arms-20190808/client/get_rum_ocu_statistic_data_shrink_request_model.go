// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumOcuStatisticDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetRumOcuStatisticDataShrinkRequest
	GetEndTime() *int64
	SetFilterShrink(v string) *GetRumOcuStatisticDataShrinkRequest
	GetFilterShrink() *string
	SetGroupShrink(v string) *GetRumOcuStatisticDataShrinkRequest
	GetGroupShrink() *string
	SetPage(v int32) *GetRumOcuStatisticDataShrinkRequest
	GetPage() *int32
	SetPageSize(v int32) *GetRumOcuStatisticDataShrinkRequest
	GetPageSize() *int32
	SetQueryType(v string) *GetRumOcuStatisticDataShrinkRequest
	GetQueryType() *string
	SetRegionId(v string) *GetRumOcuStatisticDataShrinkRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetRumOcuStatisticDataShrinkRequest
	GetStartTime() *int64
}

type GetRumOcuStatisticDataShrinkRequest struct {
	// The end of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1687849260000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter condition. Three types of filter conditions are provided:
	//
	// 	- Application name: pid (Note that the application name is displayed, but the application ID is actually specified)
	//
	// 	- Application type: siteType
	//
	// 	- Data type: dataType
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The grouping fields. Valid values:
	//
	// 	- siteType: The total number of OCUs is grouped by application type.
	//
	// 	- dataType: The total number of OCUs is grouped by data type.
	//
	// 	- pid: The total number of OCUs is grouped by application ID.
	//
	// 	- appName: The total number of OCUs is grouped by application name.
	//
	// 	- startTime: The total number of OCUs is grouped by start time.
	GroupShrink *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the query. To query non-time series data, set the value to INSTANT. To query time series data, set the value to TIME_SERIES.
	//
	// example:
	//
	// TIME_SERIES
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1600063200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetRumOcuStatisticDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumOcuStatisticDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetGroupShrink() *string {
	return s.GroupShrink
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumOcuStatisticDataShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetEndTime(v int64) *GetRumOcuStatisticDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetFilterShrink(v string) *GetRumOcuStatisticDataShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetGroupShrink(v string) *GetRumOcuStatisticDataShrinkRequest {
	s.GroupShrink = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetPage(v int32) *GetRumOcuStatisticDataShrinkRequest {
	s.Page = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetPageSize(v int32) *GetRumOcuStatisticDataShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetQueryType(v string) *GetRumOcuStatisticDataShrinkRequest {
	s.QueryType = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetRegionId(v string) *GetRumOcuStatisticDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) SetStartTime(v int64) *GetRumOcuStatisticDataShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetRumOcuStatisticDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
