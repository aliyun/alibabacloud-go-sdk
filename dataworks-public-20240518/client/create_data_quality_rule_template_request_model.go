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
	// The sample verification settings.
	CheckingConfig *CreateDataQualityRuleTemplateRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The directory path where the custom template is stored. Levels are separated by forward slashes (/). Each level name can be up to 1024 characters in length and cannot contain whitespace characters or forward slashes.
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the rule template. The name can contain digits, letters, Chinese characters, and half-width or full-width punctuation marks. The name can be up to 512 characters in length.
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
	// The settings required for sample collection.
	SamplingConfig *CreateDataQualityRuleTemplateRequestSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The visibility scope of the template. Valid values:
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

type CreateDataQualityRuleTemplateRequestCheckingConfig struct {
	// An expression that specifies how to query reference samples. Some threshold types require querying reference samples and then aggregating their values to derive the threshold for comparison.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold calculation method. Valid values:
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
	// The name of the sampling metric. Valid values:
	//
	// - Count: the number of table rows.
	//
	// - Min: the minimum value of a field.
	//
	// - Max: the maximum value of a field.
	//
	// - Avg: the average value of a field.
	//
	// - DistinctCount: the number of distinct values in a field.
	//
	// - DistinctPercent: the ratio of distinct values to the total number of rows.
	//
	// - DuplicatedCount: the number of duplicate values in a field.
	//
	// - DuplicatedPercent: the ratio of duplicate values to the total number of rows.
	//
	// - TableSize: the table size.
	//
	// - NullValueCount: the number of rows where the field value is null.
	//
	// - NullValuePercent: the ratio of rows where the field value is null.
	//
	// - GroupCount: the count of rows for each value after aggregation by field value.
	//
	// - CountNotIn: the number of rows that do not match the enumerated values.
	//
	// - CountDistinctNotIn: the number of distinct values that do not match the enumerated values.
	//
	// - UserDefinedSql: sample collection through a custom SQL statement.
	//
	// example:
	//
	// Count
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sample collection.
	//
	// example:
	//
	// {"SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The runtime parameter setting statements to execute before the sampling statement. The value can be up to 1000 characters in length. Currently, only MaxCompute is supported.
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
