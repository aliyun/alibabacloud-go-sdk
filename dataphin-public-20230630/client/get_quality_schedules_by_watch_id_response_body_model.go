// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualitySchedulesByWatchIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualitySchedulesByWatchIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualitySchedulesByWatchIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualitySchedulesByWatchIdResponseBody
	GetMessage() *string
	SetQualityScheduleList(v []*GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) *GetQualitySchedulesByWatchIdResponseBody
	GetQualityScheduleList() []*GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList
	SetRequestId(v string) *GetQualitySchedulesByWatchIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualitySchedulesByWatchIdResponseBody
	GetSuccess() *bool
}

type GetQualitySchedulesByWatchIdResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of schedule objects.
	QualityScheduleList []*GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList `json:"QualityScheduleList,omitempty" xml:"QualityScheduleList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualitySchedulesByWatchIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualitySchedulesByWatchIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualitySchedulesByWatchIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualitySchedulesByWatchIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualitySchedulesByWatchIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualitySchedulesByWatchIdResponseBody) GetQualityScheduleList() []*GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	return s.QualityScheduleList
}

func (s *GetQualitySchedulesByWatchIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualitySchedulesByWatchIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualitySchedulesByWatchIdResponseBody) SetCode(v string) *GetQualitySchedulesByWatchIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBody) SetHttpStatusCode(v int32) *GetQualitySchedulesByWatchIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBody) SetMessage(v string) *GetQualitySchedulesByWatchIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBody) SetQualityScheduleList(v []*GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) *GetQualitySchedulesByWatchIdResponseBody {
	s.QualityScheduleList = v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBody) SetRequestId(v string) *GetQualitySchedulesByWatchIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBody) SetSuccess(v bool) *GetQualitySchedulesByWatchIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBody) Validate() error {
	if s.QualityScheduleList != nil {
		for _, item := range s.QualityScheduleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList struct {
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The cron expression for timed scheduling.
	//
	// example:
	//
	// 	- 	- 1/	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The schedule object ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the schedule is referenced by a rule.
	IsRefByRule *bool `json:"IsRefByRule,omitempty" xml:"IsRefByRule,omitempty"`
	// The user ID of the last modifier.
	//
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The schedule object name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The partition expression. A custom expression.
	//
	// example:
	//
	// ds=${yyyyMMdd}
	PartitionExpression *string `json:"PartitionExpression,omitempty" xml:"PartitionExpression,omitempty"`
	// The partition type. Valid values:
	//
	// - EVERY_DAY: every day.
	//
	// - PRE_DAY: yesterday.
	//
	// - TODAY: today.
	//
	// - FIRST_DAY_OF_WEEK: first day of the week (Sunday).
	//
	// - CUSTOM: custom.
	//
	// example:
	//
	// CUSTOM
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// The interval type for timed scheduling. Valid values:
	//
	// - DAILY: day.
	//
	// - WEEKLY: week.
	//
	// - MONTHLY: month.
	//
	// - HOURLY: hour.
	//
	// - MINUTELY: minute.
	//
	// example:
	//
	// DAILY
	PeriodScheduleIntervalType *string `json:"PeriodScheduleIntervalType,omitempty" xml:"PeriodScheduleIntervalType,omitempty"`
	// The interval values for timed scheduling.
	PeriodScheduleParamList []*string `json:"PeriodScheduleParamList,omitempty" xml:"PeriodScheduleParamList,omitempty" type:"Repeated"`
	// The trigger method for fixed task triggers. Valid values:
	//
	// - ALL_TASKS_FINISHED
	//
	// - ONE_TASKS_FINISHED
	//
	// - PRE_ONE_TASKS_START.
	//
	// example:
	//
	// ONE_TASKS_FINISHED
	StaticTaskTriggerType *string `json:"StaticTaskTriggerType,omitempty" xml:"StaticTaskTriggerType,omitempty"`
	// The list of trigger nodes for trigger-based scheduling.
	TriggerNodeList []*string `json:"TriggerNodeList,omitempty" xml:"TriggerNodeList,omitempty" type:"Repeated"`
	// The trigger method for trigger-based scheduling. Valid values:
	//
	// - STATIC_TASK_TRIGGER: fixed task trigger.
	//
	// - CODE_CHECK_TRIGGER: code check trigger.
	//
	// example:
	//
	// STATIC_TASK_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The schedule type. Valid values:
	//
	// - PERIOD_SCHEDULE: timed scheduling.
	//
	// - MANUAL_SCHEDULE: manual trigger.
	//
	// - CODE_CHECK_TRIGGER: code check trigger.
	//
	// - STATIC_TASK_TRIGGER: fixed task trigger.
	//
	// - DEPENDENCY_SCHEDULE: dependency scheduling.
	//
	// example:
	//
	// PERIOD_SCHEDULE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The validation scope. Valid values:
	//
	// - TASK_REFERRED_PARTITION: task-updated partition.
	//
	// - USER_DEFINED_PARTITION: custom partition.
	//
	// example:
	//
	// TASK_REFERRED_PARTITION
	ValidatePartitionType *string `json:"ValidatePartitionType,omitempty" xml:"ValidatePartitionType,omitempty"`
	// The monitored object ID.
	//
	// example:
	//
	// 22
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) String() string {
	return dara.Prettify(s)
}

func (s GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GoString() string {
	return s.String()
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetCreator() *string {
	return s.Creator
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetId() *int64 {
	return s.Id
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetIsRefByRule() *bool {
	return s.IsRefByRule
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetName() *string {
	return s.Name
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetPartitionExpression() *string {
	return s.PartitionExpression
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetPartitionType() *string {
	return s.PartitionType
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetPeriodScheduleIntervalType() *string {
	return s.PeriodScheduleIntervalType
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetPeriodScheduleParamList() []*string {
	return s.PeriodScheduleParamList
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetStaticTaskTriggerType() *string {
	return s.StaticTaskTriggerType
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetTriggerNodeList() []*string {
	return s.TriggerNodeList
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetType() *string {
	return s.Type
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetValidatePartitionType() *string {
	return s.ValidatePartitionType
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetCreateTime(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.CreateTime = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetCreator(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.Creator = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetCronExpression(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.CronExpression = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetId(v int64) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.Id = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetIsRefByRule(v bool) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.IsRefByRule = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetModifier(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.Modifier = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetModifyTime(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.ModifyTime = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetName(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.Name = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetPartitionExpression(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.PartitionExpression = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetPartitionType(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.PartitionType = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetPeriodScheduleIntervalType(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.PeriodScheduleIntervalType = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetPeriodScheduleParamList(v []*string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.PeriodScheduleParamList = v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetStaticTaskTriggerType(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.StaticTaskTriggerType = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetTriggerNodeList(v []*string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.TriggerNodeList = v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetTriggerType(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.TriggerType = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetType(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.Type = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetValidatePartitionType(v string) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.ValidatePartitionType = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) SetWatchId(v int64) *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList {
	s.WatchId = &v
	return s
}

func (s *GetQualitySchedulesByWatchIdResponseBodyQualityScheduleList) Validate() error {
	return dara.Validate(s)
}
