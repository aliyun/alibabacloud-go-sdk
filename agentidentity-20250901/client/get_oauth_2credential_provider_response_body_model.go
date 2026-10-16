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
	// The credential provider.
	OAuth2CredentialProvider *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider `json:"OAuth2CredentialProvider,omitempty" xml:"OAuth2CredentialProvider,omitempty" type:"Struct"`
	// The unique identifier of the request.
	//
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
	// The unique callback URL assigned to the provider, which must be configured in the list of allowed callback URLs of the identity provider.
	//
	// Format: https://agentidentitydata.{{region}}.aliyuncs.com/oauth2/callback/{{UUID}}
	//
	// If a value is specified, the specified value is used.
	//
	// If no value is specified, the server automatically generates one.
	//
	// example:
	//
	// https://agentidentitydata.cn-beijing.aliyuncs.com/oauth2/callback/d51171bc-0dae-3219-8e65-6b4cdafa3beb
	CallbackURL *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-12-18T06:19:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ARN of the credential provider.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:tokenvault/default/oauth2credentialprovider/oauth-provider-aliyun
	CredentialProviderArn *string `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
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
	// The IDaaS instance ID.
	//
	// example:
	//
	// idaas_xxxx
	IDaaSInstanceId *string `json:"IDaaSInstanceId,omitempty" xml:"IDaaSInstanceId,omitempty"`
	// The name of the credential provider.
	//
	// example:
	//
	// oauth2-provider-aliyun
	OAuth2CredentialProviderName *string `json:"OAuth2CredentialProviderName,omitempty" xml:"OAuth2CredentialProviderName,omitempty"`
	// The object that contains provider-specific configurations, such as clientId and clientSecret.
	OAuth2ProviderConfig *OAuth2ProviderConfig `json:"OAuth2ProviderConfig,omitempty" xml:"OAuth2ProviderConfig,omitempty"`
	OAuthType            *string               `json:"OAuthType,omitempty" xml:"OAuthType,omitempty"`
	// The ID of the outbound application.
	//
	// example:
	//
	// app_xxxxx
	OutboundApplicationId *string `json:"OutboundApplicationId,omitempty" xml:"OutboundApplicationId,omitempty"`
	// The outbound audience.
	//
	// example:
	//
	// agent-xxxx
	OutboundAudience *string `json:"OutboundAudience,omitempty" xml:"OutboundAudience,omitempty"`
	// The client ID of the outbound application.
	//
	// example:
	//
	// client_xxxx
	OutboundClientId *string `json:"OutboundClientId,omitempty" xml:"OutboundClientId,omitempty"`
	SourcePlatform   *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// The name of the TokenVault.
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
	// The update time.
	//
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

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetIDaaSInstanceId() *string {
	return s.IDaaSInstanceId
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuth2CredentialProviderName() *string {
	return s.OAuth2CredentialProviderName
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuth2ProviderConfig() *OAuth2ProviderConfig {
	return s.OAuth2ProviderConfig
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOAuthType() *string {
	return s.OAuthType
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOutboundApplicationId() *string {
	return s.OutboundApplicationId
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOutboundAudience() *string {
	return s.OutboundAudience
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetOutboundClientId() *string {
	return s.OutboundClientId
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) GetTokenVaultName() *string {
	return s.TokenVaultName
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

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetIDaaSInstanceId(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.IDaaSInstanceId = &v
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

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOAuthType(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OAuthType = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOutboundApplicationId(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OutboundApplicationId = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOutboundAudience(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OutboundAudience = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetOutboundClientId(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.OutboundClientId = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetSourcePlatform(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.SourcePlatform = &v
	return s
}

func (s *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider) SetTokenVaultName(v string) *GetOAuth2CredentialProviderResponseBodyOAuth2CredentialProvider {
	s.TokenVaultName = &v
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
