// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskLineChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRiskLineChartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeRiskLineChartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeRiskLineChartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRiskLineChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRiskLineChartResponseBodyResultObject) *DescribeRiskLineChartResponseBody
	GetResultObject() *DescribeRiskLineChartResponseBodyResultObject
	SetSuccess(v bool) *DescribeRiskLineChartResponseBody
	GetSuccess() *bool
}

type DescribeRiskLineChartResponseBody struct {
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
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeRiskLineChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeRiskLineChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLineChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskLineChartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRiskLineChartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeRiskLineChartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRiskLineChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskLineChartResponseBody) GetResultObject() *DescribeRiskLineChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRiskLineChartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRiskLineChartResponseBody) SetCode(v string) *DescribeRiskLineChartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRiskLineChartResponseBody) SetHttpStatusCode(v string) *DescribeRiskLineChartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRiskLineChartResponseBody) SetMessage(v string) *DescribeRiskLineChartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRiskLineChartResponseBody) SetRequestId(v string) *DescribeRiskLineChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskLineChartResponseBody) SetResultObject(v *DescribeRiskLineChartResponseBodyResultObject) *DescribeRiskLineChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRiskLineChartResponseBody) SetSuccess(v bool) *DescribeRiskLineChartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRiskLineChartResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRiskLineChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeRiskLineChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// Details of xaxis node.
	Xaxis *DescribeRiskLineChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeRiskLineChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLineChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRiskLineChartResponseBodyResultObject) GetSeries() []*DescribeRiskLineChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeRiskLineChartResponseBodyResultObject) GetXaxis() *DescribeRiskLineChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeRiskLineChartResponseBodyResultObject) SetSeries(v []*DescribeRiskLineChartResponseBodyResultObjectSeries) *DescribeRiskLineChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeRiskLineChartResponseBodyResultObject) SetXaxis(v *DescribeRiskLineChartResponseBodyResultObjectXaxis) *DescribeRiskLineChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeRiskLineChartResponseBodyResultObject) Validate() error {
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

type DescribeRiskLineChartResponseBodyResultObjectSeries struct {
	// Line chart data
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Title of the line segment in the line chart
	//
	// example:
	//
	// rm0102
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeRiskLineChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLineChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeRiskLineChartResponseBodyResultObjectSeries) GetData() []*string {
	return s.Data
}

func (s *DescribeRiskLineChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeRiskLineChartResponseBodyResultObjectSeries) SetData(v []*string) *DescribeRiskLineChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeRiskLineChartResponseBodyResultObjectSeries) SetName(v string) *DescribeRiskLineChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeRiskLineChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskLineChartResponseBodyResultObjectXaxis struct {
	// Returns x-axis data points
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeRiskLineChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLineChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeRiskLineChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeRiskLineChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeRiskLineChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeRiskLineChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
