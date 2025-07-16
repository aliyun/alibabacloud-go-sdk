// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDatasHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDatasHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SearchFormDatasHeadersAccountContext) *SearchFormDatasHeaders
	GetAccountContext() *SearchFormDatasHeadersAccountContext
}

type SearchFormDatasHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SearchFormDatasHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SearchFormDatasHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDatasHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDatasHeaders) GetAccountContext() *SearchFormDatasHeadersAccountContext {
	return s.AccountContext
}

func (s *SearchFormDatasHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDatasHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDatasHeaders) SetAccountContext(v *SearchFormDatasHeadersAccountContext) *SearchFormDatasHeaders {
	s.AccountContext = v
	return s
}

func (s *SearchFormDatasHeaders) Validate() error {
	return dara.Validate(s)
}

type SearchFormDatasHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SearchFormDatasHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDatasHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SearchFormDatasHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchFormDatasHeadersAccountContext) SetAccountId(v string) *SearchFormDatasHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SearchFormDatasHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
