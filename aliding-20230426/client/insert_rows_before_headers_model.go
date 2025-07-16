// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRowsBeforeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertRowsBeforeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InsertRowsBeforeHeadersAccountContext) *InsertRowsBeforeHeaders
	GetAccountContext() *InsertRowsBeforeHeadersAccountContext
}

type InsertRowsBeforeHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertRowsBeforeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertRowsBeforeHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeHeaders) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertRowsBeforeHeaders) GetAccountContext() *InsertRowsBeforeHeadersAccountContext {
	return s.AccountContext
}

func (s *InsertRowsBeforeHeaders) SetCommonHeaders(v map[string]*string) *InsertRowsBeforeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertRowsBeforeHeaders) SetAccountContext(v *InsertRowsBeforeHeadersAccountContext) *InsertRowsBeforeHeaders {
	s.AccountContext = v
	return s
}

func (s *InsertRowsBeforeHeaders) Validate() error {
	return dara.Validate(s)
}

type InsertRowsBeforeHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertRowsBeforeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InsertRowsBeforeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertRowsBeforeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InsertRowsBeforeHeadersAccountContext) SetAccountId(v string) *InsertRowsBeforeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InsertRowsBeforeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
