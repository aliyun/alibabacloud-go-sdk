// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMinutesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMinutesHeadersAccountContext) *QueryMinutesHeaders
	GetAccountContext() *QueryMinutesHeadersAccountContext
}

type QueryMinutesHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMinutesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMinutesHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMinutesHeaders) GetAccountContext() *QueryMinutesHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMinutesHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesHeaders) SetAccountContext(v *QueryMinutesHeadersAccountContext) *QueryMinutesHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMinutesHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMinutesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMinutesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMinutesHeadersAccountContext) SetAccountId(v string) *QueryMinutesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMinutesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
