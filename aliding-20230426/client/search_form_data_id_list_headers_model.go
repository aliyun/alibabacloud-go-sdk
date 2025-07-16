// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataIdListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchFormDataIdListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SearchFormDataIdListHeadersAccountContext) *SearchFormDataIdListHeaders
	GetAccountContext() *SearchFormDataIdListHeadersAccountContext
}

type SearchFormDataIdListHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SearchFormDataIdListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SearchFormDataIdListHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataIdListHeaders) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchFormDataIdListHeaders) GetAccountContext() *SearchFormDataIdListHeadersAccountContext {
	return s.AccountContext
}

func (s *SearchFormDataIdListHeaders) SetCommonHeaders(v map[string]*string) *SearchFormDataIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchFormDataIdListHeaders) SetAccountContext(v *SearchFormDataIdListHeadersAccountContext) *SearchFormDataIdListHeaders {
	s.AccountContext = v
	return s
}

func (s *SearchFormDataIdListHeaders) Validate() error {
	return dara.Validate(s)
}

type SearchFormDataIdListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SearchFormDataIdListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataIdListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchFormDataIdListHeadersAccountContext) SetAccountId(v string) *SearchFormDataIdListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SearchFormDataIdListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
