// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentCollectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *CreateDocumentCollectionShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateDocumentCollectionShrinkRequest
	GetDBInstanceId() *string
	SetDimension(v int32) *CreateDocumentCollectionShrinkRequest
	GetDimension() *int32
	SetEmbeddingModel(v string) *CreateDocumentCollectionShrinkRequest
	GetEmbeddingModel() *string
	SetEnableGraph(v bool) *CreateDocumentCollectionShrinkRequest
	GetEnableGraph() *bool
	SetEntityTypesShrink(v string) *CreateDocumentCollectionShrinkRequest
	GetEntityTypesShrink() *string
	SetExternalStorage(v int32) *CreateDocumentCollectionShrinkRequest
	GetExternalStorage() *int32
	SetFullTextRetrievalFields(v string) *CreateDocumentCollectionShrinkRequest
	GetFullTextRetrievalFields() *string
	SetHnswEfConstruction(v string) *CreateDocumentCollectionShrinkRequest
	GetHnswEfConstruction() *string
	SetHnswM(v int32) *CreateDocumentCollectionShrinkRequest
	GetHnswM() *int32
	SetLLMModel(v string) *CreateDocumentCollectionShrinkRequest
	GetLLMModel() *string
	SetLanguage(v string) *CreateDocumentCollectionShrinkRequest
	GetLanguage() *string
	SetManagerAccount(v string) *CreateDocumentCollectionShrinkRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateDocumentCollectionShrinkRequest
	GetManagerAccountPassword() *string
	SetMetadata(v string) *CreateDocumentCollectionShrinkRequest
	GetMetadata() *string
	SetMetadataIndices(v string) *CreateDocumentCollectionShrinkRequest
	GetMetadataIndices() *string
	SetMetrics(v string) *CreateDocumentCollectionShrinkRequest
	GetMetrics() *string
	SetNamespace(v string) *CreateDocumentCollectionShrinkRequest
	GetNamespace() *string
	SetOwnerId(v int64) *CreateDocumentCollectionShrinkRequest
	GetOwnerId() *int64
	SetParser(v string) *CreateDocumentCollectionShrinkRequest
	GetParser() *string
	SetPqEnable(v int32) *CreateDocumentCollectionShrinkRequest
	GetPqEnable() *int32
	SetRegionId(v string) *CreateDocumentCollectionShrinkRequest
	GetRegionId() *string
	SetRelationshipTypesShrink(v string) *CreateDocumentCollectionShrinkRequest
	GetRelationshipTypesShrink() *string
}

