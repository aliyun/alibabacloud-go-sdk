// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisTrafficRankingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *StartNisTrafficRankingRequest
	GetBeginTime() *int64
	SetDirection(v string) *StartNisTrafficRankingRequest
	GetDirection() *string
	SetEndTime(v int64) *StartNisTrafficRankingRequest
	GetEndTime() *int64
	SetFilter(v []*StartNisTrafficRankingRequestFilter) *StartNisTrafficRankingRequest
	GetFilter() []*StartNisTrafficRankingRequestFilter
	SetGroupBy(v []*string) *StartNisTrafficRankingRequest
	GetGroupBy() []*string
	SetLanguage(v string) *StartNisTrafficRankingRequest
	GetLanguage() *string
	SetMaxResults(v int32) *StartNisTrafficRankingRequest
	GetMaxResults() *int32
	SetNextToken(v string) *StartNisTrafficRankingRequest
	GetNextToken() *string
	SetOrderBy(v string) *StartNisTrafficRankingRequest
	GetOrderBy() *string
	SetRegionNo(v string) *StartNisTrafficRankingRequest
	GetRegionNo() *string
	SetSort(v string) *StartNisTrafficRankingRequest
	GetSort() *string
	SetStorageInterval(v int32) *StartNisTrafficRankingRequest
	GetStorageInterval() *int32
	SetTopN(v int32) *StartNisTrafficRankingRequest
	GetTopN() *int32
	SetTrafficAnalyzerId(v string) *StartNisTrafficRankingRequest
	GetTrafficAnalyzerId() *string
	SetTrafficScenario(v string) *StartNisTrafficRankingRequest
	GetTrafficScenario() *string
	SetTupleDimension(v string) *StartNisTrafficRankingRequest
	GetTupleDimension() *string
}

