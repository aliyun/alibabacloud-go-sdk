// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSyntheticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommonParamShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetCommonParamShrink() *string
	SetDownloadShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetDownloadShrink() *string
	SetExtendIntervalShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetExtendIntervalShrink() *string
	SetIntervalTime(v string) *CreateSyntheticTaskShrinkRequest
	GetIntervalTime() *string
	SetIntervalType(v string) *CreateSyntheticTaskShrinkRequest
	GetIntervalType() *string
	SetIpType(v int64) *CreateSyntheticTaskShrinkRequest
	GetIpType() *int64
	SetMonitorListShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetMonitorListShrink() *string
	SetNavigationShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetNavigationShrink() *string
	SetNetShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetNetShrink() *string
	SetProtocolShrink(v string) *CreateSyntheticTaskShrinkRequest
	GetProtocolShrink() *string
	SetRegionId(v string) *CreateSyntheticTaskShrinkRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateSyntheticTaskShrinkRequest
	GetTaskName() *string
	SetTaskType(v int64) *CreateSyntheticTaskShrinkRequest
	GetTaskType() *int64
	SetUpdateTask(v bool) *CreateSyntheticTaskShrinkRequest
	GetUpdateTask() *bool
	SetUrl(v string) *CreateSyntheticTaskShrinkRequest
	GetUrl() *string
}

type CreateSyntheticTaskShrinkRequest struct {
	// The common parameters.
	CommonParamShrink *string `json:"CommonParam,omitempty" xml:"CommonParam,omitempty"`
	// The file download task.
	DownloadShrink *string `json:"Download,omitempty" xml:"Download,omitempty"`
	// The frequency.
	ExtendIntervalShrink *string `json:"ExtendInterval,omitempty" xml:"ExtendInterval,omitempty"`
	// The interval at which synthetic monitoring is performed. Unit: minutes. Valid values:
	//
	// 	- 1
	//
	// 	- 5
	//
	// 	- 10
	//
	// 	- 15
	//
	// 	- 20
	//
	// 	- 30
	//
	// 	- 60
	//
	// 	- 120
	//
	// 	- 180
	//
	// 	- 240
	//
	// 	- 360
	//
	// 	- 480
	//
	// 	- 720
	//
	// 	- 1440
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	IntervalTime *string `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	// The interval type.
	//
	// 	- 0: daily
	//
	// 	- 1: custom frequency
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IntervalType *string `json:"IntervalType,omitempty" xml:"IntervalType,omitempty"`
	// The IP address type:
	//
	// 	- 0: an automatic IP address
	//
	// 	- 1: IPv4
	//
	// 	- 2: IPv6
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IpType *int64 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The list of monitoring points.
	//
	// This parameter is required.
	MonitorListShrink *string `json:"MonitorList,omitempty" xml:"MonitorList,omitempty"`
	// The monitoring items that are associated with the browse tasks.
	NavigationShrink *string `json:"Navigation,omitempty" xml:"Navigation,omitempty"`
	// The network synthetic monitoring task.
	NetShrink *string `json:"Net,omitempty" xml:"Net,omitempty"`
	// The API performance synthetic monitoring task.
	ProtocolShrink *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the region in which the application is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the task. To update a synthetic monitoring task, enter the task name and set the **UpdateTask*	- parameter to **true**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Network synthetic monitoring task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the monitoring task. Valid values:
	//
	// 1.  3: web page performance - IE
	//
	// 2.  34: web Page Performance - Chrome
	//
	// 3.  0: network quality
	//
	// 4.  40: file download
	//
	// 5.  7:API performance
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Specifies whether to update existing synthetic monitoring tasks.
	//
	// 	- true: updates existing synthetic monitoring tasks.
	//
	// 	- false: creates new synthetic monitoring tasks.
	//
	// example:
	//
	// false
	UpdateTask *bool `json:"UpdateTask,omitempty" xml:"UpdateTask,omitempty"`
	// The URL for synthetic monitoring.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.example.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateSyntheticTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSyntheticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskShrinkRequest) GetCommonParamShrink() *string {
	return s.CommonParamShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetDownloadShrink() *string {
	return s.DownloadShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetExtendIntervalShrink() *string {
	return s.ExtendIntervalShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetIntervalTime() *string {
	return s.IntervalTime
}

func (s *CreateSyntheticTaskShrinkRequest) GetIntervalType() *string {
	return s.IntervalType
}

func (s *CreateSyntheticTaskShrinkRequest) GetIpType() *int64 {
	return s.IpType
}

func (s *CreateSyntheticTaskShrinkRequest) GetMonitorListShrink() *string {
	return s.MonitorListShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetNavigationShrink() *string {
	return s.NavigationShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetNetShrink() *string {
	return s.NetShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetProtocolShrink() *string {
	return s.ProtocolShrink
}

func (s *CreateSyntheticTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSyntheticTaskShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateSyntheticTaskShrinkRequest) GetTaskType() *int64 {
	return s.TaskType
}

func (s *CreateSyntheticTaskShrinkRequest) GetUpdateTask() *bool {
	return s.UpdateTask
}

func (s *CreateSyntheticTaskShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateSyntheticTaskShrinkRequest) SetCommonParamShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.CommonParamShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetDownloadShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.DownloadShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetExtendIntervalShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.ExtendIntervalShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetIntervalTime(v string) *CreateSyntheticTaskShrinkRequest {
	s.IntervalTime = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetIntervalType(v string) *CreateSyntheticTaskShrinkRequest {
	s.IntervalType = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetIpType(v int64) *CreateSyntheticTaskShrinkRequest {
	s.IpType = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetMonitorListShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.MonitorListShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetNavigationShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.NavigationShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetNetShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.NetShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetProtocolShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.ProtocolShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetRegionId(v string) *CreateSyntheticTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetTaskName(v string) *CreateSyntheticTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetTaskType(v int64) *CreateSyntheticTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetUpdateTask(v bool) *CreateSyntheticTaskShrinkRequest {
	s.UpdateTask = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetUrl(v string) *CreateSyntheticTaskShrinkRequest {
	s.Url = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
