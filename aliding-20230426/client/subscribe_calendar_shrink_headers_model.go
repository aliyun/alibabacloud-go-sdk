// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeCalendarShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubscribeCalendarShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SubscribeCalendarShrinkHeaders
	GetAccountContextShrink() *string
}

type SubscribeCalendarShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SubscribeCalendarShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubscribeCalendarShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubscribeCalendarShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SubscribeCalendarShrinkHeaders) SetCommonHeaders(v map[string]*string) *SubscribeCalendarShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeCalendarShrinkHeaders) SetAccountContextShrink(v string) *SubscribeCalendarShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SubscribeCalendarShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
