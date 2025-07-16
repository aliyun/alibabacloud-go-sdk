// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgTodoTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryOrgTodoTasksShrinkRequest
	GetTenantContextShrink() *string
	SetIsDone(v bool) *QueryOrgTodoTasksShrinkRequest
	GetIsDone() *bool
	SetNextToken(v string) *QueryOrgTodoTasksShrinkRequest
	GetNextToken() *string
}

type QueryOrgTodoTasksShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// true
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgTodoTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryOrgTodoTasksShrinkRequest) GetIsDone() *bool {
	return s.IsDone
}

func (s *QueryOrgTodoTasksShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryOrgTodoTasksShrinkRequest) SetTenantContextShrink(v string) *QueryOrgTodoTasksShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryOrgTodoTasksShrinkRequest) SetIsDone(v bool) *QueryOrgTodoTasksShrinkRequest {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksShrinkRequest) SetNextToken(v string) *QueryOrgTodoTasksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
