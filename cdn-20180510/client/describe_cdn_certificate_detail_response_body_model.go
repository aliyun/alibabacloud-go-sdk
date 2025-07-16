// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeCdnCertificateDetailResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeCdnCertificateDetailResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeCdnCertificateDetailResponseBody
	GetCertName() *string
	SetKey(v string) *DescribeCdnCertificateDetailResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeCdnCertificateDetailResponseBody
	GetRequestId() *string
}

type DescribeCdnCertificateDetailResponseBody struct {
	// The certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIFzDCCBLSgAwIBxxxx
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 881049
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// cert-15480655xxxx
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The key of the SSL certificate.
	//
	// example:
	//
	// xxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeCdnCertificateDetailResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeCdnCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnCertificateDetailResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeCdnCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnCertificateDetailResponseBody) SetCert(v string) *DescribeCdnCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetCertId(v int64) *DescribeCdnCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetCertName(v string) *DescribeCdnCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetKey(v string) *DescribeCdnCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) SetRequestId(v string) *DescribeCdnCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
