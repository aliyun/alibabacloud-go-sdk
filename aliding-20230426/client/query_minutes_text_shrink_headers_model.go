// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesTextShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMinutesTextShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryMinutesTextShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryMinutesTextShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryMinutesTextShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesTextShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesTextShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMinutesTextShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryMinutesTextShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesTextShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesTextShrinkHeaders) SetAccountContextShrink(v string) *QueryMinutesTextShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryMinutesTextShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
