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
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to view the details of all instances in the specified region, including instance IDs.
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
	// The name of the namespace to create, which is used for tenant isolation (the backend performs automatic creation of a user account with the same name).
	//
	// > The name must comply with PostgreSQL object naming rules.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
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
	// The ID of the workspace that consists of multiple database instances. This parameter and the DBInstanceId parameter cannot be left empty at the same time. If both parameters are specified, this parameter takes precedence.
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
