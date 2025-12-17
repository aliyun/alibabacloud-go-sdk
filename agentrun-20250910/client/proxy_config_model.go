// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProxyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEndpoints(v []*ProxyConfigEndpoints) *ProxyConfig
	GetEndpoints() []*ProxyConfigEndpoints
	SetPolicies(v *ProxyConfigPolicies) *ProxyConfig
	GetPolicies() *ProxyConfigPolicies
}

type ProxyConfig struct {
	Endpoints []*ProxyConfigEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	Policies  *ProxyConfigPolicies    `json:"policies,omitempty" xml:"policies,omitempty" type:"Struct"`
}

func (s ProxyConfig) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfig) GoString() string {
	return s.String()
}

func (s *ProxyConfig) GetEndpoints() []*ProxyConfigEndpoints {
	return s.Endpoints
}

func (s *ProxyConfig) GetPolicies() *ProxyConfigPolicies {
	return s.Policies
}

func (s *ProxyConfig) SetEndpoints(v []*ProxyConfigEndpoints) *ProxyConfig {
	s.Endpoints = v
	return s
}

func (s *ProxyConfig) SetPolicies(v *ProxyConfigPolicies) *ProxyConfig {
	s.Policies = v
	return s
}

func (s *ProxyConfig) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Policies != nil {
		if err := s.Policies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProxyConfigEndpoints struct {
	BaseUrl          *string   `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	ModelNames       []*string `json:"modelNames,omitempty" xml:"modelNames,omitempty" type:"Repeated"`
	ModelServiceName *string   `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
	Weight           *int32    `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ProxyConfigEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfigEndpoints) GoString() string {
	return s.String()
}

func (s *ProxyConfigEndpoints) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ProxyConfigEndpoints) GetModelNames() []*string {
	return s.ModelNames
}

func (s *ProxyConfigEndpoints) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *ProxyConfigEndpoints) GetWeight() *int32 {
	return s.Weight
}

func (s *ProxyConfigEndpoints) SetBaseUrl(v string) *ProxyConfigEndpoints {
	s.BaseUrl = &v
	return s
}

func (s *ProxyConfigEndpoints) SetModelNames(v []*string) *ProxyConfigEndpoints {
	s.ModelNames = v
	return s
}

func (s *ProxyConfigEndpoints) SetModelServiceName(v string) *ProxyConfigEndpoints {
	s.ModelServiceName = &v
	return s
}

func (s *ProxyConfigEndpoints) SetWeight(v int32) *ProxyConfigEndpoints {
	s.Weight = &v
	return s
}

func (s *ProxyConfigEndpoints) Validate() error {
	return dara.Validate(s)
}

type ProxyConfigPolicies struct {
	AiGuardrailConfig *ProxyConfigPoliciesAiGuardrailConfig `json:"aiGuardrailConfig,omitempty" xml:"aiGuardrailConfig,omitempty" type:"Struct"`
	Cache             *bool                                 `json:"cache,omitempty" xml:"cache,omitempty"`
	ConcurrencyLimit  *int32                                `json:"concurrencyLimit,omitempty" xml:"concurrencyLimit,omitempty"`
	Fallbacks         []*ProxyConfigPoliciesFallbacks       `json:"fallbacks,omitempty" xml:"fallbacks,omitempty" type:"Repeated"`
	NumRetries        *int32                                `json:"numRetries,omitempty" xml:"numRetries,omitempty"`
	RequestTimeout    *int32                                `json:"requestTimeout,omitempty" xml:"requestTimeout,omitempty"`
	TokenRateLimiter  *ProxyConfigPoliciesTokenRateLimiter  `json:"tokenRateLimiter,omitempty" xml:"tokenRateLimiter,omitempty" type:"Struct"`
}

func (s ProxyConfigPolicies) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfigPolicies) GoString() string {
	return s.String()
}

func (s *ProxyConfigPolicies) GetAiGuardrailConfig() *ProxyConfigPoliciesAiGuardrailConfig {
	return s.AiGuardrailConfig
}

func (s *ProxyConfigPolicies) GetCache() *bool {
	return s.Cache
}

func (s *ProxyConfigPolicies) GetConcurrencyLimit() *int32 {
	return s.ConcurrencyLimit
}

func (s *ProxyConfigPolicies) GetFallbacks() []*ProxyConfigPoliciesFallbacks {
	return s.Fallbacks
}

func (s *ProxyConfigPolicies) GetNumRetries() *int32 {
	return s.NumRetries
}

func (s *ProxyConfigPolicies) GetRequestTimeout() *int32 {
	return s.RequestTimeout
}

func (s *ProxyConfigPolicies) GetTokenRateLimiter() *ProxyConfigPoliciesTokenRateLimiter {
	return s.TokenRateLimiter
}

func (s *ProxyConfigPolicies) SetAiGuardrailConfig(v *ProxyConfigPoliciesAiGuardrailConfig) *ProxyConfigPolicies {
	s.AiGuardrailConfig = v
	return s
}

func (s *ProxyConfigPolicies) SetCache(v bool) *ProxyConfigPolicies {
	s.Cache = &v
	return s
}

func (s *ProxyConfigPolicies) SetConcurrencyLimit(v int32) *ProxyConfigPolicies {
	s.ConcurrencyLimit = &v
	return s
}

func (s *ProxyConfigPolicies) SetFallbacks(v []*ProxyConfigPoliciesFallbacks) *ProxyConfigPolicies {
	s.Fallbacks = v
	return s
}

func (s *ProxyConfigPolicies) SetNumRetries(v int32) *ProxyConfigPolicies {
	s.NumRetries = &v
	return s
}

func (s *ProxyConfigPolicies) SetRequestTimeout(v int32) *ProxyConfigPolicies {
	s.RequestTimeout = &v
	return s
}

func (s *ProxyConfigPolicies) SetTokenRateLimiter(v *ProxyConfigPoliciesTokenRateLimiter) *ProxyConfigPolicies {
	s.TokenRateLimiter = v
	return s
}

func (s *ProxyConfigPolicies) Validate() error {
	if s.AiGuardrailConfig != nil {
		if err := s.AiGuardrailConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Fallbacks != nil {
		for _, item := range s.Fallbacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TokenRateLimiter != nil {
		if err := s.TokenRateLimiter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProxyConfigPoliciesAiGuardrailConfig struct {
	BlockOnContentModeration  *bool   `json:"blockOnContentModeration,omitempty" xml:"blockOnContentModeration,omitempty"`
	BlockOnMaliciousUrl       *bool   `json:"blockOnMaliciousUrl,omitempty" xml:"blockOnMaliciousUrl,omitempty"`
	BlockOnModelHallucination *bool   `json:"blockOnModelHallucination,omitempty" xml:"blockOnModelHallucination,omitempty"`
	BlockOnPromptAttack       *bool   `json:"blockOnPromptAttack,omitempty" xml:"blockOnPromptAttack,omitempty"`
	BlockOnSensitiveData      *bool   `json:"blockOnSensitiveData,omitempty" xml:"blockOnSensitiveData,omitempty"`
	CheckRequest              *bool   `json:"checkRequest,omitempty" xml:"checkRequest,omitempty"`
	CheckResponse             *bool   `json:"checkResponse,omitempty" xml:"checkResponse,omitempty"`
	Level                     *string `json:"level,omitempty" xml:"level,omitempty"`
	MaxTextLength             *int32  `json:"maxTextLength,omitempty" xml:"maxTextLength,omitempty"`
}

func (s ProxyConfigPoliciesAiGuardrailConfig) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfigPoliciesAiGuardrailConfig) GoString() string {
	return s.String()
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetBlockOnContentModeration() *bool {
	return s.BlockOnContentModeration
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetBlockOnMaliciousUrl() *bool {
	return s.BlockOnMaliciousUrl
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetBlockOnModelHallucination() *bool {
	return s.BlockOnModelHallucination
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetBlockOnPromptAttack() *bool {
	return s.BlockOnPromptAttack
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetBlockOnSensitiveData() *bool {
	return s.BlockOnSensitiveData
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetCheckRequest() *bool {
	return s.CheckRequest
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetCheckResponse() *bool {
	return s.CheckResponse
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetLevel() *string {
	return s.Level
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) GetMaxTextLength() *int32 {
	return s.MaxTextLength
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetBlockOnContentModeration(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.BlockOnContentModeration = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetBlockOnMaliciousUrl(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.BlockOnMaliciousUrl = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetBlockOnModelHallucination(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.BlockOnModelHallucination = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetBlockOnPromptAttack(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.BlockOnPromptAttack = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetBlockOnSensitiveData(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.BlockOnSensitiveData = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetCheckRequest(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.CheckRequest = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetCheckResponse(v bool) *ProxyConfigPoliciesAiGuardrailConfig {
	s.CheckResponse = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetLevel(v string) *ProxyConfigPoliciesAiGuardrailConfig {
	s.Level = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) SetMaxTextLength(v int32) *ProxyConfigPoliciesAiGuardrailConfig {
	s.MaxTextLength = &v
	return s
}

func (s *ProxyConfigPoliciesAiGuardrailConfig) Validate() error {
	return dara.Validate(s)
}

type ProxyConfigPoliciesFallbacks struct {
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
}

func (s ProxyConfigPoliciesFallbacks) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfigPoliciesFallbacks) GoString() string {
	return s.String()
}

func (s *ProxyConfigPoliciesFallbacks) GetModelName() *string {
	return s.ModelName
}

func (s *ProxyConfigPoliciesFallbacks) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *ProxyConfigPoliciesFallbacks) SetModelName(v string) *ProxyConfigPoliciesFallbacks {
	s.ModelName = &v
	return s
}

func (s *ProxyConfigPoliciesFallbacks) SetModelServiceName(v string) *ProxyConfigPoliciesFallbacks {
	s.ModelServiceName = &v
	return s
}

func (s *ProxyConfigPoliciesFallbacks) Validate() error {
	return dara.Validate(s)
}

type ProxyConfigPoliciesTokenRateLimiter struct {
	Tpd *int32 `json:"tpd,omitempty" xml:"tpd,omitempty"`
	Tph *int32 `json:"tph,omitempty" xml:"tph,omitempty"`
	Tpm *int32 `json:"tpm,omitempty" xml:"tpm,omitempty"`
	Tps *int32 `json:"tps,omitempty" xml:"tps,omitempty"`
}

func (s ProxyConfigPoliciesTokenRateLimiter) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfigPoliciesTokenRateLimiter) GoString() string {
	return s.String()
}

func (s *ProxyConfigPoliciesTokenRateLimiter) GetTpd() *int32 {
	return s.Tpd
}

func (s *ProxyConfigPoliciesTokenRateLimiter) GetTph() *int32 {
	return s.Tph
}

func (s *ProxyConfigPoliciesTokenRateLimiter) GetTpm() *int32 {
	return s.Tpm
}

func (s *ProxyConfigPoliciesTokenRateLimiter) GetTps() *int32 {
	return s.Tps
}

func (s *ProxyConfigPoliciesTokenRateLimiter) SetTpd(v int32) *ProxyConfigPoliciesTokenRateLimiter {
	s.Tpd = &v
	return s
}

func (s *ProxyConfigPoliciesTokenRateLimiter) SetTph(v int32) *ProxyConfigPoliciesTokenRateLimiter {
	s.Tph = &v
	return s
}

func (s *ProxyConfigPoliciesTokenRateLimiter) SetTpm(v int32) *ProxyConfigPoliciesTokenRateLimiter {
	s.Tpm = &v
	return s
}

func (s *ProxyConfigPoliciesTokenRateLimiter) SetTps(v int32) *ProxyConfigPoliciesTokenRateLimiter {
	s.Tps = &v
	return s
}

func (s *ProxyConfigPoliciesTokenRateLimiter) Validate() error {
	return dara.Validate(s)
}
