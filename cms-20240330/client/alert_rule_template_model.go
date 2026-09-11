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
	SetBizType(v string) *AlertRuleTemplate
	GetBizType() *string
	SetDatasource(v string) *AlertRuleTemplate
	GetDatasource() *string
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
	SetNamespace(v string) *AlertRuleTemplate
	GetNamespace() *string
	SetProductCategory(v string) *AlertRuleTemplate
	GetProductCategory() *string
	SetRuleConfigs(v string) *AlertRuleTemplate
	GetRuleConfigs() *string
	SetScenes(v string) *AlertRuleTemplate
	GetScenes() *string
	SetSchemaVersion(v string) *AlertRuleTemplate
	GetSchemaVersion() *string
	SetSourceType(v string) *AlertRuleTemplate
	GetSourceType() *string
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
	// The alert type.
	//
	// example:
	//
	// METRIC_SET
	AlertType *string `json:"alertType,omitempty" xml:"alertType,omitempty"`
	// The number of rules that have been applied from this template.
	//
	// example:
	//
	// 5
	ApplyCount *int64 `json:"applyCount,omitempty" xml:"applyCount,omitempty"`
	// The business type.
	//
	// example:
	//
	// ALERT
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// The data source, in JSON string format.
	//
	// example:
	//
	// {"type":"SLS","project":"my-project"}
	Datasource *string `json:"datasource,omitempty" xml:"datasource,omitempty"`
	// The template description.
	//
	// example:
	//
	// Triggers an alert when the CPU usage of an ECS instance exceeds the threshold
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The creation time, in UNIX millisecond timestamp format.
	//
	// example:
	//
	// 1700000000000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modified time, in UNIX millisecond timestamp format.
	//
	// example:
	//
	// 1700000000000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The template ID.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the template is a system template. Valid values: 1: yes. 0: no.
	//
	// example:
	//
	// 1
	IsSystem *int32 `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
	// The labels, in JSON string format.
	//
	// example:
	//
	// {"env":"prod","team":"ops"}
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The namespace.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The cloud service category.
	//
	// example:
	//
	// ecs
	ProductCategory *string `json:"productCategory,omitempty" xml:"productCategory,omitempty"`
	// The rule configurations, in JSON string format.
	//
	// example:
	//
	// {"threshold":80,"duration":60}
	RuleConfigs *string `json:"ruleConfigs,omitempty" xml:"ruleConfigs,omitempty"`
	// The applicable scenarios.
	//
	// example:
	//
	// ECS
	Scenes *string `json:"scenes,omitempty" xml:"scenes,omitempty"`
	// The schema version.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The source type.
	//
	// example:
	//
	// SYSTEM
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The template status. Valid values: 1: enabled. 0: disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The subtype.
	//
	// example:
	//
	// THRESHOLD
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// The template name.
	//
	// example:
	//
	// ECS CPU Usage Alert Template
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The ID of the user to whom the template belongs.
	//
	// example:
	//
	// 1234567890
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The template UUID.
	//
	// example:
	//
	// a1b2c3d4-e5f6-7890-abcd-ef1234567890
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

func (s *AlertRuleTemplate) GetBizType() *string {
	return s.BizType
}

func (s *AlertRuleTemplate) GetDatasource() *string {
	return s.Datasource
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

func (s *AlertRuleTemplate) GetNamespace() *string {
	return s.Namespace
}

func (s *AlertRuleTemplate) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *AlertRuleTemplate) GetRuleConfigs() *string {
	return s.RuleConfigs
}

func (s *AlertRuleTemplate) GetScenes() *string {
	return s.Scenes
}

func (s *AlertRuleTemplate) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *AlertRuleTemplate) GetSourceType() *string {
	return s.SourceType
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

func (s *AlertRuleTemplate) SetBizType(v string) *AlertRuleTemplate {
	s.BizType = &v
	return s
}

func (s *AlertRuleTemplate) SetDatasource(v string) *AlertRuleTemplate {
	s.Datasource = &v
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

func (s *AlertRuleTemplate) SetNamespace(v string) *AlertRuleTemplate {
	s.Namespace = &v
	return s
}

func (s *AlertRuleTemplate) SetProductCategory(v string) *AlertRuleTemplate {
	s.ProductCategory = &v
	return s
}

func (s *AlertRuleTemplate) SetRuleConfigs(v string) *AlertRuleTemplate {
	s.RuleConfigs = &v
	return s
}

func (s *AlertRuleTemplate) SetScenes(v string) *AlertRuleTemplate {
	s.Scenes = &v
	return s
}

func (s *AlertRuleTemplate) SetSchemaVersion(v string) *AlertRuleTemplate {
	s.SchemaVersion = &v
	return s
}

func (s *AlertRuleTemplate) SetSourceType(v string) *AlertRuleTemplate {
	s.SourceType = &v
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
