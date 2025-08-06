// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *DescribeVodCertificateDetailRequest
	GetCertName() *string
	SetOwnerId(v int64) *DescribeVodCertificateDetailRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodCertificateDetailRequest
	GetSecurityToken() *string
}

type DescribeVodCertificateDetailRequest struct {
	// This parameter is required.
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodCertificateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateDetailRequest) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodCertificateDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodCertificateDetailRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodCertificateDetailRequest) SetCertName(v string) *DescribeVodCertificateDetailRequest {
	s.CertName = &v
	return s
}

func (s *DescribeVodCertificateDetailRequest) SetOwnerId(v int64) *DescribeVodCertificateDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodCertificateDetailRequest) SetSecurityToken(v string) *DescribeVodCertificateDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodCertificateDetailRequest) Validate() error {
	return dara.Validate(s)
}
