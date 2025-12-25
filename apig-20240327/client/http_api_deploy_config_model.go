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
	SetPolicyConfigs(v []*HttpApiDeployConfigPolicyConfigs) *HttpApiDeployConfig
	GetPolicyConfigs() []*HttpApiDeployConfigPolicyConfigs
	SetRouteBackend(v *Backend) *HttpApiDeployConfig
	GetRouteBackend() *Backend
	SetServiceConfigs(v []*HttpApiDeployConfigServiceConfigs) *HttpApiDeployConfig
	GetServiceConfigs() []*HttpApiDeployConfigServiceConfigs
	SetSubDomains(v []*HttpApiDeployConfigSubDomains) *HttpApiDeployConfig
	GetSubDomains() []*HttpApiDeployConfigSubDomains
}

type HttpApiDeployConfig struct {
	// example:
	//
	// true
	AutoDeploy *bool `json:"autoDeploy,omitempty" xml:"autoDeploy,omitempty"`
	// example:
	//
	// SingleService
	BackendScene      *string                                 `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	CustomDomainIds   []*string                               `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	CustomDomainInfos []*HttpApiDeployConfigCustomDomainInfos `json:"customDomainInfos,omitempty" xml:"customDomainInfos,omitempty" type:"Repeated"`
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-xx
	GatewayId   *string      `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	GatewayInfo *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// example:
	//
	// API
	GatewayType    *string                              `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	Mock           *HttpApiMockContract                 `json:"mock,omitempty" xml:"mock,omitempty"`
	PolicyConfigs  []*HttpApiDeployConfigPolicyConfigs  `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
	RouteBackend   *Backend                             `json:"routeBackend,omitempty" xml:"routeBackend,omitempty"`
	ServiceConfigs []*HttpApiDeployConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
	SubDomains     []*HttpApiDeployConfigSubDomains     `json:"subDomains,omitempty" xml:"subDomains,omitempty" type:"Repeated"`
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

func (s *HttpApiDeployConfig) GetPolicyConfigs() []*HttpApiDeployConfigPolicyConfigs {
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

func (s *HttpApiDeployConfig) SetPolicyConfigs(v []*HttpApiDeployConfigPolicyConfigs) *HttpApiDeployConfig {
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
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
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

type HttpApiDeployConfigPolicyConfigs struct {
	AiFallbackConfig *HttpApiDeployConfigPolicyConfigsAiFallbackConfig `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty" type:"Struct"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// AiFallback
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigs) GetAiFallbackConfig() *HttpApiDeployConfigPolicyConfigsAiFallbackConfig {
	return s.AiFallbackConfig
}

func (s *HttpApiDeployConfigPolicyConfigs) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiDeployConfigPolicyConfigs) GetType() *string {
	return s.Type
}

func (s *HttpApiDeployConfigPolicyConfigs) SetAiFallbackConfig(v *HttpApiDeployConfigPolicyConfigsAiFallbackConfig) *HttpApiDeployConfigPolicyConfigs {
	s.AiFallbackConfig = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigs) SetEnable(v bool) *HttpApiDeployConfigPolicyConfigs {
	s.Enable = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigs) SetType(v string) *HttpApiDeployConfigPolicyConfigs {
	s.Type = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigs) Validate() error {
	if s.AiFallbackConfig != nil {
		if err := s.AiFallbackConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiDeployConfigPolicyConfigsAiFallbackConfig struct {
	ServiceConfigs []*HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiDeployConfigPolicyConfigsAiFallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiFallbackConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfig) GetServiceConfigs() []*HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs {
	return s.ServiceConfigs
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfig) SetServiceConfigs(v []*HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) *HttpApiDeployConfigPolicyConfigsAiFallbackConfig {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfig) Validate() error {
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

type HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs struct {
	ServiceId       *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TargetModelName *string `json:"targetModelName,omitempty" xml:"targetModelName,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) GetTargetModelName() *string {
	return s.TargetModelName
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) SetServiceId(v string) *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) SetTargetModelName(v string) *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs {
	s.TargetModelName = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiFallbackConfigServiceConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigServiceConfigs struct {
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// qwen-*
	ModelNamePattern *string `json:"modelNamePattern,omitempty" xml:"modelNamePattern,omitempty"`
	// example:
	//
	// svc-xxx
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
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

func (s *HttpApiDeployConfigServiceConfigs) GetIntentCode() *string {
	return s.IntentCode
}

func (s *HttpApiDeployConfigServiceConfigs) GetModelName() *string {
	return s.ModelName
}

func (s *HttpApiDeployConfigServiceConfigs) GetModelNamePattern() *string {
	return s.ModelNamePattern
}

func (s *HttpApiDeployConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiDeployConfigServiceConfigs) GetWeight() *int64 {
	return s.Weight
}

func (s *HttpApiDeployConfigServiceConfigs) SetIntentCode(v string) *HttpApiDeployConfigServiceConfigs {
	s.IntentCode = &v
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

func (s *HttpApiDeployConfigServiceConfigs) SetServiceId(v string) *HttpApiDeployConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) SetWeight(v int64) *HttpApiDeployConfigServiceConfigs {
	s.Weight = &v
	return s
}

func (s *HttpApiDeployConfigServiceConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigSubDomains struct {
	DomainId    *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	Protocol    *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
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
