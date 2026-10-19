// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncChatMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendAsyncChatMessageShrinkRequest
	GetContent() *string
	SetContentType(v string) *SendAsyncChatMessageShrinkRequest
	GetContentType() *string
	SetDigitalEmployeeNameShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetDirectChat(v bool) *SendAsyncChatMessageShrinkRequest
	GetDirectChat() *bool
	SetEnableWebSearch(v bool) *SendAsyncChatMessageShrinkRequest
	GetEnableWebSearch() *bool
	SetFilesShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetFilesShrink() *string
	SetModel(v string) *SendAsyncChatMessageShrinkRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendAsyncChatMessageShrinkRequest
	GetReuseLastSession() *bool
	SetSessionId(v string) *SendAsyncChatMessageShrinkRequest
	GetSessionId() *string
	SetStream(v bool) *SendAsyncChatMessageShrinkRequest
	GetStream() *bool
	SetTaskExecutionShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetTaskExecutionShrink() *string
	SetTenantId(v string) *SendAsyncChatMessageShrinkRequest
	GetTenantId() *string
	SetWorkMode(v string) *SendAsyncChatMessageShrinkRequest
	GetWorkMode() *string
}

type SendAsyncChatMessageShrinkRequest struct {
	// The body of the user message.
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
	// The list of digital employee names. For compatibility with the legacy format, a single string can also be passed.
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether to enable direct connection mode. If set to true, regular scenario routing is skipped and the direct conversation scenario is used.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// Specifies whether to enable web search. Default value: false. In the task execution scenario where taskExecution is passed, the task configuration prevails.
	//
	// example:
	//
	// false
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The list of file references. Each item is an object, and the fileId parameter is required. The fileId is returned by the uploadChatFile operation.
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
	// The abstract model tier. Valid values: quick, standard, and flagship. If not specified, standard is used for new sessions, and the current tier is retained for existing sessions.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Specifies whether to reuse the most recent session under the digital employee when sessionId is not specified. This is typically used in CLI scenarios. Default value: false, which means a new session is created.
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
	// Specifies whether to use streaming generation. This operation always generates backend content in streaming mode and writes it to the message stream. The value of this parameter does not change the response structure.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by the executeScheduledTask operation. If passed, the request is processed through the task execution pipeline.
	TaskExecutionShrink *string `json:"taskExecution,omitempty" xml:"taskExecution,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The session work mode. Valid values: ask (quick Q&A, which trims tools, skills, and connectors for single-turn direct answers), work (deep work, default), and direct (direct connection mode at the request level, which does not start a sandbox and prevents context pollution, equivalent to directChat=true). The ask and work modes are session-level. The mode is fixed when a new session is created and inherited by default in subsequent chats. If the explicitly passed value is inconsistent with the session mode, a parameter error is returned. To switch modes, create a new session or fork the current one. If ask is passed in multi-digital-employee or task execution scenarios, it takes effect as work. This parameter is ignored when directChat is set to true.
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
}

func (s SendAsyncChatMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *SendAsyncChatMessageShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendAsyncChatMessageShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *SendAsyncChatMessageShrinkRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendAsyncChatMessageShrinkRequest) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendAsyncChatMessageShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *SendAsyncChatMessageShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *SendAsyncChatMessageShrinkRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendAsyncChatMessageShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendAsyncChatMessageShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *SendAsyncChatMessageShrinkRequest) GetTaskExecutionShrink() *string {
	return s.TaskExecutionShrink
}

func (s *SendAsyncChatMessageShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendAsyncChatMessageShrinkRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendAsyncChatMessageShrinkRequest) SetContent(v string) *SendAsyncChatMessageShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetContentType(v string) *SendAsyncChatMessageShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetDigitalEmployeeNameShrink(v string) *SendAsyncChatMessageShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetDirectChat(v bool) *SendAsyncChatMessageShrinkRequest {
	s.DirectChat = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetEnableWebSearch(v bool) *SendAsyncChatMessageShrinkRequest {
	s.EnableWebSearch = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetFilesShrink(v string) *SendAsyncChatMessageShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetModel(v string) *SendAsyncChatMessageShrinkRequest {
	s.Model = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetReuseLastSession(v bool) *SendAsyncChatMessageShrinkRequest {
	s.ReuseLastSession = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetSessionId(v string) *SendAsyncChatMessageShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetStream(v bool) *SendAsyncChatMessageShrinkRequest {
	s.Stream = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetTaskExecutionShrink(v string) *SendAsyncChatMessageShrinkRequest {
	s.TaskExecutionShrink = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetTenantId(v string) *SendAsyncChatMessageShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetWorkMode(v string) *SendAsyncChatMessageShrinkRequest {
	s.WorkMode = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
