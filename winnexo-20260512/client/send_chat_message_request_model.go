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
	SetContextItems(v []*SendChatMessageRequestContextItems) *SendChatMessageRequest
	GetContextItems() []*SendChatMessageRequestContextItems
	SetDigitalEmployeeName(v []*string) *SendChatMessageRequest
	GetDigitalEmployeeName() []*string
	SetDirectChat(v bool) *SendChatMessageRequest
	GetDirectChat() *bool
	SetDirectMaxHistoryRounds(v interface{}) *SendChatMessageRequest
	GetDirectMaxHistoryRounds() interface{}
	SetEnableWebSearch(v bool) *SendChatMessageRequest
	GetEnableWebSearch() *bool
	SetFiles(v []*SendChatMessageRequestFiles) *SendChatMessageRequest
	GetFiles() []*SendChatMessageRequestFiles
	SetLocale(v string) *SendChatMessageRequest
	GetLocale() *string
	SetModel(v string) *SendChatMessageRequest
	GetModel() *string
	SetReuseLastSession(v bool) *SendChatMessageRequest
	GetReuseLastSession() *bool
	SetSegments(v []*SendChatMessageRequestSegments) *SendChatMessageRequest
	GetSegments() []*SendChatMessageRequestSegments
	SetSessionId(v string) *SendChatMessageRequest
	GetSessionId() *string
	SetStream(v bool) *SendChatMessageRequest
	GetStream() *bool
	SetTaskExecution(v *SendChatMessageRequestTaskExecution) *SendChatMessageRequest
	GetTaskExecution() *SendChatMessageRequestTaskExecution
	SetTenantId(v string) *SendChatMessageRequest
	GetTenantId() *string
	SetWorkMode(v string) *SendChatMessageRequest
	GetWorkMode() *string
}

