// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMinutesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopMinutesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StopMinutesShrinkHeaders
	GetAccountContextShrink() *string
}

type StopMinutesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StopMinutesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StopMinutesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopMinutesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StopMinutesShrinkHeaders) SetCommonHeaders(v map[string]*string) *StopMinutesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopMinutesShrinkHeaders) SetAccountContextShrink(v string) *StopMinutesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StopMinutesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
