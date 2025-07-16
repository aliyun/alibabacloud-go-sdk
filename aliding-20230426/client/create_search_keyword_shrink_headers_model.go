// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchKeywordShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSearchKeywordShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateSearchKeywordShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateSearchKeywordShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateSearchKeywordShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchKeywordShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateSearchKeywordShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSearchKeywordShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateSearchKeywordShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateSearchKeywordShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSearchKeywordShrinkHeaders) SetAccountContextShrink(v string) *CreateSearchKeywordShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateSearchKeywordShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
