// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkerDetectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeWorkerDetectionResponseBody
	GetDBClusterId() *string
	SetDetectionItems(v []*DescribeWorkerDetectionResponseBodyDetectionItems) *DescribeWorkerDetectionResponseBody
	GetDetectionItems() []*DescribeWorkerDetectionResponseBodyDetectionItems
	SetRequestId(v string) *DescribeWorkerDetectionResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeWorkerDetectionResponseBody
	GetTotalCount() *string
}

type DescribeWorkerDetectionResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-xxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The queried detection items and detection results.
	DetectionItems []*DescribeWorkerDetectionResponseBodyDetectionItems `json:"DetectionItems,omitempty" xml:"DetectionItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E5B37B61-E6C9-5FE0-9374-45BAA548AEF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWorkerDetectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeWorkerDetectionResponseBody) GetDetectionItems() []*DescribeWorkerDetectionResponseBodyDetectionItems {
	return s.DetectionItems
}

func (s *DescribeWorkerDetectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWorkerDetectionResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeWorkerDetectionResponseBody) SetDBClusterId(v string) *DescribeWorkerDetectionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBody) SetDetectionItems(v []*DescribeWorkerDetectionResponseBodyDetectionItems) *DescribeWorkerDetectionResponseBody {
	s.DetectionItems = v
	return s
}

func (s *DescribeWorkerDetectionResponseBody) SetRequestId(v string) *DescribeWorkerDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBody) SetTotalCount(v string) *DescribeWorkerDetectionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItems struct {
	// The information about the detection result.
	//
	// example:
	//
	// There are a total of 10 tables with an excessive number of primary keys.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the detection item.
	//
	// example:
	//
	// Metric detection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The detection result items.
	Results *DescribeWorkerDetectionResponseBodyDetectionItemsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
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

func (s DescribeWorkerDetectionResponseBodyDetectionItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItems) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) GetMessage() *string {
	return s.Message
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) GetName() *string {
	return s.Name
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) GetResults() *DescribeWorkerDetectionResponseBodyDetectionItemsResults {
	return s.Results
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) SetMessage(v string) *DescribeWorkerDetectionResponseBodyDetectionItems {
	s.Message = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) SetName(v string) *DescribeWorkerDetectionResponseBodyDetectionItems {
	s.Name = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) SetResults(v *DescribeWorkerDetectionResponseBodyDetectionItemsResults) *DescribeWorkerDetectionResponseBodyDetectionItems {
	s.Results = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) SetStatus(v string) *DescribeWorkerDetectionResponseBodyDetectionItems {
	s.Status = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItems) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResults struct {
	// The detection result items of operator metric aggregation.
	OperatorAgg []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg `json:"OperatorAgg,omitempty" xml:"OperatorAgg,omitempty" type:"Repeated"`
	// The detection result items of abnormal operators.
	OperatorDetails []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails `json:"OperatorDetails,omitempty" xml:"OperatorDetails,omitempty" type:"Repeated"`
	// The detection result items of improper partitioned tables.
	PartitionedTables []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables `json:"PartitionedTables,omitempty" xml:"PartitionedTables,omitempty" type:"Repeated"`
	// The detection result items of skewed tables.
	SkewedTables []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables `json:"SkewedTables,omitempty" xml:"SkewedTables,omitempty" type:"Repeated"`
	// The detection result items of table access.
	TopAccessTables []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables `json:"TopAccessTables,omitempty" xml:"TopAccessTables,omitempty" type:"Repeated"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResults) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) GetOperatorAgg() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg {
	return s.OperatorAgg
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) GetOperatorDetails() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails {
	return s.OperatorDetails
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) GetPartitionedTables() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	return s.PartitionedTables
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) GetSkewedTables() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	return s.SkewedTables
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) GetTopAccessTables() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables {
	return s.TopAccessTables
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) SetOperatorAgg(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) *DescribeWorkerDetectionResponseBodyDetectionItemsResults {
	s.OperatorAgg = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) SetOperatorDetails(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) *DescribeWorkerDetectionResponseBodyDetectionItemsResults {
	s.OperatorDetails = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) SetPartitionedTables(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) *DescribeWorkerDetectionResponseBodyDetectionItemsResults {
	s.PartitionedTables = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) SetSkewedTables(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) *DescribeWorkerDetectionResponseBodyDetectionItemsResults {
	s.SkewedTables = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) SetTopAccessTables(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) *DescribeWorkerDetectionResponseBodyDetectionItemsResults {
	s.TopAccessTables = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResults) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg struct {
	// The detection result items of operator metric aggregation.
	//
	// example:
	//
	// Peak memory
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The detection result items of operator metric aggregation.
	SearchResults []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults `json:"SearchResults,omitempty" xml:"SearchResults,omitempty" type:"Repeated"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) GetSearchResults() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	return s.SearchResults
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) SetMetricName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg {
	s.MetricName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) SetSearchResults(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg {
	s.SearchResults = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAgg) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults struct {
	// The average value of the operator metric.
	//
	// example:
	//
	// 2234
	AvgValue *float64 `json:"AvgValue,omitempty" xml:"AvgValue,omitempty"`
	// The maximum value of the operator metric.
	//
	// example:
	//
	// 444
	MaxValue *int64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The number of occurrences of the operator.
	//
	// example:
	//
	// 1234
	OperatorCount *int64 `json:"OperatorCount,omitempty" xml:"OperatorCount,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// Aggregation
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The cumulative value of the operator metric.
	//
	// example:
	//
	// 123
	TotalValue *int64 `json:"TotalValue,omitempty" xml:"TotalValue,omitempty"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetAvgValue() *float64 {
	return s.AvgValue
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetOperatorCount() *int64 {
	return s.OperatorCount
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) GetTotalValue() *int64 {
	return s.TotalValue
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetAvgValue(v float64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.AvgValue = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetMaxValue(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.MaxValue = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetOperatorCount(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.OperatorCount = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetOperatorName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.OperatorName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) SetTotalValue(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults {
	s.TotalValue = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorAggSearchResults) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails struct {
	// The name of the detection metric.
	//
	// example:
	//
	// PeakMemory
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The detection result items of abnormal operators.
	SearchResults []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults `json:"SearchResults,omitempty" xml:"SearchResults,omitempty" type:"Repeated"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) GetSearchResults() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	return s.SearchResults
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) SetMetricName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails {
	s.MetricName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) SetSearchResults(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails {
	s.SearchResults = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults struct {
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
	// 123
	InputSize *int64 `json:"InputSize,omitempty" xml:"InputSize,omitempty"`
	// The total CPU time consumed by all operators in the stage, which is equivalent to the total CPU time of the stage. You can use this parameter to determine which parts of the stage consume a large amount of computing resources. Unit: milliseconds.
	//
	// example:
	//
	// 23
	OperatorCost *int64 `json:"OperatorCost,omitempty" xml:"OperatorCost,omitempty"`
	// The property information about the operator.
	//
	// example:
	//
	// GROUP BY field: id
	OperatorInfo *string `json:"OperatorInfo,omitempty" xml:"OperatorInfo,omitempty"`
	// The name of the operator.
	//
	// example:
	//
	// TableScan
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The number of rows output by the operator.
	//
	// example:
	//
	// 123
	OutputRows *int64 `json:"OutputRows,omitempty" xml:"OutputRows,omitempty"`
	// The amount of data output by the operator. Unit: bytes.
	//
	// example:
	//
	// 123
	OutputSize *int64 `json:"OutputSize,omitempty" xml:"OutputSize,omitempty"`
	// The peak memory. Unit: bytes.
	//
	// example:
	//
	// 23
	PeakMemory *int64 `json:"PeakMemory,omitempty" xml:"PeakMemory,omitempty"`
	// The query ID that can be used for diagnostics.
	//
	// example:
	//
	// 2024041909301402103302422803151411141
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The stage ID.
	//
	// example:
	//
	// Stage[2]
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetInputRows() *int64 {
	return s.InputRows
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetInputSize() *int64 {
	return s.InputSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOperatorCost() *int64 {
	return s.OperatorCost
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOperatorInfo() *string {
	return s.OperatorInfo
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOperatorName() *string {
	return s.OperatorName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOutputRows() *int64 {
	return s.OutputRows
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetOutputSize() *int64 {
	return s.OutputSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetPeakMemory() *int64 {
	return s.PeakMemory
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetProcessId() *string {
	return s.ProcessId
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) GetStageId() *string {
	return s.StageId
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetInputRows(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.InputRows = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetInputSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.InputSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOperatorCost(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OperatorCost = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOperatorInfo(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OperatorInfo = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOperatorName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OperatorName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOutputRows(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OutputRows = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetOutputSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.OutputSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetPeakMemory(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.PeakMemory = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetProcessId(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.ProcessId = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) SetStageId(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults {
	s.StageId = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsOperatorDetailsSearchResults) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables struct {
	// The SQL statement that is used to create the table.
	//
	// example:
	//
	// create table test(id varchar)
	DDL *string `json:"DDL,omitempty" xml:"DDL,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 234
	PartitionCount *string `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The ID of the improper partition.
	//
	// example:
	//
	// [2024,2025]
	PartitionIds *string `json:"PartitionIds,omitempty" xml:"PartitionIds,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// nxg
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// zhw_place_order
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total data size of the table.
	//
	// example:
	//
	// 1234
	TotalDataSize *int64 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GetDDL() *string {
	return s.DDL
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GetPartitionCount() *string {
	return s.PartitionCount
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GetPartitionIds() *string {
	return s.PartitionIds
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) GetTotalDataSize() *int64 {
	return s.TotalDataSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) SetDDL(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	s.DDL = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) SetPartitionCount(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	s.PartitionCount = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) SetPartitionIds(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	s.PartitionIds = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) SetSchemaName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	s.SchemaName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) SetTableName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	s.TableName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) SetTotalDataSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables {
	s.TotalDataSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsPartitionedTables) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables struct {
	// The SQL statement that is used to create the table.
	//
	// example:
	//
	// create table test(id varchar)
	DDL *string `json:"DDL,omitempty" xml:"DDL,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 2
	PartitionCount *int32 `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// platfunc
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The number of skewed rows in the table.
	//
	// example:
	//
	// 1234
	ShardSkewedRows *string `json:"ShardSkewedRows,omitempty" xml:"ShardSkewedRows,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// sls_log_cheat_action
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total data size of the table. Unit: bytes.
	//
	// example:
	//
	// 2345
	TotalDataSize *int64 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty"`
	// The size of hot data. Unit: bytes.
	//
	// example:
	//
	// 2345
	TotalLocalDataSize *string `json:"TotalLocalDataSize,omitempty" xml:"TotalLocalDataSize,omitempty"`
	// The data size of the primary key. Unit: bytes.
	//
	// example:
	//
	// 234
	TotalPkSize *int64 `json:"TotalPkSize,omitempty" xml:"TotalPkSize,omitempty"`
	// The size of cold data. Unit: bytes.
	//
	// example:
	//
	// 234
	TotalRemoteDataSize *int64 `json:"TotalRemoteDataSize,omitempty" xml:"TotalRemoteDataSize,omitempty"`
	// The number of rows in the table.
	//
	// example:
	//
	// 34
	TotalRowCount *int64 `json:"TotalRowCount,omitempty" xml:"TotalRowCount,omitempty"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetDDL() *string {
	return s.DDL
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetShardSkewedRows() *string {
	return s.ShardSkewedRows
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetTotalDataSize() *int64 {
	return s.TotalDataSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetTotalLocalDataSize() *string {
	return s.TotalLocalDataSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetTotalPkSize() *int64 {
	return s.TotalPkSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetTotalRemoteDataSize() *int64 {
	return s.TotalRemoteDataSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) GetTotalRowCount() *int64 {
	return s.TotalRowCount
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetDDL(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.DDL = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetPartitionCount(v int32) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.PartitionCount = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetSchemaName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.SchemaName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetShardSkewedRows(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.ShardSkewedRows = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetTableName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.TableName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetTotalDataSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.TotalDataSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetTotalLocalDataSize(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.TotalLocalDataSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetTotalPkSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.TotalPkSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetTotalRemoteDataSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.TotalRemoteDataSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) SetTotalRowCount(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables {
	s.TotalRowCount = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsSkewedTables) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables struct {
	// The name of the detection metric.
	//
	// example:
	//
	// Peak memory detection
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The detection result items of table access.
	SearchResults []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults `json:"SearchResults,omitempty" xml:"SearchResults,omitempty" type:"Repeated"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) GetSearchResults() []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	return s.SearchResults
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) SetMetricName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables {
	s.MetricName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) SetSearchResults(v []*DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables {
	s.SearchResults = v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTables) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults struct {
	// The number of accesses to the table.
	//
	// example:
	//
	// 1111
	AccessCount *int64 `json:"AccessCount,omitempty" xml:"AccessCount,omitempty"`
	// The average amount of time for scanning. Unit: milliseconds.
	//
	// example:
	//
	// 234
	AvgScanCost *float64 `json:"AvgScanCost,omitempty" xml:"AvgScanCost,omitempty"`
	// The average data size for scanning. Unit: bytes.
	//
	// example:
	//
	// 234
	AvgScanSize *float64 `json:"AvgScanSize,omitempty" xml:"AvgScanSize,omitempty"`
	// The maximum amount of time for scanning. Unit: milliseconds.
	//
	// example:
	//
	// 345
	MaxScanCost *int64 `json:"MaxScanCost,omitempty" xml:"MaxScanCost,omitempty"`
	// The maximum data size for scanning. Unit: bytes.
	//
	// example:
	//
	// 2345
	MaxScanSize *int64 `json:"MaxScanSize,omitempty" xml:"MaxScanSize,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tiberias_2copt_origin_order_goods_info
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GoString() string {
	return s.String()
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GetAccessCount() *int64 {
	return s.AccessCount
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GetAvgScanCost() *float64 {
	return s.AvgScanCost
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GetAvgScanSize() *float64 {
	return s.AvgScanSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GetMaxScanCost() *int64 {
	return s.MaxScanCost
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GetMaxScanSize() *int64 {
	return s.MaxScanSize
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) GetTableName() *string {
	return s.TableName
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) SetAccessCount(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	s.AccessCount = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) SetAvgScanCost(v float64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	s.AvgScanCost = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) SetAvgScanSize(v float64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	s.AvgScanSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) SetMaxScanCost(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	s.MaxScanCost = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) SetMaxScanSize(v int64) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	s.MaxScanSize = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) SetTableName(v string) *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults {
	s.TableName = &v
	return s
}

func (s *DescribeWorkerDetectionResponseBodyDetectionItemsResultsTopAccessTablesSearchResults) Validate() error {
	return dara.Validate(s)
}
