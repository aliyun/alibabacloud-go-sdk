// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryMonitorValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeHistoryMonitorValuesRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeHistoryMonitorValuesRequest
	GetInstanceId() *string
	SetIntervalForHistory(v string) *DescribeHistoryMonitorValuesRequest
	GetIntervalForHistory() *string
	SetMonitorKeys(v string) *DescribeHistoryMonitorValuesRequest
	GetMonitorKeys() *string
	SetNodeId(v string) *DescribeHistoryMonitorValuesRequest
	GetNodeId() *string
	SetNodeRole(v string) *DescribeHistoryMonitorValuesRequest
	GetNodeRole() *string
	SetOwnerAccount(v string) *DescribeHistoryMonitorValuesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHistoryMonitorValuesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeHistoryMonitorValuesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHistoryMonitorValuesRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeHistoryMonitorValuesRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeHistoryMonitorValuesRequest
	GetStartTime() *string
	SetType(v string) *DescribeHistoryMonitorValuesRequest
	GetType() *string
}

type DescribeHistoryMonitorValuesRequest struct {
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  You can query the monitoring data of the previous month. The maximum time range that you can specify for a query is seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-06T00:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is deprecated. Set the value to `01m`.
	//
	// The **interval at which a query is performed*	- is automatically determined based on the start time and end time of the query. For example, if the query time range is less than or equal to 10 minutes, data is aggregated at a frequency of every 5 seconds and the results are returned at 5-second intervals.
	//
	// > 	- The query result is aligned with the data aggregation frequency. If the specified StartTime value does not coincide with a point in time for data aggregation, the system returns the latest point in time for data aggregation as the first point in time. For example, if you set the StartTime parameter to 2022-01-20T12:01:48Z, the first point in time returned is 2022-01-20T12:01:45Z.
	//
	// > 	- If the number of data shards is greater than or equal to 32, the minimum data aggregation frequency is 1 minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01m
	IntervalForHistory *string `json:"IntervalForHistory,omitempty" xml:"IntervalForHistory,omitempty"`
	// The monitoring metrics. Separate the metrics with commas (,). Take CpuUsage as an example:
	//
	// 	- Cluster or read/write splitting instances
	//
	//     	- To query the overall CPU utilization of all data nodes, specify **CpuUsage$db**.
	//
	//     	- To query the CPU utilization of a single data node, specify **CpuUsage*	- and NodeId.
	//
	// 	- Standard master-replica instances: Specify only **CpuUsage**.
	//
	// For more information about monitoring metrics and their descriptions, see [Additional description of MonitorKeys](https://www.alibabacloud.com/help/zh/redis/developer-reference/api-r-kvstore-2015-01-01-describehistorymonitorvalues-redis#monitorKeys-note).
	//
	// > 	- This parameter is empty by default, which indicates that the UsedMemory and quotaMemory metrics are returned.
	//
	// > 	- To ensure query efficiency, we recommend that you specify no more than five metrics for a single node at a time, and specify only a single metric when you query aggregate metrics.
	//
	// example:
	//
	// memoryUsage
	MonitorKeys *string `json:"MonitorKeys,omitempty" xml:"MonitorKeys,omitempty"`
	// The ID of the node in the instance. You can set this parameter to query the data of a specified node.
	//
	// 	- This parameter is available only for read/write splitting or cluster instances of Tair.
	//
	// 	- You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query node IDs.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0#1679****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// If you want to query the metrics of the read replicas in a cloud-native read/write splitting instance, you must set this parameter to **READONLY*	- and specify **NodeId**.
	//
	// > In other cases, you do not need to specify this parameter or you can set this parameter to **MASTER**.
	//
	// example:
	//
	// READONLY
	NodeRole             *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-06T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeHistoryMonitorValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryMonitorValuesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryMonitorValuesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeHistoryMonitorValuesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHistoryMonitorValuesRequest) GetIntervalForHistory() *string {
	return s.IntervalForHistory
}

func (s *DescribeHistoryMonitorValuesRequest) GetMonitorKeys() *string {
	return s.MonitorKeys
}

func (s *DescribeHistoryMonitorValuesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeHistoryMonitorValuesRequest) GetNodeRole() *string {
	return s.NodeRole
}

func (s *DescribeHistoryMonitorValuesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHistoryMonitorValuesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHistoryMonitorValuesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHistoryMonitorValuesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHistoryMonitorValuesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryMonitorValuesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeHistoryMonitorValuesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeHistoryMonitorValuesRequest) SetEndTime(v string) *DescribeHistoryMonitorValuesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetInstanceId(v string) *DescribeHistoryMonitorValuesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetIntervalForHistory(v string) *DescribeHistoryMonitorValuesRequest {
	s.IntervalForHistory = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetMonitorKeys(v string) *DescribeHistoryMonitorValuesRequest {
	s.MonitorKeys = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetNodeId(v string) *DescribeHistoryMonitorValuesRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetNodeRole(v string) *DescribeHistoryMonitorValuesRequest {
	s.NodeRole = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetOwnerAccount(v string) *DescribeHistoryMonitorValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetOwnerId(v int64) *DescribeHistoryMonitorValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetResourceOwnerAccount(v string) *DescribeHistoryMonitorValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetResourceOwnerId(v int64) *DescribeHistoryMonitorValuesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetSecurityToken(v string) *DescribeHistoryMonitorValuesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetStartTime(v string) *DescribeHistoryMonitorValuesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) SetType(v string) *DescribeHistoryMonitorValuesRequest {
	s.Type = &v
	return s
}

func (s *DescribeHistoryMonitorValuesRequest) Validate() error {
	return dara.Validate(s)
}
