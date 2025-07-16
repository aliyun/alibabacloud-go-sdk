// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupLiveInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryGroupLiveInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryGroupLiveInfoHeadersAccountContext) *QueryGroupLiveInfoHeaders
	GetAccountContext() *QueryGroupLiveInfoHeadersAccountContext
}

type QueryGroupLiveInfoHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryGroupLiveInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryGroupLiveInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryGroupLiveInfoHeaders) GetAccountContext() *QueryGroupLiveInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryGroupLiveInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupLiveInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupLiveInfoHeaders) SetAccountContext(v *QueryGroupLiveInfoHeadersAccountContext) *QueryGroupLiveInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryGroupLiveInfoHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryGroupLiveInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryGroupLiveInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryGroupLiveInfoHeadersAccountContext) SetAccountId(v string) *QueryGroupLiveInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryGroupLiveInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
