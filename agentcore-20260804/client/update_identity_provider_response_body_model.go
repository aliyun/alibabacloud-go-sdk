// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateIdentityProviderResponseBody
	GetCode() *string
	SetData(v *UpdateIdentityProviderResponseBodyData) *UpdateIdentityProviderResponseBody
	GetData() *UpdateIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *UpdateIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateIdentityProviderResponseBody
	GetSuccess() *bool
}

type UpdateIdentityProviderResponseBody struct {
	// The business result code. If the request fails, use this code along with the HTTP status code and message for troubleshooting.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business result. For the result structure, refer to the response fields of this operation. For asynchronous write operations, the response indicates that the request is accepted, but does not guarantee that the background task is completed.
	Data *UpdateIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The processing status of the current request. This does not indicate the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request acceptance or error message. For asynchronous tasks, verify the task result by calling the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Retain this ID to correlate the current call and troubleshoot failures.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the current request is processed successfully. A value of true indicates success. For asynchronous write operations, you must still query the task result.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateIdentityProviderResponseBody) GetData() *UpdateIdentityProviderResponseBodyData {
	return s.Data
}

func (s *UpdateIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIdentityProviderResponseBody) SetCode(v string) *UpdateIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetData(v *UpdateIdentityProviderResponseBodyData) *UpdateIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetHttpStatusCode(v int32) *UpdateIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetMessage(v string) *UpdateIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetRequestId(v string) *UpdateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) SetSuccess(v bool) *UpdateIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderResponseBodyData struct {
	// The actual upstream configuration. The actual SAML configuration is read on demand only in GetIdentityProvider. SCIM queries and preparations return the actual reception status and URL.
	Configuration *UpdateIdentityProviderResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// General CRUD operations support DingTalk, Feishu, and CustomSSO. SP queries support only CustomSSO. The SCIM path uses the CustomSSO shared pool.
	//
	// example:
	//
	// CustomSSO
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// Do not pass true during CustomSSO initialization. A valid SAML configuration must be saved before enabling. Disabling does not delete the configuration. If an asynchronous failure occurs, this expected value does not indicate that it has taken effect upstream.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The expected security summary of the configuration. This does not include credentials. The original SAML fields and meanings are retained.
	Oidc *UpdateIdentityProviderResponseBodyDataOidc `json:"oidc,omitempty" xml:"oidc,omitempty" type:"Struct"`
	// The login protocol for CustomSSO. Currently, only SAML is supported. You can pass SAML when saving the SAML configuration. You can omit this parameter if you only modify loginEnabled. DingTalk and Feishu do not accept this field.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The SAML configuration for CustomSSO. Choose either metadataXml or the four manually entered fields: entityId, loginUrl, bindingType, and x509Certificates. If not submitted, the existing configuration is retained. DingTalk and Feishu do not accept this field.
	Saml *UpdateIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The status of the identity source resource and logon configuration. A READY status in the CustomSSO shared pool does not mean that SAML is configured. Query operationStatus independently for SCIM operation results.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether synchronization is enabled. Do not set this parameter to true during CustomSSO initialization. You can only update this parameter later by calling UpdateIdentityProviderSCIM. Evaluate the result in combination with operationStatus.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// Follows the definition of the department-based synchronization scope. If omitted during creation, the default value is AllAuthorized. The initial provider synchronization is determined by syncOnBind. If omitted or set to null during an update, the original value is retained. Saving the scope during an update does not automatically trigger a full synchronization.
	SyncScope *UpdateIdentityProviderResponseBodyDataSyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
	// Required. The ID of the destination workspace. The region of the request must be the same as the region of the destination workspace.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBodyData) GetConfiguration() *UpdateIdentityProviderResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *UpdateIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *UpdateIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *UpdateIdentityProviderResponseBodyData) GetOidc() *UpdateIdentityProviderResponseBodyDataOidc {
	return s.Oidc
}

func (s *UpdateIdentityProviderResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateIdentityProviderResponseBodyData) GetSaml() *UpdateIdentityProviderResponseBodyDataSaml {
	return s.Saml
}

func (s *UpdateIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderResponseBodyData) GetSyncScope() *UpdateIdentityProviderResponseBodyDataSyncScope {
	return s.SyncScope
}

