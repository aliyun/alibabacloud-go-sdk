// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsBarChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagsBarChartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeTagsBarChartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeTagsBarChartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagsBarChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeTagsBarChartResponseBodyResultObject) *DescribeTagsBarChartResponseBody
	GetResultObject() *DescribeTagsBarChartResponseBodyResultObject
	SetSuccess(v bool) *DescribeTagsBarChartResponseBody
	GetSuccess() *bool
}

type DescribeTagsBarChartResponseBody struct {
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
	// Return object
	ResultObject *DescribeTagsBarChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeTagsBarChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagsBarChartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeTagsBarChartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagsBarChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsBarChartResponseBody) GetResultObject() *DescribeTagsBarChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTagsBarChartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagsBarChartResponseBody) SetCode(v string) *DescribeTagsBarChartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagsBarChartResponseBody) SetHttpStatusCode(v string) *DescribeTagsBarChartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTagsBarChartResponseBody) SetMessage(v string) *DescribeTagsBarChartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagsBarChartResponseBody) SetRequestId(v string) *DescribeTagsBarChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsBarChartResponseBody) SetResultObject(v *DescribeTagsBarChartResponseBodyResultObject) *DescribeTagsBarChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTagsBarChartResponseBody) SetSuccess(v bool) *DescribeTagsBarChartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagsBarChartResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsBarChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeTagsBarChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// xaxis interface configuration.
	Xaxis *DescribeTagsBarChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeTagsBarChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartResponseBodyResultObject) GetSeries() []*DescribeTagsBarChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeTagsBarChartResponseBodyResultObject) GetXaxis() *DescribeTagsBarChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeTagsBarChartResponseBodyResultObject) SetSeries(v []*DescribeTagsBarChartResponseBodyResultObjectSeries) *DescribeTagsBarChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObject) SetXaxis(v *DescribeTagsBarChartResponseBodyResultObjectXaxis) *DescribeTagsBarChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObject) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Xaxis != nil {
		if err := s.Xaxis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsBarChartResponseBodyResultObjectSeries struct {
	// Chart data list
	Data []*DescribeTagsBarChartResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Series name.
	//
	// example:
	//
	// tag
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Chart field, same as name
	//
	// example:
	//
	// tag
	Stack *string `json:"stack,omitempty" xml:"stack,omitempty"`
}

func (s DescribeTagsBarChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) GetData() []*DescribeTagsBarChartResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) GetStack() *string {
	return s.Stack
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) SetData(v []*DescribeTagsBarChartResponseBodyResultObjectSeriesData) *DescribeTagsBarChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) SetName(v string) *DescribeTagsBarChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) SetStack(v string) *DescribeTagsBarChartResponseBodyResultObjectSeries {
	s.Stack = &v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeries) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagsBarChartResponseBodyResultObjectSeriesData struct {
	// Number.
	//
	// example:
	//
	// 200
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Scale
	//
	// example:
	//
	// 10%
	Scale *string `json:"scale,omitempty" xml:"scale,omitempty"`
}

func (s DescribeTagsBarChartResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeriesData) GetNum() *int64 {
	return s.Num
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeriesData) GetScale() *string {
	return s.Scale
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeriesData) SetNum(v int64) *DescribeTagsBarChartResponseBodyResultObjectSeriesData {
	s.Num = &v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeriesData) SetScale(v string) *DescribeTagsBarChartResponseBodyResultObjectSeriesData {
	s.Scale = &v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsBarChartResponseBodyResultObjectXaxis struct {
	// xaxis data items
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeTagsBarChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeTagsBarChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeTagsBarChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeTagsBarChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
