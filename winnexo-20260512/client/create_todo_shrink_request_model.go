// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateTodoShrinkRequest
	GetAppToken() *string
	SetAssigneeUserIdsShrink(v string) *CreateTodoShrinkRequest
	GetAssigneeUserIdsShrink() *string
	SetCollaborationGroupId(v string) *CreateTodoShrinkRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *CreateTodoShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateTodoShrinkRequest
	GetOpsUserId() *string
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
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The list of platform user IDs of the assignees. You can specify 1 to 100 IDs. Duplicate IDs are automatically removed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43
	AssigneeUserIdsShrink *string `json:"assigneeUserIds,omitempty" xml:"assigneeUserIds,omitempty"`
	// The ID of the collaboration workspace. This parameter is optional.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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
	QuestionsShrink *string `json:"questions,omitempty" xml:"questions,omitempty"`
	// The source context. The value is passed through and stored in the database as-is without validation.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	SourceContextShrink *string `json:"sourceContext,omitempty" xml:"sourceContext,omitempty"`
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

func (s CreateTodoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateTodoShrinkRequest) GetAssigneeUserIdsShrink() *string {
	return s.AssigneeUserIdsShrink
}

func (s *CreateTodoShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateTodoShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateTodoShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *CreateTodoShrinkRequest) SetAppToken(v string) *CreateTodoShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetAssigneeUserIdsShrink(v string) *CreateTodoShrinkRequest {
	s.AssigneeUserIdsShrink = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetCollaborationGroupId(v string) *CreateTodoShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetOpsUserCode(v string) *CreateTodoShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateTodoShrinkRequest) SetOpsUserId(v string) *CreateTodoShrinkRequest {
	s.OpsUserId = &v
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
