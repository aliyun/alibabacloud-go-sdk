// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type CreateFinReportSummaryTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableTable *bool `json:"enableTable,omitempty" xml:"enableTable,omitempty"`
	// example:
	//
	// 10
	EndPage     *int32  `json:"endPage,omitempty" xml:"endPage,omitempty"`
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 1
	StartPage *int32 `json:"startPage,omitempty" xml:"startPage,omitempty"`
	// example:
	//
	// custom
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateFinReportSummaryTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFinReportSummaryTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskRequest) SetDocId(v string) *CreateFinReportSummaryTaskRequest {
	s.DocId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetEnableTable(v bool) *CreateFinReportSummaryTaskRequest {
	s.EnableTable = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetEndPage(v int32) *CreateFinReportSummaryTaskRequest {
	s.EndPage = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetInstruction(v string) *CreateFinReportSummaryTaskRequest {
	s.Instruction = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetLibraryId(v string) *CreateFinReportSummaryTaskRequest {
	s.LibraryId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetModelId(v string) *CreateFinReportSummaryTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetStartPage(v int32) *CreateFinReportSummaryTaskRequest {
	s.StartPage = &v
	return s
}

func (s *CreateFinReportSummaryTaskRequest) SetTaskType(v string) *CreateFinReportSummaryTaskRequest {
	s.TaskType = &v
	return s
}

type CreateFinReportSummaryTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 3284627354
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateFinReportSummaryTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFinReportSummaryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskResponseBody) SetCost(v int64) *CreateFinReportSummaryTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetData(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetDataType(v string) *CreateFinReportSummaryTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetErrCode(v string) *CreateFinReportSummaryTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetMessage(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetRequestId(v string) *CreateFinReportSummaryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetSuccess(v bool) *CreateFinReportSummaryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponseBody) SetTime(v string) *CreateFinReportSummaryTaskResponseBody {
	s.Time = &v
	return s
}

type CreateFinReportSummaryTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFinReportSummaryTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFinReportSummaryTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFinReportSummaryTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFinReportSummaryTaskResponse) SetHeaders(v map[string]*string) *CreateFinReportSummaryTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) SetStatusCode(v int32) *CreateFinReportSummaryTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFinReportSummaryTaskResponse) SetBody(v *CreateFinReportSummaryTaskResponseBody) *CreateFinReportSummaryTaskResponse {
	s.Body = v
	return s
}

type CreateLibraryRequest struct {
	// This parameter is required.
	Description  *string                           `json:"description,omitempty" xml:"description,omitempty"`
	IndexSetting *CreateLibraryRequestIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	// This parameter is required.
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s CreateLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequest) SetDescription(v string) *CreateLibraryRequest {
	s.Description = &v
	return s
}

func (s *CreateLibraryRequest) SetIndexSetting(v *CreateLibraryRequestIndexSetting) *CreateLibraryRequest {
	s.IndexSetting = v
	return s
}

func (s *CreateLibraryRequest) SetLibraryName(v string) *CreateLibraryRequest {
	s.LibraryName = &v
	return s
}

type CreateLibraryRequestIndexSetting struct {
	ChunkStrategy      *CreateLibraryRequestIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *CreateLibraryRequestIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                             `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *CreateLibraryRequestIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *CreateLibraryRequestIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *CreateLibraryRequestIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *CreateLibraryRequestIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s CreateLibraryRequestIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSetting) SetChunkStrategy(v *CreateLibraryRequestIndexSettingChunkStrategy) *CreateLibraryRequestIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetModelConfig(v *CreateLibraryRequestIndexSettingModelConfig) *CreateLibraryRequestIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetPromptRoleStyle(v string) *CreateLibraryRequestIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetQueryEnhancer(v *CreateLibraryRequestIndexSettingQueryEnhancer) *CreateLibraryRequestIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetRecallStrategy(v *CreateLibraryRequestIndexSettingRecallStrategy) *CreateLibraryRequestIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetTextIndexSetting(v *CreateLibraryRequestIndexSettingTextIndexSetting) *CreateLibraryRequestIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *CreateLibraryRequestIndexSetting) SetVectorIndexSetting(v *CreateLibraryRequestIndexSettingVectorIndexSetting) *CreateLibraryRequestIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type CreateLibraryRequestIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 300
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 20
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 300
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 300
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s CreateLibraryRequestIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetEnhanceTable(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetOverlap(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplit(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSize(v int32) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) SetSplit(v bool) *CreateLibraryRequestIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type CreateLibraryRequestIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s CreateLibraryRequestIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingModelConfig) SetTemperature(v float64) *CreateLibraryRequestIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingModelConfig) SetTopP(v float64) *CreateLibraryRequestIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type CreateLibraryRequestIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// xxxx
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s CreateLibraryRequestIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetEnableSession(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *CreateLibraryRequestIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type CreateLibraryRequestIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s CreateLibraryRequestIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) SetDocumentRankType(v string) *CreateLibraryRequestIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) SetLimit(v int32) *CreateLibraryRequestIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type CreateLibraryRequestIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Standard
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s CreateLibraryRequestIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetCategory(v string) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetEnable(v bool) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetRankThreshold(v float64) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) SetTopK(v int32) *CreateLibraryRequestIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type CreateLibraryRequestIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s CreateLibraryRequestIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetCategory(v string) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetEnable(v bool) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) SetTopK(v int32) *CreateLibraryRequestIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type CreateLibraryResponseBody struct {
	// example:
	//
	// 300
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// a1b2c3
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// null
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreateLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponseBody) SetCost(v int64) *CreateLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateLibraryResponseBody) SetData(v string) *CreateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLibraryResponseBody) SetDataType(v string) *CreateLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *CreateLibraryResponseBody) SetErrCode(v string) *CreateLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateLibraryResponseBody) SetMessage(v string) *CreateLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLibraryResponseBody) SetRequestId(v string) *CreateLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLibraryResponseBody) SetSuccess(v bool) *CreateLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLibraryResponseBody) SetTime(v string) *CreateLibraryResponseBody {
	s.Time = &v
	return s
}

type CreateLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryResponse) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponse) SetHeaders(v map[string]*string) *CreateLibraryResponse {
	s.Headers = v
	return s
}

func (s *CreateLibraryResponse) SetStatusCode(v int32) *CreateLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLibraryResponse) SetBody(v *CreateLibraryResponseBody) *CreateLibraryResponse {
	s.Body = v
	return s
}

type CreatePredefinedDocumentRequest struct {
	Chunks []*CreatePredefinedDocumentRequestChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	// example:
	//
	// a1b2c3
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// {"a": "1"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// 测试文档
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreatePredefinedDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentRequest) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentRequest) SetChunks(v []*CreatePredefinedDocumentRequestChunks) *CreatePredefinedDocumentRequest {
	s.Chunks = v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetLibraryId(v string) *CreatePredefinedDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetMetadata(v map[string]interface{}) *CreatePredefinedDocumentRequest {
	s.Metadata = v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetTitle(v string) *CreatePredefinedDocumentRequest {
	s.Title = &v
	return s
}

type CreatePredefinedDocumentRequestChunks struct {
	// example:
	//
	// {"a": "1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// 1
	ChunkOrder *int32 `json:"chunkOrder,omitempty" xml:"chunkOrder,omitempty"`
	// example:
	//
	// 这是一段测试文本
	ChunkText *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
}

func (s CreatePredefinedDocumentRequestChunks) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentRequestChunks) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkMeta(v map[string]interface{}) *CreatePredefinedDocumentRequestChunks {
	s.ChunkMeta = v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkOrder(v int32) *CreatePredefinedDocumentRequestChunks {
	s.ChunkOrder = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkText(v string) *CreatePredefinedDocumentRequestChunks {
	s.ChunkText = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkType(v string) *CreatePredefinedDocumentRequestChunks {
	s.ChunkType = &v
	return s
}

type CreatePredefinedDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 1782981430906818562
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0a06dfe617018288881568684e2937
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s CreatePredefinedDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentResponseBody) SetCost(v int64) *CreatePredefinedDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetData(v string) *CreatePredefinedDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetDataType(v string) *CreatePredefinedDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetErrCode(v string) *CreatePredefinedDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetMessage(v string) *CreatePredefinedDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetRequestId(v string) *CreatePredefinedDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetSuccess(v bool) *CreatePredefinedDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePredefinedDocumentResponseBody) SetTime(v string) *CreatePredefinedDocumentResponseBody {
	s.Time = &v
	return s
}

type CreatePredefinedDocumentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePredefinedDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePredefinedDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePredefinedDocumentResponse) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentResponse) SetHeaders(v map[string]*string) *CreatePredefinedDocumentResponse {
	s.Headers = v
	return s
}

func (s *CreatePredefinedDocumentResponse) SetStatusCode(v int32) *CreatePredefinedDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePredefinedDocumentResponse) SetBody(v *CreatePredefinedDocumentResponseBody) *CreatePredefinedDocumentResponse {
	s.Body = v
	return s
}

type DeleteDocumentRequest struct {
	// This parameter is required.
	DocIds []*string `json:"docIds,omitempty" xml:"docIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s DeleteDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentRequest) SetDocIds(v []*string) *DeleteDocumentRequest {
	s.DocIds = v
	return s
}

func (s *DeleteDocumentRequest) SetLibraryId(v string) *DeleteDocumentRequest {
	s.LibraryId = &v
	return s
}

type DeleteDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 67C7021A-D268-553D-8C15-A087B9604028
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s DeleteDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponseBody) SetCost(v int64) *DeleteDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetData(v bool) *DeleteDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetDataType(v string) *DeleteDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetErrCode(v string) *DeleteDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetMessage(v string) *DeleteDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetRequestId(v string) *DeleteDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetSuccess(v bool) *DeleteDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDocumentResponseBody) SetTime(v string) *DeleteDocumentResponseBody {
	s.Time = &v
	return s
}

type DeleteDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDocumentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDocumentResponse) SetHeaders(v map[string]*string) *DeleteDocumentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDocumentResponse) SetStatusCode(v int32) *DeleteDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDocumentResponse) SetBody(v *DeleteDocumentResponseBody) *DeleteDocumentResponse {
	s.Body = v
	return s
}

type DeleteLibraryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// skdfefxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s DeleteLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibraryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLibraryRequest) SetLibraryId(v string) *DeleteLibraryRequest {
	s.LibraryId = &v
	return s
}

type DeleteLibraryResponseBody struct {
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 30F6AD44-F078-540D-B5A5-1E519C8E9E6D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLibraryResponseBody) SetErrCode(v string) *DeleteLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetMessage(v string) *DeleteLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetRequestId(v string) *DeleteLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLibraryResponseBody) SetSuccess(v bool) *DeleteLibraryResponseBody {
	s.Success = &v
	return s
}

type DeleteLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibraryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLibraryResponse) SetHeaders(v map[string]*string) *DeleteLibraryResponse {
	s.Headers = v
	return s
}

func (s *DeleteLibraryResponse) SetStatusCode(v int32) *DeleteLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLibraryResponse) SetBody(v *DeleteLibraryResponseBody) *DeleteLibraryResponse {
	s.Body = v
	return s
}

type EvictTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s EvictTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s EvictTaskRequest) GoString() string {
	return s.String()
}

func (s *EvictTaskRequest) SetTaskId(v string) *EvictTaskRequest {
	s.TaskId = &v
	return s
}

type EvictTaskResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 17071319
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s EvictTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvictTaskResponseBody) GoString() string {
	return s.String()
}

func (s *EvictTaskResponseBody) SetCost(v int64) *EvictTaskResponseBody {
	s.Cost = &v
	return s
}

func (s *EvictTaskResponseBody) SetData(v string) *EvictTaskResponseBody {
	s.Data = &v
	return s
}

func (s *EvictTaskResponseBody) SetDataType(v string) *EvictTaskResponseBody {
	s.DataType = &v
	return s
}

func (s *EvictTaskResponseBody) SetErrCode(v string) *EvictTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *EvictTaskResponseBody) SetMessage(v string) *EvictTaskResponseBody {
	s.Message = &v
	return s
}

func (s *EvictTaskResponseBody) SetRequestId(v string) *EvictTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *EvictTaskResponseBody) SetSuccess(v bool) *EvictTaskResponseBody {
	s.Success = &v
	return s
}

func (s *EvictTaskResponseBody) SetTime(v string) *EvictTaskResponseBody {
	s.Time = &v
	return s
}

type EvictTaskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EvictTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EvictTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s EvictTaskResponse) GoString() string {
	return s.String()
}

func (s *EvictTaskResponse) SetHeaders(v map[string]*string) *EvictTaskResponse {
	s.Headers = v
	return s
}

func (s *EvictTaskResponse) SetStatusCode(v int32) *EvictTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *EvictTaskResponse) SetBody(v *EvictTaskResponseBody) *EvictTaskResponse {
	s.Body = v
	return s
}

