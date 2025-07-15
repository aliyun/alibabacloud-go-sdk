// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicRecommendEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTopicRecommendEventListRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListTopicRecommendEventListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTopicRecommendEventListRequest
	GetNextToken() *string
}

type ListTopicRecommendEventListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 72
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTopicRecommendEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicRecommendEventListRequest) GoString() string {
	return s.String()
}

func (s *ListTopicRecommendEventListRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTopicRecommendEventListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTopicRecommendEventListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTopicRecommendEventListRequest) SetAgentKey(v string) *ListTopicRecommendEventListRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTopicRecommendEventListRequest) SetMaxResults(v int32) *ListTopicRecommendEventListRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTopicRecommendEventListRequest) SetNextToken(v string) *ListTopicRecommendEventListRequest {
	s.NextToken = &v
	return s
}

func (s *ListTopicRecommendEventListRequest) Validate() error {
	return dara.Validate(s)
}
