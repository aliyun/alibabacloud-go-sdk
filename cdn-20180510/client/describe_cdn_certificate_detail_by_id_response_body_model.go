// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateDetailByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeCdnCertificateDetailByIdResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeCdnCertificateDetailByIdResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeCdnCertificateDetailByIdResponseBody
	GetCertName() *string
	SetKey(v string) *DescribeCdnCertificateDetailByIdResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeCdnCertificateDetailByIdResponseBody
	GetRequestId() *string
}

type DescribeCdnCertificateDetailByIdResponseBody struct {
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGINCERTIFICATE-----xxx-----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 12345
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// 587f6db37e3a2f01047b032b739cbe31
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnCertificateDetailByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) SetCert(v string) *DescribeCdnCertificateDetailByIdResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) SetCertId(v int64) *DescribeCdnCertificateDetailByIdResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) SetCertName(v string) *DescribeCdnCertificateDetailByIdResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) SetKey(v string) *DescribeCdnCertificateDetailByIdResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) SetRequestId(v string) *DescribeCdnCertificateDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
