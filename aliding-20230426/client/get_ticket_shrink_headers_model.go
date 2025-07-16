// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTicketShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetTicketShrinkHeaders
	GetAccountContextShrink() *string
}

type GetTicketShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetTicketShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTicketShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetTicketShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTicketShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetTicketShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetTicketShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTicketShrinkHeaders) SetAccountContextShrink(v string) *GetTicketShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetTicketShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
