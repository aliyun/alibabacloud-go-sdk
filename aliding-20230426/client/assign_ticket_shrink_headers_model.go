// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignTicketShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AssignTicketShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *AssignTicketShrinkHeaders
	GetAccountContextShrink() *string
}

type AssignTicketShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s AssignTicketShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketShrinkHeaders) GoString() string {
	return s.String()
}

func (s *AssignTicketShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AssignTicketShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *AssignTicketShrinkHeaders) SetCommonHeaders(v map[string]*string) *AssignTicketShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AssignTicketShrinkHeaders) SetAccountContextShrink(v string) *AssignTicketShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *AssignTicketShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
