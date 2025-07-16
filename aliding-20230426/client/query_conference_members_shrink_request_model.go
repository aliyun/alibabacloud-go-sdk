// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QueryConferenceMembersShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QueryConferenceMembersShrinkRequest
	GetNextToken() *string
	SetTenantContextShrink(v string) *QueryConferenceMembersShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *QueryConferenceMembersShrinkRequest
	GetConferenceId() *string
}

type QueryConferenceMembersShrinkRequest struct {
	// example:
	//
	// 300
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 123000000
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryConferenceMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceMembersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryConferenceMembersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryConferenceMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryConferenceMembersShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryConferenceMembersShrinkRequest) SetMaxResults(v int32) *QueryConferenceMembersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryConferenceMembersShrinkRequest) SetNextToken(v string) *QueryConferenceMembersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryConferenceMembersShrinkRequest) SetTenantContextShrink(v string) *QueryConferenceMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryConferenceMembersShrinkRequest) SetConferenceId(v string) *QueryConferenceMembersShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
