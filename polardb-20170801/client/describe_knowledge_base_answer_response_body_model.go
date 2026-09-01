// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseAnswerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetAgentId() *string
	SetAnswer(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetAnswer() *string
	SetCompletionTokens(v int32) *DescribeKnowledgeBaseAnswerResponseBody
	GetCompletionTokens() *int32
	SetErrorMessage(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetErrorMessage() *string
	SetErrorType(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetErrorType() *string
	SetLLMModelId(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetLLMModelId() *string
	SetPromptTokens(v int32) *DescribeKnowledgeBaseAnswerResponseBody
	GetPromptTokens() *int32
	SetQueryId(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetQueryId() *string
	SetRequestId(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetRequestId() *string
	SetSources(v []*DescribeKnowledgeBaseAnswerResponseBodySources) *DescribeKnowledgeBaseAnswerResponseBody
	GetSources() []*DescribeKnowledgeBaseAnswerResponseBodySources
	SetStatus(v string) *DescribeKnowledgeBaseAnswerResponseBody
	GetStatus() *string
}

type DescribeKnowledgeBaseAnswerResponseBody struct {
	// example:
	//
	// ******
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 财报的内容总结如下：******
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 3935
	CompletionTokens *int32 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// Space not found: pks-xxxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// not_found
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// example:
	//
	// OO1A6p8B******_xPUC
	LLMModelId *string `json:"LLMModelId,omitempty" xml:"LLMModelId,omitempty"`
	// example:
	//
	// 2459
	PromptTokens *int32 `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// example:
	//
	// R3BGbnBqcXN******.2a5a23c9-******-179970533d30
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sources   []*DescribeKnowledgeBaseAnswerResponseBodySources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeKnowledgeBaseAnswerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAnswerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetAnswer() *string {
	return s.Answer
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetCompletionTokens() *int32 {
	return s.CompletionTokens
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetErrorType() *string {
	return s.ErrorType
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetLLMModelId() *string {
	return s.LLMModelId
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetPromptTokens() *int32 {
	return s.PromptTokens
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetQueryId() *string {
	return s.QueryId
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetSources() []*DescribeKnowledgeBaseAnswerResponseBodySources {
	return s.Sources
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetAgentId(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetAnswer(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.Answer = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetCompletionTokens(v int32) *DescribeKnowledgeBaseAnswerResponseBody {
	s.CompletionTokens = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetErrorMessage(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetErrorType(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.ErrorType = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetLLMModelId(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.LLMModelId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetPromptTokens(v int32) *DescribeKnowledgeBaseAnswerResponseBody {
	s.PromptTokens = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetQueryId(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.QueryId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetRequestId(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetSources(v []*DescribeKnowledgeBaseAnswerResponseBodySources) *DescribeKnowledgeBaseAnswerResponseBody {
	s.Sources = v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) SetStatus(v string) *DescribeKnowledgeBaseAnswerResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBody) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKnowledgeBaseAnswerResponseBodySources struct {
	// example:
	//
	// {}
	ChunkMetadata map[string]interface{} `json:"ChunkMetadata,omitempty" xml:"ChunkMetadata,omitempty"`
	// example:
	//
	// 91b97b71-xxxx-xxxx-xxxx-33c6a6341cdc
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 财报.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// example:
	//
	// {}
	Metadata    map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	PageNumbers []*int32               `json:"PageNumbers,omitempty" xml:"PageNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// 财报
	ShardContent *string `json:"ShardContent,omitempty" xml:"ShardContent,omitempty"`
	// example:
	//
	// 8
	ShardIndex *int32 `json:"ShardIndex,omitempty" xml:"ShardIndex,omitempty"`
	// example:
	//
	// 0.75
	SimilarityScore *float64 `json:"SimilarityScore,omitempty" xml:"SimilarityScore,omitempty"`
	// example:
	//
	// 1
	SourceId *int32 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s DescribeKnowledgeBaseAnswerResponseBodySources) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseAnswerResponseBodySources) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetChunkMetadata() map[string]interface{} {
	return s.ChunkMetadata
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetFileId() *string {
	return s.FileId
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetFileName() *string {
	return s.FileName
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetPageNumbers() []*int32 {
	return s.PageNumbers
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetShardContent() *string {
	return s.ShardContent
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetShardIndex() *int32 {
	return s.ShardIndex
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetSimilarityScore() *float64 {
	return s.SimilarityScore
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) GetSourceId() *int32 {
	return s.SourceId
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetChunkMetadata(v map[string]interface{}) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.ChunkMetadata = v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetFileId(v string) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.FileId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetFileName(v string) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.FileName = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetMetadata(v map[string]interface{}) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.Metadata = v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetPageNumbers(v []*int32) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.PageNumbers = v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetShardContent(v string) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.ShardContent = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetShardIndex(v int32) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.ShardIndex = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetSimilarityScore(v float64) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.SimilarityScore = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) SetSourceId(v int32) *DescribeKnowledgeBaseAnswerResponseBodySources {
	s.SourceId = &v
	return s
}

func (s *DescribeKnowledgeBaseAnswerResponseBodySources) Validate() error {
	return dara.Validate(s)
}
