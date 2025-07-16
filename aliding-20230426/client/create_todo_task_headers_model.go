// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateTodoTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateTodoTaskHeadersAccountContext) *CreateTodoTaskHeaders
	GetAccountContext() *CreateTodoTaskHeadersAccountContext
}

type CreateTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateTodoTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateTodoTaskHeaders) GetAccountContext() *CreateTodoTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTodoTaskHeaders) SetAccountContext(v *CreateTodoTaskHeadersAccountContext) *CreateTodoTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateTodoTaskHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateTodoTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateTodoTaskHeadersAccountContext) SetAccountId(v string) *CreateTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateTodoTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
