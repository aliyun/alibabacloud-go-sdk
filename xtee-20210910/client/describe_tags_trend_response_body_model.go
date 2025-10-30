// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeTagsTrendResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeTagsTrendResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeTagsTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeTagsTrendResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeTagsTrendResponseBodyResultObject) *DescribeTagsTrendResponseBody
	GetResultObject() *DescribeTagsTrendResponseBodyResultObject
	SetSuccess(v bool) *DescribeTagsTrendResponseBody
	GetSuccess() *bool
}

type DescribeTagsTrendResponseBody struct {
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
	ResultObject *DescribeTagsTrendResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeTagsTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeTagsTrendResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeTagsTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTagsTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsTrendResponseBody) GetResultObject() *DescribeTagsTrendResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTagsTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagsTrendResponseBody) SetCode(v string) *DescribeTagsTrendResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeTagsTrendResponseBody) SetHttpStatusCode(v string) *DescribeTagsTrendResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTagsTrendResponseBody) SetMessage(v string) *DescribeTagsTrendResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTagsTrendResponseBody) SetRequestId(v string) *DescribeTagsTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsTrendResponseBody) SetResultObject(v *DescribeTagsTrendResponseBodyResultObject) *DescribeTagsTrendResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTagsTrendResponseBody) SetSuccess(v bool) *DescribeTagsTrendResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagsTrendResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsTrendResponseBodyResultObject struct {
	// Data list
	Series []*DescribeTagsTrendResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// xaxis returned data
	Xaxis *DescribeTagsTrendResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeTagsTrendResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendResponseBodyResultObject) GetSeries() []*DescribeTagsTrendResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeTagsTrendResponseBodyResultObject) GetXaxis() *DescribeTagsTrendResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeTagsTrendResponseBodyResultObject) SetSeries(v []*DescribeTagsTrendResponseBodyResultObjectSeries) *DescribeTagsTrendResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObject) SetXaxis(v *DescribeTagsTrendResponseBodyResultObjectXaxis) *DescribeTagsTrendResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObject) Validate() error {
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

type DescribeTagsTrendResponseBodyResultObjectSeries struct {
	// Chart data list
	Data []*DescribeTagsTrendResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Category name.
	//
	// example:
	//
	// rm0102
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeTagsTrendResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeries) GetData() []*DescribeTagsTrendResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeries) SetData(v []*DescribeTagsTrendResponseBodyResultObjectSeriesData) *DescribeTagsTrendResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeries) SetName(v string) *DescribeTagsTrendResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeries) Validate() error {
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

type DescribeTagsTrendResponseBodyResultObjectSeriesData struct {
	// The number of items in this category.
	//
	// example:
	//
	// 100
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Proportion
	//
	// example:
	//
	// 10%
	Scale *string `json:"scale,omitempty" xml:"scale,omitempty"`
}

func (s DescribeTagsTrendResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeriesData) GetNum() *int64 {
	return s.Num
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeriesData) GetScale() *string {
	return s.Scale
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeriesData) SetNum(v int64) *DescribeTagsTrendResponseBodyResultObjectSeriesData {
	s.Num = &v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeriesData) SetScale(v string) *DescribeTagsTrendResponseBodyResultObjectSeriesData {
	s.Scale = &v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsTrendResponseBodyResultObjectXaxis struct {
	// X-axis data
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeTagsTrendResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeTagsTrendResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeTagsTrendResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeTagsTrendResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
