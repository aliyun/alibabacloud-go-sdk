// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemoriesRequest
	GetNextToken() *string
}

type ListMemoriesRequest struct {
	// The number of entries per page in a paged query. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// dc270401186b433f975d7e1faaa3xxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesRequest) GoString() string {
	return s.String()
}

func (s *ListMemoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoriesRequest) SetMaxResults(v int32) *ListMemoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoriesRequest) SetNextToken(v string) *ListMemoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoriesRequest) Validate() error {
	return dara.Validate(s)
}
