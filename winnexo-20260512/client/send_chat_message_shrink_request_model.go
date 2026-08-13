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
	// 用户消息正文
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息类型: Text / Markdown
	//
	// example:
	//
	// Text
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 数字员工名称列表（兼容旧格式可传单个字符串）
	//
	// example:
	//
	// string_value
	DigitalEmployeeNameShrink *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// 是否启用直连模式；true 时跳过常规场景路由，直接进入直连对话场景
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// 文件引用列表；每项为对象，fileId 必传（由 uploadChatFile 返回）
	FilesShrink *string `json:"files,omitempty" xml:"files,omitempty"`
	// 抽象模型档位（quick / standard / flagship）；缺省时新会话用 standard，已有会话沿用会话当前档位
	//
	// example:
	//
	// quick
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 不传 sessionId 时是否复用该数字员工下最近一个会话（CLI 场景），缺省 false 即新建会话
	//
	// example:
	//
	// false
	ReuseLastSession *bool `json:"reuseLastSession,omitempty" xml:"reuseLastSession,omitempty"`
	// 会话 ID
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 是否流式返回，默认True
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// executeScheduledTask 返回的任务执行元数据；传入后按任务执行链路处理
	TaskExecutionShrink *string `json:"taskExecution,omitempty" xml:"taskExecution,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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
