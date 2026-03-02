// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbbrevMonitorAlertRule interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *AbbrevMonitorAlertRule
	GetCreateTime() *string
	SetId(v int64) *AbbrevMonitorAlertRule
	GetId() *int64
	SetName(v string) *AbbrevMonitorAlertRule
	GetName() *string
	SetTriggerContent(v string) *AbbrevMonitorAlertRule
	GetTriggerContent() *string
	SetTriggerRule(v string) *AbbrevMonitorAlertRule
	GetTriggerRule() *string
}

type AbbrevMonitorAlertRule struct {
	// example:
	//
	// 2022-06-11T00:00:00.000Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 规则1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 最近5分钟节点机CPU使用率平均大于等于90%
	TriggerContent *string `json:"triggerContent,omitempty" xml:"triggerContent,omitempty"`
	// example:
	//
	// 10
	TriggerRule *string `json:"triggerRule,omitempty" xml:"triggerRule,omitempty"`
}

func (s AbbrevMonitorAlertRule) String() string {
	return dara.Prettify(s)
}

func (s AbbrevMonitorAlertRule) GoString() string {
	return s.String()
}

func (s *AbbrevMonitorAlertRule) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AbbrevMonitorAlertRule) GetId() *int64 {
	return s.Id
}

func (s *AbbrevMonitorAlertRule) GetName() *string {
	return s.Name
}

func (s *AbbrevMonitorAlertRule) GetTriggerContent() *string {
	return s.TriggerContent
}

func (s *AbbrevMonitorAlertRule) GetTriggerRule() *string {
	return s.TriggerRule
}

func (s *AbbrevMonitorAlertRule) SetCreateTime(v string) *AbbrevMonitorAlertRule {
	s.CreateTime = &v
	return s
}

func (s *AbbrevMonitorAlertRule) SetId(v int64) *AbbrevMonitorAlertRule {
	s.Id = &v
	return s
}

func (s *AbbrevMonitorAlertRule) SetName(v string) *AbbrevMonitorAlertRule {
	s.Name = &v
	return s
}

func (s *AbbrevMonitorAlertRule) SetTriggerContent(v string) *AbbrevMonitorAlertRule {
	s.TriggerContent = &v
	return s
}

func (s *AbbrevMonitorAlertRule) SetTriggerRule(v string) *AbbrevMonitorAlertRule {
	s.TriggerRule = &v
	return s
}

func (s *AbbrevMonitorAlertRule) Validate() error {
	return dara.Validate(s)
}
