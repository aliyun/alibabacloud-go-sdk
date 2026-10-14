// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOptimizeRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOptimizeRunsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOptimizeRunsResponseBody
	GetNextToken() *string
	SetOptimizeRuns(v []*ListOptimizeRunsResponseBodyOptimizeRuns) *ListOptimizeRunsResponseBody
	GetOptimizeRuns() []*ListOptimizeRunsResponseBodyOptimizeRuns
	SetRequestId(v string) *ListOptimizeRunsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOptimizeRunsResponseBody
	GetTotalCount() *int32
}

type ListOptimizeRunsResponseBody struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of run instances.
	OptimizeRuns []*ListOptimizeRunsResponseBodyOptimizeRuns `json:"optimizeRuns,omitempty" xml:"optimizeRuns,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListOptimizeRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOptimizeRunsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOptimizeRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOptimizeRunsResponseBody) GetOptimizeRuns() []*ListOptimizeRunsResponseBodyOptimizeRuns {
	return s.OptimizeRuns
}

func (s *ListOptimizeRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOptimizeRunsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOptimizeRunsResponseBody) SetMaxResults(v int32) *ListOptimizeRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOptimizeRunsResponseBody) SetNextToken(v string) *ListOptimizeRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOptimizeRunsResponseBody) SetOptimizeRuns(v []*ListOptimizeRunsResponseBodyOptimizeRuns) *ListOptimizeRunsResponseBody {
	s.OptimizeRuns = v
	return s
}

func (s *ListOptimizeRunsResponseBody) SetRequestId(v string) *ListOptimizeRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOptimizeRunsResponseBody) SetTotalCount(v int32) *ListOptimizeRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOptimizeRunsResponseBody) Validate() error {
	if s.OptimizeRuns != nil {
		for _, item := range s.OptimizeRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOptimizeRunsResponseBodyOptimizeRuns struct {
	// The number of executed attempts.
	//
	// example:
	//
	// 1
	AttemptCount *int32 `json:"attemptCount,omitempty" xml:"attemptCount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-12T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The current execution stage.
	//
	// example:
	//
	// attribution
	CurrentStage *string `json:"currentStage,omitempty" xml:"currentStage,omitempty"`
	// The error code.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The optimize run failed.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The execution summary and finite-state information. To obtain complete result references, frozen outputs, and publish receipts, call GetOptimizeRun.
	Extra map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	// The completion time.
	//
	// example:
	//
	// 2026-08-12T10:01:00Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// The maximum number of attempts.
	//
	// example:
	//
	// 3
	MaxAttempts *int32 `json:"maxAttempts,omitempty" xml:"maxAttempts,omitempty"`
	// The ID of the run instance.
	//
	// example:
	//
	// or-20260812-0001
	OptimizeRunId *string `json:"optimizeRunId,omitempty" xml:"optimizeRunId,omitempty"`
	// The name of the optimization task.
	//
	// example:
	//
	// efficiency-analysis
	OptimizeTaskName *string `json:"optimizeTaskName,omitempty" xml:"optimizeTaskName,omitempty"`
	// The current execution progress. The fields are defined by the task type.
	Progress map[string]interface{} `json:"progress,omitempty" xml:"progress,omitempty"`
	// The initial start time.
	//
	// example:
	//
	// 2026-08-12T10:00:01Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The execution status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of analysis trajectories.
	//
	// example:
	//
	// 100
	TrajectoryCount *int64 `json:"trajectoryCount,omitempty" xml:"trajectoryCount,omitempty"`
	// The trigger type.
	//
	// example:
	//
	// Scheduled
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// The snapshot of the task type.
	//
	// example:
	//
	// EfficiencyAnalysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The end time of the analysis window in UNIX seconds.
	//
	// example:
	//
	// 1786432690000
	WindowEndTime *int64 `json:"windowEndTime,omitempty" xml:"windowEndTime,omitempty"`
	// The start time of the analysis window in UNIX seconds.
	//
	// example:
	//
	// 1786432688000
	WindowStartTime *int64 `json:"windowStartTime,omitempty" xml:"windowStartTime,omitempty"`
}

func (s ListOptimizeRunsResponseBodyOptimizeRuns) String() string {
	return dara.Prettify(s)
}

func (s ListOptimizeRunsResponseBodyOptimizeRuns) GoString() string {
	return s.String()
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetAttemptCount() *int32 {
	return s.AttemptCount
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetMaxAttempts() *int32 {
	return s.MaxAttempts
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetOptimizeRunId() *string {
	return s.OptimizeRunId
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetOptimizeTaskName() *string {
	return s.OptimizeTaskName
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetProgress() map[string]interface{} {
	return s.Progress
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetStatus() *string {
	return s.Status
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetTrajectoryCount() *int64 {
	return s.TrajectoryCount
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetType() *string {
	return s.Type
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetWindowEndTime() *int64 {
	return s.WindowEndTime
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) GetWindowStartTime() *int64 {
	return s.WindowStartTime
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetAttemptCount(v int32) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.AttemptCount = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetCreateTime(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.CreateTime = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetCurrentStage(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.CurrentStage = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetErrorCode(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.ErrorCode = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetErrorMessage(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.ErrorMessage = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetExtra(v map[string]interface{}) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.Extra = v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetFinishTime(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.FinishTime = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetMaxAttempts(v int32) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.MaxAttempts = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetOptimizeRunId(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.OptimizeRunId = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetOptimizeTaskName(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.OptimizeTaskName = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetProgress(v map[string]interface{}) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.Progress = v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetStartTime(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.StartTime = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetStatus(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.Status = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetTrajectoryCount(v int64) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.TrajectoryCount = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetTriggerType(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.TriggerType = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetType(v string) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.Type = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetWindowEndTime(v int64) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.WindowEndTime = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) SetWindowStartTime(v int64) *ListOptimizeRunsResponseBodyOptimizeRuns {
	s.WindowStartTime = &v
	return s
}

func (s *ListOptimizeRunsResponseBodyOptimizeRuns) Validate() error {
	return dara.Validate(s)
}
