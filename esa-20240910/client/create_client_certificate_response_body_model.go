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
	CACertificateId    *string `json:"CACertificateId,omitempty" xml:"CACertificateId,omitempty"`
	Certificate        *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CommonName         *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	FingerprintSha256  *string `json:"FingerprintSha256,omitempty" xml:"FingerprintSha256,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Issuer             *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	NotAfter           *string `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	NotBefore          *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	PrivateKey         *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SerialNumber       *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	Status             *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ValidityDays       *string `json:"ValidityDays,omitempty" xml:"ValidityDays,omitempty"`
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
