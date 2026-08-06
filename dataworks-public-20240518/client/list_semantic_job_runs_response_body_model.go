// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSemanticJobRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSemanticJobRunsResponseBodyData) *ListSemanticJobRunsResponseBody
	GetData() *ListSemanticJobRunsResponseBodyData
	SetRequestId(v string) *ListSemanticJobRunsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSemanticJobRunsResponseBody
	GetSuccess() *bool
}

type ListSemanticJobRunsResponseBody struct {
	// The paginated run record results. Use the JobRunId to download the results of a specific run, and use the ExecutorJobId to query details, retrieve logs, or stop a run.
	Data *ListSemanticJobRunsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSemanticJobRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSemanticJobRunsResponseBody) GetData() *ListSemanticJobRunsResponseBodyData {
	return s.Data
}

func (s *ListSemanticJobRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSemanticJobRunsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSemanticJobRunsResponseBody) SetData(v *ListSemanticJobRunsResponseBodyData) *ListSemanticJobRunsResponseBody {
	s.Data = v
	return s
}

func (s *ListSemanticJobRunsResponseBody) SetRequestId(v string) *ListSemanticJobRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSemanticJobRunsResponseBody) SetSuccess(v bool) *ListSemanticJobRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSemanticJobRunsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSemanticJobRunsResponseBodyData struct {
	// The list of run records.
	JobRuns []*ListSemanticJobRunsResponseBodyDataJobRuns `json:"JobRuns,omitempty" xml:"JobRuns,omitempty" type:"Repeated"`
	// The page number of the returned page, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page in the current response.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of run records that match the current job criteria.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSemanticJobRunsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobRunsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSemanticJobRunsResponseBodyData) GetJobRuns() []*ListSemanticJobRunsResponseBodyDataJobRuns {
	return s.JobRuns
}

func (s *ListSemanticJobRunsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSemanticJobRunsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSemanticJobRunsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSemanticJobRunsResponseBodyData) SetJobRuns(v []*ListSemanticJobRunsResponseBodyDataJobRuns) *ListSemanticJobRunsResponseBodyData {
	s.JobRuns = v
	return s
}

func (s *ListSemanticJobRunsResponseBodyData) SetPageNumber(v int32) *ListSemanticJobRunsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyData) SetPageSize(v int32) *ListSemanticJobRunsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyData) SetTotalCount(v int64) *ListSemanticJobRunsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyData) Validate() error {
	if s.JobRuns != nil {
		for _, item := range s.JobRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSemanticJobRunsResponseBodyDataJobRuns struct {
	// The executor job ID. Pass this value as the ExecutorJobId parameter to GetSemanticJobDetail, GetSemanticJobLog, or KillSemanticJob.
	//
	// example:
	//
	// exec-job-demo
	ExecutorJobId *string `json:"ExecutorJobId,omitempty" xml:"ExecutorJobId,omitempty"`
	// The time when the run record was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1700000000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The name of the job to which this run belongs. You can use this value to rerun the job, query run records, or download results.
	//
	// example:
	//
	// semantic-job-demo
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The semantic job run ID. Pass this value as the JobRunId parameter to DownloadSemanticResults to download the results of this run.
	//
	// example:
	//
	// 01H00000000000000000000000
	JobRunId *string `json:"JobRunId,omitempty" xml:"JobRunId,omitempty"`
	// The ID of the user who submitted this run.
	//
	// example:
	//
	// user-demo
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListSemanticJobRunsResponseBodyDataJobRuns) String() string {
	return dara.Prettify(s)
}

func (s ListSemanticJobRunsResponseBodyDataJobRuns) GoString() string {
	return s.String()
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) GetExecutorJobId() *string {
	return s.ExecutorJobId
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) GetJobName() *string {
	return s.JobName
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) GetJobRunId() *string {
	return s.JobRunId
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) GetUserId() *string {
	return s.UserId
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) SetExecutorJobId(v string) *ListSemanticJobRunsResponseBodyDataJobRuns {
	s.ExecutorJobId = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) SetGmtCreate(v int64) *ListSemanticJobRunsResponseBodyDataJobRuns {
	s.GmtCreate = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) SetJobName(v string) *ListSemanticJobRunsResponseBodyDataJobRuns {
	s.JobName = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) SetJobRunId(v string) *ListSemanticJobRunsResponseBodyDataJobRuns {
	s.JobRunId = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) SetUserId(v string) *ListSemanticJobRunsResponseBodyDataJobRuns {
	s.UserId = &v
	return s
}

func (s *ListSemanticJobRunsResponseBodyDataJobRuns) Validate() error {
	return dara.Validate(s)
}
