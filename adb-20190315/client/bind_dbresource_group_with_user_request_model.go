// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *BindDBResourceGroupWithUserRequest
	GetClientToken() *string
	SetDBClusterId(v string) *BindDBResourceGroupWithUserRequest
	GetDBClusterId() *string
	SetGroupName(v string) *BindDBResourceGroupWithUserRequest
	GetGroupName() *string
	SetGroupUser(v string) *BindDBResourceGroupWithUserRequest
	GetGroupUser() *string
	SetOwnerAccount(v string) *BindDBResourceGroupWithUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BindDBResourceGroupWithUserRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *BindDBResourceGroupWithUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindDBResourceGroupWithUserRequest
	GetResourceOwnerId() *int64
}

type BindDBResourceGroupWithUserRequest struct {
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
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The database account with which to associate the resource group. It can be a standard account or a privileged account.
	//
	// This parameter is required.
	//
	// example:
	//
	// accout
	GroupUser            *string `json:"GroupUser,omitempty" xml:"GroupUser,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDBResourceGroupWithUserRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *BindDBResourceGroupWithUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BindDBResourceGroupWithUserRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *BindDBResourceGroupWithUserRequest) GetGroupUser() *string {
	return s.GroupUser
}

func (s *BindDBResourceGroupWithUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BindDBResourceGroupWithUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDBResourceGroupWithUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindDBResourceGroupWithUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindDBResourceGroupWithUserRequest) SetClientToken(v string) *BindDBResourceGroupWithUserRequest {
	s.ClientToken = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *BindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupName(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetGroupUser(v string) *BindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetOwnerAccount(v string) *BindDBResourceGroupWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetOwnerId(v int64) *BindDBResourceGroupWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetResourceOwnerAccount(v string) *BindDBResourceGroupWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) SetResourceOwnerId(v int64) *BindDBResourceGroupWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindDBResourceGroupWithUserRequest) Validate() error {
	return dara.Validate(s)
}
