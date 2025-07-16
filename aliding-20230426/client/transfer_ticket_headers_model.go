// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TransferTicketHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *TransferTicketHeadersAccountContext) *TransferTicketHeaders
	GetAccountContext() *TransferTicketHeadersAccountContext
}

type TransferTicketHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *TransferTicketHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s TransferTicketHeaders) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketHeaders) GoString() string {
	return s.String()
}

func (s *TransferTicketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TransferTicketHeaders) GetAccountContext() *TransferTicketHeadersAccountContext {
	return s.AccountContext
}

func (s *TransferTicketHeaders) SetCommonHeaders(v map[string]*string) *TransferTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransferTicketHeaders) SetAccountContext(v *TransferTicketHeadersAccountContext) *TransferTicketHeaders {
	s.AccountContext = v
	return s
}

func (s *TransferTicketHeaders) Validate() error {
	return dara.Validate(s)
}

type TransferTicketHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s TransferTicketHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *TransferTicketHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *TransferTicketHeadersAccountContext) SetAccountId(v string) *TransferTicketHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *TransferTicketHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
