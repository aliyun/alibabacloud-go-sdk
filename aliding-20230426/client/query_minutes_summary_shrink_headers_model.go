// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesSummaryShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMinutesSummaryShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMinutesSummaryShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMinutesSummaryShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMinutesSummaryShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMinutesSummaryShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMinutesSummaryShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesSummaryShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesSummaryShrinkHeaders) SetAccountContextShrink(v string) *QueryMinutesSummaryShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMinutesSummaryShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
