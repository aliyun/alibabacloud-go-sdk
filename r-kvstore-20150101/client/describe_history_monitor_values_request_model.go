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
	// > You can query monitoring data within the past month. The maximum time range to query is 7 days.
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
	// r-bp1zxszhcgatnx******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is deprecated and its value is fixed at `01m`.
	//
	// The system automatically determines the **query interval*	- based on the specified start and end times. For example, if the specified time range is 10 minutes or less, data is aggregated every 5 seconds, and the query results are returned at 5-second intervals.
	//
	// > - If the specified `StartTime` is not at a data aggregation point, the first time point returned by the system is the nearest preceding data aggregation point. For example, if you set StartTime to `2022-01-20T12:01:48Z`, the first time point returned is `2022-01-20T12:01:45Z`.
	//
	// >
	//
	// > - If the instance has 32 or more data shards, the minimum data aggregation frequency is 1 minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01m
	IntervalForHistory *string `json:"IntervalForHistory,omitempty" xml:"IntervalForHistory,omitempty"`
	// The monitoring metric to query, such as `CpuUsage`. To specify multiple metrics, separate them with a comma (,).
	//
	// - For instances that use the cluster or read/write splitting architecture:
	//
	//   - To query the overall CPU utilization of all data nodes, set this parameter to **CpuUsage$db**.
	//
	//   - To query the CPU utilization of a single data node, set this parameter to **CpuUsage*	- and specify the node in the `NodeId` parameter.
	//
	// - For instances that use the standard architecture (primary/standby), set this parameter to **CpuUsage**.
	//
	// For more information about monitoring metrics, see <props="china">[Additional information about the MonitorKeys parameter](https://help.aliyun.com/zh/redis/developer-reference/api-r-kvstore-2015-01-01-describehistorymonitorvalues-redis#monitorKeys-note)<props="intl">[Additional information about the MonitorKeys parameter](https://www.alibabacloud.com/help/zh/redis/developer-reference/api-r-kvstore-2015-01-01-describehistorymonitorvalues-redis#monitorKeys-note) below.
	//
	// > - If you do not specify this parameter, the `UsedMemory` and `quotaMemory` metrics are returned by default.
	//
	// >
	//
	// > - To ensure query efficiency, we recommend that you specify a maximum of 5 monitoring metrics for a single node and a maximum of 1 aggregate monitoring metric per query.
	//
	// example:
	//
	// CpuUsage
	MonitorKeys *string `json:"MonitorKeys,omitempty" xml:"MonitorKeys,omitempty"`
	// The ID of a node in the instance. You can use this parameter to query the monitoring data of a specific node.
	//
	// > - This parameter is available only for instances that use the read/write splitting or cluster architecture.
	//
	// >
	//
	// > - You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query node IDs.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0#1679****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// If you want to query the metrics of a read-only node in a cloud-native instance that uses a read/write splitting architecture, you must specify the **NodeId*	- and set this parameter to **READONLY**.
	//
	// > In all other cases, you do not need to specify this parameter. You can also set it to **MASTER**.
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
