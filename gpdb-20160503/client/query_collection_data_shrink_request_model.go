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
	// Collection name.
	//
	// > You can use the [ListCollections](https://help.aliyun.com/document_detail/2401503.html) API to view the list.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// Content for full-text search. When this value is empty, only vector search is used; when it is not empty, both vector and full-text search are used.
	//
	// > The Vector parameter cannot be empty at the same time.
	//
	// example:
	//
	// hello_world
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Filter conditions for the data to be queried, in SQL WHERE format. It is an expression that returns a boolean value (true or false). Conditions can be simple comparison operators such as equal (=), not equal (<> or !=), greater than (>), less than (<), greater than or equal to (>=), less than or equal to (<=), or more complex expressions combined with logical operators (AND, OR, NOT), as well as conditions using keywords like IN, BETWEEN, and LIKE.
	//
	// >
	//
	// > - For detailed syntax, refer to: https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/
	//
	// example:
	//
	// response > 200
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Dual-path recall algorithm, default is empty (i.e., directly compare and sort the scores of vectors and full-text).
	//
	// Available values:
	//
	// - RRF: Reciprocal rank fusion, with a parameter k controlling the fusion effect. See HybridSearchArgs configuration for details;
	//
	// - Weight: Weighted sorting, using a parameter alpha to control the score ratio of vectors and full-text, then sorting. See HybridSearchArgs configuration for details;
	//
	// - Cascaded: Perform full-text search first, then vector search based on the full-text results;
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The parameters of the two-way retrieval algorithm. The following parameters are supported:
	//
	// 	- When HybridSearch is set to RRF, the scores are calculated by using the `1/(k+rank_i)` formula. The constant k is a positive integer that is greater than 1.
	//
	// <!---->
	//
	//     {
	//
	//        "RRF": {
	//
	//         "k": 60
	//
	//        }
	//
	//     }
	//
	// 	- When HybridSearch is set to Weight, the scores are calculated by using the `alpha 	- vector_score + (1-alpha) 	- text_score` formula. The alpha parameter specifies the proportion of the vector search score and the full-text search score and ranges from 0 to 1. A value of 0 specifies full-text search and a value of 1 specifies vector search.
	//
	// <!---->
	//
	//     {
	//
	//        "Weight": {
	//
	//         "alpha": 0.5
	//
	//        }
	//
	//     }
	HybridSearchArgsShrink *string `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// Defaults to empty, indicating the metadata fields to return. Multiple fields should be separated by commas.
	//
	// example:
	//
	// title,content
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
	IncludeSparseValues   *bool   `json:"IncludeSparseValues,omitempty" xml:"IncludeSparseValues,omitempty"`
	// Whether to return vector data. Value descriptions:
	//
	// - **true**: Return vector data.
	//
	// - **false**: Do not return vector data, used for full-text search scenarios.
	//
	// example:
	//
	// true
	IncludeValues *bool `json:"IncludeValues,omitempty" xml:"IncludeValues,omitempty"`
	// Similarity algorithm used during retrieval. Value descriptions:
	//
	// - **l2**: Euclidean distance.
	//
	// - **ip**: Inner product (dot product) distance.
	//
	// - **cosine**: Cosine similarity.
	//
	// > If this value is empty, the algorithm specified during index creation is used.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// Namespace.
	//
	// > You can use the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) API to view the list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Password for the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// Defaults to empty, indicating the starting point for pagination queries. Does not support hybrid search scenarios.
	//
	// The value must be >= 0. When this value is not empty, it will return `Total`, which indicates the total number of hits. This parameter works with `TopK`. For example, to paginate 20 and retrieve chunks with `chunk_id` from 0 to 44, you need to make three requests:
	//
	// - `Offset=0, TopK=20` returns `chunk_id` 0~19
	//
	// - `Offset=20, TopK=20` returns `chunk_id` 20~39
	//
	// - `Offset=30, TopK=20` returns `chunk_id` 40~44
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// Defaults to empty, indicating the field for sorting. Does not support hybrid search scenarios.
	//
	// The field must belong to metadata or be a default field in the table, such as `id`. The supported formats are:
	//
	// - A single field, e.g., `chunk_id`;
	//
	// - Multiple fields, separated by commas, e.g., `block_id, chunk_id`;
	//
	// - Supports reverse order, e.g., `block_id DESC, chunk_id DESC`;
	//
	// example:
	//
	// chunk_id
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Uses another relational table to filter vector data (similar to a Join function).
	//
	// > Data from the relational table can be returned by setting the `IncludeMetadataFields` parameter. For example, `rds_table_name.id` indicates returning the `id` field from the relational table.
	RelationalTableFilterShrink *string `json:"RelationalTableFilter,omitempty" xml:"RelationalTableFilter,omitempty"`
	SparseVectorShrink          *string `json:"SparseVector,omitempty" xml:"SparseVector,omitempty"`
	// Set the number of top results to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Vector data, with the same dimension as specified in the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) API.
	//
	// > When the vector is empty, only full-text search results are returned.
	VectorShrink *string `json:"Vector,omitempty" xml:"Vector,omitempty"`
	// The ID of the Workspace composed of multiple database instances. This parameter and `DBInstanceId` cannot both be empty. If both are specified, this parameter takes precedence.
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
