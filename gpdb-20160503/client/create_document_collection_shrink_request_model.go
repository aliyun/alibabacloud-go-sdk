// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentCollectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateDocumentCollectionShrinkRequest
	GetAlgorithm() *string
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
	SetSparseRetrievalFields(v string) *CreateDocumentCollectionShrinkRequest
	GetSparseRetrievalFields() *string
	SetSparseVectorIndexConfigShrink(v string) *CreateDocumentCollectionShrinkRequest
	GetSparseVectorIndexConfigShrink() *string
	SetSupportSparse(v bool) *CreateDocumentCollectionShrinkRequest
	GetSupportSparse() *bool
	SetVectorIndexConfigShrink(v string) *CreateDocumentCollectionShrinkRequest
	GetVectorIndexConfigShrink() *string
}

type CreateDocumentCollectionShrinkRequest struct {
	// The vector index algorithm.
	//
	// Valid values:
	//
	// - hnswflat: HNSW index without quantization compression (default).
	//
	// - novam: graph index without quantization compression, suitable for high-performance scenarios such as real-time recommendations.
	//
	// - novad: partitioned index with RaBitQ quantization, suitable for large-scale low-cost retrieval scenarios.
	//
	// example:
	//
	// hnswflat
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The name of the knowledge base to create.
	//
	// > The name must comply with PostgreSQL object naming conventions.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The vector dimensions. The default value is the dimension supported by the embedding model.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The embedding model. Default value: text-embedding-v3.
	//
	// > Supported models:
	//
	// > - text-embedding-v3 (recommended, default): 1024, 768, or 512 dimensions
	//
	// > - multimodal-embedding-v1 (recommended): 1024 dimensions, multimodal embedding model
	//
	// > - text-embedding-v1: 1536 dimensions
	//
	// > - text-embedding-v2: 1536 dimensions
	//
	// > - text2vec (not recommended): 1024 dimensions
	//
	// > - m3e-base (not recommended): 768 dimensions
	//
	// > - m3e-small (not recommended): 512 dimensions
	//
	// > - clip-vit-b-32 (not recommended): CLIP ViT-B/32 model, 512 dimensions, image embedding model
	//
	// > - clip-vit-b-16 (not recommended): CLIP ViT-B/16 model, 512 dimensions, image embedding model
	//
	// > - clip-vit-l-14 (not recommended): CLIP ViT-L/14 model, 768 dimensions, image embedding model
	//
	// > - clip-vit-l-14-336px (not recommended): CLIP ViT-L/14@336px model, 768 dimensions, image embedding model
	//
	// > - clip-rn50 (not recommended): CLIP RN50 model, 1024 dimensions, image embedding model
	//
	// > - clip-rn101 (not recommended): CLIP RN101 model, 512 dimensions, image embedding model
	//
	// > - clip-rn50x4 (not recommended): CLIP RN50x4 model, 640 dimensions, image embedding model
	//
	// > - clip-rn50x16 (not recommended): CLIP RN50x16 model, 768 dimensions, image embedding model
	//
	// > - clip-rn50x64 (not recommended): CLIP RN50x64 model, 1024 dimensions, image embedding model
	//
	// example:
	//
	// text-embedding-v1
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// Specifies whether to enable knowledge graph construction. Default value: false.
	//
	// > Before using this parameter, upgrade the instance to a version that supports the graph engine. (During the public preview, submit a ticket to upgrade the version.)
	//
	// example:
	//
	// true
	EnableGraph *bool `json:"EnableGraph,omitempty" xml:"EnableGraph,omitempty"`
	// The list of entity types.
	//
	// > This parameter is required when knowledge graph construction is enabled.
	//
	// example:
	//
	// Location
	EntityTypesShrink *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
	// Specifies whether to use mmap to build the HNSW index. Default value: 0. If data does not need to be deleted and you require high upload performance, set this parameter to 1.
	//
	// Valid values:
	//
	// - 0: uses segment-page storage to build the index. This mode uses shared_buffer in PostgreSQL as cache and supports delete and update operations.
	//
	// - 1: uses mmap to build the index. This mode does not support delete or update operations.
	//
	// 	Notice: Only version 6.0 supports the ExternalStorage parameter. Version 7.0 does not support this parameter.
	//
	// example:
	//
	// 0
	ExternalStorage *int32 `json:"ExternalStorage,omitempty" xml:"ExternalStorage,omitempty"`
	// The fields used for full-text retrieval. Separate multiple fields with commas (,). The fields must be keys defined in Metadata.
	//
	// example:
	//
	// title,page
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	// The candidate set size when building an index with the HNSW algorithm. The value must be >= 2*HNSW_M.
	//
	// > Valid values:
	//
	// >- AnalyticDB for PostgreSQL 6.0 instances: 40 to 4000.
	//
	// >- AnalyticDB for PostgreSQL 7.0 instances: 4 to 1000. Default value: 64.
	//
	// example:
	//
	// 128
	HnswEfConstruction *string `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors in the HNSW algorithm. This value is automatically set based on the vector dimensions. Manual configuration is generally not required.
	//
	// > Valid values:
	//
	// >- AnalyticDB for PostgreSQL 6.0 instances: 1 to 1000.
	//
	// >- AnalyticDB for PostgreSQL 7.0 instances: 2 to 100. Default value: 16.
	//
	// > Recommended values based on vector dimensions:
	//
	// >- 384 or fewer: 16
	//
	// >- Greater than 384 and up to 768: 32
	//
	// >- Greater than 768 and up to 1024: 64
	//
	// >- Greater than 1024: 128
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	// The LLM model name. Valid values:
	//
	// - knowledge-extract-standard: default value.
	//
	// - knowledge-extract-mini
	//
	// > This parameter takes effect only when knowledge graph construction is enabled.
	//
	// example:
	//
	// knowledge-extract-standard
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	// The language used for knowledge graph construction. Valid values:
	//
	// - Simplified Chinese: Simplified Chinese. Default value.
	//
	// - English: English.
	//
	// > This parameter takes effect only when knowledge graph construction is enabled.
	//
	// example:
	//
	// Simplified Chinese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the management account that has the rds_superuser permission.
	//
	// > You can create an account in the console by navigating to Account Management, or by calling the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
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
	// The metadata of vector data, in the format of a JSON string representing a MAP. The key represents the field name, and the value represents the data type.
	//
	// > Supported data types:
	//
	// > - For the list of data types, see [Data types](https://help.aliyun.com/document_detail/424383.html).
	//
	// > - The money type is not supported.
	//
	// 	Warning: The following fields are reserved and cannot be used: id, vector, doc_name, content, loader_metadata, source, and to_tsvector.
	//
	// example:
	//
	// {"title":"text","page":"int"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The scalar index fields. Separate multiple fields with commas (,). The fields must be keys defined in Metadata.
	//
	// example:
	//
	// title
	MetadataIndices *string `json:"MetadataIndices,omitempty" xml:"MetadataIndices,omitempty"`
	// The distance metric used for building vector indexes.
	//
	// Valid values:
	//
	// - **l2**: Euclidean distance.
	//
	// - **ip**: inner product distance.
	//
	// - **cosine*	- (default): cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can create a namespace by calling the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation and query the list of namespaces by calling the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The tokenizer used for full-text retrieval. Default value: zh_cn.
	//
	// example:
	//
	// zh_cn
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// Specifies whether to enable Product Quantization (PQ) algorithm acceleration for the index. We recommend enabling this feature when the data volume exceeds 500,000. Valid values:
	//
	// - 0: disabled.
	//
	// - 1: enabled (default).
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of relationship edge types.
	//
	// > This parameter is required when knowledge graph construction is enabled.
	//
	// example:
	//
	// Occurred
	RelationshipTypesShrink *string `json:"RelationshipTypes,omitempty" xml:"RelationshipTypes,omitempty"`
	// The metadata fields used for building sparse vectors. Separate multiple fields with commas (,). The fields must be keys defined in Metadata.
	//
	// example:
	//
	// title,abstract
	SparseRetrievalFields *string `json:"SparseRetrievalFields,omitempty" xml:"SparseRetrievalFields,omitempty"`
	// The sparse vector index configuration. If specified, a sparse vector index is created.
	SparseVectorIndexConfigShrink *string `json:"SparseVectorIndexConfig,omitempty" xml:"SparseVectorIndexConfig,omitempty"`
	// Specifies whether to support sparse vectors. Default value: false.
	//
	// example:
	//
	// true
	SupportSparse *bool `json:"SupportSparse,omitempty" xml:"SupportSparse,omitempty"`
	// The dense vector index configuration.
	VectorIndexConfigShrink *string `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty"`
}

