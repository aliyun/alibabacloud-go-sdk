// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIdentityProviderResponseBody
	GetCode() *string
	SetData(v *DeleteIdentityProviderResponseBodyData) *DeleteIdentityProviderResponseBody
	GetData() *DeleteIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *DeleteIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIdentityProviderResponseBody
	GetSuccess() *bool
}

type DeleteIdentityProviderResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The external identity provider binding information after the unbind operation is accepted.
	Data *DeleteIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s DeleteIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIdentityProviderResponseBody) GetData() *DeleteIdentityProviderResponseBodyData {
	return s.Data
}

func (s *DeleteIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIdentityProviderResponseBody) SetCode(v string) *DeleteIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetData(v *DeleteIdentityProviderResponseBodyData) *DeleteIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetHttpStatusCode(v int32) *DeleteIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetMessage(v string) *DeleteIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetRequestId(v string) *DeleteIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) SetSuccess(v bool) *DeleteIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteIdentityProviderResponseBodyData struct {
	// The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM queries and preparations return the actual receiving status and URL.
	Configuration *DeleteIdentityProviderResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The type of the external identity provider. Valid values: DingTalk and Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The security summary of the expected configuration. It does not contain credentials. The original SAML fields and meanings are retained.
	Oidc *DeleteIdentityProviderResponseBodyDataOidc `json:"oidc,omitempty" xml:"oidc,omitempty" type:"Struct"`
	// Currently, only SAML is supported. This field is not returned when no valid configuration is saved. Initialization with an OIDC placeholder is not supported.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The input supports either XML or complete manual entry. Queries return the standardized configuration and do not echo the original XML. This field can be omitted when not configured.
	Saml *DeleteIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
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
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBodyData) GetConfiguration() *DeleteIdentityProviderResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *DeleteIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *DeleteIdentityProviderResponseBodyData) GetOidc() *DeleteIdentityProviderResponseBodyDataOidc {
	return s.Oidc
}

func (s *DeleteIdentityProviderResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *DeleteIdentityProviderResponseBodyData) GetSaml() *DeleteIdentityProviderResponseBodyDataSaml {
	return s.Saml
}

func (s *DeleteIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DeleteIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteIdentityProviderResponseBodyData) SetConfiguration(v *DeleteIdentityProviderResponseBodyDataConfiguration) *DeleteIdentityProviderResponseBodyData {
	s.Configuration = v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *DeleteIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetOidc(v *DeleteIdentityProviderResponseBodyDataOidc) *DeleteIdentityProviderResponseBodyData {
	s.Oidc = v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetProtocol(v string) *DeleteIdentityProviderResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetSaml(v *DeleteIdentityProviderResponseBodyDataSaml) *DeleteIdentityProviderResponseBodyData {
	s.Saml = v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetStatus(v string) *DeleteIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) SetWorkspaceId(v string) *DeleteIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyData) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.Oidc != nil {
		if err := s.Oidc.Validate(); err != nil {
			return err
		}
	}
	if s.Saml != nil {
		if err := s.Saml.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteIdentityProviderResponseBodyDataConfiguration struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations cannot be replaced. This is different from the local entity ID queried by the service provider (SP).
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual switch value is Enabled or Disabled, which differs from the expected loginEnabled value saved in the AC.
	//
	// example:
	//
	// Enabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
}

func (s DeleteIdentityProviderResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) SetBindingType(v string) *DeleteIdentityProviderResponseBodyDataConfiguration {
	s.BindingType = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) SetEntityId(v string) *DeleteIdentityProviderResponseBodyDataConfiguration {
	s.EntityId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) SetLoginUrl(v string) *DeleteIdentityProviderResponseBodyDataConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) SetSsoStatus(v string) *DeleteIdentityProviderResponseBodyDataConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type DeleteIdentityProviderResponseBodyDataOidc struct {
	// The absolute HTTPS URL required in manual entry mode. Do not enter this in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/authorize
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" xml:"authorizationEndpoint,omitempty"`
	// The default value is client_secret_post. If none is selected, PKCE must be enabled.
	//
	// example:
	//
	// client_secret_post
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty" xml:"clientAuthMethod,omitempty"`
	// The client ID registered in the external IdP, which is different from the Matrix client. In-place replacement is not supported after configuration.
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
	// The absolute HTTPS URL required in manual entry mode. Do not enter this in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/jwks
	JwksUri *string `json:"jwksUri,omitempty" xml:"jwksUri,omitempty"`
	// Specifies whether PKCE is enabled. Default value: true. It must be true when client authentication is not used.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"pkceRequired,omitempty" xml:"pkceRequired,omitempty"`
	// The string array. It must contain openid. The default values are openid, profile, and email.
	//
	// example:
	//
	// ["openid","profile","email"]
	Scopes []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// The claim name in the ID token used to match the SCIM externalId. If omitted during the first creation, the default value sub is used. If omitted during an update, the saved value is retained. For Azure, you can enter oid. It must be non-empty and contain no whitespace characters. When queried, oidc returns the saved value, and oidcConfiguration returns the actual upstream value, which may be empty.
	//
	// example:
	//
	// oid
	SsoMappingClaim *string `json:"ssoMappingClaim,omitempty" xml:"ssoMappingClaim,omitempty"`
	// The absolute HTTPS URL required in manual entry mode. Do not enter this in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/token
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" xml:"tokenEndpoint,omitempty"`
	// The optional absolute HTTPS URL in manual entry mode. Do not enter this in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/userinfo
	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" xml:"userinfoEndpoint,omitempty"`
}

