// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *CreateWorkItemShrinkRequest
	GetAcceptanceCriteria() *string
	SetAttachmentsShrink(v string) *CreateWorkItemShrinkRequest
	GetAttachmentsShrink() *string
	SetCollaborationGroupId(v string) *CreateWorkItemShrinkRequest
	GetCollaborationGroupId() *string
	SetCollaboratorUserIdsShrink(v string) *CreateWorkItemShrinkRequest
	GetCollaboratorUserIdsShrink() *string
	SetConversationId(v string) *CreateWorkItemShrinkRequest
	GetConversationId() *string
	SetDescription(v string) *CreateWorkItemShrinkRequest
	GetDescription() *string
	SetItemType(v string) *CreateWorkItemShrinkRequest
	GetItemType() *string
	SetOwnerUserId(v int64) *CreateWorkItemShrinkRequest
	GetOwnerUserId() *int64
	SetPriority(v string) *CreateWorkItemShrinkRequest
	GetPriority() *string
	SetTenantId(v string) *CreateWorkItemShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *CreateWorkItemShrinkRequest
	GetTitle() *string
	SetWorkflowTemplateIdsShrink(v string) *CreateWorkItemShrinkRequest
	GetWorkflowTemplateIdsShrink() *string
}

type CreateWorkItemShrinkRequest struct {
	// 验收标准
	//
	// example:
	//
	// example
	AcceptanceCriteria *string `json:"acceptanceCriteria,omitempty" xml:"acceptanceCriteria,omitempty"`
	// 输入附件稳定引用
	//
	// example:
	//
	// example
	AttachmentsShrink *string `json:"attachments,omitempty" xml:"attachments,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 协作人用户ID列表
	//
	// example:
	//
	// example
	CollaboratorUserIdsShrink *string `json:"collaboratorUserIds,omitempty" xml:"collaboratorUserIds,omitempty"`
	// 关联会话ID
	//
	// example:
	//
	// example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// 工作项描述
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 工作项类型；仅接受 WORKFLOW
	//
	// This parameter is required.
	//
	// example:
	//
	// WORKFLOW
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// 负责人用户ID；缺省为调用者本人
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// 排产优先级
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 工作项标题
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 绑定的工作流模板ID列表
	//
	// example:
	//
	// example
	WorkflowTemplateIdsShrink *string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty"`
}

func (s CreateWorkItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkItemShrinkRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *CreateWorkItemShrinkRequest) GetAttachmentsShrink() *string {
	return s.AttachmentsShrink
}

func (s *CreateWorkItemShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkItemShrinkRequest) GetCollaboratorUserIdsShrink() *string {
	return s.CollaboratorUserIdsShrink
}

func (s *CreateWorkItemShrinkRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateWorkItemShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkItemShrinkRequest) GetItemType() *string {
	return s.ItemType
}

func (s *CreateWorkItemShrinkRequest) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *CreateWorkItemShrinkRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateWorkItemShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkItemShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkItemShrinkRequest) GetWorkflowTemplateIdsShrink() *string {
	return s.WorkflowTemplateIdsShrink
}

func (s *CreateWorkItemShrinkRequest) SetAcceptanceCriteria(v string) *CreateWorkItemShrinkRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetAttachmentsShrink(v string) *CreateWorkItemShrinkRequest {
	s.AttachmentsShrink = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetCollaborationGroupId(v string) *CreateWorkItemShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetCollaboratorUserIdsShrink(v string) *CreateWorkItemShrinkRequest {
	s.CollaboratorUserIdsShrink = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetConversationId(v string) *CreateWorkItemShrinkRequest {
	s.ConversationId = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetDescription(v string) *CreateWorkItemShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetItemType(v string) *CreateWorkItemShrinkRequest {
	s.ItemType = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetOwnerUserId(v int64) *CreateWorkItemShrinkRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetPriority(v string) *CreateWorkItemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetTenantId(v string) *CreateWorkItemShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetTitle(v string) *CreateWorkItemShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) SetWorkflowTemplateIdsShrink(v string) *CreateWorkItemShrinkRequest {
	s.WorkflowTemplateIdsShrink = &v
	return s
}

func (s *CreateWorkItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
