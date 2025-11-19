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
	Cache            *bool                           `json:"cache,omitempty" xml:"cache,omitempty"`
	ConcurrencyLimit *int32                          `json:"concurrencyLimit,omitempty" xml:"concurrencyLimit,omitempty"`
	Fallbacks        []*ProxyConfigPoliciesFallbacks `json:"fallbacks,omitempty" xml:"fallbacks,omitempty" type:"Repeated"`
	NumRetries       *int32                          `json:"numRetries,omitempty" xml:"numRetries,omitempty"`
	RequestTimeout   *int32                          `json:"requestTimeout,omitempty" xml:"requestTimeout,omitempty"`
}

func (s ProxyConfigPolicies) String() string {
	return dara.Prettify(s)
}

func (s ProxyConfigPolicies) GoString() string {
	return s.String()
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

func (s *ProxyConfigPolicies) Validate() error {
	if s.Fallbacks != nil {
		for _, item := range s.Fallbacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
