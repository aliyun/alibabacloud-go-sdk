// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubscribedCalendarsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateSubscribedCalendarsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateSubscribedCalendarsShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateSubscribedCalendarsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateSubscribedCalendarsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubscribedCalendarsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateSubscribedCalendarsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateSubscribedCalendarsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateSubscribedCalendarsShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateSubscribedCalendarsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkHeaders) SetAccountContextShrink(v string) *UpdateSubscribedCalendarsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateSubscribedCalendarsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
