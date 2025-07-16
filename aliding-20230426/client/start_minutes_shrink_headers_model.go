// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMinutesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartMinutesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StartMinutesShrinkHeaders
	GetAccountContextShrink() *string
}

type StartMinutesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StartMinutesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StartMinutesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartMinutesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StartMinutesShrinkHeaders) SetCommonHeaders(v map[string]*string) *StartMinutesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartMinutesShrinkHeaders) SetAccountContextShrink(v string) *StartMinutesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StartMinutesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
