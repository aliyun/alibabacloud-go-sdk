// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfig(v *UpdateDataQualityRuleTemplateRequestCheckingConfig) *UpdateDataQualityRuleTemplateRequest
	GetCheckingConfig() *UpdateDataQualityRuleTemplateRequestCheckingConfig
	SetCode(v string) *UpdateDataQualityRuleTemplateRequest
	GetCode() *string
	SetDirectoryPath(v string) *UpdateDataQualityRuleTemplateRequest
	GetDirectoryPath() *string
	SetName(v string) *UpdateDataQualityRuleTemplateRequest
	GetName() *string
	SetProjectId(v int64) *UpdateDataQualityRuleTemplateRequest
	GetProjectId() *int64
	SetSamplingConfig(v *UpdateDataQualityRuleTemplateRequestSamplingConfig) *UpdateDataQualityRuleTemplateRequest
	GetSamplingConfig() *UpdateDataQualityRuleTemplateRequestSamplingConfig
}

type UpdateDataQualityRuleTemplateRequest struct {
	// The settings for sample validation.
	CheckingConfig *UpdateDataQualityRuleTemplateRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The code of the rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// USER_DEFINED:123
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The category directory in which the custom template is stored. Levels are separated by forward slashes (/). The name of each level can be up to 1024 characters in length and cannot contain whitespace characters or forward slashes (/).
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the rule template. The name can contain digits, letters, Chinese characters, and half-width and full-width punctuation marks. The name can be up to 512 characters in length.
	//
	// example:
	//
	// Table row Count Verification
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace used for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfig *UpdateDataQualityRuleTemplateRequestSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
}

func (s UpdateDataQualityRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleTemplateRequest) GetCheckingConfig() *UpdateDataQualityRuleTemplateRequestCheckingConfig {
	return s.CheckingConfig
}

func (s *UpdateDataQualityRuleTemplateRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateDataQualityRuleTemplateRequest) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *UpdateDataQualityRuleTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityRuleTemplateRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityRuleTemplateRequest) GetSamplingConfig() *UpdateDataQualityRuleTemplateRequestSamplingConfig {
	return s.SamplingConfig
}

func (s *UpdateDataQualityRuleTemplateRequest) SetCheckingConfig(v *UpdateDataQualityRuleTemplateRequestCheckingConfig) *UpdateDataQualityRuleTemplateRequest {
	s.CheckingConfig = v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequest) SetCode(v string) *UpdateDataQualityRuleTemplateRequest {
	s.Code = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequest) SetDirectoryPath(v string) *UpdateDataQualityRuleTemplateRequest {
	s.DirectoryPath = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequest) SetName(v string) *UpdateDataQualityRuleTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequest) SetProjectId(v int64) *UpdateDataQualityRuleTemplateRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequest) SetSamplingConfig(v *UpdateDataQualityRuleTemplateRequestSamplingConfig) *UpdateDataQualityRuleTemplateRequest {
	s.SamplingConfig = v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequest) Validate() error {
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

type UpdateDataQualityRuleTemplateRequestCheckingConfig struct {
	// For some types of thresholds, you must query reference samples and then aggregate the values of the reference samples to obtain the threshold used for comparison. An expression is used here to indicate the query method of reference samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The type of the monitored object. Valid values:
	//
	// - Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityRuleTemplateRequestCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleTemplateRequestCheckingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleTemplateRequestCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *UpdateDataQualityRuleTemplateRequestCheckingConfig) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityRuleTemplateRequestCheckingConfig) SetReferencedSamplesFilter(v string) *UpdateDataQualityRuleTemplateRequestCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequestCheckingConfig) SetType(v string) *UpdateDataQualityRuleTemplateRequestCheckingConfig {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequestCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityRuleTemplateRequestSamplingConfig struct {
	// The name of the metric to be sampled. Valid values:
	//
	// - Count: the number of table rows.
	//
	// - Min: the minimum value of a field.
	//
	// - Max: the maximum value of a field.
	//
	// - Avg: the average value of a field.
	//
	// - DistinctCount: the number of distinct values of a field.
	//
	// - DistinctPercent: the ratio of the number of distinct values of a field to the number of data rows.
	//
	// - DuplicatedCount: the number of duplicate values of a field.
	//
	// - DuplicatedPercent: the ratio of the number of duplicate values of a field to the number of data rows.
	//
	// - TableSize: the size of the table.
	//
	// - NullValueCount: the number of rows in which the field is null.
	//
	// - NullValuePercent: the percentage of rows in which the field is null.
	//
	// - GroupCount: the number of data rows corresponding to each value after the field values are aggregated.
	//
	// - CountNotIn: the number of rows whose values do not match the enumerated values.
	//
	// - CountDistinctNotIn: the number of distinct values that do not match the enumerated values.
	//
	// - UserDefinedSql: sample collection by using custom SQL.
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
	// The runtime parameter setting statements that are inserted and executed before the sampling statements are executed. The statements can be up to 1,000 characters in length. Only MaxCompute is supported.
	//
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s UpdateDataQualityRuleTemplateRequestSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleTemplateRequestSamplingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) SetMetric(v string) *UpdateDataQualityRuleTemplateRequestSamplingConfig {
	s.Metric = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) SetMetricParameters(v string) *UpdateDataQualityRuleTemplateRequestSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) SetSettingConfig(v string) *UpdateDataQualityRuleTemplateRequestSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *UpdateDataQualityRuleTemplateRequestSamplingConfig) Validate() error {
	return dara.Validate(s)
}
