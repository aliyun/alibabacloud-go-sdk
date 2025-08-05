// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEventCenterQueryEventsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBody(v *EventCenterQueryEventsRequestBody) *EventCenterQueryEventsRequest
  GetBody() *EventCenterQueryEventsRequestBody 
  SetBusName(v string) *EventCenterQueryEventsRequest
  GetBusName() *string 
  SetMaxResults(v int32) *EventCenterQueryEventsRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EventCenterQueryEventsRequest
  GetNextToken() *string 
}

type EventCenterQueryEventsRequest struct {
  // The request body.
  // 
  // This parameter is required.
  Body *EventCenterQueryEventsRequestBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
  // The name of the event bus.
  // 
  // example:
  // 
  // default
  BusName *string `json:"BusName,omitempty" xml:"BusName,omitempty"`
  // The number of entries per page. Valid values: 0 to 10000. Default value: 100.
  // 
  // example:
  // 
  // 100
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  // 用来标记当前开始读取的位置。置空表示从头开始。
  // 
  // example:
  // 
  // 100
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s EventCenterQueryEventsRequest) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequest) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequest) GetBody() *EventCenterQueryEventsRequestBody  {
  return s.Body
}

func (s *EventCenterQueryEventsRequest) GetBusName() *string  {
  return s.BusName
}

func (s *EventCenterQueryEventsRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EventCenterQueryEventsRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EventCenterQueryEventsRequest) SetBody(v *EventCenterQueryEventsRequestBody) *EventCenterQueryEventsRequest {
  s.Body = v
  return s
}

func (s *EventCenterQueryEventsRequest) SetBusName(v string) *EventCenterQueryEventsRequest {
  s.BusName = &v
  return s
}

func (s *EventCenterQueryEventsRequest) SetMaxResults(v int32) *EventCenterQueryEventsRequest {
  s.MaxResults = &v
  return s
}

func (s *EventCenterQueryEventsRequest) SetNextToken(v string) *EventCenterQueryEventsRequest {
  s.NextToken = &v
  return s
}

func (s *EventCenterQueryEventsRequest) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsRequestBody struct {
  // The query parameters.
  // 
  // This parameter is required.
  Parameters *EventCenterQueryEventsRequestBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
  // The query type. Valid values:
  // 
  // 	- **timeseries**: queries time series data.
  // 
  // 	- **table**: queries table data.
  // 
  // 	- **timeseries_and_table**: queries time series data and table data at the same time.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // timeseries_and_table
  QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
  // The schema ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ecs:GeneralSchema
  SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s EventCenterQueryEventsRequestBody) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequestBody) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequestBody) GetParameters() *EventCenterQueryEventsRequestBodyParameters  {
  return s.Parameters
}

func (s *EventCenterQueryEventsRequestBody) GetQueryType() *string  {
  return s.QueryType
}

func (s *EventCenterQueryEventsRequestBody) GetSchemaId() *string  {
  return s.SchemaId
}

func (s *EventCenterQueryEventsRequestBody) SetParameters(v *EventCenterQueryEventsRequestBodyParameters) *EventCenterQueryEventsRequestBody {
  s.Parameters = v
  return s
}

func (s *EventCenterQueryEventsRequestBody) SetQueryType(v string) *EventCenterQueryEventsRequestBody {
  s.QueryType = &v
  return s
}

func (s *EventCenterQueryEventsRequestBody) SetSchemaId(v string) *EventCenterQueryEventsRequestBody {
  s.SchemaId = &v
  return s
}

