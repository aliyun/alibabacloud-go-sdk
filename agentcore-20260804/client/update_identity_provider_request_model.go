// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdentityProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateIdentityProviderRequestBody) *UpdateIdentityProviderRequest
	GetBody() *UpdateIdentityProviderRequestBody
}

type UpdateIdentityProviderRequest struct {
	// Submit business parameters as described in this operation. The body parameters of public APIs are mapped to backend wrappers by the gateway. Do not mix them with backend direct connection body wrappers.
	Body *UpdateIdentityProviderRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateIdentityProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequest) GetBody() *UpdateIdentityProviderRequestBody {
	return s.Body
}

func (s *UpdateIdentityProviderRequest) SetBody(v *UpdateIdentityProviderRequestBody) *UpdateIdentityProviderRequest {
	s.Body = v
	return s
}

func (s *UpdateIdentityProviderRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateIdentityProviderRequestBody struct {
	// Do not pass true during CustomSSO initialization. A valid SAML configuration must be saved before enabling. Disabling does not delete the configuration. In case of asynchronous failure, this expected value does not mean that it has taken effect upstream.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// The provider configuration for DingTalk and Lark supports partial updates. Omitted fields retain the saved values. Submitted values cannot be null, empty strings, or whitespace strings. The complete configuration is verified after merging. CustomSSO updates use protocol and saml, and do not accept this field. The returned value does not contain provider keys.
	//
	// example:
	//
	// {"appSecret":"example-app-secret"}
	Metadata *UpdateIdentityProviderRequestBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The logon protocol for CustomSSO. Only SAML is supported. You can pass SAML when saving the SAML configuration. You can omit this field if you only modify loginEnabled. DingTalk and Lark do not accept this field.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The SAML configuration for CustomSSO. Choose either metadataXml or the four manual entry fields: entityId, loginUrl, bindingType, and x509Certificates. If not submitted, the existing configuration is retained. DingTalk and Lark do not accept this field.
	Saml *UpdateIdentityProviderRequestBodySaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// Specifies whether to enable incremental synchronization for subsequent directory change events of DingTalk or Lark. If not passed, the original value is retained. It does not control the initial synchronization or manual synchronization. CustomSSO does not accept this field. You must use UpdateIdentityProviderSCIM.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// Follows the definition of the department-based synchronization scope. If omitted during creation, the default is AllAuthorized. The initial provider synchronization is determined by syncOnBind. If omitted or set to null during an update, the original value is retained. Updating and saving the scope does not automatically execute a full synchronization.
	SyncScope *UpdateIdentityProviderRequestBodySyncScope `json:"syncScope,omitempty" xml:"syncScope,omitempty" type:"Struct"`
}

func (s UpdateIdentityProviderRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestBody) GetLoginEnabled() *bool {
	return s.LoginEnabled
}

func (s *UpdateIdentityProviderRequestBody) GetMetadata() *UpdateIdentityProviderRequestBodyMetadata {
	return s.Metadata
}

func (s *UpdateIdentityProviderRequestBody) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateIdentityProviderRequestBody) GetSaml() *UpdateIdentityProviderRequestBodySaml {
	return s.Saml
}

func (s *UpdateIdentityProviderRequestBody) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *UpdateIdentityProviderRequestBody) GetSyncScope() *UpdateIdentityProviderRequestBodySyncScope {
	return s.SyncScope
}

