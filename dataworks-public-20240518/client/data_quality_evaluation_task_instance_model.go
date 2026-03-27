// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataQualityEvaluationTaskInstance interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *DataQualityEvaluationTaskInstance
	GetCreateTime() *int64
	SetFinishTime(v int64) *DataQualityEvaluationTaskInstance
	GetFinishTime() *int64
	SetId(v int64) *DataQualityEvaluationTaskInstance
	GetId() *int64
	SetStatus(v string) *DataQualityEvaluationTaskInstance
	GetStatus() *string
	SetTask(v *DataQualityEvaluationTaskInstanceTask) *DataQualityEvaluationTaskInstance
	GetTask() *DataQualityEvaluationTaskInstanceTask
}

type DataQualityEvaluationTaskInstance struct {
	// The time at which the instance was generated.
	//
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time at which the instance finished running.
	//
	// example:
	//
	// 1710239005403
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The snapshot of the configurations for the data quality monitoring task when the task starts.
	Task *DataQualityEvaluationTaskInstanceTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s DataQualityEvaluationTaskInstance) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstance) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DataQualityEvaluationTaskInstance) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *DataQualityEvaluationTaskInstance) GetId() *int64 {
	return s.Id
}

func (s *DataQualityEvaluationTaskInstance) GetStatus() *string {
	return s.Status
}

func (s *DataQualityEvaluationTaskInstance) GetTask() *DataQualityEvaluationTaskInstanceTask {
	return s.Task
}

func (s *DataQualityEvaluationTaskInstance) SetCreateTime(v int64) *DataQualityEvaluationTaskInstance {
	s.CreateTime = &v
	return s
}

func (s *DataQualityEvaluationTaskInstance) SetFinishTime(v int64) *DataQualityEvaluationTaskInstance {
	s.FinishTime = &v
	return s
}

func (s *DataQualityEvaluationTaskInstance) SetId(v int64) *DataQualityEvaluationTaskInstance {
	s.Id = &v
	return s
}

func (s *DataQualityEvaluationTaskInstance) SetStatus(v string) *DataQualityEvaluationTaskInstance {
	s.Status = &v
	return s
}

func (s *DataQualityEvaluationTaskInstance) SetTask(v *DataQualityEvaluationTaskInstanceTask) *DataQualityEvaluationTaskInstance {
	s.Task = v
	return s
}

func (s *DataQualityEvaluationTaskInstance) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityEvaluationTaskInstanceTask struct {
	// The ID of the data source that is used for task running.
	//
	// example:
	//
	// 201
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The callback configurations of the task during the instance lifecycle. Blocking an auto triggered node is a type of callback event. Only this type is supported.
	Hooks []*DataQualityEvaluationTaskInstanceTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitoring task.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data quality monitoring task. The name can be up to 255 characters in length and can contain digits, letters, and punctuation marks.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subscription configurations for alert notifications.
	Notifications []*DataQualityEvaluationTaskInstanceTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 2626
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The configuration of the data source. The value of the queue field is default, and that of the sqlEngine field can be SPARK_SQL, KYUUBI, PRESTO_SQL, or HIVE_SQL to collect EMR data. The value default indicates the YARN queue for E-MapReduce (EMR) tasks.
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the data quality monitoring task.
	Target *DataQualityEvaluationTaskInstanceTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The DataWorks tenant ID.
	//
	// example:
	//
	// 195820716552192
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The trigger configuration of the data quality monitoring task.
	Trigger *DataQualityEvaluationTaskInstanceTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s DataQualityEvaluationTaskInstanceTask) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTask) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTask) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *DataQualityEvaluationTaskInstanceTask) GetHooks() []*DataQualityEvaluationTaskInstanceTaskHooks {
	return s.Hooks
}

func (s *DataQualityEvaluationTaskInstanceTask) GetId() *int64 {
	return s.Id
}

func (s *DataQualityEvaluationTaskInstanceTask) GetName() *string {
	return s.Name
}

func (s *DataQualityEvaluationTaskInstanceTask) GetNotifications() []*DataQualityEvaluationTaskInstanceTaskNotifications {
	return s.Notifications
}

func (s *DataQualityEvaluationTaskInstanceTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DataQualityEvaluationTaskInstanceTask) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *DataQualityEvaluationTaskInstanceTask) GetTarget() *DataQualityEvaluationTaskInstanceTaskTarget {
	return s.Target
}

func (s *DataQualityEvaluationTaskInstanceTask) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DataQualityEvaluationTaskInstanceTask) GetTrigger() *DataQualityEvaluationTaskInstanceTaskTrigger {
	return s.Trigger
}

