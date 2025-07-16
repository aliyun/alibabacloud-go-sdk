// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiDimTableFieldHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableFieldHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateMultiDimTableFieldHeadersAccountContext) *UpdateMultiDimTableFieldHeaders
	GetAccountContext() *UpdateMultiDimTableFieldHeadersAccountContext
}

type UpdateMultiDimTableFieldHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateMultiDimTableFieldHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateMultiDimTableFieldHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMultiDimTableFieldHeaders) GetAccountContext() *UpdateMultiDimTableFieldHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateMultiDimTableFieldHeaders) SetCommonHeaders(v map[string]*string) *UpdateMultiDimTableFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMultiDimTableFieldHeaders) SetAccountContext(v *UpdateMultiDimTableFieldHeadersAccountContext) *UpdateMultiDimTableFieldHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateMultiDimTableFieldHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiDimTableFieldHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateMultiDimTableFieldHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiDimTableFieldHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateMultiDimTableFieldHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateMultiDimTableFieldHeadersAccountContext) SetAccountId(v string) *UpdateMultiDimTableFieldHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateMultiDimTableFieldHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
