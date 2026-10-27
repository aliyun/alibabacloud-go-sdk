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
	// The external identity provider binding information after the disconnection operation is accepted.
	Data *DeleteIdentityProviderResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Currently, only SAML is supported. This parameter may not be returned if no valid configuration is saved. OIDC placeholders are not supported for initialization.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The input supports either XML or complete manual configuration. Queries return the standardized configuration and do not echo the original XML. This parameter can be omitted when not configured.
	Saml *DeleteIdentityProviderResponseBodyDataSaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// The status. Valid values: CONFIGURED (the configuration is accepted and the system is waiting to activate the user pool), SYNCING (synchronizing organization members), SYNCED (organization member synchronization is complete), READY (the binding is effective), SYNC_FAILED (organization member synchronization failed), UPDATING (updating the configuration), UPDATE_FAILED (configuration update failed), DISCONNECTING (disconnecting), and DISCONNECT_FAILED (disconnection failed).
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
	if s.Saml != nil {
		if err := s.Saml.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteIdentityProviderResponseBodyDataConfiguration struct {
	// HTTP-Redirect or HTTP-POST. This parameter is required when manually filled in. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// This parameter is required when manually filled in. It must be non-empty and cannot exceed 2,048 characters. Existing configurations do not support replacement. This is different from the Entity ID of the SP queried on our side.
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
	// The actual switch value, which is Enabled or Disabled. This is not equivalent to the expected value of loginEnabled saved in AC.
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

type DeleteIdentityProviderResponseBodyDataSaml struct {
	// HTTP-Redirect or HTTP-POST. This parameter is required when manually filled in. Queries return the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// This parameter is required when manually filled in. It must be non-empty and cannot exceed 2,048 characters. Existing configurations do not support replacement. This is different from the Entity ID of the SP queried on our side.
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
	// This parameter is required when manually filled in. It must be a non-empty array and use PEM X.509 certificate text. The current validity period is verified. A maximum of two certificates are allowed after deduplication.
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
