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
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the log monitoring metric.
	LogMonitor *DescribeLogMonitorAttributeResponseBodyLogMonitor `json:"LogMonitor,omitempty" xml:"LogMonitor,omitempty" type:"Struct"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6718537-E673-4A58-8EE1-24B8B38C7AAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	return dara.Validate(s)
}

type DescribeLogMonitorAttributeResponseBodyLogMonitor struct {
	// The aggregation logic.
	Aggregates []*DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates `json:"Aggregates,omitempty" xml:"Aggregates,omitempty" type:"Repeated"`
	// The time when the metric was created.
	//
	// This value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1547431398000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId  *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Groupbys []*string `json:"Groupbys,omitempty" xml:"Groupbys,omitempty" type:"Repeated"`
	// The ID of the log.
	//
	// example:
	//
	// 1234
	LogId *int64 `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The extended field. The extended field allows you to perform basic operations on the aggregation results.
	//
	// For example, if you have calculated TotalNumber and 5XXNumber by aggregating the data. TotalNumber indicates the total number of HTTP requests, and 5XXNumber indicates the number of HTTP requests whose status code is greater than 499. You can calculate the server error rate by adding the following formula to the extended field: 5XXNumber/TotalNumber\\*100.
	//
	// example:
	//
	// {"extend":{"errorPercent":"5XXNumber/TotalNumber*100"}}
	MetricExpress *string `json:"MetricExpress,omitempty" xml:"MetricExpress,omitempty"`
	// The metric name. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// test-logstore
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// test-project
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The ID of the region where the Simple Log Service (SLS) Logstore resides.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId     *string   `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	Tumblingwindows []*string `json:"Tumblingwindows,omitempty" xml:"Tumblingwindows,omitempty" type:"Repeated"`
	// The condition that is used to filter logs. The ValueFilter and ValueFilterRelation parameters are used in pair. The filter condition is equivalent to the WHERE clause in SQL statements.
	//
	// If no filter condition is specified, all logs are processed. For example, logs contain the Level and Error fields. If you need to calculate the number of times that logs of the Error level appear every minute, you can set the filter condition to Level=Error and count the number of logs that meet this condition.
	ValueFilter []*DescribeLogMonitorAttributeResponseBodyLogMonitorValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
	// The logical operator that is used between log filter conditions. The ValueFilter and ValueFilterRelation parameters must be used in pair. Valid values:
	//
	// 	- and
	//
	// 	- or
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
	return dara.Validate(s)
}

type DescribeLogMonitorAttributeResponseBodyLogMonitorAggregates struct {
	// The alias of the field.
	//
	// example:
	//
	// HostName
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the field in logs.
	//
	// example:
	//
	// hostName
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The function that is used to aggregate the monitoring data of logs within a statistical period. Valid values:
	//
	// 	- count: counts the number
	//
	// 	- sum: calculates the total value
	//
	// 	- avg: calculates the average value
	//
	// 	- max: calculates the maximum value
	//
	// 	- min: calculates the minimum value
	//
	// 	- countps: calculates the number of values of the specified field divided by the total number of seconds within a statistical period
	//
	// 	- sumps: calculates the sum of the values of the specified field divided by the total number of seconds within a statistical period
	//
	// 	- distinct: calculates the number of unique values of the specified field within a statistical period
	//
	// example:
	//
	// count
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The maximum value.
	//
	// example:
	//
	// 10
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
	// The name of the log field used for matching in the filter condition.
	//
	// example:
	//
	// hostName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The method that is used to match the field value. Valid values:
	//
	// 	- `contain`: contains
	//
	// 	- `notContain`: does not contain
	//
	// 	- `>`: greater than
	//
	// 	- `<`: less than
	//
	// 	- `>=`: greater than or equal to
	//
	// 	- `<=`: less than or equal to
	//
	// example:
	//
	// contain
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The field value to be matched in the filter condition.
	//
	// example:
	//
	// portal
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
