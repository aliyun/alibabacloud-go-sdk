// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBResourceGroupRequest
	GetDBClusterId() *string
	SetGroupName(v string) *DescribeDBResourceGroupRequest
	GetGroupName() *string
	SetOwnerAccount(v string) *DescribeDBResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBResourceGroupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDBResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBResourceGroupRequest
	GetResourceOwnerId() *int64
}

type DescribeDBResourceGroupRequest struct {
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
	// example:
	//
	// USER_DEFAULT
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDBResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBResourceGroupRequest) SetDBClusterId(v string) *DescribeDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetGroupName(v string) *DescribeDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetOwnerAccount(v string) *DescribeDBResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetOwnerId(v int64) *DescribeDBResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceOwnerAccount(v string) *DescribeDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceOwnerId(v int64) *DescribeDBResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
