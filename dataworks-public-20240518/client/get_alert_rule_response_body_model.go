// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRule(v *GetAlertRuleResponseBodyAlertRule) *GetAlertRuleResponseBody
	GetAlertRule() *GetAlertRuleResponseBodyAlertRule
	SetRequestId(v string) *GetAlertRuleResponseBody
	GetRequestId() *string
}

type GetAlertRuleResponseBody struct {
	// The details of the custom alert rule.
	AlertRule *GetAlertRuleResponseBodyAlertRule `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBody) GetAlertRule() *GetAlertRuleResponseBodyAlertRule {
	return s.AlertRule
}

func (s *GetAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertRuleResponseBody) SetAlertRule(v *GetAlertRuleResponseBodyAlertRule) *GetAlertRuleResponseBody {
	s.AlertRule = v
	return s
}

func (s *GetAlertRuleResponseBody) SetRequestId(v string) *GetAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertRuleResponseBody) Validate() error {
	if s.AlertRule != nil {
		if err := s.AlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertRuleResponseBodyAlertRule struct {
	// Indicates whether the alert rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the custom alert rule.
	//
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the custom alert rule.
	//
	// example:
	//
	// error_rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert notification configuration.
	Notification *GetAlertRuleResponseBodyAlertRuleNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The Alibaba Cloud UID of the owner of the custom alert rule.
	//
	// example:
	//
	// 279961421580845157
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The condition that triggers the alert.
	TriggerCondition *GetAlertRuleResponseBodyAlertRuleTriggerCondition `json:"TriggerCondition,omitempty" xml:"TriggerCondition,omitempty" type:"Struct"`
}

func (s GetAlertRuleResponseBodyAlertRule) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRule) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetAlertRuleResponseBodyAlertRule) GetId() *int64 {
	return s.Id
}

func (s *GetAlertRuleResponseBodyAlertRule) GetName() *string {
	return s.Name
}

func (s *GetAlertRuleResponseBodyAlertRule) GetNotification() *GetAlertRuleResponseBodyAlertRuleNotification {
	return s.Notification
}

func (s *GetAlertRuleResponseBodyAlertRule) GetOwner() *string {
	return s.Owner
}

func (s *GetAlertRuleResponseBodyAlertRule) GetTriggerCondition() *GetAlertRuleResponseBodyAlertRuleTriggerCondition {
	return s.TriggerCondition
}

func (s *GetAlertRuleResponseBodyAlertRule) SetEnabled(v bool) *GetAlertRuleResponseBodyAlertRule {
	s.Enabled = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRule) SetId(v int64) *GetAlertRuleResponseBodyAlertRule {
	s.Id = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRule) SetName(v string) *GetAlertRuleResponseBodyAlertRule {
	s.Name = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRule) SetNotification(v *GetAlertRuleResponseBodyAlertRuleNotification) *GetAlertRuleResponseBodyAlertRule {
	s.Notification = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRule) SetOwner(v string) *GetAlertRuleResponseBodyAlertRule {
	s.Owner = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRule) SetTriggerCondition(v *GetAlertRuleResponseBodyAlertRuleTriggerCondition) *GetAlertRuleResponseBodyAlertRule {
	s.TriggerCondition = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRule) Validate() error {
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerCondition != nil {
		if err := s.TriggerCondition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertRuleResponseBodyAlertRuleNotification struct {
	// The list of alert channels.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The alert interval, in minutes. Valid values: 5 to 10000.
	//
	// example:
	//
	// 30
	IntervalInMinutes *int32 `json:"IntervalInMinutes,omitempty" xml:"IntervalInMinutes,omitempty"`
	// The maximum number of alerts within a calendar day. Valid values: 1 to 10000.
	//
	// example:
	//
	// 3
	Maximum *int32 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// The alert recipients.
	Receivers []*GetAlertRuleResponseBodyAlertRuleNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The end time of the mute period. Format: HH:mm:ss.
	//
	// example:
	//
	// 00:00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The start time of the mute period. Format: HH:mm:ss.
	//
	// example:
	//
	// 00:00:00
	SilenceStartTime *string `json:"SilenceStartTime,omitempty" xml:"SilenceStartTime,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleNotification) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleNotification) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) GetChannels() []*string {
	return s.Channels
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) GetIntervalInMinutes() *int32 {
	return s.IntervalInMinutes
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) GetMaximum() *int32 {
	return s.Maximum
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) GetReceivers() []*GetAlertRuleResponseBodyAlertRuleNotificationReceivers {
	return s.Receivers
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) GetSilenceEndTime() *string {
	return s.SilenceEndTime
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) GetSilenceStartTime() *string {
	return s.SilenceStartTime
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) SetChannels(v []*string) *GetAlertRuleResponseBodyAlertRuleNotification {
	s.Channels = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) SetIntervalInMinutes(v int32) *GetAlertRuleResponseBodyAlertRuleNotification {
	s.IntervalInMinutes = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) SetMaximum(v int32) *GetAlertRuleResponseBodyAlertRuleNotification {
	s.Maximum = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) SetReceivers(v []*GetAlertRuleResponseBodyAlertRuleNotificationReceivers) *GetAlertRuleResponseBodyAlertRuleNotification {
	s.Receivers = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) SetSilenceEndTime(v string) *GetAlertRuleResponseBodyAlertRuleNotification {
	s.SilenceEndTime = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) SetSilenceStartTime(v string) *GetAlertRuleResponseBodyAlertRuleNotification {
	s.SilenceStartTime = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotification) Validate() error {
	if s.Receivers != nil {
		for _, item := range s.Receivers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlertRuleResponseBodyAlertRuleNotificationReceivers struct {
	// The additional configuration required by the alert recipient. If ReceiverType is DingdingUrl, you can set {"atAll":true} to @ all members.
	//
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient. Valid values:
	//
	// - AliUid: Alibaba Cloud UID.
	//
	// - ShiftSchedule: shift schedule.
	//
	// - TaskOwner: node owner. Applicable to custom alerting and event alerting.
	//
	// - Owner: owner. Applicable to baseline alerting.
	//
	// - WebhookUrl: custom webhook URL.
	//
	// - DingdingUrl: DingTalk webhook URL.
	//
	// - FeishuUrl: Lark webhook URL.
	//
	// - WeixinUrl: WeChat webhook URL.
	//
	// example:
	//
	// WebhookUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The values of the alert recipient.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s GetAlertRuleResponseBodyAlertRuleNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleNotificationReceivers) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) SetExtension(v string) *GetAlertRuleResponseBodyAlertRuleNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) SetReceiverType(v string) *GetAlertRuleResponseBodyAlertRuleNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) SetReceiverValues(v []*string) *GetAlertRuleResponseBodyAlertRuleNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerCondition struct {
	// The extension information. Required for certain trigger conditions.
	Extension *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The monitored object.
	Target *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The type of the alert trigger. Valid values:
	//
	// - Finished: instance completed.
	//
	// - UnFinished: instance not completed.
	//
	// - Error: instance failed.
	//
	// - CycleUnfinished: instance cycle not completed.
	//
	// - Timeout: instance timed out.
	//
	// - InstanceTransferComplete: node-to-instance conversion completed.
	//
	// - InstanceTransferFluctuate: instance count fluctuation.
	//
	// - ExhaustedError: instance still failed after automatic reruns.
	//
	// - InstanceKeyword: failed instance contains keyword.
	//
	// - InstanceErrorCount: number of failed instances.
	//
	// - InstanceErrorPercentage: percentage of failed instances.
	//
	// - ResourceGroupPercentage: schedule resource utilization.
	//
	// - ResourceGroupWaitCount: number of instances waiting for schedule resources.
	//
	// example:
	//
	// Error
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerCondition) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerCondition) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) GetExtension() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	return s.Extension
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) GetTarget() *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget {
	return s.Target
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) GetType() *string {
	return s.Type
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) SetExtension(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) *GetAlertRuleResponseBodyAlertRuleTriggerCondition {
	s.Extension = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) SetTarget(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) *GetAlertRuleResponseBodyAlertRuleTriggerCondition {
	s.Target = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) SetType(v string) *GetAlertRuleResponseBodyAlertRuleTriggerCondition {
	s.Type = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerCondition) Validate() error {
	if s.Extension != nil {
		if err := s.Extension.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension struct {
	// The cycle-not-completed alert configuration.
	CycleUnfinished *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished `json:"CycleUnfinished,omitempty" xml:"CycleUnfinished,omitempty" type:"Struct"`
	// The error alert configuration.
	Error *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The instance error count alert configuration.
	InstanceErrorCount *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount `json:"InstanceErrorCount,omitempty" xml:"InstanceErrorCount,omitempty" type:"Struct"`
	// The instance error percentage alert configuration.
	InstanceErrorPercentage *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage `json:"InstanceErrorPercentage,omitempty" xml:"InstanceErrorPercentage,omitempty" type:"Struct"`
	// The instance count fluctuation alert configuration.
	InstanceTransferFluctuate *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate `json:"InstanceTransferFluctuate,omitempty" xml:"InstanceTransferFluctuate,omitempty" type:"Struct"`
	// The timeout alert configuration.
	Timeout *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
	// The not-completed alert configuration.
	UnFinished *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished `json:"UnFinished,omitempty" xml:"UnFinished,omitempty" type:"Struct"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetCycleUnfinished() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished {
	return s.CycleUnfinished
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetError() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError {
	return s.Error
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetInstanceErrorCount() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount {
	return s.InstanceErrorCount
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetInstanceErrorPercentage() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage {
	return s.InstanceErrorPercentage
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetInstanceTransferFluctuate() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate {
	return s.InstanceTransferFluctuate
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetTimeout() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout {
	return s.Timeout
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) GetUnFinished() *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished {
	return s.UnFinished
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetCycleUnfinished(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.CycleUnfinished = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetError(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.Error = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetInstanceErrorCount(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.InstanceErrorCount = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetInstanceErrorPercentage(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.InstanceErrorPercentage = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetInstanceTransferFluctuate(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.InstanceTransferFluctuate = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetTimeout(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.Timeout = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) SetUnFinished(v *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension {
	s.UnFinished = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension) Validate() error {
	if s.CycleUnfinished != nil {
		if err := s.CycleUnfinished.Validate(); err != nil {
			return err
		}
	}
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceErrorCount != nil {
		if err := s.InstanceErrorCount.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceErrorPercentage != nil {
		if err := s.InstanceErrorPercentage.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceTransferFluctuate != nil {
		if err := s.InstanceTransferFluctuate.Validate(); err != nil {
			return err
		}
	}
	if s.Timeout != nil {
		if err := s.Timeout.Validate(); err != nil {
			return err
		}
	}
	if s.UnFinished != nil {
		if err := s.UnFinished.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished struct {
	// The list of cycle and time configurations.
	CycleAndTime []*GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime `json:"CycleAndTime,omitempty" xml:"CycleAndTime,omitempty" type:"Repeated"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished) GetCycleAndTime() []*GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	return s.CycleAndTime
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished) SetCycleAndTime(v []*GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished {
	s.CycleAndTime = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished) Validate() error {
	if s.CycleAndTime != nil {
		for _, item := range s.CycleAndTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime struct {
	// The cycle ID. Valid values: 1 to 288.
	//
	// example:
	//
	// 1
	CycleId *int32 `json:"CycleId,omitempty" xml:"CycleId,omitempty"`
	// The timeout time. Format: hh:mm. Valid values of hh: 0 to 47. Valid values of mm: 0 to 59.
	//
	// example:
	//
	// 12:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetCycleId() *int32 {
	return s.CycleId
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetTime() *string {
	return s.Time
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetCycleId(v int32) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.CycleId = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetTime(v string) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.Time = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinishedCycleAndTime) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError struct {
	// Specifies whether to generate an alert when a batch task is automatically rerun due to a failure.
	//
	// example:
	//
	// false
	AutoRerunAlertEnabled *bool `json:"AutoRerunAlertEnabled,omitempty" xml:"AutoRerunAlertEnabled,omitempty"`
	// The IDs of real-time computing nodes to monitor.
	StreamTaskIds []*int64 `json:"StreamTaskIds,omitempty" xml:"StreamTaskIds,omitempty" type:"Repeated"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) GetAutoRerunAlertEnabled() *bool {
	return s.AutoRerunAlertEnabled
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) GetStreamTaskIds() []*int64 {
	return s.StreamTaskIds
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) SetAutoRerunAlertEnabled(v bool) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError {
	s.AutoRerunAlertEnabled = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) SetStreamTaskIds(v []*int64) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError {
	s.StreamTaskIds = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount struct {
	// The number of failed instances. Valid values: 1 to 10000.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount) GetCount() *int32 {
	return s.Count
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount) SetCount(v int32) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount {
	s.Count = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage struct {
	// The percentage of failed instances. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage) GetPercentage() *int32 {
	return s.Percentage
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage) SetPercentage(v int32) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage {
	s.Percentage = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate struct {
	// The fluctuation percentage. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The fluctuation type. Valid values:
	//
	// - abs: absolute value.
	//
	// - increase: increase.
	//
	// - decrease: decrease.
	//
	// example:
	//
	// 10
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) GetPercentage() *int32 {
	return s.Percentage
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) GetTrend() *string {
	return s.Trend
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) SetPercentage(v int32) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate {
	s.Percentage = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) SetTrend(v string) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate {
	s.Trend = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout struct {
	// The timeout duration, in minutes.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout) SetTimeoutInMinutes(v int32) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout {
	s.TimeoutInMinutes = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished struct {
	// The not-completed time. Format: hh:mm. Valid values of hh: 0 to 47. Valid values of mm: 0 to 59.
	//
	// example:
	//
	// 12:00
	UnFinishedTime *string `json:"UnFinishedTime,omitempty" xml:"UnFinishedTime,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished) GetUnFinishedTime() *string {
	return s.UnFinishedTime
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished) SetUnFinishedTime(v string) *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished {
	s.UnFinishedTime = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionUnFinished) Validate() error {
	return dara.Validate(s)
}

type GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget struct {
	// The whitelist of monitored nodes.
	AllowTasks []*int64 `json:"AllowTasks,omitempty" xml:"AllowTasks,omitempty" type:"Repeated"`
	// The list of monitored object IDs.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The monitored object type. Valid values:
	//
	// - Task: node.
	//
	// - Baseline: baseline.
	//
	// - Project: workspace.
	//
	// - BizProcess: business process flow.
	//
	// example:
	//
	// Task
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) GoString() string {
	return s.String()
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) GetAllowTasks() []*int64 {
	return s.AllowTasks
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) GetType() *string {
	return s.Type
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) SetAllowTasks(v []*int64) *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget {
	s.AllowTasks = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) SetIds(v []*int64) *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget {
	s.Ids = v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) SetType(v string) *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget {
	s.Type = &v
	return s
}

func (s *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget) Validate() error {
	return dara.Validate(s)
}
