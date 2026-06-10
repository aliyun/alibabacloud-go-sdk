// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTimerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeTimerGroupResponseBodyData) *DescribeTimerGroupResponseBody
	GetData() *DescribeTimerGroupResponseBodyData
	SetRequestId(v string) *DescribeTimerGroupResponseBody
	GetRequestId() *string
}

type DescribeTimerGroupResponseBody struct {
	// The details of the timer group.
	Data *DescribeTimerGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTimerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupResponseBody) GetData() *DescribeTimerGroupResponseBodyData {
	return s.Data
}

func (s *DescribeTimerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTimerGroupResponseBody) SetData(v *DescribeTimerGroupResponseBodyData) *DescribeTimerGroupResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTimerGroupResponseBody) SetRequestId(v string) *DescribeTimerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTimerGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTimerGroupResponseBodyData struct {
	// The number of resources associated with the timer group.
	//
	// example:
	//
	// 50
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// A map of associated resource counts, categorized by resource type.
	BindCountMap map[string]*int32 `json:"BindCountMap,omitempty" xml:"BindCountMap,omitempty"`
	// The configurations of the scheduled tasks.
	ConfigTimers []*DescribeTimerGroupResponseBodyDataConfigTimers `json:"ConfigTimers,omitempty" xml:"ConfigTimers,omitempty" type:"Repeated"`
	// The description of the timer group.
	//
	// example:
	//
	// Scheduled task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the timer group.
	//
	// example:
	//
	// cg-75aazkg2tnqb2*****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// An internal code used by the frontend to display the description of a system-scheduled task.
	//
	// example:
	//
	// INNER_TIMER_10_MINUTES_HIBERNATE_NO_UPDATE_DESC
	InnerTimerDesc *string `json:"InnerTimerDesc,omitempty" xml:"InnerTimerDesc,omitempty"`
	// An internal code used by the frontend to display the name of a system-scheduled task.
	//
	// example:
	//
	// INNER_TIMER_10_MINUTES_HIBERNATE_NO_UPDATE
	InnerTimerName *string `json:"InnerTimerName,omitempty" xml:"InnerTimerName,omitempty"`
	// Indicates that resources cannot be bound to or unbound from this timer group.
	IsBind *bool `json:"IsBind,omitempty" xml:"IsBind,omitempty"`
	// Indicates that this timer group cannot be modified.
	IsUpdate *bool `json:"IsUpdate,omitempty" xml:"IsUpdate,omitempty"`
	// The name of the timer group.
	//
	// example:
	//
	// Scheduled task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The product type that the timer group supports.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The status of the timer group.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the timer group.
	//
	// example:
	//
	// Timer
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTimerGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupResponseBodyData) GetBindCount() *int32 {
	return s.BindCount
}

func (s *DescribeTimerGroupResponseBodyData) GetBindCountMap() map[string]*int32 {
	return s.BindCountMap
}

func (s *DescribeTimerGroupResponseBodyData) GetConfigTimers() []*DescribeTimerGroupResponseBodyDataConfigTimers {
	return s.ConfigTimers
}

func (s *DescribeTimerGroupResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeTimerGroupResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeTimerGroupResponseBodyData) GetInnerTimerDesc() *string {
	return s.InnerTimerDesc
}

func (s *DescribeTimerGroupResponseBodyData) GetInnerTimerName() *string {
	return s.InnerTimerName
}

func (s *DescribeTimerGroupResponseBodyData) GetIsBind() *bool {
	return s.IsBind
}

func (s *DescribeTimerGroupResponseBodyData) GetIsUpdate() *bool {
	return s.IsUpdate
}

