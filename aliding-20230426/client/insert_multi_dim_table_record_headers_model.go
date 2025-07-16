// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMultiDimTableRecordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertMultiDimTableRecordHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InsertMultiDimTableRecordHeadersAccountContext) *InsertMultiDimTableRecordHeaders
	GetAccountContext() *InsertMultiDimTableRecordHeadersAccountContext
}

type InsertMultiDimTableRecordHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertMultiDimTableRecordHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertMultiDimTableRecordHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordHeaders) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertMultiDimTableRecordHeaders) GetAccountContext() *InsertMultiDimTableRecordHeadersAccountContext {
	return s.AccountContext
}

func (s *InsertMultiDimTableRecordHeaders) SetCommonHeaders(v map[string]*string) *InsertMultiDimTableRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertMultiDimTableRecordHeaders) SetAccountContext(v *InsertMultiDimTableRecordHeadersAccountContext) *InsertMultiDimTableRecordHeaders {
	s.AccountContext = v
	return s
}

func (s *InsertMultiDimTableRecordHeaders) Validate() error {
	return dara.Validate(s)
}

type InsertMultiDimTableRecordHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertMultiDimTableRecordHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InsertMultiDimTableRecordHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertMultiDimTableRecordHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InsertMultiDimTableRecordHeadersAccountContext) SetAccountId(v string) *InsertMultiDimTableRecordHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InsertMultiDimTableRecordHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
