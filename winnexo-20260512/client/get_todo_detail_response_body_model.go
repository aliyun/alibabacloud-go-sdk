// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssigneeCount(v int64) *GetTodoDetailResponseBody
	GetAssigneeCount() *int64
	SetAssignees(v []*GetTodoDetailResponseBodyAssignees) *GetTodoDetailResponseBody
	GetAssignees() []*GetTodoDetailResponseBodyAssignees
	SetAttachmentPolicy(v map[string]interface{}) *GetTodoDetailResponseBody
	GetAttachmentPolicy() map[string]interface{}
	SetCode(v string) *GetTodoDetailResponseBody
	GetCode() *string
	SetCompletedAt(v string) *GetTodoDetailResponseBody
	GetCompletedAt() *string
	SetCompletedByUserId(v string) *GetTodoDetailResponseBody
	GetCompletedByUserId() *string
	SetCompletedByUserName(v string) *GetTodoDetailResponseBody
	GetCompletedByUserName() *string
	SetGmtCreate(v string) *GetTodoDetailResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *GetTodoDetailResponseBody
	GetGmtModified() *string
	SetIsRead(v bool) *GetTodoDetailResponseBody
	GetIsRead() *bool
	SetMessage(v string) *GetTodoDetailResponseBody
	GetMessage() *string
	SetMessageId(v string) *GetTodoDetailResponseBody
	GetMessageId() *string
	SetMyAssigneeStatus(v string) *GetTodoDetailResponseBody
	GetMyAssigneeStatus() *string
	SetMyDisplayState(v string) *GetTodoDetailResponseBody
	GetMyDisplayState() *string
	SetMyPermissions(v map[string]interface{}) *GetTodoDetailResponseBody
	GetMyPermissions() map[string]interface{}
	SetPriority(v string) *GetTodoDetailResponseBody
	GetPriority() *string
	SetQuestions(v []*GetTodoDetailResponseBodyQuestions) *GetTodoDetailResponseBody
	GetQuestions() []*GetTodoDetailResponseBodyQuestions
	SetRequestId(v string) *GetTodoDetailResponseBody
	GetRequestId() *string
	SetResultPayload(v map[string]interface{}) *GetTodoDetailResponseBody
	GetResultPayload() map[string]interface{}
	SetSourceContext(v map[string]interface{}) *GetTodoDetailResponseBody
	GetSourceContext() map[string]interface{}
	SetSourceType(v string) *GetTodoDetailResponseBody
	GetSourceType() *string
	SetSourceWorkflowId(v string) *GetTodoDetailResponseBody
	GetSourceWorkflowId() *string
	SetStatus(v string) *GetTodoDetailResponseBody
	GetStatus() *string
	SetSummary(v string) *GetTodoDetailResponseBody
	GetSummary() *string
	SetTitle(v string) *GetTodoDetailResponseBody
	GetTitle() *string
	SetTodoId(v string) *GetTodoDetailResponseBody
	GetTodoId() *string
}

