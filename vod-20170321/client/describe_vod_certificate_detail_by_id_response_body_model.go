// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateDetailByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *DescribeVodCertificateDetailByIdResponseBody
	GetCert() *string
	SetCertId(v int64) *DescribeVodCertificateDetailByIdResponseBody
	GetCertId() *int64
	SetCertName(v string) *DescribeVodCertificateDetailByIdResponseBody
	GetCertName() *string
	SetKey(v string) *DescribeVodCertificateDetailByIdResponseBody
	GetKey() *string
	SetRequestId(v string) *DescribeVodCertificateDetailByIdResponseBody
	GetRequestId() *string
}

type DescribeVodCertificateDetailByIdResponseBody struct {
	Cert      *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName  *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodCertificateDetailByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateDetailByIdResponseBody) GetCert() *string {
	return s.Cert
}

func (s *DescribeVodCertificateDetailByIdResponseBody) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeVodCertificateDetailByIdResponseBody) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodCertificateDetailByIdResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeVodCertificateDetailByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodCertificateDetailByIdResponseBody) SetCert(v string) *DescribeVodCertificateDetailByIdResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponseBody) SetCertId(v int64) *DescribeVodCertificateDetailByIdResponseBody {
	s.CertId = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponseBody) SetCertName(v string) *DescribeVodCertificateDetailByIdResponseBody {
	s.CertName = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponseBody) SetKey(v string) *DescribeVodCertificateDetailByIdResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponseBody) SetRequestId(v string) *DescribeVodCertificateDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
