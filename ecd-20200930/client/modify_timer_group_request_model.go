// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTimerGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigTimers(v []*ModifyTimerGroupRequestConfigTimers) *ModifyTimerGroupRequest
	GetConfigTimers() []*ModifyTimerGroupRequestConfigTimers
	SetDescription(v string) *ModifyTimerGroupRequest
	GetDescription() *string
	SetGroupId(v string) *ModifyTimerGroupRequest
	GetGroupId() *string
	SetName(v string) *ModifyTimerGroupRequest
	GetName() *string
	SetRegionId(v string) *ModifyTimerGroupRequest
	GetRegionId() *string
}

type ModifyTimerGroupRequest struct {
	// The scheduled tasks.
	ConfigTimers []*ModifyTimerGroupRequestConfigTimers `json:"ConfigTimers,omitempty" xml:"ConfigTimers,omitempty" type:"Repeated"`
	// The description of the configuration group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the configuration group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyTimerGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTimerGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyTimerGroupRequest) GetConfigTimers() []*ModifyTimerGroupRequestConfigTimers {
	return s.ConfigTimers
}

func (s *ModifyTimerGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyTimerGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyTimerGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyTimerGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyTimerGroupRequest) SetConfigTimers(v []*ModifyTimerGroupRequestConfigTimers) *ModifyTimerGroupRequest {
	s.ConfigTimers = v
	return s
}

