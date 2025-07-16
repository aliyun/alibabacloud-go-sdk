// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSheetContentJobIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSheetContentJobIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetSheetContentJobIdHeadersAccountContext) *GetSheetContentJobIdHeaders
	GetAccountContext() *GetSheetContentJobIdHeadersAccountContext
}

type GetSheetContentJobIdHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetSheetContentJobIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetSheetContentJobIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdHeaders) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSheetContentJobIdHeaders) GetAccountContext() *GetSheetContentJobIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetSheetContentJobIdHeaders) SetCommonHeaders(v map[string]*string) *GetSheetContentJobIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSheetContentJobIdHeaders) SetAccountContext(v *GetSheetContentJobIdHeadersAccountContext) *GetSheetContentJobIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetSheetContentJobIdHeaders) Validate() error {
	return dara.Validate(s)
}

type GetSheetContentJobIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetSheetContentJobIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetSheetContentJobIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetSheetContentJobIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSheetContentJobIdHeadersAccountContext) SetAccountId(v string) *GetSheetContentJobIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetSheetContentJobIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
