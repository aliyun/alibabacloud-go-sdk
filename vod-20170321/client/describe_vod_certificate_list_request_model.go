// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodCertificateListRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeVodCertificateListRequest
	GetSecurityToken() *string
}

type DescribeVodCertificateListRequest struct {
	// The domain name for CDN.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodCertificateListRequest) SetDomainName(v string) *DescribeVodCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodCertificateListRequest) SetOwnerId(v int64) *DescribeVodCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodCertificateListRequest) SetSecurityToken(v string) *DescribeVodCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
