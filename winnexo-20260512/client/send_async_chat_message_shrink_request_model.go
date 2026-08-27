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
}

type SendAsyncChatMessageShrinkRequest struct {
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
	// Specifies whether to enable direct chat mode. If set to true, the regular scenario routing is skipped and the direct chat scenario is entered.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The list of file references. Each item is an object in which fileId is required and is returned by uploadChatFile.
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
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
	TaskExecutionShrink *string `json:"taskExecution,omitempty" xml:"taskExecution,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
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

func (s *SendAsyncChatMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
