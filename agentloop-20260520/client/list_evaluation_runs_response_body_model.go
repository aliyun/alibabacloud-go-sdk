// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvaluationRuns(v []*ListEvaluationRunsResponseBodyEvaluationRuns) *ListEvaluationRunsResponseBody
	GetEvaluationRuns() []*ListEvaluationRunsResponseBodyEvaluationRuns
	SetMaxResults(v int32) *ListEvaluationRunsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluationRunsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEvaluationRunsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListEvaluationRunsResponseBody
	GetTotalCount() *int64
}

type ListEvaluationRunsResponseBody struct {
	// The list of run summaries.
	//
	// example:
	//
	// [{"runId":"eval-run-4fd47f3d7e684e15b1d3d178c6a5b81a","runType":"backfill","status":"Running","totalCount":100}]
	EvaluationRuns []*ListEvaluationRunsResponseBodyEvaluationRuns `json:"evaluationRuns,omitempty" xml:"evaluationRuns,omitempty" type:"Repeated"`
	// The number of entries per page used in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more pages exist.
	//
	// example:
	//
	// eyJsYXN0SWQiOjEyMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of runs that match the filter conditions.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListEvaluationRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationRunsResponseBody) GetEvaluationRuns() []*ListEvaluationRunsResponseBodyEvaluationRuns {
	return s.EvaluationRuns
}

func (s *ListEvaluationRunsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluationRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluationRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluationRunsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEvaluationRunsResponseBody) SetEvaluationRuns(v []*ListEvaluationRunsResponseBodyEvaluationRuns) *ListEvaluationRunsResponseBody {
	s.EvaluationRuns = v
	return s
}

func (s *ListEvaluationRunsResponseBody) SetMaxResults(v int32) *ListEvaluationRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluationRunsResponseBody) SetNextToken(v string) *ListEvaluationRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationRunsResponseBody) SetRequestId(v string) *ListEvaluationRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationRunsResponseBody) SetTotalCount(v int64) *ListEvaluationRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEvaluationRunsResponseBody) Validate() error {
	if s.EvaluationRuns != nil {
		for _, item := range s.EvaluationRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationRunsResponseBodyEvaluationRuns struct {
	// The creation time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The end time of the data window for this run, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782902400
	DataEndTime *int64 `json:"dataEndTime,omitempty" xml:"dataEndTime,omitempty"`
	// The start time of the data window for this run, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	DataStartTime *int64 `json:"dataStartTime,omitempty" xml:"dataStartTime,omitempty"`
	// The number of failed entries.
	//
	// example:
	//
	// 4
	FailedCount *int32 `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	// The run ID.
	//
	// example:
	//
	// eval-run-4fd47f3d7e684e15b1d3d178c6a5b81a
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The run name.
	//
	// example:
	//
	// trace_task_completion_eval-backfill
	RunName *string `json:"runName,omitempty" xml:"runName,omitempty"`
	// The run type.
	//
	// example:
	//
	// backfill
	RunType *string `json:"runType,omitempty" xml:"runType,omitempty"`
	// The run status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of successful entries.
	//
	// example:
	//
	// 96
	SuccessCount *int32 `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// The evaluation task ID.
	//
	// example:
	//
	// eval-task-8b36f2e2b1f94f9c91ce7a4b0f6d9c25
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The total number of evaluation entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The update time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816600
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListEvaluationRunsResponseBodyEvaluationRuns) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationRunsResponseBodyEvaluationRuns) GoString() string {
	return s.String()
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetDataEndTime() *int64 {
	return s.DataEndTime
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetDataStartTime() *int64 {
	return s.DataStartTime
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetRunId() *string {
	return s.RunId
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetRunName() *string {
	return s.RunName
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetRunType() *string {
	return s.RunType
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetStatus() *string {
	return s.Status
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetTaskId() *string {
	return s.TaskId
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetCreatedAt(v int64) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.CreatedAt = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetDataEndTime(v int64) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.DataEndTime = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetDataStartTime(v int64) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.DataStartTime = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetFailedCount(v int32) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.FailedCount = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetRunId(v string) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.RunId = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetRunName(v string) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.RunName = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetRunType(v string) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.RunType = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetStatus(v string) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.Status = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetSuccessCount(v int32) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.SuccessCount = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetTaskId(v string) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.TaskId = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetTotalCount(v int32) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.TotalCount = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) SetUpdatedAt(v int64) *ListEvaluationRunsResponseBodyEvaluationRuns {
	s.UpdatedAt = &v
	return s
}

func (s *ListEvaluationRunsResponseBodyEvaluationRuns) Validate() error {
	return dara.Validate(s)
}
