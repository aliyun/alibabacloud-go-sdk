// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplatesRequest
	GetNextToken() *string
	SetTeamID(v string) *ListTemplatesRequest
	GetTeamID() *string
}

type ListTemplatesRequest struct {
	// The maximum number of entries to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// eyJNYXhSZXN1bHRzIjoxMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 88a4c762-b0ce-4661-9413-578b2309e60f
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetTeamID(v string) *ListTemplatesRequest {
	s.TeamID = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
