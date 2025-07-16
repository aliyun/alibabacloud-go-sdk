// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryLiveInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryLiveInfoHeadersAccountContext) *QueryLiveInfoHeaders
	GetAccountContext() *QueryLiveInfoHeadersAccountContext
}

type QueryLiveInfoHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryLiveInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryLiveInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryLiveInfoHeaders) GetAccountContext() *QueryLiveInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryLiveInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveInfoHeaders) SetAccountContext(v *QueryLiveInfoHeadersAccountContext) *QueryLiveInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryLiveInfoHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryLiveInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryLiveInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryLiveInfoHeadersAccountContext) SetAccountId(v string) *QueryLiveInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryLiveInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
