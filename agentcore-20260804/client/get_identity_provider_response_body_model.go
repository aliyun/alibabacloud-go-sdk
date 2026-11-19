// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderResponseBodyData) *GetIdentityProviderResponseBody
	GetData() *GetIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The binding details of the external identity provider.
	Data *GetIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. Returns an error description when the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderResponseBody) GetData() *GetIdentityProviderResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderResponseBody) SetCode(v string) *GetIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetData(v *GetIdentityProviderResponseBodyData) *GetIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetMessage(v string) *GetIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetRequestId(v string) *GetIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderResponseBody) SetSuccess(v bool) *GetIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderResponseBodyData struct {
	// The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM queries and preparations return the actual received status and URLs.
	Configuration *GetIdentityProviderResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The creation time in UTC format based on RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The event subscription callback URL. You must configure this URL in the application of the external identity provider to receive organization change events. An empty string is returned when the user pool is not yet activated.
	//
	// example:
	//
	// http://auth.cn-hangzhou.agentteams.aliyuncs.com/roa/dingtalk/event/up-123456
	EventSubscriptionCallbackUrl *string `json:"eventSubscriptionCallbackUrl,omitempty" xml:"eventSubscriptionCallbackUrl,omitempty"`
	// The type of the external identity provider. Valid values: DingTalk and Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The logon callback URL. You must configure this URL in the application of the external identity provider. An empty string is returned when the user pool is not yet activated.
	//
	// example:
	//
	// https://signin-cn-hangzhou.aliyunagentid.com/up-123456/dingtalk/callback
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// Indicates whether workspace users are allowed to log on through this external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. Application secret configurations are not returned.
	Metadata *GetIdentityProviderResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The security summary of the expected configuration. It does not contain credentials. The original SAML fields and meanings are retained.
	Oidc *GetIdentityProviderResponseBodyDataOidc `json:"oidc,omitempty" xml:"oidc,omitempty" type:"Struct"`
	// The security summary queried in real time from AgentIdentity. It does not contain secrets, private keys, or JIT fields. It can be empty when it is not configured or during unbinding.
	OidcConfiguration *GetIdentityProviderResponseBodyDataOidcConfiguration `json:"oidcConfiguration,omitempty" xml:"oidcConfiguration,omitempty" type:"Struct"`
	// Currently, only SAML is supported. This parameter may not be returned when no valid configuration is saved. Placeholder initialization with OIDC is not supported.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The input supports either XML or complete manual entry. Queries return standardized configurations and do not echo the original XML. This parameter can be omitted when not configured.
	Saml *GetIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The status. Valid values:
	//
	// - CONFIGURED: The configuration is accepted and waiting for the user pool to be activated.
	//
	// - SYNCING: Organization members are being synchronized.
	//
	// - SYNCED: The synchronization of organization members is complete.
	//
	// - READY: The binding is in effect.
	//
	// - SYNC_FAILED: The synchronization of organization members failed.
	//
	// - UPDATING: The configuration is being updated.
	//
	// - UPDATE_FAILED: The configuration update failed.
	//
	// - DISCONNECTING: The binding is being disconnected.
	//
	// - DISCONNECT_FAILED: The disconnection failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether organization member synchronization is enabled. When enabled, organization members are synchronized as workspace users by this external identity provider.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The synchronization scope defined by department. If omitted during creation, the default value is AllAuthorized, and the first vendor synchronization is determined by syncOnBind. If omitted or set to null during an update, the original value is retained. Updating and saving the scope does not automatically trigger a full synchronization.
	SyncScope *GetIdentityProviderResponseBodyDataSyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
	// The time of the last modification in UTC format based on RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyData) GetConfiguration() *GetIdentityProviderResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *GetIdentityProviderResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetIdentityProviderResponseBodyData) GetEventSubscriptionCallbackUrl() *string {
	return s.EventSubscriptionCallbackUrl
}

