// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityRuleTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityRuleTemplate(v *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) *GetDataQualityRuleTemplateResponseBody
	GetDataQualityRuleTemplate() *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate
	SetRequestId(v string) *GetDataQualityRuleTemplateResponseBody
	GetRequestId() *string
}

type GetDataQualityRuleTemplateResponseBody struct {
	// The details of the rule template.
	DataQualityRuleTemplate *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate `json:"DataQualityRuleTemplate,omitempty" xml:"DataQualityRuleTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityRuleTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleTemplateResponseBody) GetDataQualityRuleTemplate() *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	return s.DataQualityRuleTemplate
}

func (s *GetDataQualityRuleTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityRuleTemplateResponseBody) SetDataQualityRuleTemplate(v *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) *GetDataQualityRuleTemplateResponseBody {
	s.DataQualityRuleTemplate = v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBody) SetRequestId(v string) *GetDataQualityRuleTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBody) Validate() error {
	if s.DataQualityRuleTemplate != nil {
		if err := s.DataQualityRuleTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate struct {
	// The sample verification settings.
	CheckingConfig *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The code of the rule template.
	//
	// example:
	//
	// USER_DEFINED:123
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The category directory in which the custom template is stored. Levels are separated by forward slashes (/). Each level name can be a maximum of 1,024 characters in length and cannot contain whitespace characters or forward slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the rule template. The name can be a combination of digits, letters, Chinese characters, and half-width or full-width punctuation marks. The name can be a maximum of 512 characters in length.
	//
	// example:
	//
	// Table row Count Verification
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 4020
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfig *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The scope in which the template is available:
	//
	// - Tenant: available to the entire tenant.
	//
	// - Project: available only in the current project.
	//
	// example:
	//
	// Project
	VisibleScope *string `json:"VisibleScope,omitempty" xml:"VisibleScope,omitempty"`
}

func (s GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetCheckingConfig() *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig {
	return s.CheckingConfig
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetCode() *string {
	return s.Code
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetName() *string {
	return s.Name
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetSamplingConfig() *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig {
	return s.SamplingConfig
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) GetVisibleScope() *string {
	return s.VisibleScope
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetCheckingConfig(v *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.CheckingConfig = v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetCode(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.Code = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetDirectoryPath(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.DirectoryPath = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetName(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.Name = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetProjectId(v int64) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetSamplingConfig(v *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.SamplingConfig = v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) SetVisibleScope(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate {
	s.VisibleScope = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate) Validate() error {
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

type GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig struct {
	// Some types of thresholds require you to query reference samples and aggregate the values of the reference samples to obtain the threshold for comparison. An expression is used to indicate the query method of reference samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold calculation method:
	//
	// - Fixed
	//
	// - Fluctation
	//
	// - FluctationDiscreate
	//
	// - Auto
	//
	// - Average
	//
	// - Variance
	//
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) GetType() *string {
	return s.Type
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) SetReferencedSamplesFilter(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) SetType(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig {
	s.Type = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig struct {
	// The name of the metric to be sampled:
	//
	// - Count: the number of rows in the table.
	//
	// - Min: the minimum value of the field.
	//
	// - Max: the maximum value of the field.
	//
	// - Avg: the average value of the field.
	//
	// - DistinctCount: the number of distinct values of the field.
	//
	// - DistinctPercent: the ratio of the number of distinct values of the field to the number of data rows.
	//
	// - DuplicatedCount: the number of duplicate values of the field.
	//
	// - DuplicatedPercent: the ratio of the number of duplicate values of the field to the number of data rows.
	//
	// - TableSize: the size of the table.
	//
	// - NullValueCount: the number of rows in which the field is null.
	//
	// - NullValuePercent: the percentage of rows in which the field is null.
	//
	// - GroupCount: the number of data rows corresponding to each value after aggregation by field value.
	//
	// - CountNotIn: the number of rows whose enumerated values do not match.
	//
	// - CountDistinctNotIn: the number of distinct values whose enumerated values do not match.
	//
	// - UserDefinedSql: collects samples by using a custom SQL statement.
	//
	// example:
	//
	// Max
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sample collection.
	//
	// example:
	//
	// {"SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The runtime parameter setting statements that are inserted and executed before the specific sampling statement is executed. The setting can be a maximum of 1,000 characters in length. Only MaxCompute is supported.
	//
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) SetMetric(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig {
	s.Metric = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) SetMetricParameters(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) SetSettingConfig(v string) *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig) Validate() error {
	return dara.Validate(s)
}
