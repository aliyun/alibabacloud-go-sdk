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
	// The details of the monitor.
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
	// The ID of the data source used for the monitor.
	//
	// example:
	//
	// 45238
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the monitor.
	//
	// example:
	//
	// The description of the quality monitoring task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook.
	Hooks []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 2178
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenAPI create a data quality monitoring test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of alert notifications.
	Notifications *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// 2626
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Extended configuration, JSON-formatted string, takes effect only for EMR-type data quality monitoring.
	//
	// - queue: the yarn queue used when performing EMR data quality verification. The default queue is the queue configured for this project.
	//
	// - sqlEngine: SQL engine used when performing EMR data verification
	//
	//     - HIVE_ SQL
	//
	//     - SPARK_ SQL
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the monitor.
	Target *GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the monitor.
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
	// The hook trigger condition. When this condition is met, the hook action is triggered. Only two conditional expressions are supported:
	//
	// 	- Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 	- Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The hook type. Only one hook type is supported.
	//
	// 	- BlockTaskInstance: Blocks the running of scheduling tasks. A monitor is triggered by scheduling tasks. After a monitor finishes running, the monitor determines whether to block the running of scheduling tasks based on the hook condition.
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
	// The notification trigger condition. When this condition is met, the alert notification is triggered. Only two conditional expressions are supported:
	//
	// 	- Specify only one group of rule strength type and rule check status, such as `${severity} == "High" AND ${status} == "Critical"`. In this expression, the hook trigger condition is met if severity is High and status is Critical.
	//
	// 	- Specify multiple groups of rule strength types and rule check status, such as `(${severity} == "High"AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")`. In this expression, the hook trigger condition is met if severity is High and status is Critical, severity is Normal and status is Critical, or severity is Normal and status is Error. The enumeration of severity in a conditional expression is the same as the enumeration of severity in DataQualityRule. The enumeration of status in a conditional expression is the same as the enumeration of status in DataQualityResult.
	//
	// example:
	//
	// (${severity} == "High" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Critical") OR (${severity} == "Normal" AND ${status} == "Error")
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The configurations of alert notifications.
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
	// The alert notification methods.
	NotificationChannels []*GetDataQualityEvaluationTaskResponseBodyDataQualityEvaluationTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The configurations of alert recipients.
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
	// The alert notification methods.
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
	// The extended information.
	//
	// example:
	//
	// {  "atAll": true }
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The additional parameters that are required when alerts are sent. The parameters are JSON-formatted strings. The following keys are supported:
	//
	// 	- atAll: specifies that all members in a group are mentioned when alerts are sent by using DingTalk. This parameter is valid only if you set ReceiverType to DingdingUrl.
	//
	// Valid values:
	//
	// 	- WebhookUrl
	//
	// 	- FeishuUrl
	//
	// 	- DingdingUrl
	//
	// 	- WeixinUrl
	//
	// 	- AliUid
	//
	// example:
	//
	// DingdingUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
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
	// The type of the database to which the table belongs. Valid values:
	//
	// 	- maxcompute
	//
	// 	- hologres
	//
	// 	- cdh
	//
	// 	- analyticdb_for_mysql
	//
	// 	- starrocks
	//
	// 	- emr
	//
	// 	- analyticdb_for_postgresql
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// Data quality monitoring partition range settings.
	//
	// example:
	//
	// pt=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The ID of the table in Data Map.
	//
	// example:
	//
	// odps.meta_open_api_test_sz.test_partition_tbl
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The type of the monitoring object.
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
	// The IDs of scheduling tasks. This parameter is valid only if you set Type to ByScheduledTaskInstance.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger type of the monitor. Valid values:
	//
	// 	- ByManual: The monitor is manually triggered.
	//
	// 	- ByScheduledTaskInstance: The monitor is triggered by associated scheduling tasks.
	//
	// 	- ByQualityNode: The monitor is triggered by created data quality monitoring nodes.
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
