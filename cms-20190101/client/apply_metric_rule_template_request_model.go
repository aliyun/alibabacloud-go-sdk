// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyMetricRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppendMode(v string) *ApplyMetricRuleTemplateRequest
	GetAppendMode() *string
	SetApplyMode(v string) *ApplyMetricRuleTemplateRequest
	GetApplyMode() *string
	SetDryRun(v bool) *ApplyMetricRuleTemplateRequest
	GetDryRun() *bool
	SetEnableEndTime(v int64) *ApplyMetricRuleTemplateRequest
	GetEnableEndTime() *int64
	SetEnableStartTime(v int64) *ApplyMetricRuleTemplateRequest
	GetEnableStartTime() *int64
	SetGroupId(v int64) *ApplyMetricRuleTemplateRequest
	GetGroupId() *int64
	SetNotifyLevel(v int64) *ApplyMetricRuleTemplateRequest
	GetNotifyLevel() *int64
	SetSilenceTime(v int64) *ApplyMetricRuleTemplateRequest
	GetSilenceTime() *int64
	SetTemplateIds(v string) *ApplyMetricRuleTemplateRequest
	GetTemplateIds() *string
	SetWebhook(v string) *ApplyMetricRuleTemplateRequest
	GetWebhook() *string
}

type ApplyMetricRuleTemplateRequest struct {
	// The application policy of the template. Valid values:
	//
	// - all: deletes all rules created by using alert templates on the selected group and then creates new alert rules based on the template. This is the default value.
	//
	// - append: deletes the rules created by using the current alert template on the selected group and then creates new alert rules based on the current template.
	//
	// example:
	//
	// all
	AppendMode *string `json:"AppendMode,omitempty" xml:"AppendMode,omitempty"`
	// The application mode of the template. Valid values:
	//
	// - GROUP_INSTANCE_FIRST: application group instances take precedence. When the alert template is applied, the instances in the application group take precedence. If an instance specified in the template does not exist in the application group, the rule for the instance in the template is ignored.
	//
	// - ALARM_TEMPLATE_FIRST: alert template instances take precedence. When the alert template is applied, alert rules are created regardless of whether the instances exist in the application group.
	//
	// example:
	//
	// GROUP_INSTANCE_FIRST
	ApplyMode *string `json:"ApplyMode,omitempty" xml:"ApplyMode,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run and does not bind a security group to the instance. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the dry run, an error is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false**: performs a normal request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The end time of the effective period for the alert, in hours. The value is a two-digit string, such as "00", "18", or "23". Valid values: 00 to 23. The value specifies the hour of the end time, and the actual effective period ends at the 59th minute of the hour. For example, "23" indicates that the effective period ends at 23:59. The default time zone is UTC+8.
	//
	// example:
	//
	// 23
	EnableEndTime *int64 `json:"EnableEndTime,omitempty" xml:"EnableEndTime,omitempty"`
	// The start time of the effective period for the alert, in hours. The value is a two-digit string, such as "00", "09", or "23". Valid values: 00 to 23, which represent 00:00 to 23:00. The default time zone is UTC+8.
	//
	// example:
	//
	// 00
	EnableStartTime *int64 `json:"EnableStartTime,omitempty" xml:"EnableStartTime,omitempty"`
	// The ID of the application group.
	//
	// For more information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The alert notification method. Valid values:
	//
	// <props="china">- 2: phone call, text message, mailbox, TradeManager, and DingTalk chatbot.
	//
	// <props="china">- 3: text message, mailbox, TradeManager, and DingTalk chatbot.
	//
	// <props="china">- 4: TradeManager and DingTalk chatbot.
	//
	//
	// <props="intl">4: TradeManager and DingTalk chatbot.
	//
	// <props="partner">4: TradeManager and DingTalk chatbot.
	//
	// example:
	//
	// 4
	NotifyLevel *int64 `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	// The mute for. Unit: seconds. Default value: 86400.
	//
	// > If the monitoring data continuously exceeds the threshold of the alert rule, the system sends only one alert notification within each epoch.
	//
	// example:
	//
	// 86400
	SilenceTime *int64 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The ID of the alert template.
	//
	// For more information about how to obtain the ID of an alert template, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 700****
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
	// The webhook URL to which a POST request is sent when an alert is triggered.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ApplyMetricRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApplyMetricRuleTemplateRequest) GetAppendMode() *string {
	return s.AppendMode
}

func (s *ApplyMetricRuleTemplateRequest) GetApplyMode() *string {
	return s.ApplyMode
}

func (s *ApplyMetricRuleTemplateRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ApplyMetricRuleTemplateRequest) GetEnableEndTime() *int64 {
	return s.EnableEndTime
}

func (s *ApplyMetricRuleTemplateRequest) GetEnableStartTime() *int64 {
	return s.EnableStartTime
}

func (s *ApplyMetricRuleTemplateRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ApplyMetricRuleTemplateRequest) GetNotifyLevel() *int64 {
	return s.NotifyLevel
}

func (s *ApplyMetricRuleTemplateRequest) GetSilenceTime() *int64 {
	return s.SilenceTime
}

func (s *ApplyMetricRuleTemplateRequest) GetTemplateIds() *string {
	return s.TemplateIds
}

func (s *ApplyMetricRuleTemplateRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *ApplyMetricRuleTemplateRequest) SetAppendMode(v string) *ApplyMetricRuleTemplateRequest {
	s.AppendMode = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetApplyMode(v string) *ApplyMetricRuleTemplateRequest {
	s.ApplyMode = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetDryRun(v bool) *ApplyMetricRuleTemplateRequest {
	s.DryRun = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetEnableEndTime(v int64) *ApplyMetricRuleTemplateRequest {
	s.EnableEndTime = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetEnableStartTime(v int64) *ApplyMetricRuleTemplateRequest {
	s.EnableStartTime = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetGroupId(v int64) *ApplyMetricRuleTemplateRequest {
	s.GroupId = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetNotifyLevel(v int64) *ApplyMetricRuleTemplateRequest {
	s.NotifyLevel = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetSilenceTime(v int64) *ApplyMetricRuleTemplateRequest {
	s.SilenceTime = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetTemplateIds(v string) *ApplyMetricRuleTemplateRequest {
	s.TemplateIds = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) SetWebhook(v string) *ApplyMetricRuleTemplateRequest {
	s.Webhook = &v
	return s
}

func (s *ApplyMetricRuleTemplateRequest) Validate() error {
	return dara.Validate(s)
}
