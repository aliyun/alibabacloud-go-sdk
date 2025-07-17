// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPromptTemplatesRequest
	GetMaxResults() *int32
	SetName(v string) *ListPromptTemplatesRequest
	GetName() *string
	SetNextToken(v string) *ListPromptTemplatesRequest
	GetNextToken() *string
	SetType(v string) *ListPromptTemplatesRequest
	GetType() *string
}

type ListPromptTemplatesRequest struct {
	// The maximum number of returned entries.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The keyword that is used to search for templates.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The token that determines the start position of the query. Set this parameter to the value of the NextToken parameter that is returned from the last call.
	//
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The type of the template. Valid values: · System · Custom
	//
	// example:
	//
	// System
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPromptTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromptTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPromptTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPromptTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListPromptTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPromptTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListPromptTemplatesRequest) SetMaxResults(v int32) *ListPromptTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPromptTemplatesRequest) SetName(v string) *ListPromptTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListPromptTemplatesRequest) SetNextToken(v string) *ListPromptTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPromptTemplatesRequest) SetType(v string) *ListPromptTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListPromptTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
