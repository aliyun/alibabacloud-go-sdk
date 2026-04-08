// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendConfig(v *UpdateHttpApiRouteRequestBackendConfig) *UpdateHttpApiRouteRequest
	GetBackendConfig() *UpdateHttpApiRouteRequestBackendConfig
	SetDescription(v string) *UpdateHttpApiRouteRequest
	GetDescription() *string
	SetDomainIds(v []*string) *UpdateHttpApiRouteRequest
	GetDomainIds() []*string
	SetEnvironmentId(v string) *UpdateHttpApiRouteRequest
	GetEnvironmentId() *string
	SetMatch(v *HttpRouteMatch) *UpdateHttpApiRouteRequest
	GetMatch() *HttpRouteMatch
	SetMcpRouteConfig(v *UpdateHttpApiRouteRequestMcpRouteConfig) *UpdateHttpApiRouteRequest
	GetMcpRouteConfig() *UpdateHttpApiRouteRequestMcpRouteConfig
	SetPolicyConfigs(v []*HttpApiPolicyConfigs) *UpdateHttpApiRouteRequest
	GetPolicyConfigs() []*HttpApiPolicyConfigs
}

type UpdateHttpApiRouteRequest struct {
	// The backend service configurations for the route.
	BackendConfig *UpdateHttpApiRouteRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// The route description.
	//
	// example:
	//
	// test route
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-cquqsollhtgid***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The route match rule.
	Match          *HttpRouteMatch                          `json:"match,omitempty" xml:"match,omitempty"`
	McpRouteConfig *UpdateHttpApiRouteRequestMcpRouteConfig `json:"mcpRouteConfig,omitempty" xml:"mcpRouteConfig,omitempty" type:"Struct"`
	PolicyConfigs  []*HttpApiPolicyConfigs                  `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequest) GetBackendConfig() *UpdateHttpApiRouteRequestBackendConfig {
	return s.BackendConfig
}

func (s *UpdateHttpApiRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateHttpApiRouteRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *UpdateHttpApiRouteRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateHttpApiRouteRequest) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *UpdateHttpApiRouteRequest) GetMcpRouteConfig() *UpdateHttpApiRouteRequestMcpRouteConfig {
	return s.McpRouteConfig
}

func (s *UpdateHttpApiRouteRequest) GetPolicyConfigs() []*HttpApiPolicyConfigs {
	return s.PolicyConfigs
}

func (s *UpdateHttpApiRouteRequest) SetBackendConfig(v *UpdateHttpApiRouteRequestBackendConfig) *UpdateHttpApiRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDescription(v string) *UpdateHttpApiRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDomainIds(v []*string) *UpdateHttpApiRouteRequest {
	s.DomainIds = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetEnvironmentId(v string) *UpdateHttpApiRouteRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetMatch(v *HttpRouteMatch) *UpdateHttpApiRouteRequest {
	s.Match = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetMcpRouteConfig(v *UpdateHttpApiRouteRequestMcpRouteConfig) *UpdateHttpApiRouteRequest {
	s.McpRouteConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetPolicyConfigs(v []*HttpApiPolicyConfigs) *UpdateHttpApiRouteRequest {
	s.PolicyConfigs = v
	return s
}

func (s *UpdateHttpApiRouteRequest) Validate() error {
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
	if s.McpRouteConfig != nil {
		if err := s.McpRouteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PolicyConfigs != nil {
		for _, item := range s.PolicyConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestBackendConfig struct {
	// The backend service scenario. Valid values:
	//
	// 	- SingleService
	//
	// 	- MultiServiceByRatio
	//
	// 	- Redirect
	//
	// 	- Mock
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The list of backend services.
	Services []*UpdateHttpApiRouteRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequestBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *UpdateHttpApiRouteRequestBackendConfig) GetServices() []*UpdateHttpApiRouteRequestBackendConfigServices {
	return s.Services
}

func (s *UpdateHttpApiRouteRequestBackendConfig) SetScene(v string) *UpdateHttpApiRouteRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfig) SetServices(v []*UpdateHttpApiRouteRequestBackendConfigServices) *UpdateHttpApiRouteRequestBackendConfig {
	s.Services = v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfig) Validate() error {
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

type UpdateHttpApiRouteRequestBackendConfigServices struct {
	// The service port (omit for dynamic ports).
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
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

func (s UpdateHttpApiRouteRequestBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetPort(v int32) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetProtocol(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetServiceId(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetVersion(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetWeight(v int32) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestMcpRouteConfig struct {
	ExposedUriPath      *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	McpStatisticsEnable *bool   `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	Protocol            *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateHttpApiRouteRequestMcpRouteConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestMcpRouteConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) SetExposedUriPath(v string) *UpdateHttpApiRouteRequestMcpRouteConfig {
	s.ExposedUriPath = &v
	return s
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) SetMcpStatisticsEnable(v bool) *UpdateHttpApiRouteRequestMcpRouteConfig {
	s.McpStatisticsEnable = &v
	return s
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) SetProtocol(v string) *UpdateHttpApiRouteRequestMcpRouteConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) Validate() error {
	return dara.Validate(s)
}
