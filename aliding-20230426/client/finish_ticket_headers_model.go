// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FinishTicketHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *FinishTicketHeadersAccountContext) *FinishTicketHeaders
	GetAccountContext() *FinishTicketHeadersAccountContext
}

type FinishTicketHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *FinishTicketHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s FinishTicketHeaders) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketHeaders) GoString() string {
	return s.String()
}

func (s *FinishTicketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FinishTicketHeaders) GetAccountContext() *FinishTicketHeadersAccountContext {
	return s.AccountContext
}

func (s *FinishTicketHeaders) SetCommonHeaders(v map[string]*string) *FinishTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishTicketHeaders) SetAccountContext(v *FinishTicketHeadersAccountContext) *FinishTicketHeaders {
	s.AccountContext = v
	return s
}

func (s *FinishTicketHeaders) Validate() error {
	return dara.Validate(s)
}

type FinishTicketHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s FinishTicketHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *FinishTicketHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *FinishTicketHeadersAccountContext) SetAccountId(v string) *FinishTicketHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *FinishTicketHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
