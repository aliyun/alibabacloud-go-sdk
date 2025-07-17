// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityEvaluationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataQualityEvaluationTasksResponseBodyPagingInfo) *ListDataQualityEvaluationTasksResponseBody
	GetPagingInfo() *ListDataQualityEvaluationTasksResponseBodyPagingInfo
	SetRequestId(v string) *ListDataQualityEvaluationTasksResponseBody
	GetRequestId() *string
}

type ListDataQualityEvaluationTasksResponseBody struct {
	// The pagination information.
	PagingInfo *ListDataQualityEvaluationTasksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBody) GetPagingInfo() *ListDataQualityEvaluationTasksResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataQualityEvaluationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityEvaluationTasksResponseBody) SetPagingInfo(v *ListDataQualityEvaluationTasksResponseBodyPagingInfo) *ListDataQualityEvaluationTasksResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBody) SetRequestId(v string) *ListDataQualityEvaluationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfo struct {
	// The data quality monitoring tasks.
	DataQualityEvaluationTasks []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks `json:"DataQualityEvaluationTasks,omitempty" xml:"DataQualityEvaluationTasks,omitempty" type:"Repeated"`
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
	// 131
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) GetDataQualityEvaluationTasks() []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	return s.DataQualityEvaluationTasks
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) GetTotalCount() *string {
	return s.TotalCount
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

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks struct {
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the data quality monitoring task. The description can be up to 65,535 characters in length.
	//
	// example:
	//
	// This is a daily run data quality evaluation plan
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The callback configurations of the task during the instance lifecycle. Blocking an auto triggered node is a type of callback event. Only this type is supported.
	Hooks []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitoring task.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data quality monitoring task. The name can be up to 255 characters in length and can contain digits, letters, and punctuation marks.
	//
	// example:
	//
	// Data quality verification task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations for alert notifications.
	Notifications *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
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
	// { "queue": "default", "sqlEngine": "SPARK-SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the task.
	Target *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the task.
	Trigger *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetDescription() *string {
	return s.Description
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetHooks() []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks {
	return s.Hooks
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetName() *string {
	return s.Name
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetNotifications() *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications {
	return s.Notifications
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetRuntimeConf() *string {
	return s.RuntimeConf
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetTarget() *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget {
	return s.Target
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) GetTrigger() *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger {
	return s.Trigger
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetDataSourceId(v int64) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.DataSourceId = &v
	return s
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

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) SetTrigger(v *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks {
	s.Trigger = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks struct {
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

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) GetCondition() *string {
	return s.Condition
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) GetType() *string {
	return s.Type
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) SetCondition(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) SetType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks {
	s.Type = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications struct {
	// The trigger condition of the alert notification.
	//
	// example:
	//
	// ${severity} == "High"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The configurations for the alert notification.
	Notifications []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) GetCondition() *string {
	return s.Condition
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) GetNotifications() []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications {
	return s.Notifications
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) SetCondition(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications {
	s.Condition = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) SetNotifications(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications {
	s.Notifications = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications struct {
	// The alert notification methods.
	NotificationChannels []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert recipients.
	NotificationReceivers []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) GetNotificationChannels() []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels {
	return s.NotificationChannels
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) GetNotificationReceivers() []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) SetNotificationChannels(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications {
	s.NotificationChannels = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) SetNotificationReceivers(v []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications {
	s.NotificationReceivers = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels struct {
	// The alert notification methods.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) SetChannels(v []*string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels {
	s.Channels = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers struct {
	// The extended information in the JSON format. For example, the DingTalk chatbot can remind all members in a DingTalk group by using the at sign (@).
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient. Valid values:
	//
	// 	- AliUid: Alibaba Cloud account ID
	//
	// 	- WebhookUrl: URL of a custom webhook
	//
	// 	- DingdingUrl: DingTalk chatbot URL
	//
	// 	- FeishuUrl: Lark chatbot URL
	//
	// 	- WeixinUrl: WeCom chatbot URL
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
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

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget struct {
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

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) GetType() *string {
	return s.Type
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

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger struct {
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

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) GoString() string {
	return s.String()
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) GetType() *string {
	return s.Type
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) SetTaskIds(v []*int64) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger {
	s.TaskIds = v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) SetType(v string) *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger {
	s.Type = &v
	return s
}

func (s *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTrigger) Validate() error {
	return dara.Validate(s)
}