type SendChatMessageRequest struct {
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
	// 作为本轮上下文引用的消息或文件条目
	//
	// example:
	//
	// []
	ContextItems []*SendChatMessageRequestContextItems `json:"contextItems,omitempty" xml:"contextItems,omitempty" type:"Repeated"`
	// The list of digital employee names. For compatibility with the legacy format, a single string can be passed.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName []*string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty" type:"Repeated"`
	// Specifies whether to enable the direct connection mode. If this parameter is set to true, the regular scenario routing is skipped and the direct connection conversation scenario is directly entered.
	//
	// example:
	//
	// false
	DirectChat *bool `json:"directChat,omitempty" xml:"directChat,omitempty"`
	// 本轮携带的最近完整 direct 问答轮数，不含本轮输入。仅 workMode=direct 或 directChat=true 时生效。有效值为 0–100 的 JSON 整数，省略时由后端按 0 处理，不携带历史但仍保存当前轮。null、布尔值、字符串、小数及越界值按原 direct 路径处理，网关保留原始值供后端判定。work/ask 忽略此字段。该参数不限制转换为 work 时的历史范围
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
	Files []*SendChatMessageRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// 本轮语言，可选 zh-CN、en-US，用于接口错误文案并透传给 CodeAgent；省略时沿用当前语言行为，不修改用户设置
	//
	// example:
	//
	// en-US
	Locale *string `json:"locale,omitempty" xml:"locale,omitempty"`
	// The abstract model tier. Valid values: quick, standard, and flagship. If this parameter is not specified, standard is used for new sessions, and the current tier is used for existing sessions.
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
	// 消息中的结构化片段，例如会话或对象提及
	//
	// example:
	//
	// []
	Segments []*SendChatMessageRequestSegments `json:"segments,omitempty" xml:"segments,omitempty" type:"Repeated"`
	// The session ID.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to enable streaming output.
	//
	// example:
	//
	// true
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// The task execution metadata returned by executeScheduledTask. If this parameter is passed, the request is processed according to the task execution link.
	TaskExecution *SendChatMessageRequestTaskExecution `json:"taskExecution,omitempty" xml:"taskExecution,omitempty" type:"Struct"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The session work mode. Valid values: ask (quick Q&A, which trims tools, skills, and connectors for single-turn direct answers), work (deep work, which is the default value), and direct (direct connection mode at the request level, which does not start the sandbox and prevents context pollution, equivalent to directChat=true). The ask and work modes are at the session level. The mode is fixed when a new session is created. Subsequent chats inherit the session mode by default. If the explicitly passed value is inconsistent with the session mode, a parameter error is returned. To switch the mode, create a new session or fork the session. If ask is passed in the multi-digital employee or task execution scenario, it takes effect as work. If directChat is set to true, this parameter is ignored.
	//
	// example:
	//
	// work
	WorkMode *string `json:"workMode,omitempty" xml:"workMode,omitempty"`
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

func (s *SendChatMessageRequest) GetContextItems() []*SendChatMessageRequestContextItems {
	return s.ContextItems
}

func (s *SendChatMessageRequest) GetDigitalEmployeeName() []*string {
	return s.DigitalEmployeeName
}

func (s *SendChatMessageRequest) GetDirectChat() *bool {
	return s.DirectChat
}

func (s *SendChatMessageRequest) GetDirectMaxHistoryRounds() interface{} {
	return s.DirectMaxHistoryRounds
}

func (s *SendChatMessageRequest) GetEnableWebSearch() *bool {
	return s.EnableWebSearch
}

func (s *SendChatMessageRequest) GetFiles() []*SendChatMessageRequestFiles {
	return s.Files
}

func (s *SendChatMessageRequest) GetLocale() *string {
	return s.Locale
}

func (s *SendChatMessageRequest) GetModel() *string {
	return s.Model
}

func (s *SendChatMessageRequest) GetReuseLastSession() *bool {
	return s.ReuseLastSession
}

func (s *SendChatMessageRequest) GetSegments() []*SendChatMessageRequestSegments {
	return s.Segments
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

func (s *SendChatMessageRequest) GetWorkMode() *string {
	return s.WorkMode
}

func (s *SendChatMessageRequest) SetContent(v string) *SendChatMessageRequest {
	s.Content = &v
	return s
}

func (s *SendChatMessageRequest) SetContentType(v string) *SendChatMessageRequest {
	s.ContentType = &v
	return s
}

func (s *SendChatMessageRequest) SetContextItems(v []*SendChatMessageRequestContextItems) *SendChatMessageRequest {
	s.ContextItems = v
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

func (s *SendChatMessageRequest) SetDirectMaxHistoryRounds(v interface{}) *SendChatMessageRequest {
	s.DirectMaxHistoryRounds = v
	return s
}

func (s *SendChatMessageRequest) SetEnableWebSearch(v bool) *SendChatMessageRequest {
	s.EnableWebSearch = &v
	return s
}

func (s *SendChatMessageRequest) SetFiles(v []*SendChatMessageRequestFiles) *SendChatMessageRequest {
	s.Files = v
	return s
}

func (s *SendChatMessageRequest) SetLocale(v string) *SendChatMessageRequest {
	s.Locale = &v
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

func (s *SendChatMessageRequest) SetSegments(v []*SendChatMessageRequestSegments) *SendChatMessageRequest {
	s.Segments = v
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

func (s *SendChatMessageRequest) SetWorkMode(v string) *SendChatMessageRequest {
	s.WorkMode = &v
	return s
}

func (s *SendChatMessageRequest) Validate() error {
	if s.ContextItems != nil {
		for _, item := range s.ContextItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Segments != nil {
		for _, item := range s.Segments {
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

type SendChatMessageRequestContextItems struct {
	// 用户对引用内容的批注
	//
	// example:
	//
	// 请核对华东区数据
	Annotation *string `json:"annotation,omitempty" xml:"annotation,omitempty"`
	// 批注关联的文件列表
	//
	// example:
	//
	// []
	AnnotationFiles []*SendChatMessageRequestContextItemsAnnotationFiles `json:"annotationFiles,omitempty" xml:"annotationFiles,omitempty" type:"Repeated"`
	// 被引用内容正文
	//
	// example:
	//
	// 本月销售收入
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 被引用文件 ID
	//
	// example:
	//
	// exampleFileId
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// 被引用文件名称
	//
	// example:
	//
	// 销售数据.xlsx
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 被引用文件类型
	//
	// example:
	//
	// xlsx
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// 被引用消息 ID
	//
	// example:
	//
	// exampleMessageId
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// 消息发送者名称
	//
	// example:
	//
	// 销售助手
	SenderName *string `json:"senderName,omitempty" xml:"senderName,omitempty"`
	// 上下文条目类型
	//
	// example:
	//
	// message_annotation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendChatMessageRequestContextItems) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestContextItems) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestContextItems) GetAnnotation() *string {
	return s.Annotation
}

func (s *SendChatMessageRequestContextItems) GetAnnotationFiles() []*SendChatMessageRequestContextItemsAnnotationFiles {
	return s.AnnotationFiles
}

func (s *SendChatMessageRequestContextItems) GetContent() *string {
	return s.Content
}

func (s *SendChatMessageRequestContextItems) GetFileId() *string {
	return s.FileId
}

func (s *SendChatMessageRequestContextItems) GetFileName() *string {
	return s.FileName
}

func (s *SendChatMessageRequestContextItems) GetFileType() *string {
	return s.FileType
}

func (s *SendChatMessageRequestContextItems) GetMessageId() *string {
	return s.MessageId
}

func (s *SendChatMessageRequestContextItems) GetSenderName() *string {
	return s.SenderName
}

func (s *SendChatMessageRequestContextItems) GetType() *string {
	return s.Type
}

func (s *SendChatMessageRequestContextItems) SetAnnotation(v string) *SendChatMessageRequestContextItems {
	s.Annotation = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetAnnotationFiles(v []*SendChatMessageRequestContextItemsAnnotationFiles) *SendChatMessageRequestContextItems {
	s.AnnotationFiles = v
	return s
}

func (s *SendChatMessageRequestContextItems) SetContent(v string) *SendChatMessageRequestContextItems {
	s.Content = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetFileId(v string) *SendChatMessageRequestContextItems {
	s.FileId = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetFileName(v string) *SendChatMessageRequestContextItems {
	s.FileName = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetFileType(v string) *SendChatMessageRequestContextItems {
	s.FileType = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetMessageId(v string) *SendChatMessageRequestContextItems {
	s.MessageId = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetSenderName(v string) *SendChatMessageRequestContextItems {
	s.SenderName = &v
	return s
}

func (s *SendChatMessageRequestContextItems) SetType(v string) *SendChatMessageRequestContextItems {
	s.Type = &v
	return s
}

func (s *SendChatMessageRequestContextItems) Validate() error {
	if s.AnnotationFiles != nil {
		for _, item := range s.AnnotationFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SendChatMessageRequestContextItemsAnnotationFiles struct {
	// 文件名称
	//
	// example:
	//
	// 补充说明.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件公开访问地址
	//
	// example:
	//
	// https://example.com/public/file
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 文件记录 ID
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// 文件内部访问地址
	//
	// example:
	//
	// https://example.com/file
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
}

func (s SendChatMessageRequestContextItemsAnnotationFiles) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestContextItemsAnnotationFiles) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) GetFileName() *string {
	return s.FileName
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) SetFileName(v string) *SendChatMessageRequestContextItemsAnnotationFiles {
	s.FileName = &v
	return s
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) SetFilePublicUrl(v string) *SendChatMessageRequestContextItemsAnnotationFiles {
	s.FilePublicUrl = &v
	return s
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) SetFileRecordId(v string) *SendChatMessageRequestContextItemsAnnotationFiles {
	s.FileRecordId = &v
	return s
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) SetFileUrl(v string) *SendChatMessageRequestContextItemsAnnotationFiles {
	s.FileUrl = &v
	return s
}

func (s *SendChatMessageRequestContextItemsAnnotationFiles) Validate() error {
	return dara.Validate(s)
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

type SendChatMessageRequestSegments struct {
	// 片段展示内容
	//
	// example:
	//
	// 旧会话
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 提及对象类别
	//
	// example:
	//
	// conversations
	MentionType *string `json:"mentionType,omitempty" xml:"mentionType,omitempty"`
	// 被引用对象名称
	//
	// example:
	//
	// 旧会话
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 被引用对象 ID
	//
	// example:
	//
	// exampleObjectId
	ObjectId *string `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// 被引用对象类型
	//
	// example:
	//
	// conversation
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// 会话类型
	//
	// example:
	//
	// personal
	SessionType *string `json:"sessionType,omitempty" xml:"sessionType,omitempty"`
	// 片段类型
	//
	// example:
	//
	// mention
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SendChatMessageRequestSegments) String() string {
	return dara.Prettify(s)
}

