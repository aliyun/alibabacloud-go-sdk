// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDataServiceResponseBody
	GetRequestId() *string
	SetResult(v *QueryDataServiceResponseBodyResult) *QueryDataServiceResponseBody
	GetResult() *QueryDataServiceResponseBodyResult
	SetSuccess(v bool) *QueryDataServiceResponseBody
	GetSuccess() *bool
}

type QueryDataServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns the result of the interface query.
	Result *QueryDataServiceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: The request was successful
	//
	// - false: The request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDataServiceResponseBody) GetResult() *QueryDataServiceResponseBodyResult {
	return s.Result
}

func (s *QueryDataServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDataServiceResponseBody) SetRequestId(v string) *QueryDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDataServiceResponseBody) SetResult(v *QueryDataServiceResponseBodyResult) *QueryDataServiceResponseBody {
	s.Result = v
	return s
}

func (s *QueryDataServiceResponseBody) SetSuccess(v bool) *QueryDataServiceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryDataServiceResponseBodyResult struct {
	// Column headers.
	Headers []*QueryDataServiceResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The SQL of the query request.
	//
	// example:
	//
	// SELECT COMPANY_T_1_.`area` AS D_AREA_2_, COMPANY_T_1_.`city` AS D_CITY_3_, SUM(COMPANY_T_1_.`profit_amt`) AS D_PROFIT_4_ FROM `quickbi_test`.`company_sales_record_copy` AS COMPANY_T_1_   LIMIT 0, 10
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The queried results returned.
	Values []map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryDataServiceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponseBodyResult) GetHeaders() []*QueryDataServiceResponseBodyResultHeaders {
	return s.Headers
}

func (s *QueryDataServiceResponseBodyResult) GetSql() *string {
	return s.Sql
}

func (s *QueryDataServiceResponseBodyResult) GetValues() []map[string]interface{} {
	return s.Values
}

func (s *QueryDataServiceResponseBodyResult) SetHeaders(v []*QueryDataServiceResponseBodyResultHeaders) *QueryDataServiceResponseBodyResult {
	s.Headers = v
	return s
}

func (s *QueryDataServiceResponseBodyResult) SetSql(v string) *QueryDataServiceResponseBodyResult {
	s.Sql = &v
	return s
}

func (s *QueryDataServiceResponseBodyResult) SetValues(v []map[string]interface{}) *QueryDataServiceResponseBodyResult {
	s.Values = v
	return s
}

func (s *QueryDataServiceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type QueryDataServiceResponseBodyResultHeaders struct {
	// Aggregation operator. Only present for measure fields, such as SUM, AVG, and MAX.
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// Field name, corresponding to the physical table field name.
	//
	// example:
	//
	// The alias of the field. The key of the map data row in the result parameter values.
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The data type of the field. Common types include number, string, date, datetime, time, and geographic.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The granularity of the dimension field.
	//
	// This field is returned only when the requested field is a date or geographic dimension, with the following possible values:
	//
	// - Date granularity: yearRegion (year), monthRegion (month), weekRegion (week), dayRegion (day), hourRegion (hour), minRegion (minute), secRegion (second)
	//
	// - Geographic granularity: COUNTRY (country level), PROVINCE (province level), CITY (city level), XIAN (district/county level), REGION (region)
	//
	// example:
	//
	// yearRegion
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// Alias for the field, serving as the key in the map data row of the values parameter.
	//
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Field type, used to distinguish between dimension and measure fields.
	//
	// example:
	//
	// StandardDimension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDataServiceResponseBodyResultHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceResponseBodyResultHeaders) GoString() string {
	return s.String()
}

func (s *QueryDataServiceResponseBodyResultHeaders) GetAggregator() *string {
	return s.Aggregator
}

func (s *QueryDataServiceResponseBodyResultHeaders) GetColumn() *string {
	return s.Column
}

func (s *QueryDataServiceResponseBodyResultHeaders) GetDataType() *string {
	return s.DataType
}

func (s *QueryDataServiceResponseBodyResultHeaders) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryDataServiceResponseBodyResultHeaders) GetLabel() *string {
	return s.Label
}

func (s *QueryDataServiceResponseBodyResultHeaders) GetType() *string {
	return s.Type
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetAggregator(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Aggregator = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetColumn(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Column = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetDataType(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.DataType = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetGranularity(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Granularity = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetLabel(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Label = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) SetType(v string) *QueryDataServiceResponseBodyResultHeaders {
	s.Type = &v
	return s
}

func (s *QueryDataServiceResponseBodyResultHeaders) Validate() error {
	return dara.Validate(s)
}
