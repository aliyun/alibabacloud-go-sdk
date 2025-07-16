// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStatisticsReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StatisticsReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *StatisticsReportShrinkHeaders
	GetAccountContextShrink() *string
}

type StatisticsReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s StatisticsReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s StatisticsReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *StatisticsReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StatisticsReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *StatisticsReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *StatisticsReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StatisticsReportShrinkHeaders) SetAccountContextShrink(v string) *StatisticsReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *StatisticsReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
