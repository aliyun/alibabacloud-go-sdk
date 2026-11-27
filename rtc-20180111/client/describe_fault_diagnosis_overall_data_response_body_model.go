// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisOverallDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricData(v *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) *DescribeFaultDiagnosisOverallDataResponseBody
	GetMetricData() *DescribeFaultDiagnosisOverallDataResponseBodyMetricData
	SetOverallData(v *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) *DescribeFaultDiagnosisOverallDataResponseBody
	GetOverallData() *DescribeFaultDiagnosisOverallDataResponseBodyOverallData
	SetRequestId(v string) *DescribeFaultDiagnosisOverallDataResponseBody
	GetRequestId() *string
}

type DescribeFaultDiagnosisOverallDataResponseBody struct {
	// The exception metric data.
	MetricData *DescribeFaultDiagnosisOverallDataResponseBodyMetricData `json:"MetricData,omitempty" xml:"MetricData,omitempty" type:"Struct"`
	// The list of overview data for exception diagnostics.
	OverallData *DescribeFaultDiagnosisOverallDataResponseBodyOverallData `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) GetMetricData() *DescribeFaultDiagnosisOverallDataResponseBodyMetricData {
	return s.MetricData
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) GetOverallData() *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	return s.OverallData
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetMetricData(v *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.MetricData = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetOverallData(v *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) Validate() error {
	if s.MetricData != nil {
		if err := s.MetricData.Validate(); err != nil {
			return err
		}
	}
	if s.OverallData != nil {
		if err := s.OverallData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFaultDiagnosisOverallDataResponseBodyMetricData struct {
	// The list of coordinate points for the metric trends chart.
	Nodes []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) GetNodes() []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	return s.Nodes
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) SetNodes(v []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) *DescribeFaultDiagnosisOverallDataResponseBodyMetricData {
	s.Nodes = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes struct {
	// The extended properties, where ratio indicates the proportion of abnormal users per unit time and totalCount indicates the total number of users per unit time.
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// The x-axis coordinate in the metric trends chart. The unit of time is 1 minute.
	//
	// example:
	//
	// 1615824000
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// The y-axis coordinate in the metric trends chart.
	//
	// example:
	//
	// 1
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) GetX() *string {
	return s.X
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) GetY() *string {
	return s.Y
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetExt(v map[string]interface{}) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.Ext = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetX(v string) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetY(v string) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.Y = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisOverallDataResponseBodyOverallData struct {
	// The number of abnormal user visits or abnormal users.
	//
	// example:
	//
	// 20
	FaultUserCount *int32 `json:"FaultUserCount,omitempty" xml:"FaultUserCount,omitempty"`
	// The total proportion of abnormal users.
	//
	// example:
	//
	// 0.5
	FaultUserRatio *float32 `json:"FaultUserRatio,omitempty" xml:"FaultUserRatio,omitempty"`
	// The total number of user visits or total users.
	//
	// example:
	//
	// 40
	TotalUserCount *int32 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyOverallData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) GetFaultUserCount() *int32 {
	return s.FaultUserCount
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) GetFaultUserRatio() *float32 {
	return s.FaultUserRatio
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) GetTotalUserCount() *int32 {
	return s.TotalUserCount
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetFaultUserCount(v int32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.FaultUserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetFaultUserRatio(v float32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.FaultUserRatio = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetTotalUserCount(v int32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.TotalUserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) Validate() error {
	return dara.Validate(s)
}
