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
	// Instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// List of documents to be re-ordered.
	DocumentsShrink *string `json:"Documents,omitempty" xml:"Documents,omitempty"`
	// Maximum number of chunks allowed when the text exceeds the model window:
	//
	// - bge-reranker-v2-m3: default value is 10.
	//
	// - bge-reranker-v2-minicpm-layerwise: default value is 5:
	//
	// > Example of splitting
	//
	// > - If using the bge-reranker-v2-minicpm-layerwise model, the maximum single inference window is 2048 tokens. If the query is 48 tokens and the content of a single document parameter is 9000 tokens, it will be divided as follows: 1-2000 for the first, 2001-4000 for the second, and so on. If the number of splits exceeds MaxChunksPerDoc, the remaining sentences will be discarded.
	//
	// example:
	//
	// 10
	MaxChunksPerDoc *int32 `json:"MaxChunksPerDoc,omitempty" xml:"MaxChunksPerDoc,omitempty"`
	// Rerank model, currently supports:
	//
	// - bge-reranker-v2-m3: (default), better performance, supports 8192 tokens per inference, if exceeded, it will be split, which may reduce the effect.
	//
	// - bge-reranker-v2-minicpm-layerwise: better performance than v2-m3, supports 2048 tokens per inference, if exceeded, it will be split, which may reduce the effect.
	//
	// example:
	//
	// bge-reranker-v2-m3
	Model   *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Query statement for Rerank.
	//
	// example:
	//
	// What is ADBPG?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Region ID where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// If set to false, does not return the Documents text, only returns the index of the document order and the rerank score.
	//
	// example:
	//
	// false
	ReturnDocuments *bool `json:"ReturnDocuments,omitempty" xml:"ReturnDocuments,omitempty"`
	// Number of most relevant documents to return.
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
