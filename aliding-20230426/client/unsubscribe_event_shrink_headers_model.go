// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UnsubscribeEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UnsubscribeEventShrinkHeaders
	GetAccountContextShrink() *string
}

type UnsubscribeEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UnsubscribeEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UnsubscribeEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UnsubscribeEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UnsubscribeEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *UnsubscribeEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UnsubscribeEventShrinkHeaders) SetAccountContextShrink(v string) *UnsubscribeEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UnsubscribeEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
