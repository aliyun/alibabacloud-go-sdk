// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMultiDimTableHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddMultiDimTableHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddMultiDimTableHeadersAccountContext) *AddMultiDimTableHeaders
	GetAccountContext() *AddMultiDimTableHeadersAccountContext
}

type AddMultiDimTableHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddMultiDimTableHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddMultiDimTableHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableHeaders) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddMultiDimTableHeaders) GetAccountContext() *AddMultiDimTableHeadersAccountContext {
	return s.AccountContext
}

func (s *AddMultiDimTableHeaders) SetCommonHeaders(v map[string]*string) *AddMultiDimTableHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddMultiDimTableHeaders) SetAccountContext(v *AddMultiDimTableHeadersAccountContext) *AddMultiDimTableHeaders {
	s.AccountContext = v
	return s
}

func (s *AddMultiDimTableHeaders) Validate() error {
	return dara.Validate(s)
}

type AddMultiDimTableHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddMultiDimTableHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddMultiDimTableHeadersAccountContext) SetAccountId(v string) *AddMultiDimTableHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddMultiDimTableHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