type StartNisTrafficRankingRequest struct {
	// The start timestamp of the query, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The network traffic direction based on Alibaba Cloud resources.
	//
	// In: Traffic flowing into the target resource.
	//
	// Out: Traffic flowing out of the target resource.
	//
	// - VPC flow log scenario (`TraffficScenario = VpcFlowLogAll` / `VpcFlowLogInternet`):
	//
	//   - In: Traffic flowing into the ENI.
	//
	//   - Out: Traffic flowing out of the ENI.
	//
	// - TR flow log scenario (`TraffficScenario = TRFlowlog`):
	//
	//   - In: Traffic flowing into the TR.
	//
	//   - Out: Traffic flowing out of the TR.
	//
	// - Internet Shared Bandwidth metric analysis scenario (`TraffficScenario = CbwpMetric`):
	//
	//   - In: Traffic flowing into the EIP.
	//
	//   - Out: Traffic flowing out of the EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// Out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end timestamp of the query, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies additional filter conditions for focused network traffic analysis.
	Filter []*StartNisTrafficRankingRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Specifies multiple traffic dimensions for aggregation and sorting.
	GroupBy []*string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty" type:"Repeated"`
	// The language. Valid values: zh-CN, en-US.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The page size. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or when no more results are available. If a next query exists, set this value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Based on the `TrafficScenario` field, the following metrics are supported for ranking traffic:
	//
	// - `TrafficScenario = VpcFlowLogAll` / `VpcFlowLogInternet` (VPC flow log scenario):
	//
	//   - `Bytes`: Bandwidth
	//
	//   - `Packets`: Packets
	//
	//   - `RoundTripTime`: TCP RTT
	//
	// - `TrafficScenario = TRFlowlog` (TR flow log scenario):
	//
	//   - `Bytes`: Bandwidth
	//
	//   - `Packets`: Packets
	//
	//   - `PacketsLostNoRoute`: Packet loss due to no routing
	//
	//   - `PacketsLostBlackhole`: Packet loss due to blackhole routing
	//
	//   - `PacketsLostTTLExpired`: Packet loss due to TTL timeout
	//
	//   - `BytesIncrease`: Bandwidth increase
	//
	//   - `BytesIncreaseRatio`: Bandwidth increase ratio
	//
	// - `TrafficScenario = CbwpMetric` (Internet Shared Bandwidth metric analysis scenario):
	//
	//   - `Bytes`: Bandwidth
	//
	//   - `Packets`: Packets
	//
	// This parameter is required.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The region where the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The sorting method for network traffic analysis. Valid values:
	//
	// - ASC: Sorts in ascending order.
	//
	// - DESC: Sorts in descending order.
	//
	// example:
	//
	// Desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The storage bucket precision property.
	//
	// The storage bucket precision specifies the storage aggregation epoch to query. Two precision levels are supported: high precision (such as 1 minute) and long epoch (such as 1 day). The specific precision is determined by the network traffic analysis sampling interval configured for high-precision traffic statistics or long-epoch traffic statistics when creating or editing the network traffic analysis instance.
	//
	// - The storage precision supported by the corresponding tuple of the network traffic analysis instance:
	//
	//   - `1`: In minutes (1 minute)
	//
	//   - `10`: In minutes (10 minutes)
	//
	//   - `60`: In minutes (60 minutes, or 1 hour)
	//
	//   - `1440`: In minutes (1440 minutes, or 1 day)
	//
	// - The storage bucket precision can be used for two typical purposes:
	//
	//   - High-precision traffic statistics: Aggregation at 1-minute, 10-minute, or 60-minute intervals.
	//
	//   - Long-epoch traffic statistics: Aggregation at 1440-minute (1-day) intervals.
	//
	// - Pass a value for this field during the query to specify the storage aggregation epoch. For example:
	//
	//   - Pass `10`: Queries short-epoch data aggregated at 10-minute granularity.
	//
	//   - Pass `1440`: Queries long-epoch data aggregated at 1-day granularity.
	//
	// Note: The active storage precision values depend on the configuration of the network traffic analysis instance.
	//
	// example:
	//
	// 10
	StorageInterval *int32 `json:"StorageInterval,omitempty" xml:"StorageInterval,omitempty"`
	// The number of entries for the network traffic analysis sorting query.
	//
	// You can specify a custom number. If this field is not specified, all traffic data that meets the specified conditions is sorted and analyzed within the performance capacity of the network traffic analysis feature.
	//
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// The ID of the network traffic analysis instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// nta-262****ca07f
	TrafficAnalyzerId *string `json:"TrafficAnalyzerId,omitempty" xml:"TrafficAnalyzerId,omitempty"`
	// Supported analysis scenarios:
	//
	// - All VPC flow log analysis
	//
	// - Public VPC flow log analysis
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
	// The storage aggregation dimension of the network traffic analysis instance.
	//
	// Based on the TraffficScenario:
	//
	// - VpcFlowLogAll/VpcFlowLog: Required. Specifies the storage aggregation view to query, which corresponds to the storage aggregation property configured in the network traffic analysis instance.
	//
	// - TRFlowLog/CbwpMetric: Optional. Automatically adapts based on the storage aggregation property of the network traffic analysis instance.
	//
	// example:
	//
	// Tuple1
	TupleDimension *string `json:"TupleDimension,omitempty" xml:"TupleDimension,omitempty"`
}

func (s StartNisTrafficRankingRequest) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingRequest) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *StartNisTrafficRankingRequest) GetDirection() *string {
	return s.Direction
}

func (s *StartNisTrafficRankingRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *StartNisTrafficRankingRequest) GetFilter() []*StartNisTrafficRankingRequestFilter {
	return s.Filter
}

func (s *StartNisTrafficRankingRequest) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *StartNisTrafficRankingRequest) GetLanguage() *string {
	return s.Language
}

func (s *StartNisTrafficRankingRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *StartNisTrafficRankingRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *StartNisTrafficRankingRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *StartNisTrafficRankingRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *StartNisTrafficRankingRequest) GetSort() *string {
	return s.Sort
}

func (s *StartNisTrafficRankingRequest) GetStorageInterval() *int32 {
	return s.StorageInterval
}

func (s *StartNisTrafficRankingRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *StartNisTrafficRankingRequest) GetTrafficAnalyzerId() *string {
	return s.TrafficAnalyzerId
}

func (s *StartNisTrafficRankingRequest) GetTrafficScenario() *string {
	return s.TrafficScenario
}

func (s *StartNisTrafficRankingRequest) GetTupleDimension() *string {
	return s.TupleDimension
}