type GetTodoDetailResponseBody struct {
	// 处理人数量
	//
	// example:
	//
	// 2
	AssigneeCount *int64 `json:"assigneeCount,omitempty" xml:"assigneeCount,omitempty"`
	// 全部处理人
	Assignees []*GetTodoDetailResponseBodyAssignees `json:"assignees,omitempty" xml:"assignees,omitempty" type:"Repeated"`
	// 附件策略（当前版本附件不落库）
	//
	// example:
	//
	// allowUpload=true
	AttachmentPolicy map[string]interface{} `json:"attachmentPolicy,omitempty" xml:"attachmentPolicy,omitempty"`
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 完成时间（北京时间 YYYY-MM-DD HH:mm:ss）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	CompletedAt *string `json:"completedAt,omitempty" xml:"completedAt,omitempty"`
	// 完成人用户ID；未完成为空
	//
	// example:
	//
	// 42
	CompletedByUserId *string `json:"completedByUserId,omitempty" xml:"completedByUserId,omitempty"`
	// 完成人姓名；未完成为空
	//
	// example:
	//
	// 张三
	CompletedByUserName *string `json:"completedByUserName,omitempty" xml:"completedByUserName,omitempty"`
	// 创建时间（北京时间 YYYY-MM-DD HH:mm:ss）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 最后修改时间（北京时间 YYYY-MM-DD HH:mm:ss）
	//
	// example:
	//
	// 2026-10-12 10:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// 调用者是否已读
	//
	// example:
	//
	// false
	IsRead *bool `json:"isRead,omitempty" xml:"isRead,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 发出该待办的 assistant 消息ID；无会话来源为空
	//
	// example:
	//
	// msg-1
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// 调用者的处理人行状态
	//
	// example:
	//
	// PENDING
	MyAssigneeStatus *string `json:"myAssigneeStatus,omitempty" xml:"myAssigneeStatus,omitempty"`
	// 调用者的列表展示状态
	//
	// example:
	//
	// TODO
	MyDisplayState *string `json:"myDisplayState,omitempty" xml:"myDisplayState,omitempty"`
	// 调用者权限（canSubmit/canTransfer）
	//
	// example:
	//
	// APP/VIEW=true
	MyPermissions map[string]interface{} `json:"myPermissions,omitempty" xml:"myPermissions,omitempty"`
	// 优先级
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 问题清单（结构化投影）
	Questions []*GetTodoDetailResponseBodyQuestions `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 作答内容（JSONB 原样；未提交为空对象缺省）
	//
	// example:
	//
	// result=同意
	ResultPayload map[string]interface{} `json:"resultPayload,omitempty" xml:"resultPayload,omitempty"`
	// 来源上下文（创建时原样透传）
	//
	// example:
	//
	// source=openapi
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// 待办来源类型
	//
	// example:
	//
	// WORKFLOW
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 工作流ID（sourceType=WORKFLOW 时必有）
	//
	// example:
	//
	// wf-1
	SourceWorkflowId *string `json:"sourceWorkflowId,omitempty" xml:"sourceWorkflowId,omitempty"`
	// 待办主状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 待办摘要
	//
	// example:
	//
	// 请审批以下事项
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 待办标题
	//
	// example:
	//
	// 治理审批
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 待办ID
	//
	// example:
	//
	// todo_abc123
	TodoId *string `json:"todoId,omitempty" xml:"todoId,omitempty"`
}

func (s GetTodoDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTodoDetailResponseBody) GetAssigneeCount() *int64 {
	return s.AssigneeCount
}

func (s *GetTodoDetailResponseBody) GetAssignees() []*GetTodoDetailResponseBodyAssignees {
	return s.Assignees
}

func (s *GetTodoDetailResponseBody) GetAttachmentPolicy() map[string]interface{} {
	return s.AttachmentPolicy
}

func (s *GetTodoDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTodoDetailResponseBody) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *GetTodoDetailResponseBody) GetCompletedByUserId() *string {
	return s.CompletedByUserId
}

func (s *GetTodoDetailResponseBody) GetCompletedByUserName() *string {
	return s.CompletedByUserName
}

func (s *GetTodoDetailResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTodoDetailResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTodoDetailResponseBody) GetIsRead() *bool {
	return s.IsRead
}

func (s *GetTodoDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTodoDetailResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *GetTodoDetailResponseBody) GetMyAssigneeStatus() *string {
	return s.MyAssigneeStatus
}

func (s *GetTodoDetailResponseBody) GetMyDisplayState() *string {
	return s.MyDisplayState
}

func (s *GetTodoDetailResponseBody) GetMyPermissions() map[string]interface{} {
	return s.MyPermissions
}

func (s *GetTodoDetailResponseBody) GetPriority() *string {
	return s.Priority
}

func (s *GetTodoDetailResponseBody) GetQuestions() []*GetTodoDetailResponseBodyQuestions {
	return s.Questions
}

