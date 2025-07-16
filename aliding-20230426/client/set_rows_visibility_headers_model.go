// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRowsVisibilityHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetRowsVisibilityHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SetRowsVisibilityHeadersAccountContext) *SetRowsVisibilityHeaders
	GetAccountContext() *SetRowsVisibilityHeadersAccountContext
}

type SetRowsVisibilityHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SetRowsVisibilityHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SetRowsVisibilityHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetRowsVisibilityHeaders) GetAccountContext() *SetRowsVisibilityHeadersAccountContext {
	return s.AccountContext
}

func (s *SetRowsVisibilityHeaders) SetCommonHeaders(v map[string]*string) *SetRowsVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetRowsVisibilityHeaders) SetAccountContext(v *SetRowsVisibilityHeadersAccountContext) *SetRowsVisibilityHeaders {
	s.AccountContext = v
	return s
}

func (s *SetRowsVisibilityHeaders) Validate() error {
	return dara.Validate(s)
}

type SetRowsVisibilityHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SetRowsVisibilityHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SetRowsVisibilityHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SetRowsVisibilityHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SetRowsVisibilityHeadersAccountContext) SetAccountId(v string) *SetRowsVisibilityHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SetRowsVisibilityHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
