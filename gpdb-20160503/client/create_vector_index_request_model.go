// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVectorIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateVectorIndexRequest
	GetAlgorithm() *string
	SetCollection(v string) *CreateVectorIndexRequest
	GetCollection() *string
	SetDBInstanceId(v string) *CreateVectorIndexRequest
	GetDBInstanceId() *string
	SetDimension(v int32) *CreateVectorIndexRequest
	GetDimension() *int32
	SetExternalStorage(v int32) *CreateVectorIndexRequest
	GetExternalStorage() *int32
	SetHnswEfConstruction(v int32) *CreateVectorIndexRequest
	GetHnswEfConstruction() *int32
	SetHnswM(v int32) *CreateVectorIndexRequest
	GetHnswM() *int32
	SetManagerAccount(v string) *CreateVectorIndexRequest
	GetManagerAccount() *string
	SetManagerAccountPassword(v string) *CreateVectorIndexRequest
	GetManagerAccountPassword() *string
	SetMetrics(v string) *CreateVectorIndexRequest
	GetMetrics() *string
	SetNamespace(v string) *CreateVectorIndexRequest
	GetNamespace() *string
	SetNlist(v int32) *CreateVectorIndexRequest
	GetNlist() *int32
	SetOwnerId(v int64) *CreateVectorIndexRequest
	GetOwnerId() *int64
	SetPqEnable(v int32) *CreateVectorIndexRequest
	GetPqEnable() *int32
	SetRabitqBits(v int32) *CreateVectorIndexRequest
	GetRabitqBits() *int32
	SetRegionId(v string) *CreateVectorIndexRequest
	GetRegionId() *string
	SetType(v string) *CreateVectorIndexRequest
	GetType() *string
}

