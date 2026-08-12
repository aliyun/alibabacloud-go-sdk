// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDataPipelinesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataPipelinesRequest
	GetNextToken() *string
}

type ListDataPipelinesRequest struct {
	// The maximum number of results per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDataPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListDataPipelinesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataPipelinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataPipelinesRequest) SetMaxResults(v int32) *ListDataPipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataPipelinesRequest) SetNextToken(v string) *ListDataPipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
