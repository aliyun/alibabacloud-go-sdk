// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *CreateAlertRuleRequest
	GetEnabled() *bool
	SetName(v string) *CreateAlertRuleRequest
	GetName() *string
	SetNotification(v *CreateAlertRuleRequestNotification) *CreateAlertRuleRequest
	GetNotification() *CreateAlertRuleRequestNotification
	SetOwner(v string) *CreateAlertRuleRequest
	GetOwner() *string
	SetTriggerCondition(v *CreateAlertRuleRequestTriggerCondition) *CreateAlertRuleRequest
	GetTriggerCondition() *CreateAlertRuleRequestTriggerCondition
}

type CreateAlertRuleRequest struct {
	// Specifies whether the alert rule is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the custom rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert notification configuration.
	Notification *CreateAlertRuleRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The Alibaba Cloud UID of the owner of the custom rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 279114181****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The condition that triggers the alert.
	//
	// This parameter is required.
	TriggerCondition *CreateAlertRuleRequestTriggerCondition `json:"TriggerCondition,omitempty" xml:"TriggerCondition,omitempty" type:"Struct"`
}

func (s CreateAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAlertRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlertRuleRequest) GetNotification() *CreateAlertRuleRequestNotification {
	return s.Notification
}

func (s *CreateAlertRuleRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateAlertRuleRequest) GetTriggerCondition() *CreateAlertRuleRequestTriggerCondition {
	return s.TriggerCondition
}

func (s *CreateAlertRuleRequest) SetEnabled(v bool) *CreateAlertRuleRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAlertRuleRequest) SetName(v string) *CreateAlertRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateAlertRuleRequest) SetNotification(v *CreateAlertRuleRequestNotification) *CreateAlertRuleRequest {
	s.Notification = v
	return s
}

func (s *CreateAlertRuleRequest) SetOwner(v string) *CreateAlertRuleRequest {
	s.Owner = &v
	return s
}

func (s *CreateAlertRuleRequest) SetTriggerCondition(v *CreateAlertRuleRequestTriggerCondition) *CreateAlertRuleRequest {
	s.TriggerCondition = v
	return s
}

func (s *CreateAlertRuleRequest) Validate() error {
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

type CreateAlertRuleRequestNotification struct {
	// The list of alert channels.
	//
	// This parameter is required.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The alert interval, in minutes. Valid values: 5 to 10000.
	//
	// example:
	//
	// 30
	IntervalInMinutes *int32 `json:"IntervalInMinutes,omitempty" xml:"IntervalInMinutes,omitempty"`
	// The maximum number of alerts within a calendar year. Valid values: 1 to 10000.
	//
	// example:
	//
	// 3
	Maximum *int32 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// The alert recipients.
	//
	// This parameter is required.
	Receivers []*CreateAlertRuleRequestNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The end time of the alert silence period, in the format of HH:mm.
	//
	// example:
	//
	// 00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The start time of the alert silence period, in the format of HH:mm.
	//
	// example:
	//
	// 00:00
	SilenceStartTime *string `json:"SilenceStartTime,omitempty" xml:"SilenceStartTime,omitempty"`
}

func (s CreateAlertRuleRequestNotification) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestNotification) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestNotification) GetChannels() []*string {
	return s.Channels
}

func (s *CreateAlertRuleRequestNotification) GetIntervalInMinutes() *int32 {
	return s.IntervalInMinutes
}

func (s *CreateAlertRuleRequestNotification) GetMaximum() *int32 {
	return s.Maximum
}

func (s *CreateAlertRuleRequestNotification) GetReceivers() []*CreateAlertRuleRequestNotificationReceivers {
	return s.Receivers
}

func (s *CreateAlertRuleRequestNotification) GetSilenceEndTime() *string {
	return s.SilenceEndTime
}

func (s *CreateAlertRuleRequestNotification) GetSilenceStartTime() *string {
	return s.SilenceStartTime
}

func (s *CreateAlertRuleRequestNotification) SetChannels(v []*string) *CreateAlertRuleRequestNotification {
	s.Channels = v
	return s
}

func (s *CreateAlertRuleRequestNotification) SetIntervalInMinutes(v int32) *CreateAlertRuleRequestNotification {
	s.IntervalInMinutes = &v
	return s
}

