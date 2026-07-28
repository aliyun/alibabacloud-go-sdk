// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisTrafficMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *GetNisTrafficMetricsRequest
	GetBeginTime() *int64
	SetDirection(v string) *GetNisTrafficMetricsRequest
	GetDirection() *string
	SetEndTime(v int64) *GetNisTrafficMetricsRequest
	GetEndTime() *int64
	SetFilter(v []*GetNisTrafficMetricsRequestFilter) *GetNisTrafficMetricsRequest
	GetFilter() []*GetNisTrafficMetricsRequestFilter
	SetMaxResults(v int32) *GetNisTrafficMetricsRequest
	GetMaxResults() *int32
	SetMetricName(v string) *GetNisTrafficMetricsRequest
	GetMetricName() *string
	SetNextToken(v string) *GetNisTrafficMetricsRequest
	GetNextToken() *string
	SetRegionNo(v string) *GetNisTrafficMetricsRequest
	GetRegionNo() *string
	SetScanBy(v string) *GetNisTrafficMetricsRequest
	GetScanBy() *string
	SetStepMinutes(v int32) *GetNisTrafficMetricsRequest
	GetStepMinutes() *int32
	SetStorageInterval(v int32) *GetNisTrafficMetricsRequest
	GetStorageInterval() *int32
	SetTrafficAnalyzerId(v string) *GetNisTrafficMetricsRequest
	GetTrafficAnalyzerId() *string
	SetTrafficScenario(v string) *GetNisTrafficMetricsRequest
	GetTrafficScenario() *string
	SetTupleDimension(v string) *GetNisTrafficMetricsRequest
	GetTupleDimension() *string
}

