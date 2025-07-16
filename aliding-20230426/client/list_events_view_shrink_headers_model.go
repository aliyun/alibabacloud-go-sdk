// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsViewShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListEventsViewShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListEventsViewShrinkHeaders
	GetAccountContextShrink() *string
}

type ListEventsViewShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListEventsViewShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsViewShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListEventsViewShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListEventsViewShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListEventsViewShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsViewShrinkHeaders) SetAccountContextShrink(v string) *ListEventsViewShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListEventsViewShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
