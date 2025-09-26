// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRuntimeEndpoint interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeEndpointArn(v string) *AgentRuntimeEndpoint
	GetAgentRuntimeEndpointArn() *string
	SetAgentRuntimeEndpointId(v string) *AgentRuntimeEndpoint
	GetAgentRuntimeEndpointId() *string
	SetAgentRuntimeEndpointName(v string) *AgentRuntimeEndpoint
	GetAgentRuntimeEndpointName() *string
	SetAgentRuntimeId(v string) *AgentRuntimeEndpoint
	GetAgentRuntimeId() *string
	SetDescription(v string) *AgentRuntimeEndpoint
	GetDescription() *string
	SetEndpointPublicUrl(v string) *AgentRuntimeEndpoint
	GetEndpointPublicUrl() *string
	SetRoutingConfiguration(v *RoutingConfiguration) *AgentRuntimeEndpoint
	GetRoutingConfiguration() *RoutingConfiguration
	SetStatus(v string) *AgentRuntimeEndpoint
	GetStatus() *string
	SetStatusReason(v string) *AgentRuntimeEndpoint
	GetStatusReason() *string
	SetTargetVersion(v string) *AgentRuntimeEndpoint
	GetTargetVersion() *string
}

type AgentRuntimeEndpoint struct {
	AgentRuntimeEndpointArn  *string `json:"agentRuntimeEndpointArn,omitempty" xml:"agentRuntimeEndpointArn,omitempty"`
	AgentRuntimeEndpointId   *string `json:"agentRuntimeEndpointId,omitempty" xml:"agentRuntimeEndpointId,omitempty"`
	AgentRuntimeEndpointName *string `json:"agentRuntimeEndpointName,omitempty" xml:"agentRuntimeEndpointName,omitempty"`
	AgentRuntimeId           *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	Description              *string `json:"description,omitempty" xml:"description,omitempty"`
	// 智能体运行时端点的公网访问地址
	EndpointPublicUrl *string `json:"endpointPublicUrl,omitempty" xml:"endpointPublicUrl,omitempty"`
	// 智能体运行时端点的路由配置，支持多版本权重分配
	RoutingConfiguration *RoutingConfiguration `json:"routingConfiguration,omitempty" xml:"routingConfiguration,omitempty"`
	Status               *string               `json:"status,omitempty" xml:"status,omitempty"`
	StatusReason         *string               `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	TargetVersion        *string               `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s AgentRuntimeEndpoint) String() string {
	return dara.Prettify(s)
}

func (s AgentRuntimeEndpoint) GoString() string {
	return s.String()
}

func (s *AgentRuntimeEndpoint) GetAgentRuntimeEndpointArn() *string {
	return s.AgentRuntimeEndpointArn
}

func (s *AgentRuntimeEndpoint) GetAgentRuntimeEndpointId() *string {
	return s.AgentRuntimeEndpointId
}

func (s *AgentRuntimeEndpoint) GetAgentRuntimeEndpointName() *string {
	return s.AgentRuntimeEndpointName
}

func (s *AgentRuntimeEndpoint) GetAgentRuntimeId() *string {
	return s.AgentRuntimeId
}

func (s *AgentRuntimeEndpoint) GetDescription() *string {
	return s.Description
}

func (s *AgentRuntimeEndpoint) GetEndpointPublicUrl() *string {
	return s.EndpointPublicUrl
}

func (s *AgentRuntimeEndpoint) GetRoutingConfiguration() *RoutingConfiguration {
	return s.RoutingConfiguration
}

func (s *AgentRuntimeEndpoint) GetStatus() *string {
	return s.Status
}

func (s *AgentRuntimeEndpoint) GetStatusReason() *string {
	return s.StatusReason
}

func (s *AgentRuntimeEndpoint) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *AgentRuntimeEndpoint) SetAgentRuntimeEndpointArn(v string) *AgentRuntimeEndpoint {
	s.AgentRuntimeEndpointArn = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetAgentRuntimeEndpointId(v string) *AgentRuntimeEndpoint {
	s.AgentRuntimeEndpointId = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetAgentRuntimeEndpointName(v string) *AgentRuntimeEndpoint {
	s.AgentRuntimeEndpointName = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetAgentRuntimeId(v string) *AgentRuntimeEndpoint {
	s.AgentRuntimeId = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetDescription(v string) *AgentRuntimeEndpoint {
	s.Description = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetEndpointPublicUrl(v string) *AgentRuntimeEndpoint {
	s.EndpointPublicUrl = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetRoutingConfiguration(v *RoutingConfiguration) *AgentRuntimeEndpoint {
	s.RoutingConfiguration = v
	return s
}

func (s *AgentRuntimeEndpoint) SetStatus(v string) *AgentRuntimeEndpoint {
	s.Status = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetStatusReason(v string) *AgentRuntimeEndpoint {
	s.StatusReason = &v
	return s
}

func (s *AgentRuntimeEndpoint) SetTargetVersion(v string) *AgentRuntimeEndpoint {
	s.TargetVersion = &v
	return s
}

func (s *AgentRuntimeEndpoint) Validate() error {
	return dara.Validate(s)
}
