// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataQualityRuleTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfig(v *DataQualityRuleTemplateCheckingConfig) *DataQualityRuleTemplate
	GetCheckingConfig() *DataQualityRuleTemplateCheckingConfig
	SetCode(v string) *DataQualityRuleTemplate
	GetCode() *string
	SetDirectoryPath(v string) *DataQualityRuleTemplate
	GetDirectoryPath() *string
	SetName(v string) *DataQualityRuleTemplate
	GetName() *string
	SetProjectId(v int64) *DataQualityRuleTemplate
	GetProjectId() *int64
	SetSamplingConfig(v *DataQualityRuleTemplateSamplingConfig) *DataQualityRuleTemplate
	GetSamplingConfig() *DataQualityRuleTemplateSamplingConfig
	SetTenantId(v int64) *DataQualityRuleTemplate
	GetTenantId() *int64
	SetVisibleScope(v string) *DataQualityRuleTemplate
	GetVisibleScope() *string
}

type DataQualityRuleTemplate struct {
	CheckingConfig *DataQualityRuleTemplateCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// /ods/订单数据
	DirectoryPath  *string                                `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	Name           *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	ProjectId      *int64                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *DataQualityRuleTemplateSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	TenantId       *int64                                 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// Project
	VisibleScope *string `json:"VisibleScope,omitempty" xml:"VisibleScope,omitempty"`
}

func (s DataQualityRuleTemplate) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleTemplate) GoString() string {
	return s.String()
}

func (s *DataQualityRuleTemplate) GetCheckingConfig() *DataQualityRuleTemplateCheckingConfig {
	return s.CheckingConfig
}

func (s *DataQualityRuleTemplate) GetCode() *string {
	return s.Code
}

func (s *DataQualityRuleTemplate) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *DataQualityRuleTemplate) GetName() *string {
	return s.Name
}

func (s *DataQualityRuleTemplate) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DataQualityRuleTemplate) GetSamplingConfig() *DataQualityRuleTemplateSamplingConfig {
	return s.SamplingConfig
}

func (s *DataQualityRuleTemplate) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DataQualityRuleTemplate) GetVisibleScope() *string {
	return s.VisibleScope
}

func (s *DataQualityRuleTemplate) SetCheckingConfig(v *DataQualityRuleTemplateCheckingConfig) *DataQualityRuleTemplate {
	s.CheckingConfig = v
	return s
}

func (s *DataQualityRuleTemplate) SetCode(v string) *DataQualityRuleTemplate {
	s.Code = &v
	return s
}

func (s *DataQualityRuleTemplate) SetDirectoryPath(v string) *DataQualityRuleTemplate {
	s.DirectoryPath = &v
	return s
}

func (s *DataQualityRuleTemplate) SetName(v string) *DataQualityRuleTemplate {
	s.Name = &v
	return s
}

func (s *DataQualityRuleTemplate) SetProjectId(v int64) *DataQualityRuleTemplate {
	s.ProjectId = &v
	return s
}

func (s *DataQualityRuleTemplate) SetSamplingConfig(v *DataQualityRuleTemplateSamplingConfig) *DataQualityRuleTemplate {
	s.SamplingConfig = v
	return s
}

func (s *DataQualityRuleTemplate) SetTenantId(v int64) *DataQualityRuleTemplate {
	s.TenantId = &v
	return s
}

func (s *DataQualityRuleTemplate) SetVisibleScope(v string) *DataQualityRuleTemplate {
	s.VisibleScope = &v
	return s
}

func (s *DataQualityRuleTemplate) Validate() error {
	if s.CheckingConfig != nil {
		if err := s.CheckingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SamplingConfig != nil {
		if err := s.SamplingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityRuleTemplateCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityRuleTemplateCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleTemplateCheckingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityRuleTemplateCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *DataQualityRuleTemplateCheckingConfig) GetType() *string {
	return s.Type
}

func (s *DataQualityRuleTemplateCheckingConfig) SetReferencedSamplesFilter(v string) *DataQualityRuleTemplateCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *DataQualityRuleTemplateCheckingConfig) SetType(v string) *DataQualityRuleTemplateCheckingConfig {
	s.Type = &v
	return s
}

func (s *DataQualityRuleTemplateCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleTemplateSamplingConfig struct {
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// { "SQL": "SELECT min(id) from table;" }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s DataQualityRuleTemplateSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleTemplateSamplingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityRuleTemplateSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *DataQualityRuleTemplateSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *DataQualityRuleTemplateSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *DataQualityRuleTemplateSamplingConfig) SetMetric(v string) *DataQualityRuleTemplateSamplingConfig {
	s.Metric = &v
	return s
}

func (s *DataQualityRuleTemplateSamplingConfig) SetMetricParameters(v string) *DataQualityRuleTemplateSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *DataQualityRuleTemplateSamplingConfig) SetSettingConfig(v string) *DataQualityRuleTemplateSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *DataQualityRuleTemplateSamplingConfig) Validate() error {
	return dara.Validate(s)
}
