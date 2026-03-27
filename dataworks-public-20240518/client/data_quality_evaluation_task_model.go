// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataQualityEvaluationTask interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int64) *DataQualityEvaluationTask
	GetDataSourceId() *int64
	SetDescription(v string) *DataQualityEvaluationTask
	GetDescription() *string
	SetHooks(v []*DataQualityEvaluationTaskHooks) *DataQualityEvaluationTask
	GetHooks() []*DataQualityEvaluationTaskHooks
	SetId(v int64) *DataQualityEvaluationTask
	GetId() *int64
	SetName(v string) *DataQualityEvaluationTask
	GetName() *string
	SetNotifications(v []*DataQualityEvaluationTaskNotifications) *DataQualityEvaluationTask
	GetNotifications() []*DataQualityEvaluationTaskNotifications
	SetProjectId(v int64) *DataQualityEvaluationTask
	GetProjectId() *int64
	SetRuntimeConf(v string) *DataQualityEvaluationTask
	GetRuntimeConf() *string
	SetTarget(v *DataQualityEvaluationTaskTarget) *DataQualityEvaluationTask
	GetTarget() *DataQualityEvaluationTaskTarget
	SetTenantId(v int64) *DataQualityEvaluationTask
	GetTenantId() *int64
	SetTrigger(v *DataQualityEvaluationTaskTrigger) *DataQualityEvaluationTask
	GetTrigger() *DataQualityEvaluationTaskTrigger
}

type DataQualityEvaluationTask struct {
	// The ID of the data source that is used for task running.
	//
	// example:
	//
	// 201
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the task. The description can be up to 65,535 characters in length.
	//
	// example:
	//
	// This is a daily run data quality evaluation plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The callback configurations of the task during the instance lifecycle. Blocking an auto triggered node is a type of callback event. Only this type is supported.
	Hooks []*DataQualityEvaluationTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitoring task.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data quality monitoring task. The name can be up to 255 characters in length and can contain digits, letters, and punctuation marks.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations for alert notifications.
	Notifications []*DataQualityEvaluationTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The configuration of the data source. The value of the queue field is default, and that of the sqlEngine field can be set to SPARK_SQL, KYUUBI, PRESTO_SQL, or HIVE_SQL. The value default indicates the YARN queue for E-MapReduce (EMR) tasks.
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the task.
	Target *DataQualityEvaluationTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The DataWorks tenant ID.
	//
	// example:
	//
	// 10
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The trigger configuration of the task.
	Trigger *DataQualityEvaluationTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s DataQualityEvaluationTask) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTask) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTask) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *DataQualityEvaluationTask) GetDescription() *string {
	return s.Description
}

func (s *DataQualityEvaluationTask) GetHooks() []*DataQualityEvaluationTaskHooks {
	return s.Hooks
}

func (s *DataQualityEvaluationTask) GetId() *int64 {
	return s.Id
}

func (s *DataQualityEvaluationTask) GetName() *string {
	return s.Name
}

func (s *DataQualityEvaluationTask) GetNotifications() []*DataQualityEvaluationTaskNotifications {
	return s.Notifications
}

func (s *DataQualityEvaluationTask) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DataQualityEvaluationTask) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *DataQualityEvaluationTask) GetTarget() *DataQualityEvaluationTaskTarget {
	return s.Target
}

func (s *DataQualityEvaluationTask) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DataQualityEvaluationTask) GetTrigger() *DataQualityEvaluationTaskTrigger {
	return s.Trigger
}

func (s *DataQualityEvaluationTask) SetDataSourceId(v int64) *DataQualityEvaluationTask {
	s.DataSourceId = &v
	return s
}

func (s *DataQualityEvaluationTask) SetDescription(v string) *DataQualityEvaluationTask {
	s.Description = &v
	return s
}

func (s *DataQualityEvaluationTask) SetHooks(v []*DataQualityEvaluationTaskHooks) *DataQualityEvaluationTask {
	s.Hooks = v
	return s
}

func (s *DataQualityEvaluationTask) SetId(v int64) *DataQualityEvaluationTask {
	s.Id = &v
	return s
}

func (s *DataQualityEvaluationTask) SetName(v string) *DataQualityEvaluationTask {
	s.Name = &v
	return s
}

func (s *DataQualityEvaluationTask) SetNotifications(v []*DataQualityEvaluationTaskNotifications) *DataQualityEvaluationTask {
	s.Notifications = v
	return s
}

func (s *DataQualityEvaluationTask) SetProjectId(v int64) *DataQualityEvaluationTask {
	s.ProjectId = &v
	return s
}

func (s *DataQualityEvaluationTask) SetRuntimeConf(v string) *DataQualityEvaluationTask {
	s.RuntimeConf = &v
	return s
}

