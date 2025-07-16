// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDocContentHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDocContentHeadersAccountContext) *GetDocContentHeaders
	GetAccountContext() *GetDocContentHeadersAccountContext
}

type GetDocContentHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDocContentHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDocContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentHeaders) GoString() string {
	return s.String()
}

func (s *GetDocContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDocContentHeaders) GetAccountContext() *GetDocContentHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDocContentHeaders) SetCommonHeaders(v map[string]*string) *GetDocContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDocContentHeaders) SetAccountContext(v *GetDocContentHeadersAccountContext) *GetDocContentHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDocContentHeaders) Validate() error {
	return dara.Validate(s)
}

type GetDocContentHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// ba3a9b612345678d8fedf544ef69d19e
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDocContentHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDocContentHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDocContentHeadersAccountContext) SetAccountId(v string) *GetDocContentHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDocContentHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
