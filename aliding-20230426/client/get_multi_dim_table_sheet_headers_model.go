// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableSheetHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableSheetHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMultiDimTableSheetHeadersAccountContext) *GetMultiDimTableSheetHeaders
	GetAccountContext() *GetMultiDimTableSheetHeadersAccountContext
}

type GetMultiDimTableSheetHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMultiDimTableSheetHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableSheetHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableSheetHeaders) GetAccountContext() *GetMultiDimTableSheetHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMultiDimTableSheetHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableSheetHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableSheetHeaders) SetAccountContext(v *GetMultiDimTableSheetHeadersAccountContext) *GetMultiDimTableSheetHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMultiDimTableSheetHeaders) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableSheetHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMultiDimTableSheetHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableSheetHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableSheetHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultiDimTableSheetHeadersAccountContext) SetAccountId(v string) *GetMultiDimTableSheetHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMultiDimTableSheetHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
