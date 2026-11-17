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
	SetContextItemsShrink(v string) *SendChatMessageShrinkRequest
	GetContextItemsShrink() *string
	SetDigitalEmployeeNameShrink(v string) *SendChatMessageShrinkRequest
	GetDigitalEmployeeNameShrink() *string
	SetDirectChat(v bool) *SendChatMessageShrinkRequest
	GetDirectChat() *bool
	SetDirectMaxHistoryRounds(v interface{}) *SendChatMessageShrinkRequest
	GetDirectMaxHistoryRounds() interface{}
	SetEnableWebSearch(v bool) *SendChatMessageShrinkRequest
	GetEnableWebSearch() *bool
	SetFilesShrink(v string) *SendChatMessageShrinkRequest
	GetFilesShrink() *string
	SetLocale(v string) *SendChatMessageShrinkRequest
	GetLocale() *string
	SetModel(v string) *SendChatMessageShrinkRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendChatMessageShrinkRequest
	GetReuseLastSession() *bool
	SetSegmentsShrink(v string) *SendChatMessageShrinkRequest
	GetSegmentsShrink() *string
	SetSessionId(v string) *SendChatMessageShrinkRequest
	GetSessionId() *string
	SetStream(v bool) *SendChatMessageShrinkRequest
	GetStream() *bool
	SetTaskExecutionShrink(v string) *SendChatMessageShrinkRequest
	GetTaskExecutionShrink() *string
	SetTenantId(v string) *SendChatMessageShrinkRequest
	GetTenantId() *string
	SetWorkMode(v string) *SendChatMessageShrinkRequest
	GetWorkMode() *string
}

type SendChatMessageShrinkRequest struct {
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
	// The list of digital employee names. For compatibility with the legacy format, a single string can be passed.
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Specifies whether to enable direct connection mode. If set to true, regular scenario routing is skipped and the direct conversation scenario is entered directly.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The number of recent complete direct Q&A rounds carried in the current turn, excluding the input of the current turn. This parameter takes effect only when workMode is set to direct or directChat is set to true. Valid values: JSON integers from 0 to 100. If omitted, the backend processes it as 0, carrying no history but still saving the current turn. Null, Boolean, string, decimal, and out-of-bounds values are processed through the original direct path, and the gateway retains the original values for backend determination. This field is ignored in work and ask modes. This parameter does not limit the history range when converted to the work mode.
	//
	// example:
	//
	// 1
	DirectMaxHistoryRounds interface{} `json:"directMaxHistoryRounds,omitempty" xml:"directMaxHistoryRounds,omitempty"`
	// Specifies whether to enable web search. Default value: false. In the task execution scenario where taskExecution is passed, the task configuration prevails.
	//
	// example:
	//
	// false
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The list of file references. Each item is an object, and the fileId parameter is required, which is returned by uploadChatFile.
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
	// The language for the current turn. Valid values: zh-CN and en-US. This parameter is used for API error messages and passed through to CodeAgent. If omitted, the current language behavior is retained and user settings are not modified.
	//
	// example:
	//
	// en-US
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The abstract model tier. Valid values: quick, standard, and flagship. If omitted, standard is used for new sessions, and the current tier is retained for existing sessions.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Specifies whether to reuse the most recent session under the digital employee when sessionId is not passed, which applies to CLI scenarios. Default value: false, which indicates that a new session is created.
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
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether streaming output is enabled.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by executeScheduledTask. If passed, the request is processed through the task execution pipeline.
	TaskExecutionShrink *string `json:"taskExecution,omitempty" xml:"taskExecution,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The session work mode. Valid values: ask (quick Q&A, which trims tools, skills, and connectors for single-turn direct answers), work (deep work, which is the default), and direct (direct connection mode at the request level, which does not start a sandbox and prevents context pollution, equivalent to directChat=true). The ask and work modes are session-level. The mode is fixed when a new session is created. Subsequent chats inherit the session mode by default. If the explicitly passed value is inconsistent with the session mode, a parameter error is returned. To switch modes, create a new session or fork the session. If ask is passed in multi-digital employee or task execution scenarios, it takes effect as work. This parameter is ignored when directChat is set to true.
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
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

func (s *SendChatMessageShrinkRequest) GetContextItemsShrink() *string {
	return s.ContextItemsShrink
}

func (s *SendChatMessageShrinkRequest) GetDigitalEmployeeNameShrink() *string {
	return s.DigitalEmployeeNameShrink
}

func (s *SendChatMessageShrinkRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendChatMessageShrinkRequest) GetDirectMaxHistoryRounds() interface{} {
	return s.DirectMaxHistoryRounds
}

func (s *SendChatMessageShrinkRequest) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendChatMessageShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *SendChatMessageShrinkRequest) GetLocale() *string {
	return s.Locale
}

func (s *SendChatMessageShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *SendChatMessageShrinkRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendChatMessageShrinkRequest) GetSegmentsShrink() *string {
	return s.SegmentsShrink
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

func (s *SendChatMessageShrinkRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendChatMessageShrinkRequest) SetContent(v string) *SendChatMessageShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetContentType(v string) *SendChatMessageShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetContextItemsShrink(v string) *SendChatMessageShrinkRequest {
	s.ContextItemsShrink = &v
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

func (s *SendChatMessageShrinkRequest) SetDirectMaxHistoryRounds(v interface{}) *SendChatMessageShrinkRequest {
	s.DirectMaxHistoryRounds = v
	return s
}

func (s *SendChatMessageShrinkRequest) SetEnableWebSearch(v bool) *SendChatMessageShrinkRequest {
	s.EnableWebSearch = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetFilesShrink(v string) *SendChatMessageShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *SendChatMessageShrinkRequest) SetLocale(v string) *SendChatMessageShrinkRequest {
	s.Locale = &v
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

func (s *SendChatMessageShrinkRequest) SetSegmentsShrink(v string) *SendChatMessageShrinkRequest {
	s.SegmentsShrink = &v
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

func (s *SendChatMessageShrinkRequest) SetWorkMode(v string) *SendChatMessageShrinkRequest {
	s.WorkMode = &v
	return s
}

func (s *SendChatMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
