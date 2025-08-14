// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaProducingJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListMediaProducingJobsRequest
	GetEndTime() *string
	SetJobType(v string) *ListMediaProducingJobsRequest
	GetJobType() *string
	SetKeyword(v string) *ListMediaProducingJobsRequest
	GetKeyword() *string
	SetMasterJobId(v string) *ListMediaProducingJobsRequest
	GetMasterJobId() *string
	SetMaxResults(v int32) *ListMediaProducingJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMediaProducingJobsRequest
	GetNextToken() *string
	SetProjectId(v string) *ListMediaProducingJobsRequest
	GetProjectId() *string
	SetSortBy(v string) *ListMediaProducingJobsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListMediaProducingJobsRequest
	GetStartTime() *string
	SetStatus(v string) *ListMediaProducingJobsRequest
	GetStatus() *string
}

type ListMediaProducingJobsRequest struct {
	// The end of the time range to query. The maximum time range between EndTime and StartTime cannot exceed 30 days. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-02-02T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The job type.
	//
	// Valid values:
	//
	// 	- LiveEditingJob: live editing job.
	//
	// 	- EditingJob: regular template-based editing job
	//
	// 	- VETemplateJob: advanced template-based editing job.
	//
	// example:
	//
	// EditingJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The search keyword. For example, you can use a job ID as the keyword to search for jobs.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the quick video production job. If this parameter is specified, the subjobs of the quick video production job are queried.
	//
	// example:
	//
	// ******8750b54e3c976a47da6f******
	MasterJobId *string `json:"MasterJobId,omitempty" xml:"MasterJobId,omitempty"`
	// The maximum number of entries to return.
	//
	// Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 8EqYpQbZ6Eh7+Zz8DxVYoQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ******927cfb53d05b96c1bfe1******
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job state.
	//
	// Valid values:
	//
	// 	- Init: The job is initialized.
	//
	// 	- Failed: The job failed.
	//
	// 	- Success: The job is successful.
	//
	// 	- Processing: The job is in progress.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaProducingJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaProducingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMediaProducingJobsRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListMediaProducingJobsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMediaProducingJobsRequest) GetMasterJobId() *string {
	return s.MasterJobId
}

func (s *ListMediaProducingJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaProducingJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaProducingJobsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListMediaProducingJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMediaProducingJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMediaProducingJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMediaProducingJobsRequest) SetEndTime(v string) *ListMediaProducingJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetJobType(v string) *ListMediaProducingJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetKeyword(v string) *ListMediaProducingJobsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetMasterJobId(v string) *ListMediaProducingJobsRequest {
	s.MasterJobId = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetMaxResults(v int32) *ListMediaProducingJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetNextToken(v string) *ListMediaProducingJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetProjectId(v string) *ListMediaProducingJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetSortBy(v string) *ListMediaProducingJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetStartTime(v string) *ListMediaProducingJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMediaProducingJobsRequest) SetStatus(v string) *ListMediaProducingJobsRequest {
	s.Status = &v
	return s
}

func (s *ListMediaProducingJobsRequest) Validate() error {
	return dara.Validate(s)
}
