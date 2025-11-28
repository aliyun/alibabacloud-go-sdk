// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DeleteVectorIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteVectorIndexRequest
	GetDBInstanceId() *string
	SetManagerAccount(v string) *DeleteVectorIndexRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *DeleteVectorIndexRequest
	GetManagerAccountPassword() *string
	SetNamespace(v string) *DeleteVectorIndexRequest
	GetNamespace() *string
	SetOwnerId(v int64) *DeleteVectorIndexRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteVectorIndexRequest
	GetRegionId() *string
	SetType(v string) *DeleteVectorIndexRequest
	GetType() *string
}

type DeleteVectorIndexRequest struct {
	// The name of the collection.
	//
	// >  You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query a list of collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the manager account that has the rds_superuser permission.
	//
	// >  You can create an account on the **Account Management*	- page of the AnalyticDB for PostgreSQL console or by calling the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
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
	// The vector type. Valid values:
	//
	// 	- Dense (default)
	//
	// 	- Sparse
	//
	// example:
	//
	// Dense
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteVectorIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteVectorIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteVectorIndexRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *DeleteVectorIndexRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *DeleteVectorIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteVectorIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVectorIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVectorIndexRequest) GetType() *string {
	return s.Type
}

func (s *DeleteVectorIndexRequest) SetCollection(v string) *DeleteVectorIndexRequest {
	s.Collection = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetDBInstanceId(v string) *DeleteVectorIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetManagerAccount(v string) *DeleteVectorIndexRequest {
	s.ManagerAccount = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetManagerAccountPassword(v string) *DeleteVectorIndexRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetNamespace(v string) *DeleteVectorIndexRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetOwnerId(v int64) *DeleteVectorIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetRegionId(v string) *DeleteVectorIndexRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVectorIndexRequest) SetType(v string) *DeleteVectorIndexRequest {
	s.Type = &v
	return s
}

func (s *DeleteVectorIndexRequest) Validate() error {
	return dara.Validate(s)
}
