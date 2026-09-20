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
	// Specifies whether alerting is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	AlertEnabled *bool `json:"AlertEnabled,omitempty" xml:"AlertEnabled,omitempty"`
	// The baseline alert margin. Unit: minutes.
	//
	// example:
	//
	// 30
	AlertMarginThreshold *int32 `json:"AlertMarginThreshold,omitempty" xml:"AlertMarginThreshold,omitempty"`
	// The baseline alert configurations.
	AlertSettings []*UpdateBaselineRequestAlertSettings `json:"AlertSettings,omitempty" xml:"AlertSettings,omitempty" type:"Repeated"`
	// The ID of the baseline. You can call [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000010800007
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline name.
	//
	// example:
	//
	// BaselineName
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The baseline type. Valid values:
	//
	// - DAILY: daily baseline.
	//
	// - HOURLY: hourly baseline.
	//
	// example:
	//
	// DAILY
	BaselineType *string `json:"BaselineType,omitempty" xml:"BaselineType,omitempty"`
	// Specifies whether the baseline is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of upstream node IDs for the baseline, separated by commas. If there are many nodes, we recommend that you add a virtual node downstream for easier management.
	//
	// example:
	//
	// 1,2,3
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The baseline committed time configurations.
	OvertimeSettings []*UpdateBaselineRequestOvertimeSettings `json:"OvertimeSettings,omitempty" xml:"OvertimeSettings,omitempty" type:"Repeated"`
	// The Alibaba Cloud UID of the baseline owner.
	//
	// example:
	//
	// 3726346****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: 1, 3, 5, 7, and 8.
	//
	// example:
	//
	// 7
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The project ID. You can call [ListBaselines](https://help.aliyun.com/document_detail/2261507.html) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2043
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The IDs of nodes to remove from the baseline. Separate multiple IDs with commas (,).
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
	if s.AlertSettings != nil {
		for _, item := range s.AlertSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OvertimeSettings != nil {
		for _, item := range s.OvertimeSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateBaselineRequestAlertSettings struct {
	// The event alerting interval. Unit: minutes. Minimum value: 5. Maximum value: 1440.
	//
	// example:
	//
	// 30
	AlertInterval *int32 `json:"AlertInterval,omitempty" xml:"AlertInterval,omitempty"`
	// The maximum number of event alerting notifications. Maximum value: 288.
	//
	// example:
	//
	// 1
	AlertMaximum *int32 `json:"AlertMaximum,omitempty" xml:"AlertMaximum,omitempty"`
	// Valid values:
	//
	// - MAIL: email.
	//
	// - SMS: text message.
	//
	// - PHONE: phone call. Only DataWorks Professional Edition and higher support phone call alerts.
	//
	// - DINGROBOTS: DingTalk chatbot. This alert method takes effect only after the RobotUrls parameter is configured.
	//
	// - Webhooks: WeCom or Lark chatbot. This alert method takes effect only after the Webhooks parameter is configured.
	AlertMethods []*string `json:"AlertMethods,omitempty" xml:"AlertMethods,omitempty" type:"Repeated"`
	// The alert recipient details. For specified users: a list of employee IDs. For on-duty schedule: the schedule name. For owner: leave empty.
	//
	// example:
	//
	// 123123
	AlertRecipient *string `json:"AlertRecipient,omitempty" xml:"AlertRecipient,omitempty"`
	// The alert recipient type. Valid values:
	//
	// - OWNER: node owner.
	//
	// - OTHER: specified users.
	//
	// - SHIFT_SCHEDULE: on-duty schedule.
	//
	// example:
	//
	// OWNER
	AlertRecipientType *string `json:"AlertRecipientType,omitempty" xml:"AlertRecipientType,omitempty"`
	// The alert type. Valid values:
	//
	// - BASELINE: baseline alerting.
	//
	// - TOPIC: event alerting.
	//
	// example:
	//
	// BASELINE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// Specifies whether baseline alerting is enabled. This is a baseline-specific configuration. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	BaselineAlertEnabled *bool `json:"BaselineAlertEnabled,omitempty" xml:"BaselineAlertEnabled,omitempty"`
	// The list of DingTalk chatbots.
	DingRobots []*UpdateBaselineRequestAlertSettingsDingRobots `json:"DingRobots,omitempty" xml:"DingRobots,omitempty" type:"Repeated"`
	// The silence end time.
	//
	// example:
	//
	// 00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The silence start time.
	//
	// example:
	//
	// 00:00
	SilenceStartTime *string `json:"SilenceStartTime,omitempty" xml:"SilenceStartTime,omitempty"`
	// The threshold configuration for event slowdown alerts.
	TopicSlowConfig *UpdateBaselineRequestAlertSettingsTopicSlowConfig `json:"TopicSlowConfig,omitempty" xml:"TopicSlowConfig,omitempty" type:"Struct"`
	// The event alerting type. This is an event-specific configuration.
	TopicTypes []*string `json:"TopicTypes,omitempty" xml:"TopicTypes,omitempty" type:"Repeated"`
	// The webhook list.
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

func (s *UpdateBaselineRequestAlertSettings) GetTopicSlowConfig() *UpdateBaselineRequestAlertSettingsTopicSlowConfig {
	return s.TopicSlowConfig
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

func (s *UpdateBaselineRequestAlertSettings) SetTopicSlowConfig(v *UpdateBaselineRequestAlertSettingsTopicSlowConfig) *UpdateBaselineRequestAlertSettings {
	s.TopicSlowConfig = v
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
	if s.DingRobots != nil {
		for _, item := range s.DingRobots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TopicSlowConfig != nil {
		if err := s.TopicSlowConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateBaselineRequestAlertSettingsDingRobots struct {
	// Specifies whether to @all members. Valid values:
	//
	// - true: Yes.
	//
	// - false: No.
	//
	// example:
	//
	// false
	AtAll *bool `json:"AtAll,omitempty" xml:"AtAll,omitempty"`
	// The webhook URL of the DingTalk group chatbot.
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

type UpdateBaselineRequestAlertSettingsTopicSlowConfig struct {
	// The minimum slowdown threshold. Unit: seconds. Valid values: 300 to 18000.
	//
	// example:
	//
	// 3600
	MinOver *int32 `json:"MinOver,omitempty" xml:"MinOver,omitempty"`
	// The ratio used to calculate the slowdown threshold based on the historical average execution duration of the node. Valid values: 0.1 to 2.
	//
	// example:
	//
	// 0.2
	OverFactor *float64 `json:"OverFactor,omitempty" xml:"OverFactor,omitempty"`
}

func (s UpdateBaselineRequestAlertSettingsTopicSlowConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineRequestAlertSettingsTopicSlowConfig) GoString() string {
	return s.String()
}

func (s *UpdateBaselineRequestAlertSettingsTopicSlowConfig) GetMinOver() *int32 {
	return s.MinOver
}

func (s *UpdateBaselineRequestAlertSettingsTopicSlowConfig) GetOverFactor() *float64 {
	return s.OverFactor
}

func (s *UpdateBaselineRequestAlertSettingsTopicSlowConfig) SetMinOver(v int32) *UpdateBaselineRequestAlertSettingsTopicSlowConfig {
	s.MinOver = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettingsTopicSlowConfig) SetOverFactor(v float64) *UpdateBaselineRequestAlertSettingsTopicSlowConfig {
	s.OverFactor = &v
	return s
}

func (s *UpdateBaselineRequestAlertSettingsTopicSlowConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateBaselineRequestOvertimeSettings struct {
	// The epoch corresponding to the committed time. For daily baselines, the value is 1. For hourly baselines, you can configure up to 24 epochs.
	//
	// example:
	//
	// 1
	Cycle *int32 `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The committed time in hh:mm format. Valid values of hh: 0 to 47. Valid values of mm: 0 to 59.
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
