// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeCalendarShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UnsubscribeCalendarShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UnsubscribeCalendarShrinkHeaders
	GetAccountContextShrink() *string
}

type UnsubscribeCalendarShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UnsubscribeCalendarShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeCalendarShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeCalendarShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UnsubscribeCalendarShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UnsubscribeCalendarShrinkHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeCalendarShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeCalendarShrinkHeaders) SetAccountContextShrink(v string) *UnsubscribeCalendarShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UnsubscribeCalendarShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
