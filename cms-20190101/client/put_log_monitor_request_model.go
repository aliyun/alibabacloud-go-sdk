// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutLogMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregates(v []*PutLogMonitorRequestAggregates) *PutLogMonitorRequest
	GetAggregates() []*PutLogMonitorRequestAggregates
	SetGroupId(v string) *PutLogMonitorRequest
	GetGroupId() *string
	SetGroupbys(v []*PutLogMonitorRequestGroupbys) *PutLogMonitorRequest
	GetGroupbys() []*PutLogMonitorRequestGroupbys
	SetLogId(v string) *PutLogMonitorRequest
	GetLogId() *string
	SetMetricExpress(v string) *PutLogMonitorRequest
	GetMetricExpress() *string
	SetMetricName(v string) *PutLogMonitorRequest
	GetMetricName() *string
	SetRegionId(v string) *PutLogMonitorRequest
	GetRegionId() *string
	SetSlsLogstore(v string) *PutLogMonitorRequest
	GetSlsLogstore() *string
	SetSlsProject(v string) *PutLogMonitorRequest
	GetSlsProject() *string
	SetSlsRegionId(v string) *PutLogMonitorRequest
	GetSlsRegionId() *string
	SetTumblingwindows(v string) *PutLogMonitorRequest
	GetTumblingwindows() *string
	SetUnit(v string) *PutLogMonitorRequest
	GetUnit() *string
	SetValueFilter(v []*PutLogMonitorRequestValueFilter) *PutLogMonitorRequest
	GetValueFilter() []*PutLogMonitorRequestValueFilter
	SetValueFilterRelation(v string) *PutLogMonitorRequest
	GetValueFilterRelation() *string
}

type PutLogMonitorRequest struct {
	// The aggregate function definitions.
	//
	// This parameter is required.
	Aggregates []*PutLogMonitorRequestAggregates `json:"Aggregates,omitempty" xml:"Aggregates,omitempty" type:"Repeated"`
	// The ID of the application group.
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The dimensions used for spatial aggregation. This is equivalent to the Group By clause in SQL, which groups monitoring data by specified dimensions. If no dimension is specified, all monitoring data is aggregated based on the aggregate function.
	Groupbys []*PutLogMonitorRequestGroupbys `json:"Groupbys,omitempty" xml:"Groupbys,omitempty" type:"Repeated"`
	// The ID of the log monitoring metric.
	//
	// example:
	//
	// 16****
	LogId *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The extended field. The extended field provides arithmetic operations on the results of the statistical methods.
	//
	// For example, if you configure the total number of HTTP status code requests (TotalNumber) and the number of requests with HTTP status codes greater than 499 (5xxNumber) in the statistical methods, you can use the extended field to calculate the server error rate: 5xxNumber/TotalNumber*100.
	//
	// JSON format: {"extend":{"errorPercent":"5xxNumber/TotalNumber*100"}}. Field description:
	//
	// - extend: required.
	//
	// - errorPercent: the alias of the new field generated from the calculation result. You can specify a custom name.
	//
	// - errorPercent: the calculation expression for existing fields.
	//
	// example:
	//
	// {"extend":{"errorPercent":"5XXNumber/TotalNumber*100"}}
	MetricExpress *string `json:"MetricExpress,omitempty" xml:"MetricExpress,omitempty"`
	// The metric name. For information about the metrics supported by CloudMonitor for Alibaba Cloud services, see [Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Log Service Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-logstore
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the Log Service project.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The region where the Log Service project resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// The tumbling window size for pre-aggregation. Unit: seconds. CloudMonitor performs an aggregation calculation on the data at the specified interval.
	//
	// example:
	//
	// 60,300
	Tumblingwindows *string `json:"Tumblingwindows,omitempty" xml:"Tumblingwindows,omitempty"`
	// The unit.
	//
	// example:
	//
	// Percent
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The filter rules, used together with ValueFilterRelation. This is equivalent to the Where clause in SQL. If this parameter is not specified, all data is processed. For example, if the log contains Level and Error fields and you want to count the number of Error occurrences per minute, you can define the statistical method to sum the Level field with the condition Level=Error.
	ValueFilter []*PutLogMonitorRequestValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
	// The logical operator used to combine log filter conditions. Valid values:
	//
	// - and
	//
	// - or
	//
	// > This parameter must be used together with `ValueFilter.N.Key`.
	//
	// This parameter is required.
	//
	// example:
	//
	// and
	ValueFilterRelation *string `json:"ValueFilterRelation,omitempty" xml:"ValueFilterRelation,omitempty"`
}

