// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTeamsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTeamsShrinkRequest
	GetNextToken() *string
	SetTenantContextShrink(v string) *ListTeamsShrinkRequest
	GetTenantContextShrink() *string
}

type ListTeamsShrinkRequest struct {
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2023-05-15T11:29Z
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s ListTeamsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTeamsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTeamsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListTeamsShrinkRequest) SetMaxResults(v int32) *ListTeamsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTeamsShrinkRequest) SetNextToken(v string) *ListTeamsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTeamsShrinkRequest) SetTenantContextShrink(v string) *ListTeamsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListTeamsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
