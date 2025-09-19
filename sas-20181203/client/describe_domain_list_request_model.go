// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDomainListRequest
	GetCurrentPage() *int32
	SetDomainType(v string) *DescribeDomainListRequest
	GetDomainType() *string
	SetFuzzyDomain(v string) *DescribeDomainListRequest
	GetFuzzyDomain() *string
	SetPageSize(v int32) *DescribeDomainListRequest
	GetPageSize() *int32
	SetSourceIp(v string) *DescribeDomainListRequest
	GetSourceIp() *string
}

type DescribeDomainListRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the domain name that you want to query. Valid values:
	//
	// 	- **root**: root domain name
	//
	// 	- **sub**: subdomain name
	//
	// example:
	//
	// root
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The keyword that is used to query domain names. Fuzzy match is supported.
	//
	// example:
	//
	// sas
	FuzzyDomain *string `json:"FuzzyDomain,omitempty" xml:"FuzzyDomain,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.122.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDomainListRequest) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDomainListRequest) GetFuzzyDomain() *string {
	return s.FuzzyDomain
}

func (s *DescribeDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainListRequest) SetCurrentPage(v int32) *DescribeDomainListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDomainListRequest) SetDomainType(v string) *DescribeDomainListRequest {
	s.DomainType = &v
	return s
}

func (s *DescribeDomainListRequest) SetFuzzyDomain(v string) *DescribeDomainListRequest {
	s.FuzzyDomain = &v
	return s
}

func (s *DescribeDomainListRequest) SetPageSize(v int32) *DescribeDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainListRequest) SetSourceIp(v string) *DescribeDomainListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainListRequest) Validate() error {
	return dara.Validate(s)
}
