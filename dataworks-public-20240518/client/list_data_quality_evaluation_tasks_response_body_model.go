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
	// The paged query result of quality evaluation nodes.
	PagingInfo *ListDataQualityEvaluationTasksResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The API request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-****
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
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfo struct {
	// The quality evaluation tasks.
	DataQualityEvaluationTasks []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks `json:"DataQualityEvaluationTasks,omitempty" xml:"DataQualityEvaluationTasks,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
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
	if s.DataQualityEvaluationTasks != nil {
		for _, item := range s.DataQualityEvaluationTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasks struct {
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The description of the data quality evaluation task. The description can be up to 65,535 characters in length.
	//
	// example:
	//
	// This is a daily run data quality evaluation plan
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The callback settings during the epoch of data quality evaluation task instances. Currently, only one hook that blocks a scheduling node instance is supported.
	Hooks []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality evaluation task.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data quality evaluation task. The name can contain digits, letters, Chinese characters, and half-width or full-width punctuation marks. The name can be up to 255 characters in length.
	//
	// example:
	//
	// Data quality verification task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert configuration.
	Notifications *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotifications `json:"Notifications,omitempty" xml:"Notifications,omitempty" type:"Struct"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The runtime settings for data sources. Currently, only specifying the EMR YARN queue and the SQL engine for collecting EMR tables is supported. Valid SQL engine values: SPARK_SQL, KYUUBI, PRESTO_SQL, and HIVE_SQL.
	//
	// example:
	//
	// { "queue": "default", "sqlEngine": "SPARK_SQL" }
	RuntimeConf *string `json:"RuntimeConf,omitempty" xml:"RuntimeConf,omitempty"`
	// The monitored object of the data quality evaluation task.
	Target *ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The trigger configuration of the data quality evaluation task.
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

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksHooks struct {
	// The cause that triggers the hook.
	//
	// example:
	//
	// ${severity} == "High" AND ${status} == "Critical"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the follow-up action. Valid values:
	//
	// - BlockTaskInstance: Blocks the execution of a DataWorks node instance.
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
	// The cause that triggers the notification.
	//
	// example:
	//
	// ${severity} == "High"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The alert settings.
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

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotifications struct {
	// The alert channel configurations.
	NotificationChannels []*ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The alert recipient configurations.
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

type ListDataQualityEvaluationTasksResponseBodyPagingInfoDataQualityEvaluationTasksNotificationsNotificationsNotificationChannels struct {
	// The alert channel.
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
	// The extension information in JSON format. For example, DingTalk chatbots support mentioning all members.
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient. Valid values:
	//
	// - AliUid: Alibaba Cloud account UID.
	//
	// - WebhookUrl: custom webhook URL.
	//
	// - DingdingUrl: DingTalk chatbot URL.
	//
	// - FeishuUrl: Lark chatbot URL.
	//
	// - WeixinUrl: WeCom chatbot URL.
	//
	// example:
	//
	// AliUid
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The specific values of the alert recipients.
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
	// The database type of the table dataset. Valid values:
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
	// The unique ID of the table in DataWorks Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The monitored object type. Valid values:
	//
	// - Table
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
	// This parameter takes effect when type is set to ByScheduledTaskInstance. Specifies the scheduled node IDs whose successful instance execution can trigger the task.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The type of event that triggers the quality evaluation task. Valid values:
	//
	// - ByScheduledTaskInstance: Triggered when a scheduled node instance runs successfully.
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
