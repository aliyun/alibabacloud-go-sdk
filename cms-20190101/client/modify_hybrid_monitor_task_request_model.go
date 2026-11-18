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
	// The labels of the metric.
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
	// SLS Log Monitoring data.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SLS log configuration.
	SLSProcessConfig *ModifyHybridMonitorTaskRequestSLSProcessConfig `json:"SLSProcessConfig,omitempty" xml:"SLSProcessConfig,omitempty" type:"Struct"`
	// The ID of the monitoring task.
	//
	// For information about how to obtain the monitoring task ID, see [DescribeHybridMonitorTaskList](https://help.aliyun.com/document_detail/428624.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 36****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the monitoring task.
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
	// The key of the metric label.
	//
	// example:
	//
	// app_service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the metric label.
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
	// The result of arithmetic operations on the extended fields of the SLS log statistics.
	Express []*ModifyHybridMonitorTaskRequestSLSProcessConfigExpress `json:"Express,omitempty" xml:"Express,omitempty" type:"Repeated"`
	// The filter conditions for parameters in SLS logs.
	Filter *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The spatial dimension for data aggregation, equivalent to the GROUP BY clause in SQL.
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
	// The alias for the result of arithmetic operations on the extended fields of the SLS log statistics.
	//
	// example:
	//
	// SuccRate
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The result of arithmetic operations on the extended fields of the SLS log statistics.
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
	// The relationship between multiple filter conditions for parameters. Valid values:
	//
	// - and: logs are processed only when all filter conditions are met. This is the default value.
	//
	// - or: logs are processed when any filter condition is met.
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
	// The filter method for parameter values in SLS logs. Valid values:
	//
	// - `contain`: includes
	//
	// - `notContain`: does not include
	//
	// - `>`: greater than
	//
	// - `<`: less than
	//
	// - `=`: equal to
	//
	// - `!=`: not equal to
	//
	// - `>=`: greater than or equal to
	//
	// - `<=`: less than or equal to
	//
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The name of the parameter to be filtered in SLS logs.
	//
	// example:
	//
	// code
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
	// The filter value of the parameter in SLS logs.
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
	// The alias for the SLS log statistics result.
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
	// The alias for the SLS log statistics result.
	//
	// example:
	//
	// level_count
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The aggregation method used to process log data within the calculation period. Valid values:
	//
	// - count: count
	//
	// - sum: sum
	//
	// - avg: average value
	//
	// - max: maximum value
	//
	// - min: minimum value
	//
	// - value: sampling within the period
	//
	// - countps: average count per second of the specified field within the statistical period
	//
	// - sumps: average sum per second of the specified field within the statistical period
	//
	// - distinct: count of occurrences of the specified field after deduplication within the statistical period
	//
	// - distribution: count of occurrences of field values within a specified range
	//
	// - percentile: percentile of field values, such as P50
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The statistical value of SLS logs.
	//
	// - If `Function` is set to `distribution`, this parameter specifies the lower limit of the statistical range. For example, to count the number of 2XX `HttpCode` status codes, enter 200.
	//
	// - If `Function` is set to `percentile`, this parameter specifies the percentile of the statistical distribution. For example, 0.5 indicates P50.
	//
	// example:
	//
	// 200
	Parameter1 *string `json:"Parameter1,omitempty" xml:"Parameter1,omitempty"`
	// The statistical value of SLS logs.
	//
	// > This parameter is required only when `Function` is set to `distribution`. It specifies the upper limit of the statistical range. For example, to count the number of 2XX `HttpCode` status codes, enter 299.
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
