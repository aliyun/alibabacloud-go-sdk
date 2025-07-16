// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInnerGroupsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SearchInnerGroupsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SearchInnerGroupsHeadersAccountContext) *SearchInnerGroupsHeaders
	GetAccountContext() *SearchInnerGroupsHeadersAccountContext
}

type SearchInnerGroupsHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SearchInnerGroupsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SearchInnerGroupsHeaders) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsHeaders) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SearchInnerGroupsHeaders) GetAccountContext() *SearchInnerGroupsHeadersAccountContext {
	return s.AccountContext
}

func (s *SearchInnerGroupsHeaders) SetCommonHeaders(v map[string]*string) *SearchInnerGroupsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SearchInnerGroupsHeaders) SetAccountContext(v *SearchInnerGroupsHeadersAccountContext) *SearchInnerGroupsHeaders {
	s.AccountContext = v
	return s
}

func (s *SearchInnerGroupsHeaders) Validate() error {
	return dara.Validate(s)
}

type SearchInnerGroupsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SearchInnerGroupsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SearchInnerGroupsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SearchInnerGroupsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SearchInnerGroupsHeadersAccountContext) SetAccountId(v string) *SearchInnerGroupsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SearchInnerGroupsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
