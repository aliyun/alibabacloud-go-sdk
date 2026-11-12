// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerankShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RerankShrinkRequest
	GetDBInstanceId() *string
	SetDocumentsShrink(v string) *RerankShrinkRequest
	GetDocumentsShrink() *string
	SetInstruct(v string) *RerankShrinkRequest
	GetInstruct() *string
	SetMaxChunksPerDoc(v int32) *RerankShrinkRequest
	GetMaxChunksPerDoc() *int32
	SetModel(v string) *RerankShrinkRequest
	GetModel() *string
	SetOwnerId(v int64) *RerankShrinkRequest
	GetOwnerId() *int64
	SetQuery(v string) *RerankShrinkRequest
	GetQuery() *string
	SetRegionId(v string) *RerankShrinkRequest
	GetRegionId() *string
	SetReturnDocuments(v bool) *RerankShrinkRequest
	GetReturnDocuments() *bool
	SetTopK(v int32) *RerankShrinkRequest
	GetTopK() *int32
}

type RerankShrinkRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to view the details of all AnalyticDB for PostgreSQL instances in the specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of documents to be reranked.
	//
	// example:
	//
	// ADBPG is an Alibaba Cloud OLAP database.
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// The description of the custom ranking task type. This parameter can be set when RerankModel.Name is qwen3-rerank. It guides the model to use different ranking strategies.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The maximum number of chunks allowed when the text exceeds the model window. Valid values:
	//
	// - bge-reranker-v2-m3: 10. This is the default value.
	//
	// - bge-reranker-v2-minicpm-layerwise: 5. This is the default value.
	//
	// > Chunking example
	//
	// > - If you use the bge-reranker-v2-minicpm-layerwise model, the maximum single inference window is 2,048 tokens. If the query is 48 tokens and a single document parameter contains 9,000 tokens, the document is split into chunks: tokens 1 to 2,000 form the first chunk, tokens 2,001 to 4,000 form the second chunk, and so on. If the number of chunks exceeds the value of MaxChunksPerDoc, the subsequent text is discarded.
	//
	// example:
	//
	// 10
	MaxChunksPerDoc *int32 `json:"MaxChunksPerDoc,omitempty" xml:"MaxChunksPerDoc,omitempty"`
	// The rerank model. Valid values:
	//
	// - bge-reranker-v2-m3: provides good performance and supports up to 8,192 tokens for a single inference. Text that exceeds this limit is split into chunks, which may degrade performance. This is the default value.
	//
	// - bge-reranker-v2-minicpm-layerwise: provides better performance than bge-reranker-v2-m3 and supports up to 2,048 tokens for a single inference. Text that exceeds this limit is split into chunks, which may degrade performance.
	//
	// example:
	//
	// bge-reranker-v2-m3
	Model   *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The query statement for reranking.
	//
	// example:
	//
	// What is ADBPG?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return the document text. If this parameter is set to false, the document text is not returned. Only the indexes and rerank scores of the sorted documents are returned.
	//
	// example:
	//
	// false
	ReturnDocuments *bool `json:"ReturnDocuments,omitempty" xml:"ReturnDocuments,omitempty"`
	// The number of most relevant documents to return.
	//
	// example:
	//
	// 3
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s RerankShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RerankShrinkRequest) GoString() string {
	return s.String()
}

func (s *RerankShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RerankShrinkRequest) GetDocumentsShrink() *string {
	return s.DocumentsShrink
}

func (s *RerankShrinkRequest) GetInstruct() *string {
	return s.Instruct
}

func (s *RerankShrinkRequest) GetMaxChunksPerDoc() *int32 {
	return s.MaxChunksPerDoc
}

func (s *RerankShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *RerankShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RerankShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *RerankShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RerankShrinkRequest) GetReturnDocuments() *bool {
	return s.ReturnDocuments
}

func (s *RerankShrinkRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RerankShrinkRequest) SetDBInstanceId(v string) *RerankShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RerankShrinkRequest) SetDocumentsShrink(v string) *RerankShrinkRequest {
	s.DocumentsShrink = &v
	return s
}

func (s *RerankShrinkRequest) SetInstruct(v string) *RerankShrinkRequest {
	s.Instruct = &v
	return s
}

func (s *RerankShrinkRequest) SetMaxChunksPerDoc(v int32) *RerankShrinkRequest {
	s.MaxChunksPerDoc = &v
	return s
}

func (s *RerankShrinkRequest) SetModel(v string) *RerankShrinkRequest {
	s.Model = &v
	return s
}

func (s *RerankShrinkRequest) SetOwnerId(v int64) *RerankShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *RerankShrinkRequest) SetQuery(v string) *RerankShrinkRequest {
	s.Query = &v
	return s
}

func (s *RerankShrinkRequest) SetRegionId(v string) *RerankShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RerankShrinkRequest) SetReturnDocuments(v bool) *RerankShrinkRequest {
	s.ReturnDocuments = &v
	return s
}

func (s *RerankShrinkRequest) SetTopK(v int32) *RerankShrinkRequest {
	s.TopK = &v
	return s
}

func (s *RerankShrinkRequest) Validate() error {
	return dara.Validate(s)
}
