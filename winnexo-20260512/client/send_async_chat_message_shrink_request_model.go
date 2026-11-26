// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncChatMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SendAsyncChatMessageShrinkRequest
	GetAppToken() *string
	SetAutoPreference(v string) *SendAsyncChatMessageShrinkRequest
	GetAutoPreference() *string
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
	SetOpsUserCode(v string) *SendAsyncChatMessageShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SendAsyncChatMessageShrinkRequest
	GetOpsUserId() *string
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
	SetThinkingEnabled(v bool) *SendAsyncChatMessageShrinkRequest
	GetThinkingEnabled() *bool
	SetThinkingLevel(v string) *SendAsyncChatMessageShrinkRequest
	GetThinkingLevel() *string
	SetWorkMode(v string) *SendAsyncChatMessageShrinkRequest
	GetWorkMode() *string
}

type SendAsyncChatMessageShrinkRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// AUTO 偏好：none=默认、performance=效果优先、credit=成本优先；缺省或 null 继承会话，空白重置为 none
	//
	// example:
	//
	// credit
	AutoPreference *string `json:"autoPreference,omitempty" xml:"autoPreference,omitempty"`
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
	// The references and annotations for the current turn.
	//
	// example:
	//
	// []
	ContextItemsShrink *string `json:"contextItems,omitempty" xml:"contextItems,omitempty"`
	// The list of digital employee names. For compatibility with the old format, a single string can also be passed.
	//
	// example:
	//
	// ["pcitc-magent"]
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// **[Deprecated]*	- Use workMode=direct instead. If set to true, winnexo_chat is retained and processed in request-level direct mode.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// The maximum number of history rounds referenced in the current turn for the direct mode. Valid values: integers from 0 to 100. Default value: 0. Invalid values fall back to the original single-turn path. This parameter is ignored in non-direct modes.
	//
	// example:
	//
	// 1
	DirectMaxHistoryRounds interface{} `json:"directMaxHistoryRounds,omitempty" xml:"directMaxHistoryRounds,omitempty"`
	// Specifies whether to enable web search. Default value: false. In the task execution scenario (when taskExecution is passed), the task configuration prevails.
	//
	// example:
	//
	// false
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// The list of file references. Each item is an object, and fileId is required (returned by uploadChatFile).
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
	// The language for the current turn. Valid values: zh-CN and en-US. This parameter is used for API error messages and passed through to CodeAgent. If omitted, the current language behavior is retained and user settings are not modified.
	//
	// example:
	//
	// en-US
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The model tier. Valid values: quick, standard, flagship, auto, or other currently available model names (case-insensitive). If not specified, the model tier of the session is inherited. For a new session, the default value is standard.
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// Specifies whether to reuse the most recent session under the digital employee when sessionId is not provided (CLI scenario). Default value: false, which means a new session is created.
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
	// Specifies whether to use streaming generation. This API always generates background content in streaming mode and writes it to the message stream. The value of this parameter does not change the response structure.
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
	// 本轮显式覆盖是否思考；AUTO 未传时由决策模型判断，固定档未传时沿用个人设置，最终遵守模型能力与租户开关
	//
	// example:
	//
	// true
	ThinkingEnabled *bool `json:"thinkingEnabled,omitempty" xml:"thinkingEnabled,omitempty"`
	// 本轮显式覆盖思考强度；AUTO 未传时由决策模型判断，固定档未传时沿用个人设置，最终遵守模型能力与租户开关
	//
	// example:
	//
	// high
	ThinkingLevel *string `json:"thinkingLevel,omitempty" xml:"thinkingLevel,omitempty"`
	// The session work mode. Valid values: ask (quick Q&A), work (deep work, default), and direct (direct connection mode, which does not start the sandbox, uses independent direct history controlled by directMaxHistoryRounds, and is equivalent to directChat=true). The ask and work modes are session-level. For continued chats, the session mode is inherited by default. To switch modes, create a new session or fork. In multi-digital employee or task execution scenarios, if ask is passed, it takes effect as work. If directChat is set to true, this parameter is ignored.
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

func (s *SendAsyncChatMessageShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SendAsyncChatMessageShrinkRequest) GetAutoPreference() *string {
	return s.AutoPreference
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

func (s *SendAsyncChatMessageShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SendAsyncChatMessageShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *SendAsyncChatMessageShrinkRequest) GetThinkingEnabled() *bool {
	return s.ThinkingEnabled
}

func (s *SendAsyncChatMessageShrinkRequest) GetThinkingLevel() *string {
	return s.ThinkingLevel
}

func (s *SendAsyncChatMessageShrinkRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendAsyncChatMessageShrinkRequest) SetAppToken(v string) *SendAsyncChatMessageShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetAutoPreference(v string) *SendAsyncChatMessageShrinkRequest {
	s.AutoPreference = &v
	return s
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

func (s *SendAsyncChatMessageShrinkRequest) SetOpsUserCode(v string) *SendAsyncChatMessageShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetOpsUserId(v string) *SendAsyncChatMessageShrinkRequest {
	s.OpsUserId = &v
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

func (s *SendAsyncChatMessageShrinkRequest) SetThinkingEnabled(v bool) *SendAsyncChatMessageShrinkRequest {
	s.ThinkingEnabled = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetThinkingLevel(v string) *SendAsyncChatMessageShrinkRequest {
	s.ThinkingLevel = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) SetWorkMode(v string) *SendAsyncChatMessageShrinkRequest {
	s.WorkMode = &v
	return s
}

func (s *SendAsyncChatMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
