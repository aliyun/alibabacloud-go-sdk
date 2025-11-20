// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssembledSources(v []*UpdateMcpServerRequestAssembledSources) *UpdateMcpServerRequest
	GetAssembledSources() []*UpdateMcpServerRequestAssembledSources
	SetBackendConfig(v *UpdateMcpServerRequestBackendConfig) *UpdateMcpServerRequest
	GetBackendConfig() *UpdateMcpServerRequestBackendConfig
	SetCreateFromType(v string) *UpdateMcpServerRequest
	GetCreateFromType() *string
	SetDescription(v string) *UpdateMcpServerRequest
	GetDescription() *string
	SetDomainIds(v []*string) *UpdateMcpServerRequest
	GetDomainIds() []*string
	SetExposedUriPath(v string) *UpdateMcpServerRequest
	GetExposedUriPath() *string
	SetMatch(v *HttpRouteMatch) *UpdateMcpServerRequest
	GetMatch() *HttpRouteMatch
	SetMcpStatisticsEnable(v bool) *UpdateMcpServerRequest
	GetMcpStatisticsEnable() *bool
	SetProtocol(v string) *UpdateMcpServerRequest
	GetProtocol() *string
	SetType(v string) *UpdateMcpServerRequest
	GetType() *string
}

type UpdateMcpServerRequest struct {
	// The list of assembly sources. This parameter is required when the type parameter is set to AssemblyMCP.
	AssembledSources []*UpdateMcpServerRequestAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	// The backend service configurations for the route.
	BackendConfig *UpdateMcpServerRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// Specifies the type of source for MCP server creation.
	CreateFromType *string `json:"createFromType,omitempty" xml:"createFromType,omitempty"`
	// The description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The exposed URI path. This parameter is required when the protocol parameter is set to SSE or StreamableHTTP and the type parameter is set to RealMCP.
	//
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// The route match rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// Specifies if MCP observability is enabled. Default value: false.
	//
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// The service protocol. Valid values: HTTP, HTTPS, SSE, and StreamableHTTP.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The MCP server type. Valid values: RealMCP and AssemblyMCP.
	//
	// This parameter is required.
	//
	// example:
	//
	// RealMCP
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerRequest) GetAssembledSources() []*UpdateMcpServerRequestAssembledSources {
	return s.AssembledSources
}

func (s *UpdateMcpServerRequest) GetBackendConfig() *UpdateMcpServerRequestBackendConfig {
	return s.BackendConfig
}

func (s *UpdateMcpServerRequest) GetCreateFromType() *string {
	return s.CreateFromType
}

func (s *UpdateMcpServerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpServerRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *UpdateMcpServerRequest) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *UpdateMcpServerRequest) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *UpdateMcpServerRequest) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *UpdateMcpServerRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateMcpServerRequest) GetType() *string {
	return s.Type
}

func (s *UpdateMcpServerRequest) SetAssembledSources(v []*UpdateMcpServerRequestAssembledSources) *UpdateMcpServerRequest {
	s.AssembledSources = v
	return s
}

func (s *UpdateMcpServerRequest) SetBackendConfig(v *UpdateMcpServerRequestBackendConfig) *UpdateMcpServerRequest {
	s.BackendConfig = v
	return s
}

func (s *UpdateMcpServerRequest) SetCreateFromType(v string) *UpdateMcpServerRequest {
	s.CreateFromType = &v
	return s
}

func (s *UpdateMcpServerRequest) SetDescription(v string) *UpdateMcpServerRequest {
	s.Description = &v
	return s
}

func (s *UpdateMcpServerRequest) SetDomainIds(v []*string) *UpdateMcpServerRequest {
	s.DomainIds = v
	return s
}

func (s *UpdateMcpServerRequest) SetExposedUriPath(v string) *UpdateMcpServerRequest {
	s.ExposedUriPath = &v
	return s
}

func (s *UpdateMcpServerRequest) SetMatch(v *HttpRouteMatch) *UpdateMcpServerRequest {
	s.Match = v
	return s
}

func (s *UpdateMcpServerRequest) SetMcpStatisticsEnable(v bool) *UpdateMcpServerRequest {
	s.McpStatisticsEnable = &v
	return s
}

func (s *UpdateMcpServerRequest) SetProtocol(v string) *UpdateMcpServerRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateMcpServerRequest) SetType(v string) *UpdateMcpServerRequest {
	s.Type = &v
	return s
}

func (s *UpdateMcpServerRequest) Validate() error {
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
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpServerRequestAssembledSources struct {
	// The MCP server ID.
	//
	// example:
	//
	// mcp-afaefaefaf
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The name of the MCP server.
	//
	// example:
	//
	// test-mcp
	McpServerName *string `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	// The MCP tools.
	Tools []*string `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s UpdateMcpServerRequestAssembledSources) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerRequestAssembledSources) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerRequestAssembledSources) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *UpdateMcpServerRequestAssembledSources) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *UpdateMcpServerRequestAssembledSources) GetTools() []*string {
	return s.Tools
}

func (s *UpdateMcpServerRequestAssembledSources) SetMcpServerId(v string) *UpdateMcpServerRequestAssembledSources {
	s.McpServerId = &v
	return s
}

func (s *UpdateMcpServerRequestAssembledSources) SetMcpServerName(v string) *UpdateMcpServerRequestAssembledSources {
	s.McpServerName = &v
	return s
}

func (s *UpdateMcpServerRequestAssembledSources) SetTools(v []*string) *UpdateMcpServerRequestAssembledSources {
	s.Tools = v
	return s
}

func (s *UpdateMcpServerRequestAssembledSources) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpServerRequestBackendConfig struct {
	// The backend service scenario.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The backend services.
	Services []*UpdateMcpServerRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s UpdateMcpServerRequestBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerRequestBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *UpdateMcpServerRequestBackendConfig) GetServices() []*UpdateMcpServerRequestBackendConfigServices {
	return s.Services
}

func (s *UpdateMcpServerRequestBackendConfig) SetScene(v string) *UpdateMcpServerRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *UpdateMcpServerRequestBackendConfig) SetServices(v []*UpdateMcpServerRequestBackendConfigServices) *UpdateMcpServerRequestBackendConfig {
	s.Services = v
	return s
}

func (s *UpdateMcpServerRequestBackendConfig) Validate() error {
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

type UpdateMcpServerRequestBackendConfigServices struct {
	// The service port (omit for dynamic ports).
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- HTTP
	//
	// 	- DUBBO
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-cr6pk4tlhtgm58e***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The traffic weight percentage.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s UpdateMcpServerRequestBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerRequestBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *UpdateMcpServerRequestBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateMcpServerRequestBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateMcpServerRequestBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *UpdateMcpServerRequestBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateMcpServerRequestBackendConfigServices) SetPort(v int32) *UpdateMcpServerRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *UpdateMcpServerRequestBackendConfigServices) SetProtocol(v string) *UpdateMcpServerRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *UpdateMcpServerRequestBackendConfigServices) SetServiceId(v string) *UpdateMcpServerRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateMcpServerRequestBackendConfigServices) SetVersion(v string) *UpdateMcpServerRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *UpdateMcpServerRequestBackendConfigServices) SetWeight(v int32) *UpdateMcpServerRequestBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *UpdateMcpServerRequestBackendConfigServices) Validate() error {
	return dara.Validate(s)
}
