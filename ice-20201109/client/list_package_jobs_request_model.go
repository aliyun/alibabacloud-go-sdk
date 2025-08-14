// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPackageJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListPackageJobsRequest
	GetEndOfCreateTime() *string
	SetJobId(v string) *ListPackageJobsRequest
	GetJobId() *string
	SetNextPageToken(v string) *ListPackageJobsRequest
	GetNextPageToken() *string
	SetOrderBy(v string) *ListPackageJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListPackageJobsRequest
	GetPageSize() *int32
	SetStartOfCreateTime(v string) *ListPackageJobsRequest
	GetStartOfCreateTime() *string
	SetStatus(v string) *ListPackageJobsRequest
	GetStatus() *string
}

type ListPackageJobsRequest struct {
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
	// 	- CreateTimeDesc: sorts the jobs by creation time in descending order.
	//
	// 	- CreateTimeAsc: sorts the jobs by creation time in ascending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Valid values: 0 to 100. Default value: 20.
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
	// The state of the job.
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

func (s ListPackageJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPackageJobsRequest) GoString() string {
	return s.String()
}

func (s *ListPackageJobsRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListPackageJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListPackageJobsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListPackageJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListPackageJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPackageJobsRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListPackageJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListPackageJobsRequest) SetEndOfCreateTime(v string) *ListPackageJobsRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListPackageJobsRequest) SetJobId(v string) *ListPackageJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListPackageJobsRequest) SetNextPageToken(v string) *ListPackageJobsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListPackageJobsRequest) SetOrderBy(v string) *ListPackageJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPackageJobsRequest) SetPageSize(v int32) *ListPackageJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPackageJobsRequest) SetStartOfCreateTime(v string) *ListPackageJobsRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListPackageJobsRequest) SetStatus(v string) *ListPackageJobsRequest {
	s.Status = &v
	return s
}

func (s *ListPackageJobsRequest) Validate() error {
	return dara.Validate(s)
}
