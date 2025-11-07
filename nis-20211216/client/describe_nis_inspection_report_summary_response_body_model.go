// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionReportSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeNisInspectionReportSummaryResponseBody
	GetEndTime() *string
	SetInspectionReportId(v string) *DescribeNisInspectionReportSummaryResponseBody
	GetInspectionReportId() *string
	SetInspectionTaskId(v string) *DescribeNisInspectionReportSummaryResponseBody
	GetInspectionTaskId() *string
	SetRequestId(v string) *DescribeNisInspectionReportSummaryResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeNisInspectionReportSummaryResponseBody
	GetStartTime() *string
	SetStatus(v string) *DescribeNisInspectionReportSummaryResponseBody
	GetStatus() *string
	SetSummary(v *DescribeNisInspectionReportSummaryResponseBodySummary) *DescribeNisInspectionReportSummaryResponseBody
	GetSummary() *DescribeNisInspectionReportSummaryResponseBodySummary
}

type DescribeNisInspectionReportSummaryResponseBody struct {
	// example:
	//
	// 2024-06-03 09:36:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// nir-38abb318b27b49cc9a01
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
	// example:
	//
	// ni-8svmpe0yso2bhzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	// example:
	//
	// 4838F3F2-30E1-5D82-B25A-B9FE33BC3E25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2024-06-03 09:35:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Active
	Status  *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Summary *DescribeNisInspectionReportSummaryResponseBodySummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
}

func (s DescribeNisInspectionReportSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNisInspectionReportSummaryResponseBody) GetSummary() *DescribeNisInspectionReportSummaryResponseBodySummary {
	return s.Summary
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetEndTime(v string) *DescribeNisInspectionReportSummaryResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetInspectionReportId(v string) *DescribeNisInspectionReportSummaryResponseBody {
	s.InspectionReportId = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetInspectionTaskId(v string) *DescribeNisInspectionReportSummaryResponseBody {
	s.InspectionTaskId = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetRequestId(v string) *DescribeNisInspectionReportSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetStartTime(v string) *DescribeNisInspectionReportSummaryResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetStatus(v string) *DescribeNisInspectionReportSummaryResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) SetSummary(v *DescribeNisInspectionReportSummaryResponseBodySummary) *DescribeNisInspectionReportSummaryResponseBody {
	s.Summary = v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBody) Validate() error {
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNisInspectionReportSummaryResponseBodySummary struct {
	// example:
	//
	// 11
	CheckItemCount *int32 `json:"CheckItemCount,omitempty" xml:"CheckItemCount,omitempty"`
	// example:
	//
	// 123
	CheckResourceCount *int32                                                                  `json:"CheckResourceCount,omitempty" xml:"CheckResourceCount,omitempty"`
	PassRateSummary    []*DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary `json:"PassRateSummary,omitempty" xml:"PassRateSummary,omitempty" type:"Repeated"`
	RiskSummary        []*DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary     `json:"RiskSummary,omitempty" xml:"RiskSummary,omitempty" type:"Repeated"`
}

func (s DescribeNisInspectionReportSummaryResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportSummaryResponseBodySummary) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) GetCheckItemCount() *int32 {
	return s.CheckItemCount
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) GetCheckResourceCount() *int32 {
	return s.CheckResourceCount
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) GetPassRateSummary() []*DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary {
	return s.PassRateSummary
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) GetRiskSummary() []*DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary {
	return s.RiskSummary
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) SetCheckItemCount(v int32) *DescribeNisInspectionReportSummaryResponseBodySummary {
	s.CheckItemCount = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) SetCheckResourceCount(v int32) *DescribeNisInspectionReportSummaryResponseBodySummary {
	s.CheckResourceCount = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) SetPassRateSummary(v []*DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) *DescribeNisInspectionReportSummaryResponseBodySummary {
	s.PassRateSummary = v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) SetRiskSummary(v []*DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) *DescribeNisInspectionReportSummaryResponseBodySummary {
	s.RiskSummary = v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummary) Validate() error {
	if s.PassRateSummary != nil {
		for _, item := range s.PassRateSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RiskSummary != nil {
		for _, item := range s.RiskSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary struct {
	// example:
	//
	// 0.98
	PassRate *float64 `json:"PassRate,omitempty" xml:"PassRate,omitempty"`
	// example:
	//
	// Stability
	PassRateScope *string `json:"PassRateScope,omitempty" xml:"PassRateScope,omitempty"`
}

func (s DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) GetPassRate() *float64 {
	return s.PassRate
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) GetPassRateScope() *string {
	return s.PassRateScope
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) SetPassRate(v float64) *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary {
	s.PassRate = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) SetPassRateScope(v string) *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary {
	s.PassRateScope = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryPassRateSummary) Validate() error {
	return dara.Validate(s)
}

type DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary struct {
	// example:
	//
	// 0
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// example:
	//
	// 3
	RiskCount *int32 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// HighRisk
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// StabilityRisk
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) GetResourceCount() *int32 {
	return s.ResourceCount
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) GetRiskCount() *int32 {
	return s.RiskCount
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) GetRiskType() *string {
	return s.RiskType
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) SetResourceCount(v int32) *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary {
	s.ResourceCount = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) SetRiskCount(v int32) *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary {
	s.RiskCount = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) SetRiskLevel(v string) *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary {
	s.RiskLevel = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) SetRiskType(v string) *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary {
	s.RiskType = &v
	return s
}

func (s *DescribeNisInspectionReportSummaryResponseBodySummaryRiskSummary) Validate() error {
	return dara.Validate(s)
}
