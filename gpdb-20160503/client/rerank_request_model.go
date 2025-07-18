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
	Documents []*string `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
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
