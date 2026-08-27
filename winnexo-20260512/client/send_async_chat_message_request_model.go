// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendAsyncChatMessageRequest
	GetContent() *string
	SetContentType(v string) *SendAsyncChatMessageRequest
	GetContentType() *string
	SetDigitalEmployeeName(v []*string) *SendAsyncChatMessageRequest
	GetDigitalEmployeeName() []*string
	SetDirectChat(v bool) *SendAsyncChatMessageRequest
	GetDirectChat() *bool
	SetFiles(v []*SendAsyncChatMessageRequestFiles) *SendAsyncChatMessageRequest
	GetFiles() []*SendAsyncChatMessageRequestFiles
	SetModel(v string) *SendAsyncChatMessageRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendAsyncChatMessageRequest
	GetReuseLastSession() *bool
	SetSessionId(v string) *SendAsyncChatMessageRequest
	GetSessionId() *string
	SetStream(v bool) *SendAsyncChatMessageRequest
	GetStream() *bool
	SetTaskExecution(v *SendAsyncChatMessageRequestTaskExecution) *SendAsyncChatMessageRequest
	GetTaskExecution() *SendAsyncChatMessageRequestTaskExecution
	SetTenantId(v string) *SendAsyncChatMessageRequest
	GetTenantId() *string
}

type SendAsyncChatMessageRequest struct {
	// The message body from the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The message type. Valid values: Text and Markdown.
	//
	// example:
	//
	// Text
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The list of digital employee names. A single string can be passed for backward compatibility with the legacy format.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// Specifies whether to enable direct chat mode. If set to true, the regular scenario routing is skipped and the direct chat scenario is entered.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The list of file references. Each item is an object in which fileId is required and is returned by uploadChatFile.
	Files []*SendAsyncChatMessageRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The abstract model tier. Valid values: quick, standard, and flagship. If not specified, new sessions use standard, and existing sessions retain their current tier.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Specifies whether to reuse the most recent session of the digital employee when sessionId is not specified. This is designed for CLI scenarios. Default value: false, which creates a new session.
	//
	// example:
	//
	// false
	ReuseLastSession *bool `json:"reuseLastSession,omitempty" xml:"reuseLastSession,omitempty"`
	// The session ID. If not specified, a new session is created.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to use streaming generation. This operation always generates backend content in streaming mode and writes it to the message stream. The value does not change the response structure.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by executeScheduledTask. When provided, the request is processed through the task execution pipeline.
	TaskExecution *SendAsyncChatMessageRequestTaskExecution `json:"taskExecution,omitempty" xml:"taskExecution,omitempty" type:"Struct"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SendAsyncChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequest) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendAsyncChatMessageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendAsyncChatMessageRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *SendAsyncChatMessageRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendAsyncChatMessageRequest) GetFiles() []*SendAsyncChatMessageRequestFiles {
	return s.Files
}

func (s *SendAsyncChatMessageRequest) GetModel() *string {
	return s.Model
}

func (s *SendAsyncChatMessageRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendAsyncChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendAsyncChatMessageRequest) GetStream() *bool {
	return s.Stream
}

func (s *SendAsyncChatMessageRequest) GetTaskExecution() *SendAsyncChatMessageRequestTaskExecution {
	return s.TaskExecution
}

func (s *SendAsyncChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendAsyncChatMessageRequest) SetContent(v string) *SendAsyncChatMessageRequest {
	s.Content = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetContentType(v string) *SendAsyncChatMessageRequest {
	s.ContentType = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetDigitalEmployeeName(v []*string) *SendAsyncChatMessageRequest {
	s.DigitalEmployeeName = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetDirectChat(v bool) *SendAsyncChatMessageRequest {
	s.DirectChat = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetFiles(v []*SendAsyncChatMessageRequestFiles) *SendAsyncChatMessageRequest {
	s.Files = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetModel(v string) *SendAsyncChatMessageRequest {
	s.Model = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetReuseLastSession(v bool) *SendAsyncChatMessageRequest {
	s.ReuseLastSession = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetSessionId(v string) *SendAsyncChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetStream(v bool) *SendAsyncChatMessageRequest {
	s.Stream = &v
	return s
}

func (s *SendAsyncChatMessageRequest) SetTaskExecution(v *SendAsyncChatMessageRequestTaskExecution) *SendAsyncChatMessageRequest {
	s.TaskExecution = v
	return s
}

func (s *SendAsyncChatMessageRequest) SetTenantId(v string) *SendAsyncChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *SendAsyncChatMessageRequest) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskExecution != nil {
		if err := s.TaskExecution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendAsyncChatMessageRequestFiles struct {
	// The file ID returned by uploadChatFile.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleFileId
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// The file type.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendAsyncChatMessageRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestFiles) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestFiles) GetFileId() *string {
	return s.FileId
}

func (s *SendAsyncChatMessageRequestFiles) GetType() *string {
	return s.Type
}

func (s *SendAsyncChatMessageRequestFiles) SetFileId(v string) *SendAsyncChatMessageRequestFiles {
	s.FileId = &v
	return s
}

func (s *SendAsyncChatMessageRequestFiles) SetType(v string) *SendAsyncChatMessageRequestFiles {
	s.Type = &v
	return s
}

func (s *SendAsyncChatMessageRequestFiles) Validate() error {
	return dara.Validate(s)
}

type SendAsyncChatMessageRequestTaskExecution struct {
	// The billing ID.
	//
	// example:
	//
	// exampleBillingId
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// Specifies whether to enable web search.
	//
	// example:
	//
	// true
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// The digital employee name.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The list of associated skill codes.
	//
	// example:
	//
	// string_value
	SkillCodes []*string `json:"skillCodes,omitempty" xml:"skillCodes,omitempty" type:"Repeated"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// string_value
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The task understanding content.
	//
	// example:
	//
	// string_value
	TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The user ID to which the task belongs.
	//
	// example:
	//
	// exampleUserId
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SendAsyncChatMessageRequestTaskExecution) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageRequestTaskExecution) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetBillingId() *string {
	return s.BillingId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetSkillCodes() []*string {
	return s.SkillCodes
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTaskId() *string {
	return s.TaskId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTaskName() *string {
	return s.TaskName
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTaskUnderstand() *string {
	return s.TaskUnderstand
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetTenantId() *string {
	return s.TenantId
}

func (s *SendAsyncChatMessageRequestTaskExecution) GetUserId() *string {
	return s.UserId
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetBillingId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.BillingId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetEnableWebSearch(v bool) *SendAsyncChatMessageRequestTaskExecution {
	s.EnableWebSearch = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetExecutionId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.ExecutionId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetOperatingObjectName(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.OperatingObjectName = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetSkillCodes(v []*string) *SendAsyncChatMessageRequestTaskExecution {
	s.SkillCodes = v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTaskId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TaskId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTaskName(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TaskName = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTaskUnderstand(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TaskUnderstand = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetTenantId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.TenantId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) SetUserId(v string) *SendAsyncChatMessageRequestTaskExecution {
	s.UserId = &v
	return s
}

func (s *SendAsyncChatMessageRequestTaskExecution) Validate() error {
	return dara.Validate(s)
}
