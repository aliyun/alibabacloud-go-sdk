// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowChart(v []*DescribeFlowChartResponseBodyFlowChart) *DescribeFlowChartResponseBody
	GetFlowChart() []*DescribeFlowChartResponseBodyFlowChart
	SetRequestId(v string) *DescribeFlowChartResponseBody
	GetRequestId() *string
}

type DescribeFlowChartResponseBody struct {
	// The traffic statistics.
	FlowChart []*DescribeFlowChartResponseBodyFlowChart `json:"FlowChart,omitempty" xml:"FlowChart,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// BFA71416-670E-585D-AAE6-E7BBEE248FAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponseBody) GetFlowChart() []*DescribeFlowChartResponseBodyFlowChart {
	return s.FlowChart
}

func (s *DescribeFlowChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowChartResponseBody) SetFlowChart(v []*DescribeFlowChartResponseBodyFlowChart) *DescribeFlowChartResponseBody {
	s.FlowChart = v
	return s
}

func (s *DescribeFlowChartResponseBody) SetRequestId(v string) *DescribeFlowChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowChartResponseBody) Validate() error {
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

type DescribeFlowChartResponseBodyFlowChart struct {
	// The number of requests that are blocked by custom access control list (ACL) rules.
	//
	// example:
	//
	// 0
	AclCustomBlockSum *int64 `json:"AclCustomBlockSum,omitempty" xml:"AclCustomBlockSum,omitempty"`
	// The number of requests that are monitored by custom ACL rules.
	//
	// example:
	//
	// 0
	AclCustomReportsSum *int64 `json:"AclCustomReportsSum,omitempty" xml:"AclCustomReportsSum,omitempty"`
	// The number of requests that are blocked by scan protection rules.
	//
	// example:
	//
	// 0
	AntiScanBlockSum *int64 `json:"AntiScanBlockSum,omitempty" xml:"AntiScanBlockSum,omitempty"`
	// The number of requests that are blocked by bot management rules.
	//
	// example:
	//
	// 0
	AntibotBlockSum *int64 `json:"AntibotBlockSum,omitempty" xml:"AntibotBlockSum,omitempty"`
	// The number of requests that are monitored by bot management rules.
	//
	// example:
	//
	// 0
	AntibotReportSum *string `json:"AntibotReportSum,omitempty" xml:"AntibotReportSum,omitempty"`
	// The number of requests that are monitored by scan protection rules.
	//
	// example:
	//
	// 0
	AntiscanReportsSum *int64 `json:"AntiscanReportsSum,omitempty" xml:"AntiscanReportsSum,omitempty"`
	// The number of requests that are blocked by the IP address blacklist.
	//
	// example:
	//
	// 0
	BlacklistBlockSum *string `json:"BlacklistBlockSum,omitempty" xml:"BlacklistBlockSum,omitempty"`
	// The number of requests that are monitored by the IP address blacklist.
	//
	// example:
	//
	// 0
	BlacklistReportsSum *int64 `json:"BlacklistReportsSum,omitempty" xml:"BlacklistReportsSum,omitempty"`
	// The number of requests that are blocked by custom HTTP flood protection rules.
	//
	// example:
	//
	// 0
	CcCustomBlockSum *int64 `json:"CcCustomBlockSum,omitempty" xml:"CcCustomBlockSum,omitempty"`
	// The number of requests that are monitored by custom HTTP flood protection rules.
	//
	// example:
	//
	// 0
	CcCustomReportsSum *int64 `json:"CcCustomReportsSum,omitempty" xml:"CcCustomReportsSum,omitempty"`
	// The number of requests that are blocked by HTTP flood protection rules created by the system.
	//
	// example:
	//
	// 0
	CcSystemBlocksSum *int64 `json:"CcSystemBlocksSum,omitempty" xml:"CcSystemBlocksSum,omitempty"`
	// The number of requests that are monitored by HTTP flood protection rules created by the system.
	//
	// example:
	//
	// 0
	CcSystemReportsSum *int64 `json:"CcSystemReportsSum,omitempty" xml:"CcSystemReportsSum,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 2932
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The total number of requests that are redirected to the WAF instance.
	//
	// example:
	//
	// 121645464
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The serial number of the time interval. The serial numbers are arranged in chronological order.
	//
	// example:
	//
	// 10
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The peak traffic.
	//
	// example:
	//
	// 2932
	MaxPv *int64 `json:"MaxPv,omitempty" xml:"MaxPv,omitempty"`
	// The total number of requests that are forwarded by the WAF instance.
	//
	// example:
	//
	// 1200540464
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The number of requests that are blocked by rate limiting rules.
	//
	// example:
	//
	// 0
	RatelimitBlockSum *int64 `json:"RatelimitBlockSum,omitempty" xml:"RatelimitBlockSum,omitempty"`
	// The number of requests that are monitored by rate limiting rules.
	//
	// example:
	//
	// 0
	RatelimitReportSum *int64 `json:"RatelimitReportSum,omitempty" xml:"RatelimitReportSum,omitempty"`
	// The number of requests that are blocked by region blacklist rules.
	//
	// example:
	//
	// 0
	RegionBlockBlocksSum *int64 `json:"RegionBlockBlocksSum,omitempty" xml:"RegionBlockBlocksSum,omitempty"`
	// The number of requests that are monitored by region blacklist rules.
	//
	// example:
	//
	// 0
	RegionBlockReportsSum *int64 `json:"RegionBlockReportsSum,omitempty" xml:"RegionBlockReportsSum,omitempty"`
	// The total number of bot requests.
	//
	// example:
	//
	// 1110
	RobotCount *int64 `json:"RobotCount,omitempty" xml:"RobotCount,omitempty"`
	// The number of requests that are blocked by basic protection rules.
	//
	// example:
	//
	// 0
	WafBlockSum *int64 `json:"WafBlockSum,omitempty" xml:"WafBlockSum,omitempty"`
	// The number of requests that are monitored by basic protection rules.
	//
	// example:
	//
	// 0
	WafReportSum *string `json:"WafReportSum,omitempty" xml:"WafReportSum,omitempty"`
}

