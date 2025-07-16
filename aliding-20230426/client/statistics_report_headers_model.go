// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StatisticsReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StatisticsReportHeadersAccountContext) *StatisticsReportHeaders
	GetAccountContext() *StatisticsReportHeadersAccountContext
}

type StatisticsReportHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StatisticsReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StatisticsReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportHeaders) GoString() string {
	return s.String()
}

func (s *StatisticsReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StatisticsReportHeaders) GetAccountContext() *StatisticsReportHeadersAccountContext {
	return s.AccountContext
}

func (s *StatisticsReportHeaders) SetCommonHeaders(v map[string]*string) *StatisticsReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StatisticsReportHeaders) SetAccountContext(v *StatisticsReportHeadersAccountContext) *StatisticsReportHeaders {
	s.AccountContext = v
	return s
}

func (s *StatisticsReportHeaders) Validate() error {
	return dara.Validate(s)
}

type StatisticsReportHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StatisticsReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StatisticsReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StatisticsReportHeadersAccountContext) SetAccountId(v string) *StatisticsReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StatisticsReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
