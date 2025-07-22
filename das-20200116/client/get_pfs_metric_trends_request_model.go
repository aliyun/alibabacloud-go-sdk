// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPfsMetricTrendsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetPfsMetricTrendsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetPfsMetricTrendsRequest
	GetInstanceId() *string
	SetMetric(v string) *GetPfsMetricTrendsRequest
	GetMetric() *string
	SetNodeId(v string) *GetPfsMetricTrendsRequest
	GetNodeId() *string
	SetStartTime(v int64) *GetPfsMetricTrendsRequest
	GetStartTime() *int64
}

type GetPfsMetricTrendsRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time. You can view the data of up to seven days in the previous 30 days.
	//
	// example:
	//
	// 1678432430967
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-m5ea73876ukci****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric whose trend you want to query. Valid values:
	//
	// 	- **count**: the number of executions.
	//
	// 	- **avgRt**: the average execution duration.
	//
	// 	- **rtRate**: the execution duration percentage.
	//
	// 	- **rowsExamined**: the total number of scanned rows.
	//
	// example:
	//
	// Count
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The node ID.
	//
	// >  This parameter is required if the database instance is an ApsaraDB RDS for MySQL Cluster Edition instance or a PolarDB for MySQL clusters.
	//
	// example:
	//
	// r-x****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1677461663092
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPfsMetricTrendsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPfsMetricTrendsRequest) GoString() string {
	return s.String()
}

func (s *GetPfsMetricTrendsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPfsMetricTrendsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPfsMetricTrendsRequest) GetMetric() *string {
	return s.Metric
}

func (s *GetPfsMetricTrendsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPfsMetricTrendsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPfsMetricTrendsRequest) SetEndTime(v int64) *GetPfsMetricTrendsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPfsMetricTrendsRequest) SetInstanceId(v string) *GetPfsMetricTrendsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPfsMetricTrendsRequest) SetMetric(v string) *GetPfsMetricTrendsRequest {
	s.Metric = &v
	return s
}

func (s *GetPfsMetricTrendsRequest) SetNodeId(v string) *GetPfsMetricTrendsRequest {
	s.NodeId = &v
	return s
}

func (s *GetPfsMetricTrendsRequest) SetStartTime(v int64) *GetPfsMetricTrendsRequest {
	s.StartTime = &v
	return s
}

func (s *GetPfsMetricTrendsRequest) Validate() error {
	return dara.Validate(s)
}
