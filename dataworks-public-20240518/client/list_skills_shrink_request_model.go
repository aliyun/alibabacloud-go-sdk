// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSkillsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillsShrinkRequest
	GetNextToken() *string
	SetQ(v string) *ListSkillsShrinkRequest
	GetQ() *string
	SetVisibilityShrink(v string) *ListSkillsShrinkRequest
	GetVisibilityShrink() *string
}

type ListSkillsShrinkRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 5
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// analysis
	Q                *string `json:"Q,omitempty" xml:"Q,omitempty"`
	VisibilityShrink *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListSkillsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsShrinkRequest) GetQ() *string {
	return s.Q
}

func (s *ListSkillsShrinkRequest) GetVisibilityShrink() *string {
	return s.VisibilityShrink
}

func (s *ListSkillsShrinkRequest) SetMaxResults(v int32) *ListSkillsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetNextToken(v string) *ListSkillsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetQ(v string) *ListSkillsShrinkRequest {
	s.Q = &v
	return s
}

func (s *ListSkillsShrinkRequest) SetVisibilityShrink(v string) *ListSkillsShrinkRequest {
	s.VisibilityShrink = &v
	return s
}

func (s *ListSkillsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
