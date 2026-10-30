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
	// The response message. An error description is returned if the request fails.
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
	// The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM queries and preparations return the actual receiving status and URL.
	Configuration *GetIdentityProviderResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The creation time in UTC format based on RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The event subscription callback URL. You must configure this URL in the application of the external identity provider to receive organization change events. An empty string is returned if the user pool is not activated.
	//
	// example:
	//
	// http://auth.cn-hangzhou.agentteams.aliyuncs.com/roa/dingtalk/event/up-123456
	EventSubscriptionCallbackUrl *string `json:"eventSubscriptionCallbackUrl,omitempty" xml:"eventSubscriptionCallbackUrl,omitempty"`
	// The external identity provider type. Valid values: DingTalk and Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The logon callback URL. You must configure this URL in the application of the external identity provider. An empty string is returned if the user pool is not activated.
	//
	// example:
	//
	// https://signin-cn-hangzhou.aliyunagentid.com/up-123456/dingtalk/callback
	LoginCallbackUrl *string `json:"loginCallbackUrl,omitempty" xml:"loginCallbackUrl,omitempty"`
	// Specifies whether workspace users are allowed to log on through the external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The application configuration of the external identity provider. Application secret configurations are not returned.
	Metadata *GetIdentityProviderResponseBodyDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The logon protocol. Only SAML is currently supported. This field may not be returned if no valid configuration is saved. OIDC placeholder initialization is not supported.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The SAML logon configuration. The input supports either XML or complete manual entry. Queries return the standardized configuration and do not echo the original XML. This field can be omitted if it is not configured.
	Saml *GetIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The status. Valid values:
	//
	// - CONFIGURED: The configuration is accepted and the system is waiting for the user pool to be activated.
	//
	// - SYNCING: Organization members are being synchronized.
	//
	// - SYNCED: Organization members are synchronized.
	//
	// - READY: The binding is in effect.
	//
	// - SYNC_FAILED: Organization member synchronization failed.
	//
	// - UPDATING: The configuration is being updated.
	//
	// - UPDATE_FAILED: Configuration update failed.
	//
	// - DISCONNECTING: The binding is being removed.
	//
	// - DISCONNECT_FAILED: Binding removal failed.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether organization member synchronization is enabled. After it is enabled, organization members are synchronized as workspace users by the external identity provider.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The configured synchronization scope, which follows the synchronization scope definition by department. If omitted during creation, the default value is AllAuthorized, and the first vendor synchronization is determined by syncOnBind. If omitted or set to null during an update, the original value is retained. Updating and saving the scope does not automatically trigger full synchronization.
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
	// The SAML binding type, which is HTTP-Redirect or HTTP-POST. Required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The IdP entity ID. Required for manual entry. Must be non-empty and no more than 2,048 characters in length. Existing configurations cannot be changed. This is different from the entity ID of the service provider (SP) returned in queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The IdP logon URL. Required for manual entry. Must be an absolute HTTPS URL without user credentials or fragments.
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
	// The App ID of the Lark application. Required when the binding type is Feishu.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The AppKey of the DingTalk application. Required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// The SAML binding URI in the historical metadata snapshot. To update CustomSSO, use saml.bindingType.
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
	// The CorpId of the DingTalk enterprise. Required when the binding type is DingTalk.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The IdP entity ID. Required for manual entry. Must be non-empty and no more than 2,048 characters in length. Existing configurations cannot be changed. This is different from the entity ID of the service provider (SP) returned in queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The IdP logon URL. Required for manual entry. Must be an absolute HTTPS URL without user credentials or fragments.
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

type GetIdentityProviderResponseBodyDataSaml struct {
	// The SAML binding type, which is HTTP-Redirect or HTTP-POST. Required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The IdP entity ID. Required for manual entry. Must be non-empty and no more than 2,048 characters in length. Existing configurations cannot be changed. This is different from the entity ID of the service provider (SP) returned in queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The IdP logon URL. Required for manual entry. Must be an absolute HTTPS URL without user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The list of IdP signing certificates. Required for manual entry. Must be a non-empty array using PEM X.509 certificate text. The current validity period is verified. A maximum of two certificates are allowed after deduplication.
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
	// The list of department IDs. When the mode is SelectedDepartments, at least one non-empty department ID must be included, which should be obtained from ListIdentityProviderDepartments. When the mode is AllAuthorized, pass an empty array or omit this field.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// The scope mode. Valid values:
	//
	// - AllAuthorized: All authorized scopes, subject to the application address book authorization limits.
	//
	// - SelectedDepartments: The selected departments and all their sub-departments.
	//
	// When explicitly submitting syncScope, you must provide a valid mode, which is verified by the backend.
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
