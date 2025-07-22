// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSqlOptimizeStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetEndTime() *string
	SetFilterEnable(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetFilterEnable() *string
	SetInstanceId(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetInstanceId() *string
	SetNodeId(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetNodeId() *string
	SetStartTime(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetStartTime() *string
	SetThreshold(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetThreshold() *string
	SetUseMerging(v string) *GetInstanceSqlOptimizeStatisticRequest
	GetUseMerging() *string
}

type GetInstanceSqlOptimizeStatisticRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1662518540764
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to filter instances for which DAS Enterprise Edition is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If you set this parameter to **true**, only database instances for which DAS Enterprise Edition is disabled are queried. If you set this parameter to **false**, all database instances are queried.
	//
	// example:
	//
	// false
	FilterEnable *string `json:"FilterEnable,omitempty" xml:"FilterEnable,omitempty"`
	// The database instance ID.
	//
	// >  The database instance must be an ApsaraDB RDS for MySQL instance or a PolarDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-wz90h9560rvdz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// >  For ApsaraDB RDS for MySQL Cluster Edition instances or PolarDB for MySQL clusters, you must specify the node ID.
	//
	// example:
	//
	// pi-bp12v7243x012****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1661308902060
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The duration threshold for automatic SQL optimization events. After this parameter is specified, the system collects statistics on automatic SQL optimization events whose duration does not exceed the specified threshold.
	//
	// >  This parameter is a reserved parameter and does not take effect.
	//
	// example:
	//
	// None
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Specifies whether to merge automatic SQL optimization events. Valid values:
	//
	// 	- **true**: merges automatic SQL optimization events.
	//
	// 	- **false**: does not merge automatic SQL optimization events.
	//
	// >  This parameter is a reserved parameter and does not take effect.
	//
	// example:
	//
	// true
	UseMerging *string `json:"UseMerging,omitempty" xml:"UseMerging,omitempty"`
}

func (s GetInstanceSqlOptimizeStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSqlOptimizeStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetFilterEnable() *string {
	return s.FilterEnable
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetThreshold() *string {
	return s.Threshold
}

func (s *GetInstanceSqlOptimizeStatisticRequest) GetUseMerging() *string {
	return s.UseMerging
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetEndTime(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetFilterEnable(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.FilterEnable = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetInstanceId(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetNodeId(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.NodeId = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetStartTime(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetThreshold(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.Threshold = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) SetUseMerging(v string) *GetInstanceSqlOptimizeStatisticRequest {
	s.UseMerging = &v
	return s
}

func (s *GetInstanceSqlOptimizeStatisticRequest) Validate() error {
	return dara.Validate(s)
}
