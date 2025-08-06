// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateDetailByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *DescribeVodCertificateDetailByIdRequest
	GetCertId() *string
	SetCertRegion(v string) *DescribeVodCertificateDetailByIdRequest
	GetCertRegion() *string
	SetOwnerId(v int64) *DescribeVodCertificateDetailByIdRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodCertificateDetailByIdRequest
	GetSecurityToken() *string
}

type DescribeVodCertificateDetailByIdRequest struct {
	// This parameter is required.
	CertId        *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertRegion    *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodCertificateDetailByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateDetailByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateDetailByIdRequest) GetCertId() *string {
	return s.CertId
}

func (s *DescribeVodCertificateDetailByIdRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeVodCertificateDetailByIdRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodCertificateDetailByIdRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodCertificateDetailByIdRequest) SetCertId(v string) *DescribeVodCertificateDetailByIdRequest {
	s.CertId = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdRequest) SetCertRegion(v string) *DescribeVodCertificateDetailByIdRequest {
	s.CertRegion = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdRequest) SetOwnerId(v int64) *DescribeVodCertificateDetailByIdRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdRequest) SetSecurityToken(v string) *DescribeVodCertificateDetailByIdRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodCertificateDetailByIdRequest) Validate() error {
	return dara.Validate(s)
}
