// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentEndpoint(v string) *CreateAgentInfo
	GetAgentEndpoint() *string
	SetDeployMethod(v string) *CreateAgentInfo
	GetDeployMethod() *string
	SetName(v string) *CreateAgentInfo
	GetName() *string
	SetTags(v string) *CreateAgentInfo
	GetTags() *string
	SetTunnelId(v string) *CreateAgentInfo
	GetTunnelId() *string
}

type CreateAgentInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// vpc
	AgentEndpoint *string `json:"AgentEndpoint,omitempty" xml:"AgentEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	DeployMethod *string `json:"DeployMethod,omitempty" xml:"DeployMethod,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_tunnel_id
	TunnelId *string `json:"TunnelId,omitempty" xml:"TunnelId,omitempty"`
}

func (s CreateAgentInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentInfo) GoString() string {
	return s.String()
}

func (s *CreateAgentInfo) GetAgentEndpoint() *string {
	return s.AgentEndpoint
}

func (s *CreateAgentInfo) GetDeployMethod() *string {
	return s.DeployMethod
}

func (s *CreateAgentInfo) GetName() *string {
	return s.Name
}

func (s *CreateAgentInfo) GetTags() *string {
	return s.Tags
}

func (s *CreateAgentInfo) GetTunnelId() *string {
	return s.TunnelId
}

func (s *CreateAgentInfo) SetAgentEndpoint(v string) *CreateAgentInfo {
	s.AgentEndpoint = &v
	return s
}

func (s *CreateAgentInfo) SetDeployMethod(v string) *CreateAgentInfo {
	s.DeployMethod = &v
	return s
}

func (s *CreateAgentInfo) SetName(v string) *CreateAgentInfo {
	s.Name = &v
	return s
}

func (s *CreateAgentInfo) SetTags(v string) *CreateAgentInfo {
	s.Tags = &v
	return s
}

func (s *CreateAgentInfo) SetTunnelId(v string) *CreateAgentInfo {
	s.TunnelId = &v
	return s
}

func (s *CreateAgentInfo) Validate() error {
	return dara.Validate(s)
}
