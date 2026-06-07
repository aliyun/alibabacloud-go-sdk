// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpServersShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMcpServersShrinkRequest
	GetNextToken() *string
	SetQ(v string) *ListMcpServersShrinkRequest
	GetQ() *string
	SetVisibilityShrink(v string) *ListMcpServersShrinkRequest
	GetVisibilityShrink() *string
}

type ListMcpServersShrinkRequest struct {
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
	// mcp
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// example:
	//
	// -
	VisibilityShrink *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListMcpServersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMcpServersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpServersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpServersShrinkRequest) GetQ() *string {
	return s.Q
}

func (s *ListMcpServersShrinkRequest) GetVisibilityShrink() *string {
	return s.VisibilityShrink
}

func (s *ListMcpServersShrinkRequest) SetMaxResults(v int32) *ListMcpServersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpServersShrinkRequest) SetNextToken(v string) *ListMcpServersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpServersShrinkRequest) SetQ(v string) *ListMcpServersShrinkRequest {
	s.Q = &v
	return s
}

func (s *ListMcpServersShrinkRequest) SetVisibilityShrink(v string) *ListMcpServersShrinkRequest {
	s.VisibilityShrink = &v
	return s
}

func (s *ListMcpServersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
