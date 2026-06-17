// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *QueryContentShrinkRequest
	GetCollection() *string
	SetContent(v string) *QueryContentShrinkRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryContentShrinkRequest
	GetDBInstanceId() *string
	SetFileName(v string) *QueryContentShrinkRequest
	GetFileName() *string
	SetFileUrl(v string) *QueryContentShrinkRequest
	GetFileUrl() *string
	SetFilter(v string) *QueryContentShrinkRequest
	GetFilter() *string
	SetGraphEnhance(v bool) *QueryContentShrinkRequest
	GetGraphEnhance() *bool
	SetGraphSearchArgsShrink(v string) *QueryContentShrinkRequest
	GetGraphSearchArgsShrink() *string
	SetHybridSearch(v string) *QueryContentShrinkRequest
	GetHybridSearch() *string
	SetHybridSearchArgsShrink(v string) *QueryContentShrinkRequest
	GetHybridSearchArgsShrink() *string
	SetIncludeFileUrl(v bool) *QueryContentShrinkRequest
	GetIncludeFileUrl() *bool
	SetIncludeMetadataFields(v string) *QueryContentShrinkRequest
	GetIncludeMetadataFields() *string
	SetIncludeVector(v bool) *QueryContentShrinkRequest
	GetIncludeVector() *bool
	SetMetrics(v string) *QueryContentShrinkRequest
	GetMetrics() *string
	SetNamespace(v string) *QueryContentShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *QueryContentShrinkRequest
	GetNamespacePassword() *string
	SetOffset(v int32) *QueryContentShrinkRequest
	GetOffset() *int32
	SetOrderBy(v string) *QueryContentShrinkRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *QueryContentShrinkRequest
	GetOwnerId() *int64
	SetRecallWindowShrink(v string) *QueryContentShrinkRequest
	GetRecallWindowShrink() *string
	SetRegionId(v string) *QueryContentShrinkRequest
	GetRegionId() *string
	SetRerankFactor(v float64) *QueryContentShrinkRequest
	GetRerankFactor() *float64
	SetRerankModelShrink(v string) *QueryContentShrinkRequest
	GetRerankModelShrink() *string
	SetTopK(v int32) *QueryContentShrinkRequest
	GetTopK() *int32
	SetUrlExpiration(v string) *QueryContentShrinkRequest
	GetUrlExpiration() *string
	SetUseFullTextRetrieval(v bool) *QueryContentShrinkRequest
	GetUseFullTextRetrieval() *bool
}

