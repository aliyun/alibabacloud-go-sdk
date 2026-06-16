// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0MemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *DeleteMem0MemoriesRequest
	GetAgentSpace() *string
	SetAgentId(v string) *DeleteMem0MemoriesRequest
	GetAgentId() *string
	SetAppId(v string) *DeleteMem0MemoriesRequest
	GetAppId() *string
	SetContextStoreId(v string) *DeleteMem0MemoriesRequest
	GetContextStoreId() *string
	SetMetadata(v map[string]interface{}) *DeleteMem0MemoriesRequest
	GetMetadata() map[string]interface{}
	SetOrgId(v string) *DeleteMem0MemoriesRequest
	GetOrgId() *string
	SetProjectId(v string) *DeleteMem0MemoriesRequest
	GetProjectId() *string
	SetRunId(v string) *DeleteMem0MemoriesRequest
	GetRunId() *string
	SetUserId(v string) *DeleteMem0MemoriesRequest
	GetUserId() *string
}

type DeleteMem0MemoriesRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// agent-001
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// example:
	//
	// app-001
	AppId *string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// example:
	//
	// cs-0001
	ContextStoreId *string `json:"context_store_id,omitempty" xml:"context_store_id,omitempty"`
	// example:
	//
	// {"channel":"app","locale":"zh-CN"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// org-001
	OrgId *string `json:"org_id,omitempty" xml:"org_id,omitempty"`
	// example:
	//
	// proj-001
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// example:
	//
	// run-001
	RunId *string `json:"run_id,omitempty" xml:"run_id,omitempty"`
	// example:
	//
	// alice
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s DeleteMem0MemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0MemoriesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMem0MemoriesRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *DeleteMem0MemoriesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteMem0MemoriesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMem0MemoriesRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *DeleteMem0MemoriesRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *DeleteMem0MemoriesRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *DeleteMem0MemoriesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteMem0MemoriesRequest) GetRunId() *string {
	return s.RunId
}

func (s *DeleteMem0MemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteMem0MemoriesRequest) SetAgentSpace(v string) *DeleteMem0MemoriesRequest {
	s.AgentSpace = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetAgentId(v string) *DeleteMem0MemoriesRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetAppId(v string) *DeleteMem0MemoriesRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetContextStoreId(v string) *DeleteMem0MemoriesRequest {
	s.ContextStoreId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetMetadata(v map[string]interface{}) *DeleteMem0MemoriesRequest {
	s.Metadata = v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetOrgId(v string) *DeleteMem0MemoriesRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetProjectId(v string) *DeleteMem0MemoriesRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetRunId(v string) *DeleteMem0MemoriesRequest {
	s.RunId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) SetUserId(v string) *DeleteMem0MemoriesRequest {
	s.UserId = &v
	return s
}

func (s *DeleteMem0MemoriesRequest) Validate() error {
	return dara.Validate(s)
}
