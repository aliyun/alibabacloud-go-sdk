// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskExecutorStatusHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateTodoTaskExecutorStatusHeadersAccountContext) *UpdateTodoTaskExecutorStatusHeaders
	GetAccountContext() *UpdateTodoTaskExecutorStatusHeadersAccountContext
}

type UpdateTodoTaskExecutorStatusHeaders struct {
	CommonHeaders  map[string]*string                                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateTodoTaskExecutorStatusHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateTodoTaskExecutorStatusHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateTodoTaskExecutorStatusHeaders) GetAccountContext() *UpdateTodoTaskExecutorStatusHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskExecutorStatusHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusHeaders) SetAccountContext(v *UpdateTodoTaskExecutorStatusHeadersAccountContext) *UpdateTodoTaskExecutorStatusHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTodoTaskExecutorStatusHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateTodoTaskExecutorStatusHeadersAccountContext) SetAccountId(v string) *UpdateTodoTaskExecutorStatusHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
