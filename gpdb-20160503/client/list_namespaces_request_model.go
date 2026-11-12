// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ListNamespacesRequest
	GetDBInstanceId() *string
	SetManagerAccount(v string) *ListNamespacesRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *ListNamespacesRequest
	GetManagerAccountPassword() *string
	SetOwnerId(v int64) *ListNamespacesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListNamespacesRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListNamespacesRequest
	GetWorkspaceId() *string
}

type ListNamespacesRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// example:
	//
	// gp-bp152460513z****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the administrator account that has the rds_superuser permission.
	//
	// > You can create an account by navigating to Account Management in the console, or call the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	// The password of the administrator account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the workspace that consists of multiple database instances. This parameter and the DBInstanceId parameter cannot be both empty. If both parameters are specified, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListNamespacesRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *ListNamespacesRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *ListNamespacesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListNamespacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNamespacesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListNamespacesRequest) SetDBInstanceId(v string) *ListNamespacesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListNamespacesRequest) SetManagerAccount(v string) *ListNamespacesRequest {
	s.ManagerAccount = &v
	return s
}

func (s *ListNamespacesRequest) SetManagerAccountPassword(v string) *ListNamespacesRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *ListNamespacesRequest) SetOwnerId(v int64) *ListNamespacesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListNamespacesRequest) SetRegionId(v string) *ListNamespacesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNamespacesRequest) SetWorkspaceId(v string) *ListNamespacesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
