// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeContainerHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeContainerHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InvokeContainerHeadersAccountContext) *InvokeContainerHeaders
	GetAccountContext() *InvokeContainerHeadersAccountContext
}

type InvokeContainerHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InvokeContainerHeadersAccountContext `json:"accountContext,omitempty" xml:"accountContext,omitempty" type:"Struct"`
}

func (s InvokeContainerHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeContainerHeaders) GoString() string {
	return s.String()
}

func (s *InvokeContainerHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeContainerHeaders) GetAccountContext() *InvokeContainerHeadersAccountContext {
	return s.AccountContext
}

func (s *InvokeContainerHeaders) SetCommonHeaders(v map[string]*string) *InvokeContainerHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeContainerHeaders) SetAccountContext(v *InvokeContainerHeadersAccountContext) *InvokeContainerHeaders {
	s.AccountContext = v
	return s
}

func (s *InvokeContainerHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeContainerHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AlidingSsoTicket *string `json:"alidingSsoTicket,omitempty" xml:"alidingSsoTicket,omitempty"`
	SsoTicket        *string `json:"ssoTicket,omitempty" xml:"ssoTicket,omitempty"`
}

func (s InvokeContainerHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InvokeContainerHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InvokeContainerHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InvokeContainerHeadersAccountContext) GetAlidingSsoTicket() *string {
	return s.AlidingSsoTicket
}

func (s *InvokeContainerHeadersAccountContext) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *InvokeContainerHeadersAccountContext) SetAccountId(v string) *InvokeContainerHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InvokeContainerHeadersAccountContext) SetAlidingSsoTicket(v string) *InvokeContainerHeadersAccountContext {
	s.AlidingSsoTicket = &v
	return s
}

func (s *InvokeContainerHeadersAccountContext) SetSsoTicket(v string) *InvokeContainerHeadersAccountContext {
	s.SsoTicket = &v
	return s
}

func (s *InvokeContainerHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
