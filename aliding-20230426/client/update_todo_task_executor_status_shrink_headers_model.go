// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskExecutorStatusShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateTodoTaskExecutorStatusShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateTodoTaskExecutorStatusShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) SetAccountContextShrink(v string) *UpdateTodoTaskExecutorStatusShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
