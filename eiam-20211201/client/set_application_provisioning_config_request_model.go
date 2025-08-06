// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationProvisioningConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *SetApplicationProvisioningConfigRequest
	GetApplicationId() *string
	SetCallbackProvisioningConfig(v *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) *SetApplicationProvisioningConfigRequest
	GetCallbackProvisioningConfig() *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig
	SetInstanceId(v string) *SetApplicationProvisioningConfigRequest
	GetInstanceId() *string
	SetNetworkAccessEndpointId(v string) *SetApplicationProvisioningConfigRequest
	GetNetworkAccessEndpointId() *string
	SetProvisionPassword(v bool) *SetApplicationProvisioningConfigRequest
	GetProvisionPassword() *bool
	SetProvisionProtocolType(v string) *SetApplicationProvisioningConfigRequest
	GetProvisionProtocolType() *string
	SetScimProvisioningConfig(v *SetApplicationProvisioningConfigRequestScimProvisioningConfig) *SetApplicationProvisioningConfigRequest
	GetScimProvisioningConfig() *SetApplicationProvisioningConfigRequestScimProvisioningConfig
}

type SetApplicationProvisioningConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The configuration of event callback synchronization. This parameter is required when the ProvisionProtocolType parameter is set to idaas_callback.
	CallbackProvisioningConfig *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig `json:"CallbackProvisioningConfig,omitempty" xml:"CallbackProvisioningConfig,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// Specifies whether to synchronize the password in IDaaS user event callbacks. Valid values:
	//
	// 	- true: synchronize the password.
	//
	// 	- false: do not synchronize the password.
	//
	// example:
	//
	// true
	ProvisionPassword *bool `json:"ProvisionPassword,omitempty" xml:"ProvisionPassword,omitempty"`
	// The synchronization protocol type of the application. Valid values:
	//
	// 	- idaas_callback: custom event callback protocol of IDaaS.
	//
	// 	- scim2: System for Cross-domain Identity Management (SCIM) protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_callback
	ProvisionProtocolType *string `json:"ProvisionProtocolType,omitempty" xml:"ProvisionProtocolType,omitempty"`
	// The configuration of SCIM-based IDaaS synchronization. This parameter is required when the ProvisionProtocolType parameter is set to scim2.
	ScimProvisioningConfig *SetApplicationProvisioningConfigRequestScimProvisioningConfig `json:"ScimProvisioningConfig,omitempty" xml:"ScimProvisioningConfig,omitempty" type:"Struct"`
}

func (s SetApplicationProvisioningConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *SetApplicationProvisioningConfigRequest) GetCallbackProvisioningConfig() *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	return s.CallbackProvisioningConfig
}

func (s *SetApplicationProvisioningConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetApplicationProvisioningConfigRequest) GetNetworkAccessEndpointId() *string {
	return s.NetworkAccessEndpointId
}

func (s *SetApplicationProvisioningConfigRequest) GetProvisionPassword() *bool {
	return s.ProvisionPassword
}

func (s *SetApplicationProvisioningConfigRequest) GetProvisionProtocolType() *string {
	return s.ProvisionProtocolType
}

func (s *SetApplicationProvisioningConfigRequest) GetScimProvisioningConfig() *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	return s.ScimProvisioningConfig
}

func (s *SetApplicationProvisioningConfigRequest) SetApplicationId(v string) *SetApplicationProvisioningConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetCallbackProvisioningConfig(v *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) *SetApplicationProvisioningConfigRequest {
	s.CallbackProvisioningConfig = v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetInstanceId(v string) *SetApplicationProvisioningConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetNetworkAccessEndpointId(v string) *SetApplicationProvisioningConfigRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetProvisionPassword(v bool) *SetApplicationProvisioningConfigRequest {
	s.ProvisionPassword = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetProvisionProtocolType(v string) *SetApplicationProvisioningConfigRequest {
	s.ProvisionProtocolType = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetScimProvisioningConfig(v *SetApplicationProvisioningConfigRequestScimProvisioningConfig) *SetApplicationProvisioningConfigRequest {
	s.ScimProvisioningConfig = v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) Validate() error {
	return dara.Validate(s)
}

type SetApplicationProvisioningConfigRequestCallbackProvisioningConfig struct {
	// The URL that the application uses to receive IDaaS event callbacks.
	//
	// example:
	//
	// https://example.com/event/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The symmetric key for IDaaS event callbacks. The key is an AES-256 encryption key in the HEX format.
	//
	// example:
	//
	// ad3b248**************************b3561a73d7
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// Specifies whether to encrypt IDaaS event callback messages. Valid values:
	//
	// 	- true: encrypt the messages.
	//
	// 	- false: transmit the messages in plaintext.
	//
	// example:
	//
	// true
	EncryptRequired *bool `json:"EncryptRequired,omitempty" xml:"EncryptRequired,omitempty"`
	// The list of types of IDaaS event callback messages that are supported by the listener.
	ListenEventScopes []*string `json:"ListenEventScopes,omitempty" xml:"ListenEventScopes,omitempty" type:"Repeated"`
}

func (s SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) GetEncryptRequired() *bool {
	return s.EncryptRequired
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) GetListenEventScopes() []*string {
	return s.ListenEventScopes
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetCallbackUrl(v string) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.CallbackUrl = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetEncryptKey(v string) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.EncryptKey = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetEncryptRequired(v bool) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.EncryptRequired = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetListenEventScopes(v []*string) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.ListenEventScopes = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) Validate() error {
	return dara.Validate(s)
}

type SetApplicationProvisioningConfigRequestScimProvisioningConfig struct {
	// The configuration parameters related to SCIM-based synchronization.
	AuthnConfiguration *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration `json:"AuthnConfiguration,omitempty" xml:"AuthnConfiguration,omitempty" type:"Struct"`
	// The full synchronization scope of the SCIM protocol. Valid value:
	//
	// 	- urn:alibaba:idaas:app:scim:User:PUSH: full account data synchronization.
	FullPushScopes []*string `json:"FullPushScopes,omitempty" xml:"FullPushScopes,omitempty" type:"Repeated"`
	// The resource operations of the SCIM protocol. Valid values:
	//
	// 	- urn:alibaba:idaas:app:scim:User:CREATE: account creation.
	//
	// 	- urn:alibaba:idaas:app:scim:User:UPDATE: account update.
	//
	// 	- urn:alibaba:idaas:app:scim:User:DELETE: account deletion.
	ProvisioningActions []*string `json:"ProvisioningActions,omitempty" xml:"ProvisioningActions,omitempty" type:"Repeated"`
	// The base URL that the application uses to receive the SCIM protocol for IDaaS synchronization.
	//
	// example:
	//
	// https://example.com/scim
	ScimBaseUrl *string `json:"ScimBaseUrl,omitempty" xml:"ScimBaseUrl,omitempty"`
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfig) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) GetAuthnConfiguration() *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	return s.AuthnConfiguration
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) GetFullPushScopes() []*string {
	return s.FullPushScopes
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) GetProvisioningActions() []*string {
	return s.ProvisioningActions
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) GetScimBaseUrl() *string {
	return s.ScimBaseUrl
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetAuthnConfiguration(v *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.AuthnConfiguration = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetFullPushScopes(v []*string) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.FullPushScopes = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetProvisioningActions(v []*string) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.ProvisioningActions = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetScimBaseUrl(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.ScimBaseUrl = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) Validate() error {
	return dara.Validate(s)
}

type SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration struct {
	// The authentication mode of the SCIM protocol. Valid value:
	//
	// 	- oauth2: OAuth2.0 mode.
	//
	// example:
	//
	// oauth2
	AuthnMode *string `json:"AuthnMode,omitempty" xml:"AuthnMode,omitempty"`
	// The configuration parameters related to authorization.
	//
	// 	- If the GrantType parameter is set to client_credentials, you can set the configuration parameters ClientId, ClientSecret, and AuthnMethod.
	//
	// 	- If the GrantType parameter is set to bearer_token, you can set the configuration parameter AccessToken.
	AuthnParam *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// The grant type of the SCIM protocol. Valid values:
	//
	// 	- client_credentials: client mode.
	//
	// 	- bearer_token: key mode.
	//
	// example:
	//
	// bearer_token
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) GetAuthnMode() *string {
	return s.AuthnMode
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) GetAuthnParam() *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	return s.AuthnParam
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) GetGrantType() *string {
	return s.GrantType
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) SetAuthnMode(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	s.AuthnMode = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) SetAuthnParam(v *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	s.AuthnParam = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) SetGrantType(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	s.GrantType = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) Validate() error {
	return dara.Validate(s)
}

type SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam struct {
	// The access token. If the GrantType parameter is set to bearer_token, you can set this parameter.
	//
	// example:
	//
	// k52x2ru63rlkflina5utgkxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The authentication mode of the SCIM protocol. Valid values:
	//
	// 	- client_secret_basic: The client secret is passed in the request header.
	//
	// 	- client_secret_post: The client secret is passed in the request body.
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

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GetAccessToken() *string {
	return s.AccessToken
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GetAuthnMethod() *string {
	return s.AuthnMethod
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GetClientId() *string {
	return s.ClientId
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetAccessToken(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AccessToken = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetAuthnMethod(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientId(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientSecret(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetTokenEndpoint(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.TokenEndpoint = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) Validate() error {
	return dara.Validate(s)
}
