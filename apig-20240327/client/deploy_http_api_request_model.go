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
	HttpApiConfig *DeployHttpApiRequestHttpApiConfig `json:"httpApiConfig,omitempty" xml:"httpApiConfig,omitempty" type:"Struct"`
	// Rest API deployment configuration. Required when deploying an HTTP API as a Rest API.
	RestApiConfig *DeployHttpApiRequestRestApiConfig `json:"restApiConfig,omitempty" xml:"restApiConfig,omitempty" type:"Struct"`
	// Route ID. This must be provided when publishing the route of an HTTP API.
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
	return dara.Validate(s)
}

type DeployHttpApiRequestHttpApiConfig struct {
	GatewayId *string   `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	RouteIds  []*string `json:"routeIds,omitempty" xml:"routeIds,omitempty" type:"Repeated"`
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
	// Publication description.
	//
	// example:
	//
	// 用户服务API发布。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Publication environment configuration.
	Environment  *DeployHttpApiRequestRestApiConfigEnvironment `json:"environment,omitempty" xml:"environment,omitempty" type:"Struct"`
	GatewayId    *string                                       `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	OperationIds []*string                                     `json:"operationIds,omitempty" xml:"operationIds,omitempty" type:"Repeated"`
	// Historical version number. If this field is specified, the publication information will be based on the historical version information.
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

func (s *DeployHttpApiRequestRestApiConfig) SetOperationIds(v []*string) *DeployHttpApiRequestRestApiConfig {
	s.OperationIds = v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) SetRevisionId(v string) *DeployHttpApiRequestRestApiConfig {
	s.RevisionId = &v
	return s
}

func (s *DeployHttpApiRequestRestApiConfig) Validate() error {
	return dara.Validate(s)
}

type DeployHttpApiRequestRestApiConfigEnvironment struct {
	// API publication scenario.
	//
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// List of user domains.
	CustomDomainIds []*string `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	// Environment ID.
	//
	// example:
	//
	// env-cpqnr6tlhtgubc***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Existing service configurations. Only one entry is allowed in a single-service scenario, while multiple entries are allowed in scenarios such as by ratio or by content.
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
	return dara.Validate(s)
}

type DeployHttpApiRequestRestApiConfigEnvironmentServiceConfigs struct {
	// Configuration of matching conditions related to API deployment.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// Service port, do not provide for dynamic ports.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Service protocol:
	//
	// - HTTP.
	//
	// - HTTPS.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Service ID.
	//
	// example:
	//
	// svc-cr6pk4tlhtgm58e***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// Service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// Weight, range [1,100], valid only in the by-ratio scenario.
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
	return dara.Validate(s)
}
