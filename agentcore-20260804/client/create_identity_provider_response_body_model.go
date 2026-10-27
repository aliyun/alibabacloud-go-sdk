// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIdentityProviderResponseBody
	GetCode() *string
	SetData(v *CreateIdentityProviderResponseBodyData) *CreateIdentityProviderResponseBody
	GetData() *CreateIdentityProviderResponseBodyData
	SetHttpStatusCode(v int32) *CreateIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIdentityProviderResponseBody
	GetSuccess() *bool
}

type CreateIdentityProviderResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The binding information of the external identity provider.
	Data *CreateIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned when the request fails.
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

func (s CreateIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIdentityProviderResponseBody) GetData() *CreateIdentityProviderResponseBodyData {
	return s.Data
}

func (s *CreateIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIdentityProviderResponseBody) SetCode(v string) *CreateIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetData(v *CreateIdentityProviderResponseBodyData) *CreateIdentityProviderResponseBody {
	s.Data = v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetHttpStatusCode(v int32) *CreateIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetMessage(v string) *CreateIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetRequestId(v string) *CreateIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) SetSuccess(v bool) *CreateIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIdentityProviderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderResponseBodyData struct {
	// The actual SAML configuration is read back on demand only in the GetIdentityProvider operation. SCIM queries and preparations return the actual receiving status and URL.
	Configuration *CreateIdentityProviderResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The type of the external identity provider. Valid values: DingTalk and Feishu.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// Specifies whether to allow workspace users to log on through the external identity provider.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// Only SAML is currently supported. This parameter may not be returned if no valid configuration is saved. Placeholder initialization with OIDC is not supported.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The input supports either XML or complete manual entry. Queries return the standardized configuration and do not echo the original XML. This parameter can be omitted if it is not configured.
	Saml *CreateIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The binding status. Valid values:
	//
	// - CONFIGURED: The configuration is accepted and waiting for the user pool to be activated.
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
	// Specifies whether to enable organization member synchronization. After this feature is enabled, the external identity provider synchronizes organization members as workspace users.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// Retains the department-based synchronization scope definition. If this parameter is omitted during creation, the default value is AllAuthorized, and the first vendor synchronization is determined by syncOnBind. If this parameter is omitted or set to null during an update, the original value is retained. Updating and saving the scope does not automatically perform full synchronization.
	SyncScope *CreateIdentityProviderResponseBodyDataSyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateIdentityProviderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBodyData) GetConfiguration() *CreateIdentityProviderResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *CreateIdentityProviderResponseBodyData) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *CreateIdentityProviderResponseBodyData) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *CreateIdentityProviderResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateIdentityProviderResponseBodyData) GetSaml() *CreateIdentityProviderResponseBodyDataSaml {
	return s.Saml
}

func (s *CreateIdentityProviderResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateIdentityProviderResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *CreateIdentityProviderResponseBodyData) GetSyncScope() *CreateIdentityProviderResponseBodyDataSyncScope {
	return s.SyncScope
}