func (s DeleteIdentityProviderResponseBodyDataOidc) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBodyDataOidc) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetClientAuthMethod() *string {
	return s.ClientAuthMethod
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetDiscoveryEndpoint() *string {
	return s.DiscoveryEndpoint
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetIssuer() *string {
	return s.Issuer
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetJwksUri() *string {
	return s.JwksUri
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetScopes() []*string {
	return s.Scopes
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetSsoMappingClaim() *string {
	return s.SsoMappingClaim
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetAuthorizationEndpoint(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetClientAuthMethod(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.ClientAuthMethod = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetClientId(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.ClientId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetDiscoveryEndpoint(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetIssuer(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.Issuer = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetJwksUri(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.JwksUri = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetPkceRequired(v bool) *DeleteIdentityProviderResponseBodyDataOidc {
	s.PkceRequired = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetScopes(v []*string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.Scopes = v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetSsoMappingClaim(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.SsoMappingClaim = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetTokenEndpoint(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.TokenEndpoint = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) SetUserinfoEndpoint(v string) *DeleteIdentityProviderResponseBodyDataOidc {
	s.UserinfoEndpoint = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataOidc) Validate() error {
	return dara.Validate(s)
}

type DeleteIdentityProviderResponseBodyDataSaml struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain a maximum of 2,048 characters. Existing configurations cannot be replaced. This is different from the local entity ID queried by the service provider (SP).
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// Required for manual entry as a non-empty array. Use PEM X.509 certificate text. The current validity period is verified. A maximum of two certificates are allowed after deduplication.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s DeleteIdentityProviderResponseBodyDataSaml) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponseBodyDataSaml) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) GetBindingType() *string {
	return s.BindingType
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) SetBindingType(v string) *DeleteIdentityProviderResponseBodyDataSaml {
	s.BindingType = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) SetEntityId(v string) *DeleteIdentityProviderResponseBodyDataSaml {
	s.EntityId = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) SetLoginUrl(v string) *DeleteIdentityProviderResponseBodyDataSaml {
	s.LoginUrl = &v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) SetX509Certificates(v []*string) *DeleteIdentityProviderResponseBodyDataSaml {
	s.X509Certificates = v
	return s
}

func (s *DeleteIdentityProviderResponseBodyDataSaml) Validate() error {
	return dara.Validate(s)
}
