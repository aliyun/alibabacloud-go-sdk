// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v map[string]interface{}) *SearchContextRequest
	GetFilter() map[string]interface{}
	SetFormatted(v bool) *SearchContextRequest
	GetFormatted() *bool
	SetLimit(v int32) *SearchContextRequest
	GetLimit() *int32
	SetQuery(v string) *SearchContextRequest
	GetQuery() *string
	SetRetrievalOption(v string) *SearchContextRequest
	GetRetrievalOption() *string
	SetThreshold(v float64) *SearchContextRequest
	GetThreshold() *float64
}

type SearchContextRequest struct {
	// Filter conditions
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	// Whether to format
	//
	// example:
	//
	// true
	Formatted *bool `json:"formatted,omitempty" xml:"formatted,omitempty"`
	// Limit the number of items returned
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// Query content
	//
	// This parameter is required.
	//
	// example:
	//
	// How is the cost for consuming SLS logs in Flink calculated?
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Retrieval option
	//
	// example:
	//
	// reranker,llm_rank
	RetrievalOption *string `json:"retrievalOption,omitempty" xml:"retrievalOption,omitempty"`
	// Similarity threshold
	//
	// example:
	//
	// 0.3
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty"`
}

func (s SearchContextRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchContextRequest) GoString() string {
	return s.String()
}

func (s *SearchContextRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *SearchContextRequest) GetFormatted() *bool {
	return s.Formatted
}

func (s *SearchContextRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchContextRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchContextRequest) GetRetrievalOption() *string {
	return s.RetrievalOption
}

func (s *SearchContextRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *SearchContextRequest) SetFilter(v map[string]interface{}) *SearchContextRequest {
	s.Filter = v
	return s
}

func (s *SearchContextRequest) SetFormatted(v bool) *SearchContextRequest {
	s.Formatted = &v
	return s
}

func (s *SearchContextRequest) SetLimit(v int32) *SearchContextRequest {
	s.Limit = &v
	return s
}

func (s *SearchContextRequest) SetQuery(v string) *SearchContextRequest {
	s.Query = &v
	return s
}

func (s *SearchContextRequest) SetRetrievalOption(v string) *SearchContextRequest {
	s.RetrievalOption = &v
	return s
}

func (s *SearchContextRequest) SetThreshold(v float64) *SearchContextRequest {
	s.Threshold = &v
	return s
}

func (s *SearchContextRequest) Validate() error {
	return dara.Validate(s)
}
