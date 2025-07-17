// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronBeginDate(v string) *UpdateTaskFlowScheduleRequest
	GetCronBeginDate() *string
	SetCronEndDate(v string) *UpdateTaskFlowScheduleRequest
	GetCronEndDate() *string
	SetCronStr(v string) *UpdateTaskFlowScheduleRequest
	GetCronStr() *string
	SetCronType(v string) *UpdateTaskFlowScheduleRequest
	GetCronType() *string
	SetDagId(v int64) *UpdateTaskFlowScheduleRequest
	GetDagId() *int64
	SetScheduleParam(v string) *UpdateTaskFlowScheduleRequest
	GetScheduleParam() *string
	SetScheduleSwitch(v bool) *UpdateTaskFlowScheduleRequest
	GetScheduleSwitch() *bool
	SetTid(v int64) *UpdateTaskFlowScheduleRequest
	GetTid() *int64
	SetTimeZoneId(v string) *UpdateTaskFlowScheduleRequest
	GetTimeZoneId() *string
	SetTriggerType(v string) *UpdateTaskFlowScheduleRequest
	GetTriggerType() *string
}

type UpdateTaskFlowScheduleRequest struct {
	// The start of the time range for scheduling.
	//
	// example:
	//
	// CronBeginDate_test
	CronBeginDate *string `json:"CronBeginDate,omitempty" xml:"CronBeginDate,omitempty"`
	// The end of the time range for scheduling.
	//
	// example:
	//
	// CronEndDate_test
	CronEndDate *string `json:"CronEndDate,omitempty" xml:"CronEndDate,omitempty"`
	// The cron expression for timed scheduling.
	//
	// example:
	//
	// CronStr_test
	CronStr *string `json:"CronStr,omitempty" xml:"CronStr,omitempty"`
	// The type of the scheduling cycle. Valid values:
	//
	// 	- **MINUTE**: scheduling by minute
	//
	// 	- **HOUR**: scheduling by hour
	//
	// 	- **DAY**: scheduling by day
	//
	// 	- **WEEK**: scheduling by week
	//
	// 	- **MONTH**: scheduling by month
	//
	// example:
	//
	// HOUR
	CronType *string `json:"CronType,omitempty" xml:"CronType,omitempty"`
	// The ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The event scheduling configuration. The value of this parameter is a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// ScheduleParam_test
	ScheduleParam *string `json:"ScheduleParam,omitempty" xml:"ScheduleParam,omitempty"`
	// Specifies whether to enable scheduling. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// This parameter is required.
	//
	// example:
	//
	// Disable
	ScheduleSwitch *bool `json:"ScheduleSwitch,omitempty" xml:"ScheduleSwitch,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The time zone. The default time zone is UTC+8 (Asia/Shanghai).
	//
	// example:
	//
	// Asia/Shanghai
	TimeZoneId *string `json:"TimeZoneId,omitempty" xml:"TimeZoneId,omitempty"`
	// The mode in which the task flow is triggered. Valid values:
	//
	// 	- **Cron**: The task flow is triggered based on timed scheduling.
	//
	// 	- **Event**: The task flow is triggered by events.
	//
	// This parameter is required.
	//
	// example:
	//
	// Event
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s UpdateTaskFlowScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowScheduleRequest) GetCronBeginDate() *string {
	return s.CronBeginDate
}

func (s *UpdateTaskFlowScheduleRequest) GetCronEndDate() *string {
	return s.CronEndDate
}

func (s *UpdateTaskFlowScheduleRequest) GetCronStr() *string {
	return s.CronStr
}

func (s *UpdateTaskFlowScheduleRequest) GetCronType() *string {
	return s.CronType
}

func (s *UpdateTaskFlowScheduleRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowScheduleRequest) GetScheduleParam() *string {
	return s.ScheduleParam
}

func (s *UpdateTaskFlowScheduleRequest) GetScheduleSwitch() *bool {
	return s.ScheduleSwitch
}

func (s *UpdateTaskFlowScheduleRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowScheduleRequest) GetTimeZoneId() *string {
	return s.TimeZoneId
}

func (s *UpdateTaskFlowScheduleRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateTaskFlowScheduleRequest) SetCronBeginDate(v string) *UpdateTaskFlowScheduleRequest {
	s.CronBeginDate = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetCronEndDate(v string) *UpdateTaskFlowScheduleRequest {
	s.CronEndDate = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetCronStr(v string) *UpdateTaskFlowScheduleRequest {
	s.CronStr = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetCronType(v string) *UpdateTaskFlowScheduleRequest {
	s.CronType = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetDagId(v int64) *UpdateTaskFlowScheduleRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetScheduleParam(v string) *UpdateTaskFlowScheduleRequest {
	s.ScheduleParam = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetScheduleSwitch(v bool) *UpdateTaskFlowScheduleRequest {
	s.ScheduleSwitch = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetTid(v int64) *UpdateTaskFlowScheduleRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetTimeZoneId(v string) *UpdateTaskFlowScheduleRequest {
	s.TimeZoneId = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) SetTriggerType(v string) *UpdateTaskFlowScheduleRequest {
	s.TriggerType = &v
	return s
}

func (s *UpdateTaskFlowScheduleRequest) Validate() error {
	return dara.Validate(s)
}
