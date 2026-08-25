// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalSAMLIdPCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListExternalSAMLIdPCertificatesResponseBody
	GetRequestId() *string
	SetSAMLIdPCertificates(v []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) *ListExternalSAMLIdPCertificatesResponseBody
	GetSAMLIdPCertificates() []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates
	SetTotalCounts(v int32) *ListExternalSAMLIdPCertificatesResponseBody
	GetTotalCounts() *int32
}

type ListExternalSAMLIdPCertificatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 400979BC-92EC-58B9-B47C-6913BD56A6FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SAML signing certificates.
	SAMLIdPCertificates []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates `json:"SAMLIdPCertificates,omitempty" xml:"SAMLIdPCertificates,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) GetSAMLIdPCertificates() []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	return s.SAMLIdPCertificates
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) SetRequestId(v string) *ListExternalSAMLIdPCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) SetSAMLIdPCertificates(v []*ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) *ListExternalSAMLIdPCertificatesResponseBody {
	s.SAMLIdPCertificates = v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) SetTotalCounts(v int32) *ListExternalSAMLIdPCertificatesResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBody) Validate() error {
	if s.SAMLIdPCertificates != nil {
		for _, item := range s.SAMLIdPCertificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates struct {
	// The ID of the certificate.
	//
	// example:
	//
	// idp-c-00dt9gnl7fmjaw9c****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// 1.2.840.113549.1.9.1=#160d696e666f406f6b74612e63****,CN=dev-xxxxxx,OU=SSOProvider,O=Okta,L=San Francisco,ST=California,C=US
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate expires.
	//
	// example:
	//
	// 2030-06-23T07:04:37Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 2020-06-23T07:03:37Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The public key of the certificate. The value of this parameter is in the PEM format and is Base64-encoded.
	//
	// example:
	//
	// MIIBIjANBgkqhkiG****
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 159289587****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256withRSA
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The subject of the certificate.
	//
	// example:
	//
	// 1.2.840.113549.1.9.1=#160d696e666f406f6b74612e63****,CN=dev-xxxxxx,OU=SSOProvider,O=Okta,L=San Francisco,ST=California,C=US
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The version of the certificate.
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The X.509 certificate in the PEM format.
	//
	// example:
	//
	// MIIDpDCCAoygAwIBAgIG****
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) String() string {
	return dara.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetIssuer() *string {
	return s.Issuer
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetNotAfter() *string {
	return s.NotAfter
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetSubject() *string {
	return s.Subject
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetVersion() *int32 {
	return s.Version
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetCertificateId(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.CertificateId = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetIssuer(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.Issuer = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetNotAfter(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.NotAfter = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetNotBefore(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.NotBefore = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetPublicKey(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.PublicKey = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetSerialNumber(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.SerialNumber = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetSignatureAlgorithm(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.SignatureAlgorithm = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetSubject(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.Subject = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetVersion(v int32) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.Version = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) SetX509Certificate(v string) *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates {
	s.X509Certificate = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponseBodySAMLIdPCertificates) Validate() error {
	return dara.Validate(s)
}
