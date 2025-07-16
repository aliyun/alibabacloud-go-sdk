// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSSLCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnSSLCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeCdnSSLCertificateListRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeCdnSSLCertificateListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCdnSSLCertificateListRequest
	GetPageSize() *int64
	SetSearchKeyword(v string) *DescribeCdnSSLCertificateListRequest
	GetSearchKeyword() *string
	SetSecurityToken(v string) *DescribeCdnSSLCertificateListRequest
	GetSecurityToken() *string
}

type DescribeCdnSSLCertificateListRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: an integer from **1*	- to **1000**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used to query the certificate.
	//
	// example:
	//
	// certabc
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnSSLCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSSLCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnSSLCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnSSLCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnSSLCertificateListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnSSLCertificateListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnSSLCertificateListRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *DescribeCdnSSLCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnSSLCertificateListRequest) SetDomainName(v string) *DescribeCdnSSLCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnSSLCertificateListRequest) SetOwnerId(v int64) *DescribeCdnSSLCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnSSLCertificateListRequest) SetPageNumber(v int64) *DescribeCdnSSLCertificateListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnSSLCertificateListRequest) SetPageSize(v int64) *DescribeCdnSSLCertificateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnSSLCertificateListRequest) SetSearchKeyword(v string) *DescribeCdnSSLCertificateListRequest {
	s.SearchKeyword = &v
	return s
}

func (s *DescribeCdnSSLCertificateListRequest) SetSecurityToken(v string) *DescribeCdnSSLCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnSSLCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
