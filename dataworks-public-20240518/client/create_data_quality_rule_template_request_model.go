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
	// The sample validation settings.
	CheckingConfig *CreateDataQualityRuleTemplateRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The category directory where the custom template is stored. Hierarchy levels are separated by slashes. Each level name can be up to 1024 characters long and cannot contain whitespace characters or slashes.
	//
	// example:
	//
	// /ods/order_data
	DirectoryPath *string `json:"DirectoryPath,omitempty" xml:"DirectoryPath,omitempty"`
	// The name of the rule template. It can be a combination of digits, English letters, Chinese characters, and half-width or full-width punctuation marks. The maximum length is 512 characters.
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
	// The visibility scope of the template:
	//
	// - Tenant: available to the entire tenant
	//
	// - Project: available only in the current project
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
	// Some threshold types require querying reference samples and then aggregating the values of those reference samples to derive the threshold used for comparison. An expression is used here to describe how the reference samples are queried.
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
	// The name of the metric to sample:
	//
	// - Count: number of table rows
	//
	// - Min: minimum field value
	//
	// - Max: maximum field value
	//
	// - Avg: average field value
	//
	// - DistinctCount: number of distinct field values
	//
	// - DistinctPercent: ratio of the number of distinct field values to the number of data rows
	//
	// - DuplicatedCount: number of duplicate field values
	//
	// - DuplicatedPercent: ratio of the number of duplicate field values to the number of data rows
	//
	// - TableSize: table size
	//
	// - NullValueCount: number of rows where the field is null
	//
	// - NullValuePercent: ratio of rows where the field is null
	//
	// - GroupCount: each value and the corresponding number of data rows after aggregating by field value
	//
	// - CountNotIn: number of rows whose enum value does not match
	//
	// - CountDistinctNotIn: number of distinct values whose enum value does not match
	//
	// - UserDefinedSql: sample collection via a custom SQL statement
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
	// Runtime parameter setting statements to be inserted and executed before the sampling statement is executed. The maximum length is 1000 characters. Currently only MaxCompute is supported.
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
