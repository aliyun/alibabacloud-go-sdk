// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceClusterListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceClusterId(v string) *DescribeInstanceClusterListRequest
	GetInstanceClusterId() *string
	SetInstanceClusterName(v string) *DescribeInstanceClusterListRequest
	GetInstanceClusterName() *string
	SetPageNumber(v int32) *DescribeInstanceClusterListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceClusterListRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeInstanceClusterListRequest
	GetSecurityToken() *string
}

type DescribeInstanceClusterListRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// apigateway-cluster-hz-xxxxxxxxxxxx
	InstanceClusterId *string `json:"InstanceClusterId,omitempty" xml:"InstanceClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test-cluster
	InstanceClusterName *string `json:"InstanceClusterName,omitempty" xml:"InstanceClusterName,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeInstanceClusterListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterListRequest) GetInstanceClusterId() *string {
	return s.InstanceClusterId
}

func (s *DescribeInstanceClusterListRequest) GetInstanceClusterName() *string {
	return s.InstanceClusterName
}

func (s *DescribeInstanceClusterListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceClusterListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceClusterListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceClusterListRequest) SetInstanceClusterId(v string) *DescribeInstanceClusterListRequest {
	s.InstanceClusterId = &v
	return s
}

func (s *DescribeInstanceClusterListRequest) SetInstanceClusterName(v string) *DescribeInstanceClusterListRequest {
	s.InstanceClusterName = &v
	return s
}

func (s *DescribeInstanceClusterListRequest) SetPageNumber(v int32) *DescribeInstanceClusterListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceClusterListRequest) SetPageSize(v int32) *DescribeInstanceClusterListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceClusterListRequest) SetSecurityToken(v string) *DescribeInstanceClusterListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceClusterListRequest) Validate() error {
	return dara.Validate(s)
}
