// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *BindDBResourceGroupWithRoleShrinkRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *BindDBResourceGroupWithRoleShrinkRequest
	GetOwnerId() *int64
	SetResourceGroupName(v string) *BindDBResourceGroupWithRoleShrinkRequest
	GetResourceGroupName() *string
	SetRoleListShrink(v string) *BindDBResourceGroupWithRoleShrinkRequest
	GetRoleListShrink() *string
}

type BindDBResourceGroupWithRoleShrinkRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// testgroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The roles.
	//
	// This parameter is required.
	RoleListShrink *string `json:"RoleList,omitempty" xml:"RoleList,omitempty"`
}

func (s BindDBResourceGroupWithRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) GetRoleListShrink() *string {
	return s.RoleListShrink
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) SetDBInstanceId(v string) *BindDBResourceGroupWithRoleShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) SetOwnerId(v int64) *BindDBResourceGroupWithRoleShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) SetResourceGroupName(v string) *BindDBResourceGroupWithRoleShrinkRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) SetRoleListShrink(v string) *BindDBResourceGroupWithRoleShrinkRequest {
	s.RoleListShrink = &v
	return s
}

func (s *BindDBResourceGroupWithRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