type GetAppConfigResponseBody struct {
	// example:
	//
	// null
	Cost *int64                        `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetAppConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// EF4B5C9B-3BC8-5171-A47B-4C5CF3DC3258
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBody) SetCost(v int64) *GetAppConfigResponseBody {
	s.Cost = &v
	return s
}

func (s *GetAppConfigResponseBody) SetData(v *GetAppConfigResponseBodyData) *GetAppConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAppConfigResponseBody) SetDataType(v string) *GetAppConfigResponseBody {
	s.DataType = &v
	return s
}

func (s *GetAppConfigResponseBody) SetErrCode(v string) *GetAppConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAppConfigResponseBody) SetMessage(v string) *GetAppConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppConfigResponseBody) SetRequestId(v string) *GetAppConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppConfigResponseBody) SetSuccess(v bool) *GetAppConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppConfigResponseBody) SetTime(v string) *GetAppConfigResponseBody {
	s.Time = &v
	return s
}

type GetAppConfigResponseBodyData struct {
	EmbeddingTypeList         []map[string]*string `json:"embeddingTypeList,omitempty" xml:"embeddingTypeList,omitempty" type:"Repeated"`
	FrontendConfig            map[string]*bool     `json:"frontendConfig,omitempty" xml:"frontendConfig,omitempty"`
	LibraryDocumentStatusList []map[string]*string `json:"libraryDocumentStatusList,omitempty" xml:"libraryDocumentStatusList,omitempty" type:"Repeated"`
	LlmHelperTypeList         []map[string]*string `json:"llmHelperTypeList,omitempty" xml:"llmHelperTypeList,omitempty" type:"Repeated"`
	TextIndexCategoryList     []*string            `json:"textIndexCategoryList,omitempty" xml:"textIndexCategoryList,omitempty" type:"Repeated"`
	VectorIndexCategoryList   []*string            `json:"vectorIndexCategoryList,omitempty" xml:"vectorIndexCategoryList,omitempty" type:"Repeated"`
}

func (s GetAppConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponseBodyData) SetEmbeddingTypeList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.EmbeddingTypeList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetFrontendConfig(v map[string]*bool) *GetAppConfigResponseBodyData {
	s.FrontendConfig = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetLibraryDocumentStatusList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.LibraryDocumentStatusList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetLlmHelperTypeList(v []map[string]*string) *GetAppConfigResponseBodyData {
	s.LlmHelperTypeList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetTextIndexCategoryList(v []*string) *GetAppConfigResponseBodyData {
	s.TextIndexCategoryList = v
	return s
}

func (s *GetAppConfigResponseBodyData) SetVectorIndexCategoryList(v []*string) *GetAppConfigResponseBodyData {
	s.VectorIndexCategoryList = v
	return s
}

type GetAppConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAppConfigResponse) SetHeaders(v map[string]*string) *GetAppConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAppConfigResponse) SetStatusCode(v int32) *GetAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppConfigResponse) SetBody(v *GetAppConfigResponseBody) *GetAppConfigResponse {
	s.Body = v
	return s
}

type GetDocumentChunkListRequest struct {
	ChunkIdList []*string `json:"chunkIdList,omitempty" xml:"chunkIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 182364872346
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dsjgfdjgfxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
}

func (s GetDocumentChunkListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListRequest) SetChunkIdList(v []*string) *GetDocumentChunkListRequest {
	s.ChunkIdList = v
	return s
}

func (s *GetDocumentChunkListRequest) SetDocId(v string) *GetDocumentChunkListRequest {
	s.DocId = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetLibraryId(v string) *GetDocumentChunkListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetOrder(v string) *GetDocumentChunkListRequest {
	s.Order = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetOrderBy(v string) *GetDocumentChunkListRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetPage(v int32) *GetDocumentChunkListRequest {
	s.Page = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetPageSize(v int32) *GetDocumentChunkListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentChunkListRequest) SetSearchQuery(v string) *GetDocumentChunkListRequest {
	s.SearchQuery = &v
	return s
}

type GetDocumentChunkListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDocumentChunkListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2B8F6DC9-6FAF-576F-9095-CCD90FB2BDDF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDocumentChunkListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBody) SetCost(v int64) *GetDocumentChunkListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetData(v *GetDocumentChunkListResponseBodyData) *GetDocumentChunkListResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetDataType(v string) *GetDocumentChunkListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetErrCode(v string) *GetDocumentChunkListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetMessage(v string) *GetDocumentChunkListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetRequestId(v string) *GetDocumentChunkListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetSuccess(v bool) *GetDocumentChunkListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentChunkListResponseBody) SetTime(v string) *GetDocumentChunkListResponseBody {
	s.Time = &v
	return s
}

type GetDocumentChunkListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetDocumentChunkListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetDocumentChunkListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyData) SetCurrentPage(v int64) *GetDocumentChunkListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetPageSize(v int64) *GetDocumentChunkListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetRecords(v []*GetDocumentChunkListResponseBodyDataRecords) *GetDocumentChunkListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetTotalPages(v int64) *GetDocumentChunkListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyData) SetTotalRecords(v int64) *GetDocumentChunkListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetDocumentChunkListResponseBodyDataRecords struct {
	// example:
	//
	// 28377468263482764
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// oss-xxxx-hangzhou.com/test.pdf
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 8947387648356
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// jhsdvne
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 947538465
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*GetDocumentChunkListResponseBodyDataRecordsPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 9848346548365
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetDocumentChunkListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkMeta(v map[string]interface{}) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkMeta = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkOssUrl(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkOssUrl = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkText(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkText = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetChunkType(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.ChunkType = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetDocId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetFileType(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetLibraryId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetLibraryName(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.LibraryName = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetNextChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.NextChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetPos(v []*GetDocumentChunkListResponseBodyDataRecordsPos) *GetDocumentChunkListResponseBodyDataRecords {
	s.Pos = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetPreChunkId(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.PreChunkId = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetScore(v float32) *GetDocumentChunkListResponseBodyDataRecords {
	s.Score = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecords) SetTitle(v string) *GetDocumentChunkListResponseBodyDataRecords {
	s.Title = &v
	return s
}

type GetDocumentChunkListResponseBodyDataRecordsPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s GetDocumentChunkListResponseBodyDataRecordsPos) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponseBodyDataRecordsPos) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetAxisArray(v []*float64) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.AxisArray = v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetPage(v int32) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.Page = &v
	return s
}

func (s *GetDocumentChunkListResponseBodyDataRecordsPos) SetTextHighlightArea(v []*int32) *GetDocumentChunkListResponseBodyDataRecordsPos {
	s.TextHighlightArea = v
	return s
}

type GetDocumentChunkListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentChunkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentChunkListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentChunkListResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentChunkListResponse) SetHeaders(v map[string]*string) *GetDocumentChunkListResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentChunkListResponse) SetStatusCode(v int32) *GetDocumentChunkListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentChunkListResponse) SetBody(v *GetDocumentChunkListResponseBody) *GetDocumentChunkListResponse {
	s.Body = v
	return s
}

type GetDocumentListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDocumentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentListRequest) SetLibraryId(v string) *GetDocumentListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentListRequest) SetPage(v int32) *GetDocumentListRequest {
	s.Page = &v
	return s
}

func (s *GetDocumentListRequest) SetPageSize(v int32) *GetDocumentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentListRequest) SetStatus(v string) *GetDocumentListRequest {
	s.Status = &v
	return s
}

type GetDocumentListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDocumentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDocumentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBody) SetCost(v int64) *GetDocumentListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentListResponseBody) SetData(v *GetDocumentListResponseBodyData) *GetDocumentListResponseBody {
	s.Data = v
	return s
}

func (s *GetDocumentListResponseBody) SetDataType(v string) *GetDocumentListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentListResponseBody) SetErrCode(v string) *GetDocumentListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentListResponseBody) SetMessage(v string) *GetDocumentListResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentListResponseBody) SetRequestId(v string) *GetDocumentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentListResponseBody) SetSuccess(v bool) *GetDocumentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentListResponseBody) SetTime(v string) *GetDocumentListResponseBody {
	s.Time = &v
	return s
}

type GetDocumentListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetDocumentListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetDocumentListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBodyData) SetCurrentPage(v int64) *GetDocumentListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetPageSize(v int64) *GetDocumentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetRecords(v []*GetDocumentListResponseBodyDataRecords) *GetDocumentListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetDocumentListResponseBodyData) SetTotalPages(v int64) *GetDocumentListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetDocumentListResponseBodyData) SetTotalRecords(v int64) *GetDocumentListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetDocumentListResponseBodyDataRecords struct {
	// example:
	//
	// 8326748346
	DocId        *string                `json:"docId,omitempty" xml:"docId,omitempty"`
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// skjdhshbv
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// WaitRefresh
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// null
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDocumentListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponseBodyDataRecords) SetDocId(v string) *GetDocumentListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetDocumentMeta(v map[string]interface{}) *GetDocumentListResponseBodyDataRecords {
	s.DocumentMeta = v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetFileType(v string) *GetDocumentListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetGmtCreate(v string) *GetDocumentListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetGmtModified(v string) *GetDocumentListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetLibraryId(v string) *GetDocumentListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetStatusCode(v string) *GetDocumentListResponseBodyDataRecords {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetTitle(v string) *GetDocumentListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetDocumentListResponseBodyDataRecords) SetUrl(v string) *GetDocumentListResponseBodyDataRecords {
	s.Url = &v
	return s
}

type GetDocumentListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentListResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentListResponse) SetHeaders(v map[string]*string) *GetDocumentListResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentListResponse) SetStatusCode(v int32) *GetDocumentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentListResponse) SetBody(v *GetDocumentListResponseBody) *GetDocumentListResponse {
	s.Body = v
	return s
}

type GetDocumentUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12681367362
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s GetDocumentUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlRequest) SetDocumentId(v string) *GetDocumentUrlRequest {
	s.DocumentId = &v
	return s
}

type GetDocumentUrlResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// https://path_to_file
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 66249B43-8C2B-5EE7-AE78-B382306621C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDocumentUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlResponseBody) SetCost(v int64) *GetDocumentUrlResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetData(v string) *GetDocumentUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetDataType(v string) *GetDocumentUrlResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetErrCode(v string) *GetDocumentUrlResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetMessage(v string) *GetDocumentUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetRequestId(v string) *GetDocumentUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetSuccess(v bool) *GetDocumentUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetDocumentUrlResponseBody) SetTime(v string) *GetDocumentUrlResponseBody {
	s.Time = &v
	return s
}

type GetDocumentUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocumentUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocumentUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDocumentUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDocumentUrlResponse) SetHeaders(v map[string]*string) *GetDocumentUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDocumentUrlResponse) SetStatusCode(v int32) *GetDocumentUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocumentUrlResponse) SetBody(v *GetDocumentUrlResponseBody) *GetDocumentUrlResponse {
	s.Body = v
	return s
}

type GetFilterDocumentListRequest struct {
	And       []*GetFilterDocumentListRequestAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	DocIdList []*string                          `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string                           `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	Or        []*GetFilterDocumentListRequestOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s GetFilterDocumentListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListRequest) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequest) SetAnd(v []*GetFilterDocumentListRequestAnd) *GetFilterDocumentListRequest {
	s.And = v
	return s
}

func (s *GetFilterDocumentListRequest) SetDocIdList(v []*string) *GetFilterDocumentListRequest {
	s.DocIdList = v
	return s
}

func (s *GetFilterDocumentListRequest) SetLibraryId(v string) *GetFilterDocumentListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetOr(v []*GetFilterDocumentListRequestOr) *GetFilterDocumentListRequest {
	s.Or = v
	return s
}

func (s *GetFilterDocumentListRequest) SetPage(v int32) *GetFilterDocumentListRequest {
	s.Page = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetPageSize(v int32) *GetFilterDocumentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetFilterDocumentListRequest) SetStatus(v []*string) *GetFilterDocumentListRequest {
	s.Status = v
	return s
}

