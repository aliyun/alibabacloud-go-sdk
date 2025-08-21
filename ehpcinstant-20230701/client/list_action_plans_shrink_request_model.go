// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlansShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionPlanIdsShrink(v string) *ListActionPlansShrinkRequest
	GetActionPlanIdsShrink() *string
	SetMaxResults(v int32) *ListActionPlansShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListActionPlansShrinkRequest
	GetNextToken() *string
}

type ListActionPlansShrinkRequest struct {
	ActionPlanIdsShrink *string `json:"ActionPlanIds,omitempty" xml:"ActionPlanIds,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1d2db86scXXXXXXXXXX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListActionPlansShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlansShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListActionPlansShrinkRequest) GetActionPlanIdsShrink() *string {
	return s.ActionPlanIdsShrink
}

func (s *ListActionPlansShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListActionPlansShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListActionPlansShrinkRequest) SetActionPlanIdsShrink(v string) *ListActionPlansShrinkRequest {
	s.ActionPlanIdsShrink = &v
	return s
}

func (s *ListActionPlansShrinkRequest) SetMaxResults(v int32) *ListActionPlansShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionPlansShrinkRequest) SetNextToken(v string) *ListActionPlansShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionPlansShrinkRequest) Validate() error {
	return dara.Validate(s)
}