func (s *DataQualityEvaluationTaskInstanceTask) SetDataSourceId(v int64) *DataQualityEvaluationTaskInstanceTask {
	s.DataSourceId = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetHooks(v []*DataQualityEvaluationTaskInstanceTaskHooks) *DataQualityEvaluationTaskInstanceTask {
	s.Hooks = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetId(v int64) *DataQualityEvaluationTaskInstanceTask {
	s.Id = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetName(v string) *DataQualityEvaluationTaskInstanceTask {
	s.Name = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetNotifications(v []*DataQualityEvaluationTaskInstanceTaskNotifications) *DataQualityEvaluationTaskInstanceTask {
	s.Notifications = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetProjectId(v int64) *DataQualityEvaluationTaskInstanceTask {
	s.ProjectId = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetRuntimeConf(v string) *DataQualityEvaluationTaskInstanceTask {
	s.RuntimeConf = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetTarget(v *DataQualityEvaluationTaskInstanceTaskTarget) *DataQualityEvaluationTaskInstanceTask {
	s.Target = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetTenantId(v int64) *DataQualityEvaluationTaskInstanceTask {
	s.TenantId = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) SetTrigger(v *DataQualityEvaluationTaskInstanceTaskTrigger) *DataQualityEvaluationTaskInstanceTask {
	s.Trigger = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTask) Validate() error {
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
		for _, item := range s.Notifications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type DataQualityEvaluationTaskInstanceTaskHooks struct {
	// The trigger configuration of the callback event.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the callback event. Valid values:
	//
	// 	- BlockTaskInstance: An auto triggered node is blocked.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskInstanceTaskHooks) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskHooks) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) GetCondition() *string {
	return s.Condition
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) GetType() *string {
	return s.Type
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) SetCondition(v string) *DataQualityEvaluationTaskInstanceTaskHooks {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) SetType(v string) *DataQualityEvaluationTaskInstanceTaskHooks {
	s.Type = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskNotifications struct {
	// The trigger condition of the alert notification.
	//
	// example:
	//
	// ${blockType} == "Strong"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The configurations for the alert notification.
	Notifications []*DataQualityEvaluationTaskInstanceTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotifications) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) GetCondition() *string {
	return s.Condition
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) GetNotifications() []*DataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	return s.Notifications
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) SetCondition(v string) *DataQualityEvaluationTaskInstanceTaskNotifications {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) SetNotifications(v []*DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) *DataQualityEvaluationTaskInstanceTaskNotifications {
	s.Notifications = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) Validate() error {
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

type DataQualityEvaluationTaskInstanceTaskNotificationsNotifications struct {
	// The alert notification methods.
	NotificationChannels []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert recipient configurations.
	NotificationReceivers []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GetNotificationChannels() []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GetNotificationReceivers() []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) SetNotificationChannels(v []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) SetNotificationReceivers(v []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) Validate() error {
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

type DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels struct {
	// The alert notification methods.
	//
	// example:
	//
	// Mail
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers struct {
	// The extended information in the JSON format. For example, the DingTalk chatbot can remind all members in a DingTalk group by using the at sign (@).
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The alert recipient configuration.
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) SetExtension(v string) *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskTarget struct {
	// The type of the database to which the table belongs. Valid values:
	//
	// 	- maxcompute
	//
	// 	- emr
	//
	// 	- cdh
	//
	// 	- hologres
	//
	// 	- analyticdb_for_postgresql
	//
	// 	- analyticdb_for_mysql
	//
	// 	- starrocks
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The configuration of the partitioned table.
	//
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The ID of the table in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The type of the monitored object. Valid values:
	//
	// 	- Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskInstanceTaskTarget) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskTarget) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) GetType() *string {
	return s.Type
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) SetDatabaseType(v string) *DataQualityEvaluationTaskInstanceTaskTarget {
	s.DatabaseType = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) SetPartitionSpec(v string) *DataQualityEvaluationTaskInstanceTaskTarget {
	s.PartitionSpec = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) SetTableGuid(v string) *DataQualityEvaluationTaskInstanceTaskTarget {
	s.TableGuid = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) SetType(v string) *DataQualityEvaluationTaskInstanceTaskTarget {
	s.Type = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTarget) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskTrigger struct {
	// The IDs of the auto triggered nodes of which the instances are successfully run. This parameter takes effect only if the Type parameter is set to ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger condition of the task. Valid values:
	//
	// 	- ByScheduledTaskInstance: The task is triggered when the instance of an auto triggered node is successfully run.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskInstanceTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskTrigger) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) GetType() *string {
	return s.Type
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) SetTaskIds(v []*int64) *DataQualityEvaluationTaskInstanceTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) SetType(v string) *DataQualityEvaluationTaskInstanceTaskTrigger {
	s.Type = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) Validate() error {
	return dara.Validate(s)
}
