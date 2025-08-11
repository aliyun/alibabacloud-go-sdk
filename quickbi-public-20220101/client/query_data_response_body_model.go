// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDataResponseBody
	GetRequestId() *string
	SetResult(v *QueryDataResponseBodyResult) *QueryDataResponseBody
	GetResult() *QueryDataResponseBodyResult
	SetSuccess(v bool) *QueryDataResponseBody
	GetSuccess() *bool
}

type QueryDataResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the interface execution. Possible values:
	//
	// - true: Execution succeeded
	//
	// - false: Execution failed
	Result *QueryDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDataResponseBody) GetResult() *QueryDataResponseBodyResult {
	return s.Result
}

func (s *QueryDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDataResponseBody) SetRequestId(v string) *QueryDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataResponseBody) SetResult(v *QueryDataResponseBodyResult) *QueryDataResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataResponseBody) SetSuccess(v bool) *QueryDataResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDataResponseBodyResult struct {
	// Column headers.
	Headers []*QueryDataResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The SQL query that was executed.
	//
	// > The filter conditions in the returned SQL statement include not only the parameters passed through this interface but also the row and column permission configurations.
	//
	// example:
	//
	// test
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The results of the query.
	Values []map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryDataResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBodyResult) GetHeaders() []*QueryDataResponseBodyResultHeaders {
	return s.Headers
}

func (s *QueryDataResponseBodyResult) GetSql() *string {
	return s.Sql
}

func (s *QueryDataResponseBodyResult) GetValues() []map[string]interface{} {
	return s.Values
}

func (s *QueryDataResponseBodyResult) SetHeaders(v []*QueryDataResponseBodyResultHeaders) *QueryDataResponseBodyResult {
	s.Headers = v
	return s
}

func (s *QueryDataResponseBodyResult) SetSql(v string) *QueryDataResponseBodyResult {
	s.Sql = &v
	return s
}

func (s *QueryDataResponseBodyResult) SetValues(v []map[string]interface{}) *QueryDataResponseBodyResult {
	s.Values = v
	return s
}

func (s *QueryDataResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryDataResponseBodyResultHeaders struct {
	// Aggregation operator. Only present for measure fields, such as SUM, AVG, and MAX.
	//
	// - SUM: Sum
	//
	// - MAX: Maximum value
	//
	// - MIN: Minimum value
	//
	// - AVG: Average
	//
	// - COUNT: Count
	//
	// - COUNTD: Distinct count
	//
	// - STDDEV_POP: Population standard deviation
	//
	// - STDDEV_SAMP: Sample standard deviation
	//
	// - VAR_POP: Population variance
	//
	// - VAR_SAMP: Sample variance
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// Field name, corresponding to the physical table field name.
	//
	// example:
	//
	// Specific physical field name
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The keyword of the sensitive field type.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The granularity of the dimension field.
	//
	// This field is returned only when the requested field is a date or geographic dimension, with the following possible values:
	//
	// - Date Granularity: yearRegion (year), monthRegion (month), weekRegion (week), dayRegion (day), hourRegion (hour), minRegion (minute), secRegion (second)
	//
	// - Geographic Granularity: COUNTRY (international level), PROVINCE (provincial level), CITY (city level), XIAN (district/county level), REGION (region)
	//
	// example:
	//
	// REGION
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// Field alias, which serves as the key in the map data rows of the `values` parameter.
	//
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// if can be null:
	// true
	OriginalColumn *string `json:"OriginalColumn,omitempty" xml:"OriginalColumn,omitempty"`
	// Field type, used to distinguish between dimension and measure fields.
	//
	// - Dimension: dimension
	//
	// - Measure: measure
	//
	// example:
	//
	// Dimension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataResponseBodyResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponseBodyResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBodyResultHeaders) GetAggregator() *string {
	return s.Aggregator
}

func (s *QueryDataResponseBodyResultHeaders) GetColumn() *string {
	return s.Column
}

func (s *QueryDataResponseBodyResultHeaders) GetDataType() *string {
	return s.DataType
}

func (s *QueryDataResponseBodyResultHeaders) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryDataResponseBodyResultHeaders) GetLabel() *string {
	return s.Label
}

func (s *QueryDataResponseBodyResultHeaders) GetOriginalColumn() *string {
	return s.OriginalColumn
}

func (s *QueryDataResponseBodyResultHeaders) GetType() *string {
	return s.Type
}

func (s *QueryDataResponseBodyResultHeaders) SetAggregator(v string) *QueryDataResponseBodyResultHeaders {
	s.Aggregator = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetColumn(v string) *QueryDataResponseBodyResultHeaders {
	s.Column = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetDataType(v string) *QueryDataResponseBodyResultHeaders {
	s.DataType = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetGranularity(v string) *QueryDataResponseBodyResultHeaders {
	s.Granularity = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetLabel(v string) *QueryDataResponseBodyResultHeaders {
	s.Label = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetOriginalColumn(v string) *QueryDataResponseBodyResultHeaders {
	s.OriginalColumn = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) SetType(v string) *QueryDataResponseBodyResultHeaders {
	s.Type = &v
	return s
}

func (s *QueryDataResponseBodyResultHeaders) Validate() error {
	return dara.Validate(s)
}
