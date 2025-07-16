// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateTicketShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateTicketShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateTicketShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateTicketShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateTicketShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateTicketShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateTicketShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateTicketShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateTicketShrinkHeaders) SetAccountContextShrink(v string) *CreateTicketShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateTicketShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
