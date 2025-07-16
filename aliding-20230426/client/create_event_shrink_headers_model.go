// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateEventShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEventShrinkHeaders) SetAccountContextShrink(v string) *CreateEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
