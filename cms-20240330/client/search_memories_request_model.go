// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *SearchMemoriesRequest
	GetAgentId() *string
	SetAppId(v string) *SearchMemoriesRequest
	GetAppId() *string
	SetFilters(v map[string]interface{}) *SearchMemoriesRequest
	GetFilters() map[string]interface{}
	SetQuery(v string) *SearchMemoriesRequest
	GetQuery() *string
	SetRerank(v bool) *SearchMemoriesRequest
	GetRerank() *bool
	SetRetrievalOption(v string) *SearchMemoriesRequest
	GetRetrievalOption() *string
	SetRunId(v string) *SearchMemoriesRequest
	GetRunId() *string
	SetSearchType(v string) *SearchMemoriesRequest
	GetSearchType() *string
	SetThreshold(v float64) *SearchMemoriesRequest
	GetThreshold() *float64
	SetTopK(v int32) *SearchMemoriesRequest
	GetTopK() *int32
	SetUserId(v string) *SearchMemoriesRequest
	GetUserId() *string
}

type SearchMemoriesRequest struct {
	// example:
	//
	// 972772996913709056
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// mm_480d961a1b5e4efe84603f4cbc0f
	AppId   *string                `json:"appId,omitempty" xml:"appId,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty" xml:"filters,omitempty"`
	// example:
	//
	// What I like
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// true
	Rerank *bool `json:"rerank,omitempty" xml:"rerank,omitempty"`
	// example:
	//
	// reranker,llm_rank
	RetrievalOption *string `json:"retrievalOption,omitempty" xml:"retrievalOption,omitempty"`
	// example:
	//
	// test_session_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// experience
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// example:
	//
	// 0.3
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// example:
	//
	// 1
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
	// example:
	//
	// test_session_001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchMemoriesRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoriesRequest) GetAppId() *string {
	return s.AppId
}

func (s *SearchMemoriesRequest) GetFilters() map[string]interface{} {
	return s.Filters
}

func (s *SearchMemoriesRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchMemoriesRequest) GetRerank() *bool {
	return s.Rerank
}

func (s *SearchMemoriesRequest) GetRetrievalOption() *string {
	return s.RetrievalOption
}

func (s *SearchMemoriesRequest) GetRunId() *string {
	return s.RunId
}

func (s *SearchMemoriesRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *SearchMemoriesRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *SearchMemoriesRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *SearchMemoriesRequest) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoriesRequest) SetAgentId(v string) *SearchMemoriesRequest {
	s.AgentId = &v
	return s
}

func (s *SearchMemoriesRequest) SetAppId(v string) *SearchMemoriesRequest {
	s.AppId = &v
	return s
}

func (s *SearchMemoriesRequest) SetFilters(v map[string]interface{}) *SearchMemoriesRequest {
	s.Filters = v
	return s
}

func (s *SearchMemoriesRequest) SetQuery(v string) *SearchMemoriesRequest {
	s.Query = &v
	return s
}

func (s *SearchMemoriesRequest) SetRerank(v bool) *SearchMemoriesRequest {
	s.Rerank = &v
	return s
}

func (s *SearchMemoriesRequest) SetRetrievalOption(v string) *SearchMemoriesRequest {
	s.RetrievalOption = &v
	return s
}

func (s *SearchMemoriesRequest) SetRunId(v string) *SearchMemoriesRequest {
	s.RunId = &v
	return s
}

func (s *SearchMemoriesRequest) SetSearchType(v string) *SearchMemoriesRequest {
	s.SearchType = &v
	return s
}

func (s *SearchMemoriesRequest) SetThreshold(v float64) *SearchMemoriesRequest {
	s.Threshold = &v
	return s
}

func (s *SearchMemoriesRequest) SetTopK(v int32) *SearchMemoriesRequest {
	s.TopK = &v
	return s
}

func (s *SearchMemoriesRequest) SetUserId(v string) *SearchMemoriesRequest {
	s.UserId = &v
	return s
}

func (s *SearchMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
