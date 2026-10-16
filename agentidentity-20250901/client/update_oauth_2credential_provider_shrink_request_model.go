// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOAuth2CredentialProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackURL(v string) *UpdateOAuth2CredentialProviderShrinkRequest
	GetCallbackURL() *string
	SetCredentialProviderVendor(v string) *UpdateOAuth2CredentialProviderShrinkRequest
	GetCredentialProviderVendor() *string
	SetDescription(v string) *UpdateOAuth2CredentialProviderShrinkRequest
	GetDescription() *string
	SetOAuth2CredentialProviderName(v string) *UpdateOAuth2CredentialProviderShrinkRequest
	GetOAuth2CredentialProviderName() *string
	SetOAuth2ProviderConfigShrink(v string) *UpdateOAuth2CredentialProviderShrinkRequest
	GetOAuth2ProviderConfigShrink() *string
	SetTokenVaultName(v string) *UpdateOAuth2CredentialProviderShrinkRequest
	GetTokenVaultName() *string
}

type UpdateOAuth2CredentialProviderShrinkRequest struct {
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
	OAuth2ProviderConfigShrink *string `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	// The name of the TokenVault.
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s UpdateOAuth2CredentialProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOAuth2CredentialProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) GetOAuth2ProviderConfigShrink() *string {
	return s.OAuth2ProviderConfigShrink
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) SetCallbackURL(v string) *UpdateOAuth2CredentialProviderShrinkRequest {
	s.CallbackURL = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) SetCredentialProviderVendor(v string) *UpdateOAuth2CredentialProviderShrinkRequest {
	s.CredentialProviderVendor = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) SetDescription(v string) *UpdateOAuth2CredentialProviderShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) SetOAuth2CredentialProviderName(v string) *UpdateOAuth2CredentialProviderShrinkRequest {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) SetOAuth2ProviderConfigShrink(v string) *UpdateOAuth2CredentialProviderShrinkRequest {
	s.OAuth2ProviderConfigShrink = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) SetTokenVaultName(v string) *UpdateOAuth2CredentialProviderShrinkRequest {
	s.TokenVaultName = &v
	return s
}

func (s *UpdateOAuth2CredentialProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
