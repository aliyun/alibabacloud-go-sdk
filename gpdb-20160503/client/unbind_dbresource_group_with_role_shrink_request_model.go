// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UnbindDBResourceGroupWithRoleShrinkRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UnbindDBResourceGroupWithRoleShrinkRequest
	GetOwnerId() *int64
	SetResourceGroupName(v string) *UnbindDBResourceGroupWithRoleShrinkRequest
	GetResourceGroupName() *string
	SetRoleListShrink(v string) *UnbindDBResourceGroupWithRoleShrinkRequest
	GetRoleListShrink() *string
}

type UnbindDBResourceGroupWithRoleShrinkRequest struct {
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

func (s UnbindDBResourceGroupWithRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) GetRoleListShrink() *string {
	return s.RoleListShrink
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) SetDBInstanceId(v string) *UnbindDBResourceGroupWithRoleShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) SetOwnerId(v int64) *UnbindDBResourceGroupWithRoleShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) SetResourceGroupName(v string) *UnbindDBResourceGroupWithRoleShrinkRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) SetRoleListShrink(v string) *UnbindDBResourceGroupWithRoleShrinkRequest {
	s.RoleListShrink = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
