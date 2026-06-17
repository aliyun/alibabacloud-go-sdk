// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *QueryContentRequest
	GetCollection() *string
	SetContent(v string) *QueryContentRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryContentRequest
	GetDBInstanceId() *string
	SetFileName(v string) *QueryContentRequest
	GetFileName() *string
	SetFileUrl(v string) *QueryContentRequest
	GetFileUrl() *string
	SetFilter(v string) *QueryContentRequest
	GetFilter() *string
	SetGraphEnhance(v bool) *QueryContentRequest
	GetGraphEnhance() *bool
	SetGraphSearchArgs(v *QueryContentRequestGraphSearchArgs) *QueryContentRequest
	GetGraphSearchArgs() *QueryContentRequestGraphSearchArgs
	SetHybridSearch(v string) *QueryContentRequest
	GetHybridSearch() *string
	SetHybridSearchArgs(v map[string]map[string]interface{}) *QueryContentRequest
	GetHybridSearchArgs() map[string]map[string]interface{}
	SetIncludeFileUrl(v bool) *QueryContentRequest
	GetIncludeFileUrl() *bool
	SetIncludeMetadataFields(v string) *QueryContentRequest
	GetIncludeMetadataFields() *string
	SetIncludeVector(v bool) *QueryContentRequest
	GetIncludeVector() *bool
	SetMetrics(v string) *QueryContentRequest
	GetMetrics() *string
	SetNamespace(v string) *QueryContentRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *QueryContentRequest
	GetNamespacePassword() *string
	SetOffset(v int32) *QueryContentRequest
	GetOffset() *int32
	SetOrderBy(v string) *QueryContentRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *QueryContentRequest
	GetOwnerId() *int64
	SetRecallWindow(v []*int32) *QueryContentRequest
	GetRecallWindow() []*int32
	SetRegionId(v string) *QueryContentRequest
	GetRegionId() *string
	SetRerankFactor(v float64) *QueryContentRequest
	GetRerankFactor() *float64
	SetRerankModel(v *QueryContentRequestRerankModel) *QueryContentRequest
	GetRerankModel() *QueryContentRequestRerankModel
	SetTopK(v int32) *QueryContentRequest
	GetTopK() *int32
	SetUrlExpiration(v string) *QueryContentRequest
	GetUrlExpiration() *string
	SetUseFullTextRetrieval(v bool) *QueryContentRequest
	GetUseFullTextRetrieval() *bool
}

type QueryContentRequest struct {
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
	GraphSearchArgs *QueryContentRequestGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
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
	HybridSearchArgs map[string]map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
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
	RecallWindow []*int32 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
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
	RerankModel *QueryContentRequestRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
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

func (s QueryContentRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContentRequest) GoString() string {
	return s.String()
}

func (s *QueryContentRequest) GetCollection() *string {
	return s.Collection
}

func (s *QueryContentRequest) GetContent() *string {
	return s.Content
}

func (s *QueryContentRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryContentRequest) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *QueryContentRequest) GetFilter() *string {
	return s.Filter
}

func (s *QueryContentRequest) GetGraphEnhance() *bool {
	return s.GraphEnhance
}

func (s *QueryContentRequest) GetGraphSearchArgs() *QueryContentRequestGraphSearchArgs {
	return s.GraphSearchArgs
}

func (s *QueryContentRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *QueryContentRequest) GetHybridSearchArgs() map[string]map[string]interface{} {
	return s.HybridSearchArgs
}

func (s *QueryContentRequest) GetIncludeFileUrl() *bool {
	return s.IncludeFileUrl
}

func (s *QueryContentRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *QueryContentRequest) GetIncludeVector() *bool {
	return s.IncludeVector
}

func (s *QueryContentRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryContentRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryContentRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *QueryContentRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *QueryContentRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryContentRequest) GetRecallWindow() []*int32 {
	return s.RecallWindow
}

func (s *QueryContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryContentRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *QueryContentRequest) GetRerankModel() *QueryContentRequestRerankModel {
	return s.RerankModel
}

func (s *QueryContentRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *QueryContentRequest) GetUrlExpiration() *string {
	return s.UrlExpiration
}

func (s *QueryContentRequest) GetUseFullTextRetrieval() *bool {
	return s.UseFullTextRetrieval
}

func (s *QueryContentRequest) SetCollection(v string) *QueryContentRequest {
	s.Collection = &v
	return s
}

func (s *QueryContentRequest) SetContent(v string) *QueryContentRequest {
	s.Content = &v
	return s
}

func (s *QueryContentRequest) SetDBInstanceId(v string) *QueryContentRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryContentRequest) SetFileName(v string) *QueryContentRequest {
	s.FileName = &v
	return s
}

func (s *QueryContentRequest) SetFileUrl(v string) *QueryContentRequest {
	s.FileUrl = &v
	return s
}

