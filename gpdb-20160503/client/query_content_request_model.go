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
	SetOwnerId(v int64) *QueryContentRequest
	GetOwnerId() *int64
	SetRecallWindow(v []*int32) *QueryContentRequest
	GetRecallWindow() []*int32
	SetRegionId(v string) *QueryContentRequest
	GetRegionId() *string
	SetRerankFactor(v float64) *QueryContentRequest
	GetRerankFactor() *float64
	SetTopK(v int32) *QueryContentRequest
	GetTopK() *int32
	SetUrlExpiration(v string) *QueryContentRequest
	GetUrlExpiration() *string
	SetUseFullTextRetrieval(v bool) *QueryContentRequest
	GetUseFullTextRetrieval() *bool
}

type QueryContentRequest struct {
	// Document collection name.
	//
	// > Created by the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) API. You can use the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) API to view the list of created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB for PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// In image search scenarios, the source file name of the image to be searched.
	//
	// > The image file must have a file extension. Currently supported image extensions: bmp, jpg, jpeg, png, tiff.
	//
	// example:
	//
	// test.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// In image search scenarios, the publicly accessible URL of the image file.
	//
	// > The image file must have a file extension. Currently supported image extensions: bmp, jpg, jpeg, png, tiff.
	//
	// example:
	//
	// https://xx/myImage.jpg
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Filter condition for the data to be queried, in SQL WHERE format. It is an expression that returns a boolean value (true or false). The conditions can be simple comparison operators such as equal (=), not equal (<> or !=), greater than (>), less than (<), greater than or equal to (>=), less than or equal to (<=), or more complex expressions combined with logical operators (AND, OR, NOT), and conditions using keywords like IN, BETWEEN, LIKE, etc.
	//
	// >
	//
	// > - For detailed syntax, refer to: https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/
	//
	// example:
	//
	// title = \\"test\\" AND name like \\"test%\\"
	Filter          *string                             `json:"Filter,omitempty" xml:"Filter,omitempty"`
	GraphEnhance    *bool                               `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	GraphSearchArgs *QueryContentRequestGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// Dual recall algorithm, default is empty (i.e., directly compare and sort the scores of vectors and full text).
	//
	// Available values:
	//
	// - RRF: Reciprocal rank fusion, with a parameter k controlling the fusion effect. See HybridSearchArgs configuration for details;
	//
	// - Weight: Weighted ranking, using a parameter alpha to control the weight of vector and full-text scores, then sorting. See HybridSearchArgs configuration for details;
	//
	// - Cascaded: Perform full-text retrieval first, then vector retrieval on top of it;
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
	// Specifies whether to return the URL of the document. Default value: false.
	//
	// example:
	//
	// false
	IncludeFileUrl *bool `json:"IncludeFileUrl,omitempty" xml:"IncludeFileUrl,omitempty"`
	// The metadata fields to be returned. Separate multiple fields with commas (,). This parameter is empty by default.
	//
	// example:
	//
	// title,page
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
	// Whether to return vectors. Default is false.
	//
	// > - **false**: Do not return vectors.
	//
	// > - **true**: Return vectors.
	//
	// example:
	//
	// true
	IncludeVector *bool `json:"IncludeVector,omitempty" xml:"IncludeVector,omitempty"`
	// Similarity algorithm used during retrieval. If this value is empty, the algorithm specified at the time of knowledge base creation is used. It is recommended not to set this unless there is a specific need.
	//
	// > Value description:
	//
	// > - **l2**: Euclidean distance.
	//
	// > - **ip**: Inner product (dot product) distance.
	//
	// > - **cosine**: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// Namespace, default is public.
	//
	// > You can create a namespace using the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API and view the list of namespaces using the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) API.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Password for the namespace.
	//
	// > This value is specified in the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Recall window. When this value is not empty, it adds context to the returned search results. The format is an array of 2 elements: List<A, B>, where -10 <= A <= 0 and 0 <= B <= 10.
	//
	// > - Recommended when documents are fragmented and retrieval may lose contextual information.
	//
	// > - Re-ranking takes precedence over windowing, i.e., re-rank first, then apply windowing.
	RecallWindow []*int32 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Re-ranking factor. When this value is not empty, it will re-rank the vector search results. The value range is 1 < RerankFactor <= 5.
	//
	// > - Re-ranking is slower when documents are sparsely split.
	//
	// > - It is recommended that the re-ranked count (TopK 	- Factor, rounded up) does not exceed 50.
	//
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// Set the number of top results to return.
	//
	// example:
	//
	// 10
	TopK          *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
	UrlExpiration *string `json:"UrlExpiration,omitempty" xml:"UrlExpiration,omitempty"`
	// Whether to use full-text retrieval (dual recall). Default is false, which means only vector retrieval is used.
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
	return dara.Validate(s)
}

type QueryContentRequestGraphSearchArgs struct {
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
