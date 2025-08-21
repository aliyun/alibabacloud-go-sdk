// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *DescribeVsCertificateDetailRequest
	GetCertName() *string
	SetOwnerId(v int64) *DescribeVsCertificateDetailRequest
	GetOwnerId() *int64
}

type DescribeVsCertificateDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cert-539xxxx
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateDetailRequest) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVsCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsCertificateDetailRequest) SetCertName(v string) *DescribeVsCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeVsCertificateDetailRequest) SetOwnerId(v int64) *DescribeVsCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
