// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobState(v string) *ListSmartJobsRequest
	GetJobState() *string
	SetJobType(v string) *ListSmartJobsRequest
	GetJobType() *string
	SetMaxResults(v int64) *ListSmartJobsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListSmartJobsRequest
	GetNextToken() *string
	SetPageNo(v int64) *ListSmartJobsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListSmartJobsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListSmartJobsRequest
	GetSortBy() *string
}

type ListSmartJobsRequest struct {
	// The job state.
	//
	// Valid values:
	//
	// 	- Finished: The job is complete.
	//
	// 	- Failed: The job failed.
	//
	// 	- Executing: The job is in progress.
	//
	// 	- Created: The job is created.
	//
	// example:
	//
	// Finished
	JobState *string `json:"JobState,omitempty" xml:"JobState,omitempty"`
	// The job type.
	//
	// Valid values:
	//
	// 	- ASR: automatic speech recognition(job) job.
	//
	// 	- DynamicChart: dynamic chart job.
	//
	// 	- VideoTranslation: video translation job.
	//
	// 	- TextToSpeech: intelligent audio production job.
	//
	// example:
	//
	// ASR
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The maximum number of entries to return.
	//
	// Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ****73f33c91-d59383e8280b****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting parameter. By default, the query results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- CreationTime:Asc: sorted by creation time in ascending order.
	//
	// 	- CreationTime:Desc: sorted by creation time in descending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListSmartJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSmartJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSmartJobsRequest) GetJobState() *string {
	return s.JobState
}

func (s *ListSmartJobsRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListSmartJobsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListSmartJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSmartJobsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListSmartJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSmartJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSmartJobsRequest) SetJobState(v string) *ListSmartJobsRequest {
	s.JobState = &v
	return s
}

func (s *ListSmartJobsRequest) SetJobType(v string) *ListSmartJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListSmartJobsRequest) SetMaxResults(v int64) *ListSmartJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSmartJobsRequest) SetNextToken(v string) *ListSmartJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSmartJobsRequest) SetPageNo(v int64) *ListSmartJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSmartJobsRequest) SetPageSize(v int64) *ListSmartJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmartJobsRequest) SetSortBy(v string) *ListSmartJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSmartJobsRequest) Validate() error {
	return dara.Validate(s)
}
