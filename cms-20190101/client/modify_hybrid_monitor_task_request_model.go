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
	// The collection period of the metric. Valid values:
	//
	// 	- 15
	//
	// 	- 60
	//
	// Unit: seconds.
	//
	// example:
	//
	// 15
	CollectInterval *string `json:"CollectInterval,omitempty" xml:"CollectInterval,omitempty"`
	// The description of the metric import task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configurations of the logs that are imported from Simple Log Service.
	SLSProcessConfig *ModifyHybridMonitorTaskRequestSLSProcessConfig `json:"SLSProcessConfig,omitempty" xml:"SLSProcessConfig,omitempty" type:"Struct"`
	// The ID of the metric import task.
	//
	// For information about how to obtain the ID of a metric import task, see [DescribeHybridMonitorTaskList](https://help.aliyun.com/document_detail/428624.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 36****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the metric import task.
	//
	// For information about how to obtain the ID of a metric import task, see [DescribeHybridMonitorTaskList](https://help.aliyun.com/document_detail/428624.html).
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
	return dara.Validate(s)
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
	// The extended fields that specify the results of basic operations performed on aggregation results.
	Express []*ModifyHybridMonitorTaskRequestSLSProcessConfigExpress `json:"Express,omitempty" xml:"Express,omitempty" type:"Repeated"`
	// The conditions that are used to filter logs imported from Simple Log Service.
	Filter *ModifyHybridMonitorTaskRequestSLSProcessConfigFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The dimension based on which data is aggregated. This parameter is equivalent to the GROUP BY clause in SQL.
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
	return dara.Validate(s)
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigExpress struct {
	// The alias of the extended field that specifies the result of basic operations performed on aggregation results.
	//
	// example:
	//
	// SuccRate
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The extended field that specifies the result of basic operations performed on aggregation results.
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
	// 	- and (default): Logs are processed only if all filter conditions are met.
	//
	// 	- or: Logs are processed if one of the filter conditions is met.
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
	return dara.Validate(s)
}

type ModifyHybridMonitorTaskRequestSLSProcessConfigFilterFilters struct {
	// The method that is used to filter logs imported from Simple Log Service. Valid values:
	//
	// 	- `contain`: contains
	//
	// 	- `notContain`: does not contain
	//
	// 	- `>`: greater than
	//
	// 	- `<`: less than
	//
	// 	- `=`: equal to
	//
	// 	- `! =`: not equal to
	//
	// 	- `>=`: greater than or equal to
	//
	// 	- `<=`: less than or equal to
	//
	// example:
	//
	// =
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The name of the key that is used to filter logs imported from Simple Log Service.
	//
	// example:
	//
	// code
	SLSKeyName *string `json:"SLSKeyName,omitempty" xml:"SLSKeyName,omitempty"`
	// The value of the key that is used to filter logs imported from Simple Log Service.
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
	// The alias of the aggregation result.
	//
	// example:
	//
	// ApiResult
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the key that is used to aggregate logs imported from Simple Log Service.
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
	// The alias of the aggregation result.
	//
	// example:
	//
	// level_count
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The function that is used to aggregate the log data of a statistical period. Valid values:
	//
	// 	- count: counts the number.
	//
	// 	- sum: calculates the total value.
	//
	// 	- avg: calculates the average value.
	//
	// 	- max: calculates the maximum value.
	//
	// 	- min: calculates the minimum value.
	//
	// 	- value: collects samples within the statistical period.
	//
	// 	- countps: calculates the number of values of the specified field divided by the total number of seconds within the statistical period.
	//
	// 	- sumps: calculates the sum of the values of the specified field divided by the total number of seconds within the statistical period.
	//
	// 	- distinct: calculates the number of unique values of the specified field within the statistical period.
	//
	// 	- distribution: calculates the number of logs that meet a specified condition within the statistical period.
	//
	// 	- percentile: sorts the values of the specified field in ascending order, and then returns the value that is at the specified percentile within the statistical period. Example: P50.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The value of the function that is used to aggregate logs imported from Simple Log Service.
	//
	// 	- If the `Function` parameter is set to `distribution`, this parameter specifies the lower limit of the statistical interval. For example, if you want to calculate the number of HTTP requests whose status code is 2XX, set this parameter to 200.
	//
	// 	- If the `Function` parameter is set to `percentile`, this parameter specifies the percentile at which the expected value is. For example, 0.5 specifies P50.
	//
	// example:
	//
	// 200
	Parameter1 *string `json:"Parameter1,omitempty" xml:"Parameter1,omitempty"`
	// The value of the function that is used to aggregate logs imported from Simple Log Service.
	//
	// >  This parameter is required only if the `Function` parameter is set to `distribution`. This parameter specifies the upper limit of the statistical interval. For example, if you want to calculate the number of HTTP requests whose status code is 2XX, set this parameter to 299.
	//
	// example:
	//
	// 299
	Parameter2 *string `json:"Parameter2,omitempty" xml:"Parameter2,omitempty"`
	// The name of the key that is used to aggregate logs imported from Simple Log Service.
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
