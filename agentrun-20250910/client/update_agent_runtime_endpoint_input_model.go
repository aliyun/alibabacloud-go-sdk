// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRuntimeEndpointInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeEndpointName(v string) *UpdateAgentRuntimeEndpointInput
	GetAgentRuntimeEndpointName() *string
	SetDescription(v string) *UpdateAgentRuntimeEndpointInput
	GetDescription() *string
	SetRoutingConfiguration(v *RoutingConfiguration) *UpdateAgentRuntimeEndpointInput
	GetRoutingConfiguration() *RoutingConfiguration
	SetTargetVersion(v string) *UpdateAgentRuntimeEndpointInput
	GetTargetVersion() *string
}

type UpdateAgentRuntimeEndpointInput struct {
	// example:
	//
	// production-endpoint
	AgentRuntimeEndpointName *string `json:"agentRuntimeEndpointName,omitempty" xml:"agentRuntimeEndpointName,omitempty"`
	// example:
	//
	// Updated endpoint configuration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 智能体运行时端点的路由配置，支持多版本权重分配
	//
	// example:
	//
	// {}
	RoutingConfiguration *RoutingConfiguration `json:"routingConfiguration,omitempty" xml:"routingConfiguration,omitempty"`
	// 智能体运行时的目标版本
	//
	// example:
	//
	// LATEST
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s UpdateAgentRuntimeEndpointInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRuntimeEndpointInput) GoString() string {
	return s.String()
}

func (s *UpdateAgentRuntimeEndpointInput) GetAgentRuntimeEndpointName() *string {
	return s.AgentRuntimeEndpointName
}

func (s *UpdateAgentRuntimeEndpointInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentRuntimeEndpointInput) GetRoutingConfiguration() *RoutingConfiguration {
	return s.RoutingConfiguration
}

func (s *UpdateAgentRuntimeEndpointInput) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UpdateAgentRuntimeEndpointInput) SetAgentRuntimeEndpointName(v string) *UpdateAgentRuntimeEndpointInput {
	s.AgentRuntimeEndpointName = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetDescription(v string) *UpdateAgentRuntimeEndpointInput {
	s.Description = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetRoutingConfiguration(v *RoutingConfiguration) *UpdateAgentRuntimeEndpointInput {
	s.RoutingConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetTargetVersion(v string) *UpdateAgentRuntimeEndpointInput {
	s.TargetVersion = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) Validate() error {
	if s.RoutingConfiguration != nil {
		if err := s.RoutingConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}
