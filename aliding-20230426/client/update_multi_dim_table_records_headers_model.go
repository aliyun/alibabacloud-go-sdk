// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableRecordsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableRecordsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateMultiDimTableRecordsHeadersAccountContext) *UpdateMultiDimTableRecordsHeaders
	GetAccountContext() *UpdateMultiDimTableRecordsHeadersAccountContext
}

type UpdateMultiDimTableRecordsHeaders struct {
	CommonHeaders  map[string]*string                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateMultiDimTableRecordsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateMultiDimTableRecordsHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMultiDimTableRecordsHeaders) GetAccountContext() *UpdateMultiDimTableRecordsHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateMultiDimTableRecordsHeaders) SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableRecordsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMultiDimTableRecordsHeaders) SetAccountContext(v *UpdateMultiDimTableRecordsHeadersAccountContext) *UpdateMultiDimTableRecordsHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateMultiDimTableRecordsHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiDimTableRecordsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateMultiDimTableRecordsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableRecordsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableRecordsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateMultiDimTableRecordsHeadersAccountContext) SetAccountId(v string) *UpdateMultiDimTableRecordsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateMultiDimTableRecordsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
