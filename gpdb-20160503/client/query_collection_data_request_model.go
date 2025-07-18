// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCollectionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *QueryCollectionDataRequest
	GetCollection() *string
	SetContent(v string) *QueryCollectionDataRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryCollectionDataRequest
	GetDBInstanceId() *string
	SetFilter(v string) *QueryCollectionDataRequest
	GetFilter() *string
	SetHybridSearch(v string) *QueryCollectionDataRequest
	GetHybridSearch() *string
	SetHybridSearchArgs(v map[string]map[string]interface{}) *QueryCollectionDataRequest
	GetHybridSearchArgs() map[string]map[string]interface{}
	SetIncludeMetadataFields(v string) *QueryCollectionDataRequest
	GetIncludeMetadataFields() *string
	SetIncludeValues(v bool) *QueryCollectionDataRequest
	GetIncludeValues() *bool
	SetMetrics(v string) *QueryCollectionDataRequest
	GetMetrics() *string
	SetNamespace(v string) *QueryCollectionDataRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *QueryCollectionDataRequest
	GetNamespacePassword() *string
	SetOffset(v int32) *QueryCollectionDataRequest
	GetOffset() *int32
	SetOrderBy(v string) *QueryCollectionDataRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *QueryCollectionDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryCollectionDataRequest
	GetRegionId() *string
	SetRelationalTableFilter(v *QueryCollectionDataRequestRelationalTableFilter) *QueryCollectionDataRequest
	GetRelationalTableFilter() *QueryCollectionDataRequestRelationalTableFilter
	SetSparseVector(v *QueryCollectionDataRequestSparseVector) *QueryCollectionDataRequest
	GetSparseVector() *QueryCollectionDataRequestSparseVector
	SetTopK(v int64) *QueryCollectionDataRequest
	GetTopK() *int64
	SetVector(v []*float64) *QueryCollectionDataRequest
	GetVector() []*float64
	SetWorkspaceId(v string) *QueryCollectionDataRequest
	GetWorkspaceId() *string
}

type QueryCollectionDataRequest struct {
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
	Filter       *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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
	HybridSearchArgs map[string]map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// Defaults to empty, indicating the metadata fields to return. Multiple fields should be separated by commas.
	//
	// example:
	//
	// title,content
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
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
	RelationalTableFilter *QueryCollectionDataRequestRelationalTableFilter `json:"RelationalTableFilter,omitempty" xml:"RelationalTableFilter,omitempty" type:"Struct"`
	SparseVector          *QueryCollectionDataRequestSparseVector          `json:"SparseVector,omitempty" xml:"SparseVector,omitempty" type:"Struct"`
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
	Vector []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
	// The ID of the Workspace composed of multiple database instances. This parameter and `DBInstanceId` cannot both be empty. If both are specified, this parameter takes precedence.
	//
	// example:
	//
	// gp-ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryCollectionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataRequest) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataRequest) GetCollection() *string {
	return s.Collection
}

func (s *QueryCollectionDataRequest) GetContent() *string {
	return s.Content
}

func (s *QueryCollectionDataRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryCollectionDataRequest) GetFilter() *string {
	return s.Filter
}

func (s *QueryCollectionDataRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *QueryCollectionDataRequest) GetHybridSearchArgs() map[string]map[string]interface{} {
	return s.HybridSearchArgs
}

func (s *QueryCollectionDataRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *QueryCollectionDataRequest) GetIncludeValues() *bool {
	return s.IncludeValues
}

func (s *QueryCollectionDataRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryCollectionDataRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryCollectionDataRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *QueryCollectionDataRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *QueryCollectionDataRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryCollectionDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryCollectionDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCollectionDataRequest) GetRelationalTableFilter() *QueryCollectionDataRequestRelationalTableFilter {
	return s.RelationalTableFilter
}

func (s *QueryCollectionDataRequest) GetSparseVector() *QueryCollectionDataRequestSparseVector {
	return s.SparseVector
}

func (s *QueryCollectionDataRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *QueryCollectionDataRequest) GetVector() []*float64 {
	return s.Vector
}

func (s *QueryCollectionDataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryCollectionDataRequest) SetCollection(v string) *QueryCollectionDataRequest {
	s.Collection = &v
	return s
}

func (s *QueryCollectionDataRequest) SetContent(v string) *QueryCollectionDataRequest {
	s.Content = &v
	return s
}

func (s *QueryCollectionDataRequest) SetDBInstanceId(v string) *QueryCollectionDataRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryCollectionDataRequest) SetFilter(v string) *QueryCollectionDataRequest {
	s.Filter = &v
	return s
}

func (s *QueryCollectionDataRequest) SetHybridSearch(v string) *QueryCollectionDataRequest {
	s.HybridSearch = &v
	return s
}

func (s *QueryCollectionDataRequest) SetHybridSearchArgs(v map[string]map[string]interface{}) *QueryCollectionDataRequest {
	s.HybridSearchArgs = v
	return s
}

