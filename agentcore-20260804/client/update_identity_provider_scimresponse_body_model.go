// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderSCIMResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateIdentityProviderSCIMResponseBody
	GetCode() *string
	SetData(v *UpdateIdentityProviderSCIMResponseBodyData) *UpdateIdentityProviderSCIMResponseBody
	GetData() *UpdateIdentityProviderSCIMResponseBodyData
	SetHttpStatusCode(v int32) *UpdateIdentityProviderSCIMResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateIdentityProviderSCIMResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateIdentityProviderSCIMResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateIdentityProviderSCIMResponseBody
	GetSuccess() *bool
}

type UpdateIdentityProviderSCIMResponseBody struct {
	// The business result code. If the request fails, use this code in combination with the HTTP status code and the result message for troubleshooting.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data object. Retains the original fields and the meaning of the status field. The new scimStatus field corresponds to the data.status field in the query API operation, and the operationStatus field retains the status of the most recent SCIM task.
	Data *UpdateIdentityProviderSCIMResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. Indicates the processing status of the current request, which does not necessarily reflect the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message. Indicates whether the request is accepted or returns an error message. For an asynchronous task, check the task result by calling the corresponding query API operation.
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
	// Indicates whether the current request is successful. A value of true indicates that the request is processed successfully. For asynchronous write operations, you must continue to query the task result.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateIdentityProviderSCIMResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateIdentityProviderSCIMResponseBody) GetData() *UpdateIdentityProviderSCIMResponseBodyData {
	return s.Data
}

func (s *UpdateIdentityProviderSCIMResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateIdentityProviderSCIMResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateIdentityProviderSCIMResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIdentityProviderSCIMResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIdentityProviderSCIMResponseBody) SetCode(v string) *UpdateIdentityProviderSCIMResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBody) SetData(v *UpdateIdentityProviderSCIMResponseBodyData) *UpdateIdentityProviderSCIMResponseBody {
	s.Data = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBody) SetHttpStatusCode(v int32) *UpdateIdentityProviderSCIMResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBody) SetMessage(v string) *UpdateIdentityProviderSCIMResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBody) SetRequestId(v string) *UpdateIdentityProviderSCIMResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBody) SetSuccess(v bool) *UpdateIdentityProviderSCIMResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderSCIMResponseBodyData struct {
	// The actual upstream configuration. The actual SAML configuration is read back on demand only by the GetIdentityProvider API operation. The SCIM query and preparation operations return the actual receiving status and URL.
	Configuration *UpdateIdentityProviderSCIMResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The time when the record was created on the server.
	//
	// example:
	//
	// 2026-10-26T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The SAML vendor event callback URL. For CustomSSO, an empty string is returned. To obtain the SCIM receiving URL, call the SCIM query or preparation API operation. The example shows a JSON empty string "", but the actual value does not include double quotation marks.
	//
	// example:
	//
	// ""
	EventSubscriptionCallbackUrl *string `json:"eventSubscriptionCallbackUrl,omitempty" xml:"eventSubscriptionCallbackUrl,omitempty"`
	// The identity provider type. General CRUD operations support DingTalk, Feishu, and CustomSSO. SP queries support only CustomSSO. The SCIM path uses the CustomSSO shared pool.
	//
	// example:
	//
	// CustomSSO
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The SAML vendor callback URL. For CustomSSO, an empty string is returned. To obtain the Assertion Consumer Service (ACS) URL for SAML, call the SP query API operation. The example shows a JSON empty string "", but the actual value does not include double quotation marks.
	//
	// example:
	//
	// ""
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// Specifies whether logon is enabled. Do not set this parameter to true during CustomSSO initialization. You must save a valid SAML configuration before enabling logon. Disabling logon does not delete the configuration. If an asynchronous task fails, this expected value does not indicate that the change has taken effect on the upstream server.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The standard identity provider (IdP) metadata. Vendor parameters are used for DingTalk and Feishu. To update CustomSSO, use the protocol and saml fields. This field is not accepted for CustomSSO updates. The return value does not include vendor keys.
	Metadata *UpdateIdentityProviderSCIMResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The security summary of the expected OIDC configuration. This summary does not include credentials. The original SAML fields and their meanings are retained.
	Oidc *UpdateIdentityProviderSCIMResponseBodyDataOidc `json:"oidc,omitempty" xml:"oidc,omitempty" type:"Struct"`
	// The original status of the most recent independent SCIM operation task. If no historical tasks exist, an empty string ("") is returned.
	//
	// example:
	//
	// PENDING
	OperationStatus *string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty"`
	// The logon protocol. Only SAML is supported. This field may not be returned if no valid configuration is saved. You cannot use OIDC as a placeholder for initialization.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The SAML logon configuration. You can provide the input by using XML or by manually entering all fields. The query returns the standardized configuration and does not echo the original XML. This field can be omitted if it is not configured.
	Saml *UpdateIdentityProviderSCIMResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The independent SCIM status. This field does not change the meaning of the original status field for shared resources and logon configurations. It corresponds to the data.status field in the query API operation.
	//
	// example:
	//
	// UPDATING
	ScimStatus *string `json:"scimStatus,omitempty" xml:"scimStatus,omitempty"`
	// The operation status of the CustomSSO user pool lifecycle and logon configuration tasks. This status does not indicate the result of SCIM operations. To query the result of the SCIM switch, use the GetIdentityProviderSCIM.operationStatus field. You cannot use this field to determine whether all protocols are configured.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether to enable SCIM synchronization. Do not set this parameter to true during CustomSSO initialization. You can only update this parameter later by calling UpdateIdentityProviderSCIM. Determine the result based on the operationStatus field.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The time when the record was last updated on the server.
	//
	// example:
	//
	// 2026-10-26T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// Required. The ID of the destination workspace. The region of the request must be the same as the region of the destination workspace.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateIdentityProviderSCIMResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetConfiguration() *UpdateIdentityProviderSCIMResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetEventSubscriptionCallbackUrl() *string {
	return s.EventSubscriptionCallbackUrl
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetLoginCallbackUrl() *string {
	return s.LoginCallbackUrl
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetMetadata() *UpdateIdentityProviderSCIMResponseBodyDataMetadata {
	return s.Metadata
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetOidc() *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	return s.Oidc
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetSaml() *UpdateIdentityProviderSCIMResponseBodyDataSaml {
	return s.Saml
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetScimStatus() *string {
	return s.ScimStatus
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetConfiguration(v *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Configuration = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetCreatedAt(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetEventSubscriptionCallbackUrl(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.EventSubscriptionCallbackUrl = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetIdentityProviderType(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetLoginCallbackUrl(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.LoginCallbackUrl = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetLoginEnabled(v bool) *UpdateIdentityProviderSCIMResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetMetadata(v *UpdateIdentityProviderSCIMResponseBodyDataMetadata) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Metadata = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetOidc(v *UpdateIdentityProviderSCIMResponseBodyDataOidc) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Oidc = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetOperationStatus(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.OperationStatus = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetProtocol(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetSaml(v *UpdateIdentityProviderSCIMResponseBodyDataSaml) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Saml = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetScimStatus(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.ScimStatus = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetStatus(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetSyncEnabled(v bool) *UpdateIdentityProviderSCIMResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetUpdatedAt(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetWorkspaceId(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) Validate() error {
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
	if s.Saml != nil {
		if err := s.Saml.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderSCIMResponseBodyDataConfiguration struct {
	// The SAML binding type. Valid values: HTTP-Redirect and HTTP-POST. This field is required when you manually enter the configuration. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The IdP entity ID. This field is required when you manually enter the metadata. It must be a non-empty string of up to 2,048 characters. You cannot change the entity ID after it is configured. This ID is different from the local entity ID returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The IdP SSO logon URL. This field is required when you manually enter the metadata. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual SSO logon status of AgentIdentity. The actual switch value is Enabled or Disabled, which is not equivalent to the expected value of loginEnabled saved in the access configuration.
	//
	// example:
	//
	// Enabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
}

func (s UpdateIdentityProviderSCIMResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) SetBindingType(v string) *UpdateIdentityProviderSCIMResponseBodyDataConfiguration {
	s.BindingType = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) SetEntityId(v string) *UpdateIdentityProviderSCIMResponseBodyDataConfiguration {
	s.EntityId = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) SetLoginUrl(v string) *UpdateIdentityProviderSCIMResponseBodyDataConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) SetSsoStatus(v string) *UpdateIdentityProviderSCIMResponseBodyDataConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderSCIMResponseBodyDataMetadata struct {
	// The SAML binding URI in the historical metadata snapshot. To update CustomSSO, use the saml.bindingType field.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect
	Binding *string `json:"binding,omitempty" xml:"binding,omitempty"`
	// The certificate string in the historical metadata snapshot. To edit the configuration, use the saml.x509Certificates field instead of submitting the old metadata.
	//
	// example:
	//
	// MIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQELBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoXDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8oGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYNW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbYEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Okkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0EBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUWld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImcaWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTqP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjocJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdVjfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5
	Certificates *string `json:"certificates,omitempty" xml:"certificates,omitempty"`
	// The IdP entity ID. This field is required when you manually enter the metadata. It must be a non-empty string of up to 2,048 characters. You cannot change the entity ID after it is configured. This ID is different from the local entity ID returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The IdP SSO logon URL. This field is required when you manually enter the metadata. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
}

func (s UpdateIdentityProviderSCIMResponseBodyDataMetadata) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponseBodyDataMetadata) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) GetBinding() *string {
	return s.Binding
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) GetCertificates() *string {
	return s.Certificates
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) SetBinding(v string) *UpdateIdentityProviderSCIMResponseBodyDataMetadata {
	s.Binding = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) SetCertificates(v string) *UpdateIdentityProviderSCIMResponseBodyDataMetadata {
	s.Certificates = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) SetEntityId(v string) *UpdateIdentityProviderSCIMResponseBodyDataMetadata {
	s.EntityId = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) SetLoginUrl(v string) *UpdateIdentityProviderSCIMResponseBodyDataMetadata {
	s.LoginUrl = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataMetadata) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderSCIMResponseBodyDataOidc struct {
	// The authorization endpoint. This field is required in manual mode and must be an absolute HTTPS URL. Do not specify this field in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/authorize
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" xml:"authorizationEndpoint,omitempty"`
	// The client authentication method. Default value: client_secret_post. If the value is none, Proof Key for Code Exchange (PKCE) must be enabled.
	//
	// example:
	//
	// client_secret_post
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty" xml:"clientAuthMethod,omitempty"`
	// The client ID registered in the external IdP. This ID is different from the Matrix client ID. You cannot change the client ID in place after it is configured.
	//
	// example:
	//
	// example-client
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The absolute HTTPS OIDC discovery endpoint. This field is mutually exclusive with manually entered endpoints.
	//
	// example:
	//
	// https://idp.example.com/.well-known/openid-configuration
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" xml:"discoveryEndpoint,omitempty"`
	// The absolute HTTPS issuer of the external IdP. You cannot change the issuer in place after it is configured.
	//
	// example:
	//
	// https://idp.example.com
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// The authorization endpoint. This field is required in manual mode and must be an absolute HTTPS URL. Do not specify this field in discovery mode.
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
	// The array of strings that must contain openid. Default values: openid, profile, and email.
	//
	// example:
	//
	// ["openid","profile","email"]
	Scopes []*string `json:"scopes,omitempty" xml:"scopes,omitempty" type:"Repeated"`
	// The name of the claim in the ID token that is used to match the SCIM externalId. If omitted during creation, the default value sub is used. If omitted during an update, the saved value is retained. For Azure, you can specify oid. The value must be non-empty and cannot contain whitespace characters. When querying oidc, the saved value is returned. When querying oidcConfiguration, the actual upstream value is returned, which may be empty.
	//
	// example:
	//
	// oid
	SsoMappingClaim *string `json:"ssoMappingClaim,omitempty" xml:"ssoMappingClaim,omitempty"`
	// The authorization endpoint. This field is required in manual mode and must be an absolute HTTPS URL. Do not specify this field in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/token
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" xml:"tokenEndpoint,omitempty"`
	// The optional absolute HTTPS URL for manual mode. Leave this field empty in discovery mode.
	//
	// example:
	//
	// https://idp.example.com/userinfo
	UserinfoEndpoint *string `json:"userinfoEndpoint,omitempty" xml:"userinfoEndpoint,omitempty"`
}

func (s UpdateIdentityProviderSCIMResponseBodyDataOidc) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponseBodyDataOidc) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetAuthorizationEndpoint() *string {
	return s.AuthorizationEndpoint
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetClientAuthMethod() *string {
	return s.ClientAuthMethod
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetClientId() *string {
	return s.ClientId
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetDiscoveryEndpoint() *string {
	return s.DiscoveryEndpoint
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetIssuer() *string {
	return s.Issuer
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetJwksUri() *string {
	return s.JwksUri
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetPkceRequired() *bool {
	return s.PkceRequired
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetScopes() []*string {
	return s.Scopes
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetSsoMappingClaim() *string {
	return s.SsoMappingClaim
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetTokenEndpoint() *string {
	return s.TokenEndpoint
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) GetUserinfoEndpoint() *string {
	return s.UserinfoEndpoint
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetAuthorizationEndpoint(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetClientAuthMethod(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.ClientAuthMethod = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetClientId(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.ClientId = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetDiscoveryEndpoint(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetIssuer(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.Issuer = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetJwksUri(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.JwksUri = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetPkceRequired(v bool) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.PkceRequired = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetScopes(v []*string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.Scopes = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetSsoMappingClaim(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.SsoMappingClaim = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetTokenEndpoint(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.TokenEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) SetUserinfoEndpoint(v string) *UpdateIdentityProviderSCIMResponseBodyDataOidc {
	s.UserinfoEndpoint = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataOidc) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderSCIMResponseBodyDataSaml struct {
	// The SAML binding type. Valid values: HTTP-Redirect and HTTP-POST. This field is required when you manually enter the configuration. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The IdP entity ID. This field is required when you manually enter the metadata. It must be a non-empty string of up to 2,048 characters. You cannot change the entity ID after it is configured. This ID is different from the local entity ID returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The IdP SSO logon URL. This field is required when you manually enter the metadata. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The list of IdP signing certificates. When you manually enter the configuration, this field is required and must be a non-empty array of PEM X.509 certificate texts. The system verifies the current validity period. A maximum of two certificates are retained after deduplication.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s UpdateIdentityProviderSCIMResponseBodyDataSaml) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderSCIMResponseBodyDataSaml) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) GetBindingType() *string {
	return s.BindingType
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) SetBindingType(v string) *UpdateIdentityProviderSCIMResponseBodyDataSaml {
	s.BindingType = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) SetEntityId(v string) *UpdateIdentityProviderSCIMResponseBodyDataSaml {
	s.EntityId = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) SetLoginUrl(v string) *UpdateIdentityProviderSCIMResponseBodyDataSaml {
	s.LoginUrl = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) SetX509Certificates(v []*string) *UpdateIdentityProviderSCIMResponseBodyDataSaml {
	s.X509Certificates = v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyDataSaml) Validate() error {
	return dara.Validate(s)
}
