// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetLibraryResponseBody
	GetCost() *int64
	SetData(v *GetLibraryResponseBodyData) *GetLibraryResponseBody
	GetData() *GetLibraryResponseBodyData
	SetDataType(v string) *GetLibraryResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetLibraryResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetLibraryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLibraryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLibraryResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetLibraryResponseBody
	GetTime() *string
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetLibraryResponseBody) GetData() *GetLibraryResponseBodyData {
	return s.Data
}

func (s *GetLibraryResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetLibraryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetLibraryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLibraryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLibraryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLibraryResponseBody) GetTime() *string {
	return s.Time
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

func (s *GetLibraryResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetLibraryResponseBodyData) GetDocumentCount() *int64 {
	return s.DocumentCount
}

func (s *GetLibraryResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetLibraryResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetLibraryResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetLibraryResponseBodyData) GetIndexSetting() *GetLibraryResponseBodyDataIndexSetting {
	return s.IndexSetting
}

func (s *GetLibraryResponseBodyData) GetLibraryName() *string {
	return s.LibraryName
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

func (s *GetLibraryResponseBodyData) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetChunkStrategy() *GetLibraryResponseBodyDataIndexSettingChunkStrategy {
	return s.ChunkStrategy
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetModelConfig() *GetLibraryResponseBodyDataIndexSettingModelConfig {
	return s.ModelConfig
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetPromptRoleStyle() *string {
	return s.PromptRoleStyle
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetQueryEnhancer() *GetLibraryResponseBodyDataIndexSettingQueryEnhancer {
	return s.QueryEnhancer
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetRecallStrategy() *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	return s.RecallStrategy
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetTextIndexSetting() *GetLibraryResponseBodyDataIndexSettingTextIndexSetting {
	return s.TextIndexSetting
}

func (s *GetLibraryResponseBodyDataIndexSetting) GetVectorIndexSetting() *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting {
	return s.VectorIndexSetting
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

func (s *GetLibraryResponseBodyDataIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetDocTreeSplit() *bool {
	return s.DocTreeSplit
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetDocTreeSplitSize() *int32 {
	return s.DocTreeSplitSize
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetEnhanceGraph() *bool {
	return s.EnhanceGraph
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetEnhanceTable() *bool {
	return s.EnhanceTable
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetOverlap() *int32 {
	return s.Overlap
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetSentenceSplit() *bool {
	return s.SentenceSplit
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetSentenceSplitSize() *int32 {
	return s.SentenceSplitSize
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetSize() *int32 {
	return s.Size
}

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) GetSplit() *bool {
	return s.Split
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

func (s *GetLibraryResponseBodyDataIndexSettingChunkStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) GetTemperature() *float64 {
	return s.Temperature
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) GetTopP() *float64 {
	return s.TopP
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) SetTemperature(v float64) *GetLibraryResponseBodyDataIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) SetTopP(v float64) *GetLibraryResponseBodyDataIndexSettingModelConfig {
	s.TopP = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingModelConfig) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetEnableFollowUp() *bool {
	return s.EnableFollowUp
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetEnableMultiQuery() *bool {
	return s.EnableMultiQuery
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetEnableOpenQa() *bool {
	return s.EnableOpenQa
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetEnableQueryRewrite() *bool {
	return s.EnableQueryRewrite
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetEnableSession() *bool {
	return s.EnableSession
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetLocalKnowledgeId() *string {
	return s.LocalKnowledgeId
}

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) GetWithDocumentReference() *bool {
	return s.WithDocumentReference
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

func (s *GetLibraryResponseBodyDataIndexSettingQueryEnhancer) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) GetDocumentRankType() *string {
	return s.DocumentRankType
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) GetLimit() *int32 {
	return s.Limit
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) SetDocumentRankType(v string) *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) SetLimit(v int32) *GetLibraryResponseBodyDataIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

func (s *GetLibraryResponseBodyDataIndexSettingRecallStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GetIndexAnalyzer() *string {
	return s.IndexAnalyzer
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GetSearchAnalyzer() *string {
	return s.SearchAnalyzer
}

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *GetLibraryResponseBodyDataIndexSettingTextIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GetEmbeddingType() *string {
	return s.EmbeddingType
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *GetLibraryResponseBodyDataIndexSettingVectorIndexSetting) Validate() error {
	return dara.Validate(s)
}
