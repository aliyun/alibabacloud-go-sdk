// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCACertificateId(v string) *CreateClientCertificateResponseBody
	GetCACertificateId() *string
	SetCertificate(v string) *CreateClientCertificateResponseBody
	GetCertificate() *string
	SetCommonName(v string) *CreateClientCertificateResponseBody
	GetCommonName() *string
	SetFingerprintSha256(v string) *CreateClientCertificateResponseBody
	GetFingerprintSha256() *string
	SetId(v string) *CreateClientCertificateResponseBody
	GetId() *string
	SetIssuer(v string) *CreateClientCertificateResponseBody
	GetIssuer() *string
	SetNotAfter(v string) *CreateClientCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *CreateClientCertificateResponseBody
	GetNotBefore() *string
	SetPrivateKey(v string) *CreateClientCertificateResponseBody
	GetPrivateKey() *string
	SetRequestId(v string) *CreateClientCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *CreateClientCertificateResponseBody
	GetSerialNumber() *string
	SetSignatureAlgorithm(v string) *CreateClientCertificateResponseBody
	GetSignatureAlgorithm() *string
	SetStatus(v string) *CreateClientCertificateResponseBody
	GetStatus() *string
	SetValidityDays(v string) *CreateClientCertificateResponseBody
	GetValidityDays() *string
}

type CreateClientCertificateResponseBody struct {
	// The ID of the CA certificate.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1dbb
	CACertificateId *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	// The certificate content.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The Common Name of the certificate.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456aca66
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The certificate ID on ESA.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d95
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The CA that issued the certificate.
	//
	// example:
	//
	// DCDN CA
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate expires.
	//
	// example:
	//
	// 2024-12-01T02:12:49Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate takes effect.
	//
	// example:
	//
	// 2023-12-01T02:12:49Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The private key of the certificate.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1daa
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256-RSA
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The status of the certificate.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The validity period of the certificate. Unit: day.
	//
	// example:
	//
	// 365
	ValidityDays *string `json:"ValidityDays,omitempty" xml:"ValidityDays,omitempty"`
}

func (s CreateClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateResponseBody) GetCACertificateId() *string {
	return s.CACertificateId
}

func (s *CreateClientCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateClientCertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateClientCertificateResponseBody) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *CreateClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateClientCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *CreateClientCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *CreateClientCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *CreateClientCertificateResponseBody) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CreateClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateClientCertificateResponseBody) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *CreateClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateClientCertificateResponseBody) GetValidityDays() *string {
	return s.ValidityDays
}

func (s *CreateClientCertificateResponseBody) SetCACertificateId(v string) *CreateClientCertificateResponseBody {
	s.CACertificateId = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetCertificate(v string) *CreateClientCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetCommonName(v string) *CreateClientCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetFingerprintSha256(v string) *CreateClientCertificateResponseBody {
	s.FingerprintSha256 = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetId(v string) *CreateClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetIssuer(v string) *CreateClientCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetNotAfter(v string) *CreateClientCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetNotBefore(v string) *CreateClientCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetPrivateKey(v string) *CreateClientCertificateResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetRequestId(v string) *CreateClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetSerialNumber(v string) *CreateClientCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetSignatureAlgorithm(v string) *CreateClientCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetStatus(v string) *CreateClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetValidityDays(v string) *CreateClientCertificateResponseBody {
	s.ValidityDays = &v
	return s
}

func (s *CreateClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
