// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIndicesResponseBody
	GetCode() *string
	SetData(v *ListIndicesResponseBodyData) *ListIndicesResponseBody
	GetData() *ListIndicesResponseBodyData
	SetMessage(v string) *ListIndicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIndicesResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListIndicesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListIndicesResponseBody
	GetSuccess() *bool
}

type ListIndicesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListIndicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIndicesResponseBody) GetData() *ListIndicesResponseBodyData {
	return s.Data
}

func (s *ListIndicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIndicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndicesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListIndicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIndicesResponseBody) SetCode(v string) *ListIndicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIndicesResponseBody) SetData(v *ListIndicesResponseBodyData) *ListIndicesResponseBody {
	s.Data = v
	return s
}

func (s *ListIndicesResponseBody) SetMessage(v string) *ListIndicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIndicesResponseBody) SetRequestId(v string) *ListIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndicesResponseBody) SetStatus(v string) *ListIndicesResponseBody {
	s.Status = &v
	return s
}

func (s *ListIndicesResponseBody) SetSuccess(v bool) *ListIndicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListIndicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIndicesResponseBodyData struct {
	// The list of knowledge bases.
	Indices []*ListIndicesResponseBodyDataIndices `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
	// The page number returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page returned.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 48
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIndicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBodyData) GetIndices() []*ListIndicesResponseBodyDataIndices {
	return s.Indices
}

func (s *ListIndicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIndicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIndicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIndicesResponseBodyData) SetIndices(v []*ListIndicesResponseBodyDataIndices) *ListIndicesResponseBodyData {
	s.Indices = v
	return s
}

func (s *ListIndicesResponseBodyData) SetPageNumber(v int32) *ListIndicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIndicesResponseBodyData) SetPageSize(v int32) *ListIndicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIndicesResponseBodyData) SetTotalCount(v int32) *ListIndicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListIndicesResponseBodyData) Validate() error {
	if s.Indices != nil {
		for _, item := range s.Indices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIndicesResponseBodyDataIndices struct {
	// The estimated chunk length. Valid values: 1 to 2048.
	//
	// example:
	//
	// 5
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The configuration mode used by this knowledge base. Valid values:
	//
	// - recommend: recommended configuration.
	//
	// - user-defined: custom configuration.
	//
	// example:
	//
	// recommend
	ConfgModel *string `json:"ConfgModel,omitempty" xml:"ConfgModel,omitempty"`
	// The description of the knowledge base.
	//
	// example:
	//
	// 清单中产品主要面向海外客户。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of file IDs.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// The name of the embedding model. Valid values:
	//
	// <props="china">
	//
	// - text-embedding-v4: the text-embedding-v4 model.
	//
	// - text-embedding-v3: the text-embedding-v3 model.
	//
	// - text-embedding-v2: the text-embedding-v2 model.
	//
	//
	//
	// <props="intl">
	//
	// - text-embedding-v2: the text-embedding-v2 model.
	//
	// .
	//
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// Indicates whether <props="china">[multi-turn conversation rewriting](https://help.aliyun.com/model-studio/use-cases/rag-optimization#b7031e2ad6cji)<props="intl">[multi-turn conversation rewriting](https://www.alibabacloud.com/help/model-studio/use-cases/rag-optimization#b7031e2ad6cji) is enabled for this knowledge base. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The knowledge base ID, which is the `Data.Id` returned by the **CreateIndex*	- operation.
	//
	// example:
	//
	// lecxr5xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the knowledge base.
	//
	// example:
	//
	// XXXX产品清单
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The chunk overlap length. Valid values: 0 to 1024.
	//
	// example:
	//
	// 10
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The similarity threshold. Valid values: 0.01 to 1.00.
	//
	// example:
	//
	// 0.01
	RerankMinScore *string `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The name of the rerank model. Valid values:
	//
	// <props="china">
	//
	// - qwen3-rerank-hybrid: qwen3-rerank (hybrid) reranking.
	//
	// - qwen3-rerank: qwen3-rerank reranking.
	//
	// - gte-rerank-hybrid: gte-rerank (hybrid) reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	// <props="intl">
	//
	// - gte-rerank-hybrid: official reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	// .
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// The sentence separator. If multiple separators are used, they are separated by |. Valid values:
	//
	// - \\
	//
	// : line break
	//
	// - ，: Chinese comma
	//
	// - ,: English comma
	//
	// - 。: Chinese period
	//
	// - .: English period
	//
	// - ！: Chinese exclamation mark
	//
	// - !: English exclamation mark
	//
	// - ；: Chinese semicolon
	//
	// - ;: English semicolon
	//
	// - ？: Chinese question mark
	//
	// - ?: English question mark.
	//
	// example:
	//
	// \\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The instance ID of the vector storage for the knowledge base.
	//
	// example:
	//
	// gp-bp1gq62t1788yxxxx
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region of the vector storage instance for the knowledge base.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The vector storage type of the knowledge base. Valid values:
	//
	// - ES: built-in vector database.
	//
	// - BUILT_IN: built-in vector database.
	//
	// - ADB: AnalyticDB for PostgreSQL database.
	//
	// example:
	//
	// BUILT_IN
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// The data type of Alibaba Cloud Model Studio <props="china">[application data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[application data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center).
	//
	//
	// For document search<props="china">/audio and video search knowledge bases, valid values:
	//
	// - DATA_CENTER_CATEGORY: category type.
	//
	// - DATA_CENTER_FILE: file type.
	//
	// For data query/image Q&A knowledge bases, valid values:
	//
	// - DATA_CENTER_STRUCTURED_TABLE: data table type.
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the knowledge base. Valid values:
	//
	// - UNSTRUCTURED: document search.
	//
	// example:
	//
	// UNSTRUCTURED
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s ListIndicesResponseBodyDataIndices) String() string {
	return dara.Prettify(s)
}

