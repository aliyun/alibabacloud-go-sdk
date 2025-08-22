// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDeletedDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnDeletedDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnDeletedDomainsRequest
	GetPageSize() *int32
}

type DescribeDcdnDeletedDomainsRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names to return on each page. Valid values: an integer from **1*	- to **500**. Default value: **20**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDcdnDeletedDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnDeletedDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnDeletedDomainsRequest) SetPageNumber(v int32) *DescribeDcdnDeletedDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsRequest) SetPageSize(v int32) *DescribeDcdnDeletedDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsRequest) Validate() error {
	return dara.Validate(s)
}
