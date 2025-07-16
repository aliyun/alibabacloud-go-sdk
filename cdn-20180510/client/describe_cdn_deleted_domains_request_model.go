// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeletedDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeCdnDeletedDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCdnDeletedDomainsRequest
	GetPageSize() *int32
}

type DescribeCdnDeletedDomainsRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names to return per page. Valid values: an integer between **1*	- and **500**. Default value: **20**.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCdnDeletedDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeletedDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeletedDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCdnDeletedDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCdnDeletedDomainsRequest) SetPageNumber(v int32) *DescribeCdnDeletedDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnDeletedDomainsRequest) SetPageSize(v int32) *DescribeCdnDeletedDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnDeletedDomainsRequest) Validate() error {
	return dara.Validate(s)
}
