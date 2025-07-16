// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDingtalkPersonalTodoTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDingtalkPersonalTodoTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateDingtalkPersonalTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDingtalkPersonalTodoTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateDingtalkPersonalTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateDingtalkPersonalTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *CreateDingtalkPersonalTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
