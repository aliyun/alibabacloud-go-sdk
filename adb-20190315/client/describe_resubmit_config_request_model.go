// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResubmitConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeResubmitConfigRequest
	GetDBClusterId() *string
	SetGroupName(v string) *DescribeResubmitConfigRequest
	GetGroupName() *string
	SetOwnerAccount(v string) *DescribeResubmitConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeResubmitConfigRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeResubmitConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeResubmitConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeResubmitConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeResubmitConfigRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-8vbyw9awuj141haf9
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group name of a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929XXXX
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeResubmitConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResubmitConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeResubmitConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeResubmitConfigRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeResubmitConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeResubmitConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeResubmitConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeResubmitConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeResubmitConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeResubmitConfigRequest) SetDBClusterId(v string) *DescribeResubmitConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetGroupName(v string) *DescribeResubmitConfigRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetOwnerAccount(v string) *DescribeResubmitConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetOwnerId(v int64) *DescribeResubmitConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetResourceGroupId(v string) *DescribeResubmitConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetResourceOwnerAccount(v string) *DescribeResubmitConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeResubmitConfigRequest) SetResourceOwnerId(v int64) *DescribeResubmitConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeResubmitConfigRequest) Validate() error {
	return dara.Validate(s)
}
