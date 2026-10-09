// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iQueryContentAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollection(v string) *QueryContentAdvanceRequest
	GetCollection() *string
	SetContent(v string) *QueryContentAdvanceRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryContentAdvanceRequest
	GetDBInstanceId() *string
	SetFileName(v string) *QueryContentAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *QueryContentAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFilter(v string) *QueryContentAdvanceRequest
	GetFilter() *string
	SetGraphEnhance(v bool) *QueryContentAdvanceRequest
	GetGraphEnhance() *bool
	SetGraphSearchArgs(v *QueryContentAdvanceRequestGraphSearchArgs) *QueryContentAdvanceRequest
	GetGraphSearchArgs() *QueryContentAdvanceRequestGraphSearchArgs
	SetHybridSearch(v string) *QueryContentAdvanceRequest
	GetHybridSearch() *string
	SetHybridSearchArgs(v map[string]map[string]interface{}) *QueryContentAdvanceRequest
	GetHybridSearchArgs() map[string]map[string]interface{}
	SetIncludeFileUrl(v bool) *QueryContentAdvanceRequest
	GetIncludeFileUrl() *bool
	SetIncludeMetadataFields(v string) *QueryContentAdvanceRequest
	GetIncludeMetadataFields() *string
	SetIncludeVector(v bool) *QueryContentAdvanceRequest
	GetIncludeVector() *bool
	SetMetrics(v string) *QueryContentAdvanceRequest
	GetMetrics() *string
	SetNamespace(v string) *QueryContentAdvanceRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *QueryContentAdvanceRequest
	GetNamespacePassword() *string
	SetOffset(v int32) *QueryContentAdvanceRequest
	GetOffset() *int32
	SetOrderBy(v string) *QueryContentAdvanceRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *QueryContentAdvanceRequest
	GetOwnerId() *int64
	SetRecallWindow(v []*int32) *QueryContentAdvanceRequest
	GetRecallWindow() []*int32
	SetRegionId(v string) *QueryContentAdvanceRequest
	GetRegionId() *string
	SetRerankFactor(v float64) *QueryContentAdvanceRequest
	GetRerankFactor() *float64
	SetRerankModel(v *QueryContentAdvanceRequestRerankModel) *QueryContentAdvanceRequest
	GetRerankModel() *QueryContentAdvanceRequestRerankModel
	SetTopK(v int32) *QueryContentAdvanceRequest
	GetTopK() *int32
	SetUrlExpiration(v string) *QueryContentAdvanceRequest
	GetUrlExpiration() *string
	SetUseFullTextRetrieval(v bool) *QueryContentAdvanceRequest
	GetUseFullTextRetrieval() *bool
}

