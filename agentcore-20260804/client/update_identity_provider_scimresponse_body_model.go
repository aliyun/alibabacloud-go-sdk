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
	// The business result code. If the request fails, handle the error based on the HTTP status code and the message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The shared binding after the request is accepted. The status indicates the status of shared resources and logon configurations, not the status of this SCIM task. Query GetIdentityProviderSCIM.operationStatus for the execution result.
	Data *UpdateIdentityProviderSCIMResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. It indicates the processing status of this request and is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Task result message. It is the request acceptance or error message. For asynchronous tasks, you must check task result by calling the corresponding query API.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. It is used to associate this call with troubleshooting. Retain it when a failure occurs.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates that the request is processed successfully. For asynchronous writes, you must continue to query the task result.
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
	// The actual upstream configuration. The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM query and prepare operations return the actual receiving status and URL.
	Configuration *UpdateIdentityProviderSCIMResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The creation time recorded by the server.
	//
	// example:
	//
	// 2026-10-26T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The event subscription callback URL. CustomSSO returns an empty string. You can obtain the SCIM receiving address through the SCIM query or prepare API. The example shows a JSON empty string "". The actual value does not contain double quotes.
	//
	// example:
	//
	// ""
	EventSubscriptionCallbackUrl *string `json:"eventSubscriptionCallbackUrl,omitempty" xml:"eventSubscriptionCallbackUrl,omitempty"`
	// The identity provider type. General CRUD operations support DingTalk, Feishu, and CustomSSO. SP queries only support CustomSSO. The SCIM path fixedly uses the CustomSSO shared pool.
	//
	// example:
	//
	// CustomSSO
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The logon callback URL. CustomSSO returns an empty string. You should obtain the ACS URL for SAML through the SP query API. The example shows a JSON empty string "". The actual value does not contain double quotes.
	//
	// example:
	//
	// ""
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// The logon switch. Do not set this parameter to true during CustomSSO initialization. You must save a valid SAML configuration before enabling it. Disabling it does not delete the configuration. When an asynchronous failure occurs, this expected value does not mean that it has taken effect on the upstream.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The standard IdP metadata. Vendor parameters are used for DingTalk and Feishu. CustomSSO updates use protocol and saml, and do not accept this field. The return value does not contain vendor keys.
	Metadata *UpdateIdentityProviderSCIMResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The logon protocol. Currently, only SAML is supported. This field may not be returned if no valid configuration is saved. OIDC placeholder initialization is not supported.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The SAML logon configuration. The input supports either XML or complete manual entry. The query returns the standardized configuration and does not echo the original XML. This field can be omitted if it is not configured.
	Saml *UpdateIdentityProviderSCIMResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The operation status of shared resources and logon configurations. It indicates the user pool lifecycle and logon configuration task status of CustomSSO, and does not represent the SCIM operation result. Query GetIdentityProviderSCIM.operationStatus for the SCIM switch result. You cannot use this field to determine whether all protocols are configured.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Do not set this parameter to true during CustomSSO initialization. You can only update it later by calling UpdateIdentityProviderSCIM. Determine the result based on operationStatus.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The last update time recorded by the server.
	//
	// example:
	//
	// 2026-10-26T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The ID of the target workspace. This parameter is required. The request region must be the same as the region of the target workspace.
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

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) GetSaml() *UpdateIdentityProviderSCIMResponseBodyDataSaml {
	return s.Saml
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

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetProtocol(v string) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *UpdateIdentityProviderSCIMResponseBodyData) SetSaml(v *UpdateIdentityProviderSCIMResponseBodyDataSaml) *UpdateIdentityProviderSCIMResponseBodyData {
	s.Saml = v
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
	if s.Saml != nil {
		if err := s.Saml.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderSCIMResponseBodyDataConfiguration struct {
	// The SAML binding type. Valid values: HTTP-Redirect and HTTP-POST. This parameter is required when you manually enter the configuration. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The entity ID of the IdP. This parameter is required when you manually enter the configuration. It cannot be empty and cannot exceed 2,048 characters. Existing configurations cannot be replaced. This is different from the entity ID of the SP returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The logon URL of the IdP. This parameter is required when you manually enter the configuration. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual SSO logon status of AgentIdentity. The actual switch value is Enabled or Disabled, which is not equivalent to the expected loginEnabled value saved in AC.
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
	// The canonical SAML binding identifier. It is the SAML binding URI in the historical metadata snapshot. For CustomSSO updates, use saml.bindingType.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect
	Binding *string `json:"binding,omitempty" xml:"binding,omitempty"`
	// The canonical SAML certificate content. It is the certificate string in the historical metadata snapshot. To edit the configuration, use saml.x509Certificates and do not submit the old metadata.
	//
	// example:
	//
	// MIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQELBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoXDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8oGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYNW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbYEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Okkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0EBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUWld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImcaWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTqP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjocJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdVjfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5
	Certificates *string `json:"certificates,omitempty" xml:"certificates,omitempty"`
	// The entity ID of the IdP. This parameter is required when you manually enter the configuration. It cannot be empty and cannot exceed 2,048 characters. Existing configurations cannot be replaced. This is different from the entity ID of the SP returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The logon URL of the IdP. This parameter is required when you manually enter the configuration. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
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

type UpdateIdentityProviderSCIMResponseBodyDataSaml struct {
	// The SAML binding type. Valid values: HTTP-Redirect and HTTP-POST. This parameter is required when you manually enter the configuration. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The entity ID of the IdP. This parameter is required when you manually enter the configuration. It cannot be empty and cannot exceed 2,048 characters. Existing configurations cannot be replaced. This is different from the entity ID of the SP returned by the SP query.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The logon URL of the IdP. This parameter is required when you manually enter the configuration. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The list of IdP signing certificates. When you manually enter the configuration, this parameter is required and must be a non-empty array. Use PEM X.509 certificate text. The current validity period is verified. A maximum of two certificates are allowed after deduplication.
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
