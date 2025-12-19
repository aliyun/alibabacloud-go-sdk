// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncludedOAuth2ProviderConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationEndpoint(v string) *IncludedOAuth2ProviderConfig
	GetAuthorizationEndpoint() *string
	SetClientId(v string) *IncludedOAuth2ProviderConfig
	GetClientId() *string
	SetClientSecret(v string) *IncludedOAuth2ProviderConfig
	GetClientSecret() *string
	SetTokenEndpoint(v string) *IncludedOAuth2ProviderConfig
	GetTokenEndpoint() *string
}

type IncludedOAuth2ProviderConfig struct {
	AuthorizationEndpoint *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	ClientId              *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientSecret          *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	TokenEndpoint         *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s IncludedOAuth2ProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s IncludedOAuth2ProviderConfig) GoString() string {
	return s.String()
}

func (s *IncludedOAuth2ProviderConfig) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *IncludedOAuth2ProviderConfig) GetClientId() *string {
	return s.ClientId
}

func (s *IncludedOAuth2ProviderConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *IncludedOAuth2ProviderConfig) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *IncludedOAuth2ProviderConfig) SetAuthorizationEndpoint(v string) *IncludedOAuth2ProviderConfig {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *IncludedOAuth2ProviderConfig) SetClientId(v string) *IncludedOAuth2ProviderConfig {
	s.ClientId = &v
	return s
}

func (s *IncludedOAuth2ProviderConfig) SetClientSecret(v string) *IncludedOAuth2ProviderConfig {
	s.ClientSecret = &v
	return s
}

func (s *IncludedOAuth2ProviderConfig) SetTokenEndpoint(v string) *IncludedOAuth2ProviderConfig {
	s.TokenEndpoint = &v
	return s
}

func (s *IncludedOAuth2ProviderConfig) Validate() error {
	return dara.Validate(s)
}
