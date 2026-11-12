// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeNamespaceRequest
	GetDBInstanceId() *string
	SetManagerAccount(v string) *DescribeNamespaceRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *DescribeNamespaceRequest
	GetManagerAccountPassword() *string
	SetNamespace(v string) *DescribeNamespaceRequest
	GetNamespace() *string
	SetOwnerId(v int64) *DescribeNamespaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeNamespaceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DescribeNamespaceRequest
	GetWorkspaceId() *string
}

type DescribeNamespaceRequest struct {
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
	// > You can create an account in the console by navigating to Account Management, or call the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
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
	// The namespace. Default value: public.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DescribeNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeNamespaceRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *DescribeNamespaceRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *DescribeNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeNamespaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNamespaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeNamespaceRequest) SetDBInstanceId(v string) *DescribeNamespaceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeNamespaceRequest) SetManagerAccount(v string) *DescribeNamespaceRequest {
	s.ManagerAccount = &v
	return s
}

func (s *DescribeNamespaceRequest) SetManagerAccountPassword(v string) *DescribeNamespaceRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *DescribeNamespaceRequest) SetNamespace(v string) *DescribeNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeNamespaceRequest) SetOwnerId(v int64) *DescribeNamespaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNamespaceRequest) SetRegionId(v string) *DescribeNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceRequest) SetWorkspaceId(v string) *DescribeNamespaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
