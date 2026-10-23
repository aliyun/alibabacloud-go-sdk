// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *UpdateWorkItemRequest
	GetAcceptanceCriteria() *string
	SetAttachments(v []*UpdateWorkItemRequestAttachments) *UpdateWorkItemRequest
	GetAttachments() []*UpdateWorkItemRequestAttachments
	SetCollaborationGroupId(v string) *UpdateWorkItemRequest
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *UpdateWorkItemRequest
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *UpdateWorkItemRequest
	GetConversationId() *string
	SetDescription(v string) *UpdateWorkItemRequest
	GetDescription() *string
	SetExpectedVersion(v int64) *UpdateWorkItemRequest
	GetExpectedVersion() *int64
	SetOwnerUserId(v int64) *UpdateWorkItemRequest
	GetOwnerUserId() *int64
	SetPriority(v string) *UpdateWorkItemRequest
	GetPriority() *string
	SetTenantId(v string) *UpdateWorkItemRequest
	GetTenantId() *string
	SetTitle(v string) *UpdateWorkItemRequest
	GetTitle() *string
	SetWorkItemId(v string) *UpdateWorkItemRequest
	GetWorkItemId() *string
	SetWorkflowTemplateIds(v []*string) *UpdateWorkItemRequest
	GetWorkflowTemplateIds() []*string
}

type UpdateWorkItemRequest struct {
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
	Attachments []*UpdateWorkItemRequestAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
	CollaboratorUserIds []*int64 `json:"collaboratorUserIds,omitempty" xml:"collaboratorUserIds,omitempty" type:"Repeated"`
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
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s UpdateWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *UpdateWorkItemRequest) GetAttachments() []*UpdateWorkItemRequestAttachments {
	return s.Attachments
}

func (s *UpdateWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkItemRequest) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *UpdateWorkItemRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *UpdateWorkItemRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkItemRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateWorkItemRequest) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *UpdateWorkItemRequest) GetPriority() *string {
	return s.Priority
}

func (s *UpdateWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkItemRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *UpdateWorkItemRequest) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *UpdateWorkItemRequest) SetAcceptanceCriteria(v string) *UpdateWorkItemRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *UpdateWorkItemRequest) SetAttachments(v []*UpdateWorkItemRequestAttachments) *UpdateWorkItemRequest {
	s.Attachments = v
	return s
}

func (s *UpdateWorkItemRequest) SetCollaborationGroupId(v string) *UpdateWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkItemRequest) SetCollaboratorUserIds(v []*int64) *UpdateWorkItemRequest {
	s.CollaboratorUserIds = v
	return s
}

func (s *UpdateWorkItemRequest) SetConversationId(v string) *UpdateWorkItemRequest {
	s.ConversationId = &v
	return s
}

func (s *UpdateWorkItemRequest) SetDescription(v string) *UpdateWorkItemRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkItemRequest) SetExpectedVersion(v int64) *UpdateWorkItemRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateWorkItemRequest) SetOwnerUserId(v int64) *UpdateWorkItemRequest {
	s.OwnerUserId = &v
	return s
}

func (s *UpdateWorkItemRequest) SetPriority(v string) *UpdateWorkItemRequest {
	s.Priority = &v
	return s
}

func (s *UpdateWorkItemRequest) SetTenantId(v string) *UpdateWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkItemRequest) SetTitle(v string) *UpdateWorkItemRequest {
	s.Title = &v
	return s
}

func (s *UpdateWorkItemRequest) SetWorkItemId(v string) *UpdateWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *UpdateWorkItemRequest) SetWorkflowTemplateIds(v []*string) *UpdateWorkItemRequest {
	s.WorkflowTemplateIds = v
	return s
}

func (s *UpdateWorkItemRequest) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWorkItemRequestAttachments struct {
	// The content type.
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The reference URL of the attachment.
	//
	// example:
	//
	// example
	RefUrl *string `json:"refUrl,omitempty" xml:"refUrl,omitempty"`
	// The resource ID of the attachment.
	//
	// This parameter is required.
	//
	// example:
	//
	// res_example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type of the attachment.
	//
	// This parameter is required.
	//
	// example:
	//
	// attachment
	ResourceKind *string `json:"resourceKind,omitempty" xml:"resourceKind,omitempty"`
	// The size of the attachment in bytes. The maximum size is 20 MB.
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// The title of the attachment.
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateWorkItemRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkItemRequestAttachments) GoString() string {
	return s.String()
}

func (s *UpdateWorkItemRequestAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *UpdateWorkItemRequestAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *UpdateWorkItemRequestAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateWorkItemRequestAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *UpdateWorkItemRequestAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *UpdateWorkItemRequestAttachments) GetTitle() *string {
	return s.Title
}

func (s *UpdateWorkItemRequestAttachments) SetContentType(v string) *UpdateWorkItemRequestAttachments {
	s.ContentType = &v
	return s
}

func (s *UpdateWorkItemRequestAttachments) SetRefUrl(v string) *UpdateWorkItemRequestAttachments {
	s.RefUrl = &v
	return s
}

func (s *UpdateWorkItemRequestAttachments) SetResourceId(v string) *UpdateWorkItemRequestAttachments {
	s.ResourceId = &v
	return s
}

func (s *UpdateWorkItemRequestAttachments) SetResourceKind(v string) *UpdateWorkItemRequestAttachments {
	s.ResourceKind = &v
	return s
}

func (s *UpdateWorkItemRequestAttachments) SetSizeBytes(v int64) *UpdateWorkItemRequestAttachments {
	s.SizeBytes = &v
	return s
}

func (s *UpdateWorkItemRequestAttachments) SetTitle(v string) *UpdateWorkItemRequestAttachments {
	s.Title = &v
	return s
}

func (s *UpdateWorkItemRequestAttachments) Validate() error {
	return dara.Validate(s)
}
