// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumOcuStatisticDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetRumOcuStatisticDataRequest
	GetEndTime() *int64
	SetFilter(v []*GetRumOcuStatisticDataRequestFilter) *GetRumOcuStatisticDataRequest
	GetFilter() []*GetRumOcuStatisticDataRequestFilter
	SetGroup(v []*string) *GetRumOcuStatisticDataRequest
	GetGroup() []*string
	SetPage(v int32) *GetRumOcuStatisticDataRequest
	GetPage() *int32
	SetPageSize(v int32) *GetRumOcuStatisticDataRequest
	GetPageSize() *int32
	SetQueryType(v string) *GetRumOcuStatisticDataRequest
	GetQueryType() *string
	SetRegionId(v string) *GetRumOcuStatisticDataRequest
	GetRegionId() *string
	SetStartTime(v int64) *GetRumOcuStatisticDataRequest
	GetStartTime() *int64
}

type GetRumOcuStatisticDataRequest struct {
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
	Filter []*GetRumOcuStatisticDataRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	Group []*string `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
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

func (s GetRumOcuStatisticDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumOcuStatisticDataRequest) GoString() string {
	return s.String()
}

func (s *GetRumOcuStatisticDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetRumOcuStatisticDataRequest) GetFilter() []*GetRumOcuStatisticDataRequestFilter {
	return s.Filter
}

func (s *GetRumOcuStatisticDataRequest) GetGroup() []*string {
	return s.Group
}

func (s *GetRumOcuStatisticDataRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetRumOcuStatisticDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRumOcuStatisticDataRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *GetRumOcuStatisticDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumOcuStatisticDataRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetRumOcuStatisticDataRequest) SetEndTime(v int64) *GetRumOcuStatisticDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetFilter(v []*GetRumOcuStatisticDataRequestFilter) *GetRumOcuStatisticDataRequest {
	s.Filter = v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetGroup(v []*string) *GetRumOcuStatisticDataRequest {
	s.Group = v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetPage(v int32) *GetRumOcuStatisticDataRequest {
	s.Page = &v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetPageSize(v int32) *GetRumOcuStatisticDataRequest {
	s.PageSize = &v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetQueryType(v string) *GetRumOcuStatisticDataRequest {
	s.QueryType = &v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetRegionId(v string) *GetRumOcuStatisticDataRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumOcuStatisticDataRequest) SetStartTime(v int64) *GetRumOcuStatisticDataRequest {
	s.StartTime = &v
	return s
}

func (s *GetRumOcuStatisticDataRequest) Validate() error {
	return dara.Validate(s)
}

type GetRumOcuStatisticDataRequestFilter struct {
	// The key of the filter condition. Three types of filter conditions are provided:
	//
	// 	- Application name: pid (Note that the application name is displayed, but the application ID is actually specified)
	//
	// 	- Application type: siteType
	//
	// 	- Data type: dataType
	//
	// example:
	//
	// pid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the operator. Valid value: in.
	//
	// example:
	//
	// in
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The value of the filter condition. The value is a JSON array of strings.
	//
	// example:
	//
	// ["b590xxxxx@2dcbxxxxx9", "b590xxxxx@2dcbxxxxx8"]
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRumOcuStatisticDataRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetRumOcuStatisticDataRequestFilter) GoString() string {
	return s.String()
}

func (s *GetRumOcuStatisticDataRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetRumOcuStatisticDataRequestFilter) GetOpType() *string {
	return s.OpType
}

func (s *GetRumOcuStatisticDataRequestFilter) GetValue() interface{} {
	return s.Value
}

func (s *GetRumOcuStatisticDataRequestFilter) SetKey(v string) *GetRumOcuStatisticDataRequestFilter {
	s.Key = &v
	return s
}

func (s *GetRumOcuStatisticDataRequestFilter) SetOpType(v string) *GetRumOcuStatisticDataRequestFilter {
	s.OpType = &v
	return s
}

func (s *GetRumOcuStatisticDataRequestFilter) SetValue(v interface{}) *GetRumOcuStatisticDataRequestFilter {
	s.Value = v
	return s
}

func (s *GetRumOcuStatisticDataRequestFilter) Validate() error {
	return dara.Validate(s)
}
