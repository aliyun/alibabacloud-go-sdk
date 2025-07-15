// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimedViewAttitudeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTimedViewAttitudeRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListTimedViewAttitudeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTimedViewAttitudeRequest
	GetNextToken() *string
	SetTopic(v string) *ListTimedViewAttitudeRequest
	GetTopic() *string
	SetTopicSource(v string) *ListTimedViewAttitudeRequest
	GetTopicSource() *string
}

type ListTimedViewAttitudeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 53
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

func (s ListTimedViewAttitudeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTimedViewAttitudeRequest) GoString() string {
	return s.String()
}

func (s *ListTimedViewAttitudeRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTimedViewAttitudeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTimedViewAttitudeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTimedViewAttitudeRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListTimedViewAttitudeRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListTimedViewAttitudeRequest) SetAgentKey(v string) *ListTimedViewAttitudeRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetMaxResults(v int32) *ListTimedViewAttitudeRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetNextToken(v string) *ListTimedViewAttitudeRequest {
	s.NextToken = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetTopic(v string) *ListTimedViewAttitudeRequest {
	s.Topic = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) SetTopicSource(v string) *ListTimedViewAttitudeRequest {
	s.TopicSource = &v
	return s
}

func (s *ListTimedViewAttitudeRequest) Validate() error {
	return dara.Validate(s)
}
