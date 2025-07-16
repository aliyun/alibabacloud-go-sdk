// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDentryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CopyDentryHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CopyDentryHeadersAccountContext) *CopyDentryHeaders
	GetAccountContext() *CopyDentryHeadersAccountContext
}

type CopyDentryHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CopyDentryHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CopyDentryHeaders) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryHeaders) GoString() string {
	return s.String()
}

func (s *CopyDentryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CopyDentryHeaders) GetAccountContext() *CopyDentryHeadersAccountContext {
	return s.AccountContext
}

func (s *CopyDentryHeaders) SetCommonHeaders(v map[string]*string) *CopyDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyDentryHeaders) SetAccountContext(v *CopyDentryHeadersAccountContext) *CopyDentryHeaders {
	s.AccountContext = v
	return s
}

func (s *CopyDentryHeaders) Validate() error {
	return dara.Validate(s)
}

type CopyDentryHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CopyDentryHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CopyDentryHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CopyDentryHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CopyDentryHeadersAccountContext) SetAccountId(v string) *CopyDentryHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CopyDentryHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
