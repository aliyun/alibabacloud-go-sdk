// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomEventRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v string) *PutCustomEventRuleRequest
	GetContactGroups() *string
	SetEffectiveInterval(v string) *PutCustomEventRuleRequest
	GetEffectiveInterval() *string
	SetEmailSubject(v string) *PutCustomEventRuleRequest
	GetEmailSubject() *string
	SetEventName(v string) *PutCustomEventRuleRequest
	GetEventName() *string
	SetGroupId(v string) *PutCustomEventRuleRequest
	GetGroupId() *string
	SetLevel(v string) *PutCustomEventRuleRequest
	GetLevel() *string
	SetPeriod(v string) *PutCustomEventRuleRequest
	GetPeriod() *string
	SetRuleId(v string) *PutCustomEventRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *PutCustomEventRuleRequest
	GetRuleName() *string
	SetThreshold(v string) *PutCustomEventRuleRequest
	GetThreshold() *string
	SetWebhook(v string) *PutCustomEventRuleRequest
	GetWebhook() *string
}

type PutCustomEventRuleRequest struct {
	// The alert contact group that receives alert notifications. Separate multiple contact groups with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The time period during which the alert rule is effective. Valid values: 00:00 to 23:59.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert notification email.
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The name of the custom event. For more information about how to obtain the event name, see [DescribeCustomEventAttribute](https://help.aliyun.com/document_detail/115262.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// HostDown
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The ID of the application group. For more information about how to obtain the group ID, see [DescribeCustomEventAttribute](https://help.aliyun.com/document_detail/115262.html).
	//
	// >  The value 0 indicates that the reported custom event does not belong to any application Group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7378****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The level of the alert. Valid values:
	//
	// 	- CRITICAL: critical issue
	//
	// 	- WARN: warning
	//
	// 	- INFO: information
	//
	// This parameter is required.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The cycle that is used to aggregate monitoring data of the custom event. Unit: seconds. Set the value to an integral multiple of 60. Default value: 300.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the alert rule.
	//
	// >  You can specify an existing ID to modify the corresponding alert rule or specify a new ID to create an alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// CustomRuleId1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// CustomeRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The alert threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutCustomEventRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCustomEventRuleRequest) GoString() string {
	return s.String()
}

func (s *PutCustomEventRuleRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutCustomEventRuleRequest) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutCustomEventRuleRequest) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutCustomEventRuleRequest) GetEventName() *string {
	return s.EventName
}

func (s *PutCustomEventRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PutCustomEventRuleRequest) GetLevel() *string {
	return s.Level
}

func (s *PutCustomEventRuleRequest) GetPeriod() *string {
	return s.Period
}

func (s *PutCustomEventRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutCustomEventRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutCustomEventRuleRequest) GetThreshold() *string {
	return s.Threshold
}

func (s *PutCustomEventRuleRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *PutCustomEventRuleRequest) SetContactGroups(v string) *PutCustomEventRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetEffectiveInterval(v string) *PutCustomEventRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetEmailSubject(v string) *PutCustomEventRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetEventName(v string) *PutCustomEventRuleRequest {
	s.EventName = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetGroupId(v string) *PutCustomEventRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetLevel(v string) *PutCustomEventRuleRequest {
	s.Level = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetPeriod(v string) *PutCustomEventRuleRequest {
	s.Period = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetRuleId(v string) *PutCustomEventRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetRuleName(v string) *PutCustomEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetThreshold(v string) *PutCustomEventRuleRequest {
	s.Threshold = &v
	return s
}

func (s *PutCustomEventRuleRequest) SetWebhook(v string) *PutCustomEventRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutCustomEventRuleRequest) Validate() error {
	return dara.Validate(s)
}
