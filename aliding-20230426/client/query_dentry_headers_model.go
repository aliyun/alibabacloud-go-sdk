// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryDentryHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryDentryHeadersAccountContext) *QueryDentryHeaders
	GetAccountContext() *QueryDentryHeadersAccountContext
}

type QueryDentryHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryDentryHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryDentryHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryHeaders) GoString() string {
	return s.String()
}

func (s *QueryDentryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryDentryHeaders) GetAccountContext() *QueryDentryHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryDentryHeaders) SetCommonHeaders(v map[string]*string) *QueryDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDentryHeaders) SetAccountContext(v *QueryDentryHeadersAccountContext) *QueryDentryHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryDentryHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryDentryHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryDentryHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryDentryHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryDentryHeadersAccountContext) SetAccountId(v string) *QueryDentryHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryDentryHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
