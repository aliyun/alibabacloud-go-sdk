// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscoveryEndpoint interface {
	dara.Model
	String() string
	GoString() string
	SetAgentEndpointConfigs(v []*AgentEndpointConfig) *DiscoveryEndpoint
	GetAgentEndpointConfigs() []*AgentEndpointConfig
	SetCredentialName(v string) *DiscoveryEndpoint
	GetCredentialName() *string
	SetName(v string) *DiscoveryEndpoint
	GetName() *string
	SetReturnAgentCredentialContent(v bool) *DiscoveryEndpoint
	GetReturnAgentCredentialContent() *bool
}

type DiscoveryEndpoint struct {
	AgentEndpointConfigs []*AgentEndpointConfig `json:"agentEndpointConfigs,omitempty" xml:"agentEndpointConfigs,omitempty" type:"Repeated"`
	// 该发现端点使用的凭证名称
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否在发现结果中返回 agent 的凭证内容
	ReturnAgentCredentialContent *bool `json:"returnAgentCredentialContent,omitempty" xml:"returnAgentCredentialContent,omitempty"`
}

func (s DiscoveryEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DiscoveryEndpoint) GoString() string {
	return s.String()
}

func (s *DiscoveryEndpoint) GetAgentEndpointConfigs() []*AgentEndpointConfig {
	return s.AgentEndpointConfigs
}

func (s *DiscoveryEndpoint) GetCredentialName() *string {
	return s.CredentialName
}

func (s *DiscoveryEndpoint) GetName() *string {
	return s.Name
}

func (s *DiscoveryEndpoint) GetReturnAgentCredentialContent() *bool {
	return s.ReturnAgentCredentialContent
}

func (s *DiscoveryEndpoint) SetAgentEndpointConfigs(v []*AgentEndpointConfig) *DiscoveryEndpoint {
	s.AgentEndpointConfigs = v
	return s
}

func (s *DiscoveryEndpoint) SetCredentialName(v string) *DiscoveryEndpoint {
	s.CredentialName = &v
	return s
}

func (s *DiscoveryEndpoint) SetName(v string) *DiscoveryEndpoint {
	s.Name = &v
	return s
}

func (s *DiscoveryEndpoint) SetReturnAgentCredentialContent(v bool) *DiscoveryEndpoint {
	s.ReturnAgentCredentialContent = &v
	return s
}

func (s *DiscoveryEndpoint) Validate() error {
	if s.AgentEndpointConfigs != nil {
		for _, item := range s.AgentEndpointConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
