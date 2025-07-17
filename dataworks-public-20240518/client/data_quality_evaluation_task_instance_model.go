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
	// example:
	//
	// 1710239005403
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1710239005403
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Passed
	Status *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Task   *DataQualityEvaluationTaskInstanceTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTask struct {
	// example:
	//
	// 201
	DataSourceId *int64                                        `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	Hooks        []*DataQualityEvaluationTaskInstanceTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 质量校验任务
	Name          *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Notifications []*DataQualityEvaluationTaskInstanceTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
	ProjectId     *int64                                                `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string                                       `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	Target      *DataQualityEvaluationTaskInstanceTaskTarget  `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	TenantId    *int64                                        `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Trigger     *DataQualityEvaluationTaskInstanceTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskHooks struct {
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
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
	// example:
	//
	// ${blockType} == "Strong"
	Condition     *string                                                            `json:"Condition,omitempty" xml:"Condition,omitempty"`
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
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskNotificationsNotifications struct {
	NotificationChannels  []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels struct {
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
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// AliUid
	ReceiverType   *string   `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
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
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
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
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
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
