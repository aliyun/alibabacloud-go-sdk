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
	AiFallbackConfig       *HttpApiDeployConfigPolicyConfigsAiFallbackConfig       `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty" type:"Struct"`
	AiSecurityGuardConfig  *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig  `json:"aiSecurityGuardConfig,omitempty" xml:"aiSecurityGuardConfig,omitempty" type:"Struct"`
	AiTokenRateLimitConfig *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig `json:"aiTokenRateLimitConfig,omitempty" xml:"aiTokenRateLimitConfig,omitempty" type:"Struct"`
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

func (s *HttpApiDeployConfigPolicyConfigs) GetAiSecurityGuardConfig() *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	return s.AiSecurityGuardConfig
}

func (s *HttpApiDeployConfigPolicyConfigs) GetAiTokenRateLimitConfig() *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig {
	return s.AiTokenRateLimitConfig
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

func (s *HttpApiDeployConfigPolicyConfigs) SetAiSecurityGuardConfig(v *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) *HttpApiDeployConfigPolicyConfigs {
	s.AiSecurityGuardConfig = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigs) SetAiTokenRateLimitConfig(v *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) *HttpApiDeployConfigPolicyConfigs {
	s.AiTokenRateLimitConfig = v
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
	if s.AiSecurityGuardConfig != nil {
		if err := s.AiSecurityGuardConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiTokenRateLimitConfig != nil {
		if err := s.AiTokenRateLimitConfig.Validate(); err != nil {
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

type HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig struct {
	BufferLimit                  *int32                                                                               `json:"bufferLimit,omitempty" xml:"bufferLimit,omitempty"`
	CheckRequest                 *bool                                                                                `json:"checkRequest,omitempty" xml:"checkRequest,omitempty"`
	CheckRequestImage            *bool                                                                                `json:"checkRequestImage,omitempty" xml:"checkRequestImage,omitempty"`
	CheckResponse                *bool                                                                                `json:"checkResponse,omitempty" xml:"checkResponse,omitempty"`
	CheckResponseImage           *bool                                                                                `json:"checkResponseImage,omitempty" xml:"checkResponseImage,omitempty"`
	ConsumerRequestCheckService  []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService  `json:"consumerRequestCheckService,omitempty" xml:"consumerRequestCheckService,omitempty" type:"Repeated"`
	ConsumerResponseCheckService []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService `json:"consumerResponseCheckService,omitempty" xml:"consumerResponseCheckService,omitempty" type:"Repeated"`
	ConsumerRiskLevel            []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel            `json:"consumerRiskLevel,omitempty" xml:"consumerRiskLevel,omitempty" type:"Repeated"`
	RequestCheckService          *string                                                                              `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService     *string                                                                              `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
	ResponseCheckService         *string                                                                              `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService    *string                                                                              `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
	RiskAlertLevel               *string                                                                              `json:"riskAlertLevel,omitempty" xml:"riskAlertLevel,omitempty"`
	RiskConfig                   []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig                   `json:"riskConfig,omitempty" xml:"riskConfig,omitempty" type:"Repeated"`
	ServiceAddress               *string                                                                              `json:"serviceAddress,omitempty" xml:"serviceAddress,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetBufferLimit() *int32 {
	return s.BufferLimit
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetCheckRequest() *bool {
	return s.CheckRequest
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetCheckRequestImage() *bool {
	return s.CheckRequestImage
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetCheckResponse() *bool {
	return s.CheckResponse
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetCheckResponseImage() *bool {
	return s.CheckResponseImage
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetConsumerRequestCheckService() []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	return s.ConsumerRequestCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetConsumerResponseCheckService() []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	return s.ConsumerResponseCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetConsumerRiskLevel() []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	return s.ConsumerRiskLevel
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetRiskAlertLevel() *string {
	return s.RiskAlertLevel
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetRiskConfig() []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig {
	return s.RiskConfig
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetBufferLimit(v int32) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.BufferLimit = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetCheckRequest(v bool) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.CheckRequest = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetCheckRequestImage(v bool) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.CheckRequestImage = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetCheckResponse(v bool) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.CheckResponse = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetCheckResponseImage(v bool) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.CheckResponseImage = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetConsumerRequestCheckService(v []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerRequestCheckService = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetConsumerResponseCheckService(v []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerResponseCheckService = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetConsumerRiskLevel(v []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerRiskLevel = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetRequestCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.RequestCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetRequestImageCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.RequestImageCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetResponseCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.ResponseCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetResponseImageCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.ResponseImageCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetRiskAlertLevel(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.RiskAlertLevel = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetRiskConfig(v []*HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.RiskConfig = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) SetServiceAddress(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig {
	s.ServiceAddress = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfig) Validate() error {
	if s.ConsumerRequestCheckService != nil {
		for _, item := range s.ConsumerRequestCheckService {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConsumerResponseCheckService != nil {
		for _, item := range s.ConsumerResponseCheckService {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConsumerRiskLevel != nil {
		for _, item := range s.ConsumerRiskLevel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskConfig != nil {
		for _, item := range s.RiskConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService struct {
	MatchType                *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType             *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                     *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestCheckService      *string `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService *string `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetName() *string {
	return s.Name
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetMatchType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.MatchType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetModalityType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.ModalityType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetName(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.Name = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetRequestCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetRequestImageCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestImageCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService struct {
	MatchType                 *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType              *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                      *string `json:"name,omitempty" xml:"name,omitempty"`
	ResponseCheckService      *string `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService *string `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetName() *string {
	return s.Name
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetMatchType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.MatchType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetModalityType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ModalityType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetName(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.Name = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetResponseCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetResponseImageCheckService(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseImageCheckService = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel struct {
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetLevel() *string {
	return s.Level
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetName() *string {
	return s.Name
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetType() *string {
	return s.Type
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetLevel(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Level = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetMatchType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.MatchType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetName(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Name = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Type = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig struct {
	ConsumerRules *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules `json:"consumerRules,omitempty" xml:"consumerRules,omitempty" type:"Struct"`
	Level         *string                                                                       `json:"level,omitempty" xml:"level,omitempty"`
	Type          *string                                                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) GetConsumerRules() *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	return s.ConsumerRules
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) GetLevel() *string {
	return s.Level
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) GetType() *string {
	return s.Type
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) SetConsumerRules(v *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.ConsumerRules = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) SetLevel(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.Level = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) SetType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.Type = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfig) Validate() error {
	if s.ConsumerRules != nil {
		if err := s.ConsumerRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules struct {
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Pattern   *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GetPattern() *string {
	return s.Pattern
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) SetMatchType(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	s.MatchType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) SetPattern(v string) *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	s.Pattern = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig struct {
	EnableGlobalRules *bool                                                                `json:"enableGlobalRules,omitempty" xml:"enableGlobalRules,omitempty"`
	GlobalRules       []*HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules `json:"globalRules,omitempty" xml:"globalRules,omitempty" type:"Repeated"`
	Rules             []*HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules       `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) GetEnableGlobalRules() *bool {
	return s.EnableGlobalRules
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) GetGlobalRules() []*HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	return s.GlobalRules
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) GetRules() []*HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	return s.Rules
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) SetEnableGlobalRules(v bool) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig {
	s.EnableGlobalRules = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) SetGlobalRules(v []*HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig {
	s.GlobalRules = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) SetRules(v []*HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig {
	s.Rules = v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfig) Validate() error {
	if s.GlobalRules != nil {
		for _, item := range s.GlobalRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *int32  `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitMode() *string {
	return s.LimitMode
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitType() *string {
	return s.LimitType
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitValue() *int32 {
	return s.LimitValue
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchKey() *string {
	return s.MatchKey
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchValue() *string {
	return s.MatchValue
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitMode(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitMode = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitType(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitValue(v int32) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitValue = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchKey(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchKey = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchType(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchValue(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchValue = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigGlobalRules) Validate() error {
	return dara.Validate(s)
}

type HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *int32  `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GoString() string {
	return s.String()
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GetLimitMode() *string {
	return s.LimitMode
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GetLimitType() *string {
	return s.LimitType
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GetLimitValue() *int32 {
	return s.LimitValue
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GetMatchKey() *string {
	return s.MatchKey
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) GetMatchValue() *string {
	return s.MatchValue
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) SetLimitMode(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitMode = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) SetLimitType(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) SetLimitValue(v int32) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitValue = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) SetMatchKey(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchKey = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) SetMatchType(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchType = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) SetMatchValue(v string) *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchValue = &v
	return s
}

func (s *HttpApiDeployConfigPolicyConfigsAiTokenRateLimitConfigRules) Validate() error {
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
