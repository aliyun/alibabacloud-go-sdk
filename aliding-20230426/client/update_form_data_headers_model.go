// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateFormDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateFormDataHeadersAccountContext) *UpdateFormDataHeaders
	GetAccountContext() *UpdateFormDataHeadersAccountContext
}

type UpdateFormDataHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateFormDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateFormDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormDataHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFormDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateFormDataHeaders) GetAccountContext() *UpdateFormDataHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateFormDataHeaders) SetCommonHeaders(v map[string]*string) *UpdateFormDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFormDataHeaders) SetAccountContext(v *UpdateFormDataHeadersAccountContext) *UpdateFormDataHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateFormDataHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateFormDataHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateFormDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateFormDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateFormDataHeadersAccountContext) SetAccountId(v string) *UpdateFormDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateFormDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
