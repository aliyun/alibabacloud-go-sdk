// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeSkillHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeSkillHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InvokeSkillHeadersAccountContext) *InvokeSkillHeaders
	GetAccountContext() *InvokeSkillHeadersAccountContext
}

type InvokeSkillHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InvokeSkillHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InvokeSkillHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillHeaders) GoString() string {
	return s.String()
}

func (s *InvokeSkillHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeSkillHeaders) GetAccountContext() *InvokeSkillHeadersAccountContext {
	return s.AccountContext
}

func (s *InvokeSkillHeaders) SetCommonHeaders(v map[string]*string) *InvokeSkillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeSkillHeaders) SetAccountContext(v *InvokeSkillHeadersAccountContext) *InvokeSkillHeaders {
	s.AccountContext = v
	return s
}

func (s *InvokeSkillHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeSkillHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	SsoTicket *string `json:"ssoTicket,omitempty" xml:"ssoTicket,omitempty"`
}

func (s InvokeSkillHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InvokeSkillHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InvokeSkillHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InvokeSkillHeadersAccountContext) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *InvokeSkillHeadersAccountContext) SetAccountId(v string) *InvokeSkillHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InvokeSkillHeadersAccountContext) SetSsoTicket(v string) *InvokeSkillHeadersAccountContext {
	s.SsoTicket = &v
	return s
}

func (s *InvokeSkillHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
