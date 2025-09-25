// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLibraryRequest
	GetDescription() *string
	SetIndexSetting(v *UpdateLibraryRequestIndexSetting) *UpdateLibraryRequest
	GetIndexSetting() *UpdateLibraryRequestIndexSetting
	SetLibraryId(v string) *UpdateLibraryRequest
	GetLibraryId() *string
	SetLibraryName(v string) *UpdateLibraryRequest
	GetLibraryName() *string
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLibraryRequest) GetIndexSetting() *UpdateLibraryRequestIndexSetting {
	return s.IndexSetting
}

func (s *UpdateLibraryRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *UpdateLibraryRequest) GetLibraryName() *string {
	return s.LibraryName
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

func (s *UpdateLibraryRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSetting) GetChunkStrategy() *UpdateLibraryRequestIndexSettingChunkStrategy {
	return s.ChunkStrategy
}

func (s *UpdateLibraryRequestIndexSetting) GetModelConfig() *UpdateLibraryRequestIndexSettingModelConfig {
	return s.ModelConfig
}

func (s *UpdateLibraryRequestIndexSetting) GetPromptRoleStyle() *string {
	return s.PromptRoleStyle
}

func (s *UpdateLibraryRequestIndexSetting) GetQueryEnhancer() *UpdateLibraryRequestIndexSettingQueryEnhancer {
	return s.QueryEnhancer
}

func (s *UpdateLibraryRequestIndexSetting) GetRecallStrategy() *UpdateLibraryRequestIndexSettingRecallStrategy {
	return s.RecallStrategy
}

func (s *UpdateLibraryRequestIndexSetting) GetTextIndexSetting() *UpdateLibraryRequestIndexSettingTextIndexSetting {
	return s.TextIndexSetting
}

func (s *UpdateLibraryRequestIndexSetting) GetVectorIndexSetting() *UpdateLibraryRequestIndexSettingVectorIndexSetting {
	return s.VectorIndexSetting
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

func (s *UpdateLibraryRequestIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetDocTreeSplit() *bool {
	return s.DocTreeSplit
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetDocTreeSplitSize() *int32 {
	return s.DocTreeSplitSize
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetEnhanceGraph() *bool {
	return s.EnhanceGraph
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetEnhanceTable() *bool {
	return s.EnhanceTable
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetOverlap() *int32 {
	return s.Overlap
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetSentenceSplit() *bool {
	return s.SentenceSplit
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetSentenceSplitSize() *int32 {
	return s.SentenceSplitSize
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetSize() *int32 {
	return s.Size
}

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) GetSplit() *bool {
	return s.Split
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

func (s *UpdateLibraryRequestIndexSettingChunkStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) GetTemperature() *float64 {
	return s.Temperature
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) GetTopP() *float64 {
	return s.TopP
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) SetTemperature(v float64) *UpdateLibraryRequestIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) SetTopP(v float64) *UpdateLibraryRequestIndexSettingModelConfig {
	s.TopP = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingModelConfig) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetEnableFollowUp() *bool {
	return s.EnableFollowUp
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetEnableMultiQuery() *bool {
	return s.EnableMultiQuery
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetEnableOpenQa() *bool {
	return s.EnableOpenQa
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetEnableQueryRewrite() *bool {
	return s.EnableQueryRewrite
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetEnableSession() *bool {
	return s.EnableSession
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetLocalKnowledgeId() *string {
	return s.LocalKnowledgeId
}

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) GetWithDocumentReference() *bool {
	return s.WithDocumentReference
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

func (s *UpdateLibraryRequestIndexSettingQueryEnhancer) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) GetDocumentRankType() *string {
	return s.DocumentRankType
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) GetLimit() *int32 {
	return s.Limit
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) SetDocumentRankType(v string) *UpdateLibraryRequestIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) SetLimit(v int32) *UpdateLibraryRequestIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

func (s *UpdateLibraryRequestIndexSettingRecallStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) GetIndexAnalyzer() *string {
	return s.IndexAnalyzer
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) GetSearchAnalyzer() *string {
	return s.SearchAnalyzer
}

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *UpdateLibraryRequestIndexSettingTextIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s UpdateLibraryRequestIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) GetEmbeddingType() *string {
	return s.EmbeddingType
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *UpdateLibraryRequestIndexSettingVectorIndexSetting) Validate() error {
	return dara.Validate(s)
}
