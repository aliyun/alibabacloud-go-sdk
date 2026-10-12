// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApigRefConfig(v *CreateProductRefRequestApigRefConfig) *CreateProductRefRequest
	GetApigRefConfig() *CreateProductRefRequestApigRefConfig
	SetGatewayId(v string) *CreateProductRefRequest
	GetGatewayId() *string
}

type CreateProductRefRequest struct {
	// The API configuration associated with the API product.
	ApigRefConfig *CreateProductRefRequestApigRefConfig `json:"apigRefConfig,omitempty" xml:"apigRefConfig,omitempty" type:"Struct"`
	// The ID of the gateway.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
}

func (s CreateProductRefRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRefRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRefRequest) GetApigRefConfig() *CreateProductRefRequestApigRefConfig {
	return s.ApigRefConfig
}

func (s *CreateProductRefRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateProductRefRequest) SetApigRefConfig(v *CreateProductRefRequestApigRefConfig) *CreateProductRefRequest {
	s.ApigRefConfig = v
	return s
}

func (s *CreateProductRefRequest) SetGatewayId(v string) *CreateProductRefRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateProductRefRequest) Validate() error {
	if s.ApigRefConfig != nil {
		if err := s.ApigRefConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProductRefRequestApigRefConfig struct {
	// The ID of the associated agent API.
	//
	// example:
	//
	// api-xxxxxxxx
	AgentApiId *string `json:"agentApiId,omitempty" xml:"agentApiId,omitempty"`
	// The name of the associated agent API.
	//
	// example:
	//
	// petstore
	AgentApiName *string `json:"agentApiName,omitempty" xml:"agentApiName,omitempty"`
	// The ID of the associated API.
	//
	// example:
	//
	// api-xxx
	ApiId *string `json:"apiId,omitempty" xml:"apiId,omitempty"`
	// The name of the associated API.
	//
	// example:
	//
	// petstore
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// The ID of the associated MCP route.
	//
	// example:
	//
	// hr-xxxxxxxx
	McpRouteId *string `json:"mcpRouteId,omitempty" xml:"mcpRouteId,omitempty"`
	// The ID of the associated MCP server.
	//
	// example:
	//
	// mcp-adfefz24afg
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name of the associated MCP server.
	//
	// example:
	//
	// test-mcp
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// The ID of the associated model API.
	//
	// example:
	//
	// api-xxxxxxxx
	ModelApiId *string `json:"modelApiId,omitempty" xml:"modelApiId,omitempty"`
	// The name of the associated model API.
	//
	// example:
	//
	// petstore
	ModelApiName *string `json:"modelApiName,omitempty" xml:"modelApiName,omitempty"`
}

func (s CreateProductRefRequestApigRefConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRefRequestApigRefConfig) GoString() string {
	return s.String()
}

func (s *CreateProductRefRequestApigRefConfig) GetAgentApiId() *string {
	return s.AgentApiId
}

func (s *CreateProductRefRequestApigRefConfig) GetAgentApiName() *string {
	return s.AgentApiName
}

func (s *CreateProductRefRequestApigRefConfig) GetApiId() *string {
	return s.ApiId
}

func (s *CreateProductRefRequestApigRefConfig) GetApiName() *string {
	return s.ApiName
}

func (s *CreateProductRefRequestApigRefConfig) GetMcpRouteId() *string {
	return s.McpRouteId
}

func (s *CreateProductRefRequestApigRefConfig) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateProductRefRequestApigRefConfig) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *CreateProductRefRequestApigRefConfig) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *CreateProductRefRequestApigRefConfig) GetModelApiName() *string {
	return s.ModelApiName
}

func (s *CreateProductRefRequestApigRefConfig) SetAgentApiId(v string) *CreateProductRefRequestApigRefConfig {
	s.AgentApiId = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetAgentApiName(v string) *CreateProductRefRequestApigRefConfig {
	s.AgentApiName = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetApiId(v string) *CreateProductRefRequestApigRefConfig {
	s.ApiId = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetApiName(v string) *CreateProductRefRequestApigRefConfig {
	s.ApiName = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetMcpRouteId(v string) *CreateProductRefRequestApigRefConfig {
	s.McpRouteId = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetMcpServerId(v string) *CreateProductRefRequestApigRefConfig {
	s.McpServerId = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetMcpServerName(v string) *CreateProductRefRequestApigRefConfig {
	s.McpServerName = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetModelApiId(v string) *CreateProductRefRequestApigRefConfig {
	s.ModelApiId = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) SetModelApiName(v string) *CreateProductRefRequestApigRefConfig {
	s.ModelApiName = &v
	return s
}

func (s *CreateProductRefRequestApigRefConfig) Validate() error {
	return dara.Validate(s)
}
