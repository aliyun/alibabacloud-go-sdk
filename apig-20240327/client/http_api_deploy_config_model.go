// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiDeployConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDeploy(v bool) *HttpApiDeployConfig
	GetAutoDeploy() *bool
	SetBackendScene(v string) *HttpApiDeployConfig
	GetBackendScene() *string
	SetBuiltinRouteNames(v []*string) *HttpApiDeployConfig
	GetBuiltinRouteNames() []*string
	SetCustomDomainIds(v []*string) *HttpApiDeployConfig
	GetCustomDomainIds() []*string
	SetCustomDomainInfos(v []*HttpApiDeployConfigCustomDomainInfos) *HttpApiDeployConfig
	GetCustomDomainInfos() []*HttpApiDeployConfigCustomDomainInfos
	SetEnvironmentId(v string) *HttpApiDeployConfig
	GetEnvironmentId() *string
	SetGatewayId(v string) *HttpApiDeployConfig
	GetGatewayId() *string
	SetGatewayInfo(v *GatewayInfo) *HttpApiDeployConfig
	GetGatewayInfo() *GatewayInfo
	SetGatewayType(v string) *HttpApiDeployConfig
	GetGatewayType() *string
	SetMock(v *HttpApiMockContract) *HttpApiDeployConfig
	GetMock() *HttpApiMockContract
	SetPolicyConfigs(v []*HttpApiPolicyConfigs) *HttpApiDeployConfig
	GetPolicyConfigs() []*HttpApiPolicyConfigs
	SetRouteBackend(v *Backend) *HttpApiDeployConfig
	GetRouteBackend() *Backend
	SetServiceConfigs(v []*HttpApiDeployConfigServiceConfigs) *HttpApiDeployConfig
	GetServiceConfigs() []*HttpApiDeployConfigServiceConfigs
	SetSubDomains(v []*HttpApiDeployConfigSubDomains) *HttpApiDeployConfig
	GetSubDomains() []*HttpApiDeployConfigSubDomains
}

type HttpApiDeployConfig struct {
	// Specifies whether to enable automatic deployment.
	//
	// example:
	//
	// true
	AutoDeploy *bool `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	// The publishing scenario.
	//
	// example:
	//
	// SingleService
	BackendScene      *string   `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	BuiltinRouteNames []*string `json:"builtinRouteNames,omitempty" xml:"builtinRouteNames,omitempty" type:"Repeated"`
	// The IDs of the custom domain names.
	CustomDomainIds []*string `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	// The information about the custom domain names.
	CustomDomainInfos []*HttpApiDeployConfigCustomDomainInfos `json:"customDomainInfos,omitempty" xml:"customDomainInfos,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gw-xx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The instance information.
	//
	// if can be null:
	// true
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// 网关类型
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The Mock settings.
	//
	// if can be null:
	// true
	Mock *HttpApiMockContract `json:"mock,omitempty" xml:"mock,omitempty"`
	// The policy configurations.
	PolicyConfigs []*HttpApiPolicyConfigs `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
	// routeBackend
	//
	// if can be null:
	// true
	RouteBackend *Backend `json:"routeBackend,omitempty" xml:"routeBackend,omitempty"`
	// The service configurations.
	ServiceConfigs []*HttpApiDeployConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	// The information about the sub-domain names.
	SubDomains []*HttpApiDeployConfigSubDomains `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
}

func (s HttpApiDeployConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfig) GetAutoDeploy() *bool {
	return s.AutoDeploy
}

func (s *HttpApiDeployConfig) GetBackendScene() *string {
	return s.BackendScene
}

func (s *HttpApiDeployConfig) GetBuiltinRouteNames() []*string {
	return s.BuiltinRouteNames
}

func (s *HttpApiDeployConfig) GetCustomDomainIds() []*string {
	return s.CustomDomainIds
}

func (s *HttpApiDeployConfig) GetCustomDomainInfos() []*HttpApiDeployConfigCustomDomainInfos {
	return s.CustomDomainInfos
}

func (s *HttpApiDeployConfig) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *HttpApiDeployConfig) GetGatewayId() *string {
	return s.GatewayId
}

func (s *HttpApiDeployConfig) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *HttpApiDeployConfig) GetGatewayType() *string {
	return s.GatewayType
}

func (s *HttpApiDeployConfig) GetMock() *HttpApiMockContract {
	return s.Mock
}

func (s *HttpApiDeployConfig) GetPolicyConfigs() []*HttpApiPolicyConfigs {
	return s.PolicyConfigs
}

func (s *HttpApiDeployConfig) GetRouteBackend() *Backend {
	return s.RouteBackend
}

func (s *HttpApiDeployConfig) GetServiceConfigs() []*HttpApiDeployConfigServiceConfigs {
	return s.ServiceConfigs
}

func (s *HttpApiDeployConfig) GetSubDomains() []*HttpApiDeployConfigSubDomains {
	return s.SubDomains
}

