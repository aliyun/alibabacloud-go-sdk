// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateCollectionRequest
	GetAlgorithm() *string
	SetCollection(v string) *CreateCollectionRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateCollectionRequest
	GetDBInstanceId() *string
	SetDimension(v int64) *CreateCollectionRequest
	GetDimension() *int64
	SetExternalStorage(v int32) *CreateCollectionRequest
	GetExternalStorage() *int32
	SetFullTextRetrievalFields(v string) *CreateCollectionRequest
	GetFullTextRetrievalFields() *string
	SetHnswEfConstruction(v string) *CreateCollectionRequest
	GetHnswEfConstruction() *string
	SetHnswM(v int32) *CreateCollectionRequest
	GetHnswM() *int32
	SetManagerAccount(v string) *CreateCollectionRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateCollectionRequest
	GetManagerAccountPassword() *string
	SetMetadata(v string) *CreateCollectionRequest
	GetMetadata() *string
	SetMetadataIndices(v string) *CreateCollectionRequest
	GetMetadataIndices() *string
	SetMetrics(v string) *CreateCollectionRequest
	GetMetrics() *string
	SetNamespace(v string) *CreateCollectionRequest
	GetNamespace() *string
	SetOwnerId(v int64) *CreateCollectionRequest
	GetOwnerId() *int64
	SetParser(v string) *CreateCollectionRequest
	GetParser() *string
	SetPqEnable(v int32) *CreateCollectionRequest
	GetPqEnable() *int32
	SetRegionId(v string) *CreateCollectionRequest
	GetRegionId() *string
	SetSparseVectorIndexConfig(v *CreateCollectionRequestSparseVectorIndexConfig) *CreateCollectionRequest
	GetSparseVectorIndexConfig() *CreateCollectionRequestSparseVectorIndexConfig
	SetSupportSparse(v bool) *CreateCollectionRequest
	GetSupportSparse() *bool
	SetVectorIndexConfig(v *CreateCollectionRequestVectorIndexConfig) *CreateCollectionRequest
	GetVectorIndexConfig() *CreateCollectionRequestVectorIndexConfig
	SetWorkspaceId(v string) *CreateCollectionRequest
	GetWorkspaceId() *string
}

