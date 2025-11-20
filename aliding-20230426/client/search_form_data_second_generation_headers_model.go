// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SearchFormDataSecondGenerationHeadersAccountContext) *SearchFormDataSecondGenerationHeaders
	GetAccountContext() *SearchFormDataSecondGenerationHeadersAccountContext
}

type SearchFormDataSecondGenerationHeaders struct {
	CommonHeaders  map[string]*string                                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SearchFormDataSecondGenerationHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SearchFormDataSecondGenerationHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDataSecondGenerationHeaders) GetAccountContext() *SearchFormDataSecondGenerationHeadersAccountContext {
	return s.AccountContext
}

func (s *SearchFormDataSecondGenerationHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationHeaders) SetAccountContext(v *SearchFormDataSecondGenerationHeadersAccountContext) *SearchFormDataSecondGenerationHeaders {
	s.AccountContext = v
	return s
}

func (s *SearchFormDataSecondGenerationHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFormDataSecondGenerationHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SearchFormDataSecondGenerationHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchFormDataSecondGenerationHeadersAccountContext) SetAccountId(v string) *SearchFormDataSecondGenerationHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SearchFormDataSecondGenerationHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