func (s *StartNisTrafficRankingRequest) SetBeginTime(v int64) *StartNisTrafficRankingRequest {
	s.BeginTime = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetDirection(v string) *StartNisTrafficRankingRequest {
	s.Direction = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetEndTime(v int64) *StartNisTrafficRankingRequest {
	s.EndTime = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetFilter(v []*StartNisTrafficRankingRequestFilter) *StartNisTrafficRankingRequest {
	s.Filter = v
	return s
}

func (s *StartNisTrafficRankingRequest) SetGroupBy(v []*string) *StartNisTrafficRankingRequest {
	s.GroupBy = v
	return s
}

func (s *StartNisTrafficRankingRequest) SetLanguage(v string) *StartNisTrafficRankingRequest {
	s.Language = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetMaxResults(v int32) *StartNisTrafficRankingRequest {
	s.MaxResults = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetNextToken(v string) *StartNisTrafficRankingRequest {
	s.NextToken = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetOrderBy(v string) *StartNisTrafficRankingRequest {
	s.OrderBy = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetRegionNo(v string) *StartNisTrafficRankingRequest {
	s.RegionNo = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetSort(v string) *StartNisTrafficRankingRequest {
	s.Sort = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetStorageInterval(v int32) *StartNisTrafficRankingRequest {
	s.StorageInterval = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTopN(v int32) *StartNisTrafficRankingRequest {
	s.TopN = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTrafficAnalyzerId(v string) *StartNisTrafficRankingRequest {
	s.TrafficAnalyzerId = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTrafficScenario(v string) *StartNisTrafficRankingRequest {
	s.TrafficScenario = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTupleDimension(v string) *StartNisTrafficRankingRequest {
	s.TupleDimension = &v
	return s
}

func (s *StartNisTrafficRankingRequest) Validate() error {
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

type StartNisTrafficRankingRequestFilter struct {
	// Based on the `TupleDimension` and `TrafficScenario` fields, the following filter condition label keys are supported:
	//
	// - `TrafficScenario = VpcFlowLogAll` / `VpcFlowLogInternet` (VPC flow log scenario):
	//
	//   - When `TupleDimension` is 1-tuple, the following keys are supported:
	//
	//     - `FlowAction`: The action type to execute on traffic after it matches a rule or policy (required, corresponding value does not support multiple selections)
	//
	//     - `VpcId`: VPC ID (corresponding value supports multiple selections)
	//
	//     - `VSwitchId`: vSwitch ID (corresponding value supports multiple selections)
	//
	//     - `NetworkInterfaceId`: Network interface controller (NIC) ID (corresponding value supports multiple selections)
	//
	//     - `EcsId`: ECS server ID (corresponding value supports multiple selections)
	//
	//     - `CloudIp`: Cloud IP address (corresponding value supports multiple selections)
	//
	//   - When `TupleDimension` is 2-tuple, the following keys are supported:
	//
	//     - `FlowAction`: The action type to execute on traffic after it matches a rule or policy (required, corresponding value does not support multiple selections)
	//
	//     - `VpcId`: VPC ID (corresponding value supports multiple selections)
	//
	//     - `VSwitchId`: vSwitch ID (corresponding value supports multiple selections)
	//
	//     - `NetworkInterfaceId`: Network interface controller (NIC) ID (corresponding value supports multiple selections)
	//
	//     - `EcsId`: ECS server ID (corresponding value supports multiple selections)
	//
	//     - `SourceIp`: Source IP address (corresponding value supports multiple selections)
	//
	//     - `DestinationIp`: Destination IP address (corresponding value supports multiple selections)
	//
	//     - `TrafficPath`: Traffic path (corresponding value supports multiple selections)
	//
	//   - When `TupleDimension` is 5-tuple, the following keys are supported:
	//
	//     - `FlowAction`: The action type to execute on traffic after it matches a rule or policy (required, corresponding value does not support multiple selections)
	//
	//     - `VpcId`: VPC ID (corresponding value supports multiple selections)
	//
	//     - `VSwitchId`: vSwitch ID (corresponding value supports multiple selections)
	//
	//     - `NetworkInterfaceId`: Network interface controller (NIC) ID (corresponding value supports multiple selections)
	//
	//     - `EcsId`: ECS server ID (corresponding value supports multiple selections)
	//
	//     - `SourceIp`: Source IP address
	//
	//     - `DestinationIp`: Destination IP address
	//
	//     - `TrafficPath`: Traffic path (corresponding value supports multiple selections)
	//
	//     - `SourcePort`: Source port (corresponding value supports multiple selections)
	//
	//     - `DestinationPort`: Destination port (corresponding value supports multiple selections)
	//
	//     - `Protocol`: Network protocol (corresponding value supports multiple selections)
	//
	//   - For VPC public network scenarios (`TrafficScenario = VpcFlowLogInternet`), the following additional keys are supported for filtering by Internet location:
	//
	//     - `ClientCountry`: Filter network traffic analysis scope by country (corresponding value supports multiple selections)
	//
	//     - `ClientCity`: Filter network traffic analysis scope by city (corresponding value supports multiple selections)
	//
	//     - `ClientAsn`: Filter network traffic analysis scope by ASN (corresponding value supports multiple selections)
	//
	//     - `ClientIsp`: Filter network traffic analysis scope by client ISP (corresponding value supports multiple selections)
	//
	//   - For all VPC scenarios, filtering by traffic metrics is supported:
	//
	//     - `MinBytes`: Specifies the minimum traffic volume for sorting, in bytes (corresponding value does not support multiple selections)
	//
	//     - `MaxBytes`: Specifies the maximum traffic volume for sorting, in bytes (corresponding value does not support multiple selections)
	//
	//     - `MinRoundTripTime`: Specifies the minimum RTT for sorting, in ms (corresponding value does not support multiple selections)
	//
	//     - `MaxRoundTripTime`: Specifies the maximum RTT for sorting, in ms (corresponding value does not support multiple selections)
	//
	//     - `MinPackages`: Specifies the minimum number of packets for sorting (corresponding value does not support multiple selections)
	//
	//     - `MaxPackages`: Specifies the maximum number of packets for sorting (corresponding value does not support multiple selections)
	//
	// ---
	//
	// - `TrafficScenario = TRFlowlog` (TR flow log scenario):
	//
	//   - When querying 2-tuple or adaptive 2-tuple, the following keys are supported:
	//
	//     - `TransitRouterAttachmentId`: Network instance connection ID (required, corresponding value does not support multiple selections)
	//
	//     - `TransitRouterPairAttachmentId`: Peer TR connection ID (corresponding value supports multiple selections)
	//
	//     - `TransitRouterId`: Forward router instance ID (corresponding value supports multiple selections)
	//
	//     - `SourceIp`: Source IP address (corresponding value does not support multiple selections when Operator = like. Corresponding value supports multiple selections when Operator != like)
	//
	//     - `DestinationIp`: Destination IP address (corresponding value does not support multiple selections when Operator = like. Corresponding value supports multiple selections when Operator != like)
	//
	//     - `Dscp`: Differentiated Services Code Point (corresponding value supports multiple selections)
	//
	//   - When querying 5-tuple or adaptive 5-tuple, the following additional keys are supported on top of 2-tuple:
	//
	//     - `Protocol`: Network protocol (corresponding value supports multiple selections)
	//
	//     - `SourcePort`: Source port (corresponding value supports multiple selections)
	//
	//     - `DestinationPort`: Destination port (corresponding value supports multiple selections)
	//
	//   - In `non-TR cross-region scenarios`, the following additional keys are supported:
	//
	//     - `TransitRouterSourceResourceId`: Source network instance ID (corresponding value supports multiple selections)
	//
	//     - `TransitRouterDestinationResourceId`: Destination network instance ID (corresponding value supports multiple selections)
	//
	//   - In `VPC connection traffic scenarios`, the following additional keys are supported:
	//
	//     - `TransitRouterSourceNetworkInterface`: Source TR ENI (corresponding value supports multiple selections)
	//
	//     - `TransitRouterDestinationNetworkInterface`: Destination TR ENI (corresponding value supports multiple selections)
	//
	//   - For all TR scenarios, filtering by traffic metrics is supported:
	//
	//     - `MinBytes`: Specifies the minimum traffic volume for sorting, in bytes (corresponding value does not support multiple selections)
	//
	//     - `MaxBytes`: Specifies the maximum traffic volume for sorting, in bytes (corresponding value does not support multiple selections)
	//
	//     - `MinPackages`: Specifies the minimum number of packets for sorting (corresponding value does not support multiple selections)
	//
	//     - `MaxPackages`: Specifies the maximum number of packets for sorting (corresponding value does not support multiple selections)
	//
	//     - `MinPacketsLostNoRoute`: Minimum packet loss due to no routing (corresponding value does not support multiple selections)
	//
	//     - `MinPacketsLostBlackhole`: Minimum packet loss due to blackhole routing (corresponding value does not support multiple selections)
	//
	//     - `MinPacketsLostTTLExpired`: Minimum packet loss due to TTL timeout (corresponding value does not support multiple selections)
	//
	//     - `MaxPacketsLostNoRoute`: Maximum packet loss due to no routing (corresponding value does not support multiple selections)
	//
	//     - `MaxPacketsLostBlackhole`: Maximum packet loss due to blackhole routing (corresponding value does not support multiple selections)
	//
	//     - `MaxPacketsLostTTLExpired`: Maximum packet loss due to TTL timeout (corresponding value does not support multiple selections)
	//
	// ---
	//
	// - `TrafficScenario = CbwpMetric` (Internet Shared Bandwidth metric analysis scenario):
	//
	//   - Filtering by conditions supports:
	//
	//     - `PublicIpAddress`: Public IP address of the bound EIP (corresponding value does not support multiple selections when Operator = like. Corresponding value supports multiple selections when Operator != like)
	//
	//     - `BindingResourceType`: Resource type of the instance bound to the EIP (corresponding value supports multiple selections)
	//
	//     - `BindingResourceId`: Resource ID of the instance bound to the EIP (corresponding value supports multiple selections)
	//
	//     - `CbwpId`: Internet Shared Bandwidth ID (required, corresponding value does not support multiple selections)
	//
	//     - `InstanceId`: EIP ID bound to the Internet Shared Bandwidth instance (corresponding value supports multiple selections)
	//
	//   - For all CBWP scenarios, filtering by traffic metrics is supported:
	//
	//     - `MinBytes`: Specifies the minimum traffic volume for sorting, in bytes (corresponding value does not support multiple selections)
	//
	//     - `MaxBytes`: Specifies the maximum traffic volume for sorting, in bytes (corresponding value does not support multiple selections)
	//
	//     - `MinPackages`: Specifies the minimum number of packets for sorting (corresponding value does not support multiple selections)
	//
	//     - `MaxPackages`: Specifies the maximum number of packets for sorting (corresponding value does not support multiple selections)
	//
	// example:
	//
	// FlowAction
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// For specified key types, some support using operators to perform string matching on the passed value. Valid values (default value: `in`):
	//
	// - `in`: Equal to.
	//
	// - `not in`: Not equal to.
	//
	// - `like`: Contains.
	//
	// Based on the `TupleDimension` and `TrafficScenario` fields, `like` is supported as follows:
	//
	// - `TrafficScenario = VpcFlowLogAll` / `VpcFlowLogInternet` (VPC flow log scenario):
	//
	//   - The `like` operator is supported when the key is one of the following:
	//
	//     - `CloudIp`
	//
	//     - `SourceIp`
	//
	//     - `DestinationIp`
	//
	// - `TrafficScenario = TRFlowlog` (TR flow log scenario):
	//
	//   - The `like` operator is supported when the key is one of the following:
	//
	//     - `SourceIp`
	//
	//     - `DestinationIp`
	//
	// - `TrafficScenario = CbwpMetric` (Internet Shared Bandwidth metric analysis scenario):
	//
	//   - The `like` operator is supported when the key is one of the following:
	//
	//     - `PublicIpAddress`
	//
	// For all other fields, only the `in` and `not in` operators are supported.
	//
	// example:
	//
	// in
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s StartNisTrafficRankingRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingRequestFilter) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingRequestFilter) GetKey() *string {
	return s.Key
}

func (s *StartNisTrafficRankingRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *StartNisTrafficRankingRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *StartNisTrafficRankingRequestFilter) SetKey(v string) *StartNisTrafficRankingRequestFilter {
	s.Key = &v
	return s
}

func (s *StartNisTrafficRankingRequestFilter) SetOperator(v string) *StartNisTrafficRankingRequestFilter {
	s.Operator = &v
	return s
}

func (s *StartNisTrafficRankingRequestFilter) SetValue(v []*string) *StartNisTrafficRankingRequestFilter {
	s.Value = v
	return s
}

func (s *StartNisTrafficRankingRequestFilter) Validate() error {
	return dara.Validate(s)
}
