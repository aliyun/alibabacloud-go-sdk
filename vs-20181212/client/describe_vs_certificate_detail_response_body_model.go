// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeVsCertificateDetailResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeVsCertificateDetailResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeVsCertificateDetailResponseBody
	GetCertName() *string
	SetKey(v string) *DescribeVsCertificateDetailResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeVsCertificateDetailResponseBody
	GetRequestId() *string
}

type DescribeVsCertificateDetailResponseBody struct {
	// example:
	//
	// -----BEGIN CERTIFICATE-----xxxxx-----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// example:
	//
	// 63000000
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// cert-539xxxxx
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// xxxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// D94D0E1E-E71B-562D-8C18-969BB3653FBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateDetailResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeVsCertificateDetailResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeVsCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVsCertificateDetailResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeVsCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsCertificateDetailResponseBody) SetCert(v string) *DescribeVsCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetCertId(v int64) *DescribeVsCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetCertName(v string) *DescribeVsCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetKey(v string) *DescribeVsCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) SetRequestId(v string) *DescribeVsCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
