// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CreateCollectionShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateCollectionShrinkRequest
	GetDBInstanceId() *string
	SetDimension(v int64) *CreateCollectionShrinkRequest
	GetDimension() *int64
	SetExternalStorage(v int32) *CreateCollectionShrinkRequest
	GetExternalStorage() *int32
	SetFullTextRetrievalFields(v string) *CreateCollectionShrinkRequest
	GetFullTextRetrievalFields() *string
	SetHnswEfConstruction(v string) *CreateCollectionShrinkRequest
	GetHnswEfConstruction() *string
	SetHnswM(v int32) *CreateCollectionShrinkRequest
	GetHnswM() *int32
	SetManagerAccount(v string) *CreateCollectionShrinkRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateCollectionShrinkRequest
	GetManagerAccountPassword() *string
	SetMetadata(v string) *CreateCollectionShrinkRequest
	GetMetadata() *string
	SetMetadataIndices(v string) *CreateCollectionShrinkRequest
	GetMetadataIndices() *string
	SetMetrics(v string) *CreateCollectionShrinkRequest
	GetMetrics() *string
	SetNamespace(v string) *CreateCollectionShrinkRequest
	GetNamespace() *string
	SetOwnerId(v int64) *CreateCollectionShrinkRequest
	GetOwnerId() *int64
	SetParser(v string) *CreateCollectionShrinkRequest
	GetParser() *string
	SetPqEnable(v int32) *CreateCollectionShrinkRequest
	GetPqEnable() *int32
	SetRegionId(v string) *CreateCollectionShrinkRequest
	GetRegionId() *string
	SetSparseVectorIndexConfigShrink(v string) *CreateCollectionShrinkRequest
	GetSparseVectorIndexConfigShrink() *string
	SetSupportSparse(v bool) *CreateCollectionShrinkRequest
	GetSupportSparse() *bool
	SetWorkspaceId(v string) *CreateCollectionShrinkRequest
	GetWorkspaceId() *string
}

type CreateCollectionShrinkRequest struct {
	// The name of the collection that you want to create.
	//
	// >  The name must comply with the naming conventions of PostgreSQL objects.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of vector dimensions.
	//
	// >  If you specify this parameter, an index is created. When you call the [UpsertCollectionData](https://help.aliyun.com/document_detail/2401493.html) operation, make sure that the length of the Rows.Vector parameter is the same as the value of this parameter. If you do not specify this parameter, you can call the [CreateVectorIndex](https://help.aliyun.com/document_detail/2401499.html) operation to create an index.
	//
	// example:
	//
	// 1024
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// Specifies whether to use the memory mapping technology to create HNSW indexes. Valid values: 0 and 1. Default value: 0. We recommend that you set the value to 1 in scenarios that require upload speed but not data deletion.
	//
	// >
	//
	// 	- 0: uses segmented paging storage to create indexes. This method uses the shared buffer of PostgreSQL for caching and supports the delete and update operations.
	//
	// 	- 1: uses the memory mapping technology to create indexes. This method does not support the delete or update operation.
	//
	// example:
	//
	// 0
	ExternalStorage *int32 `json:"ExternalStorage,omitempty" xml:"ExternalStorage,omitempty"`
	// The fields used for full-text search. Separate multiple fields with commas (,). These fields must be keys defined in Metadata.
	//
	// example:
	//
	// title,content
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	HnswEfConstruction      *string `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors for the Hierarchical Navigable Small World (HNSW) algorithm. Valid values: 1 to 1000. In most cases, this parameter is automatically configured based on the value of the Dimension parameter. You do not need to configure this parameter.
	//
	// >  We recommend that you configure this parameter based on the value of the Dimension parameter.
	//
	// *If you set Dimension to a value less than or equal to 384, set the value of HnswM to 16.
	//
	// *If you set Dimension to a value greater than 384 and less than or equal to 768, set the value of HnswM to 32.
	//
	// *If you set Dimension to a value greater than 768 and less than or equal to 1024, set the value of HnswM to 64.
	//
	// *If you set Dimension to a value greater than 1024, set the value of HnswM to 128.
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	// Name of the management account with rds_superuser permissions.
	//
	// > You can create an account through the console -> Account Management, or by using the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) API.
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
	// The metadata of the vector data, which is a JSON string in the MAP format. The key specifies the field name, and the value specifies the data type.
	//
	// >  Supported data types:
	//
	// 	- For information about the supported data types, see [Data types](https://www.alibabacloud.com/help/zh/analyticdb/analyticdb-for-postgresql/developer-reference/data-types-1/).
	//
	// 	- The money data type is not supported.
	//
	// **
	//
	// **Warning*	- Reserved fields such as id, vector, to_tsvector, and source cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"title":"text","content":"text","response":"int"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The scalar index fields. Separate multiple fields with commas (,). These fields must be keys defined in Metadata.
	//
	// example:
	//
	// title
	MetadataIndices *string `json:"MetadataIndices,omitempty" xml:"MetadataIndices,omitempty"`
	// The method that is used to create vector indexes. Valid values:
	//
	// 	- l2: Euclidean distance.
	//
	// 	- ip: inner product distance.
	//
	// 	- cosine: cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The name of the namespace.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The analyzer that is used for full-text search.
	//
	// example:
	//
	// zh_cn
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// Specifies whether to enable the product quantization (PQ) feature for index acceleration. We recommend that you enable this feature for more than 500,000 rows of data. Valid values:
	//
	// 	- 0: no.
	//
	// 	- 1 (default): yes.
	//
	// example:
	//
	// 0
	PqEnable *int32 `json:"PqEnable,omitempty" xml:"PqEnable,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SparseVectorIndexConfigShrink *string `json:"SparseVectorIndexConfig,omitempty" xml:"SparseVectorIndexConfig,omitempty"`
	SupportSparse                 *bool   `json:"SupportSparse,omitempty" xml:"SupportSparse,omitempty"`
	// The ID of the workspace that consists of multiple AnalyticDB for PostgreSQL instances. You must specify one of the WorkspaceId and DBInstanceId parameters. If you specify both parameters, the WorkspaceId parameter takes effect.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCollectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCollectionShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateCollectionShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateCollectionShrinkRequest) GetDimension() *int64 {
	return s.Dimension
}

