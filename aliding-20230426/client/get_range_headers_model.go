// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRangeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRangeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetRangeHeadersAccountContext) *GetRangeHeaders
	GetAccountContext() *GetRangeHeadersAccountContext
}

type GetRangeHeaders struct {
	CommonHeaders  map[string]*string             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetRangeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetRangeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRangeHeaders) GoString() string {
	return s.String()
}

func (s *GetRangeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRangeHeaders) GetAccountContext() *GetRangeHeadersAccountContext {
	return s.AccountContext
}

func (s *GetRangeHeaders) SetCommonHeaders(v map[string]*string) *GetRangeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRangeHeaders) SetAccountContext(v *GetRangeHeadersAccountContext) *GetRangeHeaders {
	s.AccountContext = v
	return s
}

func (s *GetRangeHeaders) Validate() error {
	return dara.Validate(s)
}

type GetRangeHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetRangeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetRangeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetRangeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetRangeHeadersAccountContext) SetAccountId(v string) *GetRangeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetRangeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
