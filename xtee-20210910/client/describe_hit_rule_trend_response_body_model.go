// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeHitRuleTrendResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeHitRuleTrendResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeHitRuleTrendResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeHitRuleTrendResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeHitRuleTrendResponseBodyResultObject) *DescribeHitRuleTrendResponseBody
	GetResultObject() *DescribeHitRuleTrendResponseBodyResultObject
	SetSuccess(v bool) *DescribeHitRuleTrendResponseBody
	GetSuccess() *bool
}

type DescribeHitRuleTrendResponseBody struct {
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
	// Response object
	ResultObject *DescribeHitRuleTrendResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Indicates whether the operation was successful, where true means success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeHitRuleTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeHitRuleTrendResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeHitRuleTrendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeHitRuleTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHitRuleTrendResponseBody) GetResultObject() *DescribeHitRuleTrendResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeHitRuleTrendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeHitRuleTrendResponseBody) SetCode(v string) *DescribeHitRuleTrendResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBody) SetHttpStatusCode(v string) *DescribeHitRuleTrendResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBody) SetMessage(v string) *DescribeHitRuleTrendResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBody) SetRequestId(v string) *DescribeHitRuleTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBody) SetResultObject(v *DescribeHitRuleTrendResponseBodyResultObject) *DescribeHitRuleTrendResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeHitRuleTrendResponseBody) SetSuccess(v bool) *DescribeHitRuleTrendResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHitRuleTrendResponseBodyResultObject struct {
	// Chart data
	Series []*DescribeHitRuleTrendResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// X-axis data
	Xaxis *DescribeHitRuleTrendResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeHitRuleTrendResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendResponseBodyResultObject) GetSeries() []*DescribeHitRuleTrendResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeHitRuleTrendResponseBodyResultObject) GetXaxis() *DescribeHitRuleTrendResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeHitRuleTrendResponseBodyResultObject) SetSeries(v []*DescribeHitRuleTrendResponseBodyResultObjectSeries) *DescribeHitRuleTrendResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObject) SetXaxis(v *DescribeHitRuleTrendResponseBodyResultObjectXaxis) *DescribeHitRuleTrendResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObject) Validate() error {
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

type DescribeHitRuleTrendResponseBodyResultObjectSeries struct {
	// Returned data object
	Data []*DescribeHitRuleTrendResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Display title
	//
	// example:
	//
	// 策略name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeHitRuleTrendResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeries) GetData() []*DescribeHitRuleTrendResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeries) SetData(v []*DescribeHitRuleTrendResponseBodyResultObjectSeriesData) *DescribeHitRuleTrendResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeries) SetName(v string) *DescribeHitRuleTrendResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeries) Validate() error {
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

type DescribeHitRuleTrendResponseBodyResultObjectSeriesData struct {
	// Number
	//
	// example:
	//
	// 50
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Scale
	//
	// example:
	//
	// 10.00%
	Scale *string `json:"scale,omitempty" xml:"scale,omitempty"`
}

func (s DescribeHitRuleTrendResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeriesData) GetNum() *int64 {
	return s.Num
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeriesData) GetScale() *string {
	return s.Scale
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeriesData) SetNum(v int64) *DescribeHitRuleTrendResponseBodyResultObjectSeriesData {
	s.Num = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeriesData) SetScale(v string) *DescribeHitRuleTrendResponseBodyResultObjectSeriesData {
	s.Scale = &v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}

type DescribeHitRuleTrendResponseBodyResultObjectXaxis struct {
	// Returned data object
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeHitRuleTrendResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeHitRuleTrendResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeHitRuleTrendResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
