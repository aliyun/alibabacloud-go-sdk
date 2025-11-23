// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowsByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTaskFlowsByPageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowsByPageResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowsByPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowsByPageResponseBody
	GetSuccess() *bool
	SetTaskFlowList(v *ListTaskFlowsByPageResponseBodyTaskFlowList) *ListTaskFlowsByPageResponseBody
	GetTaskFlowList() *ListTaskFlowsByPageResponseBodyTaskFlowList
	SetTotalCount(v int32) *ListTaskFlowsByPageResponseBody
	GetTotalCount() *int32
}

type ListTaskFlowsByPageResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6CB28697-BFE2-5739-9228-3971990E982C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the returned task flows.
	TaskFlowList *ListTaskFlowsByPageResponseBodyTaskFlowList `json:"TaskFlowList,omitempty" xml:"TaskFlowList,omitempty" type:"Struct"`
	// The total number of entries returned.
	//
	// example:
	//
	// 24
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskFlowsByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowsByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowsByPageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowsByPageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowsByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowsByPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowsByPageResponseBody) GetTaskFlowList() *ListTaskFlowsByPageResponseBodyTaskFlowList {
	return s.TaskFlowList
}

func (s *ListTaskFlowsByPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskFlowsByPageResponseBody) SetErrorCode(v string) *ListTaskFlowsByPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBody) SetErrorMessage(v string) *ListTaskFlowsByPageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBody) SetRequestId(v string) *ListTaskFlowsByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBody) SetSuccess(v bool) *ListTaskFlowsByPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBody) SetTaskFlowList(v *ListTaskFlowsByPageResponseBodyTaskFlowList) *ListTaskFlowsByPageResponseBody {
	s.TaskFlowList = v
	return s
}