type QueryContentAdvanceRequest struct {
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
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
	GraphSearchArgs *QueryContentAdvanceRequestGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
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
	HybridSearchArgs map[string]map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
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
	RecallWindow []*int32 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
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
	RerankModel *QueryContentAdvanceRequestRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
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

func (s QueryContentAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *QueryContentAdvanceRequest) GetCollection() *string {
	return s.Collection
}

func (s *QueryContentAdvanceRequest) GetContent() *string {
	return s.Content
}

func (s *QueryContentAdvanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryContentAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *QueryContentAdvanceRequest) GetFilter() *string {
	return s.Filter
}

func (s *QueryContentAdvanceRequest) GetGraphEnhance() *bool {
	return s.GraphEnhance
}

func (s *QueryContentAdvanceRequest) GetGraphSearchArgs() *QueryContentAdvanceRequestGraphSearchArgs {
	return s.GraphSearchArgs
}

func (s *QueryContentAdvanceRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *QueryContentAdvanceRequest) GetHybridSearchArgs() map[string]map[string]interface{} {
	return s.HybridSearchArgs
}

func (s *QueryContentAdvanceRequest) GetIncludeFileUrl() *bool {
	return s.IncludeFileUrl
}

func (s *QueryContentAdvanceRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *QueryContentAdvanceRequest) GetIncludeVector() *bool {
	return s.IncludeVector
}

func (s *QueryContentAdvanceRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryContentAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryContentAdvanceRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *QueryContentAdvanceRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *QueryContentAdvanceRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryContentAdvanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryContentAdvanceRequest) GetRecallWindow() []*int32 {
	return s.RecallWindow
}

func (s *QueryContentAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryContentAdvanceRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *QueryContentAdvanceRequest) GetRerankModel() *QueryContentAdvanceRequestRerankModel {
	return s.RerankModel
}

func (s *QueryContentAdvanceRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *QueryContentAdvanceRequest) GetUrlExpiration() *string {
	return s.UrlExpiration
}

func (s *QueryContentAdvanceRequest) GetUseFullTextRetrieval() *bool {
	return s.UseFullTextRetrieval
}

func (s *QueryContentAdvanceRequest) SetCollection(v string) *QueryContentAdvanceRequest {
	s.Collection = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetContent(v string) *QueryContentAdvanceRequest {
	s.Content = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetDBInstanceId(v string) *QueryContentAdvanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetFileName(v string) *QueryContentAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetFileUrlObject(v io.Reader) *QueryContentAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *QueryContentAdvanceRequest) SetFilter(v string) *QueryContentAdvanceRequest {
	s.Filter = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetGraphEnhance(v bool) *QueryContentAdvanceRequest {
	s.GraphEnhance = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetGraphSearchArgs(v *QueryContentAdvanceRequestGraphSearchArgs) *QueryContentAdvanceRequest {
	s.GraphSearchArgs = v
	return s
}

func (s *QueryContentAdvanceRequest) SetHybridSearch(v string) *QueryContentAdvanceRequest {
	s.HybridSearch = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetHybridSearchArgs(v map[string]map[string]interface{}) *QueryContentAdvanceRequest {
	s.HybridSearchArgs = v
	return s
}

func (s *QueryContentAdvanceRequest) SetIncludeFileUrl(v bool) *QueryContentAdvanceRequest {
	s.IncludeFileUrl = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetIncludeMetadataFields(v string) *QueryContentAdvanceRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetIncludeVector(v bool) *QueryContentAdvanceRequest {
	s.IncludeVector = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetMetrics(v string) *QueryContentAdvanceRequest {
	s.Metrics = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetNamespace(v string) *QueryContentAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetNamespacePassword(v string) *QueryContentAdvanceRequest {
	s.NamespacePassword = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetOffset(v int32) *QueryContentAdvanceRequest {
	s.Offset = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetOrderBy(v string) *QueryContentAdvanceRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetOwnerId(v int64) *QueryContentAdvanceRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetRecallWindow(v []*int32) *QueryContentAdvanceRequest {
	s.RecallWindow = v
	return s
}

func (s *QueryContentAdvanceRequest) SetRegionId(v string) *QueryContentAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetRerankFactor(v float64) *QueryContentAdvanceRequest {
	s.RerankFactor = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetRerankModel(v *QueryContentAdvanceRequestRerankModel) *QueryContentAdvanceRequest {
	s.RerankModel = v
	return s
}

func (s *QueryContentAdvanceRequest) SetTopK(v int32) *QueryContentAdvanceRequest {
	s.TopK = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetUrlExpiration(v string) *QueryContentAdvanceRequest {
	s.UrlExpiration = &v
	return s
}

func (s *QueryContentAdvanceRequest) SetUseFullTextRetrieval(v bool) *QueryContentAdvanceRequest {
	s.UseFullTextRetrieval = &v
	return s
}

func (s *QueryContentAdvanceRequest) Validate() error {
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

type QueryContentAdvanceRequestGraphSearchArgs struct {
	// The number of top entities and relationship edges to return. Default value: 60.
	//
	// example:
	//
	// 60
	GraphTopK *int32 `json:"GraphTopK,omitempty" xml:"GraphTopK,omitempty"`
}

func (s QueryContentAdvanceRequestGraphSearchArgs) String() string {
	return dara.Prettify(s)
}

func (s QueryContentAdvanceRequestGraphSearchArgs) GoString() string {
	return s.String()
}

func (s *QueryContentAdvanceRequestGraphSearchArgs) GetGraphTopK() *int32 {
	return s.GraphTopK
}

func (s *QueryContentAdvanceRequestGraphSearchArgs) SetGraphTopK(v int32) *QueryContentAdvanceRequestGraphSearchArgs {
	s.GraphTopK = &v
	return s
}

func (s *QueryContentAdvanceRequestGraphSearchArgs) Validate() error {
	return dara.Validate(s)
}

type QueryContentAdvanceRequestRerankModel struct {
	// This parameter can be set when RerankModel.Name is set to qwen3-rerank. You can add a custom sorting task description to guide the model to adopt different sorting strategies.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the reranking model. Valid values: qwen3-rerank and gte-rerank-v2.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The metadata fields that participate in reranking. Separate multiple fields with commas (,). By default, only the document content (content) is used for reranking. The field names must be metadata defined in the collection.
	RerankMetadataFields *string `json:"RerankMetadataFields,omitempty" xml:"RerankMetadataFields,omitempty"`
}

func (s QueryContentAdvanceRequestRerankModel) String() string {
	return dara.Prettify(s)
}

func (s QueryContentAdvanceRequestRerankModel) GoString() string {
	return s.String()
}

func (s *QueryContentAdvanceRequestRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *QueryContentAdvanceRequestRerankModel) GetName() *string {
	return s.Name
}

func (s *QueryContentAdvanceRequestRerankModel) GetRerankMetadataFields() *string {
	return s.RerankMetadataFields
}

func (s *QueryContentAdvanceRequestRerankModel) SetInstruct(v string) *QueryContentAdvanceRequestRerankModel {
	s.Instruct = &v
	return s
}

func (s *QueryContentAdvanceRequestRerankModel) SetName(v string) *QueryContentAdvanceRequestRerankModel {
	s.Name = &v
	return s
}

func (s *QueryContentAdvanceRequestRerankModel) SetRerankMetadataFields(v string) *QueryContentAdvanceRequestRerankModel {
	s.RerankMetadataFields = &v
	return s
}

func (s *QueryContentAdvanceRequestRerankModel) Validate() error {
	return dara.Validate(s)
}
