// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssigneeUserIds(v []*string) *CreateTodoRequest
	GetAssigneeUserIds() []*string
	SetCollaborationGroupId(v string) *CreateTodoRequest
	GetCollaborationGroupId() *string
	SetPriority(v string) *CreateTodoRequest
	GetPriority() *string
	SetQuestions(v []*CreateTodoRequestQuestions) *CreateTodoRequest
	GetQuestions() []*CreateTodoRequestQuestions
	SetSourceContext(v map[string]interface{}) *CreateTodoRequest
	GetSourceContext() map[string]interface{}
	SetSourceType(v string) *CreateTodoRequest
	GetSourceType() *string
	SetTenantId(v string) *CreateTodoRequest
	GetTenantId() *string
	SetTitle(v string) *CreateTodoRequest
	GetTitle() *string
}

type CreateTodoRequest struct {
	// 处理人平台用户ID列表（1~100，去重后生效）
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	AssigneeUserIds []*string `json:"assigneeUserIds,omitempty" xml:"assigneeUserIds,omitempty" type:"Repeated"`
	// 协作空间 ID（可选）
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 优先级
	//
	// example:
	//
	// MEDIUM
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 问题清单（1~N）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Questions []*CreateTodoRequestQuestions `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	// 来源上下文（原样透传落库，不校验）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// 待办来源（仅 OPEN_API）
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_API
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 待办标题（≤200 字）
	//
	// This parameter is required.
	//
	// example:
	//
	// 治理审批
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTodoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoRequest) GetAssigneeUserIds() []*string {
	return s.AssigneeUserIds
}

func (s *CreateTodoRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateTodoRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateTodoRequest) GetQuestions() []*CreateTodoRequestQuestions {
	return s.Questions
}

func (s *CreateTodoRequest) GetSourceContext() map[string]interface{} {
	return s.SourceContext
}

func (s *CreateTodoRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateTodoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTodoRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTodoRequest) SetAssigneeUserIds(v []*string) *CreateTodoRequest {
	s.AssigneeUserIds = v
	return s
}

func (s *CreateTodoRequest) SetCollaborationGroupId(v string) *CreateTodoRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateTodoRequest) SetPriority(v string) *CreateTodoRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoRequest) SetQuestions(v []*CreateTodoRequestQuestions) *CreateTodoRequest {
	s.Questions = v
	return s
}

func (s *CreateTodoRequest) SetSourceContext(v map[string]interface{}) *CreateTodoRequest {
	s.SourceContext = v
	return s
}

func (s *CreateTodoRequest) SetSourceType(v string) *CreateTodoRequest {
	s.SourceType = &v
	return s
}

func (s *CreateTodoRequest) SetTenantId(v string) *CreateTodoRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTodoRequest) SetTitle(v string) *CreateTodoRequest {
	s.Title = &v
	return s
}

func (s *CreateTodoRequest) Validate() error {
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

type CreateTodoRequestQuestions struct {
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
	// This parameter is required.
	//
	// example:
	//
	// SINGLE_CHOICE
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// 问题的选项列表；选择题必填，每项含选项文案（label）、补充说明（description）与是否推荐（recommended）
	Options []*CreateTodoRequestQuestionsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// TEXT 题占位符
	//
	// example:
	//
	// 请输入
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// 问题正文（引用键，清单内唯一）
	//
	// This parameter is required.
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

func (s CreateTodoRequestQuestions) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoRequestQuestions) GoString() string {
	return s.String()
}

func (s *CreateTodoRequestQuestions) GetAllowCustom() *bool {
	return s.AllowCustom
}

func (s *CreateTodoRequestQuestions) GetCustomLabel() *string {
	return s.CustomLabel
}

func (s *CreateTodoRequestQuestions) GetCustomMaxLength() *int64 {
	return s.CustomMaxLength
}

func (s *CreateTodoRequestQuestions) GetCustomPlaceholder() *string {
	return s.CustomPlaceholder
}

func (s *CreateTodoRequestQuestions) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *CreateTodoRequestQuestions) GetMode() *string {
	return s.Mode
}

func (s *CreateTodoRequestQuestions) GetOptions() []*CreateTodoRequestQuestionsOptions {
	return s.Options
}

func (s *CreateTodoRequestQuestions) GetPlaceholder() *string {
	return s.Placeholder
}

func (s *CreateTodoRequestQuestions) GetQuestion() *string {
	return s.Question
}

func (s *CreateTodoRequestQuestions) GetRequired() *bool {
	return s.Required
}

func (s *CreateTodoRequestQuestions) SetAllowCustom(v bool) *CreateTodoRequestQuestions {
	s.AllowCustom = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetCustomLabel(v string) *CreateTodoRequestQuestions {
	s.CustomLabel = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetCustomMaxLength(v int64) *CreateTodoRequestQuestions {
	s.CustomMaxLength = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetCustomPlaceholder(v string) *CreateTodoRequestQuestions {
	s.CustomPlaceholder = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetMaxLength(v int64) *CreateTodoRequestQuestions {
	s.MaxLength = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetMode(v string) *CreateTodoRequestQuestions {
	s.Mode = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetOptions(v []*CreateTodoRequestQuestionsOptions) *CreateTodoRequestQuestions {
	s.Options = v
	return s
}

func (s *CreateTodoRequestQuestions) SetPlaceholder(v string) *CreateTodoRequestQuestions {
	s.Placeholder = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetQuestion(v string) *CreateTodoRequestQuestions {
	s.Question = &v
	return s
}

func (s *CreateTodoRequestQuestions) SetRequired(v bool) *CreateTodoRequestQuestions {
	s.Required = &v
	return s
}

func (s *CreateTodoRequestQuestions) Validate() error {
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

type CreateTodoRequestQuestionsOptions struct {
	// 选项补充说明
	//
	// example:
	//
	// 选项补充说明
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 选项文案（引用键）
	//
	// This parameter is required.
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

func (s CreateTodoRequestQuestionsOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoRequestQuestionsOptions) GoString() string {
	return s.String()
}

func (s *CreateTodoRequestQuestionsOptions) GetDescription() *string {
	return s.Description
}

func (s *CreateTodoRequestQuestionsOptions) GetLabel() *string {
	return s.Label
}

func (s *CreateTodoRequestQuestionsOptions) GetRecommended() *bool {
	return s.Recommended
}

func (s *CreateTodoRequestQuestionsOptions) SetDescription(v string) *CreateTodoRequestQuestionsOptions {
	s.Description = &v
	return s
}

func (s *CreateTodoRequestQuestionsOptions) SetLabel(v string) *CreateTodoRequestQuestionsOptions {
	s.Label = &v
	return s
}

func (s *CreateTodoRequestQuestionsOptions) SetRecommended(v bool) *CreateTodoRequestQuestionsOptions {
	s.Recommended = &v
	return s
}

func (s *CreateTodoRequestQuestionsOptions) Validate() error {
	return dara.Validate(s)
}