func (s PutLogMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s PutLogMonitorRequest) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequest) GetAggregates() []*PutLogMonitorRequestAggregates {
	return s.Aggregates
}

func (s *PutLogMonitorRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PutLogMonitorRequest) GetGroupbys() []*PutLogMonitorRequestGroupbys {
	return s.Groupbys
}

func (s *PutLogMonitorRequest) GetLogId() *string {
	return s.LogId
}

func (s *PutLogMonitorRequest) GetMetricExpress() *string {
	return s.MetricExpress
}

func (s *PutLogMonitorRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *PutLogMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutLogMonitorRequest) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *PutLogMonitorRequest) GetSlsProject() *string {
	return s.SlsProject
}

func (s *PutLogMonitorRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *PutLogMonitorRequest) GetTumblingwindows() *string {
	return s.Tumblingwindows
}

func (s *PutLogMonitorRequest) GetUnit() *string {
	return s.Unit
}

func (s *PutLogMonitorRequest) GetValueFilter() []*PutLogMonitorRequestValueFilter {
	return s.ValueFilter
}

func (s *PutLogMonitorRequest) GetValueFilterRelation() *string {
	return s.ValueFilterRelation
}

func (s *PutLogMonitorRequest) SetAggregates(v []*PutLogMonitorRequestAggregates) *PutLogMonitorRequest {
	s.Aggregates = v
	return s
}

func (s *PutLogMonitorRequest) SetGroupId(v string) *PutLogMonitorRequest {
	s.GroupId = &v
	return s
}

func (s *PutLogMonitorRequest) SetGroupbys(v []*PutLogMonitorRequestGroupbys) *PutLogMonitorRequest {
	s.Groupbys = v
	return s
}

func (s *PutLogMonitorRequest) SetLogId(v string) *PutLogMonitorRequest {
	s.LogId = &v
	return s
}

func (s *PutLogMonitorRequest) SetMetricExpress(v string) *PutLogMonitorRequest {
	s.MetricExpress = &v
	return s
}

func (s *PutLogMonitorRequest) SetMetricName(v string) *PutLogMonitorRequest {
	s.MetricName = &v
	return s
}

func (s *PutLogMonitorRequest) SetRegionId(v string) *PutLogMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *PutLogMonitorRequest) SetSlsLogstore(v string) *PutLogMonitorRequest {
	s.SlsLogstore = &v
	return s
}

func (s *PutLogMonitorRequest) SetSlsProject(v string) *PutLogMonitorRequest {
	s.SlsProject = &v
	return s
}

func (s *PutLogMonitorRequest) SetSlsRegionId(v string) *PutLogMonitorRequest {
	s.SlsRegionId = &v
	return s
}

func (s *PutLogMonitorRequest) SetTumblingwindows(v string) *PutLogMonitorRequest {
	s.Tumblingwindows = &v
	return s
}

func (s *PutLogMonitorRequest) SetUnit(v string) *PutLogMonitorRequest {
	s.Unit = &v
	return s
}

func (s *PutLogMonitorRequest) SetValueFilter(v []*PutLogMonitorRequestValueFilter) *PutLogMonitorRequest {
	s.ValueFilter = v
	return s
}

func (s *PutLogMonitorRequest) SetValueFilterRelation(v string) *PutLogMonitorRequest {
	s.ValueFilterRelation = &v
	return s
}

