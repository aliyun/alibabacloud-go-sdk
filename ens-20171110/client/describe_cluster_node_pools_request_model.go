// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterNodePoolsRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeClusterNodePoolsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterNodePoolsRequest
	GetPageSize() *int32
}

type DescribeClusterNodePoolsRequest struct {
	// A short description of struct
	//
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeClusterNodePoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterNodePoolsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterNodePoolsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterNodePoolsRequest) SetClusterId(v string) *DescribeClusterNodePoolsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodePoolsRequest) SetPageNumber(v int32) *DescribeClusterNodePoolsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodePoolsRequest) SetPageSize(v int32) *DescribeClusterNodePoolsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodePoolsRequest) Validate() error {
	return dara.Validate(s)
}
