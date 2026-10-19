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
	// The end time of the query for historical monitoring data. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// > You can query monitoring data of the last month. The maximum query range cannot exceed 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-06T00:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. The value is fixed to `01m`.
	//
	// The system automatically determines the **query interval*	- based on the start time and end time. For example, if the query time range is less than or equal to 10 minutes, the data aggregation frequency is 5s, and the query results are returned at a frequency of 5s.
	//
	// > 	- The query results are aligned with the data aggregation frequency of the instance. If the StartTime you specify does not fall on a data aggregation point, the first time point returned by the system is the nearest data aggregation point. For example, if the StartTime you specify is 2022-01-20T12:01:48Z, the first time point returned by the system is 2022-01-20T12:01:45Z.
	//
	// > 	- If the number of data shards of the instance is greater than or equal to 32, the minimum data aggregation frequency is 1 minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01m
	IntervalForHistory *string `json:"IntervalForHistory,omitempty" xml:"IntervalForHistory,omitempty"`
	// The monitoring metrics. To specify multiple monitoring metrics, separate them with commas (,). Take CpuUsage as an example:
	//
	// 	- Cluster or read/write splitting architecture:
	//
	//     - To query the overall CPU usage of all data nodes (DB), specify **CpuUsage$db**.
	//
	//    - To query the CPU usage of a single data node, specify **CpuUsage*	- and set the NodeId parameter to the DB node.
	//
	// 	- Standard architecture (primary/secondary): Specify **CpuUsage**.
	//
	// For more information about monitoring metrics and their descriptions, see <props="china">[MonitorKeys parameter description](https://www.alibabacloud.com/help/en/redis/developer-reference/api-r-kvstore-2015-01-01-describehistorymonitorvalues-redis#monitorKeys-note)<props="intl">[MonitorKeys parameter description](https://www.alibabacloud.com/help/zh/redis/developer-reference/api-r-kvstore-2015-01-01-describehistorymonitorvalues-redis#monitorKeys-note).
	//
	// > 	- This parameter is empty by default. If this parameter is not specified, the UsedMemory and quotaMemory metrics are returned.
	//
	// > 	- To ensure query efficiency, specify no more than 5 monitoring metrics in a single request when querying metrics for a single node. Specify no more than 1 monitoring metric in a single request when querying aggregated monitoring metrics.
	//
	// example:
	//
	// CpuUsage
	MonitorKeys *string `json:"MonitorKeys,omitempty" xml:"MonitorKeys,omitempty"`
	// The node ID of the instance. Specify this parameter to query the monitoring information of a specific node.
	//
	// > 	- This parameter is available only when the instance uses a read/write splitting or cluster architecture.
	//
	// > 	- You can call the [DescribeLogicInstanceTopology](https://help.aliyun.com/document_detail/473786.html) operation to query node IDs.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****-db-0#1679****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// To query the metrics of a read-only node in a cloud-native read/write splitting instance, specify the **NodeId*	- parameter and set this parameter to **READONLY**.
	//
	// >  In other cases, you do not need to specify this parameter (or set it to **MASTER**).
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
	// The start time of the query for historical monitoring data. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
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
