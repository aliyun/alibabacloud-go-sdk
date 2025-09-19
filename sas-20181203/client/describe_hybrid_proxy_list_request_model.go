// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeHybridProxyListRequest
	GetClusterName() *string
	SetCurrentPage(v int32) *DescribeHybridProxyListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeHybridProxyListRequest
	GetPageSize() *int32
}

type DescribeHybridProxyListRequest struct {
	// The name of the proxy cluster.
	//
	// example:
	//
	// idc-sas-proxy
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeHybridProxyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridProxyListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeHybridProxyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHybridProxyListRequest) SetClusterName(v string) *DescribeHybridProxyListRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridProxyListRequest) SetCurrentPage(v int32) *DescribeHybridProxyListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeHybridProxyListRequest) SetPageSize(v int32) *DescribeHybridProxyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridProxyListRequest) Validate() error {
	return dara.Validate(s)
}