func (s *ListTaskFlowsByPageResponseBody) SetTotalCount(v int32) *ListTaskFlowsByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBody) Validate() error {
	if s.TaskFlowList != nil {
		if err := s.TaskFlowList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowsByPageResponseBodyTaskFlowList struct {
	TaskFlow []*ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow `json:"TaskFlow,omitempty" xml:"TaskFlow,omitempty" type:"Repeated"`
}

func (s ListTaskFlowsByPageResponseBodyTaskFlowList) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowsByPageResponseBodyTaskFlowList) GoString() string {
	return s.String()
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowList) GetTaskFlow() []*ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	return s.TaskFlow
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowList) SetTaskFlow(v []*ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) *ListTaskFlowsByPageResponseBodyTaskFlowList {
	s.TaskFlow = v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowList) Validate() error {
	if s.TaskFlow != nil {
		for _, item := range s.TaskFlow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow struct {
	// The ID of the user who created the task flow.
	//
	// example:
	//
	// 51****
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The username of the user who created the task flow.
	//
	// example:
	//
	// Creator_NickName
	CreatorNickName *string `json:"CreatorNickName,omitempty" xml:"CreatorNickName,omitempty"`
	// The start time of scheduled scheduling. The task flow is not scheduled before this point in time.
	//
	// example:
	//
	// 1970-01-01
	CronBeginDate *string `json:"CronBeginDate,omitempty" xml:"CronBeginDate,omitempty"`
	// The end time of scheduled scheduling. The task flow is not scheduled after this point in time.
	//
	// example:
	//
	// 2023-01-01
	CronEndDate *string `json:"CronEndDate,omitempty" xml:"CronEndDate,omitempty"`
	// Scheduled Cron.
	//
	// example:
	//
	// 0 0 1 	- 	- ? *
	CronStr *string `json:"CronStr,omitempty" xml:"CronStr,omitempty"`
	// Whether to enable scheduled scheduling.
	//
	// example:
	//
	// false
	CronSwitch *bool `json:"CronSwitch,omitempty" xml:"CronSwitch,omitempty"`
	// Scheduling cycle type. Valid values:
	//
	// - **2**: Hourly scheduling
	//
	// - **3**: Daily scheduling
	//
	// - **4**: Weekly scheduling
	//
	// - **5**: Monthly scheduling
	//
	// example:
	//
	// 2
	CronType *int32 `json:"CronType,omitempty" xml:"CronType,omitempty"`
	// The name of the task flow.
	//
	// example:
	//
	// poc_task_test
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The user ID of the task flow owner.
	//
	// example:
	//
	// 12***89
	DagOwnerId *string `json:"DagOwnerId,omitempty" xml:"DagOwnerId,omitempty"`
	// The username of the owner of the task flow.
	//
	// example:
	//
	// Owner_NickName
	DagOwnerNickName *string `json:"DagOwnerNickName,omitempty" xml:"DagOwnerNickName,omitempty"`
	// The ID of the last deployment record of the task flow.
	//
	// example:
	//
	// 65***
	DeployId *int64 `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	// The description of the task flow.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the task flow.
	//
	// example:
	//
	// 7***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the last execution of the task flow. Valid values:
	//
	// 	- **0**: invalid
	//
	// 	- **1**: scheduling disabled
	//
	// 	- **2**: waiting to be scheduled
	//
	// example:
	//
	// 0
	LatestInstanceStatus *int32 `json:"LatestInstanceStatus,omitempty" xml:"LatestInstanceStatus,omitempty"`
	// The time when the last execution record was created.
	//
	// example:
	//
	// 2022-04-13
	LatestInstanceTime *string `json:"LatestInstanceTime,omitempty" xml:"LatestInstanceTime,omitempty"`
	// The ID of the application scenario.
	//
	// example:
	//
	// 1245
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Event scheduling configuration, JSON string format.
	//
	// example:
	//
	// {\\"triggerType\\":\\"1\\",\\"specificTime\\":\\"2022-11-15 11:59\\"}
	ScheduleParam *string `json:"ScheduleParam,omitempty" xml:"ScheduleParam,omitempty"`
	// The status of the task flow. Valid values:
	//
	// 	- **0**: invalid
	//
	// 	- **1**: scheduling disabled
	//
	// 	- **2**: waiting to be scheduled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Time zone setting. Default value: East 8(Asia/Shanghai).
	//
	// example:
	//
	// Asia/Shanghai
	TimeZoneId *string `json:"TimeZoneId,omitempty" xml:"TimeZoneId,omitempty"`
	// The trigger type. Valid values:
	//
	// - **0**: Periodic scheduling
	//
	// - **1**: Run manually
	//
	// example:
	//
	// 0
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GoString() string {
	return s.String()
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCreatorNickName() *string {
	return s.CreatorNickName
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCronBeginDate() *string {
	return s.CronBeginDate
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCronEndDate() *string {
	return s.CronEndDate
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCronStr() *string {
	return s.CronStr
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCronSwitch() *bool {
	return s.CronSwitch
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetCronType() *int32 {
	return s.CronType
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetDagName() *string {
	return s.DagName
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetDagOwnerId() *string {
	return s.DagOwnerId
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetDagOwnerNickName() *string {
	return s.DagOwnerNickName
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetDeployId() *int64 {
	return s.DeployId
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetDescription() *string {
	return s.Description
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetId() *int64 {
	return s.Id
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetLatestInstanceStatus() *int32 {
	return s.LatestInstanceStatus
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetLatestInstanceTime() *string {
	return s.LatestInstanceTime
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetScheduleParam() *string {
	return s.ScheduleParam
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetStatus() *int32 {
	return s.Status
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetTimeZoneId() *string {
	return s.TimeZoneId
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCreatorId(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CreatorId = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCreatorNickName(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CreatorNickName = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCronBeginDate(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CronBeginDate = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCronEndDate(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CronEndDate = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCronStr(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CronStr = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCronSwitch(v bool) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CronSwitch = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetCronType(v int32) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.CronType = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetDagName(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.DagName = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetDagOwnerId(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.DagOwnerId = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetDagOwnerNickName(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.DagOwnerNickName = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetDeployId(v int64) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.DeployId = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetDescription(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.Description = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetId(v int64) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.Id = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetLatestInstanceStatus(v int32) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.LatestInstanceStatus = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetLatestInstanceTime(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.LatestInstanceTime = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetScenarioId(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.ScenarioId = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetScheduleParam(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.ScheduleParam = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetStatus(v int32) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.Status = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetTimeZoneId(v string) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.TimeZoneId = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) SetTriggerType(v int32) *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow {
	s.TriggerType = &v
	return s
}

func (s *ListTaskFlowsByPageResponseBodyTaskFlowListTaskFlow) Validate() error {
	return dara.Validate(s)
}