type GetNisTrafficMetricsRequest struct {
	// The start timestamp, in milliseconds. If not specified, the most recent 1 hour is queried by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The network traffic direction based on Alibaba Cloud resources.
	//
	// In: traffic flowing into the target resource.
	//
	// Out: traffic flowing out of the target resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// In
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end timestamp, in milliseconds. If not specified, the most recent 1 hour is queried by default. If only BeginTime is specified, the 1 hour after BeginTime is queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies additional filter conditions for the traffic to perform focused network traffic analysis.
	Filter []*GetNisTrafficMetricsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// In VPC scenarios, this parameter specifies the paging size. In TR and Internet Shared Bandwidth scenarios, this parameter specifies the SQL query limit. If not specified, the backend defaults to 1440.
	//
	// example:
	//
	// 1440
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The metric name.
	//
	// Common parameters supported in network traffic analysis scenarios:
	//
	//   bps: bits per second.
	//
	//   pps: packets per second.
	//
	// Parameters specific to the Internet scenario:
	//
	//   rtt: round-trip time when establishing a TCP protocol connection.
	//
	//   RetransmitRate: retransmission rate.
	//
	// Parameters specific to the area-level bandwidth scenario:
	//
	//   RatelimitDropPps: rate of packet loss due to rate limiting.
	//
	//   BandwidthUtilization: bandwidth utilization.
	//
	// Parameters specific to the NAT scenario:
	//
	//   ActiveSessionCount: number of concurrent sessions.
	//
	//   NewSessionPerSecond: number of new sessions per second.
	//
	// This parameter is required.
	//
	// example:
	//
	// bps
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The token for the next query. You do not need to specify this parameter for the first query or when no more results exist. If a next page exists, set this parameter to the NextToken value returned by the previous API invoke. This parameter is valid only in VPC scenarios. TR and Internet Shared Bandwidth scenarios do not use this parameter.
	//
	// example:
	//
	// f7zUd3gArYj/xjPttJo5L5dK0R+gSbfHElLqi8C2IPWMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The sort order. Valid values:
	//
	// TimestampAscending: sorts by time in ascending order.
	//
	// TimestampDescending: sorts by time in descending order.
	//
	// example:
	//
	// TimestampAscending
	ScanBy *string `json:"ScanBy,omitempty" xml:"ScanBy,omitempty"`
	// The aggregation step for time series data, in minutes. The final query granularity is the larger value between StepMinutes and the underlying storage granularity. The number of data points calculated by (EndTime-BeginTime)/StepMinutes cannot exceed 1440.
	//
	// example:
	//
	// 10
	StepMinutes *int32 `json:"StepMinutes,omitempty" xml:"StepMinutes,omitempty"`
	// The storage bucket precision property.
	//
	// The storage bucket precision specifies the storage aggregation epoch to query. Two precision levels are supported: high precision (such as 1 minute) or long epoch (such as 1 day). The specific precision is determined by the network traffic analysis sampling interval configured for high-precision traffic statistics or long-epoch traffic statistics when creating or editing the network traffic analysis analyzer.
	//
	// - The storage precisions active for the corresponding tuples of the network traffic analysis analyzer are:
	//
	//   - `1`: in minutes (1 minute)
	//
	//   - `10`: in minutes (10 minutes)
	//
	//   - `60`: in minutes (60 minutes, i.e., 1 hour)
	//
	//   - `1440`: in minutes (1440 minutes, i.e., 1 day)
	//
	// - The storage bucket precision can be used for two typical purposes:
	//
	//   - High-precision traffic statistics: such as 1-minute, 10-minute, or 60-minute aggregation
	//
	//   - Long-epoch traffic statistics: such as 1440-minute (1-day) aggregation
	//
	// - Specify a value for this field during the query to select the storage aggregation epoch. For example:
	//
	//   - Pass `10`: queries short-epoch data with a 10-minute aggregation granularity
	//
	//   - Pass `1440`: queries long-epoch data with a 1-day aggregation granularity
	//
	// example:
	//
	// 10
	StorageInterval *int32 `json:"StorageInterval,omitempty" xml:"StorageInterval,omitempty"`
	// The ID of the network traffic analysis analyzer.
	//
	// This parameter is required.
	//
	// example:
	//
	// nta-e093cb80c7c047afbd1d
	TrafficAnalyzerId *string `json:"TrafficAnalyzerId,omitempty" xml:"TrafficAnalyzerId,omitempty"`
	// The supported analysis scenarios:
	//
	// - All VPC flow log analysis
	//
	// - Internet VPC flow log analysis
	//
	// - All TR flow log analysis
	//
	// - Internet Shared Bandwidth metric analysis
	//
	// This parameter is required.
	//
	// example:
	//
	// VpcFlowLogAll
	TrafficScenario *string `json:"TrafficScenario,omitempty" xml:"TrafficScenario,omitempty"`
	// The traffic storage aggregation dimension.
	//
	// Based on the TrafficScenario:
	//
	// - VpcFlowLogAll/VpcFlowLog: required. Specifies the storage aggregation view to query, which corresponds to the storage aggregation property configured in the network traffic analysis analyzer.
	//
	// - TRFlowLog/CbwpMetric: optional. Automatically adapts based on the storage aggregation property of the network traffic analysis analyzer.
	//
	// example:
	//
	// Tuple2
	TupleDimension *string `json:"TupleDimension,omitempty" xml:"TupleDimension,omitempty"`
}

func (s GetNisTrafficMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisTrafficMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetNisTrafficMetricsRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisTrafficMetricsRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetNisTrafficMetricsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisTrafficMetricsRequest) GetFilter() []*GetNisTrafficMetricsRequestFilter {
	return s.Filter
}

func (s *GetNisTrafficMetricsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetNisTrafficMetricsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetNisTrafficMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetNisTrafficMetricsRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisTrafficMetricsRequest) GetScanBy() *string {
	return s.ScanBy
}

func (s *GetNisTrafficMetricsRequest) GetStepMinutes() *int32 {
	return s.StepMinutes
}

func (s *GetNisTrafficMetricsRequest) GetStorageInterval() *int32 {
	return s.StorageInterval
}

func (s *GetNisTrafficMetricsRequest) GetTrafficAnalyzerId() *string {
	return s.TrafficAnalyzerId
}

