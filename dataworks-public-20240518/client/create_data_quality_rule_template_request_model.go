// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfig(v *CreateDataQualityRuleTemplateRequestCheckingConfig) *CreateDataQualityRuleTemplateRequest
	GetCheckingConfig() *CreateDataQualityRuleTemplateRequestCheckingConfig
	SetDirectoryPath(v string) *CreateDataQualityRuleTemplateRequest
	GetDirectoryPath() *string
	SetName(v string) *CreateDataQualityRuleTemplateRequest
	GetName() *string
	SetProjectId(v int64) *CreateDataQualityRuleTemplateRequest
	GetProjectId() *int64
	SetSamplingConfig(v *CreateDataQualityRuleTemplateRequestSamplingConfig) *CreateDataQualityRuleTemplateRequest
	GetSamplingConfig() *CreateDataQualityRuleTemplateRequestSamplingConfig
	SetVisibleScope(v string) *CreateDataQualityRuleTemplateRequest
	GetVisibleScope() *string
}

type CreateDataQualityRuleTemplateRequest struct {
	// The check settings for sample data.
	CheckingConfig *CreateDataQualityRuleTemplateRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The directory in which the template is stored. Slashes (/) are used to separate directory levels. The name of each directory level can be up to 1,024 characters in length. It cannot contain whitespace characters or slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the template. The name can be up to 512 characters in length and can contain digits, letters, and punctuation marks.
	//
	// This parameter is required.
	//
	// example:
	//
	// Table row Count Verification
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
	SamplingConfig *CreateDataQualityRuleTemplateRequestSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The applicable scope of the template. Valid values:
	//
	// 	- Tenant: The template is available in all workspaces in the current tenant.
	//
	// 	- Project: The template is available only in the current workspace.
	//
	// example:
	//
	// Project
	VisibleScope *string `json:"VisibleScope,omitempty" xml:"VisibleScope,omitempty"`
}

func (s CreateDataQualityRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleTemplateRequest) GetCheckingConfig() *CreateDataQualityRuleTemplateRequestCheckingConfig {
	return s.CheckingConfig
}

func (s *CreateDataQualityRuleTemplateRequest) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *CreateDataQualityRuleTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityRuleTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityRuleTemplateRequest) GetSamplingConfig() *CreateDataQualityRuleTemplateRequestSamplingConfig {
	return s.SamplingConfig
}

func (s *CreateDataQualityRuleTemplateRequest) GetVisibleScope() *string {
	return s.VisibleScope
}

func (s *CreateDataQualityRuleTemplateRequest) SetCheckingConfig(v *CreateDataQualityRuleTemplateRequestCheckingConfig) *CreateDataQualityRuleTemplateRequest {
	s.CheckingConfig = v
	return s
}

func (s *CreateDataQualityRuleTemplateRequest) SetDirectoryPath(v string) *CreateDataQualityRuleTemplateRequest {
	s.DirectoryPath = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequest) SetName(v string) *CreateDataQualityRuleTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequest) SetProjectId(v int64) *CreateDataQualityRuleTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequest) SetSamplingConfig(v *CreateDataQualityRuleTemplateRequestSamplingConfig) *CreateDataQualityRuleTemplateRequest {
	s.SamplingConfig = v
	return s
}

func (s *CreateDataQualityRuleTemplateRequest) SetVisibleScope(v string) *CreateDataQualityRuleTemplateRequest {
	s.VisibleScope = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleTemplateRequestCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference samples and perform aggregate operations on the reference values. In this example, an expression is used to specify the query method of referenced samples.
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

func (s CreateDataQualityRuleTemplateRequestCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleTemplateRequestCheckingConfig) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleTemplateRequestCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *CreateDataQualityRuleTemplateRequestCheckingConfig) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityRuleTemplateRequestCheckingConfig) SetReferencedSamplesFilter(v string) *CreateDataQualityRuleTemplateRequestCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequestCheckingConfig) SetType(v string) *CreateDataQualityRuleTemplateRequestCheckingConfig {
	s.Type = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequestCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleTemplateRequestSamplingConfig struct {
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
	// 	- UserDefinedSql: specifies that data is sampled by executing custom SQL statements.
	//
	// example:
	//
	// Count
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

func (s CreateDataQualityRuleTemplateRequestSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleTemplateRequestSamplingConfig) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) SetMetric(v string) *CreateDataQualityRuleTemplateRequestSamplingConfig {
	s.Metric = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) SetMetricParameters(v string) *CreateDataQualityRuleTemplateRequestSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) SetSettingConfig(v string) *CreateDataQualityRuleTemplateRequestSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *CreateDataQualityRuleTemplateRequestSamplingConfig) Validate() error {
	return dara.Validate(s)
}