func (s *PutLogMonitorRequest) Validate() error {
	if s.Aggregates != nil {
		for _, item := range s.Aggregates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Groupbys != nil {
		for _, item := range s.Groupbys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ValueFilter != nil {
		for _, item := range s.ValueFilter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutLogMonitorRequestAggregates struct {
	// The alias of the aggregate function. Valid values of N: 1 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// Count
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the original field for aggregation. Valid values of N: 1 to 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// sourceCount
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The statistical method used to aggregate log data within a statistical period. Valid values of N: 1 to 10. Valid values:
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
	// - countps: calculates the average count per second for the specified field within the statistical period.
	//
	// - sumps: calculates the average sum per second for the specified field within the statistical period.
	//
	// - distinct: counts the number of occurrences of the specified field after deduplication within the statistical period.
	//
	// This parameter is required.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
}

func (s PutLogMonitorRequestAggregates) String() string {
	return dara.Prettify(s)
}

func (s PutLogMonitorRequestAggregates) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequestAggregates) GetAlias() *string {
	return s.Alias
}

func (s *PutLogMonitorRequestAggregates) GetFieldName() *string {
	return s.FieldName
}

func (s *PutLogMonitorRequestAggregates) GetFunction() *string {
	return s.Function
}

func (s *PutLogMonitorRequestAggregates) SetAlias(v string) *PutLogMonitorRequestAggregates {
	s.Alias = &v
	return s
}

func (s *PutLogMonitorRequestAggregates) SetFieldName(v string) *PutLogMonitorRequestAggregates {
	s.FieldName = &v
	return s
}

func (s *PutLogMonitorRequestAggregates) SetFunction(v string) *PutLogMonitorRequestAggregates {
	s.Function = &v
	return s
}

func (s *PutLogMonitorRequestAggregates) Validate() error {
	return dara.Validate(s)
}

type PutLogMonitorRequestGroupbys struct {
	// The alias of the Group By field. Valid values of N: 1 to 10.
	//
	// example:
	//
	// CPUUtilization
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the Group By field. Valid values of N: 1 to 10.
	//
	// example:
	//
	// cpu
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
}

func (s PutLogMonitorRequestGroupbys) String() string {
	return dara.Prettify(s)
}

func (s PutLogMonitorRequestGroupbys) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequestGroupbys) GetAlias() *string {
	return s.Alias
}

func (s *PutLogMonitorRequestGroupbys) GetFieldName() *string {
	return s.FieldName
}

func (s *PutLogMonitorRequestGroupbys) SetAlias(v string) *PutLogMonitorRequestGroupbys {
	s.Alias = &v
	return s
}

func (s *PutLogMonitorRequestGroupbys) SetFieldName(v string) *PutLogMonitorRequestGroupbys {
	s.FieldName = &v
	return s
}

func (s *PutLogMonitorRequestGroupbys) Validate() error {
	return dara.Validate(s)
}

type PutLogMonitorRequestValueFilter struct {
	// The name of the log field to match. Valid values of N: 1 to 10.
	//
	// example:
	//
	// lh_source
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method for the field value. Valid values of N: 1 to 10. Valid values:
	//
	// - `contain`: contains.
	//
	// - `notContain`: does not contain.
	//
	// - `>`: greater than.
	//
	// - `<`: less than.
	//
	// - `>=`: greater than or equal to.
	//
	// - `<=`: less than or equal to.
	//
	// example:
	//
	// contain
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the log field to match. Valid values of N: 1 to 10.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutLogMonitorRequestValueFilter) String() string {
	return dara.Prettify(s)
}

func (s PutLogMonitorRequestValueFilter) GoString() string {
	return s.String()
}

func (s *PutLogMonitorRequestValueFilter) GetKey() *string {
	return s.Key
}

func (s *PutLogMonitorRequestValueFilter) GetOperator() *string {
	return s.Operator
}

func (s *PutLogMonitorRequestValueFilter) GetValue() *string {
	return s.Value
}

func (s *PutLogMonitorRequestValueFilter) SetKey(v string) *PutLogMonitorRequestValueFilter {
	s.Key = &v
	return s
}

func (s *PutLogMonitorRequestValueFilter) SetOperator(v string) *PutLogMonitorRequestValueFilter {
	s.Operator = &v
	return s
}

func (s *PutLogMonitorRequestValueFilter) SetValue(v string) *PutLogMonitorRequestValueFilter {
	s.Value = &v
	return s
}

func (s *PutLogMonitorRequestValueFilter) Validate() error {
	return dara.Validate(s)
}
