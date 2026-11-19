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
	// The ID of the request.
	//
	// example:
	//
	// 78C1AA2D-9201-599E-A0BA-6FC462E57A95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query result returned by the API.
	Result *QueryDataServiceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDataServiceResponseBodyResult struct {
	// The column headers.
	Headers []*QueryDataServiceResponseBodyResultHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The SQL statement of the query request.
	//
	// example:
	//
	// SELECT COMPANY_T_1_.`area` AS D_AREA_2_, COMPANY_T_1_.`city` AS D_CITY_3_, SUM(COMPANY_T_1_.`profit_amt`) AS D_PROFIT_4_ FROM `quickbi_test`.`company_sales_record_copy` AS COMPANY_T_1_ WHERE COMPANY_T_1_.`area` LIKE \\"%East China%\\" GROUP BY COMPANY_T_1_.`area`, COMPANY_T_1_.`city` HAVING SUM(COMPANY_T_1_.`order_amt`) > 1 LIMIT 0, 10
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// The returned query results.
	//
	// example:
	//
	// [{"area":"East China","city":"Sanming"},{"area":"East China","city":"Shanghai"}]
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

type QueryDataServiceResponseBodyResultHeaders struct {
	// The aggregation operator. This parameter is returned only for measure fields, such as SUM, AVG, and MAX.
	//
	// example:
	//
	// SUM
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The name of the field, which corresponds to the field name in the physical table.
	//
	// example:
	//
	// A specific physical field name
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The data type of the field. Common values include number, string, date, datetime, time, and geographic.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The granularity of the dimension field. This parameter is returned only when the requested field is a date dimension or geographic dimension field. Valid values:
	//
	// - Date granularity: yearRegion (year), monthRegion (month), weekRegion (week), dayRegion (day), hourRegion (hour), minRegion (minute), and secRegion (second).
	//
	// - Geographic information granularity: COUNTRY (country level), PROVINCE (province level), CITY (city level), XIAN (district or county level), and REGION (region level).
	//
	// example:
	//
	// yearRegion
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The alias of the field, which is the key of the map data row in the Values response parameter.
	//
	// example:
	//
	// area
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The type of the field, which is used to distinguish whether the field is a dimension or a measure.
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
