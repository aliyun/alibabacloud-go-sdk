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
	SetGroupType(v string) *DescribeDBResourceGroupRequest
	GetGroupType() *string
	SetRegionId(v string) *DescribeDBResourceGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBResourceGroupRequest
	GetResourceOwnerAccount() *string
}

type DescribeDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// > If you do not specify this parameter, the information about all resource groups in the cluster is returned.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- **Interactive**
	//
	// 	- **Job**
	//
	// >  For more information about resource groups, see [Resource group overview](https://help.aliyun.com/document_detail/428610.html).
	//
	// example:
	//
	// Job
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The region ID of the cluster.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
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

func (s *DescribeDBResourceGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeDBResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBResourceGroupRequest) SetDBClusterId(v string) *DescribeDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetGroupName(v string) *DescribeDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetGroupType(v string) *DescribeDBResourceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetRegionId(v string) *DescribeDBResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) SetResourceOwnerAccount(v string) *DescribeDBResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
