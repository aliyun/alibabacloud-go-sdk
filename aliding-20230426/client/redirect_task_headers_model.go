// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedirectTaskHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RedirectTaskHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *RedirectTaskHeadersAccountContext) *RedirectTaskHeaders
	GetAccountContext() *RedirectTaskHeadersAccountContext
}

type RedirectTaskHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *RedirectTaskHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s RedirectTaskHeaders) String() string {
	return dara.Prettify(s)
}

func (s RedirectTaskHeaders) GoString() string {
	return s.String()
}

func (s *RedirectTaskHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RedirectTaskHeaders) GetAccountContext() *RedirectTaskHeadersAccountContext {
	return s.AccountContext
}

func (s *RedirectTaskHeaders) SetCommonHeaders(v map[string]*string) *RedirectTaskHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RedirectTaskHeaders) SetAccountContext(v *RedirectTaskHeadersAccountContext) *RedirectTaskHeaders {
	s.AccountContext = v
	return s
}

func (s *RedirectTaskHeaders) Validate() error {
	return dara.Validate(s)
}

type RedirectTaskHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s RedirectTaskHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s RedirectTaskHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *RedirectTaskHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *RedirectTaskHeadersAccountContext) SetAccountId(v string) *RedirectTaskHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *RedirectTaskHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
