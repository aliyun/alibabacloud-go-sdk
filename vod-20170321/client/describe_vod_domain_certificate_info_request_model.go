// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCertificateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainCertificateInfoRequest
	GetDomainName() *string
	SetHeraApiAutoVersion(v string) *DescribeVodDomainCertificateInfoRequest
	GetHeraApiAutoVersion() *string
	SetOwnerId(v int64) *DescribeVodDomainCertificateInfoRequest
	GetOwnerId() *int64
}

type DescribeVodDomainCertificateInfoRequest struct {
	// The accelerated domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName         *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	HeraApiAutoVersion *string `json:"HeraApiAutoVersion,omitempty" xml:"HeraApiAutoVersion,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodDomainCertificateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainCertificateInfoRequest) GetHeraApiAutoVersion() *string {
	return s.HeraApiAutoVersion
}

func (s *DescribeVodDomainCertificateInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainCertificateInfoRequest) SetDomainName(v string) *DescribeVodDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoRequest) SetHeraApiAutoVersion(v string) *DescribeVodDomainCertificateInfoRequest {
	s.HeraApiAutoVersion = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeVodDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoRequest) Validate() error {
	return dara.Validate(s)
}