func (s *GetTodoDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTodoDetailResponseBody) GetResultPayload() map[string]interface{} {
	return s.ResultPayload
}

func (s *GetTodoDetailResponseBody) GetSourceContext() map[string]interface{} {
	return s.SourceContext
}

func (s *GetTodoDetailResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetTodoDetailResponseBody) GetSourceWorkflowId() *string {
	return s.SourceWorkflowId
}

func (s *GetTodoDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTodoDetailResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *GetTodoDetailResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetTodoDetailResponseBody) GetTodoId() *string {
	return s.TodoId
}

func (s *GetTodoDetailResponseBody) SetAssigneeCount(v int64) *GetTodoDetailResponseBody {
	s.AssigneeCount = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetAssignees(v []*GetTodoDetailResponseBodyAssignees) *GetTodoDetailResponseBody {
	s.Assignees = v
	return s
}

func (s *GetTodoDetailResponseBody) SetAttachmentPolicy(v map[string]interface{}) *GetTodoDetailResponseBody {
	s.AttachmentPolicy = v
	return s
}

func (s *GetTodoDetailResponseBody) SetCode(v string) *GetTodoDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetCompletedAt(v string) *GetTodoDetailResponseBody {
	s.CompletedAt = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetCompletedByUserId(v string) *GetTodoDetailResponseBody {
	s.CompletedByUserId = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetCompletedByUserName(v string) *GetTodoDetailResponseBody {
	s.CompletedByUserName = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetGmtCreate(v string) *GetTodoDetailResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetGmtModified(v string) *GetTodoDetailResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetIsRead(v bool) *GetTodoDetailResponseBody {
	s.IsRead = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetMessage(v string) *GetTodoDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetMessageId(v string) *GetTodoDetailResponseBody {
	s.MessageId = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetMyAssigneeStatus(v string) *GetTodoDetailResponseBody {
	s.MyAssigneeStatus = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetMyDisplayState(v string) *GetTodoDetailResponseBody {
	s.MyDisplayState = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetMyPermissions(v map[string]interface{}) *GetTodoDetailResponseBody {
	s.MyPermissions = v
	return s
}

func (s *GetTodoDetailResponseBody) SetPriority(v string) *GetTodoDetailResponseBody {
	s.Priority = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetQuestions(v []*GetTodoDetailResponseBodyQuestions) *GetTodoDetailResponseBody {
	s.Questions = v
	return s
}

func (s *GetTodoDetailResponseBody) SetRequestId(v string) *GetTodoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetResultPayload(v map[string]interface{}) *GetTodoDetailResponseBody {
	s.ResultPayload = v
	return s
}

func (s *GetTodoDetailResponseBody) SetSourceContext(v map[string]interface{}) *GetTodoDetailResponseBody {
	s.SourceContext = v
	return s
}

func (s *GetTodoDetailResponseBody) SetSourceType(v string) *GetTodoDetailResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetSourceWorkflowId(v string) *GetTodoDetailResponseBody {
	s.SourceWorkflowId = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetStatus(v string) *GetTodoDetailResponseBody {
	s.Status = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetSummary(v string) *GetTodoDetailResponseBody {
	s.Summary = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetTitle(v string) *GetTodoDetailResponseBody {
	s.Title = &v
	return s
}

func (s *GetTodoDetailResponseBody) SetTodoId(v string) *GetTodoDetailResponseBody {
	s.TodoId = &v
	return s
}

func (s *GetTodoDetailResponseBody) Validate() error {
	if s.Assignees != nil {
		for _, item := range s.Assignees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Questions != nil {
		for _, item := range s.Questions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTodoDetailResponseBodyAssignees struct {
	// 处理时间（北京时间）；未处理为空
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	ActedAt *string `json:"actedAt,omitempty" xml:"actedAt,omitempty"`
	// 处理人来源：ORIGINAL/TRANSFERRED
	//
	// example:
	//
	// ORIGINAL
	AssigneeSource *string `json:"assigneeSource,omitempty" xml:"assigneeSource,omitempty"`
	// 处理人用户ID
	//
	// example:
	//
	// 42
	AssigneeUserId *string `json:"assigneeUserId,omitempty" xml:"assigneeUserId,omitempty"`
	// 已读时间（北京时间）；未读为空
	//
	// example:
	//
	// 2026-10-12 10:01:00
	ReadAt *string `json:"readAt,omitempty" xml:"readAt,omitempty"`
	// 处理人行状态
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 转出人用户ID；仅转办产生
	//
	// example:
	//
	// 43
	TransferredFromUserId *string `json:"transferredFromUserId,omitempty" xml:"transferredFromUserId,omitempty"`
}

func (s GetTodoDetailResponseBodyAssignees) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailResponseBodyAssignees) GoString() string {
	return s.String()
}

func (s *GetTodoDetailResponseBodyAssignees) GetActedAt() *string {
	return s.ActedAt
}

func (s *GetTodoDetailResponseBodyAssignees) GetAssigneeSource() *string {
	return s.AssigneeSource
}

func (s *GetTodoDetailResponseBodyAssignees) GetAssigneeUserId() *string {
	return s.AssigneeUserId
}

func (s *GetTodoDetailResponseBodyAssignees) GetReadAt() *string {
	return s.ReadAt
}

func (s *GetTodoDetailResponseBodyAssignees) GetStatus() *string {
	return s.Status
}

func (s *GetTodoDetailResponseBodyAssignees) GetTransferredFromUserId() *string {
	return s.TransferredFromUserId
}

func (s *GetTodoDetailResponseBodyAssignees) SetActedAt(v string) *GetTodoDetailResponseBodyAssignees {
	s.ActedAt = &v
	return s
}

func (s *GetTodoDetailResponseBodyAssignees) SetAssigneeSource(v string) *GetTodoDetailResponseBodyAssignees {
	s.AssigneeSource = &v
	return s
}

func (s *GetTodoDetailResponseBodyAssignees) SetAssigneeUserId(v string) *GetTodoDetailResponseBodyAssignees {
	s.AssigneeUserId = &v
	return s
}

func (s *GetTodoDetailResponseBodyAssignees) SetReadAt(v string) *GetTodoDetailResponseBodyAssignees {
	s.ReadAt = &v
	return s
}

func (s *GetTodoDetailResponseBodyAssignees) SetStatus(v string) *GetTodoDetailResponseBodyAssignees {
	s.Status = &v
	return s
}

func (s *GetTodoDetailResponseBodyAssignees) SetTransferredFromUserId(v string) *GetTodoDetailResponseBodyAssignees {
	s.TransferredFromUserId = &v
	return s
}

func (s *GetTodoDetailResponseBodyAssignees) Validate() error {
	return dara.Validate(s)
}

type GetTodoDetailResponseBodyQuestions struct {
	// 是否允许自定义输入
	//
	// example:
	//
	// false
	AllowCustom *bool `json:"allowCustom,omitempty" xml:"allowCustom,omitempty"`
	// 自定义输入项文案
	//
	// example:
	//
	// 其他
	CustomLabel *string `json:"customLabel,omitempty" xml:"customLabel,omitempty"`
	// 自定义输入长度上限
	//
	// example:
	//
	// 200
	CustomMaxLength *int64 `json:"customMaxLength,omitempty" xml:"customMaxLength,omitempty"`
	// 自定义输入占位符
	//
	// example:
	//
	// 请说明
	CustomPlaceholder *string `json:"customPlaceholder,omitempty" xml:"customPlaceholder,omitempty"`
	// TEXT 题长度上限
	//
	// example:
	//
	// 500
	MaxLength *int64 `json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	// 作答形态：SINGLE_CHOICE/MULTI_CHOICE/TEXT
	//
	// example:
	//
	// SINGLE_CHOICE
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 选项（选择类必有）
	Options []*GetTodoDetailResponseBodyQuestionsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// TEXT 题占位符
	//
	// example:
	//
	// 请输入
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 问题正文（引用键）
	//
	// example:
	//
	// 是否通过审批
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// 是否必答
	//
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s GetTodoDetailResponseBodyQuestions) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailResponseBodyQuestions) GoString() string {
	return s.String()
}

func (s *GetTodoDetailResponseBodyQuestions) GetAllowCustom() *bool {
	return s.AllowCustom
}

func (s *GetTodoDetailResponseBodyQuestions) GetCustomLabel() *string {
	return s.CustomLabel
}

func (s *GetTodoDetailResponseBodyQuestions) GetCustomMaxLength() *int64 {
	return s.CustomMaxLength
}

func (s *GetTodoDetailResponseBodyQuestions) GetCustomPlaceholder() *string {
	return s.CustomPlaceholder
}

func (s *GetTodoDetailResponseBodyQuestions) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *GetTodoDetailResponseBodyQuestions) GetMode() *string {
	return s.Mode
}

func (s *GetTodoDetailResponseBodyQuestions) GetOptions() []*GetTodoDetailResponseBodyQuestionsOptions {
	return s.Options
}

func (s *GetTodoDetailResponseBodyQuestions) GetPlaceholder() *string {
	return s.Placeholder
}

func (s *GetTodoDetailResponseBodyQuestions) GetQuestion() *string {
	return s.Question
}

func (s *GetTodoDetailResponseBodyQuestions) GetRequired() *bool {
	return s.Required
}

func (s *GetTodoDetailResponseBodyQuestions) SetAllowCustom(v bool) *GetTodoDetailResponseBodyQuestions {
	s.AllowCustom = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetCustomLabel(v string) *GetTodoDetailResponseBodyQuestions {
	s.CustomLabel = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetCustomMaxLength(v int64) *GetTodoDetailResponseBodyQuestions {
	s.CustomMaxLength = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetCustomPlaceholder(v string) *GetTodoDetailResponseBodyQuestions {
	s.CustomPlaceholder = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetMaxLength(v int64) *GetTodoDetailResponseBodyQuestions {
	s.MaxLength = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetMode(v string) *GetTodoDetailResponseBodyQuestions {
	s.Mode = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetOptions(v []*GetTodoDetailResponseBodyQuestionsOptions) *GetTodoDetailResponseBodyQuestions {
	s.Options = v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetPlaceholder(v string) *GetTodoDetailResponseBodyQuestions {
	s.Placeholder = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetQuestion(v string) *GetTodoDetailResponseBodyQuestions {
	s.Question = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) SetRequired(v bool) *GetTodoDetailResponseBodyQuestions {
	s.Required = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestions) Validate() error {
	if s.Options != nil {
		for _, item := range s.Options {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTodoDetailResponseBodyQuestionsOptions struct {
	// 选项补充说明
	//
	// example:
	//
	// 选项补充说明
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 选项文案（引用键）
	//
	// example:
	//
	// 通过
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 是否推荐
	//
	// example:
	//
	// false
	Recommended *bool `json:"recommended,omitempty" xml:"recommended,omitempty"`
}

func (s GetTodoDetailResponseBodyQuestionsOptions) String() string {
	return dara.Prettify(s)
}

func (s GetTodoDetailResponseBodyQuestionsOptions) GoString() string {
	return s.String()
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) GetDescription() *string {
	return s.Description
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) GetLabel() *string {
	return s.Label
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) GetRecommended() *bool {
	return s.Recommended
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) SetDescription(v string) *GetTodoDetailResponseBodyQuestionsOptions {
	s.Description = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) SetLabel(v string) *GetTodoDetailResponseBodyQuestionsOptions {
	s.Label = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) SetRecommended(v bool) *GetTodoDetailResponseBodyQuestionsOptions {
	s.Recommended = &v
	return s
}

func (s *GetTodoDetailResponseBodyQuestionsOptions) Validate() error {
	return dara.Validate(s)
}