func (s *EventCenterQueryEventsRequestBody) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsRequestBodyParameters struct {
  // Specifies whether to further split the dataset based on the column name.
  Breakdowns []*string `json:"Breakdowns,omitempty" xml:"Breakdowns,omitempty" type:"Repeated"`
  // The operator that is used to calculate the specified column.
  Calculations []*EventCenterQueryEventsRequestBodyParametersCalculations `json:"Calculations,omitempty" xml:"Calculations,omitempty" type:"Repeated"`
  // The timestamp that specifies the end of the time range to query. Unit: milliseconds.
  // 
  // example:
  // 
  // 1687861201814
  EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
  // The logic used to filter the combination of conditions.
  // 
  // example:
  // 
  // AND
  FilterCombination *string `json:"FilterCombination,omitempty" xml:"FilterCombination,omitempty"`
  // The filter conditions.
  Filters []*EventCenterQueryEventsRequestBodyParametersFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
  // The minimum time unit for querying time series data. Minimum value: 1. Unit: seconds. The value of this parameter is a recommended value. The actual value returned shall prevail.
  // 
  // example:
  // 
  // 30
  Granularity *int `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
  // The maximum number of events to query. Valid values: 1 to 10000.
  // 
  // example:
  // 
  // 100
  Limit *int `json:"Limit,omitempty" xml:"Limit,omitempty"`
  // The offset of the start position for this query. The offset starts from 0.
  // 
  // example:
  // 
  // 0
  Offset *int `json:"Offset,omitempty" xml:"Offset,omitempty"`
  // The order of the query results. This parameter is valid only if you set QueryType to table.
  Orders []*EventCenterQueryEventsRequestBodyParametersOrders `json:"Orders,omitempty" xml:"Orders,omitempty" type:"Repeated"`
  // The timestamp that specifies the beginning of the time range to query. Unit: milliseconds.
  // 
  // example:
  // 
  // 1687860901814
  StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
  // The time range during which events are queried. Minimum value: 1000. Unit: milliseconds.
  // 
  // example:
  // 
  // 1000000
  TimeRange *int64 `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s EventCenterQueryEventsRequestBodyParameters) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequestBodyParameters) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetBreakdowns() []*string  {
  return s.Breakdowns
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetCalculations() []*EventCenterQueryEventsRequestBodyParametersCalculations  {
  return s.Calculations
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetEndTime() *int64  {
  return s.EndTime
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetFilterCombination() *string  {
  return s.FilterCombination
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetFilters() []*EventCenterQueryEventsRequestBodyParametersFilters  {
  return s.Filters
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetGranularity() *int  {
  return s.Granularity
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetLimit() *int  {
  return s.Limit
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetOffset() *int  {
  return s.Offset
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetOrders() []*EventCenterQueryEventsRequestBodyParametersOrders  {
  return s.Orders
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetStartTime() *int64  {
  return s.StartTime
}

func (s *EventCenterQueryEventsRequestBodyParameters) GetTimeRange() *int64  {
  return s.TimeRange
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetBreakdowns(v []*string) *EventCenterQueryEventsRequestBodyParameters {
  s.Breakdowns = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetCalculations(v []*EventCenterQueryEventsRequestBodyParametersCalculations) *EventCenterQueryEventsRequestBodyParameters {
  s.Calculations = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetEndTime(v int64) *EventCenterQueryEventsRequestBodyParameters {
  s.EndTime = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetFilterCombination(v string) *EventCenterQueryEventsRequestBodyParameters {
  s.FilterCombination = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetFilters(v []*EventCenterQueryEventsRequestBodyParametersFilters) *EventCenterQueryEventsRequestBodyParameters {
  s.Filters = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetGranularity(v int) *EventCenterQueryEventsRequestBodyParameters {
  s.Granularity = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetLimit(v int) *EventCenterQueryEventsRequestBodyParameters {
  s.Limit = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetOffset(v int) *EventCenterQueryEventsRequestBodyParameters {
  s.Offset = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetOrders(v []*EventCenterQueryEventsRequestBodyParametersOrders) *EventCenterQueryEventsRequestBodyParameters {
  s.Orders = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetStartTime(v int64) *EventCenterQueryEventsRequestBodyParameters {
  s.StartTime = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) SetTimeRange(v int64) *EventCenterQueryEventsRequestBodyParameters {
  s.TimeRange = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParameters) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsRequestBodyParametersCalculations struct {
  // The column name.
  // 
  // example:
  // 
  // source
  Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
  // The operator.
  // 
  // example:
  // 
  // COUNT_DISTINCT
  Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
}

func (s EventCenterQueryEventsRequestBodyParametersCalculations) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequestBodyParametersCalculations) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequestBodyParametersCalculations) GetColumn() *string  {
  return s.Column
}

func (s *EventCenterQueryEventsRequestBodyParametersCalculations) GetOp() *string  {
  return s.Op
}

func (s *EventCenterQueryEventsRequestBodyParametersCalculations) SetColumn(v string) *EventCenterQueryEventsRequestBodyParametersCalculations {
  s.Column = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersCalculations) SetOp(v string) *EventCenterQueryEventsRequestBodyParametersCalculations {
  s.Op = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersCalculations) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsRequestBodyParametersFilters struct {
  // The column name.
  // 
  // example:
  // 
  // source
  Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
  // example:
  // 
  // AND
  NestedFilterCombination *string `json:"NestedFilterCombination,omitempty" xml:"NestedFilterCombination,omitempty"`
  NestedFilters []*EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters `json:"NestedFilters,omitempty" xml:"NestedFilters,omitempty" type:"Repeated"`
  // The operator.
  // 
  // example:
  // 
  // =
  Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
  // The values that are used together with the operator.
  Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s EventCenterQueryEventsRequestBodyParametersFilters) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequestBodyParametersFilters) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) GetColumn() *string  {
  return s.Column
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) GetNestedFilterCombination() *string  {
  return s.NestedFilterCombination
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) GetNestedFilters() []*EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters  {
  return s.NestedFilters
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) GetOp() *string  {
  return s.Op
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) GetValues() []*string  {
  return s.Values
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) SetColumn(v string) *EventCenterQueryEventsRequestBodyParametersFilters {
  s.Column = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) SetNestedFilterCombination(v string) *EventCenterQueryEventsRequestBodyParametersFilters {
  s.NestedFilterCombination = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) SetNestedFilters(v []*EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) *EventCenterQueryEventsRequestBodyParametersFilters {
  s.NestedFilters = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) SetOp(v string) *EventCenterQueryEventsRequestBodyParametersFilters {
  s.Op = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) SetValues(v []*string) *EventCenterQueryEventsRequestBodyParametersFilters {
  s.Values = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFilters) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters struct {
  Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
  Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
  Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) GetColumn() *string  {
  return s.Column
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) GetOp() *string  {
  return s.Op
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) GetValues() []*string  {
  return s.Values
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) SetColumn(v string) *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters {
  s.Column = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) SetOp(v string) *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters {
  s.Op = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) SetValues(v []*string) *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters {
  s.Values = v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersFiltersNestedFilters) Validate() error {
  return dara.Validate(s)
}

type EventCenterQueryEventsRequestBodyParametersOrders struct {
  // The column name.
  // 
  // example:
  // 
  // source
  Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
  // Specifies whether to sort the query results in descending order.
  // 
  // example:
  // 
  // true
  Desc *bool `json:"Desc,omitempty" xml:"Desc,omitempty"`
  // The operator.
  // 
  // example:
  // 
  // AVG
  Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
}

func (s EventCenterQueryEventsRequestBodyParametersOrders) String() string {
  return dara.Prettify(s)
}

func (s EventCenterQueryEventsRequestBodyParametersOrders) GoString() string {
  return s.String()
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) GetColumn() *string  {
  return s.Column
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) GetDesc() *bool  {
  return s.Desc
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) GetOp() *string  {
  return s.Op
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) SetColumn(v string) *EventCenterQueryEventsRequestBodyParametersOrders {
  s.Column = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) SetDesc(v bool) *EventCenterQueryEventsRequestBodyParametersOrders {
  s.Desc = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) SetOp(v string) *EventCenterQueryEventsRequestBodyParametersOrders {
  s.Op = &v
  return s
}

func (s *EventCenterQueryEventsRequestBodyParametersOrders) Validate() error {
  return dara.Validate(s)
}