func (s *CreateAlertRuleRequestNotification) SetMaximum(v int32) *CreateAlertRuleRequestNotification {
	s.Maximum = &v
	return s
}

func (s *CreateAlertRuleRequestNotification) SetReceivers(v []*CreateAlertRuleRequestNotificationReceivers) *CreateAlertRuleRequestNotification {
	s.Receivers = v
	return s
}

func (s *CreateAlertRuleRequestNotification) SetSilenceEndTime(v string) *CreateAlertRuleRequestNotification {
	s.SilenceEndTime = &v
	return s
}

func (s *CreateAlertRuleRequestNotification) SetSilenceStartTime(v string) *CreateAlertRuleRequestNotification {
	s.SilenceStartTime = &v
	return s
}

func (s *CreateAlertRuleRequestNotification) Validate() error {
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

type CreateAlertRuleRequestNotificationReceivers struct {
	// The additional configuration required for the alert recipient. If ReceiverType is DingdingUrl, you can set {"atAll":true} to @ all members.
	//
	// example:
	//
	// {"atAll":true}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The alert recipient type. Valid values:
	//
	// - AliUid: Alibaba Cloud UID
	//
	// - ShiftSchedule: shift schedule
	//
	// - TaskOwner: node owner, applicable to custom alerting and event alerting
	//
	// - Owner: owner, applicable to baseline alerting
	//
	// - WebhookUrl: custom webhook URL
	//
	// - DingdingUrl: DingTalk webhook URL
	//
	// - FeishuUrl: Lark webhook URL
	//
	// - WeixinUrl: WeCom webhook URL
	//
	// example:
	//
	// TaskOwner
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The values of the alert recipient.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s CreateAlertRuleRequestNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestNotificationReceivers) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestNotificationReceivers) GetExtension() *string {
	return s.Extension
}

func (s *CreateAlertRuleRequestNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *CreateAlertRuleRequestNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *CreateAlertRuleRequestNotificationReceivers) SetExtension(v string) *CreateAlertRuleRequestNotificationReceivers {
	s.Extension = &v
	return s
}

func (s *CreateAlertRuleRequestNotificationReceivers) SetReceiverType(v string) *CreateAlertRuleRequestNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *CreateAlertRuleRequestNotificationReceivers) SetReceiverValues(v []*string) *CreateAlertRuleRequestNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *CreateAlertRuleRequestNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerCondition struct {
	// The extension information. This parameter is required for certain trigger condition configurations.
	Extension *CreateAlertRuleRequestTriggerConditionExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The monitored object.
	Target *CreateAlertRuleRequestTriggerConditionTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The type of alert trigger. Valid values:
	//
	// - Finished: Instance completed.
	//
	// - UnFinished: Instance not completed.
	//
	// - Error: Instance error.
	//
	// - CycleUnfinished: Instance cycle not completed.
	//
	// - Timeout: Instance timeout.
	//
	// - InstanceTransferComplete: Node-to-instance conversion completed.
	//
	// - InstanceTransferFluctuate: Instance count fluctuation.
	//
	// - ExhaustedError: Error persists after automatic reruns.
	//
	// - InstanceKeyword: Error instance contains keyword.
	//
	// - InstanceErrorCount: Number of error instances.
	//
	// - InstanceErrorPercentage: Percentage of error instances.
	//
	// - ResourceGroupPercentage: Resource group utilization.
	//
	// - ResourceGroupWaitCount: Number of instances waiting for resources in the resource group.
	//
	// example:
	//
	// Error
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAlertRuleRequestTriggerCondition) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerCondition) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerCondition) GetExtension() *CreateAlertRuleRequestTriggerConditionExtension {
	return s.Extension
}

func (s *CreateAlertRuleRequestTriggerCondition) GetTarget() *CreateAlertRuleRequestTriggerConditionTarget {
	return s.Target
}

func (s *CreateAlertRuleRequestTriggerCondition) GetType() *string {
	return s.Type
}

func (s *CreateAlertRuleRequestTriggerCondition) SetExtension(v *CreateAlertRuleRequestTriggerConditionExtension) *CreateAlertRuleRequestTriggerCondition {
	s.Extension = v
	return s
}

func (s *CreateAlertRuleRequestTriggerCondition) SetTarget(v *CreateAlertRuleRequestTriggerConditionTarget) *CreateAlertRuleRequestTriggerCondition {
	s.Target = v
	return s
}

