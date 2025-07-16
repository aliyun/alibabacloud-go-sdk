// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSimpleListReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SimpleListReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SimpleListReportHeadersAccountContext) *SimpleListReportHeaders
	GetAccountContext() *SimpleListReportHeadersAccountContext
}

type SimpleListReportHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SimpleListReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SimpleListReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportHeaders) GoString() string {
	return s.String()
}

func (s *SimpleListReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SimpleListReportHeaders) GetAccountContext() *SimpleListReportHeadersAccountContext {
	return s.AccountContext
}

func (s *SimpleListReportHeaders) SetCommonHeaders(v map[string]*string) *SimpleListReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SimpleListReportHeaders) SetAccountContext(v *SimpleListReportHeadersAccountContext) *SimpleListReportHeaders {
	s.AccountContext = v
	return s
}

func (s *SimpleListReportHeaders) Validate() error {
	return dara.Validate(s)
}

type SimpleListReportHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SimpleListReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SimpleListReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SimpleListReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SimpleListReportHeadersAccountContext) SetAccountId(v string) *SimpleListReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SimpleListReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