type QueryContentShrinkRequest struct {
	// The name of the document collection.
	//
	// > A document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. Call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to list your document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The text to use for retrieval.
	//
	// example:
	//
	// What is AnalyticDB for PostgreSQL?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to find the IDs of all AnalyticDB for PostgreSQL instances in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The filename of the source image for a search-by-image query.
	//
	// > The image file must have a file extension. The supported extensions are bmp, jpg, jpeg, png, and tiff.
	//
	// example:
	//
	// test.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The publicly accessible URL of the image file for a search-by-image query.
	//
	// > The image file must have a file extension. The supported extensions are bmp, jpg, jpeg, png, and tiff.
	//
	// example:
	//
	// https://xx/myImage.jpg
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// A filter condition for the query, specified as a SQL `WHERE` clause that returns a boolean value. The clause can use comparison operators (such as `=`, `<>`, `!=`, `>`, `<`, `>=`, and `<=`), logical operators (`AND`, `OR`, and `NOT`), and keywords such as `IN`, `BETWEEN`, and `LIKE`.
	//
	// > - For details about the syntax, see https\\://www\\.postgresqltutorial.com/postgresql-tutorial/postgresql-where/.
	//
	// example:
	//
	// title = \\"test\\" AND name like \\"test%\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to enable knowledge graph enhancement. The default value is `false`.
	//
	// example:
	//
	// false
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// The parameters for knowledge graph retrieval.
	GraphSearchArgsShrink *string `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty"`
	// Specifies the hybrid search algorithm. If this parameter is not specified, the system directly compares and sorts the scores from dense vector retrieval and full-text search.
	//
	// Valid values:
	//
	// - RRF: reciprocal rank fusion. This method uses a `k` parameter to control the fusion effect. For more information, see the `HybridSearchArgs` parameter.
	//
	// - Weight: A weighted sorting method. This method uses parameters to control the score weights of vector retrieval and full-text search before sorting. For more information, see the `HybridSearchArgs` parameter.
	//
	// - Cascaded: Performs full-text search first, and then performs vector retrieval on the results.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// Parameters for the multi-channel recall algorithm. Currently, `RRF` and `Weight` are supported. `HybridPathsSetting` can be used to specify the recall paths, including dense vector search (`dense`), sparse vector search (`sparse`), and full-text search (`fulltext`). If this parameter is not specified, the system recalls dense vectors and full-text search results by default.
	//
	// - RRF: Specifies the constant `k` in the scoring formula `1/(k+rank_i)`. The value must be an integer greater than 1. Example:
	//
	// ```
	//
	// {
	//
	//   "HybridPathsSetting": {
	//
	//     "paths": "dense,fulltext"
	//
	//   },
	//
	//   "RRF": {
	//
	//     "k": 60
	//
	//   }
	//
	// }
	//
	// ```
	//
	// - Weight:
	//
	//   - For dual-channel recall (if `HybridPathsSetting` is not specified, only `alpha` is configured):
	//
	//     - The score is calculated using the formula: `alpha 	- dense_score + (1-alpha) 	- fulltext_score`. The `alpha` parameter represents the score weight of dense vector retrieval relative to full-text search. The value must be in the range of 0 to 1. A value of 0 indicates that only full-text search is used, and a value of 1 indicates that only dense vector retrieval is used.
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
	//
	// - For three-channel recall:
	//
	//   - The score is calculated using the formula: `normalized_dense 	- dense_score + normalized_sparse 	- sparse_score + normalized_fulltext 	- fulltext_score`. The `dense`, `sparse`, and `fulltext` parameters represent the weights for the dense vector, sparse vector, and full-text search results, respectively. Their values must be greater than or equal to 0. The system automatically normalizes the weights to a sum of 1 (for example, `normalized_x = x / (dense + sparse + fulltext)`).
	//
	// ```
	//
	// {
	//
	//   "HybridPathsSetting": {
	//
	//      "paths": "dense,sparse,fulltext"
	//
	//    },
	//
	//   "Weight": {
	//
	//     "dense": 0.5,
	//
	//     "sparse": 0.3,
	//
	//     "fulltext": 0.2
	//
	//   }
	//
	// }
	//
	// ```
	HybridSearchArgsShrink *string `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// Specifies whether to return the URL of the document. The default value is `false`.
	//
	// example:
	//
	// false
	IncludeFileUrl *bool `json:"IncludeFileUrl,omitempty" xml:"IncludeFileUrl,omitempty"`
	// The metadata fields to include in the response. If left empty, no metadata is returned. To specify multiple fields, use a comma-separated list.
	//
	// example:
	//
	// title,page
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
	// Specifies whether to include the vector in the results. The default value is `false`.
	//
	// > - **false**: The vector is not returned.
	//
	// >
	//
	// > - **true**: The vector is returned.
	//
	// example:
	//
	// true
	IncludeVector *bool `json:"IncludeVector,omitempty" xml:"IncludeVector,omitempty"`
	// The similarity algorithm used for retrieval. If this parameter is not specified, the algorithm that was specified when the document collection was created is used. We recommend that you do not set this parameter unless you have specific requirements.
	//
	// > Valid values:
	//
	// >
	//
	// > - **l2**: Euclidean distance.
	//
	// >
	//
	// > - **ip**: dot product (inner product) distance.
	//
	// >
	//
	// > - **cosine**: cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace. The default value is `public`.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to list existing namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > This password is set when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The offset for paginated queries.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The field to sort the results by. By default, this parameter is empty.
	//
	// The field must be a metadata field or a default field from the table, such as `id`. Supported formats include single fields (e.g., `chunk_id`), multiple comma-separated fields (e.g., `block_id, chunk_id`), and fields with a sort direction (e.g., `block_id DESC, chunk_id DESC`).
	//
	// example:
	//
	// created_at
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recall window. If this parameter is specified, additional context is returned with the retrieval results. The value must be a two-element array, `[A, B]`, where `-10 <= A <= 0` and `0 <= B <= 10`.
	//
	// > - Use this parameter when documents are finely chunked and retrieval might otherwise lose contextual information.
	//
	// >
	//
	// > - Reranking is prioritized over windowing. The system first applies reranking and then processes the window.
	RecallWindowShrink *string `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The factor for reranking vector retrieval results. The value must be greater than 1 and less than or equal to 5.
	//
	// > - Reranking may be slower if document chunks are sparse.
	//
	// >
	//
	// > - For best performance, the number of items to be reranked (`TopK` \\	- `RerankFactor`, rounded up) should not exceed 50.
	//
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The parameters for the reranking model.
	RerankModelShrink *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The number of top results to return.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The validity period of the returned image URL.
	//
	// > - The value can be specified in seconds (s) or days (d). For example, `300s` indicates that the link is valid for 300 seconds, and `60d` indicates that the link is valid for 60 days.
	//
	// >
	//
	// > - The value must be in the range of `60s` to `365d`.
	//
	// >
	//
	// > - Default value: `7200s` (2 hours).
	//
	// example:
	//
	// 7200s
	UrlExpiration *string `json:"UrlExpiration,omitempty" xml:"UrlExpiration,omitempty"`
	// (Deprecated) Specifies whether to use full-text search (dual-channel recall). The default value is `false`, which means only vector retrieval is used.
	//
	// example:
	//
	// true
	UseFullTextRetrieval *bool `json:"UseFullTextRetrieval,omitempty" xml:"UseFullTextRetrieval,omitempty"`
}

