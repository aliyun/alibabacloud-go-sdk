// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssembledSources(v []*CreateMcpServerRequestAssembledSources) *CreateMcpServerRequest
	GetAssembledSources() []*CreateMcpServerRequestAssembledSources
	SetBackendConfig(v *CreateMcpServerRequestBackendConfig) *CreateMcpServerRequest
	GetBackendConfig() *CreateMcpServerRequestBackendConfig
	SetCreateFromType(v string) *CreateMcpServerRequest
	GetCreateFromType() *string
	SetDescription(v string) *CreateMcpServerRequest
	GetDescription() *string
	SetDomainIds(v []*string) *CreateMcpServerRequest
	GetDomainIds() []*string
	SetExposedUriPath(v string) *CreateMcpServerRequest
	GetExposedUriPath() *string
	SetGatewayId(v string) *CreateMcpServerRequest
	GetGatewayId() *string
	SetGrayMcpServerConfigs(v []*CreateMcpServerRequestGrayMcpServerConfigs) *CreateMcpServerRequest
	GetGrayMcpServerConfigs() []*CreateMcpServerRequestGrayMcpServerConfigs
	SetMatch(v *HttpRouteMatch) *CreateMcpServerRequest
	GetMatch() *HttpRouteMatch
	SetMcpServerConfig(v *CreateMcpServerRequestMcpServerConfig) *CreateMcpServerRequest
	GetMcpServerConfig() *CreateMcpServerRequestMcpServerConfig
	SetMcpStatisticsEnable(v bool) *CreateMcpServerRequest
	GetMcpStatisticsEnable() *bool
	SetName(v string) *CreateMcpServerRequest
	GetName() *string
	SetProtocol(v string) *CreateMcpServerRequest
	GetProtocol() *string
	SetType(v string) *CreateMcpServerRequest
	GetType() *string
}

type CreateMcpServerRequest struct {
	// The list of assembly sources. This parameter is required when type is set to AssemblyMCP.
	AssembledSources []*CreateMcpServerRequestAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	// The backend service configuration of the route.
	BackendConfig *CreateMcpServerRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// The creation source type. Valid values:
	//
	// - ApiGatewayHttpToMCP: gateway-managed HTTP-to-MCP conversion
	//
	// - ApiGatewayProxyMcpHosting: gateway-managed direct MCP proxy
	//
	// - ApiGatewayAssembly: gateway MCP assembly
	//
	// - NacosHttpToMCP: gateway-managed Nacos-synced HTTP-to-MCP conversion
	//
	// - NacosMcpHosting: gateway-managed Nacos-synced direct MCP proxy
	//
	// example:
	//
	// ApiGatewayMcpHosting
	CreateFromType *string `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	// The description of the MCP server.
	//
	// example:
	//
	// mcp tool fetch time
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain name IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The exposed URI path. This parameter is required when protocol is set to SSE or StreamableHTTP and type is set to RealMCP.
	//
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// The gateway ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qac0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The canary release routing configurations.
	GrayMcpServerConfigs []*CreateMcpServerRequestGrayMcpServerConfigs `json:"grayMcpServerConfigs,omitempty" xml:"grayMcpServerConfigs,omitempty" type:"Repeated"`
	// The route match rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The HTTP-to-MCP configuration.
	McpServerConfig *CreateMcpServerRequestMcpServerConfig `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty" type:"Struct"`
	// Specifies whether to enable MCP observability. Default value: false.
	//
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// The MCP server name. The name must match the regular expression ^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$ and cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// fetch-time
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol type. Valid values: HTTP, HTTPS, SSE, and StreamableHTTP.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The type. Valid values:
	//
	// - RealMCP: standard MCP service
	//
	// This parameter is required.
	//
	// example:
	//
	// RealMCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequest) GetAssembledSources() []*CreateMcpServerRequestAssembledSources {
	return s.AssembledSources
}

func (s *CreateMcpServerRequest) GetBackendConfig() *CreateMcpServerRequestBackendConfig {
	return s.BackendConfig
}

func (s *CreateMcpServerRequest) GetCreateFromType() *string {
	return s.CreateFromType
}

func (s *CreateMcpServerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpServerRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *CreateMcpServerRequest) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *CreateMcpServerRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateMcpServerRequest) GetGrayMcpServerConfigs() []*CreateMcpServerRequestGrayMcpServerConfigs {
	return s.GrayMcpServerConfigs
}

