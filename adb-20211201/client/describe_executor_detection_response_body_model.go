// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutorDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeExecutorDetectionResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeExecutorDetectionResponseBodyDetectionItems) *DescribeExecutorDetectionResponseBody
	GetDetectionItems() []*DescribeExecutorDetectionResponseBodyDetectionItems
	SetRequestId(v string) *DescribeExecutorDetectionResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeExecutorDetectionResponseBody
	GetTotalCount() *string
}

type DescribeExecutorDetectionResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// example:
	//
	// am-xxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeExecutorDetectionResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9DFF5F54-162B-5860-80A5-411FF550B347
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 566
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExecutorDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeExecutorDetectionResponseBody) GetDetectionItems() []*DescribeExecutorDetectionResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeExecutorDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExecutorDetectionResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeExecutorDetectionResponseBody) SetDBClusterId(v string) *DescribeExecutorDetectionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBody) SetDetectionItems(v []*DescribeExecutorDetectionResponseBodyDetectionItems) *DescribeExecutorDetectionResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeExecutorDetectionResponseBody) SetRequestId(v string) *DescribeExecutorDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBody) SetTotalCount(v string) *DescribeExecutorDetectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExecutorDetectionResponseBodyDetectionItems struct {
	// The information about the detection result.
	//
	// example:
	//
	// Large amounts of memory resources are used by the Aggregation operator.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Metric detection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detection result items.
	Results *DescribeExecutorDetectionResponseBodyDetectionItemsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// The severity level of the detection result. Valid values:
	//
	// 	- NORMAL
	//
	// 	- WARNING
	//
	// 	- CRITICAL
	//
	// example:
	//
	// WARNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExecutorDetectionResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) GetResults() *DescribeExecutorDetectionResponseBodyDetectionItemsResults {
	return s.Results
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) SetMessage(v string) *DescribeExecutorDetectionResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) SetName(v string) *DescribeExecutorDetectionResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) SetResults(v *DescribeExecutorDetectionResponseBodyDetectionItemsResults) *DescribeExecutorDetectionResponseBodyDetectionItems {
	s.Results = v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) SetStatus(v string) *DescribeExecutorDetectionResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeExecutorDetectionResponseBodyDetectionItemsResults struct {
	// The detection result items of operator metric aggregation.
	OperatorAgg []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg `json:"OperatorAgg,omitempty" xml:"OperatorAgg,omitempty" type:"Repeated"`
	// The detection result items of abnormal operators.
	OperatorDetails []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails `json:"OperatorDetails,omitempty" xml:"OperatorDetails,omitempty" type:"Repeated"`
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResults) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResults) GetOperatorAgg() []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg {
	return s.OperatorAgg
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResults) GetOperatorDetails() []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails {
	return s.OperatorDetails
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResults) SetOperatorAgg(v []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) *DescribeExecutorDetectionResponseBodyDetectionItemsResults {
	s.OperatorAgg = v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResults) SetOperatorDetails(v []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) *DescribeExecutorDetectionResponseBodyDetectionItemsResults {
	s.OperatorDetails = v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResults) Validate() error {
	return dara.Validate(s)
}

type DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg struct {
	// The name of the detection metric.
	//
	// example:
	//
	// OperatorCost
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The detection result items of operator metric aggregation.
	SearchResults []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults `json:"SearchResults,omitempty" xml:"SearchResults,omitempty" type:"Repeated"`
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) GetSearchResults() []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	return s.SearchResults
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) SetMetricName(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg {
	s.MetricName = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) SetSearchResults(v []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg {
	s.SearchResults = v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAgg) Validate() error {
	return dara.Validate(s)
}

type DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults struct {
	// The average value of the operator metric.
	//
	// example:
	//
	// 234
	AvgValue *float64 `json:"AvgValue,omitempty" xml:"AvgValue,omitempty"`
	// The maximum value of the operator metric.
	//
	// example:
	//
	// 2345
	MaxValue *int64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The number of occurrences of the operator.
	//
	// example:
	//
	// 3
	OperatorCount *int64 `json:"OperatorCount,omitempty" xml:"OperatorCount,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// Window
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The cumulative value of the operator metric.
	//
	// example:
	//
	// 345
	TotalValue *int64 `json:"TotalValue,omitempty" xml:"TotalValue,omitempty"`
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetAvgValue() *float64 {
	return s.AvgValue
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetOperatorCount() *int64 {
	return s.OperatorCount
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetTotalValue() *int64 {
	return s.TotalValue
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetAvgValue(v float64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.AvgValue = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetMaxValue(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.MaxValue = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetOperatorCount(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.OperatorCount = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetOperatorName(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.OperatorName = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetTotalValue(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.TotalValue = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) Validate() error {
	return dara.Validate(s)
}

type DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails struct {
	// The name of the detection metric.
	//
	// example:
	//
	// PeakMemory
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The detection result items of abnormal operators.
	SearchResults []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults `json:"SearchResults,omitempty" xml:"SearchResults,omitempty" type:"Repeated"`
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) GetSearchResults() []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	return s.SearchResults
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) SetMetricName(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails {
	s.MetricName = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) SetSearchResults(v []*DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails {
	s.SearchResults = v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults struct {
	// The number of rows input by the operator.
	//
	// example:
	//
	// 123
	InputRows *int64 `json:"InputRows,omitempty" xml:"InputRows,omitempty"`
	// The amount of data input by the operator. Unit: bytes.
	//
	// example:
	//
	// 345
	InputSize *int64 `json:"InputSize,omitempty" xml:"InputSize,omitempty"`
	// The total CPU time consumed by all operators in the stage, which is equivalent to the total CPU time of the stage. You can use this parameter to determine which parts of the stage consume a large amount of computing resources. Unit: milliseconds.
	//
	// example:
	//
	// 123
	OperatorCost *float64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The property information about the operator.
	//
	// example:
	//
	// GROUP BY field: uid
	OperatorInfo *string `json:"OperatorInfo,omitempty" xml:"OperatorInfo,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// Join
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The number of rows output by the operator.
	//
	// example:
	//
	// 2345
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The amount of data output by the operator. Unit: bytes.
	//
	// example:
	//
	// 234
	OutputSize *int64 `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	// The peak memory. Unit: bytes.
	//
	// example:
	//
	// 234
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID.
	//
	// example:
	//
	// 2024080110010002102500023803151627972
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The stage ID.
	//
	// example:
	//
	// Stage[3]
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GoString() string {
	return s.String()
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetInputRows() *int64 {
	return s.InputRows
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetInputSize() *int64 {
	return s.InputSize
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOperatorCost() *float64 {
	return s.OperatorCost
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOperatorInfo() *string {
	return s.OperatorInfo
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOutputSize() *int64 {
	return s.OutputSize
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetStageId() *string {
	return s.StageId
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetInputRows(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.InputRows = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetInputSize(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.InputSize = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOperatorCost(v float64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OperatorCost = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOperatorInfo(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OperatorInfo = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOperatorName(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OperatorName = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOutputRows(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OutputRows = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOutputSize(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OutputSize = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetPeakMemory(v int64) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.PeakMemory = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetProcessId(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.ProcessId = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetStageId(v string) *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.StageId = &v
	return s
}

func (s *DescribeExecutorDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) Validate() error {
	return dara.Validate(s)
}
