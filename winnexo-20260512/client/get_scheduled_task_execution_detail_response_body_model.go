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
	SetVisibility(v string) *GetScheduledTaskExecutionDetailResponseBody
	GetVisibility() *string
}

type GetScheduledTaskExecutionDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The completion time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// The full execution content.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The creator.
	//
	// example:
	//
	// string_value
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The list of digital employee names.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The execution ID.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// The list of output files.
	Files []*GetScheduledTaskExecutionDetailResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The extended metadata.
	Metadata *GetScheduledTaskExecutionDetailResponseBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The structured output content.
	//
	// example:
	//
	// string_value
	OutputContent *string `json:"outputContent,omitempty" xml:"outputContent,omitempty"`
	// The push status of the execution result.
	//
	// example:
	//
	// succuss
	PushResult *string `json:"pushResult,omitempty" xml:"pushResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of associated skill codes.
	//
	// example:
	//
	// string_value
	SkillCodes []*string `json:"skillCodes,omitempty" xml:"skillCodes,omitempty" type:"Repeated"`
	// The start time in ISO 8601 format.
	//
	// example:
	//
	// string_value
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The execution status.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The execution result title.
	//
	// example:
	//
	// Sample title
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The trigger information.
	TriggerInfo *GetScheduledTaskExecutionDetailResponseBodyTriggerInfo `json:"triggerInfo,omitempty" xml:"triggerInfo,omitempty" type:"Struct"`
	// The trigger type.
	//
	// example:
	//
	// string_value
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
	// The visibility scope of the execution record, which is always equal to the visibility scope of the associated task. Valid values: PRIVATE, COLLABORATIVE, and PUBLIC. This field is empty for personal task executions.
	//
	// example:
	//
	// COLLABORATIVE
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
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

func (s *GetScheduledTaskExecutionDetailResponseBody) GetVisibility() *string {
	return s.Visibility
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

func (s *GetScheduledTaskExecutionDetailResponseBody) SetVisibility(v string) *GetScheduledTaskExecutionDetailResponseBody {
	s.Visibility = &v
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
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The OSS URL of the file.
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
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The token usage information.
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
	// The user identifier that triggered the execution.
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
