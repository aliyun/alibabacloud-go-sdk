// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendChatMessageShrinkRequest
	GetContent() *string
	SetContentType(v string) *SendChatMessageShrinkRequest
	GetContentType() *string
	SetDigitalEmployeeNameShrink(v string) *SendChatMessageShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetDirectChat(v bool) *SendChatMessageShrinkRequest
	GetDirectChat() *bool
	SetFilesShrink(v string) *SendChatMessageShrinkRequest
	GetFilesShrink() *string
	SetModel(v string) *SendChatMessageShrinkRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendChatMessageShrinkRequest
	GetReuseLastSession() *bool
	SetSessionId(v string) *SendChatMessageShrinkRequest
	GetSessionId() *string
	SetStream(v bool) *SendChatMessageShrinkRequest
	GetStream() *bool
	SetTaskExecutionShrink(v string) *SendChatMessageShrinkRequest
	GetTaskExecutionShrink() *string
	SetTenantId(v string) *SendChatMessageShrinkRequest
	GetTenantId() *string
}

type SendChatMessageShrinkRequest struct {
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
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether to enable direct connection mode. If set to true, the regular scenario routing is skipped and the direct conversation scenario is entered.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The list of file references. Each item is an object in which fileId is required and is returned by uploadChatFile.
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
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
	TaskExecutionShrink *string `json:"taskExecution,omitempty" xml:"taskExecution,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SendChatMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendChatMessageShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *SendChatMessageShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendChatMessageShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *SendChatMessageShrinkRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendChatMessageShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *SendChatMessageShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *SendChatMessageShrinkRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendChatMessageShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *SendChatMessageShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *SendChatMessageShrinkRequest) GetTaskExecutionShrink() *string {
	return s.TaskExecutionShrink
}

func (s *SendChatMessageShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SendChatMessageShrinkRequest) SetContent(v string) *SendChatMessageShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetContentType(v string) *SendChatMessageShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetDigitalEmployeeNameShrink(v string) *SendChatMessageShrinkRequest {
	s.DigitalEmployeeNameShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetDirectChat(v bool) *SendChatMessageShrinkRequest {
	s.DirectChat = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetFilesShrink(v string) *SendChatMessageShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetModel(v string) *SendChatMessageShrinkRequest {
	s.Model = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetReuseLastSession(v bool) *SendChatMessageShrinkRequest {
	s.ReuseLastSession = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetSessionId(v string) *SendChatMessageShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetStream(v bool) *SendChatMessageShrinkRequest {
	s.Stream = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetTaskExecutionShrink(v string) *SendChatMessageShrinkRequest {
	s.TaskExecutionShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetTenantId(v string) *SendChatMessageShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SendChatMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