func (s *DescribeTimerGroupResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeTimerGroupResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeTimerGroupResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeTimerGroupResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeTimerGroupResponseBodyData) SetBindCount(v int32) *DescribeTimerGroupResponseBodyData {
	s.BindCount = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetBindCountMap(v map[string]*int32) *DescribeTimerGroupResponseBodyData {
	s.BindCountMap = v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetConfigTimers(v []*DescribeTimerGroupResponseBodyDataConfigTimers) *DescribeTimerGroupResponseBodyData {
	s.ConfigTimers = v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetDescription(v string) *DescribeTimerGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetGroupId(v string) *DescribeTimerGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetInnerTimerDesc(v string) *DescribeTimerGroupResponseBodyData {
	s.InnerTimerDesc = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetInnerTimerName(v string) *DescribeTimerGroupResponseBodyData {
	s.InnerTimerName = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetIsBind(v bool) *DescribeTimerGroupResponseBodyData {
	s.IsBind = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetIsUpdate(v bool) *DescribeTimerGroupResponseBodyData {
	s.IsUpdate = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetName(v string) *DescribeTimerGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetProductType(v string) *DescribeTimerGroupResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetStatus(v string) *DescribeTimerGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) SetType(v string) *DescribeTimerGroupResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyData) Validate() error {
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

type DescribeTimerGroupResponseBodyDataConfigTimers struct {
	// Whether to allow end users to configure the scheduled task.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The cron expression for the scheduled task.
	//
	// example:
	//
	// 0 0 16 ? 	- 1,2,3,4,5,6,7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// Specifies whether to force the execution of the scheduled task. If set to `true`, the task runs regardless of the cloud computer\\"s status or connection state.
	//
	// example:
	//
	// false
	Enforce *bool `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The interval. Unit: minutes.
	//
	// example:
	//
	// 10
	Interval         *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	NotificationTime *int32 `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	// The operation to perform when `TimerType` is set to `NoConnect`.
	//
	// example:
	//
	// Shutdown
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The process whitelist for smart detection. A scheduled task based on user inactivity does not run if a whitelisted process is running.
	ProcessWhitelist []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	// The reset type for the scheduled reset task.
	//
	// example:
	//
	// RESET_TYPE_SYSTEM
	ResetType     *string                                                        `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	SegmentTimers []*DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers `json:"SegmentTimers,omitempty" xml:"SegmentTimers,omitempty" type:"Repeated"`
	// The type of the scheduled task.
	//
	// example:
	//
	// TimerBoot
	TimerType *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
	// The detection method for user inactivity.
	//
	// example:
	//
	// Standard
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s DescribeTimerGroupResponseBodyDataConfigTimers) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupResponseBodyDataConfigTimers) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetAllowClientSetting() *bool {
	return s.AllowClientSetting
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetProcessWhitelist() []*string {
	return s.ProcessWhitelist
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetResetType() *string {
	return s.ResetType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetSegmentTimers() []*DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	return s.SegmentTimers
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetTimerType() *string {
	return s.TimerType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetAllowClientSetting(v bool) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetCronExpression(v string) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.CronExpression = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetEnforce(v bool) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetInterval(v int32) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.Interval = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetNotificationTime(v int32) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.NotificationTime = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetOperationType(v string) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetProcessWhitelist(v []*string) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.ProcessWhitelist = v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetResetType(v string) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetSegmentTimers(v []*DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.SegmentTimers = v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetTimerType(v string) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.TimerType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) SetTriggerType(v string) *DescribeTimerGroupResponseBodyDataConfigTimers {
	s.TriggerType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimers) Validate() error {
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

type DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers struct {
	// The time to execute the scheduled task, specified as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1764660600967
	AppointmentTimer  *int64  `json:"AppointmentTimer,omitempty" xml:"AppointmentTimer,omitempty"`
	CreateSnapshot    *bool   `json:"CreateSnapshot,omitempty" xml:"CreateSnapshot,omitempty"`
	EndCronExpression *string `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	Enforce           *bool   `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	// The image ID for a scheduled image-change task.
	//
	// example:
	//
	// m-5b0vjqbiqu010XXXXXX
	ImageId    *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Interval   *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IpSegments []*string `json:"IpSegments,omitempty" xml:"IpSegments,omitempty" type:"Repeated"`
	// The duration of user inactivity, in seconds, before the screen locks. This feature applies only to cloud computers joined to an Active Directory (AD) domain.
	//
	// example:
	//
	// 1800
	LockScreenTime   *int32  `json:"LockScreenTime,omitempty" xml:"LockScreenTime,omitempty"`
	NotificationTime *int32  `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	OperationType    *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// KB5082063
	PatchId                      *string   `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	ProcessWhitelist             []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	ResetType                    *string   `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	StartCronExpression          *string   `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
	TimerOrder                   *int32    `json:"TimerOrder,omitempty" xml:"TimerOrder,omitempty"`
	Timezone                     *string   `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	TriggerType                  *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	VerificationNotificationTime *int32    `json:"VerificationNotificationTime,omitempty" xml:"VerificationNotificationTime,omitempty"`
	VerificationTime             *int32    `json:"VerificationTime,omitempty" xml:"VerificationTime,omitempty"`
}

func (s DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetAppointmentTimer() *int64 {
	return s.AppointmentTimer
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetCreateSnapshot() *bool {
	return s.CreateSnapshot
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetIpSegments() []*string {
	return s.IpSegments
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetLockScreenTime() *int32 {
	return s.LockScreenTime
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetPatchId() *string {
	return s.PatchId
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetProcessWhitelist() []*string {
	return s.ProcessWhitelist
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetResetType() *string {
	return s.ResetType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetTimerOrder() *int32 {
	return s.TimerOrder
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetTriggerType() *string {
	return s.TriggerType
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetVerificationNotificationTime() *int32 {
	return s.VerificationNotificationTime
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetVerificationTime() *int32 {
	return s.VerificationTime
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetAppointmentTimer(v int64) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.AppointmentTimer = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetCreateSnapshot(v bool) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.CreateSnapshot = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetEndCronExpression(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.EndCronExpression = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetEnforce(v bool) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetImageId(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.ImageId = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetInterval(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.Interval = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetIpSegments(v []*string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.IpSegments = v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetLockScreenTime(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.LockScreenTime = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetNotificationTime(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.NotificationTime = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetOperationType(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetPatchId(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.PatchId = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetProcessWhitelist(v []*string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.ProcessWhitelist = v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetResetType(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetStartCronExpression(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.StartCronExpression = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetTimerOrder(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.TimerOrder = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetTimezone(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.Timezone = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetTriggerType(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.TriggerType = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetVerificationNotificationTime(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.VerificationNotificationTime = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetVerificationTime(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.VerificationTime = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) Validate() error {
	return dara.Validate(s)
}
