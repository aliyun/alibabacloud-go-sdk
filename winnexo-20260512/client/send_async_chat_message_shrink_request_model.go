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
	SetContextItemsShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetContextItemsShrink() *string
	SetDigitalEmployeeNameShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetDirectChat(v bool) *SendAsyncChatMessageShrinkRequest
	GetDirectChat() *bool
	SetDirectMaxHistoryRounds(v interface{}) *SendAsyncChatMessageShrinkRequest
	GetDirectMaxHistoryRounds() interface{}
	SetEnableWebSearch(v bool) *SendAsyncChatMessageShrinkRequest
	GetEnableWebSearch() *bool
	SetFilesShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetFilesShrink() *string
	SetLocale(v string) *SendAsyncChatMessageShrinkRequest
	GetLocale() *string
	SetModel(v string) *SendAsyncChatMessageShrinkRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendAsyncChatMessageShrinkRequest
	GetReuseLastSession() *bool
	SetSegmentsShrink(v string) *SendAsyncChatMessageShrinkRequest
	GetSegmentsShrink() *string
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
	// The message or file entries referenced as context for the current turn.
	//
	// example:
	//
	// []
	ContextItemsShrink *string `json:"contextItems,omitempty" xml:"contextItems,omitempty"`
	// The list of digital employee names. For backward compatibility, a single string can also be passed.
	//
	// example:
	//
	// ["pcitc-magent"]
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether to enable direct connection mode. If set to true, regular scenario routing is skipped and the direct conversation scenario is entered directly.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The number of recent complete direct Q&A rounds carried in the current turn, excluding the current input. This parameter takes effect only when workMode is set to direct or directChat is set to true. Valid values are JSON integers from 0 to 100. If omitted, the backend processes it as 0, carrying no history but still saving the current round. Null, Boolean, string, decimal, and out-of-bounds values are processed through the original direct path, and the gateway retains the original values for backend determination. This field is ignored in work and ask modes. This parameter does not limit the history range when switching to the work mode.
	//
	// example:
	//
	// 1
	DirectMaxHistoryRounds interface{} `json:"directMaxHistoryRounds,omitempty" xml:"directMaxHistoryRounds,omitempty"`
	// Specifies whether to enable web search. Default value: false. In task execution scenarios where taskExecution is passed, the task configuration prevails.
	//
	// example:
	//
	// false
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The list of file references. Each item is an object, and fileId is required, which is returned by uploadChatFile.
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
	// The language for the current turn. Valid values: zh-CN and en-US. This parameter is used for API error messages and passed through to CodeAgent. If omitted, the current language behavior is retained and user settings are not modified.
	//
	// example:
	//
	// en-US
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The abstract model tier. Valid values: quick, standard, and flagship. If not specified, standard is used for new sessions, and the current tier is retained for existing sessions.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Specifies whether to reuse the most recent session under the digital employee when sessionId is not provided, which applies to CLI scenarios. Default value: false, which means a new session is created.
	//
	// example:
	//
	// false
	ReuseLastSession *bool `json:"reuseLastSession,omitempty" xml:"reuseLastSession,omitempty"`
	// The structured segments in the message, such as session or object mentions.
	//
	// example:
	//
	// []
	SegmentsShrink *string `json:"segments,omitempty" xml:"segments,omitempty"`
	// The session ID. If not specified, a new session is created.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to use streaming generation. This API always generates background content in streaming mode and writes it to the message stream. The value does not change the response structure.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by executeScheduledTask. If passed, the request is processed through the task execution pipeline.
	TaskExecutionShrink *string `json:"taskExecution,omitempty" xml:"taskExecution,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The session work mode. Valid values:
	//
	// - ask: Quick Q&A mode. Trims tools, skills, and connectors for single-turn direct answers.
	//
	// - work: Deep work mode. This is the default value.
	//
	// - direct: Direct connection mode at the request level. Does not start a sandbox and prevents context pollution. Equivalent to setting directChat to true.
	//
	// The ask and work modes are session-level. Once selected during session creation, the mode is fixed and inherited by default in subsequent chats. If the explicitly passed value conflicts with the session mode, a parameter error is returned. To switch modes, create a new session or fork the current one. In scenarios with multiple digital employees or task execution, passing ask takes effect as work. This parameter is ignored when directChat is set to true.
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

func (s *SendAsyncChatMessageShrinkRequest) GetContextItemsShrink() *string {
	return s.ContextItemsShrink
}

func (s *SendAsyncChatMessageShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *SendAsyncChatMessageShrinkRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendAsyncChatMessageShrinkRequest) GetDirectMaxHistoryRounds() interface{} {
	return s.DirectMaxHistoryRounds
}

func (s *SendAsyncChatMessageShrinkRequest) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendAsyncChatMessageShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *SendAsyncChatMessageShrinkRequest) GetLocale() *string {
	return s.Locale
}

func (s *SendAsyncChatMessageShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *SendAsyncChatMessageShrinkRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendAsyncChatMessageShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
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

func (s *SendAsyncChatMessageShrinkRequest) SetContextItemsShrink(v string) *SendAsyncChatMessageShrinkRequest {
	s.ContextItemsShrink = &v
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

func (s *SendAsyncChatMessageShrinkRequest) SetDirectMaxHistoryRounds(v interface{}) *SendAsyncChatMessageShrinkRequest {
	s.DirectMaxHistoryRounds = v
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

func (s *SendAsyncChatMessageShrinkRequest) SetLocale(v string) *SendAsyncChatMessageShrinkRequest {
	s.Locale = &v
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

func (s *SendAsyncChatMessageShrinkRequest) SetSegmentsShrink(v string) *SendAsyncChatMessageShrinkRequest {
	s.SegmentsShrink = &v
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
