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
	// The information about the rule.
	AlertRule *GetAlertRuleResponseBodyAlertRule `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
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
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 16035
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// error_rule
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for the alert notification.
	Notification *GetAlertRuleResponseBodyAlertRuleNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// example:
	//
	// 279961421580845157
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert triggering condition.
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
	// The alert notification channels.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The interval at which an alert notification is sent. Unit: minutes. Valid values: [5,10000].
	//
	// example:
	//
	// 30
	IntervalInMinutes *int32 `json:"IntervalInMinutes,omitempty" xml:"IntervalInMinutes,omitempty"`
	// The maximum number of times an alert notification can be sent within a calendar day. Valid values: [1, 10000].
	//
	// example:
	//
	// 3
	Maximum *int32 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// The alert recipients.
	Receivers []*GetAlertRuleResponseBodyAlertRuleNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The end time for silence. The time is in the HH:mm:ss format.
	//
	// example:
	//
	// 00:00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The start time for silence. The time is in the HH:mm:ss format.
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
	// The additional configuration of the alert recipient. If the ReceiverType parameter is set to DingdingUrl, you can set this parameter to {"atAll":true} to remind all members in a DingTalk group.
	//
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The type of the alert recipient. Valid valves:
	//
	// 	- AliUid: Alibaba Cloud account ID.
	//
	// 	- Shift Schedules: the personnel in a shift schedule.
	//
	// 	- TaskOwner: the task owner. The task owner can receive custom alerts and event alerts.
	//
	// 	- Owner: the baseline owner. The baseline owner can receive baseline alerts.
	//
	// 	- WebhookUrl: URL of a custom webhook.
	//
	// 	- DingdingUrl: DingTalk webhook URL.
	//
	// 	- FeishuUrl: Lark webhook URL.
	//
	// 	- WeixinUrl: WeCom webhook URL.
	//
	// example:
	//
	// WebhookUrl
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
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
	// The extended information about the rule. This parameter is required for specific types of alerts.
	Extension *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The monitored objects.
	Target *GetAlertRuleResponseBodyAlertRuleTriggerConditionTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The alert type. Valid values:
	//
	// 	- Finished: An instance is successfully run.
	//
	// 	- UnFinished: An instance does not finish running before a specified point in time.
	//
	// 	- Error: An error occurs on an instance.
	//
	// 	- CycleUnfinished: An instance does not finish running as expected within a specific cycle.
	//
	// 	- Timeout: An instance times out.
	//
	// 	- InstanceTransferComplete: An instance is generated by the auto triggered node.
	//
	// 	- InstanceTransferFluctuate: The number of generated instances fluctuates.
	//
	// 	- ExhaustedError: An error persists after an instance is automatically rerun.
	//
	// 	- InstanceKeyword: An instance with errors contains specified keywords.
	//
	// 	- InstanceErrorCount: The number of instances on which an error occurs reaches a specified threshold.
	//
	// 	- InstanceErrorPercentage: The proportion of instances on which an error occurs in the workspace to the total number of instances reaches a specified threshold.
	//
	// 	- ResourceGroupPercentage: The usage rate of the resource group reaches a specified threshold.
	//
	// 	- ResourceGroupWaitCount: The number of instances that are waiting for resources in the resource group reaches a specified threshold.
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
	// The configuration for an alert of the CycleUnfinished type.
	CycleUnfinished *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionCycleUnfinished `json:"CycleUnfinished,omitempty" xml:"CycleUnfinished,omitempty" type:"Struct"`
	// The configuration for an alert of the Error type.
	Error *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorCount type.
	InstanceErrorCount *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorCount `json:"InstanceErrorCount,omitempty" xml:"InstanceErrorCount,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorPercentage type.
	InstanceErrorPercentage *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceErrorPercentage `json:"InstanceErrorPercentage,omitempty" xml:"InstanceErrorPercentage,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceTransferFluctuate type.
	InstanceTransferFluctuate *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionInstanceTransferFluctuate `json:"InstanceTransferFluctuate,omitempty" xml:"InstanceTransferFluctuate,omitempty" type:"Struct"`
	// The configuration for an alert of the Timeout type.
	Timeout *GetAlertRuleResponseBodyAlertRuleTriggerConditionExtensionTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
	// The configuration for an alert of the UnFinished type.
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
	// The configurations of the scheduling cycle and timeout period of the instance.
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
	// The ID of the scheduling cycle of the instance. Valid values: [1,288].
	//
	// example:
	//
	// 1
	CycleId *int32 `json:"CycleId,omitempty" xml:"CycleId,omitempty"`
	// The latest completion time of the instance within the scheduling cycle. The time is in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
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
	// Indicates whether an alert is triggered if a batch synchronization task is automatically rerun upon a failure.
	//
	// example:
	//
	// false
	AutoRerunAlertEnabled *bool `json:"AutoRerunAlertEnabled,omitempty" xml:"AutoRerunAlertEnabled,omitempty"`
	// The IDs of the real-time computing tasks. This parameter is required when you monitor real-time computing tasks.
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
	// The maximum number of instances on which an error occurs. Valid values: [1,10000].
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
	// The maximum percentage of instances on which an error occurs in the workspace to the total number of instances. Valid values: [1-100].
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
	// The maximum percentage of fluctuation in the number of auto triggered node instances that are generated in your workspace. Valid values: [1-100].
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The way in which the number of auto triggered node instances that are generated in your workspace fluctuates. Valid values:
	//
	// 	- abs: the absolute value. The number of instances increases or decreases.
	//
	// 	- increase: The number of instances increases.
	//
	// 	- decrease: The number of instances decreases.
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
	// The timeout period. Unit: minutes. Valid values: [1, 21600].
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
	// The latest completion time of the instance. The period is in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
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
	// The nodes that are not to be monitored.
	AllowTasks []*int64 `json:"AllowTasks,omitempty" xml:"AllowTasks,omitempty" type:"Repeated"`
	// The IDs of monitored objects.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The type of the monitored objects. Valid values:
	//
	// 	- Task: node
	//
	// 	- Baseline: baseline
	//
	// 	- project: workspace
	//
	// 	- BizProcess: workflow
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
