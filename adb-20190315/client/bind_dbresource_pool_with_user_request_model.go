// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourcePoolWithUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *BindDBResourcePoolWithUserRequest
	GetClientToken() *string
	SetDBClusterId(v string) *BindDBResourcePoolWithUserRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *BindDBResourcePoolWithUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindDBResourcePoolWithUserRequest
	GetOwnerId() *int64
	SetPoolName(v string) *BindDBResourcePoolWithUserRequest
	GetPoolName() *string
	SetPoolUser(v string) *BindDBResourcePoolWithUserRequest
	GetPoolUser() *string
	SetResourceOwnerAccount(v string) *BindDBResourcePoolWithUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindDBResourcePoolWithUserRequest
	GetResourceOwnerId() *int64
}

type BindDBResourcePoolWithUserRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1ub9grke1****
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
	// The database account with which to associate the resource group. It can be a standard account or a privileged account.
	//
	// This parameter is required.
	//
	// example:
	//
	// accout
	PoolUser             *string `json:"PoolUser,omitempty" xml:"PoolUser,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDBResourcePoolWithUserRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourcePoolWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourcePoolWithUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BindDBResourcePoolWithUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BindDBResourcePoolWithUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindDBResourcePoolWithUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDBResourcePoolWithUserRequest) GetPoolName() *string {
	return s.PoolName
}

func (s *BindDBResourcePoolWithUserRequest) GetPoolUser() *string {
	return s.PoolUser
}

func (s *BindDBResourcePoolWithUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindDBResourcePoolWithUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindDBResourcePoolWithUserRequest) SetClientToken(v string) *BindDBResourcePoolWithUserRequest {
	s.ClientToken = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetDBClusterId(v string) *BindDBResourcePoolWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetPoolName(v string) *BindDBResourcePoolWithUserRequest {
	s.PoolName = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetPoolUser(v string) *BindDBResourcePoolWithUserRequest {
	s.PoolUser = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetResourceOwnerAccount(v string) *BindDBResourcePoolWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) SetResourceOwnerId(v int64) *BindDBResourcePoolWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindDBResourcePoolWithUserRequest) Validate() error {
	return dara.Validate(s)
}
