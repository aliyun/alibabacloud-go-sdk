// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteNamespaceRequest
	GetDBInstanceId() *string
	SetManagerAccount(v string) *DeleteNamespaceRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *DeleteNamespaceRequest
	GetManagerAccountPassword() *string
	SetNamespace(v string) *DeleteNamespaceRequest
	GetNamespace() *string
	SetOwnerId(v int64) *DeleteNamespaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNamespaceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteNamespaceRequest
	GetWorkspaceId() *string
}

type DeleteNamespaceRequest struct {
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
	// The name of the namespace.
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

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteNamespaceRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *DeleteNamespaceRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *DeleteNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteNamespaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNamespaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteNamespaceRequest) SetDBInstanceId(v string) *DeleteNamespaceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetManagerAccount(v string) *DeleteNamespaceRequest {
	s.ManagerAccount = &v
	return s
}

func (s *DeleteNamespaceRequest) SetManagerAccountPassword(v string) *DeleteNamespaceRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequest) SetOwnerId(v int64) *DeleteNamespaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetRegionId(v string) *DeleteNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetWorkspaceId(v string) *DeleteNamespaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
