// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadClientCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonName(v string) *UploadClientCaCertificateResponseBody
	GetCommonName() *string
	SetFingerprintSha256(v string) *UploadClientCaCertificateResponseBody
	GetFingerprintSha256() *string
	SetId(v string) *UploadClientCaCertificateResponseBody
	GetId() *string
	SetIssuer(v string) *UploadClientCaCertificateResponseBody
	GetIssuer() *string
	SetNotAfter(v string) *UploadClientCaCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *UploadClientCaCertificateResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *UploadClientCaCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *UploadClientCaCertificateResponseBody
	GetSerialNumber() *string
	SetSignatureAlgorithm(v string) *UploadClientCaCertificateResponseBody
	GetSignatureAlgorithm() *string
	SetStatus(v string) *UploadClientCaCertificateResponseBody
	GetStatus() *string
	SetValidityDays(v string) *UploadClientCaCertificateResponseBody
	GetValidityDays() *string
}

type UploadClientCaCertificateResponseBody struct {
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
	// baba39055622c008b90285a8838ed09a
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
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// babab9db65ee5efcca9f3d41d4b50d66
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
	// 300
	ValidityDays *string `json:"ValidityDays,omitempty" xml:"ValidityDays,omitempty"`
}

func (s UploadClientCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadClientCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadClientCaCertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *UploadClientCaCertificateResponseBody) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *UploadClientCaCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *UploadClientCaCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *UploadClientCaCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *UploadClientCaCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *UploadClientCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadClientCaCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UploadClientCaCertificateResponseBody) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *UploadClientCaCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UploadClientCaCertificateResponseBody) GetValidityDays() *string {
	return s.ValidityDays
}

func (s *UploadClientCaCertificateResponseBody) SetCommonName(v string) *UploadClientCaCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetFingerprintSha256(v string) *UploadClientCaCertificateResponseBody {
	s.FingerprintSha256 = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetId(v string) *UploadClientCaCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetIssuer(v string) *UploadClientCaCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetNotAfter(v string) *UploadClientCaCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetNotBefore(v string) *UploadClientCaCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetRequestId(v string) *UploadClientCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetSerialNumber(v string) *UploadClientCaCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetSignatureAlgorithm(v string) *UploadClientCaCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetStatus(v string) *UploadClientCaCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) SetValidityDays(v string) *UploadClientCaCertificateResponseBody {
	s.ValidityDays = &v
	return s
}

func (s *UploadClientCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
