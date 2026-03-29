// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataDiagnosisStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDataDiagnosisStatisticsResponseBody
	GetRequestId() *string
	SetStatistics(v *QueryDataDiagnosisStatisticsResponseBodyStatistics) *QueryDataDiagnosisStatisticsResponseBody
	GetStatistics() *QueryDataDiagnosisStatisticsResponseBodyStatistics
}

type QueryDataDiagnosisStatisticsResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics *QueryDataDiagnosisStatisticsResponseBodyStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Struct"`
}

func (s QueryDataDiagnosisStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDataDiagnosisStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataDiagnosisStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDataDiagnosisStatisticsResponseBody) GetStatistics() *QueryDataDiagnosisStatisticsResponseBodyStatistics {
	return s.Statistics
}

func (s *QueryDataDiagnosisStatisticsResponseBody) SetRequestId(v string) *QueryDataDiagnosisStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponseBody) SetStatistics(v *QueryDataDiagnosisStatisticsResponseBodyStatistics) *QueryDataDiagnosisStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponseBody) Validate() error {
	if s.Statistics != nil {
		if err := s.Statistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDataDiagnosisStatisticsResponseBodyStatistics struct {
	FailedDates []*string `json:"FailedDates,omitempty" xml:"FailedDates,omitempty" type:"Repeated"`
	NoDataDates []*string `json:"NoDataDates,omitempty" xml:"NoDataDates,omitempty" type:"Repeated"`
}

func (s QueryDataDiagnosisStatisticsResponseBodyStatistics) String() string {
	return dara.Prettify(s)
}

func (s QueryDataDiagnosisStatisticsResponseBodyStatistics) GoString() string {
	return s.String()
}

func (s *QueryDataDiagnosisStatisticsResponseBodyStatistics) GetFailedDates() []*string {
	return s.FailedDates
}

func (s *QueryDataDiagnosisStatisticsResponseBodyStatistics) GetNoDataDates() []*string {
	return s.NoDataDates
}

func (s *QueryDataDiagnosisStatisticsResponseBodyStatistics) SetFailedDates(v []*string) *QueryDataDiagnosisStatisticsResponseBodyStatistics {
	s.FailedDates = v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponseBodyStatistics) SetNoDataDates(v []*string) *QueryDataDiagnosisStatisticsResponseBodyStatistics {
	s.NoDataDates = v
	return s
}

func (s *QueryDataDiagnosisStatisticsResponseBodyStatistics) Validate() error {
	return dara.Validate(s)
}
