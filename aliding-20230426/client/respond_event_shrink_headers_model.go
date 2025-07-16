// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondEventShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RespondEventShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *RespondEventShrinkHeaders
	GetAccountContextShrink() *string
}

type RespondEventShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s RespondEventShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s RespondEventShrinkHeaders) GoString() string {
	return s.String()
}

func (s *RespondEventShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RespondEventShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *RespondEventShrinkHeaders) SetCommonHeaders(v map[string]*string) *RespondEventShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RespondEventShrinkHeaders) SetAccountContextShrink(v string) *RespondEventShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *RespondEventShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
