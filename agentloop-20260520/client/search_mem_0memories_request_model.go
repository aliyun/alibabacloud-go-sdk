// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMem0MemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *SearchMem0MemoriesRequest
	GetAgentSpace() *string
	SetBody(v map[string]interface{}) *SearchMem0MemoriesRequest
	GetBody() map[string]interface{}
	SetContextStoreId(v string) *SearchMem0MemoriesRequest
	GetContextStoreId() *string
	SetEnableGraph(v bool) *SearchMem0MemoriesRequest
	GetEnableGraph() *bool
	SetOrgId(v string) *SearchMem0MemoriesRequest
	GetOrgId() *string
	SetProjectId(v string) *SearchMem0MemoriesRequest
	GetProjectId() *string
}

type SearchMem0MemoriesRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// {"query":"用户喝什么","filters":{"AND":[{"user_id":"alice"}]},"top_k":5}
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

func (s SearchMem0MemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMem0MemoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchMem0MemoriesRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *SearchMem0MemoriesRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *SearchMem0MemoriesRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *SearchMem0MemoriesRequest) GetEnableGraph() *bool {
	return s.EnableGraph
}

func (s *SearchMem0MemoriesRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *SearchMem0MemoriesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SearchMem0MemoriesRequest) SetAgentSpace(v string) *SearchMem0MemoriesRequest {
	s.AgentSpace = &v
	return s
}

func (s *SearchMem0MemoriesRequest) SetBody(v map[string]interface{}) *SearchMem0MemoriesRequest {
	s.Body = v
	return s
}

func (s *SearchMem0MemoriesRequest) SetContextStoreId(v string) *SearchMem0MemoriesRequest {
	s.ContextStoreId = &v
	return s
}

func (s *SearchMem0MemoriesRequest) SetEnableGraph(v bool) *SearchMem0MemoriesRequest {
	s.EnableGraph = &v
	return s
}

func (s *SearchMem0MemoriesRequest) SetOrgId(v string) *SearchMem0MemoriesRequest {
	s.OrgId = &v
	return s
}

func (s *SearchMem0MemoriesRequest) SetProjectId(v string) *SearchMem0MemoriesRequest {
	s.ProjectId = &v
	return s
}

func (s *SearchMem0MemoriesRequest) Validate() error {
	return dara.Validate(s)
}
