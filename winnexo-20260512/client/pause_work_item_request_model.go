// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *PauseWorkItemRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *PauseWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *PauseWorkItemRequest
	GetWorkItemId() *string
}

type PauseWorkItemRequest struct {
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

func (s PauseWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseWorkItemRequest) GoString() string {
	return s.String()
}

func (s *PauseWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *PauseWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PauseWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *PauseWorkItemRequest) SetCollaborationGroupId(v string) *PauseWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *PauseWorkItemRequest) SetTenantId(v string) *PauseWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *PauseWorkItemRequest) SetWorkItemId(v string) *PauseWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *PauseWorkItemRequest) Validate() error {
	return dara.Validate(s)
}
