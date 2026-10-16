// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuth2CredentialProviderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackURL(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetCallbackURL() *string
	SetCredentialProviderVendor(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetCredentialProviderVendor() *string
	SetDescription(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetDescription() *string
	SetOAuth2CredentialProviderName(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetOAuth2CredentialProviderName() *string
	SetOAuth2ProviderConfigShrink(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetOAuth2ProviderConfigShrink() *string
	SetOAuthType(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetOAuthType() *string
	SetTokenVaultName(v string) *CreateOAuth2CredentialProviderShrinkRequest
	GetTokenVaultName() *string
}

type CreateOAuth2CredentialProviderShrinkRequest struct {
	// The unique callback URL assigned to the provider. You must configure this URL in the list of allowed callback URLs of the identity provider.
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
	// The vendor of the provider, such as AliyunOAuth2, AlibabaCloudOAuth2, DingTalkOAuth2, FeiShuOAuth2, or CustomOAuth2.
	//
	// example:
	//
	// AliyunOAuth2
	CredentialProviderVendor *string `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	// The description of the OAuth 2.0 credential provider.
	//
	// example:
	//
	// example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the credential provider.
	//
	// example:
	//
	// oauth-provider-aliyun
	OAuth2CredentialProviderName *string `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	// The object that contains provider-specific configurations, such as clientId and clientSecret.
	OAuth2ProviderConfigShrink *string `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	OAuthType                  *string `json:"OAuthType,omitempty" xml:"OAuthType,omitempty"`
	// The name of the token vault.
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s CreateOAuth2CredentialProviderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuth2CredentialProviderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetOAuth2ProviderConfigShrink() *string {
	return s.OAuth2ProviderConfigShrink
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetOAuthType() *string {
	return s.OAuthType
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetCallbackURL(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.CallbackURL = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetCredentialProviderVendor(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.CredentialProviderVendor = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetDescription(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetOAuth2CredentialProviderName(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetOAuth2ProviderConfigShrink(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.OAuth2ProviderConfigShrink = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetOAuthType(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.OAuthType = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) SetTokenVaultName(v string) *CreateOAuth2CredentialProviderShrinkRequest {
	s.TokenVaultName = &v
	return s
}

func (s *CreateOAuth2CredentialProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
