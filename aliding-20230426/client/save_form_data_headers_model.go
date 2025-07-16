// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveFormDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SaveFormDataHeadersAccountContext) *SaveFormDataHeaders
	GetAccountContext() *SaveFormDataHeadersAccountContext
}

type SaveFormDataHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SaveFormDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SaveFormDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveFormDataHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveFormDataHeaders) GetAccountContext() *SaveFormDataHeadersAccountContext {
	return s.AccountContext
}

func (s *SaveFormDataHeaders) SetCommonHeaders(v map[string]*string) *SaveFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormDataHeaders) SetAccountContext(v *SaveFormDataHeadersAccountContext) *SaveFormDataHeaders {
	s.AccountContext = v
	return s
}

func (s *SaveFormDataHeaders) Validate() error {
	return dara.Validate(s)
}

type SaveFormDataHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SaveFormDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SaveFormDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SaveFormDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SaveFormDataHeadersAccountContext) SetAccountId(v string) *SaveFormDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SaveFormDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
