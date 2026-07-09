// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpaceName(v string) *GetPipelineRunResponseBody
	GetAgentSpaceName() *string
	SetAttempt(v int32) *GetPipelineRunResponseBody
	GetAttempt() *int32
	SetErrorCode(v string) *GetPipelineRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineRunResponseBody
	GetErrorMessage() *string
	SetFinishTime(v string) *GetPipelineRunResponseBody
	GetFinishTime() *string
	SetFromTime(v int64) *GetPipelineRunResponseBody
	GetFromTime() *int64
	SetMaxAttempts(v int32) *GetPipelineRunResponseBody
	GetMaxAttempts() *int32
	SetNextRetryTime(v string) *GetPipelineRunResponseBody
	GetNextRetryTime() *string
	SetPipelineName(v string) *GetPipelineRunResponseBody
	GetPipelineName() *string
	SetRequestId(v string) *GetPipelineRunResponseBody
	GetRequestId() *string
	SetResults(v map[string]interface{}) *GetPipelineRunResponseBody
	GetResults() map[string]interface{}
	SetRunId(v string) *GetPipelineRunResponseBody
	GetRunId() *string
	SetStartTime(v string) *GetPipelineRunResponseBody
	GetStartTime() *string
	SetStats(v map[string]interface{}) *GetPipelineRunResponseBody
	GetStats() map[string]interface{}
	SetStatus(v string) *GetPipelineRunResponseBody
	GetStatus() *string
	SetToTime(v int64) *GetPipelineRunResponseBody
	GetToTime() *int64
	SetTriggerTime(v string) *GetPipelineRunResponseBody
	GetTriggerTime() *string
	SetTriggerType(v string) *GetPipelineRunResponseBody
	GetTriggerType() *string
}

type GetPipelineRunResponseBody struct {
	// The name of the AgentSpace to which the pipeline belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpaceName *string `json:"agentSpaceName,omitempty" xml:"agentSpaceName,omitempty"`
	// The current retry count.
	//
	// example:
	//
	// 0
	Attempt *int32 `json:"attempt,omitempty" xml:"attempt,omitempty"`
	// The semantic error code.
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The finish time, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:10.000Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// The start of the data window, in UNIX seconds.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxAttempts *int32 `json:"maxAttempts,omitempty" xml:"maxAttempts,omitempty"`
	// The next retry time, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:01:00.000Z
	NextRetryTime *string `json:"nextRetryTime,omitempty" xml:"nextRetryTime,omitempty"`
	// The name of the pipeline.
	//
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The inline run results. This field is returned only when output.inline is set to true at trigger time.
	Results map[string]interface{} `json:"results,omitempty" xml:"results,omitempty"`
	// Run Id
	//
	// example:
	//
	// run-20260101-0001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The execution start time, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:01.000Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The run statistics, including processedRows (number of processed rows), processedBytes (number of processed bytes), outputRows (number of output rows), outputBytes (number of output bytes), elapsedMs (elapsed time in milliseconds), cpuSec (CPU seconds), cpuCores (number of CPU cores), and tokenCount (number of tokens consumed).
	Stats map[string]interface{} `json:"stats,omitempty" xml:"stats,omitempty"`
	// The run status. Valid values:
	//
	// - Pending
	//
	// - Running
	//
	// - Succeeded
	//
	// - Failed
	//
	// - Cancelled
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The end of the data window, in UNIX seconds.
	//
	// example:
	//
	// 1735747200
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
	// The trigger time, in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:00.000Z
	TriggerTime *string `json:"triggerTime,omitempty" xml:"triggerTime,omitempty"`
	// The trigger type. Valid values:
	//
	// - Scheduled
	//
	// - Manual
	//
	// - RunOnce
	//
	// example:
	//
	// Scheduled
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBody) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *GetPipelineRunResponseBody) GetAttempt() *int32 {
	return s.Attempt
}

func (s *GetPipelineRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineRunResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetPipelineRunResponseBody) GetFromTime() *int64 {
	return s.FromTime
}

func (s *GetPipelineRunResponseBody) GetMaxAttempts() *int32 {
	return s.MaxAttempts
}

func (s *GetPipelineRunResponseBody) GetNextRetryTime() *string {
	return s.NextRetryTime
}

func (s *GetPipelineRunResponseBody) GetPipelineName() *string {
	return s.PipelineName
}

func (s *GetPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineRunResponseBody) GetResults() map[string]interface{} {
	return s.Results
}

func (s *GetPipelineRunResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetPipelineRunResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPipelineRunResponseBody) GetStats() map[string]interface{} {
	return s.Stats
}

func (s *GetPipelineRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineRunResponseBody) GetToTime() *int64 {
	return s.ToTime
}

func (s *GetPipelineRunResponseBody) GetTriggerTime() *string {
	return s.TriggerTime
}

func (s *GetPipelineRunResponseBody) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetPipelineRunResponseBody) SetAgentSpaceName(v string) *GetPipelineRunResponseBody {
	s.AgentSpaceName = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetAttempt(v int32) *GetPipelineRunResponseBody {
	s.Attempt = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetErrorCode(v string) *GetPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetErrorMessage(v string) *GetPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetFinishTime(v string) *GetPipelineRunResponseBody {
	s.FinishTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetFromTime(v int64) *GetPipelineRunResponseBody {
	s.FromTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetMaxAttempts(v int32) *GetPipelineRunResponseBody {
	s.MaxAttempts = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetNextRetryTime(v string) *GetPipelineRunResponseBody {
	s.NextRetryTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetPipelineName(v string) *GetPipelineRunResponseBody {
	s.PipelineName = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetRequestId(v string) *GetPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetResults(v map[string]interface{}) *GetPipelineRunResponseBody {
	s.Results = v
	return s
}

func (s *GetPipelineRunResponseBody) SetRunId(v string) *GetPipelineRunResponseBody {
	s.RunId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetStartTime(v string) *GetPipelineRunResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetStats(v map[string]interface{}) *GetPipelineRunResponseBody {
	s.Stats = v
	return s
}

func (s *GetPipelineRunResponseBody) SetStatus(v string) *GetPipelineRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetToTime(v int64) *GetPipelineRunResponseBody {
	s.ToTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetTriggerTime(v string) *GetPipelineRunResponseBody {
	s.TriggerTime = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetTriggerType(v string) *GetPipelineRunResponseBody {
	s.TriggerType = &v
	return s
}

func (s *GetPipelineRunResponseBody) Validate() error {
	return dara.Validate(s)
}