func (s *CreateCollectionShrinkRequest) GetExternalStorage() *int32 {
	return s.ExternalStorage
}

func (s *CreateCollectionShrinkRequest) GetFullTextRetrievalFields() *string {
	return s.FullTextRetrievalFields
}

func (s *CreateCollectionShrinkRequest) GetHnswEfConstruction() *string {
	return s.HnswEfConstruction
}

func (s *CreateCollectionShrinkRequest) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateCollectionShrinkRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateCollectionShrinkRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateCollectionShrinkRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateCollectionShrinkRequest) GetMetadataIndices() *string {
	return s.MetadataIndices
}

func (s *CreateCollectionShrinkRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *CreateCollectionShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateCollectionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCollectionShrinkRequest) GetParser() *string {
	return s.Parser
}

func (s *CreateCollectionShrinkRequest) GetPqEnable() *int32 {
	return s.PqEnable
}

func (s *CreateCollectionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCollectionShrinkRequest) GetSparseVectorIndexConfigShrink() *string {
	return s.SparseVectorIndexConfigShrink
}

func (s *CreateCollectionShrinkRequest) GetSupportSparse() *bool {
	return s.SupportSparse
}

func (s *CreateCollectionShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCollectionShrinkRequest) SetCollection(v string) *CreateCollectionShrinkRequest {
	s.Collection = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetDBInstanceId(v string) *CreateCollectionShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetDimension(v int64) *CreateCollectionShrinkRequest {
	s.Dimension = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetExternalStorage(v int32) *CreateCollectionShrinkRequest {
	s.ExternalStorage = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetFullTextRetrievalFields(v string) *CreateCollectionShrinkRequest {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetHnswEfConstruction(v string) *CreateCollectionShrinkRequest {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetHnswM(v int32) *CreateCollectionShrinkRequest {
	s.HnswM = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetManagerAccount(v string) *CreateCollectionShrinkRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetManagerAccountPassword(v string) *CreateCollectionShrinkRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetMetadata(v string) *CreateCollectionShrinkRequest {
	s.Metadata = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetMetadataIndices(v string) *CreateCollectionShrinkRequest {
	s.MetadataIndices = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetMetrics(v string) *CreateCollectionShrinkRequest {
	s.Metrics = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetNamespace(v string) *CreateCollectionShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetOwnerId(v int64) *CreateCollectionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetParser(v string) *CreateCollectionShrinkRequest {
	s.Parser = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetPqEnable(v int32) *CreateCollectionShrinkRequest {
	s.PqEnable = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetRegionId(v string) *CreateCollectionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetSparseVectorIndexConfigShrink(v string) *CreateCollectionShrinkRequest {
	s.SparseVectorIndexConfigShrink = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetSupportSparse(v bool) *CreateCollectionShrinkRequest {
	s.SupportSparse = &v
	return s
}

func (s *CreateCollectionShrinkRequest) SetWorkspaceId(v string) *CreateCollectionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCollectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
