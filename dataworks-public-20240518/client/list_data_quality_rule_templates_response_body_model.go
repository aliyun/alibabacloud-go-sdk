// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityRuleTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataQualityRuleTemplatesResponseBodyPagingInfo) *ListDataQualityRuleTemplatesResponseBody
	GetPagingInfo() *ListDataQualityRuleTemplatesResponseBodyPagingInfo
	SetRequestId(v string) *ListDataQualityRuleTemplatesResponseBody
	GetRequestId() *string
}

type ListDataQualityRuleTemplatesResponseBody struct {
	// The pagination information.
	PagingInfo *ListDataQualityRuleTemplatesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityRuleTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesResponseBody) GetPagingInfo() *ListDataQualityRuleTemplatesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataQualityRuleTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityRuleTemplatesResponseBody) SetPagingInfo(v *ListDataQualityRuleTemplatesResponseBodyPagingInfo) *ListDataQualityRuleTemplatesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBody) SetRequestId(v string) *ListDataQualityRuleTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityRuleTemplatesResponseBodyPagingInfo struct {
	// The templates.
	DataQualityRuleTemplates []*ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates `json:"DataQualityRuleTemplates,omitempty" xml:"DataQualityRuleTemplates,omitempty" type:"Repeated"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of entries
	//
	// example:
	//
	// 42
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) GetDataQualityRuleTemplates() []*ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	return s.DataQualityRuleTemplates
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) SetDataQualityRuleTemplates(v []*ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) *ListDataQualityRuleTemplatesResponseBodyPagingInfo {
	s.DataQualityRuleTemplates = v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataQualityRuleTemplatesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) SetPageSize(v int32) *ListDataQualityRuleTemplatesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataQualityRuleTemplatesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfo) Validate() error {
	if s.DataQualityRuleTemplates != nil {
		for _, item := range s.DataQualityRuleTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates struct {
	// Sample verification settings
	CheckingConfig *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// Rule template Code
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
	// DataWorks workspace ID
	//
	// example:
	//
	// 2043
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Settings required for sample collection
	SamplingConfig *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
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

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetCheckingConfig() *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig {
	return s.CheckingConfig
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetCode() *string {
	return s.Code
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetDirectoryPath() *string {
	return s.DirectoryPath
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetName() *string {
	return s.Name
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetSamplingConfig() *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig {
	return s.SamplingConfig
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) GetVisibleScope() *string {
	return s.VisibleScope
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetCheckingConfig(v *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.CheckingConfig = v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetCode(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.Code = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetDirectoryPath(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.DirectoryPath = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetName(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.Name = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetProjectId(v int64) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetSamplingConfig(v *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.SamplingConfig = v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) SetVisibleScope(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates {
	s.VisibleScope = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplates) Validate() error {
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

type ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig struct {
	// Some types of thresholds need to query some reference samples, and then summarize the values of the reference samples to obtain the threshold for comparison. Here, an expression is used to represent the query method of the reference samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// Threshold Calculation method
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

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) GetType() *string {
	return s.Type
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) SetReferencedSamplesFilter(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) SetType(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig {
	s.Type = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig struct {
	// The name of the sampled metric.
	//
	// - Count: number of table rows
	//
	// - Min: minimum value of the field
	//
	// - Max: The maximum value of the field.
	//
	// - Avg: field mean
	//
	// - DistinctCount: number of unique field values
	//
	// - DistinctPercent: the ratio of the number of unique field values to the number of data rows.
	//
	// - DuplicatedCount: number of duplicate field values
	//
	// - DuplicatedPercent: the ratio of the number of duplicate field values to the number of data rows.
	//
	// - TableSize: table size
	//
	// - NullValueCount: number of rows with empty fields
	//
	// - NullValuePercent: the proportion of fields that are empty.
	//
	// - GroupCount: aggregate each value by field value and the corresponding number of data rows
	//
	// - CountNotIn: the enumerated value does not match the number of rows.
	//
	// - CountDistinctNotIn: the number of unique values that the enumerated values do not match.
	//
	// - UserDefinedSql: use custom SQL to collect samples
	//
	// example:
	//
	// Max
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// Parameters required for sample collection
	//
	// example:
	//
	// {"Sql": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// Before executing the sample statement, insert some runtime parameter setting statements, which can be up to 1000 characters in length. Currently, only MaxCompute are supported.
	//
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) SetMetric(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig {
	s.Metric = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) SetMetricParameters(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) SetSettingConfig(v string) *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *ListDataQualityRuleTemplatesResponseBodyPagingInfoDataQualityRuleTemplatesSamplingConfig) Validate() error {
	return dara.Validate(s)
}