type CreateVectorIndexRequest struct {
	// The vector index algorithm. Valid values:
	//
	// - hnswflat: HNSW index without quantization compression (default).
	//
	// - novam: graph index without quantization compression, suitable for high-performance scenarios such as real-time recommendation.
	//
	// - novad: partitioned index with RaBitQ quantization, suitable for large-scale, low-cost retrieval scenarios.
	//
	// example:
	//
	// hnswflat
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to query the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The vector dimensions.
	//
	// > - This parameter is required for dense vectors.
	//
	// > - The value must be the same as the length of the vector data Rows.Vector in the [UpsertCollectionData](https://help.aliyun.com/document_detail/2401493.html) operation.
	//
	// example:
	//
	// 1024
	Dimension *int32 `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// Specifies whether to use mmap to build the HNSW index. Default value: 0. If data does not need to be deleted and you have high requirements for data upload speed, we recommend that you set this parameter to 1.
	//
	// Valid values:
	//
	// - 0: uses the default segment-page storage to build the index. In this mode, shared_buffer in PostgreSQL can be used for caching, and operations such as deletion and update are supported.
	//
	// - 1: uses mmap to build the index. This mode does not support operations such as deletion and update.
	//
	// 	Notice: The ExternalStorage parameter is supported only in V6.0. It is not supported in V7.0.</notice>
	//
	// example:
	//
	// 0
	ExternalStorage *int32 `json:"ExternalStorage,omitempty" xml:"ExternalStorage,omitempty"`
	// The size of the candidate set when the HNSW algorithm builds the index. Valid values: 4 to 1000. Default value: 64.
	//
	// > This parameter is required only for AnalyticDB for PostgreSQL V7.0 instances, and the value must be greater than or equal to 2 × HNSW_M.
	//
	// example:
	//
	// 128
	HnswEfConstruction *int32 `json:"HnswEfConstruction,omitempty" xml:"HnswEfConstruction,omitempty"`
	// The maximum number of neighbors in the HNSW algorithm. The operation automatically sets this value based on the vector dimensions. You do not need to manually set this value in most cases.
	//
	// > Valid values:
	//
	// >- AnalyticDB for PostgreSQL V6.0 instances: 1 to 1000.
	//
	// >- AnalyticDB for PostgreSQL V7.0 instances: 2 to 100. Default value: 16.
	//
	// > Recommended settings based on vector dimensions:
	//
	// > - 384 or less: 16
	//
	// > - Greater than 384 and less than or equal to 768: 32
	//
	// > - Greater than 768 and less than or equal to 1024: 64
	//
	// > - Greater than 1024: 128
	//
	// example:
	//
	// 64
	HnswM *int32 `json:"HnswM,omitempty" xml:"HnswM,omitempty"`
	// The name of the management account that has the rds_superuser permission.
	//
	// > You can create an account in the console by navigating to Account Management, or call the [CreateAccount](https://help.aliyun.com/document_detail/2361789.html) operation.
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
	// The distance metric used to build the vector index. Valid values:
	//
	// - l2: Euclidean distance.
	//
	// - ip: inner product.
	//
	// - cosine: cosine similarity.
	//
	// > Only ip is supported for sparse vectors.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query the list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of lists (partitions) for Novad. Valid values: 2 to 1073741824. Default value: 256.
	//
	// example:
	//
	// 256
	Nlist   *int32 `json:"Nlist,omitempty" xml:"Nlist,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable product quantization (PQ) algorithm acceleration for the index. We recommend that you enable this feature when the data volume exceeds 500,000. Valid values:
	//
	// - 0: disabled.
	//
	// - 1: enabled (default).
	//
	// example:
	//
	// 1
	PqEnable *int32 `json:"PqEnable,omitempty" xml:"PqEnable,omitempty"`
	// The number of bits for RaBitQ compression. Valid values: 1 to 8. Default value: 3.
	//
	// example:
	//
	// 3
	RabitqBits *int32 `json:"RabitqBits,omitempty" xml:"RabitqBits,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vector type. Valid values:
	//
	// - Dense: dense vector (default).
	//
	// - Sparse: sparse vector.
	//
	// example:
	//
	// Dense
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateVectorIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVectorIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateVectorIndexRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateVectorIndexRequest) GetCollection() *string {
	return s.Collection
}

func (s *CreateVectorIndexRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateVectorIndexRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *CreateVectorIndexRequest) GetExternalStorage() *int32 {
	return s.ExternalStorage
}

func (s *CreateVectorIndexRequest) GetHnswEfConstruction() *int32 {
	return s.HnswEfConstruction
}

func (s *CreateVectorIndexRequest) GetHnswM() *int32 {
	return s.HnswM
}

func (s *CreateVectorIndexRequest) GetManagerAccount() *string {
	return s.ManagerAccount
}

func (s *CreateVectorIndexRequest) GetManagerAccountPassword() *string {
	return s.ManagerAccountPassword
}

func (s *CreateVectorIndexRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *CreateVectorIndexRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateVectorIndexRequest) GetNlist() *int32 {
	return s.Nlist
}

func (s *CreateVectorIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVectorIndexRequest) GetPqEnable() *int32 {
	return s.PqEnable
}

func (s *CreateVectorIndexRequest) GetRabitqBits() *int32 {
	return s.RabitqBits
}

func (s *CreateVectorIndexRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVectorIndexRequest) GetType() *string {
	return s.Type
}

func (s *CreateVectorIndexRequest) SetAlgorithm(v string) *CreateVectorIndexRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateVectorIndexRequest) SetCollection(v string) *CreateVectorIndexRequest {
	s.Collection = &v
	return s
}

func (s *CreateVectorIndexRequest) SetDBInstanceId(v string) *CreateVectorIndexRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateVectorIndexRequest) SetDimension(v int32) *CreateVectorIndexRequest {
	s.Dimension = &v
	return s
}

func (s *CreateVectorIndexRequest) SetExternalStorage(v int32) *CreateVectorIndexRequest {
	s.ExternalStorage = &v
	return s
}

func (s *CreateVectorIndexRequest) SetHnswEfConstruction(v int32) *CreateVectorIndexRequest {
	s.HnswEfConstruction = &v
	return s
}

func (s *CreateVectorIndexRequest) SetHnswM(v int32) *CreateVectorIndexRequest {
	s.HnswM = &v
	return s
}

func (s *CreateVectorIndexRequest) SetManagerAccount(v string) *CreateVectorIndexRequest {
	s.ManagerAccount = &v
	return s
}

func (s *CreateVectorIndexRequest) SetManagerAccountPassword(v string) *CreateVectorIndexRequest {
	s.ManagerAccountPassword = &v
	return s
}

func (s *CreateVectorIndexRequest) SetMetrics(v string) *CreateVectorIndexRequest {
	s.Metrics = &v
	return s
}

func (s *CreateVectorIndexRequest) SetNamespace(v string) *CreateVectorIndexRequest {
	s.Namespace = &v
	return s
}

func (s *CreateVectorIndexRequest) SetNlist(v int32) *CreateVectorIndexRequest {
	s.Nlist = &v
	return s
}

func (s *CreateVectorIndexRequest) SetOwnerId(v int64) *CreateVectorIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVectorIndexRequest) SetPqEnable(v int32) *CreateVectorIndexRequest {
	s.PqEnable = &v
	return s
}

func (s *CreateVectorIndexRequest) SetRabitqBits(v int32) *CreateVectorIndexRequest {
	s.RabitqBits = &v
	return s
}

func (s *CreateVectorIndexRequest) SetRegionId(v string) *CreateVectorIndexRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVectorIndexRequest) SetType(v string) *CreateVectorIndexRequest {
	s.Type = &v
	return s
}

func (s *CreateVectorIndexRequest) Validate() error {
	return dara.Validate(s)
}
