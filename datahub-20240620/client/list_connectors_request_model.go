// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListConnectorsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListConnectorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListConnectorsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListConnectorsRequest
	GetProjectName() *string
	SetPure(v bool) *ListConnectorsRequest
	GetPure() *bool
	SetSkip(v int32) *ListConnectorsRequest
	GetSkip() *int32
	SetTopicName(v string) *ListConnectorsRequest
	GetTopicName() *string
}

type ListConnectorsRequest struct {
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
	// false
	Pure *bool `json:"Pure,omitempty" xml:"Pure,omitempty"`
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

func (s ListConnectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConnectorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConnectorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConnectorsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListConnectorsRequest) GetPure() *bool {
	return s.Pure
}

func (s *ListConnectorsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListConnectorsRequest) GetTopicName() *string {
	return s.TopicName
}

func (s *ListConnectorsRequest) SetKeyword(v string) *ListConnectorsRequest {
	s.Keyword = &v
	return s
}

func (s *ListConnectorsRequest) SetMaxResults(v int32) *ListConnectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConnectorsRequest) SetNextToken(v string) *ListConnectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListConnectorsRequest) SetProjectName(v string) *ListConnectorsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListConnectorsRequest) SetPure(v bool) *ListConnectorsRequest {
	s.Pure = &v
	return s
}

func (s *ListConnectorsRequest) SetSkip(v int32) *ListConnectorsRequest {
	s.Skip = &v
	return s
}

func (s *ListConnectorsRequest) SetTopicName(v string) *ListConnectorsRequest {
	s.TopicName = &v
	return s
}

func (s *ListConnectorsRequest) Validate() error {
	return dara.Validate(s)
}
