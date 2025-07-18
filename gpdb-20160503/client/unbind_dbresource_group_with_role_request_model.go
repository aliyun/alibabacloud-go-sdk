// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UnbindDBResourceGroupWithRoleRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *UnbindDBResourceGroupWithRoleRequest
	GetOwnerId() *int64
	SetResourceGroupName(v string) *UnbindDBResourceGroupWithRoleRequest
	GetResourceGroupName() *string
	SetRoleList(v []*string) *UnbindDBResourceGroupWithRoleRequest
	GetRoleList() []*string
}

type UnbindDBResourceGroupWithRoleRequest struct {
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
	RoleList []*string `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
}

func (s UnbindDBResourceGroupWithRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithRoleRequest) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithRoleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UnbindDBResourceGroupWithRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindDBResourceGroupWithRoleRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *UnbindDBResourceGroupWithRoleRequest) GetRoleList() []*string {
	return s.RoleList
}

func (s *UnbindDBResourceGroupWithRoleRequest) SetDBInstanceId(v string) *UnbindDBResourceGroupWithRoleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleRequest) SetOwnerId(v int64) *UnbindDBResourceGroupWithRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleRequest) SetResourceGroupName(v string) *UnbindDBResourceGroupWithRoleRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleRequest) SetRoleList(v []*string) *UnbindDBResourceGroupWithRoleRequest {
	s.RoleList = v
	return s
}

func (s *UnbindDBResourceGroupWithRoleRequest) Validate() error {
	return dara.Validate(s)
}
