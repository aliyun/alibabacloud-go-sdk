// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetWorkItemAuditRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *GetWorkItemAuditRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemAuditRequest
	GetWorkItemId() *string
}

type GetWorkItemAuditRequest struct {
	// 协作空间ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 租户ID；缺省使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 工作项ID
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_0123456789ab4cde
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetWorkItemAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAuditRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemAuditRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemAuditRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemAuditRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemAuditRequest) SetCollaborationGroupId(v string) *GetWorkItemAuditRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetTenantId(v string) *GetWorkItemAuditRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemAuditRequest) SetWorkItemId(v string) *GetWorkItemAuditRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemAuditRequest) Validate() error {
	return dara.Validate(s)
}
