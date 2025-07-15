// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeLiveCertificateDetailResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeLiveCertificateDetailResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeLiveCertificateDetailResponseBody
	GetCertName() *string
	SetRequestId(v string) *DescribeLiveCertificateDetailResponseBody
	GetRequestId() *string
}

type DescribeLiveCertificateDetailResponseBody struct {
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----****-----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 23451111
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// Cert-****
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C7C69682-7F88-40DD-A198-10D0309E439B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateDetailResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeLiveCertificateDetailResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeLiveCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveCertificateDetailResponseBody) SetCert(v string) *DescribeLiveCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponseBody) SetCertId(v int64) *DescribeLiveCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponseBody) SetCertName(v string) *DescribeLiveCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponseBody) SetRequestId(v string) *DescribeLiveCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
