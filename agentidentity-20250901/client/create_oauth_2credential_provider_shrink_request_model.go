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
}

type CreateOAuth2CredentialProviderShrinkRequest struct {
	// example:
	//
	// https://agentidentitydata.cn-beijing.aliyuncs.com/oauth2/callback/d51171bc-0dae-3219-8e65-6b4cdafa3beb
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// example:
	//
	// AliyunOAuth2
	CredentialProviderVendor *string `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	// example:
	//
	// example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// oauth-provider-aliyun
	OAuth2CredentialProviderName *string `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfigShrink   *string `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
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

func (s *CreateOAuth2CredentialProviderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
