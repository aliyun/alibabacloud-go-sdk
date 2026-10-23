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
	// The ID of the collaborative workspace. The match scope is limited to this workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The workflow ID at the run level.
	//
	// This parameter is required.
	//
	// example:
	//
	// wf-1
	SourceWorkflowId *string `json:"sourceWorkflowId,omitempty" xml:"sourceWorkflowId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
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
