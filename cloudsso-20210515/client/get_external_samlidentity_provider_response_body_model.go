// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalSAMLIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetExternalSAMLIdentityProviderResponseBody
	GetRequestId() *string
	SetSAMLIdentityProviderConfiguration(v *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) *GetExternalSAMLIdentityProviderResponseBody
	GetSAMLIdentityProviderConfiguration() *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration
}

type GetExternalSAMLIdentityProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 96D1E5FF-0301-5636-8D33-071E033CFB82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the IdP.
	SAMLIdentityProviderConfiguration *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration `json:"SAMLIdentityProviderConfiguration,omitempty" xml:"SAMLIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s GetExternalSAMLIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExternalSAMLIdentityProviderResponseBody) GetSAMLIdentityProviderConfiguration() *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	return s.SAMLIdentityProviderConfiguration
}

func (s *GetExternalSAMLIdentityProviderResponseBody) SetRequestId(v string) *GetExternalSAMLIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBody) SetSAMLIdentityProviderConfiguration(v *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) *GetExternalSAMLIdentityProviderResponseBody {
	s.SAMLIdentityProviderConfiguration = v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBody) Validate() error {
	if s.SAMLIdentityProviderConfiguration != nil {
		if err := s.SAMLIdentityProviderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration struct {
	// The binding for sending SAML requests. Valid values:
	//
	// 	- Post: HTTP Post bindings.
	//
	// 	- Redirect: HTTP Redirect bindings.
	//
	// example:
	//
	// Post
	BindingType *string `json:"BindingType,omitempty" xml:"BindingType,omitempty"`
	// The SAML signing certificates.
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	// The time when the IdP was configured for the first time.
	//
	// example:
	//
	// 2021-11-09T09:30:13Z
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
	// The ID of the IdP.
	//
	// example:
	//
	// http://www.okta.com/exk3qwgtjhetR2Od****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The logon URL of the IdP.
	LoginUrl *string `json:"LoginUrl,omitempty" xml:"LoginUrl,omitempty"`
	// Indicates whether SSO is enabled. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
	// The time when the IdP configurations were last modified.
	//
	// example:
	//
	// 2021-11-09T09:30:22Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Indicates whether CloudSSO needs to sign SAML requests. The requests are sent when users log on to the CloudSSO user portal to initiate SAML-based SSO. Valid values:
	//
	// 	- true:
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	WantRequestSigned *bool `json:"WantRequestSigned,omitempty" xml:"WantRequestSigned,omitempty"`
}

func (s GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetBindingType() *string {
	return s.BindingType
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetCertificateIds() []*string {
	return s.CertificateIds
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetEncodedMetadataDocument() *string {
	return s.EncodedMetadataDocument
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetEntityId() *string {
	return s.EntityId
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetLoginUrl() *string {
	return s.LoginUrl
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) GetWantRequestSigned() *bool {
	return s.WantRequestSigned
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetBindingType(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.BindingType = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCertificateIds(v []*string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CertificateIds = v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetCreateTime(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.CreateTime = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetDirectoryId(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.DirectoryId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEncodedMetadataDocument(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetEntityId(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.EntityId = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetLoginUrl(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.LoginUrl = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetSSOStatus(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetUpdateTime(v string) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.UpdateTime = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) SetWantRequestSigned(v bool) *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration {
	s.WantRequestSigned = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponseBodySAMLIdentityProviderConfiguration) Validate() error {
	return dara.Validate(s)
}
