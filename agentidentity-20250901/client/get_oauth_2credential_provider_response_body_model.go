// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOAuth2CredentialProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOAuth2CredentialProvider(v *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) *GetOAuth2CredentialProviderResponseBody
	GetOAuth2CredentialProvider() *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider
	SetRequestId(v string) *GetOAuth2CredentialProviderResponseBody
	GetRequestId() *string
}

type GetOAuth2CredentialProviderResponseBody struct {
	OAuth2CredentialProvider *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider `json:"OAuth2CredentialProvider,omitempty" xml:"OAuth2CredentialProvider,omitempty" type:"Struct"`
	// example:
	//
	// D9A9DC39-61BB-53FD-9ADC-B14884F21038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOAuth2CredentialProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOAuth2CredentialProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetOAuth2CredentialProviderResponseBody) GetOAuth2CredentialProvider() *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	return s.OAuth2CredentialProvider
}

func (s *GetOAuth2CredentialProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOAuth2CredentialProviderResponseBody) SetOAuth2CredentialProvider(v *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) *GetOAuth2CredentialProviderResponseBody {
	s.OAuth2CredentialProvider = v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBody) SetRequestId(v string) *GetOAuth2CredentialProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBody) Validate() error {
	if s.OAuth2CredentialProvider != nil {
		if err := s.OAuth2CredentialProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider struct {
	// example:
	//
	// https://agentidentitydata.cn-beijing.aliyuncs.com/oauth2/callback/d51171bc-0dae-3219-8e65-6b4cdafa3beb
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default/oauth2credentialprovider/oauth-provider-aliyun
	CredentialProviderArn *string `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
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
	// oauth2-provider-aliyun
	OAuth2CredentialProviderName *string               `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	OAuth2ProviderConfig         *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	// example:
	//
	// 2025-12-18T06:19:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) String() string {
	return dara.Prettify(s)
}

func (s GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GoString() string {
	return s.String()
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCallbackURL() *string {
	return s.CallbackURL
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetCredentialProviderVendor() *string {
	return s.CredentialProviderVendor
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetDescription() *string {
	return s.Description
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuth2ProviderConfig() *OAuth2ProviderConfig {
	return s.OAuth2ProviderConfig
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCallbackURL(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CallbackURL = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCreateTime(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CreateTime = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCredentialProviderArn(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CredentialProviderArn = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetCredentialProviderVendor(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.CredentialProviderVendor = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetDescription(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.Description = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOAuth2CredentialProviderName(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OAuth2CredentialProviderName = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOAuth2ProviderConfig(v *OAuth2ProviderConfig) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OAuth2ProviderConfig = v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetUpdateTime(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.UpdateTime = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) Validate() error {
	if s.OAuth2ProviderConfig != nil {
		if err := s.OAuth2ProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
