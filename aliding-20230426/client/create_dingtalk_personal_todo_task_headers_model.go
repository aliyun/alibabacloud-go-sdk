// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDingtalkPersonalTodoTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateDingtalkPersonalTodoTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateDingtalkPersonalTodoTaskHeadersAccountContext) *CreateDingtalkPersonalTodoTaskHeaders
	GetAccountContext() *CreateDingtalkPersonalTodoTaskHeadersAccountContext
}

type CreateDingtalkPersonalTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateDingtalkPersonalTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateDingtalkPersonalTodoTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateDingtalkPersonalTodoTaskHeaders) GetAccountContext() *CreateDingtalkPersonalTodoTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateDingtalkPersonalTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *CreateDingtalkPersonalTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskHeaders) SetAccountContext(v *CreateDingtalkPersonalTodoTaskHeadersAccountContext) *CreateDingtalkPersonalTodoTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateDingtalkPersonalTodoTaskHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateDingtalkPersonalTodoTaskHeadersAccountContext) SetAccountId(v string) *CreateDingtalkPersonalTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
