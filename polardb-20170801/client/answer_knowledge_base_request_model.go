// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *AnswerKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
	SetMaxContextChars(v int32) *AnswerKnowledgeBaseRequest
	GetMaxContextChars() *int32
	SetQueryText(v string) *AnswerKnowledgeBaseRequest
	GetQueryText() *string
	SetRegionId(v string) *AnswerKnowledgeBaseRequest
	GetRegionId() *string
	SetRerankEnabled(v bool) *AnswerKnowledgeBaseRequest
	GetRerankEnabled() *bool
	SetReturnSources(v bool) *AnswerKnowledgeBaseRequest
	GetReturnSources() *bool
	SetScoreThreshold(v float64) *AnswerKnowledgeBaseRequest
	GetScoreThreshold() *float64
	SetSearchMode(v string) *AnswerKnowledgeBaseRequest
	GetSearchMode() *string
	SetSystemPrompt(v string) *AnswerKnowledgeBaseRequest
	GetSystemPrompt() *string
	SetTopK(v int32) *AnswerKnowledgeBaseRequest
	GetTopK() *int32
	SetUserInstructions(v string) *AnswerKnowledgeBaseRequest
	GetUserInstructions() *string
}

type AnswerKnowledgeBaseRequest struct {
	// The unique ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The maximum number of context characters. Valid values: 1000 to 32000.
	//
	// example:
	//
	// 16000
	MaxContextChars *int32 `json:"MaxContextChars,omitempty" xml:"MaxContextChars,omitempty"`
	// The user query text.
	//
	// This parameter is required.
	//
	// example:
	//
	// Summarize this year\\"s financial report
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable reranking. Default value: false.
	//
	// example:
	//
	// false
	RerankEnabled *bool `json:"RerankEnabled,omitempty" xml:"RerankEnabled,omitempty"`
	// Specifies whether to return citation sources. Default value: true.
	//
	// example:
	//
	// true
	ReturnSources *bool `json:"ReturnSources,omitempty" xml:"ReturnSources,omitempty"`
	// The similarity score threshold.
	//
	// example:
	//
	// 0.7
	ScoreThreshold *float64 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The search mode. Valid values: knn, rrf, precise, semantic, and balanced.
	//
	// example:
	//
	// semantic
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// The system prompt.
	//
	// example:
	//
	// This is a minimal test prompt for security verification.
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The number of results to recall during retrieval.
	//
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// The supplementary user instructions.
	//
	// example:
	//
	// Return a short summary and test result.
	UserInstructions *string `json:"UserInstructions,omitempty" xml:"UserInstructions,omitempty"`
}

func (s AnswerKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s AnswerKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *AnswerKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *AnswerKnowledgeBaseRequest) GetMaxContextChars() *int32 {
	return s.MaxContextChars
}

func (s *AnswerKnowledgeBaseRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *AnswerKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AnswerKnowledgeBaseRequest) GetRerankEnabled() *bool {
	return s.RerankEnabled
}

func (s *AnswerKnowledgeBaseRequest) GetReturnSources() *bool {
	return s.ReturnSources
}

func (s *AnswerKnowledgeBaseRequest) GetScoreThreshold() *float64 {
	return s.ScoreThreshold
}

func (s *AnswerKnowledgeBaseRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *AnswerKnowledgeBaseRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *AnswerKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *AnswerKnowledgeBaseRequest) GetUserInstructions() *string {
	return s.UserInstructions
}

func (s *AnswerKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *AnswerKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetMaxContextChars(v int32) *AnswerKnowledgeBaseRequest {
	s.MaxContextChars = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetQueryText(v string) *AnswerKnowledgeBaseRequest {
	s.QueryText = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetRegionId(v string) *AnswerKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetRerankEnabled(v bool) *AnswerKnowledgeBaseRequest {
	s.RerankEnabled = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetReturnSources(v bool) *AnswerKnowledgeBaseRequest {
	s.ReturnSources = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetScoreThreshold(v float64) *AnswerKnowledgeBaseRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetSearchMode(v string) *AnswerKnowledgeBaseRequest {
	s.SearchMode = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetSystemPrompt(v string) *AnswerKnowledgeBaseRequest {
	s.SystemPrompt = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetTopK(v int32) *AnswerKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) SetUserInstructions(v string) *AnswerKnowledgeBaseRequest {
	s.UserInstructions = &v
	return s
}

func (s *AnswerKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
