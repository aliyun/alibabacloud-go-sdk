// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterProxyRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterProxyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterProxyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterProxyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBClusterProxyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterProxyRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterProxyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2zek76tdi709m6mf4
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm4ifnqnun3zq
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterProxyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterProxyRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterProxyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterProxyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterProxyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClusterProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterProxyRequest) SetDBClusterId(v string) *DescribeDBClusterProxyRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) SetOwnerAccount(v string) *DescribeDBClusterProxyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) SetOwnerId(v int64) *DescribeDBClusterProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) SetRegionId(v string) *DescribeDBClusterProxyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) SetResourceGroupId(v string) *DescribeDBClusterProxyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) SetResourceOwnerId(v int64) *DescribeDBClusterProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterProxyRequest) Validate() error {
	return dara.Validate(s)
}
