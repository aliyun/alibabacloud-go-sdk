// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendChatMessageRequest
	GetContent() *string
	SetContentType(v string) *SendChatMessageRequest
	GetContentType() *string
	SetDigitalEmployeeName(v []*string) *SendChatMessageRequest
	GetDigitalEmployeeName() []*string
	SetDirectChat(v bool) *SendChatMessageRequest
	GetDirectChat() *bool
	SetFiles(v []*SendChatMessageRequestFiles) *SendChatMessageRequest
	GetFiles() []*SendChatMessageRequestFiles
	SetModel(v string) *SendChatMessageRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendChatMessageRequest
	GetReuseLastSession() *bool
	SetSessionId(v string) *SendChatMessageRequest
	GetSessionId() *string
	SetStream(v bool) *SendChatMessageRequest
	GetStream() *bool
	SetTaskExecution(v *SendChatMessageRequestTaskExecution) *SendChatMessageRequest
	GetTaskExecution() *SendChatMessageRequestTaskExecution
	SetTenantId(v string) *SendChatMessageRequest
	GetTenantId() *string
}

type SendChatMessageRequest struct {
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
	// Specifies whether to enable direct connection mode. When set to true, the regular scenario routing is skipped and the direct conversation scenario is entered directly.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The list of file references. Each item is an object, and fileId is required (returned by uploadChatFile).
	Files []*SendChatMessageRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The abstract model tier. Valid values: quick, standard, and flagship. If not specified, new sessions use standard, and existing sessions retain the current session tier.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Specifies whether to reuse the most recent session of the digital employee when sessionId is not provided (CLI scenario). Default value: false, which creates a new session.
	//
	// example:
	//
	// false
	ReuseLastSession *bool `json:"reuseLastSession,omitempty" xml:"reuseLastSession,omitempty"`
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to use streaming output.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by executeScheduledTask. When provided, the request is processed through the task execution pipeline.
	TaskExecution *SendChatMessageRequestTaskExecution `json:"taskExecution,omitempty" xml:"taskExecution,omitempty" type:"Struct"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SendChatMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequest) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequest) GetContent() *string {
	return s.Content
}

func (s *SendChatMessageRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendChatMessageRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *SendChatMessageRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendChatMessageRequest) GetFiles() []*SendChatMessageRequestFiles {
	return s.Files
}

func (s *SendChatMessageRequest) GetModel() *string {
	return s.Model
}

func (s *SendChatMessageRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendChatMessageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendChatMessageRequest) GetStream() *bool {
	return s.Stream
}

func (s *SendChatMessageRequest) GetTaskExecution() *SendChatMessageRequestTaskExecution {
	return s.TaskExecution
}

func (s *SendChatMessageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendChatMessageRequest) SetContent(v string) *SendChatMessageRequest {
	s.Content = &v
	return s
}

func (s *SendChatMessageRequest) SetContentType(v string) *SendChatMessageRequest {
	s.ContentType = &v
	return s
}

func (s *SendChatMessageRequest) SetDigitalEmployeeName(v []*string) *SendChatMessageRequest {
	s.DigitalEmployeeName = v
	return s
}

func (s *SendChatMessageRequest) SetDirectChat(v bool) *SendChatMessageRequest {
	s.DirectChat = &v
	return s
}

func (s *SendChatMessageRequest) SetFiles(v []*SendChatMessageRequestFiles) *SendChatMessageRequest {
	s.Files = v
	return s
}

func (s *SendChatMessageRequest) SetModel(v string) *SendChatMessageRequest {
	s.Model = &v
	return s
}

func (s *SendChatMessageRequest) SetReuseLastSession(v bool) *SendChatMessageRequest {
	s.ReuseLastSession = &v
	return s
}

func (s *SendChatMessageRequest) SetSessionId(v string) *SendChatMessageRequest {
	s.SessionId = &v
	return s
}

func (s *SendChatMessageRequest) SetStream(v bool) *SendChatMessageRequest {
	s.Stream = &v
	return s
}

func (s *SendChatMessageRequest) SetTaskExecution(v *SendChatMessageRequestTaskExecution) *SendChatMessageRequest {
	s.TaskExecution = v
	return s
}

func (s *SendChatMessageRequest) SetTenantId(v string) *SendChatMessageRequest {
	s.TenantId = &v
	return s
}

func (s *SendChatMessageRequest) Validate() error {
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

type SendChatMessageRequestFiles struct {
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleFileId
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// The element type. Valid values: text, web_search, mention, and skill.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendChatMessageRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestFiles) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestFiles) GetFileId() *string {
	return s.FileId
}

func (s *SendChatMessageRequestFiles) GetType() *string {
	return s.Type
}

func (s *SendChatMessageRequestFiles) SetFileId(v string) *SendChatMessageRequestFiles {
	s.FileId = &v
	return s
}

func (s *SendChatMessageRequestFiles) SetType(v string) *SendChatMessageRequestFiles {
	s.Type = &v
	return s
}

func (s *SendChatMessageRequestFiles) Validate() error {
	return dara.Validate(s)
}

type SendChatMessageRequestTaskExecution struct {
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

func (s SendChatMessageRequestTaskExecution) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestTaskExecution) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestTaskExecution) GetBillingId() *string {
	return s.BillingId
}

func (s *SendChatMessageRequestTaskExecution) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendChatMessageRequestTaskExecution) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *SendChatMessageRequestTaskExecution) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *SendChatMessageRequestTaskExecution) GetSkillCodes() []*string {
	return s.SkillCodes
}

func (s *SendChatMessageRequestTaskExecution) GetTaskId() *string {
	return s.TaskId
}

func (s *SendChatMessageRequestTaskExecution) GetTaskName() *string {
	return s.TaskName
}

func (s *SendChatMessageRequestTaskExecution) GetTaskUnderstand() *string {
	return s.TaskUnderstand
}

func (s *SendChatMessageRequestTaskExecution) GetTenantId() *string {
	return s.TenantId
}

func (s *SendChatMessageRequestTaskExecution) GetUserId() *string {
	return s.UserId
}

func (s *SendChatMessageRequestTaskExecution) SetBillingId(v string) *SendChatMessageRequestTaskExecution {
	s.BillingId = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetEnableWebSearch(v bool) *SendChatMessageRequestTaskExecution {
	s.EnableWebSearch = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetExecutionId(v string) *SendChatMessageRequestTaskExecution {
	s.ExecutionId = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetOperatingObjectName(v string) *SendChatMessageRequestTaskExecution {
	s.OperatingObjectName = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetSkillCodes(v []*string) *SendChatMessageRequestTaskExecution {
	s.SkillCodes = v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetTaskId(v string) *SendChatMessageRequestTaskExecution {
	s.TaskId = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetTaskName(v string) *SendChatMessageRequestTaskExecution {
	s.TaskName = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetTaskUnderstand(v string) *SendChatMessageRequestTaskExecution {
	s.TaskUnderstand = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetTenantId(v string) *SendChatMessageRequestTaskExecution {
	s.TenantId = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) SetUserId(v string) *SendChatMessageRequestTaskExecution {
	s.UserId = &v
	return s
}

func (s *SendChatMessageRequestTaskExecution) Validate() error {
	return dara.Validate(s)
}
