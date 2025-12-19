// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomOAuth2ProviderConfig interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *CustomOAuth2ProviderConfig
	GetClientId() *string
	SetClientSecret(v string) *CustomOAuth2ProviderConfig
	GetClientSecret() *string
	SetOAuth2Discovery(v *OAuth2Discovery) *CustomOAuth2ProviderConfig
	GetOAuth2Discovery() *OAuth2Discovery
}

type CustomOAuth2ProviderConfig struct {
	ClientId        *string          `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret    *string          `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	OAuth2Discovery *OAuth2Discovery `json:"OAuth2Discovery,omitempty" xml:"OAuth2Discovery,omitempty"`
}

func (s CustomOAuth2ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s CustomOAuth2ProviderConfig) GoString() string {
	return s.String()
}

func (s *CustomOAuth2ProviderConfig) GetClientId() *string {
	return s.ClientId
}

func (s *CustomOAuth2ProviderConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *CustomOAuth2ProviderConfig) GetOAuth2Discovery() *OAuth2Discovery {
	return s.OAuth2Discovery
}

func (s *CustomOAuth2ProviderConfig) SetClientId(v string) *CustomOAuth2ProviderConfig {
	s.ClientId = &v
	return s
}

func (s *CustomOAuth2ProviderConfig) SetClientSecret(v string) *CustomOAuth2ProviderConfig {
	s.ClientSecret = &v
	return s
}

func (s *CustomOAuth2ProviderConfig) SetOAuth2Discovery(v *OAuth2Discovery) *CustomOAuth2ProviderConfig {
	s.OAuth2Discovery = v
	return s
}

func (s *CustomOAuth2ProviderConfig) Validate() error {
	if s.OAuth2Discovery != nil {
		if err := s.OAuth2Discovery.Validate(); err != nil {
			return err
		}
	}
	return nil
}
