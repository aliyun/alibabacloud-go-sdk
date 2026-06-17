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
	// The name of the collection.
	//
	// > You can call the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) operation to list available collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The content for full-text search. If this parameter is omitted, only vector search is performed. If this parameter is specified, the system performs a hybrid search of vector search and full-text search.
	//
	// > You must specify one of the Content and Vector parameters.
	//
	// example:
	//
	// hello_world
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details for all AnalyticDB for PostgreSQL instances in a region, including their instance IDs.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The filter conditions for data retrieval. It is in the format of a WHERE clause in SQL. This expression returns a boolean value, which can be a simple comparison operator, such as `=`, `<>`, `!=`, `>`, `<`, `>=`, and `<=`, or a more complex expression combined with logical operators, such as `AND`, `OR`, and `NOT`, and keywords such as `IN`, `BETWEEN`, and `LIKE`.
	//
	// > - For more information about the syntax, see [PostgreSQL WHERE](https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/).
	//
	// example:
	//
	// pipeline_id=\\"1yhpmo0rbn\\" AND (spu=\\"10025667796135\\" AND dept_id=\\"226\\")
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The hybrid search algorithm. If this parameter is empty, the system ranks results by directly comparing the scores from the vector search and the full-text search.
	//
	// Valid values:
	//
	// - `RRF`: Reciprocal Rank Fusion. This algorithm has a parameter k to control the fusion effect. For more information, see the description of the `HybridSearchArgs` parameter.
	//
	// - `Weight`: weighted sort. This algorithm uses a parameter alpha to control the score ratio of vector search and full-text search, and then sorts the results. For more information about the parameter, see the `HybridSearchArgs` parameter.
	//
	// - `Cascaded`: performs a full-text search, and then performs a vector search on the search results.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The parameters for the hybrid search algorithm. The following algorithms are supported: RRF and Weight.
	//
	// - For RRF, specify the constant k in the scoring algorithm `1/(k+rank_i)`. The value must be a positive integer greater than 1. The format is as follows:
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
	// - For Weight, in the formula `alpha 	- vector_score + (1-alpha) 	- text_score`, the alpha parameter indicates the score ratio of the vector search to the full-text search. The value ranges from 0 to 1. 0 indicates that only the full-text search is used, and 1 indicates that only the vector search is used.
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
	// This parameter is left empty by default. It specifies the metadata fields to be returned. You can specify multiple fields and separate them with commas (,).
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
	// The similarity algorithm for search. Valid values:
	//
	// - **l2**: the Euclidean distance.
	//
	// - **ip**: the dot product distance.
	//
	// - **cosine**: the cosine similarity.
	//
	// > If this parameter is not specified, the algorithm specified when the index is created is used.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The name of the namespace.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to list available namespaces.
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
	// This parameter is left empty by default. It specifies the start position of a paged query. This parameter is not supported in hybrid search.
	//
	// The value must be greater than or equal to 0. When this parameter is not empty, Total in the response indicates the total number of hits. This parameter is used with TopK. For example, if you want to retrieve chunks 0 to 44 with a page size of 20, you must send three requests:
	//
	// - `Offset=0, TopK=20` returns chunks 0 to 19.
	//
	// - `Offset=20, TopK=20` returns chunks 20 to 39.
	//
	// - `Offset=40, TopK=20` returns chunks 40 to 44.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// This parameter is left empty by default. It specifies the field based on which to sort the results. This parameter is not supported in hybrid search.
	//
	// The field must be a metadata field or a default field in the table, such as `id`. The following formats are supported:
	//
	// - A single field, such as `chunk_id`.
	//
	// - Multiple fields separated by commas (,), such as `block_id, chunk_id`.
	//
	// - Descending order, such as `block_id DESC, chunk_id DESC`.
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
	// Uses another relational table to filter vector data, which is similar to the JOIN operation.
	//
	// > The data of the relational table can be returned by setting the IncludeMetadataFields parameter. For example, `rds_table_name.id` indicates that the id field of the relational table is returned.
	RelationalTableFilterShrink *string `json:"RelationalTableFilter,omitempty" xml:"RelationalTableFilter,omitempty"`
	// The sparse vector data.
	SparseVectorShrink *string `json:"SparseVector,omitempty" xml:"SparseVector,omitempty"`
	// Specifies the number of top results to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The vector data. The length of the vector data must be the same as that specified in the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) operation.
	//
	// > - If `SparseVector` is empty, only the dense vector search results are returned.
	//
	// >
	//
	// > - If both `Vector` and `SparseVector` are empty, only the full-text search results are returned.
	VectorShrink *string `json:"Vector,omitempty" xml:"Vector,omitempty"`
	// The ID of the workspace that consists of multiple database instances. You must specify this parameter or the DBInstanceId parameter. If both this parameter and DBInstanceId are specified, this parameter is used.
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
