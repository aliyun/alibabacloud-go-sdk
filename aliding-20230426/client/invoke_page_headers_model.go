// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokePageHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokePageHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InvokePageHeadersAccountContext) *InvokePageHeaders
	GetAccountContext() *InvokePageHeadersAccountContext
}

type InvokePageHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InvokePageHeadersAccountContext `json:"accountContext,omitempty" xml:"accountContext,omitempty" type:"Struct"`
}

func (s InvokePageHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokePageHeaders) GoString() string {
	return s.String()
}

func (s *InvokePageHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokePageHeaders) GetAccountContext() *InvokePageHeadersAccountContext {
	return s.AccountContext
}

func (s *InvokePageHeaders) SetCommonHeaders(v map[string]*string) *InvokePageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokePageHeaders) SetAccountContext(v *InvokePageHeadersAccountContext) *InvokePageHeaders {
	s.AccountContext = v
	return s
}

func (s *InvokePageHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokePageHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AlidingSsoTicket *string `json:"alidingSsoTicket,omitempty" xml:"alidingSsoTicket,omitempty"`
	SsoTicket        *string `json:"ssoTicket,omitempty" xml:"ssoTicket,omitempty"`
}

func (s InvokePageHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InvokePageHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InvokePageHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InvokePageHeadersAccountContext) GetAlidingSsoTicket() *string {
	return s.AlidingSsoTicket
}

func (s *InvokePageHeadersAccountContext) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *InvokePageHeadersAccountContext) SetAccountId(v string) *InvokePageHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InvokePageHeadersAccountContext) SetAlidingSsoTicket(v string) *InvokePageHeadersAccountContext {
	s.AlidingSsoTicket = &v
	return s
}

func (s *InvokePageHeadersAccountContext) SetSsoTicket(v string) *InvokePageHeadersAccountContext {
	s.SsoTicket = &v
	return s
}

func (s *InvokePageHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
