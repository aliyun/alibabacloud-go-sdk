// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEnabled(v bool) *UpdateBaselineRequest
	GetAlertEnabled() *bool
	SetAlertMarginThreshold(v int32) *UpdateBaselineRequest
	GetAlertMarginThreshold() *int32
	SetAlertSettings(v []*UpdateBaselineRequestAlertSettings) *UpdateBaselineRequest
	GetAlertSettings() []*UpdateBaselineRequestAlertSettings
	SetBaselineId(v int64) *UpdateBaselineRequest
	GetBaselineId() *int64
	SetBaselineName(v string) *UpdateBaselineRequest
	GetBaselineName() *string
	SetBaselineType(v string) *UpdateBaselineRequest
	GetBaselineType() *string
	SetEnabled(v bool) *UpdateBaselineRequest
	GetEnabled() *bool
	SetNodeIds(v string) *UpdateBaselineRequest
	GetNodeIds() *string
	SetOvertimeSettings(v []*UpdateBaselineRequestOvertimeSettings) *UpdateBaselineRequest
	GetOvertimeSettings() []*UpdateBaselineRequestOvertimeSettings
	SetOwner(v string) *UpdateBaselineRequest
	GetOwner() *string
	SetPriority(v int32) *UpdateBaselineRequest
	GetPriority() *int32
	SetProjectId(v int64) *UpdateBaselineRequest
	GetProjectId() *int64
	SetRemoveNodeIds(v string) *UpdateBaselineRequest
	GetRemoveNodeIds() *string
}

