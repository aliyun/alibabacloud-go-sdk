// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *TransferTicketShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *TransferTicketShrinkHeaders
	GetAccountContextShrink() *string
}

type TransferTicketShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s TransferTicketShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketShrinkHeaders) GoString() string {
	return s.String()
}

func (s *TransferTicketShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *TransferTicketShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *TransferTicketShrinkHeaders) SetCommonHeaders(v map[string]*string) *TransferTicketShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *TransferTicketShrinkHeaders) SetAccountContextShrink(v string) *TransferTicketShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *TransferTicketShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
