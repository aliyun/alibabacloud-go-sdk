// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityEvaluationTask(v *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) *GetDataQualityEvaluationTaskResponseBody
	GetDataQualityEvaluationTask() *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask
	SetRequestId(v string) *GetDataQualityEvaluationTaskResponseBody
	GetRequestId() *string
}

type GetDataQualityEvaluationTaskResponseBody struct {
	// Data quality monitoring details.
	DataQualityEvaluationTask *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask `json:"DataQualityEvaluationTask,omitempty" xml:"DataQualityEvaluationTask,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// SDFSDFSDF-SDFSDF-SDFDSF-SDFSDF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBody) GetDataQualityEvaluationTask() *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	return s.DataQualityEvaluationTask
}

func (s *GetDataQualityEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityEvaluationTaskResponseBody) SetDataQualityEvaluationTask(v *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) *GetDataQualityEvaluationTaskResponseBody {
	s.DataQualityEvaluationTask = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBody) SetRequestId(v string) *GetDataQualityEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBody) Validate() error {
	if s.DataQualityEvaluationTask != nil {
		if err := s.DataQualityEvaluationTask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask struct {
	// Data source ID used by the quality monitoring task.
	//
	// example:
	//
	// 45238
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Description of the quality monitoring task.
	//
	// example:
	//
	// The description of the quality monitoring task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Callback settings.
	Hooks []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// Data quality monitoring ID.
	//
	// example:
	//
	// 2178
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name of the quality monitoring task.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Notification subscription configuration.
	Notifications *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// Workspace ID.
	//
	// example:
	//
	// 2626
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Extended configuration. A JSON-formatted string. Only takes effect for EMR-type data quality monitoring.
	//
	// - queue: The YARN queue used when executing EMR data quality validation. Defaults to the queue configured for the current project.
	//
	// - sqlEngine: The SQL engine used when executing EMR data validation.
	//
	//   + HIVE_SQL
	//
	//   + SPARK_SQL
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// Data quality monitoring object.
	Target *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// Trigger configuration of the data quality validation task.
	Trigger *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetDescription() *string {
	return s.Description
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetHooks() []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks {
	return s.Hooks
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetName() *string {
	return s.Name
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetNotifications() *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications {
	return s.Notifications
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetTarget() *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget {
	return s.Target
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) GetTrigger() *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger {
	return s.Trigger
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetDataSourceId(v int64) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.DataSourceId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetDescription(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Description = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetHooks(v []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Hooks = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetId(v int64) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetName(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Name = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetNotifications(v *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Notifications = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetProjectId(v int64) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetRuntimeConf(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.RuntimeConf = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetTarget(v *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Target = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) SetTrigger(v *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask {
	s.Trigger = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTask) Validate() error {
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

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks struct {
	// Hook trigger condition. When this condition is met, the hook action is triggered. Currently, only two types of conditional expressions are supported:
	//
	// - Specify a single group of rule severity type and rule validation status, e.g., `${severity} == "High" AND ${status} == "Critical"`, which means the condition is met if any executed rule with severity High has a validation result of Critical.
	//
	// - Specify multiple groups of rule severity types and rule validation statuses, e.g., `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`, which means the condition is met if any executed rule has severity High with validation result Critical, or severity Normal with validation result Critical, or severity Normal with validation result Error. The severity enum in the conditional expression is consistent with the severity enum in DataQualityRule, and the status enum is consistent with the status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// Hook type. Currently, only one type is supported:
	//
	// - BlockTaskInstance: Blocks the scheduled task from continuing to run. When data quality monitoring is triggered by a scheduled task, after the data quality monitoring completes, Hook.Condition is used to determine whether to block the scheduled task from continuing to run.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) SetCondition(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks {
	s.Condition = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) SetType(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications struct {
	// Notification trigger condition. When this condition is met, the message notification is triggered. Currently, only two types of conditional expressions are supported:
	//
	// - Specify a single group of rule severity type and rule validation status, e.g., `${severity} == "High" AND ${status} == "Critical"`, which means the condition is met if any executed rule with severity High has a validation result of Critical.
	//
	// - Specify multiple groups of rule severity types and rule validation statuses, e.g., `(${severity} == "High"AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`, which means the condition is met if any executed rule has severity High with validation result Critical, or severity Normal with validation result Critical, or severity Normal with validation result Error. The severity enum in the conditional expression is consistent with the severity enum in DataQualityRule, and the status enum is consistent with the status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// Notification settings.
	Notifications []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) GetNotifications() []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications {
	return s.Notifications
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) SetCondition(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications {
	s.Condition = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) SetNotifications(v []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications {
	s.Notifications = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications) Validate() error {
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

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications struct {
	// Notification method.
	NotificationChannels []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// Alert receiver settings.
	NotificationReceivers []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) GetNotificationChannels() []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) GetNotificationReceivers() []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) SetNotificationChannels(v []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) SetNotificationReceivers(v []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotifications) Validate() error {
	if s.NotificationChannels != nil {
		for _, item := range s.NotificationChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotificationReceivers != nil {
		for _, item := range s.NotificationReceivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels struct {
	// Notification method.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers struct {
	// Extended information.
	//
	// example:
	//
	// {  "atAll": true }
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Additional parameter settings when sending alerts. JSON format. Supported keys are as follows:
	//
	// - atAll: Whether to @everyone in the group when sending DingTalk alerts. Takes effect when ReceiverType is DingdingUrl.
	//
	// example:
	//
	// DingdingUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// Alert receiver.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) SetExtension(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget struct {
	// Database type to which the table belongs:
	//
	// - maxcompute
	//
	// - hologres
	//
	// - cdh
	//
	// - analyticdb_for_mysql
	//
	// - starrocks
	//
	// - emr
	//
	// - analyticdb_for_postgresql
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// Partition range setting for data quality monitoring.
	//
	// example:
	//
	// pt=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// Unique ID of the table in Data Map.
	//
	// example:
	//
	// odps.meta_open_api_test_sz.test_partition_tbl
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// Monitoring object type.
	//
	// - Table: Table.
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) SetDatabaseType(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget {
	s.DatabaseType = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) SetPartitionSpec(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget {
	s.PartitionSpec = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) SetTableGuid(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget {
	s.TableGuid = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) SetType(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger struct {
	// List of scheduled task IDs. Valid when Type is ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// Quality monitoring trigger type:
	//
	// - ByManual: Manual trigger
	//
	// - ByScheduledTaskInstance: Scheduled task trigger
	//
	// - ByQualityNode: Quality node trigger
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) GetType() *string {
	return s.Type
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) SetTaskIds(v []*int64) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) SetType(v string) *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger {
	s.Type = &v
	return s
}

func (s *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTrigger) Validate() error {
	return dara.Validate(s)
}
