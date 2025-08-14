// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicImageJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListDynamicImageJobsRequest
	GetEndOfCreateTime() *string
	SetJobId(v string) *ListDynamicImageJobsRequest
	GetJobId() *string
	SetNextPageToken(v string) *ListDynamicImageJobsRequest
	GetNextPageToken() *string
	SetOrderBy(v string) *ListDynamicImageJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListDynamicImageJobsRequest
	GetPageSize() *int32
	SetStartOfCreateTime(v string) *ListDynamicImageJobsRequest
	GetStartOfCreateTime() *string
	SetStatus(v string) *ListDynamicImageJobsRequest
	GetStatus() *string
}

type ListDynamicImageJobsRequest struct {
	// The end of the time range during which the jobs to be queried were created.
	//
	// example:
	//
	// 2022-07-14T00:00:00Z
	EndOfCreateTime *string `json:"EndOfCreateTime,omitempty" xml:"EndOfCreateTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// ****cdb3e74639973036bc84****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// cdb3e74639973036bc84
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The order that you use to sort the query results. Valid values:
	//
	// 1.  CreateTimeAsc: sorts the jobs by creation time in ascending order.
	//
	// 2.  CreateTimeDesc: sorts the jobs by creation time in descending order.
	//
	// example:
	//
	// CreateTimeDesc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range during which the jobs to be queried were created.
	//
	// example:
	//
	// 2022-07-12T00:00:00Z
	StartOfCreateTime *string `json:"StartOfCreateTime,omitempty" xml:"StartOfCreateTime,omitempty"`
	// The state of the job.
	//
	// Valid values:
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

func (s ListDynamicImageJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicImageJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicImageJobsRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListDynamicImageJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListDynamicImageJobsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListDynamicImageJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDynamicImageJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDynamicImageJobsRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListDynamicImageJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDynamicImageJobsRequest) SetEndOfCreateTime(v string) *ListDynamicImageJobsRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListDynamicImageJobsRequest) SetJobId(v string) *ListDynamicImageJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListDynamicImageJobsRequest) SetNextPageToken(v string) *ListDynamicImageJobsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListDynamicImageJobsRequest) SetOrderBy(v string) *ListDynamicImageJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDynamicImageJobsRequest) SetPageSize(v int32) *ListDynamicImageJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDynamicImageJobsRequest) SetStartOfCreateTime(v string) *ListDynamicImageJobsRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListDynamicImageJobsRequest) SetStatus(v string) *ListDynamicImageJobsRequest {
	s.Status = &v
	return s
}

func (s *ListDynamicImageJobsRequest) Validate() error {
	return dara.Validate(s)
}
