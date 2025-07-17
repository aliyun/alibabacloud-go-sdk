// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimingSyntheticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAssertionsShrink(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetAvailableAssertionsShrink() *string
	SetCommonSettingShrink(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetCommonSettingShrink() *string
	SetCustomPeriodShrink(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetCustomPeriodShrink() *string
	SetFrequency(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetFrequency() *string
	SetMonitorCategory(v int32) *CreateTimingSyntheticTaskShrinkRequest
	GetMonitorCategory() *int32
	SetMonitorConfShrink(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetMonitorConfShrink() *string
	SetMonitorsShrink(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetMonitorsShrink() *string
	SetName(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetName() *string
	SetRegionId(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *CreateTimingSyntheticTaskShrinkRequest
	GetTagsShrink() *string
	SetTaskType(v int32) *CreateTimingSyntheticTaskShrinkRequest
	GetTaskType() *int32
}

type CreateTimingSyntheticTaskShrinkRequest struct {
	// The list of assertions.
	AvailableAssertionsShrink *string `json:"AvailableAssertions,omitempty" xml:"AvailableAssertions,omitempty"`
	// The general settings.
	CommonSettingShrink *string `json:"CommonSetting,omitempty" xml:"CommonSetting,omitempty"`
	// The general settings.
	CustomPeriodShrink *string `json:"CustomPeriod,omitempty" xml:"CustomPeriod,omitempty"`
	// The detection frequency. Valid values: 1m, 5m, 10m, 15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, and 24h.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5m
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The detection point type. Valid values:
	//
	// - 1: PC
	//
	// - 2: mobile device
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MonitorCategory *int32 `json:"MonitorCategory,omitempty" xml:"MonitorCategory,omitempty"`
	// The monitoring configurations.
	//
	// This parameter is required.
	MonitorConfShrink *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	// The list of detection points.
	//
	// This parameter is required.
	MonitorsShrink *string `json:"Monitors,omitempty" xml:"Monitors,omitempty"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameter is optional.
	//
	// example:
	//
	// xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the task. Valid values:
	//
	// 1: ICMP. 2: TCP. 3: DNS. 4: HTTP. 5: website speed measurement. 6: file download.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateTimingSyntheticTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetAvailableAssertionsShrink() *string {
	return s.AvailableAssertionsShrink
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetCommonSettingShrink() *string {
	return s.CommonSettingShrink
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetCustomPeriodShrink() *string {
	return s.CustomPeriodShrink
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetMonitorCategory() *int32 {
	return s.MonitorCategory
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetMonitorConfShrink() *string {
	return s.MonitorConfShrink
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetMonitorsShrink() *string {
	return s.MonitorsShrink
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateTimingSyntheticTaskShrinkRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetAvailableAssertionsShrink(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.AvailableAssertionsShrink = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetCommonSettingShrink(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.CommonSettingShrink = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetCustomPeriodShrink(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.CustomPeriodShrink = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetFrequency(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetMonitorCategory(v int32) *CreateTimingSyntheticTaskShrinkRequest {
	s.MonitorCategory = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetMonitorConfShrink(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.MonitorConfShrink = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetMonitorsShrink(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.MonitorsShrink = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetName(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetRegionId(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetResourceGroupId(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetTagsShrink(v string) *CreateTimingSyntheticTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) SetTaskType(v int32) *CreateTimingSyntheticTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *CreateTimingSyntheticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
