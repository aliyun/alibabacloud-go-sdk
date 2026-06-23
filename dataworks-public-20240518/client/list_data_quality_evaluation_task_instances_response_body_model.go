// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityEvaluationTaskInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) *ListDataQualityEvaluationTaskInstancesResponseBody
	GetPagingInfo() *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo
	SetRequestId(v string) *ListDataQualityEvaluationTaskInstancesResponseBody
	GetRequestId() *string
}

type ListDataQualityEvaluationTaskInstancesResponseBody struct {
	// The pagination query result of quality evaluation task instances.
	PagingInfo *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The API request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) GetPagingInfo() *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) SetPagingInfo(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) *ListDataQualityEvaluationTaskInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) SetRequestId(v string) *ListDataQualityEvaluationTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo struct {
	// The successfully triggered TaskInstance.
	DataQualityEvaluationTaskInstances []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances `json:"DataQualityEvaluationTaskInstances,omitempty" xml:"DataQualityEvaluationTaskInstances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) GetDataQualityEvaluationTaskInstances() []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	return s.DataQualityEvaluationTaskInstances
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) SetDataQualityEvaluationTaskInstances(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo {
	s.DataQualityEvaluationTaskInstances = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) SetPageSize(v int32) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) Validate() error {
	if s.DataQualityEvaluationTaskInstances != nil {
		for _, item := range s.DataQualityEvaluationTaskInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances struct {
	// The creation time of the task instance.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The end time of the task instance.
	//
	// example:
	//
	// 1710239005403
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the quality check task instance.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The parameter settings used during the actual runtime.
	//
	// example:
	//
	// {
	//
	//   "bizdate": "20240517",
	//
	//   "triggerTime": "1710239005403"
	//
	// }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the DataWorks project workspace.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The current running status.
	//
	// - Running
	//
	// - Error
	//
	// - Passed
	//
	// - Warned
	//
	// - Critical
	//
	// example:
	//
	// Critical
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The snapshot of the data quality evaluation task when the evaluation starts.
	Task *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// The context information when the instance is triggered.
	//
	// example:
	//
	// {
	//
	//   "TriggerClientId": 10001,
	//
	//   "TriggerClient": "CWF2"
	//
	// }
	TriggerContext *string `json:"TriggerContext,omitempty" xml:"TriggerContext,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetParameters() *string {
	return s.Parameters
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetStatus() *string {
	return s.Status
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetTask() *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	return s.Task
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GetTriggerContext() *string {
	return s.TriggerContext
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetCreateTime(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.CreateTime = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetFinishTime(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.FinishTime = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetId(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.Id = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetParameters(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.Parameters = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetProjectId(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetStatus(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.Status = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetTask(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.Task = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) SetTriggerContext(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances {
	s.TriggerContext = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask struct {
	// The description of the quality monitoring task.
	//
	// example:
	//
	// This is a daily run data quality evaluation plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The callback settings in the lifecycle of the data quality evaluation task instance. Currently, only one Hook for blocking scheduled tasks is supported.
	Hooks []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality evaluation task.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the quality monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// Quality verification task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification subscription configuration.
	Notifications *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The ID of the project workspace.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Settings used when accessing data sources. Currently, only specifying the YARN queue of EMR and specifying the SQL engine as SPARK_SQL|KYUUBI|PRESTO_SQL|HIVE_SQL when collecting EMR tables are supported.
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitoring object of the data quality evaluation task.
	Target *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the data quality evaluation task.
	Trigger *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetDescription() *string {
	return s.Description
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetHooks() []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks {
	return s.Hooks
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetName() *string {
	return s.Name
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetNotifications() *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications {
	return s.Notifications
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetTarget() *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget {
	return s.Target
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GetTrigger() *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger {
	return s.Trigger
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetDescription(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Description = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetHooks(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Hooks = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetId(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Id = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetName(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Name = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetNotifications(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Notifications = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetProjectId(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetRuntimeConf(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.RuntimeConf = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetTarget(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Target = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetTrigger(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Trigger = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) Validate() error {
	if s.Hooks != nil {
		for _, item := range s.Hooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Notifications != nil {
		if err := s.Notifications.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks struct {
	// The trigger condition of the Hook.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the subsequent processing action.
	//
	// - BlockTaskInstance: Blocks the execution of the DataWorks task instance.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) GetCondition() *string {
	return s.Condition
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) GetType() *string {
	return s.Type
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) SetCondition(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) SetType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks {
	s.Type = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications struct {
	// The trigger condition of the notification.
	//
	// example:
	//
	// ${severity} == "High"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The notification settings.
	Notifications []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) GetCondition() *string {
	return s.Condition
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) GetNotifications() []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications {
	return s.Notifications
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) SetCondition(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) SetNotifications(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications {
	s.Notifications = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) Validate() error {
	if s.Notifications != nil {
		for _, item := range s.Notifications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications struct {
	// The alert recipient settings.
	NofiticationReceivers []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers `json:"NofiticationReceivers,omitempty" xml:"NofiticationReceivers,omitempty" type:"Repeated"`
	// The alert method.
	NotificationChannels []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) GetNofiticationReceivers() []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers {
	return s.NofiticationReceivers
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) GetNotificationChannels() []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) SetNofiticationReceivers(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications {
	s.NofiticationReceivers = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) SetNotificationChannels(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) Validate() error {
	if s.NofiticationReceivers != nil {
		for _, item := range s.NofiticationReceivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotificationChannels != nil {
		for _, item := range s.NotificationChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers struct {
	// The extension information, in JSON format. For example, a DingTalk bot supports at-all.
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient.
	//
	// - AliUid - The UID of the Alibaba Cloud account
	//
	// - WebhookUrl - A custom webhook URL
	//
	// - DingdingUrl - The URL of a DingTalk bot
	//
	// - FeishuUrl - The URL of a Lark bot
	//
	// - WeixinUrl - The URL of a WeCom bot
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipient.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) SetExtension(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers {
	s.Extension = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) SetReceiverType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) SetReceiverValues(v []*string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels struct {
	// The alert method.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget struct {
	// For a dataset of table type, the database type to which the table belongs.
	//
	// - maxcompute
	//
	// - emr
	//
	// - cdh
	//
	// - hologres
	//
	// - analyticdb_for_postgresql
	//
	// - analyticdb_for_mysql
	//
	// - starrocks
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The partition settings of the partitioned table.
	//
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The unique ID of the table in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The type of the monitoring object.
	//
	// - Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) GetType() *string {
	return s.Type
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) SetDatabaseType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget {
	s.DatabaseType = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) SetPartitionSpec(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget {
	s.PartitionSpec = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) SetTableGuid(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) SetType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget {
	s.Type = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger struct {
	// The IDs of scheduled task nodes.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The type of event that can trigger the execution of the quality evaluation task.
	//
	// - ByScheduledTaskInstance: A scheduled instance runs successfully.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) GetType() *string {
	return s.Type
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) SetTaskIds(v []*int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) SetType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger {
	s.Type = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) Validate() error {
	return dara.Validate(s)
}