func (s *HttpApiDeployConfig) SetAutoDeploy(v bool) *HttpApiDeployConfig {
	s.AutoDeploy = &v
	return s
}

func (s *HttpApiDeployConfig) SetBackendScene(v string) *HttpApiDeployConfig {
	s.BackendScene = &v
	return s
}

func (s *HttpApiDeployConfig) SetBuiltinRouteNames(v []*string) *HttpApiDeployConfig {
	s.BuiltinRouteNames = v
	return s
}

func (s *HttpApiDeployConfig) SetCustomDomainIds(v []*string) *HttpApiDeployConfig {
	s.CustomDomainIds = v
	return s
}

func (s *HttpApiDeployConfig) SetCustomDomainInfos(v []*HttpApiDeployConfigCustomDomainInfos) *HttpApiDeployConfig {
	s.CustomDomainInfos = v
	return s
}

func (s *HttpApiDeployConfig) SetEnvironmentId(v string) *HttpApiDeployConfig {
	s.EnvironmentId = &v
	return s
}

func (s *HttpApiDeployConfig) SetGatewayId(v string) *HttpApiDeployConfig {
	s.GatewayId = &v
	return s
}

func (s *HttpApiDeployConfig) SetGatewayInfo(v *GatewayInfo) *HttpApiDeployConfig {
	s.GatewayInfo = v
	return s
}

func (s *HttpApiDeployConfig) SetGatewayType(v string) *HttpApiDeployConfig {
	s.GatewayType = &v
	return s
}

func (s *HttpApiDeployConfig) SetMock(v *HttpApiMockContract) *HttpApiDeployConfig {
	s.Mock = v
	return s
}

func (s *HttpApiDeployConfig) SetPolicyConfigs(v []*HttpApiPolicyConfigs) *HttpApiDeployConfig {
	s.PolicyConfigs = v
	return s
}

func (s *HttpApiDeployConfig) SetRouteBackend(v *Backend) *HttpApiDeployConfig {
	s.RouteBackend = v
	return s
}

func (s *HttpApiDeployConfig) SetServiceConfigs(v []*HttpApiDeployConfigServiceConfigs) *HttpApiDeployConfig {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiDeployConfig) SetSubDomains(v []*HttpApiDeployConfigSubDomains) *HttpApiDeployConfig {
	s.SubDomains = v
	return s
}

