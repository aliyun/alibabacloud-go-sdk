// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *GetNisNetworkMetricsRequest
	GetAccountIds() []*string
	SetBeginTime(v int64) *GetNisNetworkMetricsRequest
	GetBeginTime() *int64
	SetDimensions(v []*GetNisNetworkMetricsRequestDimensions) *GetNisNetworkMetricsRequest
	GetDimensions() []*GetNisNetworkMetricsRequestDimensions
	SetEndTime(v int64) *GetNisNetworkMetricsRequest
	GetEndTime() *int64
	SetMetricName(v string) *GetNisNetworkMetricsRequest
	GetMetricName() *string
	SetRegionNo(v string) *GetNisNetworkMetricsRequest
	GetRegionNo() *string
	SetResourceType(v string) *GetNisNetworkMetricsRequest
	GetResourceType() *string
	SetScanBy(v string) *GetNisNetworkMetricsRequest
	GetScanBy() *string
	SetStepMinutes(v int32) *GetNisNetworkMetricsRequest
	GetStepMinutes() *int32
	SetUseCrossAccount(v bool) *GetNisNetworkMetricsRequest
	GetUseCrossAccount() *bool
}

type GetNisNetworkMetricsRequest struct {
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
	Dimensions []*GetNisNetworkMetricsRequestDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
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

func (s GetNisNetworkMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetNisNetworkMetricsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisNetworkMetricsRequest) GetDimensions() []*GetNisNetworkMetricsRequestDimensions {
	return s.Dimensions
}

func (s *GetNisNetworkMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisNetworkMetricsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetNisNetworkMetricsRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkMetricsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetNisNetworkMetricsRequest) GetScanBy() *string {
	return s.ScanBy
}

func (s *GetNisNetworkMetricsRequest) GetStepMinutes() *int32 {
	return s.StepMinutes
}

func (s *GetNisNetworkMetricsRequest) GetUseCrossAccount() *bool {
	return s.UseCrossAccount
}

func (s *GetNisNetworkMetricsRequest) SetAccountIds(v []*string) *GetNisNetworkMetricsRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetBeginTime(v int64) *GetNisNetworkMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetDimensions(v []*GetNisNetworkMetricsRequestDimensions) *GetNisNetworkMetricsRequest {
	s.Dimensions = v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetEndTime(v int64) *GetNisNetworkMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetMetricName(v string) *GetNisNetworkMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetRegionNo(v string) *GetNisNetworkMetricsRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetResourceType(v string) *GetNisNetworkMetricsRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetScanBy(v string) *GetNisNetworkMetricsRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetStepMinutes(v int32) *GetNisNetworkMetricsRequest {
	s.StepMinutes = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) SetUseCrossAccount(v bool) *GetNisNetworkMetricsRequest {
	s.UseCrossAccount = &v
	return s
}

func (s *GetNisNetworkMetricsRequest) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNisNetworkMetricsRequestDimensions struct {
	// The name of the filter condition.
	//
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// eip-sample*
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkMetricsRequestDimensions) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsRequestDimensions) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsRequestDimensions) GetName() *string {
	return s.Name
}

func (s *GetNisNetworkMetricsRequestDimensions) GetValue() *string {
	return s.Value
}

func (s *GetNisNetworkMetricsRequestDimensions) SetName(v string) *GetNisNetworkMetricsRequestDimensions {
	s.Name = &v
	return s
}

func (s *GetNisNetworkMetricsRequestDimensions) SetValue(v string) *GetNisNetworkMetricsRequestDimensions {
	s.Value = &v
	return s
}

func (s *GetNisNetworkMetricsRequestDimensions) Validate() error {
	return dara.Validate(s)
}
