// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSMCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertExpireTime(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetCertExpireTime() *string
	SetCertIdentifier(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetCertIdentifier() *string
	SetCertName(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetCertName() *string
	SetCertOrg(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetCertOrg() *string
	SetCommonName(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetCommonName() *string
	SetEncryptCertificate(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetEncryptCertificate() *string
	SetRequestId(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetRequestId() *string
	SetSans(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetSans() *string
	SetSignCertificate(v string) *DescribeCdnSMCertificateDetailResponseBody
	GetSignCertificate() *string
}

type DescribeCdnSMCertificateDetailResponseBody struct {
	// The expiration time of the certificate. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-31T09:42:28Z
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 648****-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// DigiCert Inc
	CertOrg *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	// The common name.
	//
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The content of the encryption certificate.
	//
	// example:
	//
	// -BEGIN CERTIFICATE-----***-----END CERTIFICATE--
	EncryptCertificate *string `json:"EncryptCertificate,omitempty" xml:"EncryptCertificate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A7C69682-7F88-40DD-A198-10D0309E439D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The subdomain name.
	//
	// example:
	//
	// ***.example.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The content of the signature certificate.
	//
	// example:
	//
	// --BEGIN CERTIFICATE-----***-----END CERTIFICATE--
	SignCertificate *string `json:"SignCertificate,omitempty" xml:"SignCertificate,omitempty"`
}

func (s DescribeCdnSMCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetEncryptCertificate() *string {
	return s.EncryptCertificate
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetSans() *string {
	return s.Sans
}

func (s *DescribeCdnSMCertificateDetailResponseBody) GetSignCertificate() *string {
	return s.SignCertificate
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetCertExpireTime(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetCertIdentifier(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetCertName(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetCertOrg(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.CertOrg = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetCommonName(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.CommonName = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetEncryptCertificate(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.EncryptCertificate = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetRequestId(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetSans(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.Sans = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) SetSignCertificate(v string) *DescribeCdnSMCertificateDetailResponseBody {
	s.SignCertificate = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
