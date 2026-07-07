// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuth2CredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackURL(v string) *CreateOAuth2CredentialProviderRequest
	GetCallbackURL() *string
	SetCredentialProviderVendor(v string) *CreateOAuth2CredentialProviderRequest
	GetCredentialProviderVendor() *string
	SetDescription(v string) *CreateOAuth2CredentialProviderRequest
	GetDescription() *string
	SetOAuth2CredentialProviderName(v string) *CreateOAuth2CredentialProviderRequest
	GetOAuth2CredentialProviderName() *string
	SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *CreateOAuth2CredentialProviderRequest
	GetOAuth2ProviderConfig() *OAuth2ProviderConfig
	SetOAuthType(v string) *CreateOAuth2CredentialProviderRequest
	GetOAuthType() *string
	SetTokenVaultName(v string) *CreateOAuth2CredentialProviderRequest
	GetTokenVaultName() *string
}

type CreateOAuth2CredentialProviderRequest struct {
	CallbackURL                  *string               `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	CredentialProviderVendor     *string               `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	Description                  *string               `json:"Description,omitempty" xml:"Description,omitempty"`
	OAuth2CredentialProviderName *string               `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfig         *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	OAuthType                    *string               `json:"OAuthType,omitempty" xml:"OAuthType,omitempty"`
	TokenVaultName               *string               `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s CreateOAuth2CredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuth2CredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateOAuth2CredentialProviderRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *CreateOAuth2CredentialProviderRequest) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *CreateOAuth2CredentialProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOAuth2CredentialProviderRequest) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *CreateOAuth2CredentialProviderRequest) GetOAuth2ProviderConfig() *OAuth2ProviderConfig {
	return s.OAuth2ProviderConfig
}

func (s *CreateOAuth2CredentialProviderRequest) GetOAuthType() *string {
	return s.OAuthType
}

func (s *CreateOAuth2CredentialProviderRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateOAuth2CredentialProviderRequest) SetCallbackURL(v string) *CreateOAuth2CredentialProviderRequest {
	s.CallbackURL = &v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) SetCredentialProviderVendor(v string) *CreateOAuth2CredentialProviderRequest {
	s.CredentialProviderVendor = &v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) SetDescription(v string) *CreateOAuth2CredentialProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) SetOAuth2CredentialProviderName(v string) *CreateOAuth2CredentialProviderRequest {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *CreateOAuth2CredentialProviderRequest {
	s.OAuth2ProviderConfig = v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) SetOAuthType(v string) *CreateOAuth2CredentialProviderRequest {
	s.OAuthType = &v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) SetTokenVaultName(v string) *CreateOAuth2CredentialProviderRequest {
	s.TokenVaultName = &v
	return s
}

func (s *CreateOAuth2CredentialProviderRequest) Validate() error {
	if s.OAuth2ProviderConfig != nil {
		if err := s.OAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
