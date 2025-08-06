// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodSSLCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodSSLCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodSSLCertificateListRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeVodSSLCertificateListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodSSLCertificateListRequest
	GetPageSize() *int64
	SetSearchKeyword(v string) *DescribeVodSSLCertificateListRequest
	GetSearchKeyword() *string
	SetSecurityToken(v string) *DescribeVodSSLCertificateListRequest
	GetSecurityToken() *string
}

type DescribeVodSSLCertificateListRequest struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: integers from 1 to 1000.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used to query certificates.
	//
	// example:
	//
	// certabc
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodSSLCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodSSLCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodSSLCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodSSLCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodSSLCertificateListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodSSLCertificateListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodSSLCertificateListRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *DescribeVodSSLCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeVodSSLCertificateListRequest) SetDomainName(v string) *DescribeVodSSLCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodSSLCertificateListRequest) SetOwnerId(v int64) *DescribeVodSSLCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodSSLCertificateListRequest) SetPageNumber(v int64) *DescribeVodSSLCertificateListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodSSLCertificateListRequest) SetPageSize(v int64) *DescribeVodSSLCertificateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodSSLCertificateListRequest) SetSearchKeyword(v string) *DescribeVodSSLCertificateListRequest {
	s.SearchKeyword = &v
	return s
}

func (s *DescribeVodSSLCertificateListRequest) SetSecurityToken(v string) *DescribeVodSSLCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodSSLCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
