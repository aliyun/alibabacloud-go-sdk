// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDiagnosisReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *ListDataDiagnosisReportsRequest
	GetEndDate() *string
	SetFeatureName(v string) *ListDataDiagnosisReportsRequest
	GetFeatureName() *string
	SetInstanceId(v string) *ListDataDiagnosisReportsRequest
	GetInstanceId() *string
	SetRemainRateType(v string) *ListDataDiagnosisReportsRequest
	GetRemainRateType() *string
	SetStartDate(v string) *ListDataDiagnosisReportsRequest
	GetStartDate() *string
	SetTopN(v int64) *ListDataDiagnosisReportsRequest
	GetTopN() *int64
}

type ListDataDiagnosisReportsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2022-02-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// feature1
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Period
	RemainRateType *string `json:"RemainRateType,omitempty" xml:"RemainRateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 3
	TopN *int64 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s ListDataDiagnosisReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataDiagnosisReportsRequest) GoString() string {
	return s.String()
}

func (s *ListDataDiagnosisReportsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListDataDiagnosisReportsRequest) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListDataDiagnosisReportsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDataDiagnosisReportsRequest) GetRemainRateType() *string {
	return s.RemainRateType
}

func (s *ListDataDiagnosisReportsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListDataDiagnosisReportsRequest) GetTopN() *int64 {
	return s.TopN
}

func (s *ListDataDiagnosisReportsRequest) SetEndDate(v string) *ListDataDiagnosisReportsRequest {
	s.EndDate = &v
	return s
}

func (s *ListDataDiagnosisReportsRequest) SetFeatureName(v string) *ListDataDiagnosisReportsRequest {
	s.FeatureName = &v
	return s
}

func (s *ListDataDiagnosisReportsRequest) SetInstanceId(v string) *ListDataDiagnosisReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDataDiagnosisReportsRequest) SetRemainRateType(v string) *ListDataDiagnosisReportsRequest {
	s.RemainRateType = &v
	return s
}

func (s *ListDataDiagnosisReportsRequest) SetStartDate(v string) *ListDataDiagnosisReportsRequest {
	s.StartDate = &v
	return s
}

func (s *ListDataDiagnosisReportsRequest) SetTopN(v int64) *ListDataDiagnosisReportsRequest {
	s.TopN = &v
	return s
}

func (s *ListDataDiagnosisReportsRequest) Validate() error {
	return dara.Validate(s)
}