func (s *ModifyTimerGroupRequest) SetDescription(v string) *ModifyTimerGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyTimerGroupRequest) SetGroupId(v string) *ModifyTimerGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyTimerGroupRequest) SetName(v string) *ModifyTimerGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyTimerGroupRequest) SetRegionId(v string) *ModifyTimerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTimerGroupRequest) Validate() error {
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

type ModifyTimerGroupRequestConfigTimers struct {
	// Specifies whether to allow end users to configure the scheduled task.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression specified in the scheduled task.
	//
	// >  The time must be in UTC. For example, if your local time is 24:00 (UTC+8), you must set the value to 0 0 16 ? \\	- 1,2,3,4,5,6,7.
	//
	// example:
	//
	// 0 0 16 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcibly execute the scheduled task. A value of true specifies the scheduled task will run forcefully, ignoring the cloud computer and connection status.
	//
	// example:
	//
	// false
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The interval at which the scheduled task is executed. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval         *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	NotificationTime *int32 `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	// The type of the scheduled operation. If you set TimerType to NoConnect, you can specify this parameter.
	//
	// Valid values:
	//
	// 	- Hibernate: scheduled hibernation.
	//
	// 	- Shutdown: scheduled shutdown.
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The process whitelist. If whitelisted processes are running, the scheduled task does not take effect.
	ProcessWhitelist []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	// The reset option.
	//
	// Valid value:
	//
	// 	- RESET_TYPE_SYSTEM: resets the system disk.
	//
	// 	- RESET_TYPE_USER_DISK: resets the data disk.
	//
	// 	- RESET_TYPE_BOTH: resets the system disk and data disk.
	//
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType     *string                                             `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	SegmentTimers []*ModifyTimerGroupRequestConfigTimersSegmentTimers `json:"SegmentTimers,omitempty" xml:"SegmentTimers,omitempty" type:"Repeated"`
	// The type of the scheduled task.
	//
	// Valid value:
	//
	// 	- NoOperationDisconnect: scheduled disconnection upon inactivity.
	//
	// 	- NoConnect: scheduled disconnection upon specified operation (OperationType).
	//
	// 	- TimerBoot: scheduled start.
	//
	// 	- TimerReset: scheduled reset.
	//
	// 	- NoOperationShutdown: scheduled shutdown upon inactivity.
	//
	// 	- NoOperationHibernate: scheduled hibernation upon inactivity.
	//
	// 	- TimerShutdown: scheduled shutdown.
	//
	// 	- NoOperationReboot: scheduled restart upon inactivity.
	//
	// 	- TimerReboot: Restarts the cloud computers on schedule.
	//
	// example:
	//
	// TIMER_BOOT
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
	// The method to trigger the scheduled task upon inactivity.
	//
	// Valid values:
	//
	// 	- Advanced: intelligent detection.
	//
	// 	- Standard: standard detection.
	//
	// example:
	//
	// Standard
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ModifyTimerGroupRequestConfigTimers) String() string {
	return dara.Prettify(s)
}

func (s ModifyTimerGroupRequestConfigTimers) GoString() string {
	return s.String()
}

func (s *ModifyTimerGroupRequestConfigTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *ModifyTimerGroupRequestConfigTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyTimerGroupRequestConfigTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *ModifyTimerGroupRequestConfigTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyTimerGroupRequestConfigTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *ModifyTimerGroupRequestConfigTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *ModifyTimerGroupRequestConfigTimers) GetProcessWhitelist() []*string {
	return s.ProcessWhitelist
}

func (s *ModifyTimerGroupRequestConfigTimers) GetResetType() *string {
	return s.ResetType
}

func (s *ModifyTimerGroupRequestConfigTimers) GetSegmentTimers() []*ModifyTimerGroupRequestConfigTimersSegmentTimers {
	return s.SegmentTimers
}

func (s *ModifyTimerGroupRequestConfigTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *ModifyTimerGroupRequestConfigTimers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ModifyTimerGroupRequestConfigTimers) SetAllowClientSetting(v bool) *ModifyTimerGroupRequestConfigTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetCronExpression(v string) *ModifyTimerGroupRequestConfigTimers {
	s.CronExpression = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetEnforce(v bool) *ModifyTimerGroupRequestConfigTimers {
	s.Enforce = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetInterval(v int32) *ModifyTimerGroupRequestConfigTimers {
	s.Interval = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetNotificationTime(v int32) *ModifyTimerGroupRequestConfigTimers {
	s.NotificationTime = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetOperationType(v string) *ModifyTimerGroupRequestConfigTimers {
	s.OperationType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetProcessWhitelist(v []*string) *ModifyTimerGroupRequestConfigTimers {
	s.ProcessWhitelist = v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetResetType(v string) *ModifyTimerGroupRequestConfigTimers {
	s.ResetType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetSegmentTimers(v []*ModifyTimerGroupRequestConfigTimersSegmentTimers) *ModifyTimerGroupRequestConfigTimers {
	s.SegmentTimers = v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetTimerType(v string) *ModifyTimerGroupRequestConfigTimers {
	s.TimerType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) SetTriggerType(v string) *ModifyTimerGroupRequestConfigTimers {
	s.TriggerType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimers) Validate() error {
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

type ModifyTimerGroupRequestConfigTimersSegmentTimers struct {
	// example:
	//
	// 1764660600967
	AppointmentTimer  *int64  `json:"AppointmentTimer,omitempty" xml:"AppointmentTimer,omitempty"`
	EndCronExpression *string `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	Enforce           *bool   `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// example:
	//
	// m-5b0vjqbiqu010XXXXXX
	ImageId  *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Interval *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 1800
	LockScreenTime      *int32    `json:"LockScreenTime,omitempty" xml:"LockScreenTime,omitempty"`
	NotificationTime    *int32    `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	OperationType       *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ProcessWhitelist    []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	ResetType           *string   `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	StartCronExpression *string   `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
	TimerOrder          *int32    `json:"TimerOrder,omitempty" xml:"TimerOrder,omitempty"`
	Timezone            *string   `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	TriggerType         *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ModifyTimerGroupRequestConfigTimersSegmentTimers) String() string {
	return dara.Prettify(s)
}

func (s ModifyTimerGroupRequestConfigTimersSegmentTimers) GoString() string {
	return s.String()
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetAppointmentTimer() *int64 {
	return s.AppointmentTimer
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetLockScreenTime() *int32 {
	return s.LockScreenTime
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetProcessWhitelist() []*string {
	return s.ProcessWhitelist
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetResetType() *string {
	return s.ResetType
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetTimerOrder() *int32 {
	return s.TimerOrder
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetTimezone() *string {
	return s.Timezone
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetAppointmentTimer(v int64) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.AppointmentTimer = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetEndCronExpression(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.EndCronExpression = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetEnforce(v bool) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.Enforce = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetImageId(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.ImageId = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetInterval(v int32) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.Interval = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetLockScreenTime(v int32) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.LockScreenTime = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetNotificationTime(v int32) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.NotificationTime = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetOperationType(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.OperationType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetProcessWhitelist(v []*string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.ProcessWhitelist = v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetResetType(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.ResetType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetStartCronExpression(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.StartCronExpression = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetTimerOrder(v int32) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.TimerOrder = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetTimezone(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.Timezone = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) SetTriggerType(v string) *ModifyTimerGroupRequestConfigTimersSegmentTimers {
	s.TriggerType = &v
	return s
}

func (s *ModifyTimerGroupRequestConfigTimersSegmentTimers) Validate() error {
	return dara.Validate(s)
}
