// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListSubscriptionsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListSubscriptionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSubscriptionsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListSubscriptionsRequest
	GetProjectName() *string
	SetSkip(v int32) *ListSubscriptionsRequest
	GetSkip() *int32
	SetTopicName(v string) *ListSubscriptionsRequest
	GetTopicName() *string
}

type ListSubscriptionsRequest struct {
	// example:
	//
	// dh
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListSubscriptionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSubscriptionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSubscriptionsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSubscriptionsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListSubscriptionsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListSubscriptionsRequest) SetKeyword(v string) *ListSubscriptionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListSubscriptionsRequest) SetMaxResults(v int32) *ListSubscriptionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSubscriptionsRequest) SetNextToken(v string) *ListSubscriptionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSubscriptionsRequest) SetProjectName(v string) *ListSubscriptionsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListSubscriptionsRequest) SetSkip(v int32) *ListSubscriptionsRequest {
	s.Skip = &v
	return s
}

func (s *ListSubscriptionsRequest) SetTopicName(v string) *ListSubscriptionsRequest {
	s.TopicName = &v
	return s
}

func (s *ListSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
