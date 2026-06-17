// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateDocumentCollectionRequest
	GetAlgorithm() *string
	SetCollection(v string) *CreateDocumentCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateDocumentCollectionRequest
	GetDBInstanceId() *string
	SetDimension(v int32) *CreateDocumentCollectionRequest
	GetDimension() *int32
	SetEmbeddingModel(v string) *CreateDocumentCollectionRequest
	GetEmbeddingModel() *string
	SetEnableGraph(v bool) *CreateDocumentCollectionRequest
	GetEnableGraph() *bool
	SetEntityTypes(v []*string) *CreateDocumentCollectionRequest
	GetEntityTypes() []*string
	SetExternalStorage(v int32) *CreateDocumentCollectionRequest
	GetExternalStorage() *int32
	SetFullTextRetrievalFields(v string) *CreateDocumentCollectionRequest
	GetFullTextRetrievalFields() *string
	SetHnswEfConstruction(v string) *CreateDocumentCollectionRequest
	GetHnswEfConstruction() *string
	SetHnswM(v int32) *CreateDocumentCollectionRequest
	GetHnswM() *int32
	SetLLMModel(v string) *CreateDocumentCollectionRequest
	GetLLMModel() *string
	SetLanguage(v string) *CreateDocumentCollectionRequest
	GetLanguage() *string
	SetManagerAccount(v string) *CreateDocumentCollectionRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateDocumentCollectionRequest
	GetManagerAccountPassword() *string
	SetMetadata(v string) *CreateDocumentCollectionRequest
	GetMetadata() *string
	SetMetadataIndices(v string) *CreateDocumentCollectionRequest
	GetMetadataIndices() *string
	SetMetrics(v string) *CreateDocumentCollectionRequest
	GetMetrics() *string
	SetNamespace(v string) *CreateDocumentCollectionRequest
	GetNamespace() *string
	SetOwnerId(v int64) *CreateDocumentCollectionRequest
	GetOwnerId() *int64
	SetParser(v string) *CreateDocumentCollectionRequest
	GetParser() *string
	SetPqEnable(v int32) *CreateDocumentCollectionRequest
	GetPqEnable() *int32
	SetRegionId(v string) *CreateDocumentCollectionRequest
	GetRegionId() *string
	SetRelationshipTypes(v []*string) *CreateDocumentCollectionRequest
	GetRelationshipTypes() []*string
	SetSparseRetrievalFields(v string) *CreateDocumentCollectionRequest
	GetSparseRetrievalFields() *string
	SetSparseVectorIndexConfig(v *CreateDocumentCollectionRequestSparseVectorIndexConfig) *CreateDocumentCollectionRequest
	GetSparseVectorIndexConfig() *CreateDocumentCollectionRequestSparseVectorIndexConfig
	SetSupportSparse(v bool) *CreateDocumentCollectionRequest
	GetSupportSparse() *bool
	SetVectorIndexConfig(v *CreateDocumentCollectionRequestVectorIndexConfig) *CreateDocumentCollectionRequest
	GetVectorIndexConfig() *CreateDocumentCollectionRequestVectorIndexConfig
}