func (s *HttpApiDeployConfig) Validate() error {
	if s.CustomDomainInfos != nil {
		for _, item := range s.CustomDomainInfos {
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
	if s.Mock != nil {
		if err := s.Mock.Validate(); err != nil {
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
	if s.RouteBackend != nil {
		if err := s.RouteBackend.Validate(); err != nil {
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

type HttpApiDeployConfigCustomDomainInfos struct {
	// The domain name ID.
	//
	// example:
	//
	// d-cshee6dlhtgkf4muio3g
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// hello-server.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpApiDeployConfigCustomDomainInfos) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigCustomDomainInfos) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigCustomDomainInfos) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpApiDeployConfigCustomDomainInfos) GetName() *string {
	return s.Name
}

func (s *HttpApiDeployConfigCustomDomainInfos) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiDeployConfigCustomDomainInfos) SetDomainId(v string) *HttpApiDeployConfigCustomDomainInfos {
	s.DomainId = &v
	return s
}

func (s *HttpApiDeployConfigCustomDomainInfos) SetName(v string) *HttpApiDeployConfigCustomDomainInfos {
	s.Name = &v
	return s
}

func (s *HttpApiDeployConfigCustomDomainInfos) SetProtocol(v string) *HttpApiDeployConfigCustomDomainInfos {
	s.Protocol = &v
	return s
}

func (s *HttpApiDeployConfigCustomDomainInfos) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigServiceConfigs struct {
	// Legacy gateway service ID for backward compatibility
	//
	// example:
	//
	// gw-svc-abc123
	GatewayServiceId *string `json:"gatewayServiceId,omitempty" xml:"gatewayServiceId,omitempty"`
	// Intent classification code
	//
	// example:
	//
	// INQUIRY
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	// Match conditions
	Match *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The model name matching rule.
	//
	// example:
	//
	// qwen-*
	ModelNamePattern *string `json:"modelNamePattern,omitempty" xml:"modelNamePattern,omitempty"`
	// Multi-service routing strategy type
	//
	// example:
	//
	// ByWeight
	MultiServiceRouteStrategy *string `json:"multiServiceRouteStrategy,omitempty" xml:"multiServiceRouteStrategy,omitempty"`
	// Service display name
	//
	// example:
	//
	// Qwen-Max-Service
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Observability metrics-based routing config
	//
	// if can be null:
	// true
	ObservabilityRouteConfig *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig `json:"observabilityRouteConfig,omitempty" xml:"observabilityRouteConfig,omitempty" type:"Struct"`
	// Service port number
	//
	// example:
	//
	// 80
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// Service protocol
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
	// Service version tag for tag-based routing scenarios
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The service weight.
	//
	// example:
	//
	// 100
	Weight *int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HttpApiDeployConfigServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigServiceConfigs) GetGatewayServiceId() *string {
	return s.GatewayServiceId
}

func (s *HttpApiDeployConfigServiceConfigs) GetIntentCode() *string {
	return s.IntentCode
}

func (s *HttpApiDeployConfigServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *HttpApiDeployConfigServiceConfigs) GetModelName() *string {
	return s.ModelName
}

func (s *HttpApiDeployConfigServiceConfigs) GetModelNamePattern() *string {
	return s.ModelNamePattern
}

func (s *HttpApiDeployConfigServiceConfigs) GetMultiServiceRouteStrategy() *string {
	return s.MultiServiceRouteStrategy
}

func (s *HttpApiDeployConfigServiceConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiDeployConfigServiceConfigs) GetObservabilityRouteConfig() *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig {
	return s.ObservabilityRouteConfig
}

func (s *HttpApiDeployConfigServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiDeployConfigServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiDeployConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiDeployConfigServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *HttpApiDeployConfigServiceConfigs) GetWeight() *int64 {
	return s.Weight
}

func (s *HttpApiDeployConfigServiceConfigs) SetGatewayServiceId(v string) *HttpApiDeployConfigServiceConfigs {
	s.GatewayServiceId = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetIntentCode(v string) *HttpApiDeployConfigServiceConfigs {
	s.IntentCode = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *HttpApiDeployConfigServiceConfigs {
	s.Match = v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetModelName(v string) *HttpApiDeployConfigServiceConfigs {
	s.ModelName = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetModelNamePattern(v string) *HttpApiDeployConfigServiceConfigs {
	s.ModelNamePattern = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetMultiServiceRouteStrategy(v string) *HttpApiDeployConfigServiceConfigs {
	s.MultiServiceRouteStrategy = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetName(v string) *HttpApiDeployConfigServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetObservabilityRouteConfig(v *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) *HttpApiDeployConfigServiceConfigs {
	s.ObservabilityRouteConfig = v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetPort(v int32) *HttpApiDeployConfigServiceConfigs {
	s.Port = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetProtocol(v string) *HttpApiDeployConfigServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetServiceId(v string) *HttpApiDeployConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetVersion(v string) *HttpApiDeployConfigServiceConfigs {
	s.Version = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetWeight(v int64) *HttpApiDeployConfigServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.ObservabilityRouteConfig != nil {
		if err := s.ObservabilityRouteConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiDeployConfigServiceConfigsObservabilityRouteConfig struct {
	// Routing mode
	//
	// example:
	//
	// LeastBusy
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// Queue size
	//
	// example:
	//
	// 100
	QueueSize *int32 `json:"queueSize,omitempty" xml:"queueSize,omitempty"`
	// Max traffic ratio per single service
	//
	// example:
	//
	// 0.8
	RateLimit *float32 `json:"rateLimit,omitempty" xml:"rateLimit,omitempty"`
}

func (s HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) GetMode() *string {
	return s.Mode
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) GetQueueSize() *int32 {
	return s.QueueSize
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) GetRateLimit() *float32 {
	return s.RateLimit
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) SetMode(v string) *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig {
	s.Mode = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) SetQueueSize(v int32) *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig {
	s.QueueSize = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) SetRateLimit(v float32) *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig {
	s.RateLimit = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigsObservabilityRouteConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigSubDomains struct {
	// The domain name ID.
	//
	// example:
	//
	// d-csmn42um1hksudfk9eng
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network type.
	//
	// example:
	//
	// Intranet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpApiDeployConfigSubDomains) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigSubDomains) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigSubDomains) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpApiDeployConfigSubDomains) GetName() *string {
	return s.Name
}

func (s *HttpApiDeployConfigSubDomains) GetNetworkType() *string {
	return s.NetworkType
}

func (s *HttpApiDeployConfigSubDomains) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiDeployConfigSubDomains) SetDomainId(v string) *HttpApiDeployConfigSubDomains {
	s.DomainId = &v
	return s
}

func (s *HttpApiDeployConfigSubDomains) SetName(v string) *HttpApiDeployConfigSubDomains {
	s.Name = &v
	return s
}

func (s *HttpApiDeployConfigSubDomains) SetNetworkType(v string) *HttpApiDeployConfigSubDomains {
	s.NetworkType = &v
	return s
}

func (s *HttpApiDeployConfigSubDomains) SetProtocol(v string) *HttpApiDeployConfigSubDomains {
	s.Protocol = &v
	return s
}

func (s *HttpApiDeployConfigSubDomains) Validate() error {
	return dara.Validate(s)
}
