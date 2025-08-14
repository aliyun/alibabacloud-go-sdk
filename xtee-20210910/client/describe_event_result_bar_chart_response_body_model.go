// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventResultBarChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEventResultBarChartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeEventResultBarChartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeEventResultBarChartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEventResultBarChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeEventResultBarChartResponseBodyResultObject) *DescribeEventResultBarChartResponseBody
	GetResultObject() *DescribeEventResultBarChartResponseBodyResultObject
	SetSuccess(v bool) *DescribeEventResultBarChartResponseBody
	GetSuccess() *bool
}

type DescribeEventResultBarChartResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeEventResultBarChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeEventResultBarChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEventResultBarChartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeEventResultBarChartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEventResultBarChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventResultBarChartResponseBody) GetResultObject() *DescribeEventResultBarChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeEventResultBarChartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEventResultBarChartResponseBody) SetCode(v string) *DescribeEventResultBarChartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBody) SetHttpStatusCode(v string) *DescribeEventResultBarChartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBody) SetMessage(v string) *DescribeEventResultBarChartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBody) SetRequestId(v string) *DescribeEventResultBarChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBody) SetResultObject(v *DescribeEventResultBarChartResponseBodyResultObject) *DescribeEventResultBarChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeEventResultBarChartResponseBody) SetSuccess(v bool) *DescribeEventResultBarChartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEventResultBarChartResponseBodyResultObject struct {
	// Chart data
	Series []*DescribeEventResultBarChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// xaxis interface configuration.
	Xaxis *DescribeEventResultBarChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeEventResultBarChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartResponseBodyResultObject) GetSeries() []*DescribeEventResultBarChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeEventResultBarChartResponseBodyResultObject) GetXaxis() *DescribeEventResultBarChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeEventResultBarChartResponseBodyResultObject) SetSeries(v []*DescribeEventResultBarChartResponseBodyResultObjectSeries) *DescribeEventResultBarChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObject) SetXaxis(v *DescribeEventResultBarChartResponseBodyResultObjectXaxis) *DescribeEventResultBarChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeEventResultBarChartResponseBodyResultObjectSeries struct {
	// Returned data object
	Data []*DescribeEventResultBarChartResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Series name.
	//
	// example:
	//
	// 通过
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Label key
	//
	// example:
	//
	// 通过
	Stack *string `json:"stack,omitempty" xml:"stack,omitempty"`
}

func (s DescribeEventResultBarChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) GetData() []*DescribeEventResultBarChartResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) GetStack() *string {
	return s.Stack
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) SetData(v []*DescribeEventResultBarChartResponseBodyResultObjectSeriesData) *DescribeEventResultBarChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) SetName(v string) *DescribeEventResultBarChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) SetStack(v string) *DescribeEventResultBarChartResponseBodyResultObjectSeries {
	s.Stack = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeEventResultBarChartResponseBodyResultObjectSeriesData struct {
	// Number.
	//
	// example:
	//
	// 100
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Scale
	//
	// example:
	//
	// 50.00%
	Scale *string `json:"scale,omitempty" xml:"scale,omitempty"`
}

func (s DescribeEventResultBarChartResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeriesData) GetNum() *int64 {
	return s.Num
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeriesData) GetScale() *string {
	return s.Scale
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeriesData) SetNum(v int64) *DescribeEventResultBarChartResponseBodyResultObjectSeriesData {
	s.Num = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeriesData) SetScale(v string) *DescribeEventResultBarChartResponseBodyResultObjectSeriesData {
	s.Scale = &v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeEventResultBarChartResponseBodyResultObjectXaxis struct {
	// Returned data object
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeEventResultBarChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeEventResultBarChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeEventResultBarChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