type GetFilterDocumentListRequestAnd struct {
	// example:
	//
	// 1
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// company
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// alibaba
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFilterDocumentListRequestAnd) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListRequestAnd) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequestAnd) SetBoost(v float32) *GetFilterDocumentListRequestAnd {
	s.Boost = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetKey(v string) *GetFilterDocumentListRequestAnd {
	s.Key = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetOperator(v string) *GetFilterDocumentListRequestAnd {
	s.Operator = &v
	return s
}

func (s *GetFilterDocumentListRequestAnd) SetValue(v string) *GetFilterDocumentListRequestAnd {
	s.Value = &v
	return s
}

type GetFilterDocumentListRequestOr struct {
	// example:
	//
	// 1
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// company
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// alibaba
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetFilterDocumentListRequestOr) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListRequestOr) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListRequestOr) SetBoost(v float32) *GetFilterDocumentListRequestOr {
	s.Boost = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetKey(v string) *GetFilterDocumentListRequestOr {
	s.Key = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetOperator(v string) *GetFilterDocumentListRequestOr {
	s.Operator = &v
	return s
}

func (s *GetFilterDocumentListRequestOr) SetValue(v string) *GetFilterDocumentListRequestOr {
	s.Value = &v
	return s
}

type GetFilterDocumentListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                 `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetFilterDocumentListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7ADF010C-FD89-569D-A079-2D4D5247E943
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetFilterDocumentListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponseBody) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBody) SetCost(v int64) *GetFilterDocumentListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetData(v *GetFilterDocumentListResponseBodyData) *GetFilterDocumentListResponseBody {
	s.Data = v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetDataType(v string) *GetFilterDocumentListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetErrCode(v string) *GetFilterDocumentListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetMessage(v string) *GetFilterDocumentListResponseBody {
	s.Message = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetRequestId(v string) *GetFilterDocumentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetSuccess(v bool) *GetFilterDocumentListResponseBody {
	s.Success = &v
	return s
}

func (s *GetFilterDocumentListResponseBody) SetTime(v string) *GetFilterDocumentListResponseBody {
	s.Time = &v
	return s
}

type GetFilterDocumentListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                          `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetFilterDocumentListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetFilterDocumentListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBodyData) SetCurrentPage(v int64) *GetFilterDocumentListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetPageSize(v int64) *GetFilterDocumentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetRecords(v []*GetFilterDocumentListResponseBodyDataRecords) *GetFilterDocumentListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetTotalPages(v int64) *GetFilterDocumentListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyData) SetTotalRecords(v int64) *GetFilterDocumentListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetFilterDocumentListResponseBodyDataRecords struct {
	// example:
	//
	// 29368126816
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// {"a": "1"}
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// sdfgsjdfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// WaitRefresh
	StatusCode *string `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Title      *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// null
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetFilterDocumentListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetDocId(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.DocId = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetDocumentMeta(v map[string]interface{}) *GetFilterDocumentListResponseBodyDataRecords {
	s.DocumentMeta = v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetFileType(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.FileType = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetGmtCreate(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetGmtModified(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetLibraryId(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.LibraryId = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetStatusCode(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.StatusCode = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetTitle(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.Title = &v
	return s
}

func (s *GetFilterDocumentListResponseBodyDataRecords) SetUrl(v string) *GetFilterDocumentListResponseBodyDataRecords {
	s.Url = &v
	return s
}

type GetFilterDocumentListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFilterDocumentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFilterDocumentListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFilterDocumentListResponse) GoString() string {
	return s.String()
}

func (s *GetFilterDocumentListResponse) SetHeaders(v map[string]*string) *GetFilterDocumentListResponse {
	s.Headers = v
	return s
}

func (s *GetFilterDocumentListResponse) SetStatusCode(v int32) *GetFilterDocumentListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFilterDocumentListResponse) SetBody(v *GetFilterDocumentListResponseBody) *GetFilterDocumentListResponse {
	s.Body = v
	return s
}

type GetHistoryListByBizTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GysYBsxx
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LibraryChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetHistoryListByBizTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeRequest) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeRequest) SetBizId(v string) *GetHistoryListByBizTypeRequest {
	s.BizId = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetBizType(v string) *GetHistoryListByBizTypeRequest {
	s.BizType = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetPage(v int32) *GetHistoryListByBizTypeRequest {
	s.Page = &v
	return s
}

func (s *GetHistoryListByBizTypeRequest) SetPageSize(v int32) *GetHistoryListByBizTypeRequest {
	s.PageSize = &v
	return s
}

type GetHistoryListByBizTypeResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                   `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetHistoryListByBizTypeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9DF9B3F3-9FFE-52CB-A8DC-F7BD5F842F0E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetHistoryListByBizTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponseBody) SetCost(v int64) *GetHistoryListByBizTypeResponseBody {
	s.Cost = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetData(v *GetHistoryListByBizTypeResponseBodyData) *GetHistoryListByBizTypeResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetDataType(v string) *GetHistoryListByBizTypeResponseBody {
	s.DataType = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetErrCode(v string) *GetHistoryListByBizTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetMessage(v string) *GetHistoryListByBizTypeResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetRequestId(v string) *GetHistoryListByBizTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetSuccess(v bool) *GetHistoryListByBizTypeResponseBody {
	s.Success = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBody) SetTime(v string) *GetHistoryListByBizTypeResponseBody {
	s.Time = &v
	return s
}

type GetHistoryListByBizTypeResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetHistoryListByBizTypeResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetHistoryListByBizTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetCurrentPage(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetPageSize(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetRecords(v []*GetHistoryListByBizTypeResponseBodyDataRecords) *GetHistoryListByBizTypeResponseBodyData {
	s.Records = v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetTotalPages(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyData) SetTotalRecords(v int64) *GetHistoryListByBizTypeResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetHistoryListByBizTypeResponseBodyDataRecords struct {
	// example:
	//
	// GysYBsxx
	BizId *string `json:"bizId,omitempty" xml:"bizId,omitempty"`
	// example:
	//
	// LibraryChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// example:
	//
	// null
	ExtraMessage interface{} `json:"extraMessage,omitempty" xml:"extraMessage,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 210
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LlmAnswer *string `json:"llmAnswer,omitempty" xml:"llmAnswer,omitempty"`
	LlmPrompt *string `json:"llmPrompt,omitempty" xml:"llmPrompt,omitempty"`
	// example:
	//
	// qwen-max
	LlmType *string `json:"llmType,omitempty" xml:"llmType,omitempty"`
	// example:
	//
	// null
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	UserQuery *string `json:"userQuery,omitempty" xml:"userQuery,omitempty"`
}

func (s GetHistoryListByBizTypeResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetBizId(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.BizId = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetBizType(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.BizType = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetExtraMessage(v interface{}) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.ExtraMessage = v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetGmtCreate(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetGmtModified(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetId(v int64) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetLlmAnswer(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.LlmAnswer = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetLlmPrompt(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.LlmPrompt = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetLlmType(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.LlmType = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetSessionId(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.SessionId = &v
	return s
}

func (s *GetHistoryListByBizTypeResponseBodyDataRecords) SetUserQuery(v string) *GetHistoryListByBizTypeResponseBodyDataRecords {
	s.UserQuery = &v
	return s
}

type GetHistoryListByBizTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryListByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryListByBizTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoryListByBizTypeResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponse) SetHeaders(v map[string]*string) *GetHistoryListByBizTypeResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryListByBizTypeResponse) SetStatusCode(v int32) *GetHistoryListByBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryListByBizTypeResponse) SetBody(v *GetHistoryListByBizTypeResponseBody) *GetHistoryListByBizTypeResponse {
	s.Body = v
	return s
}

type GetLibraryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s GetLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryRequest) SetLibraryId(v string) *GetLibraryRequest {
	s.LibraryId = &v
	return s
}

type GetLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64                      `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetLibraryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 44BD277A-87F9-5310-8D63-3E6645F1DA85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBody) SetCost(v int64) *GetLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *GetLibraryResponseBody) SetData(v *GetLibraryResponseBodyData) *GetLibraryResponseBody {
	s.Data = v
	return s
}

func (s *GetLibraryResponseBody) SetDataType(v string) *GetLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *GetLibraryResponseBody) SetErrCode(v string) *GetLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetLibraryResponseBody) SetMessage(v string) *GetLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *GetLibraryResponseBody) SetRequestId(v string) *GetLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryResponseBody) SetSuccess(v bool) *GetLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *GetLibraryResponseBody) SetTime(v string) *GetLibraryResponseBody {
	s.Time = &v
	return s
}

type GetLibraryResponseBodyData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 10
	DocumentCount *int64 `json:"documentCount,omitempty" xml:"documentCount,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 19386728376
	Id           *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	IndexSetting *GetLibraryResponseBodyDataIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	LibraryName  *string                                 `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s GetLibraryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyData) SetDescription(v string) *GetLibraryResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetDocumentCount(v int64) *GetLibraryResponseBodyData {
	s.DocumentCount = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetGmtCreate(v string) *GetLibraryResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetGmtModified(v string) *GetLibraryResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetId(v string) *GetLibraryResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetLibraryResponseBodyData) SetIndexSetting(v *GetLibraryResponseBodyDataIndexSetting) *GetLibraryResponseBodyData {
	s.IndexSetting = v
	return s
}

func (s *GetLibraryResponseBodyData) SetLibraryName(v string) *GetLibraryResponseBodyData {
	s.LibraryName = &v
	return s
}

type GetLibraryResponseBodyDataIndexSetting struct {
	ChunkStrategy      *GetLibraryResponseBodyDataIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *GetLibraryResponseBodyDataIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                                   `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *GetLibraryResponseBodyDataIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *GetLibraryResponseBodyDataIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *GetLibraryResponseBodyDataIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s GetLibraryResponseBodyDataIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetChunkStrategy(v *GetLibraryResponseBodyDataIndexSettingChunkStrategy) *GetLibraryResponseBodyDataIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetModelConfig(v *GetLibraryResponseBodyDataIndexSettingModelConfig) *GetLibraryResponseBodyDataIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetPromptRoleStyle(v string) *GetLibraryResponseBodyDataIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetQueryEnhancer(v *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) *GetLibraryResponseBodyDataIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetRecallStrategy(v *GetLibraryResponseBodyDataIndexSettingRecallStrategy) *GetLibraryResponseBodyDataIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetTextIndexSetting(v *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) *GetLibraryResponseBodyDataIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSetting) SetVectorIndexSetting(v *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) *GetLibraryResponseBodyDataIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type GetLibraryResponseBodyDataIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 160
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 40
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 160
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 256
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetEnhanceTable(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetOverlap(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSentenceSplit(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSize(v int32) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) SetSplit(v bool) *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// topP
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) SetTemperature(v float64) *GetLibraryResponseBodyDataIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) SetTopP(v float64) *GetLibraryResponseBodyDataIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// 2836482634
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetEnableSession(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) SetDocumentRankType(v string) *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) SetLimit(v int32) *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// IkMaxWord
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetCategory(v string) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetEnable(v bool) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetRankThreshold(v float64) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) SetTopK(v int32) *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryResponseBodyDataIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetCategory(v string) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetEnable(v bool) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) SetTopK(v int32) *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryResponse) SetHeaders(v map[string]*string) *GetLibraryResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryResponse) SetStatusCode(v int32) *GetLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryResponse) SetBody(v *GetLibraryResponseBody) *GetLibraryResponse {
	s.Body = v
	return s
}

type GetLibraryListRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query    *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetLibraryListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryListRequest) SetPage(v int32) *GetLibraryListRequest {
	s.Page = &v
	return s
}

func (s *GetLibraryListRequest) SetPageSize(v int32) *GetLibraryListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLibraryListRequest) SetQuery(v string) *GetLibraryListRequest {
	s.Query = &v
	return s
}

type GetLibraryListResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetLibraryListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0a06dfe817156528535968405edce3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetLibraryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBody) SetCost(v int64) *GetLibraryListResponseBody {
	s.Cost = &v
	return s
}

func (s *GetLibraryListResponseBody) SetData(v *GetLibraryListResponseBodyData) *GetLibraryListResponseBody {
	s.Data = v
	return s
}

func (s *GetLibraryListResponseBody) SetDataType(v string) *GetLibraryListResponseBody {
	s.DataType = &v
	return s
}

func (s *GetLibraryListResponseBody) SetErrCode(v string) *GetLibraryListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetLibraryListResponseBody) SetMessage(v string) *GetLibraryListResponseBody {
	s.Message = &v
	return s
}

func (s *GetLibraryListResponseBody) SetRequestId(v string) *GetLibraryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryListResponseBody) SetSuccess(v bool) *GetLibraryListResponseBody {
	s.Success = &v
	return s
}

func (s *GetLibraryListResponseBody) SetTime(v string) *GetLibraryListResponseBody {
	s.Time = &v
	return s
}

type GetLibraryListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Records  []*GetLibraryListResponseBodyDataRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalPages *int64 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
	// example:
	//
	// 100
	TotalRecords *int64 `json:"totalRecords,omitempty" xml:"totalRecords,omitempty"`
}

func (s GetLibraryListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyData) SetCurrentPage(v int64) *GetLibraryListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetLibraryListResponseBodyData) SetPageSize(v int64) *GetLibraryListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetLibraryListResponseBodyData) SetRecords(v []*GetLibraryListResponseBodyDataRecords) *GetLibraryListResponseBodyData {
	s.Records = v
	return s
}

func (s *GetLibraryListResponseBodyData) SetTotalPages(v int64) *GetLibraryListResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetLibraryListResponseBodyData) SetTotalRecords(v int64) *GetLibraryListResponseBodyData {
	s.TotalRecords = &v
	return s
}

type GetLibraryListResponseBodyDataRecords struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 10
	DocumentCount *int64 `json:"documentCount,omitempty" xml:"documentCount,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 24vs4aa42jv1rg7
	Id           *string                                            `json:"id,omitempty" xml:"id,omitempty"`
	IndexSetting *GetLibraryListResponseBodyDataRecordsIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	LibraryName  *string                                            `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecords) SetDescription(v string) *GetLibraryListResponseBodyDataRecords {
	s.Description = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetDocumentCount(v int64) *GetLibraryListResponseBodyDataRecords {
	s.DocumentCount = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetGmtCreate(v string) *GetLibraryListResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetGmtModified(v string) *GetLibraryListResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetId(v string) *GetLibraryListResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetIndexSetting(v *GetLibraryListResponseBodyDataRecordsIndexSetting) *GetLibraryListResponseBodyDataRecords {
	s.IndexSetting = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecords) SetLibraryName(v string) *GetLibraryListResponseBodyDataRecords {
	s.LibraryName = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSetting struct {
	ChunkStrategy      *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                                              `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetChunkStrategy(v *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetModelConfig(v *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetPromptRoleStyle(v string) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetQueryEnhancer(v *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetRecallStrategy(v *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetTextIndexSetting(v *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) SetVectorIndexSetting(v *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) *GetLibraryListResponseBodyDataRecordsIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 160
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 40
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 160
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 256
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetEnhanceTable(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetOverlap(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSentenceSplit(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSize(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) SetSplit(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// topP
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) SetTemperature(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) SetTopP(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// sdbcjsbc
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetEnableSession(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) SetDocumentRankType(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) SetLimit(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Standard
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetCategory(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetEnable(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetRankThreshold(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) SetTopK(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// null
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetCategory(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetEnable(v bool) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) SetTopK(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type GetLibraryListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLibraryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryListResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponse) SetHeaders(v map[string]*string) *GetLibraryListResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryListResponse) SetStatusCode(v int32) *GetLibraryListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryListResponse) SetBody(v *GetLibraryListResponseBody) *GetLibraryListResponse {
	s.Body = v
	return s
}

type GetParseResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 873648346573245
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s GetParseResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultRequest) GoString() string {
	return s.String()
}

func (s *GetParseResultRequest) SetDocId(v string) *GetParseResultRequest {
	s.DocId = &v
	return s
}

func (s *GetParseResultRequest) SetLibraryId(v string) *GetParseResultRequest {
	s.LibraryId = &v
	return s
}

type GetParseResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetParseResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0abb793617204049360065953ec6dd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetParseResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetParseResultResponseBody) SetCost(v int64) *GetParseResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetParseResultResponseBody) SetData(v *GetParseResultResponseBodyData) *GetParseResultResponseBody {
	s.Data = v
	return s
}

func (s *GetParseResultResponseBody) SetDataType(v string) *GetParseResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetParseResultResponseBody) SetErrCode(v string) *GetParseResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetParseResultResponseBody) SetMessage(v string) *GetParseResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetParseResultResponseBody) SetRequestId(v string) *GetParseResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParseResultResponseBody) SetSuccess(v bool) *GetParseResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetParseResultResponseBody) SetTime(v string) *GetParseResultResponseBody {
	s.Time = &v
	return s
}

type GetParseResultResponseBodyData struct {
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// null
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	// example:
	//
	// b0a202e2-5031-4589-a6d7-39185f0d8d01
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {
	//
	//           "Status": "Success",
	//
	//           "Data": {},
	//
	//           "Message": null,
	//
	//           "TaskId": "docmind-20240601-123abc"
	//
	//         }
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// WaitRefresh
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetParseResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetParseResultResponseBodyData) SetFileType(v string) *GetParseResultResponseBodyData {
	s.FileType = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetProviderType(v string) *GetParseResultResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetRequestId(v string) *GetParseResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetParseResultResponseBodyData) SetResult(v map[string]interface{}) *GetParseResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetParseResultResponseBodyData) SetStatus(v string) *GetParseResultResponseBodyData {
	s.Status = &v
	return s
}

type GetParseResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParseResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParseResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParseResultResponse) GoString() string {
	return s.String()
}

func (s *GetParseResultResponse) SetHeaders(v map[string]*string) *GetParseResultResponse {
	s.Headers = v
	return s
}

func (s *GetParseResultResponse) SetStatusCode(v int32) *GetParseResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParseResultResponse) SetBody(v *GetParseResultResponseBody) *GetParseResultResponse {
	s.Body = v
	return s
}

type GetSummaryTaskResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 17071319
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetSummaryTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultRequest) SetTaskId(v string) *GetSummaryTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetSummaryTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetSummaryTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetSummaryTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBody) SetCost(v int64) *GetSummaryTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetData(v *GetSummaryTaskResultResponseBodyData) *GetSummaryTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetDataType(v string) *GetSummaryTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetErrCode(v string) *GetSummaryTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetMessage(v string) *GetSummaryTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetRequestId(v string) *GetSummaryTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetSuccess(v bool) *GetSummaryTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSummaryTaskResultResponseBody) SetTime(v string) *GetSummaryTaskResultResponseBody {
	s.Time = &v
	return s
}

