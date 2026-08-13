// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetCode() *string
	SetCompletedAt(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetCompletedAt() *string
	SetContent(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetContent() *string
	SetCreator(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetCreator() *string
	SetDigitalEmployeeName(v []*string) *GetScheduledTaskExecutionDetailResponseBody
	GetDigitalEmployeeName() []*string
	SetErrorMessage(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetErrorMessage() *string
	SetExecutionId(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetExecutionId() *string
	SetFiles(v []*GetScheduledTaskExecutionDetailResponseBodyFiles) *GetScheduledTaskExecutionDetailResponseBody
	GetFiles() []*GetScheduledTaskExecutionDetailResponseBodyFiles
	SetGmtCreate(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetMessage() *string
	SetMetadata(v *GetScheduledTaskExecutionDetailResponseBodyMetadata) *GetScheduledTaskExecutionDetailResponseBody
	GetMetadata() *GetScheduledTaskExecutionDetailResponseBodyMetadata
	SetOutputContent(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetOutputContent() *string
	SetPushResult(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetPushResult() *string
	SetRequestId(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetRequestId() *string
	SetSkillCodes(v []*string) *GetScheduledTaskExecutionDetailResponseBody
	GetSkillCodes() []*string
	SetStartedAt(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetStartedAt() *string
	SetStatus(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetTaskId() *string
	SetTitle(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetTitle() *string
	SetTriggerInfo(v *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) *GetScheduledTaskExecutionDetailResponseBody
	GetTriggerInfo() *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo
	SetTriggerType(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetTriggerType() *string
}

type GetScheduledTaskExecutionDetailResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 完成时间 ISO8601
	//
	// example:
	//
	// string_value
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// 执行完整内容
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建人
	//
	// example:
	//
	// string_value
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// digitalEmployeeName
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// 错误信息
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// 执行 ID
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string                                             `json:"executionId,omitempty" xml:"executionId,omitempty"`
	Files       []*GetScheduledTaskExecutionDetailResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// 创建时间 ISO8601
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 错误描述，成功时为空
	Message  *string                                              `json:"message,omitempty" xml:"message,omitempty"`
	Metadata *GetScheduledTaskExecutionDetailResponseBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// 结构化输出内容
	//
	// example:
	//
	// string_value
	OutputContent *string `json:"outputContent,omitempty" xml:"outputContent,omitempty"`
	PushResult    *string `json:"pushResult,omitempty" xml:"pushResult,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// skillCodes
	//
	// example:
	//
	// string_value
	SkillCodes []*string `json:"skillCodes,omitempty" xml:"skillCodes,omitempty" type:"Repeated"`
	// 开始时间 ISO8601
	//
	// example:
	//
	// string_value
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// 执行状态
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 任务 ID
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 执行结果标题
	//
	// example:
	//
	// 示例标题
	Title       *string                                                 `json:"title,omitempty" xml:"title,omitempty"`
	TriggerInfo *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo `json:"triggerInfo,omitempty" xml:"triggerInfo,omitempty" type:"Struct"`
	// 触发类型
	//
	// example:
	//
	// string_value
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s GetScheduledTaskExecutionDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetFiles() []*GetScheduledTaskExecutionDetailResponseBodyFiles {
	return s.Files
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetMetadata() *GetScheduledTaskExecutionDetailResponseBodyMetadata {
	return s.Metadata
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetOutputContent() *string {
	return s.OutputContent
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetPushResult() *string {
	return s.PushResult
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetSkillCodes() []*string {
	return s.SkillCodes
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetTriggerInfo() *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo {
	return s.TriggerInfo
}

func (s *GetScheduledTaskExecutionDetailResponseBody) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetCode(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetCompletedAt(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetContent(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Content = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetCreator(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Creator = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetDigitalEmployeeName(v []*string) *GetScheduledTaskExecutionDetailResponseBody {
	s.DigitalEmployeeName = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetErrorMessage(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetExecutionId(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.ExecutionId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetFiles(v []*GetScheduledTaskExecutionDetailResponseBodyFiles) *GetScheduledTaskExecutionDetailResponseBody {
	s.Files = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetGmtCreate(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetMessage(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetMetadata(v *GetScheduledTaskExecutionDetailResponseBodyMetadata) *GetScheduledTaskExecutionDetailResponseBody {
	s.Metadata = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetOutputContent(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.OutputContent = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetPushResult(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.PushResult = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetRequestId(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetSkillCodes(v []*string) *GetScheduledTaskExecutionDetailResponseBody {
	s.SkillCodes = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetStartedAt(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.StartedAt = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetStatus(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Status = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetTaskId(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetTitle(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Title = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetTriggerInfo(v *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) *GetScheduledTaskExecutionDetailResponseBody {
	s.TriggerInfo = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) SetTriggerType(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.TriggerType = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBody) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerInfo != nil {
		if err := s.TriggerInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScheduledTaskExecutionDetailResponseBodyFiles struct {
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件 OSS URL
	//
	// example:
	//
	// https://example.com/oss/file.pdf
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetScheduledTaskExecutionDetailResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailResponseBodyFiles) GetName() *string {
	return s.Name
}

func (s *GetScheduledTaskExecutionDetailResponseBodyFiles) GetPath() *string {
	return s.Path
}

func (s *GetScheduledTaskExecutionDetailResponseBodyFiles) SetName(v string) *GetScheduledTaskExecutionDetailResponseBodyFiles {
	s.Name = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBodyFiles) SetPath(v string) *GetScheduledTaskExecutionDetailResponseBodyFiles {
	s.Path = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}

type GetScheduledTaskExecutionDetailResponseBodyMetadata struct {
	// 会话 ID
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 执行结果推送状态（多频道时为列表）
	//
	// example:
	//
	// string_value
	Usage map[string]interface{} `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s GetScheduledTaskExecutionDetailResponseBodyMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailResponseBodyMetadata) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailResponseBodyMetadata) GetSessionId() *string {
	return s.SessionId
}

func (s *GetScheduledTaskExecutionDetailResponseBodyMetadata) GetUsage() map[string]interface{} {
	return s.Usage
}

func (s *GetScheduledTaskExecutionDetailResponseBodyMetadata) SetSessionId(v string) *GetScheduledTaskExecutionDetailResponseBodyMetadata {
	s.SessionId = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBodyMetadata) SetUsage(v map[string]interface{}) *GetScheduledTaskExecutionDetailResponseBodyMetadata {
	s.Usage = v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBodyMetadata) Validate() error {
	return dara.Validate(s)
}

type GetScheduledTaskExecutionDetailResponseBodyTriggerInfo struct {
	// 触发执行的用户标识
	//
	// example:
	//
	// user_10001
	TriggeredBy *string `json:"triggeredBy,omitempty" xml:"triggeredBy,omitempty"`
}

func (s GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) GetTriggeredBy() *string {
	return s.TriggeredBy
}

func (s *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) SetTriggeredBy(v string) *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo {
	s.TriggeredBy = &v
	return s
}

func (s *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo) Validate() error {
	return dara.Validate(s)
}
