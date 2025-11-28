// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *DeleteIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteIndexRequest
	GetDBInstanceId() *string
	SetIndexName(v string) *DeleteIndexRequest
	GetIndexName() *string
	SetNamespace(v string) *DeleteIndexRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteIndexRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *DeleteIndexRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteIndexRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *DeleteIndexRequest
	GetWorkspaceId() *string
}

type DeleteIndexRequest struct {
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
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the Elasticsearch index.
	//
	// This parameter is required.
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
	// > The value of this parameter is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available region list.
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

func (s DeleteIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *DeleteIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteIndexRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIndexRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteIndexRequest) SetCollection(v string) *DeleteIndexRequest {
	s.Collection = &v
	return s
}

func (s *DeleteIndexRequest) SetDBInstanceId(v string) *DeleteIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteIndexRequest) SetIndexName(v string) *DeleteIndexRequest {
	s.IndexName = &v
	return s
}

func (s *DeleteIndexRequest) SetNamespace(v string) *DeleteIndexRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteIndexRequest) SetNamespacePassword(v string) *DeleteIndexRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteIndexRequest) SetOwnerId(v int64) *DeleteIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteIndexRequest) SetRegionId(v string) *DeleteIndexRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIndexRequest) SetWorkspaceId(v string) *DeleteIndexRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteIndexRequest) Validate() error {
	return dara.Validate(s)
}
