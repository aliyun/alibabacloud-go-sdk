// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *AiServiceConfig
	GetAddress() *string
	SetApiKeys(v []*string) *AiServiceConfig
	GetApiKeys() []*string
	SetEnableHealthCheck(v bool) *AiServiceConfig
	GetEnableHealthCheck() *bool
	SetProtocols(v []*string) *AiServiceConfig
	GetProtocols() []*string
	SetProvider(v string) *AiServiceConfig
	GetProvider() *string
}

type AiServiceConfig struct {
	Address           *string   `json:"address,omitempty" xml:"address,omitempty"`
	ApiKeys           []*string `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
	EnableHealthCheck *bool     `json:"enableHealthCheck,omitempty" xml:"enableHealthCheck,omitempty"`
	Protocols         []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	Provider          *string   `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s AiServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s AiServiceConfig) GoString() string {
	return s.String()
}

func (s *AiServiceConfig) GetAddress() *string {
	return s.Address
}

func (s *AiServiceConfig) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *AiServiceConfig) GetEnableHealthCheck() *bool {
	return s.EnableHealthCheck
}

func (s *AiServiceConfig) GetProtocols() []*string {
	return s.Protocols
}

func (s *AiServiceConfig) GetProvider() *string {
	return s.Provider
}

func (s *AiServiceConfig) SetAddress(v string) *AiServiceConfig {
	s.Address = &v
	return s
}

func (s *AiServiceConfig) SetApiKeys(v []*string) *AiServiceConfig {
	s.ApiKeys = v
	return s
}

func (s *AiServiceConfig) SetEnableHealthCheck(v bool) *AiServiceConfig {
	s.EnableHealthCheck = &v
	return s
}

func (s *AiServiceConfig) SetProtocols(v []*string) *AiServiceConfig {
	s.Protocols = v
	return s
}

func (s *AiServiceConfig) SetProvider(v string) *AiServiceConfig {
	s.Provider = &v
	return s
}

func (s *AiServiceConfig) Validate() error {
	return dara.Validate(s)
}
