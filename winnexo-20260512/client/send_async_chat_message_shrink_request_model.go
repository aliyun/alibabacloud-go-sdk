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
	// 用户消息正文
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 消息类型：Text / Markdown
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
	// 会话ID，不传则新建会话
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// 是否流式生成；本接口固定按流式生成后台内容并写入消息流，取值不改变返回结构
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
