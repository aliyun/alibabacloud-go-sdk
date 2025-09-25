// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetLibraryListResponseBody
	GetCost() *int64
	SetData(v *GetLibraryListResponseBodyData) *GetLibraryListResponseBody
	GetData() *GetLibraryListResponseBodyData
	SetDataType(v string) *GetLibraryListResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetLibraryListResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetLibraryListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLibraryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLibraryListResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetLibraryListResponseBody
	GetTime() *string
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetLibraryListResponseBody) GetData() *GetLibraryListResponseBodyData {
	return s.Data
}

func (s *GetLibraryListResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetLibraryListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetLibraryListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLibraryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLibraryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLibraryListResponseBody) GetTime() *string {
	return s.Time
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

func (s *GetLibraryListResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetLibraryListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetLibraryListResponseBodyData) GetRecords() []*GetLibraryListResponseBodyDataRecords {
	return s.Records
}

func (s *GetLibraryListResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *GetLibraryListResponseBodyData) GetTotalRecords() *int64 {
	return s.TotalRecords
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

func (s *GetLibraryListResponseBodyData) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecords) GetDescription() *string {
	return s.Description
}

func (s *GetLibraryListResponseBodyDataRecords) GetDocumentCount() *int64 {
	return s.DocumentCount
}

func (s *GetLibraryListResponseBodyDataRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetLibraryListResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetLibraryListResponseBodyDataRecords) GetId() *string {
	return s.Id
}

func (s *GetLibraryListResponseBodyDataRecords) GetIndexSetting() *GetLibraryListResponseBodyDataRecordsIndexSetting {
	return s.IndexSetting
}

func (s *GetLibraryListResponseBodyDataRecords) GetLibraryName() *string {
	return s.LibraryName
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

func (s *GetLibraryListResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetChunkStrategy() *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy {
	return s.ChunkStrategy
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetModelConfig() *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	return s.ModelConfig
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetPromptRoleStyle() *string {
	return s.PromptRoleStyle
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetQueryEnhancer() *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer {
	return s.QueryEnhancer
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetRecallStrategy() *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	return s.RecallStrategy
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetTextIndexSetting() *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting {
	return s.TextIndexSetting
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) GetVectorIndexSetting() *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting {
	return s.VectorIndexSetting
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

func (s *GetLibraryListResponseBodyDataRecordsIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetDocTreeSplit() *bool {
	return s.DocTreeSplit
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetDocTreeSplitSize() *int32 {
	return s.DocTreeSplitSize
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetEnhanceGraph() *bool {
	return s.EnhanceGraph
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetEnhanceTable() *bool {
	return s.EnhanceTable
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetOverlap() *int32 {
	return s.Overlap
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetSentenceSplit() *bool {
	return s.SentenceSplit
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetSentenceSplitSize() *int32 {
	return s.SentenceSplitSize
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetSize() *int32 {
	return s.Size
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) GetSplit() *bool {
	return s.Split
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

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingChunkStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) GetTemperature() *float64 {
	return s.Temperature
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) GetTopP() *float64 {
	return s.TopP
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) SetTemperature(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) SetTopP(v float64) *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig {
	s.TopP = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingModelConfig) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetEnableFollowUp() *bool {
	return s.EnableFollowUp
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetEnableMultiQuery() *bool {
	return s.EnableMultiQuery
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetEnableOpenQa() *bool {
	return s.EnableOpenQa
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetEnableQueryRewrite() *bool {
	return s.EnableQueryRewrite
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetEnableSession() *bool {
	return s.EnableSession
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetLocalKnowledgeId() *string {
	return s.LocalKnowledgeId
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) GetWithDocumentReference() *bool {
	return s.WithDocumentReference
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

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingQueryEnhancer) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) GetDocumentRankType() *string {
	return s.DocumentRankType
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) GetLimit() *int32 {
	return s.Limit
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) SetDocumentRankType(v string) *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) SetLimit(v int32) *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingRecallStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GetIndexAnalyzer() *string {
	return s.IndexAnalyzer
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GetSearchAnalyzer() *string {
	return s.SearchAnalyzer
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingTextIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GetEmbeddingType() *string {
	return s.EmbeddingType
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *GetLibraryListResponseBodyDataRecordsIndexSettingVectorIndexSetting) Validate() error {
	return dara.Validate(s)
}
