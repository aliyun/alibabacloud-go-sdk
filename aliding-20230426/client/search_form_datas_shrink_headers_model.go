// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDatasShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDatasShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SearchFormDatasShrinkHeaders
	GetAccountContextShrink() *string
}

type SearchFormDatasShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SearchFormDatasShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDatasShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDatasShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SearchFormDatasShrinkHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDatasShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDatasShrinkHeaders) SetAccountContextShrink(v string) *SearchFormDatasShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SearchFormDatasShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
