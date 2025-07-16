// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateTodoTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateTodoTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateTodoTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateTodoTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *CreateTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
