// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *CreateAppGroupRequest
	GetAppKey() *string
	SetAppName(v string) *CreateAppGroupRequest
	GetAppName() *string
	SetAppType(v int32) *CreateAppGroupRequest
	GetAppType() *int32
	SetAppVersion(v int32) *CreateAppGroupRequest
	GetAppVersion() *int32
	SetDescription(v string) *CreateAppGroupRequest
	GetDescription() *string
	SetEnableLog(v bool) *CreateAppGroupRequest
	GetEnableLog() *bool
	SetGroupId(v string) *CreateAppGroupRequest
	GetGroupId() *string
	SetMaxJobs(v int32) *CreateAppGroupRequest
	GetMaxJobs() *int32
	SetMonitorConfigJson(v string) *CreateAppGroupRequest
	GetMonitorConfigJson() *string
	SetMonitorContactsJson(v string) *CreateAppGroupRequest
	GetMonitorContactsJson() *string
	SetNamespace(v string) *CreateAppGroupRequest
	GetNamespace() *string
	SetNamespaceName(v string) *CreateAppGroupRequest
	GetNamespaceName() *string
	SetNamespaceSource(v string) *CreateAppGroupRequest
	GetNamespaceSource() *string
	SetNotificationPolicyName(v string) *CreateAppGroupRequest
	GetNotificationPolicyName() *string
	SetRegionId(v string) *CreateAppGroupRequest
	GetRegionId() *string
	SetScheduleBusyWorkers(v bool) *CreateAppGroupRequest
	GetScheduleBusyWorkers() *bool
}

type CreateAppGroupRequest struct {
	// The AppKey of the application.
	//
	// example:
	//
	// adcExHZviLcl****
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// DocTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// - `1`: Standard application.
	//
	// - `2`: Kubernetes (K8s) application.
	//
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The application version. Valid values: `1` (Basic Edition) and `2` (Professional Edition).
	//
	// example:
	//
	// 2
	AppVersion *int32 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The application description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable logging.
	//
	// - `true`: Enable logging.
	//
	// - `false`: Disable logging.
	//
	// example:
	//
	// true
	EnableLog *bool `json:"EnableLog,omitempty" xml:"EnableLog,omitempty"`
	// The ID of the application group. You can find this ID on the **Application Management*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The maximum number of jobs.
	//
	// example:
	//
	// 1000
	MaxJobs *int32 `json:"MaxJobs,omitempty" xml:"MaxJobs,omitempty"`
	// Specifies the alert notification configuration as a JSON string. The string can contain the following properties: `sendChannel`, `alarmType`, and `webhookIsAtAll`.
	//
	// > For more information, see the **Additional information about request parameters*	- section.
	//
	// example:
	//
	// {
	//
	//     "sendChannel": "ding,sms,mail,phone",
	//
	//     "alarmType": "Contacts",
	//
	//     "webhookIsAtAll": false
	//
	// }
	MonitorConfigJson *string `json:"MonitorConfigJson,omitempty" xml:"MonitorConfigJson,omitempty"`
	// The alert contacts. This can include individual contacts and contact groups.
	//
	// > For more information, see the **Additional information about request parameters*	- section.
	//
	// example:
	//
	// [
	//
	//     {"name": "Alice Johnson"},
	//
	//     {"name": "Lee Smith"}
	//
	// ]
	MonitorContactsJson *string `json:"MonitorContactsJson,omitempty" xml:"MonitorContactsJson,omitempty"`
	// The ID of the namespace. You can find this ID on the **Namespace*	- page in the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// Test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// This parameter is currently unsupported and can be left unspecified.
	//
	// example:
	//
	// schedulerx
	NamespaceSource *string `json:"NamespaceSource,omitempty" xml:"NamespaceSource,omitempty"`
	// The notification policy name.
	//
	// example:
	//
	// test-workday-notification
	NotificationPolicyName *string `json:"NotificationPolicyName,omitempty" xml:"NotificationPolicyName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to schedule jobs on a busy worker.
	//
	// example:
	//
	// false
	ScheduleBusyWorkers *bool `json:"ScheduleBusyWorkers,omitempty" xml:"ScheduleBusyWorkers,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateAppGroupRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppGroupRequest) GetAppType() *int32 {
	return s.AppType
}

func (s *CreateAppGroupRequest) GetAppVersion() *int32 {
	return s.AppVersion
}

func (s *CreateAppGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppGroupRequest) GetEnableLog() *bool {
	return s.EnableLog
}

func (s *CreateAppGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateAppGroupRequest) GetMaxJobs() *int32 {
	return s.MaxJobs
}

func (s *CreateAppGroupRequest) GetMonitorConfigJson() *string {
	return s.MonitorConfigJson
}

func (s *CreateAppGroupRequest) GetMonitorContactsJson() *string {
	return s.MonitorContactsJson
}

func (s *CreateAppGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateAppGroupRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateAppGroupRequest) GetNamespaceSource() *string {
	return s.NamespaceSource
}

func (s *CreateAppGroupRequest) GetNotificationPolicyName() *string {
	return s.NotificationPolicyName
}

func (s *CreateAppGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAppGroupRequest) GetScheduleBusyWorkers() *bool {
	return s.ScheduleBusyWorkers
}

func (s *CreateAppGroupRequest) SetAppKey(v string) *CreateAppGroupRequest {
	s.AppKey = &v
	return s
}

func (s *CreateAppGroupRequest) SetAppName(v string) *CreateAppGroupRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppGroupRequest) SetAppType(v int32) *CreateAppGroupRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppGroupRequest) SetAppVersion(v int32) *CreateAppGroupRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppGroupRequest) SetDescription(v string) *CreateAppGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateAppGroupRequest) SetEnableLog(v bool) *CreateAppGroupRequest {
	s.EnableLog = &v
	return s
}

func (s *CreateAppGroupRequest) SetGroupId(v string) *CreateAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateAppGroupRequest) SetMaxJobs(v int32) *CreateAppGroupRequest {
	s.MaxJobs = &v
	return s
}

func (s *CreateAppGroupRequest) SetMonitorConfigJson(v string) *CreateAppGroupRequest {
	s.MonitorConfigJson = &v
	return s
}

func (s *CreateAppGroupRequest) SetMonitorContactsJson(v string) *CreateAppGroupRequest {
	s.MonitorContactsJson = &v
	return s
}

func (s *CreateAppGroupRequest) SetNamespace(v string) *CreateAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *CreateAppGroupRequest) SetNamespaceName(v string) *CreateAppGroupRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateAppGroupRequest) SetNamespaceSource(v string) *CreateAppGroupRequest {
	s.NamespaceSource = &v
	return s
}

func (s *CreateAppGroupRequest) SetNotificationPolicyName(v string) *CreateAppGroupRequest {
	s.NotificationPolicyName = &v
	return s
}

func (s *CreateAppGroupRequest) SetRegionId(v string) *CreateAppGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAppGroupRequest) SetScheduleBusyWorkers(v bool) *CreateAppGroupRequest {
	s.ScheduleBusyWorkers = &v
	return s
}

func (s *CreateAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
