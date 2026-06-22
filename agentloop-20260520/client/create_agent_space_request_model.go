// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *CreateAgentSpaceRequest
	GetAgentSpace() *string
	SetCmsWorkspace(v string) *CreateAgentSpaceRequest
	GetCmsWorkspace() *string
	SetDescription(v string) *CreateAgentSpaceRequest
	GetDescription() *string
	SetClientToken(v string) *CreateAgentSpaceRequest
	GetClientToken() *string
}

type CreateAgentSpaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// example:
	//
	// test-cms-workspace
	CmsWorkspace *string `json:"cmsWorkspace,omitempty" xml:"cmsWorkspace,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// a1b2c3d4-1234-5678-90ab-cdef12345678
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateAgentSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSpaceRequest) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *CreateAgentSpaceRequest) GetCmsWorkspace() *string {
	return s.CmsWorkspace
}

func (s *CreateAgentSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentSpaceRequest) SetAgentSpace(v string) *CreateAgentSpaceRequest {
	s.AgentSpace = &v
	return s
}

func (s *CreateAgentSpaceRequest) SetCmsWorkspace(v string) *CreateAgentSpaceRequest {
	s.CmsWorkspace = &v
	return s
}

func (s *CreateAgentSpaceRequest) SetDescription(v string) *CreateAgentSpaceRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentSpaceRequest) SetClientToken(v string) *CreateAgentSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentSpaceRequest) Validate() error {
	return dara.Validate(s)
}
