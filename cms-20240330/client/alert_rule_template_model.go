// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAlertType(v string) *AlertRuleTemplate
	GetAlertType() *string
	SetApplyCount(v int64) *AlertRuleTemplate
	GetApplyCount() *int64
	SetDescription(v string) *AlertRuleTemplate
	GetDescription() *string
	SetGmtCreate(v int64) *AlertRuleTemplate
	GetGmtCreate() *int64
	SetGmtModified(v int64) *AlertRuleTemplate
	GetGmtModified() *int64
	SetId(v int64) *AlertRuleTemplate
	GetId() *int64
	SetIsSystem(v int32) *AlertRuleTemplate
	GetIsSystem() *int32
	SetLabels(v string) *AlertRuleTemplate
	GetLabels() *string
	SetRuleConfigs(v string) *AlertRuleTemplate
	GetRuleConfigs() *string
	SetStatus(v int32) *AlertRuleTemplate
	GetStatus() *int32
	SetSubType(v string) *AlertRuleTemplate
	GetSubType() *string
	SetTemplateName(v string) *AlertRuleTemplate
	GetTemplateName() *string
	SetUserId(v string) *AlertRuleTemplate
	GetUserId() *string
	SetUuid(v string) *AlertRuleTemplate
	GetUuid() *string
}

type AlertRuleTemplate struct {
	// The type of the alert.
	AlertType *string `json:"alertType,omitempty" xml:"alertType,omitempty"`
	// The number of alert rules created from this template.
	ApplyCount *int64 `json:"applyCount,omitempty" xml:"applyCount,omitempty"`
	// The description of the template.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The creation time of the template, as a UNIX timestamp.
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time the template was last modified, as a UNIX timestamp.
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ID of the alert rule template.
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the template is system-defined. Valid values: `0` (user-defined) and `1` (system-defined).
	IsSystem *int32 `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
	// The labels associated with the template, formatted as a JSON string.
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The rule configuration, formatted as a JSON string.
	RuleConfigs *string `json:"ruleConfigs,omitempty" xml:"ruleConfigs,omitempty"`
	// The status of the template.
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The subtype of the alert.
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// The name of the alert rule template.
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The ID of the user who owns the template.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The universally unique identifier (UUID) of the template.
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s AlertRuleTemplate) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleTemplate) GoString() string {
	return s.String()
}

func (s *AlertRuleTemplate) GetAlertType() *string {
	return s.AlertType
}

func (s *AlertRuleTemplate) GetApplyCount() *int64 {
	return s.ApplyCount
}

func (s *AlertRuleTemplate) GetDescription() *string {
	return s.Description
}

func (s *AlertRuleTemplate) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *AlertRuleTemplate) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *AlertRuleTemplate) GetId() *int64 {
	return s.Id
}

func (s *AlertRuleTemplate) GetIsSystem() *int32 {
	return s.IsSystem
}

func (s *AlertRuleTemplate) GetLabels() *string {
	return s.Labels
}

func (s *AlertRuleTemplate) GetRuleConfigs() *string {
	return s.RuleConfigs
}

func (s *AlertRuleTemplate) GetStatus() *int32 {
	return s.Status
}

func (s *AlertRuleTemplate) GetSubType() *string {
	return s.SubType
}

func (s *AlertRuleTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AlertRuleTemplate) GetUserId() *string {
	return s.UserId
}

func (s *AlertRuleTemplate) GetUuid() *string {
	return s.Uuid
}

func (s *AlertRuleTemplate) SetAlertType(v string) *AlertRuleTemplate {
	s.AlertType = &v
	return s
}

func (s *AlertRuleTemplate) SetApplyCount(v int64) *AlertRuleTemplate {
	s.ApplyCount = &v
	return s
}

func (s *AlertRuleTemplate) SetDescription(v string) *AlertRuleTemplate {
	s.Description = &v
	return s
}

func (s *AlertRuleTemplate) SetGmtCreate(v int64) *AlertRuleTemplate {
	s.GmtCreate = &v
	return s
}

func (s *AlertRuleTemplate) SetGmtModified(v int64) *AlertRuleTemplate {
	s.GmtModified = &v
	return s
}

func (s *AlertRuleTemplate) SetId(v int64) *AlertRuleTemplate {
	s.Id = &v
	return s
}

func (s *AlertRuleTemplate) SetIsSystem(v int32) *AlertRuleTemplate {
	s.IsSystem = &v
	return s
}

func (s *AlertRuleTemplate) SetLabels(v string) *AlertRuleTemplate {
	s.Labels = &v
	return s
}

func (s *AlertRuleTemplate) SetRuleConfigs(v string) *AlertRuleTemplate {
	s.RuleConfigs = &v
	return s
}

func (s *AlertRuleTemplate) SetStatus(v int32) *AlertRuleTemplate {
	s.Status = &v
	return s
}

func (s *AlertRuleTemplate) SetSubType(v string) *AlertRuleTemplate {
	s.SubType = &v
	return s
}

func (s *AlertRuleTemplate) SetTemplateName(v string) *AlertRuleTemplate {
	s.TemplateName = &v
	return s
}

func (s *AlertRuleTemplate) SetUserId(v string) *AlertRuleTemplate {
	s.UserId = &v
	return s
}

func (s *AlertRuleTemplate) SetUuid(v string) *AlertRuleTemplate {
	s.Uuid = &v
	return s
}

func (s *AlertRuleTemplate) Validate() error {
	return dara.Validate(s)
}
