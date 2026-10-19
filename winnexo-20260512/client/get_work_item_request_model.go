// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetWorkItemRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *GetWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemRequest
	GetWorkItemId() *string
}

type GetWorkItemRequest struct {
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

func (s GetWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemRequest) SetCollaborationGroupId(v string) *GetWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemRequest) SetTenantId(v string) *GetWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemRequest) SetWorkItemId(v string) *GetWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
