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
	// The unique callback URL assigned to the provider. You must add this URL to the list of allowed callback URLs of the identity provider.
	//
	// Format: https://agentidentitydata.{{region}}.aliyuncs.com/oauth2/callback/{{UUID}}
	//
	// If you specify this parameter, the specified value is used.
	//
	// If you do not specify this parameter, the server automatically generates a value.
	//
	// example:
	//
	// https://agentidentitydata.cn-beijing.aliyuncs.com/oauth2/callback/d51171bc-0dae-3219-8e65-6b4cdafa3beb
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// AliyunOAuth2
	//
	// example:
	//
	// AliyunOAuth2
	CredentialProviderVendor *string `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	// The description of the OAuth 2.0 credential provider.
	//
	// example:
	//
	// new example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the credential provider.
	//
	// example:
	//
	// oauth2-provider-aliyun
	OAuth2CredentialProviderName *string `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	// The object that contains provider-specific configurations, such as clientId and clientSecret.
	OAuth2ProviderConfig *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	// The name of the TokenVault.
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
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
