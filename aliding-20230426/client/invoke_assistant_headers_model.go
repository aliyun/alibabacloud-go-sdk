// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAssistantHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InvokeAssistantHeaders
	GetCommonHeaders() map[string]*string
	SetAccountId(v string) *InvokeAssistantHeaders
	GetAccountId() *string
	SetAlidingSsoTicket(v string) *InvokeAssistantHeaders
	GetAlidingSsoTicket() *string
}

type InvokeAssistantHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// example:
	//
	// 123456
	AccountId        *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	AlidingSsoTicket *string `json:"alidingSsoTicket,omitempty" xml:"alidingSsoTicket,omitempty"`
}

func (s InvokeAssistantHeaders) String() string {
	return dara.Prettify(s)
}

func (s InvokeAssistantHeaders) GoString() string {
	return s.String()
}

func (s *InvokeAssistantHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InvokeAssistantHeaders) GetAccountId() *string {
	return s.AccountId
}

func (s *InvokeAssistantHeaders) GetAlidingSsoTicket() *string {
	return s.AlidingSsoTicket
}

func (s *InvokeAssistantHeaders) SetCommonHeaders(v map[string]*string) *InvokeAssistantHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InvokeAssistantHeaders) SetAccountId(v string) *InvokeAssistantHeaders {
	s.AccountId = &v
	return s
}

func (s *InvokeAssistantHeaders) SetAlidingSsoTicket(v string) *InvokeAssistantHeaders {
	s.AlidingSsoTicket = &v
	return s
}

func (s *InvokeAssistantHeaders) Validate() error {
	return dara.Validate(s)
}
