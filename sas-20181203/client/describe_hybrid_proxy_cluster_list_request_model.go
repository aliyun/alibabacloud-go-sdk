// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyClusterListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeHybridProxyClusterListRequest
	GetClusterName() *string
	SetCurrentPage(v int32) *DescribeHybridProxyClusterListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeHybridProxyClusterListRequest
	GetPageSize() *int32
}

type DescribeHybridProxyClusterListRequest struct {
	// The name of the proxy cluster.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeHybridProxyClusterListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyClusterListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyClusterListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridProxyClusterListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeHybridProxyClusterListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridProxyClusterListRequest) SetClusterName(v string) *DescribeHybridProxyClusterListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridProxyClusterListRequest) SetCurrentPage(v int32) *DescribeHybridProxyClusterListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHybridProxyClusterListRequest) SetPageSize(v int32) *DescribeHybridProxyClusterListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridProxyClusterListRequest) Validate() error {
	return dara.Validate(s)
}
