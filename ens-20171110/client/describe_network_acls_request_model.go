// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclId(v string) *DescribeNetworkAclsRequest
	GetNetworkAclId() *string
	SetNetworkAclName(v string) *DescribeNetworkAclsRequest
	GetNetworkAclName() *string
	SetPageNumber(v string) *DescribeNetworkAclsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeNetworkAclsRequest
	GetPageSize() *string
	SetResourceId(v string) *DescribeNetworkAclsRequest
	GetResourceId() *string
}

type DescribeNetworkAclsRequest struct {
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-bp1lhl0taikrbgnh****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The name of the network ACL.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the associated instance.
	//
	// example:
	//
	// n-5****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DescribeNetworkAclsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAclsRequest) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *DescribeNetworkAclsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeNetworkAclsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNetworkAclsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetworkAclsRequest) SetNetworkAclId(v string) *DescribeNetworkAclsRequest {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetNetworkAclName(v string) *DescribeNetworkAclsRequest {
	s.NetworkAclName = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetPageNumber(v string) *DescribeNetworkAclsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetPageSize(v string) *DescribeNetworkAclsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetResourceId(v string) *DescribeNetworkAclsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) Validate() error {
	return dara.Validate(s)
}
