// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *GrantCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *GrantCollectionRequest
	GetDBInstanceId() *string
	SetGrantToNamespace(v string) *GrantCollectionRequest
	GetGrantToNamespace() *string
	SetGrantType(v string) *GrantCollectionRequest
	GetGrantType() *string
	SetManagerAccount(v string) *GrantCollectionRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *GrantCollectionRequest
	GetManagerAccountPassword() *string
	SetNamespace(v string) *GrantCollectionRequest
	GetNamespace() *string
	SetOwnerId(v int64) *GrantCollectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GrantCollectionRequest
	GetRegionId() *string
}

type GrantCollectionRequest struct {
	// The name of the collection.
	//
	// >  You can call the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) operation to create a vector collection and call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query a list of vector collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The ID of the instance in reserved storage mode.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the namespace to which you want to grant the vector collection permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// othernamespace
	GrantToNamespace *string `json:"GrantToNamespace,omitempty" xml:"GrantToNamespace,omitempty"`
	// The type of the permissions that you want to grant. Valid values:
	//
	// 	- rw: the read and write permissions.
	//
	// 	- ro: the read-only permission.
	//
	// 	- none: the delete permission.
	//
	// This parameter is required.
	//
	// example:
	//
	// rw
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
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
}

func (s GrantCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantCollectionRequest) GoString() string {
	return s.String()
}

func (s *GrantCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *GrantCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *GrantCollectionRequest) GetGrantToNamespace() *string {
	return s.GrantToNamespace
}

func (s *GrantCollectionRequest) GetGrantType() *string {
	return s.GrantType
}

func (s *GrantCollectionRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *GrantCollectionRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *GrantCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GrantCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GrantCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GrantCollectionRequest) SetCollection(v string) *GrantCollectionRequest {
	s.Collection = &v
	return s
}

func (s *GrantCollectionRequest) SetDBInstanceId(v string) *GrantCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *GrantCollectionRequest) SetGrantToNamespace(v string) *GrantCollectionRequest {
	s.GrantToNamespace = &v
	return s
}

func (s *GrantCollectionRequest) SetGrantType(v string) *GrantCollectionRequest {
	s.GrantType = &v
	return s
}

func (s *GrantCollectionRequest) SetManagerAccount(v string) *GrantCollectionRequest {
	s.ManagerAccount = &v
	return s
}

func (s *GrantCollectionRequest) SetManagerAccountPassword(v string) *GrantCollectionRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *GrantCollectionRequest) SetNamespace(v string) *GrantCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *GrantCollectionRequest) SetOwnerId(v int64) *GrantCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantCollectionRequest) SetRegionId(v string) *GrantCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *GrantCollectionRequest) Validate() error {
	return dara.Validate(s)
}
