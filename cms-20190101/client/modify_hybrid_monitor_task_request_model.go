// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachLabels(v []*ModifyHybridMonitorTaskRequestAttachLabels) *ModifyHybridMonitorTaskRequest
	GetAttachLabels() []*ModifyHybridMonitorTaskRequestAttachLabels
	SetCollectInterval(v string) *ModifyHybridMonitorTaskRequest
	GetCollectInterval() *string
	SetDescription(v string) *ModifyHybridMonitorTaskRequest
	GetDescription() *string
	SetRegionId(v string) *ModifyHybridMonitorTaskRequest
	GetRegionId() *string
	SetSLSProcessConfig(v *ModifyHybridMonitorTaskRequestSLSProcessConfig) *ModifyHybridMonitorTaskRequest
	GetSLSProcessConfig() *ModifyHybridMonitorTaskRequestSLSProcessConfig
	SetTaskId(v string) *ModifyHybridMonitorTaskRequest
	GetTaskId() *string
	SetTaskName(v string) *ModifyHybridMonitorTaskRequest
	GetTaskName() *string
}

type ModifyHybridMonitorTaskRequest struct {
	// The tags of the metric.
	AttachLabels []*ModifyHybridMonitorTaskRequestAttachLabels `json:"AttachLabels,omitempty" xml:"AttachLabels,omitempty" type:"Repeated"`
	// The collection interval of the metric. Valid values:
	//
	// - 15
	//
	// - 60
	//
	// Unit: seconds.
	//
	// example:
	//
	// 15
	CollectInterval *string `json:"CollectInterval,omitempty" xml:"CollectInterval,omitempty"`
	// The description of the monitoring task.
	//
	// example:
	//
	// SLS log monitoring data.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SLS log configuration.
	SLSProcessConfig *ModifyHybridMonitorTaskRequestSLSProcessConfig `json:"SLSProcessConfig,omitempty" xml:"SLSProcessConfig,omitempty" type:"Struct"`
	// The monitoring task ID.
	//
	// For information about how to obtain the monitoring task ID, see [DescribeHybridMonitorTaskList](https://help.aliyun.com/document_detail/428624.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 36****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The monitoring task name.
	//
	// For information about how to obtain the monitoring task ID, see [DescribeHybridMonitorTaskList](https://help.aliyun.com/document_detail/428624.html).
	//
	// example:
	//
	// SLS_task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ModifyHybridMonitorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequest) GetAttachLabels() []*ModifyHybridMonitorTaskRequestAttachLabels {
	return s.AttachLabels
}

func (s *ModifyHybridMonitorTaskRequest) GetCollectInterval() *string {
	return s.CollectInterval
}

func (s *ModifyHybridMonitorTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyHybridMonitorTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridMonitorTaskRequest) GetSLSProcessConfig() *ModifyHybridMonitorTaskRequestSLSProcessConfig {
	return s.SLSProcessConfig
}

