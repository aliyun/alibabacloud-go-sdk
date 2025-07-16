// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetEventShrinkHeaders
	GetAccountContextShrink() *string
}

type GetEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetEventShrinkHeaders) SetAccountContextShrink(v string) *GetEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
