// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDataNetworkListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeGlobalDataNetworkListRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeGlobalDataNetworkListRequest
	GetPageSize() *int64
}

type DescribeGlobalDataNetworkListRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGlobalDataNetworkListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDataNetworkListRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDataNetworkListRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeGlobalDataNetworkListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeGlobalDataNetworkListRequest) SetPageNumber(v int64) *DescribeGlobalDataNetworkListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDataNetworkListRequest) SetPageSize(v int64) *DescribeGlobalDataNetworkListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDataNetworkListRequest) Validate() error {
	return dara.Validate(s)
}
