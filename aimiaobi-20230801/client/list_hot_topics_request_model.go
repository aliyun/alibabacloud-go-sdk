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
	SetCreateTimeEnd(v string) *ListHotTopicsRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListHotTopicsRequest
	GetCreateTimeStart() *string
	SetCustomField(v string) *ListHotTopicsRequest
	GetCustomField() *string
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
	TopicIds []*string `json:"TopicIds,omitempty" xml:"TopicIds,omitempty" type:"Repeated"`
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
	Topics []*string `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// Specifies whether to include news in the response.
	//
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

func (s *ListHotTopicsRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListHotTopicsRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListHotTopicsRequest) GetCustomField() *string {
	return s.CustomField
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

func (s *ListHotTopicsRequest) SetCreateTimeEnd(v string) *ListHotTopicsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListHotTopicsRequest) SetCreateTimeStart(v string) *ListHotTopicsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListHotTopicsRequest) SetCustomField(v string) *ListHotTopicsRequest {
	s.CustomField = &v
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
