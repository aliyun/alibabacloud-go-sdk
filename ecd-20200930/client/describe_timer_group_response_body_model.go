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
	// The configuration group.
	Data *DescribeTimerGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
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
	// The number of resources that are bound to the configuration group.
	//
	// example:
	//
	// 50
	BindCount *int32 `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	// The number of bound resources.
	BindCountMap map[string]*int32 `json:"BindCountMap,omitempty" xml:"BindCountMap,omitempty"`
	// The scheduled tasks.
	ConfigTimers []*DescribeTimerGroupResponseBodyDataConfigTimers `json:"ConfigTimers,omitempty" xml:"ConfigTimers,omitempty" type:"Repeated"`
	// The description of the configuration group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the configuration group.
	//
	// example:
	//
	// cg-75aazkg2tnqb2*****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the configuration group.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The service type of the configuration group.
	//
	// Valid value:
	//
	// 	- CLOUD_DESKTOP: the cloud computer service.
	//
	// example:
	//
	// CLOUD_DESKTOP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The state of the configuration group.
	//
	// Valid values:
	//
	// 	- AVAILABLE: The configuration group is available.
	//
	// 	- UNAVAILABLE: The configuration group is deleted.
	//
	// 	- DELETING: The configuration group is being deleted.
	//
	// 	- UPDATING: The configuration group is being modified.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the configuration group.
	//
	// Valid value:
	//
	// 	- Timer: the scheduled task type.
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
	// Indicates whether end users can configure scheduled tasks.
	//
	// example:
	//
	// true
	AllowClientSetting *bool `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	// The CRON expression for the scheduled task.
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
	// The type of the scheduled disconnection task.
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
	// The process whitelist. If whitelisted processes are running, the scheduled task upon inactivity does not take effect.
	ProcessWhitelist []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	// The reset operation of the scheduled task.
	//
	// Valid values:
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
	ResetType     *string                                                        `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	SegmentTimers []*DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers `json:"SegmentTimers,omitempty" xml:"SegmentTimers,omitempty" type:"Repeated"`
	// The type of the scheduled task.
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
	// example:
	//
	// TimerBoot
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
	EndCronExpression   *string   `json:"EndCronExpression,omitempty" xml:"EndCronExpression,omitempty"`
	Enforce             *bool     `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	Interval            *int32    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	NotificationTime    *int32    `json:"NotificationTime,omitempty" xml:"NotificationTime,omitempty"`
	OperationType       *string   `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ProcessWhitelist    []*string `json:"ProcessWhitelist,omitempty" xml:"ProcessWhitelist,omitempty" type:"Repeated"`
	ResetType           *string   `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	StartCronExpression *string   `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
	TimerOrder          *int32    `json:"TimerOrder,omitempty" xml:"TimerOrder,omitempty"`
	Timezone            *string   `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	TriggerType         *string   `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) String() string {
	return dara.Prettify(s)
}

func (s DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GoString() string {
	return s.String()
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetEndCronExpression() *string {
	return s.EndCronExpression
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetEnforce() *bool {
	return s.Enforce
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetInterval() *int32 {
	return s.Interval
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetNotificationTime() *int32 {
	return s.NotificationTime
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) GetOperationType() *string {
	return s.OperationType
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

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetEndCronExpression(v string) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.EndCronExpression = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetEnforce(v bool) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) SetInterval(v int32) *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers {
	s.Interval = &v
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

func (s *DescribeTimerGroupResponseBodyDataConfigTimersSegmentTimers) Validate() error {
	return dara.Validate(s)
}
