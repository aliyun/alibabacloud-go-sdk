// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *RetrieveKnowledgeBaseRequest
	GetFilter() *string
	SetHybridSearch(v string) *RetrieveKnowledgeBaseRequest
	GetHybridSearch() *string
	SetHybridSearchArgs(v string) *RetrieveKnowledgeBaseRequest
	GetHybridSearchArgs() *string
	SetIncludeMetadataFields(v string) *RetrieveKnowledgeBaseRequest
	GetIncludeMetadataFields() *string
	SetIncludeVector(v bool) *RetrieveKnowledgeBaseRequest
	GetIncludeVector() *bool
	SetKbUuid(v string) *RetrieveKnowledgeBaseRequest
	GetKbUuid() *string
	SetMetrics(v string) *RetrieveKnowledgeBaseRequest
	GetMetrics() *string
	SetOffset(v int32) *RetrieveKnowledgeBaseRequest
	GetOffset() *int32
	SetOrderBy(v string) *RetrieveKnowledgeBaseRequest
	GetOrderBy() *string
	SetQuery(v string) *RetrieveKnowledgeBaseRequest
	GetQuery() *string
	SetRecallWindow(v string) *RetrieveKnowledgeBaseRequest
	GetRecallWindow() *string
	SetRerankFactor(v float64) *RetrieveKnowledgeBaseRequest
	GetRerankFactor() *float64
	SetTopK(v int32) *RetrieveKnowledgeBaseRequest
	GetTopK() *int32
}

type RetrieveKnowledgeBaseRequest struct {
	// A filter for the data, specified as a SQL `WHERE` clause.
	//
	// example:
	//
	// title = \\"test\\" AND name like \\"test%\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The hybrid search algorithm. If this parameter is not set, the system directly compares and ranks the scores from the dense vector and full-text searches.
	//
	// Valid values:
	//
	// - `RRF`: Reciprocal Rank Fusion. This method uses a parameter `k` to control the fusion effect. For more information, see the `HybridSearchArgs` configuration.
	//
	// - `Weight`: Weighted ranking. This method applies weights to the vector and full-text search scores before ranking. For more information, see the `HybridSearchArgs` configuration.
	//
	// - `Cascaded`: Performs a full-text search first, followed by a vector search on the results of the full-text search.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// Parameters for the specified `HybridSearch` algorithm. Both `RRF` and `Weight` are supported. You can use the `HybridPathsSetting` object to specify the retrieval paths: dense vector (`dense`), sparse vector (`sparse`), and full-text search (`fulltext`). If this object is not provided, the default retrieval paths are `dense` and `fulltext`.
	//
	// - `RRF`: Specifies the constant `k` in the scoring formula `1/(k+rank_i)`. The value of `k` must be an integer greater than 1. The format is as follows:
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
	// - `Weight`:
	//
	//   - Two-path recall (do not specify `HybridPathsSetting`; specify only `alpha`):
	//
	//     - The score is calculated using the formula: `alpha 	- dense_score + (1-alpha) 	- fulltext_score`. The `alpha` parameter balances the scores from the dense vector and full-text searches. Its value must be in the range [0, 1], where 0 relies solely on full-text search, and 1 relies solely on dense vector search.
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
	// - Three-path recall:
	//
	//   - The score is calculated using the formula: `normalized_dense 	- dense_score + normalized_sparse 	- sparse_score + normalized_fulltext 	- fulltext_score`. The `dense`, `sparse`, and `fulltext` parameters are the weights for the dense vector, sparse vector, and full-text searches, respectively. Their values must be 0 or greater. The system automatically normalizes the weights to sum to 1 (for example, `normalized_x = x / (dense + sparse + fulltext)`).
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
	//
	// example:
	//
	// { \\"Weight\\": { \\"alpha\\": 0.5 } }
	HybridSearchArgs *string `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// The metadata fields to return, separated by commas. By default, no metadata fields are returned.
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
	// false
	IncludeVector *bool `json:"IncludeVector,omitempty" xml:"IncludeVector,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The distance metric for retrieval. If unspecified, this defaults to the metric configured for the knowledge base. Only set this parameter if you have specific requirements.
	//
	// Valid values:
	//
	// - `l2`: Euclidean distance.
	//
	// - `ip`: Inner product.
	//
	// - `cosine`: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The offset for pagination.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The field to use for sorting the results. By default, this parameter is empty.
	//
	// The field must be a metadata field or a default table field, such as `id`. Supported formats include:
	//
	// You can specify a single field (for example, `chunk_id`), multiple comma-separated fields (for example, `block_id, chunk_id`), or fields with descending order (for example, `block_id DESC, chunk_id DESC`).
	//
	// example:
	//
	// created_at
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The query text.
	//
	// This parameter is required.
	//
	// example:
	//
	// What is GraphRAG?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The recall window. If specified, this parameter expands the context of the retrieved results. The format is a two-element array `[A, B]`, where `-10 <= A <= 0` and `0 <= B <= 10`.
	//
	// > - Recommended when document chunks are highly fragmented, which might cause context loss during retrieval.
	//
	// >
	//
	// > - Reranking occurs before windowing is applied.
	//
	// example:
	//
	// [-5,5]
	RecallWindow *string `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty"`
	// The factor used to rerank vector search results. The value must be in the range (1, 5].
	//
	// > - Reranking may be slow if document chunks are sparse.
	//
	// >
	//
	// > - The number of items to rerank, calculated as `ceil(TopK 	- RerankFactor)`, should not exceed 50.
	//
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The number of top-ranked results to return.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s RetrieveKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseRequest) GetFilter() *string {
	return s.Filter
}

func (s *RetrieveKnowledgeBaseRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *RetrieveKnowledgeBaseRequest) GetHybridSearchArgs() *string {
	return s.HybridSearchArgs
}

func (s *RetrieveKnowledgeBaseRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *RetrieveKnowledgeBaseRequest) GetIncludeVector() *bool {
	return s.IncludeVector
}

func (s *RetrieveKnowledgeBaseRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *RetrieveKnowledgeBaseRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *RetrieveKnowledgeBaseRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *RetrieveKnowledgeBaseRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *RetrieveKnowledgeBaseRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveKnowledgeBaseRequest) GetRecallWindow() *string {
	return s.RecallWindow
}

func (s *RetrieveKnowledgeBaseRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *RetrieveKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RetrieveKnowledgeBaseRequest) SetFilter(v string) *RetrieveKnowledgeBaseRequest {
	s.Filter = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetHybridSearch(v string) *RetrieveKnowledgeBaseRequest {
	s.HybridSearch = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetHybridSearchArgs(v string) *RetrieveKnowledgeBaseRequest {
	s.HybridSearchArgs = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetIncludeMetadataFields(v string) *RetrieveKnowledgeBaseRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetIncludeVector(v bool) *RetrieveKnowledgeBaseRequest {
	s.IncludeVector = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetKbUuid(v string) *RetrieveKnowledgeBaseRequest {
	s.KbUuid = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetMetrics(v string) *RetrieveKnowledgeBaseRequest {
	s.Metrics = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetOffset(v int32) *RetrieveKnowledgeBaseRequest {
	s.Offset = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetOrderBy(v string) *RetrieveKnowledgeBaseRequest {
	s.OrderBy = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetQuery(v string) *RetrieveKnowledgeBaseRequest {
	s.Query = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetRecallWindow(v string) *RetrieveKnowledgeBaseRequest {
	s.RecallWindow = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetRerankFactor(v float64) *RetrieveKnowledgeBaseRequest {
	s.RerankFactor = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetTopK(v int32) *RetrieveKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