func (s *QueryCollectionDataRequest) SetIncludeMetadataFields(v string) *QueryCollectionDataRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *QueryCollectionDataRequest) SetIncludeValues(v bool) *QueryCollectionDataRequest {
	s.IncludeValues = &v
	return s
}

func (s *QueryCollectionDataRequest) SetMetrics(v string) *QueryCollectionDataRequest {
	s.Metrics = &v
	return s
}

func (s *QueryCollectionDataRequest) SetNamespace(v string) *QueryCollectionDataRequest {
	s.Namespace = &v
	return s
}

func (s *QueryCollectionDataRequest) SetNamespacePassword(v string) *QueryCollectionDataRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryCollectionDataRequest) SetOffset(v int32) *QueryCollectionDataRequest {
	s.Offset = &v
	return s
}

func (s *QueryCollectionDataRequest) SetOrderBy(v string) *QueryCollectionDataRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryCollectionDataRequest) SetOwnerId(v int64) *QueryCollectionDataRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCollectionDataRequest) SetRegionId(v string) *QueryCollectionDataRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCollectionDataRequest) SetRelationalTableFilter(v *QueryCollectionDataRequestRelationalTableFilter) *QueryCollectionDataRequest {
	s.RelationalTableFilter = v
	return s
}

func (s *QueryCollectionDataRequest) SetSparseVector(v *QueryCollectionDataRequestSparseVector) *QueryCollectionDataRequest {
	s.SparseVector = v
	return s
}

func (s *QueryCollectionDataRequest) SetTopK(v int64) *QueryCollectionDataRequest {
	s.TopK = &v
	return s
}

func (s *QueryCollectionDataRequest) SetVector(v []*float64) *QueryCollectionDataRequest {
	s.Vector = v
	return s
}

func (s *QueryCollectionDataRequest) SetWorkspaceId(v string) *QueryCollectionDataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryCollectionDataRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataRequestRelationalTableFilter struct {
	// The Metadata field of the vector collection, used to associate with the fields in the vector table.
	//
	// example:
	//
	// doc_id
	CollectionMetadataField *string `json:"CollectionMetadataField,omitempty" xml:"CollectionMetadataField,omitempty"`
	// The filtering condition for the relational table.
	//
	// example:
	//
	// tags @> ARRAY[\\"art\\"]
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The field in the relational table, used to associate with the Metadata field of the vector collection.
	//
	// example:
	//
	// id
	TableField *string `json:"TableField,omitempty" xml:"TableField,omitempty"`
	// The name of the relational table.
	//
	// example:
	//
	// my_rds_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s QueryCollectionDataRequestRelationalTableFilter) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataRequestRelationalTableFilter) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataRequestRelationalTableFilter) GetCollectionMetadataField() *string {
	return s.CollectionMetadataField
}

func (s *QueryCollectionDataRequestRelationalTableFilter) GetCondition() *string {
	return s.Condition
}

func (s *QueryCollectionDataRequestRelationalTableFilter) GetTableField() *string {
	return s.TableField
}

func (s *QueryCollectionDataRequestRelationalTableFilter) GetTableName() *string {
	return s.TableName
}

func (s *QueryCollectionDataRequestRelationalTableFilter) SetCollectionMetadataField(v string) *QueryCollectionDataRequestRelationalTableFilter {
	s.CollectionMetadataField = &v
	return s
}

func (s *QueryCollectionDataRequestRelationalTableFilter) SetCondition(v string) *QueryCollectionDataRequestRelationalTableFilter {
	s.Condition = &v
	return s
}

func (s *QueryCollectionDataRequestRelationalTableFilter) SetTableField(v string) *QueryCollectionDataRequestRelationalTableFilter {
	s.TableField = &v
	return s
}

func (s *QueryCollectionDataRequestRelationalTableFilter) SetTableName(v string) *QueryCollectionDataRequestRelationalTableFilter {
	s.TableName = &v
	return s
}

func (s *QueryCollectionDataRequestRelationalTableFilter) Validate() error {
	return dara.Validate(s)
}

type QueryCollectionDataRequestSparseVector struct {
	Indices []*int64   `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
	Values  []*float64 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryCollectionDataRequestSparseVector) String() string {
	return dara.Prettify(s)
}

func (s QueryCollectionDataRequestSparseVector) GoString() string {
	return s.String()
}

func (s *QueryCollectionDataRequestSparseVector) GetIndices() []*int64 {
	return s.Indices
}

func (s *QueryCollectionDataRequestSparseVector) GetValues() []*float64 {
	return s.Values
}

func (s *QueryCollectionDataRequestSparseVector) SetIndices(v []*int64) *QueryCollectionDataRequestSparseVector {
	s.Indices = v
	return s
}

func (s *QueryCollectionDataRequestSparseVector) SetValues(v []*float64) *QueryCollectionDataRequestSparseVector {
	s.Values = v
	return s
}

func (s *QueryCollectionDataRequestSparseVector) Validate() error {
	return dara.Validate(s)
}
