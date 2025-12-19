// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOAuth2ProviderConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCustomOAuth2ProviderConfig(v *CustomOAuth2ProviderConfig) *OAuth2ProviderConfig
	GetCustomOAuth2ProviderConfig() *CustomOAuth2ProviderConfig
	SetIncludedOAuth2ProviderConfig(v *IncludedOAuth2ProviderConfig) *OAuth2ProviderConfig
	GetIncludedOAuth2ProviderConfig() *IncludedOAuth2ProviderConfig
}

type OAuth2ProviderConfig struct {
	CustomOAuth2ProviderConfig   *CustomOAuth2ProviderConfig   `json:"CustomOAuth2ProviderConfig,omitempty" xml:"CustomOAuth2ProviderConfig,omitempty"`
	IncludedOAuth2ProviderConfig *IncludedOAuth2ProviderConfig `json:"IncludedOAuth2ProviderConfig,omitempty" xml:"IncludedOAuth2ProviderConfig,omitempty"`
}

func (s OAuth2ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s OAuth2ProviderConfig) GoString() string {
	return s.String()
}

func (s *OAuth2ProviderConfig) GetCustomOAuth2ProviderConfig() *CustomOAuth2ProviderConfig {
	return s.CustomOAuth2ProviderConfig
}

func (s *OAuth2ProviderConfig) GetIncludedOAuth2ProviderConfig() *IncludedOAuth2ProviderConfig {
	return s.IncludedOAuth2ProviderConfig
}

func (s *OAuth2ProviderConfig) SetCustomOAuth2ProviderConfig(v *CustomOAuth2ProviderConfig) *OAuth2ProviderConfig {
	s.CustomOAuth2ProviderConfig = v
	return s
}

func (s *OAuth2ProviderConfig) SetIncludedOAuth2ProviderConfig(v *IncludedOAuth2ProviderConfig) *OAuth2ProviderConfig {
	s.IncludedOAuth2ProviderConfig = v
	return s
}

func (s *OAuth2ProviderConfig) Validate() error {
	if s.CustomOAuth2ProviderConfig != nil {
		if err := s.CustomOAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IncludedOAuth2ProviderConfig != nil {
		if err := s.IncludedOAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
