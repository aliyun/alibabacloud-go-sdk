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
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// 是否启用直连模式；true 时跳过常规场景路由，直接进入直连对话场景
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// 文件引用列表；每项为对象，fileId 必传（由 uploadChatFile 返回）
	Files []*SendAsyncChatMessageRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
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
	TaskExecution *SendAsyncChatMessageRequestTaskExecution `json:"taskExecution,omitempty" xml:"taskExecution,omitempty" type:"Struct"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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
	// 文件 ID，由 uploadChatFile 返回
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleFileId
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 文件类型
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
	// 计费 ID
	//
	// example:
	//
	// exampleBillingId
	BillingId *string `json:"billingId,omitempty" xml:"billingId,omitempty"`
	// 是否启用联网搜索
	//
	// example:
	//
	// true
	EnableWebSearch *bool `json:"enableWebSearch,omitempty" xml:"enableWebSearch,omitempty"`
	// 执行记录 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleExecutionId
	ExecutionId *string `json:"executionId,omitempty" xml:"executionId,omitempty"`
	// 数字员工名称
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// 关联技能编码列表
	//
	// example:
	//
	// string_value
	SkillCodes []*string `json:"skillCodes,omitempty" xml:"skillCodes,omitempty" type:"Repeated"`
	// 任务 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// 任务名称
	//
	// example:
	//
	// string_value
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// 任务理解内容
	//
	// example:
	//
	// string_value
	TaskUnderstand *string `json:"taskUnderstand,omitempty" xml:"taskUnderstand,omitempty"`
	// 任务所属租户 ID
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 任务所属用户 ID
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
