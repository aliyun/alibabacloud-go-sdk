// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListGroupsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListGroupsRequest
	GetProjectName() *string
	SetPure(v bool) *ListGroupsRequest
	GetPure() *bool
	SetSkip(v int32) *ListGroupsRequest
	GetSkip() *int32
}

type ListGroupsRequest struct {
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
}

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListGroupsRequest) GetPure() *bool {
	return s.Pure
}

func (s *ListGroupsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListGroupsRequest) SetKeyword(v string) *ListGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListGroupsRequest) SetMaxResults(v int32) *ListGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsRequest) SetNextToken(v string) *ListGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsRequest) SetProjectName(v string) *ListGroupsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListGroupsRequest) SetPure(v bool) *ListGroupsRequest {
	s.Pure = &v
	return s
}

func (s *ListGroupsRequest) SetSkip(v int32) *ListGroupsRequest {
	s.Skip = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
