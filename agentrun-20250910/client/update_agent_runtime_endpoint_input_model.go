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
	SetDeleteScalingConfig(v bool) *UpdateAgentRuntimeEndpointInput
	GetDeleteScalingConfig() *bool
	SetDescription(v string) *UpdateAgentRuntimeEndpointInput
	GetDescription() *string
	SetDisablePublicNetworkAccess(v bool) *UpdateAgentRuntimeEndpointInput
	GetDisablePublicNetworkAccess() *bool
	SetRoutingConfiguration(v *RoutingConfiguration) *UpdateAgentRuntimeEndpointInput
	GetRoutingConfiguration() *RoutingConfiguration
	SetScalingConfig(v *ScalingConfig) *UpdateAgentRuntimeEndpointInput
	GetScalingConfig() *ScalingConfig
	SetTargetVersion(v string) *UpdateAgentRuntimeEndpointInput
	GetTargetVersion() *string
}

type UpdateAgentRuntimeEndpointInput struct {
	// Agent runtime endpoint Name
	//
	// example:
	//
	// production-endpoint
	AgentRuntimeEndpointName *string `json:"agentRuntimeEndpointName,omitempty" xml:"agentRuntimeEndpointName,omitempty"`
	// 为 true 时删除该端点的弹性配置
	DeleteScalingConfig *bool `json:"deleteScalingConfig,omitempty" xml:"deleteScalingConfig,omitempty"`
	// Description
	//
	// example:
	//
	// Updated endpoint configuration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Whether to disable public network access for this endpoint
	DisablePublicNetworkAccess *bool `json:"disablePublicNetworkAccess,omitempty" xml:"disablePublicNetworkAccess,omitempty"`
	// The Ingress configuration of the agent runtime endpoint, which supports weight assignment across multiple versions
	//
	// example:
	//
	// {}
	RoutingConfiguration *RoutingConfiguration `json:"routingConfiguration,omitempty" xml:"routingConfiguration,omitempty"`
	// 端点的弹性伸缩配置，包括最小实例数和定时扩容策略（复用 ScalingConfig）
	ScalingConfig *ScalingConfig `json:"scalingConfig,omitempty" xml:"scalingConfig,omitempty"`
	// Target Version of the agent runtime
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

func (s *UpdateAgentRuntimeEndpointInput) GetDeleteScalingConfig() *bool {
	return s.DeleteScalingConfig
}

func (s *UpdateAgentRuntimeEndpointInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentRuntimeEndpointInput) GetDisablePublicNetworkAccess() *bool {
	return s.DisablePublicNetworkAccess
}

func (s *UpdateAgentRuntimeEndpointInput) GetRoutingConfiguration() *RoutingConfiguration {
	return s.RoutingConfiguration
}

func (s *UpdateAgentRuntimeEndpointInput) GetScalingConfig() *ScalingConfig {
	return s.ScalingConfig
}

func (s *UpdateAgentRuntimeEndpointInput) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UpdateAgentRuntimeEndpointInput) SetAgentRuntimeEndpointName(v string) *UpdateAgentRuntimeEndpointInput {
	s.AgentRuntimeEndpointName = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetDeleteScalingConfig(v bool) *UpdateAgentRuntimeEndpointInput {
	s.DeleteScalingConfig = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetDescription(v string) *UpdateAgentRuntimeEndpointInput {
	s.Description = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetDisablePublicNetworkAccess(v bool) *UpdateAgentRuntimeEndpointInput {
	s.DisablePublicNetworkAccess = &v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetRoutingConfiguration(v *RoutingConfiguration) *UpdateAgentRuntimeEndpointInput {
	s.RoutingConfiguration = v
	return s
}

func (s *UpdateAgentRuntimeEndpointInput) SetScalingConfig(v *ScalingConfig) *UpdateAgentRuntimeEndpointInput {
	s.ScalingConfig = v
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
	if s.ScalingConfig != nil {
		if err := s.ScalingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
