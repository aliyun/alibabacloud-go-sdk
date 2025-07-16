// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiDimTableRecordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMultiDimTableRecordHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMultiDimTableRecordHeadersAccountContext) *GetMultiDimTableRecordHeaders
	GetAccountContext() *GetMultiDimTableRecordHeadersAccountContext
}

type GetMultiDimTableRecordHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMultiDimTableRecordHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMultiDimTableRecordHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordHeaders) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMultiDimTableRecordHeaders) GetAccountContext() *GetMultiDimTableRecordHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMultiDimTableRecordHeaders) SetCommonHeaders(v map[string]*string) *GetMultiDimTableRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMultiDimTableRecordHeaders) SetAccountContext(v *GetMultiDimTableRecordHeadersAccountContext) *GetMultiDimTableRecordHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMultiDimTableRecordHeaders) Validate() error {
	return dara.Validate(s)
}

type GetMultiDimTableRecordHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMultiDimTableRecordHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMultiDimTableRecordHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMultiDimTableRecordHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultiDimTableRecordHeadersAccountContext) SetAccountId(v string) *GetMultiDimTableRecordHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMultiDimTableRecordHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
