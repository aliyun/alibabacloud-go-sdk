// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsListByTypeReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StatisticsListByTypeReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StatisticsListByTypeReportHeadersAccountContext) *StatisticsListByTypeReportHeaders
	GetAccountContext() *StatisticsListByTypeReportHeadersAccountContext
}

type StatisticsListByTypeReportHeaders struct {
	CommonHeaders  map[string]*string                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StatisticsListByTypeReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StatisticsListByTypeReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportHeaders) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StatisticsListByTypeReportHeaders) GetAccountContext() *StatisticsListByTypeReportHeadersAccountContext {
	return s.AccountContext
}

func (s *StatisticsListByTypeReportHeaders) SetCommonHeaders(v map[string]*string) *StatisticsListByTypeReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StatisticsListByTypeReportHeaders) SetAccountContext(v *StatisticsListByTypeReportHeadersAccountContext) *StatisticsListByTypeReportHeaders {
	s.AccountContext = v
	return s
}

func (s *StatisticsListByTypeReportHeaders) Validate() error {
	return dara.Validate(s)
}

type StatisticsListByTypeReportHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StatisticsListByTypeReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StatisticsListByTypeReportHeadersAccountContext) SetAccountId(v string) *StatisticsListByTypeReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StatisticsListByTypeReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
