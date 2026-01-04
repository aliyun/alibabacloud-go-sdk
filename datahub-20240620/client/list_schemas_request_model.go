// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSchemasRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSchemasRequest
	GetNextToken() *string
	SetProjectName(v string) *ListSchemasRequest
	GetProjectName() *string
	SetSkip(v int32) *ListSchemasRequest
	GetSkip() *int32
	SetTopicName(v string) *ListSchemasRequest
	GetTopicName() *string
}

type ListSchemasRequest struct {
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

func (s ListSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListSchemasRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSchemasRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSchemasRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSchemasRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListSchemasRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListSchemasRequest) SetMaxResults(v int32) *ListSchemasRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSchemasRequest) SetNextToken(v string) *ListSchemasRequest {
	s.NextToken = &v
	return s
}

func (s *ListSchemasRequest) SetProjectName(v string) *ListSchemasRequest {
	s.ProjectName = &v
	return s
}

func (s *ListSchemasRequest) SetSkip(v int32) *ListSchemasRequest {
	s.Skip = &v
	return s
}

func (s *ListSchemasRequest) SetTopicName(v string) *ListSchemasRequest {
	s.TopicName = &v
	return s
}

func (s *ListSchemasRequest) Validate() error {
	return dara.Validate(s)
}
