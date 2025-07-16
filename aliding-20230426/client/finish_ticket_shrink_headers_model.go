// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *FinishTicketShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *FinishTicketShrinkHeaders
	GetAccountContextShrink() *string
}

type FinishTicketShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s FinishTicketShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketShrinkHeaders) GoString() string {
	return s.String()
}

func (s *FinishTicketShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *FinishTicketShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *FinishTicketShrinkHeaders) SetCommonHeaders(v map[string]*string) *FinishTicketShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *FinishTicketShrinkHeaders) SetAccountContextShrink(v string) *FinishTicketShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *FinishTicketShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
