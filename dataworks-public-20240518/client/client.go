// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataQualityEvaluationTask struct {
	// example:
	//
	// 201
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// This is a daily run data quality evaluation plan.
	Description *string                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks       []*DataQualityEvaluationTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 质量校验任务
	Name          *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Notifications []*DataQualityEvaluationTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string                          `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	Target      *DataQualityEvaluationTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// 10
	TenantId *int64                            `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Trigger  *DataQualityEvaluationTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s DataQualityEvaluationTask) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTask) GoString() string {
	return s.String()
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

type DataQualityEvaluationTaskHooks struct {
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskHooks) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskHooks) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskHooks) SetCondition(v string) *DataQualityEvaluationTaskHooks {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskHooks) SetType(v string) *DataQualityEvaluationTaskHooks {
	s.Type = &v
	return s
}

type DataQualityEvaluationTaskNotifications struct {
	// example:
	//
	// ${blockType} == "Strong"
	Condition     *string                                                `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Notifications []*DataQualityEvaluationTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotifications) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotifications) SetCondition(v string) *DataQualityEvaluationTaskNotifications {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskNotifications) SetNotifications(v []*DataQualityEvaluationTaskNotificationsNotifications) *DataQualityEvaluationTaskNotifications {
	s.Notifications = v
	return s
}

type DataQualityEvaluationTaskNotificationsNotifications struct {
	NotificationChannels  []*DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	NotificationReceivers []*DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotificationsNotifications) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) SetNotificationChannels(v []*DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) *DataQualityEvaluationTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *DataQualityEvaluationTaskNotificationsNotifications) SetNotificationReceivers(v []*DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) *DataQualityEvaluationTaskNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

type DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels struct {
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *DataQualityEvaluationTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

type DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers struct {
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

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
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

type DataQualityEvaluationTaskTarget struct {
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

func (s DataQualityEvaluationTaskTarget) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskTarget) GoString() string {
	return s.String()
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

type DataQualityEvaluationTaskTrigger struct {
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskTrigger) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskTrigger) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskTrigger) SetTaskIds(v []*int64) *DataQualityEvaluationTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *DataQualityEvaluationTaskTrigger) SetType(v string) *DataQualityEvaluationTaskTrigger {
	s.Type = &v
	return s
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
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstance) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTask) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskHooks) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) SetCondition(v string) *DataQualityEvaluationTaskInstanceTaskHooks {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskHooks) SetType(v string) *DataQualityEvaluationTaskInstanceTaskHooks {
	s.Type = &v
	return s
}

type DataQualityEvaluationTaskInstanceTaskNotifications struct {
	// example:
	//
	// ${blockType} == "Strong"
	Condition     *string                                                            `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Notifications []*DataQualityEvaluationTaskInstanceTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotifications) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) SetCondition(v string) *DataQualityEvaluationTaskInstanceTaskNotifications {
	s.Condition = &v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotifications) SetNotifications(v []*DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) *DataQualityEvaluationTaskInstanceTaskNotifications {
	s.Notifications = v
	return s
}

type DataQualityEvaluationTaskInstanceTaskNotificationsNotifications struct {
	NotificationChannels  []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	NotificationReceivers []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) SetNotificationChannels(v []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications) SetNotificationReceivers(v []*DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) *DataQualityEvaluationTaskInstanceTaskNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

type DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels struct {
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
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
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskTarget) GoString() string {
	return s.String()
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

type DataQualityEvaluationTaskInstanceTaskTrigger struct {
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityEvaluationTaskInstanceTaskTrigger) String() string {
	return tea.Prettify(s)
}

func (s DataQualityEvaluationTaskInstanceTaskTrigger) GoString() string {
	return s.String()
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) SetTaskIds(v []*int64) *DataQualityEvaluationTaskInstanceTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *DataQualityEvaluationTaskInstanceTaskTrigger) SetType(v string) *DataQualityEvaluationTaskInstanceTaskTrigger {
	s.Type = &v
	return s
}

type DataQualityResult struct {
	Details []*DataQualityResultDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id   *int64                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Rule *DataQualityResultRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// example:
	//
	// [   {     "gender": "male",     "_count": 100   }, {     "gender": "female",     "_count": 100   } ]
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20001
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s DataQualityResult) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResult) GoString() string {
	return s.String()
}

func (s *DataQualityResult) SetDetails(v []*DataQualityResultDetails) *DataQualityResult {
	s.Details = v
	return s
}

func (s *DataQualityResult) SetId(v int64) *DataQualityResult {
	s.Id = &v
	return s
}

func (s *DataQualityResult) SetRule(v *DataQualityResultRule) *DataQualityResult {
	s.Rule = v
	return s
}

func (s *DataQualityResult) SetSample(v string) *DataQualityResult {
	s.Sample = &v
	return s
}

func (s *DataQualityResult) SetStatus(v string) *DataQualityResult {
	s.Status = &v
	return s
}

func (s *DataQualityResult) SetTaskInstanceId(v int64) *DataQualityResult {
	s.TaskInstanceId = &v
	return s
}

type DataQualityResultDetails struct {
	// example:
	//
	// 100.0
	CheckedValue *string `json:"CheckedValue,omitempty" xml:"CheckedValue,omitempty"`
	// example:
	//
	// 0.0
	ReferencedValue *string `json:"ReferencedValue,omitempty" xml:"ReferencedValue,omitempty"`
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DataQualityResultDetails) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultDetails) GoString() string {
	return s.String()
}

func (s *DataQualityResultDetails) SetCheckedValue(v string) *DataQualityResultDetails {
	s.CheckedValue = &v
	return s
}

func (s *DataQualityResultDetails) SetReferencedValue(v string) *DataQualityResultDetails {
	s.ReferencedValue = &v
	return s
}

func (s *DataQualityResultDetails) SetStatus(v string) *DataQualityResultDetails {
	s.Status = &v
	return s
}

type DataQualityResultRule struct {
	CheckingConfig *DataQualityResultRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled       *bool                                 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers []*DataQualityResultRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 表不能为空
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	ProjectId      *int64                               `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *DataQualityResultRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// example:
	//
	// High
	Severity *string                      `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target   *DataQualityResultRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 1
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DataQualityResultRule) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRule) GoString() string {
	return s.String()
}

func (s *DataQualityResultRule) SetCheckingConfig(v *DataQualityResultRuleCheckingConfig) *DataQualityResultRule {
	s.CheckingConfig = v
	return s
}

func (s *DataQualityResultRule) SetDescription(v string) *DataQualityResultRule {
	s.Description = &v
	return s
}

func (s *DataQualityResultRule) SetEnabled(v bool) *DataQualityResultRule {
	s.Enabled = &v
	return s
}

func (s *DataQualityResultRule) SetErrorHandlers(v []*DataQualityResultRuleErrorHandlers) *DataQualityResultRule {
	s.ErrorHandlers = v
	return s
}

func (s *DataQualityResultRule) SetId(v int64) *DataQualityResultRule {
	s.Id = &v
	return s
}

func (s *DataQualityResultRule) SetName(v string) *DataQualityResultRule {
	s.Name = &v
	return s
}

func (s *DataQualityResultRule) SetProjectId(v int64) *DataQualityResultRule {
	s.ProjectId = &v
	return s
}

func (s *DataQualityResultRule) SetSamplingConfig(v *DataQualityResultRuleSamplingConfig) *DataQualityResultRule {
	s.SamplingConfig = v
	return s
}

func (s *DataQualityResultRule) SetSeverity(v string) *DataQualityResultRule {
	s.Severity = &v
	return s
}

func (s *DataQualityResultRule) SetTarget(v *DataQualityResultRuleTarget) *DataQualityResultRule {
	s.Target = v
	return s
}

func (s *DataQualityResultRule) SetTemplateCode(v string) *DataQualityResultRule {
	s.TemplateCode = &v
	return s
}

func (s *DataQualityResultRule) SetTenantId(v int64) *DataQualityResultRule {
	s.TenantId = &v
	return s
}

type DataQualityResultRuleCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string                                        `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *DataQualityResultRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityResultRuleCheckingConfig) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfig) SetReferencedSamplesFilter(v string) *DataQualityResultRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfig) SetThresholds(v *DataQualityResultRuleCheckingConfigThresholds) *DataQualityResultRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *DataQualityResultRuleCheckingConfig) SetType(v string) *DataQualityResultRuleCheckingConfig {
	s.Type = &v
	return s
}

type DataQualityResultRuleCheckingConfigThresholds struct {
	Critical *DataQualityResultRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *DataQualityResultRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *DataQualityResultRuleCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s DataQualityResultRuleCheckingConfigThresholds) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholds) SetCritical(v *DataQualityResultRuleCheckingConfigThresholdsCritical) *DataQualityResultRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholds) SetExpected(v *DataQualityResultRuleCheckingConfigThresholdsExpected) *DataQualityResultRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholds) SetWarned(v *DataQualityResultRuleCheckingConfigThresholdsWarned) *DataQualityResultRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

type DataQualityResultRuleCheckingConfigThresholdsCritical struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityResultRuleCheckingConfigThresholdsCritical) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) SetOperator(v string) *DataQualityResultRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) SetValue(v string) *DataQualityResultRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

type DataQualityResultRuleCheckingConfigThresholdsExpected struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityResultRuleCheckingConfigThresholdsExpected) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) SetOperator(v string) *DataQualityResultRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) SetValue(v string) *DataQualityResultRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

type DataQualityResultRuleCheckingConfigThresholdsWarned struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityResultRuleCheckingConfigThresholdsWarned) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) SetOperator(v string) *DataQualityResultRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) SetValue(v string) *DataQualityResultRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

type DataQualityResultRuleErrorHandlers struct {
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityResultRuleErrorHandlers) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleErrorHandlers) SetErrorDataFilter(v string) *DataQualityResultRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *DataQualityResultRuleErrorHandlers) SetType(v string) *DataQualityResultRuleErrorHandlers {
	s.Type = &v
	return s
}

type DataQualityResultRuleSamplingConfig struct {
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// { "Columns": [ "id", "name" ] }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s DataQualityResultRuleSamplingConfig) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleSamplingConfig) SetMetric(v string) *DataQualityResultRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) SetMetricParameters(v string) *DataQualityResultRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) SetSamplingFilter(v string) *DataQualityResultRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) SetSettingConfig(v string) *DataQualityResultRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

type DataQualityResultRuleTarget struct {
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

func (s DataQualityResultRuleTarget) String() string {
	return tea.Prettify(s)
}

func (s DataQualityResultRuleTarget) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleTarget) SetDatabaseType(v string) *DataQualityResultRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *DataQualityResultRuleTarget) SetPartitionSpec(v string) *DataQualityResultRuleTarget {
	s.PartitionSpec = &v
	return s
}

func (s *DataQualityResultRuleTarget) SetTableGuid(v string) *DataQualityResultRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *DataQualityResultRuleTarget) SetType(v string) *DataQualityResultRuleTarget {
	s.Type = &v
	return s
}

type DataQualityRule struct {
	CheckingConfig *DataQualityRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled       *bool                           `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers []*DataQualityRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 表不能为空
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	ProjectId      *int64                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *DataQualityRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// example:
	//
	// High
	Severity *string                `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target   *DataQualityRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 1
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DataQualityRule) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRule) GoString() string {
	return s.String()
}

func (s *DataQualityRule) SetCheckingConfig(v *DataQualityRuleCheckingConfig) *DataQualityRule {
	s.CheckingConfig = v
	return s
}

func (s *DataQualityRule) SetDescription(v string) *DataQualityRule {
	s.Description = &v
	return s
}

func (s *DataQualityRule) SetEnabled(v bool) *DataQualityRule {
	s.Enabled = &v
	return s
}

func (s *DataQualityRule) SetErrorHandlers(v []*DataQualityRuleErrorHandlers) *DataQualityRule {
	s.ErrorHandlers = v
	return s
}

func (s *DataQualityRule) SetId(v int64) *DataQualityRule {
	s.Id = &v
	return s
}

func (s *DataQualityRule) SetName(v string) *DataQualityRule {
	s.Name = &v
	return s
}

func (s *DataQualityRule) SetProjectId(v int64) *DataQualityRule {
	s.ProjectId = &v
	return s
}

func (s *DataQualityRule) SetSamplingConfig(v *DataQualityRuleSamplingConfig) *DataQualityRule {
	s.SamplingConfig = v
	return s
}

func (s *DataQualityRule) SetSeverity(v string) *DataQualityRule {
	s.Severity = &v
	return s
}

func (s *DataQualityRule) SetTarget(v *DataQualityRuleTarget) *DataQualityRule {
	s.Target = v
	return s
}

func (s *DataQualityRule) SetTemplateCode(v string) *DataQualityRule {
	s.TemplateCode = &v
	return s
}

func (s *DataQualityRule) SetTenantId(v int64) *DataQualityRule {
	s.TenantId = &v
	return s
}

type DataQualityRuleCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string                                  `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *DataQualityRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityRuleCheckingConfig) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfig) SetReferencedSamplesFilter(v string) *DataQualityRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *DataQualityRuleCheckingConfig) SetThresholds(v *DataQualityRuleCheckingConfigThresholds) *DataQualityRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *DataQualityRuleCheckingConfig) SetType(v string) *DataQualityRuleCheckingConfig {
	s.Type = &v
	return s
}

type DataQualityRuleCheckingConfigThresholds struct {
	Critical *DataQualityRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *DataQualityRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *DataQualityRuleCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s DataQualityRuleCheckingConfigThresholds) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholds) SetCritical(v *DataQualityRuleCheckingConfigThresholdsCritical) *DataQualityRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholds) SetExpected(v *DataQualityRuleCheckingConfigThresholdsExpected) *DataQualityRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholds) SetWarned(v *DataQualityRuleCheckingConfigThresholdsWarned) *DataQualityRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

type DataQualityRuleCheckingConfigThresholdsCritical struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityRuleCheckingConfigThresholdsCritical) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) SetOperator(v string) *DataQualityRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) SetValue(v string) *DataQualityRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

type DataQualityRuleCheckingConfigThresholdsExpected struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityRuleCheckingConfigThresholdsExpected) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) SetOperator(v string) *DataQualityRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) SetValue(v string) *DataQualityRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

type DataQualityRuleCheckingConfigThresholdsWarned struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityRuleCheckingConfigThresholdsWarned) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) SetOperator(v string) *DataQualityRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) SetValue(v string) *DataQualityRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

type DataQualityRuleErrorHandlers struct {
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityRuleErrorHandlers) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *DataQualityRuleErrorHandlers) SetErrorDataFilter(v string) *DataQualityRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *DataQualityRuleErrorHandlers) SetType(v string) *DataQualityRuleErrorHandlers {
	s.Type = &v
	return s
}

type DataQualityRuleSamplingConfig struct {
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// { "Columns": [ "id", "name" ] }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s DataQualityRuleSamplingConfig) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityRuleSamplingConfig) SetMetric(v string) *DataQualityRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) SetMetricParameters(v string) *DataQualityRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) SetSamplingFilter(v string) *DataQualityRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) SetSettingConfig(v string) *DataQualityRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

type DataQualityRuleTarget struct {
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

func (s DataQualityRuleTarget) String() string {
	return tea.Prettify(s)
}

func (s DataQualityRuleTarget) GoString() string {
	return s.String()
}

func (s *DataQualityRuleTarget) SetDatabaseType(v string) *DataQualityRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *DataQualityRuleTarget) SetPartitionSpec(v string) *DataQualityRuleTarget {
	s.PartitionSpec = &v
	return s
}

func (s *DataQualityRuleTarget) SetTableGuid(v string) *DataQualityRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *DataQualityRuleTarget) SetType(v string) *DataQualityRuleTarget {
	s.Type = &v
	return s
}

type AbolishDeploymentRequest struct {
	// The ID of the process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1606087c-9ac4-43f0-83a8-0b5ced21XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AbolishDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentRequest) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentRequest) SetId(v string) *AbolishDeploymentRequest {
	s.Id = &v
	return s
}

func (s *AbolishDeploymentRequest) SetProjectId(v string) *AbolishDeploymentRequest {
	s.ProjectId = &v
	return s
}

type AbolishDeploymentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 55D786C9-DD57-524D-884C-C5239278XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbolishDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentResponseBody) SetRequestId(v string) *AbolishDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishDeploymentResponseBody) SetSuccess(v bool) *AbolishDeploymentResponseBody {
	s.Success = &v
	return s
}

type AbolishDeploymentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentResponse) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentResponse) SetHeaders(v map[string]*string) *AbolishDeploymentResponse {
	s.Headers = v
	return s
}

func (s *AbolishDeploymentResponse) SetStatusCode(v int32) *AbolishDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishDeploymentResponse) SetBody(v *AbolishDeploymentResponseBody) *AbolishDeploymentResponse {
	s.Body = v
	return s
}

type AssociateProjectToResourceGroupRequest struct {
	// The ID of the DataWorks workspace with which you want to associate the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AssociateProjectToResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateProjectToResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *AssociateProjectToResourceGroupRequest) SetProjectId(v int64) *AssociateProjectToResourceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateProjectToResourceGroupRequest) SetResourceGroupId(v string) *AssociateProjectToResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type AssociateProjectToResourceGroupResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssociateProjectToResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateProjectToResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateProjectToResourceGroupResponseBody) SetRequestId(v string) *AssociateProjectToResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateProjectToResourceGroupResponseBody) SetSuccess(v bool) *AssociateProjectToResourceGroupResponseBody {
	s.Success = &v
	return s
}

type AssociateProjectToResourceGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateProjectToResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateProjectToResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateProjectToResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *AssociateProjectToResourceGroupResponse) SetHeaders(v map[string]*string) *AssociateProjectToResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *AssociateProjectToResourceGroupResponse) SetStatusCode(v int32) *AssociateProjectToResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateProjectToResourceGroupResponse) SetBody(v *AssociateProjectToResourceGroupResponseBody) *AssociateProjectToResourceGroupResponse {
	s.Body = v
	return s
}

type CloneDataSourceRequest struct {
	// example:
	//
	// demo_holo_datasource
	CloneDataSourceName *string `json:"CloneDataSourceName,omitempty" xml:"CloneDataSourceName,omitempty"`
	// example:
	//
	// 16036
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CloneDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CloneDataSourceRequest) SetCloneDataSourceName(v string) *CloneDataSourceRequest {
	s.CloneDataSourceName = &v
	return s
}

func (s *CloneDataSourceRequest) SetId(v int64) *CloneDataSourceRequest {
	s.Id = &v
	return s
}

type CloneDataSourceResponseBody struct {
	// example:
	//
	// 19715
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// FCD583B9-346B-5E75-82C1-4A7C192C48DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CloneDataSourceResponseBody) SetId(v int64) *CloneDataSourceResponseBody {
	s.Id = &v
	return s
}

func (s *CloneDataSourceResponseBody) SetRequestId(v string) *CloneDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CloneDataSourceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CloneDataSourceResponse) SetHeaders(v map[string]*string) *CloneDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CloneDataSourceResponse) SetStatusCode(v int32) *CloneDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneDataSourceResponse) SetBody(v *CloneDataSourceResponseBody) *CloneDataSourceResponse {
	s.Body = v
	return s
}

type CreateDIAlarmRuleRequest struct {
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 任务ID，是告警规则关联的任务ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// 描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 告警规则是否启用，默认不开启。
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 告警指标类型，可选的枚举值：
	//
	// - Heartbeat（任务状态报警）
	//
	// - FailoverCount（failover次数报警）
	//
	// - Delay（任务延迟报警）
	//
	// This parameter is required.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alartRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NotificationSettings *CreateDIAlarmRuleRequestNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// This parameter is required.
	TriggerConditions []*CreateDIAlarmRuleRequestTriggerConditions `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s CreateDIAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequest) SetClientToken(v string) *CreateDIAlarmRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetDIJobId(v int64) *CreateDIAlarmRuleRequest {
	s.DIJobId = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetDescription(v string) *CreateDIAlarmRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetEnabled(v bool) *CreateDIAlarmRuleRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetMetricType(v string) *CreateDIAlarmRuleRequest {
	s.MetricType = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetName(v string) *CreateDIAlarmRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetNotificationSettings(v *CreateDIAlarmRuleRequestNotificationSettings) *CreateDIAlarmRuleRequest {
	s.NotificationSettings = v
	return s
}

func (s *CreateDIAlarmRuleRequest) SetTriggerConditions(v []*CreateDIAlarmRuleRequestTriggerConditions) *CreateDIAlarmRuleRequest {
	s.TriggerConditions = v
	return s
}

type CreateDIAlarmRuleRequestNotificationSettings struct {
	// example:
	//
	// 5
	InhibitionInterval    *int32                                                               `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	NotificationChannels  []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	NotificationReceivers []*CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s CreateDIAlarmRuleRequestNotificationSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleRequestNotificationSettings) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetInhibitionInterval(v int32) *CreateDIAlarmRuleRequestNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetNotificationChannels(v []*CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) *CreateDIAlarmRuleRequestNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettings) SetNotificationReceivers(v []*CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) *CreateDIAlarmRuleRequestNotificationSettings {
	s.NotificationReceivers = v
	return s
}

type CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels struct {
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetChannels(v []*string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetSeverity(v string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

type CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers struct {
	// example:
	//
	// DingToken
	ReceiverType   *string   `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverType(v string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *CreateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

type CreateDIAlarmRuleRequestTriggerConditions struct {
	DdlReportTags []*string `json:"DdlReportTags,omitempty" xml:"DdlReportTags,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 10
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreateDIAlarmRuleRequestTriggerConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleRequestTriggerConditions) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetDdlReportTags(v []*string) *CreateDIAlarmRuleRequestTriggerConditions {
	s.DdlReportTags = v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetDuration(v int64) *CreateDIAlarmRuleRequestTriggerConditions {
	s.Duration = &v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetSeverity(v string) *CreateDIAlarmRuleRequestTriggerConditions {
	s.Severity = &v
	return s
}

func (s *CreateDIAlarmRuleRequestTriggerConditions) SetThreshold(v int64) *CreateDIAlarmRuleRequestTriggerConditions {
	s.Threshold = &v
	return s
}

type CreateDIAlarmRuleShrinkRequest struct {
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 任务ID，是告警规则关联的任务ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// 描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 告警规则是否启用，默认不开启。
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 告警指标类型，可选的枚举值：
	//
	// - Heartbeat（任务状态报警）
	//
	// - FailoverCount（failover次数报警）
	//
	// - Delay（任务延迟报警）
	//
	// This parameter is required.
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alartRule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NotificationSettingsShrink *string `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty"`
	// This parameter is required.
	TriggerConditionsShrink *string `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty"`
}

func (s CreateDIAlarmRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleShrinkRequest) SetClientToken(v string) *CreateDIAlarmRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetDIJobId(v int64) *CreateDIAlarmRuleShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetDescription(v string) *CreateDIAlarmRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetEnabled(v bool) *CreateDIAlarmRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetMetricType(v string) *CreateDIAlarmRuleShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetName(v string) *CreateDIAlarmRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetNotificationSettingsShrink(v string) *CreateDIAlarmRuleShrinkRequest {
	s.NotificationSettingsShrink = &v
	return s
}

func (s *CreateDIAlarmRuleShrinkRequest) SetTriggerConditionsShrink(v string) *CreateDIAlarmRuleShrinkRequest {
	s.TriggerConditionsShrink = &v
	return s
}

type CreateDIAlarmRuleResponseBody struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 1
	DIAlarmRuleId *string `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// example:
	//
	// C636A747-7E4E-594D-94CD-2B4F8A9A9A63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDIAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleResponseBody) SetDIAlarmRuleId(v string) *CreateDIAlarmRuleResponseBody {
	s.DIAlarmRuleId = &v
	return s
}

func (s *CreateDIAlarmRuleResponseBody) SetRequestId(v string) *CreateDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDIAlarmRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDIAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDIAlarmRuleResponse) SetHeaders(v map[string]*string) *CreateDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDIAlarmRuleResponse) SetStatusCode(v int32) *CreateDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDIAlarmRuleResponse) SetBody(v *CreateDIAlarmRuleResponseBody) *CreateDIAlarmRuleResponse {
	s.Body = v
	return s
}

type CreateDIJobRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DestinationDataSourceSettings []*CreateDIJobRequestDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName     *string                        `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettings *CreateDIJobRequestJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	ResourceSettings *CreateDIJobRequestResourceSettings `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	// This parameter is required.
	SourceDataSourceSettings []*CreateDIJobRequestSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// This parameter is required.
	TableMappings       []*CreateDIJobRequestTableMappings       `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	TransformationRules []*CreateDIJobRequestTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequest) SetDescription(v string) *CreateDIJobRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobRequest) SetDestinationDataSourceSettings(v []*CreateDIJobRequestDestinationDataSourceSettings) *CreateDIJobRequest {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetDestinationDataSourceType(v string) *CreateDIJobRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobRequest) SetJobName(v string) *CreateDIJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobRequest) SetJobSettings(v *CreateDIJobRequestJobSettings) *CreateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *CreateDIJobRequest) SetMigrationType(v string) *CreateDIJobRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobRequest) SetProjectId(v int64) *CreateDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobRequest) SetResourceSettings(v *CreateDIJobRequestResourceSettings) *CreateDIJobRequest {
	s.ResourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetSourceDataSourceSettings(v []*CreateDIJobRequestSourceDataSourceSettings) *CreateDIJobRequest {
	s.SourceDataSourceSettings = v
	return s
}

func (s *CreateDIJobRequest) SetSourceDataSourceType(v string) *CreateDIJobRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobRequest) SetTableMappings(v []*CreateDIJobRequestTableMappings) *CreateDIJobRequest {
	s.TableMappings = v
	return s
}

func (s *CreateDIJobRequest) SetTransformationRules(v []*CreateDIJobRequestTransformationRules) *CreateDIJobRequest {
	s.TransformationRules = v
	return s
}

type CreateDIJobRequestDestinationDataSourceSettings struct {
	// example:
	//
	// holo_datasource_1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
}

func (s CreateDIJobRequestDestinationDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestDestinationDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

type CreateDIJobRequestJobSettings struct {
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings        *string                                                `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	ColumnDataTypeSettings []*CreateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	CycleScheduleSettings  *CreateDIJobRequestJobSettingsCycleScheduleSettings    `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	DdlHandlingSettings    []*CreateDIJobRequestJobSettingsDdlHandlingSettings    `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	RuntimeSettings        []*CreateDIJobRequestJobSettingsRuntimeSettings        `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequestJobSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettings) SetChannelSettings(v string) *CreateDIJobRequestJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetColumnDataTypeSettings(v []*CreateDIJobRequestJobSettingsColumnDataTypeSettings) *CreateDIJobRequestJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetCycleScheduleSettings(v *CreateDIJobRequestJobSettingsCycleScheduleSettings) *CreateDIJobRequestJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetDdlHandlingSettings(v []*CreateDIJobRequestJobSettingsDdlHandlingSettings) *CreateDIJobRequestJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *CreateDIJobRequestJobSettings) SetRuntimeSettings(v []*CreateDIJobRequestJobSettingsRuntimeSettings) *CreateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

type CreateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s CreateDIJobRequestJobSettingsColumnDataTypeSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *CreateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

type CreateDIJobRequestJobSettingsCycleScheduleSettings struct {
	// example:
	//
	// Full
	CycleMigrationType *string `json:"CycleMigrationType,omitempty" xml:"CycleMigrationType,omitempty"`
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s CreateDIJobRequestJobSettingsCycleScheduleSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *CreateDIJobRequestJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

type CreateDIJobRequestJobSettingsDdlHandlingSettings struct {
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// AddColumn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDIJobRequestJobSettingsDdlHandlingSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) SetAction(v string) *CreateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsDdlHandlingSettings) SetType(v string) *CreateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

type CreateDIJobRequestJobSettingsRuntimeSettings struct {
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDIJobRequestJobSettingsRuntimeSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) SetName(v string) *CreateDIJobRequestJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *CreateDIJobRequestJobSettingsRuntimeSettings) SetValue(v string) *CreateDIJobRequestJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

type CreateDIJobRequestResourceSettings struct {
	OfflineResourceSettings  *CreateDIJobRequestResourceSettingsOfflineResourceSettings  `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	RealtimeResourceSettings *CreateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	ScheduleResourceSettings *CreateDIJobRequestResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s CreateDIJobRequestResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *CreateDIJobRequestResourceSettingsOfflineResourceSettings) *CreateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) *CreateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *CreateDIJobRequestResourceSettings) SetScheduleResourceSettings(v *CreateDIJobRequestResourceSettingsScheduleResourceSettings) *CreateDIJobRequestResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

type CreateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsOfflineResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type CreateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsRealtimeResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type CreateDIJobRequestResourceSettingsScheduleResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1579085295030
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s CreateDIJobRequestResourceSettingsScheduleResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *CreateDIJobRequestResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *CreateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type CreateDIJobRequestSourceDataSourceSettings struct {
	// example:
	//
	// mysql_datasource_1
	DataSourceName       *string                                                         `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceProperties *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
}

func (s CreateDIJobRequestSourceDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceName(v string) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettings) SetDataSourceProperties(v *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) *CreateDIJobRequestSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

type CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties struct {
	// example:
	//
	// UTF-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetEncoding(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.Encoding = &v
	return s
}

func (s *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties) SetTimezone(v string) *CreateDIJobRequestSourceDataSourceSettingsDataSourceProperties {
	s.Timezone = &v
	return s
}

type CreateDIJobRequestTableMappings struct {
	SourceObjectSelectionRules []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	TransformationRules        []*CreateDIJobRequestTableMappingsTransformationRules        `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s CreateDIJobRequestTableMappings) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTableMappings) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappings) SetSourceObjectSelectionRules(v []*CreateDIJobRequestTableMappingsSourceObjectSelectionRules) *CreateDIJobRequestTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *CreateDIJobRequestTableMappings) SetTransformationRules(v []*CreateDIJobRequestTableMappingsTransformationRules) *CreateDIJobRequestTableMappings {
	s.TransformationRules = v
	return s
}

type CreateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetAction(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *CreateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

type CreateDIJobRequestTableMappingsTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s CreateDIJobRequestTableMappingsTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleActionType(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleName(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *CreateDIJobRequestTableMappingsTransformationRules) SetRuleTargetType(v string) *CreateDIJobRequestTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

type CreateDIJobRequestTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s CreateDIJobRequestTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobRequestTransformationRules) GoString() string {
	return s.String()
}

func (s *CreateDIJobRequestTransformationRules) SetRuleActionType(v string) *CreateDIJobRequestTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleExpression(v string) *CreateDIJobRequestTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleName(v string) *CreateDIJobRequestTransformationRules {
	s.RuleName = &v
	return s
}

func (s *CreateDIJobRequestTransformationRules) SetRuleTargetType(v string) *CreateDIJobRequestTransformationRules {
	s.RuleTargetType = &v
	return s
}

type CreateDIJobShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DestinationDataSourceSettingsShrink *string `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName           *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// This parameter is required.
	SourceDataSourceSettingsShrink *string `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// This parameter is required.
	TableMappingsShrink       *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s CreateDIJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobShrinkRequest) SetDescription(v string) *CreateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobName(v string) *CreateDIJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetMigrationType(v string) *CreateDIJobShrinkRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetProjectId(v int64) *CreateDIJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

type CreateDIJobResponseBody struct {
	// example:
	//
	// 11792
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// 4F6AB6B3-41FB-5EBB-AFB2-0C98D49DA2BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDIJobResponseBody) SetDIJobId(v int64) *CreateDIJobResponseBody {
	s.DIJobId = &v
	return s
}

func (s *CreateDIJobResponseBody) SetRequestId(v string) *CreateDIJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDIJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDIJobResponse) SetHeaders(v map[string]*string) *CreateDIJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDIJobResponse) SetStatusCode(v int32) *CreateDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDIJobResponse) SetBody(v *CreateDIJobResponseBody) *CreateDIJobResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// this is a holo datasource
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// demo_holo_datasource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetConnectionProperties(v string) *CreateDataSourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *CreateDataSourceRequest) SetConnectionPropertiesMode(v string) *CreateDataSourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetProjectId(v int64) *CreateDataSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

type CreateDataSourceResponseBody struct {
	// example:
	//
	// 22130
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// B62EC203-B39E-5DC1-B5B8-EB3C61707009
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetId(v int64) *CreateDataSourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetStatusCode(v int32) *CreateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type CreateDataSourceSharedRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 144544
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 1107550004253538
	SharedUser *string `json:"SharedUser,omitempty" xml:"SharedUser,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 106560
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s CreateDataSourceSharedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceSharedRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleRequest) SetDataSourceId(v int64) *CreateDataSourceSharedRuleRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetEnvType(v string) *CreateDataSourceSharedRuleRequest {
	s.EnvType = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetSharedUser(v string) *CreateDataSourceSharedRuleRequest {
	s.SharedUser = &v
	return s
}

func (s *CreateDataSourceSharedRuleRequest) SetTargetProjectId(v int64) *CreateDataSourceSharedRuleRequest {
	s.TargetProjectId = &v
	return s
}

type CreateDataSourceSharedRuleResponseBody struct {
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 46F594E6-84AB-5FA5-8144-6F3D149961E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceSharedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceSharedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleResponseBody) SetId(v int64) *CreateDataSourceSharedRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponseBody) SetRequestId(v string) *CreateDataSourceSharedRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataSourceSharedRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceSharedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceSharedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceSharedRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleResponse) SetHeaders(v map[string]*string) *CreateDataSourceSharedRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) SetStatusCode(v int32) *CreateDataSourceSharedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) SetBody(v *CreateDataSourceSharedRuleResponseBody) *CreateDataSourceSharedRuleResponse {
	s.Body = v
	return s
}

