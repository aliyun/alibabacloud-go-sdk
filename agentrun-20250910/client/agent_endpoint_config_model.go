// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentEndpointConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *AgentEndpointConfig
	GetAgentName() *string
	SetCustomDomainUrl(v string) *AgentEndpointConfig
	GetCustomDomainUrl() *string
	SetEndpointName(v string) *AgentEndpointConfig
	GetEndpointName() *string
	SetEndpointUrl(v string) *AgentEndpointConfig
	GetEndpointUrl() *string
}

type AgentEndpointConfig struct {
	AgentName       *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	CustomDomainUrl *string `json:"customDomainUrl,omitempty" xml:"customDomainUrl,omitempty"`
	// 端点名称
	EndpointName *string `json:"endpointName,omitempty" xml:"endpointName,omitempty"`
	EndpointUrl  *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
}

func (s AgentEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s AgentEndpointConfig) GoString() string {
	return s.String()
}

func (s *AgentEndpointConfig) GetAgentName() *string {
	return s.AgentName
}

func (s *AgentEndpointConfig) GetCustomDomainUrl() *string {
	return s.CustomDomainUrl
}

func (s *AgentEndpointConfig) GetEndpointName() *string {
	return s.EndpointName
}

func (s *AgentEndpointConfig) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *AgentEndpointConfig) SetAgentName(v string) *AgentEndpointConfig {
	s.AgentName = &v
	return s
}

func (s *AgentEndpointConfig) SetCustomDomainUrl(v string) *AgentEndpointConfig {
	s.CustomDomainUrl = &v
	return s
}

func (s *AgentEndpointConfig) SetEndpointName(v string) *AgentEndpointConfig {
	s.EndpointName = &v
	return s
}

func (s *AgentEndpointConfig) SetEndpointUrl(v string) *AgentEndpointConfig {
	s.EndpointUrl = &v
	return s
}

func (s *AgentEndpointConfig) Validate() error {
	return dara.Validate(s)
}
