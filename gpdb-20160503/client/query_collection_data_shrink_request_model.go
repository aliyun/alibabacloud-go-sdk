// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCollectionDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *QueryCollectionDataShrinkRequest
	GetCollection() *string
	SetContent(v string) *QueryCollectionDataShrinkRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryCollectionDataShrinkRequest
	GetDBInstanceId() *string
	SetFilter(v string) *QueryCollectionDataShrinkRequest
	GetFilter() *string
	SetHybridSearch(v string) *QueryCollectionDataShrinkRequest
	GetHybridSearch() *string
	SetHybridSearchArgsShrink(v string) *QueryCollectionDataShrinkRequest
	GetHybridSearchArgsShrink() *string
	SetIncludeMetadataFields(v string) *QueryCollectionDataShrinkRequest
	GetIncludeMetadataFields() *string
	SetIncludeSparseValues(v bool) *QueryCollectionDataShrinkRequest
	GetIncludeSparseValues() *bool
	SetIncludeValues(v bool) *QueryCollectionDataShrinkRequest
	GetIncludeValues() *bool
	SetMetrics(v string) *QueryCollectionDataShrinkRequest
	GetMetrics() *string
	SetNamespace(v string) *QueryCollectionDataShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *QueryCollectionDataShrinkRequest
	GetNamespacePassword() *string
	SetOffset(v int32) *QueryCollectionDataShrinkRequest
	GetOffset() *int32
	SetOrderBy(v string) *QueryCollectionDataShrinkRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *QueryCollectionDataShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryCollectionDataShrinkRequest
	GetRegionId() *string
	SetRelationalTableFilterShrink(v string) *QueryCollectionDataShrinkRequest
	GetRelationalTableFilterShrink() *string
	SetSparseVectorShrink(v string) *QueryCollectionDataShrinkRequest
	GetSparseVectorShrink() *string
	SetTopK(v int64) *QueryCollectionDataShrinkRequest
	GetTopK() *int64
	SetVectorShrink(v string) *QueryCollectionDataShrinkRequest
	GetVectorShrink() *string
	SetWorkspaceId(v string) *QueryCollectionDataShrinkRequest
	GetWorkspaceId() *string
}

