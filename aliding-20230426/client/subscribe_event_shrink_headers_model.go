// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SubscribeEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SubscribeEventShrinkHeaders
	GetAccountContextShrink() *string
}

type SubscribeEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SubscribeEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SubscribeEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SubscribeEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SubscribeEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *SubscribeEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SubscribeEventShrinkHeaders) SetAccountContextShrink(v string) *SubscribeEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SubscribeEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