type CreateDocumentCollectionShrinkRequest struct {
	// The name of the document collection that you want to create.
	//
	// > The name must comply with PostgreSQL object naming restrictions.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB for PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Dimension    *int32  `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The vectorization algorithm.
	//
	// >  Supported algorithms:
	//
	// 	- text-embedding-v1: the algorithm that produces 1536-dimensional vectors.
	//
	// 	- text-embedding-v2: the algorithm that produces 1536-dimensional vectors.
	//
	// 	- text2vec: the algorithm that produces 1024-dimensional vectors.
	//
	// 	- m3e-base: the algorithm that produces 768-dimensional vectors.
	//
	// 	- m3e-small: the algorithm that produces 512-dimensional vectors.
	//
	// 	- clip-vit-b-32: the image vectorization algorithm that uses the Contrastive Language-Image Pre-Training (CLIP) ViT-B/32 model and produces 512-dimensional vectors.
	//
	// 	- clip-vit-b-16: the image vectorization algorithm that uses the CLIP ViT-B/16 model and produces 512-dimensional vectors.
	//
	// 	- clip-vit-l-14: the image vectorization algorithm that uses the CLIP ViT-L/14 model and produces 768-dimensional vectors.
	//
	// 	- clip-vit-l-14-336px: the image vectorization algorithm that uses the CLIP ViT-L/14@336px model and produces 768-dimensional vectors.
	//
	// 	- clip-rn50: the image vectorization algorithm that uses the CLIP RN50 model and produces 1024-dimensional vectors.
	//
	// 	- clip-rn101: the image vectorization algorithm that uses the CLIP RN101 model and produces 512-dimensional vectors.
	//
	// 	- clip-rn50x4: the image vectorization algorithm that uses the CLIP RN50x4 model and produces 640-dimensional vectors.
	//
	// 	- clip-rn50x16: the image vectorization algorithm that uses the CLIP RN50x16 model and produces 768-dimensional vectors.
	//
	// 	- clip-rn50x64: the image vectorization algorithm that uses the CLIP RN50x64 model and produces 1024-dimensional vectors.
	//
	// example:
	//
	// text-embedding-v1
	EmbeddingModel    *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	EnableGraph       *bool   `json:"EnableGraph,omitempty" xml:"EnableGraph,omitempty"`
	EntityTypesShrink *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
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
	// title,page
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	HnswEfConstruction      *string `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors for the Hierarchical Navigable Small World (HNSW) algorithm. Valid values: 1 to 1000. In most cases, this parameter is automatically configured based on the value of the Dimension parameter. You do not need to configure this parameter.
	//
	// >  We recommend that you configure this parameter based on the value of the Dimension parameter.
	//
	// 	- If you set Dimension to a value less than or equal to 384, set the value of HnswM to 16.
	//
	// 	- If you set Dimension to a value greater than 384 and less than or equal to 768, set the value of HnswM to 32.
	//
	// 	- If you set Dimension to a value greater than 768 and less than or equal to 1024, set the value of HnswM to 64.
	//
	// 	- If you set Dimension to a value greater than 1024, set the value of HnswM to 128.
	//
	// example:
	//
	// 64
	HnswM    *int32  `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the manager account that has the rds_superuser permission.
	//
	// > You can create an account through the console -> Account Management, or by using the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	// The password of the management account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	// The metadata of the vector data, which is a JSON string in the MAP format. The key specifies the field name, and the value specifies the data type.
	//
	// > Supported data types:
	//
	// > - For information about data types, see: [Data Types](https://www.alibabacloud.com/help/en/analyticdb/analyticdb-for-postgresql/developer-reference/data-types-1/).
	//
	// > - The money type is not supported.
	//
	// 	Warning: The fields id, vector, doc_name, content, loader_metadata, source, and to_tsvector are reserved and should not be used.
	//
	// example:
	//
	// {"title":"text","page":"int"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// title
	MetadataIndices *string `json:"MetadataIndices,omitempty" xml:"MetadataIndices,omitempty"`
	// The method that is used to create vector indexes.
	//
	// Valid values:
	//
	// 	- **l2**: Euclidean distance.
	//
	// 	- **ip**: inner product distance.
	//
	// 	- **cosine*	- (default): cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The analyzer that is used for full-text search. Default value: zh_cn.
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
	// 1
	PqEnable *int32 `json:"PqEnable,omitempty" xml:"PqEnable,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RelationshipTypesShrink *string `json:"RelationshipTypes,omitempty" xml:"RelationshipTypes,omitempty"`
}

func (s CreateDocumentCollectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateDocumentCollectionShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDocumentCollectionShrinkRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *CreateDocumentCollectionShrinkRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateDocumentCollectionShrinkRequest) GetEnableGraph() *bool {
	return s.EnableGraph
}

func (s *CreateDocumentCollectionShrinkRequest) GetEntityTypesShrink() *string {
	return s.EntityTypesShrink
}

func (s *CreateDocumentCollectionShrinkRequest) GetExternalStorage() *int32 {
	return s.ExternalStorage
}

func (s *CreateDocumentCollectionShrinkRequest) GetFullTextRetrievalFields() *string {
	return s.FullTextRetrievalFields
}

func (s *CreateDocumentCollectionShrinkRequest) GetHnswEfConstruction() *string {
	return s.HnswEfConstruction
}

func (s *CreateDocumentCollectionShrinkRequest) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateDocumentCollectionShrinkRequest) GetLLMModel() *string {
	return s.LLMModel
}

func (s *CreateDocumentCollectionShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateDocumentCollectionShrinkRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateDocumentCollectionShrinkRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateDocumentCollectionShrinkRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateDocumentCollectionShrinkRequest) GetMetadataIndices() *string {
	return s.MetadataIndices
}

func (s *CreateDocumentCollectionShrinkRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *CreateDocumentCollectionShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateDocumentCollectionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDocumentCollectionShrinkRequest) GetParser() *string {
	return s.Parser
}

func (s *CreateDocumentCollectionShrinkRequest) GetPqEnable() *int32 {
	return s.PqEnable
}

func (s *CreateDocumentCollectionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocumentCollectionShrinkRequest) GetRelationshipTypesShrink() *string {
	return s.RelationshipTypesShrink
}

func (s *CreateDocumentCollectionShrinkRequest) SetCollection(v string) *CreateDocumentCollectionShrinkRequest {
	s.Collection = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetDBInstanceId(v string) *CreateDocumentCollectionShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetDimension(v int32) *CreateDocumentCollectionShrinkRequest {
	s.Dimension = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetEmbeddingModel(v string) *CreateDocumentCollectionShrinkRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetEnableGraph(v bool) *CreateDocumentCollectionShrinkRequest {
	s.EnableGraph = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetEntityTypesShrink(v string) *CreateDocumentCollectionShrinkRequest {
	s.EntityTypesShrink = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetExternalStorage(v int32) *CreateDocumentCollectionShrinkRequest {
	s.ExternalStorage = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetFullTextRetrievalFields(v string) *CreateDocumentCollectionShrinkRequest {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetHnswEfConstruction(v string) *CreateDocumentCollectionShrinkRequest {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetHnswM(v int32) *CreateDocumentCollectionShrinkRequest {
	s.HnswM = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetLLMModel(v string) *CreateDocumentCollectionShrinkRequest {
	s.LLMModel = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetLanguage(v string) *CreateDocumentCollectionShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetManagerAccount(v string) *CreateDocumentCollectionShrinkRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetManagerAccountPassword(v string) *CreateDocumentCollectionShrinkRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetMetadata(v string) *CreateDocumentCollectionShrinkRequest {
	s.Metadata = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetMetadataIndices(v string) *CreateDocumentCollectionShrinkRequest {
	s.MetadataIndices = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetMetrics(v string) *CreateDocumentCollectionShrinkRequest {
	s.Metrics = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetNamespace(v string) *CreateDocumentCollectionShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetOwnerId(v int64) *CreateDocumentCollectionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetParser(v string) *CreateDocumentCollectionShrinkRequest {
	s.Parser = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetPqEnable(v int32) *CreateDocumentCollectionShrinkRequest {
	s.PqEnable = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetRegionId(v string) *CreateDocumentCollectionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetRelationshipTypesShrink(v string) *CreateDocumentCollectionShrinkRequest {
	s.RelationshipTypesShrink = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
