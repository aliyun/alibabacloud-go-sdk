// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateTodoTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateTodoTaskHeadersAccountContext) *UpdateTodoTaskHeaders
	GetAccountContext() *UpdateTodoTaskHeadersAccountContext
}

type UpdateTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateTodoTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateTodoTaskHeaders) GetAccountContext() *UpdateTodoTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *UpdateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateTodoTaskHeaders) SetAccountContext(v *UpdateTodoTaskHeadersAccountContext) *UpdateTodoTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateTodoTaskHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTodoTaskHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateTodoTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateTodoTaskHeadersAccountContext) SetAccountId(v string) *UpdateTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateTodoTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
