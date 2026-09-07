// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *UpdateAlertRuleRequest
	GetAlertId() *int64
	SetContactGroupIds(v string) *UpdateAlertRuleRequest
	GetContactGroupIds() *string
	SetIsAutoStart(v bool) *UpdateAlertRuleRequest
	GetIsAutoStart() *bool
	SetRegionId(v string) *UpdateAlertRuleRequest
	GetRegionId() *string
	SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest
	GetTemplageAlertConfig() *string
}

type UpdateAlertRuleRequest struct {
	// The ID of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The IDs of the alert contact groups. The value must be a JSON array.
	//
	// example:
	//
	// [123, 234]
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	// Specifies whether to enable the alert rule after it is created. Default value: `false`.
	//
	// 	- `true`: enables the alert rule.
	//
	// 	- `false`: disables the alert rule.
	//
	// example:
	//
	// true
	IsAutoStart *bool `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configurations of the alert template based on which you want to create an alert rule. The value must be a JSON string. You must set at least one of the **TemplateAlertId*	- and **TemplageAlertConfig*	- parameters. If you set both parameters, the **TemplateAlertId*	- parameter prevails. For more information about the TemplageAlertConfig parameter, see the following **additional information about the TemplageAlertConfig parameter**.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ { "contactGroupIds": "381", "alertType": 5, "alarmContext": { "subTitle": "", "content": "报警名称：$报警名称\\n筛选条件：$筛选\\n报警时间：$报警时间\\n报警内容：$报警内容\\n注意：该报警未收到恢复邮件之前，正在持续报警中，24小时后会再次提醒您！" }, "alertLevel": "WARN", "metricParam": { "appId": "70901", "pid": "atc889zkcf@d8deedfa9bf****", "type": "TXN", "dimensions": [ { "type": "STATIC", "value": "\\\\/hello_test_api_address\\\\/test1", "key": "rpc" } ] }, "alertWay": [ "SMS", "MAIL", "DING_ROBOT" ], "alertRule": { "rules": [ { "measure": "appstat.txn.rt", "alias": "入口调用响应时间_ms", "aggregates": "AVG", "nValue": 1, "value": 1, "operator": "CURRENT_GTE" } ], "operator": "|" }, "title": "报警模板报警名", "config": "{\\"continuous\\":false,\\"dataRevision\\":2,\\"ownerId\\":\\"123412341234\\"}", "notice": { "noticeStartTime": 1480521600000, "startTime": 1480521600000, "endTime": 1480607940000, "noticeEndTime": 1480607940000 }, "status": "NON"  } ]
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
}

func (s UpdateAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *UpdateAlertRuleRequest) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *UpdateAlertRuleRequest) GetIsAutoStart() *bool {
	return s.IsAutoStart
}

func (s *UpdateAlertRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAlertRuleRequest) GetTemplageAlertConfig() *string {
	return s.TemplageAlertConfig
}

func (s *UpdateAlertRuleRequest) SetAlertId(v int64) *UpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetContactGroupIds(v string) *UpdateAlertRuleRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetIsAutoStart(v bool) *UpdateAlertRuleRequest {
	s.IsAutoStart = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetRegionId(v string) *UpdateAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *UpdateAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}