func (s SendChatMessageRequestSegments) GoString() string {
	return s.String()
}

func (s *SendChatMessageRequestSegments) GetContent() *string {
	return s.Content
}

func (s *SendChatMessageRequestSegments) GetMentionType() *string {
	return s.MentionType
}

func (s *SendChatMessageRequestSegments) GetName() *string {
	return s.Name
}

func (s *SendChatMessageRequestSegments) GetObjectId() *string {
	return s.ObjectId
}

func (s *SendChatMessageRequestSegments) GetObjectType() *string {
	return s.ObjectType
}

func (s *SendChatMessageRequestSegments) GetSessionType() *string {
	return s.SessionType
}

func (s *SendChatMessageRequestSegments) GetType() *string {
	return s.Type
}

func (s *SendChatMessageRequestSegments) SetContent(v string) *SendChatMessageRequestSegments {
	s.Content = &v
	return s
}

func (s *SendChatMessageRequestSegments) SetMentionType(v string) *SendChatMessageRequestSegments {
	s.MentionType = &v
	return s
}

func (s *SendChatMessageRequestSegments) SetName(v string) *SendChatMessageRequestSegments {
	s.Name = &v
	return s
}

func (s *SendChatMessageRequestSegments) SetObjectId(v string) *SendChatMessageRequestSegments {
	s.ObjectId = &v
	return s
}

func (s *SendChatMessageRequestSegments) SetObjectType(v string) *SendChatMessageRequestSegments {
	s.ObjectType = &v
	return s
}

func (s *SendChatMessageRequestSegments) SetSessionType(v string) *SendChatMessageRequestSegments {
	s.SessionType = &v
	return s
}

func (s *SendChatMessageRequestSegments) SetType(v string) *SendChatMessageRequestSegments {
	s.Type = &v
	return s
}

func (s *SendChatMessageRequestSegments) Validate() error {
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
	// The name of the digital employee.
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
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the user to which the task belongs.
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
