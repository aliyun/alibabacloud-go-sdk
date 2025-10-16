// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *ModifyDesktopTimerRequest
	GetDesktopId() []*string
	SetDesktopTimers(v []*ModifyDesktopTimerRequestDesktopTimers) *ModifyDesktopTimerRequest
	GetDesktopTimers() []*ModifyDesktopTimerRequestDesktopTimers
	SetRegionId(v string) *ModifyDesktopTimerRequest
	GetRegionId() *string
	SetUseDesktopTimers(v bool) *ModifyDesktopTimerRequest
	GetUseDesktopTimers() *bool
}

type ModifyDesktopTimerRequest struct {
	// The IDs of the cloud computers.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The details of the scheduled task on cloud computers.
	DesktopTimers []*ModifyDesktopTimerRequestDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to use the `DesktopTimers`*	- parameter. Set the value to `true`**.
	//
	// example:
	//
	// true
	UseDesktopTimers *bool `json:"UseDesktopTimers,omitempty" xml:"UseDesktopTimers,omitempty"`
}

func (s ModifyDesktopTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopTimerRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopTimerRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *ModifyDesktopTimerRequest) GetDesktopTimers() []*ModifyDesktopTimerRequestDesktopTimers {
	return s.DesktopTimers
}

func (s *ModifyDesktopTimerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDesktopTimerRequest) GetUseDesktopTimers() *bool {
	return s.UseDesktopTimers
}

func (s *ModifyDesktopTimerRequest) SetDesktopId(v []*string) *ModifyDesktopTimerRequest {
	s.DesktopId = v
	return s
}

func (s *ModifyDesktopTimerRequest) SetDesktopTimers(v []*ModifyDesktopTimerRequestDesktopTimers) *ModifyDesktopTimerRequest {
	s.DesktopTimers = v
	return s
}

func (s *ModifyDesktopTimerRequest) SetRegionId(v string) *ModifyDesktopTimerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDesktopTimerRequest) SetUseDesktopTimers(v bool) *ModifyDesktopTimerRequest {
	s.UseDesktopTimers = &v
	return s
}

func (s *ModifyDesktopTimerRequest) Validate() error {
	if s.DesktopTimers != nil {
		for _, item := range s.DesktopTimers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDesktopTimerRequestDesktopTimers struct {
	// Specifies whether to allow end users to configure the scheduled task.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression of the schedule.
	//
	// > The time must be in UTC. For example, for 24:00 (UTC+8), you must set the value to 0 0 16 ? \\	- 1,2,3,4,5,6,7
	//
	// example:
	//
	// 0 0 16 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcibly execute the scheduled task.
	//
	// Valid values:
	//
	// 	- true: forcibly executes the scheduled task regardless of the status and connection of the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false: does not forcibly execute the scheduled task.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The interval at which the scheduled task is executed. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The operations that scheduled tasks support. This parameter is valid only when TimerType is set to NoConnect.
	//
	// Valid values:
	//
	// 	- Hibernate: hibernates the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Shutdown: stops the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The reset type of the cloud computers.
	//
	// Valid values:
	//
	// 	- RESET_TYPE_SYSTE: resets the system disk.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- RESET_TYPE_BOTH: resets data and user disks.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The type of the scheduled task.
	//
	// Valid values:
	//
	// 	- NoOperationDisconnect: Disconnects the cloud computers without performing operations on the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- LogoutShutdown: Stops the cloud computers when end users log out Alibaba Cloud Workspace clients.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NoConnect: Disconnects the cloud computers when end users perform one of the actions that is specified by the OperationType parameter.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TimerBoot: Starts the cloud computers on schedule.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TimerReset: Resets the cloud computers on schedule.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- LoginAutoConnect: automatically connects to the cloud computers when end users log on to Alibaba Cloud Workspace clients.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NoOperationShutdown: Stops the cloud computers without performing operations on the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TimerShutdown: Stops the cloud computers on schedule.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NoOperationReboot: Restarts the cloud computers without performing operations on the cloud computers.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- TimerReboot: Restarts the cloud computers on schedule.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// TimerBoot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s ModifyDesktopTimerRequestDesktopTimers) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopTimerRequestDesktopTimers) GoString() string {
	return s.String()
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetResetType() *string {
	return s.ResetType
}

func (s *ModifyDesktopTimerRequestDesktopTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetAllowClientSetting(v bool) *ModifyDesktopTimerRequestDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetCronExpression(v string) *ModifyDesktopTimerRequestDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetEnforce(v bool) *ModifyDesktopTimerRequestDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetInterval(v int32) *ModifyDesktopTimerRequestDesktopTimers {
	s.Interval = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetOperationType(v string) *ModifyDesktopTimerRequestDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetResetType(v string) *ModifyDesktopTimerRequestDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) SetTimerType(v string) *ModifyDesktopTimerRequestDesktopTimers {
	s.TimerType = &v
	return s
}

func (s *ModifyDesktopTimerRequestDesktopTimers) Validate() error {
	return dara.Validate(s)
}
