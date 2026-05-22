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
	CommonName         *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	FingerprintSha256  *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer             *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	NotAfter           *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore          *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidityDays       *string `json:"ValidityDays,omitempty" xml:"ValidityDays,omitempty"`
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
