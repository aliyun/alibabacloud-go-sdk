// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketMemoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddTicketMemoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddTicketMemoHeadersAccountContext) *AddTicketMemoHeaders
	GetAccountContext() *AddTicketMemoHeadersAccountContext
}

type AddTicketMemoHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddTicketMemoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddTicketMemoHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoHeaders) GoString() string {
	return s.String()
}

func (s *AddTicketMemoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddTicketMemoHeaders) GetAccountContext() *AddTicketMemoHeadersAccountContext {
	return s.AccountContext
}

func (s *AddTicketMemoHeaders) SetCommonHeaders(v map[string]*string) *AddTicketMemoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddTicketMemoHeaders) SetAccountContext(v *AddTicketMemoHeadersAccountContext) *AddTicketMemoHeaders {
	s.AccountContext = v
	return s
}

func (s *AddTicketMemoHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddTicketMemoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddTicketMemoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddTicketMemoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddTicketMemoHeadersAccountContext) SetAccountId(v string) *AddTicketMemoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddTicketMemoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
