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
	// The business status code of the API request. SUCCESS indicates that the request is successful, but does not mean that the asynchronous configuration or synchronization task is completed.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The currently saved identity provider binding information, including the expected configuration and resource status. The actual update may still be executing in the background.
	Data *UpdateIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the API request, which does not represent the execution result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message of the API request. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request. Provide this value when you troubleshoot issues.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API request is successful. To check whether the background configuration has taken effect, query the identity provider status separately.
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
	// The actual SAML configuration read back from the upstream. This update operation does not perform the read-back and usually does not return this field. Use GetIdentityProvider to obtain it.
	Configuration *UpdateIdentityProviderResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The type of the identity provider. DingTalk: DingTalk. Feishu: Lark. CustomSSO: custom enterprise identity provider connected through SAML. The value must be the same as the type of the identity provider bound to the current workspace.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"identityProviderType,omitempty" xml:"identityProviderType,omitempty"`
	// The currently saved expected value of the logon switch, which alone does not prove that the upstream logon configuration has taken effect.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The currently saved logon protocol for CustomSSO, which is SAML when configured. This field is not returned for DingTalk or Lark.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The saved standardized SAML configuration for CustomSSO. The original metadata XML is not echoed. This field is not returned if no valid configuration is saved or if the identity provider is DingTalk or Lark. This value is the locally saved configuration and is not equivalent to the actual upstream configuration.
	Saml *UpdateIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The status of the identity provider resource and logon configuration. Valid values:
	//
	// - CONFIGURED: The configuration is accepted.
	//
	// - UPDATING: The configuration is being updated.
	//
	// - SYNCING: Synchronization is in progress.
	//
	// - SYNCED: The synchronization phase is completed.
	//
	// - READY: The resource is ready.
	//
	// - SYNC_FAILED: Synchronization failed.
	//
	// - UPDATE_FAILED: The update failed.
	//
	// - DISCONNECTING: Unbinding is in progress.
	//
	// - DISCONNECT_FAILED: Unbinding failed.
	//
	// READY alone does not indicate that users have been imported. Query the SCIM operation results for CustomSSO separately.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The currently saved expected value of the synchronization switch. For DingTalk and Lark, it is used for subsequent address book change events. The SCIM switch for CustomSSO is managed through UpdateIdentityProviderSCIM, and the operation results are queried through GetIdentityProviderSCIM.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The saved synchronization scope for DingTalk and Lark. The default value for valid historical configurations is AllAuthorized, and this field may not be returned for corrupted historical configurations. This field is not returned for CustomSSO.
	SyncScope *UpdateIdentityProviderResponseBodyDataSyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
	// The ID of the workspace to which the identity provider belongs.
	//
	// example:
	//
	// ws-123456
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
	// The actual SAML binding type saved in the upstream, which is HTTP-Redirect or HTTP-POST.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The actual enterprise IdP entity ID saved in the upstream.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The actual enterprise IdP SSO logon URL saved in the upstream.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The actual SSO logon status in the upstream. Enabled indicates that it is turned on, and Disabled indicates that it is turned off. This should be evaluated separately from the expected value of loginEnabled.
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

type UpdateIdentityProviderResponseBodyDataSaml struct {
	// The saved SAML binding type, which is HTTP-Redirect or HTTP-POST.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The saved enterprise IdP entity ID, which is different from the entity ID of the service provider (SP).
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The saved enterprise IdP SSO logon URL.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The list of saved enterprise IdP X.509 signing certificates.
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
	// The list of department IDs in the synchronization scope, obtained from ListIdentityProviderDepartments. In SelectedDepartments mode, at least one non-empty department ID must be provided. In AllAuthorized mode, this parameter should be omitted or an empty array should be passed.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// The synchronization scope mode. AllAuthorized: all members within the authorization scope of the application address book. SelectedDepartments: the selected departments and all their sub-departments. A valid mode must be provided when syncScope is explicitly submitted.
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
