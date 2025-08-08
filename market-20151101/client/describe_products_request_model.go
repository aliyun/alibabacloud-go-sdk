// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeProductsRequestFilter) *DescribeProductsRequest
	GetFilter() []*DescribeProductsRequestFilter
	SetPageNumber(v int32) *DescribeProductsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProductsRequest
	GetPageSize() *int32
	SetSearchTerm(v string) *DescribeProductsRequest
	GetSearchTerm() *string
}

type DescribeProductsRequest struct {
	Filter []*DescribeProductsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchTerm *string `json:"SearchTerm,omitempty" xml:"SearchTerm,omitempty"`
}

func (s DescribeProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequest) GetFilter() []*DescribeProductsRequestFilter {
	return s.Filter
}

func (s *DescribeProductsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProductsRequest) GetSearchTerm() *string {
	return s.SearchTerm
}

func (s *DescribeProductsRequest) SetFilter(v []*DescribeProductsRequestFilter) *DescribeProductsRequest {
	s.Filter = v
	return s
}

func (s *DescribeProductsRequest) SetPageNumber(v int32) *DescribeProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProductsRequest) SetPageSize(v int32) *DescribeProductsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProductsRequest) SetSearchTerm(v string) *DescribeProductsRequest {
	s.SearchTerm = &v
	return s
}

func (s *DescribeProductsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsRequestFilter struct {
	// example:
	//
	// categoryId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 53366009
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProductsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeProductsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeProductsRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeProductsRequestFilter) SetKey(v string) *DescribeProductsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeProductsRequestFilter) SetValue(v string) *DescribeProductsRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeProductsRequestFilter) Validate() error {
	return dara.Validate(s)
}
