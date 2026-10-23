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
	// The list of platform user IDs of the assignees. You can specify 1 to 100 IDs. Duplicate IDs are automatically removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	AssigneeUserIds []*string `json:"assigneeUserIds,omitempty" xml:"assigneeUserIds,omitempty" type:"Repeated"`
	// The ID of the collaboration workspace. This parameter is optional.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The priority.
	//
	// example:
	//
	// MEDIUM
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The list of questions.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Questions []*CreateTodoRequestQuestions `json:"questions,omitempty" xml:"questions,omitempty" type:"Repeated"`
	// The source context. The value is passed through and stored in the database as-is without validation.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceContext map[string]interface{} `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
	// The source of the to-do item. Only OPEN_API is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// OPEN_API
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The title of the to-do item. The maximum length is 200 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Governance approval
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
	// Specifies whether custom input is allowed.
	//
	// example:
	//
	// false
	AllowCustom *bool `json:"allowCustom,omitempty" xml:"allowCustom,omitempty"`
	// The text for the custom input item.
	//
	// example:
	//
	// Other
	CustomLabel *string `json:"customLabel,omitempty" xml:"customLabel,omitempty"`
	// The maximum length of the custom input.
	//
	// example:
	//
	// 200
	CustomMaxLength *int64 `json:"customMaxLength,omitempty" xml:"customMaxLength,omitempty"`
	// The placeholder for the custom input.
	//
	// example:
	//
	// Please specify
	CustomPlaceholder *string `json:"customPlaceholder,omitempty" xml:"customPlaceholder,omitempty"`
	// The maximum length for text questions.
	//
	// example:
	//
	// 500
	MaxLength *int64 `json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	// The response mode. Valid values: SINGLE_CHOICE, MULTI_CHOICE, and TEXT.
	//
	// This parameter is required.
	//
	// example:
	//
	// SINGLE_CHOICE
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The list of options for the question. This parameter is required for choice questions. Each option includes the option text (label), supplementary description (description), and whether it is recommended (recommended).
	Options []*CreateTodoRequestQuestionsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// The placeholder for text questions.
	//
	// example:
	//
	// Please enter
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// The question text. This is a reference key and must be unique within the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// Whether to approve the request
	Question *string `json:"question,omitempty" xml:"question,omitempty"`
	// Specifies whether the question is required.
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
	// The supplementary description of the option.
	//
	// example:
	//
	// Supplementary description of the option
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The option text. This is a reference key.
	//
	// This parameter is required.
	//
	// example:
	//
	// Approve
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// Specifies whether the option is recommended.
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
