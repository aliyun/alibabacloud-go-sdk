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
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the manager account that has the rds_superuser permission.
	//
	// >  You can create an account on the Account Management page of the AnalyticDB for PostgreSQL console or by calling the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	// The password of the manager account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
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
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes effect.
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
