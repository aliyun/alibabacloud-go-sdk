// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDentryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDentryHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteDentryHeadersAccountContext) *DeleteDentryHeaders
	GetAccountContext() *DeleteDentryHeadersAccountContext
}

type DeleteDentryHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteDentryHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteDentryHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDentryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDentryHeaders) GetAccountContext() *DeleteDentryHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteDentryHeaders) SetCommonHeaders(v map[string]*string) *DeleteDentryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDentryHeaders) SetAccountContext(v *DeleteDentryHeadersAccountContext) *DeleteDentryHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteDentryHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteDentryHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteDentryHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteDentryHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteDentryHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteDentryHeadersAccountContext) SetAccountId(v string) *DeleteDentryHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteDentryHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
