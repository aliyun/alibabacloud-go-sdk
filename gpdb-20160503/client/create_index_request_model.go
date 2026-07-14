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
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a region.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The index parameters. If this parameter is not specified, a B-tree index is created by default.
	//
	// >
	//
	// >- b-tree: The fillFactor parameter. Valid values: 10 to 100. Default value: 90.
	//
	// >- gin: The fastUpdate parameter. Valid values: true and false. Default value: true.
	//
	// example:
	//
	// {"b-tree":{"fillFactor":90}} or {"gin":{"fastUpdate":false}}
	IndexConfig *string `json:"IndexConfig,omitempty" xml:"IndexConfig,omitempty"`
	// The index field. Only a single field is supported, and the field must be a key defined in Metadata.
	//
	// example:
	//
	// title
	IndexField *string `json:"IndexField,omitempty" xml:"IndexField,omitempty"`
	// The index name.
	//
	// example:
	//
	// testindex
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The namespace.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > This value is specified by the CreateNamespace operation.
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
	// The ID of the workspace that consists of multiple database instances. This parameter and DBInstanceId cannot both be empty. If both this parameter and DBInstanceId are specified, this parameter takes precedence.
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
