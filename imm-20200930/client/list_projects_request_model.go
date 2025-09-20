// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListProjectsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListProjectsRequest
	GetNextToken() *string
	SetPrefix(v string) *ListProjectsRequest
	GetPrefix() *string
	SetTag(v []*ListProjectsRequestTag) *ListProjectsRequest
	GetTag() []*ListProjectsRequestTag
}

type ListProjectsRequest struct {
	// The maximum number of entries to return. Valid values: 0 to 200. Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken. The operation returns the projects in lexicographical order starting from the location specified by NextToken.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDAx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The prefix used by the projects that you want to query. The prefix must be up to 128 characters in length.
	//
	// example:
	//
	// immtest
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The tags.
	Tag []*ListProjectsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectsRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListProjectsRequest) GetTag() []*ListProjectsRequestTag {
	return s.Tag
}

func (s *ListProjectsRequest) SetMaxResults(v int64) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetPrefix(v string) *ListProjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListProjectsRequest) SetTag(v []*ListProjectsRequestTag) *ListProjectsRequest {
	s.Tag = v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}

type ListProjectsRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequestTag) GoString() string {
	return s.String()
}

func (s *ListProjectsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListProjectsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListProjectsRequestTag) SetKey(v string) *ListProjectsRequestTag {
	s.Key = &v
	return s
}

func (s *ListProjectsRequestTag) SetValue(v string) *ListProjectsRequestTag {
	s.Value = &v
	return s
}

func (s *ListProjectsRequestTag) Validate() error {
	return dara.Validate(s)
}
