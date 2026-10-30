// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGpuClusterListAlarmRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmCluster(v string) *GetGpuClusterListAlarmRecordRequest
	GetAlarmCluster() *string
	SetAlarmExtInfo(v string) *GetGpuClusterListAlarmRecordRequest
	GetAlarmExtInfo() *string
	SetAlarmTask(v string) *GetGpuClusterListAlarmRecordRequest
	GetAlarmTask() *string
	SetAlarmTime(v string) *GetGpuClusterListAlarmRecordRequest
	GetAlarmTime() *string
	SetAlarmType(v string) *GetGpuClusterListAlarmRecordRequest
	GetAlarmType() *string
	SetCurrent(v int64) *GetGpuClusterListAlarmRecordRequest
	GetCurrent() *int64
	SetPageSize(v int64) *GetGpuClusterListAlarmRecordRequest
	GetPageSize() *int64
	SetStatus(v string) *GetGpuClusterListAlarmRecordRequest
	GetStatus() *string
}

type GetGpuClusterListAlarmRecordRequest struct {
	// The cluster to query. Specify a cluster, such as "cluster1". To query multiple clusters, separate them with commas, such as "cluster1,cluster2,cluster3".
	//
	// example:
	//
	// "cluster1"
	AlarmCluster *string `json:"alarmCluster,omitempty" xml:"alarmCluster,omitempty"`
	// The instance ID or container used to filter anomalies.
	//
	// example:
	//
	// instance id
	AlarmExtInfo *string `json:"alarmExtInfo,omitempty" xml:"alarmExtInfo,omitempty"`
	// The task to query. Specify a task, such as "task1". To query multiple tasks, separate them with commas, such as "task1,task2,task3".
	//
	// example:
	//
	// task id
	AlarmTask *string `json:"alarmTask,omitempty" xml:"alarmTask,omitempty"`
	// The specific time or time range to query. Specify a specific time, such as "2025-08-22 18:16:25". To specify a time range, use the "start,end" format, such as "2025-08-22 18:16:25,2025-08-23 18:16:25".
	//
	// example:
	//
	// "2025-08-22 18:16:25,2025-08-23 18:16:25"
	AlarmTime *string `json:"alarmTime,omitempty" xml:"alarmTime,omitempty"`
	// The anomaly type used to filter results, such as GPU memory access out of bounds or task hang.
	//
	// example:
	//
	// cuda error
	AlarmType *string `json:"alarmType,omitempty" xml:"alarmType,omitempty"`
	// The current page number for pagination.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The number of entries per page for pagination.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The task status used to filter results. Valid values: blocking, running, exited, and restart.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetGpuClusterListAlarmRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGpuClusterListAlarmRecordRequest) GoString() string {
	return s.String()
}

func (s *GetGpuClusterListAlarmRecordRequest) GetAlarmCluster() *string {
	return s.AlarmCluster
}

func (s *GetGpuClusterListAlarmRecordRequest) GetAlarmExtInfo() *string {
	return s.AlarmExtInfo
}

func (s *GetGpuClusterListAlarmRecordRequest) GetAlarmTask() *string {
	return s.AlarmTask
}

func (s *GetGpuClusterListAlarmRecordRequest) GetAlarmTime() *string {
	return s.AlarmTime
}

func (s *GetGpuClusterListAlarmRecordRequest) GetAlarmType() *string {
	return s.AlarmType
}

func (s *GetGpuClusterListAlarmRecordRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *GetGpuClusterListAlarmRecordRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetGpuClusterListAlarmRecordRequest) GetStatus() *string {
	return s.Status
}

func (s *GetGpuClusterListAlarmRecordRequest) SetAlarmCluster(v string) *GetGpuClusterListAlarmRecordRequest {
	s.AlarmCluster = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetAlarmExtInfo(v string) *GetGpuClusterListAlarmRecordRequest {
	s.AlarmExtInfo = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetAlarmTask(v string) *GetGpuClusterListAlarmRecordRequest {
	s.AlarmTask = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetAlarmTime(v string) *GetGpuClusterListAlarmRecordRequest {
	s.AlarmTime = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetAlarmType(v string) *GetGpuClusterListAlarmRecordRequest {
	s.AlarmType = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetCurrent(v int64) *GetGpuClusterListAlarmRecordRequest {
	s.Current = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetPageSize(v int64) *GetGpuClusterListAlarmRecordRequest {
	s.PageSize = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) SetStatus(v string) *GetGpuClusterListAlarmRecordRequest {
	s.Status = &v
	return s
}

func (s *GetGpuClusterListAlarmRecordRequest) Validate() error {
	return dara.Validate(s)
}
