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
	DefenseCountStatistics *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics `json:"DefenseCountStatistics,omitempty" xml:"DefenseCountStatistics,omitempty" type:"Struct"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
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
	return dara.Validate(s)
}

type DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics struct {
	// example:
	//
	// 0
	DefenseCountTotalUsageOfCurrentMonth *int32 `json:"DefenseCountTotalUsageOfCurrentMonth,omitempty" xml:"DefenseCountTotalUsageOfCurrentMonth,omitempty"`
	// example:
	//
	// 10
	FlowPackCountRemain *int32 `json:"FlowPackCountRemain,omitempty" xml:"FlowPackCountRemain,omitempty"`
	// example:
	//
	// 0
	MaxUsableDefenseCountCurrentMonth *int32 `json:"MaxUsableDefenseCountCurrentMonth,omitempty" xml:"MaxUsableDefenseCountCurrentMonth,omitempty"`
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

func (s *DescribeDefenseCountStatisticsResponseBodyDefenseCountStatistics) Validate() error {
	return dara.Validate(s)
}