type UpdateBaselineRequest struct {
	// Specifies whether to enable the alerting feature. Valid values: true and false.
	//
	// example:
	//
	// true
	AlertEnabled *bool `json:"AlertEnabled,omitempty" xml:"AlertEnabled,omitempty"`
	// The alert margin threshold of the baseline. Unit: minutes.
	//
	// example:
	//
	// 30
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The alert settings of the baseline.
	AlertSettings []*UpdateBaselineRequestAlertSettings `json:"AlertSettings,omitempty" xml:"AlertSettings,omitempty" type:"Repeated"`
	// The baseline ID. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000010800007
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// BaselineName
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The type of the baseline. Valid values: DAILY and HOURLY.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// Specifies whether to enable the baseline. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ancestor nodes of nodes in the baseline. Separate the ancestor nodes with commas (,). If a large number of ancestor nodes exist, we recommend that you create a zero load node and configure the zero load node as the descendant node of nodes in the baseline to facilitate node management.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The settings of the committed completion time of the baseline.
	OvertimeSettings []*UpdateBaselineRequestOvertimeSettings `json:"OvertimeSettings,omitempty" xml:"OvertimeSettings,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account used by the baseline owner.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: {1,3,5,7,8}.
	//
	// example:
	//
	// 7
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The workspace ID. You can call the [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2043
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the node that you want to disassociate from the baseline. You can specify multiple node IDs. Separate multiple node IDs with commas (,).
	//
	// example:
	//
	// 123,456
	RemoveNodeIds *string `json:"RemoveNodeIds,omitempty" xml:"RemoveNodeIds,omitempty"`
}

func (s UpdateBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdateBaselineRequest) GetAlertEnabled() *bool {
	return s.AlertEnabled
}

func (s *UpdateBaselineRequest) GetAlertMarginThreshold() *int32 {
	return s.AlertMarginThreshold
}

func (s *UpdateBaselineRequest) GetAlertSettings() []*UpdateBaselineRequestAlertSettings {
	return s.AlertSettings
}

func (s *UpdateBaselineRequest) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *UpdateBaselineRequest) GetBaselineName() *string {
	return s.BaselineName
}

func (s *UpdateBaselineRequest) GetBaselineType() *string {
	return s.BaselineType
}

func (s *UpdateBaselineRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateBaselineRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *UpdateBaselineRequest) GetOvertimeSettings() []*UpdateBaselineRequestOvertimeSettings {
	return s.OvertimeSettings
}

func (s *UpdateBaselineRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateBaselineRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateBaselineRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateBaselineRequest) GetRemoveNodeIds() *string {
	return s.RemoveNodeIds
}

func (s *UpdateBaselineRequest) SetAlertEnabled(v bool) *UpdateBaselineRequest {
	s.AlertEnabled = &v
	return s
}

func (s *UpdateBaselineRequest) SetAlertMarginThreshold(v int32) *UpdateBaselineRequest {
	s.AlertMarginThreshold = &v
	return s
}

func (s *UpdateBaselineRequest) SetAlertSettings(v []*UpdateBaselineRequestAlertSettings) *UpdateBaselineRequest {
	s.AlertSettings = v
	return s
}

func (s *UpdateBaselineRequest) SetBaselineId(v int64) *UpdateBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *UpdateBaselineRequest) SetBaselineName(v string) *UpdateBaselineRequest {
	s.BaselineName = &v
	return s
}

func (s *UpdateBaselineRequest) SetBaselineType(v string) *UpdateBaselineRequest {
	s.BaselineType = &v
	return s
}

func (s *UpdateBaselineRequest) SetEnabled(v bool) *UpdateBaselineRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateBaselineRequest) SetNodeIds(v string) *UpdateBaselineRequest {
	s.NodeIds = &v
	return s
}

func (s *UpdateBaselineRequest) SetOvertimeSettings(v []*UpdateBaselineRequestOvertimeSettings) *UpdateBaselineRequest {
	s.OvertimeSettings = v
	return s
}

func (s *UpdateBaselineRequest) SetOwner(v string) *UpdateBaselineRequest {
	s.Owner = &v
	return s
}

func (s *UpdateBaselineRequest) SetPriority(v int32) *UpdateBaselineRequest {
	s.Priority = &v
	return s
}

func (s *UpdateBaselineRequest) SetProjectId(v int64) *UpdateBaselineRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateBaselineRequest) SetRemoveNodeIds(v string) *UpdateBaselineRequest {
	s.RemoveNodeIds = &v
	return s
}

func (s *UpdateBaselineRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateBaselineRequestAlertSettings struct {
	// The interval at which an event alert notification is sent. Unit: minutes. Minimum value: 5. Maximum value: 1,440.
	//
	// example:
	//
	// 1800
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The maximum number of times an event alert notification is sent. Maximum value: 24.
	//
	// example:
	//
	// 1
	AlertMaximum *int32 `json:"AlertMaximum,omitempty" xml:"AlertMaximum,omitempty"`
	// The alert notification methods. Valid values: MAIL, SMS, PHONE, DINGROBOTS, and Webhooks. The value MAIL indicates that alert notifications are sent by email. The value SMS indicates that alert notifications are sent by text message. The value PHONE indicates that alert notifications are sent by phone call. You can use this notification method only in DataWorks Professional Edition or a more advanced edition. The value DINGROBOTS indicates that alert notifications are sent by using a DingTalk chatbot. You can use this notification method only if the RobotUrls parameter is configured. The value Webhooks indicates that alert notifications are sent by WeCom or Lark. You can use this notification method only if the Webhooks parameter is configured.
	AlertMethods []*string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty" type:"Repeated"`
	// The details of the alert recipient. If you set AlertRecipientType to OWNER, leave this parameter empty. If you set AlertRecipientType to SHIFT_SCHEDULE, set this parameter to the name of the shift schedule. If you set AlertRecipientType to OTHER, set this parameter to the employee IDs of specified personnel.
	//
	// example:
	//
	// 123123
	AlertRecipient *string `json:"AlertRecipient,omitempty" xml:"AlertRecipient,omitempty"`
	// The type of the alert recipient. Valid values: OWNER, OTHER, and SHIFT_SCHEDULE. The value OWNER indicates the node owner. The value OTHER indicates specified personnel. The value SHIFT_SCHEDULE indicates personnel in a shift schedule.
	//
	// example:
	//
	// OWNER
	AlertRecipientType *string `json:"AlertRecipientType,omitempty" xml:"AlertRecipientType,omitempty"`
	// The type of the alert. Valid values: BASELINE and TOPIC. The value BASELINE indicates a baseline alert. The value TOPIC indicates an event alert.
	//
	// example:
	//
	// BASELINE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// Specifies whether to enable the baseline alerting feature. This feature is specific to baselines. Valid values: true and false.
	//
	// example:
	//
	// true
	BaselineAlertEnabled *bool `json:"BaselineAlertEnabled,omitempty" xml:"BaselineAlertEnabled,omitempty"`
	// The DingTalk chatbots.
	DingRobots []*UpdateBaselineRequestAlertSettingsDingRobots `json:"DingRobots,omitempty" xml:"DingRobots,omitempty" type:"Repeated"`
	// The end time of silence.
	//
	// example:
	//
	// 00:00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The start time of silence.
	//
	// example:
	//
	// 00:00:00
	SilenceStartTime *string `json:"SilenceStartTime,omitempty" xml:"SilenceStartTime,omitempty"`
	// The types of event alerts, which are event-specific configurations.
	TopicTypes []*string `json:"TopicTypes,omitempty" xml:"TopicTypes,omitempty" type:"Repeated"`
	// The webhook URLs.
	Webhooks []*string `json:"Webhooks,omitempty" xml:"Webhooks,omitempty" type:"Repeated"`
}

func (s UpdateBaselineRequestAlertSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineRequestAlertSettings) GoString() string {
	return s.String()
}

func (s *UpdateBaselineRequestAlertSettings) GetAlertInterval() *int32 {
	return s.AlertInterval
}

func (s *UpdateBaselineRequestAlertSettings) GetAlertMaximum() *int32 {
	return s.AlertMaximum
}

func (s *UpdateBaselineRequestAlertSettings) GetAlertMethods() []*string {
	return s.AlertMethods
}

func (s *UpdateBaselineRequestAlertSettings) GetAlertRecipient() *string {
	return s.AlertRecipient
}

