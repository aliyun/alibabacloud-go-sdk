// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourcePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBResourcePoolRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDBResourcePoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBResourcePoolRequest
	GetOwnerId() *int64
	SetPoolName(v string) *DescribeDBResourcePoolRequest
	GetPoolName() *string
	SetResourceOwnerAccount(v string) *DescribeDBResourcePoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBResourcePoolRequest
	GetResourceOwnerId() *int64
}

type DescribeDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// USER_DEFAULT
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBResourcePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourcePoolRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBResourcePoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBResourcePoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBResourcePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *DescribeDBResourcePoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBResourcePoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBResourcePoolRequest) SetDBClusterId(v string) *DescribeDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetPoolName(v string) *DescribeDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetResourceOwnerAccount(v string) *DescribeDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) SetResourceOwnerId(v int64) *DescribeDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBResourcePoolRequest) Validate() error {
	return dara.Validate(s)
}
