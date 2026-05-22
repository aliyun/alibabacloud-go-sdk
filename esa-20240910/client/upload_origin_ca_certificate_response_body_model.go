// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOriginCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonName(v string) *UploadOriginCaCertificateResponseBody
	GetCommonName() *string
	SetFingerprintSha256(v string) *UploadOriginCaCertificateResponseBody
	GetFingerprintSha256() *string
	SetId(v string) *UploadOriginCaCertificateResponseBody
	GetId() *string
	SetIssuer(v string) *UploadOriginCaCertificateResponseBody
	GetIssuer() *string
	SetNotAfter(v string) *UploadOriginCaCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *UploadOriginCaCertificateResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *UploadOriginCaCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *UploadOriginCaCertificateResponseBody
	GetSerialNumber() *string
	SetSignatureAlgorithm(v string) *UploadOriginCaCertificateResponseBody
	GetSignatureAlgorithm() *string
	SetStatus(v string) *UploadOriginCaCertificateResponseBody
	GetStatus() *string
	SetValidityDays(v string) *UploadOriginCaCertificateResponseBody
	GetValidityDays() *string
}

type UploadOriginCaCertificateResponseBody struct {
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
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
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
	// Indicates whether the operation is successful.
	//
	// 	- OK
	//
	// 	- Fail
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

func (s UploadOriginCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadOriginCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadOriginCaCertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *UploadOriginCaCertificateResponseBody) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *UploadOriginCaCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *UploadOriginCaCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *UploadOriginCaCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *UploadOriginCaCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *UploadOriginCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadOriginCaCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UploadOriginCaCertificateResponseBody) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *UploadOriginCaCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UploadOriginCaCertificateResponseBody) GetValidityDays() *string {
	return s.ValidityDays
}

func (s *UploadOriginCaCertificateResponseBody) SetCommonName(v string) *UploadOriginCaCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetFingerprintSha256(v string) *UploadOriginCaCertificateResponseBody {
	s.FingerprintSha256 = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetId(v string) *UploadOriginCaCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetIssuer(v string) *UploadOriginCaCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetNotAfter(v string) *UploadOriginCaCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetNotBefore(v string) *UploadOriginCaCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetRequestId(v string) *UploadOriginCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetSerialNumber(v string) *UploadOriginCaCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetSignatureAlgorithm(v string) *UploadOriginCaCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetStatus(v string) *UploadOriginCaCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) SetValidityDays(v string) *UploadOriginCaCertificateResponseBody {
	s.ValidityDays = &v
	return s
}

func (s *UploadOriginCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
