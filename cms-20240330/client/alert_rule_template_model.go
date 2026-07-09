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
	AlertType     *string `json:"alertType,omitempty" xml:"alertType,omitempty"`
	ApplyCount    *int64  `json:"applyCount,omitempty" xml:"applyCount,omitempty"`
	BizType       *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	Datasource    *string `json:"datasource,omitempty" xml:"datasource,omitempty"`
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate     *int64  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *int64  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsSystem      *int32  `json:"isSystem,omitempty" xml:"isSystem,omitempty"`
	Labels        *string `json:"labels,omitempty" xml:"labels,omitempty"`
	RuleConfigs   *string `json:"ruleConfigs,omitempty" xml:"ruleConfigs,omitempty"`
	Scenes        *string `json:"scenes,omitempty" xml:"scenes,omitempty"`
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	SourceType    *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Status        *int32  `json:"status,omitempty" xml:"status,omitempty"`
	SubType       *string `json:"subType,omitempty" xml:"subType,omitempty"`
	TemplateName  *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	UserId        *string `json:"userId,omitempty" xml:"userId,omitempty"`
	Uuid          *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
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
