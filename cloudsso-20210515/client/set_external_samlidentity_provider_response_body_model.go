// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetExternalSAMLIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetExternalSAMLIdentityProviderResponseBody
	GetRequestId() *string
	SetSAMLIdentityProviderConfiguration(v *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) *SetExternalSAMLIdentityProviderResponseBody
	GetSAMLIdentityProviderConfiguration() *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration
}

type SetExternalSAMLIdentityProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 63160579-2E1B-57B0-8273-B27427172385
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the IdP.
	SAMLIdentityProviderConfiguration *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration `json:"SAMLIdentityProviderConfiguration,omitempty" xml:"SAMLIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s SetExternalSAMLIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetExternalSAMLIdentityProviderResponseBody) GetSAMLIdentityProviderConfiguration() *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	return s.SAMLIdentityProviderConfiguration
}

func (s *SetExternalSAMLIdentityProviderResponseBody) SetRequestId(v string) *SetExternalSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBody) SetSAMLIdentityProviderConfiguration(v *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) *SetExternalSAMLIdentityProviderResponseBody {
	s.SAMLIdentityProviderConfiguration = v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBody) Validate() error {
	if s.SAMLIdentityProviderConfiguration != nil {
		if err := s.SAMLIdentityProviderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration struct {
	// The binding for sending SAML requests. Valid values:
	//
	// 	- Post: HTTP Post bindings.
	//
	// 	- Redirect: HTTP Redirect bindings.
	//
	// example:
	//
	// Redirect
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The IDs of the SAML signing certificates.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The time when the IdP was configured for the first time.
	//
	// example:
	//
	// 2021-11-10T02:57:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the IdP. The value of this parameter is Base64-encoded.
	//
	// example:
	//
	// PD94bWwgdmVyc2lvbj0iMS4****
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the IdP.
	//
	// example:
	//
	// http://www.okta.com/exk3qwgtjhetR2Od****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The logon URL of the IdP.
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// The status of SSO logon. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Disabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// The time when the IdP configurations were last modified.
	//
	// example:
	//
	// 2021-11-10T02:57:16Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Indicates whether CloudSSO needs to sign SAML requests. The requests are sent when users log on to the CloudSSO user portal to initiate SAML-based SSO. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	WantRequestSigned *bool `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
}

func (s SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetCertificateIds() []*string {
	return s.CertificateIds
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetEncodedMetadataDocument() *string {
	return s.EncodedMetadataDocument
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetWantRequestSigned() *bool {
	return s.WantRequestSigned
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetBindingType(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.BindingType = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCertificateIds(v []*string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CertificateIds = v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCreateTime(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CreateTime = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetDirectoryId(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEncodedMetadataDocument(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEntityId(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EntityId = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetLoginUrl(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetSSOStatus(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetUpdateTime(v string) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetWantRequestSigned(v bool) *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.WantRequestSigned = &v
	return s
}

func (s *SetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) Validate() error {
	return dara.Validate(s)
}
