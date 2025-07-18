// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateNamespaceRequest
	GetDBInstanceId() *string
	SetManagerAccount(v string) *CreateNamespaceRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateNamespaceRequest
	GetManagerAccountPassword() *string
	SetNamespace(v string) *CreateNamespaceRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CreateNamespaceRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CreateNamespaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNamespaceRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CreateNamespaceRequest
	GetWorkspaceId() *string
}

type CreateNamespaceRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	// The name of the namespace. After the namespace is created, the system automatically creates an account that has the same name.
	//
	// >  The name must comply with the naming conventions of PostgreSQL objects.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword2
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateNamespaceRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateNamespaceRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateNamespaceRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CreateNamespaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNamespaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateNamespaceRequest) SetDBInstanceId(v string) *CreateNamespaceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetManagerAccount(v string) *CreateNamespaceRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateNamespaceRequest) SetManagerAccountPassword(v string) *CreateNamespaceRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespace(v string) *CreateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespacePassword(v string) *CreateNamespaceRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CreateNamespaceRequest) SetOwnerId(v int64) *CreateNamespaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNamespaceRequest) SetRegionId(v string) *CreateNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNamespaceRequest) SetWorkspaceId(v string) *CreateNamespaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
