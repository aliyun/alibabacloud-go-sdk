// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDecisionResultTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDecisionResultTrendResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeDecisionResultTrendResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeDecisionResultTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDecisionResultTrendResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeDecisionResultTrendResponseBodyResultObject) *DescribeDecisionResultTrendResponseBody
	GetResultObject() *DescribeDecisionResultTrendResponseBodyResultObject
	SetSuccess(v bool) *DescribeDecisionResultTrendResponseBody
	GetSuccess() *bool
}

type DescribeDecisionResultTrendResponseBody struct {
	// Status code
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
	// Error details
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeDecisionResultTrendResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeDecisionResultTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDecisionResultTrendResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeDecisionResultTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDecisionResultTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDecisionResultTrendResponseBody) GetResultObject() *DescribeDecisionResultTrendResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeDecisionResultTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDecisionResultTrendResponseBody) SetCode(v string) *DescribeDecisionResultTrendResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBody) SetHttpStatusCode(v string) *DescribeDecisionResultTrendResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBody) SetMessage(v string) *DescribeDecisionResultTrendResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBody) SetRequestId(v string) *DescribeDecisionResultTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBody) SetResultObject(v *DescribeDecisionResultTrendResponseBodyResultObject) *DescribeDecisionResultTrendResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeDecisionResultTrendResponseBody) SetSuccess(v bool) *DescribeDecisionResultTrendResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDecisionResultTrendResponseBodyResultObject struct {
	// Chart data
	Series []*DescribeDecisionResultTrendResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// X-axis data
	Xaxis *DescribeDecisionResultTrendResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeDecisionResultTrendResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendResponseBodyResultObject) GetSeries() []*DescribeDecisionResultTrendResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeDecisionResultTrendResponseBodyResultObject) GetXaxis() *DescribeDecisionResultTrendResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeDecisionResultTrendResponseBodyResultObject) SetSeries(v []*DescribeDecisionResultTrendResponseBodyResultObjectSeries) *DescribeDecisionResultTrendResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObject) SetXaxis(v *DescribeDecisionResultTrendResponseBodyResultObjectXaxis) *DescribeDecisionResultTrendResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeDecisionResultTrendResponseBodyResultObjectSeries struct {
	// Returned data object
	Data []*DescribeDecisionResultTrendResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Name.
	//
	// example:
	//
	// IpTag_FFF
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeDecisionResultTrendResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeries) GetData() []*DescribeDecisionResultTrendResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeries) SetData(v []*DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) *DescribeDecisionResultTrendResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeries) SetName(v string) *DescribeDecisionResultTrendResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeDecisionResultTrendResponseBodyResultObjectSeriesData struct {
	// Number
	//
	// example:
	//
	// 10
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// ratio
	//
	// example:
	//
	// 5.56%
	Scale *string `json:"scale,omitempty" xml:"scale,omitempty"`
}

func (s DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) GetNum() *int64 {
	return s.Num
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) GetScale() *string {
	return s.Scale
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) SetNum(v int64) *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData {
	s.Num = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) SetScale(v string) *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData {
	s.Scale = &v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeDecisionResultTrendResponseBodyResultObjectXaxis struct {
	// X-axis data structure.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeDecisionResultTrendResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeDecisionResultTrendResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeDecisionResultTrendResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
