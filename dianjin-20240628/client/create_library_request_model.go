// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateLibraryRequest
	GetDescription() *string
	SetIndexSetting(v *CreateLibraryRequestIndexSetting) *CreateLibraryRequest
	GetIndexSetting() *CreateLibraryRequestIndexSetting
	SetLibraryName(v string) *CreateLibraryRequest
	GetLibraryName() *string
}

type CreateLibraryRequest struct {
	// This parameter is required.
	Description  *string                           `json:"description,omitempty" xml:"description,omitempty"`
	IndexSetting *CreateLibraryRequestIndexSetting `json:"indexSetting,omitempty" xml:"indexSetting,omitempty" type:"Struct"`
	// This parameter is required.
	LibraryName *string `json:"libraryName,omitempty" xml:"libraryName,omitempty"`
}

func (s CreateLibraryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLibraryRequest) GetIndexSetting() *CreateLibraryRequestIndexSetting {
	return s.IndexSetting
}

func (s *CreateLibraryRequest) GetLibraryName() *string {
	return s.LibraryName
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

func (s *CreateLibraryRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSetting) GetChunkStrategy() *CreateLibraryRequestIndexSettingChunkStrategy {
	return s.ChunkStrategy
}

func (s *CreateLibraryRequestIndexSetting) GetModelConfig() *CreateLibraryRequestIndexSettingModelConfig {
	return s.ModelConfig
}

func (s *CreateLibraryRequestIndexSetting) GetPromptRoleStyle() *string {
	return s.PromptRoleStyle
}

func (s *CreateLibraryRequestIndexSetting) GetQueryEnhancer() *CreateLibraryRequestIndexSettingQueryEnhancer {
	return s.QueryEnhancer
}

func (s *CreateLibraryRequestIndexSetting) GetRecallStrategy() *CreateLibraryRequestIndexSettingRecallStrategy {
	return s.RecallStrategy
}

func (s *CreateLibraryRequestIndexSetting) GetTextIndexSetting() *CreateLibraryRequestIndexSettingTextIndexSetting {
	return s.TextIndexSetting
}

func (s *CreateLibraryRequestIndexSetting) GetVectorIndexSetting() *CreateLibraryRequestIndexSettingVectorIndexSetting {
	return s.VectorIndexSetting
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

func (s *CreateLibraryRequestIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingChunkStrategy) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetDocTreeSplit() *bool {
	return s.DocTreeSplit
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetDocTreeSplitSize() *int32 {
	return s.DocTreeSplitSize
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetEnhanceGraph() *bool {
	return s.EnhanceGraph
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetEnhanceTable() *bool {
	return s.EnhanceTable
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetOverlap() *int32 {
	return s.Overlap
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetSentenceSplit() *bool {
	return s.SentenceSplit
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetSentenceSplitSize() *int32 {
	return s.SentenceSplitSize
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetSize() *int32 {
	return s.Size
}

func (s *CreateLibraryRequestIndexSettingChunkStrategy) GetSplit() *bool {
	return s.Split
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

func (s *CreateLibraryRequestIndexSettingChunkStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingModelConfig) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingModelConfig) GetTemperature() *float64 {
	return s.Temperature
}

func (s *CreateLibraryRequestIndexSettingModelConfig) GetTopP() *float64 {
	return s.TopP
}

func (s *CreateLibraryRequestIndexSettingModelConfig) SetTemperature(v float64) *CreateLibraryRequestIndexSettingModelConfig {
	s.Temperature = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingModelConfig) SetTopP(v float64) *CreateLibraryRequestIndexSettingModelConfig {
	s.TopP = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingModelConfig) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingQueryEnhancer) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetEnableFollowUp() *bool {
	return s.EnableFollowUp
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetEnableMultiQuery() *bool {
	return s.EnableMultiQuery
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetEnableOpenQa() *bool {
	return s.EnableOpenQa
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetEnableQueryRewrite() *bool {
	return s.EnableQueryRewrite
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetEnableSession() *bool {
	return s.EnableSession
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetLocalKnowledgeId() *string {
	return s.LocalKnowledgeId
}

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) GetWithDocumentReference() *bool {
	return s.WithDocumentReference
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

func (s *CreateLibraryRequestIndexSettingQueryEnhancer) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingRecallStrategy) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) GetDocumentRankType() *string {
	return s.DocumentRankType
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) GetLimit() *int32 {
	return s.Limit
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) SetDocumentRankType(v string) *CreateLibraryRequestIndexSettingRecallStrategy {
	s.DocumentRankType = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) SetLimit(v int32) *CreateLibraryRequestIndexSettingRecallStrategy {
	s.Limit = &v
	return s
}

func (s *CreateLibraryRequestIndexSettingRecallStrategy) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingTextIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) GetIndexAnalyzer() *string {
	return s.IndexAnalyzer
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) GetSearchAnalyzer() *string {
	return s.SearchAnalyzer
}

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *CreateLibraryRequestIndexSettingTextIndexSetting) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateLibraryRequestIndexSettingVectorIndexSetting) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) GetCategory() *string {
	return s.Category
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) GetEmbeddingType() *string {
	return s.EmbeddingType
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) GetEnable() *bool {
	return s.Enable
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) GetRankThreshold() *float64 {
	return s.RankThreshold
}

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) GetTopK() *int32 {
	return s.TopK
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

func (s *CreateLibraryRequestIndexSettingVectorIndexSetting) Validate() error {
	return dara.Validate(s)
}
