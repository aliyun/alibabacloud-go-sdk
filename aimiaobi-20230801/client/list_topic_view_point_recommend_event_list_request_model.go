// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicViewPointRecommendEventListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListTopicViewPointRecommendEventListRequest
	GetAgentKey() *string
	SetId(v string) *ListTopicViewPointRecommendEventListRequest
	GetId() *string
	SetMaxResults(v int32) *ListTopicViewPointRecommendEventListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTopicViewPointRecommendEventListRequest
	GetNextToken() *string
}

type ListTopicViewPointRecommendEventListRequest struct {
	// Unique identifier for the workspace: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Hot spot event ID.
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Maximum number of results to return. Defaults to 5.
	//
	// example:
	//
	// 66
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next page.
	//
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTopicViewPointRecommendEventListRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicViewPointRecommendEventListRequest) GoString() string {
	return s.String()
}

func (s *ListTopicViewPointRecommendEventListRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListTopicViewPointRecommendEventListRequest) GetId() *string {
	return s.Id
}

func (s *ListTopicViewPointRecommendEventListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTopicViewPointRecommendEventListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTopicViewPointRecommendEventListRequest) SetAgentKey(v string) *ListTopicViewPointRecommendEventListRequest {
	s.AgentKey = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) SetId(v string) *ListTopicViewPointRecommendEventListRequest {
	s.Id = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) SetMaxResults(v int32) *ListTopicViewPointRecommendEventListRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) SetNextToken(v string) *ListTopicViewPointRecommendEventListRequest {
	s.NextToken = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListRequest) Validate() error {
	return dara.Validate(s)
}
