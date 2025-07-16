// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateMultiDimTableHeadersAccountContext) *UpdateMultiDimTableHeaders
	GetAccountContext() *UpdateMultiDimTableHeadersAccountContext
}

type UpdateMultiDimTableHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateMultiDimTableHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateMultiDimTableHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMultiDimTableHeaders) GetAccountContext() *UpdateMultiDimTableHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateMultiDimTableHeaders) SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMultiDimTableHeaders) SetAccountContext(v *UpdateMultiDimTableHeadersAccountContext) *UpdateMultiDimTableHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateMultiDimTableHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiDimTableHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateMultiDimTableHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateMultiDimTableHeadersAccountContext) SetAccountId(v string) *UpdateMultiDimTableHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateMultiDimTableHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
