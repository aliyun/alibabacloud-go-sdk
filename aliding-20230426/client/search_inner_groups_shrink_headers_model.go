// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInnerGroupsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchInnerGroupsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SearchInnerGroupsShrinkHeaders
	GetAccountContextShrink() *string
}

type SearchInnerGroupsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SearchInnerGroupsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchInnerGroupsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SearchInnerGroupsShrinkHeaders) SetCommonHeaders(v map[string]*string) *SearchInnerGroupsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchInnerGroupsShrinkHeaders) SetAccountContextShrink(v string) *SearchInnerGroupsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SearchInnerGroupsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
