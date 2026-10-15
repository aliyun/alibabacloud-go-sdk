// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleTemplateApplyRecord interface {
	dara.Model
	String() string
	GoString() string
	SetApplyConfig(v string) *AlertRuleTemplateApplyRecord
	GetApplyConfig() *string
	SetGmtCreate(v int64) *AlertRuleTemplateApplyRecord
	GetGmtCreate() *int64
	SetGmtModified(v int64) *AlertRuleTemplateApplyRecord
	GetGmtModified() *int64
	SetId(v int64) *AlertRuleTemplateApplyRecord
	GetId() *int64
	SetRegionId(v string) *AlertRuleTemplateApplyRecord
	GetRegionId() *string
	SetRuleId(v string) *AlertRuleTemplateApplyRecord
	GetRuleId() *string
	SetStatus(v string) *AlertRuleTemplateApplyRecord
	GetStatus() *string
	SetTemplateUuid(v string) *AlertRuleTemplateApplyRecord
	GetTemplateUuid() *string
	SetUserId(v string) *AlertRuleTemplateApplyRecord
	GetUserId() *string
}

type AlertRuleTemplateApplyRecord struct {
	// The effective configuration for the rule application.
	//
	// example:
	//
	// {"workspace":"ws-1","displayName":"CPU alert"}
	ApplyConfig *string `json:"applyConfig,omitempty" xml:"applyConfig,omitempty"`
	// The time when the record was created.
	//
	// example:
	//
	// 1680861352600
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the record was updated.
	//
	// example:
	//
	// 1680861352600
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The record ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 12345
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The status.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The UUID of the alert rule template.
	//
	// example:
	//
	// template-prometheus-ack-node-cpu-usage-exceeds-85
	TemplateUuid *string `json:"templateUuid,omitempty" xml:"templateUuid,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1654218965343050
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AlertRuleTemplateApplyRecord) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleTemplateApplyRecord) GoString() string {
	return s.String()
}

func (s *AlertRuleTemplateApplyRecord) GetApplyConfig() *string {
	return s.ApplyConfig
}

func (s *AlertRuleTemplateApplyRecord) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *AlertRuleTemplateApplyRecord) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *AlertRuleTemplateApplyRecord) GetId() *int64 {
	return s.Id
}

func (s *AlertRuleTemplateApplyRecord) GetRegionId() *string {
	return s.RegionId
}

func (s *AlertRuleTemplateApplyRecord) GetRuleId() *string {
	return s.RuleId
}

func (s *AlertRuleTemplateApplyRecord) GetStatus() *string {
	return s.Status
}

func (s *AlertRuleTemplateApplyRecord) GetTemplateUuid() *string {
	return s.TemplateUuid
}

func (s *AlertRuleTemplateApplyRecord) GetUserId() *string {
	return s.UserId
}

func (s *AlertRuleTemplateApplyRecord) SetApplyConfig(v string) *AlertRuleTemplateApplyRecord {
	s.ApplyConfig = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetGmtCreate(v int64) *AlertRuleTemplateApplyRecord {
	s.GmtCreate = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetGmtModified(v int64) *AlertRuleTemplateApplyRecord {
	s.GmtModified = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetId(v int64) *AlertRuleTemplateApplyRecord {
	s.Id = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetRegionId(v string) *AlertRuleTemplateApplyRecord {
	s.RegionId = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetRuleId(v string) *AlertRuleTemplateApplyRecord {
	s.RuleId = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetStatus(v string) *AlertRuleTemplateApplyRecord {
	s.Status = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetTemplateUuid(v string) *AlertRuleTemplateApplyRecord {
	s.TemplateUuid = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) SetUserId(v string) *AlertRuleTemplateApplyRecord {
	s.UserId = &v
	return s
}

func (s *AlertRuleTemplateApplyRecord) Validate() error {
	return dara.Validate(s)
}
