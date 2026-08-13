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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 任务创建时间，ISO8601
	//
	// example:
	//
	// string_value
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 错误码，仅 Failed 时返回
	//
	// example:
	//
	// string_value
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 错误描述，仅 Failed 时返回
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 任务结束时间，ISO8601；仅终态（Succeeded/Failed/Cancelled）有值
	//
	// example:
	//
	// string_value
	FinishedAt *string                  `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	Logs       []map[string]interface{} `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 进度百分比（仅 Running 时有意义）
	//
	// example:
	//
	// 1
	Progress *int64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// 进度描述
	//
	// example:
	//
	// string_value
	ProgressMessage *string `json:"progressMessage,omitempty" xml:"progressMessage,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// 异步任务 ID
	//
	// example:
	//
	// exampleRunId
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// 技能编码
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// 技能名称
	//
	// example:
	//
	// string_value
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// 任务开始执行时间，ISO8601
	//
	// example:
	//
	// string_value
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// 执行状态：Running / Succeeded / Failed / Cancelled
	//
	// example:
	//
	// READY
	Status *string                `json:"status,omitempty" xml:"status,omitempty"`
	Usage  map[string]interface{} `json:"usage,omitempty" xml:"usage,omitempty"`
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
