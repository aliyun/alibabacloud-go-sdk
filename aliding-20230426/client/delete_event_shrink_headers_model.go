// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteEventShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteEventShrinkHeaders) SetAccountContextShrink(v string) *DeleteEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