func (s *CreateMcpServerRequest) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *CreateMcpServerRequest) GetMcpServerConfig() *CreateMcpServerRequestMcpServerConfig {
	return s.McpServerConfig
}

func (s *CreateMcpServerRequest) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *CreateMcpServerRequest) GetName() *string {
	return s.Name
}

func (s *CreateMcpServerRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpServerRequest) GetType() *string {
	return s.Type
}

func (s *CreateMcpServerRequest) SetAssembledSources(v []*CreateMcpServerRequestAssembledSources) *CreateMcpServerRequest {
	s.AssembledSources = v
	return s
}

func (s *CreateMcpServerRequest) SetBackendConfig(v *CreateMcpServerRequestBackendConfig) *CreateMcpServerRequest {
	s.BackendConfig = v
	return s
}

func (s *CreateMcpServerRequest) SetCreateFromType(v string) *CreateMcpServerRequest {
	s.CreateFromType = &v
	return s
}

func (s *CreateMcpServerRequest) SetDescription(v string) *CreateMcpServerRequest {
	s.Description = &v
	return s
}

func (s *CreateMcpServerRequest) SetDomainIds(v []*string) *CreateMcpServerRequest {
	s.DomainIds = v
	return s
}

func (s *CreateMcpServerRequest) SetExposedUriPath(v string) *CreateMcpServerRequest {
	s.ExposedUriPath = &v
	return s
}

func (s *CreateMcpServerRequest) SetGatewayId(v string) *CreateMcpServerRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateMcpServerRequest) SetGrayMcpServerConfigs(v []*CreateMcpServerRequestGrayMcpServerConfigs) *CreateMcpServerRequest {
	s.GrayMcpServerConfigs = v
	return s
}

func (s *CreateMcpServerRequest) SetMatch(v *HttpRouteMatch) *CreateMcpServerRequest {
	s.Match = v
	return s
}

func (s *CreateMcpServerRequest) SetMcpServerConfig(v *CreateMcpServerRequestMcpServerConfig) *CreateMcpServerRequest {
	s.McpServerConfig = v
	return s
}

func (s *CreateMcpServerRequest) SetMcpStatisticsEnable(v bool) *CreateMcpServerRequest {
	s.McpStatisticsEnable = &v
	return s
}

func (s *CreateMcpServerRequest) SetName(v string) *CreateMcpServerRequest {
	s.Name = &v
	return s
}

func (s *CreateMcpServerRequest) SetProtocol(v string) *CreateMcpServerRequest {
	s.Protocol = &v
	return s
}

func (s *CreateMcpServerRequest) SetType(v string) *CreateMcpServerRequest {
	s.Type = &v
	return s
}

