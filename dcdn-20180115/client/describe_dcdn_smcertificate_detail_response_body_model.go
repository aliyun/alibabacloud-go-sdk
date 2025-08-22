// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSMCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertExpireTime(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetCertExpireTime() *string
	SetCertIdentifier(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetCertIdentifier() *string
	SetCertName(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetCertName() *string
	SetCertOrg(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetCertOrg() *string
	SetCommonName(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetCommonName() *string
	SetEncryptCertificate(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetEncryptCertificate() *string
	SetRequestId(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetRequestId() *string
	SetSans(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetSans() *string
	SetSignCertificate(v string) *DescribeDcdnSMCertificateDetailResponseBody
	GetSignCertificate() *string
}

type DescribeDcdnSMCertificateDetailResponseBody struct {
	// The time when the certificate expires. The time is displayed in UTC.
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
	// The top-level domain name.
	//
	// example:
	//
	// example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The content of the encryption certificate.
	//
	// example:
	//
	// --BEGIN CERTIFICATE-----***-----END CERTIFICATE--
	EncryptCertificate *string `json:"EncryptCertificate,omitempty" xml:"EncryptCertificate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A7C69682-7F88-40DD-A198-10D0309E439D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The domain name on the additional certificate.
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

func (s DescribeDcdnSMCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetEncryptCertificate() *string {
	return s.EncryptCertificate
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetSans() *string {
	return s.Sans
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) GetSignCertificate() *string {
	return s.SignCertificate
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertExpireTime(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertIdentifier(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertName(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCertOrg(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CertOrg = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetCommonName(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.CommonName = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetEncryptCertificate(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.EncryptCertificate = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetRequestId(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetSans(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.Sans = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) SetSignCertificate(v string) *DescribeDcdnSMCertificateDetailResponseBody {
	s.SignCertificate = &v
	return s
}

func (s *DescribeDcdnSMCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
