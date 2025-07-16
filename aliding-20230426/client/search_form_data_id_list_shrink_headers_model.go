// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataIdListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDataIdListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SearchFormDataIdListShrinkHeaders
	GetAccountContextShrink() *string
}

type SearchFormDataIdListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SearchFormDataIdListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataIdListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDataIdListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SearchFormDataIdListShrinkHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataIdListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataIdListShrinkHeaders) SetAccountContextShrink(v string) *SearchFormDataIdListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SearchFormDataIdListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
