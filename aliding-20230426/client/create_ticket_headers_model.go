// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateTicketHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateTicketHeadersAccountContext) *CreateTicketHeaders
	GetAccountContext() *CreateTicketHeadersAccountContext
}

type CreateTicketHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateTicketHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateTicketHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketHeaders) GoString() string {
	return s.String()
}

func (s *CreateTicketHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateTicketHeaders) GetAccountContext() *CreateTicketHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateTicketHeaders) SetCommonHeaders(v map[string]*string) *CreateTicketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTicketHeaders) SetAccountContext(v *CreateTicketHeadersAccountContext) *CreateTicketHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateTicketHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTicketHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateTicketHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateTicketHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateTicketHeadersAccountContext) SetAccountId(v string) *CreateTicketHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateTicketHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
