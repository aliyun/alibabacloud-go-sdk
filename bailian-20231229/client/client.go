// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847
	LeaseId *string `json:"LeaseId,omitempty" xml:"LeaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) SetCategoryId(v string) *AddFileRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileRequest) SetLeaseId(v string) *AddFileRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileRequest) SetParser(v string) *AddFileRequest {
	s.Parser = &v
	return s
}

type AddFileResponseBody struct {
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileResponseBody) SetCode(v string) *AddFileResponseBody {
	s.Code = &v
	return s
}

func (s *AddFileResponseBody) SetData(v *AddFileResponseBodyData) *AddFileResponseBody {
	s.Data = v
	return s
}

func (s *AddFileResponseBody) SetMessage(v string) *AddFileResponseBody {
	s.Message = &v
	return s
}

func (s *AddFileResponseBody) SetRequestId(v string) *AddFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFileResponseBody) SetStatus(v string) *AddFileResponseBody {
	s.Status = &v
	return s
}

func (s *AddFileResponseBody) SetSuccess(v string) *AddFileResponseBody {
	s.Success = &v
	return s
}

type AddFileResponseBodyData struct {
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFileResponseBodyData) SetFileId(v string) *AddFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *AddFileResponseBodyData) SetParser(v string) *AddFileResponseBodyData {
	s.Parser = &v
	return s
}

type AddFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponse) GoString() string {
	return s.String()
}

func (s *AddFileResponse) SetHeaders(v map[string]*string) *AddFileResponse {
	s.Headers = v
	return s
}

func (s *AddFileResponse) SetStatusCode(v int32) *AddFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFileResponse) SetBody(v *AddFileResponseBody) *AddFileResponse {
	s.Body = v
	return s
}

type ApplyFileUploadLeaseRequest struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19657c391f6c70bcea63c154d8606bb3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	SizeInBytes *string `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
}

