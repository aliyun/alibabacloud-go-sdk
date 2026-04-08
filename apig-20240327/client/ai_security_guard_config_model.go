// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiSecurityGuardConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBufferLimit(v int32) *AiSecurityGuardConfig
	GetBufferLimit() *int32
	SetCheckRequest(v bool) *AiSecurityGuardConfig
	GetCheckRequest() *bool
	SetCheckRequestImage(v bool) *AiSecurityGuardConfig
	GetCheckRequestImage() *bool
	SetCheckResponse(v bool) *AiSecurityGuardConfig
	GetCheckResponse() *bool
	SetCheckResponseImage(v bool) *AiSecurityGuardConfig
	GetCheckResponseImage() *bool
	SetConsumerRequestCheckService(v []*AiSecurityGuardConfigConsumerRequestCheckService) *AiSecurityGuardConfig
	GetConsumerRequestCheckService() []*AiSecurityGuardConfigConsumerRequestCheckService
	SetConsumerResponseCheckService(v []*AiSecurityGuardConfigConsumerResponseCheckService) *AiSecurityGuardConfig
	GetConsumerResponseCheckService() []*AiSecurityGuardConfigConsumerResponseCheckService
	SetConsumerRiskLevel(v []*AiSecurityGuardConfigConsumerRiskLevel) *AiSecurityGuardConfig
	GetConsumerRiskLevel() []*AiSecurityGuardConfigConsumerRiskLevel
	SetPluginStatus(v *AiPluginStatus) *AiSecurityGuardConfig
	GetPluginStatus() *AiPluginStatus
	SetRequestCheckService(v string) *AiSecurityGuardConfig
	GetRequestCheckService() *string
	SetRequestImageCheckService(v string) *AiSecurityGuardConfig
	GetRequestImageCheckService() *string
	SetResponseCheckService(v string) *AiSecurityGuardConfig
	GetResponseCheckService() *string
	SetResponseImageCheckService(v string) *AiSecurityGuardConfig
	GetResponseImageCheckService() *string
	SetRiskAlertLevel(v string) *AiSecurityGuardConfig
	GetRiskAlertLevel() *string
	SetRiskConfig(v []*AiSecurityGuardConfigRiskConfig) *AiSecurityGuardConfig
	GetRiskConfig() []*AiSecurityGuardConfigRiskConfig
	SetServiceAddress(v string) *AiSecurityGuardConfig
	GetServiceAddress() *string
}

