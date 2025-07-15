// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListHotTopicsShrinkRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListHotTopicsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHotTopicsShrinkRequest
	GetNextToken() *string
	SetTopicIdsShrink(v string) *ListHotTopicsShrinkRequest
	GetTopicIdsShrink() *string
	SetTopicQuery(v string) *ListHotTopicsShrinkRequest
	GetTopicQuery() *string
	SetTopicSource(v string) *ListHotTopicsShrinkRequest
	GetTopicSource() *string
	SetTopicVersion(v string) *ListHotTopicsShrinkRequest
	GetTopicVersion() *string
	SetTopicsShrink(v string) *ListHotTopicsShrinkRequest
	GetTopicsShrink() *string
	SetWithNews(v bool) *ListHotTopicsShrinkRequest
	GetWithNews() *bool
}

type ListHotTopicsShrinkRequest struct {
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
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TopicIdsShrink *string `json:"TopicIds,omitempty" xml:"TopicIds,omitempty"`
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
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// example:
	//
	// true
	WithNews *bool `json:"WithNews,omitempty" xml:"WithNews,omitempty"`
}

func (s ListHotTopicsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotTopicsShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListHotTopicsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotTopicsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotTopicsShrinkRequest) GetTopicIdsShrink() *string {
	return s.TopicIdsShrink
}

func (s *ListHotTopicsShrinkRequest) GetTopicQuery() *string {
	return s.TopicQuery
}

func (s *ListHotTopicsShrinkRequest) GetTopicSource() *string {
	return s.TopicSource
}

func (s *ListHotTopicsShrinkRequest) GetTopicVersion() *string {
	return s.TopicVersion
}

func (s *ListHotTopicsShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *ListHotTopicsShrinkRequest) GetWithNews() *bool {
	return s.WithNews
}

func (s *ListHotTopicsShrinkRequest) SetAgentKey(v string) *ListHotTopicsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetMaxResults(v int32) *ListHotTopicsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetNextToken(v string) *ListHotTopicsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicIdsShrink(v string) *ListHotTopicsShrinkRequest {
	s.TopicIdsShrink = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicQuery(v string) *ListHotTopicsShrinkRequest {
	s.TopicQuery = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicSource(v string) *ListHotTopicsShrinkRequest {
	s.TopicSource = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicVersion(v string) *ListHotTopicsShrinkRequest {
	s.TopicVersion = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetTopicsShrink(v string) *ListHotTopicsShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetWithNews(v bool) *ListHotTopicsShrinkRequest {
	s.WithNews = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
