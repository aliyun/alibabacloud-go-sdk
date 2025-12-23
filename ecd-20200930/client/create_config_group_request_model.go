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
	// The scheduled task groups.
	ConfigTimers []*CreateConfigGroupRequestConfigTimers `json:"ConfigTimers,omitempty" xml:"ConfigTimers,omitempty" type:"Repeated"`
	// The description of the configuration group.
	//
	// example:
	//
	// ScheduledTask
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the configuration group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ScheduledTask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service type of the configuration group.
	//
	// Valid value:
	//
	// 	- CLOUD_DESKTOP: the cloud computer service.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the region. Set the value to `cn-shanghai`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The group type.
	//
	// Valid value:
	//
	// 	- Timer: a scheduled task group.
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
	// Specifies whether to allow end users to configure the scheduled task.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression specified in the scheduled task.
	//
	// >  The time must be in UTC. For example, for 24:00 (UTC+8), you must set the value to 0 0 16 ? \\	- 1,2,3,4,5,6,7
	//
	// example:
	//
	// 0 0 16 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to forcefully execute the scheduled task.
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
	// Valid values:
	//
	// 	- RESET_TYPE_SYSTEM: resets only the system disk.
	//
	// 	- RESET_TYPE_USER_DISK: resets only the data disk.
	//
	// 	- RESET_TYPE_BOTH: resets the system and data disks.
	//
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType     *string                                              `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	SegmentTimers []*CreateConfigGroupRequestConfigTimersSegmentTimers `json:"SegmentTimers,omitempty" xml:"SegmentTimers,omitempty" type:"Repeated"`
	// The scheduled task type.
	//
	// Valid values:
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
	// 	- TimerReboot: scheduled restart.
	//
	// This parameter is required.
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

func (s CreateConfigGroupRequestConfigTimersSegmentTimers) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigGroupRequestConfigTimersSegmentTimers) GoString() string {
	return s.String()
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetAppointmentTimer() *int64 {
	return s.AppointmentTimer
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

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetLockScreenTime() *int32 {
	return s.LockScreenTime
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) GetOperationType() *string {
	return s.OperationType
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

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) SetAppointmentTimer(v int64) *CreateConfigGroupRequestConfigTimersSegmentTimers {
	s.AppointmentTimer = &v
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

func (s *CreateConfigGroupRequestConfigTimersSegmentTimers) Validate() error {
	return dara.Validate(s)
}
