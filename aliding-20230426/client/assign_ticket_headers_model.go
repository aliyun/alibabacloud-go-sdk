// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignTicketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AssignTicketHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AssignTicketHeadersAccountContext) *AssignTicketHeaders
	GetAccountContext() *AssignTicketHeadersAccountContext
}

type AssignTicketHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AssignTicketHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AssignTicketHeaders) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketHeaders) GoString() string {
	return s.String()
}

func (s *AssignTicketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AssignTicketHeaders) GetAccountContext() *AssignTicketHeadersAccountContext {
	return s.AccountContext
}

func (s *AssignTicketHeaders) SetCommonHeaders(v map[string]*string) *AssignTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignTicketHeaders) SetAccountContext(v *AssignTicketHeadersAccountContext) *AssignTicketHeaders {
	s.AccountContext = v
	return s
}

func (s *AssignTicketHeaders) Validate() error {
	return dara.Validate(s)
}

type AssignTicketHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AssignTicketHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AssignTicketHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AssignTicketHeadersAccountContext) SetAccountId(v string) *AssignTicketHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AssignTicketHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
