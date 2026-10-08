// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceRequest
	GetEndTime() *string
	SetKey(v string) *DescribeDBInstancePerformanceRequest
	GetKey() *string
	SetNodeId(v string) *DescribeDBInstancePerformanceRequest
	GetNodeId() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDBInstancePerformanceRequest
	GetStartTime() *string
}

type DescribeDBInstancePerformanceRequest struct {
	// The instance ID. You can call DescribeDBInstances to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the query. Format: <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z (UTC).
	//
	// > The interval between the start time and end time must be greater than the monitoring frequency of your instance. Otherwise, an empty list may be returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2012-06-18T15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance metrics that you want to query. Separate multiple values with commas (,). You can specify up to 30 metrics. For more information, see [Performance parameters](https://help.aliyun.com/document_detail/26316.html).
	//
	// > If **Key*	- is set to **MySQL_SpaceUsage*	- or **SQLServer_SpaceUsage**, only monitoring data within the last day can be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_NetworkTraffic
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The unique identifier of the instance.
	//
	// example:
	//
	// 339****
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the query. Format: <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z (UTC).
	//
	// > The interval between the start time and end time must be greater than the monitoring frequency of your instance. Otherwise, an empty list may be returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2012-06-08T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancePerformanceRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstancePerformanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancePerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetNodeId(v string) *DescribeDBInstancePerformanceRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetResourceOwnerId(v int64) *DescribeDBInstancePerformanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) Validate() error {
	return dara.Validate(s)
}
