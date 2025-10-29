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
	// The information about the template.
	DataQualityRuleTemplate *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplate `json:"DataQualityRuleTemplate,omitempty" xml:"DataQualityRuleTemplate,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
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
	// The check settings for sample data.
	CheckingConfig *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The code for the template.
	//
	// example:
	//
	// USER_DEFINED:123
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The directory in which the template is stored. Slashes (/) are used to separate directory levels. The name of each directory level can be up to 1,024 characters in length. It cannot contain whitespace characters or slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the template. The name can be up to 512 characters in length and can contain digits, letters, and punctuation marks.
	//
	// example:
	//
	// Table row Count Verification
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 4020
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
	SamplingConfig *GetDataQualityRuleTemplateResponseBodyDataQualityRuleTemplateSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// Available range of templates:
	//
	// - Tenant: all tenants are available
	//
	// - Project: only available in the current Project
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
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference samples and perform aggregate operations on the reference values. In this example, an expression is used to indicate the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold calculation method. Valid values:
	//
	// 	- Fixed
	//
	// 	- Fluctation
	//
	// 	- FluctationDiscreate
	//
	// 	- Auto
	//
	// 	- Average
	//
	// 	- Variance
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
	// The metrics used for sampling. Valid values:
	//
	// 	- Count: the number of rows in the table.
	//
	// 	- Min: the minimum value of the field.
	//
	// 	- Max: the maximum value of the field.
	//
	// 	- Avg: the average value of the field.
	//
	// 	- DistinctCount: the number of unique values of the field after deduplication.
	//
	// 	- DistinctPercent: the proportion of the number of unique values of the field after deduplication to the number of rows in the table.
	//
	// 	- DuplicatedCount: the number of duplicated values of the field.
	//
	// 	- DuplicatedPercent: the proportion of the number of duplicated values of the field to the number of rows in the table.
	//
	// 	- TableSize: the table size.
	//
	// 	- NullValueCount: the number of rows in which the field value is null.
	//
	// 	- NullValuePercent: the proportion of the number of rows in which the field value is null to the number of rows in the table.
	//
	// 	- GroupCount: the field value and the number of rows for each field value.
	//
	// 	- CountNotIn: the number of rows in which the field values are different from the referenced values that you specified in the rule.
	//
	// 	- CountDistinctNotIn: the number of unique values that are different from the referenced values that you specified in the rule after deduplication.
	//
	// 	- UserDefinedSql: indicates that data is sampled by executing custom SQL statements.
	//
	// example:
	//
	// Max
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sampling.
	//
	// example:
	//
	// {"SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The statements that are used to configure the parameters required for sampling before you execute the sampling statements. The statements can be up to 1,000 characters in length. Only the MaxCompute database is supported.
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