func (s *UpdateIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateIdentityProviderResponseBodyData) SetConfiguration(v *UpdateIdentityProviderResponseBodyDataConfiguration) *UpdateIdentityProviderResponseBodyData {
	s.Configuration = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *UpdateIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *UpdateIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetOidc(v *UpdateIdentityProviderResponseBodyDataOidc) *UpdateIdentityProviderResponseBodyData {
	s.Oidc = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetProtocol(v string) *UpdateIdentityProviderResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetSaml(v *UpdateIdentityProviderResponseBodyDataSaml) *UpdateIdentityProviderResponseBodyData {
	s.Saml = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetStatus(v string) *UpdateIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *UpdateIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetSyncScope(v *UpdateIdentityProviderResponseBodyDataSyncScope) *UpdateIdentityProviderResponseBodyData {
	s.SyncScope = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) SetWorkspaceId(v string) *UpdateIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyData) Validate() error {
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
	if s.SyncScope != nil {
		if err := s.SyncScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderResponseBodyDataConfiguration struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain up to 2,048 characters. The existing configuration does not support replacement. It is different from the local entity ID returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL and must not contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual switch value, which is Enabled or Disabled. This is not equivalent to the expected loginEnabled value saved in AC.
	//
	// example:
	//
	// Enabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
}

func (s UpdateIdentityProviderResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) SetBindingType(v string) *UpdateIdentityProviderResponseBodyDataConfiguration {
	s.BindingType = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) SetEntityId(v string) *UpdateIdentityProviderResponseBodyDataConfiguration {
	s.EntityId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) SetLoginUrl(v string) *UpdateIdentityProviderResponseBodyDataConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) SetSsoStatus(v string) *UpdateIdentityProviderResponseBodyDataConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderResponseBodyDataOidc struct {
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
	// The client identifier registered in the external IdP, which is different from the Matrix client. This value cannot be changed in place after it is configured.
	//
	// example:
	//
	// example-client
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The absolute HTTPS discovery endpoint URL. This parameter is mutually exclusive with manually specified endpoints.
	//
	// example:
	//
	// https://idp.example.com/.well-known/openid-configuration
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" xml:"discoveryEndpoint,omitempty"`
	// The absolute HTTPS issuer URL of the external IdP. This value cannot be changed in place after it is configured.
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
	// Specifies whether PKCE is enabled. Default value: true. This parameter must be set to true when no client authentication is used.
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"pkceRequired,omitempty" xml:"pkceRequired,omitempty"`
	// The array of scope strings. It must include openid. Default values: openid, profile, and email.
	//
	// example:
	//
	// ["openid","profile","email"]
	Scopes []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// The name of the claim in the ID token used to match the SCIM externalId. If omitted during the initial configuration, the default value is sub. If omitted during an update, the saved value is retained. For Azure, you can specify oid. The value must be non-empty and cannot contain whitespace characters. When queried, oidc returns the saved value, while oidcConfiguration returns the actual upstream value, which may be empty.
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

func (s UpdateIdentityProviderResponseBodyDataOidc) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBodyDataOidc) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetClientAuthMethod() *string {
	return s.ClientAuthMethod
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetDiscoveryEndpoint() *string {
	return s.DiscoveryEndpoint
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetIssuer() *string {
	return s.Issuer
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetJwksUri() *string {
	return s.JwksUri
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetScopes() []*string {
	return s.Scopes
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetSsoMappingClaim() *string {
	return s.SsoMappingClaim
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetAuthorizationEndpoint(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetClientAuthMethod(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.ClientAuthMethod = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetClientId(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.ClientId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetDiscoveryEndpoint(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetIssuer(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.Issuer = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetJwksUri(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.JwksUri = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetPkceRequired(v bool) *UpdateIdentityProviderResponseBodyDataOidc {
	s.PkceRequired = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetScopes(v []*string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.Scopes = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetSsoMappingClaim(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.SsoMappingClaim = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetTokenEndpoint(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.TokenEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) SetUserinfoEndpoint(v string) *UpdateIdentityProviderResponseBodyDataOidc {
	s.UserinfoEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataOidc) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderResponseBodyDataSaml struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain up to 2,048 characters. The existing configuration does not support replacement. It is different from the local entity ID returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// Required for manual entry. It must be an absolute HTTPS URL and must not contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// Required for manual entry. It must be a non-empty array using PEM X.509 certificate text. The current validity period is verified. Up to two certificates are allowed after deduplication.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s UpdateIdentityProviderResponseBodyDataSaml) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBodyDataSaml) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) GetBindingType() *string {
	return s.BindingType
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) SetBindingType(v string) *UpdateIdentityProviderResponseBodyDataSaml {
	s.BindingType = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) SetEntityId(v string) *UpdateIdentityProviderResponseBodyDataSaml {
	s.EntityId = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) SetLoginUrl(v string) *UpdateIdentityProviderResponseBodyDataSaml {
	s.LoginUrl = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) SetX509Certificates(v []*string) *UpdateIdentityProviderResponseBodyDataSaml {
	s.X509Certificates = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataSaml) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderResponseBodyDataSyncScope struct {
	// When set to SelectedDepartments, it must contain at least one non-empty department ID, which should be obtained from ListIdentityProviderDepartments. When set to AllAuthorized, pass an empty array or omit this parameter.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// AllAuthorized: the full authorization scope, subject to the application address book authorization limits. SelectedDepartments: the selected departments and all their sub-departments. When explicitly submitting syncScope, a valid mode must be provided and verified by the backend.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateIdentityProviderResponseBodyDataSyncScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderResponseBodyDataSyncScope) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderResponseBodyDataSyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *UpdateIdentityProviderResponseBodyDataSyncScope) GetMode() *string {
	return s.Mode
}

func (s *UpdateIdentityProviderResponseBodyDataSyncScope) SetDepartmentIds(v []*string) *UpdateIdentityProviderResponseBodyDataSyncScope {
	s.DepartmentIds = v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataSyncScope) SetMode(v string) *UpdateIdentityProviderResponseBodyDataSyncScope {
	s.Mode = &v
	return s
}

func (s *UpdateIdentityProviderResponseBodyDataSyncScope) Validate() error {
	return dara.Validate(s)
}
