// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketMemoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddTicketMemoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AddTicketMemoShrinkHeaders
	GetAccountContextShrink() *string
}

type AddTicketMemoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AddTicketMemoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AddTicketMemoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddTicketMemoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AddTicketMemoShrinkHeaders) SetCommonHeaders(v map[string]*string) *AddTicketMemoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddTicketMemoShrinkHeaders) SetAccountContextShrink(v string) *AddTicketMemoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AddTicketMemoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
