// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMinutesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMinutesShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMinutesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMinutesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMinutesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMinutesShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesShrinkHeaders) SetAccountContextShrink(v string) *QueryMinutesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMinutesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
