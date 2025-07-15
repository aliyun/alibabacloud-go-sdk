// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFreshViewPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListFreshViewPointsRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListFreshViewPointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFreshViewPointsRequest
	GetNextToken() *string
	SetTopic(v string) *ListFreshViewPointsRequest
	GetTopic() *string
	SetTopicSource(v string) *ListFreshViewPointsRequest
	GetTopicSource() *string
}

type ListFreshViewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 6
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

func (s ListFreshViewPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFreshViewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListFreshViewPointsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListFreshViewPointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFreshViewPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreshViewPointsRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListFreshViewPointsRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListFreshViewPointsRequest) SetAgentKey(v string) *ListFreshViewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetMaxResults(v int32) *ListFreshViewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetNextToken(v string) *ListFreshViewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetTopic(v string) *ListFreshViewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListFreshViewPointsRequest) SetTopicSource(v string) *ListFreshViewPointsRequest {
	s.TopicSource = &v
	return s
}

func (s *ListFreshViewPointsRequest) Validate() error {
	return dara.Validate(s)
}
