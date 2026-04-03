// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRuntimeEndpointInput interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeEndpointName(v string) *CreateAgentRuntimeEndpointInput
	GetAgentRuntimeEndpointName() *string
	SetDescription(v string) *CreateAgentRuntimeEndpointInput
	GetDescription() *string
	SetDisablePublicNetworkAccess(v bool) *CreateAgentRuntimeEndpointInput
	GetDisablePublicNetworkAccess() *bool
	SetRoutingConfiguration(v *RoutingConfiguration) *CreateAgentRuntimeEndpointInput
	GetRoutingConfiguration() *RoutingConfiguration
	SetScalingConfig(v *ScalingConfig) *CreateAgentRuntimeEndpointInput
	GetScalingConfig() *ScalingConfig
	SetTargetVersion(v string) *CreateAgentRuntimeEndpointInput
	GetTargetVersion() *string
}

type CreateAgentRuntimeEndpointInput struct {
	// example:
	//
	// production-endpoint
	AgentRuntimeEndpointName *string `json:"agentRuntimeEndpointName,omitempty" xml:"agentRuntimeEndpointName,omitempty"`
	// example:
	//
	// Production endpoint for customer service agent
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 是否禁用该端点的公网访问
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// 智能体运行时端点的路由配置，支持多版本权重分配
	//
	// example:
	//
	// {}
	RoutingConfiguration *RoutingConfiguration `json:"routingConfiguration,omitempty" xml:"routingConfiguration,omitempty"`
	// 端点的弹性伸缩配置，包括最小实例数和定时扩容策略（复用 ScalingConfig）
	ScalingConfig *ScalingConfig `json:"scalingConfig,omitempty" xml:"scalingConfig,omitempty"`
	// 智能体运行时的目标版本
	//
	// example:
	//
	// LATEST
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s CreateAgentRuntimeEndpointInput) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRuntimeEndpointInput) GoString() string {
	return s.String()
}

func (s *CreateAgentRuntimeEndpointInput) GetAgentRuntimeEndpointName() *string {
	return s.AgentRuntimeEndpointName
}

func (s *CreateAgentRuntimeEndpointInput) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRuntimeEndpointInput) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
}

func (s *CreateAgentRuntimeEndpointInput) GetRoutingConfiguration() *RoutingConfiguration {
	return s.RoutingConfiguration
}

func (s *CreateAgentRuntimeEndpointInput) GetScalingConfig() *ScalingConfig {
	return s.ScalingConfig
}

func (s *CreateAgentRuntimeEndpointInput) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *CreateAgentRuntimeEndpointInput) SetAgentRuntimeEndpointName(v string) *CreateAgentRuntimeEndpointInput {
	s.AgentRuntimeEndpointName = &v
	return s
}

func (s *CreateAgentRuntimeEndpointInput) SetDescription(v string) *CreateAgentRuntimeEndpointInput {
	s.Description = &v
	return s
}

func (s *CreateAgentRuntimeEndpointInput) SetDisablePublicNetworkAccess(v bool) *CreateAgentRuntimeEndpointInput {
	s.DisablePublicNetworkAccess = &v
	return s
}

func (s *CreateAgentRuntimeEndpointInput) SetRoutingConfiguration(v *RoutingConfiguration) *CreateAgentRuntimeEndpointInput {
	s.RoutingConfiguration = v
	return s
}

func (s *CreateAgentRuntimeEndpointInput) SetScalingConfig(v *ScalingConfig) *CreateAgentRuntimeEndpointInput {
	s.ScalingConfig = v
	return s
}

func (s *CreateAgentRuntimeEndpointInput) SetTargetVersion(v string) *CreateAgentRuntimeEndpointInput {
	s.TargetVersion = &v
	return s
}

func (s *CreateAgentRuntimeEndpointInput) Validate() error {
	if s.RoutingConfiguration != nil {
		if err := s.RoutingConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingConfig != nil {
		if err := s.ScalingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