func (s *QueryContentRequest) SetFilter(v string) *QueryContentRequest {
	s.Filter = &v
	return s
}

func (s *QueryContentRequest) SetGraphEnhance(v bool) *QueryContentRequest {
	s.GraphEnhance = &v
	return s
}

func (s *QueryContentRequest) SetGraphSearchArgs(v *QueryContentRequestGraphSearchArgs) *QueryContentRequest {
	s.GraphSearchArgs = v
	return s
}

func (s *QueryContentRequest) SetHybridSearch(v string) *QueryContentRequest {
	s.HybridSearch = &v
	return s
}

func (s *QueryContentRequest) SetHybridSearchArgs(v map[string]map[string]interface{}) *QueryContentRequest {
	s.HybridSearchArgs = v
	return s
}

func (s *QueryContentRequest) SetIncludeFileUrl(v bool) *QueryContentRequest {
	s.IncludeFileUrl = &v
	return s
}

func (s *QueryContentRequest) SetIncludeMetadataFields(v string) *QueryContentRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *QueryContentRequest) SetIncludeVector(v bool) *QueryContentRequest {
	s.IncludeVector = &v
	return s
}

func (s *QueryContentRequest) SetMetrics(v string) *QueryContentRequest {
	s.Metrics = &v
	return s
}

func (s *QueryContentRequest) SetNamespace(v string) *QueryContentRequest {
	s.Namespace = &v
	return s
}

func (s *QueryContentRequest) SetNamespacePassword(v string) *QueryContentRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryContentRequest) SetOffset(v int32) *QueryContentRequest {
	s.Offset = &v
	return s
}

func (s *QueryContentRequest) SetOrderBy(v string) *QueryContentRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryContentRequest) SetOwnerId(v int64) *QueryContentRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryContentRequest) SetRecallWindow(v []*int32) *QueryContentRequest {
	s.RecallWindow = v
	return s
}

func (s *QueryContentRequest) SetRegionId(v string) *QueryContentRequest {
	s.RegionId = &v
	return s
}

func (s *QueryContentRequest) SetRerankFactor(v float64) *QueryContentRequest {
	s.RerankFactor = &v
	return s
}

func (s *QueryContentRequest) SetRerankModel(v *QueryContentRequestRerankModel) *QueryContentRequest {
	s.RerankModel = v
	return s
}

func (s *QueryContentRequest) SetTopK(v int32) *QueryContentRequest {
	s.TopK = &v
	return s
}

func (s *QueryContentRequest) SetUrlExpiration(v string) *QueryContentRequest {
	s.UrlExpiration = &v
	return s
}

func (s *QueryContentRequest) SetUseFullTextRetrieval(v bool) *QueryContentRequest {
	s.UseFullTextRetrieval = &v
	return s
}

func (s *QueryContentRequest) Validate() error {
	if s.GraphSearchArgs != nil {
		if err := s.GraphSearchArgs.Validate(); err != nil {
			return err
		}
	}
	if s.RerankModel != nil {
		if err := s.RerankModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryContentRequestGraphSearchArgs struct {
	// The number of top entities and relationship edges to return. The default value is `60`.
	//
	// example:
	//
	// 60
	GraphTopK *int32 `json:"GraphTopK,omitempty" xml:"GraphTopK,omitempty"`
}

func (s QueryContentRequestGraphSearchArgs) String() string {
	return dara.Prettify(s)
}

func (s QueryContentRequestGraphSearchArgs) GoString() string {
	return s.String()
}

func (s *QueryContentRequestGraphSearchArgs) GetGraphTopK() *int32 {
	return s.GraphTopK
}

func (s *QueryContentRequestGraphSearchArgs) SetGraphTopK(v int32) *QueryContentRequestGraphSearchArgs {
	s.GraphTopK = &v
	return s
}

func (s *QueryContentRequestGraphSearchArgs) Validate() error {
	return dara.Validate(s)
}

type QueryContentRequestRerankModel struct {
	// This parameter is available only when `RerankModel.Name` is set to `qwen3-rerank`.
	//
	// You can add a custom task description to guide the model\\"s sorting strategy.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the reranking model. Valid values: `qwen3-rerank`, `gte-rerank-v2`.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryContentRequestRerankModel) String() string {
	return dara.Prettify(s)
}

func (s QueryContentRequestRerankModel) GoString() string {
	return s.String()
}

func (s *QueryContentRequestRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *QueryContentRequestRerankModel) GetName() *string {
	return s.Name
}

func (s *QueryContentRequestRerankModel) SetInstruct(v string) *QueryContentRequestRerankModel {
	s.Instruct = &v
	return s
}

func (s *QueryContentRequestRerankModel) SetName(v string) *QueryContentRequestRerankModel {
	s.Name = &v
	return s
}

func (s *QueryContentRequestRerankModel) Validate() error {
	return dara.Validate(s)
}