type CreateCollectionRequest struct {
	// The vector index algorithm.
	//
	// Valid values:
	//
	// - `hnswflat`: (Default) An HNSW index without quantization compression.
	//
	// - `novam`: A graph index without quantization compression. This algorithm is suitable for high-performance scenarios, such as real-time recommendations.
	//
	// - `novad`: A partitioned index with `rabitq` quantization. This algorithm is suitable for large-scale, low-cost retrieval scenarios.
	//
	// example:
	//
	// hnswflat
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The name of the collection to create.
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
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// example:
	//
	// gp-bp152460513z****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The vector dimension.
	//
	// > If you specify this parameter, a vector index is created. In subsequent calls to the [UpsertCollectionData](https://help.aliyun.com/document_detail/2401493.html) operation, the length of `Rows.Vector` must match this dimension. If you do not specify this parameter, you must call the [CreateVectorIndex](https://help.aliyun.com/document_detail/2401499.html) operation to create an index later.
	//
	// example:
	//
	// 1024
	Dimension *int64 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// Specifies whether to use `mmap` to build the HNSW index. The default value is 0. We recommend setting this to 1 if your data does not require deletion and you need high-performance data ingestion.
	//
	// Valid values:
	//
	// - `0`: (Default) Builds the index by using segmented page storage. This mode can use the `shared_buffer` in PostgreSQL for caching and supports `DELETE` and `UPDATE` operations.
	//
	// - `1`: Builds the index by using `mmap`. This mode does not support `DELETE` or `UPDATE` operations.
	//
	// 	Notice:
	//
	// The `ExternalStorage` parameter is available only for AnalyticDB for PostgreSQL v6.0 instances and is not supported in v7.0.
	//
	// example:
	//
	// 0
	ExternalStorage *int32 `json:"ExternalStorage,omitempty" xml:"ExternalStorage,omitempty"`
	// The fields to use for full-text search. Use commas (`,`) to separate multiple field names. These fields must be keys defined in the `Metadata` parameter.
	//
	// example:
	//
	// title,content
	FullTextRetrievalFields *string `json:"FullTextRetrievalFields,omitempty" xml:"FullTextRetrievalFields,omitempty"`
	// The size of the candidate set for HNSW index construction. The value must be greater than or equal to `2 	- HnswM`.
	//
	// > Value range:
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V6.0 instances: 40 to 4000.
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V7.0 instances: 4 to 1000. The default value is 64.
	//
	// example:
	//
	// 128
	HnswEfConstruction *string `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors for the HNSW algorithm. You do not typically need to set this parameter, as the system automatically determines a value based on the vector dimension.
	//
	// > Value range:
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V6.0 instances: 1 to 1000.
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V7.0 instances: 2 to 100. The default value is 16.
	//
	// > We recommend that you set this parameter based on the vector dimension:
	//
	// >
	//
	// > - 16 for dimensions less than or equal to 384.
	//
	// >
	//
	// > - 32 for dimensions greater than 384 and less than or equal to 768.
	//
	// >
	//
	// > - 64 for dimensions greater than 768 and less than or equal to 1024.
	//
	// >
	//
	// > - 128 for dimensions greater than 1024.
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	// The name of the management account that has the `rds_superuser` privilege.
	//
	// > You can call the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation to create an account.
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
	// A JSON string that defines the metadata schema as a map. The keys are field names, and the values are their corresponding data types.
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
	// The field names `id`, `vector`, `to_tsvector`, and `source` are reserved and cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"title":"text","content":"text","response":"int"}
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The scalar index fields. Separate multiple fields with commas (`,`). The fields must be keys that are defined in `Metadata`.
	//
	// example:
	//
	// title
	MetadataIndices *string `json:"MetadataIndices,omitempty" xml:"MetadataIndices,omitempty"`
	// The distance metric used to build the vector index. Valid values:
	//
	// - `l2`: Euclidean distance.
	//
	// - `ip`: dot product.
	//
	// - `cosine`: cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace or the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to list existing namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parser for full-text search. The default is `zh_cn`.
	//
	// example:
	//
	// zh_cn
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// Specifies whether to enable Product Quantization (PQ) for index acceleration. This is recommended for datasets with more than 500,000 entries. Valid values:
	//
	// - `0`: Disabled.
	//
	// - `1`: (Default) Enabled.
	//
	// example:
	//
	// 1
	PqEnable *int32 `json:"PqEnable,omitempty" xml:"PqEnable,omitempty"`
	// The ID of the region where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration for the sparse vector index. If specified, a sparse vector index is created.
	SparseVectorIndexConfig *CreateCollectionRequestSparseVectorIndexConfig `json:"SparseVectorIndexConfig,omitempty" xml:"SparseVectorIndexConfig,omitempty" type:"Struct"`
	// Specifies whether to enable support for sparse vectors. The default value is `false`.
	//
	// example:
	//
	// true
	SupportSparse *bool `json:"SupportSparse,omitempty" xml:"SupportSparse,omitempty"`
	// The configuration for the dense vector index.
	VectorIndexConfig *CreateCollectionRequestVectorIndexConfig `json:"VectorIndexConfig,omitempty" xml:"VectorIndexConfig,omitempty" type:"Struct"`
	// The ID of the workspace, which contains multiple database instances. You must specify either `WorkspaceId` or `DBInstanceId`. If both are specified, `WorkspaceId` takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectionRequest) GoString() string {
	return s.String()
}

func (s *CreateCollectionRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateCollectionRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateCollectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateCollectionRequest) GetDimension() *int64 {
	return s.Dimension
}

func (s *CreateCollectionRequest) GetExternalStorage() *int32 {
	return s.ExternalStorage
}

func (s *CreateCollectionRequest) GetFullTextRetrievalFields() *string {
	return s.FullTextRetrievalFields
}

func (s *CreateCollectionRequest) GetHnswEfConstruction() *string {
	return s.HnswEfConstruction
}

func (s *CreateCollectionRequest) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateCollectionRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateCollectionRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateCollectionRequest) GetMetadata() *string {
	return s.Metadata
}

func (s *CreateCollectionRequest) GetMetadataIndices() *string {
	return s.MetadataIndices
}

func (s *CreateCollectionRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *CreateCollectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCollectionRequest) GetParser() *string {
	return s.Parser
}

func (s *CreateCollectionRequest) GetPqEnable() *int32 {
	return s.PqEnable
}

func (s *CreateCollectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCollectionRequest) GetSparseVectorIndexConfig() *CreateCollectionRequestSparseVectorIndexConfig {
	return s.SparseVectorIndexConfig
}

func (s *CreateCollectionRequest) GetSupportSparse() *bool {
	return s.SupportSparse
}

func (s *CreateCollectionRequest) GetVectorIndexConfig() *CreateCollectionRequestVectorIndexConfig {
	return s.VectorIndexConfig
}

func (s *CreateCollectionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCollectionRequest) SetAlgorithm(v string) *CreateCollectionRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateCollectionRequest) SetCollection(v string) *CreateCollectionRequest {
	s.Collection = &v
	return s
}

func (s *CreateCollectionRequest) SetDBInstanceId(v string) *CreateCollectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateCollectionRequest) SetDimension(v int64) *CreateCollectionRequest {
	s.Dimension = &v
	return s
}

func (s *CreateCollectionRequest) SetExternalStorage(v int32) *CreateCollectionRequest {
	s.ExternalStorage = &v
	return s
}

func (s *CreateCollectionRequest) SetFullTextRetrievalFields(v string) *CreateCollectionRequest {
	s.FullTextRetrievalFields = &v
	return s
}

func (s *CreateCollectionRequest) SetHnswEfConstruction(v string) *CreateCollectionRequest {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateCollectionRequest) SetHnswM(v int32) *CreateCollectionRequest {
	s.HnswM = &v
	return s
}

func (s *CreateCollectionRequest) SetManagerAccount(v string) *CreateCollectionRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateCollectionRequest) SetManagerAccountPassword(v string) *CreateCollectionRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateCollectionRequest) SetMetadata(v string) *CreateCollectionRequest {
	s.Metadata = &v
	return s
}

func (s *CreateCollectionRequest) SetMetadataIndices(v string) *CreateCollectionRequest {
	s.MetadataIndices = &v
	return s
}

func (s *CreateCollectionRequest) SetMetrics(v string) *CreateCollectionRequest {
	s.Metrics = &v
	return s
}

func (s *CreateCollectionRequest) SetNamespace(v string) *CreateCollectionRequest {
	s.Namespace = &v
	return s
}

func (s *CreateCollectionRequest) SetOwnerId(v int64) *CreateCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCollectionRequest) SetParser(v string) *CreateCollectionRequest {
	s.Parser = &v
	return s
}

func (s *CreateCollectionRequest) SetPqEnable(v int32) *CreateCollectionRequest {
	s.PqEnable = &v
	return s
}

func (s *CreateCollectionRequest) SetRegionId(v string) *CreateCollectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCollectionRequest) SetSparseVectorIndexConfig(v *CreateCollectionRequestSparseVectorIndexConfig) *CreateCollectionRequest {
	s.SparseVectorIndexConfig = v
	return s
}

func (s *CreateCollectionRequest) SetSupportSparse(v bool) *CreateCollectionRequest {
	s.SupportSparse = &v
	return s
}

func (s *CreateCollectionRequest) SetVectorIndexConfig(v *CreateCollectionRequestVectorIndexConfig) *CreateCollectionRequest {
	s.VectorIndexConfig = v
	return s
}

func (s *CreateCollectionRequest) SetWorkspaceId(v string) *CreateCollectionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCollectionRequest) Validate() error {
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

type CreateCollectionRequestSparseVectorIndexConfig struct {
	// The vector index algorithm.
	//
	// Valid values:
	//
	// - `hnswflat`: (Default) An HNSW index without quantization compression.
	//
	// - `novam`: A graph index without quantization compression. This algorithm is suitable for high-performance scenarios, such as real-time recommendations.
	//
	// example:
	//
	// hnswflat
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The size of the candidate set for HNSW index construction. The value must be an integer from 4 to 1,000. The default is 64.
	//
	// > This parameter is required only for AnalyticDB for PostgreSQL V7.0 instances, and its value must be greater than or equal to `2 	- HnswM`.
	//
	// example:
	//
	// 128
	HnswEfConstruction *int32 `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors for the HNSW algorithm. You do not typically need to set this parameter, as the system automatically determines a value based on the vector dimension.
	//
	// > Value range:
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V6.0 instances: 1 to 1000.
	//
	// >
	//
	// > - For AnalyticDB for PostgreSQL V7.0 instances: 2 to 100. The default value is 16.
	//
	// > We recommend that you set this parameter based on the vector dimension:
	//
	// >
	//
	// > - 16 for dimensions less than or equal to 384.
	//
	// >
	//
	// > - 32 for dimensions greater than 384 and less than or equal to 768.
	//
	// >
	//
	// > - 64 for dimensions greater than 768 and less than or equal to 1024.
	//
	// >
	//
	// > - 128 for dimensions greater than 1024.
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
}

func (s CreateCollectionRequestSparseVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectionRequestSparseVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) GetHnswEfConstruction() *int32 {
	return s.HnswEfConstruction
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) SetAlgorithm(v string) *CreateCollectionRequestSparseVectorIndexConfig {
	s.Algorithm = &v
	return s
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) SetHnswEfConstruction(v int32) *CreateCollectionRequestSparseVectorIndexConfig {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) SetHnswM(v int32) *CreateCollectionRequestSparseVectorIndexConfig {
	s.HnswM = &v
	return s
}

func (s *CreateCollectionRequestSparseVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCollectionRequestVectorIndexConfig struct {
	// The number of lists (partitions) for a `novad` index. The value must be an integer from 2 to 1,073,741,824. The default is 256.
	//
	// example:
	//
	// 256
	Nlist *int32 `json:"Nlist,omitempty" xml:"Nlist,omitempty"`
	// The number of bits for `rabitq` compression. The value must be an integer from 1 to 8. The default is 3.
	//
	// example:
	//
	// 3
	RabitqBits *int32 `json:"RabitqBits,omitempty" xml:"RabitqBits,omitempty"`
}

func (s CreateCollectionRequestVectorIndexConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCollectionRequestVectorIndexConfig) GoString() string {
	return s.String()
}

func (s *CreateCollectionRequestVectorIndexConfig) GetNlist() *int32 {
	return s.Nlist
}

func (s *CreateCollectionRequestVectorIndexConfig) GetRabitqBits() *int32 {
	return s.RabitqBits
}

func (s *CreateCollectionRequestVectorIndexConfig) SetNlist(v int32) *CreateCollectionRequestVectorIndexConfig {
	s.Nlist = &v
	return s
}

func (s *CreateCollectionRequestVectorIndexConfig) SetRabitqBits(v int32) *CreateCollectionRequestVectorIndexConfig {
	s.RabitqBits = &v
	return s
}

func (s *CreateCollectionRequestVectorIndexConfig) Validate() error {
	return dara.Validate(s)
}
