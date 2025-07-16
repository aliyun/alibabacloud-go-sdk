// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchKeywordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSearchKeywordHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateSearchKeywordHeadersAccountContext) *CreateSearchKeywordHeaders
	GetAccountContext() *CreateSearchKeywordHeadersAccountContext
}

type CreateSearchKeywordHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateSearchKeywordHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateSearchKeywordHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordHeaders) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSearchKeywordHeaders) GetAccountContext() *CreateSearchKeywordHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateSearchKeywordHeaders) SetCommonHeaders(v map[string]*string) *CreateSearchKeywordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSearchKeywordHeaders) SetAccountContext(v *CreateSearchKeywordHeadersAccountContext) *CreateSearchKeywordHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateSearchKeywordHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateSearchKeywordHeadersAccountContext struct {
	// example:
	//
	// ba3a9b612345678d8fedf544ef69d19e
	UserToken *string `json:"userToken,omitempty" xml:"userToken,omitempty"`
}

func (s CreateSearchKeywordHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordHeadersAccountContext) GetUserToken() *string {
	return s.UserToken
}

func (s *CreateSearchKeywordHeadersAccountContext) SetUserToken(v string) *CreateSearchKeywordHeadersAccountContext {
	s.UserToken = &v
	return s
}

func (s *CreateSearchKeywordHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
