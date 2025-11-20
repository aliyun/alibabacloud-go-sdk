// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataSecondGenerationNoTableFieldHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationNoTableFieldHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) *SearchFormDataSecondGenerationNoTableFieldHeaders
	GetAccountContext() *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext
}

type SearchFormDataSecondGenerationNoTableFieldHeaders struct {
	CommonHeaders  map[string]*string                                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SearchFormDataSecondGenerationNoTableFieldHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) GetAccountContext() *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext {
	return s.AccountContext
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataSecondGenerationNoTableFieldHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) SetAccountContext(v *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) *SearchFormDataSecondGenerationNoTableFieldHeaders {
	s.AccountContext = v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) SetAccountId(v string) *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SearchFormDataSecondGenerationNoTableFieldHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
