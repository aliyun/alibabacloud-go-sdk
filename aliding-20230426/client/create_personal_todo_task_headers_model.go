// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTodoTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreatePersonalTodoTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreatePersonalTodoTaskHeadersAccountContext) *CreatePersonalTodoTaskHeaders
	GetAccountContext() *CreatePersonalTodoTaskHeadersAccountContext
}

type CreatePersonalTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreatePersonalTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreatePersonalTodoTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreatePersonalTodoTaskHeaders) GetAccountContext() *CreatePersonalTodoTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *CreatePersonalTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreatePersonalTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreatePersonalTodoTaskHeaders) SetAccountContext(v *CreatePersonalTodoTaskHeadersAccountContext) *CreatePersonalTodoTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *CreatePersonalTodoTaskHeaders) Validate() error {
	return dara.Validate(s)
}

type CreatePersonalTodoTaskHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// ba3a9b612345678d8fedf544ef69d19e
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreatePersonalTodoTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreatePersonalTodoTaskHeadersAccountContext) SetAccountId(v string) *CreatePersonalTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreatePersonalTodoTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
