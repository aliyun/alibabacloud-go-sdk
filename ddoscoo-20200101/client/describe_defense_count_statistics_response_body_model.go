// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseCountStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseCountStatistics(v *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) *DescribeDefenseCountStatisticsResponseBody
	GetDefenseCountStatistics() *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics
	SetRequestId(v string) *DescribeDefenseCountStatisticsResponseBody
	GetRequestId() *string
}

type DescribeDefenseCountStatisticsResponseBody struct {
	// The statistics on the number of advanced mitigation sessions.
	DefenseCountStatistics *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics `json:"DefenseCountStatistics,omitempty" xml:"DefenseCountStatistics,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 310A41FD-0990-5610-92E0-A6A55D7C6444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBody) GetDefenseCountStatistics() *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	return s.DefenseCountStatistics
}

func (s *DescribeDefenseCountStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetDefenseCountStatistics(v *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) *DescribeDefenseCountStatisticsResponseBody {
	s.DefenseCountStatistics = v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBody) SetRequestId(v string) *DescribeDefenseCountStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBody) Validate() error {
	if s.DefenseCountStatistics != nil {
		if err := s.DefenseCountStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics struct {
	// The number of advanced mitigation sessions that are used within the current calendar month.
	//
	// example:
	//
	// 0
	DefenseCountTotalUsageOfCurrentMonth *int32 `json:"DefenseCountTotalUsageOfCurrentMonth,omitempty" xml:"DefenseCountTotalUsageOfCurrentMonth,omitempty"`
	// The number of available global advanced mitigation sessions for the Insurance mitigation plan.
	//
	// example:
	//
	// 0
	FlowPackCountRemain *int32 `json:"FlowPackCountRemain,omitempty" xml:"FlowPackCountRemain,omitempty"`
	// The maximum number of advanced mitigation sessions available for the current calendar month. The advanced mitigation sessions include the advanced mitigation sessions that are provided free of charge and the global advanced mitigation sessions that you purchase.
	//
	// example:
	//
	// 20
	MaxUsableDefenseCountCurrentMonth *int32 `json:"MaxUsableDefenseCountCurrentMonth,omitempty" xml:"MaxUsableDefenseCountCurrentMonth,omitempty"`
	// The number of available global advanced mitigation sessions for the Secure Chinese Mainland Acceleration (Sec-CMA) mitigation plan.
	//
	// example:
	//
	// 0
	SecHighSpeedCountRemain *int32 `json:"SecHighSpeedCountRemain,omitempty" xml:"SecHighSpeedCountRemain,omitempty"`
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GoString() string {
	return s.String()
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GetDefenseCountTotalUsageOfCurrentMonth() *int32 {
	return s.DefenseCountTotalUsageOfCurrentMonth
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GetFlowPackCountRemain() *int32 {
	return s.FlowPackCountRemain
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GetMaxUsableDefenseCountCurrentMonth() *int32 {
	return s.MaxUsableDefenseCountCurrentMonth
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) GetSecHighSpeedCountRemain() *int32 {
	return s.SecHighSpeedCountRemain
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetDefenseCountTotalUsageOfCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.DefenseCountTotalUsageOfCurrentMonth = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetFlowPackCountRemain(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.FlowPackCountRemain = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetMaxUsableDefenseCountCurrentMonth(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.MaxUsableDefenseCountCurrentMonth = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) SetSecHighSpeedCountRemain(v int32) *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics {
	s.SecHighSpeedCountRemain = &v
	return s
}

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) Validate() error {
	return dara.Validate(s)
}
