// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptanceCriteria(v string) *CreateWorkItemRequest
	GetAcceptanceCriteria() *string
	SetAttachments(v []*CreateWorkItemRequestAttachments) *CreateWorkItemRequest
	GetAttachments() []*CreateWorkItemRequestAttachments
	SetCollaborationGroupId(v string) *CreateWorkItemRequest
	GetCollaborationGroupId() *string
	SetCollaboratorUserIds(v []*int64) *CreateWorkItemRequest
	GetCollaboratorUserIds() []*int64
	SetConversationId(v string) *CreateWorkItemRequest
	GetConversationId() *string
	SetDescription(v string) *CreateWorkItemRequest
	GetDescription() *string
	SetItemType(v string) *CreateWorkItemRequest
	GetItemType() *string
	SetOwnerUserId(v int64) *CreateWorkItemRequest
	GetOwnerUserId() *int64
	SetPriority(v string) *CreateWorkItemRequest
	GetPriority() *string
	SetTenantId(v string) *CreateWorkItemRequest
	GetTenantId() *string
	SetTitle(v string) *CreateWorkItemRequest
	GetTitle() *string
	SetWorkflowTemplateIds(v []*string) *CreateWorkItemRequest
	GetWorkflowTemplateIds() []*string
}

type CreateWorkItemRequest struct {
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
	Attachments []*CreateWorkItemRequestAttachments `json:"attachments,omitempty" xml:"attachments,omitempty" type:"Repeated"`
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
	CollaboratorUserIds []*int64 `json:"collaboratorUserIds,omitempty" xml:"collaboratorUserIds,omitempty" type:"Repeated"`
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
	WorkflowTemplateIds []*string `json:"workflowTemplateIds,omitempty" xml:"workflowTemplateIds,omitempty" type:"Repeated"`
}

func (s CreateWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkItemRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkItemRequest) GetAcceptanceCriteria() *string {
	return s.AcceptanceCriteria
}

func (s *CreateWorkItemRequest) GetAttachments() []*CreateWorkItemRequestAttachments {
	return s.Attachments
}

func (s *CreateWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkItemRequest) GetCollaboratorUserIds() []*int64 {
	return s.CollaboratorUserIds
}

func (s *CreateWorkItemRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateWorkItemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkItemRequest) GetItemType() *string {
	return s.ItemType
}

func (s *CreateWorkItemRequest) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *CreateWorkItemRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkItemRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkItemRequest) GetWorkflowTemplateIds() []*string {
	return s.WorkflowTemplateIds
}

func (s *CreateWorkItemRequest) SetAcceptanceCriteria(v string) *CreateWorkItemRequest {
	s.AcceptanceCriteria = &v
	return s
}

func (s *CreateWorkItemRequest) SetAttachments(v []*CreateWorkItemRequestAttachments) *CreateWorkItemRequest {
	s.Attachments = v
	return s
}

func (s *CreateWorkItemRequest) SetCollaborationGroupId(v string) *CreateWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkItemRequest) SetCollaboratorUserIds(v []*int64) *CreateWorkItemRequest {
	s.CollaboratorUserIds = v
	return s
}

func (s *CreateWorkItemRequest) SetConversationId(v string) *CreateWorkItemRequest {
	s.ConversationId = &v
	return s
}

func (s *CreateWorkItemRequest) SetDescription(v string) *CreateWorkItemRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkItemRequest) SetItemType(v string) *CreateWorkItemRequest {
	s.ItemType = &v
	return s
}

func (s *CreateWorkItemRequest) SetOwnerUserId(v int64) *CreateWorkItemRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateWorkItemRequest) SetPriority(v string) *CreateWorkItemRequest {
	s.Priority = &v
	return s
}

func (s *CreateWorkItemRequest) SetTenantId(v string) *CreateWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *CreateWorkItemRequest) SetTitle(v string) *CreateWorkItemRequest {
	s.Title = &v
	return s
}

func (s *CreateWorkItemRequest) SetWorkflowTemplateIds(v []*string) *CreateWorkItemRequest {
	s.WorkflowTemplateIds = v
	return s
}

func (s *CreateWorkItemRequest) Validate() error {
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

type CreateWorkItemRequestAttachments struct {
	// 内容类型
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 附件引用地址
	//
	// example:
	//
	// example
	RefUrl *string `json:"refUrl,omitempty" xml:"refUrl,omitempty"`
	// 附件资源ID
	//
	// This parameter is required.
	//
	// example:
	//
	// res_example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 附件资源类型
	//
	// This parameter is required.
	//
	// example:
	//
	// attachment
	ResourceKind *string `json:"resourceKind,omitempty" xml:"resourceKind,omitempty"`
	// 附件大小(字节)，上限 20MB
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"sizeBytes,omitempty" xml:"sizeBytes,omitempty"`
	// 附件标题
	//
	// example:
	//
	// example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateWorkItemRequestAttachments) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkItemRequestAttachments) GoString() string {
	return s.String()
}

func (s *CreateWorkItemRequestAttachments) GetContentType() *string {
	return s.ContentType
}

func (s *CreateWorkItemRequestAttachments) GetRefUrl() *string {
	return s.RefUrl
}

func (s *CreateWorkItemRequestAttachments) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateWorkItemRequestAttachments) GetResourceKind() *string {
	return s.ResourceKind
}

func (s *CreateWorkItemRequestAttachments) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateWorkItemRequestAttachments) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkItemRequestAttachments) SetContentType(v string) *CreateWorkItemRequestAttachments {
	s.ContentType = &v
	return s
}

func (s *CreateWorkItemRequestAttachments) SetRefUrl(v string) *CreateWorkItemRequestAttachments {
	s.RefUrl = &v
	return s
}

func (s *CreateWorkItemRequestAttachments) SetResourceId(v string) *CreateWorkItemRequestAttachments {
	s.ResourceId = &v
	return s
}

func (s *CreateWorkItemRequestAttachments) SetResourceKind(v string) *CreateWorkItemRequestAttachments {
	s.ResourceKind = &v
	return s
}

func (s *CreateWorkItemRequestAttachments) SetSizeBytes(v int64) *CreateWorkItemRequestAttachments {
	s.SizeBytes = &v
	return s
}

func (s *CreateWorkItemRequestAttachments) SetTitle(v string) *CreateWorkItemRequestAttachments {
	s.Title = &v
	return s
}

func (s *CreateWorkItemRequestAttachments) Validate() error {
	return dara.Validate(s)
}
