// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterShardNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterShardNumberRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBClusterShardNumberRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterShardNumberRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBClusterShardNumberRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterShardNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterShardNumberRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClusterShardNumberRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClusterShardNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterShardNumberRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterShardNumberRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterShardNumberRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterShardNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterShardNumberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterShardNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterShardNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterShardNumberRequest) SetDBClusterId(v string) *DescribeDBClusterShardNumberRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterShardNumberRequest) SetOwnerAccount(v string) *DescribeDBClusterShardNumberRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterShardNumberRequest) SetOwnerId(v int64) *DescribeDBClusterShardNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterShardNumberRequest) SetRegionId(v string) *DescribeDBClusterShardNumberRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterShardNumberRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterShardNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterShardNumberRequest) SetResourceOwnerId(v int64) *DescribeDBClusterShardNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterShardNumberRequest) Validate() error {
	return dara.Validate(s)
}
