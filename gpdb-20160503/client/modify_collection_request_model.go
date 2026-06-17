// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *ModifyCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *ModifyCollectionRequest
	GetDBInstanceId() *string
	SetMetadata(v string) *ModifyCollectionRequest
	GetMetadata() *string
	SetNamespace(v string) *ModifyCollectionRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *ModifyCollectionRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *ModifyCollectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCollectionRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ModifyCollectionRequest
	GetWorkspaceId() *string
}

type ModifyCollectionRequest struct {
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to list all collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details for all AnalyticDB for PostgreSQL instances in a region, including their instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A JSON string that specifies the operations to add or modify metadata fields. For the required format, see the example.
	//
	// Use this parameter to add new metadata fields, rename existing metadata fields, or perform implicit data type conversion on existing fields.
	//
	// Details:
	//
	// To add a new metadata field, set `operations[*].operator = add`. Then, use `operations[*].newMetaName` to specify the field\\"s name, `operations[*].newMetaType` for its data type, and `operations[*].fullTextRetrieval` to enable full-text retrieval for it.
	//
	// To modify an existing metadata field, set `operations[*].operator = replace`. You must specify the current field name in `operations[*].oldMetaName`. To rename the field, provide the new name in `operations[*].newMetaName`. To change its data type, provide the new type in `operations[*].newMetaType`.
	//
	// > - For a list of supported data types, see [Data types](https://help.aliyun.com/document_detail/424383.html). The money data type is not supported.
	//
	// >
	//
	// > - Full-text retrieval can be enabled for a field only during an `add` operation, not a `replace` operation.
	//
	// 	Warning:
	//
	// The field names `id`, `vector`, `to_tsvector`, and `source` are reserved.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"operations":[
	//
	// {"operator":"add","newMetaType":"int","newMetaName":"ext1","fullTextRetrieval":true},
	//
	// {"operator":"replace","oldMetaName":"ext2","newMetaName":"ext3"},
	//
	// {"operator":"replace","newMetaType":"bigint","oldMetaName":"ext4"},
	//
	// {"operator":"replace","newMetaType":"int","oldMetaName":"ext5","newMetaName":"ext6"}
	//
	// ]}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The namespace. The default value is `public`.
	//
	// > You can call the CreateNamespace operation to create a namespace and the ListNamespaces operation to list existing namespaces.
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
	// The ID of the workspace that contains multiple database instances. You must specify either this parameter or `DBInstanceId`. If you specify both, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ModifyCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCollectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *ModifyCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyCollectionRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *ModifyCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyCollectionRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ModifyCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCollectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCollectionRequest) SetCollection(v string) *ModifyCollectionRequest {
	s.Collection = &v
	return s
}

func (s *ModifyCollectionRequest) SetDBInstanceId(v string) *ModifyCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyCollectionRequest) SetMetadata(v string) *ModifyCollectionRequest {
	s.Metadata = &v
	return s
}

func (s *ModifyCollectionRequest) SetNamespace(v string) *ModifyCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyCollectionRequest) SetNamespacePassword(v string) *ModifyCollectionRequest {
	s.NamespacePassword = &v
	return s
}

func (s *ModifyCollectionRequest) SetOwnerId(v int64) *ModifyCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCollectionRequest) SetRegionId(v string) *ModifyCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCollectionRequest) SetWorkspaceId(v string) *ModifyCollectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCollectionRequest) Validate() error {
	return dara.Validate(s)
}
