// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePeakTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowChart(v []*DescribePeakTrendResponseBodyFlowChart) *DescribePeakTrendResponseBody
	GetFlowChart() []*DescribePeakTrendResponseBodyFlowChart
	SetRequestId(v string) *DescribePeakTrendResponseBody
	GetRequestId() *string
}

type DescribePeakTrendResponseBody struct {
	// An array of the QPS statistics of the WAF instance.
	FlowChart []*DescribePeakTrendResponseBodyFlowChart `json:"FlowChart,omitempty" xml:"FlowChart,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 9D11AC3A-A10C-56E7-A342-E87EC892BAE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePeakTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePeakTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponseBody) GetFlowChart() []*DescribePeakTrendResponseBodyFlowChart {
	return s.FlowChart
}

func (s *DescribePeakTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePeakTrendResponseBody) SetFlowChart(v []*DescribePeakTrendResponseBodyFlowChart) *DescribePeakTrendResponseBody {
	s.FlowChart = v
	return s
}

func (s *DescribePeakTrendResponseBody) SetRequestId(v string) *DescribePeakTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePeakTrendResponseBody) Validate() error {
	if s.FlowChart != nil {
		for _, item := range s.FlowChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePeakTrendResponseBodyFlowChart struct {
	// The number of requests that are monitored or blocked by the custom rule (access control) module.
	//
	// example:
	//
	// 0
	AclSum *int64 `json:"AclSum,omitempty" xml:"AclSum,omitempty"`
	// The number of requests that are monitored or blocked by the scan protection module.
	//
	// example:
	//
	// 0
	AntiScanSum *int64 `json:"AntiScanSum,omitempty" xml:"AntiScanSum,omitempty"`
	// The number of requests that are monitored or blocked by the HTTP flood protection module.
	//
	// example:
	//
	// 0
	CcSum *int64 `json:"CcSum,omitempty" xml:"CcSum,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 2622
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The serial number of the time interval. The serial numbers are arranged in chronological order.
	//
	// example:
	//
	// 10
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The number of requests that are monitored or blocked by the regular expression protection engine.
	//
	// example:
	//
	// 0
	WafSum *int64 `json:"WafSum,omitempty" xml:"WafSum,omitempty"`
}

func (s DescribePeakTrendResponseBodyFlowChart) String() string {
	return dara.Prettify(s)
}

func (s DescribePeakTrendResponseBodyFlowChart) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendResponseBodyFlowChart) GetAclSum() *int64 {
	return s.AclSum
}

func (s *DescribePeakTrendResponseBodyFlowChart) GetAntiScanSum() *int64 {
	return s.AntiScanSum
}

func (s *DescribePeakTrendResponseBodyFlowChart) GetCcSum() *int64 {
	return s.CcSum
}

func (s *DescribePeakTrendResponseBodyFlowChart) GetCount() *int64 {
	return s.Count
}

func (s *DescribePeakTrendResponseBodyFlowChart) GetIndex() *int64 {
	return s.Index
}

func (s *DescribePeakTrendResponseBodyFlowChart) GetWafSum() *int64 {
	return s.WafSum
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetAclSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.AclSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetAntiScanSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.AntiScanSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetCcSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.CcSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetCount(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.Count = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetIndex(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.Index = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) SetWafSum(v int64) *DescribePeakTrendResponseBodyFlowChart {
	s.WafSum = &v
	return s
}

func (s *DescribePeakTrendResponseBodyFlowChart) Validate() error {
	return dara.Validate(s)
}
