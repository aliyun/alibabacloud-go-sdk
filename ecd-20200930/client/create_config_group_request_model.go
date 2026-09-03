// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigTimers(v []*CreateConfigGroupRequestConfigTimers) *CreateConfigGroupRequest
	GetConfigTimers() []*CreateConfigGroupRequestConfigTimers
	SetDescription(v string) *CreateConfigGroupRequest
	GetDescription() *string
	SetName(v string) *CreateConfigGroupRequest
	GetName() *string
	SetProductType(v string) *CreateConfigGroupRequest
	GetProductType() *string
	SetRegionId(v string) *CreateConfigGroupRequest
	GetRegionId() *string
	SetType(v string) *CreateConfigGroupRequest
	GetType() *string
}

type CreateConfigGroupRequest struct {
	// The configuration information of scheduled tasks. This parameter is a list.
	ConfigTimers []*CreateConfigGroupRequestConfigTimers `json:"ConfigTimers,omitempty" xml:"ConfigTimers,omitempty" type:"Repeated"`
	// The description of the configuration group.
	//
	// example:
	//
	// Description of the scheduled task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ScheduledTaskGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The product type used by the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID. This feature is not region-specific. Set this parameter to `cn-shanghai`.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Timer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateConfigGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigGroupRequest) GetConfigTimers() []*CreateConfigGroupRequestConfigTimers {
	return s.ConfigTimers
}

func (s *CreateConfigGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateConfigGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *CreateConfigGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateConfigGroupRequest) GetType() *string {
	return s.Type
}

func (s *CreateConfigGroupRequest) SetConfigTimers(v []*CreateConfigGroupRequestConfigTimers) *CreateConfigGroupRequest {
	s.ConfigTimers = v
	return s
}

func (s *CreateConfigGroupRequest) SetDescription(v string) *CreateConfigGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigGroupRequest) SetName(v string) *CreateConfigGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateConfigGroupRequest) SetProductType(v string) *CreateConfigGroupRequest {
	s.ProductType = &v
	return s
}

func (s *CreateConfigGroupRequest) SetRegionId(v string) *CreateConfigGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConfigGroupRequest) SetType(v string) *CreateConfigGroupRequest {
	s.Type = &v
	return s
}

