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
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The metadata of the addition or modification operation, which is in the JSON string format.
	//
	// You can specify this parameter to add a metadata definition, or rename an existing metadata definition and perform implicit type conversion.
	//
	// If you specify `operations[*].operator = add` to add a metadata definition, `operations[*].newMetaName` specifies the name of the metadata definition, and `operations[*].newMetaType` specifies the data type of the metadata definition.
	//
	// If you specify `operations[*].operator = replace` to modify an existing metadata definition, `operations[*].oldMetaName` specifies the current name of the metadata definition, `operations[*].newMetaName` specifies the new name of the metadata definition, and `operations[*].newMetaType` specifies the new data type of the metadata definition. If you only want to rename the metadata definition, you do not need to specify the `operations[*].newMetaType` field. If you only want to perform implicit type conversion, you do not need to specify the `operations[*].newMetaName` field.
	//
	// >
	//
	// 	- For information about the supported data types, see [Data types](https://help.aliyun.com/document_detail/424383.html).
	//
	// 	- The money data type is not supported.
	//
	// **
	//
	// **Warning**Reserved fields such as id, vector, to_tsvector, and source cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"operations":[
	//
	// {"operator":"add","newMetaType":"int","newMetaName":"ext1"},
	//
	// {"operator":"replace","oldMetaName":"ext2","newMetaName":"ext3"},
	//
	// {"operator":"replace","newMetaType":"bigint","oldMetaName":"ext4"},
	//
	// {"operator":"replace","newMetaType":"int","oldMetaName":"ext5","newMetaName":"ext6"}
	//
	// ]}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to query a list of namespaces.
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
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes precedence.
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
