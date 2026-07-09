// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *GetEvaluationRunResponseBody
	GetCreatedAt() *int64
	SetDataEndTime(v int64) *GetEvaluationRunResponseBody
	GetDataEndTime() *int64
	SetDataStartTime(v int64) *GetEvaluationRunResponseBody
	GetDataStartTime() *int64
	SetEvaluatorProgress(v []*GetEvaluationRunResponseBodyEvaluatorProgress) *GetEvaluationRunResponseBody
	GetEvaluatorProgress() []*GetEvaluationRunResponseBodyEvaluatorProgress
	SetEvaluators(v string) *GetEvaluationRunResponseBody
	GetEvaluators() *string
	SetFailedCount(v int32) *GetEvaluationRunResponseBody
	GetFailedCount() *int32
	SetRequestId(v string) *GetEvaluationRunResponseBody
	GetRequestId() *string
	SetRunId(v string) *GetEvaluationRunResponseBody
	GetRunId() *string
	SetRunName(v string) *GetEvaluationRunResponseBody
	GetRunName() *string
	SetRunType(v string) *GetEvaluationRunResponseBody
	GetRunType() *string
	SetStatus(v string) *GetEvaluationRunResponseBody
	GetStatus() *string
	SetSuccessCount(v int32) *GetEvaluationRunResponseBody
	GetSuccessCount() *int32
	SetTaskId(v string) *GetEvaluationRunResponseBody
	GetTaskId() *string
	SetTotalCount(v int32) *GetEvaluationRunResponseBody
	GetTotalCount() *int32
	SetUpdatedAt(v int64) *GetEvaluationRunResponseBody
	GetUpdatedAt() *int64
}

type GetEvaluationRunResponseBody struct {
	// The creation time, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The end time of the data window for the run, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782902400
	DataEndTime *int64 `json:"dataEndTime,omitempty" xml:"dataEndTime,omitempty"`
	// The start time of the data window for the run, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1782816000
	DataStartTime *int64 `json:"dataStartTime,omitempty" xml:"dataStartTime,omitempty"`
	// The list of progress details by evaluator.
	//
	// example:
	//
	// [{"evaluatorName":"Builtin.agent_task_completion","totalCount":100,"successCount":96,"failedCount":4}]
	EvaluatorProgress []*GetEvaluationRunResponseBodyEvaluatorProgress `json:"evaluatorProgress,omitempty" xml:"evaluatorProgress,omitempty" type:"Repeated"`
	// The evaluator configuration snapshot at the time the run was created, in JSON string format.
	//
	// example:
	//
	// [{"evaluatorRef":"Builtin.agent_task_completion"}]
	Evaluators *string `json:"evaluators,omitempty" xml:"evaluators,omitempty"`
	// The number of failed entries.
	//
	// example:
	//
	// 4
	FailedCount *int32 `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s GetEvaluationRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetEvaluationRunResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetEvaluationRunResponseBody) GetDataEndTime() *int64 {
	return s.DataEndTime
}

func (s *GetEvaluationRunResponseBody) GetDataStartTime() *int64 {
	return s.DataStartTime
}

func (s *GetEvaluationRunResponseBody) GetEvaluatorProgress() []*GetEvaluationRunResponseBodyEvaluatorProgress {
	return s.EvaluatorProgress
}

func (s *GetEvaluationRunResponseBody) GetEvaluators() *string {
	return s.Evaluators
}

func (s *GetEvaluationRunResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *GetEvaluationRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEvaluationRunResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetEvaluationRunResponseBody) GetRunName() *string {
	return s.RunName
}

func (s *GetEvaluationRunResponseBody) GetRunType() *string {
	return s.RunType
}

func (s *GetEvaluationRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEvaluationRunResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetEvaluationRunResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEvaluationRunResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetEvaluationRunResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetEvaluationRunResponseBody) SetCreatedAt(v int64) *GetEvaluationRunResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetDataEndTime(v int64) *GetEvaluationRunResponseBody {
	s.DataEndTime = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetDataStartTime(v int64) *GetEvaluationRunResponseBody {
	s.DataStartTime = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetEvaluatorProgress(v []*GetEvaluationRunResponseBodyEvaluatorProgress) *GetEvaluationRunResponseBody {
	s.EvaluatorProgress = v
	return s
}

func (s *GetEvaluationRunResponseBody) SetEvaluators(v string) *GetEvaluationRunResponseBody {
	s.Evaluators = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetFailedCount(v int32) *GetEvaluationRunResponseBody {
	s.FailedCount = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetRequestId(v string) *GetEvaluationRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetRunId(v string) *GetEvaluationRunResponseBody {
	s.RunId = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetRunName(v string) *GetEvaluationRunResponseBody {
	s.RunName = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetRunType(v string) *GetEvaluationRunResponseBody {
	s.RunType = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetStatus(v string) *GetEvaluationRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetSuccessCount(v int32) *GetEvaluationRunResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetTaskId(v string) *GetEvaluationRunResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetTotalCount(v int32) *GetEvaluationRunResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetEvaluationRunResponseBody) SetUpdatedAt(v int64) *GetEvaluationRunResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetEvaluationRunResponseBody) Validate() error {
	if s.EvaluatorProgress != nil {
		for _, item := range s.EvaluatorProgress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEvaluationRunResponseBodyEvaluatorProgress struct {
	// The evaluator name.
	//
	// example:
	//
	// Builtin.agent_task_completion
	EvaluatorName *string `json:"evaluatorName,omitempty" xml:"evaluatorName,omitempty"`
	// The number of failed entries for this evaluator.
	//
	// example:
	//
	// 4
	FailedCount *int32 `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	// The number of successful entries for this evaluator.
	//
	// example:
	//
	// 96
	SuccessCount *int32 `json:"successCount,omitempty" xml:"successCount,omitempty"`
	// The total number of entries for this evaluator.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetEvaluationRunResponseBodyEvaluatorProgress) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationRunResponseBodyEvaluatorProgress) GoString() string {
	return s.String()
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) GetEvaluatorName() *string {
	return s.EvaluatorName
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) SetEvaluatorName(v string) *GetEvaluationRunResponseBodyEvaluatorProgress {
	s.EvaluatorName = &v
	return s
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) SetFailedCount(v int32) *GetEvaluationRunResponseBodyEvaluatorProgress {
	s.FailedCount = &v
	return s
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) SetSuccessCount(v int32) *GetEvaluationRunResponseBodyEvaluatorProgress {
	s.SuccessCount = &v
	return s
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) SetTotalCount(v int32) *GetEvaluationRunResponseBodyEvaluatorProgress {
	s.TotalCount = &v
	return s
}

func (s *GetEvaluationRunResponseBodyEvaluatorProgress) Validate() error {
	return dara.Validate(s)
}
