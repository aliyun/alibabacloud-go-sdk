// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiApiInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProtocols(v []*string) *HttpApiApiInfo
	GetAgentProtocols() []*string
	SetAiProtocols(v []*string) *HttpApiApiInfo
	GetAiProtocols() []*string
	SetAuthConfig(v *AuthConfig) *HttpApiApiInfo
	GetAuthConfig() *AuthConfig
	SetBasePath(v string) *HttpApiApiInfo
	GetBasePath() *string
	SetDeployCntMap(v map[string]*HttpApiApiInfoDeployCntMapValue) *HttpApiApiInfo
	GetDeployCntMap() map[string]*HttpApiApiInfoDeployCntMapValue
	SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiApiInfo
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *HttpApiApiInfo
	GetDescription() *string
	SetEnabelAuth(v bool) *HttpApiApiInfo
	GetEnabelAuth() *bool
	SetEnvironments(v []*HttpApiApiInfoEnvironments) *HttpApiApiInfo
	GetEnvironments() []*HttpApiApiInfoEnvironments
	SetGatewayId(v string) *HttpApiApiInfo
	GetGatewayId() *string
	SetHttpApiId(v string) *HttpApiApiInfo
	GetHttpApiId() *string
	SetIngressInfo(v *HttpApiApiInfoIngressInfo) *HttpApiApiInfo
	GetIngressInfo() *HttpApiApiInfoIngressInfo
	SetModelCategory(v string) *HttpApiApiInfo
	GetModelCategory() *string
	SetName(v string) *HttpApiApiInfo
	GetName() *string
	SetProtocols(v []*string) *HttpApiApiInfo
	GetProtocols() []*string
	SetResourceGroupId(v string) *HttpApiApiInfo
	GetResourceGroupId() *string
	SetType(v string) *HttpApiApiInfo
	GetType() *string
	SetVersionInfo(v *HttpApiVersionInfo) *HttpApiApiInfo
	GetVersionInfo() *HttpApiVersionInfo
}

type HttpApiApiInfo struct {
	// Agent protocol list
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The AI protocols.
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The authentication configurations.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The base path of the API.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The API publishing status.
	DeployCntMap map[string]*HttpApiApiInfoDeployCntMapValue `json:"deployCntMap,omitempty" xml:"deployCntMap,omitempty"`
	// The API deployment configurations.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The API description.
	//
	// example:
	//
	// for test only
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable authentication.
	EnabelAuth *bool `json:"enabelAuth,omitempty" xml:"enabelAuth,omitempty"`
	// The API environment information.
	Environments []*HttpApiApiInfoEnvironments `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// gw-cpv4sqdl****
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The HTTP API ID.
	//
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// The information about the HTTP Ingress API.
	IngressInfo *HttpApiApiInfoIngressInfo `json:"ingressInfo,omitempty" xml:"ingressInfo,omitempty" type:"Struct"`
	// Model category
	//
	// example:
	//
	// Text
	ModelCategory *string `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	// The API name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocols.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The API type.
	//
	// Valid values:
	//
	// 	- HttpIngress
	//
	// 	- Rest
	//
	// 	- Websocket
	//
	// 	- AI
	//
	// 	- Http
	//
	// example:
	//
	// Rest
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The API versioning information.
	VersionInfo *HttpApiVersionInfo `json:"versionInfo,omitempty" xml:"versionInfo,omitempty"`
}

func (s HttpApiApiInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfo) GetAgentProtocols() []*string {
	return s.AgentProtocols
}

func (s *HttpApiApiInfo) GetAiProtocols() []*string {
	return s.AiProtocols
}

func (s *HttpApiApiInfo) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *HttpApiApiInfo) GetBasePath() *string {
	return s.BasePath
}

func (s *HttpApiApiInfo) GetDeployCntMap() map[string]*HttpApiApiInfoDeployCntMapValue {
	return s.DeployCntMap
}

func (s *HttpApiApiInfo) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *HttpApiApiInfo) GetDescription() *string {
	return s.Description
}

func (s *HttpApiApiInfo) GetEnabelAuth() *bool {
	return s.EnabelAuth
}

func (s *HttpApiApiInfo) GetEnvironments() []*HttpApiApiInfoEnvironments {
	return s.Environments
}

