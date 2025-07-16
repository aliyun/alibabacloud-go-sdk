// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateTodoTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateTodoTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateTodoTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateTodoTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *UpdateTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateTodoTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
