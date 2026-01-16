// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterNodesRequest
	GetClusterId() *string
	SetNodepoolId(v string) *DescribeClusterNodesRequest
	GetNodepoolId() *string
	SetPageNumber(v int32) *DescribeClusterNodesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClusterNodesRequest
	GetPageSize() *int32
}

type DescribeClusterNodesRequest struct {
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
	// np751e9d5464cf4ef081087137f33e4528
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
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

func (s DescribeClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterNodesRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *DescribeClusterNodesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClusterNodesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClusterNodesRequest) SetClusterId(v string) *DescribeClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetNodepoolId(v string) *DescribeClusterNodesRequest {
	s.NodepoolId = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetPageNumber(v int32) *DescribeClusterNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClusterNodesRequest) SetPageSize(v int32) *DescribeClusterNodesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}
