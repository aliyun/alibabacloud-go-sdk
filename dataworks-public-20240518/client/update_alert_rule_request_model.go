// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateAlertRuleRequest
	GetEnabled() *bool
	SetId(v int64) *UpdateAlertRuleRequest
	GetId() *int64
	SetName(v string) *UpdateAlertRuleRequest
	GetName() *string
	SetNotification(v *UpdateAlertRuleRequestNotification) *UpdateAlertRuleRequest
	GetNotification() *UpdateAlertRuleRequestNotification
	SetOwner(v string) *UpdateAlertRuleRequest
	GetOwner() *string
	SetTriggerCondition(v *UpdateAlertRuleRequestTriggerCondition) *UpdateAlertRuleRequest
	GetTriggerCondition() *UpdateAlertRuleRequestTriggerCondition
}

type UpdateAlertRuleRequest struct {
	// Specifies whether to enable the rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 105412
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// collection_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for the alert notification.
	Notification *UpdateAlertRuleRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// example:
	//
	// 193379****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert triggering condition.
	TriggerCondition *UpdateAlertRuleRequestTriggerCondition `json:"TriggerCondition,omitempty" xml:"TriggerCondition,omitempty" type:"Struct"`
}

func (s UpdateAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAlertRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateAlertRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlertRuleRequest) GetNotification() *UpdateAlertRuleRequestNotification {
	return s.Notification
}

func (s *UpdateAlertRuleRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateAlertRuleRequest) GetTriggerCondition() *UpdateAlertRuleRequestTriggerCondition {
	return s.TriggerCondition
}

func (s *UpdateAlertRuleRequest) SetEnabled(v bool) *UpdateAlertRuleRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetId(v int64) *UpdateAlertRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetName(v string) *UpdateAlertRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetNotification(v *UpdateAlertRuleRequestNotification) *UpdateAlertRuleRequest {
	s.Notification = v
	return s
}

func (s *UpdateAlertRuleRequest) SetOwner(v string) *UpdateAlertRuleRequest {
	s.Owner = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetTriggerCondition(v *UpdateAlertRuleRequestTriggerCondition) *UpdateAlertRuleRequest {
	s.TriggerCondition = v
	return s
}

func (s *UpdateAlertRuleRequest) Validate() error {
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

type UpdateAlertRuleRequestNotification struct {
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
	Receivers []*UpdateAlertRuleRequestNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
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

func (s UpdateAlertRuleRequestNotification) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestNotification) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestNotification) GetChannels() []*string {
	return s.Channels
}

func (s *UpdateAlertRuleRequestNotification) GetIntervalInMinutes() *int32 {
	return s.IntervalInMinutes
}

func (s *UpdateAlertRuleRequestNotification) GetMaximum() *int32 {
	return s.Maximum
}

func (s *UpdateAlertRuleRequestNotification) GetReceivers() []*UpdateAlertRuleRequestNotificationReceivers {
	return s.Receivers
}

func (s *UpdateAlertRuleRequestNotification) GetSilenceEndTime() *string {
	return s.SilenceEndTime
}

func (s *UpdateAlertRuleRequestNotification) GetSilenceStartTime() *string {
	return s.SilenceStartTime
}

func (s *UpdateAlertRuleRequestNotification) SetChannels(v []*string) *UpdateAlertRuleRequestNotification {
	s.Channels = v
	return s
}

func (s *UpdateAlertRuleRequestNotification) SetIntervalInMinutes(v int32) *UpdateAlertRuleRequestNotification {
	s.IntervalInMinutes = &v
	return s
}

func (s *UpdateAlertRuleRequestNotification) SetMaximum(v int32) *UpdateAlertRuleRequestNotification {
	s.Maximum = &v
	return s
}

func (s *UpdateAlertRuleRequestNotification) SetReceivers(v []*UpdateAlertRuleRequestNotificationReceivers) *UpdateAlertRuleRequestNotification {
	s.Receivers = v
	return s
}

func (s *UpdateAlertRuleRequestNotification) SetSilenceEndTime(v string) *UpdateAlertRuleRequestNotification {
	s.SilenceEndTime = &v
	return s
}

func (s *UpdateAlertRuleRequestNotification) SetSilenceStartTime(v string) *UpdateAlertRuleRequestNotification {
	s.SilenceStartTime = &v
	return s
}

func (s *UpdateAlertRuleRequestNotification) Validate() error {
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

type UpdateAlertRuleRequestNotificationReceivers struct {
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
	// TaskOwner
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The alert recipients.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s UpdateAlertRuleRequestNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestNotificationReceivers) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *UpdateAlertRuleRequestNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *UpdateAlertRuleRequestNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *UpdateAlertRuleRequestNotificationReceivers) SetExtension(v string) *UpdateAlertRuleRequestNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *UpdateAlertRuleRequestNotificationReceivers) SetReceiverType(v string) *UpdateAlertRuleRequestNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *UpdateAlertRuleRequestNotificationReceivers) SetReceiverValues(v []*string) *UpdateAlertRuleRequestNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *UpdateAlertRuleRequestNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerCondition struct {
	// The extended information about the rule. This parameter is required for specific types of alerts.
	Extension *UpdateAlertRuleRequestTriggerConditionExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The monitored objects.
	Target *UpdateAlertRuleRequestTriggerConditionTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
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
	// ERROR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerCondition) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerCondition) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerCondition) GetExtension() *UpdateAlertRuleRequestTriggerConditionExtension {
	return s.Extension
}

