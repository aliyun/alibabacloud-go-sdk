// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *BindDBResourceGroupWithRoleRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *BindDBResourceGroupWithRoleRequest
	GetOwnerId() *int64
	SetResourceGroupName(v string) *BindDBResourceGroupWithRoleRequest
	GetResourceGroupName() *string
	SetRoleList(v []*string) *BindDBResourceGroupWithRoleRequest
	GetRoleList() []*string
}

type BindDBResourceGroupWithRoleRequest struct {
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

func (s BindDBResourceGroupWithRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithRoleRequest) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithRoleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *BindDBResourceGroupWithRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDBResourceGroupWithRoleRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *BindDBResourceGroupWithRoleRequest) GetRoleList() []*string {
	return s.RoleList
}

func (s *BindDBResourceGroupWithRoleRequest) SetDBInstanceId(v string) *BindDBResourceGroupWithRoleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *BindDBResourceGroupWithRoleRequest) SetOwnerId(v int64) *BindDBResourceGroupWithRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDBResourceGroupWithRoleRequest) SetResourceGroupName(v string) *BindDBResourceGroupWithRoleRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *BindDBResourceGroupWithRoleRequest) SetRoleList(v []*string) *BindDBResourceGroupWithRoleRequest {
	s.RoleList = v
	return s
}

func (s *BindDBResourceGroupWithRoleRequest) Validate() error {
	return dara.Validate(s)
}
