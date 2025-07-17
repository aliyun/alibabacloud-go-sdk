// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAppAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v string) *ImportAppAlertRulesRequest
	GetContactGroupIds() *string
	SetIsAutoStart(v bool) *ImportAppAlertRulesRequest
	GetIsAutoStart() *bool
	SetPids(v string) *ImportAppAlertRulesRequest
	GetPids() *string
	SetRegionId(v string) *ImportAppAlertRulesRequest
	GetRegionId() *string
	SetTags(v []*ImportAppAlertRulesRequestTags) *ImportAppAlertRulesRequest
	GetTags() []*ImportAppAlertRulesRequestTags
	SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest
	GetTemplageAlertConfig() *string
	SetTemplateAlertId(v string) *ImportAppAlertRulesRequest
	GetTemplateAlertId() *string
}

type ImportAppAlertRulesRequest struct {
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
	// The process identifiers (PIDs) of the applications associated with the alert rule. The value must be a JSON array. For more information about how to obtain the PID, see [Obtain the PID of an application](~~186100#section-bkl-3j6-ezg~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["atc889zkcf@d8deedfa9bfxxxx", "acd129bfcf@d5daebfa6cdxxxx"]
	Pids *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	// The ID of the region where the associated applications reside.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags.
	Tags []*ImportAppAlertRulesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The configurations of the alert template based on which you want to create an alert rule. The value must be a JSON string. You must set at least one of the **TemplateAlertId*	- and **TemplageAlertConfig*	- parameters. If you set both parameters, the **TemplateAlertId*	- parameter prevails. For more information about the TemplageAlertConfig parameter, see the following **additional information about the TemplageAlertConfig parameter**.
	//
	// example:
	//
	// [ { "contactGroupIds": "381", "alertType": 5, "alarmContext": { "subTitle": "", "content": "Alarm name: $alarm name\\nFilter condition: $filter\\nAlarm time : $Alarm time\\nAlarm content: $Alarm content\\nNote: Before the recovery email is received, the alarm is in continuous alarm, and you will be reminded again after 24 hours!" }, "alertLevel": "WARN", " metricParam": { "appId": "70901", "pid": "atc889zkcf@d8deedfa9bf****", "type": "TXN", "dimensions": [ { "type": "STATIC", "value ": "\\\\/hello_test_api_address\\\\/test1", "key": "rpc" } ] }, "alertWay": [ "SMS", "MAIL", "DING_ROBOT" ], "alertRule": { "rules" : [ { "measure": "appstat.txn.rt", "alias": "Entry call response time_ms", "aggregates": "AVG", "nValue": 1, "value": 1, "operator ": "CURRENT_GTE" } ], "operator": "|" }, "title": "Alarm template alarm name", "config": "{\\"continuous\\":false,\\"dataRevision\\":2, \\"ownerId\\":\\"123412341234\\"}", "notice": { "noticeStartTime": 1480521600000, "startTime": 1480521600000, "endTime": 1480607940000, "noticeEndTime": 1480607940000 }, "stat us": "NON " } ]
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	// The ID of the alert template. You must set at least one of the **TemplateAlertId*	- and **TemplageAlertConfig*	- parameters. If you set both parameters, the **TemplateAlertId*	- parameter prevails.
	//
	// example:
	//
	// 324324234
	TemplateAlertId *string `json:"TemplateAlertId,omitempty" xml:"TemplateAlertId,omitempty"`
}

func (s ImportAppAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportAppAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequest) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *ImportAppAlertRulesRequest) GetIsAutoStart() *bool {
	return s.IsAutoStart
}

func (s *ImportAppAlertRulesRequest) GetPids() *string {
	return s.Pids
}

func (s *ImportAppAlertRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportAppAlertRulesRequest) GetTags() []*ImportAppAlertRulesRequestTags {
	return s.Tags
}

func (s *ImportAppAlertRulesRequest) GetTemplageAlertConfig() *string {
	return s.TemplageAlertConfig
}

func (s *ImportAppAlertRulesRequest) GetTemplateAlertId() *string {
	return s.TemplateAlertId
}

func (s *ImportAppAlertRulesRequest) SetContactGroupIds(v string) *ImportAppAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetIsAutoStart(v bool) *ImportAppAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetPids(v string) *ImportAppAlertRulesRequest {
	s.Pids = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetRegionId(v string) *ImportAppAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTags(v []*ImportAppAlertRulesRequestTags) *ImportAppAlertRulesRequest {
	s.Tags = v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplateAlertId(v string) *ImportAppAlertRulesRequest {
	s.TemplateAlertId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}

type ImportAppAlertRulesRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImportAppAlertRulesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ImportAppAlertRulesRequestTags) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ImportAppAlertRulesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ImportAppAlertRulesRequestTags) SetKey(v string) *ImportAppAlertRulesRequestTags {
	s.Key = &v
	return s
}

func (s *ImportAppAlertRulesRequestTags) SetValue(v string) *ImportAppAlertRulesRequestTags {
	s.Value = &v
	return s
}

func (s *ImportAppAlertRulesRequestTags) Validate() error {
	return dara.Validate(s)
}
