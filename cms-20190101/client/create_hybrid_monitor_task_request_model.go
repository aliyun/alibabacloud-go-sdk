// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachLabels(v []*CreateHybridMonitorTaskRequestAttachLabels) *CreateHybridMonitorTaskRequest
	GetAttachLabels() []*CreateHybridMonitorTaskRequestAttachLabels
	SetCloudAccessId(v []*string) *CreateHybridMonitorTaskRequest
	GetCloudAccessId() []*string
	SetCollectInterval(v string) *CreateHybridMonitorTaskRequest
	GetCollectInterval() *string
	SetCollectTargetType(v string) *CreateHybridMonitorTaskRequest
	GetCollectTargetType() *string
	SetDescription(v string) *CreateHybridMonitorTaskRequest
	GetDescription() *string
	SetGroupId(v string) *CreateHybridMonitorTaskRequest
	GetGroupId() *string
	SetNamespace(v string) *CreateHybridMonitorTaskRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateHybridMonitorTaskRequest
	GetRegionId() *string
	SetSLSProcessConfig(v *CreateHybridMonitorTaskRequestSLSProcessConfig) *CreateHybridMonitorTaskRequest
	GetSLSProcessConfig() *CreateHybridMonitorTaskRequestSLSProcessConfig
	SetTargetUserId(v string) *CreateHybridMonitorTaskRequest
	GetTargetUserId() *string
	SetTargetUserIdList(v string) *CreateHybridMonitorTaskRequest
	GetTargetUserIdList() *string
	SetTaskName(v string) *CreateHybridMonitorTaskRequest
	GetTaskName() *string
	SetTaskType(v string) *CreateHybridMonitorTaskRequest
	GetTaskType() *string
	SetYARMConfig(v string) *CreateHybridMonitorTaskRequest
	GetYARMConfig() *string
}

type CreateHybridMonitorTaskRequest struct {
	// The labels of the metric.
	//
	// > This parameter is required only when `TaskType` is set to `aliyun_sls`.
	AttachLabels []*CreateHybridMonitorTaskRequestAttachLabels `json:"AttachLabels,omitempty" xml:"AttachLabels,omitempty" type:"Repeated"`
	// The list of primary account IDs of other cloud providers that are imported into CloudMonitor.
	CloudAccessId []*string `json:"CloudAccessId,omitempty" xml:"CloudAccessId,omitempty" type:"Repeated"`
	// The collection period of the metric. Valid values:
	//
	// - 15
	//
	// - 60 (default)
	//
	// Unit: seconds.
	//
	// > This parameter is required only when `TaskType` is set to `aliyun_sls`.
	//
	// example:
	//
	// 60
	CollectInterval *string `json:"CollectInterval,omitempty" xml:"CollectInterval,omitempty"`
	// The type of the collection target for monitoring data.
	//
	// - If `TaskType` is set to `aliyun_fc`, enter `aliyun_fc`.
	//
	// - If `TaskType` is set to `aliyun_sls`, enter the name of the Logstore group for SLS logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun_fc
	CollectTargetType *string `json:"CollectTargetType,omitempty" xml:"CollectTargetType,omitempty"`
	// The description of the monitoring task.
	//
	// example:
	//
	// Alibaba Cloud service data import task.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group ID.
	//
	// For information about how to obtain the application group ID, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// > This parameter is required only when `TaskType` is set to `aliyun_sls`.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the namespace.
	//
	// For information about how to obtain the namespace name, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SLS log configuration.
	//
	// > This parameter is required only when `TaskType` is set to `aliyun_sls`.
	SLSProcessConfig *CreateHybridMonitorTaskRequestSLSProcessConfig `json:"SLSProcessConfig,omitempty" xml:"SLSProcessConfig,omitempty" type:"Struct"`
	// The member ID.
	//
	// When you invoke this operation by using a management account, you can use cloud native mode to import Alibaba Cloud services of any member in the resource folder into Hybrid Cloud Monitoring. This allows you to use the resource directory to monitor Alibaba Cloud services across accounts in CloudMonitor. You can go to **Settings*	- to configure this feature.
	//
	// > This parameter is required only when `TaskType` is set to `aliyun_fc`.
	//
	// example:
	//
	// 120886317861****
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
	// The member IDs. Separate multiple member IDs with commas (,).
	//
	// > This parameter is required only when you call this operation by using a management account.
	//
	// example:
	//
	// 120886317861****
	TargetUserIdList *string `json:"TargetUserIdList,omitempty" xml:"TargetUserIdList,omitempty"`
	// The name of the monitoring task.
	//
	// - If `TaskType` is set to `aliyun_fc`, enter the name of the data import task for the Alibaba Cloud service.
	//
	// - If `TaskType` is set to `aliyun_sls`, enter the metric name for SLS logs.
	//
	// example:
	//
	// aliyun_task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the monitoring task. Valid values:
	//
	// - aliyun_fc: a data import task for Alibaba Cloud services.
	//
	// - aliyun_sls: a metric for SLS logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun_fc
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The configuration file of the Alibaba Cloud service to be imported into Hybrid Cloud Monitoring.
	//
	// - namespace: the namespace of the Alibaba Cloud service. For information about how to query the namespace, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html).
	//
	// - metric_list: the metrics of the Alibaba Cloud service. For information about how to query the metrics, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html).
	//
	// The following example shows the configuration file:
	//
	// ```
	//
	// products:
	//
	// - namespace: acs_ecs_dashboard
	//
	//   metric_info:
	//
	//   - metric_list:
	//
	//     - cpu_total
	//
	//     - cpu_idle
	//
	//     - diskusage_utilization
	//
	//     - CPUUtilization
	//
	//     - DiskReadBPS
	//
	//     - InternetOut
	//
	//     - IntranetOut
	//
	//     - cpu_system
	//
	// - namespace: acs_rds_dashboard
	//
	//   metric_info:
	//
	//   - metric_list:
	//
	//     - MySQL_QPS
	//
	//     - MySQL_TPS
	//
	// ```
	//
	// > This parameter is required when `TaskType` is set to `aliyun_fc`.
	//
	// example:
	//
	// products:- namespace: acs_ecs_dashboard  metric_info:  - metric_list:    - cpu_total
	YARMConfig *string `json:"YARMConfig,omitempty" xml:"YARMConfig,omitempty"`
}

