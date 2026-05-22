// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSiteOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonName(v string) *UploadSiteOriginClientCertificateResponseBody
	GetCommonName() *string
	SetFingerprintSha256(v string) *UploadSiteOriginClientCertificateResponseBody
	GetFingerprintSha256() *string
	SetId(v string) *UploadSiteOriginClientCertificateResponseBody
	GetId() *string
	SetIssuer(v string) *UploadSiteOriginClientCertificateResponseBody
	GetIssuer() *string
	SetNotAfter(v string) *UploadSiteOriginClientCertificateResponseBody
	GetNotAfter() *string
	SetNotBefore(v string) *UploadSiteOriginClientCertificateResponseBody
	GetNotBefore() *string
	SetRequestId(v string) *UploadSiteOriginClientCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *UploadSiteOriginClientCertificateResponseBody
	GetSerialNumber() *string
	SetSignatureAlgorithm(v string) *UploadSiteOriginClientCertificateResponseBody
	GetSignatureAlgorithm() *string
	SetStatus(v string) *UploadSiteOriginClientCertificateResponseBody
	GetStatus() *string
	SetValidityDays(v string) *UploadSiteOriginClientCertificateResponseBody
	GetValidityDays() *string
}

type UploadSiteOriginClientCertificateResponseBody struct {
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
	// 1dc5fc9af4eead2570c70d94b416130baeb6d4429b51fd3557379588456aca**
	FingerprintSha256 *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	// The certificate ID on ESA.
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
	// 2024-12-01T02:12:49Z
	NotAfter *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The time when the certificate takes effect.
	//
	// example:
	//
	// 2023-12-01T02:13:07Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d**
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

func (s UploadSiteOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadSiteOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetFingerprintSha256() *string {
	return s.FingerprintSha256
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetNotAfter() *string {
	return s.NotAfter
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetNotBefore() *string {
	return s.NotBefore
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetSignatureAlgorithm() *string {
	return s.SignatureAlgorithm
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UploadSiteOriginClientCertificateResponseBody) GetValidityDays() *string {
	return s.ValidityDays
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetCommonName(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.CommonName = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetFingerprintSha256(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.FingerprintSha256 = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetId(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetIssuer(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.Issuer = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetNotAfter(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.NotAfter = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetNotBefore(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.NotBefore = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetRequestId(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetSerialNumber(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetSignatureAlgorithm(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.SignatureAlgorithm = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetStatus(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.Status = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) SetValidityDays(v string) *UploadSiteOriginClientCertificateResponseBody {
	s.ValidityDays = &v
	return s
}

func (s *UploadSiteOriginClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
