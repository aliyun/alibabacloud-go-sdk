// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetUserHeadersAccountContext) *GetUserHeaders
	GetAccountContext() *GetUserHeadersAccountContext
}

type GetUserHeaders struct {
	CommonHeaders  map[string]*string            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserHeaders) GoString() string {
	return s.String()
}

func (s *GetUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserHeaders) GetAccountContext() *GetUserHeadersAccountContext {
	return s.AccountContext
}

func (s *GetUserHeaders) SetCommonHeaders(v map[string]*string) *GetUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserHeaders) SetAccountContext(v *GetUserHeadersAccountContext) *GetUserHeaders {
	s.AccountContext = v
	return s
}

func (s *GetUserHeaders) Validate() error {
	return dara.Validate(s)
}

type GetUserHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserHeadersAccountContext) SetAccountId(v string) *GetUserHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetUserHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