func (s *GetIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetIdentityProviderResponseBodyData) GetLoginCallbackUrl() *string {
	return s.LoginCallbackUrl
}

func (s *GetIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *GetIdentityProviderResponseBodyData) GetMetadata() *GetIdentityProviderResponseBodyDataMetadata {
	return s.Metadata
}

func (s *GetIdentityProviderResponseBodyData) GetOidc() *GetIdentityProviderResponseBodyDataOidc {
	return s.Oidc
}

func (s *GetIdentityProviderResponseBodyData) GetOidcConfiguration() *GetIdentityProviderResponseBodyDataOidcConfiguration {
	return s.OidcConfiguration
}

func (s *GetIdentityProviderResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetIdentityProviderResponseBodyData) GetSaml() *GetIdentityProviderResponseBodyDataSaml {
	return s.Saml
}

func (s *GetIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *GetIdentityProviderResponseBodyData) GetSyncScope() *GetIdentityProviderResponseBodyDataSyncScope {
	return s.SyncScope
}

func (s *GetIdentityProviderResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetIdentityProviderResponseBodyData) SetConfiguration(v *GetIdentityProviderResponseBodyDataConfiguration) *GetIdentityProviderResponseBodyData {
	s.Configuration = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetCreatedAt(v string) *GetIdentityProviderResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetEventSubscriptionCallbackUrl(v string) *GetIdentityProviderResponseBodyData {
	s.EventSubscriptionCallbackUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *GetIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetLoginCallbackUrl(v string) *GetIdentityProviderResponseBodyData {
	s.LoginCallbackUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *GetIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetMetadata(v *GetIdentityProviderResponseBodyDataMetadata) *GetIdentityProviderResponseBodyData {
	s.Metadata = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetOidc(v *GetIdentityProviderResponseBodyDataOidc) *GetIdentityProviderResponseBodyData {
	s.Oidc = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetOidcConfiguration(v *GetIdentityProviderResponseBodyDataOidcConfiguration) *GetIdentityProviderResponseBodyData {
	s.OidcConfiguration = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetProtocol(v string) *GetIdentityProviderResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetSaml(v *GetIdentityProviderResponseBodyDataSaml) *GetIdentityProviderResponseBodyData {
	s.Saml = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetStatus(v string) *GetIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *GetIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetSyncScope(v *GetIdentityProviderResponseBodyDataSyncScope) *GetIdentityProviderResponseBodyData {
	s.SyncScope = v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetUpdatedAt(v string) *GetIdentityProviderResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) SetWorkspaceId(v string) *GetIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyData) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	if s.Oidc != nil {
		if err := s.Oidc.Validate(); err != nil {
			return err
		}
	}
	if s.OidcConfiguration != nil {
		if err := s.OidcConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Saml != nil {
		if err := s.Saml.Validate(); err != nil {
			return err
		}
	}
	if s.SyncScope != nil {
		if err := s.SyncScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderResponseBodyDataConfiguration struct {
	// The SAML binding type. Valid values: HTTP-Redirect and HTTP-POST. This parameter is required when manually filled in. Queries return short names.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// This parameter is required when manually filled in. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations do not support replacement. It is different from the Entity ID queried by the service provider (SP).
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required when manually filled in. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual switch value is Enabled or Disabled, which is not equivalent to the expected loginEnabled value saved in AC.
	//
	// example:
	//
	// Enabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
}

func (s GetIdentityProviderResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) SetBindingType(v string) *GetIdentityProviderResponseBodyDataConfiguration {
	s.BindingType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) SetEntityId(v string) *GetIdentityProviderResponseBodyDataConfiguration {
	s.EntityId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) SetLoginUrl(v string) *GetIdentityProviderResponseBodyDataConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) SetSsoStatus(v string) *GetIdentityProviderResponseBodyDataConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyDataMetadata struct {
	// The App ID of the Lark application. This parameter is required when the binding type is Feishu.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The AppKey of the DingTalk application. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// The SAML Binding URI in the historical metadata snapshot. For CustomSSO updates, use saml.bindingType.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect
	Binding *string `json:"binding,omitempty" xml:"binding,omitempty"`
	// The certificate string in the historical metadata snapshot. To edit the configuration, use saml.x509Certificates and do not submit the old metadata.
	//
	// example:
	//
	// MIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQELBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoXDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8oGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYNW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbYEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Okkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0EBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUWld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImcaWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTqP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjocJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdVjfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5
	Certificates *string `json:"certificates,omitempty" xml:"certificates,omitempty"`
	// The CorpId of the DingTalk enterprise. This parameter is required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// This parameter is required when manually filled in. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations do not support replacement. It is different from the Entity ID queried by the service provider (SP).
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required when manually filled in. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
}

func (s GetIdentityProviderResponseBodyDataMetadata) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetAppId() *string {
	return s.AppId
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetAppKey() *string {
	return s.AppKey
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetBinding() *string {
	return s.Binding
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetCertificates() *string {
	return s.Certificates
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetEntityId() *string {
	return s.EntityId
}

func (s *GetIdentityProviderResponseBodyDataMetadata) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetAppId(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.AppId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetAppKey(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.AppKey = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetBinding(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.Binding = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetCertificates(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.Certificates = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetCorpId(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.CorpId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetEntityId(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.EntityId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) SetLoginUrl(v string) *GetIdentityProviderResponseBodyDataMetadata {
	s.LoginUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataMetadata) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyDataOidc struct {
	// The absolute HTTPS URL required in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/authorize
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" xml:"authorizationEndpoint,omitempty"`
	// The default value is client_secret_post. If the value is none, PKCE must be enabled.
	//
	// example:
	//
	// client_secret_post
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty" xml:"clientAuthMethod,omitempty"`
	// The client identifier registered in the external IdP, which is different from the Matrix client. In-place replacement is not supported after configuration.
	//
	// example:
	//
	// example-client
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The absolute HTTPS discovery endpoint. It is mutually exclusive with manually entered endpoints.
	//
	// example:
	//
	// https://idp.example.com/.well-known/openid-configuration
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" xml:"discoveryEndpoint,omitempty"`
	// The absolute HTTPS issuer of the external IdP. In-place replacement is not supported after configuration.
	//
	// example:
	//
	// https://idp.example.com
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// The absolute HTTPS URL required in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/jwks
	JwksUri *string `json:"jwksUri,omitempty" xml:"jwksUri,omitempty"`
	// Specifies whether PKCE is enabled. Default value: true. It must be set to true when no client authentication is used.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"pkceRequired,omitempty" xml:"pkceRequired,omitempty"`
	// An array of strings that must contain openid. The default values are openid, profile, and email.
	//
	// example:
	//
	// ["openid","profile","email"]
	Scopes []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// The claim name in the ID token used to match the SCIM externalId. If omitted during initial creation, the default value sub is used. If omitted during an update, the saved value is retained. For Azure, you can enter oid. The value must be non-empty and cannot contain whitespace characters. When you query oidc, the saved value is returned. When you query oidcConfiguration, the actual upstream value is returned, which may be empty.
	//
	// example:
	//
	// oid
	SsoMappingClaim *string `json:"ssoMappingClaim,omitempty" xml:"ssoMappingClaim,omitempty"`
	// The absolute HTTPS URL required in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/token
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" xml:"tokenEndpoint,omitempty"`
	// The optional absolute HTTPS URL in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/userinfo
	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" xml:"userinfoEndpoint,omitempty"`
}

func (s GetIdentityProviderResponseBodyDataOidc) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataOidc) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetClientAuthMethod() *string {
	return s.ClientAuthMethod
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetClientId() *string {
	return s.ClientId
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetDiscoveryEndpoint() *string {
	return s.DiscoveryEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetIssuer() *string {
	return s.Issuer
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetJwksUri() *string {
	return s.JwksUri
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetScopes() []*string {
	return s.Scopes
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetSsoMappingClaim() *string {
	return s.SsoMappingClaim
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidc) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetAuthorizationEndpoint(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetClientAuthMethod(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.ClientAuthMethod = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetClientId(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.ClientId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetDiscoveryEndpoint(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetIssuer(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.Issuer = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetJwksUri(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.JwksUri = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetPkceRequired(v bool) *GetIdentityProviderResponseBodyDataOidc {
	s.PkceRequired = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetScopes(v []*string) *GetIdentityProviderResponseBodyDataOidc {
	s.Scopes = v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetSsoMappingClaim(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.SsoMappingClaim = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetTokenEndpoint(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.TokenEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) SetUserinfoEndpoint(v string) *GetIdentityProviderResponseBodyDataOidc {
	s.UserinfoEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidc) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyDataOidcConfiguration struct {
	// The absolute HTTPS URL required in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/authorize
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" xml:"authorizationEndpoint,omitempty"`
	// The default value is client_secret_post. If the value is none, PKCE must be enabled.
	//
	// example:
	//
	// client_secret_post
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty" xml:"clientAuthMethod,omitempty"`
	// The client identifier registered in the external IdP, which is different from the Matrix client. In-place replacement is not supported after configuration.
	//
	// example:
	//
	// example-client
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The absolute HTTPS discovery endpoint. It is mutually exclusive with manually entered endpoints.
	//
	// example:
	//
	// https://idp.example.com/.well-known/openid-configuration
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" xml:"discoveryEndpoint,omitempty"`
	// Indicates whether the client secret is saved. Only the existence status is returned, not the plaintext.
	//
	// example:
	//
	// false
	HasClientPrivateKey *bool `json:"hasClientPrivateKey,omitempty" xml:"hasClientPrivateKey,omitempty"`
	// Indicates whether the client secret is saved. Only the existence status is returned, not the plaintext.
	//
	// example:
	//
	// true
	HasClientSecret *bool `json:"hasClientSecret,omitempty" xml:"hasClientSecret,omitempty"`
	// The absolute HTTPS issuer of the external IdP. In-place replacement is not supported after configuration.
	//
	// example:
	//
	// https://idp.example.com
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// The absolute HTTPS URL required in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/jwks
	JwksUri *string `json:"jwksUri,omitempty" xml:"jwksUri,omitempty"`
	// Specifies whether PKCE is enabled. Default value: true. It must be set to true when no client authentication is used.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"pkceRequired,omitempty" xml:"pkceRequired,omitempty"`
	// An array of strings that must contain openid. The default values are openid, profile, and email.
	//
	// example:
	//
	// ["openid","profile","email"]
	Scopes []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// The claim name in the ID token used to match the SCIM externalId. If omitted during initial creation, the default value sub is used. If omitted during an update, the saved value is retained. For Azure, you can enter oid. The value must be non-empty and cannot contain whitespace characters. When you query oidc, the saved value is returned. When you query oidcConfiguration, the actual upstream value is returned, which may be empty.
	//
	// example:
	//
	// oid
	SsoMappingClaim *string `json:"ssoMappingClaim,omitempty" xml:"ssoMappingClaim,omitempty"`
	// The current upstream logon switch, which is Enabled or Disabled.
	//
	// example:
	//
	// Disabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
	// The absolute HTTPS URL required in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/token
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" xml:"tokenEndpoint,omitempty"`
	// The optional absolute HTTPS URL in manual mode. Do not specify this parameter in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/userinfo
	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" xml:"userinfoEndpoint,omitempty"`
}

func (s GetIdentityProviderResponseBodyDataOidcConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataOidcConfiguration) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetClientAuthMethod() *string {
	return s.ClientAuthMethod
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetClientId() *string {
	return s.ClientId
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetDiscoveryEndpoint() *string {
	return s.DiscoveryEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetHasClientPrivateKey() *bool {
	return s.HasClientPrivateKey
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetHasClientSecret() *bool {
	return s.HasClientSecret
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetIssuer() *string {
	return s.Issuer
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetJwksUri() *string {
	return s.JwksUri
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetScopes() []*string {
	return s.Scopes
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetSsoMappingClaim() *string {
	return s.SsoMappingClaim
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetAuthorizationEndpoint(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetClientAuthMethod(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.ClientAuthMethod = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetClientId(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.ClientId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetDiscoveryEndpoint(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetHasClientPrivateKey(v bool) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.HasClientPrivateKey = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetHasClientSecret(v bool) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.HasClientSecret = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetIssuer(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.Issuer = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetJwksUri(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.JwksUri = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetPkceRequired(v bool) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.PkceRequired = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetScopes(v []*string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.Scopes = v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetSsoMappingClaim(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.SsoMappingClaim = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetSsoStatus(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetTokenEndpoint(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.TokenEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) SetUserinfoEndpoint(v string) *GetIdentityProviderResponseBodyDataOidcConfiguration {
	s.UserinfoEndpoint = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataOidcConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyDataSaml struct {
	// The SAML binding type. Valid values: HTTP-Redirect and HTTP-POST. This parameter is required when manually filled in. Queries return short names.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// This parameter is required when manually filled in. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations do not support replacement. It is different from the Entity ID queried by the service provider (SP).
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required when manually filled in. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// When manually filled in, a non-empty array is required, using PEM X.509 certificate text. The current validity period is verified, and a maximum of two certificates are allowed after deduplication.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s GetIdentityProviderResponseBodyDataSaml) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataSaml) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataSaml) GetBindingType() *string {
	return s.BindingType
}

func (s *GetIdentityProviderResponseBodyDataSaml) GetEntityId() *string {
	return s.EntityId
}

func (s *GetIdentityProviderResponseBodyDataSaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *GetIdentityProviderResponseBodyDataSaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *GetIdentityProviderResponseBodyDataSaml) SetBindingType(v string) *GetIdentityProviderResponseBodyDataSaml {
	s.BindingType = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataSaml) SetEntityId(v string) *GetIdentityProviderResponseBodyDataSaml {
	s.EntityId = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataSaml) SetLoginUrl(v string) *GetIdentityProviderResponseBodyDataSaml {
	s.LoginUrl = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataSaml) SetX509Certificates(v []*string) *GetIdentityProviderResponseBodyDataSaml {
	s.X509Certificates = v
	return s
}

func (s *GetIdentityProviderResponseBodyDataSaml) Validate() error {
	return dara.Validate(s)
}

type GetIdentityProviderResponseBodyDataSyncScope struct {
	// When the mode is SelectedDepartments, this list must contain at least one non-empty department ID, which should be obtained from ListIdentityProviderDepartments. When the mode is AllAuthorized, pass an empty array or omit this parameter.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// The scope mode. Valid values: AllAuthorized (all authorized scopes, limited by application address book authorization) and SelectedDepartments (selected departments and all their sub-departments). A valid mode must be provided when explicitly submitting syncScope and is verified by the backend.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s GetIdentityProviderResponseBodyDataSyncScope) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponseBodyDataSyncScope) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponseBodyDataSyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *GetIdentityProviderResponseBodyDataSyncScope) GetMode() *string {
	return s.Mode
}

func (s *GetIdentityProviderResponseBodyDataSyncScope) SetDepartmentIds(v []*string) *GetIdentityProviderResponseBodyDataSyncScope {
	s.DepartmentIds = v
	return s
}

func (s *GetIdentityProviderResponseBodyDataSyncScope) SetMode(v string) *GetIdentityProviderResponseBodyDataSyncScope {
	s.Mode = &v
	return s
}

func (s *GetIdentityProviderResponseBodyDataSyncScope) Validate() error {
	return dara.Validate(s)
}
