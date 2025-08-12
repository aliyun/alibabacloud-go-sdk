// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMonitorListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeLogMonitorListResponseBody
	GetCode() *string
	SetLogMonitorList(v []*DescribeLogMonitorListResponseBodyLogMonitorList) *DescribeLogMonitorListResponseBody
	GetLogMonitorList() []*DescribeLogMonitorListResponseBodyLogMonitorList
	SetMessage(v string) *DescribeLogMonitorListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeLogMonitorListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLogMonitorListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLogMonitorListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeLogMonitorListResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *DescribeLogMonitorListResponseBody
	GetTotal() *int64
}

type DescribeLogMonitorListResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The log monitoring metrics.
	LogMonitorList []*DescribeLogMonitorListResponseBodyLogMonitorList `json:"LogMonitorList,omitempty" xml:"LogMonitorList,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01E90080-4300-4FAA-B9AE-161956BC350D
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
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLogMonitorListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeLogMonitorListResponseBody) GetLogMonitorList() []*DescribeLogMonitorListResponseBodyLogMonitorList {
	return s.LogMonitorList
}

func (s *DescribeLogMonitorListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeLogMonitorListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLogMonitorListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLogMonitorListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogMonitorListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeLogMonitorListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLogMonitorListResponseBody) SetCode(v string) *DescribeLogMonitorListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetLogMonitorList(v []*DescribeLogMonitorListResponseBodyLogMonitorList) *DescribeLogMonitorListResponseBody {
	s.LogMonitorList = v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetMessage(v string) *DescribeLogMonitorListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetPageNumber(v int32) *DescribeLogMonitorListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetPageSize(v int32) *DescribeLogMonitorListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetRequestId(v string) *DescribeLogMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetSuccess(v bool) *DescribeLogMonitorListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) SetTotal(v int64) *DescribeLogMonitorListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLogMonitorListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLogMonitorListResponseBodyLogMonitorList struct {
	// The time when the log monitoring metric was created.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1577766395000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 12345
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the log monitoring metric.
	//
	// example:
	//
	// 12345
	LogId *int64 `json:"LogId,omitempty" xml:"LogId,omitempty"`
	// The metric name. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The name of the Simple Log Service (SLS) Logstore.
	//
	// example:
	//
	// testSlS****
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// sls-project-test****
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// The ID of the region where the SLS Logstore resides.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// The condition that is used to filter logs. The ValueFilter and ValueFilterRelation parameters are used in pair. The filter condition is equivalent to the WHERE clause in SQL statements. If no filter condition is specified, all logs are processed. For example, logs contain the Level and Error fields. If you need to calculate the number of times that logs of the Error level appear every minute, you can set the filter condition to Level=Error and count the number of logs that meet this condition.
	ValueFilter []*DescribeLogMonitorListResponseBodyLogMonitorListValueFilter `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty" type:"Repeated"`
	// The logical operator that is used between log filter conditions. The ValueFilter and ValueFilterRelation parameters are used in pair. Valid values:
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

func (s DescribeLogMonitorListResponseBodyLogMonitorList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorListResponseBodyLogMonitorList) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetLogId() *int64 {
	return s.LogId
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetSlsProject() *string {
	return s.SlsProject
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetValueFilter() []*DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	return s.ValueFilter
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) GetValueFilterRelation() *string {
	return s.ValueFilterRelation
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetGmtCreate(v int64) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetGroupId(v int64) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.GroupId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetLogId(v int64) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.LogId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetMetricName(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.MetricName = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetSlsLogstore(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetSlsProject(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.SlsProject = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetSlsRegionId(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetValueFilter(v []*DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.ValueFilter = v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) SetValueFilterRelation(v string) *DescribeLogMonitorListResponseBodyLogMonitorList {
	s.ValueFilterRelation = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorList) Validate() error {
	return dara.Validate(s)
}

type DescribeLogMonitorListResponseBodyLogMonitorListValueFilter struct {
	// The name of the log field used for matching in the filter condition.
	//
	// example:
	//
	// hostName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The method that is used to match the field value. Valid values:
	//
	// 	- contain: contains
	//
	// 	- notContain: does not contain
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

func (s DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) GetOperator() *string {
	return s.Operator
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) SetKey(v string) *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	s.Key = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) SetOperator(v string) *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	s.Operator = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) SetValue(v string) *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter {
	s.Value = &v
	return s
}

func (s *DescribeLogMonitorListResponseBodyLogMonitorListValueFilter) Validate() error {
	return dara.Validate(s)
}
