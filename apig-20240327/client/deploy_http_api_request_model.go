// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHttpApiConfig(v *DeployHttpApiRequestHttpApiConfig) *DeployHttpApiRequest
	GetHttpApiConfig() *DeployHttpApiRequestHttpApiConfig
	SetRestApiConfig(v *DeployHttpApiRequestRestApiConfig) *DeployHttpApiRequest
	GetRestApiConfig() *DeployHttpApiRequestRestApiConfig
	SetRouteId(v string) *DeployHttpApiRequest
	GetRouteId() *string
}

type DeployHttpApiRequest struct {
	// Deprecated
	//
	// httpApiConfig
	HttpApiConfig *DeployHttpApiRequestHttpApiConfig `json:"httpApiConfig,omitempty" xml:"httpApiConfig,omitempty" type:"Struct"`
	// The REST API deployment configuration. This parameter is required when you publish a REST API.
	RestApiConfig *DeployHttpApiRequestRestApiConfig `json:"restApiConfig,omitempty" xml:"restApiConfig,omitempty" type:"Struct"`
	// The route ID. You must specify this parameter when you publish an HTTP API.
	//
	// example:
	//
	// hr-cr82undlhtgrl***
	RouteId *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s DeployHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiRequest) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequest) GetHttpApiConfig() *DeployHttpApiRequestHttpApiConfig {
	return s.HttpApiConfig
}

func (s *DeployHttpApiRequest) GetRestApiConfig() *DeployHttpApiRequestRestApiConfig {
	return s.RestApiConfig
}

func (s *DeployHttpApiRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *DeployHttpApiRequest) SetHttpApiConfig(v *DeployHttpApiRequestHttpApiConfig) *DeployHttpApiRequest {
	s.HttpApiConfig = v
	return s
}

func (s *DeployHttpApiRequest) SetRestApiConfig(v *DeployHttpApiRequestRestApiConfig) *DeployHttpApiRequest {
	s.RestApiConfig = v
	return s
}

func (s *DeployHttpApiRequest) SetRouteId(v string) *DeployHttpApiRequest {
	s.RouteId = &v
	return s
}

