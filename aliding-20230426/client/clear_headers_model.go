// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ClearHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ClearHeadersAccountContext) *ClearHeaders
	GetAccountContext() *ClearHeadersAccountContext
}

type ClearHeaders struct {
	CommonHeaders  map[string]*string          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ClearHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ClearHeaders) String() string {
	return dara.Prettify(s)
}

func (s ClearHeaders) GoString() string {
	return s.String()
}

func (s *ClearHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ClearHeaders) GetAccountContext() *ClearHeadersAccountContext {
	return s.AccountContext
}

func (s *ClearHeaders) SetCommonHeaders(v map[string]*string) *ClearHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ClearHeaders) SetAccountContext(v *ClearHeadersAccountContext) *ClearHeaders {
	s.AccountContext = v
	return s
}

func (s *ClearHeaders) Validate() error {
	return dara.Validate(s)
}

type ClearHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ClearHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ClearHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ClearHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ClearHeadersAccountContext) SetAccountId(v string) *ClearHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ClearHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
