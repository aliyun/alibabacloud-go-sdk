// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourcePoolWithUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnbindDBResourcePoolWithUserRequest
	GetClientToken() *string
	SetDBClusterId(v string) *UnbindDBResourcePoolWithUserRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest
	GetOwnerId() *int64
	SetPoolName(v string) *UnbindDBResourcePoolWithUserRequest
	GetPoolName() *string
	SetPoolUser(v string) *UnbindDBResourcePoolWithUserRequest
	GetPoolUser() *string
	SetResourceOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest
	GetResourceOwnerId() *int64
}

type UnbindDBResourcePoolWithUserRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bpxxxxxxxx47
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The database account with which the resource group is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// testb
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnbindDBResourcePoolWithUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourcePoolWithUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnbindDBResourcePoolWithUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UnbindDBResourcePoolWithUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindDBResourcePoolWithUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindDBResourcePoolWithUserRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *UnbindDBResourcePoolWithUserRequest) GetPoolUser() *string {
	return s.PoolUser
}

func (s *UnbindDBResourcePoolWithUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindDBResourcePoolWithUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindDBResourcePoolWithUserRequest) SetClientToken(v string) *UnbindDBResourcePoolWithUserRequest {
	s.ClientToken = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *UnbindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetPoolName(v string) *UnbindDBResourcePoolWithUserRequest {
	s.PoolName = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetPoolUser(v string) *UnbindDBResourcePoolWithUserRequest {
	s.PoolUser = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetResourceOwnerAccount(v string) *UnbindDBResourcePoolWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) SetResourceOwnerId(v int64) *UnbindDBResourcePoolWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindDBResourcePoolWithUserRequest) Validate() error {
	return dara.Validate(s)
}
