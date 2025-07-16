// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnHttpsDomainListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeCdnHttpsDomainListRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DescribeCdnHttpsDomainListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCdnHttpsDomainListRequest
	GetPageSize() *int32
}

type DescribeCdnHttpsDomainListRequest struct {
	// The keyword that is used to search for certificates.
	//
	// example:
	//
	// com
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of the page to return. Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCdnHttpsDomainListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnHttpsDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeCdnHttpsDomainListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCdnHttpsDomainListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCdnHttpsDomainListRequest) SetKeyword(v string) *DescribeCdnHttpsDomainListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeCdnHttpsDomainListRequest) SetPageNumber(v int32) *DescribeCdnHttpsDomainListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnHttpsDomainListRequest) SetPageSize(v int32) *DescribeCdnHttpsDomainListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnHttpsDomainListRequest) Validate() error {
	return dara.Validate(s)
}
