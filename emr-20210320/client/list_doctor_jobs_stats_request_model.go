// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorJobsStatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorJobsStatsRequest
	GetClusterId() *string
	SetEndRange(v *ListDoctorJobsStatsRequestEndRange) *ListDoctorJobsStatsRequest
	GetEndRange() *ListDoctorJobsStatsRequestEndRange
	SetGroupBy(v []*string) *ListDoctorJobsStatsRequest
	GetGroupBy() []*string
	SetMaxResults(v int32) *ListDoctorJobsStatsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorJobsStatsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorJobsStatsRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorJobsStatsRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorJobsStatsRequest
	GetRegionId() *string
	SetStartRange(v *ListDoctorJobsStatsRequestStartRange) *ListDoctorJobsStatsRequest
	GetStartRange() *ListDoctorJobsStatsRequestStartRange
}

type ListDoctorJobsStatsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The range of end time. You can filter jobs whose end time falls within the specified time range.
	EndRange *ListDoctorJobsStatsRequestEndRange `json:"EndRange,omitempty" xml:"EndRange,omitempty" type:"Struct"`
	// The fields that are used for grouping data.
	//
	// Currently, only the first value is used for grouping data. Combinations of multiple values will be supported in the future.
	//
	// example:
	//
	// null
	GroupBy []*string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field that you use to sort the query results. Valid values:
	//
	// 	- vcoreSeconds: the aggregated number of vCPUs that are allocated to the job multiplied by the number of seconds the job has been running
	//
	// 	- memSeconds: the aggregated amount of memory that is allocated to the job multiplied by the number of seconds the job has been running
	//
	// example:
	//
	// vcoreSeconds
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- ASC: in ascending order
	//
	// 	- DESC: in descending order
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The range of start time. You can filter jobs whose start time falls within the specified time range.
	StartRange *ListDoctorJobsStatsRequestStartRange `json:"StartRange,omitempty" xml:"StartRange,omitempty" type:"Struct"`
}

func (s ListDoctorJobsStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorJobsStatsRequest) GetEndRange() *ListDoctorJobsStatsRequestEndRange {
	return s.EndRange
}

func (s *ListDoctorJobsStatsRequest) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *ListDoctorJobsStatsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorJobsStatsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorJobsStatsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorJobsStatsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorJobsStatsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorJobsStatsRequest) GetStartRange() *ListDoctorJobsStatsRequestStartRange {
	return s.StartRange
}

func (s *ListDoctorJobsStatsRequest) SetClusterId(v string) *ListDoctorJobsStatsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetEndRange(v *ListDoctorJobsStatsRequestEndRange) *ListDoctorJobsStatsRequest {
	s.EndRange = v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetGroupBy(v []*string) *ListDoctorJobsStatsRequest {
	s.GroupBy = v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetMaxResults(v int32) *ListDoctorJobsStatsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetNextToken(v string) *ListDoctorJobsStatsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetOrderBy(v string) *ListDoctorJobsStatsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetOrderType(v string) *ListDoctorJobsStatsRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetRegionId(v string) *ListDoctorJobsStatsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorJobsStatsRequest) SetStartRange(v *ListDoctorJobsStatsRequestStartRange) *ListDoctorJobsStatsRequest {
	s.StartRange = v
	return s
}

func (s *ListDoctorJobsStatsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsStatsRequestEndRange struct {
	// The end of the time range during which jobs ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1680019200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range during which jobs ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1675180800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDoctorJobsStatsRequestEndRange) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsRequestEndRange) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsRequestEndRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDoctorJobsStatsRequestEndRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDoctorJobsStatsRequestEndRange) SetEndTime(v int64) *ListDoctorJobsStatsRequestEndRange {
	s.EndTime = &v
	return s
}

func (s *ListDoctorJobsStatsRequestEndRange) SetStartTime(v int64) *ListDoctorJobsStatsRequestEndRange {
	s.StartTime = &v
	return s
}

func (s *ListDoctorJobsStatsRequestEndRange) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsStatsRequestStartRange struct {
	// The end of the time range during which jobs were submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1666406820000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range during which jobs were submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1679036826987
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDoctorJobsStatsRequestStartRange) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsRequestStartRange) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsRequestStartRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDoctorJobsStatsRequestStartRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDoctorJobsStatsRequestStartRange) SetEndTime(v int64) *ListDoctorJobsStatsRequestStartRange {
	s.EndTime = &v
	return s
}

func (s *ListDoctorJobsStatsRequestStartRange) SetStartTime(v int64) *ListDoctorJobsStatsRequestStartRange {
	s.StartTime = &v
	return s
}

func (s *ListDoctorJobsStatsRequestStartRange) Validate() error {
	return dara.Validate(s)
}
