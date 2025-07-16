// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTicketHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetTicketHeadersAccountContext) *GetTicketHeaders
	GetAccountContext() *GetTicketHeadersAccountContext
}

type GetTicketHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetTicketHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetTicketHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTicketHeaders) GoString() string {
	return s.String()
}

func (s *GetTicketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTicketHeaders) GetAccountContext() *GetTicketHeadersAccountContext {
	return s.AccountContext
}

func (s *GetTicketHeaders) SetCommonHeaders(v map[string]*string) *GetTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTicketHeaders) SetAccountContext(v *GetTicketHeadersAccountContext) *GetTicketHeaders {
	s.AccountContext = v
	return s
}

func (s *GetTicketHeaders) Validate() error {
	return dara.Validate(s)
}

type GetTicketHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetTicketHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetTicketHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetTicketHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTicketHeadersAccountContext) SetAccountId(v string) *GetTicketHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetTicketHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
