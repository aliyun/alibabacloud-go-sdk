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
	// Indicates whether the rule is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// xm_create_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration for the alert notification.
	Notification *CreateAlertRuleRequestNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account used by the owner of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 279114181716147735
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The alert triggering condition.
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
	return dara.Validate(s)
}

type CreateAlertRuleRequestNotification struct {
	// The alert notification channels.
	//
	// This parameter is required.
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
	//
	// This parameter is required.
	Receivers []*CreateAlertRuleRequestNotificationReceivers `json:"Receivers,omitempty" xml:"Receivers,omitempty" type:"Repeated"`
	// The end time for silence. The time is in the HH:mm format.
	//
	// example:
	//
	// 00:00:00
	SilenceEndTime *string `json:"SilenceEndTime,omitempty" xml:"SilenceEndTime,omitempty"`
	// The start time for silence. The time is in the HH:mm format.
	//
	// example:
	//
	// 00:00:00
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
	return dara.Validate(s)
}

type CreateAlertRuleRequestNotificationReceivers struct {
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
	// The ID of the alert recipient.
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
	// The extended information about the rule. This parameter is required for specific types of alerts.
	Extension *CreateAlertRuleRequestTriggerConditionExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The monitored objects.
	Target *CreateAlertRuleRequestTriggerConditionTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtension struct {
	// The configuration for an alert of the CycleUnfinished type.
	CycleUnfinished *CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished `json:"CycleUnfinished,omitempty" xml:"CycleUnfinished,omitempty" type:"Struct"`
	// The configuration for an alert of the Error type.
	Error *CreateAlertRuleRequestTriggerConditionExtensionError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorCount type.
	InstanceErrorCount *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorCount `json:"InstanceErrorCount,omitempty" xml:"InstanceErrorCount,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceErrorPercentage type.
	InstanceErrorPercentage *CreateAlertRuleRequestTriggerConditionExtensionInstanceErrorPercentage `json:"InstanceErrorPercentage,omitempty" xml:"InstanceErrorPercentage,omitempty" type:"Struct"`
	// The configuration for an alert of the InstanceTransferFluctuate type.
	InstanceTransferFluctuate *CreateAlertRuleRequestTriggerConditionExtensionInstanceTransferFluctuate `json:"InstanceTransferFluctuate,omitempty" xml:"InstanceTransferFluctuate,omitempty" type:"Struct"`
	// The configuration for an alert of the Timeout type.
	Timeout *CreateAlertRuleRequestTriggerConditionExtensionTimeout `json:"Timeout,omitempty" xml:"Timeout,omitempty" type:"Struct"`
	// The configuration for an alert of the UnFinished type.
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
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinished struct {
	// The configurations of the scheduling cycle and timeout period of the instance.
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
	return dara.Validate(s)
}

type CreateAlertRuleRequestTriggerConditionExtensionCycleUnfinishedCycleAndTime struct {
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
	// Specifies whether to trigger an alert if a batch synchronization task is automatically rerun upon a failure.
	//
	// example:
	//
	// false
	AutoRerunAlertEnabled *bool `json:"AutoRerunAlertEnabled,omitempty" xml:"AutoRerunAlertEnabled,omitempty"`
	// The IDs of the real-time computing tasks. This parameter is required when you monitor real-time computing tasks.
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
	// The maximum number of instances on which an error occurs. Valid values: [1,10000].
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
	// The maximum percentage of instances on which an error occurs in the workspace to the total number of instances. Valid values: [1-100].
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
	// The timeout period. Unit: minutes. Valid values: [1, 21600].
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
	// The latest completion time of the instance. The period is in the hh:mm format. Valid values of hh: [0,47]. Valid values of mm: [0,59].
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
	// 	- Project: workspace
	//
	// 	- BizProcess: workflow
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
