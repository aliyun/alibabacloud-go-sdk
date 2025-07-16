// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgTodoTasksShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryOrgTodoTasksShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryOrgTodoTasksShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryOrgTodoTasksShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryOrgTodoTasksShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryOrgTodoTasksShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgTodoTasksShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgTodoTasksShrinkHeaders) SetAccountContextShrink(v string) *QueryOrgTodoTasksShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryOrgTodoTasksShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
