// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeClusterTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterTasksRequest
	GetPageSize() *int32
}

type DescribeClusterTasksRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s DescribeClusterTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterTasksRequest) SetPageNumber(v int32) *DescribeClusterTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterTasksRequest) SetPageSize(v int32) *DescribeClusterTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterTasksRequest) Validate() error {
	return dara.Validate(s)
}
