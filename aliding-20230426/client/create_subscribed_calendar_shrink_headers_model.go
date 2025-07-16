// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscribedCalendarShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSubscribedCalendarShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateSubscribedCalendarShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateSubscribedCalendarShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateSubscribedCalendarShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSubscribedCalendarShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateSubscribedCalendarShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateSubscribedCalendarShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSubscribedCalendarShrinkHeaders) SetAccountContextShrink(v string) *CreateSubscribedCalendarShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateSubscribedCalendarShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
