// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeNamespacesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeNamespacesRequest
	GetPageSize() *int32
}

type DescribeNamespacesRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Valid values: 0 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNamespacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNamespacesRequest) SetCurrentPage(v int32) *DescribeNamespacesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNamespacesRequest) SetPageSize(v int32) *DescribeNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
