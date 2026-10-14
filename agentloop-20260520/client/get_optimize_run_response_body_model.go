// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptimizeRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttemptCount(v int32) *GetOptimizeRunResponseBody
	GetAttemptCount() *int32
	SetCreateTime(v string) *GetOptimizeRunResponseBody
	GetCreateTime() *string
	SetCurrentStage(v string) *GetOptimizeRunResponseBody
	GetCurrentStage() *string
	SetErrorCode(v string) *GetOptimizeRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOptimizeRunResponseBody
	GetErrorMessage() *string
	SetExtra(v map[string]interface{}) *GetOptimizeRunResponseBody
	GetExtra() map[string]interface{}
	SetFinishTime(v string) *GetOptimizeRunResponseBody
	GetFinishTime() *string
	SetMaxAttempts(v int32) *GetOptimizeRunResponseBody
	GetMaxAttempts() *int32
	SetOptimizeRunId(v string) *GetOptimizeRunResponseBody
	GetOptimizeRunId() *string
	SetOptimizeTaskName(v string) *GetOptimizeRunResponseBody
	GetOptimizeTaskName() *string
	SetProgress(v map[string]interface{}) *GetOptimizeRunResponseBody
	GetProgress() map[string]interface{}
	SetRequestId(v string) *GetOptimizeRunResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetOptimizeRunResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetOptimizeRunResponseBody
	GetStatus() *string
	SetTrajectoryCount(v int64) *GetOptimizeRunResponseBody
	GetTrajectoryCount() *int64
	SetTriggerType(v string) *GetOptimizeRunResponseBody
	GetTriggerType() *string
	SetType(v string) *GetOptimizeRunResponseBody
	GetType() *string
	SetWindowEndTime(v int64) *GetOptimizeRunResponseBody
	GetWindowEndTime() *int64
	SetWindowStartTime(v int64) *GetOptimizeRunResponseBody
	GetWindowStartTime() *int64
}

type GetOptimizeRunResponseBody struct {
	// example:
	//
	// 1
	AttemptCount *int32 `json:"attemptCount,omitempty" xml:"attemptCount,omitempty"`
	// example:
	//
	// 2026-08-12T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// attribution
	CurrentStage *string `json:"currentStage,omitempty" xml:"currentStage,omitempty"`
	// example:
	//
	// InternalError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// The optimize run failed.
	ErrorMessage *string                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Extra        map[string]interface{} `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// 2026-08-12T10:01:00Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// example:
	//
	// 3
	MaxAttempts *int32 `json:"maxAttempts,omitempty" xml:"maxAttempts,omitempty"`
	// example:
	//
	// or-20260812-0001
	OptimizeRunId *string `json:"optimizeRunId,omitempty" xml:"optimizeRunId,omitempty"`
	// example:
	//
	// efficiency-analysis
	OptimizeTaskName *string                `json:"optimizeTaskName,omitempty" xml:"optimizeTaskName,omitempty"`
	Progress         map[string]interface{} `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2026-08-12T10:00:01Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 100
	TrajectoryCount *int64 `json:"trajectoryCount,omitempty" xml:"trajectoryCount,omitempty"`
	// example:
	//
	// Scheduled
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// example:
	//
	// EfficiencyAnalysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1786432690000
	WindowEndTime *int64 `json:"windowEndTime,omitempty" xml:"windowEndTime,omitempty"`
	// example:
	//
	// 1786432688000
	WindowStartTime *int64 `json:"windowStartTime,omitempty" xml:"windowStartTime,omitempty"`
}

func (s GetOptimizeRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOptimizeRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetOptimizeRunResponseBody) GetAttemptCount() *int32 {
	return s.AttemptCount
}

func (s *GetOptimizeRunResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOptimizeRunResponseBody) GetCurrentStage() *string {
	return s.CurrentStage
}

func (s *GetOptimizeRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOptimizeRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOptimizeRunResponseBody) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *GetOptimizeRunResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetOptimizeRunResponseBody) GetMaxAttempts() *int32 {
	return s.MaxAttempts
}

func (s *GetOptimizeRunResponseBody) GetOptimizeRunId() *string {
	return s.OptimizeRunId
}

func (s *GetOptimizeRunResponseBody) GetOptimizeTaskName() *string {
	return s.OptimizeTaskName
}

func (s *GetOptimizeRunResponseBody) GetProgress() map[string]interface{} {
	return s.Progress
}

func (s *GetOptimizeRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOptimizeRunResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOptimizeRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetOptimizeRunResponseBody) GetTrajectoryCount() *int64 {
	return s.TrajectoryCount
}

func (s *GetOptimizeRunResponseBody) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetOptimizeRunResponseBody) GetType() *string {
	return s.Type
}

func (s *GetOptimizeRunResponseBody) GetWindowEndTime() *int64 {
	return s.WindowEndTime
}

func (s *GetOptimizeRunResponseBody) GetWindowStartTime() *int64 {
	return s.WindowStartTime
}

func (s *GetOptimizeRunResponseBody) SetAttemptCount(v int32) *GetOptimizeRunResponseBody {
	s.AttemptCount = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetCreateTime(v string) *GetOptimizeRunResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetCurrentStage(v string) *GetOptimizeRunResponseBody {
	s.CurrentStage = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetErrorCode(v string) *GetOptimizeRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetErrorMessage(v string) *GetOptimizeRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetExtra(v map[string]interface{}) *GetOptimizeRunResponseBody {
	s.Extra = v
	return s
}

func (s *GetOptimizeRunResponseBody) SetFinishTime(v string) *GetOptimizeRunResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetMaxAttempts(v int32) *GetOptimizeRunResponseBody {
	s.MaxAttempts = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetOptimizeRunId(v string) *GetOptimizeRunResponseBody {
	s.OptimizeRunId = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetOptimizeTaskName(v string) *GetOptimizeRunResponseBody {
	s.OptimizeTaskName = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetProgress(v map[string]interface{}) *GetOptimizeRunResponseBody {
	s.Progress = v
	return s
}

func (s *GetOptimizeRunResponseBody) SetRequestId(v string) *GetOptimizeRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetStartTime(v string) *GetOptimizeRunResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetStatus(v string) *GetOptimizeRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetTrajectoryCount(v int64) *GetOptimizeRunResponseBody {
	s.TrajectoryCount = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetTriggerType(v string) *GetOptimizeRunResponseBody {
	s.TriggerType = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetType(v string) *GetOptimizeRunResponseBody {
	s.Type = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetWindowEndTime(v int64) *GetOptimizeRunResponseBody {
	s.WindowEndTime = &v
	return s
}

func (s *GetOptimizeRunResponseBody) SetWindowStartTime(v int64) *GetOptimizeRunResponseBody {
	s.WindowStartTime = &v
	return s
}

func (s *GetOptimizeRunResponseBody) Validate() error {
	return dara.Validate(s)
}