func (s *CreateMcpServerRequest) Validate() error {
	if s.AssembledSources != nil {
		for _, item := range s.AssembledSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BackendConfig != nil {
		if err := s.BackendConfig.Validate(); err != nil {
			return err
		}
	}
	if s.GrayMcpServerConfigs != nil {
		for _, item := range s.GrayMcpServerConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.McpServerConfig != nil {
		if err := s.McpServerConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpServerRequestAssembledSources struct {
	// MCP Server ID
	//
	// example:
	//
	// mcp-sdfa3qgavz
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The MCP server name.
	//
	// example:
	//
	// test-mcp
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// The list of MCP tools.
	Tools []*string `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CreateMcpServerRequestAssembledSources) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestAssembledSources) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestAssembledSources) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateMcpServerRequestAssembledSources) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *CreateMcpServerRequestAssembledSources) GetTools() []*string {
	return s.Tools
}

func (s *CreateMcpServerRequestAssembledSources) SetMcpServerId(v string) *CreateMcpServerRequestAssembledSources {
	s.McpServerId = &v
	return s
}

func (s *CreateMcpServerRequestAssembledSources) SetMcpServerName(v string) *CreateMcpServerRequestAssembledSources {
	s.McpServerName = &v
	return s
}

func (s *CreateMcpServerRequestAssembledSources) SetTools(v []*string) *CreateMcpServerRequestAssembledSources {
	s.Tools = v
	return s
}

func (s *CreateMcpServerRequestAssembledSources) Validate() error {
	return dara.Validate(s)
}

type CreateMcpServerRequestBackendConfig struct {
	// The backend service scenario.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The backend services.
	Services []*CreateMcpServerRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s CreateMcpServerRequestBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *CreateMcpServerRequestBackendConfig) GetServices() []*CreateMcpServerRequestBackendConfigServices {
	return s.Services
}

func (s *CreateMcpServerRequestBackendConfig) SetScene(v string) *CreateMcpServerRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *CreateMcpServerRequestBackendConfig) SetServices(v []*CreateMcpServerRequestBackendConfigServices) *CreateMcpServerRequestBackendConfig {
	s.Services = v
	return s
}

func (s *CreateMcpServerRequestBackendConfig) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMcpServerRequestBackendConfigServices struct {
	// The backend node port of the service.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol. Valid values:
	//
	// - HTTP
	//
	// - HTTPS
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-crbgq0dlhtgr***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 2.1.6
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The weight.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateMcpServerRequestBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpServerRequestBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpServerRequestBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateMcpServerRequestBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *CreateMcpServerRequestBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateMcpServerRequestBackendConfigServices) SetPort(v int32) *CreateMcpServerRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *CreateMcpServerRequestBackendConfigServices) SetProtocol(v string) *CreateMcpServerRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *CreateMcpServerRequestBackendConfigServices) SetServiceId(v string) *CreateMcpServerRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *CreateMcpServerRequestBackendConfigServices) SetVersion(v string) *CreateMcpServerRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *CreateMcpServerRequestBackendConfigServices) SetWeight(v int32) *CreateMcpServerRequestBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *CreateMcpServerRequestBackendConfigServices) Validate() error {
	return dara.Validate(s)
}

type CreateMcpServerRequestGrayMcpServerConfigs struct {
	// The backend configuration.
	BackendConfig *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// The route match rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The route ID.
	//
	// example:
	//
	// gray-route-123
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s CreateMcpServerRequestGrayMcpServerConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestGrayMcpServerConfigs) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) GetBackendConfig() *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig {
	return s.BackendConfig
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) GetRouteId() *string {
	return s.RouteId
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) SetBackendConfig(v *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) *CreateMcpServerRequestGrayMcpServerConfigs {
	s.BackendConfig = v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) SetMatch(v *HttpRouteMatch) *CreateMcpServerRequestGrayMcpServerConfigs {
	s.Match = v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) SetRouteId(v string) *CreateMcpServerRequestGrayMcpServerConfigs {
	s.RouteId = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigs) Validate() error {
	if s.BackendConfig != nil {
		if err := s.BackendConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpServerRequestGrayMcpServerConfigsBackendConfig struct {
	// The backend scenario.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The backend services.
	Services []*CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) GetServices() []*CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices {
	return s.Services
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) SetScene(v string) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig {
	s.Scene = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) SetServices(v []*CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig {
	s.Services = v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices struct {
	// The service port.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-gray
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// v2.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) SetPort(v int32) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices {
	s.Port = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) SetProtocol(v string) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) SetServiceId(v string) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) SetVersion(v string) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices {
	s.Version = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) SetWeight(v int32) *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *CreateMcpServerRequestGrayMcpServerConfigsBackendConfigServices) Validate() error {
	return dara.Validate(s)
}

type CreateMcpServerRequestMcpServerConfig struct {
	// The MCP server spec configuration. This parameter is mutually exclusive with swaggerConfig.
	//
	// example:
	//
	// mcp-spec.yaml
	McpServerSpec *string `json:"mcpServerSpec,omitempty" xml:"mcpServerSpec,omitempty"`
	// The Swagger document for HTTP-to-MCP conversion. The document must comply with the OpenAPI 3.0 specification.
	//
	// example:
	//
	// swagger.yaml
	SwaggerConfig *string `json:"swaggerConfig,omitempty" xml:"swaggerConfig,omitempty"`
}

func (s CreateMcpServerRequestMcpServerConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerRequestMcpServerConfig) GoString() string {
	return s.String()
}

func (s *CreateMcpServerRequestMcpServerConfig) GetMcpServerSpec() *string {
	return s.McpServerSpec
}

func (s *CreateMcpServerRequestMcpServerConfig) GetSwaggerConfig() *string {
	return s.SwaggerConfig
}

func (s *CreateMcpServerRequestMcpServerConfig) SetMcpServerSpec(v string) *CreateMcpServerRequestMcpServerConfig {
	s.McpServerSpec = &v
	return s
}

func (s *CreateMcpServerRequestMcpServerConfig) SetSwaggerConfig(v string) *CreateMcpServerRequestMcpServerConfig {
	s.SwaggerConfig = &v
	return s
}

func (s *CreateMcpServerRequestMcpServerConfig) Validate() error {
	return dara.Validate(s)
}
