// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotViewPointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListHotViewPointsRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListHotViewPointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHotViewPointsRequest
	GetNextToken() *string
	SetTopic(v string) *ListHotViewPointsRequest
	GetTopic() *string
	SetTopicSource(v string) *ListHotViewPointsRequest
	GetTopicSource() *string
}

type ListHotViewPointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 56
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

func (s ListHotViewPointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotViewPointsRequest) GoString() string {
	return s.String()
}

func (s *ListHotViewPointsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListHotViewPointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotViewPointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotViewPointsRequest) GetTopic() *string {
	return s.Topic
}

func (s *ListHotViewPointsRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListHotViewPointsRequest) SetAgentKey(v string) *ListHotViewPointsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotViewPointsRequest) SetMaxResults(v int32) *ListHotViewPointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotViewPointsRequest) SetNextToken(v string) *ListHotViewPointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotViewPointsRequest) SetTopic(v string) *ListHotViewPointsRequest {
	s.Topic = &v
	return s
}

func (s *ListHotViewPointsRequest) SetTopicSource(v string) *ListHotViewPointsRequest {
	s.TopicSource = &v
	return s
}

func (s *ListHotViewPointsRequest) Validate() error {
	return dara.Validate(s)
}
