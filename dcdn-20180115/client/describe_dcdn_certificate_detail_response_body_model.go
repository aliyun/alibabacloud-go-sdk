// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeDcdnCertificateDetailResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeDcdnCertificateDetailResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeDcdnCertificateDetailResponseBody
	GetCertName() *string
	SetKey(v string) *DescribeDcdnCertificateDetailResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeDcdnCertificateDetailResponseBody
	GetRequestId() *string
}

type DescribeDcdnCertificateDetailResponseBody struct {
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----xxx-----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 123
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// 123
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The key of the certificate.
	//
	// example:
	//
	// ak1htyxxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7C69682-7F88-40DD-A198-10D0309E439B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateDetailResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeDcdnCertificateDetailResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeDcdnCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnCertificateDetailResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeDcdnCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetCert(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetCertId(v int64) *DescribeDcdnCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetCertName(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetKey(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) SetRequestId(v string) *DescribeDcdnCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