func (s *UpdateIdentityProviderRequestBody) SetLoginEnabled(v bool) *UpdateIdentityProviderRequestBody {
	s.LoginEnabled = &v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetMetadata(v *UpdateIdentityProviderRequestBodyMetadata) *UpdateIdentityProviderRequestBody {
	s.Metadata = v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetProtocol(v string) *UpdateIdentityProviderRequestBody {
	s.Protocol = &v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetSaml(v *UpdateIdentityProviderRequestBodySaml) *UpdateIdentityProviderRequestBody {
	s.Saml = v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetSyncEnabled(v bool) *UpdateIdentityProviderRequestBody {
	s.SyncEnabled = &v
	return s
}

func (s *UpdateIdentityProviderRequestBody) SetSyncScope(v *UpdateIdentityProviderRequestBodySyncScope) *UpdateIdentityProviderRequestBody {
	s.SyncScope = v
	return s
}

func (s *UpdateIdentityProviderRequestBody) Validate() error {
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

type UpdateIdentityProviderRequestBodyMetadata struct {
	// The Lark App ID. It can be omitted during an update to retain the saved value. The merged Lark configuration of the old and new values must contain this item.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The DingTalk AppKey. It can be omitted during an update to retain the saved value. The merged DingTalk configuration of the old and new values must contain this item.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// The provider application secret. If omitted during an update, the saved value is retained. You can submit this field separately to update the secret. The explicit value cannot be empty. It is sensitive input and is not echoed in the response.
	//
	// example:
	//
	// example-app-secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// The DingTalk enterprise CorpId. It can be omitted during an update to retain the saved value. The merged DingTalk configuration of the old and new values must contain this item.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// Required when synchronization is enabled for DingTalk and Lark. It is sensitive input and is unrelated to the SCIM token.
	//
	// example:
	//
	// example-encrypt-key
	EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
	// Required when synchronization is enabled for DingTalk and Lark. It is sensitive input and is unrelated to the SCIM token.
	//
	// example:
	//
	// example-verification-token
	VerificationToken *string `json:"verificationToken,omitempty" xml:"verificationToken,omitempty"`
}

func (s UpdateIdentityProviderRequestBodyMetadata) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestBodyMetadata) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetAppId() *string {
	return s.AppId
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetAppKey() *string {
	return s.AppKey
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetAppSecret() *string {
	return s.AppSecret
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetCorpId() *string {
	return s.CorpId
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetEncryptKey() *string {
	return s.EncryptKey
}

func (s *UpdateIdentityProviderRequestBodyMetadata) GetVerificationToken() *string {
	return s.VerificationToken
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetAppId(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.AppId = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetAppKey(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.AppKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetAppSecret(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.AppSecret = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetCorpId(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.CorpId = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetEncryptKey(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.EncryptKey = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) SetVerificationToken(v string) *UpdateIdentityProviderRequestBodyMetadata {
	s.VerificationToken = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodyMetadata) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestBodySaml struct {
	// HTTP-Redirect or HTTP-POST. Required for manual entry. The query returns the short name.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// Required for manual entry. It must be non-empty and contain up to 2,048 characters. The existing configuration cannot be replaced. It is different from the local entity ID queried by the SP.
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
	// The input is the XML text of a single IdP EntityDescriptor. The maximum size is 1 MiB. It cannot be a URL or file path. It is mutually exclusive with the four manual entry fields. DTD and external entities are prohibited. If both SSO bindings are declared at the same time, HTTP-POST takes precedence.
	//
	// example:
	//
	// <EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://idp.example.com/entity"><IDPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><KeyDescriptor use="signing"><KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#"><X509Data><X509Certificate>MIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQELBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoXDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8oGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYNW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbYEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Okkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0EBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUWld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImcaWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTqP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjocJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdVjfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5</X509Certificate></X509Data></KeyInfo></KeyDescriptor><SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://idp.example.com/sso"/></IDPSSODescriptor></EntityDescriptor>
	MetadataXml *string `json:"metadataXml,omitempty" xml:"metadataXml,omitempty"`
	// Required for manual entry. It must be a non-empty array using PEM X.509 certificate text. The current validity period is verified. Up to two certificates are allowed after deduplication.
	//
	// example:
	//
	// ["-----BEGIN CERTIFICATE-----\\\\nMIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQEL\\\\nBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoX\\\\nDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjAN\\\\nBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8\\\\noGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYN\\\\nW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbY\\\\nEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Ok\\\\nkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0E\\\\nBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQAB\\\\nMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUW\\\\nld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImca\\\\nWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTq\\\\nP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjoc\\\\nJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdV\\\\njfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5\\\\n-----END CERTIFICATE-----"]
	X509Certificates []*string `json:"x509Certificates,omitempty" xml:"x509Certificates,omitempty" type:"Repeated"`
}

func (s UpdateIdentityProviderRequestBodySaml) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestBodySaml) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestBodySaml) GetBindingType() *string {
	return s.BindingType
}

func (s *UpdateIdentityProviderRequestBodySaml) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateIdentityProviderRequestBodySaml) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *UpdateIdentityProviderRequestBodySaml) GetMetadataXml() *string {
	return s.MetadataXml
}

func (s *UpdateIdentityProviderRequestBodySaml) GetX509Certificates() []*string {
	return s.X509Certificates
}

func (s *UpdateIdentityProviderRequestBodySaml) SetBindingType(v string) *UpdateIdentityProviderRequestBodySaml {
	s.BindingType = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodySaml) SetEntityId(v string) *UpdateIdentityProviderRequestBodySaml {
	s.EntityId = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodySaml) SetLoginUrl(v string) *UpdateIdentityProviderRequestBodySaml {
	s.LoginUrl = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodySaml) SetMetadataXml(v string) *UpdateIdentityProviderRequestBodySaml {
	s.MetadataXml = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodySaml) SetX509Certificates(v []*string) *UpdateIdentityProviderRequestBodySaml {
	s.X509Certificates = v
	return s
}

func (s *UpdateIdentityProviderRequestBodySaml) Validate() error {
	return dara.Validate(s)
}

type UpdateIdentityProviderRequestBodySyncScope struct {
	// For SelectedDepartments, it must contain at least one non-empty department ID, which should be obtained from ListIdentityProviderDepartments. For AllAuthorized, pass an empty array or omit it.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// AllAuthorized: All authorized scopes, subject to application directory authorization restrictions. SelectedDepartments: Selected departments and all sub-departments. When explicitly submitting syncScope, a valid mode must be provided and verified by the backend.
	//
	// example:
	//
	// SelectedDepartments
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateIdentityProviderRequestBodySyncScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdentityProviderRequestBodySyncScope) GoString() string {
	return s.String()
}

func (s *UpdateIdentityProviderRequestBodySyncScope) GetDepartmentIds() []*string {
	return s.DepartmentIds
}

func (s *UpdateIdentityProviderRequestBodySyncScope) GetMode() *string {
	return s.Mode
}

func (s *UpdateIdentityProviderRequestBodySyncScope) SetDepartmentIds(v []*string) *UpdateIdentityProviderRequestBodySyncScope {
	s.DepartmentIds = v
	return s
}

func (s *UpdateIdentityProviderRequestBodySyncScope) SetMode(v string) *UpdateIdentityProviderRequestBodySyncScope {
	s.Mode = &v
	return s
}

func (s *UpdateIdentityProviderRequestBodySyncScope) Validate() error {
	return dara.Validate(s)
}