func (s *HttpApiApiInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiApiInfo) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *HttpApiApiInfo) GetIngressInfo() *HttpApiApiInfoIngressInfo {
	return s.IngressInfo
}

func (s *HttpApiApiInfo) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *HttpApiApiInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiApiInfo) GetProtocols() []*string {
	return s.Protocols
}

func (s *HttpApiApiInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *HttpApiApiInfo) GetType() *string {
	return s.Type
}

func (s *HttpApiApiInfo) GetVersionInfo() *HttpApiVersionInfo {
	return s.VersionInfo
}

func (s *HttpApiApiInfo) SetAgentProtocols(v []*string) *HttpApiApiInfo {
	s.AgentProtocols = v
	return s
}

func (s *HttpApiApiInfo) SetAiProtocols(v []*string) *HttpApiApiInfo {
	s.AiProtocols = v
	return s
}

func (s *HttpApiApiInfo) SetAuthConfig(v *AuthConfig) *HttpApiApiInfo {
	s.AuthConfig = v
	return s
}

func (s *HttpApiApiInfo) SetBasePath(v string) *HttpApiApiInfo {
	s.BasePath = &v
	return s
}

func (s *HttpApiApiInfo) SetDeployCntMap(v map[string]*HttpApiApiInfoDeployCntMapValue) *HttpApiApiInfo {
	s.DeployCntMap = v
	return s
}

func (s *HttpApiApiInfo) SetDeployConfigs(v []*HttpApiDeployConfig) *HttpApiApiInfo {
	s.DeployConfigs = v
	return s
}

func (s *HttpApiApiInfo) SetDescription(v string) *HttpApiApiInfo {
	s.Description = &v
	return s
}

func (s *HttpApiApiInfo) SetEnabelAuth(v bool) *HttpApiApiInfo {
	s.EnabelAuth = &v
	return s
}

func (s *HttpApiApiInfo) SetEnvironments(v []*HttpApiApiInfoEnvironments) *HttpApiApiInfo {
	s.Environments = v
	return s
}

func (s *HttpApiApiInfo) SetGatewayId(v string) *HttpApiApiInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpApiApiInfo) SetHttpApiId(v string) *HttpApiApiInfo {
	s.HttpApiId = &v
	return s
}

func (s *HttpApiApiInfo) SetIngressInfo(v *HttpApiApiInfoIngressInfo) *HttpApiApiInfo {
	s.IngressInfo = v
	return s
}

func (s *HttpApiApiInfo) SetModelCategory(v string) *HttpApiApiInfo {
	s.ModelCategory = &v
	return s
}

func (s *HttpApiApiInfo) SetName(v string) *HttpApiApiInfo {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfo) SetProtocols(v []*string) *HttpApiApiInfo {
	s.Protocols = v
	return s
}

func (s *HttpApiApiInfo) SetResourceGroupId(v string) *HttpApiApiInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *HttpApiApiInfo) SetType(v string) *HttpApiApiInfo {
	s.Type = &v
	return s
}

func (s *HttpApiApiInfo) SetVersionInfo(v *HttpApiVersionInfo) *HttpApiApiInfo {
	s.VersionInfo = v
	return s
}

