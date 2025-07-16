// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContentHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveContentHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SaveContentHeadersAccountContext) *SaveContentHeaders
	GetAccountContext() *SaveContentHeadersAccountContext
}

type SaveContentHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SaveContentHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SaveContentHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveContentHeaders) GoString() string {
	return s.String()
}

func (s *SaveContentHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveContentHeaders) GetAccountContext() *SaveContentHeadersAccountContext {
	return s.AccountContext
}

func (s *SaveContentHeaders) SetCommonHeaders(v map[string]*string) *SaveContentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveContentHeaders) SetAccountContext(v *SaveContentHeadersAccountContext) *SaveContentHeaders {
	s.AccountContext = v
	return s
}

func (s *SaveContentHeaders) Validate() error {
	return dara.Validate(s)
}

type SaveContentHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SaveContentHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SaveContentHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SaveContentHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SaveContentHeadersAccountContext) SetAccountId(v string) *SaveContentHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SaveContentHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
