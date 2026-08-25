// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTeamsRequest
	GetMaxResults() *int32
	SetNameLike(v string) *ListTeamsRequest
	GetNameLike() *string
	SetNextToken(v string) *ListTeamsRequest
	GetNextToken() *string
}

type ListTeamsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// team
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// example:
	//
	// dGVhbS1vZmZzZXQ6MTA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListTeamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsRequest) GoString() string {
	return s.String()
}

func (s *ListTeamsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListTeamsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamsRequest) SetMaxResults(v int32) *ListTeamsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamsRequest) SetNameLike(v string) *ListTeamsRequest {
	s.NameLike = &v
	return s
}

func (s *ListTeamsRequest) SetNextToken(v string) *ListTeamsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTeamsRequest) Validate() error {
	return dara.Validate(s)
}