func (s *UpdateBaselineRequestAlertSettings) GetAlertRecipientType() *string {
	return s.AlertRecipientType
}

func (s *UpdateBaselineRequestAlertSettings) GetAlertType() *string {
	return s.AlertType
}

func (s *UpdateBaselineRequestAlertSettings) GetBaselineAlertEnabled() *bool {
	return s.BaselineAlertEnabled
}

func (s *UpdateBaselineRequestAlertSettings) GetDingRobots() []*UpdateBaselineRequestAlertSettingsDingRobots {
	return s.DingRobots
}

func (s *UpdateBaselineRequestAlertSettings) GetSilenceEndTime() *string {
	return s.SilenceEndTime
}

func (s *UpdateBaselineRequestAlertSettings) GetSilenceStartTime() *string {
	return s.SilenceStartTime
}

func (s *UpdateBaselineRequestAlertSettings) GetTopicTypes() []*string {
	return s.TopicTypes
}

func (s *UpdateBaselineRequestAlertSettings) GetWebhooks() []*string {
	return s.Webhooks
}

func (s *UpdateBaselineRequestAlertSettings) SetAlertInterval(v int32) *UpdateBaselineRequestAlertSettings {
	s.AlertInterval = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetAlertMaximum(v int32) *UpdateBaselineRequestAlertSettings {
	s.AlertMaximum = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetAlertMethods(v []*string) *UpdateBaselineRequestAlertSettings {
	s.AlertMethods = v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetAlertRecipient(v string) *UpdateBaselineRequestAlertSettings {
	s.AlertRecipient = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetAlertRecipientType(v string) *UpdateBaselineRequestAlertSettings {
	s.AlertRecipientType = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetAlertType(v string) *UpdateBaselineRequestAlertSettings {
	s.AlertType = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetBaselineAlertEnabled(v bool) *UpdateBaselineRequestAlertSettings {
	s.BaselineAlertEnabled = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetDingRobots(v []*UpdateBaselineRequestAlertSettingsDingRobots) *UpdateBaselineRequestAlertSettings {
	s.DingRobots = v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetSilenceEndTime(v string) *UpdateBaselineRequestAlertSettings {
	s.SilenceEndTime = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetSilenceStartTime(v string) *UpdateBaselineRequestAlertSettings {
	s.SilenceStartTime = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetTopicTypes(v []*string) *UpdateBaselineRequestAlertSettings {
	s.TopicTypes = v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) SetWebhooks(v []*string) *UpdateBaselineRequestAlertSettings {
	s.Webhooks = v
	return s
}

func (s *UpdateBaselineRequestAlertSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateBaselineRequestAlertSettingsDingRobots struct {
	// Specifies whether to remind all members by using the at sign (@). Valid values: true and false.
	//
	// example:
	//
	// false
	AtAll *bool `json:"AtAll,omitempty" xml:"AtAll,omitempty"`
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=xxx
	WebUrl *string `json:"WebUrl,omitempty" xml:"WebUrl,omitempty"`
}

func (s UpdateBaselineRequestAlertSettingsDingRobots) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineRequestAlertSettingsDingRobots) GoString() string {
	return s.String()
}

func (s *UpdateBaselineRequestAlertSettingsDingRobots) GetAtAll() *bool {
	return s.AtAll
}

func (s *UpdateBaselineRequestAlertSettingsDingRobots) GetWebUrl() *string {
	return s.WebUrl
}

func (s *UpdateBaselineRequestAlertSettingsDingRobots) SetAtAll(v bool) *UpdateBaselineRequestAlertSettingsDingRobots {
	s.AtAll = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettingsDingRobots) SetWebUrl(v string) *UpdateBaselineRequestAlertSettingsDingRobots {
	s.WebUrl = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettingsDingRobots) Validate() error {
	return dara.Validate(s)
}

type UpdateBaselineRequestOvertimeSettings struct {
	// The cycle that corresponds to the committed completion time. For a day-level baseline, set this parameter to 1. For an hour-level baseline, set this parameter to a value that is no more than 24.
	//
	// example:
	//
	// 1
	Cycle *int32 `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The committed completion time in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 00:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s UpdateBaselineRequestOvertimeSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineRequestOvertimeSettings) GoString() string {
	return s.String()
}

func (s *UpdateBaselineRequestOvertimeSettings) GetCycle() *int32 {
	return s.Cycle
}

func (s *UpdateBaselineRequestOvertimeSettings) GetTime() *string {
	return s.Time
}

func (s *UpdateBaselineRequestOvertimeSettings) SetCycle(v int32) *UpdateBaselineRequestOvertimeSettings {
	s.Cycle = &v
	return s
}

func (s *UpdateBaselineRequestOvertimeSettings) SetTime(v string) *UpdateBaselineRequestOvertimeSettings {
	s.Time = &v
	return s
}

func (s *UpdateBaselineRequestOvertimeSettings) Validate() error {
	return dara.Validate(s)
}