func (s ListIndicesResponseBodyDataIndices) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBodyDataIndices) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *ListIndicesResponseBodyDataIndices) GetConfgModel() *string {
	return s.ConfgModel
}

func (s *ListIndicesResponseBodyDataIndices) GetDescription() *string {
	return s.Description
}

func (s *ListIndicesResponseBodyDataIndices) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *ListIndicesResponseBodyDataIndices) GetEmbeddingModelName() *string {
	return s.EmbeddingModelName
}

func (s *ListIndicesResponseBodyDataIndices) GetEnableRewrite() *bool {
	return s.EnableRewrite
}

func (s *ListIndicesResponseBodyDataIndices) GetId() *string {
	return s.Id
}

func (s *ListIndicesResponseBodyDataIndices) GetName() *string {
	return s.Name
}

func (s *ListIndicesResponseBodyDataIndices) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *ListIndicesResponseBodyDataIndices) GetRerankMinScore() *string {
	return s.RerankMinScore
}

func (s *ListIndicesResponseBodyDataIndices) GetRerankModelName() *string {
	return s.RerankModelName
}

func (s *ListIndicesResponseBodyDataIndices) GetSeparator() *string {
	return s.Separator
}

func (s *ListIndicesResponseBodyDataIndices) GetSinkInstanceId() *string {
	return s.SinkInstanceId
}

func (s *ListIndicesResponseBodyDataIndices) GetSinkRegion() *string {
	return s.SinkRegion
}

func (s *ListIndicesResponseBodyDataIndices) GetSinkType() *string {
	return s.SinkType
}

func (s *ListIndicesResponseBodyDataIndices) GetSourceType() *string {
	return s.SourceType
}

func (s *ListIndicesResponseBodyDataIndices) GetStructureType() *string {
	return s.StructureType
}

func (s *ListIndicesResponseBodyDataIndices) SetChunkSize(v int32) *ListIndicesResponseBodyDataIndices {
	s.ChunkSize = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetConfgModel(v string) *ListIndicesResponseBodyDataIndices {
	s.ConfgModel = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetDescription(v string) *ListIndicesResponseBodyDataIndices {
	s.Description = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetDocumentIds(v []*string) *ListIndicesResponseBodyDataIndices {
	s.DocumentIds = v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetEmbeddingModelName(v string) *ListIndicesResponseBodyDataIndices {
	s.EmbeddingModelName = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetEnableRewrite(v bool) *ListIndicesResponseBodyDataIndices {
	s.EnableRewrite = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetId(v string) *ListIndicesResponseBodyDataIndices {
	s.Id = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetName(v string) *ListIndicesResponseBodyDataIndices {
	s.Name = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetOverlapSize(v int32) *ListIndicesResponseBodyDataIndices {
	s.OverlapSize = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetRerankMinScore(v string) *ListIndicesResponseBodyDataIndices {
	s.RerankMinScore = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetRerankModelName(v string) *ListIndicesResponseBodyDataIndices {
	s.RerankModelName = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSeparator(v string) *ListIndicesResponseBodyDataIndices {
	s.Separator = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSinkInstanceId(v string) *ListIndicesResponseBodyDataIndices {
	s.SinkInstanceId = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSinkRegion(v string) *ListIndicesResponseBodyDataIndices {
	s.SinkRegion = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSinkType(v string) *ListIndicesResponseBodyDataIndices {
	s.SinkType = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSourceType(v string) *ListIndicesResponseBodyDataIndices {
	s.SourceType = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetStructureType(v string) *ListIndicesResponseBodyDataIndices {
	s.StructureType = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) Validate() error {
	return dara.Validate(s)
}
