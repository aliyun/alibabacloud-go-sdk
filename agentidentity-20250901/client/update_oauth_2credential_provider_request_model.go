// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOAuth2CredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackURL(v string) *UpdateOAuth2CredentialProviderRequest
	GetCallbackURL() *string
	SetCredentialProviderVendor(v string) *UpdateOAuth2CredentialProviderRequest
	GetCredentialProviderVendor() *string
	SetDescription(v string) *UpdateOAuth2CredentialProviderRequest
	GetDescription() *string
	SetOAuth2CredentialProviderName(v string) *UpdateOAuth2CredentialProviderRequest
	GetOAuth2CredentialProviderName() *string
	SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *UpdateOAuth2CredentialProviderRequest
	GetOAuth2ProviderConfig() *OAuth2ProviderConfig
	SetTokenVaultName(v string) *UpdateOAuth2CredentialProviderRequest
	GetTokenVaultName() *string
}

type UpdateOAuth2CredentialProviderRequest struct {
	CallbackURL                  *string               `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	CredentialProviderVendor     *string               `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	Description                  *string               `json:"Description,omitempty" xml:"Description,omitempty"`
	OAuth2CredentialProviderName *string               `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfig         *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	TokenVaultName               *string               `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s UpdateOAuth2CredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOAuth2CredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateOAuth2CredentialProviderRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *UpdateOAuth2CredentialProviderRequest) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *UpdateOAuth2CredentialProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOAuth2CredentialProviderRequest) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *UpdateOAuth2CredentialProviderRequest) GetOAuth2ProviderConfig() *OAuth2ProviderConfig {
	return s.OAuth2ProviderConfig
}

func (s *UpdateOAuth2CredentialProviderRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *UpdateOAuth2CredentialProviderRequest) SetCallbackURL(v string) *UpdateOAuth2CredentialProviderRequest {
	s.CallbackURL = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderRequest) SetCredentialProviderVendor(v string) *UpdateOAuth2CredentialProviderRequest {
	s.CredentialProviderVendor = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderRequest) SetDescription(v string) *UpdateOAuth2CredentialProviderRequest {
	s.Description = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderRequest) SetOAuth2CredentialProviderName(v string) *UpdateOAuth2CredentialProviderRequest {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderRequest) SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *UpdateOAuth2CredentialProviderRequest {
	s.OAuth2ProviderConfig = v
	return s
}

func (s *UpdateOAuth2CredentialProviderRequest) SetTokenVaultName(v string) *UpdateOAuth2CredentialProviderRequest {
	s.TokenVaultName = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderRequest) Validate() error {
	if s.OAuth2ProviderConfig != nil {
		if err := s.OAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
