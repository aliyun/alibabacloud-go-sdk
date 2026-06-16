// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMem0MemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *UpdateMem0MemoryRequest
	GetAgentSpace() *string
	SetBody(v map[string]interface{}) *UpdateMem0MemoryRequest
	GetBody() map[string]interface{}
	SetContextStoreId(v string) *UpdateMem0MemoryRequest
	GetContextStoreId() *string
	SetOrgId(v string) *UpdateMem0MemoryRequest
	GetOrgId() *string
	SetProjectId(v string) *UpdateMem0MemoryRequest
	GetProjectId() *string
}

type UpdateMem0MemoryRequest struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// {"text":"用户改成喜欢喝美式","metadata":{"channel":"app"}}
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
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

func (s UpdateMem0MemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMem0MemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMem0MemoryRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *UpdateMem0MemoryRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *UpdateMem0MemoryRequest) GetContextStoreId() *string {
	return s.ContextStoreId
}

func (s *UpdateMem0MemoryRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *UpdateMem0MemoryRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateMem0MemoryRequest) SetAgentSpace(v string) *UpdateMem0MemoryRequest {
	s.AgentSpace = &v
	return s
}

func (s *UpdateMem0MemoryRequest) SetBody(v map[string]interface{}) *UpdateMem0MemoryRequest {
	s.Body = v
	return s
}

func (s *UpdateMem0MemoryRequest) SetContextStoreId(v string) *UpdateMem0MemoryRequest {
	s.ContextStoreId = &v
	return s
}

func (s *UpdateMem0MemoryRequest) SetOrgId(v string) *UpdateMem0MemoryRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateMem0MemoryRequest) SetProjectId(v string) *UpdateMem0MemoryRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateMem0MemoryRequest) Validate() error {
	return dara.Validate(s)
}
