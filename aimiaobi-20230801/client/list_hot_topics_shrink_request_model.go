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
	SetCreateTimeEnd(v string) *ListHotTopicsShrinkRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListHotTopicsShrinkRequest
	GetCreateTimeStart() *string
	SetCustomField(v string) *ListHotTopicsShrinkRequest
	GetCustomField() *string
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
	// The unique identifier of the business space.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The end of the creation time filter range (inclusive). The value must be in the `yyyy-MM-dd HH:mm:ss` format.
	//
	// example:
	//
	// 2026-06-04 23:59:59
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// The start of the creation time filter range (inclusive). The value must be in the `yyyy-MM-dd HH:mm:ss` format.
	//
	// example:
	//
	// 2026-06-01 00:00:00
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// Filters the results by a custom business field. The service performs an exact keyword match on this field. The value can be up to 255 characters long.
	//
	// example:
	//
	// biz-tag-001
	CustomField *string `json:"CustomField,omitempty" xml:"CustomField,omitempty"`
	// The maximum number of results to return for a single request. If this parameter is not specified, the service uses a default value.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. If you do not specify this parameter, the service returns the first page of results. You can get this token from the `NextToken` response parameter of the previous request.
	//
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// A list of topic IDs.
	TopicIdsShrink *string `json:"TopicIds,omitempty" xml:"TopicIds,omitempty"`
	// The keywords for a full-text search on hot topics.
	//
	// example:
	//
	// 根据热榜主题全文检索
	TopicQuery *string `json:"TopicQuery,omitempty" xml:"TopicQuery,omitempty"`
	// Filters the results by hot topic source. For a list of supported hot topic sources, call the `ListHotSources` operation.
	//
	// `Aggregation`: represents the aggregated list of national hot topics.
	//
	// example:
	//
	// Quark
	TopicSource *string `json:"TopicSource,omitempty" xml:"TopicSource,omitempty"`
	// Filters the results by data version.
	//
	// example:
	//
	// 数据版本筛选
	TopicVersion *string `json:"TopicVersion,omitempty" xml:"TopicVersion,omitempty"`
	// Filters the results by hot topic.
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// Specifies whether to include news in the response.
	//
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

func (s *ListHotTopicsShrinkRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListHotTopicsShrinkRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListHotTopicsShrinkRequest) GetCustomField() *string {
	return s.CustomField
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

func (s *ListHotTopicsShrinkRequest) SetCreateTimeEnd(v string) *ListHotTopicsShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetCreateTimeStart(v string) *ListHotTopicsShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListHotTopicsShrinkRequest) SetCustomField(v string) *ListHotTopicsShrinkRequest {
	s.CustomField = &v
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