func (s CreateDocumentCollectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionShrinkRequest) GetAlgorithm() *string {
	return s.Algorithm
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

func (s *CreateDocumentCollectionShrinkRequest) GetSparseRetrievalFields() *string {
	return s.SparseRetrievalFields
}

func (s *CreateDocumentCollectionShrinkRequest) GetSparseVectorIndexConfigShrink() *string {
	return s.SparseVectorIndexConfigShrink
}

func (s *CreateDocumentCollectionShrinkRequest) GetSupportSparse() *bool {
	return s.SupportSparse
}

func (s *CreateDocumentCollectionShrinkRequest) GetVectorIndexConfigShrink() *string {
	return s.VectorIndexConfigShrink
}

func (s *CreateDocumentCollectionShrinkRequest) SetAlgorithm(v string) *CreateDocumentCollectionShrinkRequest {
	s.Algorithm = &v
	return s
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

func (s *CreateDocumentCollectionShrinkRequest) SetSparseRetrievalFields(v string) *CreateDocumentCollectionShrinkRequest {
	s.SparseRetrievalFields = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetSparseVectorIndexConfigShrink(v string) *CreateDocumentCollectionShrinkRequest {
	s.SparseVectorIndexConfigShrink = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetSupportSparse(v bool) *CreateDocumentCollectionShrinkRequest {
	s.SupportSparse = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) SetVectorIndexConfigShrink(v string) *CreateDocumentCollectionShrinkRequest {
	s.VectorIndexConfigShrink = &v
	return s
}

func (s *CreateDocumentCollectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
