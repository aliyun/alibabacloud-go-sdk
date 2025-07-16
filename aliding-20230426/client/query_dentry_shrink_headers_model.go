// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDentryShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryDentryShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryDentryShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryDentryShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryDentryShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDentryShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryDentryShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryDentryShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryDentryShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryDentryShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryDentryShrinkHeaders) SetAccountContextShrink(v string) *QueryDentryShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryDentryShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