func (s *DataQualityEvaluationTask) SetTarget(v *DataQualityEvaluationTaskTarget) *DataQualityEvaluationTask {
	s.Target = v
	return s
}

func (s *DataQualityEvaluationTask) SetTenantId(v int64) *DataQualityEvaluationTask {
	s.TenantId = &v
	return s
}

func (s *DataQualityEvaluationTask) SetTrigger(v *DataQualityEvaluationTaskTrigger) *DataQualityEvaluationTask {
	s.Trigger = v
	return s
}

func (s *DataQualityEvaluationTask) Validate() error {
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

type DataQualityEvaluationTaskHooks struct {
	// The trigger configuration of the callback event.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the callback event. Valid values:
	//
	// 	- BlockTaskInstance. The value indicates that an auto triggered node is blocked.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskHooks) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskHooks) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskHooks) GetCondition() *string {
	return s.Condition
}

func (s *DataQualityEvaluationTaskHooks) GetType() *string {
	return s.Type
}

func (s *DataQualityEvaluationTaskHooks) SetCondition(v string) *DataQualityEvaluationTaskHooks {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskHooks) SetType(v string) *DataQualityEvaluationTaskHooks {
	s.Type = &v
	return s
}

func (s *DataQualityEvaluationTaskHooks) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskNotifications struct {
	// The trigger condition of the alert notification.
	//
	// example:
	//
	// ${blockType} == "Strong"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The configurations for alert notifications.
	Notifications []*DataQualityEvaluationTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotifications) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotifications) GetCondition() *string {
	return s.Condition
}

func (s *DataQualityEvaluationTaskNotifications) GetNotifications() []*DataQualityEvaluationTaskNotificationsNotifications {
	return s.Notifications
}

func (s *DataQualityEvaluationTaskNotifications) SetCondition(v string) *DataQualityEvaluationTaskNotifications {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskNotifications) SetNotifications(v []*DataQualityEvaluationTaskNotificationsNotifications) *DataQualityEvaluationTaskNotifications {
	s.Notifications = v
	return s
}

func (s *DataQualityEvaluationTaskNotifications) Validate() error {
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

type DataQualityEvaluationTaskNotificationsNotifications struct {
	// The alert notification methods.
	NotificationChannels []*DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert recipients.
	NotificationReceivers []*DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) GetNotificationChannels() []*DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) GetNotificationReceivers() []*DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) SetNotificationChannels(v []*DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) *DataQualityEvaluationTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) SetNotificationReceivers(v []*DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) *DataQualityEvaluationTaskNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) Validate() error {
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

type DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels struct {
	// The alert notification method.
	//
	// example:
	//
	// Mail
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers struct {
	// The extended information in the JSON format. For example, the DingTalk chatbot can remind all members in a DingTalk group by using the at sign (@).
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient.
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) SetExtension(v string) *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskTarget struct {
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

func (s DataQualityEvaluationTaskTarget) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskTarget) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DataQualityEvaluationTaskTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *DataQualityEvaluationTaskTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *DataQualityEvaluationTaskTarget) GetType() *string {
	return s.Type
}

func (s *DataQualityEvaluationTaskTarget) SetDatabaseType(v string) *DataQualityEvaluationTaskTarget {
	s.DatabaseType = &v
	return s
}

func (s *DataQualityEvaluationTaskTarget) SetPartitionSpec(v string) *DataQualityEvaluationTaskTarget {
	s.PartitionSpec = &v
	return s
}

func (s *DataQualityEvaluationTaskTarget) SetTableGuid(v string) *DataQualityEvaluationTaskTarget {
	s.TableGuid = &v
	return s
}

func (s *DataQualityEvaluationTaskTarget) SetType(v string) *DataQualityEvaluationTaskTarget {
	s.Type = &v
	return s
}

func (s *DataQualityEvaluationTaskTarget) Validate() error {
	return dara.Validate(s)
}

type DataQualityEvaluationTaskTrigger struct {
	// The IDs of the auto triggered nodes of which the instances are successfully run. This parameter takes effect only if the Type parameter is set to ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger condition of the task. Valid values:
	//
	// 	- ByScheduledTaskInstance. The value indicates that the task is triggered when the instance of an auto triggered node is successfully run.
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskTrigger) String() string {
	return dara.Prettify(s)
}

func (s DataQualityEvaluationTaskTrigger) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *DataQualityEvaluationTaskTrigger) GetType() *string {
	return s.Type
}

func (s *DataQualityEvaluationTaskTrigger) SetTaskIds(v []*int64) *DataQualityEvaluationTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *DataQualityEvaluationTaskTrigger) SetType(v string) *DataQualityEvaluationTaskTrigger {
	s.Type = &v
	return s
}

func (s *DataQualityEvaluationTaskTrigger) Validate() error {
	return dara.Validate(s)
}
