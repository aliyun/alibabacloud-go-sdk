// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeVodCertificateDetailResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeVodCertificateDetailResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeVodCertificateDetailResponseBody
	GetCertName() *string
	SetKey(v string) *DescribeVodCertificateDetailResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeVodCertificateDetailResponseBody
	GetRequestId() *string
}

type DescribeVodCertificateDetailResponseBody struct {
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateDetailResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeVodCertificateDetailResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeVodCertificateDetailResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodCertificateDetailResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeVodCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodCertificateDetailResponseBody) SetCert(v string) *DescribeVodCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeVodCertificateDetailResponseBody) SetCertId(v int64) *DescribeVodCertificateDetailResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeVodCertificateDetailResponseBody) SetCertName(v string) *DescribeVodCertificateDetailResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeVodCertificateDetailResponseBody) SetKey(v string) *DescribeVodCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeVodCertificateDetailResponseBody) SetRequestId(v string) *DescribeVodCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodCertificateDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
