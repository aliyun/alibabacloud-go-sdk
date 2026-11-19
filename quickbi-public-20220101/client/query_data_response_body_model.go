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
	// The ID of the request.
	//
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the API. Valid values:
	//
	// - true: The execution is successful.
	//
	// - false: The execution fails.
	//
	// example:
	//
	// true
	Result *QueryDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDataResponseBodyResult struct {
	// The column headers.
	Headers []*QueryDataResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The SQL statement of the query.
	//
	// > In addition to the parameter conditions of the current API, the filter conditions in the returned SQL statement also include the row-level and column-level permission configurations.
	//
	// example:
	//
	// SELECT COMPANY_T_1_.`area` AS D_AREA_2_, COMPANY_T_1_.`city` AS D_CITY_3_, SUM(COMPANY_T_1_.`profit_amt`) AS D_PROFIT_4_ FROM `quickbi_test`.`company_sales_record_copy` AS COMPANY_T_1_ WHERE COMPANY_T_1_.`area` LIKE \\"%test%\\" GROUP BY COMPANY_T_1_.`area`, COMPANY_T_1_.`city` HAVING SUM(COMPANY_T_1_.`order_amt`) > 1 LIMIT 0, 10
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The returned query results.
	//
	// example:
	//
	// [{"area":"test","city":"test"},{"area":"test","city":"test"}]
	Values []map[string]interface{} `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	// The returned query results, including the original value fields.
	ValuesV2 []*QueryDataResponseBodyResultValuesV2 `json:"ValuesV2,omitempty" xml:"ValuesV2,omitempty" type:"Repeated"`
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

func (s *QueryDataResponseBodyResult) GetValuesV2() []*QueryDataResponseBodyResultValuesV2 {
	return s.ValuesV2
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

func (s *QueryDataResponseBodyResult) SetValuesV2(v []*QueryDataResponseBodyResultValuesV2) *QueryDataResponseBodyResult {
	s.ValuesV2 = v
	return s
}

func (s *QueryDataResponseBodyResult) Validate() error {
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ValuesV2 != nil {
		for _, item := range s.ValuesV2 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDataResponseBodyResultHeaders struct {
	// The aggregation operator. This field is available only for measure fields, such as SUM, AVG, and MAX. Valid values:
	//
	// - SUM: Sum.
	//
	// - MAX: Maximum value.
	//
	// - MIN: Minimum value.
	//
	// - AVG: Average value.
	//
	// - COUNT: Count.
	//
	// - COUNTD: Deduplicate count.
	//
	// - STDDEV_POP: Population standard deviation.
	//
	// - STDDEV_SAMP: Sample standard deviation.
	//
	// - VAR_POP: Population variance.
	//
	// - VAR_SAMP: Sample variance.
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The name of the field, which corresponds to the field name in the physical table.
	//
	// example:
	//
	// test
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The data type of the field. Valid values:
	//
	// - number: Numeric.
	//
	// - string: String.
	//
	// - date: Date.
	//
	// - time: Time.
	//
	// - datetime: Date and time.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The granularity of the dimension field. This field is returned only when the requested field is a date dimension or a geographic dimension. Valid values:
	//
	// - Date granularity: yearRegion (year), monthRegion (month), weekRegion (week), dayRegion (day), hourRegion (hour), minRegion (minute), and secRegion (second).
	//
	// - Geographic granularity: COUNTRY (country), PROVINCE (province), CITY (city), XIAN (district/county), and REGION (region).
	//
	// example:
	//
	// REGION
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The alias of the field, which is the key of the map data row in the values result parameter.
	//
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The original name of the corresponding field in the dataset.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// area
	OriginalColumn *string `json:"OriginalColumn,omitempty" xml:"OriginalColumn,omitempty"`
	// The type of the field, which is used to distinguish whether the field is a dimension or a measure. Valid values:
	//
	// - Dimension: Dimension.
	//
	// - Measure: Measure.
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

type QueryDataResponseBodyResultValuesV2 struct {
	// The row values.
	RowValue []*QueryDataResponseBodyResultValuesV2RowValue `json:"RowValue,omitempty" xml:"RowValue,omitempty" type:"Repeated"`
}

func (s QueryDataResponseBodyResultValuesV2) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponseBodyResultValuesV2) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBodyResultValuesV2) GetRowValue() []*QueryDataResponseBodyResultValuesV2RowValue {
	return s.RowValue
}

func (s *QueryDataResponseBodyResultValuesV2) SetRowValue(v []*QueryDataResponseBodyResultValuesV2RowValue) *QueryDataResponseBodyResultValuesV2 {
	s.RowValue = v
	return s
}

func (s *QueryDataResponseBodyResultValuesV2) Validate() error {
	if s.RowValue != nil {
		for _, item := range s.RowValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDataResponseBodyResultValuesV2RowValue struct {
	// The field alias set in Data Service.
	//
	// example:
	//
	// test
	L *string `json:"L,omitempty" xml:"L,omitempty"`
	// The original value.
	//
	// example:
	//
	// 0
	R *string `json:"R,omitempty" xml:"R,omitempty"`
	// The display value.
	//
	// example:
	//
	// 0
	V *string `json:"V,omitempty" xml:"V,omitempty"`
}

func (s QueryDataResponseBodyResultValuesV2RowValue) String() string {
	return dara.Prettify(s)
}

func (s QueryDataResponseBodyResultValuesV2RowValue) GoString() string {
	return s.String()
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) GetL() *string {
	return s.L
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) GetR() *string {
	return s.R
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) GetV() *string {
	return s.V
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) SetL(v string) *QueryDataResponseBodyResultValuesV2RowValue {
	s.L = &v
	return s
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) SetR(v string) *QueryDataResponseBodyResultValuesV2RowValue {
	s.R = &v
	return s
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) SetV(v string) *QueryDataResponseBodyResultValuesV2RowValue {
	s.V = &v
	return s
}

func (s *QueryDataResponseBodyResultValuesV2RowValue) Validate() error {
	return dara.Validate(s)
}
