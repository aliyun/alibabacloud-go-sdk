// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainCertificateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsDomainCertificateInfoRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVsDomainCertificateInfoRequest
	GetOwnerId() *int64
}

type DescribeVsDomainCertificateInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsDomainCertificateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsDomainCertificateInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsDomainCertificateInfoRequest) SetDomainName(v string) *DescribeVsDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeVsDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoRequest) Validate() error {
	return dara.Validate(s)
}
