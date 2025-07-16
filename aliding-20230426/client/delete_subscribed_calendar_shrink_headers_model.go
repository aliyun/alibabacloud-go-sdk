// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscribedCalendarShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteSubscribedCalendarShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteSubscribedCalendarShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteSubscribedCalendarShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteSubscribedCalendarShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscribedCalendarShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteSubscribedCalendarShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteSubscribedCalendarShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteSubscribedCalendarShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteSubscribedCalendarShrinkHeaders) SetAccountContextShrink(v string) *DeleteSubscribedCalendarShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteSubscribedCalendarShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
