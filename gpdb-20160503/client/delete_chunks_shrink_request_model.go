// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIdsShrink(v string) *DeleteChunksShrinkRequest
	GetChunkIdsShrink() *string
	SetCollection(v string) *DeleteChunksShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *DeleteChunksShrinkRequest
	GetDBInstanceId() *string
	SetNamespace(v string) *DeleteChunksShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *DeleteChunksShrinkRequest
	GetNamespacePassword() *string
	SetRegionId(v string) *DeleteChunksShrinkRequest
	GetRegionId() *string
}

type DeleteChunksShrinkRequest struct {
	// A list of chunk IDs.
	//
	// This parameter is required.
	ChunkIdsShrink *string `json:"ChunkIds,omitempty" xml:"ChunkIds,omitempty"`
	// The name of the document collection.
	//
	// > You create this document collection by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. To view existing document collections, call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testcollection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > Call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to view details for all AnalyticDB for PostgreSQL instances in a specific region, including their instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the namespace. The default value is public.
	//
	// > Call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > This password is set when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteChunksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunksShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteChunksShrinkRequest) GetChunkIdsShrink() *string {
	return s.ChunkIdsShrink
}

func (s *DeleteChunksShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *DeleteChunksShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteChunksShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteChunksShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *DeleteChunksShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteChunksShrinkRequest) SetChunkIdsShrink(v string) *DeleteChunksShrinkRequest {
	s.ChunkIdsShrink = &v
	return s
}

func (s *DeleteChunksShrinkRequest) SetCollection(v string) *DeleteChunksShrinkRequest {
	s.Collection = &v
	return s
}

func (s *DeleteChunksShrinkRequest) SetDBInstanceId(v string) *DeleteChunksShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteChunksShrinkRequest) SetNamespace(v string) *DeleteChunksShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteChunksShrinkRequest) SetNamespacePassword(v string) *DeleteChunksShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *DeleteChunksShrinkRequest) SetRegionId(v string) *DeleteChunksShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteChunksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
