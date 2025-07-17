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
	SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRouteRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *UpdateHttpApiRouteRequest
	GetDescription() *string
	SetDomainIds(v []*string) *UpdateHttpApiRouteRequest
	GetDomainIds() []*string
	SetEnvironmentId(v string) *UpdateHttpApiRouteRequest
	GetEnvironmentId() *string
	SetMatch(v *HttpRouteMatch) *UpdateHttpApiRouteRequest
	GetMatch() *HttpRouteMatch
}

type UpdateHttpApiRouteRequest struct {
	// The backend service configurations of the route.
	BackendConfig *UpdateHttpApiRouteRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	DeployConfigs []*HttpApiDeployConfig                  `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The route description.
	//
	// example:
	//
	// test route
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-cquqsollhtgid***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The rules for matching the route.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
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

func (s *UpdateHttpApiRouteRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
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

func (s *UpdateHttpApiRouteRequest) SetBackendConfig(v *UpdateHttpApiRouteRequestBackendConfig) *UpdateHttpApiRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRouteRequest {
	s.DeployConfigs = v
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

func (s *UpdateHttpApiRouteRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestBackendConfig struct {
	// The backend service scenario.
	//
	// Valid values:
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
	// The backend services.
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
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestBackendConfigServices struct {
	// The service port. If you want to use a dynamic port, do not pass this parameter.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. Valid values:
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
	// The percentage value of traffic.
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
