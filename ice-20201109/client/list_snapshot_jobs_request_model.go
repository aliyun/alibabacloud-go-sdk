// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndOfCreateTime(v string) *ListSnapshotJobsRequest
	GetEndOfCreateTime() *string
	SetJobId(v string) *ListSnapshotJobsRequest
	GetJobId() *string
	SetNextPageToken(v string) *ListSnapshotJobsRequest
	GetNextPageToken() *string
	SetOrderBy(v string) *ListSnapshotJobsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListSnapshotJobsRequest
	GetPageSize() *int32
	SetStartOfCreateTime(v string) *ListSnapshotJobsRequest
	GetStartOfCreateTime() *string
	SetStatus(v string) *ListSnapshotJobsRequest
	GetStatus() *string
}

type ListSnapshotJobsRequest struct {
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
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// The order that you use to sort the query results.
	//
	// 1.  CreateTimeDesc
	//
	// 2.  CreateTimeAsc
	//
	// Valid values:
	//
	// 	- CreateTimeDesc: sorts the jobs by creation time in descending order
	//
	// 	- CreateTimeAsc: sorts the jobs by creation time in ascending order.
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

func (s ListSnapshotJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotJobsRequest) GetEndOfCreateTime() *string {
	return s.EndOfCreateTime
}

func (s *ListSnapshotJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListSnapshotJobsRequest) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListSnapshotJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSnapshotJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSnapshotJobsRequest) GetStartOfCreateTime() *string {
	return s.StartOfCreateTime
}

func (s *ListSnapshotJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListSnapshotJobsRequest) SetEndOfCreateTime(v string) *ListSnapshotJobsRequest {
	s.EndOfCreateTime = &v
	return s
}

func (s *ListSnapshotJobsRequest) SetJobId(v string) *ListSnapshotJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListSnapshotJobsRequest) SetNextPageToken(v string) *ListSnapshotJobsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListSnapshotJobsRequest) SetOrderBy(v string) *ListSnapshotJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSnapshotJobsRequest) SetPageSize(v int32) *ListSnapshotJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotJobsRequest) SetStartOfCreateTime(v string) *ListSnapshotJobsRequest {
	s.StartOfCreateTime = &v
	return s
}

func (s *ListSnapshotJobsRequest) SetStatus(v string) *ListSnapshotJobsRequest {
	s.Status = &v
	return s
}

func (s *ListSnapshotJobsRequest) Validate() error {
	return dara.Validate(s)
}
