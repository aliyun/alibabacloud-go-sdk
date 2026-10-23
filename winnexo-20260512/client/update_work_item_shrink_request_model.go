// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkItemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *UpdateWorkItemShrinkRequest
	GetAcceptanceCriteria() *string
	SetAttachmentsShrink(v string) *UpdateWorkItemShrinkRequest
	GetAttachmentsShrink() *string
	SetCollaborationGroupId(v string) *UpdateWorkItemShrinkRequest
	GetCollaborationGroupId() *string
	SetCollaboratorUserIdsShrink(v string) *UpdateWorkItemShrinkRequest
	GetCollaboratorUserIdsShrink() *string
	SetConversationId(v string) *UpdateWorkItemShrinkRequest
	GetConversationId() *string
	SetDescription(v string) *UpdateWorkItemShrinkRequest
	GetDescription() *string
	SetExpectedVersion(v int64) *UpdateWorkItemShrinkRequest
	GetExpectedVersion() *int64
	SetOwnerUserId(v int64) *UpdateWorkItemShrinkRequest
	GetOwnerUserId() *int64
	SetPriority(v string) *UpdateWorkItemShrinkRequest
	GetPriority() *string
	SetTenantId(v string) *UpdateWorkItemShrinkRequest
	GetTenantId() *string
	SetTitle(v string) *UpdateWorkItemShrinkRequest
	GetTitle() *string
	SetWorkItemId(v string) *UpdateWorkItemShrinkRequest
	GetWorkItemId() *string
	SetWorkflowTemplateIdsShrink(v string) *UpdateWorkItemShrinkRequest
	GetWorkflowTemplateIdsShrink() *string
}

type UpdateWorkItemShrinkRequest struct {
	// The acceptance criteria.
	//
	// example:
	//
	// example
	AcceptanceCriteria *string `json:"acceptanceCriteria,omitempty" xml:"acceptanceCriteria,omitempty"`
	// The stable references for the attachments.
	//
	// example:
	//
	// example
	AttachmentsShrink *string `json:"attachments,omitempty" xml:"attachments,omitempty"`
	// The ID of the collaboration space.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The list of user IDs of the collaborators.
	//
	// example:
	//
	// example
	CollaboratorUserIdsShrink *string `json:"collaboratorUserIds,omitempty" xml:"collaboratorUserIds,omitempty"`
	// The ID of the associated conversation.
	//
	// example:
	//
	// example
	ConversationId *string `json:"conversationId,omitempty" xml:"conversationId,omitempty"`
	// The description of the work item.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The optimistic version number. Obtain this value from the version field in the previous response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ExpectedVersion *int64 `json:"expectedVersion,omitempty" xml:"expectedVersion,omitempty"`
	// The user ID of the owner. If not specified, the user ID of the caller is used.
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// The scheduling priority.
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The title of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
	// The list of bound workflow template IDs. If not specified, the existing bindings are retained. If an empty array is provided, the bindings are cleared.
	//
	// example:
	//
	// example
	WorkflowTemplateIdsShrink *string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty"`
}

func (s UpdateWorkItemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemShrinkRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *UpdateWorkItemShrinkRequest) GetAttachmentsShrink() *string {
	return s.AttachmentsShrink
}

func (s *UpdateWorkItemShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkItemShrinkRequest) GetCollaboratorUserIdsShrink() *string {
	return s.CollaboratorUserIdsShrink
}

func (s *UpdateWorkItemShrinkRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateWorkItemShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkItemShrinkRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateWorkItemShrinkRequest) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *UpdateWorkItemShrinkRequest) GetPriority() *string {
	return s.Priority
}

func (s *UpdateWorkItemShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkItemShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateWorkItemShrinkRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *UpdateWorkItemShrinkRequest) GetWorkflowTemplateIdsShrink() *string {
	return s.WorkflowTemplateIdsShrink
}

func (s *UpdateWorkItemShrinkRequest) SetAcceptanceCriteria(v string) *UpdateWorkItemShrinkRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetAttachmentsShrink(v string) *UpdateWorkItemShrinkRequest {
	s.AttachmentsShrink = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetCollaborationGroupId(v string) *UpdateWorkItemShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetCollaboratorUserIdsShrink(v string) *UpdateWorkItemShrinkRequest {
	s.CollaboratorUserIdsShrink = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetConversationId(v string) *UpdateWorkItemShrinkRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetDescription(v string) *UpdateWorkItemShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetExpectedVersion(v int64) *UpdateWorkItemShrinkRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetOwnerUserId(v int64) *UpdateWorkItemShrinkRequest {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetPriority(v string) *UpdateWorkItemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetTenantId(v string) *UpdateWorkItemShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetTitle(v string) *UpdateWorkItemShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetWorkItemId(v string) *UpdateWorkItemShrinkRequest {
	s.WorkItemId = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) SetWorkflowTemplateIdsShrink(v string) *UpdateWorkItemShrinkRequest {
	s.WorkflowTemplateIdsShrink = &v
	return s
}

func (s *UpdateWorkItemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
