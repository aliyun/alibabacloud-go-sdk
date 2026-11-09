// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentInstanceVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListAgentInstanceVersionsRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ListAgentInstanceVersionsRequest
	GetAgentProvider() *string
}

type ListAgentInstanceVersionsRequest struct {
	// The agent platform used to filter version numbers by business type. If not specified, the default value is ENTERPRISE.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The name of the agent provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
}

func (s ListAgentInstanceVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentInstanceVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentInstanceVersionsRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListAgentInstanceVersionsRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListAgentInstanceVersionsRequest) SetAgentPlatform(v string) *ListAgentInstanceVersionsRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListAgentInstanceVersionsRequest) SetAgentProvider(v string) *ListAgentInstanceVersionsRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListAgentInstanceVersionsRequest) Validate() error {
	return dara.Validate(s)
}
