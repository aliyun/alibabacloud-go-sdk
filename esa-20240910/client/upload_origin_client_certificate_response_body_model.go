// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonName(v string) *UploadOriginClientCertificateResponseBody
	GetCommonName() *string
	SetFingerprintSha256(v string) *UploadOriginClientCertificateResponseBody
	GetFingerprintSha256() *string
	SetId(v string) *UploadOriginClientCertificateResponseBody
	GetId() *string
	SetIssuer(v string) *UploadOriginClientCertificateResponseBody
	GetIssuer() *string
	SetNotAfter(v string) *UploadOriginClientCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *UploadOriginClientCertificateResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *UploadOriginClientCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *UploadOriginClientCertificateResponseBody
	GetSerialNumber() *string
	SetSignatureAlgorithm(v string) *UploadOriginClientCertificateResponseBody
	GetSignatureAlgorithm() *string
	SetStatus(v string) *UploadOriginClientCertificateResponseBody
	GetStatus() *string
	SetValidityDays(v string) *UploadOriginClientCertificateResponseBody
	GetValidityDays() *string
}

type UploadOriginClientCertificateResponseBody struct {
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
	// The certificate ID.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The CA that issued the certificate.
	//
	// example:
	//
	// GlobalSign nv-sa
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate expires.
	//
	// example:
	//
	// 2024-12-01T02:13:07Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate takes effect.
	//
	// example:
	//
	// 2023-12-01T02:13:07Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
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
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The validity period of the certificate. Unit: day.
	//
	// example:
	//
	// 365
	ValidityDays *string `json:"ValidityDays,omitempty" xml:"ValidityDays,omitempty"`
}

func (s UploadOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadOriginClientCertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *UploadOriginClientCertificateResponseBody) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *UploadOriginClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *UploadOriginClientCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *UploadOriginClientCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *UploadOriginClientCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *UploadOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadOriginClientCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UploadOriginClientCertificateResponseBody) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *UploadOriginClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UploadOriginClientCertificateResponseBody) GetValidityDays() *string {
	return s.ValidityDays
}

func (s *UploadOriginClientCertificateResponseBody) SetCommonName(v string) *UploadOriginClientCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetFingerprintSha256(v string) *UploadOriginClientCertificateResponseBody {
	s.FingerprintSha256 = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetId(v string) *UploadOriginClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetIssuer(v string) *UploadOriginClientCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetNotAfter(v string) *UploadOriginClientCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetNotBefore(v string) *UploadOriginClientCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetRequestId(v string) *UploadOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetSerialNumber(v string) *UploadOriginClientCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetSignatureAlgorithm(v string) *UploadOriginClientCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetStatus(v string) *UploadOriginClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) SetValidityDays(v string) *UploadOriginClientCertificateResponseBody {
	s.ValidityDays = &v
	return s
}

func (s *UploadOriginClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
