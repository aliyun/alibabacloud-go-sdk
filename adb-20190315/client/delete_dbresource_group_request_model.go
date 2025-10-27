// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBResourceGroupRequest
	GetDBClusterId() *string
	SetGroupName(v string) *DeleteDBResourceGroupRequest
	GetGroupName() *string
	SetOwnerAccount(v string) *DeleteDBResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDBResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBResourceGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteDBResourceGroupRequest struct {
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
	// test_group
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteDBResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDBResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBResourceGroupRequest) SetDBClusterId(v string) *DeleteDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetGroupName(v string) *DeleteDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetOwnerAccount(v string) *DeleteDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetOwnerId(v int64) *DeleteDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetResourceOwnerAccount(v string) *DeleteDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetResourceOwnerId(v int64) *DeleteDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