func (s *ModifyHybridMonitorTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyHybridMonitorTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyHybridMonitorTaskRequest) SetAttachLabels(v []*ModifyHybridMonitorTaskRequestAttachLabels) *ModifyHybridMonitorTaskRequest {
	s.AttachLabels = v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) SetCollectInterval(v string) *ModifyHybridMonitorTaskRequest {
	s.CollectInterval = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) SetDescription(v string) *ModifyHybridMonitorTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) SetRegionId(v string) *ModifyHybridMonitorTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) SetSLSProcessConfig(v *ModifyHybridMonitorTaskRequestSLSProcessConfig) *ModifyHybridMonitorTaskRequest {
	s.SLSProcessConfig = v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) SetTaskId(v string) *ModifyHybridMonitorTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) SetTaskName(v string) *ModifyHybridMonitorTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequest) Validate() error {
	if s.AttachLabels != nil {
		for _, item := range s.AttachLabels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SLSProcessConfig != nil {
		if err := s.SLSProcessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyHybridMonitorTaskRequestAttachLabels struct {
	// The tag key of the metric.
	//
	// example:
	//
	// app_service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tag value of the metric.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyHybridMonitorTaskRequestAttachLabels) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestAttachLabels) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestAttachLabels) GetName() *string {
	return s.Name
}

func (s *ModifyHybridMonitorTaskRequestAttachLabels) GetValue() *string {
	return s.Value
}

func (s *ModifyHybridMonitorTaskRequestAttachLabels) SetName(v string) *ModifyHybridMonitorTaskRequestAttachLabels {
	s.Name = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestAttachLabels) SetValue(v string) *ModifyHybridMonitorTaskRequestAttachLabels {
	s.Value = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestAttachLabels) Validate() error {
	return dara.Validate(s)
}

type ModifyHybridMonitorTaskRequestSLSProcessConfig struct {
	// The arithmetic operation result of the extended field in the SLS log statistics result.
	Express []*ModifyHybridMonitorTaskRequestSLSProcessConfigExpress `json:"Express,omitempty" xml:"Express,omitempty" type:"Repeated"`
	// The filter conditions for parameters in the SLS log.
	Filter *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// Aggregates data by spatial dimension, which is equivalent to GROUP BY in SQL.
	GroupBy []*ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy `json:"GroupBy,omitempty" xml:"GroupBy,omitempty" type:"Repeated"`
	// None.
	Statistics []*ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfig) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) GetExpress() []*ModifyHybridMonitorTaskRequestSLSProcessConfigExpress {
	return s.Express
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) GetFilter() *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter {
	return s.Filter
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) GetGroupBy() []*ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy {
	return s.GroupBy
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) GetStatistics() []*ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics {
	return s.Statistics
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) SetExpress(v []*ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) *ModifyHybridMonitorTaskRequestSLSProcessConfig {
	s.Express = v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) SetFilter(v *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) *ModifyHybridMonitorTaskRequestSLSProcessConfig {
	s.Filter = v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) SetGroupBy(v []*ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) *ModifyHybridMonitorTaskRequestSLSProcessConfig {
	s.GroupBy = v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) SetStatistics(v []*ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) *ModifyHybridMonitorTaskRequestSLSProcessConfig {
	s.Statistics = v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfig) Validate() error {
	if s.Express != nil {
		for _, item := range s.Express {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	if s.GroupBy != nil {
		for _, item := range s.GroupBy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigExpress struct {
	// The alias of the arithmetic operation result of the extended field in the SLS log statistics result.
	//
	// example:
	//
	// SuccRate
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The arithmetic operation result of the extended field in the SLS log statistics result.
	//
	// example:
	//
	// success_count
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) GetAlias() *string {
	return s.Alias
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) GetExpress() *string {
	return s.Express
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) SetAlias(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress {
	s.Alias = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) SetExpress(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress {
	s.Express = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigExpress) Validate() error {
	return dara.Validate(s)
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigFilter struct {
	// None.
	Filters []*ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The relationship between multiple filter conditions. Valid values:
	//
	// - and (default): Logs are processed only when all filter conditions are met.
	//
	// - or: Logs are processed when any filter condition is met.
	//
	// example:
	//
	// and
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) GetFilters() []*ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	return s.Filters
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) GetRelation() *string {
	return s.Relation
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) SetFilters(v []*ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter {
	s.Filters = v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) SetRelation(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter {
	s.Relation = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters struct {
	// The method used to filter parameter values in the SLS log. Valid values:
	//
	// - `contain`: contains.
	//
	// - `notContain`: does not contain.
	//
	// - `>`: greater than.
	//
	// - `<`: less than.
	//
	// - `=`: equal to.
	//
	// - `!=`: not equal to.
	//
	// - `>=`: greater than or equal to.
	//
	// - `<=`: less than or equal to.
	//
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The name of the parameter to filter in the SLS log.
	//
	// example:
	//
	// code
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
	// The filter value of the parameter in the SLS log.
	//
	// example:
	//
	// 200
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GetOperator() *string {
	return s.Operator
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GetValue() *string {
	return s.Value
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) SetOperator(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	s.Operator = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) SetSLSKeyName(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	s.SLSKeyName = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) SetValue(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	s.Value = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters) Validate() error {
	return dara.Validate(s)
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy struct {
	// The alias of the SLS log statistics result.
	//
	// example:
	//
	// ApiResult
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the parameter for SLS log statistics.
	//
	// example:
	//
	// code
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) GetAlias() *string {
	return s.Alias
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) SetAlias(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy {
	s.Alias = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) SetSLSKeyName(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy {
	s.SLSKeyName = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigGroupBy) Validate() error {
	return dara.Validate(s)
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics struct {
	// The alias of the SLS log statistics result.
	//
	// example:
	//
	// level_count
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The statistical method used to aggregate log data within a statistical period. Valid values:
	//
	// - count: counts the number of occurrences.
	//
	// - sum: calculates the sum.
	//
	// - avg: calculates the average.
	//
	// - max: returns the maximum value.
	//
	// - min: returns the minimum value.
	//
	// - value: samples within the period.
	//
	// - countps: calculates the per-second average of the count for the specified field within the statistical period.
	//
	// - sumps: calculates the per-second average of the sum for the specified field within the statistical period.
	//
	// - distinct: calculates the number of occurrences of the specified field after deduplication within the statistical period.
	//
	// - distribution: calculates the number of occurrences of field values within a specified range.
	//
	// - percentile: calculates the distribution value of field values, such as P50.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The statistical value of the SLS log.
	//
	// - If `Function` is set to `distribution`, this parameter specifies the lower limit of the statistical range. For example, to count the number of 2XX HTTP status codes, set this parameter to 200.
	//
	// - If `Function` is set to `percentile`, this parameter specifies the percentile of the statistical distribution. For example, 0.5 indicates P50.
	//
	// example:
	//
	// 200
	Parameter1 *string `json:"Parameter1,omitempty" xml:"Parameter1,omitempty"`
	// The statistical value of the SLS log.
	//
	// > This parameter is required only when `Function` is set to `distribution`. It specifies the upper limit of the statistical range. For example, to count the number of 2XX HTTP status codes, set this parameter to 299.
	//
	// example:
	//
	// 299
	Parameter2 *string `json:"Parameter2,omitempty" xml:"Parameter2,omitempty"`
	// The name of the parameter for SLS log statistics.
	//
	// example:
	//
	// name
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) GetAlias() *string {
	return s.Alias
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) GetFunction() *string {
	return s.Function
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) GetParameter1() *string {
	return s.Parameter1
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) GetParameter2() *string {
	return s.Parameter2
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) SetAlias(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Alias = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) SetFunction(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Function = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) SetParameter1(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Parameter1 = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) SetParameter2(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Parameter2 = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) SetSLSKeyName(v string) *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.SLSKeyName = &v
	return s
}

func (s *ModifyHybridMonitorTaskRequestSLSProcessConfigStatistics) Validate() error {
	return dara.Validate(s)
}
