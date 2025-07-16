// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesSummaryHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMinutesSummaryHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMinutesSummaryHeadersAccountContext) *QueryMinutesSummaryHeaders
	GetAccountContext() *QueryMinutesSummaryHeadersAccountContext
}

type QueryMinutesSummaryHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMinutesSummaryHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMinutesSummaryHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryHeaders) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMinutesSummaryHeaders) GetAccountContext() *QueryMinutesSummaryHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMinutesSummaryHeaders) SetCommonHeaders(v map[string]*string) *QueryMinutesSummaryHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMinutesSummaryHeaders) SetAccountContext(v *QueryMinutesSummaryHeadersAccountContext) *QueryMinutesSummaryHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMinutesSummaryHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryMinutesSummaryHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMinutesSummaryHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesSummaryHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMinutesSummaryHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMinutesSummaryHeadersAccountContext) SetAccountId(v string) *QueryMinutesSummaryHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMinutesSummaryHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
