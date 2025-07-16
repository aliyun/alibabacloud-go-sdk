// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SearchFormDataSecondGenerationShrinkHeaders
	GetAccountContextShrink() *string
}

type SearchFormDataSecondGenerationShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SearchFormDataSecondGenerationShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDataSecondGenerationShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SearchFormDataSecondGenerationShrinkHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationShrinkHeaders) SetAccountContextShrink(v string) *SearchFormDataSecondGenerationShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SearchFormDataSecondGenerationShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
