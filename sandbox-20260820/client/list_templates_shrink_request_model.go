// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTemplatesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplatesShrinkRequest
	GetNextToken() *string
	SetResourceTagsShrink(v string) *ListTemplatesShrinkRequest
	GetResourceTagsShrink() *string
	SetTeamID(v string) *ListTemplatesShrinkRequest
	GetTeamID() *string
}

type ListTemplatesShrinkRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Use the nextToken value returned in the previous response.
	//
	// example:
	//
	// cae5f900-8b1d-4c0e-9c2a-1a2b3c4d5e6f
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource tags.
	ResourceTagsShrink *string `json:"resourceTags,omitempty" xml:"resourceTags,omitempty"`
	// The team ID.
	//
	// example:
	//
	// 13b721e6-8cc8-5df2-af13-80316f7508af
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s ListTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesShrinkRequest) GetResourceTagsShrink() *string {
	return s.ResourceTagsShrink
}

func (s *ListTemplatesShrinkRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *ListTemplatesShrinkRequest) SetMaxResults(v int32) *ListTemplatesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetNextToken(v string) *ListTemplatesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetResourceTagsShrink(v string) *ListTemplatesShrinkRequest {
	s.ResourceTagsShrink = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTeamID(v string) *ListTemplatesShrinkRequest {
	s.TeamID = &v
	return s
}

func (s *ListTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
