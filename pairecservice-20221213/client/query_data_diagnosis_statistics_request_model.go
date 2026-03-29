// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataDiagnosisStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *QueryDataDiagnosisStatisticsRequest
	GetEndDate() *string
	SetInstanceId(v string) *QueryDataDiagnosisStatisticsRequest
	GetInstanceId() *string
	SetRemainRateType(v string) *QueryDataDiagnosisStatisticsRequest
	GetRemainRateType() *string
	SetStartDate(v string) *QueryDataDiagnosisStatisticsRequest
	GetStartDate() *string
}

type QueryDataDiagnosisStatisticsRequest struct {
	// example:
	//
	// 2023-08-08
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// learn-pairec-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Period
	RemainRateType *string `json:"RemainRateType,omitempty" xml:"RemainRateType,omitempty"`
	// example:
	//
	// 2023-08-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryDataDiagnosisStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDataDiagnosisStatisticsRequest) GoString() string {
	return s.String()
}

func (s *QueryDataDiagnosisStatisticsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryDataDiagnosisStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDataDiagnosisStatisticsRequest) GetRemainRateType() *string {
	return s.RemainRateType
}

func (s *QueryDataDiagnosisStatisticsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryDataDiagnosisStatisticsRequest) SetEndDate(v string) *QueryDataDiagnosisStatisticsRequest {
	s.EndDate = &v
	return s
}

func (s *QueryDataDiagnosisStatisticsRequest) SetInstanceId(v string) *QueryDataDiagnosisStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryDataDiagnosisStatisticsRequest) SetRemainRateType(v string) *QueryDataDiagnosisStatisticsRequest {
	s.RemainRateType = &v
	return s
}

func (s *QueryDataDiagnosisStatisticsRequest) SetStartDate(v string) *QueryDataDiagnosisStatisticsRequest {
	s.StartDate = &v
	return s
}

func (s *QueryDataDiagnosisStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