func (s DescribeFlowChartResponseBodyFlowChart) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowChartResponseBodyFlowChart) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetAclCustomBlockSum() *int64 {
	return s.AclCustomBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetAclCustomReportsSum() *int64 {
	return s.AclCustomReportsSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetAntiScanBlockSum() *int64 {
	return s.AntiScanBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetAntibotBlockSum() *int64 {
	return s.AntibotBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetAntibotReportSum() *string {
	return s.AntibotReportSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetAntiscanReportsSum() *int64 {
	return s.AntiscanReportsSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetBlacklistBlockSum() *string {
	return s.BlacklistBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetBlacklistReportsSum() *int64 {
	return s.BlacklistReportsSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetCcCustomBlockSum() *int64 {
	return s.CcCustomBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetCcCustomReportsSum() *int64 {
	return s.CcCustomReportsSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetCcSystemBlocksSum() *int64 {
	return s.CcSystemBlocksSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetCcSystemReportsSum() *int64 {
	return s.CcSystemReportsSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetCount() *int64 {
	return s.Count
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetMaxPv() *int64 {
	return s.MaxPv
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetRatelimitBlockSum() *int64 {
	return s.RatelimitBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetRatelimitReportSum() *int64 {
	return s.RatelimitReportSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetRegionBlockBlocksSum() *int64 {
	return s.RegionBlockBlocksSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetRegionBlockReportsSum() *int64 {
	return s.RegionBlockReportsSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetRobotCount() *int64 {
	return s.RobotCount
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetWafBlockSum() *int64 {
	return s.WafBlockSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) GetWafReportSum() *string {
	return s.WafReportSum
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAclCustomBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AclCustomBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAclCustomReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AclCustomReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntiScanBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntiScanBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntibotBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntibotBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntibotReportSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.AntibotReportSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetAntiscanReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.AntiscanReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetBlacklistBlockSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.BlacklistBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetBlacklistReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.BlacklistReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcCustomBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcCustomBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcCustomReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcCustomReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcSystemBlocksSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcSystemBlocksSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCcSystemReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.CcSystemReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetCount(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.Count = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetInBytes(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.InBytes = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetIndex(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.Index = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetMaxPv(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.MaxPv = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetOutBytes(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.OutBytes = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRatelimitBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RatelimitBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRatelimitReportSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RatelimitReportSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRegionBlockBlocksSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RegionBlockBlocksSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRegionBlockReportsSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RegionBlockReportsSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetRobotCount(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.RobotCount = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetWafBlockSum(v int64) *DescribeFlowChartResponseBodyFlowChart {
	s.WafBlockSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) SetWafReportSum(v string) *DescribeFlowChartResponseBodyFlowChart {
	s.WafReportSum = &v
	return s
}

func (s *DescribeFlowChartResponseBodyFlowChart) Validate() error {
	return dara.Validate(s)
}
