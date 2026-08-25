// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectorySAMLServiceProviderInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDirectorySAMLServiceProviderInfoResponseBody
	GetRequestId() *string
	SetSAMLServiceProvider(v *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) *GetDirectorySAMLServiceProviderInfoResponseBody
	GetSAMLServiceProvider() *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider
}

type GetDirectorySAMLServiceProviderInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4632107D-BCE1-5A96-B30B-182EE0709625
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SP.
	SAMLServiceProvider *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider `json:"SAMLServiceProvider,omitempty" xml:"SAMLServiceProvider,omitempty" type:"Struct"`
}

func (s GetDirectorySAMLServiceProviderInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) GetSAMLServiceProvider() *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	return s.SAMLServiceProvider
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) SetRequestId(v string) *GetDirectorySAMLServiceProviderInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) SetSAMLServiceProvider(v *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) *GetDirectorySAMLServiceProviderInfoResponseBody {
	s.SAMLServiceProvider = v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBody) Validate() error {
	if s.SAMLServiceProvider != nil {
		if err := s.SAMLServiceProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider struct {
	// The Assertion Consumer Service (ACS) URL of the SP.
	//
	// example:
	//
	// https://signin-cn-shanghai.alibabacloudsso.com/saml/acs/51d298a9-2a3f-4e23-97c7-7ad1cfa9****
	AcsUrl *string `json:"AcsUrl,omitempty" xml:"AcsUrl,omitempty"`
	// The signature algorithm supported by the AuthNRequest initiated by Alibaba Cloud. Value:
	//
	// - rsa-sha256
	//
	// - rsa-sha1
	//
	// example:
	//
	// rsa-sha256
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The certificate type used by Alibaba Cloud for signing during the SSO process. Value:
	//
	// - self-signed: Use a self-signed certificate.
	//
	// - public: Use a certificate issued by CA.
	//
	// example:
	//
	// public
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The metadata file of the SP. The value of this parameter is Base64-encoded.
	//
	// example:
	//
	// PD94bWwgdmVyc2lv****
	EncodedMetadataDocument *string `json:"EncodedMetadataDocument,omitempty" xml:"EncodedMetadataDocument,omitempty"`
	// The entity ID of the SP.
	//
	// example:
	//
	// https://signin-cn-shanghai.alibabacloudsso.com/saml/sp/d-00fc2p61****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// Whether to support Assertion encryption on the IdP side.
	//
	// example:
	//
	// true
	SupportEncryptedAssertion *bool `json:"SupportEncryptedAssertion,omitempty" xml:"SupportEncryptedAssertion,omitempty"`
}

func (s GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) String() string {
	return dara.Prettify(s)
}

func (s GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GoString() string {
	return s.String()
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetAcsUrl() *string {
	return s.AcsUrl
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetCertificateType() *string {
	return s.CertificateType
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetEncodedMetadataDocument() *string {
	return s.EncodedMetadataDocument
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetEntityId() *string {
	return s.EntityId
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) GetSupportEncryptedAssertion() *bool {
	return s.SupportEncryptedAssertion
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetAcsUrl(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.AcsUrl = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetAuthnSignAlgo(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.AuthnSignAlgo = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetCertificateType(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.CertificateType = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetDirectoryId(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetEncodedMetadataDocument(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.EncodedMetadataDocument = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetEntityId(v string) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.EntityId = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) SetSupportEncryptedAssertion(v bool) *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider {
	s.SupportEncryptedAssertion = &v
	return s
}

func (s *GetDirectorySAMLServiceProviderInfoResponseBodySAMLServiceProvider) Validate() error {
	return dara.Validate(s)
}
