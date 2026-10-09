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
	// > The document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to view the created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The text content used for retrieval.
	//
	// example:
	//
	// What is AnalyticDB for PostgreSQL?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// The source file name of the image to search in image-to-image search scenarios.
	//
	// > The image file must have a file name extension. Supported image file name extensions: bmp, jpg, jpeg, png, and tiff.
	//
	// example:
	//
	// test.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The publicly accessible URL of the image file in image-to-image search scenarios.
	//
	// > The image file must have a file name extension. Supported image file name extensions: bmp, jpg, jpeg, png, and tiff.
	//
	// example:
	//
	// https://xx/myImage.jpg
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The filter conditions for the data to query, formatted as an SQL WHERE clause. This is an expression that returns a Boolean value (true or false). The conditions can be simple comparison operators such as equal to (=), not equal to (<> or !=), greater than (>), less than (<), greater than or equal to (>=), and less than or equal to (<=). They can also be more complex expressions combined with logical operators (AND, OR, NOT), or conditions using keywords such as IN, BETWEEN, and LIKE.
	//
	// >
	//
	// > - For detailed syntax, refer to https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/
	//
	// example:
	//
	// title = \\"test\\" AND name like \\"test%\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to enable knowledge graph enhancement. Default value: false.
	//
	// example:
	//
	// false
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// The knowledge graph retrieval parameters.
	GraphSearchArgsShrink *string `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty"`
	// The multi-channel recall algorithm. Default value: empty. If this parameter is empty, the scores of dense vectors and full text are directly compared, and sorting is performed.
	//
	// Valid values:
	//
	// - RRF: Reciprocal Rank Fusion. A parameter k is used to control the fusion effect. For more information, see the HybridSearchArgs configuration.
	//
	// - Weight: Weighted sorting. Parameters are used to control the score weights of vectors and full text before sorting. For more information, see the HybridSearchArgs configuration.
	//
	// - Cascaded: Full-text retrieval is performed first, and then vector retrieval is performed based on the full-text retrieval results.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The algorithm parameters for multi-channel recall. RRF and Weight are supported. You can use HybridPathsSetting to specify the recall of dense vectors (dense), sparse vectors (sparse), and full-text retrieval (fulltext). If this value is empty, dense vectors (dense) and full-text retrieval (fulltext) are recalled by default.
	//
	// - RRF: Specifies the constant k in the score calculation formula `1/(k+rank_i)`. The value must be a positive integer greater than 1. Format:
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
	//    - Dual-channel recall (HybridPathsSetting is not specified, and only alpha is specified):
	//
	//       - Formula: alpha 	- dense_score + (1-alpha) 	- fulltext_score. The alpha parameter indicates the score weights of dense vectors and full-text retrieval. Valid values: 0 to 1. A value of 0 indicates full-text retrieval only, and a value of 1 indicates dense vectors only:
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
	//   - Three-channel recall pattern:
	//
	//      - Formula: normalized_dense 	- dense_score + normalized_sparse 	- sparse_score + normalized_fulltext 	- fulltext_score. The dense, sparse, and fulltext parameters represent the weights of dense vectors, sparse vectors, and full-text retrieval, respectively. Valid values: greater than or equal to 0. The system automatically performs normalization on the weights to 0 to 1 (that is, normalized_x = x / (dense + sparse + fulltext)).
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
	// Specifies whether to synchronously return the URL of the document. By default, the URL is not returned.
	//
	// example:
	//
	// false
	IncludeFileUrl *bool `json:"IncludeFileUrl,omitempty" xml:"IncludeFileUrl,omitempty"`
	// The metadata fields to return. Default value: empty. Separate multiple fields with commas (,).
	//
	// example:
	//
	// title,page
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
	// Specifies whether to return vectors. Default value: false.
	//
	// > - **false**: Vectors are not returned.
	//
	// > - **true**: Vectors are returned.
	//
	// example:
	//
	// true
	IncludeVector *bool `json:"IncludeVector,omitempty" xml:"IncludeVector,omitempty"`
	// The similarity algorithm used during retrieval. If this value is empty, the algorithm specified when the knowledge base is created is used. You do not need to set this parameter unless you have special requirements.
	//
	// > Valid values:
	//
	// > - **l2**: Euclidean distance.
	//
	// > - **ip**: Dot product (inner product) distance.
	//
	// > - **cosine**: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view the namespace list.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > This value is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The offset used for a paged query.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The field based on which sorting is performed. Default value: empty. The field must belong to the metadata or be a default field in the table, such as id. Supported formats: a single field, such as chunk_id; multiple fields separated by commas (,), such as block_id, chunk_id; and reverse order, such as block_id DESC, chunk_id DESC.
	//
	// example:
	//
	// created_at
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The recall window. When this value is not empty, the context of the retrieval results is additionally returned. The format is an array of two elements: List<A, B>, where -10 <= A <= 0 and 0 <= B <= 10.
	//
	// > - Use this parameter when documents are split into excessively small chunks and retrieval may lose context information.
	//
	// > - Reranking takes precedence over windowing. That is, reranking is performed before windowing.
	RecallWindowShrink *string `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reranking factor. When this value is not empty, the vector retrieval results are reranked. Valid values: 1 < RerankFactor <= 5.
	//
	// > - When document chunks are sparse, reranking is slow.
	//
	// > - The number of reranked results (TopK × Factor, rounded up) should not exceed 50.
	//
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The reranking model parameters.
	RerankModelShrink *string `json:"RerankModel,omitempty" xml:"RerankModel,omitempty"`
	// The number of top results to return.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The validity period of the returned image URL.
	//
	// > Valid values:
	//
	// > - The unit can be seconds (s) or days (d). For example, 300s indicates a validity period of 300 seconds, and 60d indicates a validity period of 60 days.
	//
	// > - Valid values: 60s to 365d.
	//
	// > - Default value: 7200s, which is 2 hours.
	//
	// example:
	//
	// 7200s
	UrlExpiration *string `json:"UrlExpiration,omitempty" xml:"UrlExpiration,omitempty"`
	// **[Deprecated]*	- Specifies whether to use full-text retrieval (dual-channel recall). Default value: false, which indicates that only vector retrieval is used.
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
