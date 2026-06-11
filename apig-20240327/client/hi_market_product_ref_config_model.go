// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketProductRefConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApigRefConfig(v *HiMarketProductRefConfigApigRefConfig) *HiMarketProductRefConfig
	GetApigRefConfig() *HiMarketProductRefConfigApigRefConfig
	SetGatewayId(v string) *HiMarketProductRefConfig
	GetGatewayId() *string
}

type HiMarketProductRefConfig struct {
	// Reference settings for the API gateway.
	ApigRefConfig *HiMarketProductRefConfigApigRefConfig `json:"apigRefConfig,omitempty" xml:"apigRefConfig,omitempty" type:"Struct"`
	// Unique identifier for the gateway.
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s HiMarketProductRefConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketProductRefConfig) GoString() string {
	return s.String()
}

func (s *HiMarketProductRefConfig) GetApigRefConfig() *HiMarketProductRefConfigApigRefConfig {
	return s.ApigRefConfig
}

func (s *HiMarketProductRefConfig) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HiMarketProductRefConfig) SetApigRefConfig(v *HiMarketProductRefConfigApigRefConfig) *HiMarketProductRefConfig {
	s.ApigRefConfig = v
	return s
}

func (s *HiMarketProductRefConfig) SetGatewayId(v string) *HiMarketProductRefConfig {
	s.GatewayId = &v
	return s
}

func (s *HiMarketProductRefConfig) Validate() error {
	if s.ApigRefConfig != nil {
		if err := s.ApigRefConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HiMarketProductRefConfigApigRefConfig struct {
	// Unique identifier for the agent API.
	AgentApiId *string `json:"agentApiId,omitempty" xml:"agentApiId,omitempty"`
	// The name of the agent API.
	AgentApiName *string `json:"agentApiName,omitempty" xml:"agentApiName,omitempty"`
	// Unique identifier for the MCP route.
	McpRouteId *string `json:"mcpRouteId,omitempty" xml:"mcpRouteId,omitempty"`
	// Unique identifier for the MCP server.
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name of the MCP server.
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// Unique identifier for the model API.
	ModelApiId *string `json:"modelApiId,omitempty" xml:"modelApiId,omitempty"`
	// The name of the model API.
	ModelApiName *string `json:"modelApiName,omitempty" xml:"modelApiName,omitempty"`
}

func (s HiMarketProductRefConfigApigRefConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketProductRefConfigApigRefConfig) GoString() string {
	return s.String()
}

func (s *HiMarketProductRefConfigApigRefConfig) GetAgentApiId() *string {
	return s.AgentApiId
}

func (s *HiMarketProductRefConfigApigRefConfig) GetAgentApiName() *string {
	return s.AgentApiName
}

func (s *HiMarketProductRefConfigApigRefConfig) GetMcpRouteId() *string {
	return s.McpRouteId
}

func (s *HiMarketProductRefConfigApigRefConfig) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *HiMarketProductRefConfigApigRefConfig) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *HiMarketProductRefConfigApigRefConfig) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *HiMarketProductRefConfigApigRefConfig) GetModelApiName() *string {
	return s.ModelApiName
}

func (s *HiMarketProductRefConfigApigRefConfig) SetAgentApiId(v string) *HiMarketProductRefConfigApigRefConfig {
	s.AgentApiId = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) SetAgentApiName(v string) *HiMarketProductRefConfigApigRefConfig {
	s.AgentApiName = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) SetMcpRouteId(v string) *HiMarketProductRefConfigApigRefConfig {
	s.McpRouteId = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) SetMcpServerId(v string) *HiMarketProductRefConfigApigRefConfig {
	s.McpServerId = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) SetMcpServerName(v string) *HiMarketProductRefConfigApigRefConfig {
	s.McpServerName = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) SetModelApiId(v string) *HiMarketProductRefConfigApigRefConfig {
	s.ModelApiId = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) SetModelApiName(v string) *HiMarketProductRefConfigApigRefConfig {
	s.ModelApiName = &v
	return s
}

func (s *HiMarketProductRefConfigApigRefConfig) Validate() error {
	return dara.Validate(s)
}
