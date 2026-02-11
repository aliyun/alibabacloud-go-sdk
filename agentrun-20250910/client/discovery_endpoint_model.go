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
	SetName(v string) *DiscoveryEndpoint
	GetName() *string
}

type DiscoveryEndpoint struct {
	AgentEndpointConfigs []*AgentEndpointConfig `json:"agentEndpointConfigs,omitempty" xml:"agentEndpointConfigs,omitempty" type:"Repeated"`
	Name                 *string                `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *DiscoveryEndpoint) GetName() *string {
	return s.Name
}

func (s *DiscoveryEndpoint) SetAgentEndpointConfigs(v []*AgentEndpointConfig) *DiscoveryEndpoint {
	s.AgentEndpointConfigs = v
	return s
}

func (s *DiscoveryEndpoint) SetName(v string) *DiscoveryEndpoint {
	s.Name = &v
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
