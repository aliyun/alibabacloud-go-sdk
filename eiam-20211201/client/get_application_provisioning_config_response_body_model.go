// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) *GetApplicationProvisioningConfigResponseBody
	GetApplicationProvisioningConfig() *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig
	SetRequestId(v string) *GetApplicationProvisioningConfigResponseBody
	GetRequestId() *string
}

type GetApplicationProvisioningConfigResponseBody struct {
	// The account synchronization configuration of the application.
	ApplicationProvisioningConfig *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig `json:"ApplicationProvisioningConfig,omitempty" xml:"ApplicationProvisioningConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBody) GetApplicationProvisioningConfig() *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	return s.ApplicationProvisioningConfig
}

func (s *GetApplicationProvisioningConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationProvisioningConfigResponseBody) SetApplicationProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) *GetApplicationProvisioningConfigResponseBody {
	s.ApplicationProvisioningConfig = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBody) SetRequestId(v string) *GetApplicationProvisioningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBody) Validate() error {
	if s.ApplicationProvisioningConfig != nil {
		if err := s.ApplicationProvisioningConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig struct {
	// The application ID.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The configuration parameters for the custom event callback protocol of IDaaS.
	CallbackProvisioningConfig *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig `json:"CallbackProvisioningConfig,omitempty" xml:"CallbackProvisioningConfig,omitempty" type:"Struct"`
	// The rendering mode of the account synchronization page. Valid values:
	//
	// - standard: standard mode.
	//
	// - template: template mode.
	//
	// example:
	//
	// standard
	ConfigOperateMode *string `json:"ConfigOperateMode,omitempty" xml:"ConfigOperateMode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The network endpoint ID.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// The public key endpoint for signature verification of application synchronization callback messages.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk267xxxxx/app_mkv7rgt4d7i4u7zqtzev2mxxxx/provisioning/jwks
	ProvisionJwksEndpoint *string `json:"ProvisionJwksEndpoint,omitempty" xml:"ProvisionJwksEndpoint,omitempty"`
	// Indicates whether to synchronize passwords in IDaaS user event callbacks. Valid values:
	//
	// - true: Passwords are synchronized.
	//
	// - false: Passwords are not synchronized.
	//
	// example:
	//
	// true
	ProvisionPassword *bool `json:"ProvisionPassword,omitempty" xml:"ProvisionPassword,omitempty"`
	// The protocol type for application synchronization. Valid values:
	//
	// - idaas_callback: the custom event callback protocol of IDaaS.
	//
	// - scim2: the System for Cross-domain Identity Management (SCIM) protocol.
	//
	// example:
	//
	// idaas_callback
	ProvisionProtocolType *string `json:"ProvisionProtocolType,omitempty" xml:"ProvisionProtocolType,omitempty"`
	// The configuration parameters for IDaaS SCIM protocol synchronization.
	ScimProvisioningConfig *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig `json:"ScimProvisioningConfig,omitempty" xml:"ScimProvisioningConfig,omitempty" type:"Struct"`
	// The status of the IDaaS account synchronization feature. Valid values:
	//
	// - enabled: The feature is enabled.
	//
	// - disabled: The feature is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetCallbackProvisioningConfig() *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	return s.CallbackProvisioningConfig
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetConfigOperateMode() *string {
	return s.ConfigOperateMode
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetProvisionJwksEndpoint() *string {
	return s.ProvisionJwksEndpoint
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetProvisionPassword() *bool {
	return s.ProvisionPassword
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetProvisionProtocolType() *string {
	return s.ProvisionProtocolType
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetScimProvisioningConfig() *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	return s.ScimProvisioningConfig
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GetStatus() *string {
	return s.Status
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetApplicationId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetCallbackProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.CallbackProvisioningConfig = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetConfigOperateMode(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ConfigOperateMode = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetInstanceId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetNetworkAccessEndpointId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetProvisionJwksEndpoint(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ProvisionJwksEndpoint = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetProvisionPassword(v bool) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ProvisionPassword = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetProvisionProtocolType(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ProvisionProtocolType = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetScimProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ScimProvisioningConfig = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetStatus(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.Status = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) Validate() error {
	if s.CallbackProvisioningConfig != nil {
		if err := s.CallbackProvisioningConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ScimProvisioningConfig != nil {
		if err := s.ScimProvisioningConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig struct {
	// The destination address where the application receives IDaaS event callbacks.
	//
	// example:
	//
	// https://example.com/event/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The symmetric key for encrypting and decrypting IDaaS event callbacks. This key uses the AES-256 algorithm and is in hexadecimal format.
	//
	// example:
	//
	// 1adfdfdfd******111
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// Indicates whether to encrypt IDaaS event callback messages. Valid values:
	//
	// - true: The messages are encrypted.
	//
	// - false: The messages are not encrypted and are transmitted in plaintext.
	//
	// example:
	//
	// true
	EncryptRequired *bool `json:"EncryptRequired,omitempty" xml:"EncryptRequired,omitempty"`
	// The list of message types for the IDaaS event callback listener.
	ListenEventScopes []*string `json:"ListenEventScopes,omitempty" xml:"ListenEventScopes,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) GetEncryptRequired() *bool {
	return s.EncryptRequired
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) GetListenEventScopes() []*string {
	return s.ListenEventScopes
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetCallbackUrl(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.CallbackUrl = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetEncryptKey(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.EncryptKey = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetEncryptRequired(v bool) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.EncryptRequired = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetListenEventScopes(v []*string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.ListenEventScopes = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) Validate() error {
	return dara.Validate(s)
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig struct {
	// The configuration parameters for SCIM protocol synchronization.
	AuthnConfiguration *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration `json:"AuthnConfiguration,omitempty" xml:"AuthnConfiguration,omitempty" type:"Struct"`
	// The scope of a full push for the SCIM protocol. Valid values:
	//
	// - urn:alibaba:idaas:app:scim:User:PUSH: performs a full synchronization of users.
	FullPushScopes []*string `json:"FullPushScopes,omitempty" xml:"FullPushScopes,omitempty" type:"Repeated"`
	// The operations on the target resource for the SCIM protocol. Valid values:
	//
	// - urn:alibaba:idaas:app:scim:User:CREATE: creates an account.
	//
	// - urn:alibaba:idaas:app:scim:User:UPDATE: updates an account.
	//
	// - urn:alibaba:idaas:app:scim:User:DELETE: deletes an account.
	ProvisioningActions []*string `json:"ProvisioningActions,omitempty" xml:"ProvisioningActions,omitempty" type:"Repeated"`
	// The base URL where the application receives IDaaS SCIM protocol synchronization.
	//
	// example:
	//
	// https://example.com/scim
	ScimBaseUrl *string `json:"ScimBaseUrl,omitempty" xml:"ScimBaseUrl,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) GetAuthnConfiguration() *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	return s.AuthnConfiguration
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) GetFullPushScopes() []*string {
	return s.FullPushScopes
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) GetProvisioningActions() []*string {
	return s.ProvisioningActions
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) GetScimBaseUrl() *string {
	return s.ScimBaseUrl
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetAuthnConfiguration(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.AuthnConfiguration = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetFullPushScopes(v []*string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.FullPushScopes = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetProvisioningActions(v []*string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.ProvisioningActions = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetScimBaseUrl(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.ScimBaseUrl = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) Validate() error {
	if s.AuthnConfiguration != nil {
		if err := s.AuthnConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration struct {
	// The authorization mode for the SCIM protocol interface. Valid values:
	//
	// - oauth2: OAuth2 mode.
	//
	// example:
	//
	// oauth2
	AuthnMode *string `json:"AuthnMode,omitempty" xml:"AuthnMode,omitempty"`
	// The authorization configuration parameters. Usage:
	//
	// - If GrantType is set to client_credentials, you must specify ClientId, ClientSecret, and AuthnMethod.
	//
	// - If GrantType is set to bearer_token, you must specify AccessToken.
	AuthnParam *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// The authorization mode for the SCIM protocol. Valid values:
	//
	// - client_credentials: client credentials mode.
	//
	// - bearer_token: bearer token mode.
	//
	// example:
	//
	// bearer_token
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) GetAuthnMode() *string {
	return s.AuthnMode
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) GetAuthnParam() *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	return s.AuthnParam
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) GetGrantType() *string {
	return s.GrantType
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) SetAuthnMode(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	s.AuthnMode = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) SetAuthnParam(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	s.AuthnParam = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) SetGrantType(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	s.GrantType = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) Validate() error {
	if s.AuthnParam != nil {
		if err := s.AuthnParam.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam struct {
	// The access token. This parameter is required when GrantType is set to bearer_token.
	//
	// example:
	//
	// k52x2ru63rlkflina5utgkxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The authentication mode for the SCIM protocol. Valid values:
	//
	// - client_secret_basic: The key is passed in the request header.
	//
	// - client_secret_post: The key is passed in the request body.
	//
	// example:
	//
	// client_secret_basic
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The token endpoint.
	//
	// example:
	//
	// https://www.example.com/oauth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GetAuthnMethod() *string {
	return s.AuthnMethod
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GetClientId() *string {
	return s.ClientId
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetAccessToken(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AccessToken = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetAuthnMethod(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientSecret(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetTokenEndpoint(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.TokenEndpoint = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) Validate() error {
	return dara.Validate(s)
}
