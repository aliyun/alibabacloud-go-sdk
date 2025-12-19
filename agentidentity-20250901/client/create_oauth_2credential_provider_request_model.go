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
}

type CreateOAuth2CredentialProviderRequest struct {
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
	OAuth2CredentialProviderName *string               `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfig         *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
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

func (s *CreateOAuth2CredentialProviderRequest) Validate() error {
	if s.OAuth2ProviderConfig != nil {
		if err := s.OAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
