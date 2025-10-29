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
	SetDescription(v string) *CreateMcpServerRequest
	GetDescription() *string
	SetDomainIds(v []*string) *CreateMcpServerRequest
	GetDomainIds() []*string
	SetExposedUriPath(v string) *CreateMcpServerRequest
	GetExposedUriPath() *string
	SetGatewayId(v string) *CreateMcpServerRequest
	GetGatewayId() *string
	SetMatch(v *HttpRouteMatch) *CreateMcpServerRequest
	GetMatch() *HttpRouteMatch
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
	AssembledSources []*CreateMcpServerRequestAssembledSources `json:"assembledSources,omitempty" xml:"assembledSources,omitempty" type:"Repeated"`
	BackendConfig    *CreateMcpServerRequestBackendConfig      `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// example:
	//
	// mcp tool fetch time
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	DomainIds   []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// example:
	//
	// /sse
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qac0
	GatewayId *string         `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	Match     *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fetch-time
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
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

func (s *CreateMcpServerRequest) GetMatch() *HttpRouteMatch {
	return s.Match
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

func (s *CreateMcpServerRequest) SetMatch(v *HttpRouteMatch) *CreateMcpServerRequest {
	s.Match = v
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
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
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
	// example:
	//
	// test-mcp
	McpServerName *string   `json:"mcpServerName,omitempty" xml:"mcpServerName,omitempty"`
	Tools         []*string `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
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
	// example:
	//
	// SingleService
	Scene    *string                                        `json:"scene,omitempty" xml:"scene,omitempty"`
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
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// svc-crbgq0dlhtgr***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// 2.1.6
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