type AiSecurityGuardConfig struct {
	BufferLimit                  *int32                                               `json:"bufferLimit,omitempty" xml:"bufferLimit,omitempty"`
	CheckRequest                 *bool                                                `json:"checkRequest,omitempty" xml:"checkRequest,omitempty"`
	CheckRequestImage            *bool                                                `json:"checkRequestImage,omitempty" xml:"checkRequestImage,omitempty"`
	CheckResponse                *bool                                                `json:"checkResponse,omitempty" xml:"checkResponse,omitempty"`
	CheckResponseImage           *bool                                                `json:"checkResponseImage,omitempty" xml:"checkResponseImage,omitempty"`
	ConsumerRequestCheckService  []*AiSecurityGuardConfigConsumerRequestCheckService  `json:"consumerRequestCheckService,omitempty" xml:"consumerRequestCheckService,omitempty" type:"Repeated"`
	ConsumerResponseCheckService []*AiSecurityGuardConfigConsumerResponseCheckService `json:"consumerResponseCheckService,omitempty" xml:"consumerResponseCheckService,omitempty" type:"Repeated"`
	ConsumerRiskLevel            []*AiSecurityGuardConfigConsumerRiskLevel            `json:"consumerRiskLevel,omitempty" xml:"consumerRiskLevel,omitempty" type:"Repeated"`
	// if can be null:
	// true
	PluginStatus              *AiPluginStatus                    `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	RequestCheckService       *string                            `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService  *string                            `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
	ResponseCheckService      *string                            `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService *string                            `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
	RiskAlertLevel            *string                            `json:"riskAlertLevel,omitempty" xml:"riskAlertLevel,omitempty"`
	RiskConfig                []*AiSecurityGuardConfigRiskConfig `json:"riskConfig,omitempty" xml:"riskConfig,omitempty" type:"Repeated"`
	ServiceAddress            *string                            `json:"serviceAddress,omitempty" xml:"serviceAddress,omitempty"`
}

func (s AiSecurityGuardConfig) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfig) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfig) GetBufferLimit() *int32 {
	return s.BufferLimit
}

func (s *AiSecurityGuardConfig) GetCheckRequest() *bool {
	return s.CheckRequest
}

func (s *AiSecurityGuardConfig) GetCheckRequestImage() *bool {
	return s.CheckRequestImage
}

func (s *AiSecurityGuardConfig) GetCheckResponse() *bool {
	return s.CheckResponse
}

func (s *AiSecurityGuardConfig) GetCheckResponseImage() *bool {
	return s.CheckResponseImage
}

func (s *AiSecurityGuardConfig) GetConsumerRequestCheckService() []*AiSecurityGuardConfigConsumerRequestCheckService {
	return s.ConsumerRequestCheckService
}

func (s *AiSecurityGuardConfig) GetConsumerResponseCheckService() []*AiSecurityGuardConfigConsumerResponseCheckService {
	return s.ConsumerResponseCheckService
}

func (s *AiSecurityGuardConfig) GetConsumerRiskLevel() []*AiSecurityGuardConfigConsumerRiskLevel {
	return s.ConsumerRiskLevel
}

func (s *AiSecurityGuardConfig) GetPluginStatus() *AiPluginStatus {
	return s.PluginStatus
}

func (s *AiSecurityGuardConfig) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *AiSecurityGuardConfig) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *AiSecurityGuardConfig) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *AiSecurityGuardConfig) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *AiSecurityGuardConfig) GetRiskAlertLevel() *string {
	return s.RiskAlertLevel
}

func (s *AiSecurityGuardConfig) GetRiskConfig() []*AiSecurityGuardConfigRiskConfig {
	return s.RiskConfig
}

func (s *AiSecurityGuardConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *AiSecurityGuardConfig) SetBufferLimit(v int32) *AiSecurityGuardConfig {
	s.BufferLimit = &v
	return s
}

func (s *AiSecurityGuardConfig) SetCheckRequest(v bool) *AiSecurityGuardConfig {
	s.CheckRequest = &v
	return s
}

func (s *AiSecurityGuardConfig) SetCheckRequestImage(v bool) *AiSecurityGuardConfig {
	s.CheckRequestImage = &v
	return s
}

func (s *AiSecurityGuardConfig) SetCheckResponse(v bool) *AiSecurityGuardConfig {
	s.CheckResponse = &v
	return s
}

func (s *AiSecurityGuardConfig) SetCheckResponseImage(v bool) *AiSecurityGuardConfig {
	s.CheckResponseImage = &v
	return s
}

func (s *AiSecurityGuardConfig) SetConsumerRequestCheckService(v []*AiSecurityGuardConfigConsumerRequestCheckService) *AiSecurityGuardConfig {
	s.ConsumerRequestCheckService = v
	return s
}

func (s *AiSecurityGuardConfig) SetConsumerResponseCheckService(v []*AiSecurityGuardConfigConsumerResponseCheckService) *AiSecurityGuardConfig {
	s.ConsumerResponseCheckService = v
	return s
}

func (s *AiSecurityGuardConfig) SetConsumerRiskLevel(v []*AiSecurityGuardConfigConsumerRiskLevel) *AiSecurityGuardConfig {
	s.ConsumerRiskLevel = v
	return s
}

func (s *AiSecurityGuardConfig) SetPluginStatus(v *AiPluginStatus) *AiSecurityGuardConfig {
	s.PluginStatus = v
	return s
}

func (s *AiSecurityGuardConfig) SetRequestCheckService(v string) *AiSecurityGuardConfig {
	s.RequestCheckService = &v
	return s
}

func (s *AiSecurityGuardConfig) SetRequestImageCheckService(v string) *AiSecurityGuardConfig {
	s.RequestImageCheckService = &v
	return s
}

func (s *AiSecurityGuardConfig) SetResponseCheckService(v string) *AiSecurityGuardConfig {
	s.ResponseCheckService = &v
	return s
}

func (s *AiSecurityGuardConfig) SetResponseImageCheckService(v string) *AiSecurityGuardConfig {
	s.ResponseImageCheckService = &v
	return s
}

func (s *AiSecurityGuardConfig) SetRiskAlertLevel(v string) *AiSecurityGuardConfig {
	s.RiskAlertLevel = &v
	return s
}

func (s *AiSecurityGuardConfig) SetRiskConfig(v []*AiSecurityGuardConfigRiskConfig) *AiSecurityGuardConfig {
	s.RiskConfig = v
	return s
}

func (s *AiSecurityGuardConfig) SetServiceAddress(v string) *AiSecurityGuardConfig {
	s.ServiceAddress = &v
	return s
}

func (s *AiSecurityGuardConfig) Validate() error {
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
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
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

type AiSecurityGuardConfigConsumerRequestCheckService struct {
	MatchType                *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType             *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                     *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestCheckService      *string `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService *string `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
}

func (s AiSecurityGuardConfigConsumerRequestCheckService) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfigConsumerRequestCheckService) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) GetName() *string {
	return s.Name
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) SetMatchType(v string) *AiSecurityGuardConfigConsumerRequestCheckService {
	s.MatchType = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) SetModalityType(v string) *AiSecurityGuardConfigConsumerRequestCheckService {
	s.ModalityType = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) SetName(v string) *AiSecurityGuardConfigConsumerRequestCheckService {
	s.Name = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) SetRequestCheckService(v string) *AiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestCheckService = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) SetRequestImageCheckService(v string) *AiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestImageCheckService = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRequestCheckService) Validate() error {
	return dara.Validate(s)
}

type AiSecurityGuardConfigConsumerResponseCheckService struct {
	MatchType                 *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType              *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                      *string `json:"name,omitempty" xml:"name,omitempty"`
	ResponseCheckService      *string `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService *string `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
}

func (s AiSecurityGuardConfigConsumerResponseCheckService) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfigConsumerResponseCheckService) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) GetName() *string {
	return s.Name
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) SetMatchType(v string) *AiSecurityGuardConfigConsumerResponseCheckService {
	s.MatchType = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) SetModalityType(v string) *AiSecurityGuardConfigConsumerResponseCheckService {
	s.ModalityType = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) SetName(v string) *AiSecurityGuardConfigConsumerResponseCheckService {
	s.Name = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) SetResponseCheckService(v string) *AiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseCheckService = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) SetResponseImageCheckService(v string) *AiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseImageCheckService = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerResponseCheckService) Validate() error {
	return dara.Validate(s)
}

type AiSecurityGuardConfigConsumerRiskLevel struct {
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AiSecurityGuardConfigConsumerRiskLevel) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfigConsumerRiskLevel) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) GetLevel() *string {
	return s.Level
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) GetMatchType() *string {
	return s.MatchType
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) GetName() *string {
	return s.Name
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) GetType() *string {
	return s.Type
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) SetLevel(v string) *AiSecurityGuardConfigConsumerRiskLevel {
	s.Level = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) SetMatchType(v string) *AiSecurityGuardConfigConsumerRiskLevel {
	s.MatchType = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) SetName(v string) *AiSecurityGuardConfigConsumerRiskLevel {
	s.Name = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) SetType(v string) *AiSecurityGuardConfigConsumerRiskLevel {
	s.Type = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerRiskLevel) Validate() error {
	return dara.Validate(s)
}

type AiSecurityGuardConfigRiskConfig struct {
	ConsumerRules *AiSecurityGuardConfigRiskConfigConsumerRules `json:"consumerRules,omitempty" xml:"consumerRules,omitempty" type:"Struct"`
	Level         *string                                       `json:"level,omitempty" xml:"level,omitempty"`
	Type          *string                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AiSecurityGuardConfigRiskConfig) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfigRiskConfig) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfigRiskConfig) GetConsumerRules() *AiSecurityGuardConfigRiskConfigConsumerRules {
	return s.ConsumerRules
}

func (s *AiSecurityGuardConfigRiskConfig) GetLevel() *string {
	return s.Level
}

func (s *AiSecurityGuardConfigRiskConfig) GetType() *string {
	return s.Type
}

func (s *AiSecurityGuardConfigRiskConfig) SetConsumerRules(v *AiSecurityGuardConfigRiskConfigConsumerRules) *AiSecurityGuardConfigRiskConfig {
	s.ConsumerRules = v
	return s
}

func (s *AiSecurityGuardConfigRiskConfig) SetLevel(v string) *AiSecurityGuardConfigRiskConfig {
	s.Level = &v
	return s
}

func (s *AiSecurityGuardConfigRiskConfig) SetType(v string) *AiSecurityGuardConfigRiskConfig {
	s.Type = &v
	return s
}

func (s *AiSecurityGuardConfigRiskConfig) Validate() error {
	if s.ConsumerRules != nil {
		if err := s.ConsumerRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AiSecurityGuardConfigRiskConfigConsumerRules struct {
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Pattern   *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
}

func (s AiSecurityGuardConfigRiskConfigConsumerRules) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfigRiskConfigConsumerRules) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfigRiskConfigConsumerRules) GetMatchType() *string {
	return s.MatchType
}

func (s *AiSecurityGuardConfigRiskConfigConsumerRules) GetPattern() *string {
	return s.Pattern
}

func (s *AiSecurityGuardConfigRiskConfigConsumerRules) SetMatchType(v string) *AiSecurityGuardConfigRiskConfigConsumerRules {
	s.MatchType = &v
	return s
}

func (s *AiSecurityGuardConfigRiskConfigConsumerRules) SetPattern(v string) *AiSecurityGuardConfigRiskConfigConsumerRules {
	s.Pattern = &v
	return s
}

func (s *AiSecurityGuardConfigRiskConfigConsumerRules) Validate() error {
	return dara.Validate(s)
}
