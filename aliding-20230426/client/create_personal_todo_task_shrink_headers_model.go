// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTodoTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreatePersonalTodoTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreatePersonalTodoTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type CreatePersonalTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreatePersonalTodoTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreatePersonalTodoTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreatePersonalTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreatePersonalTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePersonalTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *CreatePersonalTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreatePersonalTodoTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