func (s *CreateConfigGroupRequest) Validate() error {
	if s.ConfigTimers != nil {
		for _, item := range s.ConfigTimers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateConfigGroupRequestConfigTimers struct {
	// Specifies whether to allow end users to configure scheduled tasks.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression of the scheduled task.
	//
	// 	Notice: Specify the time in UTC. For example, to schedule a task at 00:00 (UTC+8) every day, use 0 0 16 ? 	- 1,2,3,4,5,6,7.</notice>
	//
	// example:
	//
	// 0 0 16 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcefully execute the task.
	//
	// example:
	//
	// true
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The time interval. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The advance notification time before the scheduled task is executed. Unit: seconds.
	//
	// example:
	//
	// 300
	NotificationTime *int32 `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	// The operation type of the scheduled task. Currently, only disconnect scheduled tasks support this parameter.
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The process whitelist for intelligent detection of no-operation scheduled tasks. If a specified process is running, the no-operation scheduled task is not triggered.
	ProcessWhitelist []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	// The reset type of the cloud desktop.
	//
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The list of segment timer configurations.
	SegmentTimers []*CreateConfigGroupRequestConfigTimersSegmentTimers `json:"SegmentTimers,omitempty" xml:"SegmentTimers,omitempty" type:"Repeated"`
	// The type of the scheduled task.
	//
	// This parameter is required.
	//
	// example:
	//
	// TIMER_BOOT
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
	// The trigger configuration type for no-operation scheduled tasks.
	//
	// example:
	//
	// Standard
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s CreateConfigGroupRequestConfigTimers) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigGroupRequestConfigTimers) GoString() string {
	return s.String()
}

func (s *CreateConfigGroupRequestConfigTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *CreateConfigGroupRequestConfigTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateConfigGroupRequestConfigTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *CreateConfigGroupRequestConfigTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateConfigGroupRequestConfigTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *CreateConfigGroupRequestConfigTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateConfigGroupRequestConfigTimers) GetProcessWhitelist() []*string {
	return s.ProcessWhitelist
}

func (s *CreateConfigGroupRequestConfigTimers) GetResetType() *string {
	return s.ResetType
}

func (s *CreateConfigGroupRequestConfigTimers) GetSegmentTimers() []*CreateConfigGroupRequestConfigTimersSegmentTimers {
	return s.SegmentTimers
}

func (s *CreateConfigGroupRequestConfigTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *CreateConfigGroupRequestConfigTimers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateConfigGroupRequestConfigTimers) SetAllowClientSetting(v bool) *CreateConfigGroupRequestConfigTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetCronExpression(v string) *CreateConfigGroupRequestConfigTimers {
	s.CronExpression = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetEnforce(v bool) *CreateConfigGroupRequestConfigTimers {
	s.Enforce = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetInterval(v int32) *CreateConfigGroupRequestConfigTimers {
	s.Interval = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetNotificationTime(v int32) *CreateConfigGroupRequestConfigTimers {
	s.NotificationTime = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetOperationType(v string) *CreateConfigGroupRequestConfigTimers {
	s.OperationType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetProcessWhitelist(v []*string) *CreateConfigGroupRequestConfigTimers {
	s.ProcessWhitelist = v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetResetType(v string) *CreateConfigGroupRequestConfigTimers {
	s.ResetType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetSegmentTimers(v []*CreateConfigGroupRequestConfigTimersSegmentTimers) *CreateConfigGroupRequestConfigTimers {
	s.SegmentTimers = v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetTimerType(v string) *CreateConfigGroupRequestConfigTimers {
	s.TimerType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) SetTriggerType(v string) *CreateConfigGroupRequestConfigTimers {
	s.TriggerType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimers) Validate() error {
	if s.SegmentTimers != nil {
		for _, item := range s.SegmentTimers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateConfigGroupRequestConfigTimersSegmentTimers struct {
	// The appointment timer used for executing scheduled tasks at specified time points.
	//
	// example:
	//
	// 1764660600967
	AppointmentTimer *int64 `json:"AppointmentTimer,omitempty" xml:"AppointmentTimer,omitempty"`
	// Specifies whether to create a snapshot.
	CreateSnapshot *bool `json:"CreateSnapshot,omitempty" xml:"CreateSnapshot,omitempty"`
	// The cron expression for the end of the scheduled task execution.
	//
	// example:
	//
	// 0 0 18 ? 	- 1-5
	EndCronExpression *string `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	// Specifies whether to forcefully execute the task. If set to true, the scheduled task is forcefully executed regardless of the desktop and connection status.
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The time interval. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The list of effective IP CIDR blocks.
	IpSegments []*string `json:"IpSegments,omitempty" xml:"IpSegments,omitempty" type:"Repeated"`
	// The duration of inactivity before the screen is locked, used by the no-operation lock screen feature. Unit: minutes. Only AD cloud desktops are supported.
	//
	// example:
	//
	// 5
	LockScreenTime *int32 `json:"LockScreenTime,omitempty" xml:"LockScreenTime,omitempty"`
	// The advance notification time before the scheduled task is executed. Unit: seconds.
	//
	// example:
	//
	// 300
	NotificationTime *int32 `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	// The operation type of the scheduled task. Currently, only disconnect scheduled tasks support this parameter.
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The patch ID.
	//
	// example:
	//
	// KB5082063
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The process whitelist for intelligent detection of no-operation scheduled tasks. If a specified process is running, the no-operation scheduled task is not triggered.
	ProcessWhitelist []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	// The reset type, which determines whether to reset and the scope of cloud disks to reset.
	//
	// example:
	//
	// 1
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The cron expression for the start of the scheduled task execution.
	//
	// example:
	//
	// 0 0 8 ? 	- 1-5
	StartCronExpression *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
	// The execution order number of the timer.
	//
	// example:
	//
	// 1
	TimerOrder *int32 `json:"TimerOrder,omitempty" xml:"TimerOrder,omitempty"`
	// The time zone used by the scheduled task.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The trigger configuration type for no-operation scheduled tasks.
	//
	// example:
	//
	// Standard
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The advance notification time before verification is executed. Unit: seconds.
	//
	// example:
	//
	// 300
	VerificationNotificationTime *int32 `json:"VerificationNotificationTime,omitempty" xml:"VerificationNotificationTime,omitempty"`
	// The verification wait duration. Unit: seconds.
	//
	// example:
	//
	// 600
	VerificationTime *int32 `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s CreateConfigGroupRequestConfigTimersSegmentTimers) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigGroupRequestConfigTimersSegmentTimers) GoString() string {
	return s.String()
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetAppointmentTimer() *int64 {
	return s.AppointmentTimer
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetCreateSnapshot() *bool {
	return s.CreateSnapshot
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetImageId() *string {
	return s.ImageId
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetIpSegments() []*string {
	return s.IpSegments
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetLockScreenTime() *int32 {
	return s.LockScreenTime
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetPatchId() *string {
	return s.PatchId
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetProcessWhitelist() []*string {
	return s.ProcessWhitelist
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetResetType() *string {
	return s.ResetType
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetTimerOrder() *int32 {
	return s.TimerOrder
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetTimezone() *string {
	return s.Timezone
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetVerificationNotificationTime() *int32 {
	return s.VerificationNotificationTime
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetVerificationTime() *int32 {
	return s.VerificationTime
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetAppointmentTimer(v int64) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.AppointmentTimer = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetCreateSnapshot(v bool) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.CreateSnapshot = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetEndCronExpression(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.EndCronExpression = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetEnforce(v bool) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.Enforce = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetImageId(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.ImageId = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetInterval(v int32) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.Interval = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetIpSegments(v []*string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.IpSegments = v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetLockScreenTime(v int32) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.LockScreenTime = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetNotificationTime(v int32) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.NotificationTime = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetOperationType(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.OperationType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetPatchId(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.PatchId = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetProcessWhitelist(v []*string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.ProcessWhitelist = v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetResetType(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.ResetType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetStartCronExpression(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.StartCronExpression = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetTimerOrder(v int32) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.TimerOrder = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetTimezone(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.Timezone = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetTriggerType(v string) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.TriggerType = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetVerificationNotificationTime(v int32) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.VerificationNotificationTime = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetVerificationTime(v int32) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.VerificationTime = &v
	return s
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) Validate() error {
	return dara.Validate(s)
}