func (s ApplyFileUploadLeaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseRequest) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseRequest) SetFileName(v string) *ApplyFileUploadLeaseRequest {
	s.FileName = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetMd5(v string) *ApplyFileUploadLeaseRequest {
	s.Md5 = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest {
	s.SizeInBytes = &v
	return s
}

type ApplyFileUploadLeaseResponseBody struct {
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ApplyFileUploadLeaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBody) SetCode(v string) *ApplyFileUploadLeaseResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetData(v *ApplyFileUploadLeaseResponseBodyData) *ApplyFileUploadLeaseResponseBody {
	s.Data = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetMessage(v string) *ApplyFileUploadLeaseResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetRequestId(v string) *ApplyFileUploadLeaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetStatus(v string) *ApplyFileUploadLeaseResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetSuccess(v bool) *ApplyFileUploadLeaseResponseBody {
	s.Success = &v
	return s
}

type ApplyFileUploadLeaseResponseBodyData struct {
	// example:
	//
	// 1e6a159107384782be5e45ac4759b247.1719325231035
	FileUploadLeaseId *string                                    `json:"FileUploadLeaseId,omitempty" xml:"FileUploadLeaseId,omitempty"`
	Param             *ApplyFileUploadLeaseResponseBodyDataParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetFileUploadLeaseId(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.FileUploadLeaseId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetParam(v *ApplyFileUploadLeaseResponseBodyDataParam) *ApplyFileUploadLeaseResponseBodyData {
	s.Param = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetType(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.Type = &v
	return s
}

type ApplyFileUploadLeaseResponseBodyDataParam struct {
	// example:
	//
	// "X-bailian-extra": "MTAwNTQyNjQ5NTE2OTE3OA==",
	//
	//         "Content-Type": "application/pdf"
	Headers interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// PUT
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// https://bailian-datahub-data-origin-prod.oss-cn-hangzhou.aliyuncs.com/1005426495169178/10024405/68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847.pdf?Expires=1716699536&OSSAccessKeyId=TestID&Signature=HfwPUZo4pR6DatSDym0zFKVh9Wg%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetHeaders(v interface{}) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetMethod(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Method = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetUrl(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Url = &v
	return s
}

type ApplyFileUploadLeaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFileUploadLeaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFileUploadLeaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponse) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponse) SetHeaders(v map[string]*string) *ApplyFileUploadLeaseResponse {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetStatusCode(v int32) *ApplyFileUploadLeaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetBody(v *ApplyFileUploadLeaseResponseBody) *ApplyFileUploadLeaseResponse {
	s.Body = v
	return s
}

type CreateIndexRequest struct {
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 128
	ChunkSize   *int32                       `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	Columns     []*CreateIndexRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Description *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentIds []*string                    `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// example:
	//
	// ,
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// example:
	//
	// gp-bp321093j84
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// structured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) SetCategoryIds(v []*string) *CreateIndexRequest {
	s.CategoryIds = v
	return s
}

func (s *CreateIndexRequest) SetChunkSize(v int32) *CreateIndexRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateIndexRequest) SetColumns(v []*CreateIndexRequestColumns) *CreateIndexRequest {
	s.Columns = v
	return s
}

func (s *CreateIndexRequest) SetDescription(v string) *CreateIndexRequest {
	s.Description = &v
	return s
}

func (s *CreateIndexRequest) SetDocumentIds(v []*string) *CreateIndexRequest {
	s.DocumentIds = v
	return s
}

func (s *CreateIndexRequest) SetEmbeddingModelName(v string) *CreateIndexRequest {
	s.EmbeddingModelName = &v
	return s
}

func (s *CreateIndexRequest) SetName(v string) *CreateIndexRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexRequest) SetOverlapSize(v int32) *CreateIndexRequest {
	s.OverlapSize = &v
	return s
}

func (s *CreateIndexRequest) SetRerankMinScore(v float64) *CreateIndexRequest {
	s.RerankMinScore = &v
	return s
}

func (s *CreateIndexRequest) SetRerankModelName(v string) *CreateIndexRequest {
	s.RerankModelName = &v
	return s
}

func (s *CreateIndexRequest) SetSeparator(v string) *CreateIndexRequest {
	s.Separator = &v
	return s
}

func (s *CreateIndexRequest) SetSinkInstanceId(v string) *CreateIndexRequest {
	s.SinkInstanceId = &v
	return s
}

func (s *CreateIndexRequest) SetSinkRegion(v string) *CreateIndexRequest {
	s.SinkRegion = &v
	return s
}

func (s *CreateIndexRequest) SetSinkType(v string) *CreateIndexRequest {
	s.SinkType = &v
	return s
}

func (s *CreateIndexRequest) SetSourceType(v string) *CreateIndexRequest {
	s.SourceType = &v
	return s
}

func (s *CreateIndexRequest) SetStructureType(v string) *CreateIndexRequest {
	s.StructureType = &v
	return s
}

type CreateIndexRequestColumns struct {
	Column   *string `json:"Column,omitempty" xml:"Column,omitempty"`
	IsRecall *bool   `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	IsSearch *bool   `json:"IsSearch,omitempty" xml:"IsSearch,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIndexRequestColumns) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestColumns) SetColumn(v string) *CreateIndexRequestColumns {
	s.Column = &v
	return s
}

func (s *CreateIndexRequestColumns) SetIsRecall(v bool) *CreateIndexRequestColumns {
	s.IsRecall = &v
	return s
}

func (s *CreateIndexRequestColumns) SetIsSearch(v bool) *CreateIndexRequestColumns {
	s.IsSearch = &v
	return s
}

func (s *CreateIndexRequestColumns) SetName(v string) *CreateIndexRequestColumns {
	s.Name = &v
	return s
}

func (s *CreateIndexRequestColumns) SetType(v string) *CreateIndexRequestColumns {
	s.Type = &v
	return s
}

type CreateIndexShrinkRequest struct {
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// 128
	ChunkSize         *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ColumnsShrink     *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// example:
	//
	// ,
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// example:
	//
	// gp-bp321093j84
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// structured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s CreateIndexShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexShrinkRequest) SetCategoryIdsShrink(v string) *CreateIndexShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetChunkSize(v int32) *CreateIndexShrinkRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetColumnsShrink(v string) *CreateIndexShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDescription(v string) *CreateIndexShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDocumentIdsShrink(v string) *CreateIndexShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetEmbeddingModelName(v string) *CreateIndexShrinkRequest {
	s.EmbeddingModelName = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetName(v string) *CreateIndexShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetOverlapSize(v int32) *CreateIndexShrinkRequest {
	s.OverlapSize = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankMinScore(v float64) *CreateIndexShrinkRequest {
	s.RerankMinScore = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankModelName(v string) *CreateIndexShrinkRequest {
	s.RerankModelName = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSeparator(v string) *CreateIndexShrinkRequest {
	s.Separator = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkInstanceId(v string) *CreateIndexShrinkRequest {
	s.SinkInstanceId = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkRegion(v string) *CreateIndexShrinkRequest {
	s.SinkRegion = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkType(v string) *CreateIndexShrinkRequest {
	s.SinkType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSourceType(v string) *CreateIndexShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetStructureType(v string) *CreateIndexShrinkRequest {
	s.StructureType = &v
	return s
}

type CreateIndexResponseBody struct {
	// example:
	//
	// Forbidden
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid input, variable name is missing
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) SetCode(v string) *CreateIndexResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIndexResponseBody) SetData(v *CreateIndexResponseBodyData) *CreateIndexResponseBody {
	s.Data = v
	return s
}

func (s *CreateIndexResponseBody) SetMessage(v string) *CreateIndexResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetStatus(v string) *CreateIndexResponseBody {
	s.Status = &v
	return s
}

func (s *CreateIndexResponseBody) SetSuccess(v bool) *CreateIndexResponseBody {
	s.Success = &v
	return s
}

type CreateIndexResponseBodyData struct {
	// example:
	//
	// jkurxhju6b
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateIndexResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBodyData) SetId(v string) *CreateIndexResponseBodyData {
	s.Id = &v
	return s
}

type CreateIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndexResponse) SetBody(v *CreateIndexResponseBody) *CreateIndexResponse {
	s.Body = v
	return s
}

type DescribeFileResponseBody struct {
	// example:
	//
	// Success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBody) SetCode(v string) *DescribeFileResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFileResponseBody) SetData(v *DescribeFileResponseBodyData) *DescribeFileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFileResponseBody) SetMessage(v string) *DescribeFileResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFileResponseBody) SetRequestId(v string) *DescribeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileResponseBody) SetStatus(v string) *DescribeFileResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFileResponseBody) SetSuccess(v bool) *DescribeFileResponseBody {
	s.Success = &v
	return s
}

type DescribeFileResponseBodyData struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2024-05-26 12:45:43
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// test.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// example:
	//
	// 1234
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// example:
	//
	// PARSE_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBodyData) SetCategoryId(v string) *DescribeFileResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetCreateTime(v string) *DescribeFileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileId(v string) *DescribeFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileName(v string) *DescribeFileResponseBodyData {
	s.FileName = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileType(v string) *DescribeFileResponseBodyData {
	s.FileType = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetParser(v string) *DescribeFileResponseBodyData {
	s.Parser = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetSizeInBytes(v int64) *DescribeFileResponseBodyData {
	s.SizeInBytes = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetStatus(v string) *DescribeFileResponseBodyData {
	s.Status = &v
	return s
}

type DescribeFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileResponse) SetHeaders(v map[string]*string) *DescribeFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileResponse) SetStatusCode(v int32) *DescribeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileResponse) SetBody(v *DescribeFileResponseBody) *DescribeFileResponse {
	s.Body = v
	return s
}

type GetIndexJobStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230718xxxx-146c93bf
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetIndexJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusRequest) SetIndexId(v string) *GetIndexJobStatusRequest {
	s.IndexId = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetJobId(v string) *GetIndexJobStatusRequest {
	s.JobId = &v
	return s
}

type GetIndexJobStatusResponseBody struct {
	// example:
	//
	// Index.Forbidden
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetIndexJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIndexJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBody) SetCode(v string) *GetIndexJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetData(v *GetIndexJobStatusResponseBodyData) *GetIndexJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetMessage(v string) *GetIndexJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetRequestId(v string) *GetIndexJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetStatus(v string) *GetIndexJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetSuccess(v bool) *GetIndexJobStatusResponseBody {
	s.Success = &v
	return s
}

type GetIndexJobStatusResponseBodyData struct {
	Documents []*GetIndexJobStatusResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// 66122af12a4e45ddae6bd6c845556647
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIndexJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBodyData) SetDocuments(v []*GetIndexJobStatusResponseBodyDataDocuments) *GetIndexJobStatusResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) SetJobId(v string) *GetIndexJobStatusResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) SetStatus(v string) *GetIndexJobStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetIndexJobStatusResponseBodyDataDocuments struct {
	// example:
	//
	// Index.Document.ChunkError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	DocId   *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// document parse error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIndexJobStatusResponseBodyDataDocuments) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetCode(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Code = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetDocId(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetDocName(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.DocName = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetMessage(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Message = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetStatus(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Status = &v
	return s
}

type GetIndexJobStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponse) SetHeaders(v map[string]*string) *GetIndexJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetIndexJobStatusResponse) SetStatusCode(v int32) *GetIndexJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexJobStatusResponse) SetBody(v *GetIndexJobStatusResponseBody) *GetIndexJobStatusResponse {
	s.Body = v
	return s
}

type RetrieveRequest struct {
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// example:
	//
	// false
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0m2g6t
	IndexId *string                  `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Query   *string                  `json:"Query,omitempty" xml:"Query,omitempty"`
	Rerank  []*RetrieveRequestRerank `json:"Rerank,omitempty" xml:"Rerank,omitempty" type:"Repeated"`
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// 5
	RerankTopN *int32                    `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	Rewrite    []*RetrieveRequestRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SaveRetrieverHistory *bool `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	// example:
	//
	// 100
	SparseSimilarityTopK *int32 `json:"SparseSimilarityTopK,omitempty" xml:"SparseSimilarityTopK,omitempty"`
}

func (s RetrieveRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequest) GoString() string {
	return s.String()
}

func (s *RetrieveRequest) SetDenseSimilarityTopK(v int32) *RetrieveRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *RetrieveRequest) SetEnableReranking(v bool) *RetrieveRequest {
	s.EnableReranking = &v
	return s
}

func (s *RetrieveRequest) SetEnableRewrite(v bool) *RetrieveRequest {
	s.EnableRewrite = &v
	return s
}

func (s *RetrieveRequest) SetIndexId(v string) *RetrieveRequest {
	s.IndexId = &v
	return s
}

func (s *RetrieveRequest) SetQuery(v string) *RetrieveRequest {
	s.Query = &v
	return s
}

func (s *RetrieveRequest) SetRerank(v []*RetrieveRequestRerank) *RetrieveRequest {
	s.Rerank = v
	return s
}

func (s *RetrieveRequest) SetRerankMinScore(v float32) *RetrieveRequest {
	s.RerankMinScore = &v
	return s
}

func (s *RetrieveRequest) SetRerankTopN(v int32) *RetrieveRequest {
	s.RerankTopN = &v
	return s
}

func (s *RetrieveRequest) SetRewrite(v []*RetrieveRequestRewrite) *RetrieveRequest {
	s.Rewrite = v
	return s
}

func (s *RetrieveRequest) SetSaveRetrieverHistory(v bool) *RetrieveRequest {
	s.SaveRetrieverHistory = &v
	return s
}

func (s *RetrieveRequest) SetSparseSimilarityTopK(v int32) *RetrieveRequest {
	s.SparseSimilarityTopK = &v
	return s
}

type RetrieveRequestRerank struct {
	// example:
	//
	// gte-rerank-hybrid
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s RetrieveRequestRerank) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequestRerank) GoString() string {
	return s.String()
}

func (s *RetrieveRequestRerank) SetModelName(v string) *RetrieveRequestRerank {
	s.ModelName = &v
	return s
}

type RetrieveRequestRewrite struct {
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s RetrieveRequestRewrite) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequestRewrite) GoString() string {
	return s.String()
}

func (s *RetrieveRequestRewrite) SetModelName(v string) *RetrieveRequestRewrite {
	s.ModelName = &v
	return s
}

type RetrieveShrinkRequest struct {
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// example:
	//
	// false
	EnableRewrite *bool `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0m2g6t
	IndexId      *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Query        *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RerankShrink *string `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// 5
	RerankTopN    *int32  `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	RewriteShrink *string `json:"Rewrite,omitempty" xml:"Rewrite,omitempty"`
	// example:
	//
	// false
	SaveRetrieverHistory *bool `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	// example:
	//
	// 100
	SparseSimilarityTopK *int32 `json:"SparseSimilarityTopK,omitempty" xml:"SparseSimilarityTopK,omitempty"`
}

func (s RetrieveShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetrieveShrinkRequest) SetDenseSimilarityTopK(v int32) *RetrieveShrinkRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *RetrieveShrinkRequest) SetEnableReranking(v bool) *RetrieveShrinkRequest {
	s.EnableReranking = &v
	return s
}

func (s *RetrieveShrinkRequest) SetEnableRewrite(v bool) *RetrieveShrinkRequest {
	s.EnableRewrite = &v
	return s
}

func (s *RetrieveShrinkRequest) SetIndexId(v string) *RetrieveShrinkRequest {
	s.IndexId = &v
	return s
}

func (s *RetrieveShrinkRequest) SetQuery(v string) *RetrieveShrinkRequest {
	s.Query = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankShrink(v string) *RetrieveShrinkRequest {
	s.RerankShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankMinScore(v float32) *RetrieveShrinkRequest {
	s.RerankMinScore = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankTopN(v int32) *RetrieveShrinkRequest {
	s.RerankTopN = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRewriteShrink(v string) *RetrieveShrinkRequest {
	s.RewriteShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSaveRetrieverHistory(v bool) *RetrieveShrinkRequest {
	s.SaveRetrieverHistory = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSparseSimilarityTopK(v int32) *RetrieveShrinkRequest {
	s.SparseSimilarityTopK = &v
	return s
}

type RetrieveResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RetrieveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetrieveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBody) SetCode(v string) *RetrieveResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveResponseBody) SetData(v *RetrieveResponseBodyData) *RetrieveResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveResponseBody) SetMessage(v string) *RetrieveResponseBody {
	s.Message = &v
	return s
}

func (s *RetrieveResponseBody) SetRequestId(v string) *RetrieveResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveResponseBody) SetStatus(v string) *RetrieveResponseBody {
	s.Status = &v
	return s
}

func (s *RetrieveResponseBody) SetSuccess(v bool) *RetrieveResponseBody {
	s.Success = &v
	return s
}

type RetrieveResponseBodyData struct {
	Nodes []*RetrieveResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RetrieveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyData) SetNodes(v []*RetrieveResponseBodyDataNodes) *RetrieveResponseBodyData {
	s.Nodes = v
	return s
}

type RetrieveResponseBodyDataNodes struct {
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 0.3
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text  *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RetrieveResponseBodyDataNodes) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyDataNodes) SetMetadata(v interface{}) *RetrieveResponseBodyDataNodes {
	s.Metadata = v
	return s
}

func (s *RetrieveResponseBodyDataNodes) SetScore(v float64) *RetrieveResponseBodyDataNodes {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyDataNodes) SetText(v string) *RetrieveResponseBodyDataNodes {
	s.Text = &v
	return s
}

type RetrieveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponse) GoString() string {
	return s.String()
}

func (s *RetrieveResponse) SetHeaders(v map[string]*string) *RetrieveResponse {
	s.Headers = v
	return s
}

func (s *RetrieveResponse) SetStatusCode(v int32) *RetrieveResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveResponse) SetBody(v *RetrieveResponseBody) *RetrieveResponse {
	s.Body = v
	return s
}

type SubmitIndexJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s SubmitIndexJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobRequest) SetIndexId(v string) *SubmitIndexJobRequest {
	s.IndexId = &v
	return s
}

type SubmitIndexJobResponseBody struct {
	// example:
	//
	// InvalidParameter
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitIndexJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIndexJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponseBody) SetCode(v string) *SubmitIndexJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetData(v *SubmitIndexJobResponseBodyData) *SubmitIndexJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitIndexJobResponseBody) SetMessage(v string) *SubmitIndexJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetRequestId(v string) *SubmitIndexJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetStatus(v string) *SubmitIndexJobResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetSuccess(v bool) *SubmitIndexJobResponseBody {
	s.Success = &v
	return s
}

type SubmitIndexJobResponseBodyData struct {
	// example:
	//
	// eFDr2fGRzP9gdDZWAdo3YQ==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// khdyak1uuj
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s SubmitIndexJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponseBodyData) SetId(v string) *SubmitIndexJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitIndexJobResponseBodyData) SetIndexId(v string) *SubmitIndexJobResponseBodyData {
	s.IndexId = &v
	return s
}

type SubmitIndexJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIndexJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIndexJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponse) SetHeaders(v map[string]*string) *SubmitIndexJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIndexJobResponse) SetStatusCode(v int32) *SubmitIndexJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIndexJobResponse) SetBody(v *SubmitIndexJobResponseBody) *SubmitIndexJobResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将临时上传的文档导入百炼数据中心，导入成功之后会自动触发文档解析。
//
// @param request - AddFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFileResponse
func (client *Client) AddFileWithOptions(WorkspaceId *string, request *AddFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LeaseId)) {
		body["LeaseId"] = request.LeaseId
	}

	if !tea.BoolValue(util.IsUnset(request.Parser)) {
		body["Parser"] = request.Parser
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将临时上传的文档导入百炼数据中心，导入成功之后会自动触发文档解析。
//
// @param request - AddFileRequest
//
// @return AddFileResponse
func (client *Client) AddFile(WorkspaceId *string, request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 请求文档上传租约，进行文档上传。
//
// @param request - ApplyFileUploadLeaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLeaseWithOptions(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyFileUploadLeaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		body["Md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.SizeInBytes)) {
		body["SizeInBytes"] = request.SizeInBytes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyFileUploadLease"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/category/" + tea.StringValue(openapiutil.GetEncodeParam(CategoryId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 请求文档上传租约，进行文档上传。
//
// @param request - ApplyFileUploadLeaseRequest
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLease(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest) (_result *ApplyFileUploadLeaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.ApplyFileUploadLeaseWithOptions(CategoryId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建并运行pipeline
//
// @param tmpReq - CreateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
func (client *Client) CreateIndexWithOptions(WorkspaceId *string, tmpReq *CreateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIndexShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CategoryIds)) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, tea.String("CategoryIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Columns)) {
		request.ColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Columns, tea.String("Columns"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DocumentIds)) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, tea.String("DocumentIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryIdsShrink)) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChunkSize)) {
		query["ChunkSize"] = request.ChunkSize
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnsShrink)) {
		query["Columns"] = request.ColumnsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentIdsShrink)) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EmbeddingModelName)) {
		query["EmbeddingModelName"] = request.EmbeddingModelName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OverlapSize)) {
		query["OverlapSize"] = request.OverlapSize
	}

	if !tea.BoolValue(util.IsUnset(request.RerankMinScore)) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !tea.BoolValue(util.IsUnset(request.RerankModelName)) {
		query["RerankModelName"] = request.RerankModelName
	}

	if !tea.BoolValue(util.IsUnset(request.Separator)) {
		query["Separator"] = request.Separator
	}

	if !tea.BoolValue(util.IsUnset(request.SinkInstanceId)) {
		query["SinkInstanceId"] = request.SinkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SinkRegion)) {
		query["SinkRegion"] = request.SinkRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SinkType)) {
		query["SinkType"] = request.SinkType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StructureType)) {
		query["StructureType"] = request.StructureType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndex"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并运行pipeline
//
// @param request - CreateIndexRequest
//
// @return CreateIndexResponse
func (client *Client) CreateIndex(WorkspaceId *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档基本信息，包括文档名称、类型、状态等。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileResponse
func (client *Client) DescribeFileWithOptions(WorkspaceId *string, FileId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFileResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file/" + tea.StringValue(openapiutil.GetEncodeParam(FileId)) + "/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档基本信息，包括文档名称、类型、状态等。
//
// @return DescribeFileResponse
func (client *Client) DescribeFile(WorkspaceId *string, FileId *string) (_result *DescribeFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFileResponse{}
	_body, _err := client.DescribeFileWithOptions(WorkspaceId, FileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Index运行状态
//
// @param request - GetIndexJobStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexJobStatusResponse
func (client *Client) GetIndexJobStatusWithOptions(WorkspaceId *string, request *GetIndexJobStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndexJobStatus"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/job/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Index运行状态
//
// @param request - GetIndexJobStatusRequest
//
// @return GetIndexJobStatusResponse
func (client *Client) GetIndexJobStatus(WorkspaceId *string, request *GetIndexJobStatusRequest) (_result *GetIndexJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexJobStatusResponse{}
	_body, _err := client.GetIndexJobStatusWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 召回测试
//
// @param tmpReq - RetrieveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveResponse
func (client *Client) RetrieveWithOptions(WorkspaceId *string, tmpReq *RetrieveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetrieveResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RetrieveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Rerank)) {
		request.RerankShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rerank, tea.String("Rerank"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Rewrite)) {
		request.RewriteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rewrite, tea.String("Rewrite"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DenseSimilarityTopK)) {
		query["DenseSimilarityTopK"] = request.DenseSimilarityTopK
	}

	if !tea.BoolValue(util.IsUnset(request.EnableReranking)) {
		query["EnableReranking"] = request.EnableReranking
	}

	if !tea.BoolValue(util.IsUnset(request.EnableRewrite)) {
		query["EnableRewrite"] = request.EnableRewrite
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RerankShrink)) {
		query["Rerank"] = request.RerankShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RerankMinScore)) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !tea.BoolValue(util.IsUnset(request.RerankTopN)) {
		query["RerankTopN"] = request.RerankTopN
	}

	if !tea.BoolValue(util.IsUnset(request.RewriteShrink)) {
		query["Rewrite"] = request.RewriteShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SaveRetrieverHistory)) {
		query["SaveRetrieverHistory"] = request.SaveRetrieverHistory
	}

	if !tea.BoolValue(util.IsUnset(request.SparseSimilarityTopK)) {
		query["SparseSimilarityTopK"] = request.SparseSimilarityTopK
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Retrieve"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/retrieve"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 召回测试
//
// @param request - RetrieveRequest
//
// @return RetrieveResponse
func (client *Client) Retrieve(WorkspaceId *string, request *RetrieveRequest) (_result *RetrieveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetrieveResponse{}
	_body, _err := client.RetrieveWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交索引任务
//
// @param request - SubmitIndexJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIndexJobResponse
func (client *Client) SubmitIndexJobWithOptions(WorkspaceId *string, request *SubmitIndexJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitIndexJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitIndexJob"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/submit_index_job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitIndexJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交索引任务
//
// @param request - SubmitIndexJobRequest
//
// @return SubmitIndexJobResponse
func (client *Client) SubmitIndexJob(WorkspaceId *string, request *SubmitIndexJobRequest) (_result *SubmitIndexJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitIndexJobResponse{}
	_body, _err := client.SubmitIndexJobWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
