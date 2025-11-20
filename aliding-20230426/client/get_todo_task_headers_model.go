// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTodoTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetTodoTaskHeadersAccountContext) *GetTodoTaskHeaders
	GetAccountContext() *GetTodoTaskHeadersAccountContext
}

type GetTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetTodoTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *GetTodoTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTodoTaskHeaders) GetAccountContext() *GetTodoTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *GetTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *GetTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTodoTaskHeaders) SetAccountContext(v *GetTodoTaskHeadersAccountContext) *GetTodoTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *GetTodoTaskHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTodoTaskHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetTodoTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetTodoTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTodoTaskHeadersAccountContext) SetAccountId(v string) *GetTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetTodoTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
