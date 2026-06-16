// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMem0MemoriesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *DeleteMem0MemoriesShrinkRequest
	GetAgentSpace() *string
	SetAgentId(v string) *DeleteMem0MemoriesShrinkRequest
	GetAgentId() *string
	SetAppId(v string) *DeleteMem0MemoriesShrinkRequest
	GetAppId() *string
	SetContextStoreId(v string) *DeleteMem0MemoriesShrinkRequest
	GetContextStoreId() *string
	SetMetadataShrink(v string) *DeleteMem0MemoriesShrinkRequest
	GetMetadataShrink() *string
	SetOrgId(v string) *DeleteMem0MemoriesShrinkRequest
	GetOrgId() *string
	SetProjectId(v string) *DeleteMem0MemoriesShrinkRequest
	GetProjectId() *string
	SetRunId(v string) *DeleteMem0MemoriesShrinkRequest
	GetRunId() *string
	SetUserId(v string) *DeleteMem0MemoriesShrinkRequest
	GetUserId() *string
}

type DeleteMem0MemoriesShrinkRequest struct {
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
	MetadataShrink *string `json:"metadata,omitempty" xml:"metadata,omitempty"`
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

func (s DeleteMem0MemoriesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMem0MemoriesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteMem0MemoriesShrinkRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *DeleteMem0MemoriesShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteMem0MemoriesShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteMem0MemoriesShrinkRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *DeleteMem0MemoriesShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *DeleteMem0MemoriesShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *DeleteMem0MemoriesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteMem0MemoriesShrinkRequest) GetRunId() *string {
	return s.RunId
}

func (s *DeleteMem0MemoriesShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteMem0MemoriesShrinkRequest) SetAgentSpace(v string) *DeleteMem0MemoriesShrinkRequest {
	s.AgentSpace = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetAgentId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetAppId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetContextStoreId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.ContextStoreId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetMetadataShrink(v string) *DeleteMem0MemoriesShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetOrgId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetProjectId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetRunId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.RunId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) SetUserId(v string) *DeleteMem0MemoriesShrinkRequest {
	s.UserId = &v
	return s
}

func (s *DeleteMem0MemoriesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
