// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemAttachmentAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetAttachmentId() *string
	SetCollaborationGroupId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ResolveWorkItemAttachmentAccessRequest
	GetWorkItemId() *string
}

type ResolveWorkItemAttachmentAccessRequest struct {
	// 附件 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 工作项 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ResolveWorkItemAttachmentAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemAttachmentAccessRequest) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResolveWorkItemAttachmentAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetAttachmentId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.AttachmentId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetCollaborationGroupId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetTenantId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.TenantId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) SetWorkItemId(v string) *ResolveWorkItemAttachmentAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *ResolveWorkItemAttachmentAccessRequest) Validate() error {
	return dara.Validate(s)
}
