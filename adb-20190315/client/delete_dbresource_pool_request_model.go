// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourcePoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBResourcePoolRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DeleteDBResourcePoolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBResourcePoolRequest
	GetOwnerId() *int64
	SetPoolName(v string) *DeleteDBResourcePoolRequest
	GetPoolName() *string
	SetResourceOwnerAccount(v string) *DeleteDBResourcePoolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBResourcePoolRequest
	GetResourceOwnerId() *int64
}

type DeleteDBResourcePoolRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1r053byu48p****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	PoolName             *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBResourcePoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourcePoolRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourcePoolRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBResourcePoolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBResourcePoolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBResourcePoolRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *DeleteDBResourcePoolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBResourcePoolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBResourcePoolRequest) SetDBClusterId(v string) *DeleteDBResourcePoolRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetPoolName(v string) *DeleteDBResourcePoolRequest {
	s.PoolName = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetResourceOwnerAccount(v string) *DeleteDBResourcePoolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) SetResourceOwnerId(v int64) *DeleteDBResourcePoolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBResourcePoolRequest) Validate() error {
	return dara.Validate(s)
}
