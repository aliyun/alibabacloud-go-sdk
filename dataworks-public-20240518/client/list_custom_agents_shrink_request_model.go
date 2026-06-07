// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCustomAgentsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCustomAgentsShrinkRequest
	GetNextToken() *string
	SetQ(v string) *ListCustomAgentsShrinkRequest
	GetQ() *string
	SetVisibilityShrink(v string) *ListCustomAgentsShrinkRequest
	GetVisibilityShrink() *string
}

type ListCustomAgentsShrinkRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// analysis
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// example:
	//
	// -
	VisibilityShrink *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListCustomAgentsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomAgentsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomAgentsShrinkRequest) GetQ() *string {
	return s.Q
}

func (s *ListCustomAgentsShrinkRequest) GetVisibilityShrink() *string {
	return s.VisibilityShrink
}

func (s *ListCustomAgentsShrinkRequest) SetMaxResults(v int32) *ListCustomAgentsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomAgentsShrinkRequest) SetNextToken(v string) *ListCustomAgentsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomAgentsShrinkRequest) SetQ(v string) *ListCustomAgentsShrinkRequest {
	s.Q = &v
	return s
}

func (s *ListCustomAgentsShrinkRequest) SetVisibilityShrink(v string) *ListCustomAgentsShrinkRequest {
	s.VisibilityShrink = &v
	return s
}

func (s *ListCustomAgentsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