func (s *CreateAlertRuleRequestTriggerCondition) SetType(v string) *CreateAlertRuleRequestTriggerCondition {
	s.Type = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerCondition) Validate() error {
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

type CreateAlertRuleRequestTriggerConditionExtension struct {
	// The cycle unfinished alert configuration.
	CycleUnfinished *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished `json:"CycleUnfinished,omitempty" xml:"CycleUnfinished,omitempty" type:"Struct"`
	// The error alert configuration.
	Error *CreateAlertRuleRequestTriggerConditionExtensionError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The instance error count alert configuration.
	InstanceErrorCount *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount `json:"InstanceErrorCount,omitempty" xml:"InstanceErrorCount,omitempty" type:"Struct"`
	// The instance error percentage alert configuration.
	InstanceErrorPercentage *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage `json:"InstanceErrorPercentage,omitempty" xml:"InstanceErrorPercentage,omitempty" type:"Struct"`
	// The instance transfer fluctuation alert configuration.
	InstanceTransferFluctuate *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate `json:"InstanceTransferFluctuate,omitempty" xml:"InstanceTransferFluctuate,omitempty" type:"Struct"`
	// The timeout alert configuration.
	Timeout *CreateAlertRuleRequestTriggerConditionExtensionTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
	// The unfinished alert configuration.
	UnFinished *CreateAlertRuleRequestTriggerConditionExtensionUnFinished `json:"UnFinished,omitempty" xml:"UnFinished,omitempty" type:"Struct"`
}

func (s CreateAlertRuleRequestTriggerConditionExtension) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtension) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetCycleUnfinished() *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished {
	return s.CycleUnfinished
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetError() *CreateAlertRuleRequestTriggerConditionExtensionError {
	return s.Error
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetInstanceErrorCount() *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount {
	return s.InstanceErrorCount
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetInstanceErrorPercentage() *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage {
	return s.InstanceErrorPercentage
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetInstanceTransferFluctuate() *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate {
	return s.InstanceTransferFluctuate
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetTimeout() *CreateAlertRuleRequestTriggerConditionExtensionTimeout {
	return s.Timeout
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) GetUnFinished() *CreateAlertRuleRequestTriggerConditionExtensionUnFinished {
	return s.UnFinished
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetCycleUnfinished(v *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) *CreateAlertRuleRequestTriggerConditionExtension {
	s.CycleUnfinished = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetError(v *CreateAlertRuleRequestTriggerConditionExtensionError) *CreateAlertRuleRequestTriggerConditionExtension {
	s.Error = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetInstanceErrorCount(v *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) *CreateAlertRuleRequestTriggerConditionExtension {
	s.InstanceErrorCount = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetInstanceErrorPercentage(v *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) *CreateAlertRuleRequestTriggerConditionExtension {
	s.InstanceErrorPercentage = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetInstanceTransferFluctuate(v *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) *CreateAlertRuleRequestTriggerConditionExtension {
	s.InstanceTransferFluctuate = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetTimeout(v *CreateAlertRuleRequestTriggerConditionExtensionTimeout) *CreateAlertRuleRequestTriggerConditionExtension {
	s.Timeout = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) SetUnFinished(v *CreateAlertRuleRequestTriggerConditionExtensionUnFinished) *CreateAlertRuleRequestTriggerConditionExtension {
	s.UnFinished = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtension) Validate() error {
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

type CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished struct {
	// The list of cycle and time configurations.
	CycleAndTime []*CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime `json:"CycleAndTime,omitempty" xml:"CycleAndTime,omitempty" type:"Repeated"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) GetCycleAndTime() []*CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	return s.CycleAndTime
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) SetCycleAndTime(v []*CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished {
	s.CycleAndTime = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished) Validate() error {
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

type CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime struct {
	// The cycle ID. Valid values: 1 to 288.
	//
	// example:
	//
	// 1
	CycleId *int32 `json:"CycleId,omitempty" xml:"CycleId,omitempty"`
	// The timeout time, in the format of hh:mm. Valid values of hh: 0 to 47. Valid values of mm: 0 to 59.
	//
	// example:
	//
	// 12:00
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetCycleId() *int32 {
	return s.CycleId
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) GetTime() *string {
	return s.Time
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetCycleId(v int32) *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.CycleId = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) SetTime(v string) *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime {
	s.Time = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionError struct {
	// Specifies whether to generate an alert when an offline task is automatically rerun due to failure.
	//
	// example:
	//
	// false
	AutoRerunAlertEnabled *bool `json:"AutoRerunAlertEnabled,omitempty" xml:"AutoRerunAlertEnabled,omitempty"`
	// The IDs of real-time computing tasks to monitor.
	StreamTaskIds []*int64 `json:"StreamTaskIds,omitempty" xml:"StreamTaskIds,omitempty" type:"Repeated"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionError) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionError) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionError) GetAutoRerunAlertEnabled() *bool {
	return s.AutoRerunAlertEnabled
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionError) GetStreamTaskIds() []*int64 {
	return s.StreamTaskIds
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionError) SetAutoRerunAlertEnabled(v bool) *CreateAlertRuleRequestTriggerConditionExtensionError {
	s.AutoRerunAlertEnabled = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionError) SetStreamTaskIds(v []*int64) *CreateAlertRuleRequestTriggerConditionExtensionError {
	s.StreamTaskIds = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionError) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount struct {
	// The number of error instances. Valid values: 1 to 10000.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) GetCount() *int32 {
	return s.Count
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) SetCount(v int32) *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount {
	s.Count = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage struct {
	// The percentage of error instances. Valid values: 1 to 100.
	//
	// example:
	//
	// 5
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) SetPercentage(v int32) *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage {
	s.Percentage = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate struct {
	// The percentage of instance transfer fluctuation. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The type of instance transfer fluctuation. Valid values:
	//
	// - abs: absolute value
	//
	// - increase: increase
	//
	// - decrease: decrease
	//
	// example:
	//
	// abs
	Trend *string `json:"Trend,omitempty" xml:"Trend,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) GetPercentage() *int32 {
	return s.Percentage
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) GetTrend() *string {
	return s.Trend
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) SetPercentage(v int32) *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate {
	s.Percentage = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) SetTrend(v string) *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate {
	s.Trend = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionTimeout struct {
	// The timeout duration, in minutes. Valid values: 1 to 21600.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionTimeout) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionTimeout) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionTimeout) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionTimeout) SetTimeoutInMinutes(v int32) *CreateAlertRuleRequestTriggerConditionExtensionTimeout {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionTimeout) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionUnFinished struct {
	// The unfinished time, in the format of hh:mm. Valid values of hh: 0 to 47. Valid values of mm: 0 to 59.
	//
	// example:
	//
	// 30:00
	UnFinishedTime *string `json:"UnFinishedTime,omitempty" xml:"UnFinishedTime,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionExtensionUnFinished) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionExtensionUnFinished) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionUnFinished) GetUnFinishedTime() *string {
	return s.UnFinishedTime
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionUnFinished) SetUnFinishedTime(v string) *CreateAlertRuleRequestTriggerConditionExtensionUnFinished {
	s.UnFinishedTime = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionExtensionUnFinished) Validate() error {
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionTarget struct {
	// The whitelist of monitored tasks.
	AllowTasks []*int64 `json:"AllowTasks,omitempty" xml:"AllowTasks,omitempty" type:"Repeated"`
	// The list of monitored object IDs.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The monitored object type. Valid values:
	//
	// - Task: node
	//
	// - Baseline: baseline
	//
	// - Project: workspace
	//
	// - BizProcess: business process
	//
	// example:
	//
	// Task
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAlertRuleRequestTriggerConditionTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertRuleRequestTriggerConditionTarget) GoString() string {
	return s.String()
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) GetAllowTasks() []*int64 {
	return s.AllowTasks
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) GetIds() []*int64 {
	return s.Ids
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) GetType() *string {
	return s.Type
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) SetAllowTasks(v []*int64) *CreateAlertRuleRequestTriggerConditionTarget {
	s.AllowTasks = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) SetIds(v []*int64) *CreateAlertRuleRequestTriggerConditionTarget {
	s.Ids = v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) SetType(v string) *CreateAlertRuleRequestTriggerConditionTarget {
	s.Type = &v
	return s
}

func (s *CreateAlertRuleRequestTriggerConditionTarget) Validate() error {
	return dara.Validate(s)
}
