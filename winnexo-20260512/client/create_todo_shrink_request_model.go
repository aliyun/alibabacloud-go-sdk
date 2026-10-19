// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssigneeUserIdsShrink(v string) *CreateTodoShrinkRequest
	GetAssigneeUserIdsShrink() *string
	SetCollaborationGroupId(v string) *CreateTodoShrinkRequest
	GetCollaborationGroupId() *string
	SetPriority(v string) *CreateTodoShrinkRequest
	GetPriority() *string
	SetQuestionsShrink(v string) *CreateTodoShrinkRequest
	GetQuestionsShrink() *string
	SetSourceContextShrink(v string) *CreateTodoShrinkRequest
	GetSourceContextShrink() *string
	SetSourceType(v string) *CreateTodoShrinkRequest
	GetSourceType() *string
	SetTenantId(v string) *CreateTodoShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateTodoShrinkRequest
	GetTitle() *string
}

type CreateTodoShrinkRequest struct {
	// 处理人平台用户ID列表（1~100，去重后生效）
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	AssigneeUserIdsShrink *string `json:"assigneeUserIds,omitempty" xml:"assigneeUserIds,omitempty"`
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
	QuestionsShrink *string `json:"questions,omitempty" xml:"questions,omitempty"`
	// 来源上下文（原样透传落库，不校验）
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceContextShrink *string `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
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

func (s CreateTodoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoShrinkRequest) GetAssigneeUserIdsShrink() *string {
	return s.AssigneeUserIdsShrink
}

func (s *CreateTodoShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateTodoShrinkRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateTodoShrinkRequest) GetQuestionsShrink() *string {
	return s.QuestionsShrink
}

func (s *CreateTodoShrinkRequest) GetSourceContextShrink() *string {
	return s.SourceContextShrink
}

func (s *CreateTodoShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateTodoShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTodoShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTodoShrinkRequest) SetAssigneeUserIdsShrink(v string) *CreateTodoShrinkRequest {
	s.AssigneeUserIdsShrink = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetCollaborationGroupId(v string) *CreateTodoShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetPriority(v string) *CreateTodoShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetQuestionsShrink(v string) *CreateTodoShrinkRequest {
	s.QuestionsShrink = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetSourceContextShrink(v string) *CreateTodoShrinkRequest {
	s.SourceContextShrink = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetSourceType(v string) *CreateTodoShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetTenantId(v string) *CreateTodoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetTitle(v string) *CreateTodoShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateTodoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
