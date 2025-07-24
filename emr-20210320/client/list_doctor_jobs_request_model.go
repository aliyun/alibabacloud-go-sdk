// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v []*string) *ListDoctorJobsRequest
	GetAppIds() []*string
	SetClusterId(v string) *ListDoctorJobsRequest
	GetClusterId() *string
	SetEndRange(v *ListDoctorJobsRequestEndRange) *ListDoctorJobsRequest
	GetEndRange() *ListDoctorJobsRequestEndRange
	SetMaxResults(v int32) *ListDoctorJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorJobsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorJobsRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorJobsRequest
	GetOrderType() *string
	SetQueues(v []*string) *ListDoctorJobsRequest
	GetQueues() []*string
	SetRegionId(v string) *ListDoctorJobsRequest
	GetRegionId() *string
	SetStartRange(v *ListDoctorJobsRequestStartRange) *ListDoctorJobsRequest
	GetStartRange() *ListDoctorJobsRequestStartRange
	SetTypes(v []*string) *ListDoctorJobsRequest
	GetTypes() []*string
	SetUsers(v []*string) *ListDoctorJobsRequest
	GetUsers() []*string
}

type ListDoctorJobsRequest struct {
	// The IDs of the jobs that are submitted to YARN.
	//
	// example:
	//
	// null
	AppIds []*string `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The range of end time. You can filter jobs whose end time falls within the specified time range.
	EndRange *ListDoctorJobsRequestEndRange `json:"EndRange,omitempty" xml:"EndRange,omitempty" type:"Struct"`
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
	// 	- ASC: the ascending order
	//
	// 	- DESC: the descending order
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The YARN queues to which the jobs are submitted.
	//
	// example:
	//
	// null
	Queues []*string `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The range of start time. You can filter jobs whose start time falls within the specified time range.
	StartRange *ListDoctorJobsRequestStartRange `json:"StartRange,omitempty" xml:"StartRange,omitempty" type:"Struct"`
	// The YARN engines to which the jobs are submitted.
	//
	// example:
	//
	// null
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	// The users who submit the jobs.
	//
	// example:
	//
	// null
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListDoctorJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsRequest) GetAppIds() []*string {
	return s.AppIds
}

func (s *ListDoctorJobsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorJobsRequest) GetEndRange() *ListDoctorJobsRequestEndRange {
	return s.EndRange
}

func (s *ListDoctorJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorJobsRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorJobsRequest) GetQueues() []*string {
	return s.Queues
}

func (s *ListDoctorJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorJobsRequest) GetStartRange() *ListDoctorJobsRequestStartRange {
	return s.StartRange
}

func (s *ListDoctorJobsRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListDoctorJobsRequest) GetUsers() []*string {
	return s.Users
}

func (s *ListDoctorJobsRequest) SetAppIds(v []*string) *ListDoctorJobsRequest {
	s.AppIds = v
	return s
}

func (s *ListDoctorJobsRequest) SetClusterId(v string) *ListDoctorJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorJobsRequest) SetEndRange(v *ListDoctorJobsRequestEndRange) *ListDoctorJobsRequest {
	s.EndRange = v
	return s
}

func (s *ListDoctorJobsRequest) SetMaxResults(v int32) *ListDoctorJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorJobsRequest) SetNextToken(v string) *ListDoctorJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorJobsRequest) SetOrderBy(v string) *ListDoctorJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorJobsRequest) SetOrderType(v string) *ListDoctorJobsRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorJobsRequest) SetQueues(v []*string) *ListDoctorJobsRequest {
	s.Queues = v
	return s
}

func (s *ListDoctorJobsRequest) SetRegionId(v string) *ListDoctorJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorJobsRequest) SetStartRange(v *ListDoctorJobsRequestStartRange) *ListDoctorJobsRequest {
	s.StartRange = v
	return s
}

func (s *ListDoctorJobsRequest) SetTypes(v []*string) *ListDoctorJobsRequest {
	s.Types = v
	return s
}

func (s *ListDoctorJobsRequest) SetUsers(v []*string) *ListDoctorJobsRequest {
	s.Users = v
	return s
}

func (s *ListDoctorJobsRequest) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsRequestEndRange struct {
	// The end of the time range during which jobs ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1666865137099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range during which jobs ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1679135111960
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDoctorJobsRequestEndRange) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsRequestEndRange) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsRequestEndRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDoctorJobsRequestEndRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDoctorJobsRequestEndRange) SetEndTime(v int64) *ListDoctorJobsRequestEndRange {
	s.EndTime = &v
	return s
}

func (s *ListDoctorJobsRequestEndRange) SetStartTime(v int64) *ListDoctorJobsRequestEndRange {
	s.StartTime = &v
	return s
}

func (s *ListDoctorJobsRequestEndRange) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsRequestStartRange struct {
	// The end of the time range during which jobs were submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1683340662020
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range during which jobs were submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1683340662016
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDoctorJobsRequestStartRange) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsRequestStartRange) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsRequestStartRange) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDoctorJobsRequestStartRange) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDoctorJobsRequestStartRange) SetEndTime(v int64) *ListDoctorJobsRequestStartRange {
	s.EndTime = &v
	return s
}

func (s *ListDoctorJobsRequestStartRange) SetStartTime(v int64) *ListDoctorJobsRequestStartRange {
	s.StartTime = &v
	return s
}

func (s *ListDoctorJobsRequestStartRange) Validate() error {
	return dara.Validate(s)
}
