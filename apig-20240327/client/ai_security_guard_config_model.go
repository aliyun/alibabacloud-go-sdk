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
	SetConsumerConfigBindings(v []*AiSecurityGuardConfigConsumerConfigBindings) *AiSecurityGuardConfig
	GetConsumerConfigBindings() []*AiSecurityGuardConfigConsumerConfigBindings
	SetConsumerRequestCheckService(v []*AiSecurityGuardConfigConsumerRequestCheckService) *AiSecurityGuardConfig
	GetConsumerRequestCheckService() []*AiSecurityGuardConfigConsumerRequestCheckService
	SetConsumerResponseCheckService(v []*AiSecurityGuardConfigConsumerResponseCheckService) *AiSecurityGuardConfig
	GetConsumerResponseCheckService() []*AiSecurityGuardConfigConsumerResponseCheckService
	SetConsumerRiskLevel(v []*AiSecurityGuardConfigConsumerRiskLevel) *AiSecurityGuardConfig
	GetConsumerRiskLevel() []*AiSecurityGuardConfigConsumerRiskLevel
	SetDefaultGenerated(v bool) *AiSecurityGuardConfig
	GetDefaultGenerated() *bool
	SetLongTermEnabled(v bool) *AiSecurityGuardConfig
	GetLongTermEnabled() *bool
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
	// The size of the response buffer in KB. Default value: 1000. Valid values: 1 to 1500.
	//
	// example:
	//
	// 1000
	BufferLimit *int32 `json:"bufferLimit,omitempty" xml:"bufferLimit,omitempty"`
	// Specifies whether to check the request content.
	//
	// example:
	//
	// true
	CheckRequest *bool `json:"checkRequest,omitempty" xml:"checkRequest,omitempty"`
	// Specifies whether to check the request images.
	//
	// example:
	//
	// false
	CheckRequestImage *bool `json:"checkRequestImage,omitempty" xml:"checkRequestImage,omitempty"`
	// Specifies whether to check the response content.
	//
	// example:
	//
	// true
	CheckResponse *bool `json:"checkResponse,omitempty" xml:"checkResponse,omitempty"`
	// Specifies whether to check the response images.
	//
	// example:
	//
	// false
	CheckResponseImage *bool `json:"checkResponseImage,omitempty" xml:"checkResponseImage,omitempty"`
	// The bindings between consumers and security guardrail configurations. An empty array explicitly clears the pending configurations.
	//
	// if can be null:
	// true
	ConsumerConfigBindings []*AiSecurityGuardConfigConsumerConfigBindings `json:"consumerConfigBindings,omitempty" xml:"consumerConfigBindings,omitempty" type:"Repeated"`
	// The consumer-level request check service configuration.
	ConsumerRequestCheckService []*AiSecurityGuardConfigConsumerRequestCheckService `json:"consumerRequestCheckService,omitempty" xml:"consumerRequestCheckService,omitempty" type:"Repeated"`
	// The consumer-level response check service configuration.
	ConsumerResponseCheckService []*AiSecurityGuardConfigConsumerResponseCheckService `json:"consumerResponseCheckService,omitempty" xml:"consumerResponseCheckService,omitempty" type:"Repeated"`
	// The consumer-level risk level configuration.
	ConsumerRiskLevel []*AiSecurityGuardConfigConsumerRiskLevel `json:"consumerRiskLevel,omitempty" xml:"consumerRiskLevel,omitempty" type:"Repeated"`
	// Specifies whether the entire security guardrail policy is populated by default for this query. Returns true only for synthesized policies. For existing policies, returns false or is not returned. This is a read-only field. It is ignored upon submission, is not saved, and does not affect the start or stop status.
	//
	// if can be null:
	// true
	DefaultGenerated *bool `json:"defaultGenerated,omitempty" xml:"defaultGenerated,omitempty"`
	// Specifies whether to use the long-term solution. If omitted or set to null, the original value is retained during an update.
	//
	// if can be null:
	// true
	LongTermEnabled *bool `json:"longTermEnabled,omitempty" xml:"longTermEnabled,omitempty"`
	// The running status of the plug-in.
	//
	// if can be null:
	// true
	PluginStatus *AiPluginStatus `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	// The ServiceCode of the request text check service. The system injects the default value.
	//
	// example:
	//
	// query_security_check
	RequestCheckService *string `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	// The ServiceCode of the request image check service. The system injects the default value.
	//
	// example:
	//
	// img_query_guard
	RequestImageCheckService *string `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
	// The ServiceCode of the response text check service. The system injects the default value.
	//
	// example:
	//
	// response_security_check
	ResponseCheckService *string `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	// The ServiceCode of the response image check service. The system injects the default value.
	//
	// example:
	//
	// img_response_guard
	ResponseImageCheckService *string `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
	// The global risk alert level.
	//
	// example:
	//
	// high
	RiskAlertLevel *string `json:"riskAlertLevel,omitempty" xml:"riskAlertLevel,omitempty"`
	// The list of risk dimension configurations. The system injects this list, which is generated by normalizing ConsumerRiskLevel.
	RiskConfig []*AiSecurityGuardConfigRiskConfig `json:"riskConfig,omitempty" xml:"riskConfig,omitempty" type:"Repeated"`
	// The address of the security guardrail service, which is the green-cip endpoint. If the gateway and the security guardrail are in the same region, use the VPC internal network address.
	//
	// example:
	//
	// https://green-cip-vpc.cn-shanghai.aliyuncs.com
	ServiceAddress *string `json:"serviceAddress,omitempty" xml:"serviceAddress,omitempty"`
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

func (s *AiSecurityGuardConfig) GetConsumerConfigBindings() []*AiSecurityGuardConfigConsumerConfigBindings {
	return s.ConsumerConfigBindings
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

func (s *AiSecurityGuardConfig) GetDefaultGenerated() *bool {
	return s.DefaultGenerated
}

func (s *AiSecurityGuardConfig) GetLongTermEnabled() *bool {
	return s.LongTermEnabled
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

func (s *AiSecurityGuardConfig) SetConsumerConfigBindings(v []*AiSecurityGuardConfigConsumerConfigBindings) *AiSecurityGuardConfig {
	s.ConsumerConfigBindings = v
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

func (s *AiSecurityGuardConfig) SetDefaultGenerated(v bool) *AiSecurityGuardConfig {
	s.DefaultGenerated = &v
	return s
}

func (s *AiSecurityGuardConfig) SetLongTermEnabled(v bool) *AiSecurityGuardConfig {
	s.LongTermEnabled = &v
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
	if s.ConsumerConfigBindings != nil {
		for _, item := range s.ConsumerConfigBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type AiSecurityGuardConfigConsumerConfigBindings struct {
	// The match type for consumer names. Valid values: exact (exact match), prefix (prefix match), regexp (regular expression match), and all (matches any consumer). Rules of the all type must be placed at the end of the binding list.
	//
	// example:
	//
	// exact
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The match condition for the consumer name. If matchType is exact, enter the full consumer name. If matchType is prefix, enter the name prefix. If matchType is regexp, enter a regular expression. The value cannot be empty for these three types. If matchType is all, the value must be an empty string.
	//
	// example:
	//
	// example-consumer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the bound security guardrail configuration. This parameter cannot be empty. You can query the ID by calling the ListAiSecurityGuardConfigs operation and use the returned templateId.
	//
	// example:
	//
	// apig_0123456789abcdef
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s AiSecurityGuardConfigConsumerConfigBindings) String() string {
	return dara.Prettify(s)
}

func (s AiSecurityGuardConfigConsumerConfigBindings) GoString() string {
	return s.String()
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) GetMatchType() *string {
	return s.MatchType
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) GetName() *string {
	return s.Name
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) GetTemplateId() *string {
	return s.TemplateId
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) SetMatchType(v string) *AiSecurityGuardConfigConsumerConfigBindings {
	s.MatchType = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) SetName(v string) *AiSecurityGuardConfigConsumerConfigBindings {
	s.Name = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) SetTemplateId(v string) *AiSecurityGuardConfigConsumerConfigBindings {
	s.TemplateId = &v
	return s
}

func (s *AiSecurityGuardConfigConsumerConfigBindings) Validate() error {
	return dara.Validate(s)
}

type AiSecurityGuardConfigConsumerRequestCheckService struct {
	// The consumer match type.
	//
	// example:
	//
	// exact
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The modality type.
	//
	// example:
	//
	// Text
	ModalityType *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// consumer-a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The check service.
	//
	// example:
	//
	// query_security_check
	RequestCheckService *string `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	// The image check service.
	//
	// example:
	//
	// img_query_guard
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
	// The consumer match type.
	//
	// example:
	//
	// exact
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The modality type.
	//
	// example:
	//
	// Text
	ModalityType *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// consumer-a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The check service.
	//
	// example:
	//
	// query_security_check
	ResponseCheckService *string `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	// The image check service.
	//
	// example:
	//
	// img_query_guard
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
	// The risk level.
	//
	// example:
	//
	// high
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The consumer match type.
	//
	// example:
	//
	// exact
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// consumer-a
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The risk dimension type.
	//
	// example:
	//
	// ContentModeration
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// The consumer-level match rules.
	ConsumerRules *AiSecurityGuardConfigRiskConfigConsumerRules `json:"consumerRules,omitempty" xml:"consumerRules,omitempty" type:"Struct"`
	// The risk level.
	//
	// example:
	//
	// high
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// The risk dimension type.
	//
	// example:
	//
	// ContentModeration
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// The match type.
	//
	// example:
	//
	// Exact
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	// The consumer match pattern value.
	//
	// example:
	//
	// consumer-a
	Pattern *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
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
