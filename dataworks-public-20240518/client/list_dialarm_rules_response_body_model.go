// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIAlarmRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDIAlarmRulesResponseBodyPagingInfo) *ListDIAlarmRulesResponseBody
	GetPagingInfo() *ListDIAlarmRulesResponseBodyPagingInfo
	SetRequestId(v string) *ListDIAlarmRulesResponseBody
	GetRequestId() *string
}

type ListDIAlarmRulesResponseBody struct {
	// The pagination information.
	PagingInfo *ListDIAlarmRulesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 74C2FECD-5B3A-554A-BCF5-351A36DE9815
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIAlarmRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBody) GetPagingInfo() *ListDIAlarmRulesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDIAlarmRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIAlarmRulesResponseBody) SetPagingInfo(v *ListDIAlarmRulesResponseBodyPagingInfo) *ListDIAlarmRulesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIAlarmRulesResponseBody) SetRequestId(v string) *ListDIAlarmRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyPagingInfo struct {
	// The alert rules.
	DIJobAlarmRules []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules `json:"DIJobAlarmRules,omitempty" xml:"DIJobAlarmRules,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 90
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) GetDIJobAlarmRules() []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	return s.DIJobAlarmRules
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetDIJobAlarmRules(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.DIJobAlarmRules = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetPageSize(v int64) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIAlarmRulesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfo) Validate() error {
	if s.DIJobAlarmRules != nil {
		for _, item := range s.DIJobAlarmRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 72402
	DIAlarmRuleId *int64 `json:"DIAlarmRuleId,omitempty" xml:"DIAlarmRuleId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 32594
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The description of the alert rule.
	//
	// example:
	//
	// rule descrition
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the alert rule is enabled. Valid values: True and False.
	//
	// example:
	//
	// True
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// 72402
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric type in the alert rule. Valid values:
	//
	// 	- Heartbeat
	//
	// 	- FailoverCount
	//
	// 	- Delay
	//
	// 	- DdlReport
	//
	// 	- ResourceUtilization
	//
	// example:
	//
	// Heartbeat
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// rule_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert notification method and recipient settings.
	NotificationSettings *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings `json:"NotificationSettings,omitempty" xml:"NotificationSettings,omitempty" type:"Struct"`
	// The conditions that are used to trigger the alert rule.
	TriggerConditions []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions `json:"TriggerConditions,omitempty" xml:"TriggerConditions,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetDIAlarmRuleId() *int64 {
	return s.DIAlarmRuleId
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetDescription() *string {
	return s.Description
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetId() *int64 {
	return s.Id
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetMetricType() *string {
	return s.MetricType
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetName() *string {
	return s.Name
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetNotificationSettings() *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	return s.NotificationSettings
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) GetTriggerConditions() []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	return s.TriggerConditions
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetDIAlarmRuleId(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.DIAlarmRuleId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetDIJobId(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.DIJobId = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetDescription(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Description = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetEnabled(v bool) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Enabled = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetId(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Id = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetMetricType(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.MetricType = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetName(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.Name = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetNotificationSettings(v *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.NotificationSettings = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) SetTriggerConditions(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules {
	s.TriggerConditions = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRules) Validate() error {
	if s.NotificationSettings != nil {
		if err := s.NotificationSettings.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConditions != nil {
		for _, item := range s.TriggerConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings struct {
	// Deprecated
	//
	// This parameter is deprecated and replaced by the MuteInterval parameter.
	//
	// example:
	//
	// 5
	InhibitionInterval *int64 `json:"InhibitionInterval,omitempty" xml:"InhibitionInterval,omitempty"`
	// The duration of the alert suppression interval. Unit: minutes.
	//
	// example:
	//
	// 5
	MuteInterval *int64 `json:"MuteInterval,omitempty" xml:"MuteInterval,omitempty"`
	// The alert notification methods.
	NotificationChannels []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels `json:"NotificationChannels,omitempty" xml:"NotificationChannels,omitempty" type:"Repeated"`
	// The settings of alert notification recipients.
	NotificationReceivers []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers `json:"NotificationReceivers,omitempty" xml:"NotificationReceivers,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) GetInhibitionInterval() *int64 {
	return s.InhibitionInterval
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) GetMuteInterval() *int64 {
	return s.MuteInterval
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) GetNotificationChannels() []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels {
	return s.NotificationChannels
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) GetNotificationReceivers() []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	return s.NotificationReceivers
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetInhibitionInterval(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.InhibitionInterval = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetMuteInterval(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.MuteInterval = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetNotificationChannels(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.NotificationChannels = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) SetNotificationReceivers(v []*ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings {
	s.NotificationReceivers = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettings) Validate() error {
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

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels struct {
	// The alert notification methods.
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
	// The severity level. Valid values:
	//
	// 	- Warning
	//
	// 	- Critical
	//
	// example:
	//
	// Critical
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) GetChannels() []*string {
	return s.Channels
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) GetSeverity() *string {
	return s.Severity
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) SetChannels(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels {
	s.Channels = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) SetSeverity(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels {
	s.Severity = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationChannels) Validate() error {
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers struct {
	// The recipient type. Valid values: AliyunUid, DingToken, FeishuToken, and WebHookUrl.
	//
	// example:
	//
	// DingToken
	ReceiverType *string `json:"ReceiverType,omitempty" xml:"ReceiverType,omitempty"`
	// The recipient.
	//
	// 	- If the value of the ReceiverType parameter is AliyunUid, the value of this parameter is the Alibaba Cloud account ID of a user.
	//
	// 	- If the value of the ReceiverType parameter is DingToken, the value of this parameter is the token of a DingTalk chatbot.
	ReceiverValues []*string `json:"ReceiverValues,omitempty" xml:"ReceiverValues,omitempty" type:"Repeated"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) GetReceiverType() *string {
	return s.ReceiverType
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) GetReceiverValues() []*string {
	return s.ReceiverValues
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) SetReceiverType(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	s.ReceiverType = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) SetReceiverValues(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers {
	s.ReceiverValues = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesNotificationSettingsNotificationReceivers) Validate() error {
	return dara.Validate(s)
}

type ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions struct {
	// Deprecated
	//
	// This parameter is deprecated and replaced by the DdlTypes parameter.
	DdlReportTags []*string `json:"DdlReportTags,omitempty" xml:"DdlReportTags,omitempty" type:"Repeated"`
	// The types of DDL operations for which the alert rule takes effect. This parameter is returned only if the MetricType parameter is set to DdlReport.
	DdlTypes []*string `json:"DdlTypes,omitempty" xml:"DdlTypes,omitempty" type:"Repeated"`
	// The time interval for alert calculation. Unit: minutes.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The severity level. Valid values:
	//
	// 	- Warning
	//
	// 	- Critical
	//
	// example:
	//
	// Critical
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The alert threshold.
	//
	// 	- If the alert rule is for task status, no threshold is used.
	//
	// 	- If the alert rule is for failovers, the threshold is the number of failovers.
	//
	// 	- If the alert rule is for latency, the threshold is the latency duration, in seconds.
	//
	// example:
	//
	// 5
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GoString() string {
	return s.String()
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GetDdlReportTags() []*string {
	return s.DdlReportTags
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GetDdlTypes() []*string {
	return s.DdlTypes
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GetDuration() *int64 {
	return s.Duration
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GetSeverity() *string {
	return s.Severity
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) GetThreshold() *int64 {
	return s.Threshold
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetDdlReportTags(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.DdlReportTags = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetDdlTypes(v []*string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.DdlTypes = v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetDuration(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.Duration = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetSeverity(v string) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.Severity = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) SetThreshold(v int64) *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions {
	s.Threshold = &v
	return s
}

func (s *ListDIAlarmRulesResponseBodyPagingInfoDIJobAlarmRulesTriggerConditions) Validate() error {
	return dara.Validate(s)
}
