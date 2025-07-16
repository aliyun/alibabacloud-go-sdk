// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTodoTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetTodoTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type GetTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetTodoTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTodoTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *GetTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetTodoTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
