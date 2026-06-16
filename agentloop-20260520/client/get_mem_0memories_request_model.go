// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMem0MemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetMem0MemoriesRequest
	GetAgentSpace() *string
	SetBody(v map[string]interface{}) *GetMem0MemoriesRequest
	GetBody() map[string]interface{}
	SetContextStoreId(v string) *GetMem0MemoriesRequest
	GetContextStoreId() *string
	SetEnableGraph(v bool) *GetMem0MemoriesRequest
	GetEnableGraph() *bool
	SetOrgId(v string) *GetMem0MemoriesRequest
	GetOrgId() *string
	SetProjectId(v string) *GetMem0MemoriesRequest
	GetProjectId() *string
}

type GetMem0MemoriesRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// {"filters":{"AND":[{"user_id":"alice"}]}}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// cs-0001
	ContextStoreId *string `json:"context_store_id,omitempty" xml:"context_store_id,omitempty"`
	// example:
	//
	// false
	EnableGraph *bool `json:"enable_graph,omitempty" xml:"enable_graph,omitempty"`
	// example:
	//
	// org-001
	OrgId *string `json:"org_id,omitempty" xml:"org_id,omitempty"`
	// example:
	//
	// proj-001
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
}

func (s GetMem0MemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMem0MemoriesRequest) GoString() string {
	return s.String()
}

func (s *GetMem0MemoriesRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetMem0MemoriesRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *GetMem0MemoriesRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *GetMem0MemoriesRequest) GetEnableGraph() *bool {
	return s.EnableGraph
}

func (s *GetMem0MemoriesRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetMem0MemoriesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetMem0MemoriesRequest) SetAgentSpace(v string) *GetMem0MemoriesRequest {
	s.AgentSpace = &v
	return s
}

func (s *GetMem0MemoriesRequest) SetBody(v map[string]interface{}) *GetMem0MemoriesRequest {
	s.Body = v
	return s
}

func (s *GetMem0MemoriesRequest) SetContextStoreId(v string) *GetMem0MemoriesRequest {
	s.ContextStoreId = &v
	return s
}

func (s *GetMem0MemoriesRequest) SetEnableGraph(v bool) *GetMem0MemoriesRequest {
	s.EnableGraph = &v
	return s
}

func (s *GetMem0MemoriesRequest) SetOrgId(v string) *GetMem0MemoriesRequest {
	s.OrgId = &v
	return s
}

func (s *GetMem0MemoriesRequest) SetProjectId(v string) *GetMem0MemoriesRequest {
	s.ProjectId = &v
	return s
}

func (s *GetMem0MemoriesRequest) Validate() error {
	return dara.Validate(s)
}
