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
	SetIncludeSparseValues(v bool) *QueryCollectionDataRequest
	GetIncludeSparseValues() *bool
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
	HybridSearchArgs map[string]map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
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
	RelationalTableFilter *QueryCollectionDataRequestRelationalTableFilter `json:"RelationalTableFilter,omitempty" xml:"RelationalTableFilter,omitempty" type:"Struct"`
	// The sparse vector data.
	SparseVector *QueryCollectionDataRequestSparseVector `json:"SparseVector,omitempty" xml:"SparseVector,omitempty" type:"Struct"`
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
	Vector []*float64 `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Repeated"`
	// The ID of the workspace that consists of multiple database instances. You must specify this parameter or the DBInstanceId parameter. If both this parameter and DBInstanceId are specified, this parameter is used.
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

func (s *QueryCollectionDataRequest) GetIncludeSparseValues() *bool {
	return s.IncludeSparseValues
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

func (s *QueryCollectionDataRequest) SetIncludeSparseValues(v bool) *QueryCollectionDataRequest {
	s.IncludeSparseValues = &v
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
	if s.RelationalTableFilter != nil {
		if err := s.RelationalTableFilter.Validate(); err != nil {
			return err
		}
	}
	if s.SparseVector != nil {
		if err := s.SparseVector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCollectionDataRequestRelationalTableFilter struct {
	// The metadata field of the vector collection, which is used to associate with the fields of the vector table.
	//
	// example:
	//
	// doc_id
	CollectionMetadataField *string `json:"CollectionMetadataField,omitempty" xml:"CollectionMetadataField,omitempty"`
	// The filter conditions for the relational table.
	//
	// example:
	//
	// tags @> ARRAY[\\"art\\"]
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The field of the relational table, which is used to associate with the metadata field of the vector collection.
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
	// The array of indexes.
	//
	// > The number of elements in the array cannot exceed 4,000.
	Indices []*int64 `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
	// The array of sparse vectors.
	Values []*float64 `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
