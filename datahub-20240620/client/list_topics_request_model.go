// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListTopicsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListTopicsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTopicsRequest
	GetNextToken() *string
	SetProjectName(v string) *ListTopicsRequest
	GetProjectName() *string
	SetPure(v bool) *ListTopicsRequest
	GetPure() *bool
	SetSkip(v int32) *ListTopicsRequest
	GetSkip() *int32
}

type ListTopicsRequest struct {
	// The filter keyword for a paged query.
	//
	// example:
	//
	// dh
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The maximum number of records to return in a paged query.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. If NextToken is empty, paged query starts from the beginning. Otherwise, paged query starts from the position where the previous query ended.
	//
	// example:
	//
	// 9892074a2a89600ae4b0d5a34fb99a3f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to return only primary key information.
	//
	// example:
	//
	// false
	Pure *bool `json:"Pure,omitempty" xml:"Pure,omitempty"`
	// The number of records to skip in a paged query.
	//
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTopicsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTopicsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTopicsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTopicsRequest) GetPure() *bool {
	return s.Pure
}

func (s *ListTopicsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListTopicsRequest) SetKeyword(v string) *ListTopicsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTopicsRequest) SetMaxResults(v int32) *ListTopicsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTopicsRequest) SetNextToken(v string) *ListTopicsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTopicsRequest) SetProjectName(v string) *ListTopicsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTopicsRequest) SetPure(v bool) *ListTopicsRequest {
	s.Pure = &v
	return s
}

func (s *ListTopicsRequest) SetSkip(v int32) *ListTopicsRequest {
	s.Skip = &v
	return s
}

func (s *ListTopicsRequest) Validate() error {
	return dara.Validate(s)
}
