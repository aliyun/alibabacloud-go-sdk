// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListProjectsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListProjectsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectsRequest
	GetNextToken() *string
	SetPure(v bool) *ListProjectsRequest
	GetPure() *bool
	SetSkip(v int32) *ListProjectsRequest
	GetSkip() *int32
}

type ListProjectsRequest struct {
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
	// example:
	//
	// false
	Pure *bool `json:"Pure,omitempty" xml:"Pure,omitempty"`
	// example:
	//
	// 1
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsRequest) GetPure() *bool {
	return s.Pure
}

func (s *ListProjectsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListProjectsRequest) SetKeyword(v string) *ListProjectsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetPure(v bool) *ListProjectsRequest {
	s.Pure = &v
	return s
}

func (s *ListProjectsRequest) SetSkip(v int32) *ListProjectsRequest {
	s.Skip = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
