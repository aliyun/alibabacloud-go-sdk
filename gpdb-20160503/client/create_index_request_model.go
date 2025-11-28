// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CreateIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateIndexRequest
	GetDBInstanceId() *string
	SetIndexConfig(v string) *CreateIndexRequest
	GetIndexConfig() *string
	SetIndexField(v string) *CreateIndexRequest
	GetIndexField() *string
	SetIndexName(v string) *CreateIndexRequest
	GetIndexName() *string
	SetNamespace(v string) *CreateIndexRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *CreateIndexRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *CreateIndexRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIndexRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CreateIndexRequest
	GetWorkspaceId() *string
}

type CreateIndexRequest struct {
	// The name of the collection.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query a list of collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in the specified region.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The index parameter. If you do not specify this parameter, a B-tree index is created.
	//
	// >
	//
	// 	- b-tree: To create a B-tree index, set the fillFactor parameter to a value between 10 and 100. Default value: 90.
	//
	// 	- gin: To create a GIN index, set the fastUpdate parameter to true or false. Default value: true.
	IndexConfig *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// The index field. Only a single field is supported, and it must be a key defined in metadata.
	//
	// example:
	//
	// title
	IndexField *string `json:"IndexField,omitempty" xml:"IndexField,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// testindex
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The namespace name.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > The value of this parameter is specified by the CreateNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
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
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. This parameter and DBInstanceId cannot both be empty. If both parameters are specified, this value takes precedence.
	//
	// Valid values:
	//
	// 	- ip
	//
	// 	- l2
	//
	// 	- cosine
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateIndexRequest) GetIndexConfig() *string {
	return s.IndexConfig
}

func (s *CreateIndexRequest) GetIndexField() *string {
	return s.IndexField
}

func (s *CreateIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *CreateIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIndexRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *CreateIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIndexRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateIndexRequest) SetCollection(v string) *CreateIndexRequest {
	s.Collection = &v
	return s
}

func (s *CreateIndexRequest) SetDBInstanceId(v string) *CreateIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateIndexRequest) SetIndexConfig(v string) *CreateIndexRequest {
	s.IndexConfig = &v
	return s
}

func (s *CreateIndexRequest) SetIndexField(v string) *CreateIndexRequest {
	s.IndexField = &v
	return s
}

func (s *CreateIndexRequest) SetIndexName(v string) *CreateIndexRequest {
	s.IndexName = &v
	return s
}

func (s *CreateIndexRequest) SetNamespace(v string) *CreateIndexRequest {
	s.Namespace = &v
	return s
}

func (s *CreateIndexRequest) SetNamespacePassword(v string) *CreateIndexRequest {
	s.NamespacePassword = &v
	return s
}

func (s *CreateIndexRequest) SetOwnerId(v int64) *CreateIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIndexRequest) SetRegionId(v string) *CreateIndexRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIndexRequest) SetWorkspaceId(v string) *CreateIndexRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateIndexRequest) Validate() error {
	return dara.Validate(s)
}
