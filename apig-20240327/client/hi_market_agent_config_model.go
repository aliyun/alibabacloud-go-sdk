// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketAgentConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgentAPIConfig(v *HiMarketAgentConfigAgentAPIConfig) *HiMarketAgentConfig
	GetAgentAPIConfig() *HiMarketAgentConfigAgentAPIConfig
}

type HiMarketAgentConfig struct {
	AgentAPIConfig *HiMarketAgentConfigAgentAPIConfig `json:"agentAPIConfig,omitempty" xml:"agentAPIConfig,omitempty" type:"Struct"`
}

func (s HiMarketAgentConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketAgentConfig) GoString() string {
	return s.String()
}

func (s *HiMarketAgentConfig) GetAgentAPIConfig() *HiMarketAgentConfigAgentAPIConfig {
	return s.AgentAPIConfig
}

func (s *HiMarketAgentConfig) SetAgentAPIConfig(v *HiMarketAgentConfigAgentAPIConfig) *HiMarketAgentConfig {
	s.AgentAPIConfig = v
	return s
}

func (s *HiMarketAgentConfig) Validate() error {
	if s.AgentAPIConfig != nil {
		if err := s.AgentAPIConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketAgentConfigAgentAPIConfig struct {
	AgentProtocols []*string            `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	Routes         []*HiMarketHttpRoute `json:"routes,omitempty" xml:"routes,omitempty" type:"Repeated"`
}

func (s HiMarketAgentConfigAgentAPIConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketAgentConfigAgentAPIConfig) GoString() string {
	return s.String()
}

func (s *HiMarketAgentConfigAgentAPIConfig) GetAgentProtocols() []*string {
	return s.AgentProtocols
}

func (s *HiMarketAgentConfigAgentAPIConfig) GetRoutes() []*HiMarketHttpRoute {
	return s.Routes
}

func (s *HiMarketAgentConfigAgentAPIConfig) SetAgentProtocols(v []*string) *HiMarketAgentConfigAgentAPIConfig {
	s.AgentProtocols = v
	return s
}

func (s *HiMarketAgentConfigAgentAPIConfig) SetRoutes(v []*HiMarketHttpRoute) *HiMarketAgentConfigAgentAPIConfig {
	s.Routes = v
	return s
}

func (s *HiMarketAgentConfigAgentAPIConfig) Validate() error {
	if s.Routes != nil {
		for _, item := range s.Routes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
