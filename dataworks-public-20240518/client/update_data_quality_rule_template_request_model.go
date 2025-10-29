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
	// The check settings for sample data.
	CheckingConfig *UpdateDataQualityRuleTemplateRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The code for the template.
	//
	// This parameter is required.
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
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
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
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference samples and perform aggregate operations on the reference values. In this example, an expression is used to specify the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The type of the monitored object. Valid values:
	//
	// 	- Table
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
