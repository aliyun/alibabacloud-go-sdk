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
	// HTTP status code
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
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	return dara.Validate(s)
}

type ListIndicesResponseBodyData struct {
	// The list of knowledge bases.
	Indices []*ListIndicesResponseBodyDataIndices `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
	// The specified page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The specified number of documents on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of knowledge bases returned.
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
	return dara.Validate(s)
}

type ListIndicesResponseBodyDataIndices struct {
	// The estimated length of chunks. Valid values: [1-2048].
	//
	// example:
	//
	// 5
	ChunkSize  *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ConfgModel *string `json:"ConfgModel,omitempty" xml:"ConfgModel,omitempty"`
	// The description of the knowledge base.
	//
	// example:
	//
	// If each RAM user belongs to a RAM group, the configuration is considered compliant.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of the primary key IDs of the documents.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// The name of the embedding model. Valid values:
	//
	// 	- text-embedding-v2
	//
	// example:
	//
	// conv-rewrite-qwen-1.8b
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	EnableRewrite      *bool   `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// example:
	//
	// 259899
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the knowledge base.
	//
	// example:
	//
	// temp_mUB4j
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The overlap length. Valid values: [0-1024].
	//
	// example:
	//
	// 10
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// Similarity Threshold Valid values: [0.01-1.00].
	//
	// example:
	//
	// 0.01
	RerankMinScore *string `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The name of the rank model. Valid values:
	//
	// 	- gte-rerank-hybrid
	//
	// 	- gte-rerank
	//
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// The clause identifier. Separate multiple clause identifiers with |. Valid values:
	//
	// 	- \\n: line break
	//
	// 	- ，: Chinese comma
	//
	// 	- ,: English comma
	//
	// 	- 。 : Chinese full stop
	//
	// 	- .: English full stop
	//
	// 	- ！ : Chinese exclamation point
	//
	// 	- ! : English exclamation point
	//
	// 	- ；: Chinese semicolon
	//
	// 	- ;: English semicolon
	//
	// 	- ？ : Chinese question mark
	//
	// 	- ?: English question mark
	//
	// example:
	//
	// \\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The ID of the vector storage instance.
	//
	// example:
	//
	// gp-bp1gq62t1788yw2ol
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// The region of the vector storage instance.
	//
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// The vector storage type of the knowledge base. Valid values:
	//
	// 	- ES: Built-in vector database.
	//
	// 	- BUILT_IN: Built-in vector database.
	//
	// 	- ADB: AnalyticDB for PostgreSQL database.
	//
	// example:
	//
	// es
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// The data type of [Data Management](https://bailian.console.aliyun.com/#/data-center). For unstructured knowledge base, possible values:
	//
	// 	- DATA_CENTER_CATEGORY: The category type.
	//
	// 	- DATA_CENTER_FILE: The document type.
	//
	// For structured knowledge base, possible values:
	//
	// 	- DATA_CENTER_STRUCTURED_TABLE: The data table type.
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The vector storage type of the knowledge base. Valid values:
	//
	// 	- UNSTRUCTURED
	//
	// example:
	//
	// structured
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
