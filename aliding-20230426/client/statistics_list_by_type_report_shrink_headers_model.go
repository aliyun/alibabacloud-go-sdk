// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsListByTypeReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StatisticsListByTypeReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StatisticsListByTypeReportShrinkHeaders
	GetAccountContextShrink() *string
}

type StatisticsListByTypeReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StatisticsListByTypeReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StatisticsListByTypeReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StatisticsListByTypeReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StatisticsListByTypeReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StatisticsListByTypeReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *StatisticsListByTypeReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StatisticsListByTypeReportShrinkHeaders) SetAccountContextShrink(v string) *StatisticsListByTypeReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StatisticsListByTypeReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
