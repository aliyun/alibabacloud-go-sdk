// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSSLCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnSSLCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeDcdnSSLCertificateListRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeDcdnSSLCertificateListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDcdnSSLCertificateListRequest
	GetPageSize() *int64
	SetSearchKeyword(v string) *DescribeDcdnSSLCertificateListRequest
	GetSearchKeyword() *string
	SetSecurityToken(v string) *DescribeDcdnSSLCertificateListRequest
	GetSecurityToken() *string
}

type DescribeDcdnSSLCertificateListRequest struct {
	// The accelerated domain secured by the SSL certificate. HTTPS secure acceleration is enabled for the accelerated domain name.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page. Valid values: **1 to 1000**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword that is used to query the certificate.
	//
	// example:
	//
	// taobao
	SearchKeyword *string `json:"SearchKeyword,omitempty" xml:"SearchKeyword,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnSSLCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSSLCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSSLCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnSSLCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnSSLCertificateListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnSSLCertificateListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnSSLCertificateListRequest) GetSearchKeyword() *string {
	return s.SearchKeyword
}

func (s *DescribeDcdnSSLCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnSSLCertificateListRequest) SetDomainName(v string) *DescribeDcdnSSLCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListRequest) SetOwnerId(v int64) *DescribeDcdnSSLCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListRequest) SetPageNumber(v int64) *DescribeDcdnSSLCertificateListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListRequest) SetPageSize(v int64) *DescribeDcdnSSLCertificateListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListRequest) SetSearchKeyword(v string) *DescribeDcdnSSLCertificateListRequest {
	s.SearchKeyword = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListRequest) SetSecurityToken(v string) *DescribeDcdnSSLCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