func (s QueryContentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryContentShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *QueryContentShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *QueryContentShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryContentShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *QueryContentShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *QueryContentShrinkRequest) GetGraphEnhance() *bool {
	return s.GraphEnhance
}

func (s *QueryContentShrinkRequest) GetGraphSearchArgsShrink() *string {
	return s.GraphSearchArgsShrink
}

func (s *QueryContentShrinkRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *QueryContentShrinkRequest) GetHybridSearchArgsShrink() *string {
	return s.HybridSearchArgsShrink
}

func (s *QueryContentShrinkRequest) GetIncludeFileUrl() *bool {
	return s.IncludeFileUrl
}

func (s *QueryContentShrinkRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *QueryContentShrinkRequest) GetIncludeVector() *bool {
	return s.IncludeVector
}

func (s *QueryContentShrinkRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryContentShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryContentShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *QueryContentShrinkRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *QueryContentShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryContentShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryContentShrinkRequest) GetRecallWindowShrink() *string {
	return s.RecallWindowShrink
}

func (s *QueryContentShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryContentShrinkRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *QueryContentShrinkRequest) GetRerankModelShrink() *string {
	return s.RerankModelShrink
}

func (s *QueryContentShrinkRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *QueryContentShrinkRequest) GetUrlExpiration() *string {
	return s.UrlExpiration
}

func (s *QueryContentShrinkRequest) GetUseFullTextRetrieval() *bool {
	return s.UseFullTextRetrieval
}

func (s *QueryContentShrinkRequest) SetCollection(v string) *QueryContentShrinkRequest {
	s.Collection = &v
	return s
}

func (s *QueryContentShrinkRequest) SetContent(v string) *QueryContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *QueryContentShrinkRequest) SetDBInstanceId(v string) *QueryContentShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryContentShrinkRequest) SetFileName(v string) *QueryContentShrinkRequest {
	s.FileName = &v
	return s
}

func (s *QueryContentShrinkRequest) SetFileUrl(v string) *QueryContentShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *QueryContentShrinkRequest) SetFilter(v string) *QueryContentShrinkRequest {
	s.Filter = &v
	return s
}

func (s *QueryContentShrinkRequest) SetGraphEnhance(v bool) *QueryContentShrinkRequest {
	s.GraphEnhance = &v
	return s
}

func (s *QueryContentShrinkRequest) SetGraphSearchArgsShrink(v string) *QueryContentShrinkRequest {
	s.GraphSearchArgsShrink = &v
	return s
}

func (s *QueryContentShrinkRequest) SetHybridSearch(v string) *QueryContentShrinkRequest {
	s.HybridSearch = &v
	return s
}

func (s *QueryContentShrinkRequest) SetHybridSearchArgsShrink(v string) *QueryContentShrinkRequest {
	s.HybridSearchArgsShrink = &v
	return s
}

func (s *QueryContentShrinkRequest) SetIncludeFileUrl(v bool) *QueryContentShrinkRequest {
	s.IncludeFileUrl = &v
	return s
}

func (s *QueryContentShrinkRequest) SetIncludeMetadataFields(v string) *QueryContentShrinkRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *QueryContentShrinkRequest) SetIncludeVector(v bool) *QueryContentShrinkRequest {
	s.IncludeVector = &v
	return s
}

func (s *QueryContentShrinkRequest) SetMetrics(v string) *QueryContentShrinkRequest {
	s.Metrics = &v
	return s
}

func (s *QueryContentShrinkRequest) SetNamespace(v string) *QueryContentShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *QueryContentShrinkRequest) SetNamespacePassword(v string) *QueryContentShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryContentShrinkRequest) SetOffset(v int32) *QueryContentShrinkRequest {
	s.Offset = &v
	return s
}

func (s *QueryContentShrinkRequest) SetOrderBy(v string) *QueryContentShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryContentShrinkRequest) SetOwnerId(v int64) *QueryContentShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryContentShrinkRequest) SetRecallWindowShrink(v string) *QueryContentShrinkRequest {
	s.RecallWindowShrink = &v
	return s
}

func (s *QueryContentShrinkRequest) SetRegionId(v string) *QueryContentShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryContentShrinkRequest) SetRerankFactor(v float64) *QueryContentShrinkRequest {
	s.RerankFactor = &v
	return s
}

func (s *QueryContentShrinkRequest) SetRerankModelShrink(v string) *QueryContentShrinkRequest {
	s.RerankModelShrink = &v
	return s
}

func (s *QueryContentShrinkRequest) SetTopK(v int32) *QueryContentShrinkRequest {
	s.TopK = &v
	return s
}

func (s *QueryContentShrinkRequest) SetUrlExpiration(v string) *QueryContentShrinkRequest {
	s.UrlExpiration = &v
	return s
}

func (s *QueryContentShrinkRequest) SetUseFullTextRetrieval(v bool) *QueryContentShrinkRequest {
	s.UseFullTextRetrieval = &v
	return s
}

func (s *QueryContentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
