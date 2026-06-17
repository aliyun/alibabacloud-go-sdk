// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMonitorAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeLogMonitorAttributeResponseBody
	GetCode() *string
	SetLogMonitor(v *DescribeLogMonitorAttributeResponseBodyLogMonitor) *DescribeLogMonitorAttributeResponseBody
	GetLogMonitor() *DescribeLogMonitorAttributeResponseBodyLogMonitor
	SetMessage(v string) *DescribeLogMonitorAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeLogMonitorAttributeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLogMonitorAttributeResponseBody
	GetSuccess() *bool
}

type DescribeLogMonitorAttributeResponseBody struct {
	// The status code.
	//
	// > A status code of 200 indicates a successful request.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the Log Monitoring task.
	LogMonitor *DescribeLogMonitorAttributeResponseBodyLogMonitor `json:"LogMonitor,omitempty" xml:"LogMonitor,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6718537-E673-4A58-8EE1-24B8B38C7AAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeLogMonitorAttributeResponseBody) GetLogMonitor() *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	return s.LogMonitor
}

func (s *DescribeLogMonitorAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeLogMonitorAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogMonitorAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLogMonitorAttributeResponseBody) SetCode(v string) *DescribeLogMonitorAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetLogMonitor(v *DescribeLogMonitorAttributeResponseBodyLogMonitor) *DescribeLogMonitorAttributeResponseBody {
	s.LogMonitor = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetMessage(v string) *DescribeLogMonitorAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetRequestId(v string) *DescribeLogMonitorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) SetSuccess(v bool) *DescribeLogMonitorAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBody) Validate() error {
	if s.LogMonitor != nil {
		if err := s.LogMonitor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogMonitorAttributeResponseBodyLogMonitor struct {
	// The definitions of aggregations.
	Aggregates []*DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates `json:"Aggregates,omitempty" xml:"Aggregates,omitempty" type:"Repeated"`
	// The time when the task was created.
	//
	// This value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970.
	//
	// example:
	//
	// 1678440033000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 123******
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The dimension based on which log data is aggregated. This parameter is equivalent to the \\`GROUP BY\\` clause in an SQL statement. You can specify a dimension to group monitoring data. If you do not specify this parameter, all monitoring data is aggregated based on the aggregation method.
	Groupbys []*string `json:"Groupbys,omitempty" xml:"Groupbys,omitempty" type:"Repeated"`
	// The ID of the Log Monitoring task.
	//
	// example:
	//
	// 123******
	LogId *int64 `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The metric expression.
	//
	// example:
	//
	// {}
	MetricExpress *string `json:"MetricExpress,omitempty" xml:"MetricExpress,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// cpu_total_******
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// logstore_******
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// project_******
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The ID of the region where Simple Log Service resides.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// The pre-aggregation window. Unit: seconds. Cloud Monitor aggregates data in the specified pre-aggregation window.
	Tumblingwindows []*string `json:"Tumblingwindows,omitempty" xml:"Tumblingwindows,omitempty" type:"Repeated"`
	// The filter conditions. This parameter is used with \\`ValueFilterRelation\\`. This parameter is equivalent to the \\`WHERE\\` clause in an SQL statement.
	//
	// If you do not specify this parameter, all data is processed. For example, if a log contains a \\`Level\\` field and you want to count the number of logs where the value of \\`Level\\` is \\`Error\\`, you can set the aggregation function to \\`count\\` and specify a filter condition where \\`Level\\` equals \\`Error\\`.
	ValueFilter []*DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
	// The logical operator for the filter conditions. This parameter is used with \\`ValueFilter\\`. Valid values:
	//
	// - and: The logical AND operator.
	//
	// - or: The logical OR operator.
	//
	// example:
	//
	// and
	ValueFilterRelation *string `json:"ValueFilterRelation,omitempty" xml:"ValueFilterRelation,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitor) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitor) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetAggregates() []*DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	return s.Aggregates
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetGroupbys() []*string {
	return s.Groupbys
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetLogId() *int64 {
	return s.LogId
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetMetricExpress() *string {
	return s.MetricExpress
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetSlsProject() *string {
	return s.SlsProject
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetTumblingwindows() []*string {
	return s.Tumblingwindows
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetValueFilter() []*DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	return s.ValueFilter
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) GetValueFilterRelation() *string {
	return s.ValueFilterRelation
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetAggregates(v []*DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.Aggregates = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetGmtCreate(v int64) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.GmtCreate = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetGroupId(v int64) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.GroupId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetGroupbys(v []*string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.Groupbys = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetLogId(v int64) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.LogId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetMetricExpress(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.MetricExpress = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetMetricName(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.MetricName = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetSlsLogstore(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetSlsProject(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.SlsProject = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetSlsRegionId(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetTumblingwindows(v []*string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.Tumblingwindows = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetValueFilter(v []*DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.ValueFilter = v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) SetValueFilterRelation(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitor {
	s.ValueFilterRelation = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitor) Validate() error {
	if s.Aggregates != nil {
		for _, item := range s.Aggregates {
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

type DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates struct {
	// The alias of the field.
	//
	// example:
	//
	// alias_******
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The original name of the field in the log.
	//
	// example:
	//
	// field_******
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The function that is used to aggregate log data in a statistical period. Valid values:
	//
	// - count: Counts the number of logs.
	//
	// - sum: Calculates the sum of values in a field.
	//
	// - avg: Calculates the average of values in a field.
	//
	// - max: Selects the maximum value in a field.
	//
	// - min: Selects the minimum value in a field.
	//
	// - countps: Calculates the average number of logs that are generated per second in a statistical period.
	//
	// - sumps: Calculates the average sum of values in a field per second in a statistical period.
	//
	// - distinct: Counts the number of unique values in a field in a statistical period.
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The maximum value.
	//
	// example:
	//
	// 0
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum value.
	//
	// example:
	//
	// 0
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GetAlias() *string {
	return s.Alias
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GetFunction() *string {
	return s.Function
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GetMax() *string {
	return s.Max
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) GetMin() *string {
	return s.Min
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetAlias(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Alias = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetFieldName(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.FieldName = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetFunction(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Function = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetMax(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Max = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) SetMin(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates {
	s.Min = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates) Validate() error {
	return dara.Validate(s)
}

type DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter struct {
	// The key.
	//
	// example:
	//
	// key_******
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator that is used to match the field value. Valid values:
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
	// The value.
	//
	// example:
	//
	// value_******
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) GetOperator() *string {
	return s.Operator
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) SetKey(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	s.Key = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) SetOperator(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	s.Operator = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) SetValue(v string) *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter {
	s.Value = &v
	return s
}

func (s *DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter) Validate() error {
	return dara.Validate(s)
}
