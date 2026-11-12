// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *RerankRequest
	GetDBInstanceId() *string
	SetDocuments(v []*string) *RerankRequest
	GetDocuments() []*string
	SetInstruct(v string) *RerankRequest
	GetInstruct() *string
	SetMaxChunksPerDoc(v int32) *RerankRequest
	GetMaxChunksPerDoc() *int32
	SetModel(v string) *RerankRequest
	GetModel() *string
	SetOwnerId(v int64) *RerankRequest
	GetOwnerId() *int64
	SetQuery(v string) *RerankRequest
	GetQuery() *string
	SetRegionId(v string) *RerankRequest
	GetRegionId() *string
	SetReturnDocuments(v bool) *RerankRequest
	GetReturnDocuments() *bool
	SetTopK(v int32) *RerankRequest
	GetTopK() *int32
}

type RerankRequest struct {
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
	Documents []*string `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
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

func (s RerankRequest) String() string {
	return dara.Prettify(s)
}

func (s RerankRequest) GoString() string {
	return s.String()
}

func (s *RerankRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *RerankRequest) GetDocuments() []*string {
	return s.Documents
}

func (s *RerankRequest) GetInstruct() *string {
	return s.Instruct
}

func (s *RerankRequest) GetMaxChunksPerDoc() *int32 {
	return s.MaxChunksPerDoc
}

func (s *RerankRequest) GetModel() *string {
	return s.Model
}

func (s *RerankRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RerankRequest) GetQuery() *string {
	return s.Query
}

func (s *RerankRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RerankRequest) GetReturnDocuments() *bool {
	return s.ReturnDocuments
}

func (s *RerankRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RerankRequest) SetDBInstanceId(v string) *RerankRequest {
	s.DBInstanceId = &v
	return s
}

func (s *RerankRequest) SetDocuments(v []*string) *RerankRequest {
	s.Documents = v
	return s
}

func (s *RerankRequest) SetInstruct(v string) *RerankRequest {
	s.Instruct = &v
	return s
}

func (s *RerankRequest) SetMaxChunksPerDoc(v int32) *RerankRequest {
	s.MaxChunksPerDoc = &v
	return s
}

func (s *RerankRequest) SetModel(v string) *RerankRequest {
	s.Model = &v
	return s
}

func (s *RerankRequest) SetOwnerId(v int64) *RerankRequest {
	s.OwnerId = &v
	return s
}

func (s *RerankRequest) SetQuery(v string) *RerankRequest {
	s.Query = &v
	return s
}

func (s *RerankRequest) SetRegionId(v string) *RerankRequest {
	s.RegionId = &v
	return s
}

func (s *RerankRequest) SetReturnDocuments(v bool) *RerankRequest {
	s.ReturnDocuments = &v
	return s
}

func (s *RerankRequest) SetTopK(v int32) *RerankRequest {
	s.TopK = &v
	return s
}

func (s *RerankRequest) Validate() error {
	return dara.Validate(s)
}
