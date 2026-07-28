// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkMetricsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *GetNisNetworkMetricsShrinkRequest
	GetAccountIds() []*string
	SetBeginTime(v int64) *GetNisNetworkMetricsShrinkRequest
	GetBeginTime() *int64
	SetDimensionsShrink(v string) *GetNisNetworkMetricsShrinkRequest
	GetDimensionsShrink() *string
	SetEndTime(v int64) *GetNisNetworkMetricsShrinkRequest
	GetEndTime() *int64
	SetMetricName(v string) *GetNisNetworkMetricsShrinkRequest
	GetMetricName() *string
	SetRegionNo(v string) *GetNisNetworkMetricsShrinkRequest
	GetRegionNo() *string
	SetResourceType(v string) *GetNisNetworkMetricsShrinkRequest
	GetResourceType() *string
	SetScanBy(v string) *GetNisNetworkMetricsShrinkRequest
	GetScanBy() *string
	SetStepMinutes(v int32) *GetNisNetworkMetricsShrinkRequest
	GetStepMinutes() *int32
	SetUseCrossAccount(v bool) *GetNisNetworkMetricsShrinkRequest
	GetUseCrossAccount() *bool
}

type GetNisNetworkMetricsShrinkRequest struct {
	// Explicitly passes sub-account IDs.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The start time, in **ms**, in **UNIX*	- timestamp format. If not specified, the most recent 1 hour is queried by default. The earliest start time is 7 days ago.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The collection of metric query parameters for specific business scenarios. For metric description of each scenario, see [GetNisNetworkMetrics](https://help.aliyun.com/document_detail/2833348.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	DimensionsShrink *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The end time, in **ms**, in **UNIX*	- timestamp format. If not specified, the most recent 1 hour is queried by default. If only BeginTime is specified, the 1 hour after BeginTime is queried. The maximum time span between the end time and start time is 24 hours.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The metric name. Valid values:
	//
	// -   bps: bits per second.
	//
	// -   pps: packets per second.
	//
	// -   rtt: round-trip time when establishing a TCP connection.
	//
	// -   RetransmitRate: retransmission rate.
	//
	// -   RatelimitDropPps: rate of packets dropped due to throttling.
	//
	// -   ActiveSessionCount: concurrent sessions.
	//
	// -   NewSessionPerSecond: new sessions per second.
	//
	// -   BandwidthUtilization: bandwidth utilization.
	//
	// -   passRate: inspection pass rate.
	//
	// > If no RTT data is available within the selected time range, the connection is a persistent connection and no initial connection was established during that period.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// Analyzes traffic by the Alibaba Cloud network resource type used for traffic forwarding. Valid values:
	//
	// - AccessInternetIpV4: all Alibaba Cloud public IPv4 addresses.
	//
	// - AccessInternetIpV4Limited: all region-throttled Alibaba Cloud public IPv4 addresses.
	//
	// - ElasticIP: Elastic IP Address (EIP) (IPv4).
	//
	// - PublicIpEcs: static public IP address bound to an ECS instance (IPv4).
	//
	// - PublicIpClb: static public IP address bound to a CLB instance (IPv4).
	//
	// - NAT: public traffic through SNAT.
	//
	// - TR: traffic through Cloud Enterprise Network (CEN) transit routers.
	//
	// - TRAttachment: traffic through CEN connection instances, including intra-region and inter-region connections. Intra-region connections have inbound and outbound directions. Inter-region connections have only the outbound direction.
	//
	// - VBR: traffic through virtual border routers.
	//
	// - GA: traffic through Global Accelerator.
	//
	// - InternetProbing: Internet quality probing data.
	//
	// - IntranetProbing: internal network quality probing data.
	//
	// - NisInspectionHistoryReportScore: inspection history scores.
	//
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIPV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The sort order. Default value: TimestampAscending. Valid values:
	//
	// - TimestampAscending: sorts by time in ascending order.
	//
	// - TimestampDescending: sorts by time in descending order.
	//
	// example:
	//
	// TimestampAscending
	ScanBy      *string `json:"ScanBy,omitempty" xml:"ScanBy,omitempty"`
	StepMinutes *int32  `json:"StepMinutes,omitempty" xml:"StepMinutes,omitempty"`
	// Specifies whether to use cross-account access mode. This is a reserved parameter and is not currently supported.
	//
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkMetricsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsShrinkRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetNisNetworkMetricsShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisNetworkMetricsShrinkRequest) GetDimensionsShrink() *string {
	return s.DimensionsShrink
}

func (s *GetNisNetworkMetricsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisNetworkMetricsShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetNisNetworkMetricsShrinkRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkMetricsShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetNisNetworkMetricsShrinkRequest) GetScanBy() *string {
	return s.ScanBy
}

func (s *GetNisNetworkMetricsShrinkRequest) GetStepMinutes() *int32 {
	return s.StepMinutes
}

func (s *GetNisNetworkMetricsShrinkRequest) GetUseCrossAccount() *bool {
	return s.UseCrossAccount
}

func (s *GetNisNetworkMetricsShrinkRequest) SetAccountIds(v []*string) *GetNisNetworkMetricsShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetBeginTime(v int64) *GetNisNetworkMetricsShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetDimensionsShrink(v string) *GetNisNetworkMetricsShrinkRequest {
	s.DimensionsShrink = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetEndTime(v int64) *GetNisNetworkMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetMetricName(v string) *GetNisNetworkMetricsShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetRegionNo(v string) *GetNisNetworkMetricsShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetResourceType(v string) *GetNisNetworkMetricsShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetScanBy(v string) *GetNisNetworkMetricsShrinkRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetStepMinutes(v int32) *GetNisNetworkMetricsShrinkRequest {
	s.StepMinutes = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) SetUseCrossAccount(v bool) *GetNisNetworkMetricsShrinkRequest {
	s.UseCrossAccount = &v
	return s
}

func (s *GetNisNetworkMetricsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