type QueryCollectionDataShrinkRequest struct {
	// The collection name.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to view the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The content used for full-text retrieval. If this parameter is left empty, only vector retrieval is used. If this parameter is specified, both vector and full-text retrieval are used.
	//
	// > This parameter and the Vector parameter cannot be left empty at the same time.
	//
	// example:
	//
	// hello_world
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including the instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The filter condition for the data to query, in the SQL WHERE clause format. It is an expression that returns a Boolean value (true or false). The condition can be a simple comparison operator, such as equal to (=), not equal to (<> or !=), greater than (>), less than (<), greater than or equal to (>=), or less than or equal to (<=). It can also be a more complex expression combined with logical operators (AND, OR, NOT), or a condition that uses keywords such as IN, BETWEEN, and LIKE.
	//
	// >
	//
	// > - For detailed syntax, refer to https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/
	//
	// example:
	//
	// pipeline_id=\\"1yhpmo0rbn\\" AND (spu=\\"10025667796135\\" AND dept_id=\\"226\\")
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The dual-channel recall algorithm. Default value: empty (which means the vector and full-text scores are directly compared and sorting is applied). Valid values:
	//
	// - RRF: reciprocal rank fusion. A parameter k is used to control the fusion effect. For more information, see the HybridSearchArgs configuration.
	//
	// - Weight: weighted sorting. A parameter alpha is used to control the weight of the vector and full-text scores before sorting. For more information, see the HybridSearchArgs configuration.
	//
	// - Cascaded: full-text retrieval is performed first, and then vector retrieval is performed based on the results.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The algorithm parameters for hybrid search. RRF and Weight are supported:
	//
	// - RRF: specifies the constant k in the score calculation formula `1/(k+rank_i)`. The value must be a positive integer greater than 1. Format:
	//
	// ```
	//
	// {
	//
	//    "RRF": {
	//
	//     "k": 60
	//
	//    }
	//
	// }
	//
	// ```
	//
	// - Weight: The calculation formula is `alpha 	- vector_score + (1-alpha) 	- text_score`. The parameter alpha represents the weight of the vector and full-text search scores. The value ranges from 0 to 1, where 0 indicates full-text search only and 1 indicates vector search only. Format:
	//
	// ```
	//
	// {
	//
	//    "Weight": {
	//
	//     "alpha": 0.5
	//
	//    }
	//
	// }
	//
	// ```
	HybridSearchArgsShrink *string `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// The metadata fields to return. Default value: empty. Separate multiple fields with commas.
	//
	// example:
	//
	// title,content
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
	// Specifies whether to return sparse vector data. Valid values:
	//
	// - **true**: returns sparse vector data.
	//
	// - **false**: does not return sparse vector data.
	//
	// example:
	//
	// false
	IncludeSparseValues *bool `json:"IncludeSparseValues,omitempty" xml:"IncludeSparseValues,omitempty"`
	// Specifies whether to return dense vector data. Valid values:
	//
	// - **true**: returns dense vector data.
	//
	// - **false**: does not return dense vector data.
	//
	// example:
	//
	// true
	IncludeValues *bool `json:"IncludeValues,omitempty" xml:"IncludeValues,omitempty"`
	// The similarity algorithm used during search. Valid values:
	//
	// - **l2**: Euclidean distance.
	//
	// - **ip**: Dot product (inner product) distance.
	//
	// - **cosine**: Cosine similarity.
	//
	// > If this parameter is left empty, the algorithm specified when the index is built is used.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view the list.
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
	// The retrieval starting point during paged queries. Default value: empty. This parameter is not supported in dual-channel recall scenarios.
	//
	// The value must be greater than or equal to 0. If this parameter is specified, the Total parameter is returned to indicate the total number of hits. This parameter is used together with TopK. For example, to retrieve chunks with chunk_id ranging from 0 to 44 with a paging size of 20, you must make three requests:
	//
	// - Offset=0, TopK=20 returns chunks with chunk_id ranging from 0 to 19.
	//
	// - Offset=20, TopK=20 returns chunks with chunk_id ranging from 20 to 39.
	//
	// - Offset=30, TopK=20 returns chunks with chunk_id ranging from 40 to 44.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The field based on which the results are sorted. Default value: empty. This parameter is not supported in hybrid search scenarios.
	//
	// The field must belong to the metadata or be a default field in the table, such as id. Supported formats:
	//
	// - A single field, such as chunk_id.
	//
	// - Multiple fields separated by commas, such as block_id, chunk_id.
	//
	// - Descending order, such as block_id DESC, chunk_id DESC.
	//
	// example:
	//
	// chunk_id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Uses another relational table to filter vector data, similar to the JOIN operation.
	//
	// > You can set the IncludeMetadataFields parameter to return data from the relational table. For example, rds_table_name.id indicates that the id field of the relational table is returned.
	RelationalTableFilterShrink *string `json:"RelationalTableFilter,omitempty" xml:"RelationalTableFilter,omitempty"`
	// The list of sparse vector data.
	SparseVectorShrink *string `json:"SparseVector,omitempty" xml:"SparseVector,omitempty"`
	// The number of top results to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The vector data. The length must be consistent with the dimension specified in the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) operation.
	//
	// >
	//
	// > - If SparseVector is left empty, the system uses it to retrieve only dense vector results.
	//
	// > - If both Vector and SparseVector are left empty, the system uses it to retrieve only full-text results.
	VectorShrink *string `json:"Vector,omitempty" xml:"Vector,omitempty"`
	// The ID of the workspace that consists of multiple database instances. This parameter and the DBInstanceId parameter cannot be left empty at the same time. If both parameters are specified, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCollectionDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *QueryCollectionDataShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *QueryCollectionDataShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryCollectionDataShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *QueryCollectionDataShrinkRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *QueryCollectionDataShrinkRequest) GetHybridSearchArgsShrink() *string {
	return s.HybridSearchArgsShrink
}

func (s *QueryCollectionDataShrinkRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *QueryCollectionDataShrinkRequest) GetIncludeSparseValues() *bool {
	return s.IncludeSparseValues
}

func (s *QueryCollectionDataShrinkRequest) GetIncludeValues() *bool {
	return s.IncludeValues
}

func (s *QueryCollectionDataShrinkRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryCollectionDataShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryCollectionDataShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *QueryCollectionDataShrinkRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *QueryCollectionDataShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryCollectionDataShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCollectionDataShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCollectionDataShrinkRequest) GetRelationalTableFilterShrink() *string {
	return s.RelationalTableFilterShrink
}

func (s *QueryCollectionDataShrinkRequest) GetSparseVectorShrink() *string {
	return s.SparseVectorShrink
}

func (s *QueryCollectionDataShrinkRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *QueryCollectionDataShrinkRequest) GetVectorShrink() *string {
	return s.VectorShrink
}

func (s *QueryCollectionDataShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCollectionDataShrinkRequest) SetCollection(v string) *QueryCollectionDataShrinkRequest {
	s.Collection = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetContent(v string) *QueryCollectionDataShrinkRequest {
	s.Content = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetDBInstanceId(v string) *QueryCollectionDataShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetFilter(v string) *QueryCollectionDataShrinkRequest {
	s.Filter = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetHybridSearch(v string) *QueryCollectionDataShrinkRequest {
	s.HybridSearch = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetHybridSearchArgsShrink(v string) *QueryCollectionDataShrinkRequest {
	s.HybridSearchArgsShrink = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetIncludeMetadataFields(v string) *QueryCollectionDataShrinkRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetIncludeSparseValues(v bool) *QueryCollectionDataShrinkRequest {
	s.IncludeSparseValues = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetIncludeValues(v bool) *QueryCollectionDataShrinkRequest {
	s.IncludeValues = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetMetrics(v string) *QueryCollectionDataShrinkRequest {
	s.Metrics = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetNamespace(v string) *QueryCollectionDataShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetNamespacePassword(v string) *QueryCollectionDataShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetOffset(v int32) *QueryCollectionDataShrinkRequest {
	s.Offset = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetOrderBy(v string) *QueryCollectionDataShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetOwnerId(v int64) *QueryCollectionDataShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetRegionId(v string) *QueryCollectionDataShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetRelationalTableFilterShrink(v string) *QueryCollectionDataShrinkRequest {
	s.RelationalTableFilterShrink = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetSparseVectorShrink(v string) *QueryCollectionDataShrinkRequest {
	s.SparseVectorShrink = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetTopK(v int64) *QueryCollectionDataShrinkRequest {
	s.TopK = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetVectorShrink(v string) *QueryCollectionDataShrinkRequest {
	s.VectorShrink = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) SetWorkspaceId(v string) *QueryCollectionDataShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCollectionDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
