// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchMediaProducingJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListBatchMediaProducingJobsRequest
	GetEndTime() *string
	SetJobId(v string) *ListBatchMediaProducingJobsRequest
	GetJobId() *string
	SetJobType(v string) *ListBatchMediaProducingJobsRequest
	GetJobType() *string
	SetMaxResults(v int32) *ListBatchMediaProducingJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBatchMediaProducingJobsRequest
	GetNextToken() *string
	SetSortBy(v string) *ListBatchMediaProducingJobsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListBatchMediaProducingJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListBatchMediaProducingJobsRequest
	GetStatus() *string
}

type ListBatchMediaProducingJobsRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-05T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the quick video production job.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job type.
	//
	// Valid values:
	//
	// 	- Script: script-based editing job that mixes media assets.
	//
	// 	- Smart_Mix: intelligent editing job that mixes media assets.
	//
	// example:
	//
	// Script
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// mRZkKAovub0xWVfH14he4Q==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sorting parameter. Valid values:
	//
	// 	- desc (default): sorted by creation time in descending order.
	//
	// 	- asc: sorted by creation time in ascending order.
	//
	// <!---->
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-02-02T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state.
	//
	// Valid values:
	//
	// 	- Finished
	//
	// 	- Init
	//
	// 	- Failed
	//
	// 	- Processing
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBatchMediaProducingJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchMediaProducingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListBatchMediaProducingJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBatchMediaProducingJobsRequest) GetJobId() *string {
	return s.JobId
}

func (s *ListBatchMediaProducingJobsRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListBatchMediaProducingJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBatchMediaProducingJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBatchMediaProducingJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListBatchMediaProducingJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBatchMediaProducingJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListBatchMediaProducingJobsRequest) SetEndTime(v string) *ListBatchMediaProducingJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetJobId(v string) *ListBatchMediaProducingJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetJobType(v string) *ListBatchMediaProducingJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetMaxResults(v int32) *ListBatchMediaProducingJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetNextToken(v string) *ListBatchMediaProducingJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetSortBy(v string) *ListBatchMediaProducingJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetStartTime(v string) *ListBatchMediaProducingJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) SetStatus(v string) *ListBatchMediaProducingJobsRequest {
	s.Status = &v
	return s
}

func (s *ListBatchMediaProducingJobsRequest) Validate() error {
	return dara.Validate(s)
}
