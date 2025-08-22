// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafDomainsRequest
	GetPageSize() *int32
	SetQueryArgs(v string) *DescribeDcdnWafDomainsRequest
	GetQueryArgs() *string
}

type DescribeDcdnWafDomainsRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names to return on each page. Valid values: an integer from **1*	- to **500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. You can filter domain names by name. Fuzzy match is supported `QueryArgs={"DomainName":"Accelerated domain name"}`
	//
	// example:
	//
	// {"DomainName":"example.com"}
	QueryArgs *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s DescribeDcdnWafDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafDomainsRequest) GetQueryArgs() *string {
	return s.QueryArgs
}

func (s *DescribeDcdnWafDomainsRequest) SetPageNumber(v int32) *DescribeDcdnWafDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafDomainsRequest) SetPageSize(v int32) *DescribeDcdnWafDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafDomainsRequest) SetQueryArgs(v string) *DescribeDcdnWafDomainsRequest {
	s.QueryArgs = &v
	return s
}

func (s *DescribeDcdnWafDomainsRequest) Validate() error {
	return dara.Validate(s)
}