func (s *DeployHttpApiRequest) Validate() error {
	if s.HttpApiConfig != nil {
		if err := s.HttpApiConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RestApiConfig != nil {
		if err := s.RestApiConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployHttpApiRequestHttpApiConfig struct {
	// The gateway ID.
	//
	// example:
	//
	// gw-csrhgfmm1hknf0hu6o3g
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// routeIds
	RouteIds []*string `json:"routeIds,omitempty" xml:"routeIds,omitempty" type:"Repeated"`
}

func (s DeployHttpApiRequestHttpApiConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiRequestHttpApiConfig) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestHttpApiConfig) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeployHttpApiRequestHttpApiConfig) GetRouteIds() []*string {
	return s.RouteIds
}

func (s *DeployHttpApiRequestHttpApiConfig) SetGatewayId(v string) *DeployHttpApiRequestHttpApiConfig {
	s.GatewayId = &v
	return s
}

func (s *DeployHttpApiRequestHttpApiConfig) SetRouteIds(v []*string) *DeployHttpApiRequestHttpApiConfig {
	s.RouteIds = v
	return s
}

func (s *DeployHttpApiRequestHttpApiConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHttpApiRequestRestApiConfig struct {
	// The publish description.
	//
	// example:
	//
	// The user service API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The environment configurations.
	Environment *DeployHttpApiRequestRestApiConfigEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	// The gateway ID.
	//
	// example:
	//
	// gw-cvn2u46m1hkun04oll8g
	GatewayId            *string                                                  `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	OperationDeployments []*DeployHttpApiRequestRestApiConfigOperationDeployments `json:"operationDeployments,omitempty" xml:"operationDeployments,omitempty" type:"Repeated"`
	// Deprecated
	//
	// operationIds
	OperationIds []*string `json:"operationIds,omitempty" xml:"operationIds,omitempty" type:"Repeated"`
	// The historical version of the API. If you specify this parameter, the corresponding version of the API is published.
	//
	// example:
	//
	// apr-xxx
	RevisionId *string `json:"revisionId,omitempty" xml:"revisionId,omitempty"`
}

func (s DeployHttpApiRequestRestApiConfig) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfig) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfig) GetDescription() *string {
	return s.Description
}

func (s *DeployHttpApiRequestRestApiConfig) GetEnvironment() *DeployHttpApiRequestRestApiConfigEnvironment {
	return s.Environment
}

func (s *DeployHttpApiRequestRestApiConfig) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DeployHttpApiRequestRestApiConfig) GetOperationDeployments() []*DeployHttpApiRequestRestApiConfigOperationDeployments {
	return s.OperationDeployments
}

func (s *DeployHttpApiRequestRestApiConfig) GetOperationIds() []*string {
	return s.OperationIds
}

func (s *DeployHttpApiRequestRestApiConfig) GetRevisionId() *string {
	return s.RevisionId
}

func (s *DeployHttpApiRequestRestApiConfig) SetDescription(v string) *DeployHttpApiRequestRestApiConfig {
	s.Description = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetEnvironment(v *DeployHttpApiRequestRestApiConfigEnvironment) *DeployHttpApiRequestRestApiConfig {
	s.Environment = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetGatewayId(v string) *DeployHttpApiRequestRestApiConfig {
	s.GatewayId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetOperationDeployments(v []*DeployHttpApiRequestRestApiConfigOperationDeployments) *DeployHttpApiRequestRestApiConfig {
	s.OperationDeployments = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetOperationIds(v []*string) *DeployHttpApiRequestRestApiConfig {
	s.OperationIds = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetRevisionId(v string) *DeployHttpApiRequestRestApiConfig {
	s.RevisionId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	if s.OperationDeployments != nil {
		for _, item := range s.OperationDeployments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeployHttpApiRequestRestApiConfigEnvironment struct {
	// Deprecated
	//
	// The publishing scenario.
	//
	// Valid values:
	//
	// 	- SingleService
	//
	// 	- MultiServiceByRatio
	//
	// 	- MultiServiceByContent
	//
	// 	- Mock
	//
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// The custom domain names.
	CustomDomainIds []*string `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The environment ID.
	//
	// example:
	//
	// env-cpqnr6tlhtgubc***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Deprecated
	//
	// The configurations of existing services. For single-service publishing, only one entry is allowed. For other scenarios, multiple entries are allowed.
	ServiceConfigs []*DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s DeployHttpApiRequestRestApiConfigEnvironment) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfigEnvironment) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) GetBackendScene() *string {
	return s.BackendScene
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) GetCustomDomainIds() []*string {
	return s.CustomDomainIds
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) GetServiceConfigs() []*DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	return s.ServiceConfigs
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetBackendScene(v string) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.BackendScene = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetCustomDomainIds(v []*string) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.CustomDomainIds = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetEnvironmentId(v string) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.EnvironmentId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) SetServiceConfigs(v []*DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) *DeployHttpApiRequestRestApiConfigEnvironment {
	s.ServiceConfigs = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironment) Validate() error {
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs struct {
	// The matching condition configurations related to API publishing.
	//
	// example:
	//
	// {\\"change_order_revision\\":\\"3.657.33_fc-hz-yunqi.1662568293908382_faas-eerouter\\"}
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The service port. If you want to use a dynamic port, do not pass this parameter.
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
	// The version of the microservice.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The weight. Valid values: [1,100]. This parameter is valid only in proportional routing.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Match = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetPort(v int32) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Port = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetProtocol(v string) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetServiceId(v string) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetVersion(v string) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Version = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) SetWeight(v int32) *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs {
	s.Weight = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeployHttpApiRequestRestApiConfigOperationDeployments struct {
	// example:
	//
	// Publish
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// op-d5s57hmm1hks653u9dkg
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s DeployHttpApiRequestRestApiConfigOperationDeployments) String() string {
	return dara.Prettify(s)
}

func (s DeployHttpApiRequestRestApiConfigOperationDeployments) GoString() string {
	return s.String()
}

func (s *DeployHttpApiRequestRestApiConfigOperationDeployments) GetAction() *string {
	return s.Action
}

func (s *DeployHttpApiRequestRestApiConfigOperationDeployments) GetOperationId() *string {
	return s.OperationId
}

func (s *DeployHttpApiRequestRestApiConfigOperationDeployments) SetAction(v string) *DeployHttpApiRequestRestApiConfigOperationDeployments {
	s.Action = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigOperationDeployments) SetOperationId(v string) *DeployHttpApiRequestRestApiConfigOperationDeployments {
	s.OperationId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfigOperationDeployments) Validate() error {
	return dara.Validate(s)
}