func (s *UpdateAlertRuleRequestTriggerCondition) GetTarget() *UpdateAlertRuleRequestTriggerConditionTarget {
	return s.Target
}

func (s *UpdateAlertRuleRequestTriggerCondition) GetType() *string {
	return s.Type
}

func (s *UpdateAlertRuleRequestTriggerCondition) SetExtension(v *UpdateAlertRuleRequestTriggerConditionExtension) *UpdateAlertRuleRequestTriggerCondition {
	s.Extension = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerCondition) SetTarget(v *UpdateAlertRuleRequestTriggerConditionTarget) *UpdateAlertRuleRequestTriggerCondition {
	s.Target = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerCondition) SetType(v string) *UpdateAlertRuleRequestTriggerCondition {
	s.Type = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerCondition) Validate() error {
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

type UpdateAlertRuleRequestTriggerConditionExtension struct {
	// The configuration for an alert of the CycleUnfinished type.
	CycleUnfinished *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished `json:"CycleUnfinished,omitempty" xml:"CycleUnfinished,omitempty" type:"Struct"`
	// The configuration for an alert of the Error type.
	Error *UpdateAlertRuleRequestTriggerConditionExtensionError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorCount type.
	InstanceErrorCount *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount `json:"InstanceErrorCount,omitempty" xml:"InstanceErrorCount,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorPercentage type.
	InstanceErrorPercentage *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage `json:"InstanceErrorPercentage,omitempty" xml:"InstanceErrorPercentage,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceTransferFluctuate type.
	InstanceTransferFluctuate *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate `json:"InstanceTransferFluctuate,omitempty" xml:"InstanceTransferFluctuate,omitempty" type:"Struct"`
	// The configuration for an alert of the Timeout type.
	Timeout *UpdateAlertRuleRequestTriggerConditionExtensionTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
	// The configuration for an alert of the UnFinished type.
	UnFinished *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished `json:"UnFinished,omitempty" xml:"UnFinished,omitempty" type:"Struct"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtension) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtension) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetCycleUnfinished() *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished {
	return s.CycleUnfinished
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetError() *UpdateAlertRuleRequestTriggerConditionExtensionError {
	return s.Error
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetInstanceErrorCount() *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount {
	return s.InstanceErrorCount
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetInstanceErrorPercentage() *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage {
	return s.InstanceErrorPercentage
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetInstanceTransferFluctuate() *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate {
	return s.InstanceTransferFluctuate
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetTimeout() *UpdateAlertRuleRequestTriggerConditionExtensionTimeout {
	return s.Timeout
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) GetUnFinished() *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished {
	return s.UnFinished
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetCycleUnfinished(v *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.CycleUnfinished = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetError(v *UpdateAlertRuleRequestTriggerConditionExtensionError) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.Error = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetInstanceErrorCount(v *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.InstanceErrorCount = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetInstanceErrorPercentage(v *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.InstanceErrorPercentage = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetInstanceTransferFluctuate(v *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.InstanceTransferFluctuate = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetTimeout(v *UpdateAlertRuleRequestTriggerConditionExtensionTimeout) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.Timeout = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) SetUnFinished(v *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished) *UpdateAlertRuleRequestTriggerConditionExtension {
	s.UnFinished = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtension) Validate() error {
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

type UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished struct {
	// The configurations of the scheduling cycle and timeout period of the instance.
	CycleAndTime []*UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime `json:"CycleAndTime,omitempty" xml:"CycleAndTime,omitempty" type:"Repeated"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) GetCycleAndTime() []*UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	return s.CycleAndTime
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) SetCycleAndTime(v []*UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished {
	s.CycleAndTime = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) Validate() error {
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

type UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime struct {
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
	// 01:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetCycleId() *int32 {
	return s.CycleId
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetTime() *string {
	return s.Time
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetCycleId(v int32) *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.CycleId = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetTime(v string) *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.Time = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionExtensionError struct {
	// Specifies whether to trigger an alert if a batch synchronization task is automatically rerun upon a failure.
	//
	// example:
	//
	// false
	AutoRerunAlertEnabled *bool `json:"AutoRerunAlertEnabled,omitempty" xml:"AutoRerunAlertEnabled,omitempty"`
	// The IDs of the real-time computing tasks. This parameter is required when you monitor real-time computing tasks.
	StreamTaskIds []*int64 `json:"StreamTaskIds,omitempty" xml:"StreamTaskIds,omitempty" type:"Repeated"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionError) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionError) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionError) GetAutoRerunAlertEnabled() *bool {
	return s.AutoRerunAlertEnabled
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionError) GetStreamTaskIds() []*int64 {
	return s.StreamTaskIds
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionError) SetAutoRerunAlertEnabled(v bool) *UpdateAlertRuleRequestTriggerConditionExtensionError {
	s.AutoRerunAlertEnabled = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionError) SetStreamTaskIds(v []*int64) *UpdateAlertRuleRequestTriggerConditionExtensionError {
	s.StreamTaskIds = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionError) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount struct {
	// The maximum number of instances on which an error occurs. Valid values: [1,10000].
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) GetCount() *int32 {
	return s.Count
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) SetCount(v int32) *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount {
	s.Count = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage struct {
	// The maximum percentage of instances on which an error occurs in the workspace to the total number of instances. Valid values: [1-100].
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) GetPercentage() *int32 {
	return s.Percentage
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) SetPercentage(v int32) *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage {
	s.Percentage = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate struct {
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
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) GetPercentage() *int32 {
	return s.Percentage
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) GetTrend() *string {
	return s.Trend
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) SetPercentage(v int32) *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate {
	s.Percentage = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) SetTrend(v string) *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate {
	s.Trend = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionExtensionTimeout struct {
	// The timeout period. Unit: minutes.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionTimeout) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionTimeout) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionTimeout) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionTimeout) SetTimeoutInMinutes(v int32) *UpdateAlertRuleRequestTriggerConditionExtensionTimeout {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionTimeout) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionExtensionUnFinished struct {
	// The latest completion time of the instance. The period is in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
	//
	// example:
	//
	// 12:00
	UnFinishedTime *string `json:"UnFinishedTime,omitempty" xml:"UnFinishedTime,omitempty"`
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionUnFinished) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionExtensionUnFinished) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished) GetUnFinishedTime() *string {
	return s.UnFinishedTime
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished) SetUnFinishedTime(v string) *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished {
	s.UnFinishedTime = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionExtensionUnFinished) Validate() error {
	return dara.Validate(s)
}

type UpdateAlertRuleRequestTriggerConditionTarget struct {
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

func (s UpdateAlertRuleRequestTriggerConditionTarget) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequestTriggerConditionTarget) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) GetAllowTasks() []*int64 {
	return s.AllowTasks
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) GetType() *string {
	return s.Type
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) SetAllowTasks(v []*int64) *UpdateAlertRuleRequestTriggerConditionTarget {
	s.AllowTasks = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) SetIds(v []*int64) *UpdateAlertRuleRequestTriggerConditionTarget {
	s.Ids = v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) SetType(v string) *UpdateAlertRuleRequestTriggerConditionTarget {
	s.Type = &v
	return s
}

func (s *UpdateAlertRuleRequestTriggerConditionTarget) Validate() error {
	return dara.Validate(s)
}
