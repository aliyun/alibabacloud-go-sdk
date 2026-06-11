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
	// The request ID.
	//
	// example:
	//
	// a4d1a221d-41za1-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the API call. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	Result *QueryDataResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Whether the request succeeded. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
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
	// The SQL statement for the query.
	//
	// > The returned SQL includes both the filter conditions from this call and any row-level or column-level permission rules.
	//
	// example:
	//
	// SELECT COMPANY_T_1_.`area` AS D_AREA_2_, COMPANY_T_1_.`city` AS D_CITY_3_, SUM(COMPANY_T_1_.`profit_amt`) AS D_PROFIT_4_ FROM `quickbi_test`.`company_sales_record_copy` AS COMPANY_T_1_ WHERE COMPANY_T_1_.`area` LIKE \\"%test%\\" GROUP BY COMPANY_T_1_.`area`, COMPANY_T_1_.`city` HAVING SUM(COMPANY_T_1_.`order_amt`) > 1 LIMIT 0, 10
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The query results.
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
	if s.Headers != nil {
		for _, item := range s.Headers {
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
	// The aggregate operator. Returned only for measure fields.
	//
	// - SUM: The sum.
	//
	// - MAX: The maximum value.
	//
	// - MIN: The minimum value.
	//
	// - AVG: The average value.
	//
	// - COUNT: The count.
	//
	// - COUNTD: The count of unique values.
	//
	// - STDDEV_POP: The population standard deviation.
	//
	// - STDDEV_SAMP: The sample standard deviation.
	//
	// - VAR_POP: The population variance.
	//
	// - VAR_SAMP: The sample variance.
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The physical table field name.
	//
	// example:
	//
	// test
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The field data type. Common types:
	//
	// - number
	//
	// - string
	//
	// - date
	//
	// - time
	//
	// - datetime
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The dimension granularity. Returned only for date or geographic dimensions. Valid values:
	//
	// - Date granularity: yearRegion (year), monthRegion (month), weekRegion (week), dayRegion (day), hourRegion (hour), minRegion (minute), secRegion (second)
	//
	// - Geographic granularity: COUNTRY (country), PROVINCE (province), CITY (city), XIAN (county/district), REGION (region)
	//
	// example:
	//
	// REGION
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The field alias. Used as the key in each Values map entry.
	//
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The original field name in the dataset.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// area
	OriginalColumn *string `json:"OriginalColumn,omitempty" xml:"OriginalColumn,omitempty"`
	// Whether the field is a dimension or measure.
	//
	// - Dimension
	//
	// - Measure
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
