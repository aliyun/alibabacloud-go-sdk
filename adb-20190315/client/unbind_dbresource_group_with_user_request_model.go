// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnbindDBResourceGroupWithUserRequest
	GetClientToken() *string
	SetDBClusterId(v string) *UnbindDBResourceGroupWithUserRequest
	GetDBClusterId() *string
	SetGroupName(v string) *UnbindDBResourceGroupWithUserRequest
	GetGroupName() *string
	SetGroupUser(v string) *UnbindDBResourceGroupWithUserRequest
	GetGroupUser() *string
	SetOwnerAccount(v string) *UnbindDBResourceGroupWithUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnbindDBResourceGroupWithUserRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnbindDBResourceGroupWithUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnbindDBResourceGroupWithUserRequest
	GetResourceOwnerId() *int64
}

type UnbindDBResourceGroupWithUserRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
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
	// The database account with which the resource group is associated.
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

func (s UnbindDBResourceGroupWithUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnbindDBResourceGroupWithUserRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UnbindDBResourceGroupWithUserRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UnbindDBResourceGroupWithUserRequest) GetGroupUser() *string {
	return s.GroupUser
}

func (s *UnbindDBResourceGroupWithUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnbindDBResourceGroupWithUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindDBResourceGroupWithUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnbindDBResourceGroupWithUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnbindDBResourceGroupWithUserRequest) SetClientToken(v string) *UnbindDBResourceGroupWithUserRequest {
	s.ClientToken = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetDBClusterId(v string) *UnbindDBResourceGroupWithUserRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupName(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupName = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetGroupUser(v string) *UnbindDBResourceGroupWithUserRequest {
	s.GroupUser = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetOwnerAccount(v string) *UnbindDBResourceGroupWithUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetOwnerId(v int64) *UnbindDBResourceGroupWithUserRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetResourceOwnerAccount(v string) *UnbindDBResourceGroupWithUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) SetResourceOwnerId(v int64) *UnbindDBResourceGroupWithUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindDBResourceGroupWithUserRequest) Validate() error {
	return dara.Validate(s)
}
