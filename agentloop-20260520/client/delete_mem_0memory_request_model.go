// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0MemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *DeleteMem0MemoryRequest
	GetAgentSpace() *string
	SetContextStoreId(v string) *DeleteMem0MemoryRequest
	GetContextStoreId() *string
	SetOrgId(v string) *DeleteMem0MemoryRequest
	GetOrgId() *string
	SetProjectId(v string) *DeleteMem0MemoryRequest
	GetProjectId() *string
}

type DeleteMem0MemoryRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// cs-0001
	ContextStoreId *string `json:"context_store_id,omitempty" xml:"context_store_id,omitempty"`
	// example:
	//
	// org-001
	OrgId *string `json:"org_id,omitempty" xml:"org_id,omitempty"`
	// example:
	//
	// proj-001
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

func (s DeleteMem0MemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0MemoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteMem0MemoryRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *DeleteMem0MemoryRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *DeleteMem0MemoryRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *DeleteMem0MemoryRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteMem0MemoryRequest) SetAgentSpace(v string) *DeleteMem0MemoryRequest {
	s.AgentSpace = &v
	return s
}

func (s *DeleteMem0MemoryRequest) SetContextStoreId(v string) *DeleteMem0MemoryRequest {
	s.ContextStoreId = &v
	return s
}

func (s *DeleteMem0MemoryRequest) SetOrgId(v string) *DeleteMem0MemoryRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteMem0MemoryRequest) SetProjectId(v string) *DeleteMem0MemoryRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteMem0MemoryRequest) Validate() error {
	return dara.Validate(s)
}
