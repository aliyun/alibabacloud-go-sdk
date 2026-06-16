// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMem0MemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetMem0MemoryRequest
	GetAgentSpace() *string
	SetContextStoreId(v string) *GetMem0MemoryRequest
	GetContextStoreId() *string
	SetOrgId(v string) *GetMem0MemoryRequest
	GetOrgId() *string
	SetProjectId(v string) *GetMem0MemoryRequest
	GetProjectId() *string
}

type GetMem0MemoryRequest struct {
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

func (s GetMem0MemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMem0MemoryRequest) GoString() string {
	return s.String()
}

func (s *GetMem0MemoryRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetMem0MemoryRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *GetMem0MemoryRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetMem0MemoryRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetMem0MemoryRequest) SetAgentSpace(v string) *GetMem0MemoryRequest {
	s.AgentSpace = &v
	return s
}

func (s *GetMem0MemoryRequest) SetContextStoreId(v string) *GetMem0MemoryRequest {
	s.ContextStoreId = &v
	return s
}

func (s *GetMem0MemoryRequest) SetOrgId(v string) *GetMem0MemoryRequest {
	s.OrgId = &v
	return s
}

func (s *GetMem0MemoryRequest) SetProjectId(v string) *GetMem0MemoryRequest {
	s.ProjectId = &v
	return s
}

func (s *GetMem0MemoryRequest) Validate() error {
	return dara.Validate(s)
}
