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
	// The configuration fields to be modified. For DingTalk and Feishu, you must submit at least one of the following fields: loginEnabled, syncEnabled, metadata, and syncScope. For CustomSSO, use the protocol, saml, and loginEnabled fields. You cannot submit the vendor metadata, syncEnabled, or syncScope fields.
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
	// Specifies whether to allow logon by using this identity provider. If this field is omitted or set to null during an update, the original value is retained. For CustomSSO, a valid SAML configuration must exist before enabling logon. Disabling logon does not delete the saved SAML or SCIM configuration.
	LoginEnabled *bool `json:"loginEnabled,omitempty" xml:"loginEnabled,omitempty"`
	// Applies only to DingTalk and Feishu. The application configuration is merged by field. Omitted fields and existing keys retain their original values. Explicitly submitted fields must not be null, empty strings, or blank strings. The merged configuration must meet the complete configuration requirements of the vendor. This field is not accepted for CustomSSO.
	//
	// example:
	//
	// {"appSecret":"example-app-secret"}
	Metadata *UpdateIdentityProviderRequestBodyMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The logon protocol, which applies only to CustomSSO. Currently, only SAML is supported. This field is not accepted for DingTalk and Feishu. When submitted with a non-empty saml configuration, this field can be omitted and is processed as SAML by the server.
	//
	// example:
	//
	// SAML
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// Applies only to CustomSSO. You must choose between providing metadataXml and providing the four complete manually entered configuration fields. The two methods cannot be mixed. If a valid configuration already exists, you can also modify only the JIT switches. Saving the configuration does not automatically enable logon.
	Saml *UpdateIdentityProviderRequestBodySaml `json:"saml,omitempty" xml:"saml,omitempty" type:"Struct"`
	// Applies only to DingTalk and Feishu. This field controls the incremental synchronization of subsequent address book change events. If this field is omitted or set to null during an update, the original value is retained. This field does not control the initial synchronization or manual synchronization. When enabled, the merged metadata must include the encryptKey and verificationToken fields. For CustomSSO, use UpdateIdentityProviderSCIM to manage the synchronization switch.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
	// The address book synchronization scope, which applies only to DingTalk and Feishu. If this field is omitted or set to null, the current scope is retained. Saving the scope does not automatically trigger a full synchronization. This field is not accepted for CustomSSO.
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
	// The App ID of the Feishu application. If this field is omitted during an update, the existing value is retained. If explicitly submitted, the value must be non-empty. The merged configuration for Feishu must include this field.
	//
	// example:
	//
	// cli_exampleappid01
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The AppKey of the DingTalk application. If this field is omitted during an update, the existing value is retained. If explicitly submitted, the value must be non-empty. The merged configuration for DingTalk must include this field.
	//
	// example:
	//
	// dingexampleappkey01
	AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty"`
	// The AppSecret of the DingTalk or Feishu application. If this field is omitted during an update, the existing secret is retained and does not need to be retransmitted each time. If explicitly submitted, the value must be non-empty, and the merged configuration must include this field. This is a sensitive input and is not echoed in queries. The example is a fictitious value.
	//
	// example:
	//
	// example-app-secret
	AppSecret *string `json:"appSecret,omitempty" xml:"appSecret,omitempty"`
	// The CorpId of the organization to which the DingTalk application belongs. If this field is omitted during an update, the existing value is retained. If explicitly submitted, the value must be non-empty and consistent with the organization to which the application belongs. The merged configuration for DingTalk must include this field.
	//
	// example:
	//
	// dingexamplecorpid01
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The encryption key for vendor event callbacks, which must be consistent with the vendor application configuration. If this field is omitted during an update, the existing value is retained. When incremental synchronization is enabled, the merged configuration must include this field. This is a sensitive input and is not echoed in queries.
	//
	// example:
	//
	// example-encrypt-key
	EncryptKey *string `json:"encryptKey,omitempty" xml:"encryptKey,omitempty"`
	// The verification token for vendor event callbacks, which must be consistent with the vendor application configuration. If this field is omitted during an update, the existing value is retained. When incremental synchronization is enabled, the merged configuration must include this field. This is a sensitive input and is not echoed in queries.
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
	// The binding type when manually entering the SAML configuration. Valid values: HTTP-Redirect and HTTP-POST. This field must be submitted together with entityId, loginUrl, and x509Certificates, and cannot be mixed with metadataXml.
	//
	// example:
	//
	// HTTP-Redirect
	BindingType *string `json:"bindingType,omitempty" xml:"bindingType,omitempty"`
	// The entity identifier of the enterprise IdP. This field is required when manually entering the configuration. The value must be non-empty and contain up to 2,048 characters, and cannot be changed after saving. This field must be submitted together with loginUrl, bindingType, and x509Certificates, and cannot be mixed with metadataXml.
	//
	// example:
	//
	// https://idp.example.com/entity
	EntityId *string `json:"entityId,omitempty" xml:"entityId,omitempty"`
	// The SSO logon URL of the enterprise IdP. This field is required when manually entering the configuration. The value must be an absolute HTTPS URL without user credentials and fragments. This field must be submitted together with entityId, bindingType, and x509Certificates, and cannot be mixed with metadataXml.
	//
	// example:
	//
	// https://idp.example.com/sso
	LoginUrl *string `json:"loginUrl,omitempty" xml:"loginUrl,omitempty"`
	// The original SAML metadata XML text of the enterprise IdP, not a URL or file path. The maximum size is 1 MiB. The text must contain a single EntityDescriptor. DTD and external entities are prohibited. This field cannot be mixed with the four manually entered configuration fields. After saving, queries return the standardized configuration and do not echo the original XML text.
	//
	// example:
	//
	// <EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://idp.example.com/entity"><IDPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"><KeyDescriptor use="signing"><KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#"><X509Data><X509Certificate>MIICwDCCAaigAwIBAgIULlsHLfcADyBax5ZtFlj+YGzZVo0wDQYJKoZIhvcNAQELBQAwGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMB4XDTI2MDEwMTAwMDAwMFoXDTM2MDEwMTAwMDAwMFowGjEYMBYGA1UEAwwPaWRwLmV4YW1wbGUuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAto0DN0W1CrPRzK9QZ9A1aApb4Lg8oGlp2c0PbN3LokXe/HzwxCMXx65IDnrenttxzTwp9BgUnUN/3LRzWaOL2UQ/vNYNW9LFLkI1T64UPyI7eyi28YlALhb41xkYhEp7rz+MjqifBzVZeTDUxBYtbK2hAhbYEL1+udwsXEx0KB7xuQ/rbI2OKeqAEAD1tfYg7T/SM2wrXNYhx+o1HtchM4/0I8Okkk27JQ4y3ofzC5sTySmFF1Y7XcpSd+5WSyazedhkAieo8AB01DkJV6nnKfH8sH0EBsnH3FP8FhqCLrIeEkjX9B9OcIRLGjiZqx4u/ryW1QhMldnv+pFfNjWy/wIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQC1hSQG8AwtGOoX2gKBK4Hi5XBrBreij8WlLTUWld0FINQ0Rg84f/27CjDthywj6wgxCesCuOVFhpxQa1UacsP/Gcai1aKABtLWImcaWECRYK5gDkwhTMFKM9mXA26hH1l4v/UR2iTpEK9H78egC4W7bs2LDTpPryuaNKTqP2O0tMbNkAHB8yecz5LxHXnxadL9Egow25g2RClvpifWqDnuS9TSbKkIWLcPLjocJj2K9ILpBwhXiXi/BiSgyYTLiYvVvih9vX4s+hKGZNJc0j92YLiOHDF0XH1oxgdVjfTmRiedeb2VZJZElptD92z0AqFR7z2d5B+OqpPA2gjuMLV5</X509Certificate></X509Data></KeyInfo></KeyDescriptor><SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://idp.example.com/sso"/></IDPSSODescriptor></EntityDescriptor>
	MetadataXml *string `json:"metadataXml,omitempty" xml:"metadataXml,omitempty"`
	// The list of X.509 signing certificates of the enterprise IdP. When manually entering the configuration, you must provide a non-empty list of up to 10 certificates, and the certificates must be valid. This field must be submitted together with entityId, loginUrl, and bindingType, and cannot be mixed with metadataXml.
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
	// The list of department IDs in the synchronization scope, which are obtained from ListIdentityProviderDepartments. In SelectedDepartments mode, you must provide at least one non-empty department ID. In AllAuthorized mode, this field should be omitted or an empty array should be passed.
	//
	// example:
	//
	// ["10001"]
	DepartmentIds []*string `json:"departmentIds,omitempty" xml:"departmentIds,omitempty" type:"Repeated"`
	// The synchronization scope mode. Valid values: AllAuthorized and SelectedDepartments. AllAuthorized indicates all members within the authorization scope of the application address book. SelectedDepartments indicates the selected departments and all their sub-departments. When explicitly submitting syncScope, you must provide a valid mode.
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
