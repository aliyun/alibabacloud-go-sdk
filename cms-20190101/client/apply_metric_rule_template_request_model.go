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
	// The template application policy. Valid values:
	//
	// 	- all (default): deletes all the rules that are created by using the alert template from the selected application group, and then creates alert rules based on the template.
	//
	// 	- append: deletes the rules that are created by using the alert template from the selected application group, and then creates alert rules based on the existing template.
	//
	// example:
	//
	// all
	AppendMode *string `json:"AppendMode,omitempty" xml:"AppendMode,omitempty"`
	// The mode in which the alert template is applied. Valid values:
	//
	// 	- GROUP_INSTANCE_FIRST: The metrics in the application group take precedence. If a metric specified in the alert template does not exist in the application group, the system does not generate an alert rule for the metric based on the alert template.
	//
	// 	- ALARM_TEMPLATE_FIRST: The metrics specified in the alert template take precedence. If a metric specified in the alert template does not exist in the application group, the system still generates an alert rule for the metric based on the alert template.
	//
	// example:
	//
	// GROUP_INSTANCE_FIRST
	ApplyMode *string `json:"ApplyMode,omitempty" xml:"ApplyMode,omitempty"`
	// The end of the time period during which the alert rule is effective. Valid values: 00 to 23. A value of 00 indicates 00:59 and a value of 23 indicates 23:59.
	//
	// example:
	//
	// 23
	EnableEndTime *int64 `json:"EnableEndTime,omitempty" xml:"EnableEndTime,omitempty"`
	// The beginning of the time period during which the alert rule is effective. Valid values: 00 to 23. A value of 00 indicates 00:00 and a value of 23 indicates 23:00.
	//
	// example:
	//
	// 00
	EnableStartTime *int64 `json:"EnableStartTime,omitempty" xml:"EnableStartTime,omitempty"`
	// The ID of the application group to which the alert template is applied.
	//
	// For more information about how to query the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The alert notification method. Valid values:
	//
	// Set the value to 4. A value of 4 indicates that alert notifications are sent by using TradeManager and DingTalk chatbots.
	//
	// example:
	//
	// 4
	NotifyLevel *int64 `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	// The mute period during which notifications are not repeatedly sent for an alert. Unit: seconds. Default value: 86400.
	//
	// >  Only one alert notification is sent during each mute period even if the metric value exceeds the alert threshold several times.
	//
	// example:
	//
	// 86400
	SilenceTime *int64 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The ID of the alert template.
	//
	// For more information about how to query the IDs of alert templates, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 700****
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
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