type GetSummaryTaskResultResponseBodyData struct {
	Choices []*GetSummaryTaskResultResponseBodyDataChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1726285125915
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// 1202
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32                                     `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage       *GetSummaryTaskResultResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetSummaryTaskResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyData) SetChoices(v []*GetSummaryTaskResultResponseBodyDataChoices) *GetSummaryTaskResultResponseBodyData {
	s.Choices = v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetCreated(v int64) *GetSummaryTaskResultResponseBodyData {
	s.Created = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetId(v string) *GetSummaryTaskResultResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetModelId(v string) *GetSummaryTaskResultResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetRequestId(v string) *GetSummaryTaskResultResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetTime(v string) *GetSummaryTaskResultResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetTotalTokens(v int32) *GetSummaryTaskResultResponseBodyData {
	s.TotalTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyData) SetUsage(v *GetSummaryTaskResultResponseBodyDataUsage) *GetSummaryTaskResultResponseBodyData {
	s.Usage = v
	return s
}

type GetSummaryTaskResultResponseBodyDataChoices struct {
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                              `json:"index,omitempty" xml:"index,omitempty"`
	Message *GetSummaryTaskResultResponseBodyDataChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s GetSummaryTaskResultResponseBodyDataChoices) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataChoices) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetFinishReason(v string) *GetSummaryTaskResultResponseBodyDataChoices {
	s.FinishReason = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetIndex(v int32) *GetSummaryTaskResultResponseBodyDataChoices {
	s.Index = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoices) SetMessage(v *GetSummaryTaskResultResponseBodyDataChoicesMessage) *GetSummaryTaskResultResponseBodyDataChoices {
	s.Message = v
	return s
}

type GetSummaryTaskResultResponseBodyDataChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// assistant
	Role      *string                  `json:"role,omitempty" xml:"role,omitempty"`
	ToolCalls []map[string]interface{} `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s GetSummaryTaskResultResponseBodyDataChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataChoicesMessage) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetContent(v string) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.Content = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetRole(v string) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.Role = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataChoicesMessage) SetToolCalls(v []map[string]interface{}) *GetSummaryTaskResultResponseBodyDataChoicesMessage {
	s.ToolCalls = v
	return s
}

type GetSummaryTaskResultResponseBodyDataUsage struct {
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 0
	ImageTokens *int32 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 100
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 200
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetSummaryTaskResultResponseBodyDataUsage) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetImageCount(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.ImageCount = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetImageTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.ImageTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetInputTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.InputTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetOutputTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetSummaryTaskResultResponseBodyDataUsage) SetTotalTokens(v int32) *GetSummaryTaskResultResponseBodyDataUsage {
	s.TotalTokens = &v
	return s
}

type GetSummaryTaskResultResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSummaryTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSummaryTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSummaryTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetSummaryTaskResultResponse) SetHeaders(v map[string]*string) *GetSummaryTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetSummaryTaskResultResponse) SetStatusCode(v int32) *GetSummaryTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSummaryTaskResultResponse) SetBody(v *GetSummaryTaskResultResponseBody) *GetSummaryTaskResultResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	Cost      *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	DataType  *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Time      *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetCost(v int64) *GetTaskStatusResponseBody {
	s.Cost = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetData(v string) *GetTaskStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetDataType(v string) *GetTaskStatusResponseBody {
	s.DataType = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrCode(v string) *GetTaskStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTime(v string) *GetTaskStatusResponseBody {
	s.Time = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type InvokePluginRequest struct {
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// 3mj87da7zr
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s InvokePluginRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokePluginRequest) GoString() string {
	return s.String()
}

func (s *InvokePluginRequest) SetParams(v map[string]interface{}) *InvokePluginRequest {
	s.Params = v
	return s
}

func (s *InvokePluginRequest) SetPluginId(v string) *InvokePluginRequest {
	s.PluginId = &v
	return s
}

type InvokePluginResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// {\\"jobWaiting\\": [0, 0], \\"timestamps\\": [1713383820, 1713383880], \\"jobUsage\\": [0, 0], \\"quotaUsage\\": [123, 32]}
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 915AAAB9-4908-5224-9E53-9E9D7D0AA94B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s InvokePluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokePluginResponseBody) GoString() string {
	return s.String()
}

func (s *InvokePluginResponseBody) SetCost(v int64) *InvokePluginResponseBody {
	s.Cost = &v
	return s
}

func (s *InvokePluginResponseBody) SetData(v map[string]interface{}) *InvokePluginResponseBody {
	s.Data = v
	return s
}

func (s *InvokePluginResponseBody) SetDataType(v string) *InvokePluginResponseBody {
	s.DataType = &v
	return s
}

func (s *InvokePluginResponseBody) SetErrCode(v string) *InvokePluginResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InvokePluginResponseBody) SetMessage(v string) *InvokePluginResponseBody {
	s.Message = &v
	return s
}

func (s *InvokePluginResponseBody) SetRequestId(v string) *InvokePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokePluginResponseBody) SetSuccess(v bool) *InvokePluginResponseBody {
	s.Success = &v
	return s
}

func (s *InvokePluginResponseBody) SetTime(v string) *InvokePluginResponseBody {
	s.Time = &v
	return s
}

type InvokePluginResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokePluginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokePluginResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokePluginResponse) GoString() string {
	return s.String()
}

func (s *InvokePluginResponse) SetHeaders(v map[string]*string) *InvokePluginResponse {
	s.Headers = v
	return s
}

func (s *InvokePluginResponse) SetStatusCode(v int32) *InvokePluginResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokePluginResponse) SetBody(v *InvokePluginResponseBody) *InvokePluginResponse {
	s.Body = v
	return s
}

type PreviewDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8326472354762354
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s PreviewDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentRequest) GoString() string {
	return s.String()
}

func (s *PreviewDocumentRequest) SetDocumentId(v string) *PreviewDocumentRequest {
	s.DocumentId = &v
	return s
}

type PreviewDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *PreviewDocumentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// ff551395-1c8a-4f30-8ffd-ef7e87c70b4c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s PreviewDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponseBody) SetCost(v int64) *PreviewDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetData(v *PreviewDocumentResponseBodyData) *PreviewDocumentResponseBody {
	s.Data = v
	return s
}

func (s *PreviewDocumentResponseBody) SetDataType(v string) *PreviewDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetErrCode(v string) *PreviewDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetMessage(v string) *PreviewDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetRequestId(v string) *PreviewDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetSuccess(v bool) *PreviewDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *PreviewDocumentResponseBody) SetTime(v string) *PreviewDocumentResponseBody {
	s.Time = &v
	return s
}

type PreviewDocumentResponseBodyData struct {
	// example:
	//
	// pdf
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	UploadTime *string `json:"uploadTime,omitempty" xml:"uploadTime,omitempty"`
	// example:
	//
	// https://agi.alicdn.com/user/d0o/d3c1f50d-a6c2-49b3-b0c8-3e613c3f20ee_16872_3236784461.png
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s PreviewDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponseBodyData) SetPreviewType(v string) *PreviewDocumentResponseBodyData {
	s.PreviewType = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetTitle(v string) *PreviewDocumentResponseBodyData {
	s.Title = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetUploadTime(v string) *PreviewDocumentResponseBodyData {
	s.UploadTime = &v
	return s
}

func (s *PreviewDocumentResponseBodyData) SetUrl(v string) *PreviewDocumentResponseBodyData {
	s.Url = &v
	return s
}

type PreviewDocumentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewDocumentResponse) GoString() string {
	return s.String()
}

func (s *PreviewDocumentResponse) SetHeaders(v map[string]*string) *PreviewDocumentResponse {
	s.Headers = v
	return s
}

func (s *PreviewDocumentResponse) SetStatusCode(v int32) *PreviewDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewDocumentResponse) SetBody(v *PreviewDocumentResponseBody) *PreviewDocumentResponse {
	s.Body = v
	return s
}

type ReIndexRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8326472354762354
	DocumentId *string `json:"documentId,omitempty" xml:"documentId,omitempty"`
}

func (s ReIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s ReIndexRequest) GoString() string {
	return s.String()
}

func (s *ReIndexRequest) SetDocumentId(v string) *ReIndexRequest {
	s.DocumentId = &v
	return s
}

type ReIndexResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// True
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 32FFC91D-0A9F-585A-B84F-8A54C5187035
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ReIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReIndexResponseBody) GoString() string {
	return s.String()
}

func (s *ReIndexResponseBody) SetCost(v int64) *ReIndexResponseBody {
	s.Cost = &v
	return s
}

func (s *ReIndexResponseBody) SetData(v string) *ReIndexResponseBody {
	s.Data = &v
	return s
}

func (s *ReIndexResponseBody) SetDataType(v string) *ReIndexResponseBody {
	s.DataType = &v
	return s
}

func (s *ReIndexResponseBody) SetErrCode(v string) *ReIndexResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReIndexResponseBody) SetMessage(v string) *ReIndexResponseBody {
	s.Message = &v
	return s
}

func (s *ReIndexResponseBody) SetRequestId(v string) *ReIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReIndexResponseBody) SetSuccess(v bool) *ReIndexResponseBody {
	s.Success = &v
	return s
}

func (s *ReIndexResponseBody) SetTime(v string) *ReIndexResponseBody {
	s.Time = &v
	return s
}

type ReIndexResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s ReIndexResponse) GoString() string {
	return s.String()
}

func (s *ReIndexResponse) SetHeaders(v map[string]*string) *ReIndexResponse {
	s.Headers = v
	return s
}

func (s *ReIndexResponse) SetStatusCode(v int32) *ReIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *ReIndexResponse) SetBody(v *ReIndexResponseBody) *ReIndexResponse {
	s.Body = v
	return s
}

type RecallDocumentRequest struct {
	Filters []*RecallDocumentRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// false
	Rearrangement *bool `json:"rearrangement,omitempty" xml:"rearrangement,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s RecallDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequest) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequest) SetFilters(v []*RecallDocumentRequestFilters) *RecallDocumentRequest {
	s.Filters = v
	return s
}

func (s *RecallDocumentRequest) SetQuery(v string) *RecallDocumentRequest {
	s.Query = &v
	return s
}

func (s *RecallDocumentRequest) SetRearrangement(v bool) *RecallDocumentRequest {
	s.Rearrangement = &v
	return s
}

func (s *RecallDocumentRequest) SetTopK(v int32) *RecallDocumentRequest {
	s.TopK = &v
	return s
}

type RecallDocumentRequestFilters struct {
	And []*RecallDocumentRequestFiltersAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	// example:
	//
	// Text
	ChunkType *string   `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sdbjhvs
	LibraryId *string                           `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	Or        []*RecallDocumentRequestFiltersOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	Status    []*string                         `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s RecallDocumentRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequestFilters) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFilters) SetAnd(v []*RecallDocumentRequestFiltersAnd) *RecallDocumentRequestFilters {
	s.And = v
	return s
}

func (s *RecallDocumentRequestFilters) SetChunkType(v string) *RecallDocumentRequestFilters {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentRequestFilters) SetDocIdList(v []*string) *RecallDocumentRequestFilters {
	s.DocIdList = v
	return s
}

func (s *RecallDocumentRequestFilters) SetLibraryId(v string) *RecallDocumentRequestFilters {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentRequestFilters) SetOr(v []*RecallDocumentRequestFiltersOr) *RecallDocumentRequestFilters {
	s.Or = v
	return s
}

func (s *RecallDocumentRequestFilters) SetStatus(v []*string) *RecallDocumentRequestFilters {
	s.Status = v
	return s
}

