// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormRemarkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveFormRemarkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SaveFormRemarkHeadersAccountContext) *SaveFormRemarkHeaders
	GetAccountContext() *SaveFormRemarkHeadersAccountContext
}

type SaveFormRemarkHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SaveFormRemarkHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SaveFormRemarkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveFormRemarkHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveFormRemarkHeaders) GetAccountContext() *SaveFormRemarkHeadersAccountContext {
	return s.AccountContext
}

func (s *SaveFormRemarkHeaders) SetCommonHeaders(v map[string]*string) *SaveFormRemarkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormRemarkHeaders) SetAccountContext(v *SaveFormRemarkHeadersAccountContext) *SaveFormRemarkHeaders {
	s.AccountContext = v
	return s
}

func (s *SaveFormRemarkHeaders) Validate() error {
	return dara.Validate(s)
}

type SaveFormRemarkHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SaveFormRemarkHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SaveFormRemarkHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SaveFormRemarkHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SaveFormRemarkHeadersAccountContext) SetAccountId(v string) *SaveFormRemarkHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SaveFormRemarkHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
