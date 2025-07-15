// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebReviewPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListWebReviewPointsRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListWebReviewPointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListWebReviewPointsRequest
	GetNextToken() *string
	SetTopic(v string) *ListWebReviewPointsRequest
	GetTopic() *string
	SetTopicSource(v string) *ListWebReviewPointsRequest
	GetTopicSource() *string
}

type ListWebReviewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 81
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜主题
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 热榜源
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
}

func (s ListWebReviewPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListWebReviewPointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWebReviewPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWebReviewPointsRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListWebReviewPointsRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListWebReviewPointsRequest) SetAgentKey(v string) *ListWebReviewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetMaxResults(v int32) *ListWebReviewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetNextToken(v string) *ListWebReviewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetTopic(v string) *ListWebReviewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListWebReviewPointsRequest) SetTopicSource(v string) *ListWebReviewPointsRequest {
	s.TopicSource = &v
	return s
}

func (s *ListWebReviewPointsRequest) Validate() error {
	return dara.Validate(s)
}
