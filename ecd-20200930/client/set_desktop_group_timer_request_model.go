// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *SetDesktopGroupTimerRequest
	GetCronExpression() *string
	SetDesktopGroupId(v string) *SetDesktopGroupTimerRequest
	GetDesktopGroupId() *string
	SetForce(v bool) *SetDesktopGroupTimerRequest
	GetForce() *bool
	SetRegionId(v string) *SetDesktopGroupTimerRequest
	GetRegionId() *string
	SetResetType(v int32) *SetDesktopGroupTimerRequest
	GetResetType() *int32
	SetTimerType(v int32) *SetDesktopGroupTimerRequest
	GetTimerType() *int32
}

type SetDesktopGroupTimerRequest struct {
	// The cron expression for the scheduled task. This parameter is required when `TimerType` is set to 2, 3, or 4.
	//
	// example:
	//
	// 0 0 2 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-dbdkfmh883****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// Specifies whether to forcefully execute the scheduled task.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the disk that you want to reset.
	//
	// Valid values:
	//
	// - does not reset disks.
	//
	// - resets only the system disk.
	//
	// - resets only the user disk.
	//
	// - resets the system disk and the user disk.
	//
	// example:
	//
	// 1
	ResetType *int32 `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The type of the scheduled task.
	//
	// Valid values:
	//
	// 	- 1: scheduled reset
	//
	// 	- 2: scheduled startup
	//
	// 	- 3: scheduled stop
	//
	// 	- 4: scheduled restart
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimerType *int32 `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s SetDesktopGroupTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupTimerRequest) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupTimerRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *SetDesktopGroupTimerRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *SetDesktopGroupTimerRequest) GetForce() *bool {
	return s.Force
}

func (s *SetDesktopGroupTimerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDesktopGroupTimerRequest) GetResetType() *int32 {
	return s.ResetType
}

func (s *SetDesktopGroupTimerRequest) GetTimerType() *int32 {
	return s.TimerType
}

func (s *SetDesktopGroupTimerRequest) SetCronExpression(v string) *SetDesktopGroupTimerRequest {
	s.CronExpression = &v
	return s
}

func (s *SetDesktopGroupTimerRequest) SetDesktopGroupId(v string) *SetDesktopGroupTimerRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *SetDesktopGroupTimerRequest) SetForce(v bool) *SetDesktopGroupTimerRequest {
	s.Force = &v
	return s
}

func (s *SetDesktopGroupTimerRequest) SetRegionId(v string) *SetDesktopGroupTimerRequest {
	s.RegionId = &v
	return s
}

func (s *SetDesktopGroupTimerRequest) SetResetType(v int32) *SetDesktopGroupTimerRequest {
	s.ResetType = &v
	return s
}

func (s *SetDesktopGroupTimerRequest) SetTimerType(v int32) *SetDesktopGroupTimerRequest {
	s.TimerType = &v
	return s
}

func (s *SetDesktopGroupTimerRequest) Validate() error {
	return dara.Validate(s)
}
