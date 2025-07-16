// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscribedCalendarShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSubscribedCalendarShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetSubscribedCalendarShrinkHeaders
	GetAccountContextShrink() *string
}

type GetSubscribedCalendarShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetSubscribedCalendarShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSubscribedCalendarShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetSubscribedCalendarShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetSubscribedCalendarShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSubscribedCalendarShrinkHeaders) SetAccountContextShrink(v string) *GetSubscribedCalendarShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetSubscribedCalendarShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
