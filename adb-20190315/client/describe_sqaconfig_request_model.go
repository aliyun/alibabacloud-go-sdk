// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQAConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSQAConfigRequest
	GetDBClusterId() *string
	SetGroupName(v string) *DescribeSQAConfigRequest
	GetGroupName() *string
	SetOwnerAccount(v string) *DescribeSQAConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSQAConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSQAConfigRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSQAConfigRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSQAConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSQAConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeSQAConfigRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-8vbyw9awuj141haf9
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/612410.html) operation to query the resource group name of a cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeSQAConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQAConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQAConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSQAConfigRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSQAConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSQAConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSQAConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSQAConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSQAConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSQAConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSQAConfigRequest) SetDBClusterId(v string) *DescribeSQAConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetGroupName(v string) *DescribeSQAConfigRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetOwnerAccount(v string) *DescribeSQAConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetOwnerId(v int64) *DescribeSQAConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetRegionId(v string) *DescribeSQAConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetResourceGroupId(v string) *DescribeSQAConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetResourceOwnerAccount(v string) *DescribeSQAConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSQAConfigRequest) SetResourceOwnerId(v int64) *DescribeSQAConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSQAConfigRequest) Validate() error {
	return dara.Validate(s)
}
