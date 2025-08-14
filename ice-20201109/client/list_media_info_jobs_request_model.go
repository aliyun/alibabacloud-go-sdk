// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaInfoJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListMediaInfoJobsRequest
	GetEndOfCreateTime() *string
	SetJobId(v string) *ListMediaInfoJobsRequest
	GetJobId() *string
	SetNextPageToken(v string) *ListMediaInfoJobsRequest
	GetNextPageToken() *string
	SetOrderBy(v string) *ListMediaInfoJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListMediaInfoJobsRequest
	GetPageSize() *int32
	SetStartOfCreateTime(v string) *ListMediaInfoJobsRequest
	GetStartOfCreateTime() *string
	SetStatus(v string) *ListMediaInfoJobsRequest
	GetStatus() *string
}

type ListMediaInfoJobsRequest struct {
	// The end of the time range during which the jobs to be queried were created. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-07-15T00:00:00Z
	EndOfCreateTime *string `json:"EndOfCreateTime,omitempty" xml:"EndOfCreateTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 7b38a5d86f1e47838927b6e7ccb11cbe
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// ab4802364a2e49208c99efab82dfa8e8
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The order that you use to sort the query results. Valid values:
	//
	// 	- CreateTimeDesc: sorts the query results by creation time in descending order.
	//
	// 	- CreateTimeAsc: sorts the query results by creation time in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range during which the jobs to be queried were created. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-07-01T00:00:00Z
	StartOfCreateTime *string `json:"StartOfCreateTime,omitempty" xml:"StartOfCreateTime,omitempty"`
	// The state of the job. Valid values:
	//
	// 	- Init: The job is submitted.
	//
	// 	- Success: The job is successful.
	//
	// 	- Fail: The job failed.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaInfoJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaInfoJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaInfoJobsRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListMediaInfoJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListMediaInfoJobsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListMediaInfoJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMediaInfoJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMediaInfoJobsRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListMediaInfoJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMediaInfoJobsRequest) SetEndOfCreateTime(v string) *ListMediaInfoJobsRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListMediaInfoJobsRequest) SetJobId(v string) *ListMediaInfoJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListMediaInfoJobsRequest) SetNextPageToken(v string) *ListMediaInfoJobsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListMediaInfoJobsRequest) SetOrderBy(v string) *ListMediaInfoJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMediaInfoJobsRequest) SetPageSize(v int32) *ListMediaInfoJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMediaInfoJobsRequest) SetStartOfCreateTime(v string) *ListMediaInfoJobsRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListMediaInfoJobsRequest) SetStatus(v string) *ListMediaInfoJobsRequest {
	s.Status = &v
	return s
}

func (s *ListMediaInfoJobsRequest) Validate() error {
	return dara.Validate(s)
}