type CreateDeploymentRequest struct {
	// The description of the process.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of entities to which you want to apply the process.
	//
	// >  A process can be applied to only a single entity and its child entities. If you specify multiple entities in the array, the process is applied only to the first entity in the array and its child entities. Make sure that the array in your request contains only one element. Extra elements will be ignored.
	//
	// This parameter is required.
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether to deploy or undeploy the entity. Valid values:
	//
	// 	- Online: deploys the entity.
	//
	// 	- Offline: undeploys the entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentRequest) SetDescription(v string) *CreateDeploymentRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentRequest) SetObjectIds(v []*string) *CreateDeploymentRequest {
	s.ObjectIds = v
	return s
}

func (s *CreateDeploymentRequest) SetProjectId(v string) *CreateDeploymentRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeploymentRequest) SetType(v string) *CreateDeploymentRequest {
	s.Type = &v
	return s
}

type CreateDeploymentShrinkRequest struct {
	// The description of the process.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of entities to which you want to apply the process.
	//
	// >  A process can be applied to only a single entity and its child entities. If you specify multiple entities in the array, the process is applied only to the first entity in the array and its child entities. Make sure that the array in your request contains only one element. Extra elements will be ignored.
	//
	// This parameter is required.
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether to deploy or undeploy the entity. Valid values:
	//
	// 	- Online: deploys the entity.
	//
	// 	- Offline: undeploys the entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentShrinkRequest) SetDescription(v string) *CreateDeploymentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetObjectIdsShrink(v string) *CreateDeploymentShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetProjectId(v string) *CreateDeploymentShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetType(v string) *CreateDeploymentShrinkRequest {
	s.Type = &v
	return s
}

type CreateDeploymentResponseBody struct {
	// The ID of the process.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponseBody) SetId(v string) *CreateDeploymentResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetRequestId(v string) *CreateDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponse) SetHeaders(v map[string]*string) *CreateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentResponse) SetStatusCode(v int32) *CreateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentResponse) SetBody(v *CreateDeploymentResponseBody) *CreateDeploymentResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the UDF. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetProjectId(v string) *CreateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFunctionRequest) SetSpec(v string) *CreateFunctionRequest {
	s.Spec = &v
	return s
}

type CreateFunctionResponseBody struct {
	// The ID of the UDF.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AE49C88D-5BEE-5ADD-8B8C-C4BBC0D7XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetId(v string) *CreateFunctionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

type CreateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetStatusCode(v int32) *CreateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateNetworkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eb870033-74c8-4b1b-9664-04c4e7cc3465
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRequest) SetClientToken(v string) *CreateNetworkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkRequest) SetResourceGroupId(v string) *CreateNetworkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateNetworkRequest) SetVpcId(v string) *CreateNetworkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkRequest) SetVswitchId(v string) *CreateNetworkRequest {
	s.VswitchId = &v
	return s
}

type CreateNetworkResponseBody struct {
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponseBody) SetNetworkId(v int64) *CreateNetworkResponseBody {
	s.NetworkId = &v
	return s
}

func (s *CreateNetworkResponseBody) SetRequestId(v string) *CreateNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkResponseBody) SetSuccess(v bool) *CreateNetworkResponseBody {
	s.Success = &v
	return s
}

type CreateNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkResponse) SetHeaders(v map[string]*string) *CreateNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkResponse) SetStatusCode(v int32) *CreateNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkResponse) SetBody(v *CreateNetworkResponseBody) *CreateNetworkResponse {
	s.Body = v
	return s
}

type CreateNodeRequest struct {
	// The container ID. If you want to create a node in a container, you must configure this parameter to specify the container. The container can be a workflow or a node in a container.
	//
	// >  If you configure this parameter, the path field defined in FlowSpec becomes invalid.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scene of the node. This parameter determines the location (the DataStudio pane or the Manual pane) of the node. You can set this parameter to DATAWORKS_MANUAL_WORKFLOW only if the ContainerId parameter is configured and the container specified by ContainerId is a manually triggered workflow.
	//
	// Valid values:
	//
	// 	- DATAWORKS_PROJECT
	//
	// 	- DATAWORKS_MANUAL_WORKFLOW
	//
	// 	- DATAWORKS_MANUAL_TASK
	//
	// This parameter is required.
	//
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The FlowSpec field information about the node. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRequest) SetContainerId(v string) *CreateNodeRequest {
	s.ContainerId = &v
	return s
}

func (s *CreateNodeRequest) SetProjectId(v string) *CreateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateNodeRequest) SetScene(v string) *CreateNodeRequest {
	s.Scene = &v
	return s
}

func (s *CreateNodeRequest) SetSpec(v string) *CreateNodeRequest {
	s.Spec = &v
	return s
}

type CreateNodeResponseBody struct {
	// The ID of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeResponseBody) SetId(v string) *CreateNodeResponseBody {
	s.Id = &v
	return s
}

func (s *CreateNodeResponseBody) SetRequestId(v string) *CreateNodeResponseBody {
	s.RequestId = &v
	return s
}

type CreateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeResponse) SetHeaders(v map[string]*string) *CreateNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeResponse) SetStatusCode(v int32) *CreateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeResponse) SetBody(v *CreateNodeResponseBody) *CreateNodeResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string                                   `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*CreateProjectRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	Description           *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// true
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetAliyunResourceGroupId(v string) *CreateProjectRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateProjectRequest) SetAliyunResourceTags(v []*CreateProjectRequestAliyunResourceTags) *CreateProjectRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetDevEnvironmentEnabled(v bool) *CreateProjectRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *CreateProjectRequest) SetDevRoleDisabled(v bool) *CreateProjectRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *CreateProjectRequest) SetDisplayName(v string) *CreateProjectRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetPaiTaskEnabled(v bool) *CreateProjectRequest {
	s.PaiTaskEnabled = &v
	return s
}

type CreateProjectRequestAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProjectRequestAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestAliyunResourceTags) SetKey(v string) *CreateProjectRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *CreateProjectRequestAliyunResourceTags) SetValue(v string) *CreateProjectRequestAliyunResourceTags {
	s.Value = &v
	return s
}

type CreateProjectShrinkRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId    *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// true
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// This parameter is required.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) SetAliyunResourceGroupId(v string) *CreateProjectShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetAliyunResourceTagsShrink(v string) *CreateProjectShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDescription(v string) *CreateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDevEnvironmentEnabled(v bool) *CreateProjectShrinkRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDevRoleDisabled(v bool) *CreateProjectShrinkRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDisplayName(v string) *CreateProjectShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetName(v string) *CreateProjectShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetPaiTaskEnabled(v bool) *CreateProjectShrinkRequest {
	s.PaiTaskEnabled = &v
	return s
}

type CreateProjectResponseBody struct {
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectId(v int64) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateProjectMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 24054
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberRequest) SetProjectId(v int64) *CreateProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectMemberRequest) SetRoleCodes(v []*string) *CreateProjectMemberRequest {
	s.RoleCodes = v
	return s
}

func (s *CreateProjectMemberRequest) SetUserId(v string) *CreateProjectMemberRequest {
	s.UserId = &v
	return s
}

type CreateProjectMemberShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 24054
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateProjectMemberShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMemberShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberShrinkRequest) SetProjectId(v int64) *CreateProjectMemberShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectMemberShrinkRequest) SetRoleCodesShrink(v string) *CreateProjectMemberShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *CreateProjectMemberShrinkRequest) SetUserId(v string) *CreateProjectMemberShrinkRequest {
	s.UserId = &v
	return s
}

type CreateProjectMemberResponseBody struct {
	// example:
	//
	// 2B2F0B26-9253-5780-B6DB-F1A886D44D6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberResponseBody) SetRequestId(v string) *CreateProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberResponse) SetHeaders(v map[string]*string) *CreateProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectMemberResponse) SetStatusCode(v int32) *CreateProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectMemberResponse) SetBody(v *CreateProjectMemberResponseBody) *CreateProjectMemberResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the file resource. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetProjectId(v string) *CreateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceRequest) SetSpec(v string) *CreateResourceRequest {
	s.Spec = &v
	return s
}

type CreateResourceResponseBody struct {
	// The ID of the file resource.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5B97987-66EA-5563-9599-A2752292XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetId(v string) *CreateResourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type CreateResourceGroupRequest struct {
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eb870033-74c8-4b1b-9664-04c4e7cc3465
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PaymentDuration *int32 `json:"PaymentDuration,omitempty" xml:"PaymentDuration,omitempty"`
	// example:
	//
	// Month
	PaymentDurationUnit *string `json:"PaymentDurationUnit,omitempty" xml:"PaymentDurationUnit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 创建用于普通任务的通用资源组
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 2
	Spec *int32 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) SetAutoRenew(v bool) *CreateResourceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateResourceGroupRequest) SetClientToken(v string) *CreateResourceGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPaymentDuration(v int32) *CreateResourceGroupRequest {
	s.PaymentDuration = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPaymentDurationUnit(v string) *CreateResourceGroupRequest {
	s.PaymentDurationUnit = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPaymentType(v string) *CreateResourceGroupRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateResourceGroupRequest) SetRemark(v string) *CreateResourceGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateResourceGroupRequest) SetSpec(v int32) *CreateResourceGroupRequest {
	s.Spec = &v
	return s
}

func (s *CreateResourceGroupRequest) SetVpcId(v string) *CreateResourceGroupRequest {
	s.VpcId = &v
	return s
}

func (s *CreateResourceGroupRequest) SetVswitchId(v string) *CreateResourceGroupRequest {
	s.VswitchId = &v
	return s
}

type CreateResourceGroupResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId          *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupOrder *CreateResourceGroupResponseBodyResourceGroupOrder `json:"ResourceGroupOrder,omitempty" xml:"ResourceGroupOrder,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroupOrder(v *CreateResourceGroupResponseBodyResourceGroupOrder) *CreateResourceGroupResponseBody {
	s.ResourceGroupOrder = v
	return s
}

func (s *CreateResourceGroupResponseBody) SetSuccess(v bool) *CreateResourceGroupResponseBody {
	s.Success = &v
	return s
}

type CreateResourceGroupResponseBodyResourceGroupOrder struct {
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2391982058XXXXX
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// c442b330-3b10-4584-959e-736e4edXXXXX
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
}

func (s CreateResourceGroupResponseBodyResourceGroupOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroupOrder) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) SetId(v string) *CreateResourceGroupResponseBodyResourceGroupOrder {
	s.Id = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) SetOrderId(v int64) *CreateResourceGroupResponseBodyResourceGroupOrder {
	s.OrderId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupOrder) SetOrderInstanceId(v string) *CreateResourceGroupResponseBodyResourceGroupOrder {
	s.OrderInstanceId = &v
	return s
}

type CreateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponse) SetHeaders(v map[string]*string) *CreateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceGroupResponse) SetStatusCode(v int32) *CreateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceGroupResponse) SetBody(v *CreateResourceGroupResponseBody) *CreateResourceGroupResponse {
	s.Body = v
	return s
}

type CreateRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s CreateRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteRequest) SetDestinationCidr(v string) *CreateRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *CreateRouteRequest) SetNetworkId(v int64) *CreateRouteRequest {
	s.NetworkId = &v
	return s
}

type CreateRouteResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteResponseBody) SetRequestId(v string) *CreateRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteResponseBody) SetRouteId(v int64) *CreateRouteResponseBody {
	s.RouteId = &v
	return s
}

func (s *CreateRouteResponseBody) SetSuccess(v bool) *CreateRouteResponseBody {
	s.Success = &v
	return s
}

type CreateRouteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateRouteResponse) SetHeaders(v map[string]*string) *CreateRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateRouteResponse) SetStatusCode(v int32) *CreateRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRouteResponse) SetBody(v *CreateRouteResponseBody) *CreateRouteResponse {
	s.Body = v
	return s
}

type CreateWorkflowDefinitionRequest struct {
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the workflow. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow/).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionRequest) SetProjectId(v string) *CreateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkflowDefinitionRequest) SetSpec(v string) *CreateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type CreateWorkflowDefinitionResponseBody struct {
	// The ID of the workflow.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0EF298E5-0940-5AC7-9CB0-65025070XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponseBody) SetId(v string) *CreateWorkflowDefinitionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWorkflowDefinitionResponseBody) SetRequestId(v string) *CreateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

type CreateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *CreateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetStatusCode(v int32) *CreateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetBody(v *CreateWorkflowDefinitionResponseBody) *CreateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type DeleteDIAlarmRuleRequest struct {
	// The ID of the alert rule.
	//
	// example:
	//
	// 2
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 1
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
}

func (s DeleteDIAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDIAlarmRuleRequest) SetDIAlarmRuleId(v int64) *DeleteDIAlarmRuleRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *DeleteDIAlarmRuleRequest) SetDIJobId(v int64) *DeleteDIAlarmRuleRequest {
	s.DIJobId = &v
	return s
}

type DeleteDIAlarmRuleResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDIAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDIAlarmRuleResponseBody) SetRequestId(v string) *DeleteDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDIAlarmRuleResponseBody) SetSuccess(v bool) *DeleteDIAlarmRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteDIAlarmRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDIAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDIAlarmRuleResponse) SetHeaders(v map[string]*string) *DeleteDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDIAlarmRuleResponse) SetStatusCode(v int32) *DeleteDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDIAlarmRuleResponse) SetBody(v *DeleteDIAlarmRuleResponseBody) *DeleteDIAlarmRuleResponse {
	s.Body = v
	return s
}

type DeleteDIJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11126
	DIJobId   *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDIJobRequest) SetDIJobId(v int64) *DeleteDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *DeleteDIJobRequest) SetProjectId(v int64) *DeleteDIJobRequest {
	s.ProjectId = &v
	return s
}

type DeleteDIJobResponseBody struct {
	// example:
	//
	// D33D4A51-5845-579A-B4BA-FAADD0F83D53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDIJobResponseBody) SetRequestId(v string) *DeleteDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDIJobResponseBody) SetSuccess(v bool) *DeleteDIJobResponseBody {
	s.Success = &v
	return s
}

type DeleteDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDIJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDIJobResponse) SetHeaders(v map[string]*string) *DeleteDIJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDIJobResponse) SetStatusCode(v int32) *DeleteDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDIJobResponse) SetBody(v *DeleteDIJobResponseBody) *DeleteDIJobResponse {
	s.Body = v
	return s
}

type DeleteDataSourceRequest struct {
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) SetId(v int64) *DeleteDataSourceRequest {
	s.Id = &v
	return s
}

type DeleteDataSourceResponseBody struct {
	// example:
	//
	// B56432E0-2112-5C97-88D0-AA0AE5C75C74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetSuccess(v bool) *DeleteDataSourceResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type DeleteDataSourceSharedRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 22127
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataSourceSharedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceSharedRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleRequest) SetId(v int64) *DeleteDataSourceSharedRuleRequest {
	s.Id = &v
	return s
}

type DeleteDataSourceSharedRuleResponseBody struct {
	// example:
	//
	// 64B-587A-8CED-969E1973887FXXX-TT
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSourceSharedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceSharedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleResponseBody) SetRequestId(v string) *DeleteDataSourceSharedRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponseBody) SetSuccess(v bool) *DeleteDataSourceSharedRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteDataSourceSharedRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataSourceSharedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataSourceSharedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceSharedRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceSharedRuleResponse) SetHeaders(v map[string]*string) *DeleteDataSourceSharedRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) SetStatusCode(v int32) *DeleteDataSourceSharedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceSharedRuleResponse) SetBody(v *DeleteDataSourceSharedRuleResponseBody) *DeleteDataSourceSharedRuleResponse {
	s.Body = v
	return s
}

type DeleteFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) SetId(v string) *DeleteFunctionRequest {
	s.Id = &v
	return s
}

func (s *DeleteFunctionRequest) SetProjectId(v string) *DeleteFunctionRequest {
	s.ProjectId = &v
	return s
}

type DeleteFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetSuccess(v bool) *DeleteFunctionResponseBody {
	s.Success = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetStatusCode(v int32) *DeleteFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResponse) SetBody(v *DeleteFunctionResponseBody) *DeleteFunctionResponse {
	s.Body = v
	return s
}

type DeleteNetworkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRequest) SetId(v int64) *DeleteNetworkRequest {
	s.Id = &v
	return s
}

type DeleteNetworkResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponseBody) SetRequestId(v string) *DeleteNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkResponseBody) SetSuccess(v bool) *DeleteNetworkResponseBody {
	s.Success = &v
	return s
}

type DeleteNetworkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkResponse) SetHeaders(v map[string]*string) *DeleteNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkResponse) SetStatusCode(v int32) *DeleteNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkResponse) SetBody(v *DeleteNetworkResponseBody) *DeleteNetworkResponse {
	s.Body = v
	return s
}

type DeleteNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) SetId(v string) *DeleteNodeRequest {
	s.Id = &v
	return s
}

func (s *DeleteNodeRequest) SetProjectId(v string) *DeleteNodeRequest {
	s.ProjectId = &v
	return s
}

type DeleteNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A1E54497-5122-505E-91C6-BAC14980XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true\\
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetSuccess(v bool) *DeleteNodeResponseBody {
	s.Success = &v
	return s
}

type DeleteNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponse) SetHeaders(v map[string]*string) *DeleteNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeResponse) SetStatusCode(v int32) *DeleteNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeResponse) SetBody(v *DeleteNodeResponseBody) *DeleteNodeResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetId(v int64) *DeleteProjectRequest {
	s.Id = &v
	return s
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteProjectMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 534752
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberRequest) SetProjectId(v int64) *DeleteProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectMemberRequest) SetUserId(v string) *DeleteProjectMemberRequest {
	s.UserId = &v
	return s
}

type DeleteProjectMemberResponseBody struct {
	// example:
	//
	// 1FF0465F-209C-5964-8F30-FAF21B677CC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberResponseBody) SetRequestId(v string) *DeleteProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberResponse) SetHeaders(v map[string]*string) *DeleteProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectMemberResponse) SetStatusCode(v int32) *DeleteProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectMemberResponse) SetBody(v *DeleteProjectMemberResponseBody) *DeleteProjectMemberResponse {
	s.Body = v
	return s
}

type DeleteResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) SetId(v string) *DeleteResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteResourceRequest) SetProjectId(v string) *DeleteResourceRequest {
	s.ProjectId = &v
	return s
}

type DeleteResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetSuccess(v bool) *DeleteResourceResponseBody {
	s.Success = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type DeleteResourceGroupRequest struct {
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupRequest) SetId(v string) *DeleteResourceGroupRequest {
	s.Id = &v
	return s
}

type DeleteResourceGroupResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetSuccess(v bool) *DeleteResourceGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupResponse) SetBody(v *DeleteResourceGroupResponseBody) *DeleteResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteRequest) SetId(v int64) *DeleteRouteRequest {
	s.Id = &v
	return s
}

type DeleteRouteResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRouteResponseBody) SetRequestId(v string) *DeleteRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRouteResponseBody) SetSuccess(v bool) *DeleteRouteResponseBody {
	s.Success = &v
	return s
}

type DeleteRouteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteResponse) SetHeaders(v map[string]*string) *DeleteRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteResponse) SetStatusCode(v int32) *DeleteRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteResponse) SetBody(v *DeleteRouteResponseBody) *DeleteRouteResponse {
	s.Body = v
	return s
}

type DeleteWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionRequest) SetId(v string) *DeleteWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *DeleteWorkflowDefinitionRequest) SetProjectId(v string) *DeleteWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type DeleteWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B17730C0-D959-548A-AE23-E754177CXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponseBody) SetRequestId(v string) *DeleteWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponseBody) SetSuccess(v bool) *DeleteWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *DeleteWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetStatusCode(v int32) *DeleteWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetBody(v *DeleteWorkflowDefinitionResponseBody) *DeleteWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type DissociateProjectFromResourceGroupRequest struct {
	// The ID of the workspace from which you want to disassociate the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DissociateProjectFromResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateProjectFromResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromResourceGroupRequest) SetProjectId(v int64) *DissociateProjectFromResourceGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DissociateProjectFromResourceGroupRequest) SetResourceGroupId(v string) *DissociateProjectFromResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

type DissociateProjectFromResourceGroupResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DissociateProjectFromResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateProjectFromResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromResourceGroupResponseBody) SetRequestId(v string) *DissociateProjectFromResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateProjectFromResourceGroupResponseBody) SetSuccess(v bool) *DissociateProjectFromResourceGroupResponseBody {
	s.Success = &v
	return s
}

type DissociateProjectFromResourceGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateProjectFromResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateProjectFromResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateProjectFromResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DissociateProjectFromResourceGroupResponse) SetHeaders(v map[string]*string) *DissociateProjectFromResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DissociateProjectFromResourceGroupResponse) SetStatusCode(v int32) *DissociateProjectFromResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateProjectFromResourceGroupResponse) SetBody(v *DissociateProjectFromResourceGroupResponseBody) *DissociateProjectFromResourceGroupResponse {
	s.Body = v
	return s
}

type ExecDeploymentStageRequest struct {
	// The code of the stage in the process. You can call the GetDeployment operation to query the code.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the process.
	//
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecDeploymentStageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageRequest) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageRequest) SetCode(v string) *ExecDeploymentStageRequest {
	s.Code = &v
	return s
}

func (s *ExecDeploymentStageRequest) SetId(v string) *ExecDeploymentStageRequest {
	s.Id = &v
	return s
}

func (s *ExecDeploymentStageRequest) SetProjectId(v string) *ExecDeploymentStageRequest {
	s.ProjectId = &v
	return s
}

type ExecDeploymentStageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	//     **
	//
	//     **Note:*	- The value of this parameter indicates only whether the stage is triggered but does not indicate whether the execution of the stage is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecDeploymentStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageResponseBody) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageResponseBody) SetRequestId(v string) *ExecDeploymentStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecDeploymentStageResponseBody) SetSuccess(v bool) *ExecDeploymentStageResponseBody {
	s.Success = &v
	return s
}

type ExecDeploymentStageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecDeploymentStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDeploymentStageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageResponse) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageResponse) SetHeaders(v map[string]*string) *ExecDeploymentStageResponse {
	s.Headers = v
	return s
}

func (s *ExecDeploymentStageResponse) SetStatusCode(v int32) *ExecDeploymentStageResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecDeploymentStageResponse) SetBody(v *ExecDeploymentStageResponseBody) *ExecDeploymentStageResponse {
	s.Body = v
	return s
}

type GetDIJobRequest struct {
	// example:
	//
	// 11588
	DIJobId   *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// true
	WithDetails *bool `json:"WithDetails,omitempty" xml:"WithDetails,omitempty"`
}

func (s GetDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobRequest) SetDIJobId(v int64) *GetDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobRequest) SetProjectId(v int64) *GetDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *GetDIJobRequest) SetWithDetails(v bool) *GetDIJobRequest {
	s.WithDetails = &v
	return s
}

type GetDIJobResponseBody struct {
	PagingInfo *GetDIJobResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADE490
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBody) SetPagingInfo(v *GetDIJobResponseBodyPagingInfo) *GetDIJobResponseBody {
	s.PagingInfo = v
	return s
}

func (s *GetDIJobResponseBody) SetRequestId(v string) *GetDIJobResponseBody {
	s.RequestId = &v
	return s
}

type GetDIJobResponseBodyPagingInfo struct {
	// example:
	//
	// 32601
	DIJobId *string `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// description
	Description                   *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationDataSourceSettings []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty" type:"Repeated"`
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// example:
	//
	// imp_ods_dms_det_dealer_info_df
	JobName     *string                                    `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettings *GetDIJobResponseBodyPagingInfoJobSettings `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	JobStatus   *string                                    `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// example:
	//
	// 98330
	ProjectId                *int64                                                    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceSettings         *GetDIJobResponseBodyPagingInfoResourceSettings           `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	SourceDataSourceSettings []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty" type:"Repeated"`
	// example:
	//
	// Mysql
	SourceDataSourceType *string                                              `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	TableMappings        []*GetDIJobResponseBodyPagingInfoTableMappings       `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	TransformationRules  []*GetDIJobResponseBodyPagingInfoTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfo) SetDIJobId(v string) *GetDIJobResponseBodyPagingInfo {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDescription(v string) *GetDIJobResponseBodyPagingInfo {
	s.Description = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDestinationDataSourceSettings(v []*GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.DestinationDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetDestinationDataSourceType(v string) *GetDIJobResponseBodyPagingInfo {
	s.DestinationDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobName(v string) *GetDIJobResponseBodyPagingInfo {
	s.JobName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobSettings(v *GetDIJobResponseBodyPagingInfoJobSettings) *GetDIJobResponseBodyPagingInfo {
	s.JobSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetJobStatus(v string) *GetDIJobResponseBodyPagingInfo {
	s.JobStatus = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetMigrationType(v string) *GetDIJobResponseBodyPagingInfo {
	s.MigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetProjectId(v int64) *GetDIJobResponseBodyPagingInfo {
	s.ProjectId = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.ResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetSourceDataSourceSettings(v []*GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) *GetDIJobResponseBodyPagingInfo {
	s.SourceDataSourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetSourceDataSourceType(v string) *GetDIJobResponseBodyPagingInfo {
	s.SourceDataSourceType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetTableMappings(v []*GetDIJobResponseBodyPagingInfoTableMappings) *GetDIJobResponseBodyPagingInfo {
	s.TableMappings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfo) SetTransformationRules(v []*GetDIJobResponseBodyPagingInfoTransformationRules) *GetDIJobResponseBodyPagingInfo {
	s.TransformationRules = v
	return s
}

type GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings struct {
	// example:
	//
	// dw_mysql
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyPagingInfoDestinationDataSourceSettings {
	s.DataSourceName = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettings struct {
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings        *string                                                            `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	ColumnDataTypeSettings []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	CycleScheduleSettings  *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings    `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	DdlHandlingSettings    []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings    `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	RuntimeSettings        []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings        `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetChannelSettings(v string) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetColumnDataTypeSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetCycleScheduleSettings(v *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetDdlHandlingSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettings) SetRuntimeSettings(v []*GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) *GetDIJobResponseBodyPagingInfoJobSettings {
	s.RuntimeSettings = v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings struct {
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings struct {
	// example:
	//
	// Full
	CycleMigrationType *string `json:"CycleMigrationType,omitempty" xml:"CycleMigrationType,omitempty"`
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) SetCycleMigrationType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	s.CycleMigrationType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *GetDIJobResponseBodyPagingInfoJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings struct {
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// CreateTable
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) SetAction(v string) *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings) SetType(v string) *GetDIJobResponseBodyPagingInfoJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

type GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings struct {
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) SetName(v string) *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings) SetValue(v string) *GetDIJobResponseBodyPagingInfoJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettings struct {
	OfflineResourceSettings  *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings  `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	RealtimeResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	ScheduleResourceSettings *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetOfflineResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetRealtimeResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettings) SetScheduleResourceSettings(v *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) *GetDIJobResponseBodyPagingInfoResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_7708_1667792816832
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1579085295030
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *float64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1718359176885
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) SetRequestedCu(v float64) *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *GetDIJobResponseBodyPagingInfoResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettings struct {
	// example:
	//
	// dw_mysql
	DataSourceName       *string                                                                     `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	DataSourceProperties *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties `json:"DataSourceProperties,omitempty" xml:"DataSourceProperties,omitempty" type:"Struct"`
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) SetDataSourceName(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	s.DataSourceName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings) SetDataSourceProperties(v *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettings {
	s.DataSourceProperties = v
	return s
}

type GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties struct {
	// example:
	//
	// UTF-8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) SetEncoding(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	s.Encoding = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties) SetTimezone(v string) *GetDIJobResponseBodyPagingInfoSourceDataSourceSettingsDataSourceProperties {
	s.Timezone = &v
	return s
}

type GetDIJobResponseBodyPagingInfoTableMappings struct {
	SourceObjectSelectionRules []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	TransformationRules        []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules        `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappings) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappings) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) SetSourceObjectSelectionRules(v []*GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) *GetDIJobResponseBodyPagingInfoTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappings) SetTransformationRules(v []*GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) *GetDIJobResponseBodyPagingInfoTableMappings {
	s.TransformationRules = v
	return s
}

type GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules struct {
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetAction(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetExpression(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

type GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules struct {
	// example:
	//
	// AddColumn
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyPagingInfoTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

type GetDIJobResponseBodyPagingInfoTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s GetDIJobResponseBodyPagingInfoTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponseBodyPagingInfoTransformationRules) GoString() string {
	return s.String()
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleActionType(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleExpression(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleName(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleName = &v
	return s
}

func (s *GetDIJobResponseBodyPagingInfoTransformationRules) SetRuleTargetType(v string) *GetDIJobResponseBodyPagingInfoTransformationRules {
	s.RuleTargetType = &v
	return s
}

type GetDIJobResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobResponse) GoString() string {
	return s.String()
}

func (s *GetDIJobResponse) SetHeaders(v map[string]*string) *GetDIJobResponse {
	s.Headers = v
	return s
}

func (s *GetDIJobResponse) SetStatusCode(v int32) *GetDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIJobResponse) SetBody(v *GetDIJobResponseBody) *GetDIJobResponse {
	s.Body = v
	return s
}

type GetDIJobLogRequest struct {
	// The ID of the synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The failover ID.
	//
	// example:
	//
	// 10
	FailoverId *int64 `json:"FailoverId,omitempty" xml:"FailoverId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 6153616438
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDIJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetDIJobLogRequest) SetDIJobId(v int64) *GetDIJobLogRequest {
	s.DIJobId = &v
	return s
}

func (s *GetDIJobLogRequest) SetFailoverId(v int64) *GetDIJobLogRequest {
	s.FailoverId = &v
	return s
}

func (s *GetDIJobLogRequest) SetInstanceId(v int64) *GetDIJobLogRequest {
	s.InstanceId = &v
	return s
}

type GetDIJobLogResponseBody struct {
	// The log.
	//
	// example:
	//
	// >>>>>>>> stdout:n++++++++++++++++++executing sql: create database if not exists jindo_test location \\"oss://pangbei-hdfs/tmp/hive\\" n++n
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobLogResponseBody) SetLog(v string) *GetDIJobLogResponseBody {
	s.Log = &v
	return s
}

func (s *GetDIJobLogResponseBody) SetRequestId(v string) *GetDIJobLogResponseBody {
	s.RequestId = &v
	return s
}

type GetDIJobLogResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDIJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDIJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDIJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetDIJobLogResponse) SetHeaders(v map[string]*string) *GetDIJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetDIJobLogResponse) SetStatusCode(v int32) *GetDIJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDIJobLogResponse) SetBody(v *GetDIJobLogResponseBody) *GetDIJobLogResponse {
	s.Body = v
	return s
}

type GetDataSourceRequest struct {
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceRequest) SetId(v int64) *GetDataSourceRequest {
	s.Id = &v
	return s
}

type GetDataSourceResponseBody struct {
	DataSource *GetDataSourceResponseBodyDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) SetDataSource(v *GetDataSourceResponseBodyDataSource) *GetDataSourceResponseBody {
	s.DataSource = v
	return s
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type GetDataSourceResponseBodyDataSource struct {
	// The connection properties of the data source.
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The mode in which the data source is added. The mode varies based on the data source type. Valid values:
	//
	// 	- InstanceMode: instance mode
	//
	// 	- UrlMode: connection string mode
	//
	// 	- CdhMode: CDH cluster mode
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The time when the data source was added. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698286929333
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who adds the data source.
	//
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 16738
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the data source was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698286929333
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the user who modifies the data source.
	//
	// example:
	//
	// 1107550004253538
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace with which the data source is associated.
	//
	// example:
	//
	// 52660
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique business key of the data source. For example, the unique business key of a Hologres data source is in the ${tenantOwnerId}:${regionId}:${type}:${instanceId}:${database} format.
	//
	// example:
	//
	// 1107550004253538:cn-beijing:holo:hgprecn-cn-x0r3oun4k001:testdb
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataSourceResponseBodyDataSource) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBodyDataSource) SetConnectionProperties(v interface{}) *GetDataSourceResponseBodyDataSource {
	s.ConnectionProperties = v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetConnectionPropertiesMode(v string) *GetDataSourceResponseBodyDataSource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetCreateTime(v int64) *GetDataSourceResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetCreateUser(v string) *GetDataSourceResponseBodyDataSource {
	s.CreateUser = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetDescription(v string) *GetDataSourceResponseBodyDataSource {
	s.Description = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetId(v int64) *GetDataSourceResponseBodyDataSource {
	s.Id = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetModifyTime(v int64) *GetDataSourceResponseBodyDataSource {
	s.ModifyTime = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetModifyUser(v string) *GetDataSourceResponseBodyDataSource {
	s.ModifyUser = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetName(v string) *GetDataSourceResponseBodyDataSource {
	s.Name = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetProjectId(v int64) *GetDataSourceResponseBodyDataSource {
	s.ProjectId = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetQualifiedName(v string) *GetDataSourceResponseBodyDataSource {
	s.QualifiedName = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetType(v string) *GetDataSourceResponseBodyDataSource {
	s.Type = &v
	return s
}

type GetDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponse) SetHeaders(v map[string]*string) *GetDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceResponse) SetStatusCode(v int32) *GetDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceResponse) SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse {
	s.Body = v
	return s
}

type GetDeploymentRequest struct {
	// The ID of the process.
	//
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentRequest) SetId(v string) *GetDeploymentRequest {
	s.Id = &v
	return s
}

func (s *GetDeploymentRequest) SetProjectId(v string) *GetDeploymentRequest {
	s.ProjectId = &v
	return s
}

type GetDeploymentResponseBody struct {
	// The information about the process.
	Pipeline *GetDeploymentResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 08468352-032C-5262-AEDC-68C9FA05XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBody) SetPipeline(v *GetDeploymentResponseBodyPipeline) *GetDeploymentResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetDeploymentResponseBody) SetRequestId(v string) *GetDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type GetDeploymentResponseBodyPipeline struct {
	// The time when the process was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724984066000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the process.
	//
	// example:
	//
	// 137946317766XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned when the process fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the process was modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724984066000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 56160
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The information about stages in the process.
	Stages []*GetDeploymentResponseBodyPipelineStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// The status of the process.
	//
	// Valid values:
	//
	// 	- INIT
	//
	// 	- RUNNING
	//
	// 	- SUCCESS
	//
	// 	- FAIL
	//
	// 	- TERMINATION
	//
	// 	- CANCEL
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeploymentResponseBodyPipeline) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyPipeline) SetCreateTime(v int64) *GetDeploymentResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetCreator(v string) *GetDeploymentResponseBodyPipeline {
	s.Creator = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetId(v string) *GetDeploymentResponseBodyPipeline {
	s.Id = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetMessage(v string) *GetDeploymentResponseBodyPipeline {
	s.Message = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetModifyTime(v int64) *GetDeploymentResponseBodyPipeline {
	s.ModifyTime = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetProjectId(v string) *GetDeploymentResponseBodyPipeline {
	s.ProjectId = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetStages(v []*GetDeploymentResponseBodyPipelineStages) *GetDeploymentResponseBodyPipeline {
	s.Stages = v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetStatus(v string) *GetDeploymentResponseBodyPipeline {
	s.Status = &v
	return s
}

type GetDeploymentResponseBodyPipelineStages struct {
	// The code of the stage.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the stage.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the stage.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The error message returned for the stage.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the stage.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the stage.
	//
	// Valid values:
	//
	// 	- INIT
	//
	// 	- RUNNING
	//
	// 	- SUCCESS
	//
	// 	- FAIL
	//
	// 	- TERMINATION
	//
	// 	- CANCEL
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The step number of the stage.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// The type of the stage.
	//
	// Valid values:
	//
	// 	- DELETE
	//
	// 	- BUILD
	//
	// 	- CHECK
	//
	// 	- DEPLOY
	//
	// 	- OFFLINE
	//
	// example:
	//
	// BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDeploymentResponseBodyPipelineStages) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBodyPipelineStages) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyPipelineStages) SetCode(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Code = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetDescription(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Description = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetDetail(v map[string]interface{}) *GetDeploymentResponseBodyPipelineStages {
	s.Detail = v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetMessage(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Message = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetName(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Name = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetStatus(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Status = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetStep(v int32) *GetDeploymentResponseBodyPipelineStages {
	s.Step = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetType(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Type = &v
	return s
}

type GetDeploymentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponse) SetHeaders(v map[string]*string) *GetDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentResponse) SetStatusCode(v int32) *GetDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentResponse) SetBody(v *GetDeploymentResponseBody) *GetDeploymentResponse {
	s.Body = v
	return s
}

type GetFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) SetId(v string) *GetFunctionRequest {
	s.Id = &v
	return s
}

func (s *GetFunctionRequest) SetProjectId(v string) *GetFunctionRequest {
	s.ProjectId = &v
	return s
}

type GetFunctionResponseBody struct {
	// The information about the UDF.
	Function *GetFunctionResponseBodyFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6CF95929-6D12-5A88-8CC3-4B2F4C2EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBody) SetFunction(v *GetFunctionResponseBodyFunction) *GetFunctionResponseBody {
	s.Function = v
	return s
}

func (s *GetFunctionResponseBody) SetRequestId(v string) *GetFunctionResponseBody {
	s.RequestId = &v
	return s
}

type GetFunctionResponseBodyFunction struct {
	// The time when the UDF was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the UDF.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the UDF was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724506661000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the UDF.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the UDF.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace to which the UDF belongs.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the UDF. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetFunctionResponseBodyFunction) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBodyFunction) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBodyFunction) SetCreateTime(v int64) *GetFunctionResponseBodyFunction {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetId(v string) *GetFunctionResponseBodyFunction {
	s.Id = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetModifyTime(v int64) *GetFunctionResponseBodyFunction {
	s.ModifyTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetName(v string) *GetFunctionResponseBodyFunction {
	s.Name = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetOwner(v string) *GetFunctionResponseBodyFunction {
	s.Owner = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetProjectId(v string) *GetFunctionResponseBodyFunction {
	s.ProjectId = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetSpec(v string) *GetFunctionResponseBodyFunction {
	s.Spec = &v
	return s
}

type GetFunctionResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetStatusCode(v int32) *GetFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResponse) SetBody(v *GetFunctionResponseBody) *GetFunctionResponse {
	s.Body = v
	return s
}

type GetJobStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 70ecdaec-bf21-4c11-8ecb-4f77453ceea8
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

type GetJobStatusResponseBody struct {
	JobStatus *GetJobStatusResponseBodyJobStatus `json:"JobStatus,omitempty" xml:"JobStatus,omitempty" type:"Struct"`
	// example:
	//
	// 5E2BFE96-C0E0-5A98-85C8-633EC803198D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) SetJobStatus(v *GetJobStatusResponseBodyJobStatus) *GetJobStatusResponseBody {
	s.JobStatus = v
	return s
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetJobStatusResponseBodyJobStatus struct {
	// example:
	//
	// False
	Completed *string `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1729063449802
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Not Found
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// C664CDE3-9C0B-5792-B17F-6C543783BBBC
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// Create
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobStatusResponseBodyJobStatus) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBodyJobStatus) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBodyJobStatus) SetCompleted(v string) *GetJobStatusResponseBodyJobStatus {
	s.Completed = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetCreateTime(v string) *GetJobStatusResponseBodyJobStatus {
	s.CreateTime = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetError(v string) *GetJobStatusResponseBodyJobStatus {
	s.Error = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetJobId(v string) *GetJobStatusResponseBodyJobStatus {
	s.JobId = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetJobType(v string) *GetJobStatusResponseBodyJobStatus {
	s.JobType = &v
	return s
}

func (s *GetJobStatusResponseBodyJobStatus) SetStatus(v string) *GetJobStatusResponseBodyJobStatus {
	s.Status = &v
	return s
}

type GetJobStatusResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetHeaders(v map[string]*string) *GetJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetJobStatusResponse) SetStatusCode(v int32) *GetJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobStatusResponse) SetBody(v *GetJobStatusResponseBody) *GetJobStatusResponse {
	s.Body = v
	return s
}

type GetNetworkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkRequest) SetId(v int64) *GetNetworkRequest {
	s.Id = &v
	return s
}

type GetNetworkResponseBody struct {
	Network *GetNetworkResponseBodyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkResponseBody) SetNetwork(v *GetNetworkResponseBodyNetwork) *GetNetworkResponseBody {
	s.Network = v
	return s
}

func (s *GetNetworkResponseBody) SetRequestId(v string) *GetNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkResponseBody) SetSuccess(v bool) *GetNetworkResponseBody {
	s.Success = &v
	return s
}

type GetNetworkResponseBodyNetwork struct {
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-2ze13vamugr7jenXXXXX
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s GetNetworkResponseBodyNetwork) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkResponseBodyNetwork) GoString() string {
	return s.String()
}

func (s *GetNetworkResponseBodyNetwork) SetCreateTime(v int64) *GetNetworkResponseBodyNetwork {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetCreateUser(v string) *GetNetworkResponseBodyNetwork {
	s.CreateUser = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetId(v int64) *GetNetworkResponseBodyNetwork {
	s.Id = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetResourceGroupId(v string) *GetNetworkResponseBodyNetwork {
	s.ResourceGroupId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetSecurityGroupId(v string) *GetNetworkResponseBodyNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetStatus(v string) *GetNetworkResponseBodyNetwork {
	s.Status = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetVpcId(v string) *GetNetworkResponseBodyNetwork {
	s.VpcId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetVswitchId(v string) *GetNetworkResponseBodyNetwork {
	s.VswitchId = &v
	return s
}

type GetNetworkResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkResponse) SetHeaders(v map[string]*string) *GetNetworkResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkResponse) SetStatusCode(v int32) *GetNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkResponse) SetBody(v *GetNetworkResponseBody) *GetNetworkResponse {
	s.Body = v
	return s
}

type GetNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the workspace ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) SetId(v string) *GetNodeRequest {
	s.Id = &v
	return s
}

func (s *GetNodeRequest) SetProjectId(v string) *GetNodeRequest {
	s.ProjectId = &v
	return s
}

type GetNodeResponseBody struct {
	// The information about the node.
	Node *GetNodeResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody {
	s.Node = v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

type GetNodeResponseBodyNode struct {
	// The time when the node was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the node was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the node.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the node.
	//
	// example:
	//
	// 196596664824XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about this node. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow).
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetNodeResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNode) SetCreateTime(v int64) *GetNodeResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetId(v string) *GetNodeResponseBodyNode {
	s.Id = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifyTime(v int64) *GetNodeResponseBodyNode {
	s.ModifyTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetName(v string) *GetNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetOwner(v string) *GetNodeResponseBodyNode {
	s.Owner = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetProjectId(v string) *GetNodeResponseBodyNode {
	s.ProjectId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetSpec(v string) *GetNodeResponseBodyNode {
	s.Spec = &v
	return s
}

type GetNodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeResponse) SetHeaders(v map[string]*string) *GetNodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeResponse) SetStatusCode(v int32) *GetNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeResponse) SetBody(v *GetNodeResponseBody) *GetNodeResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetId(v int64) *GetProjectRequest {
	s.Id = &v
	return s
}

type GetProjectResponseBody struct {
	Project *GetProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectResponseBodyProject struct {
	// example:
	//
	// rg-acfmzbn7pti3zfa
	AliyunResourceGroupId *string                                            `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*GetProjectResponseBodyProjectAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	Description           *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 28477242
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 207947397706614299
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProject) SetAliyunResourceGroupId(v string) *GetProjectResponseBodyProject {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetAliyunResourceTags(v []*GetProjectResponseBodyProjectAliyunResourceTags) *GetProjectResponseBodyProject {
	s.AliyunResourceTags = v
	return s
}

func (s *GetProjectResponseBodyProject) SetDescription(v string) *GetProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDevEnvironmentEnabled(v bool) *GetProjectResponseBodyProject {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDevRoleDisabled(v bool) *GetProjectResponseBodyProject {
	s.DevRoleDisabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDisplayName(v string) *GetProjectResponseBodyProject {
	s.DisplayName = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetId(v int64) *GetProjectResponseBodyProject {
	s.Id = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetName(v string) *GetProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetOwner(v string) *GetProjectResponseBodyProject {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetPaiTaskEnabled(v bool) *GetProjectResponseBodyProject {
	s.PaiTaskEnabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetStatus(v string) *GetProjectResponseBodyProject {
	s.Status = &v
	return s
}

type GetProjectResponseBodyProjectAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProjectResponseBodyProjectAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyProjectAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) SetKey(v string) *GetProjectResponseBodyProjectAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) SetValue(v string) *GetProjectResponseBodyProjectAliyunResourceTags {
	s.Value = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetProjectMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 88757
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetProjectMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMemberRequest) SetProjectId(v int64) *GetProjectMemberRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectMemberRequest) SetUserId(v string) *GetProjectMemberRequest {
	s.UserId = &v
	return s
}

type GetProjectMemberResponseBody struct {
	ProjectMember *GetProjectMemberResponseBodyProjectMember `json:"ProjectMember,omitempty" xml:"ProjectMember,omitempty" type:"Struct"`
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBody) SetProjectMember(v *GetProjectMemberResponseBodyProjectMember) *GetProjectMemberResponseBody {
	s.ProjectMember = v
	return s
}

func (s *GetProjectMemberResponseBody) SetRequestId(v string) *GetProjectMemberResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectMemberResponseBodyProjectMember struct {
	// example:
	//
	// 88757
	ProjectId *int64                                            `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Roles     []*GetProjectMemberResponseBodyProjectMemberRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetProjectMemberResponseBodyProjectMember) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBodyProjectMember) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyProjectMember) SetProjectId(v int64) *GetProjectMemberResponseBodyProjectMember {
	s.ProjectId = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetRoles(v []*GetProjectMemberResponseBodyProjectMemberRoles) *GetProjectMemberResponseBodyProjectMember {
	s.Roles = v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetStatus(v string) *GetProjectMemberResponseBodyProjectMember {
	s.Status = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMember) SetUserId(v string) *GetProjectMemberResponseBodyProjectMember {
	s.UserId = &v
	return s
}

type GetProjectMemberResponseBodyProjectMemberRoles struct {
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProjectMemberResponseBodyProjectMemberRoles) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponseBodyProjectMemberRoles) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) SetCode(v string) *GetProjectMemberResponseBodyProjectMemberRoles {
	s.Code = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) SetName(v string) *GetProjectMemberResponseBodyProjectMemberRoles {
	s.Name = &v
	return s
}

func (s *GetProjectMemberResponseBodyProjectMemberRoles) SetType(v string) *GetProjectMemberResponseBodyProjectMemberRoles {
	s.Type = &v
	return s
}

type GetProjectMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *GetProjectMemberResponse) SetHeaders(v map[string]*string) *GetProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *GetProjectMemberResponse) SetStatusCode(v int32) *GetProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectMemberResponse) SetBody(v *GetProjectMemberResponseBody) *GetProjectMemberResponse {
	s.Body = v
	return s
}

type GetProjectRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRoleRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRoleRequest) SetCode(v string) *GetProjectRoleRequest {
	s.Code = &v
	return s
}

func (s *GetProjectRoleRequest) SetProjectId(v int64) *GetProjectRoleRequest {
	s.ProjectId = &v
	return s
}

type GetProjectRoleResponseBody struct {
	ProjectRole *GetProjectRoleResponseBodyProjectRole `json:"ProjectRole,omitempty" xml:"ProjectRole,omitempty" type:"Struct"`
	// example:
	//
	// 82F28E60-CF48-5EDF-AB25-D806847B97D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponseBody) SetProjectRole(v *GetProjectRoleResponseBodyProjectRole) *GetProjectRoleResponseBody {
	s.ProjectRole = v
	return s
}

func (s *GetProjectRoleResponseBody) SetRequestId(v string) *GetProjectRoleResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectRoleResponseBodyProjectRole struct {
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProjectRoleResponseBodyProjectRole) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRoleResponseBodyProjectRole) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponseBodyProjectRole) SetCode(v string) *GetProjectRoleResponseBodyProjectRole {
	s.Code = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetName(v string) *GetProjectRoleResponseBodyProjectRole {
	s.Name = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetProjectId(v int64) *GetProjectRoleResponseBodyProjectRole {
	s.ProjectId = &v
	return s
}

func (s *GetProjectRoleResponseBodyProjectRole) SetType(v string) *GetProjectRoleResponseBodyProjectRole {
	s.Type = &v
	return s
}

type GetProjectRoleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRoleResponse) GoString() string {
	return s.String()
}

func (s *GetProjectRoleResponse) SetHeaders(v map[string]*string) *GetProjectRoleResponse {
	s.Headers = v
	return s
}

func (s *GetProjectRoleResponse) SetStatusCode(v int32) *GetProjectRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectRoleResponse) SetBody(v *GetProjectRoleResponseBody) *GetProjectRoleResponse {
	s.Body = v
	return s
}

type GetResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) SetId(v string) *GetResourceRequest {
	s.Id = &v
	return s
}

func (s *GetResourceRequest) SetProjectId(v string) *GetResourceRequest {
	s.ProjectId = &v
	return s
}

type GetResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E871F6C0-2EFF-5790-A00D-C57543EEXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the file resource.
	Resource *GetResourceResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetResource(v *GetResourceResponseBodyResource) *GetResourceResponseBody {
	s.Resource = v
	return s
}

type GetResourceResponseBodyResource struct {
	// The time when the file resource was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the file resource.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the file resource was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1700539206000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the file resource.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the file resource.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the file resource belongs.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the file resource. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow).
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetResourceResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyResource) SetCreateTime(v int64) *GetResourceResponseBodyResource {
	s.CreateTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetId(v string) *GetResourceResponseBodyResource {
	s.Id = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetModifyTime(v int64) *GetResourceResponseBodyResource {
	s.ModifyTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetName(v string) *GetResourceResponseBodyResource {
	s.Name = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetOwner(v string) *GetResourceResponseBodyResource {
	s.Owner = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetProjectId(v string) *GetResourceResponseBodyResource {
	s.ProjectId = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetSpec(v string) *GetResourceResponseBodyResource {
	s.Spec = &v
	return s
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type GetResourceGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) SetId(v string) *GetResourceGroupRequest {
	s.Id = &v
	return s
}

type GetResourceGroupResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroup *GetResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBody) SetRequestId(v string) *GetResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetResourceGroup(v *GetResourceGroupResponseBodyResourceGroup) *GetResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *GetResourceGroupResponseBody) SetSuccess(v bool) *GetResourceGroupResponseBody {
	s.Success = &v
	return s
}

type GetResourceGroupResponseBodyResourceGroup struct {
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	DefaultVpcId *string `json:"DefaultVpcId,omitempty" xml:"DefaultVpcId,omitempty"`
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	DefaultVswitchId *string `json:"DefaultVswitchId,omitempty" xml:"DefaultVswitchId,omitempty"`
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// c442b330-3b10-4584-959e-736e4edXXXXX
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 创建用于普通任务的通用资源组
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// CommonV2
	ResourceGroupType *string                                        `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	Spec              *GetResourceGroupResponseBodyResourceGroupSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetCreateTime(v int64) *GetResourceGroupResponseBodyResourceGroup {
	s.CreateTime = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetCreateUser(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.CreateUser = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetDefaultVpcId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.DefaultVpcId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetDefaultVswitchId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.DefaultVswitchId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetName(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetOrderInstanceId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.OrderInstanceId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetPaymentType(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.PaymentType = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetRemark(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Remark = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetResourceGroupType(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.ResourceGroupType = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetSpec(v *GetResourceGroupResponseBodyResourceGroupSpec) *GetResourceGroupResponseBodyResourceGroup {
	s.Spec = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetStatus(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

type GetResourceGroupResponseBodyResourceGroupSpec struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2CU
	Standard *string `json:"Standard,omitempty" xml:"Standard,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupSpec) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupSpec) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) SetAmount(v int32) *GetResourceGroupResponseBodyResourceGroupSpec {
	s.Amount = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) SetStandard(v string) *GetResourceGroupResponseBodyResourceGroupSpec {
	s.Standard = &v
	return s
}

type GetResourceGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponse) SetHeaders(v map[string]*string) *GetResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *GetResourceGroupResponse) SetStatusCode(v int32) *GetResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceGroupResponse) SetBody(v *GetResourceGroupResponseBody) *GetResourceGroupResponse {
	s.Body = v
	return s
}

type GetRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRouteRequest) GoString() string {
	return s.String()
}

func (s *GetRouteRequest) SetId(v int64) *GetRouteRequest {
	s.Id = &v
	return s
}

type GetRouteResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Route     *GetRouteResponseBodyRoute `json:"Route,omitempty" xml:"Route,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetRouteResponseBody) SetRequestId(v string) *GetRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRouteResponseBody) SetRoute(v *GetRouteResponseBodyRoute) *GetRouteResponseBody {
	s.Route = v
	return s
}

func (s *GetRouteResponseBody) SetSuccess(v bool) *GetRouteResponseBody {
	s.Success = &v
	return s
}

type GetRouteResponseBodyRoute struct {
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ns-679XXXXX
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetRouteResponseBodyRoute) String() string {
	return tea.Prettify(s)
}

func (s GetRouteResponseBodyRoute) GoString() string {
	return s.String()
}

func (s *GetRouteResponseBodyRoute) SetCreateTime(v int64) *GetRouteResponseBodyRoute {
	s.CreateTime = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetDestinationCidr(v string) *GetRouteResponseBodyRoute {
	s.DestinationCidr = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetId(v int64) *GetRouteResponseBodyRoute {
	s.Id = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetNetworkId(v int64) *GetRouteResponseBodyRoute {
	s.NetworkId = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetResourceGroupId(v string) *GetRouteResponseBodyRoute {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRouteResponseBodyRoute) SetResourceId(v string) *GetRouteResponseBodyRoute {
	s.ResourceId = &v
	return s
}

type GetRouteResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRouteResponse) GoString() string {
	return s.String()
}

func (s *GetRouteResponse) SetHeaders(v map[string]*string) *GetRouteResponse {
	s.Headers = v
	return s
}

func (s *GetRouteResponse) SetStatusCode(v int32) *GetRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRouteResponse) SetBody(v *GetRouteResponseBody) *GetRouteResponse {
	s.Body = v
	return s
}

type GetWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionRequest) SetId(v string) *GetWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionRequest) SetProjectId(v string) *GetWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type GetWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F2BDD628-8A21-5BD1-B930-1A2D5989XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the workflow.
	WorkflowDefinition *GetWorkflowDefinitionResponseBodyWorkflowDefinition `json:"WorkflowDefinition,omitempty" xml:"WorkflowDefinition,omitempty" type:"Struct"`
}

func (s GetWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBody) SetRequestId(v string) *GetWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBody) SetWorkflowDefinition(v *GetWorkflowDefinitionResponseBodyWorkflowDefinition) *GetWorkflowDefinitionResponseBody {
	s.WorkflowDefinition = v
	return s
}

type GetWorkflowDefinitionResponseBodyWorkflowDefinition struct {
	// The time when the workflow was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1708481905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the workflow was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1708481905000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the workflow.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the workflow.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the workflow belongs.
	//
	// example:
	//
	// 307XXX
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the workflow. For more information, see [FlowSpec](https://github.com/aliyun/alibabacloud-dataworks-tool-dflow/).
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetCreateTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetModifyTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetName(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Name = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetOwner(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Owner = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetProjectId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetSpec(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Spec = &v
	return s
}

type GetWorkflowDefinitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *GetWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetStatusCode(v int32) *GetWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetBody(v *GetWorkflowDefinitionResponseBody) *GetWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type GrantMemberProjectRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantMemberProjectRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantMemberProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesRequest) SetProjectId(v int64) *GrantMemberProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *GrantMemberProjectRolesRequest) SetRoleCodes(v []*string) *GrantMemberProjectRolesRequest {
	s.RoleCodes = v
	return s
}

func (s *GrantMemberProjectRolesRequest) SetUserId(v string) *GrantMemberProjectRolesRequest {
	s.UserId = &v
	return s
}

type GrantMemberProjectRolesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantMemberProjectRolesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantMemberProjectRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesShrinkRequest) SetProjectId(v int64) *GrantMemberProjectRolesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GrantMemberProjectRolesShrinkRequest) SetRoleCodesShrink(v string) *GrantMemberProjectRolesShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *GrantMemberProjectRolesShrinkRequest) SetUserId(v string) *GrantMemberProjectRolesShrinkRequest {
	s.UserId = &v
	return s
}

type GrantMemberProjectRolesResponseBody struct {
	// example:
	//
	// 2d9ced66-38ef-4923-baf6-391dd3a7e656
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantMemberProjectRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantMemberProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesResponseBody) SetRequestId(v string) *GrantMemberProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

type GrantMemberProjectRolesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantMemberProjectRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantMemberProjectRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantMemberProjectRolesResponse) GoString() string {
	return s.String()
}

func (s *GrantMemberProjectRolesResponse) SetHeaders(v map[string]*string) *GrantMemberProjectRolesResponse {
	s.Headers = v
	return s
}

func (s *GrantMemberProjectRolesResponse) SetStatusCode(v int32) *GrantMemberProjectRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantMemberProjectRolesResponse) SetBody(v *GrantMemberProjectRolesResponseBody) *GrantMemberProjectRolesResponse {
	s.Body = v
	return s
}

type ImportWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ImportWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionRequest) SetProjectId(v string) *ImportWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *ImportWorkflowDefinitionRequest) SetSpec(v string) *ImportWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type ImportWorkflowDefinitionResponseBody struct {
	AsyncJob *ImportWorkflowDefinitionResponseBodyAsyncJob `json:"AsyncJob,omitempty" xml:"AsyncJob,omitempty" type:"Struct"`
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF020E7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionResponseBody) SetAsyncJob(v *ImportWorkflowDefinitionResponseBodyAsyncJob) *ImportWorkflowDefinitionResponseBody {
	s.AsyncJob = v
	return s
}

func (s *ImportWorkflowDefinitionResponseBody) SetRequestId(v string) *ImportWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

type ImportWorkflowDefinitionResponseBodyAsyncJob struct {
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 1706581425000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// target folder already exists: XXXX
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 1234567691239009XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 632647691239009XXXX
	Response *string `json:"Response,omitempty" xml:"Response,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Create
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ImportWorkflowDefinitionResponseBodyAsyncJob) String() string {
	return tea.Prettify(s)
}

func (s ImportWorkflowDefinitionResponseBodyAsyncJob) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetCompleted(v bool) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Completed = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetCreateTime(v int64) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.CreateTime = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetError(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Error = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetId(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Id = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetProgress(v int32) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Progress = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetResponse(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Response = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetStatus(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Status = &v
	return s
}

func (s *ImportWorkflowDefinitionResponseBodyAsyncJob) SetType(v string) *ImportWorkflowDefinitionResponseBodyAsyncJob {
	s.Type = &v
	return s
}

type ImportWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *ImportWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *ImportWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *ImportWorkflowDefinitionResponse) SetStatusCode(v int32) *ImportWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportWorkflowDefinitionResponse) SetBody(v *ImportWorkflowDefinitionResponseBody) *ImportWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type ListDIAlarmRulesRequest struct {
	// example:
	//
	// 34988
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// example:
	//
	// 1000001
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDIAlarmRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesRequest) SetDIAlarmRuleId(v int64) *ListDIAlarmRulesRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetJobId(v int64) *ListDIAlarmRulesRequest {
	s.JobId = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetPageNumber(v int32) *ListDIAlarmRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIAlarmRulesRequest) SetPageSize(v int32) *ListDIAlarmRulesRequest {
	s.PageSize = &v
	return s
}

type ListDIAlarmRulesResponseBody struct {
	PagingInfo *ListDIAlarmRulesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 74C2FECD-5B3A-554A-BCF5-351A36DE9815
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIAlarmRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBody) SetPagingInfo(v *ListDIAlarmRulesResponseBodyPagingInfo) *ListDIAlarmRulesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIAlarmRulesResponseBody) SetRequestId(v string) *ListDIAlarmRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListDIAlarmRulesResponseBodyPagingInfo struct {
	DIJobAlarmRules []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules `json:"DIJobAlarmRules,omitempty" xml:"DIJobAlarmRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 90
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetDIJobAlarmRules(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.DIJobAlarmRules = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetPageSize(v int64) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules struct {
	// example:
	//
	// 72402
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// example:
	//
	// 32594
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// rule descrition
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// True
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// rule_name
	Name                 *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationSettings *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	TriggerConditions    []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions  `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetDIAlarmRuleId(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.DIAlarmRuleId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetDIJobId(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.DIJobId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetDescription(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Description = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetEnabled(v bool) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Enabled = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetMetricType(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.MetricType = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetName(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Name = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetNotificationSettings(v *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.NotificationSettings = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetTriggerConditions(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.TriggerConditions = v
	return s
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings struct {
	// example:
	//
	// 5
	InhibitionInterval    *int64                                                                                            `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	NotificationChannels  []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	NotificationReceivers []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetInhibitionInterval(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetNotificationChannels(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetNotificationReceivers(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.NotificationReceivers = v
	return s
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels struct {
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// example:
	//
	// Critical
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) SetChannels(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) SetSeverity(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers struct {
	// example:
	//
	// DingToken
	ReceiverType   *string   `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) SetReceiverType(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions struct {
	DdlReportTags []*string `json:"DdlReportTags,omitempty" xml:"DdlReportTags,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// Critical
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetDdlReportTags(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.DdlReportTags = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetDuration(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.Duration = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetSeverity(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.Severity = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetThreshold(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.Threshold = &v
	return s
}

type ListDIAlarmRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIAlarmRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIAlarmRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIAlarmRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponse) SetHeaders(v map[string]*string) *ListDIAlarmRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDIAlarmRulesResponse) SetStatusCode(v int32) *ListDIAlarmRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIAlarmRulesResponse) SetBody(v *ListDIAlarmRulesResponseBody) *ListDIAlarmRulesResponse {
	s.Body = v
	return s
}

type ListDIJobEventsRequest struct {
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1717971005
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Alarm
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1716971005
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsRequest) SetDIJobId(v int64) *ListDIJobEventsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobEventsRequest) SetEndTime(v int64) *ListDIJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobEventsRequest) SetEventType(v string) *ListDIJobEventsRequest {
	s.EventType = &v
	return s
}

func (s *ListDIJobEventsRequest) SetPageNumber(v int64) *ListDIJobEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobEventsRequest) SetPageSize(v int64) *ListDIJobEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobEventsRequest) SetStartTime(v int64) *ListDIJobEventsRequest {
	s.StartTime = &v
	return s
}

type ListDIJobEventsResponseBody struct {
	PagingInfo *ListDIJobEventsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 645F6D68-9C29-5961-80B1-BDD4B794C22D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBody) SetPagingInfo(v *ListDIJobEventsResponseBodyPagingInfo) *ListDIJobEventsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobEventsResponseBody) SetRequestId(v string) *ListDIJobEventsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobEventsResponseBodyPagingInfo struct {
	DIJobEvent []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent `json:"DIJobEvent,omitempty" xml:"DIJobEvent,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2524
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobEventsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetDIJobEvent(v []*ListDIJobEventsResponseBodyPagingInfoDIJobEvent) *ListDIJobEventsResponseBodyPagingInfo {
	s.DIJobEvent = v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetPageSize(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIJobEventsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDIJobEventsResponseBodyPagingInfoDIJobEvent struct {
	// example:
	//
	// Ignore
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// Phone
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// example:
	//
	// 1663573162
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Detail     *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// alter table table2 ***
	DstSql *string `json:"DstSql,omitempty" xml:"DstSql,omitempty"`
	// example:
	//
	// table2
	DstTable *string `json:"DstTable,omitempty" xml:"DstTable,omitempty"`
	// example:
	//
	// 2024-05-29 15:11:31,377 [main] INFO com.*.**.di.core.metrics.:21 []  {****}
	//
	// 2024-05-29 15:11:31,384 [main] INFO *.aliyun.*.di.*.*.metrics.*:27 [] - Open MarioDiReporter
	//
	// 2024-05-29 15:11:33,248 [flink-akka.*.*-dispatcher-17] INFO
	FailoverMessage *string `json:"FailoverMessage,omitempty" xml:"FailoverMessage,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// alter table table1 ***
	SrcSql *string `json:"SrcSql,omitempty" xml:"SrcSql,omitempty"`
	// example:
	//
	// table1
	SrcTable *string `json:"SrcTable,omitempty" xml:"SrcTable,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Delay
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDIJobEventsResponseBodyPagingInfoDIJobEvent) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponseBodyPagingInfoDIJobEvent) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetAction(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Action = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetChannels(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Channels = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetCreateTime(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.CreateTime = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDetail(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Detail = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDstSql(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.DstSql = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetDstTable(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.DstTable = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetFailoverMessage(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.FailoverMessage = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetId(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Id = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSeverity(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Severity = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSrcSql(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.SrcSql = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetSrcTable(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.SrcTable = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetStatus(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Status = &v
	return s
}

func (s *ListDIJobEventsResponseBodyPagingInfoDIJobEvent) SetType(v string) *ListDIJobEventsResponseBodyPagingInfoDIJobEvent {
	s.Type = &v
	return s
}

type ListDIJobEventsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobEventsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsResponse) SetHeaders(v map[string]*string) *ListDIJobEventsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobEventsResponse) SetStatusCode(v int32) *ListDIJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobEventsResponse) SetBody(v *ListDIJobEventsResponseBody) *ListDIJobEventsResponse {
	s.Body = v
	return s
}

type ListDIJobMetricsRequest struct {
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1712205941
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	MetricName []*string `json:"MetricName,omitempty" xml:"MetricName,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1586509407
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsRequest) SetDIJobId(v int64) *ListDIJobMetricsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobMetricsRequest) SetEndTime(v int64) *ListDIJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobMetricsRequest) SetMetricName(v []*string) *ListDIJobMetricsRequest {
	s.MetricName = v
	return s
}

func (s *ListDIJobMetricsRequest) SetStartTime(v int64) *ListDIJobMetricsRequest {
	s.StartTime = &v
	return s
}

type ListDIJobMetricsShrinkRequest struct {
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1712205941
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	MetricNameShrink *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1586509407
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobMetricsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsShrinkRequest) SetDIJobId(v int64) *ListDIJobMetricsShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetEndTime(v int64) *ListDIJobMetricsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetMetricNameShrink(v string) *ListDIJobMetricsShrinkRequest {
	s.MetricNameShrink = &v
	return s
}

func (s *ListDIJobMetricsShrinkRequest) SetStartTime(v int64) *ListDIJobMetricsShrinkRequest {
	s.StartTime = &v
	return s
}

type ListDIJobMetricsResponseBody struct {
	PagingInfo *ListDIJobMetricsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBody) SetPagingInfo(v *ListDIJobMetricsResponseBodyPagingInfo) *ListDIJobMetricsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobMetricsResponseBody) SetRequestId(v string) *ListDIJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobMetricsResponseBodyPagingInfo struct {
	JobMetrics []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics `json:"JobMetrics,omitempty" xml:"JobMetrics,omitempty" type:"Repeated"`
}

func (s ListDIJobMetricsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfo) SetJobMetrics(v []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics) *ListDIJobMetricsResponseBodyPagingInfo {
	s.JobMetrics = v
	return s
}

type ListDIJobMetricsResponseBodyPagingInfoJobMetrics struct {
	// example:
	//
	// JobDelay
	Name       *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	SeriesList []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList `json:"SeriesList,omitempty" xml:"SeriesList,omitempty" type:"Repeated"`
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetrics) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) SetName(v string) *ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	s.Name = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) SetSeriesList(v []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) *ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	s.SeriesList = v
	return s
}

type ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList struct {
	// example:
	//
	// 1716881141
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 10
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) SetTime(v int64) *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	s.Time = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) SetValue(v float64) *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	s.Value = &v
	return s
}

type ListDIJobMetricsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponse) SetHeaders(v map[string]*string) *ListDIJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobMetricsResponse) SetStatusCode(v int32) *ListDIJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobMetricsResponse) SetBody(v *ListDIJobMetricsResponseBody) *ListDIJobMetricsResponse {
	s.Body = v
	return s
}

type ListDIJobRunDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ds_name
	SourceDataSourceName *string `json:"SourceDataSourceName,omitempty" xml:"SourceDataSourceName,omitempty"`
	// example:
	//
	// db_name
	SourceDatabaseName *string `json:"SourceDatabaseName,omitempty" xml:"SourceDatabaseName,omitempty"`
	// example:
	//
	// schema_name
	SourceSchemaName *string `json:"SourceSchemaName,omitempty" xml:"SourceSchemaName,omitempty"`
	// example:
	//
	// table_name
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
}

func (s ListDIJobRunDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobRunDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsRequest) SetDIJobId(v int64) *ListDIJobRunDetailsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetInstanceId(v int64) *ListDIJobRunDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetPageNumber(v int64) *ListDIJobRunDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetPageSize(v int64) *ListDIJobRunDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceDataSourceName(v string) *ListDIJobRunDetailsRequest {
	s.SourceDataSourceName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceDatabaseName(v string) *ListDIJobRunDetailsRequest {
	s.SourceDatabaseName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceSchemaName(v string) *ListDIJobRunDetailsRequest {
	s.SourceSchemaName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceTableName(v string) *ListDIJobRunDetailsRequest {
	s.SourceTableName = &v
	return s
}

type ListDIJobRunDetailsResponseBody struct {
	PagingInfo *ListDIJobRunDetailsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobRunDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobRunDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponseBody) SetPagingInfo(v *ListDIJobRunDetailsResponseBodyPagingInfo) *ListDIJobRunDetailsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobRunDetailsResponseBody) SetRequestId(v string) *ListDIJobRunDetailsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobRunDetailsResponseBodyPagingInfo struct {
	JobRunInfos []*ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos `json:"JobRunInfos,omitempty" xml:"JobRunInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 131
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobRunDetailsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobRunDetailsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetJobRunInfos(v []*ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.JobRunInfos = v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetPageNumber(v string) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetPageSize(v string) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetTotalCount(v string) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos struct {
	// example:
	//
	// dst_db
	DestinationDatabaseName *string `json:"DestinationDatabaseName,omitempty" xml:"DestinationDatabaseName,omitempty"`
	// example:
	//
	// dst_name
	DestinationDatasourceName *string `json:"DestinationDatasourceName,omitempty" xml:"DestinationDatasourceName,omitempty"`
	// example:
	//
	// dst_schema
	DestinationSchemaName *string `json:"DestinationSchemaName,omitempty" xml:"DestinationSchemaName,omitempty"`
	// example:
	//
	// dst_name
	DestinationTableName *string `json:"DestinationTableName,omitempty" xml:"DestinationTableName,omitempty"`
	// example:
	//
	// sync table t1 fail.
	FullMigrationErrorMessage *string `json:"FullMigrationErrorMessage,omitempty" xml:"FullMigrationErrorMessage,omitempty"`
	// example:
	//
	// Finished
	FullMigrationStatus *string `json:"FullMigrationStatus,omitempty" xml:"FullMigrationStatus,omitempty"`
	// example:
	//
	// 0
	OfflineErrorRecords *int64 `json:"OfflineErrorRecords,omitempty" xml:"OfflineErrorRecords,omitempty"`
	// example:
	//
	// 100
	OfflineTotalBytes *int64 `json:"OfflineTotalBytes,omitempty" xml:"OfflineTotalBytes,omitempty"`
	// example:
	//
	// 10
	OfflineTotalRecords *int64 `json:"OfflineTotalRecords,omitempty" xml:"OfflineTotalRecords,omitempty"`
	// example:
	//
	// sync table t1 fail.
	RealtimeMigrationErrorMessage *string `json:"RealtimeMigrationErrorMessage,omitempty" xml:"RealtimeMigrationErrorMessage,omitempty"`
	// example:
	//
	// Running
	RealtimeMigrationStatus *string `json:"RealtimeMigrationStatus,omitempty" xml:"RealtimeMigrationStatus,omitempty"`
	// example:
	//
	// db_name
	SourceDatabaseName *string `json:"SourceDatabaseName,omitempty" xml:"SourceDatabaseName,omitempty"`
	// example:
	//
	// ds_name
	SourceDatasourceName *string `json:"SourceDatasourceName,omitempty" xml:"SourceDatasourceName,omitempty"`
	// example:
	//
	// schema_name
	SourceSchemaName *string `json:"SourceSchemaName,omitempty" xml:"SourceSchemaName,omitempty"`
	// example:
	//
	// table_name
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// example:
	//
	// create table t1 fail.
	StructureMigrationErrorMessage *string `json:"StructureMigrationErrorMessage,omitempty" xml:"StructureMigrationErrorMessage,omitempty"`
	// example:
	//
	// Finished
	StructureMigrationStatus *string `json:"StructureMigrationStatus,omitempty" xml:"StructureMigrationStatus,omitempty"`
}

func (s ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationDatabaseName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationDatabaseName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationDatasourceName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationDatasourceName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationSchemaName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationSchemaName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationTableName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationTableName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetFullMigrationErrorMessage(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.FullMigrationErrorMessage = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetFullMigrationStatus(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.FullMigrationStatus = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetOfflineErrorRecords(v int64) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.OfflineErrorRecords = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetOfflineTotalBytes(v int64) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.OfflineTotalBytes = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetOfflineTotalRecords(v int64) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.OfflineTotalRecords = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetRealtimeMigrationErrorMessage(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.RealtimeMigrationErrorMessage = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetRealtimeMigrationStatus(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.RealtimeMigrationStatus = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceDatabaseName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceDatabaseName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceDatasourceName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceDatasourceName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceSchemaName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceSchemaName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceTableName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceTableName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetStructureMigrationErrorMessage(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.StructureMigrationErrorMessage = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetStructureMigrationStatus(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.StructureMigrationStatus = &v
	return s
}

type ListDIJobRunDetailsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobRunDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobRunDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobRunDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponse) SetHeaders(v map[string]*string) *ListDIJobRunDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobRunDetailsResponse) SetStatusCode(v int32) *ListDIJobRunDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobRunDetailsResponse) SetBody(v *ListDIJobRunDetailsResponseBody) *ListDIJobRunDetailsResponse {
	s.Body = v
	return s
}

type ListDIJobsRequest struct {
	// The destination type. If you do not configure this parameter, no limits are imposed on the tasks.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental: one-time full synchronization and real-time incremental synchronization
	//
	// 	- RealtimeIncremental: real-time incremental synchronization
	//
	// 	- Full: full synchronization
	//
	// 	- OfflineIncremental: batch incremental synchronization
	//
	// 	- FullAndOfflineIncremental: one-time full synchronization and batch incremental synchronization
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The name of the export task.
	//
	// The name of each export task must be unique. You must make sure that the names of the export tasks in the current workspace are unique.
	//
	// example:
	//
	// test_export_01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1967
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The source type. If you do not configure this parameter, no limits are imposed on the tasks.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobsRequest) SetDestinationDataSourceType(v string) *ListDIJobsRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsRequest) SetMigrationType(v string) *ListDIJobsRequest {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsRequest) SetName(v string) *ListDIJobsRequest {
	s.Name = &v
	return s
}

func (s *ListDIJobsRequest) SetPageNumber(v int64) *ListDIJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsRequest) SetPageSize(v int64) *ListDIJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsRequest) SetProjectId(v int64) *ListDIJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsRequest) SetSourceDataSourceType(v string) *ListDIJobsRequest {
	s.SourceDataSourceType = &v
	return s
}

type ListDIJobsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDIJobsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7263E4AC-9D2E-5B29-B8AF-7C5012E92A41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBody) SetPagingInfo(v *ListDIJobsResponseBodyPagingInfo) *ListDIJobsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobsResponseBody) SetRequestId(v string) *ListDIJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListDIJobsResponseBodyPagingInfo struct {
	// The synchronization tasks returned.
	DIJobs []*ListDIJobsResponseBodyPagingInfoDIJobs `json:"DIJobs,omitempty" xml:"DIJobs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyPagingInfo) SetDIJobs(v []*ListDIJobsResponseBodyPagingInfoDIJobs) *ListDIJobsResponseBodyPagingInfo {
	s.DIJobs = v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetPageSize(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDIJobsResponseBodyPagingInfoDIJobs struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 32599
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The destination type. Valid values: Hologres and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_35197
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The status of the synchronization task. Valid values:
	//
	// 	- Finished
	//
	// 	- Initialized
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental: one-time full synchronization and real-time incremental synchronization
	//
	// 	- RealtimeIncremental: real-time incremental synchronization
	//
	// 	- Full: full synchronization
	//
	// 	- OfflineIncremental: batch incremental synchronization
	//
	// 	- FullAndOfflineIncremental: one-time full synchronization and batch incremental synchronization
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The ID of the DataWorks workspace to which the synchronization task belongs.
	//
	// example:
	//
	// 26442
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The source type. The value MySQL is returned.
	//
	// example:
	//
	// Mysql
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsResponseBodyPagingInfoDIJobs) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponseBodyPagingInfoDIJobs) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetDIJobId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetDestinationDataSourceType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetJobName(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.JobName = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetJobStatus(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.JobStatus = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetMigrationType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetProjectId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetSourceDataSourceType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.SourceDataSourceType = &v
	return s
}

type ListDIJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDIJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDIJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDIJobsResponse) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponse) SetHeaders(v map[string]*string) *ListDIJobsResponse {
	s.Headers = v
	return s
}

func (s *ListDIJobsResponse) SetStatusCode(v int32) *ListDIJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDIJobsResponse) SetBody(v *ListDIJobsResponseBody) *ListDIJobsResponse {
	s.Body = v
	return s
}

type ListDataQualityEvaluationTaskInstancesRequest struct {
	// example:
	//
	// 2024-04-01
	BizdateFrom *string `json:"BizdateFrom,omitempty" xml:"BizdateFrom,omitempty"`
	// example:
	//
	// 2024-05-01
	BizdateTo *string `json:"BizdateTo,omitempty" xml:"BizdateTo,omitempty"`
	// example:
	//
	// 1710239005403
	CreateTimeFrom *int64 `json:"CreateTimeFrom,omitempty" xml:"CreateTimeFrom,omitempty"`
	// example:
	//
	// 1710239005403
	CreateTimeTo *int64 `json:"CreateTimeTo,omitempty" xml:"CreateTimeTo,omitempty"`
	// example:
	//
	// 10000
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// example:
	//
	// CWF2
	TriggerClient *string `json:"TriggerClient,omitempty" xml:"TriggerClient,omitempty"`
	// example:
	//
	// 1001
	TriggerClientId *string `json:"TriggerClientId,omitempty" xml:"TriggerClientId,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetBizdateFrom(v string) *ListDataQualityEvaluationTaskInstancesRequest {
	s.BizdateFrom = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetBizdateTo(v string) *ListDataQualityEvaluationTaskInstancesRequest {
	s.BizdateTo = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetCreateTimeFrom(v int64) *ListDataQualityEvaluationTaskInstancesRequest {
	s.CreateTimeFrom = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetCreateTimeTo(v int64) *ListDataQualityEvaluationTaskInstancesRequest {
	s.CreateTimeTo = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetDataQualityEvaluationTaskId(v int64) *ListDataQualityEvaluationTaskInstancesRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetPageNumber(v int32) *ListDataQualityEvaluationTaskInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetPageSize(v int32) *ListDataQualityEvaluationTaskInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetProjectId(v int64) *ListDataQualityEvaluationTaskInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetTableGuid(v string) *ListDataQualityEvaluationTaskInstancesRequest {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetTriggerClient(v string) *ListDataQualityEvaluationTaskInstancesRequest {
	s.TriggerClient = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesRequest) SetTriggerClientId(v string) *ListDataQualityEvaluationTaskInstancesRequest {
	s.TriggerClientId = &v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBody struct {
	// example:
	//
	// 200
	Code       *string                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PagingInfo *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) SetCode(v string) *ListDataQualityEvaluationTaskInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) SetMessage(v string) *ListDataQualityEvaluationTaskInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) SetPagingInfo(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) *ListDataQualityEvaluationTaskInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBody) SetRequestId(v string) *ListDataQualityEvaluationTaskInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo struct {
	DataQualityEvaluationTaskInstances []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances `json:"DataQualityEvaluationTaskInstances,omitempty" xml:"DataQualityEvaluationTaskInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
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

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances struct {
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
	// {
	//
	//   "bizdate": "20240517",
	//
	//   "triggerTime": "1710239005403"
	//
	// }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Critical
	Status *string                                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Task   *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
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
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstances) GoString() string {
	return s.String()
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

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask struct {
	// 质量监控任务描述
	//
	// example:
	//
	// This is a daily run data quality evaluation plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据质量校验任务实例生命周期中的回调设置，目前只支持一个阻塞调度任务的Hook
	Hooks []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 质量监控任务名称
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据质量校验任务通知订阅配置
	Notifications *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// 项目空间Id
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 代表region的资源属性字段
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 使用数据源时的一些设置，目前只支持指定EMR的yarn队列、采集EMR表时把SQL引擎指定为SPARK-SQL
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK-SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// 参看 DataQualityTarget示例	数据质量校验任务的监控对象，参考 DataQualityTarget
	Target *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// 租户Id
	//
	// example:
	//
	// 10
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 数据质量校验任务的触发配置
	Trigger *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) GoString() string {
	return s.String()
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

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetRegionId(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.RegionId = &v
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

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetTenantId(v int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.TenantId = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask) SetTrigger(v *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTask {
	s.Trigger = v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks struct {
	// Hook触发条件
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// Hook类型
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) SetCondition(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks) SetType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskHooks {
	s.Type = &v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications struct {
	// 通知触发条件
	//
	// example:
	//
	// ${severity} == "High"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// 具体的消息通知设置
	Notifications []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) SetCondition(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications) SetNotifications(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotifications {
	s.Notifications = v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications struct {
	// 告警接收人设置
	NofiticationReceivers []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers `json:"NofiticationReceivers,omitempty" xml:"NofiticationReceivers,omitempty" type:"Repeated"`
	// 通知方式
	NotificationChannels []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) SetNofiticationReceivers(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications {
	s.NofiticationReceivers = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications) SetNotificationChannels(v []*ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers struct {
	// 扩展信息，格式为 json，例如钉钉机器人支持 at 所有人
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// 告警接收人类型
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// 告警接收人
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNofiticationReceivers) GoString() string {
	return s.String()
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

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels struct {
	// 通知方式
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget struct {
	// 表所属的数据库类型
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// 分区表的分区设置
	//
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// 表在数据地图中的唯一ID
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// 监控对象类型
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTarget) GoString() string {
	return s.String()
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

type ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger struct {
	// 具体指明哪些调度节点的实例执行成功后可以触发
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// 何种事件可以触发质量校验任务执行
	//
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) SetTaskIds(v []*int64) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger {
	s.TaskIds = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger) SetType(v string) *ListDataQualityEvaluationTaskInstancesResponseBodyPagingInfoDataQualityEvaluationTaskInstancesTaskTrigger {
	s.Type = &v
	return s
}

type ListDataQualityEvaluationTaskInstancesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityEvaluationTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityEvaluationTaskInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) SetHeaders(v map[string]*string) *ListDataQualityEvaluationTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) SetStatusCode(v int32) *ListDataQualityEvaluationTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityEvaluationTaskInstancesResponse) SetBody(v *ListDataQualityEvaluationTaskInstancesResponseBody) *ListDataQualityEvaluationTaskInstancesResponse {
	s.Body = v
	return s
}

type ListDataQualityEvaluationTasksRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s ListDataQualityEvaluationTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksRequest) SetName(v string) *ListDataQualityEvaluationTasksRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetPageNumber(v int32) *ListDataQualityEvaluationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetPageSize(v int32) *ListDataQualityEvaluationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetProjectId(v int64) *ListDataQualityEvaluationTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityEvaluationTasksRequest) SetTableGuid(v string) *ListDataQualityEvaluationTasksRequest {
	s.TableGuid = &v
	return s
}

type ListDataQualityEvaluationTasksResponseBody struct {
	// example:
	//
	// 200
	Code       *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PagingInfo *ListDataQualityEvaluationTasksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBody) SetCode(v string) *ListDataQualityEvaluationTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBody) SetMessage(v string) *ListDataQualityEvaluationTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBody) SetPagingInfo(v *ListDataQualityEvaluationTasksResponseBodyPagingInfo) *ListDataQualityEvaluationTasksResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBody) SetRequestId(v string) *ListDataQualityEvaluationTasksResponseBody {
	s.RequestId = &v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfo struct {
	DataQualityEvaluationTasks []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks `json:"DataQualityEvaluationTasks,omitempty" xml:"DataQualityEvaluationTasks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 131
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) SetDataQualityEvaluationTasks(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) *ListDataQualityEvaluationTasksResponseBodyPagingInfo {
	s.DataQualityEvaluationTasks = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) SetPageNumber(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) SetPageSize(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) SetTotalCount(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks struct {
	// example:
	//
	// This is a daily run data quality evaluation plan
	Description *string                                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks       []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id            *int64                                                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Name          *string                                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Notifications *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK-SQL" }
	RuntimeConf *string                                                                               `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	Target      *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// 10
	TenantId *int64                                                                                 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Trigger  *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetDescription(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Description = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetHooks(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Hooks = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetId(v int64) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Id = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetName(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Name = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetNotifications(v *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Notifications = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetProjectId(v int64) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetRuntimeConf(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.RuntimeConf = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetTarget(v *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Target = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetTenantId(v int64) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.TenantId = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetTrigger(v *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Trigger = v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks struct {
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) SetCondition(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) SetType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks {
	s.Type = &v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications struct {
	// example:
	//
	// ${severity} == "High"
	Condition     *string                                                                                                     `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Notifications []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) SetCondition(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) SetNotifications(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications {
	s.Notifications = v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications struct {
	NotificationChannels  []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	NotificationReceivers []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) SetNotificationChannels(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) SetNotificationReceivers(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels struct {
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers struct {
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

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) SetExtension(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) SetReceiverType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) SetReceiverValues(v []*string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget struct {
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

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) SetDatabaseType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget {
	s.DatabaseType = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) SetPartitionSpec(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget {
	s.PartitionSpec = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) SetTableGuid(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) SetType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget {
	s.Type = &v
	return s
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger struct {
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// ByScheduledTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) SetTaskIds(v []*int64) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger {
	s.TaskIds = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) SetType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger {
	s.Type = &v
	return s
}

type ListDataQualityEvaluationTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityEvaluationTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponse) SetHeaders(v map[string]*string) *ListDataQualityEvaluationTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponse) SetStatusCode(v int32) *ListDataQualityEvaluationTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponse) SetBody(v *ListDataQualityEvaluationTasksResponseBody) *ListDataQualityEvaluationTasksResponse {
	s.Body = v
	return s
}

type ListDataQualityResultsRequest struct {
	// example:
	//
	// 2024-05-01
	BizdateFrom *string `json:"BizdateFrom,omitempty" xml:"BizdateFrom,omitempty"`
	// example:
	//
	// 2024-05-04
	BizdateTo *string `json:"BizdateTo,omitempty" xml:"BizdateTo,omitempty"`
	// example:
	//
	// 1710239005403
	CreateTimeFrom *int64 `json:"CreateTimeFrom,omitempty" xml:"CreateTimeFrom,omitempty"`
	// example:
	//
	// 1710239005403
	CreateTimeTo *int64 `json:"CreateTimeTo,omitempty" xml:"CreateTimeTo,omitempty"`
	// example:
	//
	// 200001
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// example:
	//
	// 10001
	DataQualityEvaluationTaskInstanceId *int64 `json:"DataQualityEvaluationTaskInstanceId,omitempty" xml:"DataQualityEvaluationTaskInstanceId,omitempty"`
	// example:
	//
	// 100001
	DataQualityRuleId *int64 `json:"DataQualityRuleId,omitempty" xml:"DataQualityRuleId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataQualityResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsRequest) SetBizdateFrom(v string) *ListDataQualityResultsRequest {
	s.BizdateFrom = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetBizdateTo(v string) *ListDataQualityResultsRequest {
	s.BizdateTo = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetCreateTimeFrom(v int64) *ListDataQualityResultsRequest {
	s.CreateTimeFrom = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetCreateTimeTo(v int64) *ListDataQualityResultsRequest {
	s.CreateTimeTo = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetDataQualityEvaluationTaskId(v int64) *ListDataQualityResultsRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetDataQualityEvaluationTaskInstanceId(v int64) *ListDataQualityResultsRequest {
	s.DataQualityEvaluationTaskInstanceId = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetDataQualityRuleId(v int64) *ListDataQualityResultsRequest {
	s.DataQualityRuleId = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetPageNumber(v int32) *ListDataQualityResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetPageSize(v int32) *ListDataQualityResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityResultsRequest) SetProjectId(v int64) *ListDataQualityResultsRequest {
	s.ProjectId = &v
	return s
}

type ListDataQualityResultsResponseBody struct {
	// example:
	//
	// 200
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PagingInfo *ListDataQualityResultsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBody) SetCode(v string) *ListDataQualityResultsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataQualityResultsResponseBody) SetMessage(v string) *ListDataQualityResultsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataQualityResultsResponseBody) SetPagingInfo(v *ListDataQualityResultsResponseBodyPagingInfo) *ListDataQualityResultsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityResultsResponseBody) SetRequestId(v string) *ListDataQualityResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfo struct {
	DataQualityResults []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResults `json:"DataQualityResults,omitempty" xml:"DataQualityResults,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 219
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetDataQualityResults(v []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) *ListDataQualityResultsResponseBodyPagingInfo {
	s.DataQualityResults = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataQualityResultsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetPageSize(v int32) *ListDataQualityResultsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataQualityResultsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResults struct {
	// example:
	//
	// 1708284916414
	CreateTime *int64                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Details    []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// 16033
	Id   *int64                                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Rule *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// example:
	//
	// [
	//
	//   {
	//
	//     "gender": "male",
	//
	//     "_count": 100
	//
	//   }, {
	//
	//     "gender": "female",
	//
	//     "_count": 100
	//
	//   }
	//
	// ]
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// 	- RUNNING
	//
	// 	- ERROR
	//
	// 	- PASSED
	//
	// 	- WARNED
	//
	// 	- CRITICAL
	//
	// example:
	//
	// PASSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 200001
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetCreateTime(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.CreateTime = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetDetails(v []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Details = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Id = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetRule(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Rule = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetSample(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Sample = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetStatus(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Status = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetTaskInstanceId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.TaskInstanceId = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails struct {
	// example:
	//
	// 100.0
	CheckedValue *string `json:"CheckedValue,omitempty" xml:"CheckedValue,omitempty"`
	// example:
	//
	// 0.0
	ReferencedValue *string `json:"ReferencedValue,omitempty" xml:"ReferencedValue,omitempty"`
	// 	- ERROR
	//
	// 	- PASSED
	//
	// 	- WARNED
	//
	// 	- CRITICAL
	//
	// example:
	//
	// PASSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) SetCheckedValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	s.CheckedValue = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) SetReferencedValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	s.ReferencedValue = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) SetStatus(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	s.Status = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule struct {
	CheckingConfig *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled       *bool                                                                              `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// example:
	//
	// 100001
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	ProjectId      *int64                                                                            `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// 	- HIGH
	//
	// 	- NORMAL
	//
	// example:
	//
	// NORMAL
	Severity *string                                                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target   *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 1
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetCheckingConfig(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.CheckingConfig = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetDescription(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Description = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetEnabled(v bool) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Enabled = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetErrorHandlers(v []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.ErrorHandlers = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Id = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetName(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Name = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetProjectId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetSamplingConfig(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.SamplingConfig = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetSeverity(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Severity = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetTarget(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Target = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetTemplateCode(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.TemplateCode = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetTenantId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.TenantId = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string                                                                                     `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// 	- FIXED
	//
	// 	- FLUCTATION
	//
	// 	- AUTO
	//
	// 	- AVERAGE
	//
	// 	- VARIANCE
	//
	// 	- FLUCTATION_DISCREATE
	//
	// example:
	//
	// FIXED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) SetReferencedSamplesFilter(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) SetThresholds(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) SetType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	s.Type = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds struct {
	Critical *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) SetCritical(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) SetExpected(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) SetWarned(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical struct {
	// 	- \\>
	//
	// 	- \\>=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) SetOperator(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) SetValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected struct {
	// 	- \\>
	//
	// 	- \\>=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) SetOperator(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) SetValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned struct {
	// 	- \\>
	//
	// 	- \\>=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) SetOperator(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) SetValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers struct {
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// 	- SAVE_ERROR_DATA
	//
	// example:
	//
	// SAVE_ERROR_DATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) SetErrorDataFilter(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) SetType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers {
	s.Type = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig struct {
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// *
	//
	// example:
	//
	// COUNT
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// { "columns": [ "id", "name" ] }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetMetric(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetMetricParameters(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetSamplingFilter(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetSettingConfig(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget struct {
	// 	- MAX_COMPUTE
	//
	// 	- EMR
	//
	// 	- CDH
	//
	// 	- HOLOGRES
	//
	// 	- ANALYTICDB_FOR_POSTGRESQL
	//
	// 	- ANALYTICDB_FOR_MYSQL
	//
	// 	- STAR_ROCKS
	//
	// example:
	//
	// MAX_COMPUTE
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// 	- TABLE
	//
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetDatabaseType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetPartitionSpec(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.PartitionSpec = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetTableGuid(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.Type = &v
	return s
}

type ListDataQualityResultsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponse) SetHeaders(v map[string]*string) *ListDataQualityResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityResultsResponse) SetStatusCode(v int32) *ListDataQualityResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityResultsResponse) SetBody(v *ListDataQualityResultsResponseBody) *ListDataQualityResultsResponse {
	s.Body = v
	return s
}

type ListDataQualityRulesRequest struct {
	// example:
	//
	// 10000
	DataQualityEvaluationTaskId *int64 `json:"DataQualityEvaluationTaskId,omitempty" xml:"DataQualityEvaluationTaskId,omitempty"`
	// example:
	//
	// unit_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10002
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s ListDataQualityRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesRequest) SetDataQualityEvaluationTaskId(v int64) *ListDataQualityRulesRequest {
	s.DataQualityEvaluationTaskId = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetName(v string) *ListDataQualityRulesRequest {
	s.Name = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetPageNumber(v int32) *ListDataQualityRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetPageSize(v int32) *ListDataQualityRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetProjectId(v int64) *ListDataQualityRulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityRulesRequest) SetTableGuid(v string) *ListDataQualityRulesRequest {
	s.TableGuid = &v
	return s
}

type ListDataQualityRulesResponseBody struct {
	// example:
	//
	// 200
	Code       *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PagingInfo *ListDataQualityRulesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBody) SetCode(v string) *ListDataQualityRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataQualityRulesResponseBody) SetMessage(v string) *ListDataQualityRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataQualityRulesResponseBody) SetPagingInfo(v *ListDataQualityRulesResponseBodyPagingInfo) *ListDataQualityRulesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityRulesResponseBody) SetRequestId(v string) *ListDataQualityRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfo struct {
	DataQualityRules []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRules `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetDataQualityRules(v []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) *ListDataQualityRulesResponseBodyPagingInfo {
	s.DataQualityRules = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataQualityRulesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetPageSize(v int32) *ListDataQualityRulesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataQualityRulesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRules struct {
	CheckingConfig *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled       *bool                                                                      `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// example:
	//
	// 22130
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100001
	ProjectId      *int64                                                                    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// example:
	//
	// High
	Severity *string                                                           `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target   *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 100001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetCheckingConfig(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.CheckingConfig = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetDescription(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Description = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetEnabled(v bool) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Enabled = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetErrorHandlers(v []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.ErrorHandlers = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetId(v int64) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Id = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetName(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Name = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetProjectId(v int64) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetSamplingConfig(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.SamplingConfig = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetSeverity(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Severity = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetTarget(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Target = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetTemplateCode(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.TemplateCode = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetTenantId(v int64) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.TenantId = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string                                                                             `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) SetReferencedSamplesFilter(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) SetThresholds(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) SetType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	s.Type = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds struct {
	Critical *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) SetCritical(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) SetExpected(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) SetWarned(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	s.Warned = v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) SetOperator(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) SetValue(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) SetOperator(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) SetValue(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) SetOperator(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) SetValue(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers struct {
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) SetErrorDataFilter(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) SetType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers {
	s.Type = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig struct {
	// example:
	//
	// Max
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetMetric(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.Metric = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetMetricParameters(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetSamplingFilter(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetSettingConfig(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.SettingConfig = &v
	return s
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget struct {
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

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetDatabaseType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.DatabaseType = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetPartitionSpec(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.PartitionSpec = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetTableGuid(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.Type = &v
	return s
}

type ListDataQualityRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataQualityRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataQualityRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataQualityRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponse) SetHeaders(v map[string]*string) *ListDataQualityRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDataQualityRulesResponse) SetStatusCode(v int32) *ListDataQualityRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataQualityRulesResponse) SetBody(v *ListDataQualityRulesResponseBody) *ListDataQualityRulesResponse {
	s.Body = v
	return s
}

type ListDataSourceSharedRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 1
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ListDataSourceSharedRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesRequest) SetDataSourceId(v int64) *ListDataSourceSharedRulesRequest {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourceSharedRulesRequest) SetTargetProjectId(v int64) *ListDataSourceSharedRulesRequest {
	s.TargetProjectId = &v
	return s
}

type ListDataSourceSharedRulesResponseBody struct {
	DataSourceSharedRules []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules `json:"DataSourceSharedRules,omitempty" xml:"DataSourceSharedRules,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceSharedRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponseBody) SetDataSourceSharedRules(v []*ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) *ListDataSourceSharedRulesResponseBody {
	s.DataSourceSharedRules = v
	return s
}

func (s *ListDataSourceSharedRulesResponseBody) SetRequestId(v string) *ListDataSourceSharedRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourceSharedRulesResponseBodyDataSourceSharedRules struct {
	// example:
	//
	// 1724379762000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// targetProject.datasource
	SharedDataSourceName *string `json:"SharedDataSourceName,omitempty" xml:"SharedDataSourceName,omitempty"`
	// example:
	//
	// 1
	SharedUser *string `json:"SharedUser,omitempty" xml:"SharedUser,omitempty"`
	// example:
	//
	// 1
	SourceProjectId *int64 `json:"SourceProjectId,omitempty" xml:"SourceProjectId,omitempty"`
	// example:
	//
	// 1
	TargetProjectId *int64 `json:"TargetProjectId,omitempty" xml:"TargetProjectId,omitempty"`
}

func (s ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetCreateTime(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetCreateUser(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.CreateUser = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetDataSourceId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.DataSourceId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetEnvType(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.EnvType = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.Id = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSharedDataSourceName(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SharedDataSourceName = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSharedUser(v string) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SharedUser = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetSourceProjectId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.SourceProjectId = &v
	return s
}

func (s *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules) SetTargetProjectId(v int64) *ListDataSourceSharedRulesResponseBodyDataSourceSharedRules {
	s.TargetProjectId = &v
	return s
}

type ListDataSourceSharedRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceSharedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceSharedRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceSharedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceSharedRulesResponse) SetHeaders(v map[string]*string) *ListDataSourceSharedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceSharedRulesResponse) SetStatusCode(v int32) *ListDataSourceSharedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceSharedRulesResponse) SetBody(v *ListDataSourceSharedRulesResponseBody) *ListDataSourceSharedRulesResponse {
	s.Body = v
	return s
}

type ListDataSourcesRequest struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17820
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// ["tag1", "tag2", "tag3"]
	Tags  *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) SetEnvType(v string) *ListDataSourcesRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesRequest) SetName(v string) *ListDataSourcesRequest {
	s.Name = &v
	return s
}

func (s *ListDataSourcesRequest) SetOrder(v string) *ListDataSourcesRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageNumber(v int32) *ListDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageSize(v int32) *ListDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesRequest) SetProjectId(v int64) *ListDataSourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesRequest) SetSortBy(v string) *ListDataSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataSourcesRequest) SetTags(v string) *ListDataSourcesRequest {
	s.Tags = &v
	return s
}

func (s *ListDataSourcesRequest) SetTypes(v []*string) *ListDataSourcesRequest {
	s.Types = v
	return s
}

type ListDataSourcesShrinkRequest struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17820
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// ["tag1", "tag2", "tag3"]
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListDataSourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesShrinkRequest) SetEnvType(v string) *ListDataSourcesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetName(v string) *ListDataSourcesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetOrder(v string) *ListDataSourcesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageNumber(v int32) *ListDataSourcesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetPageSize(v int32) *ListDataSourcesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetProjectId(v int64) *ListDataSourcesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetSortBy(v string) *ListDataSourcesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetTags(v string) *ListDataSourcesShrinkRequest {
	s.Tags = &v
	return s
}

func (s *ListDataSourcesShrinkRequest) SetTypesShrink(v string) *ListDataSourcesShrinkRequest {
	s.TypesShrink = &v
	return s
}

type ListDataSourcesResponseBody struct {
	PagingInfo *ListDataSourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 7BE1433F-6D55-5D86-9344-CA6F7DD19B13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetPagingInfo(v *ListDataSourcesResponseBodyPagingInfo) *ListDataSourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourcesResponseBodyPagingInfo struct {
	DataSources []*ListDataSourcesResponseBodyPagingInfoDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 131
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetDataSources(v []*ListDataSourcesResponseBodyPagingInfoDataSources) *ListDataSourcesResponseBodyPagingInfo {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetPageNumber(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetPageSize(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfo) SetTotalCount(v int64) *ListDataSourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDataSourcesResponseBodyPagingInfoDataSources struct {
	DataSource []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// mysql
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfoDataSources) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfoDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetDataSource(v []*ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.DataSource = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetName(v string) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.Name = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSources) SetType(v string) *ListDataSourcesResponseBodyPagingInfoDataSources {
	s.Type = &v
	return s
}

type ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource struct {
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// 1648711113000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1624387842781448
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1648711113000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 1624387842781448
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// 1648711121000:cn-beijing:odps:yongxunQA_beijing_standard
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetConnectionProperties(v interface{}) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ConnectionProperties = v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetConnectionPropertiesMode(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetCreateTime(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.CreateTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetCreateUser(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.CreateUser = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetDescription(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.Description = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetId(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetModifyTime(v int64) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ModifyTime = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetModifyUser(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.ModifyUser = &v
	return s
}

func (s *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource) SetQualifiedName(v string) *ListDataSourcesResponseBodyPagingInfoDataSourcesDataSource {
	s.QualifiedName = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetStatusCode(v int32) *ListDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListDeploymentsRequest struct {
	// The ID of the user who creates the processes. This parameter specifies a filter condition.
	//
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the processes. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- INIT
	//
	// 	- RUNNING
	//
	// 	- SUCCESS
	//
	// 	- FAIL
	//
	// 	- TERMINATION
	//
	// 	- CANCEL
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) SetCreator(v string) *ListDeploymentsRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageNumber(v int32) *ListDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetProjectId(v string) *ListDeploymentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v string) *ListDeploymentsRequest {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDeploymentsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBody) SetPagingInfo(v *ListDeploymentsResponseBodyPagingInfo) *ListDeploymentsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDeploymentsResponseBody) SetRequestId(v string) *ListDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfo struct {
	// The processes.
	Deployments []*ListDeploymentsResponseBodyPagingInfoDeployments `json:"Deployments,omitempty" xml:"Deployments,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2524
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetDeployments(v []*ListDeploymentsResponseBodyPagingInfoDeployments) *ListDeploymentsResponseBodyPagingInfo {
	s.Deployments = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetPageNumber(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetPageSize(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetTotalCount(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfoDeployments struct {
	// The time when the process was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1702736654000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who creates the process.
	//
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The process ID.
	//
	// example:
	//
	// ddf354a5-03df-48fc-94c1-cc973f79XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned if the process fails.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the process was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1702736654000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 44683
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The stages of the process.
	Stages []*ListDeploymentsResponseBodyPagingInfoDeploymentsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// The status of the process.
	//
	// Valid values:
	//
	// 	- INIT
	//
	// 	- RUNNING
	//
	// 	- FAIL
	//
	// 	- SUCCESS
	//
	// 	- TERMINATION
	//
	// 	- CANCEL
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfoDeployments) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfoDeployments) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetCreateTime(v int64) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.CreateTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetCreator(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetId(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Id = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetMessage(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Message = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetModifyTime(v int64) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.ModifyTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetProjectId(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetStages(v []*ListDeploymentsResponseBodyPagingInfoDeploymentsStages) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Stages = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetStatus(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfoDeploymentsStages struct {
	// The code of the stage.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the stage.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The additional information about the stage.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The error message returned during the stage.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the stage.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the stage.
	//
	// Valid values:
	//
	// 	- INIT
	//
	// 	- RUNNING
	//
	// 	- SUCCESS
	//
	// 	- FAIL
	//
	// 	- TERMINATION
	//
	// 	- CANCEL
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The step number of the stage.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// The type of the stage. This parameter indicates the operation type in the stage.
	//
	// Valid values:
	//
	// 	- DEPLOY
	//
	// 	- CHECK
	//
	// 	- OFFLINE.
	//
	// 	- BUILD
	//
	// 	- DELETE
	//
	// example:
	//
	// CHECK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfoDeploymentsStages) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfoDeploymentsStages) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetCode(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Code = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetDescription(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Description = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetDetail(v map[string]interface{}) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Detail = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetMessage(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Message = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetName(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Name = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetStatus(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Status = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetStep(v int32) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Step = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetType(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Type = &v
	return s
}

type ListDeploymentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponse) SetHeaders(v map[string]*string) *ListDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentsResponse) SetStatusCode(v int32) *ListDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentsResponse) SetBody(v *ListDeploymentsResponseBody) *ListDeploymentsResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	// The ID of the owner of the UDF. This parameter specifies a filter condition.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The UDF type. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- MATH: mathematical operation function
	//
	// 	- AGGREGATE: aggregate function
	//
	// 	- STRING: string processing function
	//
	// 	- DATE: date function
	//
	// 	- ANALYTIC: window function
	//
	// 	- OTHER: others
	//
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetOwner(v string) *ListFunctionsRequest {
	s.Owner = &v
	return s
}

func (s *ListFunctionsRequest) SetPageNumber(v int32) *ListFunctionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsRequest) SetPageSize(v int32) *ListFunctionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsRequest) SetProjectId(v string) *ListFunctionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsRequest) SetType(v string) *ListFunctionsRequest {
	s.Type = &v
	return s
}

type ListFunctionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListFunctionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 89FB2BF0-EB00-5D03-9C34-05931001XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetPagingInfo(v *ListFunctionsResponseBodyPagingInfo) *ListFunctionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionsResponseBodyPagingInfo struct {
	// The UDFs.
	Functions []*ListFunctionsResponseBodyPagingInfoFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfo) SetFunctions(v []*ListFunctionsResponseBodyPagingInfoFunctions) *ListFunctionsResponseBodyPagingInfo {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageSize(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctions struct {
	// The file resources in an Advanced RISC Machines (ARM) cluster.
	//
	// example:
	//
	// xxx.jar,yyy.jar
	ArmResource *string `json:"ArmResource,omitempty" xml:"ArmResource,omitempty"`
	// The fully qualified class name of the UDF.
	//
	// example:
	//
	// com.demo.Main
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// The description of the command.
	//
	// example:
	//
	// testUdf(xx,yy)
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The time when the UDF was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1655953028000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source information about the UDF.
	DataSource *ListFunctionsResponseBodyPagingInfoFunctionsDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The name of the database. This parameter is returned for E-MapReduce (EMR) functions.
	//
	// example:
	//
	// odps_first
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The overall description of the UDF.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The code of the embedded UDF.
	//
	// example:
	//
	// print(\\"hello,world!\\")
	EmbeddedCode *string `json:"EmbeddedCode,omitempty" xml:"EmbeddedCode,omitempty"`
	// The type of the nested code.
	//
	// Valid values:
	//
	// 	- Python2
	//
	// 	- Python3
	//
	// 	- Java8
	//
	// 	- Java11
	//
	// 	- Java17
	//
	// example:
	//
	// Python2
	EmbeddedCodeType *string `json:"EmbeddedCodeType,omitempty" xml:"EmbeddedCodeType,omitempty"`
	// The type of the nested resource.
	//
	// Valid values:
	//
	// 	- File: general resources
	//
	// 	- Embedded: embedded resources
	//
	// example:
	//
	// File
	EmbeddedResourceType *string `json:"EmbeddedResourceType,omitempty" xml:"EmbeddedResourceType,omitempty"`
	// The description of the example.
	ExampleDescription *string `json:"ExampleDescription,omitempty" xml:"ExampleDescription,omitempty"`
	// The files resources.
	//
	// example:
	//
	// xxx.jar,yyy.jar
	FileResource *string `json:"FileResource,omitempty" xml:"FileResource,omitempty"`
	// The ID of the UDF.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the UDF was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1655953028000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the UDF.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the UDF.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The description of the parameter.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The ID of the workspace to which the UDF belongs.
	//
	// example:
	//
	// 307XXX
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The description of the return value.
	ReturnValueDescription *string `json:"ReturnValueDescription,omitempty" xml:"ReturnValueDescription,omitempty"`
	// The information about the resource group used when you run the UDF.
	RuntimeResource *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information about the UDF.
	Script *ListFunctionsResponseBodyPagingInfoFunctionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The UDF type.
	//
	// Valid values:
	//
	// 	- MATH: mathematical operation function
	//
	// 	- AGGREGATE: aggregate function
	//
	// 	- STRING: string processing function
	//
	// 	- DATE: date function
	//
	// 	- ANALYTIC: window function
	//
	// 	- OTHER: others
	//
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetArmResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ArmResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetClassName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ClassName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCommandDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CommandDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCreateTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDataSource(v *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DataSource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDatabaseName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DatabaseName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Description = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCode(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCode = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCodeType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCodeType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedResourceType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedResourceType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetExampleDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ExampleDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetFileResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.FileResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetModifyTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ModifyTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetOwner(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetParameterDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ParameterDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetProjectId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetReturnValueDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ReturnValueDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetRuntimeResource(v *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.RuntimeResource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetScript(v *ListFunctionsResponseBodyPagingInfoFunctionsScript) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Script = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Type = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Type = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource struct {
	// The ID of the resource group used when you run the UDF.
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) SetResourceGroupId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsScript struct {
	// The script ID.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetPath(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Path = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetRuntime(v *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Runtime = v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime struct {
	// The command.
	//
	// example:
	//
	// ODPS_FUNCTION
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) SetCommand(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime {
	s.Command = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListNetworksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNetworksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetworksRequest) GoString() string {
	return s.String()
}

func (s *ListNetworksRequest) SetResourceGroupId(v string) *ListNetworksRequest {
	s.ResourceGroupId = &v
	return s
}

type ListNetworksResponseBody struct {
	NetworkList []*ListNetworksResponseBodyNetworkList `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNetworksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworksResponseBody) SetNetworkList(v []*ListNetworksResponseBodyNetworkList) *ListNetworksResponseBody {
	s.NetworkList = v
	return s
}

func (s *ListNetworksResponseBody) SetRequestId(v string) *ListNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworksResponseBody) SetSuccess(v bool) *ListNetworksResponseBody {
	s.Success = &v
	return s
}

type ListNetworksResponseBodyNetworkList struct {
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-2ze13vamugr7jenXXXXX
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListNetworksResponseBodyNetworkList) String() string {
	return tea.Prettify(s)
}

func (s ListNetworksResponseBodyNetworkList) GoString() string {
	return s.String()
}

func (s *ListNetworksResponseBodyNetworkList) SetCreateTime(v int64) *ListNetworksResponseBodyNetworkList {
	s.CreateTime = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetCreateUser(v string) *ListNetworksResponseBodyNetworkList {
	s.CreateUser = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetId(v int64) *ListNetworksResponseBodyNetworkList {
	s.Id = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetResourceGroupId(v string) *ListNetworksResponseBodyNetworkList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetSecurityGroupId(v string) *ListNetworksResponseBodyNetworkList {
	s.SecurityGroupId = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetStatus(v string) *ListNetworksResponseBodyNetworkList {
	s.Status = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetVpcId(v string) *ListNetworksResponseBodyNetworkList {
	s.VpcId = &v
	return s
}

func (s *ListNetworksResponseBodyNetworkList) SetVswitchId(v string) *ListNetworksResponseBodyNetworkList {
	s.VswitchId = &v
	return s
}

type ListNetworksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworksResponse) GoString() string {
	return s.String()
}

func (s *ListNetworksResponse) SetHeaders(v map[string]*string) *ListNetworksResponse {
	s.Headers = v
	return s
}

func (s *ListNetworksResponse) SetStatusCode(v int32) *ListNetworksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworksResponse) SetBody(v *ListNetworksResponseBody) *ListNetworksResponse {
	s.Body = v
	return s
}

type ListNodeDependenciesRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodeDependenciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesRequest) SetId(v string) *ListNodeDependenciesRequest {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageNumber(v int32) *ListNodeDependenciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageSize(v int32) *ListNodeDependenciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetProjectId(v string) *ListNodeDependenciesRequest {
	s.ProjectId = &v
	return s
}

type ListNodeDependenciesResponseBody struct {
	// The pagination information.
	PagingInfo *ListNodeDependenciesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 204EAF68-CCE3-5112-8DA0-E7A60F02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeDependenciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBody) SetPagingInfo(v *ListNodeDependenciesResponseBodyPagingInfo) *ListNodeDependenciesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodeDependenciesResponseBody) SetRequestId(v string) *ListNodeDependenciesResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfo struct {
	// The descendant nodes.
	Nodes []*ListNodeDependenciesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 90
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetNodes(v []*ListNodeDependenciesResponseBodyPagingInfoNodes) *ListNodeDependenciesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageSize(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodes struct {
	// The time when the node was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data source.
	DataSource *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the node.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 723932906364267XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input of the node.
	Inputs *ListNodeDependenciesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The time when the node was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the node.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the node.
	Outputs *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The owner of the node.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 65133
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type.
	//
	// Valid values:
	//
	// 	- Normal: The node is scheduled as expected.
	//
	// 	- Pause: The node is paused, and the running of its descendant nodes is blocked.
	//
	// 	- Skip: The node is dry run. The system does not actually run the node but directly prompts that the node is successfully run. The running duration of the node is 0 seconds. In addition, the node does not occupy resources or block the running of its descendant nodes.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The information about the resource group.
	RuntimeResource *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *ListNodeDependenciesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The scheduling policy.
	Strategy *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// The tags. This parameter is not in use.
	Tags []*ListNodeDependenciesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The scheduling task ID.
	//
	// example:
	//
	// 580667964888595XXXX
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trigger.
	Trigger *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetInputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetProjectId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetScript(v *ListNodeDependenciesResponseBodyPagingInfoNodesScript) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTags(v []*ListNodeDependenciesResponseBodyPagingInfoNodesTags) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTaskId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputs struct {
	// The node outputs.
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The tables.
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variables.
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables struct {
	// The artifact type.
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The variable ID.
	//
	// example:
	//
	// 543218872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable.
	//
	// Valid values:
	//
	// 	- NodeParameter
	//
	// 	- NodeContext
	//
	// 	- Workflow
	//
	// 	- Workspace
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the variable.
	//
	// Valid values:
	//
	// 	- NoKvVariableExpression
	//
	// 	- Constant
	//
	// 	- PassThrough
	//
	// 	- System
	//
	// 	- NodeOutput
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// The output of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputs struct {
	// The node outputs.
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The tables.
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variables.
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 463497880880954XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables struct {
	// The artifact type.
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The variable ID.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable.
	//
	// Valid values:
	//
	// 	- NodeParameter
	//
	// 	- NodeContext
	//
	// 	- Workflow
	//
	// 	- Workspace
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the variable.
	//
	// Valid values:
	//
	// 	- NoKvVariableExpression
	//
	// 	- Constant
	//
	// 	- PassThrough
	//
	// 	- System
	//
	// 	- NodeOutput
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// The output of the node to which the variable belongs.
	//
	// example:
	//
	// 463497880880954XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource struct {
	// The resource group ID.
	//
	// example:
	//
	// S_res_group_XXXX_XXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScript struct {
	// The script ID.
	//
	// example:
	//
	// 853573334108680XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime struct {
	// The command used to distinguish node types.
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesStrategy struct {
	// The instance generation mode.
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The rerun interval after a failure. Unit: milliseconds.
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun mode.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of reruns after a failure.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTags struct {
	// The tag key.
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTrigger struct {
	// The CRON expression for scheduling.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the validity period of the scheduling. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The trigger ID.
	//
	// example:
	//
	// 543680677872062XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The start time of the validity period of the scheduling. The time is in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The type of the trigger.
	//
	// Valid values:
	//
	// 	- Scheduler
	//
	// 	- Manual
	//
	// 	- Streaming
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

type ListNodeDependenciesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeDependenciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponse) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponse) SetHeaders(v map[string]*string) *ListNodeDependenciesResponse {
	s.Headers = v
	return s
}

func (s *ListNodeDependenciesResponse) SetStatusCode(v int32) *ListNodeDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeDependenciesResponse) SetBody(v *ListNodeDependenciesResponseBody) *ListNodeDependenciesResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// The container ID. This parameter specifies a filter condition.
	//
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The rerun mode. Valid values:
	//
	// 	- Allowed: The nodes can be rerun regardless of whether they are successfully run or fail to run.
	//
	// 	- FailureAllowed: The nodes can be rerun only after they fail to run.
	//
	// 	- Denied: The nodes cannot be rerun regardless of whether they are successfully run or fail to run.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The scene of nodes. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- DATAWORKS_PROJECT
	//
	// 	- MANUAL_WORKFLOW
	//
	// 	- MANUAL_NODE
	//
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetContainerId(v string) *ListNodesRequest {
	s.ContainerId = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetProjectId(v string) *ListNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequest) SetRecurrence(v string) *ListNodesRequest {
	s.Recurrence = &v
	return s
}

func (s *ListNodesRequest) SetRerunMode(v string) *ListNodesRequest {
	s.RerunMode = &v
	return s
}

func (s *ListNodesRequest) SetScene(v string) *ListNodesRequest {
	s.Scene = &v
	return s
}

type ListNodesResponseBody struct {
	// The pagination information.
	PagingInfo *ListNodesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2197B9C4-39CE-55EA-8EEA-FDBAE52DXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetPagingInfo(v *ListNodesResponseBodyPagingInfo) *ListNodesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListNodesResponseBodyPagingInfo struct {
	// The nodes.
	Nodes []*ListNodesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 42
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfo) SetNodes(v []*ListNodesResponseBodyPagingInfoNodes) *ListNodesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageSize(v string) *ListNodesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodes struct {
	// The time when the node was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1722910655000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data source.
	DataSource *ListNodesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the node.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input of the node.
	Inputs *ListNodesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The time when the node was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1722910655000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the node.
	Outputs *ListNodesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The owner of the node.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// example:
	//
	// 33233
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type.
	//
	// Valid values:
	//
	// 	- Normal: The node is scheduled as expected.
	//
	// 	- Pause: The node is paused, and the running of its descendant nodes is blocked.
	//
	// 	- Skip: The node is dry run. The system does not actually run the node but directly prompts that the node is successfully run. The running duration of the node is 0 seconds. In addition, the node does not occupy resources or block the running of its descendant nodes.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The information about the resource group.
	RuntimeResource *ListNodesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *ListNodesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The scheduling policy.
	Strategy *ListNodesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// The tags. This parameter is not in use.
	Tags []*ListNodesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The scheduling task ID.
	//
	// example:
	//
	// 88888888888
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trigger.
	Trigger *ListNodesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodesResponseBodyPagingInfoNodesDataSource) *ListNodesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetInputs(v *ListNodesResponseBodyPagingInfoNodesInputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetName(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodesResponseBodyPagingInfoNodesOutputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetProjectId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodesResponseBodyPagingInfoNodesRuntimeResource) *ListNodesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetScript(v *ListNodesResponseBodyPagingInfoNodesScript) *ListNodesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodesResponseBodyPagingInfoNodesStrategy) *ListNodesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTags(v []*ListNodesResponseBodyPagingInfoNodesTags) *ListNodesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTaskId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodesResponseBodyPagingInfoNodesTrigger) *ListNodesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputs struct {
	// The node outputs.
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The tables.
	Tables []*ListNodesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variables.
	Variables []*ListNodesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesInputsTables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesInputsVariables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 623731286945488XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsVariables struct {
	// The artifact type.
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The variable ID.
	//
	// example:
	//
	// 543211286945488XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable.
	//
	// Valid values:
	//
	// 	- WorkSpace
	//
	// 	- NodeParameter
	//
	// 	- NodeContext
	//
	// 	- Workflow
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the variable.
	//
	// Valid values:
	//
	// 	- NoKvVariableExpression
	//
	// 	- Constant
	//
	// 	- PassThrough
	//
	// 	- System
	//
	// 	- NodeOutput
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// 222
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetId(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// The output of the node.
	//
	// example:
	//
	// 623731286945488XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputs struct {
	// The node outputs.
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The tables.
	Tables []*ListNodesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variables.
	Variables []*ListNodesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesOutputsTables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesOutputsVariables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariables struct {
	// The artifact type.
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The variable ID.
	//
	// example:
	//
	// 623731286945488XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable.
	//
	// Valid values:
	//
	// 	- NodeParameter
	//
	// 	- NodeContext
	//
	// 	- Workflow
	//
	// 	- Workspace
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the variable.
	//
	// Valid values:
	//
	// 	- NoKvVariableExpression
	//
	// 	- Constant
	//
	// 	- PassThrough
	//
	// 	- System
	//
	// 	- NodeOutput
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetId(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// The output of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesRuntimeResource struct {
	// The resource group ID.
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesScript struct {
	// The script ID.
	//
	// example:
	//
	// 853573334108680XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListNodesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodesScript) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetId(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodesResponseBodyPagingInfoNodesScriptRuntime) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesScriptRuntime struct {
	// The command used to distinguish node types.
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesStrategy struct {
	// The instance generation mode.
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The rerun interval. Unit: milliseconds.
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun mode.
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of reruns.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesTags struct {
	// The tag key.
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesTrigger struct {
	// The CRON expression for scheduling.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the validity period of the trigger.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The trigger ID.
	//
	// example:
	//
	// 543680677872062XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The start time of the validity period of the trigger.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The type of the trigger.
	//
	// Valid values:
	//
	// 	- Scheduler
	//
	// 	- Manual
	//
	// 	- Steaming
	//
	// <!---->
	//
	// *
	//
	// *
	//
	// *
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetId(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListProjectMembersRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62136
	ProjectId *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	UserIds   []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequest) SetPageNumber(v int32) *ListProjectMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersRequest) SetPageSize(v int32) *ListProjectMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersRequest) SetProjectId(v int64) *ListProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersRequest) SetRoleCodes(v []*string) *ListProjectMembersRequest {
	s.RoleCodes = v
	return s
}

func (s *ListProjectMembersRequest) SetUserIds(v []*string) *ListProjectMembersRequest {
	s.UserIds = v
	return s
}

type ListProjectMembersShrinkRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62136
	ProjectId       *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	UserIdsShrink   *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s ListProjectMembersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersShrinkRequest) SetPageNumber(v int32) *ListProjectMembersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetPageSize(v int32) *ListProjectMembersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetProjectId(v int64) *ListProjectMembersShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetRoleCodesShrink(v string) *ListProjectMembersShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *ListProjectMembersShrinkRequest) SetUserIdsShrink(v string) *ListProjectMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

type ListProjectMembersResponseBody struct {
	PagingInfo *ListProjectMembersResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 9FBBBB1F-DD5E-5D8E-8F50-37F77460F056
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBody) SetPagingInfo(v *ListProjectMembersResponseBodyPagingInfo) *ListProjectMembersResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectMembersResponseBody) SetRequestId(v string) *ListProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectMembersResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize       *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectMembers []*ListProjectMembersResponseBodyPagingInfoProjectMembers `json:"ProjectMembers,omitempty" xml:"ProjectMembers,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectMembersResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetPageNumber(v int32) *ListProjectMembersResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetPageSize(v int32) *ListProjectMembersResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetProjectMembers(v []*ListProjectMembersResponseBodyPagingInfoProjectMembers) *ListProjectMembersResponseBodyPagingInfo {
	s.ProjectMembers = v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfo) SetTotalCount(v int32) *ListProjectMembersResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListProjectMembersResponseBodyPagingInfoProjectMembers struct {
	// example:
	//
	// 62136
	ProjectId *int64                                                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Roles     []*ListProjectMembersResponseBodyPagingInfoProjectMembersRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembers) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembers) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetProjectId(v int64) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetRoles(v []*ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.Roles = v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetStatus(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.Status = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembers) SetUserId(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembers {
	s.UserId = &v
	return s
}

type ListProjectMembersResponseBodyPagingInfoProjectMembersRoles struct {
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) SetCode(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	s.Code = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) SetName(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	s.Name = &v
	return s
}

func (s *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles) SetType(v string) *ListProjectMembersResponseBodyPagingInfoProjectMembersRoles {
	s.Type = &v
	return s
}

type ListProjectMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponse) SetHeaders(v map[string]*string) *ListProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *ListProjectMembersResponse) SetStatusCode(v int32) *ListProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectMembersResponse) SetBody(v *ListProjectMembersResponseBody) *ListProjectMembersResponse {
	s.Body = v
	return s
}

type ListProjectRolesRequest struct {
	Codes []*string `json:"Codes,omitempty" xml:"Codes,omitempty" type:"Repeated"`
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRolesRequest) SetCodes(v []*string) *ListProjectRolesRequest {
	s.Codes = v
	return s
}

func (s *ListProjectRolesRequest) SetNames(v []*string) *ListProjectRolesRequest {
	s.Names = v
	return s
}

func (s *ListProjectRolesRequest) SetPageNumber(v int32) *ListProjectRolesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRolesRequest) SetPageSize(v int32) *ListProjectRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRolesRequest) SetProjectId(v int64) *ListProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesRequest) SetType(v string) *ListProjectRolesRequest {
	s.Type = &v
	return s
}

type ListProjectRolesShrinkRequest struct {
	CodesShrink *string `json:"Codes,omitempty" xml:"Codes,omitempty"`
	NamesShrink *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectRolesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRolesShrinkRequest) SetCodesShrink(v string) *ListProjectRolesShrinkRequest {
	s.CodesShrink = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetNamesShrink(v string) *ListProjectRolesShrinkRequest {
	s.NamesShrink = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetPageNumber(v int32) *ListProjectRolesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetPageSize(v int32) *ListProjectRolesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetProjectId(v int64) *ListProjectRolesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesShrinkRequest) SetType(v string) *ListProjectRolesShrinkRequest {
	s.Type = &v
	return s
}

type ListProjectRolesResponseBody struct {
	PagingInfo *ListProjectRolesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 61649187-0BCF-5E75-8D4B-64FDBEBBB447
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBody) SetPagingInfo(v *ListProjectRolesResponseBodyPagingInfo) *ListProjectRolesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectRolesResponseBody) SetRequestId(v string) *ListProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectRolesResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *string                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectRoles []*ListProjectRolesResponseBodyPagingInfoProjectRoles `json:"ProjectRoles,omitempty" xml:"ProjectRoles,omitempty" type:"Repeated"`
	// example:
	//
	// 42
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectRolesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRolesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetPageNumber(v string) *ListProjectRolesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetPageSize(v string) *ListProjectRolesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetProjectRoles(v []*ListProjectRolesResponseBodyPagingInfoProjectRoles) *ListProjectRolesResponseBodyPagingInfo {
	s.ProjectRoles = v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfo) SetTotalCount(v string) *ListProjectRolesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListProjectRolesResponseBodyPagingInfoProjectRoles struct {
	// example:
	//
	// role_project_guest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 21229
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// System
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectRolesResponseBodyPagingInfoProjectRoles) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRolesResponseBodyPagingInfoProjectRoles) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetCode(v string) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.Code = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetName(v string) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.Name = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetProjectId(v int64) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.ProjectId = &v
	return s
}

func (s *ListProjectRolesResponseBodyPagingInfoProjectRoles) SetType(v string) *ListProjectRolesResponseBodyPagingInfoProjectRoles {
	s.Type = &v
	return s
}

type ListProjectRolesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRolesResponse) GoString() string {
	return s.String()
}

func (s *ListProjectRolesResponse) SetHeaders(v map[string]*string) *ListProjectRolesResponse {
	s.Headers = v
	return s
}

func (s *ListProjectRolesResponse) SetStatusCode(v int32) *ListProjectRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectRolesResponse) SetBody(v *ListProjectRolesResponseBody) *ListProjectRolesResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string                                  `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*ListProjectsRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool     `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	Ids             []*int64  `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	Names           []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetAliyunResourceGroupId(v string) *ListProjectsRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsRequest) SetAliyunResourceTags(v []*ListProjectsRequestAliyunResourceTags) *ListProjectsRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *ListProjectsRequest) SetDevEnvironmentEnabled(v bool) *ListProjectsRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsRequest) SetDevRoleDisabled(v bool) *ListProjectsRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsRequest) SetIds(v []*int64) *ListProjectsRequest {
	s.Ids = v
	return s
}

func (s *ListProjectsRequest) SetNames(v []*string) *ListProjectsRequest {
	s.Names = v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetPaiTaskEnabled(v bool) *ListProjectsRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsRequest) SetStatus(v string) *ListProjectsRequest {
	s.Status = &v
	return s
}

type ListProjectsRequestAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsRequestAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListProjectsRequestAliyunResourceTags) SetKey(v string) *ListProjectsRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListProjectsRequestAliyunResourceTags) SetValue(v string) *ListProjectsRequestAliyunResourceTags {
	s.Value = &v
	return s
}

type ListProjectsShrinkRequest struct {
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId    *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	IdsShrink       *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	NamesShrink     *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsShrinkRequest) SetAliyunResourceGroupId(v string) *ListProjectsShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetAliyunResourceTagsShrink(v string) *ListProjectsShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetDevEnvironmentEnabled(v bool) *ListProjectsShrinkRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetDevRoleDisabled(v bool) *ListProjectsShrinkRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetIdsShrink(v string) *ListProjectsShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetNamesShrink(v string) *ListProjectsShrinkRequest {
	s.NamesShrink = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageNumber(v int32) *ListProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPageSize(v int32) *ListProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetPaiTaskEnabled(v bool) *ListProjectsShrinkRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsShrinkRequest) SetStatus(v string) *ListProjectsShrinkRequest {
	s.Status = &v
	return s
}

type ListProjectsResponseBody struct {
	PagingInfo *ListProjectsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 6D24AD9A-652F-59E2-AC1F-05029300F8A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetPagingInfo(v *ListProjectsResponseBodyPagingInfo) *ListProjectsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyPagingInfo struct {
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Projects []*ListProjectsResponseBodyPagingInfoProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfo) SetPageNumber(v int32) *ListProjectsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetPageSize(v int32) *ListProjectsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetProjects(v []*ListProjectsResponseBodyPagingInfoProjects) *ListProjectsResponseBodyPagingInfo {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBodyPagingInfo) SetTotalCount(v int32) *ListProjectsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListProjectsResponseBodyPagingInfoProjects struct {
	// example:
	//
	// rg-acfmzbn7pti3zfa
	AliyunResourceGroupId *string                                                         `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	AliyunResourceTags    []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	Description           *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// false
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123532153125
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfoProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfoProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetAliyunResourceGroupId(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetAliyunResourceTags(v []*ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) *ListProjectsResponseBodyPagingInfoProjects {
	s.AliyunResourceTags = v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDescription(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDevEnvironmentEnabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDevRoleDisabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.DevRoleDisabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetDisplayName(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.DisplayName = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetId(v int64) *ListProjectsResponseBodyPagingInfoProjects {
	s.Id = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetName(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetOwner(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetPaiTaskEnabled(v bool) *ListProjectsResponseBodyPagingInfoProjects {
	s.PaiTaskEnabled = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjects) SetStatus(v string) *ListProjectsResponseBodyPagingInfoProjects {
	s.Status = &v
	return s
}

type ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags struct {
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) SetKey(v string) *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags) SetValue(v string) *ListProjectsResponseBodyPagingInfoProjectsAliyunResourceTags {
	s.Value = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListResourceGroupsRequest struct {
	// example:
	//
	// Resource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 1000
	ProjectId          *int64    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceGroupTypes []*string `json:"ResourceGroupTypes,omitempty" xml:"ResourceGroupTypes,omitempty" type:"Repeated"`
	Statuses           []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsRequest) SetName(v string) *ListResourceGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsRequest) SetPaymentType(v string) *ListResourceGroupsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsRequest) SetProjectId(v int64) *ListResourceGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourceGroupsRequest) SetResourceGroupTypes(v []*string) *ListResourceGroupsRequest {
	s.ResourceGroupTypes = v
	return s
}

func (s *ListResourceGroupsRequest) SetStatuses(v []*string) *ListResourceGroupsRequest {
	s.Statuses = v
	return s
}

type ListResourceGroupsShrinkRequest struct {
	// example:
	//
	// Resource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 1000
	ProjectId                *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceGroupTypesShrink *string `json:"ResourceGroupTypes,omitempty" xml:"ResourceGroupTypes,omitempty"`
	StatusesShrink           *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListResourceGroupsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsShrinkRequest) SetName(v string) *ListResourceGroupsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetPaymentType(v string) *ListResourceGroupsShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetProjectId(v int64) *ListResourceGroupsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetResourceGroupTypesShrink(v string) *ListResourceGroupsShrinkRequest {
	s.ResourceGroupTypesShrink = &v
	return s
}

func (s *ListResourceGroupsShrinkRequest) SetStatusesShrink(v string) *ListResourceGroupsShrinkRequest {
	s.StatusesShrink = &v
	return s
}

type ListResourceGroupsResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupList []*ListResourceGroupsResponseBodyResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetResourceGroupList(v []*ListResourceGroupsResponseBodyResourceGroupList) *ListResourceGroupsResponseBody {
	s.ResourceGroupList = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetSuccess(v bool) *ListResourceGroupsResponseBody {
	s.Success = &v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupList struct {
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	DefaultVpcId *string `json:"DefaultVpcId,omitempty" xml:"DefaultVpcId,omitempty"`
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	DefaultVswicthId *string `json:"DefaultVswicthId,omitempty" xml:"DefaultVswicthId,omitempty"`
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// c442b330-3b10-4584-959e-736e4edXXXXX
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// example:
	//
	// 创建用于普通任务的通用资源组
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// CommonV2
	ResourceGroupType *string                                              `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	Spec              *ListResourceGroupsResponseBodyResourceGroupListSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceGroupsResponseBodyResourceGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetCreateTime(v int64) *ListResourceGroupsResponseBodyResourceGroupList {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetCreateUser(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.CreateUser = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetDefaultVpcId(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.DefaultVpcId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetDefaultVswicthId(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.DefaultVswicthId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetId(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetName(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetOrderInstanceId(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.OrderInstanceId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetPaymentType(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetRemark(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.Remark = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetResourceGroupType(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetSpec(v *ListResourceGroupsResponseBodyResourceGroupListSpec) *ListResourceGroupsResponseBodyResourceGroupList {
	s.Spec = v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupList) SetStatus(v string) *ListResourceGroupsResponseBodyResourceGroupList {
	s.Status = &v
	return s
}

type ListResourceGroupsResponseBodyResourceGroupListSpec struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 2CU
	Standard *string `json:"Standard,omitempty" xml:"Standard,omitempty"`
}

func (s ListResourceGroupsResponseBodyResourceGroupListSpec) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponseBodyResourceGroupListSpec) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyResourceGroupListSpec) SetAmount(v int32) *ListResourceGroupsResponseBodyResourceGroupListSpec {
	s.Amount = &v
	return s
}

func (s *ListResourceGroupsResponseBodyResourceGroupListSpec) SetStandard(v string) *ListResourceGroupsResponseBodyResourceGroupListSpec {
	s.Standard = &v
	return s
}

type ListResourceGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponse) SetHeaders(v map[string]*string) *ListResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupsResponse) SetStatusCode(v int32) *ListResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupsResponse) SetBody(v *ListResourceGroupsResponseBody) *ListResourceGroupsResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// The ID of the Alibaba Cloud account used by the workspace administrator. You can log on to the Alibaba Cloud Management Console and view the ID on the Security Settings page.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource type. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- Python
	//
	// 	- Jar
	//
	// 	- Archive
	//
	// 	- File
	//
	// example:
	//
	// python
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetOwner(v string) *ListResourcesRequest {
	s.Owner = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProjectId(v string) *ListResourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesRequest) SetType(v string) *ListResourcesRequest {
	s.Type = &v
	return s
}

type ListResourcesResponseBody struct {
	// The pagination information.
	PagingInfo *ListResourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPagingInfo(v *ListResourcesResponseBodyPagingInfo) *ListResourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The file resources.
	Resources []*ListResourcesResponseBodyPagingInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 131
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageNumber(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageSize(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetResources(v []*ListResourcesResponseBodyPagingInfoResources) *ListResourcesResponseBodyPagingInfo {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetTotalCount(v int32) *ListResourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResources struct {
	// The time when the file resource was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data source.
	DataSource *ListResourcesResponseBodyPagingInfoResourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The ID of the file resource.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The times when the file resource was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the file resource.
	//
	// example:
	//
	// math.py
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the file resource.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// example:
	//
	// 344247
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The script information.
	Script *ListResourcesResponseBodyPagingInfoResourcesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The path of the source of the file resource. If the SourecType parameter is set to Local, this parameter is left empty.
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The storage type of the source of the file resource.
	//
	// Valid values:
	//
	// 	- Local
	//
	// 	- Oss
	//
	// example:
	//
	// local
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The storage path of the destination of the file resource.
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The storage type of the destination of the file resource.
	//
	// Valid values:
	//
	// 	- Gateway
	//
	// 	- Oss
	//
	// 	- Hdfs
	//
	// example:
	//
	// oss
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The type of the file resource.
	//
	// Valid values:
	//
	// 	- Python
	//
	// 	- Jar
	//
	// 	- Archive
	//
	// 	- File
	//
	// example:
	//
	// jar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetCreateTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetDataSource(v *ListResourcesResponseBodyPagingInfoResourcesDataSource) *ListResourcesResponseBodyPagingInfoResources {
	s.DataSource = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetId(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetModifyTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.ModifyTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetName(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetOwner(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetProjectId(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetScript(v *ListResourcesResponseBodyPagingInfoResourcesScript) *ListResourcesResponseBodyPagingInfoResources {
	s.Script = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourcePath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourcePath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourceType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourceType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetPath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetPath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Type = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetName(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetType(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Type = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesScript struct {
	// The script ID.
	//
	// example:
	//
	// 123348864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetId(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetPath(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Path = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetRuntime(v *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Runtime = v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesScriptRuntime struct {
	// The command used to distinguish file resource types.
	//
	// example:
	//
	// ODPS_PYTHON
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) SetCommand(v string) *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime {
	s.Command = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListRoutesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
}

func (s ListRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListRoutesRequest) SetNetworkId(v int64) *ListRoutesRequest {
	s.NetworkId = &v
	return s
}

type ListRoutesResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteList []*ListRoutesResponseBodyRouteList `json:"RouteList,omitempty" xml:"RouteList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoutesResponseBody) SetRequestId(v string) *ListRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoutesResponseBody) SetRouteList(v []*ListRoutesResponseBodyRouteList) *ListRoutesResponseBody {
	s.RouteList = v
	return s
}

func (s *ListRoutesResponseBody) SetSuccess(v bool) *ListRoutesResponseBody {
	s.Success = &v
	return s
}

type ListRoutesResponseBodyRouteList struct {
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1000
	NetworkId *int64 `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ns-679XXXXXX
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ListRoutesResponseBodyRouteList) String() string {
	return tea.Prettify(s)
}

func (s ListRoutesResponseBodyRouteList) GoString() string {
	return s.String()
}

func (s *ListRoutesResponseBodyRouteList) SetCreateTime(v int64) *ListRoutesResponseBodyRouteList {
	s.CreateTime = &v
	return s
}

func (s *ListRoutesResponseBodyRouteList) SetDestinationCidr(v string) *ListRoutesResponseBodyRouteList {
	s.DestinationCidr = &v
	return s
}

func (s *ListRoutesResponseBodyRouteList) SetId(v int64) *ListRoutesResponseBodyRouteList {
	s.Id = &v
	return s
}

func (s *ListRoutesResponseBodyRouteList) SetNetworkId(v int64) *ListRoutesResponseBodyRouteList {
	s.NetworkId = &v
	return s
}

func (s *ListRoutesResponseBodyRouteList) SetResourceGroupId(v string) *ListRoutesResponseBodyRouteList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRoutesResponseBodyRouteList) SetResourceId(v string) *ListRoutesResponseBodyRouteList {
	s.ResourceId = &v
	return s
}

type ListRoutesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListRoutesResponse) SetHeaders(v map[string]*string) *ListRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListRoutesResponse) SetStatusCode(v int32) *ListRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutesResponse) SetBody(v *ListRoutesResponseBody) *ListRoutesResponse {
	s.Body = v
	return s
}

type ListWorkflowDefinitionsRequest struct {
	// The ID of the Alibaba Cloud account used by the workspace administrator. You can log on to the Alibaba Cloud Management Console and view the ID on the Security Settings page.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The workflow type. This parameter specifies a filter condition.
	//
	// Valid values:
	//
	// 	- CycleWorkflow
	//
	// 	- ManualWorkflow
	//
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsRequest) SetOwner(v string) *ListWorkflowDefinitionsRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageNumber(v int32) *ListWorkflowDefinitionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageSize(v int32) *ListWorkflowDefinitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetProjectId(v string) *ListWorkflowDefinitionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetType(v string) *ListWorkflowDefinitionsRequest {
	s.Type = &v
	return s
}

type ListWorkflowDefinitionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListWorkflowDefinitionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8C3ED0C5-ABAB-55E1-854B-DAC02B11XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBody) SetPagingInfo(v *ListWorkflowDefinitionsResponseBodyPagingInfo) *ListWorkflowDefinitionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBody) SetRequestId(v string) *ListWorkflowDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 227
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The workflows.
	WorkflowDefinitions []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions `json:"WorkflowDefinitions,omitempty" xml:"WorkflowDefinitions,omitempty" type:"Repeated"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageSize(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetWorkflowDefinitions(v []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.WorkflowDefinitions = v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions struct {
	// The time when the workflow was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698057323000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the workflow.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The times when the workflow was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698057323000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the workflow.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace to which the workflow belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4710
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The script information.
	Script *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The type of the workflow.
	//
	// Valid values:
	//
	// 	- CycleWorkflow
	//
	// 	- ManualWorkflow
	//
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetCreateTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetDescription(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Description = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetModifyTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetName(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Name = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetOwner(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetProjectId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetScript(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Script = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetType(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Type = &v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript struct {
	// The script ID.
	//
	// example:
	//
	// 698002781368644XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetPath(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Path = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetRuntime(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Runtime = v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime struct {
	// The command.
	//
	// example:
	//
	// WORKFLOW
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) SetCommand(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime {
	s.Command = &v
	return s
}

type ListWorkflowDefinitionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowDefinitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponse) SetHeaders(v map[string]*string) *ListWorkflowDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetStatusCode(v int32) *ListWorkflowDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetBody(v *ListWorkflowDefinitionsResponseBody) *ListWorkflowDefinitionsResponse {
	s.Body = v
	return s
}

type MoveFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the UDF. You do not need to specify a UDF name in the path.
	//
	// For example, if you want to move the test UDF to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionRequest) GoString() string {
	return s.String()
}

func (s *MoveFunctionRequest) SetId(v string) *MoveFunctionRequest {
	s.Id = &v
	return s
}

func (s *MoveFunctionRequest) SetPath(v string) *MoveFunctionRequest {
	s.Path = &v
	return s
}

func (s *MoveFunctionRequest) SetProjectId(v string) *MoveFunctionRequest {
	s.ProjectId = &v
	return s
}

type MoveFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 48C0E2F0-52BA-5888-BDFA-28F1B9AFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponseBody) SetRequestId(v string) *MoveFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveFunctionResponseBody) SetSuccess(v bool) *MoveFunctionResponseBody {
	s.Success = &v
	return s
}

type MoveFunctionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionResponse) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponse) SetHeaders(v map[string]*string) *MoveFunctionResponse {
	s.Headers = v
	return s
}

func (s *MoveFunctionResponse) SetStatusCode(v int32) *MoveFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveFunctionResponse) SetBody(v *MoveFunctionResponseBody) *MoveFunctionResponse {
	s.Body = v
	return s
}

type MoveNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the node. You do not need to specify a node name in the path.
	//
	// For example, if you want to move the test node to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeRequest) GoString() string {
	return s.String()
}

func (s *MoveNodeRequest) SetId(v string) *MoveNodeRequest {
	s.Id = &v
	return s
}

func (s *MoveNodeRequest) SetPath(v string) *MoveNodeRequest {
	s.Path = &v
	return s
}

func (s *MoveNodeRequest) SetProjectId(v string) *MoveNodeRequest {
	s.ProjectId = &v
	return s
}

type MoveNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeResponseBody) GoString() string {
	return s.String()
}

func (s *MoveNodeResponseBody) SetRequestId(v string) *MoveNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveNodeResponseBody) SetSuccess(v bool) *MoveNodeResponseBody {
	s.Success = &v
	return s
}

type MoveNodeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeResponse) GoString() string {
	return s.String()
}

func (s *MoveNodeResponse) SetHeaders(v map[string]*string) *MoveNodeResponse {
	s.Headers = v
	return s
}

func (s *MoveNodeResponse) SetStatusCode(v int32) *MoveNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveNodeResponse) SetBody(v *MoveNodeResponseBody) *MoveNodeResponse {
	s.Body = v
	return s
}

type MoveResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the file resource. You do not need to specify a file resource name in the path.
	//
	// For example, if you want to move the test file resource to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) SetId(v string) *MoveResourceRequest {
	s.Id = &v
	return s
}

func (s *MoveResourceRequest) SetPath(v string) *MoveResourceRequest {
	s.Path = &v
	return s
}

func (s *MoveResourceRequest) SetProjectId(v string) *MoveResourceRequest {
	s.ProjectId = &v
	return s
}

type MoveResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F332BED4-DD73-5972-A9C2-642BA6CFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceResponseBody) SetRequestId(v string) *MoveResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSuccess(v bool) *MoveResourceResponseBody {
	s.Success = &v
	return s
}

type MoveResourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceResponse) SetHeaders(v map[string]*string) *MoveResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceResponse) SetStatusCode(v int32) *MoveResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceResponse) SetBody(v *MoveResourceResponseBody) *MoveResourceResponse {
	s.Body = v
	return s
}

type MoveWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The path to which you want to move the workflow. You do not need to specify a workflow name in the path.
	//
	// For example, if you want to move the test workflow to root/demo/test, you must set this parameter to root/demo.
	//
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionRequest) SetId(v string) *MoveWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetPath(v string) *MoveWorkflowDefinitionRequest {
	s.Path = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetProjectId(v string) *MoveWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type MoveWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05ADAF4F-7709-5FB1-B606-3513483FXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponseBody) SetRequestId(v string) *MoveWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveWorkflowDefinitionResponseBody) SetSuccess(v bool) *MoveWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type MoveWorkflowDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *MoveWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetStatusCode(v int32) *MoveWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetBody(v *MoveWorkflowDefinitionResponseBody) *MoveWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type RenameFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionRequest) GoString() string {
	return s.String()
}

func (s *RenameFunctionRequest) SetId(v string) *RenameFunctionRequest {
	s.Id = &v
	return s
}

func (s *RenameFunctionRequest) SetName(v string) *RenameFunctionRequest {
	s.Name = &v
	return s
}

func (s *RenameFunctionRequest) SetProjectId(v string) *RenameFunctionRequest {
	s.ProjectId = &v
	return s
}

type RenameFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1ED4C97F-BA2A-57C5-BA7C-8853627EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponseBody) SetRequestId(v string) *RenameFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameFunctionResponseBody) SetSuccess(v string) *RenameFunctionResponseBody {
	s.Success = &v
	return s
}

type RenameFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionResponse) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponse) SetHeaders(v map[string]*string) *RenameFunctionResponse {
	s.Headers = v
	return s
}

func (s *RenameFunctionResponse) SetStatusCode(v int32) *RenameFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameFunctionResponse) SetBody(v *RenameFunctionResponseBody) *RenameFunctionResponse {
	s.Body = v
	return s
}

type RenameNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeRequest) GoString() string {
	return s.String()
}

func (s *RenameNodeRequest) SetId(v string) *RenameNodeRequest {
	s.Id = &v
	return s
}

func (s *RenameNodeRequest) SetName(v string) *RenameNodeRequest {
	s.Name = &v
	return s
}

func (s *RenameNodeRequest) SetProjectId(v string) *RenameNodeRequest {
	s.ProjectId = &v
	return s
}

type RenameNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RenameNodeResponseBody) SetRequestId(v string) *RenameNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameNodeResponseBody) SetSuccess(v bool) *RenameNodeResponseBody {
	s.Success = &v
	return s
}

type RenameNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeResponse) GoString() string {
	return s.String()
}

func (s *RenameNodeResponse) SetHeaders(v map[string]*string) *RenameNodeResponse {
	s.Headers = v
	return s
}

func (s *RenameNodeResponse) SetStatusCode(v int32) *RenameNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameNodeResponse) SetBody(v *RenameNodeResponseBody) *RenameNodeResponse {
	s.Body = v
	return s
}

type RenameResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceRequest) GoString() string {
	return s.String()
}

func (s *RenameResourceRequest) SetId(v string) *RenameResourceRequest {
	s.Id = &v
	return s
}

func (s *RenameResourceRequest) SetName(v string) *RenameResourceRequest {
	s.Name = &v
	return s
}

func (s *RenameResourceRequest) SetProjectId(v string) *RenameResourceRequest {
	s.ProjectId = &v
	return s
}

type RenameResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA8XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameResourceResponseBody) SetRequestId(v string) *RenameResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameResourceResponseBody) SetSuccess(v bool) *RenameResourceResponseBody {
	s.Success = &v
	return s
}

type RenameResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceResponse) GoString() string {
	return s.String()
}

func (s *RenameResourceResponse) SetHeaders(v map[string]*string) *RenameResourceResponse {
	s.Headers = v
	return s
}

func (s *RenameResourceResponse) SetStatusCode(v int32) *RenameResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameResourceResponse) SetBody(v *RenameResourceResponseBody) *RenameResourceResponse {
	s.Body = v
	return s
}

type RenameWorkflowDefinitionRequest struct {
	// The unique identifier of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name.
	//
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionRequest) SetId(v string) *RenameWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetName(v string) *RenameWorkflowDefinitionRequest {
	s.Name = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetProjectId(v string) *RenameWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type RenameWorkflowDefinitionResponseBody struct {
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// 975BD43D-C421-595C-A29C-565A8AD5XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponseBody) SetRequestId(v string) *RenameWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameWorkflowDefinitionResponseBody) SetSuccess(v bool) *RenameWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type RenameWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *RenameWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetStatusCode(v int32) *RenameWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetBody(v *RenameWorkflowDefinitionResponseBody) *RenameWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type RevokeMemberProjectRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RoleCodes []*string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeMemberProjectRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeMemberProjectRolesRequest) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesRequest) SetProjectId(v int64) *RevokeMemberProjectRolesRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeMemberProjectRolesRequest) SetRoleCodes(v []*string) *RevokeMemberProjectRolesRequest {
	s.RoleCodes = v
	return s
}

func (s *RevokeMemberProjectRolesRequest) SetUserId(v string) *RevokeMemberProjectRolesRequest {
	s.UserId = &v
	return s
}

type RevokeMemberProjectRolesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 105149
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RoleCodesShrink *string `json:"RoleCodes,omitempty" xml:"RoleCodes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123422344899
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeMemberProjectRolesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeMemberProjectRolesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesShrinkRequest) SetProjectId(v int64) *RevokeMemberProjectRolesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeMemberProjectRolesShrinkRequest) SetRoleCodesShrink(v string) *RevokeMemberProjectRolesShrinkRequest {
	s.RoleCodesShrink = &v
	return s
}

func (s *RevokeMemberProjectRolesShrinkRequest) SetUserId(v string) *RevokeMemberProjectRolesShrinkRequest {
	s.UserId = &v
	return s
}

type RevokeMemberProjectRolesResponseBody struct {
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeMemberProjectRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeMemberProjectRolesResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesResponseBody) SetRequestId(v string) *RevokeMemberProjectRolesResponseBody {
	s.RequestId = &v
	return s
}

type RevokeMemberProjectRolesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeMemberProjectRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeMemberProjectRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeMemberProjectRolesResponse) GoString() string {
	return s.String()
}

func (s *RevokeMemberProjectRolesResponse) SetHeaders(v map[string]*string) *RevokeMemberProjectRolesResponse {
	s.Headers = v
	return s
}

func (s *RevokeMemberProjectRolesResponse) SetStatusCode(v int32) *RevokeMemberProjectRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeMemberProjectRolesResponse) SetBody(v *RevokeMemberProjectRolesResponseBody) *RevokeMemberProjectRolesResponse {
	s.Body = v
	return s
}

type StartDIJobRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 10000
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// false
	ForceToRerun          *bool                                   `json:"ForceToRerun,omitempty" xml:"ForceToRerun,omitempty"`
	RealtimeStartSettings *StartDIJobRequestRealtimeStartSettings `json:"RealtimeStartSettings,omitempty" xml:"RealtimeStartSettings,omitempty" type:"Struct"`
}

func (s StartDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobRequest) GoString() string {
	return s.String()
}

func (s *StartDIJobRequest) SetDIJobId(v int64) *StartDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *StartDIJobRequest) SetForceToRerun(v bool) *StartDIJobRequest {
	s.ForceToRerun = &v
	return s
}

func (s *StartDIJobRequest) SetRealtimeStartSettings(v *StartDIJobRequestRealtimeStartSettings) *StartDIJobRequest {
	s.RealtimeStartSettings = v
	return s
}

type StartDIJobRequestRealtimeStartSettings struct {
	FailoverSettings *StartDIJobRequestRealtimeStartSettingsFailoverSettings `json:"FailoverSettings,omitempty" xml:"FailoverSettings,omitempty" type:"Struct"`
	// example:
	//
	// 1671516776
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s StartDIJobRequestRealtimeStartSettings) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobRequestRealtimeStartSettings) GoString() string {
	return s.String()
}

func (s *StartDIJobRequestRealtimeStartSettings) SetFailoverSettings(v *StartDIJobRequestRealtimeStartSettingsFailoverSettings) *StartDIJobRequestRealtimeStartSettings {
	s.FailoverSettings = v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettings) SetStartTime(v int64) *StartDIJobRequestRealtimeStartSettings {
	s.StartTime = &v
	return s
}

type StartDIJobRequestRealtimeStartSettingsFailoverSettings struct {
	// example:
	//
	// 10
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// 30
	UpperLimit *int64 `json:"UpperLimit,omitempty" xml:"UpperLimit,omitempty"`
}

func (s StartDIJobRequestRealtimeStartSettingsFailoverSettings) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobRequestRealtimeStartSettingsFailoverSettings) GoString() string {
	return s.String()
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) SetInterval(v int64) *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	s.Interval = &v
	return s
}

func (s *StartDIJobRequestRealtimeStartSettingsFailoverSettings) SetUpperLimit(v int64) *StartDIJobRequestRealtimeStartSettingsFailoverSettings {
	s.UpperLimit = &v
	return s
}

type StartDIJobShrinkRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 10000
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// example:
	//
	// false
	ForceToRerun                *bool   `json:"ForceToRerun,omitempty" xml:"ForceToRerun,omitempty"`
	RealtimeStartSettingsShrink *string `json:"RealtimeStartSettings,omitempty" xml:"RealtimeStartSettings,omitempty"`
}

func (s StartDIJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartDIJobShrinkRequest) SetDIJobId(v int64) *StartDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetForceToRerun(v bool) *StartDIJobShrinkRequest {
	s.ForceToRerun = &v
	return s
}

func (s *StartDIJobShrinkRequest) SetRealtimeStartSettingsShrink(v string) *StartDIJobShrinkRequest {
	s.RealtimeStartSettingsShrink = &v
	return s
}

type StartDIJobResponseBody struct {
	// example:
	//
	// 999431B2-6013-577F-B684-36F7433C753B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartDIJobResponseBody) SetRequestId(v string) *StartDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDIJobResponseBody) SetSuccess(v bool) *StartDIJobResponseBody {
	s.Success = &v
	return s
}

type StartDIJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDIJobResponse) GoString() string {
	return s.String()
}

func (s *StartDIJobResponse) SetHeaders(v map[string]*string) *StartDIJobResponse {
	s.Headers = v
	return s
}

func (s *StartDIJobResponse) SetStatusCode(v int32) *StartDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDIJobResponse) SetBody(v *StartDIJobResponseBody) *StartDIJobResponse {
	s.Body = v
	return s
}

type StopDIJobRequest struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11668
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDIJobRequest) GoString() string {
	return s.String()
}

func (s *StopDIJobRequest) SetDIJobId(v int64) *StopDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *StopDIJobRequest) SetInstanceId(v int64) *StopDIJobRequest {
	s.InstanceId = &v
	return s
}

type StopDIJobResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 92F778C7-8F00-53B1-AE1A-B3B17101247D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopDIJobResponseBody) SetRequestId(v string) *StopDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDIJobResponseBody) SetSuccess(v bool) *StopDIJobResponseBody {
	s.Success = &v
	return s
}

type StopDIJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDIJobResponse) GoString() string {
	return s.String()
}

func (s *StopDIJobResponse) SetHeaders(v map[string]*string) *StopDIJobResponse {
	s.Headers = v
	return s
}

func (s *StopDIJobResponse) SetStatusCode(v int32) *StopDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDIJobResponse) SetBody(v *StopDIJobResponseBody) *StopDIJobResponse {
	s.Body = v
	return s
}

type UpdateDIAlarmRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 34982
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// example:
	//
	// 1
	DIJobId     *int64  `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// alarm_rule_name
	Name                 *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationSettings *UpdateDIAlarmRuleRequestNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	TriggerConditions    []*UpdateDIAlarmRuleRequestTriggerConditions  `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s UpdateDIAlarmRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequest) SetDIAlarmRuleId(v int64) *UpdateDIAlarmRuleRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetDIJobId(v int64) *UpdateDIAlarmRuleRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetDescription(v string) *UpdateDIAlarmRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetEnabled(v bool) *UpdateDIAlarmRuleRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetMetricType(v string) *UpdateDIAlarmRuleRequest {
	s.MetricType = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetName(v string) *UpdateDIAlarmRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetNotificationSettings(v *UpdateDIAlarmRuleRequestNotificationSettings) *UpdateDIAlarmRuleRequest {
	s.NotificationSettings = v
	return s
}

func (s *UpdateDIAlarmRuleRequest) SetTriggerConditions(v []*UpdateDIAlarmRuleRequestTriggerConditions) *UpdateDIAlarmRuleRequest {
	s.TriggerConditions = v
	return s
}

type UpdateDIAlarmRuleRequestNotificationSettings struct {
	// example:
	//
	// 5
	InhibitionInterval    *int64                                                               `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	NotificationChannels  []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels  `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	NotificationReceivers []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s UpdateDIAlarmRuleRequestNotificationSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestNotificationSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) SetInhibitionInterval(v int64) *UpdateDIAlarmRuleRequestNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) SetNotificationChannels(v []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) *UpdateDIAlarmRuleRequestNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettings) SetNotificationReceivers(v []*UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) *UpdateDIAlarmRuleRequestNotificationSettings {
	s.NotificationReceivers = v
	return s
}

type UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels struct {
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetChannels(v []*string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels) SetSeverity(v string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

type UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers struct {
	// example:
	//
	// DingToken
	ReceiverType   *string   `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverType(v string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *UpdateDIAlarmRuleRequestNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

type UpdateDIAlarmRuleRequestTriggerConditions struct {
	DdlReportTags []*string `json:"DdlReportTags,omitempty" xml:"DdlReportTags,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// Warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s UpdateDIAlarmRuleRequestTriggerConditions) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleRequestTriggerConditions) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetDdlReportTags(v []*string) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.DdlReportTags = v
	return s
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetDuration(v int64) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.Duration = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetSeverity(v string) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.Severity = &v
	return s
}

func (s *UpdateDIAlarmRuleRequestTriggerConditions) SetThreshold(v int64) *UpdateDIAlarmRuleRequestTriggerConditions {
	s.Threshold = &v
	return s
}

type UpdateDIAlarmRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 34982
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// example:
	//
	// 1
	DIJobId     *int64  `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// alarm_rule_name
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationSettingsShrink *string `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty"`
	TriggerConditionsShrink    *string `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty"`
}

func (s UpdateDIAlarmRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetDIAlarmRuleId(v int64) *UpdateDIAlarmRuleShrinkRequest {
	s.DIAlarmRuleId = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetDIJobId(v int64) *UpdateDIAlarmRuleShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetDescription(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetEnabled(v bool) *UpdateDIAlarmRuleShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetMetricType(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.MetricType = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetName(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetNotificationSettingsShrink(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.NotificationSettingsShrink = &v
	return s
}

func (s *UpdateDIAlarmRuleShrinkRequest) SetTriggerConditionsShrink(v string) *UpdateDIAlarmRuleShrinkRequest {
	s.TriggerConditionsShrink = &v
	return s
}

type UpdateDIAlarmRuleResponseBody struct {
	// example:
	//
	// A6C6B486-E3A2-5D52-9E76-D9380485D946
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDIAlarmRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleResponseBody) SetRequestId(v string) *UpdateDIAlarmRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIAlarmRuleResponseBody) SetSuccess(v bool) *UpdateDIAlarmRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateDIAlarmRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDIAlarmRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDIAlarmRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIAlarmRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDIAlarmRuleResponse) SetHeaders(v map[string]*string) *UpdateDIAlarmRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDIAlarmRuleResponse) SetStatusCode(v int32) *UpdateDIAlarmRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDIAlarmRuleResponse) SetBody(v *UpdateDIAlarmRuleResponseBody) *UpdateDIAlarmRuleResponse {
	s.Body = v
	return s
}

type UpdateDIJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId             *int64                                   `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description         *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	JobSettings         *UpdateDIJobRequestJobSettings           `json:"JobSettings,omitempty" xml:"JobSettings,omitempty" type:"Struct"`
	ProjectId           *int64                                   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceSettings    *UpdateDIJobRequestResourceSettings      `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty" type:"Struct"`
	TableMappings       []*UpdateDIJobRequestTableMappings       `json:"TableMappings,omitempty" xml:"TableMappings,omitempty" type:"Repeated"`
	TransformationRules []*UpdateDIJobRequestTransformationRules `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequest) SetDIJobId(v int64) *UpdateDIJobRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobRequest) SetDescription(v string) *UpdateDIJobRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobRequest) SetJobSettings(v *UpdateDIJobRequestJobSettings) *UpdateDIJobRequest {
	s.JobSettings = v
	return s
}

func (s *UpdateDIJobRequest) SetProjectId(v int64) *UpdateDIJobRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDIJobRequest) SetResourceSettings(v *UpdateDIJobRequestResourceSettings) *UpdateDIJobRequest {
	s.ResourceSettings = v
	return s
}

func (s *UpdateDIJobRequest) SetTableMappings(v []*UpdateDIJobRequestTableMappings) *UpdateDIJobRequest {
	s.TableMappings = v
	return s
}

func (s *UpdateDIJobRequest) SetTransformationRules(v []*UpdateDIJobRequestTransformationRules) *UpdateDIJobRequest {
	s.TransformationRules = v
	return s
}

type UpdateDIJobRequestJobSettings struct {
	// example:
	//
	// {"structInfo":"MANAGED","storageType":"TEXTFILE","writeMode":"APPEND","partitionColumns":[{"columnName":"pt","columnType":"STRING","comment":""}],"fieldDelimiter":""}
	ChannelSettings        *string                                                `json:"ChannelSettings,omitempty" xml:"ChannelSettings,omitempty"`
	ColumnDataTypeSettings []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings `json:"ColumnDataTypeSettings,omitempty" xml:"ColumnDataTypeSettings,omitempty" type:"Repeated"`
	CycleScheduleSettings  *UpdateDIJobRequestJobSettingsCycleScheduleSettings    `json:"CycleScheduleSettings,omitempty" xml:"CycleScheduleSettings,omitempty" type:"Struct"`
	DdlHandlingSettings    []*UpdateDIJobRequestJobSettingsDdlHandlingSettings    `json:"DdlHandlingSettings,omitempty" xml:"DdlHandlingSettings,omitempty" type:"Repeated"`
	RuntimeSettings        []*UpdateDIJobRequestJobSettingsRuntimeSettings        `json:"RuntimeSettings,omitempty" xml:"RuntimeSettings,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequestJobSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettings) SetChannelSettings(v string) *UpdateDIJobRequestJobSettings {
	s.ChannelSettings = &v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetColumnDataTypeSettings(v []*UpdateDIJobRequestJobSettingsColumnDataTypeSettings) *UpdateDIJobRequestJobSettings {
	s.ColumnDataTypeSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetCycleScheduleSettings(v *UpdateDIJobRequestJobSettingsCycleScheduleSettings) *UpdateDIJobRequestJobSettings {
	s.CycleScheduleSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetDdlHandlingSettings(v []*UpdateDIJobRequestJobSettingsDdlHandlingSettings) *UpdateDIJobRequestJobSettings {
	s.DdlHandlingSettings = v
	return s
}

func (s *UpdateDIJobRequestJobSettings) SetRuntimeSettings(v []*UpdateDIJobRequestJobSettingsRuntimeSettings) *UpdateDIJobRequestJobSettings {
	s.RuntimeSettings = v
	return s
}

type UpdateDIJobRequestJobSettingsColumnDataTypeSettings struct {
	// example:
	//
	// text
	DestinationDataType *string `json:"DestinationDataType,omitempty" xml:"DestinationDataType,omitempty"`
	// example:
	//
	// bigint
	SourceDataType *string `json:"SourceDataType,omitempty" xml:"SourceDataType,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsColumnDataTypeSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsColumnDataTypeSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) SetDestinationDataType(v string) *UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.DestinationDataType = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsColumnDataTypeSettings) SetSourceDataType(v string) *UpdateDIJobRequestJobSettingsColumnDataTypeSettings {
	s.SourceDataType = &v
	return s
}

type UpdateDIJobRequestJobSettingsCycleScheduleSettings struct {
	// example:
	//
	// bizdate=$bizdate
	ScheduleParameters *string `json:"ScheduleParameters,omitempty" xml:"ScheduleParameters,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsCycleScheduleSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsCycleScheduleSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsCycleScheduleSettings) SetScheduleParameters(v string) *UpdateDIJobRequestJobSettingsCycleScheduleSettings {
	s.ScheduleParameters = &v
	return s
}

type UpdateDIJobRequestJobSettingsDdlHandlingSettings struct {
	// example:
	//
	// Critical
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// AddColumn
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsDdlHandlingSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsDdlHandlingSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) SetAction(v string) *UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Action = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsDdlHandlingSettings) SetType(v string) *UpdateDIJobRequestJobSettingsDdlHandlingSettings {
	s.Type = &v
	return s
}

type UpdateDIJobRequestJobSettingsRuntimeSettings struct {
	// example:
	//
	// runtime.offline.concurrent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDIJobRequestJobSettingsRuntimeSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestJobSettingsRuntimeSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) SetName(v string) *UpdateDIJobRequestJobSettingsRuntimeSettings {
	s.Name = &v
	return s
}

func (s *UpdateDIJobRequestJobSettingsRuntimeSettings) SetValue(v string) *UpdateDIJobRequestJobSettingsRuntimeSettings {
	s.Value = &v
	return s
}

type UpdateDIJobRequestResourceSettings struct {
	OfflineResourceSettings  *UpdateDIJobRequestResourceSettingsOfflineResourceSettings  `json:"OfflineResourceSettings,omitempty" xml:"OfflineResourceSettings,omitempty" type:"Struct"`
	RealtimeResourceSettings *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings `json:"RealtimeResourceSettings,omitempty" xml:"RealtimeResourceSettings,omitempty" type:"Struct"`
	ScheduleResourceSettings *UpdateDIJobRequestResourceSettingsScheduleResourceSettings `json:"ScheduleResourceSettings,omitempty" xml:"ScheduleResourceSettings,omitempty" type:"Struct"`
}

func (s UpdateDIJobRequestResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettings) SetOfflineResourceSettings(v *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.OfflineResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetRealtimeResourceSettings(v *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.RealtimeResourceSettings = v
	return s
}

func (s *UpdateDIJobRequestResourceSettings) SetScheduleResourceSettings(v *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) *UpdateDIJobRequestResourceSettings {
	s.ScheduleResourceSettings = v
	return s
}

type UpdateDIJobRequestResourceSettingsOfflineResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *int64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsOfflineResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsOfflineResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetRequestedCu(v int64) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsOfflineResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsOfflineResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type UpdateDIJobRequestResourceSettingsRealtimeResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *int64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_111_222
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetRequestedCu(v int64) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsRealtimeResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type UpdateDIJobRequestResourceSettingsScheduleResourceSettings struct {
	// example:
	//
	// 2.0
	RequestedCu *int64 `json:"RequestedCu,omitempty" xml:"RequestedCu,omitempty"`
	// example:
	//
	// S_res_group_235454102432001_1721021993437
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
}

func (s UpdateDIJobRequestResourceSettingsScheduleResourceSettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestResourceSettingsScheduleResourceSettings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) SetRequestedCu(v int64) *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.RequestedCu = &v
	return s
}

func (s *UpdateDIJobRequestResourceSettingsScheduleResourceSettings) SetResourceGroupIdentifier(v string) *UpdateDIJobRequestResourceSettingsScheduleResourceSettings {
	s.ResourceGroupIdentifier = &v
	return s
}

type UpdateDIJobRequestTableMappings struct {
	SourceObjectSelectionRules []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules `json:"SourceObjectSelectionRules,omitempty" xml:"SourceObjectSelectionRules,omitempty" type:"Repeated"`
	TransformationRules        []*UpdateDIJobRequestTableMappingsTransformationRules        `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty" type:"Repeated"`
}

func (s UpdateDIJobRequestTableMappings) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTableMappings) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappings) SetSourceObjectSelectionRules(v []*UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) *UpdateDIJobRequestTableMappings {
	s.SourceObjectSelectionRules = v
	return s
}

func (s *UpdateDIJobRequestTableMappings) SetTransformationRules(v []*UpdateDIJobRequestTableMappingsTransformationRules) *UpdateDIJobRequestTableMappings {
	s.TransformationRules = v
	return s
}

type UpdateDIJobRequestTableMappingsSourceObjectSelectionRules struct {
	// example:
	//
	// Include
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// mysql_table_1
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Exact
	ExpressionType *string `json:"ExpressionType,omitempty" xml:"ExpressionType,omitempty"`
	// example:
	//
	// Table
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetAction(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Action = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpression(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.Expression = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetExpressionType(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ExpressionType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules) SetObjectType(v string) *UpdateDIJobRequestTableMappingsSourceObjectSelectionRules {
	s.ObjectType = &v
	return s
}

type UpdateDIJobRequestTableMappingsTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s UpdateDIJobRequestTableMappingsTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTableMappingsTransformationRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleActionType(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleName(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleName = &v
	return s
}

func (s *UpdateDIJobRequestTableMappingsTransformationRules) SetRuleTargetType(v string) *UpdateDIJobRequestTableMappingsTransformationRules {
	s.RuleTargetType = &v
	return s
}

type UpdateDIJobRequestTransformationRules struct {
	// example:
	//
	// Rename
	RuleActionType *string `json:"RuleActionType,omitempty" xml:"RuleActionType,omitempty"`
	// example:
	//
	// {"expression":"${srcDatasoureName}_${srcDatabaseName}"}
	RuleExpression *string `json:"RuleExpression,omitempty" xml:"RuleExpression,omitempty"`
	// example:
	//
	// rename_rule_1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// Table
	RuleTargetType *string `json:"RuleTargetType,omitempty" xml:"RuleTargetType,omitempty"`
}

func (s UpdateDIJobRequestTransformationRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobRequestTransformationRules) GoString() string {
	return s.String()
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleActionType(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleActionType = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleExpression(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleExpression = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleName(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleName = &v
	return s
}

func (s *UpdateDIJobRequestTransformationRules) SetRuleTargetType(v string) *UpdateDIJobRequestTransformationRules {
	s.RuleTargetType = &v
	return s
}

type UpdateDIJobShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11588
	DIJobId                   *int64  `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	JobSettingsShrink         *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	ProjectId                 *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceSettingsShrink    *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	TableMappingsShrink       *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s UpdateDIJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDIJobShrinkRequest) SetDIJobId(v int64) *UpdateDIJobShrinkRequest {
	s.DIJobId = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetDescription(v string) *UpdateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetJobSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetProjectId(v int64) *UpdateDIJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTableMappingsShrink(v string) *UpdateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *UpdateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *UpdateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

type UpdateDIJobResponseBody struct {
	// example:
	//
	// AAC30B35-820D-5F3E-A42C-E96BB6379325
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponseBody) SetRequestId(v string) *UpdateDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDIJobResponseBody) SetSuccess(v bool) *UpdateDIJobResponseBody {
	s.Success = &v
	return s
}

type UpdateDIJobResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDIJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDIJobResponse) SetHeaders(v map[string]*string) *UpdateDIJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDIJobResponse) SetStatusCode(v int32) *UpdateDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDIJobResponse) SetBody(v *UpdateDIJobResponseBody) *UpdateDIJobResponse {
	s.Body = v
	return s
}

type UpdateDataSourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5678
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) SetConnectionProperties(v string) *UpdateDataSourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *UpdateDataSourceRequest) SetConnectionPropertiesMode(v string) *UpdateDataSourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDescription(v string) *UpdateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceRequest) SetId(v int64) *UpdateDataSourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceRequest) SetProjectId(v int64) *UpdateDataSourceRequest {
	s.ProjectId = &v
	return s
}

type UpdateDataSourceResponseBody struct {
	// example:
	//
	// 102E8E24-0387-531D-8A75-1C0AE7DD03E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponseBody) SetRequestId(v string) *UpdateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSourceResponseBody) SetSuccess(v bool) *UpdateDataSourceResponseBody {
	s.Success = &v
	return s
}

type UpdateDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceResponse) SetHeaders(v map[string]*string) *UpdateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDataSourceResponse) SetStatusCode(v int32) *UpdateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDataSourceResponse) SetBody(v *UpdateDataSourceResponseBody) *UpdateDataSourceResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	// The ID of the UDF.
	//
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the UDF. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetId(v string) *UpdateFunctionRequest {
	s.Id = &v
	return s
}

func (s *UpdateFunctionRequest) SetProjectId(v string) *UpdateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFunctionRequest) SetSpec(v string) *UpdateFunctionRequest {
	s.Spec = &v
	return s
}

type UpdateFunctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 12123960-CB2C-5086-868E-C6C1D024XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSuccess(v bool) *UpdateFunctionResponseBody {
	s.Success = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetStatusCode(v int32) *UpdateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateNodeRequest struct {
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the node. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeRequest) SetId(v string) *UpdateNodeRequest {
	s.Id = &v
	return s
}

func (s *UpdateNodeRequest) SetProjectId(v string) *UpdateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateNodeRequest) SetSpec(v string) *UpdateNodeRequest {
	s.Spec = &v
	return s
}

type UpdateNodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponseBody) SetRequestId(v string) *UpdateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeResponseBody) SetSuccess(v bool) *UpdateNodeResponseBody {
	s.Success = &v
	return s
}

type UpdateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponse) SetHeaders(v map[string]*string) *UpdateNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeResponse) SetStatusCode(v int32) *UpdateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeResponse) SetBody(v *UpdateNodeResponseBody) *UpdateNodeResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// example:
	//
	// true
	DevRoleDisabled *bool   `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// example:
	//
	// Forbidden
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetDevEnvironmentEnabled(v bool) *UpdateProjectRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *UpdateProjectRequest) SetDevRoleDisabled(v bool) *UpdateProjectRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *UpdateProjectRequest) SetDisplayName(v string) *UpdateProjectRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateProjectRequest) SetId(v int64) *UpdateProjectRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectRequest) SetPaiTaskEnabled(v bool) *UpdateProjectRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *UpdateProjectRequest) SetStatus(v string) *UpdateProjectRequest {
	s.Status = &v
	return s
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	// The ID of the file resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the file resource. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetId(v string) *UpdateResourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceRequest) SetProjectId(v string) *UpdateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateResourceRequest) SetSpec(v string) *UpdateResourceRequest {
	s.Spec = &v
	return s
}

type UpdateResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetSuccess(v bool) *UpdateResourceResponseBody {
	s.Success = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
	s.Body = v
	return s
}

type UpdateResourceGroupRequest struct {
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name that you want to change for the resource group.
	//
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new remarks that you want to modify for the resource group.
	//
	// example:
	//
	// 创建用于普通任务的通用资源组
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) SetId(v string) *UpdateResourceGroupRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetName(v string) *UpdateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetRemark(v string) *UpdateResourceGroupRequest {
	s.Remark = &v
	return s
}

type UpdateResourceGroupResponseBody struct {
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponseBody) SetRequestId(v string) *UpdateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceGroupResponseBody) SetSuccess(v bool) *UpdateResourceGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupResponse) SetHeaders(v map[string]*string) *UpdateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceGroupResponse) SetStatusCode(v int32) *UpdateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceGroupResponse) SetBody(v *UpdateResourceGroupResponseBody) *UpdateResourceGroupResponse {
	s.Body = v
	return s
}

type UpdateRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/16
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateRouteRequest) SetDestinationCidr(v string) *UpdateRouteRequest {
	s.DestinationCidr = &v
	return s
}

func (s *UpdateRouteRequest) SetId(v int64) *UpdateRouteRequest {
	s.Id = &v
	return s
}

type UpdateRouteResponseBody struct {
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRouteResponseBody) SetRequestId(v string) *UpdateRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRouteResponseBody) SetSuccess(v bool) *UpdateRouteResponseBody {
	s.Success = &v
	return s
}

type UpdateRouteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateRouteResponse) SetHeaders(v map[string]*string) *UpdateRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateRouteResponse) SetStatusCode(v int32) *UpdateRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRouteResponse) SetBody(v *UpdateRouteResponseBody) *UpdateRouteResponse {
	s.Body = v
	return s
}

type UpdateWorkflowDefinitionRequest struct {
	// The ID of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The FlowSpec field information about the workflow. For more information, see [FlowSpec](https://github.com/aliyun/dataworks-spec/blob/master/README_zh_CN.md).
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionRequest) SetId(v string) *UpdateWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetProjectId(v string) *UpdateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetSpec(v string) *UpdateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type UpdateWorkflowDefinitionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20BF7E80-668A-5620-8AD8-879B8FEAXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponseBody) SetRequestId(v string) *UpdateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponseBody) SetSuccess(v bool) *UpdateWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetStatusCode(v int32) *UpdateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetBody(v *UpdateWorkflowDefinitionResponseBody) *UpdateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("dataworks.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("dataworks.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("dataworks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("dataworks.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("dataworks.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("dataworks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            tea.String("dataworks.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            tea.String("dataworks.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou":           tea.String("dataworks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("dataworks.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          tea.String("dataworks.aliyuncs.com"),
		"cn-qingdao":            tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai":           tea.String("dataworks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("dataworks.cn-shenzhen.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("dataworks.aliyuncs.com"),
		"eu-central-1":          tea.String("dataworks.eu-central-1.aliyuncs.com"),
		"eu-west-1":             tea.String("dataworks.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("dataworks.me-east-1.aliyuncs.com"),
		"us-east-1":             tea.String("dataworks.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("dataworks.us-west-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("dataworks.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("dataworks.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataworks-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates the process for deploying or undeploying an entity. The process is not deleted and can still be queried by calling query operations.
//
// @param request - AbolishDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishDeploymentResponse
func (client *Client) AbolishDeploymentWithOptions(request *AbolishDeploymentRequest, runtime *util.RuntimeOptions) (_result *AbolishDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AbolishDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbolishDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates the process for deploying or undeploying an entity. The process is not deleted and can still be queried by calling query operations.
//
// @param request - AbolishDeploymentRequest
//
// @return AbolishDeploymentResponse
func (client *Client) AbolishDeployment(request *AbolishDeploymentRequest) (_result *AbolishDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbolishDeploymentResponse{}
	_body, _err := client.AbolishDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a resource group with a workspace.
//
// @param request - AssociateProjectToResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateProjectToResourceGroupResponse
func (client *Client) AssociateProjectToResourceGroupWithOptions(request *AssociateProjectToResourceGroupRequest, runtime *util.RuntimeOptions) (_result *AssociateProjectToResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateProjectToResourceGroup"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateProjectToResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a resource group with a workspace.
//
// @param request - AssociateProjectToResourceGroupRequest
//
// @return AssociateProjectToResourceGroupResponse
func (client *Client) AssociateProjectToResourceGroup(request *AssociateProjectToResourceGroupRequest) (_result *AssociateProjectToResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateProjectToResourceGroupResponse{}
	_body, _err := client.AssociateProjectToResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CloneDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSourceWithOptions(request *CloneDataSourceRequest, runtime *util.RuntimeOptions) (_result *CloneDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloneDataSourceName)) {
		query["CloneDataSourceName"] = request.CloneDataSourceName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CloneDataSourceRequest
//
// @return CloneDataSourceResponse
func (client *Client) CloneDataSource(request *CloneDataSourceRequest) (_result *CloneDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneDataSourceResponse{}
	_body, _err := client.CloneDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集成报警规则
//
// @param tmpReq - CreateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRuleWithOptions(tmpReq *CreateDIAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDIAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NotificationSettings)) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, tea.String("NotificationSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TriggerConditions)) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, tea.String("TriggerConditions"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDIAlarmRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集成报警规则
//
// @param request - CreateDIAlarmRuleRequest
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRule(request *CreateDIAlarmRuleRequest) (_result *CreateDIAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CreateDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集成任务
//
// @param tmpReq - CreateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJobWithOptions(tmpReq *CreateDIJobRequest, runtime *util.RuntimeOptions) (_result *CreateDIJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DestinationDataSourceSettings)) {
		request.DestinationDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationDataSourceSettings, tea.String("DestinationDataSourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.JobSettings)) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, tea.String("JobSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSettings)) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, tea.String("ResourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceDataSourceSettings)) {
		request.SourceDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceDataSourceSettings, tea.String("SourceDataSourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TableMappings)) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, tea.String("TableMappings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TransformationRules)) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, tea.String("TransformationRules"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集成任务
//
// @param request - CreateDIJobRequest
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJob(request *CreateDIJobRequest) (_result *CreateDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CreateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionProperties)) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPropertiesMode)) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRuleWithOptions(request *CreateDataSourceSharedRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["EnvType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.SharedUser)) {
		query["SharedUser"] = request.SharedUser
	}

	if !tea.BoolValue(util.IsUnset(request.TargetProjectId)) {
		query["TargetProjectId"] = request.TargetProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSourceSharedRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - CreateDataSourceSharedRuleRequest
//
// @return CreateDataSourceSharedRuleResponse
func (client *Client) CreateDataSourceSharedRule(request *CreateDataSourceSharedRuleRequest) (_result *CreateDataSourceSharedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDataSourceSharedRuleResponse{}
	_body, _err := client.CreateDataSourceSharedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a process for deploying or undeploying an entity in DataStudio.
//
// Description:
//
// >  You cannot use this API operation to create a process for multiple entities at a time. If you specify multiple entities in a request, the system creates a process only for the first entity.
//
// @param tmpReq - CreateDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeploymentWithOptions(tmpReq *CreateDeploymentRequest, runtime *util.RuntimeOptions) (_result *CreateDeploymentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDeploymentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ObjectIds)) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, tea.String("ObjectIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectIdsShrink)) {
		body["ObjectIds"] = request.ObjectIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a process for deploying or undeploying an entity in DataStudio.
//
// Description:
//
// >  You cannot use this API operation to create a process for multiple entities at a time. If you specify multiple entities in a request, the system creates a process only for the first entity.
//
// @param request - CreateDeploymentRequest
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeployment(request *CreateDeploymentRequest) (_result *CreateDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CreateDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user-defined function (UDF) in DataStudio. The information about the UDF is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple UDFs at a time. If you specify multiple UDFs by using FlowSpec, the system creates only the first specified UDF.
//
// @param request - CreateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user-defined function (UDF) in DataStudio. The information about the UDF is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple UDFs at a time. If you specify multiple UDFs by using FlowSpec, the system creates only the first specified UDF.
//
// @param request - CreateFunctionRequest
//
// @return CreateFunctionResponse
func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建并绑定通用资源组网络资源。
//
// @param request - CreateNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkResponse
func (client *Client) CreateNetworkWithOptions(request *CreateNetworkRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetwork"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并绑定通用资源组网络资源。
//
// @param request - CreateNetworkRequest
//
// @return CreateNetworkResponse
func (client *Client) CreateNetwork(request *CreateNetworkRequest) (_result *CreateNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkResponse{}
	_body, _err := client.CreateNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a node in DataStudio. The information about the node is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple nodes at a time. If you specify multiple nodes by using FlowSpec, the system creates only the first specified node.
//
// @param request - CreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeResponse
func (client *Client) CreateNodeWithOptions(request *CreateNodeRequest, runtime *util.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerId)) {
		body["ContainerId"] = request.ContainerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node in DataStudio. The information about the node is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple nodes at a time. If you specify multiple nodes by using FlowSpec, the system creates only the first specified node.
//
// @param request - CreateNodeRequest
//
// @return CreateNodeResponse
func (client *Client) CreateNode(request *CreateNodeRequest) (_result *CreateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeResponse{}
	_body, _err := client.CreateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AliyunResourceTags)) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, tea.String("AliyunResourceTags"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunResourceGroupId)) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunResourceTagsShrink)) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DevEnvironmentEnabled)) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DevRoleDisabled)) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PaiTaskEnabled)) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加工作空间成员
//
// @param tmpReq - CreateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMemberWithOptions(tmpReq *CreateProjectMemberRequest, runtime *util.RuntimeOptions) (_result *CreateProjectMemberResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProjectMemberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoleCodes)) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, tea.String("RoleCodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleCodesShrink)) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProjectMember"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加工作空间成员
//
// @param request - CreateProjectMemberRequest
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMember(request *CreateProjectMemberRequest) (_result *CreateProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CreateProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file resource in DataStudio. The information about the file resource is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple file resources at a time. If you specify multiple file resources by using FlowSpec, the system creates only the first specified resource.
//
// @param request - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file resource in DataStudio. The information about the file resource is described by using FlowSpec.
//
// Description:
//
// >  You cannot use this API operation to create multiple file resources at a time. If you specify multiple file resources by using FlowSpec, the system creates only the first specified resource.
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建通用资源组。
//
// @param request - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDuration)) {
		body["PaymentDuration"] = request.PaymentDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentDurationUnit)) {
		body["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		body["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceGroup"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建通用资源组。
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建网络资源的路由。
//
// @param request - CreateRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteResponse
func (client *Client) CreateRouteWithOptions(request *CreateRouteRequest, runtime *util.RuntimeOptions) (_result *CreateRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidr)) {
		body["DestinationCidr"] = request.DestinationCidr
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkId)) {
		body["NetworkId"] = request.NetworkId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRoute"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建网络资源的路由。
//
// @param request - CreateRouteRequest
//
// @return CreateRouteResponse
func (client *Client) CreateRoute(request *CreateRouteRequest) (_result *CreateRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRouteResponse{}
	_body, _err := client.CreateRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow in a directory of DataStudio.
//
// Description:
//
// > You cannot use this API operation to create multiple workflows at a time. If you specify multiple workflows by using FlowSpec, the system creates only the first specified workflow. Other specified workflows and the nodes in the workflows are ignored. You can call the CreateNode operation to create a node.
//
// @param request - CreateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinitionWithOptions(request *CreateWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *CreateWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow in a directory of DataStudio.
//
// Description:
//
// > You cannot use this API operation to create multiple workflows at a time. If you specify multiple workflows by using FlowSpec, the system creates only the first specified workflow. Other specified workflows and the nodes in the workflows are ignored. You can call the CreateNode operation to create a node.
//
// @param request - CreateWorkflowDefinitionRequest
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinition(request *CreateWorkflowDefinitionRequest) (_result *CreateWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CreateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert rule configured for a synchronization task.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRuleWithOptions(request *DeleteDIAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDIAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDIAlarmRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule configured for a synchronization task.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRule(request *DeleteDIAlarmRuleRequest) (_result *DeleteDIAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.DeleteDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集成任务
//
// @param request - DeleteDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJobWithOptions(request *DeleteDIJobRequest, runtime *util.RuntimeOptions) (_result *DeleteDIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集成任务
//
// @param request - DeleteDIJobRequest
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJob(request *DeleteDIJobRequest) (_result *DeleteDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.DeleteDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRuleWithOptions(request *DeleteDataSourceSharedRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSourceSharedRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - DeleteDataSourceSharedRuleRequest
//
// @return DeleteDataSourceSharedRuleResponse
func (client *Client) DeleteDataSourceSharedRule(request *DeleteDataSourceSharedRuleRequest) (_result *DeleteDataSourceSharedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDataSourceSharedRuleResponse{}
	_body, _err := client.DeleteDataSourceSharedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function (UDF) in DataStudio.
//
// Description:
//
// >  A UDF that is deployed cannot be deleted. If you want to delete such a UDF, you must first undeploy the UDF.
//
// @param request - DeleteFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function (UDF) in DataStudio.
//
// Description:
//
// >  A UDF that is deployed cannot be deleted. If you want to delete such a UDF, you must first undeploy the UDF.
//
// @param request - DeleteFunctionRequest
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑并删除通用资源组网络资源。
//
// @param request - DeleteNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkResponse
func (client *Client) DeleteNetworkWithOptions(request *DeleteNetworkRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNetwork"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑并删除通用资源组网络资源。
//
// @param request - DeleteNetworkRequest
//
// @return DeleteNetworkResponse
func (client *Client) DeleteNetwork(request *DeleteNetworkRequest) (_result *DeleteNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkResponse{}
	_body, _err := client.DeleteNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a node from DataStudio.
//
// Description:
//
// >  A node that is deployed cannot be deleted. If you want to delete such a node, you must first undeploy the node.
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithOptions(request *DeleteNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node from DataStudio.
//
// Description:
//
// >  A node that is deployed cannot be deleted. If you want to delete such a node, you must first undeploy the node.
//
// @param request - DeleteNodeRequest
//
// @return DeleteNodeResponse
func (client *Client) DeleteNode(request *DeleteNodeRequest) (_result *DeleteNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeResponse{}
	_body, _err := client.DeleteNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 销毁工作空间
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 销毁工作空间
//
// @param request - DeleteProjectRequest
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除工作空间成员
//
// @param request - DeleteProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMemberWithOptions(request *DeleteProjectMemberRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProjectMember"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除工作空间成员
//
// @param request - DeleteProjectMemberRequest
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMember(request *DeleteProjectMemberRequest) (_result *DeleteProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.DeleteProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file resource from DataStudio.
//
// Description:
//
// >  A file resource that is deployed cannot be deleted. If you want to delete such a file resource, you must first undeploy the file resource.
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file resource from DataStudio.
//
// Description:
//
// >  A file resource that is deployed cannot be deleted. If you want to delete such a file resource, you must first undeploy the file resource.
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a serverless resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceGroup"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a serverless resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网络资源的路由。
//
// @param request - DeleteRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteResponse
func (client *Client) DeleteRouteWithOptions(request *DeleteRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRoute"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网络资源的路由。
//
// @param request - DeleteRouteRequest
//
// @return DeleteRouteResponse
func (client *Client) DeleteRoute(request *DeleteRouteRequest) (_result *DeleteRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRouteResponse{}
	_body, _err := client.DeleteRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workflow from DataStudio.
//
// Description:
//
// >  A workflow that is deployed cannot be deleted. If you want to delete such a workflow, you must first undeploy the workflow.
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinitionWithOptions(request *DeleteWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workflow from DataStudio.
//
// Description:
//
// >  A workflow that is deployed cannot be deleted. If you want to delete such a workflow, you must first undeploy the workflow.
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinition(request *DeleteWorkflowDefinitionRequest) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.DeleteWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a resource group from a workspace.
//
// @param request - DissociateProjectFromResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateProjectFromResourceGroupResponse
func (client *Client) DissociateProjectFromResourceGroupWithOptions(request *DissociateProjectFromResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DissociateProjectFromResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateProjectFromResourceGroup"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateProjectFromResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a resource group from a workspace.
//
// @param request - DissociateProjectFromResourceGroupRequest
//
// @return DissociateProjectFromResourceGroupResponse
func (client *Client) DissociateProjectFromResourceGroup(request *DissociateProjectFromResourceGroupRequest) (_result *DissociateProjectFromResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DissociateProjectFromResourceGroupResponse{}
	_body, _err := client.DissociateProjectFromResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a stage in a process.
//
// Description:
//
// >  The stages in a process are sequential. For more information, see the GetDeployment operation. Skipping or repeating a stage is not allowed.
//
// >  The execution of a stage is asynchronous. The response of this operation indicates only whether a stage is triggered but does not indicate whether the execution of the stage is successful. You can call the GetDeployment operation to check whether the execution is successful.
//
// @param request - ExecDeploymentStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecDeploymentStageResponse
func (client *Client) ExecDeploymentStageWithOptions(request *ExecDeploymentStageRequest, runtime *util.RuntimeOptions) (_result *ExecDeploymentStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecDeploymentStage"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecDeploymentStageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a stage in a process.
//
// Description:
//
// >  The stages in a process are sequential. For more information, see the GetDeployment operation. Skipping or repeating a stage is not allowed.
//
// >  The execution of a stage is asynchronous. The response of this operation indicates only whether a stage is triggered but does not indicate whether the execution of the stage is successful. You can call the GetDeployment operation to check whether the execution is successful.
//
// @param request - ExecDeploymentStageRequest
//
// @return ExecDeploymentStageResponse
func (client *Client) ExecDeploymentStage(request *ExecDeploymentStageRequest) (_result *ExecDeploymentStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecDeploymentStageResponse{}
	_body, _err := client.ExecDeploymentStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看数据集成任务
//
// @param request - GetDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobResponse
func (client *Client) GetDIJobWithOptions(request *GetDIJobRequest, runtime *util.RuntimeOptions) (_result *GetDIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据集成任务
//
// @param request - GetDIJobRequest
//
// @return GetDIJobResponse
func (client *Client) GetDIJob(request *GetDIJobRequest) (_result *GetDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDIJobResponse{}
	_body, _err := client.GetDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains logs generated for a synchronization task.
//
// @param request - GetDIJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLogWithOptions(request *GetDIJobLogRequest, runtime *util.RuntimeOptions) (_result *GetDIJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDIJobLog"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains logs generated for a synchronization task.
//
// @param request - GetDIJobLogRequest
//
// @return GetDIJobLogResponse
func (client *Client) GetDIJobLog(request *GetDIJobLogRequest) (_result *GetDIJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDIJobLogResponse{}
	_body, _err := client.GetDIJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a data source by ID.
//
// Description:
//
// You can call this operation only if you are assigned one of the following roles in DataWorks:
//
// 	- Tenant Owner, Workspace Administrator, Deployment, Development, Project Owner, and O\\&M
//
// @param request - GetDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceResponse
func (client *Client) GetDataSourceWithOptions(request *GetDataSourceRequest, runtime *util.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a data source by ID.
//
// Description:
//
// You can call this operation only if you are assigned one of the following roles in DataWorks:
//
// 	- Tenant Owner, Workspace Administrator, Deployment, Development, Project Owner, and O\\&M
//
// @param request - GetDataSourceRequest
//
// @return GetDataSourceResponse
func (client *Client) GetDataSource(request *GetDataSourceRequest) (_result *GetDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataSourceResponse{}
	_body, _err := client.GetDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a process for deploying or undeploying an entity.
//
// @param request - GetDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithOptions(request *GetDeploymentRequest, runtime *util.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a process for deploying or undeploying an entity.
//
// @param request - GetDeploymentRequest
//
// @return GetDeploymentResponse
func (client *Client) GetDeployment(request *GetDeploymentRequest) (_result *GetDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeploymentResponse{}
	_body, _err := client.GetDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a user-defined function (UDF) in DataStudio.
//
// @param request - GetFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithOptions(request *GetFunctionRequest, runtime *util.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user-defined function (UDF) in DataStudio.
//
// @param request - GetFunctionRequest
//
// @return GetFunctionResponse
func (client *Client) GetFunction(request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 返回异步任务的状态信息
//
// @param request - GetJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobStatusResponse
func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobStatus"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 返回异步任务的状态信息
//
// @param request - GetJobStatusRequest
//
// @return GetJobStatusResponse
func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某个网络资源详细信息。
//
// @param request - GetNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkResponse
func (client *Client) GetNetworkWithOptions(request *GetNetworkRequest, runtime *util.RuntimeOptions) (_result *GetNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNetwork"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNetworkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个网络资源详细信息。
//
// @param request - GetNetworkRequest
//
// @return GetNetworkResponse
func (client *Client) GetNetwork(request *GetNetworkRequest) (_result *GetNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNetworkResponse{}
	_body, _err := client.GetNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a node in DataStudio.
//
// @param request - GetNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithOptions(request *GetNodeRequest, runtime *util.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a node in DataStudio.
//
// @param request - GetNodeRequest
//
// @return GetNodeResponse
func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作空间详情
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作空间详情
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作空间成员详情
//
// @param request - GetProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectMemberResponse
func (client *Client) GetProjectMemberWithOptions(request *GetProjectMemberRequest, runtime *util.RuntimeOptions) (_result *GetProjectMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectMember"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作空间成员详情
//
// @param request - GetProjectMemberRequest
//
// @return GetProjectMemberResponse
func (client *Client) GetProjectMember(request *GetProjectMemberRequest) (_result *GetProjectMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectMemberResponse{}
	_body, _err := client.GetProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作空间角色详情
//
// @param request - GetProjectRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectRoleResponse
func (client *Client) GetProjectRoleWithOptions(request *GetProjectRoleRequest, runtime *util.RuntimeOptions) (_result *GetProjectRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectRole"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作空间角色详情
//
// @param request - GetProjectRoleRequest
//
// @return GetProjectRoleResponse
func (client *Client) GetProjectRole(request *GetProjectRoleRequest) (_result *GetProjectRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectRoleResponse{}
	_body, _err := client.GetProjectRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file resource.
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file resource.
//
// @param request - GetResourceRequest
//
// @return GetResourceResponse
func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据id获取指定资源组。
//
// @param request - GetResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroupWithOptions(request *GetResourceGroupRequest, runtime *util.RuntimeOptions) (_result *GetResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceGroup"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据id获取指定资源组。
//
// @param request - GetResourceGroupRequest
//
// @return GetResourceGroupResponse
func (client *Client) GetResourceGroup(request *GetResourceGroupRequest) (_result *GetResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceGroupResponse{}
	_body, _err := client.GetResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据id获取指定路由。
//
// @param request - GetRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRouteResponse
func (client *Client) GetRouteWithOptions(request *GetRouteRequest, runtime *util.RuntimeOptions) (_result *GetRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoute"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据id获取指定路由。
//
// @param request - GetRouteRequest
//
// @return GetRouteResponse
func (client *Client) GetRoute(request *GetRouteRequest) (_result *GetRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRouteResponse{}
	_body, _err := client.GetRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the infomation about a workflow.
//
// @param request - GetWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinitionWithOptions(request *GetWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *GetWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the infomation about a workflow.
//
// @param request - GetWorkflowDefinitionRequest
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinition(request *GetWorkflowDefinitionRequest) (_result *GetWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.GetWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授予工作空间成员角色
//
// @param tmpReq - GrantMemberProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantMemberProjectRolesResponse
func (client *Client) GrantMemberProjectRolesWithOptions(tmpReq *GrantMemberProjectRolesRequest, runtime *util.RuntimeOptions) (_result *GrantMemberProjectRolesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GrantMemberProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoleCodes)) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, tea.String("RoleCodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleCodesShrink)) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantMemberProjectRoles"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantMemberProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授予工作空间成员角色
//
// @param request - GrantMemberProjectRolesRequest
//
// @return GrantMemberProjectRolesResponse
func (client *Client) GrantMemberProjectRoles(request *GrantMemberProjectRolesRequest) (_result *GrantMemberProjectRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantMemberProjectRolesResponse{}
	_body, _err := client.GrantMemberProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用此接口，可以将通过FlowSpec定义的工作流节点和其内部的子节点都导入到数据开发中
//
// @param request - ImportWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportWorkflowDefinitionResponse
func (client *Client) ImportWorkflowDefinitionWithOptions(request *ImportWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *ImportWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用此接口，可以将通过FlowSpec定义的工作流节点和其内部的子节点都导入到数据开发中
//
// @param request - ImportWorkflowDefinitionRequest
//
// @return ImportWorkflowDefinitionResponse
func (client *Client) ImportWorkflowDefinition(request *ImportWorkflowDefinitionRequest) (_result *ImportWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportWorkflowDefinitionResponse{}
	_body, _err := client.ImportWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看数据集成报警规则
//
// @param request - ListDIAlarmRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRulesWithOptions(request *ListDIAlarmRulesRequest, runtime *util.RuntimeOptions) (_result *ListDIAlarmRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIAlarmRules"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据集成报警规则
//
// @param request - ListDIAlarmRulesRequest
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRules(request *ListDIAlarmRulesRequest) (_result *ListDIAlarmRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.ListDIAlarmRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成任务事件
//
// @param request - ListDIJobEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEventsWithOptions(request *ListDIJobEventsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobEvents"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成任务事件
//
// @param request - ListDIJobEventsRequest
//
// @return ListDIJobEventsResponse
func (client *Client) ListDIJobEvents(request *ListDIJobEventsRequest) (_result *ListDIJobEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobEventsResponse{}
	_body, _err := client.ListDIJobEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成任务指标
//
// @param tmpReq - ListDIJobMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetricsWithOptions(tmpReq *ListDIJobMetricsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobMetricsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDIJobMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MetricName)) {
		request.MetricNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricName, tea.String("MetricName"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobMetrics"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成任务指标
//
// @param request - ListDIJobMetricsRequest
//
// @return ListDIJobMetricsResponse
func (client *Client) ListDIJobMetrics(request *ListDIJobMetricsRequest) (_result *ListDIJobMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobMetricsResponse{}
	_body, _err := client.ListDIJobMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集成运行信息
//
// @param request - ListDIJobRunDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobRunDetailsResponse
func (client *Client) ListDIJobRunDetailsWithOptions(request *ListDIJobRunDetailsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobRunDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobRunDetails"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobRunDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集成运行信息
//
// @param request - ListDIJobRunDetailsRequest
//
// @return ListDIJobRunDetailsResponse
func (client *Client) ListDIJobRunDetails(request *ListDIJobRunDetailsRequest) (_result *ListDIJobRunDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobRunDetailsResponse{}
	_body, _err := client.ListDIJobRunDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of synchronization tasks.
//
// @param request - ListDIJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobsWithOptions(request *ListDIJobsRequest, runtime *util.RuntimeOptions) (_result *ListDIJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDIJobs"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDIJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of synchronization tasks.
//
// @param request - ListDIJobsRequest
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobs(request *ListDIJobsRequest) (_result *ListDIJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDIJobsResponse{}
	_body, _err := client.ListDIJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ListDataQualityEvaluationTaskInstances
//
// @param request - ListDataQualityEvaluationTaskInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityEvaluationTaskInstancesResponse
func (client *Client) ListDataQualityEvaluationTaskInstancesWithOptions(request *ListDataQualityEvaluationTaskInstancesRequest, runtime *util.RuntimeOptions) (_result *ListDataQualityEvaluationTaskInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataQualityEvaluationTaskInstances"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataQualityEvaluationTaskInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ListDataQualityEvaluationTaskInstances
//
// @param request - ListDataQualityEvaluationTaskInstancesRequest
//
// @return ListDataQualityEvaluationTaskInstancesResponse
func (client *Client) ListDataQualityEvaluationTaskInstances(request *ListDataQualityEvaluationTaskInstancesRequest) (_result *ListDataQualityEvaluationTaskInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataQualityEvaluationTaskInstancesResponse{}
	_body, _err := client.ListDataQualityEvaluationTaskInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDataQualityEvaluationTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityEvaluationTasksResponse
func (client *Client) ListDataQualityEvaluationTasksWithOptions(request *ListDataQualityEvaluationTasksRequest, runtime *util.RuntimeOptions) (_result *ListDataQualityEvaluationTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataQualityEvaluationTasks"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataQualityEvaluationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDataQualityEvaluationTasksRequest
//
// @return ListDataQualityEvaluationTasksResponse
func (client *Client) ListDataQualityEvaluationTasks(request *ListDataQualityEvaluationTasksRequest) (_result *ListDataQualityEvaluationTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataQualityEvaluationTasksResponse{}
	_body, _err := client.ListDataQualityEvaluationTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListDataQualityResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityResultsResponse
func (client *Client) ListDataQualityResultsWithOptions(request *ListDataQualityResultsRequest, runtime *util.RuntimeOptions) (_result *ListDataQualityResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataQualityResults"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataQualityResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListDataQualityResultsRequest
//
// @return ListDataQualityResultsResponse
func (client *Client) ListDataQualityResults(request *ListDataQualityResultsRequest) (_result *ListDataQualityResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataQualityResultsResponse{}
	_body, _err := client.ListDataQualityResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 质量监控规则分页查询
//
// @param request - ListDataQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataQualityRulesResponse
func (client *Client) ListDataQualityRulesWithOptions(request *ListDataQualityRulesRequest, runtime *util.RuntimeOptions) (_result *ListDataQualityRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataQualityRules"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataQualityRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 质量监控规则分页查询
//
// @param request - ListDataQualityRulesRequest
//
// @return ListDataQualityRulesResponse
func (client *Client) ListDataQualityRules(request *ListDataQualityRulesRequest) (_result *ListDataQualityRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataQualityRulesResponse{}
	_body, _err := client.ListDataQualityRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - ListDataSourceSharedRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRulesWithOptions(request *ListDataSourceSharedRulesRequest, runtime *util.RuntimeOptions) (_result *ListDataSourceSharedRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceSharedRules"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - ListDataSourceSharedRulesRequest
//
// @return ListDataSourceSharedRulesResponse
func (client *Client) ListDataSourceSharedRules(request *ListDataSourceSharedRulesRequest) (_result *ListDataSourceSharedRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourceSharedRulesResponse{}
	_body, _err := client.ListDataSourceSharedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param tmpReq - ListDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(tmpReq *ListDataSourcesRequest, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Types)) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, tea.String("Types"), tea.String("simple"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - ListDataSourcesRequest
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of processes that are used to deploy or undeploy entities in DataStudio. You can also specify filter conditions to query specific processes.
//
// @param request - ListDeploymentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithOptions(request *ListDeploymentsRequest, runtime *util.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployments"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of processes that are used to deploy or undeploy entities in DataStudio. You can also specify filter conditions to query specific processes.
//
// @param request - ListDeploymentsRequest
//
// @return ListDeploymentsResponse
func (client *Client) ListDeployments(request *ListDeploymentsRequest) (_result *ListDeploymentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.ListDeploymentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of user-defined functions (UDFs) in DataStudio. You can also specify filter conditions to query specific UDFs.
//
// @param request - ListFunctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of user-defined functions (UDFs) in DataStudio. You can also specify filter conditions to query specific UDFs.
//
// @param request - ListFunctionsRequest
//
// @return ListFunctionsResponse
func (client *Client) ListFunctions(request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取通用资源组网络资源列表。
//
// @param request - ListNetworksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworksResponse
func (client *Client) ListNetworksWithOptions(request *ListNetworksRequest, runtime *util.RuntimeOptions) (_result *ListNetworksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetworks"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetworksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取通用资源组网络资源列表。
//
// @param request - ListNetworksRequest
//
// @return ListNetworksResponse
func (client *Client) ListNetworks(request *ListNetworksRequest) (_result *ListNetworksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworksResponse{}
	_body, _err := client.ListNetworksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of descendant nodes of a node in DataStudio.
//
// @param request - ListNodeDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependenciesWithOptions(request *ListNodeDependenciesRequest, runtime *util.RuntimeOptions) (_result *ListNodeDependenciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeDependencies"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of descendant nodes of a node in DataStudio.
//
// @param request - ListNodeDependenciesRequest
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependencies(request *ListNodeDependenciesRequest) (_result *ListNodeDependenciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.ListNodeDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in DataStudio. You can also specify filter conditions to query specific nodes.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of nodes in DataStudio. You can also specify filter conditions to query specific nodes.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询工作空间成员详情
//
// @param tmpReq - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithOptions(tmpReq *ListProjectMembersRequest, runtime *util.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoleCodes)) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, tea.String("RoleCodes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserIds)) {
		request.UserIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserIds, tea.String("UserIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleCodesShrink)) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdsShrink)) {
		body["UserIds"] = request.UserIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectMembers"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询工作空间成员详情
//
// @param request - ListProjectMembersRequest
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembers(request *ListProjectMembersRequest) (_result *ListProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.ListProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询工作空间角色详情
//
// @param tmpReq - ListProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRolesWithOptions(tmpReq *ListProjectRolesRequest, runtime *util.RuntimeOptions) (_result *ListProjectRolesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Codes)) {
		request.CodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Codes, tea.String("Codes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Names)) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, tea.String("Names"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodesShrink)) {
		body["Codes"] = request.CodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NamesShrink)) {
		body["Names"] = request.NamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectRoles"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询工作空间角色详情
//
// @param request - ListProjectRolesRequest
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRoles(request *ListProjectRolesRequest) (_result *ListProjectRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.ListProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询工作空间详情
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AliyunResourceTags)) {
		request.AliyunResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AliyunResourceTags, tea.String("AliyunResourceTags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Ids)) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, tea.String("Ids"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Names)) {
		request.NamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Names, tea.String("Names"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunResourceGroupId)) {
		body["AliyunResourceGroupId"] = request.AliyunResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AliyunResourceTagsShrink)) {
		body["AliyunResourceTags"] = request.AliyunResourceTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DevEnvironmentEnabled)) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DevRoleDisabled)) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !tea.BoolValue(util.IsUnset(request.IdsShrink)) {
		body["Ids"] = request.IdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NamesShrink)) {
		body["Names"] = request.NamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PaiTaskEnabled)) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询工作空间详情
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源组列表。
//
// @param tmpReq - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(tmpReq *ListResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListResourceGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceGroupTypes)) {
		request.ResourceGroupTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceGroupTypes, tea.String("ResourceGroupTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Statuses)) {
		request.StatusesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Statuses, tea.String("Statuses"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceGroups"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源组列表。
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of file resources in DataStudio. You can also specify filter conditions to query specific file resources.
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of file resources in DataStudio. You can also specify filter conditions to query specific file resources.
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网络资源的路由列表。
//
// @param request - ListRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoutesResponse
func (client *Client) ListRoutesWithOptions(request *ListRoutesRequest, runtime *util.RuntimeOptions) (_result *ListRoutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoutes"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网络资源的路由列表。
//
// @param request - ListRoutesRequest
//
// @return ListRoutesResponse
func (client *Client) ListRoutes(request *ListRoutesRequest) (_result *ListRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRoutesResponse{}
	_body, _err := client.ListRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workflows in DataStudio. You can also specify filter conditions to query specific workflows.
//
// @param request - ListWorkflowDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitionsWithOptions(request *ListWorkflowDefinitionsRequest, runtime *util.RuntimeOptions) (_result *ListWorkflowDefinitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkflowDefinitions"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows in DataStudio. You can also specify filter conditions to query specific workflows.
//
// @param request - ListWorkflowDefinitionsRequest
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitions(request *ListWorkflowDefinitionsRequest) (_result *ListWorkflowDefinitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.ListWorkflowDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a user-defined function (UDF) to a path in DataStudio.
//
// @param request - MoveFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFunctionResponse
func (client *Client) MoveFunctionWithOptions(request *MoveFunctionRequest, runtime *util.RuntimeOptions) (_result *MoveFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a user-defined function (UDF) to a path in DataStudio.
//
// @param request - MoveFunctionRequest
//
// @return MoveFunctionResponse
func (client *Client) MoveFunction(request *MoveFunctionRequest) (_result *MoveFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveFunctionResponse{}
	_body, _err := client.MoveFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a node to a path in DataStudio.
//
// @param request - MoveNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveNodeResponse
func (client *Client) MoveNodeWithOptions(request *MoveNodeRequest, runtime *util.RuntimeOptions) (_result *MoveNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a node to a path in DataStudio.
//
// @param request - MoveNodeRequest
//
// @return MoveNodeResponse
func (client *Client) MoveNode(request *MoveNodeRequest) (_result *MoveNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveNodeResponse{}
	_body, _err := client.MoveNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a file resource to a path in DataStudio.
//
// @param request - MoveResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceResponse
func (client *Client) MoveResourceWithOptions(request *MoveResourceRequest, runtime *util.RuntimeOptions) (_result *MoveResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a file resource to a path in DataStudio.
//
// @param request - MoveResourceRequest
//
// @return MoveResourceResponse
func (client *Client) MoveResource(request *MoveResourceRequest) (_result *MoveResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceResponse{}
	_body, _err := client.MoveResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a workflow to a path in DataStudio.
//
// @param request - MoveWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinitionWithOptions(request *MoveWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *MoveWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a workflow to a path in DataStudio.
//
// @param request - MoveWorkflowDefinitionRequest
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinition(request *MoveWorkflowDefinitionRequest) (_result *MoveWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.MoveWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a user-defined function (UDF) in DataStudio.
//
// @param request - RenameFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameFunctionResponse
func (client *Client) RenameFunctionWithOptions(request *RenameFunctionRequest, runtime *util.RuntimeOptions) (_result *RenameFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a user-defined function (UDF) in DataStudio.
//
// @param request - RenameFunctionRequest
//
// @return RenameFunctionResponse
func (client *Client) RenameFunction(request *RenameFunctionRequest) (_result *RenameFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameFunctionResponse{}
	_body, _err := client.RenameFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a node in DataStudio.
//
// @param request - RenameNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameNodeResponse
func (client *Client) RenameNodeWithOptions(request *RenameNodeRequest, runtime *util.RuntimeOptions) (_result *RenameNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a node in DataStudio.
//
// @param request - RenameNodeRequest
//
// @return RenameNodeResponse
func (client *Client) RenameNode(request *RenameNodeRequest) (_result *RenameNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameNodeResponse{}
	_body, _err := client.RenameNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a file resource in DataStudio.
//
// @param request - RenameResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameResourceResponse
func (client *Client) RenameResourceWithOptions(request *RenameResourceRequest, runtime *util.RuntimeOptions) (_result *RenameResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a file resource in DataStudio.
//
// @param request - RenameResourceRequest
//
// @return RenameResourceResponse
func (client *Client) RenameResource(request *RenameResourceRequest) (_result *RenameResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameResourceResponse{}
	_body, _err := client.RenameResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a workflow in DataStudio.
//
// @param request - RenameWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinitionWithOptions(request *RenameWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *RenameWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renames a workflow in DataStudio.
//
// @param request - RenameWorkflowDefinitionRequest
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinition(request *RenameWorkflowDefinitionRequest) (_result *RenameWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.RenameWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 撤销工作空间成员的角色
//
// @param tmpReq - RevokeMemberProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeMemberProjectRolesResponse
func (client *Client) RevokeMemberProjectRolesWithOptions(tmpReq *RevokeMemberProjectRolesRequest, runtime *util.RuntimeOptions) (_result *RevokeMemberProjectRolesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RevokeMemberProjectRolesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RoleCodes)) {
		request.RoleCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RoleCodes, tea.String("RoleCodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleCodesShrink)) {
		body["RoleCodes"] = request.RoleCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeMemberProjectRoles"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeMemberProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 撤销工作空间成员的角色
//
// @param request - RevokeMemberProjectRolesRequest
//
// @return RevokeMemberProjectRolesResponse
func (client *Client) RevokeMemberProjectRoles(request *RevokeMemberProjectRolesRequest) (_result *RevokeMemberProjectRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeMemberProjectRolesResponse{}
	_body, _err := client.RevokeMemberProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动数据集成任务
//
// @param tmpReq - StartDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDIJobResponse
func (client *Client) StartDIJobWithOptions(tmpReq *StartDIJobRequest, runtime *util.RuntimeOptions) (_result *StartDIJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StartDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RealtimeStartSettings)) {
		request.RealtimeStartSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeStartSettings, tea.String("RealtimeStartSettings"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动数据集成任务
//
// @param request - StartDIJobRequest
//
// @return StartDIJobResponse
func (client *Client) StartDIJob(request *StartDIJobRequest) (_result *StartDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDIJobResponse{}
	_body, _err := client.StartDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a synchronization task.
//
// @param request - StopDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDIJobResponse
func (client *Client) StopDIJobWithOptions(request *StopDIJobRequest, runtime *util.RuntimeOptions) (_result *StopDIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a synchronization task.
//
// @param request - StopDIJobRequest
//
// @return StopDIJobResponse
func (client *Client) StopDIJob(request *StopDIJobRequest) (_result *StopDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDIJobResponse{}
	_body, _err := client.StopDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集成报警规则
//
// @param tmpReq - UpdateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRuleWithOptions(tmpReq *UpdateDIAlarmRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateDIAlarmRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NotificationSettings)) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, tea.String("NotificationSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TriggerConditions)) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, tea.String("TriggerConditions"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDIAlarmRule"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集成报警规则
//
// @param request - UpdateDIAlarmRuleRequest
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRule(request *UpdateDIAlarmRuleRequest) (_result *UpdateDIAlarmRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.UpdateDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集成任务
//
// @param tmpReq - UpdateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJobWithOptions(tmpReq *UpdateDIJobRequest, runtime *util.RuntimeOptions) (_result *UpdateDIJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobSettings)) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, tea.String("JobSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceSettings)) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, tea.String("ResourceSettings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TableMappings)) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, tea.String("TableMappings"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TransformationRules)) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, tea.String("TransformationRules"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDIJob"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集成任务
//
// @param request - UpdateDIJobRequest
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJob(request *UpdateDIJobRequest) (_result *UpdateDIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.UpdateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithOptions(request *UpdateDataSourceRequest, runtime *util.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectionProperties)) {
		query["ConnectionProperties"] = request.ConnectionProperties
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionPropertiesMode)) {
		query["ConnectionPropertiesMode"] = request.ConnectionPropertiesMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataSource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证用
//
// @param request - UpdateDataSourceRequest
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSource(request *UpdateDataSourceRequest) (_result *UpdateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.UpdateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a user-defined function (UDF) in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a user-defined function (UDF) in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateFunctionRequest
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a node in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeResponse
func (client *Client) UpdateNodeWithOptions(request *UpdateNodeRequest, runtime *util.RuntimeOptions) (_result *UpdateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a node in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateNodeRequest
//
// @return UpdateNodeResponse
func (client *Client) UpdateNode(request *UpdateNodeRequest) (_result *UpdateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodeResponse{}
	_body, _err := client.UpdateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工作空间
//
// @param request - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DevEnvironmentEnabled)) {
		body["DevEnvironmentEnabled"] = request.DevEnvironmentEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DevRoleDisabled)) {
		body["DevRoleDisabled"] = request.DevRoleDisabled
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PaiTaskEnabled)) {
		body["PaiTaskEnabled"] = request.PaiTaskEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作空间
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a file resource in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(request *UpdateResourceRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a file resource in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates basic information about a resource group.
//
// @param request - UpdateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroupWithOptions(request *UpdateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceGroup"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates basic information about a resource group.
//
// @param request - UpdateResourceGroupRequest
//
// @return UpdateResourceGroupResponse
func (client *Client) UpdateResourceGroup(request *UpdateResourceGroupRequest) (_result *UpdateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceGroupResponse{}
	_body, _err := client.UpdateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网络资源的路由。
//
// @param request - UpdateRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteResponse
func (client *Client) UpdateRouteWithOptions(request *UpdateRouteRequest, runtime *util.RuntimeOptions) (_result *UpdateRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCidr)) {
		body["DestinationCidr"] = request.DestinationCidr
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoute"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网络资源的路由。
//
// @param request - UpdateRouteRequest
//
// @return UpdateRouteResponse
func (client *Client) UpdateRoute(request *UpdateRouteRequest) (_result *UpdateRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRouteResponse{}
	_body, _err := client.UpdateRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinitionWithOptions(request *UpdateWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a workflow in DataStudio. This API operation performs an incremental update. The update information is described by using FlowSpec.
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinition(request *UpdateWorkflowDefinitionRequest) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.UpdateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