type RecallDocumentRequestFiltersAnd struct {
	// example:
	//
	// 20
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// docType
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallDocumentRequestFiltersAnd) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequestFiltersAnd) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFiltersAnd) SetBoost(v float32) *RecallDocumentRequestFiltersAnd {
	s.Boost = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetKey(v string) *RecallDocumentRequestFiltersAnd {
	s.Key = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetOperator(v string) *RecallDocumentRequestFiltersAnd {
	s.Operator = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetValue(v string) *RecallDocumentRequestFiltersAnd {
	s.Value = &v
	return s
}

type RecallDocumentRequestFiltersOr struct {
	// example:
	//
	// 30
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// researcher
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallDocumentRequestFiltersOr) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentRequestFiltersOr) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFiltersOr) SetBoost(v float32) *RecallDocumentRequestFiltersOr {
	s.Boost = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetKey(v string) *RecallDocumentRequestFiltersOr {
	s.Key = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetOperator(v string) *RecallDocumentRequestFiltersOr {
	s.Operator = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetValue(v string) *RecallDocumentRequestFiltersOr {
	s.Value = &v
	return s
}

type RecallDocumentResponseBody struct {
	// example:
	//
	// 0
	Cost *int64                          `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RecallDocumentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0bc13a9517168617617186457e401f
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RecallDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBody) SetCost(v int64) *RecallDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *RecallDocumentResponseBody) SetData(v *RecallDocumentResponseBodyData) *RecallDocumentResponseBody {
	s.Data = v
	return s
}

func (s *RecallDocumentResponseBody) SetDataType(v string) *RecallDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *RecallDocumentResponseBody) SetErrCode(v string) *RecallDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RecallDocumentResponseBody) SetMessage(v string) *RecallDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *RecallDocumentResponseBody) SetRequestId(v string) *RecallDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecallDocumentResponseBody) SetSuccess(v bool) *RecallDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *RecallDocumentResponseBody) SetTime(v string) *RecallDocumentResponseBody {
	s.Time = &v
	return s
}

type RecallDocumentResponseBodyData struct {
	ChunkList     []*RecallDocumentResponseBodyDataChunkList     `json:"chunkList,omitempty" xml:"chunkList,omitempty" type:"Repeated"`
	ChunkPartList []*RecallDocumentResponseBodyDataChunkPartList `json:"chunkPartList,omitempty" xml:"chunkPartList,omitempty" type:"Repeated"`
	ChunkTextList []*string                                      `json:"chunkTextList,omitempty" xml:"chunkTextList,omitempty" type:"Repeated"`
	Documents     []*RecallDocumentResponseBodyDataDocuments     `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	EmbeddingElapsedMs *int64                                         `json:"embeddingElapsedMs,omitempty" xml:"embeddingElapsedMs,omitempty"`
	TextChunkList      []*RecallDocumentResponseBodyDataTextChunkList `json:"textChunkList,omitempty" xml:"textChunkList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TextSearchElapsedMs *int64 `json:"textSearchElapsedMs,omitempty" xml:"textSearchElapsedMs,omitempty"`
	// example:
	//
	// 400
	TotalElapsedMs  *int64                                           `json:"totalElapsedMs,omitempty" xml:"totalElapsedMs,omitempty"`
	VectorChunkList []*RecallDocumentResponseBodyDataVectorChunkList `json:"vectorChunkList,omitempty" xml:"vectorChunkList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	VectorSearchElapsedMs *int64 `json:"vectorSearchElapsedMs,omitempty" xml:"vectorSearchElapsedMs,omitempty"`
}

func (s RecallDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyData) SetChunkList(v []*RecallDocumentResponseBodyDataChunkList) *RecallDocumentResponseBodyData {
	s.ChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetChunkPartList(v []*RecallDocumentResponseBodyDataChunkPartList) *RecallDocumentResponseBodyData {
	s.ChunkPartList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetChunkTextList(v []*string) *RecallDocumentResponseBodyData {
	s.ChunkTextList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetDocuments(v []*RecallDocumentResponseBodyDataDocuments) *RecallDocumentResponseBodyData {
	s.Documents = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetEmbeddingElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.EmbeddingElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTextChunkList(v []*RecallDocumentResponseBodyDataTextChunkList) *RecallDocumentResponseBodyData {
	s.TextChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTextSearchElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.TextSearchElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetTotalElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.TotalElapsedMs = &v
	return s
}

func (s *RecallDocumentResponseBodyData) SetVectorChunkList(v []*RecallDocumentResponseBodyDataVectorChunkList) *RecallDocumentResponseBodyData {
	s.VectorChunkList = v
	return s
}

func (s *RecallDocumentResponseBodyData) SetVectorSearchElapsedMs(v int64) *RecallDocumentResponseBodyData {
	s.VectorSearchElapsedMs = &v
	return s
}

type RecallDocumentResponseBodyDataChunkList struct {
	// example:
	//
	// 823746762354
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 839468263472
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// dscsbdsk
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 982374872364
	NextChunkId *string                                       `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 827364827364832
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataChunkList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetPos(v []*RecallDocumentResponseBodyDataChunkListPos) *RecallDocumentResponseBodyDataChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataChunkList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataChunkListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataChunkListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponseBodyDataChunkPartList struct {
	// example:
	//
	// 98327482364
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 92837482364
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// sjdhgjsd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 2387648263542
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataChunkPartListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 32874682764
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	Title *string  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataChunkPartList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkPartList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkText(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetChunkType(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetDocId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetFileType(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetLibraryId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetLibraryName(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetPos(v []*RecallDocumentResponseBodyDataChunkPartListPos) *RecallDocumentResponseBodyDataChunkPartList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetScore(v float32) *RecallDocumentResponseBodyDataChunkPartList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartList) SetTitle(v string) *RecallDocumentResponseBodyDataChunkPartList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataChunkPartListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataChunkPartListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataChunkPartListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetPage(v int32) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataChunkPartListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataChunkPartListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponseBodyDataDocuments struct {
	// example:
	//
	// 92837482364
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// {"a":"1"}
	DocumentMeta map[string]interface{} `json:"documentMeta,omitempty" xml:"documentMeta,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// sjdhgjsd
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/test.pdf
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RecallDocumentResponseBodyDataDocuments) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataDocuments) SetDocId(v string) *RecallDocumentResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetDocumentMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataDocuments {
	s.DocumentMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetFileType(v string) *RecallDocumentResponseBodyDataDocuments {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetGmtCreate(v string) *RecallDocumentResponseBodyDataDocuments {
	s.GmtCreate = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetLibraryId(v string) *RecallDocumentResponseBodyDataDocuments {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetTitle(v string) *RecallDocumentResponseBodyDataDocuments {
	s.Title = &v
	return s
}

func (s *RecallDocumentResponseBodyDataDocuments) SetUrl(v string) *RecallDocumentResponseBodyDataDocuments {
	s.Url = &v
	return s
}

type RecallDocumentResponseBodyDataTextChunkList struct {
	// example:
	//
	// 32874682364
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// http://oss-xxx-hangzhou.com/xxx
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 8372467263542
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// djsgfsjd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 23874682432
	NextChunkId *string                                           `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataTextChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 89473868346
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	Title *string  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataTextChunkList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataTextChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetPos(v []*RecallDocumentResponseBodyDataTextChunkListPos) *RecallDocumentResponseBodyDataTextChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataTextChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataTextChunkList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataTextChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataTextChunkListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataTextChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataTextChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataTextChunkListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponseBodyDataVectorChunkList struct {
	// example:
	//
	// 8723642345276
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// example:
	//
	// {"a":"1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// https://oss-xxxx-hangzhou.com/test.pdf
	ChunkOssUrl *string `json:"chunkOssUrl,omitempty" xml:"chunkOssUrl,omitempty"`
	ChunkText   *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// example:
	//
	// 78326476235675372
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// example:
	//
	// djsgfsjd
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
	// example:
	//
	// 293846872343
	NextChunkId *string                                             `json:"nextChunkId,omitempty" xml:"nextChunkId,omitempty"`
	Pos         []*RecallDocumentResponseBodyDataVectorChunkListPos `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
	// example:
	//
	// 873647326542
	PreChunkId *string `json:"preChunkId,omitempty" xml:"preChunkId,omitempty"`
	// example:
	//
	// 0.5
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s RecallDocumentResponseBodyDataVectorChunkList) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataVectorChunkList) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkMeta(v map[string]interface{}) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkMeta = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkOssUrl(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkOssUrl = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkText(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkText = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetChunkType(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetDocId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.DocId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetFileType(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.FileType = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetLibraryId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetLibraryName(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.LibraryName = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetNextChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.NextChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetPos(v []*RecallDocumentResponseBodyDataVectorChunkListPos) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Pos = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetPreChunkId(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.PreChunkId = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetScore(v float32) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Score = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkList) SetTitle(v string) *RecallDocumentResponseBodyDataVectorChunkList {
	s.Title = &v
	return s
}

type RecallDocumentResponseBodyDataVectorChunkListPos struct {
	AxisArray []*float64 `json:"axisArray,omitempty" xml:"axisArray,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page              *int32   `json:"page,omitempty" xml:"page,omitempty"`
	TextHighlightArea []*int32 `json:"textHighlightArea,omitempty" xml:"textHighlightArea,omitempty" type:"Repeated"`
}

func (s RecallDocumentResponseBodyDataVectorChunkListPos) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponseBodyDataVectorChunkListPos) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetAxisArray(v []*float64) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.AxisArray = v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetPage(v int32) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.Page = &v
	return s
}

func (s *RecallDocumentResponseBodyDataVectorChunkListPos) SetTextHighlightArea(v []*int32) *RecallDocumentResponseBodyDataVectorChunkListPos {
	s.TextHighlightArea = v
	return s
}

type RecallDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecallDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecallDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s RecallDocumentResponse) GoString() string {
	return s.String()
}

func (s *RecallDocumentResponse) SetHeaders(v map[string]*string) *RecallDocumentResponse {
	s.Headers = v
	return s
}

func (s *RecallDocumentResponse) SetStatusCode(v int32) *RecallDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *RecallDocumentResponse) SetBody(v *RecallDocumentResponseBody) *RecallDocumentResponse {
	s.Body = v
	return s
}

type RecognizeIntentionRequest struct {
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// common
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	Conversation              *string                                               `json:"conversation,omitempty" xml:"conversation,omitempty"`
	GlobalIntentionList       []*RecognizeIntentionRequestGlobalIntentionList       `json:"globalIntentionList,omitempty" xml:"globalIntentionList,omitempty" type:"Repeated"`
	HierarchicalIntentionList []*RecognizeIntentionRequestHierarchicalIntentionList `json:"hierarchicalIntentionList,omitempty" xml:"hierarchicalIntentionList,omitempty" type:"Repeated"`
	IntentionList             []*RecognizeIntentionRequestIntentionList             `json:"intentionList,omitempty" xml:"intentionList,omitempty" type:"Repeated"`
	// example:
	//
	// common
	OpType *string `json:"opType,omitempty" xml:"opType,omitempty"`
	// example:
	//
	// false
	Recommend *bool `json:"recommend,omitempty" xml:"recommend,omitempty"`
}

func (s RecognizeIntentionRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequest) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequest) SetAnalysis(v bool) *RecognizeIntentionRequest {
	s.Analysis = &v
	return s
}

func (s *RecognizeIntentionRequest) SetBizType(v string) *RecognizeIntentionRequest {
	s.BizType = &v
	return s
}

func (s *RecognizeIntentionRequest) SetConversation(v string) *RecognizeIntentionRequest {
	s.Conversation = &v
	return s
}

func (s *RecognizeIntentionRequest) SetGlobalIntentionList(v []*RecognizeIntentionRequestGlobalIntentionList) *RecognizeIntentionRequest {
	s.GlobalIntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetHierarchicalIntentionList(v []*RecognizeIntentionRequestHierarchicalIntentionList) *RecognizeIntentionRequest {
	s.HierarchicalIntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetIntentionList(v []*RecognizeIntentionRequestIntentionList) *RecognizeIntentionRequest {
	s.IntentionList = v
	return s
}

func (s *RecognizeIntentionRequest) SetOpType(v string) *RecognizeIntentionRequest {
	s.OpType = &v
	return s
}

func (s *RecognizeIntentionRequest) SetRecommend(v bool) *RecognizeIntentionRequest {
	s.Recommend = &v
	return s
}

type RecognizeIntentionRequestGlobalIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1810566978021232640
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
}

func (s RecognizeIntentionRequestGlobalIntentionList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequestGlobalIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetDescription(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntention(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestGlobalIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestGlobalIntentionList {
	s.IntentionCode = &v
	return s
}

type RecognizeIntentionRequestHierarchicalIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1810929291010150400
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
}

func (s RecognizeIntentionRequestHierarchicalIntentionList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequestHierarchicalIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetDescription(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntention(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestHierarchicalIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestHierarchicalIntentionList {
	s.IntentionCode = &v
	return s
}

type RecognizeIntentionRequestIntentionList struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Intention   *string `json:"intention,omitempty" xml:"intention,omitempty"`
	// example:
	//
	// 1808766224000262144
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
}

func (s RecognizeIntentionRequestIntentionList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionRequestIntentionList) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionRequestIntentionList) SetDescription(v string) *RecognizeIntentionRequestIntentionList {
	s.Description = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntention(v string) *RecognizeIntentionRequestIntentionList {
	s.Intention = &v
	return s
}

func (s *RecognizeIntentionRequestIntentionList) SetIntentionCode(v string) *RecognizeIntentionRequestIntentionList {
	s.IntentionCode = &v
	return s
}

type RecognizeIntentionResponseBody struct {
	// example:
	//
	// null
	Cost *int64                              `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RecognizeIntentionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 003D019A-1BB3-53EC-A0D2-CE76DA5D73B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RecognizeIntentionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponseBody) SetCost(v int64) *RecognizeIntentionResponseBody {
	s.Cost = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetData(v *RecognizeIntentionResponseBodyData) *RecognizeIntentionResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeIntentionResponseBody) SetDataType(v string) *RecognizeIntentionResponseBody {
	s.DataType = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetErrCode(v string) *RecognizeIntentionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetMessage(v string) *RecognizeIntentionResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetRequestId(v string) *RecognizeIntentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetSuccess(v bool) *RecognizeIntentionResponseBody {
	s.Success = &v
	return s
}

func (s *RecognizeIntentionResponseBody) SetTime(v string) *RecognizeIntentionResponseBody {
	s.Time = &v
	return s
}

type RecognizeIntentionResponseBodyData struct {
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// example:
	//
	// 1
	IntentionCode      *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	IntentionName      *string `json:"intentionName,omitempty" xml:"intentionName,omitempty"`
	RecommendIntention *string `json:"recommendIntention,omitempty" xml:"recommendIntention,omitempty"`
	RecommendScript    *string `json:"recommendScript,omitempty" xml:"recommendScript,omitempty"`
}

func (s RecognizeIntentionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponseBodyData) SetAnalysisProcess(v string) *RecognizeIntentionResponseBodyData {
	s.AnalysisProcess = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionCode(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionCode = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetIntentionName(v string) *RecognizeIntentionResponseBodyData {
	s.IntentionName = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetRecommendIntention(v string) *RecognizeIntentionResponseBodyData {
	s.RecommendIntention = &v
	return s
}

func (s *RecognizeIntentionResponseBodyData) SetRecommendScript(v string) *RecognizeIntentionResponseBodyData {
	s.RecommendScript = &v
	return s
}

type RecognizeIntentionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecognizeIntentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecognizeIntentionResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeIntentionResponse) GoString() string {
	return s.String()
}

func (s *RecognizeIntentionResponse) SetHeaders(v map[string]*string) *RecognizeIntentionResponse {
	s.Headers = v
	return s
}

func (s *RecognizeIntentionResponse) SetStatusCode(v int32) *RecognizeIntentionResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeIntentionResponse) SetBody(v *RecognizeIntentionResponseBody) *RecognizeIntentionResponse {
	s.Body = v
	return s
}

type RunChatResultGenerationRequest struct {
	// example:
	//
	// {"topP": 0.8}
	InferenceParameters map[string]interface{} `json:"inferenceParameters,omitempty" xml:"inferenceParameters,omitempty"`
	// This parameter is required.
	Messages []*RunChatResultGenerationRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream *bool                                  `json:"stream,omitempty" xml:"stream,omitempty"`
	Tools  []*RunChatResultGenerationRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequest) SetInferenceParameters(v map[string]interface{}) *RunChatResultGenerationRequest {
	s.InferenceParameters = v
	return s
}

func (s *RunChatResultGenerationRequest) SetMessages(v []*RunChatResultGenerationRequestMessages) *RunChatResultGenerationRequest {
	s.Messages = v
	return s
}

func (s *RunChatResultGenerationRequest) SetModelId(v string) *RunChatResultGenerationRequest {
	s.ModelId = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetSessionId(v string) *RunChatResultGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetStream(v bool) *RunChatResultGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetTools(v []*RunChatResultGenerationRequestTools) *RunChatResultGenerationRequest {
	s.Tools = v
	return s
}

type RunChatResultGenerationRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunChatResultGenerationRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestMessages) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestMessages) SetContent(v string) *RunChatResultGenerationRequestMessages {
	s.Content = &v
	return s
}

func (s *RunChatResultGenerationRequestMessages) SetRole(v string) *RunChatResultGenerationRequestMessages {
	s.Role = &v
	return s
}

type RunChatResultGenerationRequestTools struct {
	Function *RunChatResultGenerationRequestToolsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunChatResultGenerationRequestTools) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestTools) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestTools) SetFunction(v *RunChatResultGenerationRequestToolsFunction) *RunChatResultGenerationRequestTools {
	s.Function = v
	return s
}

func (s *RunChatResultGenerationRequestTools) SetType(v string) *RunChatResultGenerationRequestTools {
	s.Type = &v
	return s
}

type RunChatResultGenerationRequestToolsFunction struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// get_time
	Name       *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	Parameters *RunChatResultGenerationRequestToolsFunctionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Required   []*string                                              `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationRequestToolsFunction) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestToolsFunction) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestToolsFunction) SetDescription(v string) *RunChatResultGenerationRequestToolsFunction {
	s.Description = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetName(v string) *RunChatResultGenerationRequestToolsFunction {
	s.Name = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetParameters(v *RunChatResultGenerationRequestToolsFunctionParameters) *RunChatResultGenerationRequestToolsFunction {
	s.Parameters = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetRequired(v []*string) *RunChatResultGenerationRequestToolsFunction {
	s.Required = v
	return s
}

type RunChatResultGenerationRequestToolsFunctionParameters struct {
	// example:
	//
	// {
	//
	//                             "location": {
	//
	//                                 "type": "string",
	//
	//                                 "description": "The city and state, e.g. San Francisco, CA"
	//
	//                             },
	//
	//                             "unit": {
	//
	//                                 "type": "string",
	//
	//                                 "enum": [
	//
	//                                     "celsius",
	//
	//                                     "fahrenheit"
	//
	//                                 ]
	//
	//                             }
	//
	//                         }
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunChatResultGenerationRequestToolsFunctionParameters) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationRequestToolsFunctionParameters) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) SetProperties(v map[string]interface{}) *RunChatResultGenerationRequestToolsFunctionParameters {
	s.Properties = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) SetType(v string) *RunChatResultGenerationRequestToolsFunctionParameters {
	s.Type = &v
	return s
}

type RunChatResultGenerationResponseBody struct {
	Choices []*RunChatResultGenerationResponseBodyChoices `json:"choices,omitempty" xml:"choices,omitempty" type:"Repeated"`
	// example:
	//
	// 1720602203
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 500
	TotalTokens *int32                                    `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
	Usage       *RunChatResultGenerationResponseBodyUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunChatResultGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBody) SetChoices(v []*RunChatResultGenerationResponseBodyChoices) *RunChatResultGenerationResponseBody {
	s.Choices = v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetCreated(v int64) *RunChatResultGenerationResponseBody {
	s.Created = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetId(v string) *RunChatResultGenerationResponseBody {
	s.Id = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetModelId(v string) *RunChatResultGenerationResponseBody {
	s.ModelId = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetRequestId(v string) *RunChatResultGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetTime(v string) *RunChatResultGenerationResponseBody {
	s.Time = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetTotalTokens(v int32) *RunChatResultGenerationResponseBody {
	s.TotalTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBody) SetUsage(v *RunChatResultGenerationResponseBodyUsage) *RunChatResultGenerationResponseBody {
	s.Usage = v
	return s
}

type RunChatResultGenerationResponseBodyChoices struct {
	// example:
	//
	// null
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 0
	Index   *int32                                             `json:"index,omitempty" xml:"index,omitempty"`
	Message *RunChatResultGenerationResponseBodyChoicesMessage `json:"message,omitempty" xml:"message,omitempty" type:"Struct"`
}

func (s RunChatResultGenerationResponseBodyChoices) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyChoices) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyChoices) SetFinishReason(v string) *RunChatResultGenerationResponseBodyChoices {
	s.FinishReason = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) SetIndex(v int32) *RunChatResultGenerationResponseBodyChoices {
	s.Index = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoices) SetMessage(v *RunChatResultGenerationResponseBodyChoicesMessage) *RunChatResultGenerationResponseBodyChoices {
	s.Message = v
	return s
}

type RunChatResultGenerationResponseBodyChoicesMessage struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role      *string                  `json:"role,omitempty" xml:"role,omitempty"`
	ToolCalls []map[string]interface{} `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationResponseBodyChoicesMessage) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyChoicesMessage) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetContent(v string) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.Content = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetRole(v string) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.Role = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyChoicesMessage) SetToolCalls(v []map[string]interface{}) *RunChatResultGenerationResponseBodyChoicesMessage {
	s.ToolCalls = v
	return s
}

type RunChatResultGenerationResponseBodyUsage struct {
	// example:
	//
	// 0
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// 0
	ImageTokens *int32 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 200
	InputTokens *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 300
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 500
	TotalTokens *int32 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunChatResultGenerationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponseBodyUsage) SetImageCount(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.ImageCount = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetImageTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.ImageTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetInputTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetOutputTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunChatResultGenerationResponseBodyUsage) SetTotalTokens(v int32) *RunChatResultGenerationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

type RunChatResultGenerationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunChatResultGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunChatResultGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunChatResultGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationResponse) SetHeaders(v map[string]*string) *RunChatResultGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunChatResultGenerationResponse) SetStatusCode(v int32) *RunChatResultGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunChatResultGenerationResponse) SetBody(v *RunChatResultGenerationResponseBody) *RunChatResultGenerationResponse {
	s.Body = v
	return s
}

type RunLibraryChatGenerationRequest struct {
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// false
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// false
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// qwen-max
	FollowUpLlm *string `json:"followUpLlm,omitempty" xml:"followUpLlm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	LlmType *string `json:"llmType,omitempty" xml:"llmType,omitempty"`
	// example:
	//
	// qwen-max
	MultiQueryLlm *string `json:"multiQueryLlm,omitempty" xml:"multiQueryLlm,omitempty"`
	// This parameter is required.
	Query         *string                                       `json:"query,omitempty" xml:"query,omitempty"`
	QueryCriteria *RunLibraryChatGenerationRequestQueryCriteria `json:"queryCriteria,omitempty" xml:"queryCriteria,omitempty" type:"Struct"`
	// example:
	//
	// linear
	RerankType *string `json:"rerankType,omitempty" xml:"rerankType,omitempty"`
	// sessionId
	//
	// example:
	//
	// null
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream              *bool                                               `json:"stream,omitempty" xml:"stream,omitempty"`
	SubQueryList        []*string                                           `json:"subQueryList,omitempty" xml:"subQueryList,omitempty" type:"Repeated"`
	TextSearchParameter *RunLibraryChatGenerationRequestTextSearchParameter `json:"textSearchParameter,omitempty" xml:"textSearchParameter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TopK                  *int32                                                `json:"topK,omitempty" xml:"topK,omitempty"`
	VectorSearchParameter *RunLibraryChatGenerationRequestVectorSearchParameter `json:"vectorSearchParameter,omitempty" xml:"vectorSearchParameter,omitempty" type:"Struct"`
	// example:
	//
	// false
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s RunLibraryChatGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequest) SetDocIdList(v []*string) *RunLibraryChatGenerationRequest {
	s.DocIdList = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableFollowUp(v bool) *RunLibraryChatGenerationRequest {
	s.EnableFollowUp = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableMultiQuery(v bool) *RunLibraryChatGenerationRequest {
	s.EnableMultiQuery = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableOpenQa(v bool) *RunLibraryChatGenerationRequest {
	s.EnableOpenQa = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetFollowUpLlm(v string) *RunLibraryChatGenerationRequest {
	s.FollowUpLlm = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetLibraryId(v string) *RunLibraryChatGenerationRequest {
	s.LibraryId = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetLlmType(v string) *RunLibraryChatGenerationRequest {
	s.LlmType = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetMultiQueryLlm(v string) *RunLibraryChatGenerationRequest {
	s.MultiQueryLlm = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetQuery(v string) *RunLibraryChatGenerationRequest {
	s.Query = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetQueryCriteria(v *RunLibraryChatGenerationRequestQueryCriteria) *RunLibraryChatGenerationRequest {
	s.QueryCriteria = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetRerankType(v string) *RunLibraryChatGenerationRequest {
	s.RerankType = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetSessionId(v string) *RunLibraryChatGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetStream(v bool) *RunLibraryChatGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetSubQueryList(v []*string) *RunLibraryChatGenerationRequest {
	s.SubQueryList = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetTextSearchParameter(v *RunLibraryChatGenerationRequestTextSearchParameter) *RunLibraryChatGenerationRequest {
	s.TextSearchParameter = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetTopK(v int32) *RunLibraryChatGenerationRequest {
	s.TopK = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetVectorSearchParameter(v *RunLibraryChatGenerationRequestVectorSearchParameter) *RunLibraryChatGenerationRequest {
	s.VectorSearchParameter = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetWithDocumentReference(v bool) *RunLibraryChatGenerationRequest {
	s.WithDocumentReference = &v
	return s
}

type RunLibraryChatGenerationRequestQueryCriteria struct {
	And []*RunLibraryChatGenerationRequestQueryCriteriaAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	Or  []*RunLibraryChatGenerationRequestQueryCriteriaOr  `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
}

func (s RunLibraryChatGenerationRequestQueryCriteria) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteria) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) SetAnd(v []*RunLibraryChatGenerationRequestQueryCriteriaAnd) *RunLibraryChatGenerationRequestQueryCriteria {
	s.And = v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) SetOr(v []*RunLibraryChatGenerationRequestQueryCriteriaOr) *RunLibraryChatGenerationRequestQueryCriteria {
	s.Or = v
	return s
}

type RunLibraryChatGenerationRequestQueryCriteriaAnd struct {
	// example:
	//
	// 0.5
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// city
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunLibraryChatGenerationRequestQueryCriteriaAnd) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteriaAnd) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetBoost(v float32) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Boost = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetKey(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Key = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetOperator(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Operator = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetValue(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Value = &v
	return s
}

type RunLibraryChatGenerationRequestQueryCriteriaOr struct {
	// example:
	//
	// 0.5
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// city
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunLibraryChatGenerationRequestQueryCriteriaOr) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteriaOr) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetBoost(v float32) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Boost = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetKey(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Key = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetOperator(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Operator = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetValue(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Value = &v
	return s
}

type RunLibraryChatGenerationRequestTextSearchParameter struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// IkMaxWord
	SearchAnalyzerType *string `json:"searchAnalyzerType,omitempty" xml:"searchAnalyzerType,omitempty"`
}

func (s RunLibraryChatGenerationRequestTextSearchParameter) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestTextSearchParameter) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) SetLimit(v int32) *RunLibraryChatGenerationRequestTextSearchParameter {
	s.Limit = &v
	return s
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) SetSearchAnalyzerType(v string) *RunLibraryChatGenerationRequestTextSearchParameter {
	s.SearchAnalyzerType = &v
	return s
}

type RunLibraryChatGenerationRequestVectorSearchParameter struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s RunLibraryChatGenerationRequestVectorSearchParameter) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationRequestVectorSearchParameter) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestVectorSearchParameter) SetLimit(v int32) *RunLibraryChatGenerationRequestVectorSearchParameter {
	s.Limit = &v
	return s
}

type RunLibraryChatGenerationResponseBody struct {
	// example:
	//
	// null
	Cost *int64      `json:"cost,omitempty" xml:"cost,omitempty"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RunLibraryChatGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponseBody) SetCost(v int64) *RunLibraryChatGenerationResponseBody {
	s.Cost = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetData(v interface{}) *RunLibraryChatGenerationResponseBody {
	s.Data = v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetDataType(v string) *RunLibraryChatGenerationResponseBody {
	s.DataType = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetErrCode(v string) *RunLibraryChatGenerationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetMessage(v string) *RunLibraryChatGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetRequestId(v string) *RunLibraryChatGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetSuccess(v bool) *RunLibraryChatGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLibraryChatGenerationResponseBody) SetTime(v string) *RunLibraryChatGenerationResponseBody {
	s.Time = &v
	return s
}

type RunLibraryChatGenerationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLibraryChatGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLibraryChatGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunLibraryChatGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponse) SetHeaders(v map[string]*string) *RunLibraryChatGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunLibraryChatGenerationResponse) SetStatusCode(v int32) *RunLibraryChatGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponse) SetBody(v *RunLibraryChatGenerationResponseBody) *RunLibraryChatGenerationResponse {
	s.Body = v
	return s
}

type UpdateDocumentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc123
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// {
	//
	//         "businessId": "12321"
	//
	//     }
	Meta map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentRequest) SetDocId(v string) *UpdateDocumentRequest {
	s.DocId = &v
	return s
}

func (s *UpdateDocumentRequest) SetLibraryId(v string) *UpdateDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *UpdateDocumentRequest) SetMeta(v map[string]interface{}) *UpdateDocumentRequest {
	s.Meta = v
	return s
}

func (s *UpdateDocumentRequest) SetTitle(v string) *UpdateDocumentRequest {
	s.Title = &v
	return s
}

type UpdateDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// null
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UpdateDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponseBody) SetCost(v int64) *UpdateDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetData(v string) *UpdateDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetDataType(v string) *UpdateDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetErrCode(v string) *UpdateDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetMessage(v string) *UpdateDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetRequestId(v string) *UpdateDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetSuccess(v bool) *UpdateDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDocumentResponseBody) SetTime(v string) *UpdateDocumentResponseBody {
	s.Time = &v
	return s
}

type UpdateDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocumentResponse) SetHeaders(v map[string]*string) *UpdateDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocumentResponse) SetStatusCode(v int32) *UpdateDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocumentResponse) SetBody(v *UpdateDocumentResponseBody) *UpdateDocumentResponse {
	s.Body = v
	return s
}

type UpdateLibraryRequest struct {
	Description  *string                           `json:"description,omitempty" xml:"description,omitempty"`
	IndexSetting *UpdateLibraryRequestIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// dsfbashdbb
	LibraryId   *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s UpdateLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequest) SetDescription(v string) *UpdateLibraryRequest {
	s.Description = &v
	return s
}

func (s *UpdateLibraryRequest) SetIndexSetting(v *UpdateLibraryRequestIndexSetting) *UpdateLibraryRequest {
	s.IndexSetting = v
	return s
}

func (s *UpdateLibraryRequest) SetLibraryId(v string) *UpdateLibraryRequest {
	s.LibraryId = &v
	return s
}

func (s *UpdateLibraryRequest) SetLibraryName(v string) *UpdateLibraryRequest {
	s.LibraryName = &v
	return s
}

type UpdateLibraryRequestIndexSetting struct {
	ChunkStrategy      *UpdateLibraryRequestIndexSettingChunkStrategy      `json:"chunkStrategy,omitempty" xml:"chunkStrategy,omitempty" type:"Struct"`
	ModelConfig        *UpdateLibraryRequestIndexSettingModelConfig        `json:"modelConfig,omitempty" xml:"modelConfig,omitempty" type:"Struct"`
	PromptRoleStyle    *string                                             `json:"promptRoleStyle,omitempty" xml:"promptRoleStyle,omitempty"`
	QueryEnhancer      *UpdateLibraryRequestIndexSettingQueryEnhancer      `json:"queryEnhancer,omitempty" xml:"queryEnhancer,omitempty" type:"Struct"`
	RecallStrategy     *UpdateLibraryRequestIndexSettingRecallStrategy     `json:"recallStrategy,omitempty" xml:"recallStrategy,omitempty" type:"Struct"`
	TextIndexSetting   *UpdateLibraryRequestIndexSettingTextIndexSetting   `json:"textIndexSetting,omitempty" xml:"textIndexSetting,omitempty" type:"Struct"`
	VectorIndexSetting *UpdateLibraryRequestIndexSettingVectorIndexSetting `json:"vectorIndexSetting,omitempty" xml:"vectorIndexSetting,omitempty" type:"Struct"`
}

func (s UpdateLibraryRequestIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSetting) SetChunkStrategy(v *UpdateLibraryRequestIndexSettingChunkStrategy) *UpdateLibraryRequestIndexSetting {
	s.ChunkStrategy = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetModelConfig(v *UpdateLibraryRequestIndexSettingModelConfig) *UpdateLibraryRequestIndexSetting {
	s.ModelConfig = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetPromptRoleStyle(v string) *UpdateLibraryRequestIndexSetting {
	s.PromptRoleStyle = &v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetQueryEnhancer(v *UpdateLibraryRequestIndexSettingQueryEnhancer) *UpdateLibraryRequestIndexSetting {
	s.QueryEnhancer = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetRecallStrategy(v *UpdateLibraryRequestIndexSettingRecallStrategy) *UpdateLibraryRequestIndexSetting {
	s.RecallStrategy = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetTextIndexSetting(v *UpdateLibraryRequestIndexSettingTextIndexSetting) *UpdateLibraryRequestIndexSetting {
	s.TextIndexSetting = v
	return s
}

func (s *UpdateLibraryRequestIndexSetting) SetVectorIndexSetting(v *UpdateLibraryRequestIndexSettingVectorIndexSetting) *UpdateLibraryRequestIndexSetting {
	s.VectorIndexSetting = v
	return s
}

type UpdateLibraryRequestIndexSettingChunkStrategy struct {
	// example:
	//
	// true
	DocTreeSplit *bool `json:"docTreeSplit,omitempty" xml:"docTreeSplit,omitempty"`
	// example:
	//
	// 160
	DocTreeSplitSize *int32 `json:"docTreeSplitSize,omitempty" xml:"docTreeSplitSize,omitempty"`
	// example:
	//
	// true
	EnhanceGraph *bool `json:"enhanceGraph,omitempty" xml:"enhanceGraph,omitempty"`
	// example:
	//
	// true
	EnhanceTable *bool `json:"enhanceTable,omitempty" xml:"enhanceTable,omitempty"`
	// example:
	//
	// 20
	Overlap *int32 `json:"overlap,omitempty" xml:"overlap,omitempty"`
	// example:
	//
	// true
	SentenceSplit *bool `json:"sentenceSplit,omitempty" xml:"sentenceSplit,omitempty"`
	// example:
	//
	// 160
	SentenceSplitSize *int32 `json:"sentenceSplitSize,omitempty" xml:"sentenceSplitSize,omitempty"`
	// example:
	//
	// 256
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// true
	Split *bool `json:"split,omitempty" xml:"split,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingChunkStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplit(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplit = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetDocTreeSplitSize(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.DocTreeSplitSize = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetEnhanceGraph(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceGraph = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetEnhanceTable(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.EnhanceTable = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetOverlap(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.Overlap = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplit(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplit = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSentenceSplitSize(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.SentenceSplitSize = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSize(v int32) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.Size = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) SetSplit(v bool) *UpdateLibraryRequestIndexSettingChunkStrategy {
	s.Split = &v
	return s
}

type UpdateLibraryRequestIndexSettingModelConfig struct {
	// example:
	//
	// 0.8
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// topP
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"topP,omitempty" xml:"topP,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingModelConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) SetTemperature(v float64) *UpdateLibraryRequestIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) SetTopP(v float64) *UpdateLibraryRequestIndexSettingModelConfig {
	s.TopP = &v
	return s
}

type UpdateLibraryRequestIndexSettingQueryEnhancer struct {
	// example:
	//
	// true
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// true
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// true
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// true
	EnableQueryRewrite *bool `json:"enableQueryRewrite,omitempty" xml:"enableQueryRewrite,omitempty"`
	// example:
	//
	// true
	EnableSession *bool `json:"enableSession,omitempty" xml:"enableSession,omitempty"`
	// example:
	//
	// sjdhgfc
	LocalKnowledgeId *string `json:"localKnowledgeId,omitempty" xml:"localKnowledgeId,omitempty"`
	// example:
	//
	// true
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingQueryEnhancer) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableFollowUp(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableFollowUp = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableMultiQuery(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableMultiQuery = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableOpenQa(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableOpenQa = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableQueryRewrite(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableQueryRewrite = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetEnableSession(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.EnableSession = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetLocalKnowledgeId(v string) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.LocalKnowledgeId = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) SetWithDocumentReference(v bool) *UpdateLibraryRequestIndexSettingQueryEnhancer {
	s.WithDocumentReference = &v
	return s
}

type UpdateLibraryRequestIndexSettingRecallStrategy struct {
	// example:
	//
	// model
	DocumentRankType *string `json:"documentRankType,omitempty" xml:"documentRankType,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingRecallStrategy) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) SetDocumentRankType(v string) *UpdateLibraryRequestIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) SetLimit(v int32) *UpdateLibraryRequestIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

type UpdateLibraryRequestIndexSettingTextIndexSetting struct {
	// example:
	//
	// ElasticSearch
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// Standard
	IndexAnalyzer *string `json:"indexAnalyzer,omitempty" xml:"indexAnalyzer,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// Standard
	SearchAnalyzer *string `json:"searchAnalyzer,omitempty" xml:"searchAnalyzer,omitempty"`
	// example:
	//
	// 50
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingTextIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetCategory(v string) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.Category = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetEnable(v bool) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.Enable = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetIndexAnalyzer(v string) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.IndexAnalyzer = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetRankThreshold(v float64) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetSearchAnalyzer(v string) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.SearchAnalyzer = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) SetTopK(v int32) *UpdateLibraryRequestIndexSettingTextIndexSetting {
	s.TopK = &v
	return s
}

type UpdateLibraryRequestIndexSettingVectorIndexSetting struct {
	// example:
	//
	// ADB
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// DashScope
	EmbeddingType *string `json:"embeddingType,omitempty" xml:"embeddingType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 0.5
	RankThreshold *float64 `json:"rankThreshold,omitempty" xml:"rankThreshold,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateLibraryRequestIndexSettingVectorIndexSetting) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetCategory(v string) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.Category = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetEmbeddingType(v string) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.EmbeddingType = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetEnable(v bool) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.Enable = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetRankThreshold(v float64) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.RankThreshold = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) SetTopK(v int32) *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	s.TopK = &v
	return s
}

type UpdateLibraryResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// null
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// eb2b6139-ddf1-91a0-a47f-df7617ae9032
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UpdateLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLibraryResponseBody) SetCost(v int64) *UpdateLibraryResponseBody {
	s.Cost = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetData(v string) *UpdateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetDataType(v string) *UpdateLibraryResponseBody {
	s.DataType = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetErrCode(v string) *UpdateLibraryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetMessage(v string) *UpdateLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetRequestId(v string) *UpdateLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetSuccess(v bool) *UpdateLibraryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLibraryResponseBody) SetTime(v string) *UpdateLibraryResponseBody {
	s.Time = &v
	return s
}

type UpdateLibraryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryResponse) GoString() string {
	return s.String()
}

func (s *UpdateLibraryResponse) SetHeaders(v map[string]*string) *UpdateLibraryResponse {
	s.Headers = v
	return s
}

func (s *UpdateLibraryResponse) SetStatusCode(v int32) *UpdateLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLibraryResponse) SetBody(v *UpdateLibraryResponseBody) *UpdateLibraryResponse {
	s.Body = v
	return s
}

type UploadDocumentRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-xxx.hangzhou.com/test.pdf
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdhbcsj
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UploadDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentRequest) SetData(v string) *UploadDocumentRequest {
	s.Data = &v
	return s
}

func (s *UploadDocumentRequest) SetFileName(v string) *UploadDocumentRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentRequest) SetFileUrl(v string) *UploadDocumentRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadDocumentRequest) SetLibraryId(v string) *UploadDocumentRequest {
	s.LibraryId = &v
	return s
}

type UploadDocumentAdvanceRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-xxx.hangzhou.com/test.pdf
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdhbcsj
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UploadDocumentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentAdvanceRequest) SetData(v string) *UploadDocumentAdvanceRequest {
	s.Data = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetFileName(v string) *UploadDocumentAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetFileUrlObject(v io.Reader) *UploadDocumentAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetLibraryId(v string) *UploadDocumentAdvanceRequest {
	s.LibraryId = &v
	return s
}

type UploadDocumentResponseBody struct {
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// 1782981430906818562
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// ff3fef67-48d9-4379-a237-9ba8143fe739
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s UploadDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponseBody) SetCost(v int64) *UploadDocumentResponseBody {
	s.Cost = &v
	return s
}

func (s *UploadDocumentResponseBody) SetData(v string) *UploadDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UploadDocumentResponseBody) SetDataType(v string) *UploadDocumentResponseBody {
	s.DataType = &v
	return s
}

func (s *UploadDocumentResponseBody) SetErrCode(v string) *UploadDocumentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UploadDocumentResponseBody) SetMessage(v string) *UploadDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UploadDocumentResponseBody) SetRequestId(v string) *UploadDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocumentResponseBody) SetSuccess(v bool) *UploadDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *UploadDocumentResponseBody) SetTime(v string) *UploadDocumentResponseBody {
	s.Time = &v
	return s
}

type UploadDocumentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentResponse) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponse) SetHeaders(v map[string]*string) *UploadDocumentResponse {
	s.Headers = v
	return s
}

func (s *UploadDocumentResponse) SetStatusCode(v int32) *UploadDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadDocumentResponse) SetBody(v *UploadDocumentResponseBody) *UploadDocumentResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dianjin"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTaskWithOptions(workspaceId *string, request *CreateFinReportSummaryTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTable)) {
		body["enableTable"] = request.EnableTable
	}

	if !tea.BoolValue(util.IsUnset(request.EndPage)) {
		body["endPage"] = request.EndPage
	}

	if !tea.BoolValue(util.IsUnset(request.Instruction)) {
		body["instruction"] = request.Instruction
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.StartPage)) {
		body["startPage"] = request.StartPage
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFinReportSummaryTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/summary"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFinReportSummaryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTask(workspaceId *string, request *CreateFinReportSummaryTaskRequest) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFinReportSummaryTaskResponse{}
	_body, _err := client.CreateFinReportSummaryTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryResponse
func (client *Client) CreateLibraryWithOptions(workspaceId *string, request *CreateLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IndexSetting)) {
		body["indexSetting"] = request.IndexSetting
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryName)) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @return CreateLibraryResponse
func (client *Client) CreateLibrary(workspaceId *string, request *CreateLibraryRequest) (_result *CreateLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLibraryResponse{}
	_body, _err := client.CreateLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocumentWithOptions(workspaceId *string, request *CreatePredefinedDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePredefinedDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Chunks)) {
		body["chunks"] = request.Chunks
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Metadata)) {
		body["metadata"] = request.Metadata
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePredefinedDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/createPredefinedDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePredefinedDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocument(workspaceId *string, request *CreatePredefinedDocumentRequest) (_result *CreatePredefinedDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePredefinedDocumentResponse{}
	_body, _err := client.CreatePredefinedDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithOptions(workspaceId *string, request *DeleteDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocIds)) {
		body["docIds"] = request.DocIds
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocument(workspaceId *string, request *DeleteDocumentRequest) (_result *DeleteDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDocumentResponse{}
	_body, _err := client.DeleteDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibraryWithOptions(workspaceId *string, request *DeleteLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibrary(workspaceId *string, request *DeleteLibraryRequest) (_result *DeleteLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLibraryResponse{}
	_body, _err := client.DeleteLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvictTaskResponse
func (client *Client) EvictTaskWithOptions(workspaceId *string, request *EvictTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EvictTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EvictTask"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/evict"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EvictTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @return EvictTaskResponse
func (client *Client) EvictTask(workspaceId *string, request *EvictTaskRequest) (_result *EvictTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EvictTaskResponse{}
	_body, _err := client.EvictTaskWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取app配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfigWithOptions(workspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppConfig"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/app/config"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取app配置
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfig(workspaceId *string) (_result *GetAppConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppConfigResponse{}
	_body, _err := client.GetAppConfigWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkListWithOptions(workspaceId *string, request *GetDocumentChunkListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentChunkListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChunkIdList)) {
		body["chunkIdList"] = request.ChunkIdList
	}

	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["orderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["searchQuery"] = request.SearchQuery
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentChunkList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/getDocumentChunk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentChunkListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkList(workspaceId *string, request *GetDocumentChunkListRequest) (_result *GetDocumentChunkListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentChunkListResponse{}
	_body, _err := client.GetDocumentChunkListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentListWithOptions(workspaceId *string, request *GetDocumentListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		query["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/listDocument"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentList(workspaceId *string, request *GetDocumentListRequest) (_result *GetDocumentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentListResponse{}
	_body, _err := client.GetDocumentListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrlWithOptions(workspaceId *string, request *GetDocumentUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDocumentUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentId)) {
		query["documentId"] = request.DocumentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDocumentUrl"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/url"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDocumentUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrl(workspaceId *string, request *GetDocumentUrlRequest) (_result *GetDocumentUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentUrlResponse{}
	_body, _err := client.GetDocumentUrlWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentListWithOptions(workspaceId *string, request *GetFilterDocumentListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFilterDocumentListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.And)) {
		body["and"] = request.And
	}

	if !tea.BoolValue(util.IsUnset(request.DocIdList)) {
		body["docIdList"] = request.DocIdList
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Or)) {
		body["or"] = request.Or
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFilterDocumentList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/filterDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFilterDocumentListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentList(workspaceId *string, request *GetFilterDocumentListRequest) (_result *GetFilterDocumentListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFilterDocumentListResponse{}
	_body, _err := client.GetFilterDocumentListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizTypeWithOptions(workspaceId *string, request *GetHistoryListByBizTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHistoryListByBizTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["bizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoryListByBizType"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/history/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoryListByBizTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizType(workspaceId *string, request *GetHistoryListByBizTypeRequest) (_result *GetHistoryListByBizTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHistoryListByBizTypeResponse{}
	_body, _err := client.GetHistoryListByBizTypeWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryResponse
func (client *Client) GetLibraryWithOptions(workspaceId *string, request *GetLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/get"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @return GetLibraryResponse
func (client *Client) GetLibrary(workspaceId *string, request *GetLibraryRequest) (_result *GetLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLibraryResponse{}
	_body, _err := client.GetLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryListWithOptions(workspaceId *string, request *GetLibraryListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLibraryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLibraryList"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLibraryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryList(workspaceId *string, request *GetLibraryListRequest) (_result *GetLibraryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLibraryListResponse{}
	_body, _err := client.GetLibraryListWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseResultResponse
func (client *Client) GetParseResultWithOptions(workspaceId *string, request *GetParseResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetParseResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetParseResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/getParseResult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetParseResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @return GetParseResultResponse
func (client *Client) GetParseResult(workspaceId *string, request *GetParseResultRequest) (_result *GetParseResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetParseResultResponse{}
	_body, _err := client.GetParseResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResultWithOptions(workspaceId *string, request *GetSummaryTaskResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSummaryTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSummaryTaskResult"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/summary/result"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSummaryTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResult(workspaceId *string, request *GetSummaryTaskResultRequest) (_result *GetSummaryTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSummaryTaskResultResponse{}
	_body, _err := client.GetSummaryTaskResultWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithOptions(workspaceId *string, request *GetTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/task/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatus(workspaceId *string, request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokePluginResponse
func (client *Client) InvokePluginWithOptions(workspaceId *string, request *InvokePluginRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvokePluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		body["pluginId"] = request.PluginId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokePlugin"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/plugin/invoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @return InvokePluginResponse
func (client *Client) InvokePlugin(workspaceId *string, request *InvokePluginRequest) (_result *InvokePluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokePluginResponse{}
	_body, _err := client.InvokePluginWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocumentWithOptions(workspaceId *string, request *PreviewDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PreviewDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentId)) {
		query["documentId"] = request.DocumentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviewDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/preview"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PreviewDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocument(workspaceId *string, request *PreviewDocumentRequest) (_result *PreviewDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewDocumentResponse{}
	_body, _err := client.PreviewDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReIndexResponse
func (client *Client) ReIndexWithOptions(workspaceId *string, request *ReIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentId)) {
		query["documentId"] = request.DocumentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReIndex"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/reIndex"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @return ReIndexResponse
func (client *Client) ReIndex(workspaceId *string, request *ReIndexRequest) (_result *ReIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReIndexResponse{}
	_body, _err := client.ReIndexWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档召回
//
// @param request - RecallDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallDocumentResponse
func (client *Client) RecallDocumentWithOptions(workspaceId *string, request *RecallDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecallDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		body["filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Rearrangement)) {
		body["rearrangement"] = request.Rearrangement
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["topK"] = request.TopK
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecallDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/recallDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecallDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档召回
//
// @param request - RecallDocumentRequest
//
// @return RecallDocumentResponse
func (client *Client) RecallDocument(workspaceId *string, request *RecallDocumentRequest) (_result *RecallDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecallDocumentResponse{}
	_body, _err := client.RecallDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntentionWithOptions(workspaceId *string, request *RecognizeIntentionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecognizeIntentionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Analysis)) {
		body["analysis"] = request.Analysis
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["bizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Conversation)) {
		body["conversation"] = request.Conversation
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalIntentionList)) {
		body["globalIntentionList"] = request.GlobalIntentionList
	}

	if !tea.BoolValue(util.IsUnset(request.HierarchicalIntentionList)) {
		body["hierarchicalIntentionList"] = request.HierarchicalIntentionList
	}

	if !tea.BoolValue(util.IsUnset(request.IntentionList)) {
		body["intentionList"] = request.IntentionList
	}

	if !tea.BoolValue(util.IsUnset(request.OpType)) {
		body["opType"] = request.OpType
	}

	if !tea.BoolValue(util.IsUnset(request.Recommend)) {
		body["recommend"] = request.Recommend
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeIntention"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/recog/intent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeIntentionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntention(workspaceId *string, request *RecognizeIntentionRequest) (_result *RecognizeIntentionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecognizeIntentionResponse{}
	_body, _err := client.RecognizeIntentionWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGenerationWithOptions(workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunChatResultGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InferenceParameters)) {
		body["inferenceParameters"] = request.InferenceParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		body["messages"] = request.Messages
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.Tools)) {
		body["tools"] = request.Tools
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunChatResultGeneration"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/run/chat/generation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunChatResultGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGeneration(workspaceId *string, request *RunChatResultGenerationRequest) (_result *RunChatResultGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunChatResultGenerationResponse{}
	_body, _err := client.RunChatResultGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGenerationWithOptions(workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunLibraryChatGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocIdList)) {
		body["docIdList"] = request.DocIdList
	}

	if !tea.BoolValue(util.IsUnset(request.EnableFollowUp)) {
		body["enableFollowUp"] = request.EnableFollowUp
	}

	if !tea.BoolValue(util.IsUnset(request.EnableMultiQuery)) {
		body["enableMultiQuery"] = request.EnableMultiQuery
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOpenQa)) {
		body["enableOpenQa"] = request.EnableOpenQa
	}

	if !tea.BoolValue(util.IsUnset(request.FollowUpLlm)) {
		body["followUpLlm"] = request.FollowUpLlm
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.LlmType)) {
		body["llmType"] = request.LlmType
	}

	if !tea.BoolValue(util.IsUnset(request.MultiQueryLlm)) {
		body["multiQueryLlm"] = request.MultiQueryLlm
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCriteria)) {
		body["queryCriteria"] = request.QueryCriteria
	}

	if !tea.BoolValue(util.IsUnset(request.RerankType)) {
		body["rerankType"] = request.RerankType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.SubQueryList)) {
		body["subQueryList"] = request.SubQueryList
	}

	if !tea.BoolValue(util.IsUnset(request.TextSearchParameter)) {
		body["textSearchParameter"] = request.TextSearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.TopK)) {
		body["topK"] = request.TopK
	}

	if !tea.BoolValue(util.IsUnset(request.VectorSearchParameter)) {
		body["vectorSearchParameter"] = request.VectorSearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.WithDocumentReference)) {
		body["withDocumentReference"] = request.WithDocumentReference
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunLibraryChatGeneration"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/run/library/chat/generation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunLibraryChatGenerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGeneration(workspaceId *string, request *RunLibraryChatGenerationRequest) (_result *RunLibraryChatGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunLibraryChatGenerationResponse{}
	_body, _err := client.RunLibraryChatGenerationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocumentWithOptions(workspaceId *string, request *UpdateDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		body["docId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.Meta)) {
		body["meta"] = request.Meta
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/updateDocument"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocument(workspaceId *string, request *UpdateDocumentRequest) (_result *UpdateDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDocumentResponse{}
	_body, _err := client.UpdateDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibraryWithOptions(workspaceId *string, request *UpdateLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IndexSetting)) {
		body["indexSetting"] = request.IndexSetting
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryName)) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLibrary"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/update"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibrary(workspaceId *string, request *UpdateLibraryRequest) (_result *UpdateLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLibraryResponse{}
	_body, _err := client.UpdateLibraryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocumentResponse
func (client *Client) UploadDocumentWithOptions(workspaceId *string, request *UploadDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["fileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["fileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.LibraryId)) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadDocument"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/api/library/document/upload"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @return UploadDocumentResponse
func (client *Client) UploadDocument(workspaceId *string, request *UploadDocumentRequest) (_result *UploadDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadDocumentResponse{}
	_body, _err := client.UploadDocumentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDocumentAdvance(workspaceId *string, request *UploadDocumentAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("DianJin"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	uploadDocumentReq := &UploadDocumentRequest{}
	openapiutil.Convert(request, uploadDocumentReq)
	if !tea.BoolValue(util.IsUnset(request.FileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.FileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		uploadDocumentReq.FileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	uploadDocumentResp, _err := client.UploadDocumentWithOptions(workspaceId, uploadDocumentReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = uploadDocumentResp
	return _result, _err
}
