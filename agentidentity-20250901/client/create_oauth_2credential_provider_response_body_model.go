// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOAuth2CredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOAuth2CredentialProvider(v *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) *CreateOAuth2CredentialProviderResponseBody
	GetOAuth2CredentialProvider() *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider
	SetRequestId(v string) *CreateOAuth2CredentialProviderResponseBody
	GetRequestId() *string
}

type CreateOAuth2CredentialProviderResponseBody struct {
	OAuth2CredentialProvider *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider `json:"OAuth2CredentialProvider,omitempty" xml:"OAuth2CredentialProvider,omitempty" type:"Struct"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOAuth2CredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuth2CredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOAuth2CredentialProviderResponseBody) GetOAuth2CredentialProvider() *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	return s.OAuth2CredentialProvider
}

func (s *CreateOAuth2CredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOAuth2CredentialProviderResponseBody) SetOAuth2CredentialProvider(v *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) *CreateOAuth2CredentialProviderResponseBody {
	s.OAuth2CredentialProvider = v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBody) SetRequestId(v string) *CreateOAuth2CredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBody) Validate() error {
	if s.OAuth2CredentialProvider != nil {
		if err := s.OAuth2CredentialProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider struct {
	CallbackURL                  *string               `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	CreateTime                   *string               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CredentialProviderArn        *string               `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
	CredentialProviderVendor     *string               `json:"CredentialProviderVendor,omitempty" xml:"CredentialProviderVendor,omitempty"`
	Description                  *string               `json:"Description,omitempty" xml:"Description,omitempty"`
	OAuth2CredentialProviderName *string               `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfig         *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	OAuthType                    *string               `json:"OAuthType,omitempty" xml:"OAuthType,omitempty"`
	TokenVaultName               *string               `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
	UpdateTime                   *string               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) String() string {
	return dara.Prettify(s)
}

func (s CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GoString() string {
	return s.String()
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetDescription() *string {
	return s.Description
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuth2ProviderConfig() *OAuth2ProviderConfig {
	return s.OAuth2ProviderConfig
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuthType() *string {
	return s.OAuthType
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCallbackURL(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CallbackURL = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCreateTime(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CreateTime = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCredentialProviderArn(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CredentialProviderArn = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCredentialProviderVendor(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CredentialProviderVendor = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetDescription(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.Description = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOAuth2CredentialProviderName(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OAuth2ProviderConfig = v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOAuthType(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OAuthType = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetTokenVaultName(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.TokenVaultName = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetUpdateTime(v string) *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.UpdateTime = &v
	return s
}

func (s *CreateOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) Validate() error {
	if s.OAuth2ProviderConfig != nil {
		if err := s.OAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
