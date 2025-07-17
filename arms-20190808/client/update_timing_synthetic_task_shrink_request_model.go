// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimingSyntheticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAssertionsShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetAvailableAssertionsShrink() *string
	SetCommonSettingShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetCommonSettingShrink() *string
	SetCustomPeriodShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetCustomPeriodShrink() *string
	SetFrequency(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetFrequency() *string
	SetMonitorConfShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetMonitorConfShrink() *string
	SetMonitorsShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetMonitorsShrink() *string
	SetName(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetName() *string
	SetRegionId(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetTagsShrink() *string
	SetTaskId(v string) *UpdateTimingSyntheticTaskShrinkRequest
	GetTaskId() *string
}

type UpdateTimingSyntheticTaskShrinkRequest struct {
	// The list of assertions.
	AvailableAssertionsShrink *string `json:"AvailableAssertions,omitempty" xml:"AvailableAssertions,omitempty"`
	// The general settings.
	CommonSettingShrink *string `json:"CommonSetting,omitempty" xml:"CommonSetting,omitempty"`
	// The custom cycle.
	CustomPeriodShrink *string `json:"CustomPeriod,omitempty" xml:"CustomPeriod,omitempty"`
	// The detection frequency. Valid values: 1m, 5m, 10m, 15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, and 24h.
	//
	// example:
	//
	// 5m
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The monitoring configurations.
	MonitorConfShrink *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	// The list of monitoring points.
	MonitorsShrink *string `json:"Monitors,omitempty" xml:"Monitors,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// AlibabaCloud DNS Task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 5308a2691f59422c8c3b7aeccxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTimingSyntheticTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetAvailableAssertionsShrink() *string {
	return s.AvailableAssertionsShrink
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetCommonSettingShrink() *string {
	return s.CommonSettingShrink
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetCustomPeriodShrink() *string {
	return s.CustomPeriodShrink
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetMonitorConfShrink() *string {
	return s.MonitorConfShrink
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetMonitorsShrink() *string {
	return s.MonitorsShrink
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetAvailableAssertionsShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.AvailableAssertionsShrink = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetCommonSettingShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.CommonSettingShrink = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetCustomPeriodShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.CustomPeriodShrink = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetFrequency(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetMonitorConfShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.MonitorConfShrink = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetMonitorsShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.MonitorsShrink = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetName(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetRegionId(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetResourceGroupId(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetTagsShrink(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) SetTaskId(v string) *UpdateTimingSyntheticTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