func (s *HttpApiApiInfo) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeployConfigs != nil {
		for _, item := range s.DeployConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Environments != nil {
		for _, item := range s.Environments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IngressInfo != nil {
		if err := s.IngressInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VersionInfo != nil {
		if err := s.VersionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiApiInfoEnvironments struct {
	// The environment alias.
	//
	// example:
	//
	// test
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
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
	// 	- MultiServiceByTag
	//
	// 	- Mock
	//
	// example:
	//
	// SingleService
	BackendScene *string `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	// The type of the backend service.
	//
	// Valid values:
	//
	// 	- DNS: a DNS domain name
	//
	// 	- Service: an existing service
	//
	// 	- VIP: a fixed IP address
	//
	// 	- CloudProduct: a cloud service
	//
	// example:
	//
	// Service
	BackendType *string `json:"backendType,omitempty" xml:"backendType,omitempty"`
	// The custom domain names.
	CustomDomains []*HttpApiDomainInfo `json:"customDomains,omitempty" xml:"customDomains,omitempty" type:"Repeated"`
	// The publishing status of the API in the current environment.
	//
	// example:
	//
	// Deployed
	DeployStatus *string `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The information about the gateway instance to which the environment belongs.
	GatewayInfo *HttpApiApiInfoEnvironmentsGatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty" type:"Struct"`
	// The environment name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configurations of an existing service.
	ServiceConfigs []*HttpApiApiInfoEnvironmentsServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// The default domain names of the environment.
	SubDomains []*HttpApiApiInfoEnvironmentsSubDomains `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
}

func (s HttpApiApiInfoEnvironments) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoEnvironments) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironments) GetAlias() *string {
	return s.Alias
}

func (s *HttpApiApiInfoEnvironments) GetBackendScene() *string {
	return s.BackendScene
}

func (s *HttpApiApiInfoEnvironments) GetBackendType() *string {
	return s.BackendType
}

func (s *HttpApiApiInfoEnvironments) GetCustomDomains() []*HttpApiDomainInfo {
	return s.CustomDomains
}

func (s *HttpApiApiInfoEnvironments) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *HttpApiApiInfoEnvironments) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpApiApiInfoEnvironments) GetGatewayInfo() *HttpApiApiInfoEnvironmentsGatewayInfo {
	return s.GatewayInfo
}

func (s *HttpApiApiInfoEnvironments) GetName() *string {
	return s.Name
}

func (s *HttpApiApiInfoEnvironments) GetServiceConfigs() []*HttpApiApiInfoEnvironmentsServiceConfigs {
	return s.ServiceConfigs
}

func (s *HttpApiApiInfoEnvironments) GetSubDomains() []*HttpApiApiInfoEnvironmentsSubDomains {
	return s.SubDomains
}

func (s *HttpApiApiInfoEnvironments) SetAlias(v string) *HttpApiApiInfoEnvironments {
	s.Alias = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetBackendScene(v string) *HttpApiApiInfoEnvironments {
	s.BackendScene = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetBackendType(v string) *HttpApiApiInfoEnvironments {
	s.BackendType = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetCustomDomains(v []*HttpApiDomainInfo) *HttpApiApiInfoEnvironments {
	s.CustomDomains = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetDeployStatus(v string) *HttpApiApiInfoEnvironments {
	s.DeployStatus = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetEnvironmentId(v string) *HttpApiApiInfoEnvironments {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetGatewayInfo(v *HttpApiApiInfoEnvironmentsGatewayInfo) *HttpApiApiInfoEnvironments {
	s.GatewayInfo = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetName(v string) *HttpApiApiInfoEnvironments {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetServiceConfigs(v []*HttpApiApiInfoEnvironmentsServiceConfigs) *HttpApiApiInfoEnvironments {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiApiInfoEnvironments) SetSubDomains(v []*HttpApiApiInfoEnvironmentsSubDomains) *HttpApiApiInfoEnvironments {
	s.SubDomains = v
	return s
}

func (s *HttpApiApiInfoEnvironments) Validate() error {
	if s.CustomDomains != nil {
		for _, item := range s.CustomDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubDomains != nil {
		for _, item := range s.SubDomains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiApiInfoEnvironmentsGatewayInfo struct {
	// The instance ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsGatewayInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsGatewayInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) SetGatewayId(v string) *HttpApiApiInfoEnvironmentsGatewayInfo {
	s.GatewayId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) SetName(v string) *HttpApiApiInfoEnvironmentsGatewayInfo {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsGatewayInfo) Validate() error {
	return dara.Validate(s)
}

type HttpApiApiInfoEnvironmentsServiceConfigs struct {
	// The service ID.
	//
	// example:
	//
	// gs-xxx
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// The matching conditions.
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The service name.
	//
	// example:
	//
	// demo-service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service port.
	//
	// example:
	//
	// 8080
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The version of the microservice.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetPort() *string {
	return s.Port
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetGatewayServiceId(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetName(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetPort(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetProtocol(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetServiceId(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetVersion(v string) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Version = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) SetWeight(v int32) *HttpApiApiInfoEnvironmentsServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiApiInfoEnvironmentsSubDomains struct {
	// The domain name ID.
	//
	// example:
	//
	// d-xxx
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network type.
	//
	// example:
	//
	// Internet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The communication protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpApiApiInfoEnvironmentsSubDomains) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoEnvironmentsSubDomains) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) GetName() *string {
	return s.Name
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) GetNetworkType() *string {
	return s.NetworkType
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetDomainId(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.DomainId = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetName(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.Name = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetNetworkType(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.NetworkType = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) SetProtocol(v string) *HttpApiApiInfoEnvironmentsSubDomains {
	s.Protocol = &v
	return s
}

func (s *HttpApiApiInfoEnvironmentsSubDomains) Validate() error {
	return dara.Validate(s)
}

type HttpApiApiInfoIngressInfo struct {
	// The environment information.
	EnvironmentInfo *HttpApiApiInfoIngressInfoEnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty" type:"Struct"`
	// The Ingress Class for listening.
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// The information about the Kubernetes cluster.
	K8sClusterInfo *HttpApiApiInfoIngressInfoK8sClusterInfo `json:"k8sClusterInfo,omitempty" xml:"k8sClusterInfo,omitempty" type:"Struct"`
	// Specifies whether to update the address in Ingress Status.
	//
	// example:
	//
	// true
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src-xxx
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The namespace for listening.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s HttpApiApiInfoIngressInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoIngressInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoIngressInfo) GetEnvironmentInfo() *HttpApiApiInfoIngressInfoEnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *HttpApiApiInfoIngressInfo) GetIngressClass() *string {
	return s.IngressClass
}

func (s *HttpApiApiInfoIngressInfo) GetK8sClusterInfo() *HttpApiApiInfoIngressInfoK8sClusterInfo {
	return s.K8sClusterInfo
}

func (s *HttpApiApiInfoIngressInfo) GetOverrideIngressIp() *bool {
	return s.OverrideIngressIp
}

func (s *HttpApiApiInfoIngressInfo) GetSourceId() *string {
	return s.SourceId
}

func (s *HttpApiApiInfoIngressInfo) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *HttpApiApiInfoIngressInfo) SetEnvironmentInfo(v *HttpApiApiInfoIngressInfoEnvironmentInfo) *HttpApiApiInfoIngressInfo {
	s.EnvironmentInfo = v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetIngressClass(v string) *HttpApiApiInfoIngressInfo {
	s.IngressClass = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetK8sClusterInfo(v *HttpApiApiInfoIngressInfoK8sClusterInfo) *HttpApiApiInfoIngressInfo {
	s.K8sClusterInfo = v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetOverrideIngressIp(v bool) *HttpApiApiInfoIngressInfo {
	s.OverrideIngressIp = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetSourceId(v string) *HttpApiApiInfoIngressInfo {
	s.SourceId = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) SetWatchNamespace(v string) *HttpApiApiInfoIngressInfo {
	s.WatchNamespace = &v
	return s
}

func (s *HttpApiApiInfoIngressInfo) Validate() error {
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.K8sClusterInfo != nil {
		if err := s.K8sClusterInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiApiInfoIngressInfoEnvironmentInfo struct {
	// The environment ID.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s HttpApiApiInfoIngressInfoEnvironmentInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoIngressInfoEnvironmentInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoIngressInfoEnvironmentInfo) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpApiApiInfoIngressInfoEnvironmentInfo) SetEnvironmentId(v string) *HttpApiApiInfoIngressInfoEnvironmentInfo {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiApiInfoIngressInfoEnvironmentInfo) Validate() error {
	return dara.Validate(s)
}

type HttpApiApiInfoIngressInfoK8sClusterInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// ca435c77cba1547cca9311957bcxxxxxx
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s HttpApiApiInfoIngressInfoK8sClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiApiInfoIngressInfoK8sClusterInfo) GoString() string {
	return s.String()
}

func (s *HttpApiApiInfoIngressInfoK8sClusterInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *HttpApiApiInfoIngressInfoK8sClusterInfo) SetClusterId(v string) *HttpApiApiInfoIngressInfoK8sClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *HttpApiApiInfoIngressInfoK8sClusterInfo) Validate() error {
	return dara.Validate(s)
}