func (s CreateHybridMonitorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequest) GetAttachLabels() []*CreateHybridMonitorTaskRequestAttachLabels {
	return s.AttachLabels
}

func (s *CreateHybridMonitorTaskRequest) GetCloudAccessId() []*string {
	return s.CloudAccessId
}

func (s *CreateHybridMonitorTaskRequest) GetCollectInterval() *string {
	return s.CollectInterval
}

func (s *CreateHybridMonitorTaskRequest) GetCollectTargetType() *string {
	return s.CollectTargetType
}

func (s *CreateHybridMonitorTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHybridMonitorTaskRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateHybridMonitorTaskRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateHybridMonitorTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHybridMonitorTaskRequest) GetSLSProcessConfig() *CreateHybridMonitorTaskRequestSLSProcessConfig {
	return s.SLSProcessConfig
}

func (s *CreateHybridMonitorTaskRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *CreateHybridMonitorTaskRequest) GetTargetUserIdList() *string {
	return s.TargetUserIdList
}

func (s *CreateHybridMonitorTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateHybridMonitorTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateHybridMonitorTaskRequest) GetYARMConfig() *string {
	return s.YARMConfig
}

func (s *CreateHybridMonitorTaskRequest) SetAttachLabels(v []*CreateHybridMonitorTaskRequestAttachLabels) *CreateHybridMonitorTaskRequest {
	s.AttachLabels = v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetCloudAccessId(v []*string) *CreateHybridMonitorTaskRequest {
	s.CloudAccessId = v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetCollectInterval(v string) *CreateHybridMonitorTaskRequest {
	s.CollectInterval = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetCollectTargetType(v string) *CreateHybridMonitorTaskRequest {
	s.CollectTargetType = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetDescription(v string) *CreateHybridMonitorTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetGroupId(v string) *CreateHybridMonitorTaskRequest {
	s.GroupId = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetNamespace(v string) *CreateHybridMonitorTaskRequest {
	s.Namespace = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetRegionId(v string) *CreateHybridMonitorTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetSLSProcessConfig(v *CreateHybridMonitorTaskRequestSLSProcessConfig) *CreateHybridMonitorTaskRequest {
	s.SLSProcessConfig = v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetTargetUserId(v string) *CreateHybridMonitorTaskRequest {
	s.TargetUserId = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetTargetUserIdList(v string) *CreateHybridMonitorTaskRequest {
	s.TargetUserIdList = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetTaskName(v string) *CreateHybridMonitorTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetTaskType(v string) *CreateHybridMonitorTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) SetYARMConfig(v string) *CreateHybridMonitorTaskRequest {
	s.YARMConfig = &v
	return s
}

func (s *CreateHybridMonitorTaskRequest) Validate() error {
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

type CreateHybridMonitorTaskRequestAttachLabels struct {
	// The label key of the metric.
	//
	// example:
	//
	// app_service
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The label value of the metric.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHybridMonitorTaskRequestAttachLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestAttachLabels) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestAttachLabels) GetName() *string {
	return s.Name
}

func (s *CreateHybridMonitorTaskRequestAttachLabels) GetValue() *string {
	return s.Value
}

func (s *CreateHybridMonitorTaskRequestAttachLabels) SetName(v string) *CreateHybridMonitorTaskRequestAttachLabels {
	s.Name = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestAttachLabels) SetValue(v string) *CreateHybridMonitorTaskRequestAttachLabels {
	s.Value = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestAttachLabels) Validate() error {
	return dara.Validate(s)
}

type CreateHybridMonitorTaskRequestSLSProcessConfig struct {
	// The arithmetic operation results of the extended fields for SLS log statistical results.
	Express []*CreateHybridMonitorTaskRequestSLSProcessConfigExpress `json:"Express,omitempty" xml:"Express,omitempty" type:"Repeated"`
	// The filter conditions for parameters in SLS logs.
	Filter *CreateHybridMonitorTaskRequestSLSProcessConfigFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// Aggregates data by spatial dimensions, equivalent to Group By in SQL.
	GroupBy []*CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy `json:"GroupBy,omitempty" xml:"GroupBy,omitempty" type:"Repeated"`
	// The statistical method for SLS logs.
	Statistics []*CreateHybridMonitorTaskRequestSLSProcessConfigStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfig) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) GetExpress() []*CreateHybridMonitorTaskRequestSLSProcessConfigExpress {
	return s.Express
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) GetFilter() *CreateHybridMonitorTaskRequestSLSProcessConfigFilter {
	return s.Filter
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) GetGroupBy() []*CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy {
	return s.GroupBy
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) GetStatistics() []*CreateHybridMonitorTaskRequestSLSProcessConfigStatistics {
	return s.Statistics
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) SetExpress(v []*CreateHybridMonitorTaskRequestSLSProcessConfigExpress) *CreateHybridMonitorTaskRequestSLSProcessConfig {
	s.Express = v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) SetFilter(v *CreateHybridMonitorTaskRequestSLSProcessConfigFilter) *CreateHybridMonitorTaskRequestSLSProcessConfig {
	s.Filter = v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) SetGroupBy(v []*CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) *CreateHybridMonitorTaskRequestSLSProcessConfig {
	s.GroupBy = v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) SetStatistics(v []*CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) *CreateHybridMonitorTaskRequestSLSProcessConfig {
	s.Statistics = v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfig) Validate() error {
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

type CreateHybridMonitorTaskRequestSLSProcessConfigExpress struct {
	// The alias of the arithmetic operation result of the extended field for SLS log statistical results.
	//
	// example:
	//
	// SuccRate
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The arithmetic operation results of the extended fields for SLS log statistical results.
	//
	// example:
	//
	// success_count
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfigExpress) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfigExpress) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigExpress) GetAlias() *string {
	return s.Alias
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigExpress) GetExpress() *string {
	return s.Express
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigExpress) SetAlias(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigExpress {
	s.Alias = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigExpress) SetExpress(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigExpress {
	s.Express = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigExpress) Validate() error {
	return dara.Validate(s)
}

type CreateHybridMonitorTaskRequestSLSProcessConfigFilter struct {
	// The list of filter conditions for parameters in SLS logs.
	Filters []*CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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

func (s CreateHybridMonitorTaskRequestSLSProcessConfigFilter) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfigFilter) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilter) GetFilters() []*CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	return s.Filters
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilter) GetRelation() *string {
	return s.Relation
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilter) SetFilters(v []*CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) *CreateHybridMonitorTaskRequestSLSProcessConfigFilter {
	s.Filters = v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilter) SetRelation(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigFilter {
	s.Relation = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilter) Validate() error {
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

type CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters struct {
	// The filter method for parameter values in SLS logs. Valid values:
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
	// The name of the parameter to filter in SLS logs.
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

func (s CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GetOperator() *string {
	return s.Operator
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) GetValue() *string {
	return s.Value
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) SetOperator(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	s.Operator = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) SetSLSKeyName(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	s.SLSKeyName = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) SetValue(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters {
	s.Value = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigFilterFilters) Validate() error {
	return dara.Validate(s)
}

type CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy struct {
	// The alias of the SLS log statistical result.
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

func (s CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) GetAlias() *string {
	return s.Alias
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) SetAlias(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy {
	s.Alias = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) SetSLSKeyName(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy {
	s.SLSKeyName = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigGroupBy) Validate() error {
	return dara.Validate(s)
}

type CreateHybridMonitorTaskRequestSLSProcessConfigStatistics struct {
	// The alias of the SLS log statistical result.
	//
	// example:
	//
	// level_count
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The statistical method used to aggregate log data within the statistical period. Valid values:
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
	// - countps: calculates the average count per second for the specified field within the statistical period.
	//
	// - sumps: calculates the average sum per second for the specified field within the statistical period.
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
	// The statistical value of SLS logs.
	//
	// - If `Function` is set to `distribution`, this parameter specifies the lower limit of the statistical range. For example, to count the number of 2XX HTTP status codes, set this parameter to 200.
	//
	// - If `Function` is set to `percentile`, this parameter specifies the percentile of the statistical distribution. For example, 0.5 indicates P50.
	//
	// example:
	//
	// 200
	Parameter1 *string `json:"Parameter1,omitempty" xml:"Parameter1,omitempty"`
	// The statistical value of SLS logs.
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

func (s CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) GetAlias() *string {
	return s.Alias
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) GetFunction() *string {
	return s.Function
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) GetParameter1() *string {
	return s.Parameter1
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) GetParameter2() *string {
	return s.Parameter2
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) GetSLSKeyName() *string {
	return s.SLSKeyName
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) SetAlias(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Alias = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) SetFunction(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Function = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) SetParameter1(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Parameter1 = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) SetParameter2(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.Parameter2 = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) SetSLSKeyName(v string) *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics {
	s.SLSKeyName = &v
	return s
}

func (s *CreateHybridMonitorTaskRequestSLSProcessConfigStatistics) Validate() error {
	return dara.Validate(s)
}
