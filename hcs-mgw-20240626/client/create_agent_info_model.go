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
	// The method that is used to access the agent.\\
	//
	// Valid values: public and vpc.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	AgentEndpoint *string `json:"AgentEndpoint,omitempty" xml:"AgentEndpoint,omitempty"`
	// The deployment mode of the agent.\\
	//
	// Set the value to default, which specifies the independent process mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	DeployMethod *string `json:"DeployMethod,omitempty" xml:"DeployMethod,omitempty"`
	// The name of the agent.\\
	//
	// The name can contain lowercase letters, digits, hyphens (-), and underscores (_). The name must be 3 to 63 characters in length. The name is case-sensitive and encoded in UTF-8. The name cannot start with a hyphen (-) or an underscore (_). You must specify a name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags in the key:value format.\\
	//
	// The value can contain letters, digits, hyphens (-), underscores (_), and commas (,). The value can be up to 1,024 characters in length.
	//
	// example:
	//
	// K1:V1,K2:V2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The tunnel ID.
	//
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
