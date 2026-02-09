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
	// Assembled MCP server sources
	AssembledSources []*CreateMcpServerRequestAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	// Backend configuration
	BackendConfig *CreateMcpServerRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// Creation source type
	//
	// example:
	//
	// ApiGatewayMcpHosting
	CreateFromType *string `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	// MCP server description
	//
	// example:
	//
	// mcp tool fetch time
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// List of domain IDs for the MCP server
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// Exposed URI path for SSE/StreamableHTTP protocols
	//
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// Gateway ID
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qac0
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Gray route configurations
	GrayMcpServerConfigs []*CreateMcpServerRequestGrayMcpServerConfigs `json:"grayMcpServerConfigs,omitempty" xml:"grayMcpServerConfigs,omitempty" type:"Repeated"`
	// Route matching conditions
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// MCP server specification
	McpServerConfig *CreateMcpServerRequestMcpServerConfig `json:"mcpServerConfig,omitempty" xml:"mcpServerConfig,omitempty" type:"Struct"`
	// Enable MCP statistics
	//
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// MCP server name
	//
	// This parameter is required.
	//
	// example:
	//
	// fetch-time
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// MCP protocol
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// MCP server type
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
	// Source MCP server name
	//
	// example:
	//
	// test-mcp
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// List of tool names to include
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
	// Backend scene type
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// List of backend services
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
	// Service port
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Service protocol
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Service ID
	//
	// example:
	//
	// svc-crbgq0dlhtgr***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// Service version
	//
	// example:
	//
	// 2.1.6
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Service weight
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
	// Backend configuration for gray route
	BackendConfig *CreateMcpServerRequestGrayMcpServerConfigsBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// Route matching rules
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// Route ID for update operations
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
	// Must be SingleService
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// Exactly one service
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
	// Service port number
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Service protocol type
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Service ID
	//
	// example:
	//
	// svc-gray
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// Service version
	//
	// example:
	//
	// v2.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Service weight for load balancing
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
	// Converted MCP server spec YAML
	//
	// example:
	//
	// mcp-spec.yaml
	McpServerSpec *string `json:"mcpServerSpec,omitempty" xml:"mcpServerSpec,omitempty"`
	// Raw Swagger/OpenAPI document
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
