// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSheetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSheetHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteSheetHeadersAccountContext) *DeleteSheetHeaders
	GetAccountContext() *DeleteSheetHeadersAccountContext
}

type DeleteSheetHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteSheetHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteSheetHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSheetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSheetHeaders) GetAccountContext() *DeleteSheetHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteSheetHeaders) SetCommonHeaders(v map[string]*string) *DeleteSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSheetHeaders) SetAccountContext(v *DeleteSheetHeadersAccountContext) *DeleteSheetHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteSheetHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteSheetHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteSheetHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteSheetHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteSheetHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteSheetHeadersAccountContext) SetAccountId(v string) *DeleteSheetHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteSheetHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
