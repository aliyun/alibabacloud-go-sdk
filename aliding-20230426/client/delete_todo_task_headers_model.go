// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTodoTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteTodoTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteTodoTaskHeadersAccountContext) *DeleteTodoTaskHeaders
	GetAccountContext() *DeleteTodoTaskHeadersAccountContext
}

type DeleteTodoTaskHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteTodoTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteTodoTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskHeaders) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteTodoTaskHeaders) GetAccountContext() *DeleteTodoTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteTodoTaskHeaders) SetCommonHeaders(v map[string]*string) *DeleteTodoTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteTodoTaskHeaders) SetAccountContext(v *DeleteTodoTaskHeadersAccountContext) *DeleteTodoTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteTodoTaskHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteTodoTaskHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteTodoTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteTodoTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteTodoTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteTodoTaskHeadersAccountContext) SetAccountId(v string) *DeleteTodoTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteTodoTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
