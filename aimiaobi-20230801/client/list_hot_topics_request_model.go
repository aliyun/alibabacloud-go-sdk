// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListHotTopicsRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListHotTopicsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHotTopicsRequest
	GetNextToken() *string
	SetTopicIds(v []*string) *ListHotTopicsRequest
	GetTopicIds() []*string
	SetTopicQuery(v string) *ListHotTopicsRequest
	GetTopicQuery() *string
	SetTopicSource(v string) *ListHotTopicsRequest
	GetTopicSource() *string
	SetTopicVersion(v string) *ListHotTopicsRequest
	GetTopicVersion() *string
	SetTopics(v []*string) *ListHotTopicsRequest
	GetTopics() []*string
	SetWithNews(v bool) *ListHotTopicsRequest
	GetWithNews() *bool
}

type ListHotTopicsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TopicIds  []*string `json:"TopicIds,omitempty" xml:"TopicIds,omitempty" type:"Repeated"`
	// example:
	//
	// 根据热榜主题全文检索
	TopicQuery *string `json:"TopicQuery,omitempty" xml:"TopicQuery,omitempty"`
	// example:
	//
	// 热榜源筛选，支持的热榜源。热榜源详见API：ListHotSources
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// example:
	//
	// 数据版本筛选
	TopicVersion *string   `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	Topics       []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// example:
	//
	// true
	WithNews *bool `json:"WithNews,omitempty" xml:"WithNews,omitempty"`
}

func (s ListHotTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListHotTopicsRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListHotTopicsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotTopicsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotTopicsRequest) GetTopicIds() []*string {
	return s.TopicIds
}

func (s *ListHotTopicsRequest) GetTopicQuery() *string {
	return s.TopicQuery
}

func (s *ListHotTopicsRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListHotTopicsRequest) GetTopicVersion() *string {
	return s.TopicVersion
}

func (s *ListHotTopicsRequest) GetTopics() []*string {
	return s.Topics
}

func (s *ListHotTopicsRequest) GetWithNews() *bool {
	return s.WithNews
}

func (s *ListHotTopicsRequest) SetAgentKey(v string) *ListHotTopicsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotTopicsRequest) SetMaxResults(v int32) *ListHotTopicsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicsRequest) SetNextToken(v string) *ListHotTopicsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopicIds(v []*string) *ListHotTopicsRequest {
	s.TopicIds = v
	return s
}

func (s *ListHotTopicsRequest) SetTopicQuery(v string) *ListHotTopicsRequest {
	s.TopicQuery = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopicSource(v string) *ListHotTopicsRequest {
	s.TopicSource = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopicVersion(v string) *ListHotTopicsRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListHotTopicsRequest) SetTopics(v []*string) *ListHotTopicsRequest {
	s.Topics = v
	return s
}

func (s *ListHotTopicsRequest) SetWithNews(v bool) *ListHotTopicsRequest {
	s.WithNews = &v
	return s
}

func (s *ListHotTopicsRequest) Validate() error {
	return dara.Validate(s)
}
