// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAttachmentAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentId(v string) *GetWorkItemAttachmentAccessRequest
	GetAttachmentId() *string
	SetCollaborationGroupId(v string) *GetWorkItemAttachmentAccessRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *GetWorkItemAttachmentAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemAttachmentAccessRequest
	GetWorkItemId() *string
}

type GetWorkItemAttachmentAccessRequest struct {
	// The attachment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	AttachmentId *string `json:"attachmentId,omitempty" xml:"attachmentId,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The work item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetWorkItemAttachmentAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAttachmentAccessRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemAttachmentAccessRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *GetWorkItemAttachmentAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemAttachmentAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemAttachmentAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemAttachmentAccessRequest) SetAttachmentId(v string) *GetWorkItemAttachmentAccessRequest {
	s.AttachmentId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetCollaborationGroupId(v string) *GetWorkItemAttachmentAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetTenantId(v string) *GetWorkItemAttachmentAccessRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) SetWorkItemId(v string) *GetWorkItemAttachmentAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessRequest) Validate() error {
	return dara.Validate(s)
}
