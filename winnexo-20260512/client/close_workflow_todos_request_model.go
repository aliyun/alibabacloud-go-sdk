// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseWorkflowTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CloseWorkflowTodosRequest
	GetCollaborationGroupId() *string
	SetSourceWorkflowId(v string) *CloseWorkflowTodosRequest
	GetSourceWorkflowId() *string
	SetTenantId(v string) *CloseWorkflowTodosRequest
	GetTenantId() *string
}

type CloseWorkflowTodosRequest struct {
	// 协作空间 ID（命中范围限定该空间内）
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 工作流 ID（run 粒度）
	//
	// This parameter is required.
	//
	// example:
	//
	// wf-1
	SourceWorkflowId *string `json:"sourceWorkflowId,omitempty" xml:"sourceWorkflowId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CloseWorkflowTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseWorkflowTodosRequest) GoString() string {
	return s.String()
}

func (s *CloseWorkflowTodosRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CloseWorkflowTodosRequest) GetSourceWorkflowId() *string {
	return s.SourceWorkflowId
}

func (s *CloseWorkflowTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CloseWorkflowTodosRequest) SetCollaborationGroupId(v string) *CloseWorkflowTodosRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetSourceWorkflowId(v string) *CloseWorkflowTodosRequest {
	s.SourceWorkflowId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) SetTenantId(v string) *CloseWorkflowTodosRequest {
	s.TenantId = &v
	return s
}

func (s *CloseWorkflowTodosRequest) Validate() error {
	return dara.Validate(s)
}