func (s *CreateIdentityProviderResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateIdentityProviderResponseBodyData) SetConfiguration(v *CreateIdentityProviderResponseBodyDataConfiguration) *CreateIdentityProviderResponseBodyData {
	s.Configuration = v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetIdentityProviderType(v string) *CreateIdentityProviderResponseBodyData {
	s.IdentityProviderType = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetLoginEnabled(v bool) *CreateIdentityProviderResponseBodyData {
	s.LoginEnabled = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetProtocol(v string) *CreateIdentityProviderResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetSaml(v *CreateIdentityProviderResponseBodyDataSaml) *CreateIdentityProviderResponseBodyData {
	s.Saml = v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetStatus(v string) *CreateIdentityProviderResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetSyncEnabled(v bool) *CreateIdentityProviderResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetSyncScope(v *CreateIdentityProviderResponseBodyDataSyncScope) *CreateIdentityProviderResponseBodyData {
	s.SyncScope = v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) SetWorkspaceId(v string) *CreateIdentityProviderResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyData) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
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

type CreateIdentityProviderResponseBodyDataConfiguration struct {
	// HTTP-Redirect or HTTP-POST. This parameter is required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// This parameter is required for manual entry. It must be non-empty and cannot exceed 2,048 characters. Existing configurations cannot be replaced. It is different from the Entity ID of the service provider (SP) returned in queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required for manual entry. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual switch value is Enabled or Disabled. It is not equivalent to the expected loginEnabled value saved in the access configuration.
	//
	// example:
	//
	// Enabled
	SsoStatus *string `json:"ssoStatus,omitempty" xml:"ssoStatus,omitempty"`
}

func (s CreateIdentityProviderResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) GetSsoStatus() *string {
	return s.SsoStatus
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) SetBindingType(v string) *CreateIdentityProviderResponseBodyDataConfiguration {
	s.BindingType = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) SetEntityId(v string) *CreateIdentityProviderResponseBodyDataConfiguration {
	s.EntityId = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) SetLoginUrl(v string) *CreateIdentityProviderResponseBodyDataConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) SetSsoStatus(v string) *CreateIdentityProviderResponseBodyDataConfiguration {
	s.SsoStatus = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderResponseBodyDataSaml struct {
	// HTTP-Redirect or HTTP-POST. This parameter is required for manual entry. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// This parameter is required for manual entry. It must be non-empty and cannot exceed 2,048 characters. Existing configurations cannot be replaced. It is different from the Entity ID of the service provider (SP) returned in queries.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// This parameter is required for manual entry. It must be an absolute HTTPS URL and cannot contain user credentials or fragments.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// This parameter is required for manual entry and must be a non-empty array. Use PEM X.509 certificate text. The current validity period is verified. A maximum of 10 certificates are supported.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s CreateIdentityProviderResponseBodyDataSaml) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBodyDataSaml) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBodyDataSaml) GetBindingType() *string {
	return s.BindingType
}

func (s *CreateIdentityProviderResponseBodyDataSaml) GetEntityId() *string {
	return s.EntityId
}

func (s *CreateIdentityProviderResponseBodyDataSaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *CreateIdentityProviderResponseBodyDataSaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *CreateIdentityProviderResponseBodyDataSaml) SetBindingType(v string) *CreateIdentityProviderResponseBodyDataSaml {
	s.BindingType = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataSaml) SetEntityId(v string) *CreateIdentityProviderResponseBodyDataSaml {
	s.EntityId = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataSaml) SetLoginUrl(v string) *CreateIdentityProviderResponseBodyDataSaml {
	s.LoginUrl = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataSaml) SetX509Certificates(v []*string) *CreateIdentityProviderResponseBodyDataSaml {
	s.X509Certificates = v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataSaml) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderResponseBodyDataSyncScope struct {
	// When the mode is SelectedDepartments, at least one non-empty department ID must be included, which should be obtained from ListIdentityProviderDepartments. When the mode is AllAuthorized, pass an empty array or omit this parameter.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// Valid values:
	//
	// - AllAuthorized: The full authorization scope, which is limited by the application address book authorization.
	//
	// - SelectedDepartments: The selected departments and all their sub-departments.
	//
	// When you explicitly submit syncScope, you must provide a valid mode, which is verified by the backend.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s CreateIdentityProviderResponseBodyDataSyncScope) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponseBodyDataSyncScope) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponseBodyDataSyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *CreateIdentityProviderResponseBodyDataSyncScope) GetMode() *string {
	return s.Mode
}

func (s *CreateIdentityProviderResponseBodyDataSyncScope) SetDepartmentIds(v []*string) *CreateIdentityProviderResponseBodyDataSyncScope {
	s.DepartmentIds = v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataSyncScope) SetMode(v string) *CreateIdentityProviderResponseBodyDataSyncScope {
	s.Mode = &v
	return s
}

func (s *CreateIdentityProviderResponseBodyDataSyncScope) Validate() error {
	return dara.Validate(s)
}