type CreateDocumentCollectionRequest struct {
	// The vector index algorithm.
	//
	// Valid values:
	//
	// - `hnswflat`: An HNSW index without quantization compression. This is the default value.
	//
	// - `novam`: A graph index without quantization compression. This algorithm is suitable for high-performance scenarios such as real-time recommendation.
	//
	// - `novad`: A partitioned index with rabitq quantization. This algorithm is suitable for large-scale, low-cost retrieval scenarios.
	//
	// example:
	//
	// hnswflat
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The name of the document collection to create.
	//
	// > The name must comply with PostgreSQL object naming conventions.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in the target region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The vector dimension. If you omit this parameter, the system uses a default dimension for the selected `EmbeddingModel`.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The embedding model. The default value is `text-embedding-v3`.
	//
	// > Supported models:
	//
	// >
	//
	// > - `text-embedding-v3` (Recommended, Default): 1,024, 768, or 512 dimensions
	//
	// >
	//
	// > - `multimodal-embedding-v1` (Recommended): 1,024 dimensions, a multimodal embedding model
	//
	// >
	//
	// > - `text-embedding-v1`: 1,536 dimensions
	//
	// >
	//
	// > - `text-embedding-v2`: 1,536 dimensions
	//
	// >
	//
	// > - `text2vec` (Not recommended): 1,024 dimensions
	//
	// >
	//
	// > - `m3e-base` (Not recommended): 768 dimensions
	//
	// >
	//
	// > - `m3e-small` (Not recommended): 512 dimensions
	//
	// >
	//
	// > - `clip-vit-b-32` (Not recommended): CLIP ViT-B/32 model, 512 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-vit-b-16` (Not recommended): CLIP ViT-B/16 model, 512 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-vit-l-14` (Not recommended): CLIP ViT-L/14 model, 768 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-vit-l-14-336px` (Not recommended): CLIP ViT-L/14\\@336px model, 768 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-rn50` (Not recommended): CLIP RN50 model, 1,024 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-rn101` (Not recommended): CLIP RN101 model, 512 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-rn50x4` (Not recommended): CLIP RN50x4 model, 640 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-rn50x16` (Not recommended): CLIP RN50x16 model, 768 dimensions, an image embedding model
	//
	// >
	//
	// > - `clip-rn50x64` (Not recommended): CLIP RN50x64 model, 1,024 dimensions, an image embedding model
	//
	// example:
	//
	// text-embedding-v1
	EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
	// Specifies whether to build a knowledge graph. The default value is `false`.
	//
	// > To use this parameter, you must first upgrade your instance to a version that supports the graph engine. During the public preview period, submit a ticket to request an upgrade.
	//
	// example:
	//
	// true
	EnableGraph *bool `json:"EnableGraph,omitempty" xml:"EnableGraph,omitempty"`
	// A list of entity types.
	//
	// > This parameter is required when `EnableGraph` is set to `true`.
	//
	// example:
	//
	// Location
	EntityTypes []*string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty" type:"Repeated"`
	// Specifies whether to use memory-mapped files (mmap) to build the HNSW index. The default value is 0. Setting this to `1` is recommended if you do not need to delete data and require high upload performance.
	//
	// Valid values:
	//
	// - `0`: Builds the index by using segmented page storage. This mode supports delete and update operations and can use the `shared_buffer` in PostgreSQL for caching. This is the default value.
	//
	// - `1`: Builds the index by using mmap. This mode does not support delete or update operations.
	//
	// 	Notice:
	//
	// The `ExternalStorage` parameter is supported only by AnalyticDB for PostgreSQL V6.0 instances. It is not supported by V7.0 instances.
	//
	// example:
	//
	// 0
	ExternalStorage *int32 `json:"ExternalStorage,omitempty" xml:"ExternalStorage,omitempty"`
	// The metadata fields to use for full-text search. These fields must be keys defined in `Metadata`. Separate multiple fields with a comma (,).
	//
	// example:
	//
	// title,page
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	// The size of the candidate set (`ef_construction`) for HNSW index construction. The value must be greater than or equal to `2 	- HnswM`.
	//
	// > Value range:
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V6.0 instances: 40 to 4,000.
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V7.0 instances: 4 to 1,000. The default value is 64.
	//
	// example:
	//
	// 128
	HnswEfConstruction *string `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors (M) for the HNSW algorithm. You do not typically need to set this parameter, as the system automatically sets it based on the vector dimension.
	//
	// > Value range:
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V6.0 instances: 1 to 1,000.
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V7.0 instances: 2 to 100. The default value is 16.
	//
	// > We recommend that you set this parameter based on the vector dimension:
	//
	// >
	//
	// > - If the dimension is 384 or less: 16
	//
	// >
	//
	// > - If the dimension is greater than 384 and less than or equal to 768: 32
	//
	// >
	//
	// > - If the dimension is greater than 768 and less than or equal to 1,024: 64
	//
	// >
	//
	// > - If the dimension is greater than 1,024: 128
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	// The name of the LLM model. Valid values:
	//
	// - `knowledge-extract-standard`: The default value.
	//
	// - `knowledge-extract-mini`
	//
	// > This parameter takes effect only when `EnableGraph` is set to `true`.
	//
	// example:
	//
	// knowledge-extract-standard
	LLMModel *string `json:"LLMModel,omitempty" xml:"LLMModel,omitempty"`
	// The language used to build the knowledge graph. Valid values:
	//
	// - `Simplified Chinese`: The default value.
	//
	// - `English`
	//
	// > This parameter takes effect only when `EnableGraph` is set to `true`.
	//
	// example:
	//
	// Simplified Chinese
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The name of the manager account that has `rds_superuser` permissions.
	//
	// > You can create an account in the console on the \\*\\*Account Management\\*\\	- page or by calling the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	ManagerAccount *string `json:"ManagerAccount,omitempty" xml:"ManagerAccount,omitempty"`
	// The password for the manager account.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	ManagerAccountPassword *string `json:"ManagerAccountPassword,omitempty" xml:"ManagerAccountPassword,omitempty"`
	// The metadata schema for the vector data, specified as a JSON map where keys are field names and values are data types.
	//
	// > Supported data types
	//
	// >
	//
	// > - For a list of supported data types, see [Data types](https://help.aliyun.com/document_detail/424383.html).
	//
	// >
	//
	// > - The `money` data type is not supported.
	//
	// 	Warning:
	//
	// The following fields are reserved and cannot be used: `id`, `vector`, `doc_name`, `content`, `loader_metadata`, `source`, and `to_tsvector`.
	//
	// example:
	//
	// {"title":"text","page":"int"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The metadata fields on which to create scalar indexes. These fields must be keys defined in `Metadata`. Separate multiple fields with a comma (,).
	//
	// example:
	//
	// title
	MetadataIndices *string `json:"MetadataIndices,omitempty" xml:"MetadataIndices,omitempty"`
	// The distance metric for the vector index.
	//
	// Valid values:
	//
	// - **`l2`**: Euclidean distance.
	//
	// - **`ip`**: dot product (inner product) distance.
	//
	// - **`cosine`*	- (Default): cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace. The default value is `public`.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to list namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The tokenizer for full-text search. The default value is `zh_cn`.
	//
	// example:
	//
	// zh_cn
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// Specifies whether to enable the PQ (product quantization) algorithm to accelerate indexing. This is recommended for datasets with over 500,000 entries. Valid values:
	//
	// - `0`: Disables the feature.
	//
	// - `1`: Enables the feature. This is the default value.
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
	// A list of relationship types.
	//
	// > This parameter is required when `EnableGraph` is set to `true`.
	//
	// example:
	//
	// Occurred
	RelationshipTypes []*string `json:"RelationshipTypes,omitempty" xml:"RelationshipTypes,omitempty" type:"Repeated"`
	// The metadata fields used to build the sparse vector. These fields must be keys defined in `Metadata`. Separate multiple fields with a comma (,).
	//
	// example:
	//
	// title,abstract
	SparseRetrievalFields *string `json:"SparseRetrievalFields,omitempty" xml:"SparseRetrievalFields,omitempty"`
	// Configuration for the sparse vector index. Specifying this parameter creates the index.
	SparseVectorIndexConfig *CreateDocumentCollectionRequestSparseVectorIndexConfig `json:"SparseVectorIndexConfig,omitempty" xml:"SparseVectorIndexConfig,omitempty" type:"Struct"`
	// Specifies whether to support sparse vectors. The default value is `false`.
	//
	// example:
	//
	// true
	SupportSparse *bool `json:"SupportSparse,omitempty" xml:"SupportSparse,omitempty"`
	// Configuration for the dense vector index.
	VectorIndexConfig *CreateDocumentCollectionRequestVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
}

func (s CreateDocumentCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionRequest) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateDocumentCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateDocumentCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDocumentCollectionRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *CreateDocumentCollectionRequest) GetEmbeddingModel() *string {
	return s.EmbeddingModel
}

func (s *CreateDocumentCollectionRequest) GetEnableGraph() *bool {
	return s.EnableGraph
}

func (s *CreateDocumentCollectionRequest) GetEntityTypes() []*string {
	return s.EntityTypes
}

func (s *CreateDocumentCollectionRequest) GetExternalStorage() *int32 {
	return s.ExternalStorage
}

func (s *CreateDocumentCollectionRequest) GetFullTextRetrievalFields() *string {
	return s.FullTextRetrievalFields
}

func (s *CreateDocumentCollectionRequest) GetHnswEfConstruction() *string {
	return s.HnswEfConstruction
}

func (s *CreateDocumentCollectionRequest) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateDocumentCollectionRequest) GetLLMModel() *string {
	return s.LLMModel
}

func (s *CreateDocumentCollectionRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateDocumentCollectionRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateDocumentCollectionRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateDocumentCollectionRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateDocumentCollectionRequest) GetMetadataIndices() *string {
	return s.MetadataIndices
}

func (s *CreateDocumentCollectionRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *CreateDocumentCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateDocumentCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDocumentCollectionRequest) GetParser() *string {
	return s.Parser
}

func (s *CreateDocumentCollectionRequest) GetPqEnable() *int32 {
	return s.PqEnable
}

func (s *CreateDocumentCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDocumentCollectionRequest) GetRelationshipTypes() []*string {
	return s.RelationshipTypes
}

func (s *CreateDocumentCollectionRequest) GetSparseRetrievalFields() *string {
	return s.SparseRetrievalFields
}

func (s *CreateDocumentCollectionRequest) GetSparseVectorIndexConfig() *CreateDocumentCollectionRequestSparseVectorIndexConfig {
	return s.SparseVectorIndexConfig
}

func (s *CreateDocumentCollectionRequest) GetSupportSparse() *bool {
	return s.SupportSparse
}

func (s *CreateDocumentCollectionRequest) GetVectorIndexConfig() *CreateDocumentCollectionRequestVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *CreateDocumentCollectionRequest) SetAlgorithm(v string) *CreateDocumentCollectionRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetCollection(v string) *CreateDocumentCollectionRequest {
	s.Collection = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetDBInstanceId(v string) *CreateDocumentCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetDimension(v int32) *CreateDocumentCollectionRequest {
	s.Dimension = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetEmbeddingModel(v string) *CreateDocumentCollectionRequest {
	s.EmbeddingModel = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetEnableGraph(v bool) *CreateDocumentCollectionRequest {
	s.EnableGraph = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetEntityTypes(v []*string) *CreateDocumentCollectionRequest {
	s.EntityTypes = v
	return s
}

func (s *CreateDocumentCollectionRequest) SetExternalStorage(v int32) *CreateDocumentCollectionRequest {
	s.ExternalStorage = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetFullTextRetrievalFields(v string) *CreateDocumentCollectionRequest {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetHnswEfConstruction(v string) *CreateDocumentCollectionRequest {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetHnswM(v int32) *CreateDocumentCollectionRequest {
	s.HnswM = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetLLMModel(v string) *CreateDocumentCollectionRequest {
	s.LLMModel = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetLanguage(v string) *CreateDocumentCollectionRequest {
	s.Language = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetManagerAccount(v string) *CreateDocumentCollectionRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetManagerAccountPassword(v string) *CreateDocumentCollectionRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetMetadata(v string) *CreateDocumentCollectionRequest {
	s.Metadata = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetMetadataIndices(v string) *CreateDocumentCollectionRequest {
	s.MetadataIndices = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetMetrics(v string) *CreateDocumentCollectionRequest {
	s.Metrics = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetNamespace(v string) *CreateDocumentCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetOwnerId(v int64) *CreateDocumentCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetParser(v string) *CreateDocumentCollectionRequest {
	s.Parser = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetPqEnable(v int32) *CreateDocumentCollectionRequest {
	s.PqEnable = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetRegionId(v string) *CreateDocumentCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetRelationshipTypes(v []*string) *CreateDocumentCollectionRequest {
	s.RelationshipTypes = v
	return s
}

func (s *CreateDocumentCollectionRequest) SetSparseRetrievalFields(v string) *CreateDocumentCollectionRequest {
	s.SparseRetrievalFields = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetSparseVectorIndexConfig(v *CreateDocumentCollectionRequestSparseVectorIndexConfig) *CreateDocumentCollectionRequest {
	s.SparseVectorIndexConfig = v
	return s
}

func (s *CreateDocumentCollectionRequest) SetSupportSparse(v bool) *CreateDocumentCollectionRequest {
	s.SupportSparse = &v
	return s
}

func (s *CreateDocumentCollectionRequest) SetVectorIndexConfig(v *CreateDocumentCollectionRequestVectorIndexConfig) *CreateDocumentCollectionRequest {
	s.VectorIndexConfig = v
	return s
}

func (s *CreateDocumentCollectionRequest) Validate() error {
	if s.SparseVectorIndexConfig != nil {
		if err := s.SparseVectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VectorIndexConfig != nil {
		if err := s.VectorIndexConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDocumentCollectionRequestSparseVectorIndexConfig struct {
	// The vector index algorithm.
	//
	// Valid values:
	//
	// - `hnswflat`: An HNSW index without quantization compression. This is the default value.
	//
	// - `novam`: A graph index without quantization compression. This algorithm is suitable for high-performance scenarios such as real-time recommendation.
	//
	// example:
	//
	// hnswflat
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The size of the candidate set (`ef_construction`) for HNSW index construction. Valid values: 4 to 1,000. The default value is 64.
	//
	// > This parameter is applicable only to AnalyticDB for PostgreSQL V7.0 instances, and its value must be greater than or equal to `2 	- HnswM`.
	//
	// example:
	//
	// 128
	HnswEfConstruction *int32 `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors (M) for the HNSW algorithm. You do not typically need to set this parameter, as the system automatically sets it based on the vector dimension.
	//
	// > Value range:
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V6.0 instances: 1 to 1,000.
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V7.0 instances: 2 to 100. The default value is 16.
	//
	// > We recommend that you set this parameter based on the vector dimension:
	//
	// >
	//
	// > - If the dimension is 384 or less: 16
	//
	// >
	//
	// > - If the dimension is greater than 384 and less than or equal to 768: 32
	//
	// >
	//
	// > - If the dimension is greater than 768 and less than or equal to 1,024: 64
	//
	// >
	//
	// > - If the dimension is greater than 1,024: 128
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
}

func (s CreateDocumentCollectionRequestSparseVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionRequestSparseVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) GetHnswEfConstruction() *int32 {
	return s.HnswEfConstruction
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) SetAlgorithm(v string) *CreateDocumentCollectionRequestSparseVectorIndexConfig {
	s.Algorithm = &v
	return s
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) SetHnswEfConstruction(v int32) *CreateDocumentCollectionRequestSparseVectorIndexConfig {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) SetHnswM(v int32) *CreateDocumentCollectionRequestSparseVectorIndexConfig {
	s.HnswM = &v
	return s
}

func (s *CreateDocumentCollectionRequestSparseVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDocumentCollectionRequestVectorIndexConfig struct {
	// The number of lists (partitions) for the novad algorithm. Valid values: [2, 1073741824]. The default value is 256.
	//
	// example:
	//
	// 256
	Nlist *int32 `json:"Nlist,omitempty" xml:"Nlist,omitempty"`
	// The number of bits used for rabitq compression. Valid values: [1, 8]. The default value is 3.
	//
	// example:
	//
	// 3
	RabitqBits *int32 `json:"RabitqBits,omitempty" xml:"RabitqBits,omitempty"`
}

func (s CreateDocumentCollectionRequestVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentCollectionRequestVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *CreateDocumentCollectionRequestVectorIndexConfig) GetNlist() *int32 {
	return s.Nlist
}

func (s *CreateDocumentCollectionRequestVectorIndexConfig) GetRabitqBits() *int32 {
	return s.RabitqBits
}

func (s *CreateDocumentCollectionRequestVectorIndexConfig) SetNlist(v int32) *CreateDocumentCollectionRequestVectorIndexConfig {
	s.Nlist = &v
	return s
}

func (s *CreateDocumentCollectionRequestVectorIndexConfig) SetRabitqBits(v int32) *CreateDocumentCollectionRequestVectorIndexConfig {
	s.RabitqBits = &v
	return s
}

func (s *CreateDocumentCollectionRequestVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
