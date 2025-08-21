// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVsCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVsCertificateListRequest
	GetOwnerId() *int64
}

type DescribeVsCertificateListRequest struct {
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVsCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsCertificateListRequest) SetDomainName(v string) *DescribeVsCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsCertificateListRequest) SetOwnerId(v int64) *DescribeVsCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
