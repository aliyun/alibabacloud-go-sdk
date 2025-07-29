// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotTopicSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListHotTopicSummariesRequest
	GetCategory() *string
	SetHotTopic(v string) *ListHotTopicSummariesRequest
	GetHotTopic() *string
	SetHotTopicVersion(v string) *ListHotTopicSummariesRequest
	GetHotTopicVersion() *string
	SetMaxResults(v int32) *ListHotTopicSummariesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHotTopicSummariesRequest
	GetNextToken() *string
}

type ListHotTopicSummariesRequest struct {
	// example:
	//
	// xx
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// xx
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_12
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// JlroP3CjgQh5PQDlH3ArzADkBTPZgVqo+64jhZRglNq0mEYoV5SlGb/Juvo8CdfYE9rlwEr2pIJQwdaYotak9g==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListHotTopicSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotTopicSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListHotTopicSummariesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListHotTopicSummariesRequest) GetHotTopic() *string {
	return s.HotTopic
}

func (s *ListHotTopicSummariesRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *ListHotTopicSummariesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotTopicSummariesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotTopicSummariesRequest) SetCategory(v string) *ListHotTopicSummariesRequest {
	s.Category = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetHotTopic(v string) *ListHotTopicSummariesRequest {
	s.HotTopic = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetHotTopicVersion(v string) *ListHotTopicSummariesRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetMaxResults(v int32) *ListHotTopicSummariesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotTopicSummariesRequest) SetNextToken(v string) *ListHotTopicSummariesRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotTopicSummariesRequest) Validate() error {
	return dara.Validate(s)
}
