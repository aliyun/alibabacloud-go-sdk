// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnHttpsDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeDcdnHttpsDomainListRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeDcdnHttpsDomainListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnHttpsDomainListRequest
	GetPageSize() *int32
}

type DescribeDcdnHttpsDomainListRequest struct {
	// The keyword that is used to search for certificates.
	//
	// example:
	//
	// cert
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of returned pages. Valid values: **1 to 100000**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1 to 500**. Default value: **20**.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDcdnHttpsDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeDcdnHttpsDomainListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnHttpsDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnHttpsDomainListRequest) SetKeyword(v string) *DescribeDcdnHttpsDomainListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListRequest) SetPageNumber(v int32) *DescribeDcdnHttpsDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListRequest) SetPageSize(v int32) *DescribeDcdnHttpsDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListRequest) Validate() error {
	return dara.Validate(s)
}
