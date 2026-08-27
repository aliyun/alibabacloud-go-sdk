// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillRunResponseBody
	GetCode() *string
	SetCreatedAt(v string) *GetSkillRunResponseBody
	GetCreatedAt() *string
	SetErrorCode(v string) *GetSkillRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSkillRunResponseBody
	GetErrorMessage() *string
	SetFinishedAt(v string) *GetSkillRunResponseBody
	GetFinishedAt() *string
	SetLogs(v []map[string]interface{}) *GetSkillRunResponseBody
	GetLogs() []map[string]interface{}
	SetMessage(v string) *GetSkillRunResponseBody
	GetMessage() *string
	SetProgress(v int64) *GetSkillRunResponseBody
	GetProgress() *int64
	SetProgressMessage(v string) *GetSkillRunResponseBody
	GetProgressMessage() *string
	SetRequestId(v string) *GetSkillRunResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *GetSkillRunResponseBody
	GetResult() map[string]interface{}
	SetRunId(v string) *GetSkillRunResponseBody
	GetRunId() *string
	SetSkillCode(v string) *GetSkillRunResponseBody
	GetSkillCode() *string
	SetSkillName(v string) *GetSkillRunResponseBody
	GetSkillName() *string
	SetStartedAt(v string) *GetSkillRunResponseBody
	GetStartedAt() *string
	SetStatus(v string) *GetSkillRunResponseBody
	GetStatus() *string
	SetUsage(v map[string]interface{}) *GetSkillRunResponseBody
	GetUsage() map[string]interface{}
}

type GetSkillRunResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The task creation time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The error code. This parameter is returned only when the status is Failed.
	//
	// example:
	//
	// string_value
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error description. This parameter is returned only when the status is Failed.
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The task end time in ISO 8601 format. This parameter has a value only in desired states (Succeeded, Failed, or Cancelled).
	//
	// example:
	//
	// string_value
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The execution log list. This parameter is returned only when IncludeLogs is set to true.
	Logs []map[string]interface{} `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The progress percentage. This parameter is meaningful only when the status is Running.
	//
	// example:
	//
	// 1
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The progress description.
	//
	// example:
	//
	// string_value
	ProgressMessage *string `json:"progressMessage,omitempty" xml:"progressMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The execution result. This parameter is returned only when the status is Succeeded. It contains a content list.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// The asynchronous task ID.
	//
	// example:
	//
	// exampleRunId
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The skill code.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The skill name.
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The task execution start time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The execution status. Valid values: Running, Succeeded, Failed, and Cancelled.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The LLM token usage statistics. This parameter is returned only when the status is Succeeded.
	Usage map[string]interface{} `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s GetSkillRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillRunResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillRunResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetSkillRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSkillRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSkillRunResponseBody) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetSkillRunResponseBody) GetLogs() []map[string]interface{} {
	return s.Logs
}

func (s *GetSkillRunResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillRunResponseBody) GetProgress() *int64 {
	return s.Progress
}

func (s *GetSkillRunResponseBody) GetProgressMessage() *string {
	return s.ProgressMessage
}

func (s *GetSkillRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillRunResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *GetSkillRunResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetSkillRunResponseBody) GetSkillCode() *string {
	return s.SkillCode
}

func (s *GetSkillRunResponseBody) GetSkillName() *string {
	return s.SkillName
}

func (s *GetSkillRunResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetSkillRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSkillRunResponseBody) GetUsage() map[string]interface{} {
	return s.Usage
}

func (s *GetSkillRunResponseBody) SetCode(v string) *GetSkillRunResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillRunResponseBody) SetCreatedAt(v string) *GetSkillRunResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetSkillRunResponseBody) SetErrorCode(v string) *GetSkillRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSkillRunResponseBody) SetErrorMessage(v string) *GetSkillRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSkillRunResponseBody) SetFinishedAt(v string) *GetSkillRunResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetSkillRunResponseBody) SetLogs(v []map[string]interface{}) *GetSkillRunResponseBody {
	s.Logs = v
	return s
}

func (s *GetSkillRunResponseBody) SetMessage(v string) *GetSkillRunResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillRunResponseBody) SetProgress(v int64) *GetSkillRunResponseBody {
	s.Progress = &v
	return s
}

func (s *GetSkillRunResponseBody) SetProgressMessage(v string) *GetSkillRunResponseBody {
	s.ProgressMessage = &v
	return s
}

func (s *GetSkillRunResponseBody) SetRequestId(v string) *GetSkillRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillRunResponseBody) SetResult(v map[string]interface{}) *GetSkillRunResponseBody {
	s.Result = v
	return s
}

func (s *GetSkillRunResponseBody) SetRunId(v string) *GetSkillRunResponseBody {
	s.RunId = &v
	return s
}

func (s *GetSkillRunResponseBody) SetSkillCode(v string) *GetSkillRunResponseBody {
	s.SkillCode = &v
	return s
}

func (s *GetSkillRunResponseBody) SetSkillName(v string) *GetSkillRunResponseBody {
	s.SkillName = &v
	return s
}

func (s *GetSkillRunResponseBody) SetStartedAt(v string) *GetSkillRunResponseBody {
	s.StartedAt = &v
	return s
}

func (s *GetSkillRunResponseBody) SetStatus(v string) *GetSkillRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetSkillRunResponseBody) SetUsage(v map[string]interface{}) *GetSkillRunResponseBody {
	s.Usage = v
	return s
}

func (s *GetSkillRunResponseBody) Validate() error {
	return dara.Validate(s)
}
