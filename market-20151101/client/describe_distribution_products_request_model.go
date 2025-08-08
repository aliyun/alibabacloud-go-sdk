// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeDistributionProductsRequestFilter) *DescribeDistributionProductsRequest
	GetFilter() []*DescribeDistributionProductsRequestFilter
	SetPageNumber(v int64) *DescribeDistributionProductsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDistributionProductsRequest
	GetPageSize() *int64
}

type DescribeDistributionProductsRequest struct {
	Filter []*DescribeDistributionProductsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDistributionProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsRequest) GetFilter() []*DescribeDistributionProductsRequestFilter {
	return s.Filter
}

func (s *DescribeDistributionProductsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDistributionProductsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDistributionProductsRequest) SetFilter(v []*DescribeDistributionProductsRequestFilter) *DescribeDistributionProductsRequest {
	s.Filter = v
	return s
}

func (s *DescribeDistributionProductsRequest) SetPageNumber(v int64) *DescribeDistributionProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDistributionProductsRequest) SetPageSize(v int64) *DescribeDistributionProductsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDistributionProductsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDistributionProductsRequestFilter struct {
	// This parameter is required.
	//
	// example:
	//
	// supplierName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cmj0000000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDistributionProductsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeDistributionProductsRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeDistributionProductsRequestFilter) SetKey(v string) *DescribeDistributionProductsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeDistributionProductsRequestFilter) SetValue(v string) *DescribeDistributionProductsRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeDistributionProductsRequestFilter) Validate() error {
	return dara.Validate(s)
}
