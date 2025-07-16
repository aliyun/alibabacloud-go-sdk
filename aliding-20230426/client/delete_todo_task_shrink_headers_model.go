// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTodoTaskShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteTodoTaskShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteTodoTaskShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteTodoTaskShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteTodoTaskShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteTodoTaskShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteTodoTaskShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoTaskShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoTaskShrinkHeaders) SetAccountContextShrink(v string) *DeleteTodoTaskShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteTodoTaskShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