func (s *GetNisTrafficMetricsRequest) GetTrafficScenario() *string {
	return s.TrafficScenario
}

func (s *GetNisTrafficMetricsRequest) GetTupleDimension() *string {
	return s.TupleDimension
}

func (s *GetNisTrafficMetricsRequest) SetBeginTime(v int64) *GetNisTrafficMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetDirection(v string) *GetNisTrafficMetricsRequest {
	s.Direction = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetEndTime(v int64) *GetNisTrafficMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetFilter(v []*GetNisTrafficMetricsRequestFilter) *GetNisTrafficMetricsRequest {
	s.Filter = v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetMaxResults(v int32) *GetNisTrafficMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetMetricName(v string) *GetNisTrafficMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetNextToken(v string) *GetNisTrafficMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetRegionNo(v string) *GetNisTrafficMetricsRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetScanBy(v string) *GetNisTrafficMetricsRequest {
	s.ScanBy = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetStepMinutes(v int32) *GetNisTrafficMetricsRequest {
	s.StepMinutes = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetStorageInterval(v int32) *GetNisTrafficMetricsRequest {
	s.StorageInterval = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetTrafficAnalyzerId(v string) *GetNisTrafficMetricsRequest {
	s.TrafficAnalyzerId = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetTrafficScenario(v string) *GetNisTrafficMetricsRequest {
	s.TrafficScenario = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) SetTupleDimension(v string) *GetNisTrafficMetricsRequest {
	s.TupleDimension = &v
	return s
}

func (s *GetNisTrafficMetricsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNisTrafficMetricsRequestFilter struct {
	// Based on the TupleDimension field and TrafficScenario field, the supported filter condition label keys are as follows:
	//
	// - `TrafficScenario = VpcFlowLogAll` / `VpcFlowLogInternet` (VPC flow log scenario):
	//
	//   - When `TupleDimension` is a 1-tuple, the following keys are supported:
	//
	//     - `FlowAction`: the action type to execute on traffic after it matches the corresponding rule or policy (required, the corresponding value does not support multiple selections)
	//
	//     - `VpcId`: VPC ID (the corresponding value supports multiple selections)
	//
	//     - `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
	//
	//     - `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
	//
	//     - `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
	//
	//     - `CloudIp`: cloud IP address (the corresponding value supports multiple selections)
	//
	//   - When `TupleDimension` is a 2-tuple, the following keys are supported:
	//
	//     - `FlowAction`: the action type to execute on traffic after it matches the corresponding rule or policy (required, the corresponding value does not support multiple selections)
	//
	//     - `VpcId`: VPC ID (the corresponding value supports multiple selections)
	//
	//     - `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
	//
	//     - `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
	//
	//     - `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
	//
	//     - `SourceIp`: source IP address (the corresponding value supports multiple selections)
	//
	//     - `DestinationIp`: destination IP address (the corresponding value supports multiple selections)
	//
	//     - `TrafficPath`: traffic path (the corresponding value supports multiple selections)
	//
	//   - When `TupleDimension` is a 5-tuple, the following keys are supported:
	//
	//     - `FlowAction`: the action type to execute on traffic after it matches the corresponding rule or policy (required, the corresponding value does not support multiple selections)
	//
	//     - `VpcId`: VPC ID (the corresponding value supports multiple selections)
	//
	//     - `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
	//
	//     - `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
	//
	//     - `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
	//
	//     - `SourceIp`: source IP address
	//
	//     - `DestinationIp`: destination IP address
	//
	//     - `TrafficPath`: traffic path (the corresponding value supports multiple selections)
	//
	//     - `SourcePort`: source port (the corresponding value supports multiple selections)
	//
	//     - `DestinationPort`: destination port (the corresponding value supports multiple selections)
	//
	//     - `Protocol`: network protocol (the corresponding value supports multiple selections)
	//
	//   - In the VPC Internet scenario (`TrafficScenario = VpcFlowLogInternet`), the following additional keys are supported for filtering by Internet location:
	//
	//     - `ClientCountry`: filters network traffic analysis scope by country (the corresponding value supports multiple selections)
	//
	//     - `ClientCity`: filters network traffic analysis scope by city (the corresponding value supports multiple selections)
	//
	//     - `ClientAsn`: filters network traffic analysis scope by ASN (the corresponding value supports multiple selections)
	//
	//     - `ClientIsp`: filters network traffic analysis scope by client ISP (the corresponding value supports multiple selections)
	//
	//   - In VPC scenarios, the following traffic metrics filters are supported:
	//
	//     - `MinBytes`: specifies the minimum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
	//
	//     - `MaxBytes`: specifies the maximum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
	//
	//     - `MinRoundTripTime`: specifies the minimum RTT for sorting, in ms (the corresponding value does not support multiple selections)
	//
	//     - `MaxRoundTripTime`: specifies the maximum RTT for sorting, in ms (the corresponding value does not support multiple selections)
	//
	//     - `MinPackages`: specifies the minimum number of packets for sorting (the corresponding value does not support multiple selections)
	//
	//     - `MaxPackages`: specifies the maximum number of packets for sorting (the corresponding value does not support multiple selections)
	//
	// ---
	//
	// - `TrafficScenario = TRFlowlog` (TR flow log scenario):
	//
	//   - When querying 2-tuples or adaptively using 2-tuples, the following keys are supported:
	//
	//     - `TransitRouterAttachmentId`: network instance connection ID (required, the corresponding value does not support multiple selections)
	//
	//     - `TransitRouterPairAttachmentId`: peer TR connection ID (the corresponding value supports multiple selections)
	//
	//     - `TransitRouterId`: transit router instance ID (the corresponding value supports multiple selections)
	//
	//     - `SourceIp`: source IP address (the corresponding value does not support multiple selections when Operator is like, and supports multiple selections when Operator is not like)
	//
	//     - `DestinationIp`: destination IP address (the corresponding value does not support multiple selections when Operator is like, and supports multiple selections when Operator is not like)
	//
	//     - `Dscp`: Differentiated Services Code Point (the corresponding value supports multiple selections)
	//
	//   - When querying 5-tuples or adaptively using 5-tuples, the following additional keys are supported in addition to the 2-tuple keys:
	//
	//     - `Protocol`: network protocol (the corresponding value supports multiple selections)
	//
	//     - `SourcePort`: source port (the corresponding value supports multiple selections)
	//
	//     - `DestinationPort`: destination port (the corresponding value supports multiple selections)
	//
	//   - In `non-TR cross-region scenarios`, the following additional keys are supported:
	//
	//     - `TransitRouterSourceResourceId`: source network instance ID (the corresponding value supports multiple selections)
	//
	//     - `TransitRouterDestinationResourceId`: destination network instance ID (the corresponding value supports multiple selections)
	//
	//   - In `VPC connection traffic scenarios`, the following additional keys are supported:
	//
	//     - `TransitRouterSourceNetworkInterface`: source TR network interface controller (NIC) (the corresponding value supports multiple selections)
	//
	//     - `TransitRouterDestinationNetworkInterface`: destination TR network interface controller (NIC) (the corresponding value supports multiple selections)
	//
	//   - In TR scenarios, the following traffic metrics filters are supported:
	//
	//     - `MinBytes`: specifies the minimum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
	//
	//     - `MaxBytes`: specifies the maximum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
	//
	//     - `MinPackages`: specifies the minimum number of packets for sorting (the corresponding value does not support multiple selections)
	//
	//     - `MaxPackages`: specifies the maximum number of packets for sorting (the corresponding value does not support multiple selections)
	//
	//     - `MinPacketsLostNoRoute`: minimum number of packets dropped due to no route (the corresponding value does not support multiple selections)
	//
	//     - `MinPacketsLostBlackhole`: minimum number of packets dropped due to blackhole route (the corresponding value does not support multiple selections)
	//
	//     - `MinPacketsLostTTLExpired`: minimum number of packets dropped due to TTL timeout (the corresponding value does not support multiple selections)
	//
	//     - `MaxPacketsLostNoRoute`: maximum number of packets dropped due to no route (the corresponding value does not support multiple selections)
	//
	//     - `MaxPacketsLostBlackhole`: maximum number of packets dropped due to blackhole route (the corresponding value does not support multiple selections)
	//
	//     - `MaxPacketsLostTTLExpired`: maximum number of packets dropped due to TTL timeout (the corresponding value does not support multiple selections)
	//
	// ---
	//
	// - `TrafficScenario = CbwpMetric` (Internet Shared Bandwidth metric analysis scenario):
	//
	//   - The following filter condition keys are supported:
	//
	//     - `PublicIpAddress`: the public IP address of the associated EIP (the corresponding value does not support multiple selections when Operator is like, and supports multiple selections when Operator is not like)
	//
	//     - `BindingResourceType`: the type of the instance resource to which the EIP is bound (the corresponding value supports multiple selections)
	//
	//     - `BindingResourceId`: the ID of the instance resource to which the EIP is bound (the corresponding value supports multiple selections)
	//
	//     - `CbwpId`: the Internet Shared Bandwidth instance ID (required, the corresponding value does not support multiple selections)
	//
	//     - `InstanceId`: the EIP ID bound to the Internet Shared Bandwidth instance (the corresponding value supports multiple selections)
	//
	//   - In CBWP scenarios, the following traffic metrics filters are supported:
	//
	//     - `MinBytes`: specifies the minimum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
	//
	//     - `MaxBytes`: specifies the maximum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
	//
	//     - `MinPackages`: specifies the minimum number of packets for sorting (the corresponding value does not support multiple selections)
	//
	//     - `MaxPackages`: specifies the maximum number of packets for sorting (the corresponding value does not support multiple selections)
	//
	// example:
	//
	// NetworkInterfaceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter operator.
	//
	// - TR and Internet Shared Bandwidth scenarios:
	//
	//   - Defaults to in if not specified.
	//
	//   - like performs prefix matching and only one Value can be specified.
	//
	// - VPC scenarios currently ignore this parameter and uniformly process it as IN.
	//
	// example:
	//
	// in
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The filter value corresponding to the specified key type.
	//
	// Based on the `TupleDimension` field and `TrafficScenario` field, the supported values are as follows:
	//
	// - `TrafficScenario = VpcFlowLogAll` / `VpcFlowLogInternet` (VPC flow log scenario)
	//
	//   - When the key is `FlowAction`, the valid values are:
	//
	//     - `ACCEPT` (pass `Accept` by default): traffic allowed by security groups and network ACLs
	//
	//     - `REJECT`: traffic denied by security groups and network ACLs
	//
	// - `TrafficScenario = TRFlowlog` (TR flow log scenario)
	//
	//   - When the key is `TransitRouterAttachmentId`, this is a required field, and the corresponding value is also required (specify the specific VPC connection / VPN connection / VBR connection / ECR connection / inter-region connection or network instance connection ID).
	//
	// - `TrafficScenario = CbwpMetric` (shared bandwidth metric analysis scenario)
	//
	//   - When the key is `CbwpId`, this is a required field, and the corresponding value is also required (specify the specific Internet Shared Bandwidth instance ID).
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetNisTrafficMetricsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetNisTrafficMetricsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetNisTrafficMetricsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetNisTrafficMetricsRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *GetNisTrafficMetricsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *GetNisTrafficMetricsRequestFilter) SetKey(v string) *GetNisTrafficMetricsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetNisTrafficMetricsRequestFilter) SetOperator(v string) *GetNisTrafficMetricsRequestFilter {
	s.Operator = &v
	return s
}

func (s *GetNisTrafficMetricsRequestFilter) SetValue(v []*string) *GetNisTrafficMetricsRequestFilter {
	s.Value = v
	return s
}

func (s *GetNisTrafficMetricsRequestFilter) Validate() error {
	return dara.Validate(s)
}
