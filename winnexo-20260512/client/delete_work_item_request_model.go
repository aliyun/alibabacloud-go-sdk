// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *DeleteWorkItemRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *DeleteWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *DeleteWorkItemRequest
	GetWorkItemId() *string
}

type DeleteWorkItemRequest struct {
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

func (s DeleteWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *DeleteWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *DeleteWorkItemRequest) SetCollaborationGroupId(v string) *DeleteWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *DeleteWorkItemRequest) SetTenantId(v string) *DeleteWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteWorkItemRequest) SetWorkItemId(v string) *DeleteWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *DeleteWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
